package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
	"../models"
)

// GET /api/books
func GetBooks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Library)
}

// GET /api/book/{id}
func GetBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	
	// get params from request
	params := mux.Vars(r)

	for _, item := range models.Library {
		// look for book by id
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

		// else output output book w/ empty values
		json.NewEncoder(w).Encode(&models.Book{})
	}
}

// POST /api/books 
// - TODO: data validation
func CreateBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book

	// get data from request body and parse it as JSON of a book and initialize the book
	_ = json.NewDecoder(r.Body).Decode(&book)

	// set ID - not for prod
	book.ID = strconv.Itoa(rand.Intn(100000))
	book.Author.ID = strconv.Itoa(rand.Intn(100000))

	// append book to library
	models.Library = append(models.Library, book)

	// display created book
	json.NewEncoder(w).Encode(book)
}

// GET /api/books/{id} 
// - TODO: handle partial update
// - TODO: data validation
func UpdateBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	// get params from request
	params := mux.Vars(r)


	for index , item := range models.Library {
		// look for book by id
		if item.ID == params["id"] {
			// create updated book
			var book models.Book
			_ = json.NewDecoder(r.Body).Decode(&book)

			// keep the original ID
			book.ID = item.ID
			
			// delete outdated version of book from library
			models.Library = append(models.Library[:index], models.Library[index+1:]...)

			// add updated book
			models.Library = append(models.Library, book)
			break
		}
	}

	// display updated library
	json.NewEncoder(w).Encode(models.Library)
}

// DELETE /api/books/{id}
func DeleteBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	// get params from request
	params := mux.Vars(r)

	for index , item := range models.Library {
		// look for book by id
		if item.ID == params["id"] {
			// delete from library
			models.Library = append(models.Library[:index], models.Library[index+1:]...)
			break
		}
	}

	// display updated library
	json.NewEncoder(w).Encode(models.Library)
}
