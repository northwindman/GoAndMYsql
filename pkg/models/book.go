package models

import (
	"github.com/jinzhu/gorm"
	"github.com/northwindman/GoAndMYsql/pkg/config"
	"log"
	"net/http"
)

var db *gorm.DB

type Page struct {
	gorm.Model
	Number int    `json:"number"`
	Text   string `json:"text"`
	BookID uint   `json:"book_id"`
}

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Pages       []Page `gorm:"foreignkey:BookID" json:"pages"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}, &Page{})
}

// --------------------- для страниц ---------------------
func getPagesForBook(bookId uint) []Page {
	var Pages []Page
	db.Where("book_id = ?", bookId).Find(&Pages)
	return Pages
}

// ---------------------------------------------------------

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	for i := 0; i < len(Books); i++ {
		Books[i].Pages = getPagesForBook(Books[i].ID)
	}
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	getBook.Pages = getPagesForBook(getBook.ID)
	return &getBook, db
}

func DeleteBook(ID int64) int {
	var book Book
	books := GetAllBooks()
	for i := 0; i < len(books); i++ {
		if books[i].ID == uint(ID) {
			db.Where("ID=?", ID).Delete(book)

			pages := getPagesForBook(uint(ID))

			if len(pages) > 0 {
				for i := 0; i < len(pages); i++ {
					db.Where("book_id=?", uint(ID)).Delete(pages[i])
				}
				log.Println("All pages have been deleted")
				return http.StatusOK
			} else {
				log.Println("The book does not have any pages")
				return http.StatusOK
			}

		}
	}
	return http.StatusNotFound
}
