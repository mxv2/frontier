package easymodels

import (
	"fmt"
)

type Dog struct {
	NameField string `json:"name"`
	AgeField  int    `json:"age"`

	FavoriteToy string `json:"fav_toy"`
}

func (d *Dog) Name() string {
	return d.NameField
}

func (d *Dog) SetName(name string) {
	d.NameField = name
}

func (d *Dog) Age() int {
	return d.AgeField
}

func (d *Dog) SetAge(age int) {
	d.AgeField = age
}

func (d *Dog) Kind() string {
	return "dog"
}

func (d *Dog) String() string {
	return fmt.Sprintf("{nameField:%s ageField:%d FavoriteToy:%s}", d.NameField, d.AgeField, d.FavoriteToy)
}
