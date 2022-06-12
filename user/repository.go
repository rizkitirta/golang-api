package user

import "gorm.io/gorm"

type UserRepository interface {
	FindAll() ([]User, error)
	FindById(ID int) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(user User) (User, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) *repository {
	return &repository{DB}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.DB.Find(&users).Error

	return users, err
}

func (r *repository) FindById(ID int) (User, error) {
	var user User
	err := r.DB.Find(&user, ID).Error

	return user, err
}

func (r *repository) Create(user User) (User, error) {
	err := r.DB.Create(&user).Error

	return user, err
}

func (r *repository) Update(user User) (User, error) {
	err := r.DB.Save(&user).Error

	return user, err
}

func (r *repository) Delete(user User) (User, error) {
	err := r.DB.Delete(&user).Error

	return user, err
}
