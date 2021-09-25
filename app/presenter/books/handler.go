package books

import (
	"clean-arc-go/app/presenter/books/request"
	"clean-arc-go/bussiness/books"
	"github.com/labstack/echo/v4"
	"net/http"
)

//apa yg dia butuhkan
type presenter struct{
	serviceBook books.Service
}

func NewHandler(bookService books.Service) *presenter {
	return &presenter{
		serviceBook: bookService,
	}
}

func (handler *presenter) Insert(echoContext echo.Context) error {
	request := new(request.Book)
	err := echoContext.Bind(&request)
	if err != nil{
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	handler.serviceBook.Append(request)
}
