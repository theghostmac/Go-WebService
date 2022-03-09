package controllers

import "net/http"

type userController struct {
	// empty for now
}

func (userCon userController) ServeHTTP(writer http.ResponseWriter, reader *http.Request) {
	write, err := writer.Write([]byte("This is a web service."))
	if err != nil {
		return
	}
}
