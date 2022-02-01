package Error

import(
	"fmt"
	"testing"
	"errors"
)

func TestError(t *testing.T){

	e1 := NewNoRepositoryError()
	fmt.Println(e1.Error(), e1.(NoRepositoryError).Caller())
	e2 := NewNoServiceError()
	fmt.Println(e2.Error(), e2.(NoServiceError).Caller())
	e3 := NewNotFoundError("src", errors.New("new"))
	fmt.Println(e3.Error(), e3.(NotFoundError).Caller())



}