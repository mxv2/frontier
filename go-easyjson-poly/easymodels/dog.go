package easymodels

import (
	"fmt"
)

type Dog struct {
	nameField string `json:"name"`
	ageField  int    `json:"age"`

	FavoriteToy string `json:"fav_toy"`
}

func (d *Dog) Name() string {
	return d.nameField
}

func (d *Dog) SetName(name string) {
	d.nameField = name
}

func (d *Dog) Age() int {
	return d.ageField
}

func (d *Dog) SetAge(age int) {
	d.ageField = age
}

func (d *Dog) Kind() string {
	return "dog"
}

func (d *Dog) String() string {
	return fmt.Sprintf("{nameField:%s ageField:%d FavoriteToy:%s}", d.nameField, d.ageField, d.FavoriteToy)
}
