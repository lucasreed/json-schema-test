package main

import (
	"fmt"
	"io/ioutil"

	"github.com/qri-io/jsonschema"
)

func main() {
	schemaData, err := ioutil.ReadFile("schema.json")
	if err != nil {
		fmt.Println("Got an error reading schema json file")
	}
	rs := &jsonschema.RootSchema{}
	if err := rs.UnmarshalJSON(schemaData); err != nil {
		panic("unmarshal schema: " + err.Error())
	}

	podData, err := ioutil.ReadFile("pod.json")
	if err != nil {
		fmt.Println("Got an error reading pod json file")
	}

	errors, _ := rs.ValidateBytes(podData)
	for _, err := range errors {
		fmt.Println(err)
	}
}
