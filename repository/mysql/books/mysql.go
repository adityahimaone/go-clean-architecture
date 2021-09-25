package books

import (
	"clean-arc-go/bussiness/books"
	"gorm.io/gorm"
)

type repoBooks struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) books.Repository {
	return &repoBooks{
		DB: db,
	}
}

func (r *repoBooks) Insert(book *books.Domain) (*books.Domain, error) {
	recordBook := fromDomain(*book)
	err := r.DB.Create(&recordBook).Error
	if err != nil {
		return &books.Domain{}, err
	}
	result := toDomain(recordBook)
	return result, nil
}

func (r *repoBooks) Update(book *books.Domain, id int) (*books.Domain, error) {
	return &books.Domain{}, nil
}

func (r *repoBooks) FindByID(id int) (*books.Domain, error) {
	return &books.Domain{}, nil
}

func (r *repoBooks) FindAll(generalSearch string, availability bool) []books.Domain {
	return []books.Domain{}
}

func (r *repoBooks) Delete(book *books.Domain, id int) (*books.Domain, error) {
	return &books.Domain{}, nil
}
