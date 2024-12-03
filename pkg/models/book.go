package models

import (
	"github.com/dimnyan/bookstore-api/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name   string `gorm:"json:name"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

//func DeleteBook(Id int64) Book {
//	var book Book
//	db.Where("ID=?", Id).Delete(book)
//	return book
//}
