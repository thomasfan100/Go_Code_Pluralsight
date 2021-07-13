//setup routing
package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()

	//allows handling of files with no extension and those that have more after /users
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
