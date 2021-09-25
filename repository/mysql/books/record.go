package books

import (
	"clean-arc-go/bussiness/books"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Title  string
	ISBN   string
	Writer string
	Status bool
}

//dari record ->
func toDomain(rec Record) books.Domain {
	return books.Domain{
		ID:        int(rec.ID),
		Title:     rec.Title,
		ISBN:      rec.ISBN,
		Writer:    rec.Writer,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(domain books.Domain) Record {
	return Record{
		Title:  domain.Title,
		ISBN:   domain.ISBN,
		Writer: domain.Writer,
		Status: domain.Status,
	}
}
