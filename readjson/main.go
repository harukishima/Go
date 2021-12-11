package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
  Name string `json:"personName"`
  Age int `json:"personAge"`
}

func main() {
  b := []byte(`{"personName": "Obama", "personAge": 57}`)
  var p Person
  // Unmarshalling
  json.Unmarshal(b, &p)
  fmt.Println(p)
  // Marshalling
  js, _ := json.Marshal(p)
  fmt.Printf("%s\n", js)
}