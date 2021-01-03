package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type PetStore struct {
	pets []Pet

	Staff []*User
	Count int
}

func (p *PetStore) Pets() []Pet {
	return p.pets
}

func (p *PetStore) SetPets(pets []Pet) {
	p.pets = pets
}

func (p *PetStore) String() string {
	petString := "["
	for i, pet := range p.pets {
		if i > 0 {
			petString += ", "
		}
		petString += fmt.Sprintf("%+v", pet)
	}
	petString += "]"

	return fmt.Sprintf("{pets:%s Staff:%+v Count:%d}", petString, p.Staff, p.Count)
}

func (p *PetStore) UnmarshalJSON(raw []byte) error {
	// unmarshal polymorphic field
	var data struct {
		Staff []*User         `json:"staff"`
		Pets  json.RawMessage `json:"pets"`
		Count int             `json:"count"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)

	if err := dec.Decode(&data); err != nil {
		return err
	}

	pets, err := UnmarshalPetItemSlice(bytes.NewBuffer(data.Pets))
	if err != nil && err != io.EOF {
		return err
	}

	var result PetStore
	result.SetPets(pets)

	result.Staff = data.Staff
	result.Count = data.Count

	*p = result

	return nil
}
