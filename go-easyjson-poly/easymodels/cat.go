package easymodels

import (
	"fmt"
)

type Cat struct {
	NameField string `json:"name"`
	AgeField  int    `json:"age"`

	FavoriteFood string `json:"fav_food"`
}

func (c *Cat) Name() string {
	return c.NameField
}

func (c *Cat) SetName(name string) {
	c.NameField = name
}

func (c *Cat) Age() int {
	return c.AgeField
}

func (c *Cat) SetAge(age int) {
	c.AgeField = age
}

func (c *Cat) Kind() string {
	return "cat"
}

func (c *Cat) String() string {
	return fmt.Sprintf("{nameField:%s ageField:%d FavoriteFood:%s}", c.NameField, c.AgeField, c.FavoriteFood)
}
