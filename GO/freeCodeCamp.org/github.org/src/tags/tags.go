package main

import (
	"fmt"
	"reflect"
)

type animalia struct {
	Name string ` required max: "150"`
	Place string
}

func main()  {
	t := reflect.TypeOf(animalia{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

}
