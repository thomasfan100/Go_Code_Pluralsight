//setup routing
package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	//allows handling of files with no extension and those that have more after /users
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

//creates an encoder and passes in data to that encoder
func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
