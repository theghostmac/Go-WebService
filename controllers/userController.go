package controllers

import (
	"golang.org/x/crypto/openpgp/packet"
	"net/http"
)

type userController struct {
	packet.UserID	*regexp.Regexp
}

func (userCon userController) ServeHTTP(writer http.ResponseWriter, reader *http.Request) {
	write, err := writer.Write([]byte("This is a web service."))
	if err != nil {
		return
	}
}

func newUserController()  {
	
}