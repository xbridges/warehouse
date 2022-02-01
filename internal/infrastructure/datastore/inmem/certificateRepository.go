package inmem

import(
	"context"
	"time"
	"github.com/xbridges/warehouse/internal/domain/model"
	"github.com/xbridges/warehouse/internal/domain/model/value"
	"github.com/xbridges/warehouse/internal/domain/model/Error"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/entity"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter"
)

type certificateRepository struct {
	adapter adapter.KVSAdapter
}

func NewCertificateRepository(adapter adapter.KVSAdapter) *certificateRepository{
	return &certificateRepository{adapter: adapter}
}

func (r *certificateRepository) FindCertificateBy(ctx context.Context, cert string) (model.UserAuth, error){
	
	if r == nil {
		return nil, Error.NewNoRepositoryError()
	}
	if r.adapter == nil {
		return nil, Error.NewNoAdapterError("KVSAdapter")
	}
	userJson, err := r.adapter.Get(ctx, cert)
	if err != nil {
		return nil, Error.NewNotFoundError(Error.CertificateMismatch, "kvs: certificate", err)
	}
	
	user := entity.NewUserEntity()
	if err := user.Unmarshal([]byte(userJson.(string))); err != nil {
		return nil, err
	} else {
		return model.NewUserAuth(user.GetModel()), nil
	}
}

func (r *certificateRepository) CreateCertificate(ctx context.Context, user model.UserAuth, param value.AuthParameter) (model.Certificate, error){
	if r == nil {
		return nil, Error.NewNoRepositoryError()
	}
	if r.adapter == nil {
		return nil, Error.NewNoAdapterError("KVSAdapter")
	}
	if user == nil {
		return nil, Error.NewNoModelError("UserAuth")
	}
	token, err := user.GetOnetimeToken()
	if err != nil {
		return nil, err
	}
	json, err := user.MarshalJson(value.MarshalALL)
	if err != nil {
		return nil, err	
	}
	if err := r.adapter.Set(ctx, token, json, time.Duration(param.GetExpiration())*time.Second); err != nil {
		return nil, Error.NewSystemFailedError("KVS Set failed", err)
	}
	return model.NewCertificate(token, param.GetExpiration(), param.GetDomain()), nil
}