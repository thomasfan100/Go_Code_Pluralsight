package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//[]byte turns string into byte data type
	w.Write([]byte("Hello from the User Controller!"))
}

//constructor function
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/user/(\d+)/?`),
	}
}
