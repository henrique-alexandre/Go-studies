package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {

	router := mux.NewRouter()

	books = append(books,

		Book{ID: 1, Title: "Golang pointers", Author: "Mr. Golang", Year: "2010"},
		Book{ID: 2, Title: "Goroutines", Author: "Mr. Goroutine", Year: "2011"},
		Book{ID: 3, Title: "Golang routers", Author: "Mr. Router", Year: "2012"},
		Book{ID: 4, Title: "Golang concurrency", Author: "Mr. Currency", Year: "2013"},
		Book{ID: 5, Title: "Golang good parts", Author: "Mr. Good", Year: "2014"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}

}

func getBooks(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

	p := mux.Vars(r)
	i, _ := strconv.Atoi(p["id"])

	for _, b := range books {
		if b.ID == i {
			json.NewEncoder(w).Encode(&b)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {

	var b Book
	json.NewDecoder(r.Body).Decode(&b)
	books = append(books, b)
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {

	var newBook Book
	json.NewDecoder(r.Body).Decode(&newBook)
	for i, book := range books {
		if book.ID == newBook.ID {
			books[i] = newBook
		}
	}
	json.NewEncoder(w).Encode(books)

}
func removeBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, j := range books {
		if j.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(books)
}
