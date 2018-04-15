package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("SECRET"))
var name = "MANENSE_SESSION"

// GetSession func
func GetSession(r *http.Request) *sessions.Session {
	var s, _ = store.Get(r, name)
	return s
}

// New Session
func New(w http.ResponseWriter, r *http.Request, i string) {
	var s = GetSession(r)
	s.Values["ID"] = i
	s.Save(r, w)
}

// Erase Session
func Erase(w http.ResponseWriter, r *http.Request) {
	var s = GetSession(r)
	s.Options = &sessions.Options{MaxAge: -1}
	s.Save(r, w)
}
