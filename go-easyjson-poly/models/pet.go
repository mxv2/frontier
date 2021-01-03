package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Pet interface {
	Name() string
	SetName(name string)

	Age() int
	SetAge(age int)

	Kind() string
}

func UnmarshalPetItemSlice(reader io.Reader) ([]Pet, error) {
	var elements []json.RawMessage

	dec := json.NewDecoder(reader)
	if err := dec.Decode(&elements); err != nil {
		return nil, err
	}

	var result []Pet
	for _, element := range elements {
		obj, err := unmarshalPetItem(element)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

func unmarshalPetItem(data []byte) (Pet, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	var getType struct {
		Kind string `json:"kind"`
	}
	dec := json.NewDecoder(buf)
	if err := dec.Decode(&getType); err != nil {
		return nil, err
	}

	switch getType.Kind {
	case "cat":
		var result Cat
		dec := json.NewDecoder(buf2)
		if err := dec.Decode(&result); err != nil {
			return nil, err
		}
		return &result, nil
	case "dog":
		var result Dog
		dec := json.NewDecoder(buf2)
		if err := dec.Decode(&result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, fmt.Errorf("unexpected kind %q", getType)
}
