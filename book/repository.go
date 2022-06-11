package book

import "gorm.io/gorm"

type BookRepository interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book Book) (Book, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) *repository {
	return &repository{DB}
}

func (r *repository) FindAll() ([]Book,error) {
	var books []Book
	err := r.DB.Find(&books).Error
	
	return books, err
}

func (r *repository) FindById(ID int) (Book,error) {
	var book Book
	err := r.DB.First(&book, ID).Error

	return book, err
}

func (r *repository) Create(book Book) (Book,error) {
	err := r.DB.Create(&book).Error

	return book, err
}