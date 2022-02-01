package driver

import(
	"fmt"
	"testing"
	"sync"
	"time"
)

// DB接続とデータ抽出確認
func TestPostgresqlDriver( t *testing.T ){

	connStr := "user=warehouse password=warehouse host=172.22.10.110 port=5432 dbname=warehousedb sslmode=disable"
	driver := &PostgresqlDriver{}
	driver.Open(connStr)
	defer driver.Close()

	wcaluse := driver.NewWhereCaluses()
	wcaluse.Append("display_id", 7, []interface{}{"test%"})
	
	wc, values := wcaluse.String()

	query := fmt.Sprintf("select display_id from users %s", wc)
	fmt.Printf("query: %s\n%+v\n", query, values)
	rows, err := driver.Query(query, values...)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			t.Fatal()
		}
		fmt.Println("display_id: " + name)
	}

}

// セッションの利用確認用
func TestPostgresPing( t *testing.T ){

	connStr := "user=ezkps password=ezkps host=192.168.172.15 port=5432 dbname=ezkpsdb sslmode=disable"
	driver := &PostgresqlDriver{}
	driver.Open(connStr)
	defer driver.Close()
	
	// n秒利用していないセッションは削除する。
	// セッションが作られてから、設定時間を超えたセッションがクローズされた時点で消える（再接続されない）。
	n := 20
	driver.SetConnMaxLifetime(time.Duration(n)*time.Second)

	wg := sync.WaitGroup{}

	wg.Add(1)
	cnt := 0
	
	go func(c int){
		for {
			if c >= 100 {
				wg.Done()
				break
			}
			err := driver.Ping()
			// セッションの作成
			rows, _ := driver.Query("select subscriber_import_upload_path from system_master")
			rows.Close()	//ここがないと永遠にセッションが増えるので、要注意。意識的にcloseさせる作りにしなければ！

			fmt.Printf("cnt: %d -> %+v\n", c, err)
			time.Sleep(time.Duration(10*time.Second))
			c++
		}
	}(cnt)

	wg.Wait()
}