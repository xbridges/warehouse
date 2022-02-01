package driver

import(
	"time"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

type RedisDriver struct {
	*redis.Client
}


func(driver *RedisDriver) Open( host string, passwd string, selector int ) error {

	if driver != nil {
		client := redis.NewClient(&redis.Options{
			Addr:     host,
			Password: passwd,
			DB:       selector,
		})

		if client != nil {
			driver.Client = client
		}
		return nil
	}
	return errors.New("open failed(no driver object.)")
}

func(driver *RedisDriver) Set( ctx context.Context, key string, value interface{}, expiration time.Duration ) error {
	if driver != nil {
		if driver.Client != nil {
			err := driver.Client.Set(ctx, key, value, expiration).Err()
			if err != nil {
				return err
			}
			return nil
		} else {
			return errors.New("no connection")		
		}
	}
	return errors.New("no driver")
}

func(driver *RedisDriver) Get( ctx context.Context, key string ) (interface{}, error) {
	if driver != nil {
		if driver.Client != nil {
			value, err := driver.Client.Get(ctx, key).Result()
			if err != nil {
				return nil, err
			}
			return value, nil
		} else {
			return nil, errors.New("no connection")	
		}
	}
	return nil, errors.New("no driver")
}

func(driver *RedisDriver) Delete( ctx context.Context, key ...string ) error {
	if driver != nil {
		if driver.Client != nil {
			err := driver.Del( ctx, key... ).Err()
			if err != nil {
				return err
			}
			return nil
		} else {
			return errors.New("no connection")	
		}
	}
	return errors.New("no driver")
}

func(driver *RedisDriver) MapSet( ctx context.Context, mapKey string, key string, value interface{} ) error {
	if driver != nil {
		if driver.Client != nil {
			err := driver.Client.HMSet(ctx, mapKey, key, value).Err()
			if err != nil {
				return err
			}
			return nil
		} else {
			return errors.New("no connection")		
		}
	}
	return errors.New("no driver")
}

func(driver *RedisDriver) MapGet( ctx context.Context, mapKey string, key string ) (interface{}, error) {
	if driver != nil {
		if driver.Client != nil {
			value, err := driver.Client.HMGet(ctx, mapKey, key).Result()
			if err != nil {
				return nil, err
			}
			return value, nil
		} else {
			return nil, errors.New("no connection")
		}
	}
	return nil, errors.New("no driver")
}
