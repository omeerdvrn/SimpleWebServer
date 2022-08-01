package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Person struct {
	name    string `json: "name"`
	age     string `json: "age"`
	address string `json: "address"`
	salary  string `json: "salary"`
}

func (p Person) SaveToJSON(name string) {
	file, _ := json.Marshal(p)
	_ = ioutil.WriteFile(name+".json", file, os.ModePerm)
}
