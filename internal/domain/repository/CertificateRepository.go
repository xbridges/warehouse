package repository

import(
	"context"
	"github.com/xbridges/warehouse/internal/domain/model"
	"github.com/xbridges/warehouse/internal/domain/model/value"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/inmem"
)

type CertificateRepository interface {
	FindCertificateBy(ctx context.Context, cert string) (model.UserAuth, error)
	CreateCertificate(ctx context.Context, user model.UserAuth, param value.AuthParameter) (model.Certificate, error)
}

func NewCertificateRepository(adapter adapter.KVSAdapter) CertificateRepository {
	return inmem.NewCertificateRepository(adapter)
}
