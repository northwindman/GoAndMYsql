package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/northwindman/GoAndMYsql/pkg/models"
	"github.com/northwindman/GoAndMYsql/pkg/utils"
	"log"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, err := json.Marshal(newBooks)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write(res)
	if errWrite != nil {
		log.Println(errWrite)
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write(res)
	if errWrite != nil {
		log.Println(errWrite)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	log.Println(CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		log.Println(err)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, errJson := json.Marshal(book)
	if errJson != nil {
		log.Println(errJson)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//func UpdateBook(w http.ResponseWriter, r *http.Request) {
//	var updateBook = &models.Book{}
//	utils.ParseBody(r, updateBook)
//	vars := mux.Vars(r)
//	bookId := vars["bookId"]
//	ID, err := strconv.ParseInt(bookId, 0, 0)
//	if err != nil {
//		fmt.Println("error while parse")
//	}
//	bookDetails, db := models.GetBookById(ID)
//	if updateBook.Name != "" {
//		bookDetails.Name = updateBook.Name
//	}
//	if updateBook.Author != "" {
//		bookDetails.Author = updateBook.Author
//	}
//	db.Save(&bookDetails)
//	res, _ := json.Marshal(bookDetails)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	_, errWrite := w.Write(res)
//	if errWrite != nil {
//		log.Println(errWrite)
//	}
//}
