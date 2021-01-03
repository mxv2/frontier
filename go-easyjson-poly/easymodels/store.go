package easymodels

import (
	"fmt"
)

type PetStore struct {
	PetsField PetSliceAdapter `json:"pets"`

	Staff []*User `json:"staff"`
	Count int     `json:"count"`
}

func (p *PetStore) Pets() []Pet {
	return p.PetsField
}

func (p *PetStore) SetPets(pets []Pet) {
	p.PetsField = pets
}

func (p *PetStore) String() string {
	petString := "["
	for i, pet := range p.PetsField {
		if i > 0 {
			petString += ", "
		}
		petString += fmt.Sprintf("%+v", pet)
	}
	petString += "]"

	return fmt.Sprintf("{petsField:%s Staff:%+v Count:%d}", petString, p.Staff, p.Count)
}
