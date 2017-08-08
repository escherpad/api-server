//Package handlers implements route handlers.
package handlers

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not implemented")

}

func ResetPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not implemented")
}

func SetPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not implemented")
}

func CreateTeam(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not implemented")
}

func SignupForTeam(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not implemented")
}

func Authenticate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not implemented")
}

func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not implemented")
}

func TestDummy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not implemented")
}
