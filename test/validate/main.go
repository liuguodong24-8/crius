package main

import (
	"fmt"
	"reflect"

	"gitlab.omytech.com.cn/micro-service/Crius/test/validate/tool"
)

// User ....
type User struct {
	ID       int
	Name     string
	UserName string
	Email    string
	Phone    string
}

type UserValidate struct {
	ID       int    `validate:"number,min=1,max=1000"`
	Name     string `validate:"string,min=2,max=10"`
	UserName string `validate:"string"`
	Email    string `validate:"email"`
}

func main() {
	fmt.Println("main")

	userData := &User{
		ID:       10,
		Name:     "adminaaaaaaa",
		UserName: "",
		Email:    "foobar",
		Phone:    "123456",
	}

	validate := &UserValidate{}

	dataKey := reflect.TypeOf(*userData)
	validateKey := reflect.TypeOf(*validate)

	dataElem := reflect.ValueOf(userData).Elem()
	validateElem := reflect.ValueOf(validate).Elem()

	for i := 0; i < dataKey.NumField(); i++ {
		for j := 0; j < validateKey.NumField(); j++ {
			if dataKey.Field(i).Name == validateKey.Field(j).Name {
				validateElem.Field(j).Set(dataElem.Field(i))
			}
		}
	}

	fmt.Println(userData)

	fmt.Println(validate)

	fmt.Println("Errors:")
	for i, err := range tool.ValidateStruct(*validate) {
		fmt.Printf("\t%d. %s\n", i+1, err.Error())
	}
}
