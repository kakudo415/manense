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

// Get Session
func Get(w http.ResponseWriter, r *http.Request) string {
	var s = GetSession(r)
	return s.Values["ID"].(string)
}

// Erase Session
func Erase(w http.ResponseWriter, r *http.Request) {
	var s = GetSession(r)
	s.Options = &sessions.Options{MaxAge: -1}
	s.Save(r, w)
}

// Exist Session
func Exist(w http.ResponseWriter, r *http.Request) bool {
	var s = GetSession(r)
	return s.Values["ID"] != nil
}

func init() {
	store.Options = &sessions.Options{Path: "/", MaxAge: 60 * 60 * 24 * 7}
}
