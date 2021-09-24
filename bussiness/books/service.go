package books

//service butuh repository intrface
type serviceBooks struct {
	repository Repository
}

//init service : service -> butuh Repository -> return Service
func NewService(repoBook Repository) Service {
	return &serviceBooks{
		repository: repoBook,
	}
}

func (servBook *serviceBooks)Append(book *Domain) (*Domain, error){

}
func (servBook *serviceBooks)Update(book *Domain, id int) (*Domain, error){

}
func (servBook *serviceBooks)FIndByID(id int) (*Domain, error){

}
func (servBook *serviceBooks)Available(generalSearch string) []Domain{

}
func (servBook *serviceBooks)Delete(book *Domain, id int) (*Domain, error){

}
