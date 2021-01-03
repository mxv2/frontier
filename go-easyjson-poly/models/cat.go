package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Cat struct {
	nameField string
	ageField  int

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

func (c *Cat) UnmarshalJSON(raw []byte) error {
	var data struct {
		Name         string `json:"name"`
		Age          int    `json:"age"`
		FavoriteFood string `json:"fav_food"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result Cat
	result.nameField = data.Name
	result.ageField = data.Age
	result.FavoriteFood = data.FavoriteFood

	*c = result

	return nil
}
