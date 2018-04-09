package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("SECRET"))
var name = "MANENSE"

// Exist session
func Exist(w http.ResponseWriter, r *http.Request) bool {
	var s = Get(w, r)
	return s.Values["ID"] != nil
}

// Get session
func Get(w http.ResponseWriter, r *http.Request) *sessions.Session {
	var s, e = store.Get(r, name)
	if e != nil {
		s, e = store.New(r, name)
	}
	s.Save(r, w)
	return s
}

// Create session
func Create(w http.ResponseWriter, r *http.Request, id string) {
	var s = Get(w, r)
	s.Values["ID"] = id
	s.Save(r, w)
}

// Erase session
func Erase(w http.ResponseWriter, r *http.Request) {
	var s = Get(w, r)
	s.Options = &sessions.Options{MaxAge: -1}
	s.Save(r, w)
}

// UserID from session
func UserID(w http.ResponseWriter, r *http.Request) string {
	var s = Get(w, r)
	if s.Values["ID"] == nil {
		return ""
	}
	return s.Values["ID"].(string)
}
