package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dimnyan/bookstore-api/pkg/models"
	"github.com/dimnyan/bookstore-api/pkg/utils"
	"github.com/gorilla/mux"
)

//var NewBook models.Book

func GetBook(w http.ResponseWriter) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		fmt.Println(err)
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book")
	}
	bookDetail, db := models.GetBookById(int64(ID))
	if db.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write(res)
	if err != nil {
		fmt.Println(err)
	}
}

// func DeleteBook(w http.ResponseWriter, r *http.Request){}

// func UpdateBooks(w http.ResponseWriter, r *http.Request){}
