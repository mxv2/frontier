package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mailru/easyjson"

	"github.com/mxv2/frontier/go-easyjson-poly/easymodels"
	"github.com/mxv2/frontier/go-easyjson-poly/models"
)

func main() {
	raw, err := ioutil.ReadFile("message.json")
	if err != nil {
		log.Fatal(err)
	}

	unmarshalJson(raw)
	unmarshalEasyJson(raw)
}

func unmarshalJson(raw []byte) {
	buf := bytes.NewBuffer(raw)

	var store models.PetStore
	dec := json.NewDecoder(buf)
	err := dec.Decode(&store)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("    json: %+v\n", &store)
}

func unmarshalEasyJson(raw []byte) {
	buf := bytes.NewBuffer(raw)

	var store easymodels.PetStore
	err := easyjson.UnmarshalFromReader(buf, &store)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("easyjson: %+v\n", &store)
}
