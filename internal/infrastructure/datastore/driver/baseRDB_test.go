package driver

import (
	"fmt"
	"testing"
)

func TestBaseRDB(t *testing.T){

	base := &baseRDB{}

	wcaluses := base.NewWhereCaluses()
	wcaluses.Append("aaaa", 7, []interface{}{"%%test%%"})

	whereStr, values := wcaluses.String()
	fmt.Printf("caluse: %s => %+v\n", whereStr, values)

	ocaluses := base.NewOrderCaluses()
	ocaluses.Append("aaaa", 1)
	fmt.Printf("caluse: %s\n", ocaluses.String())	

}
