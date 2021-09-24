package books

import "time"

// kebutuhan domain book -> domain layer / entity layer -> acuan utama dalam domain
type Domain struct {
	ID        int
	Title     string
	ISBN      string
	Writer    string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//fungsi apa saja dari Domain domain ngelakuin hal apa saja-> interface of data layer -> fungsi yg dibutuhkan oleh domain (business logic)
type Service interface {
	Append(book *Domain) (*Domain, error)
	Update(book *Domain, id int) (*Domain, error)
	FIndByID(id int) (*Domain, error)
	Available(generalSearch string) []Domain
	Delete(book *Domain, id int) (*Domain, error)
}

//fungsi untuk database -> interface of data layer -> fungsi untuk membaca / memrintah database, 3rd party, ataupun package lain
type Repository interface {
	Insert(book *Domain) (*Domain, error)
	Update(book *Domain, id int) (*Domain, error)
	FindByID(id int) (*Domain, error)
	FindAll(generalSearch string, availability bool) []Domain
	Delete(book *Domain, id int) (*Domain, error)
}
