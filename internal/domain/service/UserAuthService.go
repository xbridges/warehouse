package service

import(
	"github.com/xbridges/warehouse/internal/domain/model"
	"github.com/xbridges/warehouse/internal/domain/model/Error"
	"github.com/xbridges/warehouse/internal/domain/repository"
)

type userAuthService struct {
	repository repository.UserAuthRepository
}

type UserAuthService interface{
	ConfirmUser(show_id string, passwd string) (model.UserAuth, error)
}

func NewUserAuthService(repo repository.UserAuthRepository) UserAuthService {
	return &userAuthService{repository: repo}
}

// TODO: AccountLock状態を検出する機構が必要＋ロックカウント更新処理が必要
func (s *userAuthService) ConfirmUser(show_id string, passwd string) (model.UserAuth, error) {
	if s == nil {
		return nil, Error.NewNoServiceError()
	}
	// ユーザを抽出
	u, err := s.repository.FindUserBy(show_id)
	if err != nil {
		return nil, err
	}
	// 期限切れを検出
	if u.IsExpired() {
		return nil, Error.NewUserExpiredError()
	}
	// パスワードハッシュを作成
	passphrase, err := u.GetPassphrase(passwd)
	if err != nil {
		return nil, err
	}
	// パスワードハッシュを抽出
	p, t, err := s.repository.FindAuthorizationBy(passphrase) 
	if err != nil {
		return nil, err
	}
	// 期限切れを検出
	if p == nil && t != nil {
		return nil, Error.NewPasswordExpiredError(*t)
	}
	return u, nil
}