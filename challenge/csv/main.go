package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Book struct {
	title    string
	price    float64
	quantity int
}

func main() {
	file, err := os.OpenFile("./products.txt", os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	var books []Book
	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, book := range rawCSVdata {
		var newBook Book
		newBook.title = book[0]
		price, err := strconv.ParseFloat(book[1], 64)
		if err != nil {
			fmt.Println(err)
			price = -1
		}
		quantity, err := strconv.Atoi(book[2])
		if err != nil {
			fmt.Println(err)
			quantity = -1
		}
		newBook.price = price
		newBook.quantity = quantity
		books = append(books, newBook)
	}
	fmt.Println(books)

}