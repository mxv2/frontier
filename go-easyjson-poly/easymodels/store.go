package easymodels

import (
	"fmt"
)

type PetStore struct {
	petsField []Pet

	Staff []*User `json:"staff"`
	Count int     `json:"count"`
}

func (p *PetStore) Pets() []Pet {
	return p.petsField
}

func (p *PetStore) SetPets(pets []Pet) {
	p.petsField = pets
}

func (p *PetStore) String() string {
	petString := "["
	for i, pet := range p.petsField {
		if i > 0 {
			petString += ", "
		}
		petString += fmt.Sprintf("%+v", pet)
	}
	petString += "]"

	return fmt.Sprintf("{petsField:%s Staff:%+v Count:%d}", petString, p.Staff, p.Count)
}
