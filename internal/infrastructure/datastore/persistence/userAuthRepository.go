package persistence

import(
	"time"
	"database/sql"
	"github.com/xbridges/warehouse/internal/domain/model"
	"github.com/xbridges/warehouse/internal/domain/model/Error"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/entity"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter"
)

type userAuthRepository struct {
	adapter adapter.RDBAdapter
}

func NewUserAuthRepository(adapter adapter.RDBAdapter) *userAuthRepository{
	return &userAuthRepository{adapter: adapter}
}

func (r *userAuthRepository) FindUserBy(showid string) (model.UserAuth, error){
	
	if r != nil {
		if r.adapter != nil {
			query := "select id, display_id, name, permission, uuid, email_address, expire_at, recovery_at, status from users where display_id = $1 and status = $2"
			row := r.adapter.QueryRow(query, showid, entity.UserEnabled) // 1 is user enabled
			if row != nil {
				user := entity.NewUserEntity()
				if err := row.Scan(&user.ID, &user.DisplayID, &user.Name, &user.Permission, &user.UUID, &user.EmailAddress, &user.ExpireAt, &user.RecoveryAt, &user.Status); err != nil {
					return nil, Error.NewNotFoundError(Error.UserMismatch, "users", err)
				}
				return model.NewUserAuth(user.GetModel()), nil
			}
		}
		return nil, Error.NewNoAdapterError("RDBAdapter")
	}
    return nil, Error.NewNoRepositoryError()
}

func (r *userAuthRepository) FindAuthorizationBy(passphraes string) (*string, *time.Time, error){
	
	if r != nil {
		if r.adapter != nil {
			query := "select hash_key, expire_at from authorizations where hash_key = $1 and status = 1"
			row := r.adapter.QueryRow(query, passphraes)
			if row != nil {
				var hash_key string
				var expire_at sql.NullTime
				if err := row.Scan(&hash_key, &expire_at); err != nil {
					return nil, nil, Error.NewNotFoundError(Error.PasswordMismatch, "authorizations", err)
				}
				if expire_at.Valid {
					if !expire_at.Time.Before(time.Now()) {
						return &hash_key, &expire_at.Time, nil
					} else {
						return nil, &expire_at.Time, Error.NewPasswordExpiredError(expire_at.Time)
					}
				} else {
					return &hash_key, nil, nil // timeがnullなら有効期限なしのため、正常応答
				}
			}
		}
		return nil, nil, Error.NewNoAdapterError("RDBAdapter")
	}
	return nil, nil, Error.NewNoRepositoryError()
}