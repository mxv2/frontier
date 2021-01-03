package easymodels

import (
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) String() string {
	return fmt.Sprintf("{Name:%s Age:%d}", u.Name, u.Age)
}
