package main

import (
	"bufio"
	"encoding/gob"
	"log"
	"os"
)

type Address struct {
	Type             string
	City             string
	Country          string
}

type VCard struct {
	FirstName	string
	LastName	string
	Addresses	[]*Address
	Remark		string
}


func main() {
	// write your implementation here
	file, err := os.Open("./vcard.gob")
	if err != nil {
		log.Println("Can not open file", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := gob.NewDecoder(reader)
	var card *VCard
	card = new(VCard)
	decoder.Decode(card)
	log.Println(*card)
}