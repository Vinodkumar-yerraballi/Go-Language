package main

import (
	"fmt"
	"reflect"
)

type User struct {
	UserName string `maxsize:"20"`
	Password string `minsize:"8"`
}

func main() {
	u := User{}
	t := reflect.TypeOf(u)
	// field := t.Field(0)
	// fmt.Println(field.Tag.Get("maxsize"))
	field, _ := t.FieldByName("Password")
	fmt.Println(field.Tag)

}
