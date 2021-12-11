package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	err := ioutil.WriteFile(p.Title+".txt", p.Body, 0666)
	return err
}

func load(title string) (*Page, error) {
	content, err := ioutil.ReadFile(title + ".txt")
	return &Page{title, content}, err
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := load("TestPage")
	fmt.Println(string(p2.Body))
}