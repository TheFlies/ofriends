package login

import (
	"fmt"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
	"io"
	"net/http"
	"os"
)

type (
	// Handler is index web handler
	Handler struct{}
)

func New() *Handler {
	return &Handler{}
}
func (web *Handler) Login(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("web/login.html")
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	if _, err := io.Copy(w, f); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
}
func (web *Handler) Authenticaiton(w http.ResponseWriter, r *http.Request) {
	usename := r.FormValue("usename")
	password := r.FormValue("password")
	fmt.Print(usename)
	token, ok := Authenticationuser(usename, password)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(token))
	}
}
