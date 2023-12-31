package model

import (
	"reflect"
)

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func (u *User) GetUserDataTypes(data string) (string, error) {
	switch data {
	case reflect.TypeOf(u.Id).Name():
		return "int", nil
	case reflect.TypeOf(u.Name).Name():
		return "string", nil
	default:
		return "", nil
	}
}
