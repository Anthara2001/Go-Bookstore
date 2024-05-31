package controllers

import (
	"Bookstore/pkg/models"
	"Bookstore/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(res http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()
	bookData, _ := json.Marshal(newBooks)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(bookData)
}

func GetBookById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	newBook, _ := models.GetBookById(ID)
	bookData, _ := json.Marshal(newBook)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(bookData)
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(req, CreateBook)
	b := createBook.CreateBook()
	bookData, _ := json.Marshal(b)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(bookData)
}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	book := models.DeleteBook(ID)
	bookData, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(bookData)
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	updateBook := &models.Book{}
	utils.ParseBody(req, updateBook)
	bookData, _ := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookData.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookData.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookData.Publication = updateBook.Publication
	}

}
