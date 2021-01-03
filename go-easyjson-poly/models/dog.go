package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Dog struct {
	nameField string
	ageField  int

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

func (d *Dog) UnmarshalJSON(raw []byte) error {
	var data struct {
		Name        string `json:"name"`
		Age         int    `json:"age"`
		FavoriteToy string `json:"fav_toy"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result Dog
	result.nameField = data.Name
	result.ageField = data.Age
	result.FavoriteToy = data.FavoriteToy

	*d = result

	return nil
}
