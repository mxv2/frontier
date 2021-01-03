package easymodels

import (
	"fmt"
)

type PetStore struct {
	pets []Pet `json:"pets"`

	Staff []*User `json:"staff"`
	Count int     `json:"count"`
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
