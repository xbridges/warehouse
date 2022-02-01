package service

import(
	"context"
	"github.com/xbridges/warehouse/internal/domain/model"
	"github.com/xbridges/warehouse/internal/domain/model/value"
	"github.com/xbridges/warehouse/internal/domain/model/Error"
	"github.com/xbridges/warehouse/internal/domain/repository"
)

type certificateService struct {
	parameter value.AuthParameter
	repository repository.CertificateRepository
}

type CertificateService interface{
	IssueCertificate(ctx context.Context, authUser model.UserAuth) (model.Certificate, error)
}

func NewCertificateService(repo repository.CertificateRepository, authParam value.AuthParameter) CertificateService {
	return &certificateService{repository: repo, parameter: authParam}
}

func (s *certificateService) IssueCertificate(ctx context.Context, authUser model.UserAuth) (model.Certificate, error) {
	if s == nil {
		return nil, Error.NewNoServiceError()
	}
	// regist onetime token to redis db
	cert, err := s.repository.CreateCertificate(ctx, authUser, s.parameter)
	if err != nil {
		return nil, err
	}
	return cert, nil
}