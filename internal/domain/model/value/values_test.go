package value

import(
	"fmt"
//	"encoding/json"
//	"errors"
	"time"
	"testing"
)

func TestValues(t *testing.T){

	expr := time.Now()
	recv := time.Now().Add(1)
	u := NewUser("test_1", 1, "uuid1", 1, nil, nil, nil, &expr, &recv )
	passphrase, err := u.GetPassphrase("password1")
	if err != nil {
		t.Fatal(err)
	}
	token, err := u.GetOnetimeToken()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("user: %+v, passphrase: %s, token: %s\n", u, string(passphrase), token)
/*
	vt := NewValueTime(nil)
	fmt.Printf("vtime: %v\n", vt)

	var data = `
    {"Time": "2022/01/26 11:27:18"}
	`
	type ttime struct {
		Time *valueTime
	}
	tt := ttime{}
	fmt.Println(json.Unmarshal([]byte(data), &tt))
	tt.Time.Add(10 * time.Hour)

	fmt.Println(tt.Time.String())
	jstr, err := json.Marshal(&tt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(jstr))
*/
	json, err := u.MarshalJson(MarshalSafe)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(json))

	jsonALL, err := u.MarshalJson(MarshalALL)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(jsonALL))


	var eu valueUser
	err = eu.UnmarshalJson(jsonALL)
	fmt.Printf("valueTime: %+v\n", &eu)

	if err != nil {
		t.Fatal(err)
	}

/*
	e1 := NewError(ErrorNothing, ErrorMsgNothing, nil)
	fmt.Printf("%s\n", e1.Error())
	e2 := NewError(ValueErrorFaildJson, ValueErrorMsgFaildJson, errors.New("test 2"))
	fmt.Printf("%s\n", e2.Error())
	e3 := NewError(ValueErrorFaildGetPassphrase, ValueErrorMsgFaildGetPassphrase, errors.New("test 3"))
	fmt.Printf("%s\n", e3.Error())
*/
}
