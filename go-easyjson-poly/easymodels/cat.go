package easymodels

import (
	"fmt"
)

type Cat struct {
	nameField string `json:"name"`
	ageField  int    `json:"age"`

	FavoriteFood string `json:"fav_food"`
}

func (c *Cat) Name() string {
	return c.nameField
}

func (c *Cat) SetName(name string) {
	c.nameField = name
}

func (c *Cat) Age() int {
	return c.ageField
}

func (c *Cat) SetAge(age int) {
	c.ageField = age
}

func (c *Cat) Kind() string {
	return "cat"
}

func (c *Cat) String() string {
	return fmt.Sprintf("{nameField:%s ageField:%d FavoriteFood:%s}", c.nameField, c.ageField, c.FavoriteFood)
}
