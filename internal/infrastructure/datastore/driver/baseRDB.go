package driver

import (
	"fmt"
	"strings"
)

type RDBFilterOperation int
const (
	OpEqual RDBFilterOperation = iota
	OpMore
	OpMorethan
	OpLess 
	OpLessthan
	OpBetween
	OpIn
	OpLike
)

type RDBSortDirection int
const (
	OrderByASC RDBSortDirection = iota
	OrderByDESC
)

type baseRDB struct {
}

type whereCaluse struct {
	Column string
	Op     int
	Values []interface{} 
}

type whereCaluses []*whereCaluse


type orderCaluse struct {
	Column    string
	Direction RDBSortDirection
}

type orderCaluses []*orderCaluse
type WhereCaluses interface{
	String() (string, []interface{})
	Append(column string, op int, values []interface{})
}
type OrderCaluses interface {
	String() string
	Append(col string, dir int)
}

func (base *baseRDB) NewWhereCaluses() WhereCaluses{
	return &whereCaluses{}
}

func (base *baseRDB) NewOrderCaluses() OrderCaluses{
	return &orderCaluses{}
}

func (w *whereCaluse) string(cnt int) (string, int) {
	var op string
	var include int
	switch RDBFilterOperation(w.Op) {
	case OpLess:
		op = ">"
	case OpLessthan:
		op = ">="
	case OpMore:
		op = "<"
	case OpMorethan:
		op = "<="
	case OpBetween:
		op = "between"
	case OpIn:
		op = "in"
	case OpLike:
		op = "like"
	default:
		op = "="
	}

	var whereStr string
	if RDBFilterOperation(w.Op) == OpBetween{
		// 2以外は無視
		if len(w.Values) == 2 {
			whereStr = fmt.Sprintf("%s %s $%d and $%d", w.Column, op, cnt, cnt+1)
			include = 2
		}
	} else if RDBFilterOperation(w.Op) == OpIn {
		args := []string{}
		for i, _ := range w.Values{
			args = append(args, fmt.Sprintf("$%d", cnt+i ))
			include = i+1
		}
		if len(w.Values) != 0 {
			whereStr = fmt.Sprintf("%s %s ( %s )", w.Column, op, strings.Join(args,", ") )
		}
	} else {
		whereStr = fmt.Sprintf("%s %s $%d", w.Column, op, cnt)
		include = 1
	}
	return whereStr, cnt+include
}

func(w *whereCaluses) String() (string, []interface{}){

	var whereStr string
	var offset int
	var values []interface{}
	if len(*w) != 0 {
		for i, v := range *w {
			if i == 0 {
				w, o := v.string(1)
				whereStr = fmt.Sprintf("where %s", w)
				offset = o
			} else {
				w, o := v.string(offset)
				whereStr = fmt.Sprintf("%s and %s", whereStr, w)
				offset = o
			}
			for _, val := range v.Values {
				values = append(values, val)
			}
		}
		
	}
	return whereStr, values
}

func (w *whereCaluses) Append(column string, op int, values []interface{}) {
	*w = append(*w, &whereCaluse{Column: column, Op: op, Values: values})
}

func (o *orderCaluse) string() string {
	if RDBSortDirection(o.Direction) == OrderByDESC {
		return fmt.Sprintf("%s desc", o.Column)
	}
	return o.Column
}

func (o *orderCaluses) Append(col string, dir int) {
	*o = append(*o, &orderCaluse{Column: col, Direction: RDBSortDirection(dir)})
}

func (o *orderCaluses) String() string {

	orderString := ""
	for cnt, lo := range *o {
		if cnt == 0 {
			orderString = fmt.Sprintf("order by %s", lo.string())
		} else {
			orderString = fmt.Sprintf("%s, %s", orderString, lo.string())
		}
	}
	return orderString
}
