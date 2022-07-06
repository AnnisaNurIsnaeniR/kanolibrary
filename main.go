package main

import (
	"fmt"
	"kanolibrary/router"
	"net/http"
)

type books struct {
	Title     string `bson:"title"`
	Author    string `bson:"author"`
	Publisher string `bson:"publisher"`
	Synopsis  string `bson:"synopsis"`
}

func main() {
	// crud.Insert("books", books{Title: "Mandi biar ga bau", Author: "Lala", Publisher: "Bintang Pelita", Synopsis: "Blalala"})
	//crud.Find("books", books{})
	//crud.Update()
	//crud.Remove()
	http.HandleFunc("/books", router.Books)
	http.HandleFunc("/books/find", router.Book)
	http.HandleFunc("/books/insert", router.InsertBook)
	http.HandleFunc("/books/update", router.UpdateBook)
	http.HandleFunc("/books/delete", router.DeleteBook)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
