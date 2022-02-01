package driver

import(
	"fmt"
	"testing"
	"context"
	"time"
)

// DB接続とデータ抽出確認
func TestRedisDriver( t *testing.T ){

	ctx := context.Background()

	driver := &RedisDriver{}
	driver.Open("172.22.10.110:6379", "", 0)

	fmt.Println("Set")
	if err := driver.Set(ctx, "test:redisdriver", "!!test!!", time.Duration(0)); err != nil {
		t.Fatal()
	}

	fmt.Println("Get")
	if value, err := driver.Get(ctx, "test:redisdriver"); err != nil {
		t.Fatal()
	} else {
		fmt.Printf("result: %+v\n", value)
	}

	fmt.Println("MapSet")
	if err := driver.MapSet(ctx, "test:redisdriver:map", "map1", 10001); err != nil {
		t.Fatal()
	}

	fmt.Println("MapGet")
	if value, err := driver.MapGet(ctx, "test:redisdriver:map", "map1"); err != nil {
		t.Fatal()
	} else {
		fmt.Printf("result: %+v\n", value)
	}

	fmt.Println("Delete")
	if err := driver.Delete(ctx, "test:redisdriver:map"); err != nil {
		t.Fatal()
	}

	fmt.Println("Set with expiration (3sec)")
	if err := driver.Set(ctx, "test:redisdriver", "!!test!!", time.Duration(3)*time.Second); err != nil {
		t.Fatal()
	}

	time.Sleep(5*time.Second)

	defer func() {
		fmt.Println("Get after 5 seconds")
		if value, err := driver.Get(ctx, "test:redisdriver"); err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("result: %+v\n", value)
			t.Fatal()
		}
	}()

}