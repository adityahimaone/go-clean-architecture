package books

//service butuh repository interface
type serviceBooks struct {
	repository Repository
}

//init service : service -> butuh Repository -> return Service
func NewService(repoBook Repository) Service {
	return &serviceBooks{
		repository: repoBook,
	}
}

func (servBook *serviceBooks) Append(book *Domain) (*Domain, error) {
	respon, err := servBook.repository.Insert(book)
	if err != nil {
		return &Domain{}, err
	}
	return respon, nil
}
func (servBook *serviceBooks) Update(book *Domain, id int) (*Domain, error) {
	respon, err := servBook.repository.Update(book, id)
	if err != nil {
		return &Domain{}, err
	}
	return respon, nil
}
func (servBook *serviceBooks) FindByID(id int) (*Domain, error) {
	respon, err := servBook.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return respon, nil
}
func (servBook *serviceBooks) Available(generalSearch string, availability bool) (*[]Domain, error) {
	findAvailability := false
	if availability {
		findAvailability = true
	}
	respon, err := servBook.repository.FindAll(generalSearch, findAvailability)
	if err != nil {
		return &[]Domain{}, err
	}
	return respon, nil
}
func (servBook *serviceBooks) Delete(book *Domain, id int) (*Domain, error) {
	respon, err := servBook.repository.Delete(book, id)
	if err != nil {
		return &Domain{}, err
	}
	return respon, nil
}
