package usecase

import(
	"context"

	"github.com/xbridges/warehouse/internal/domain/model"
	"github.com/xbridges/warehouse/internal/domain/service"
	"github.com/xbridges/warehouse/internal/domain/model/Error"
)

type loginUsecase struct {
	userAuthService service.UserAuthService
	certificateService service.CertificateService

}

type LoginUsecase interface{
	Login(ctx context.Context, id string, password string) (model.Certificate, model.UserAuth, error)
}

func NewLoginUsecase(usrserv service.UserAuthService, certserv service.CertificateService) LoginUsecase{
	return &loginUsecase{userAuthService: usrserv, certificateService: certserv}
}

func(u *loginUsecase) Login(ctx context.Context, id string, password string) (model.Certificate, model.UserAuth, error) {
	if u == nil {
		return nil, nil, Error.NewNoUsecaseError()
	}
	user, err := u.userAuthService.ConfirmUser(id, password)
	if err != nil {
		return nil, nil, err
	}
	cert, err := u.certificateService.IssueCertificate(ctx, user)
	if err != nil {
		return nil, nil, err
	}
	return cert, user, nil
}