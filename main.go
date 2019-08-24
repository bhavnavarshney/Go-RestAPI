package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Name   string  `json:"name"`
	Author *Author `json:"author"`
}

var books []Book

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//getBooks: To get collection of books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&books)
}

//getBook: To get a single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameter := mux.Vars(r) // To get the parameter

	for _, item := range books {
		if item.ID == parameter["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

//createBook: To delete Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	_ = json.NewDecoder(r.Body).Decode(&newBook)
	newBook.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, newBook)
	json.NewEncoder(w).Encode(newBook)
}

//updateBook: To delete Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	parameter := mux.Vars(r)
	for index, item := range books{
		if item.ID == parameter["id"]{
			books = append(books[:index],books[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&newBook)
			newBook.ID = strconv.Itoa(rand.Intn(1000000))
			books = append(books, newBook)
			json.NewEncoder(w).Encode(newBook)
			return 
		}
	}
	

}

//deleteBook: To delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameter := mux.Vars(r)

	for key, item := range books {
		if item.ID == parameter["id"] {
			books = append(books[:key], books[key+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

func main() {
	router := mux.NewRouter()

	//Mock Data
	books = append(books, Book{ID: "1", Isbn: "123424", Name: "The Monk Who Sold His Ferrari", Author: &Author{Firstname: "Robin", Lastname: "Sharma"}})
	books = append(books, Book{ID: "2", Isbn: "43123", Name: "The Alchemist", Author: &Author{Firstname: "Paulo", Lastname: "Coelho"}})

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
