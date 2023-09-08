package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/grim-firefly/golang-api/pkg/models"
	"github.com/grim-firefly/golang-api/pkg/utils"
)

var NewBook models.Book

// all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// getting single books
func GetBookById(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error in parsing id")
	}
	book, _ := models.GetBookById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// creating books
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	book := CreateBook.CreateBook()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// update books
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook := &models.Book{}
	utils.ParseBody(r, UpdateBook)

	bookId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error in parsing id")
	}
	bookDetails, db := models.GetBookById(id)
	if UpdateBook.Name != "" {
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		bookDetails.Publication = UpdateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// deleting books
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error in parsing id")
	}
	book := models.DeleteBook(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
