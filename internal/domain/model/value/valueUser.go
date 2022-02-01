package value

import(
	"time"
	"encoding/json"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"

	"github.com/xbridges/warehouse/internal/domain/model/Error"
)

// user 値オブジェクト
type valueUser struct {
	ID				UserID			`json:"id"`
	DisplayID		UserDisplayID	`json:"display_id"`
	Name			UserName		`json:"name"`
	EmailAddress	UserEmail		`json:"email"`
	Permission		UserPermission	`json:"permission"`
	UUID			UserUUID		`json:"uuid,omitempty"`
	Expiration		*valueTime		`json:"expiration"`
	Recovery		*valueTime		`json:"recovery"`
	Status			UserStatus		`json:"status"`
}

type UserID			*int
type UserDisplayID	string
type UserName 		*string
type UserEmail		*string
type UserPermission int
type UserUUID		string
type UserExpiration *valueTime
type UserRecovery	*valueTime
type UserStatus		int

func (v UserDisplayID) String() string{
	return string(v)
}

func (v UserUUID) String() string{
	return string(v)
}

type User interface {
	MarshalJson(mode int) ([]byte, error)
	UnmarshalJson(data []byte) (error) 
	GetPassphrase(passwd string) (string, error)
	GetOnetimeToken() (string, error)
	IsExpired() bool
}

const(
	MarshalSafe int = iota
	MarshalALL
)

// userは値オブジェクトなので、内容は変更させない。
// 内容変更が必要な場合は、作り直してrepositoryで処理する。
func NewUser(show_id string, permission int, uuid string, status int, id *int, name *string, email *string, expr *time.Time, recv *time.Time) User {	

	return &valueUser{
		ID:				UserID(id),
		DisplayID:		UserDisplayID(show_id),
		Name:			UserName(name),
		EmailAddress:	UserEmail(email),
		Permission:		UserPermission(permission),
		UUID:			UserUUID(uuid),
		Expiration:		NewValueTime(expr),
		Recovery:		NewValueTime(recv),
		Status:			UserStatus(status),
	}
}

func (v *valueUser) MarshalJson(mode int) ([]byte, error) {
	if v != nil {
		// ここは別インスタンスでコピーを作る
		outp := &valueUser{}
    	copier.CopyWithOption(outp, v, copier.Option{
			IgnoreEmpty: false,
			DeepCopy:    true,
    	})
		if mode != MarshalALL {
			outp.UUID = UserUUID("") // uuidは外に出さない
		}
		data, err := json.Marshal(outp)
		if err != nil {
			return nil, Error.NewSystemFailedError("json.Marshal", err)
		}
		return data, nil
	}
	return nil, Error.NewNoModelError("valueUser")
}

func (v *valueUser) UnmarshalJson(data []byte) (error) {
	if v != nil {
		err := json.Unmarshal(data, v)
		if err != nil {
			return Error.NewSystemFailedError("json.Unmarshal", err)
		}
		return nil
	}
	return Error.NewNoModelError("valueUser")
}


func (v *valueUser) GetPassphrase(passwd string) (string, error) {
	if v != nil {
		dec, err := translate(passwd, "nil")
		if err != nil {
			return "", err
		}
		dec, err = translate(v.DisplayID.String(), dec)
		if err != nil {
			return "", err
		}
		dec, err = translate(v.UUID.String(), dec)
		if err != nil {
			return "", err
		}
		return dec, nil
	}
	return "", Error.NewNoModelError("valueUser")
}

func (v *valueUser) GetOnetimeToken() (string, error) {
	if v != nil {
		dec, err := translate(v.DisplayID.String(), time.Now().Format("2006-01-02T15:04:05.000Z07:00"))
		if err != nil {
			return "", err
		}
		uid, err := newUUID()
		if err != nil {
			return "", err	
		} 
		dec, err = translate(dec, *uid)
		if err != nil {
			return "", err
		}
		return dec, nil
	}
	return "", Error.NewNoModelError("valueUser")
}

func (v *valueUser) IsExpired() bool {
	if v != nil {
		if v.Expiration == nil {
			return false	// 期限無制限
		}
		if time.Now().After(v.Expiration.Time) {
			return true	// 期限が現時点を過ぎている場合
		}
		return false // 期限前
	}
	return true // そもそもユーザオブジェクトがなければ期限もくそもないけど、期限切れと判断
}

func translate(src string, key string) (string, error) {
	keyhash := sha256.Sum256([]byte(key))
	mac := hmac.New(sha256.New, keyhash[:])
	_, err := mac.Write([]byte(src))
	if err != nil {
		return "", Error.NewSystemFailedError("mac.Write", err)
	}
	keyedHash := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(keyedHash), nil
}

func newUUID() (*string, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, Error.NewSystemFailedError("uuid.NewRandom", err)
	}
	uidStr := uid.String()
	return &uidStr, nil
}