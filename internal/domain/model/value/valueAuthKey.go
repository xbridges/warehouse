package value

import(
	"time"
)

// auth 値オブジェクト
type valueAuthKey struct {
	Hash			KeyHash			`json:"hash"`
	Expiration		KeyExpiration	`json:"expiration"`
	Status			KeyStatus		`json:"status"`
}

type KeyHash		string
type KeyExpiration *time.Time
type KeyStatus		int

func NewAuthKey(hash string, status int, expir *time.Time) *valueAuthKey{
	return &valueAuthKey{
		Hash: KeyHash(hash),
		Expiration: KeyExpiration(expir),
		Status: KeyStatus(status),
	}
}
