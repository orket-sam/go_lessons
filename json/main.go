package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Language string `json:"language"`
}

func main() {
	var book Book = Book{Title: "Digital fortress", Author: "Orket", Language: "English"}
	byteArray, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		fmt.Println("Oopsy daisy")
	}
	fmt.Println(string(byteArray))
	fmt.Printf("I enjoyed reading %s  by %s \n", book.Title, book.Author)
}
