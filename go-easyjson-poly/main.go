package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mxv2/frontier/go-easyjson-poly/models"
)

func main() {
	raw, err := ioutil.ReadFile("message.json")
	if err != nil {
		log.Fatal(err)
	}

	unmarshalJson(raw)
}

func unmarshalJson(raw []byte) {
	buf := bytes.NewBuffer(raw)

	var store models.PetStore
	dec := json.NewDecoder(buf)
	err := dec.Decode(&store)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("json: %+v\n", &store)
}
