package adapter

import(
	"time"
	"context"
	
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/driver"
)

type KVSAdapter interface {
	Open(host string, passwd string, selector int) error
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string ) (interface{}, error)
	Delete(ctx context.Context, key ...string) error
	MapSet(ctx context.Context, mapKey string, key string, value interface{}) error
	MapGet(ctx context.Context, mapKey string, key string) (interface{}, error)
}

func NewKVSAdapter() KVSAdapter {
	return &driver.RedisDriver{}
}