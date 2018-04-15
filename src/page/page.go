package page

import (
	"html/template"
	"net/http"

	"../session"
)

// HomeView data
type HomeView struct {
	Common struct {
		Name string
	}
}

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("view/home.html"))
	var v = new(HomeView)
	v.Common.Name = "manense"
	t.Execute(w, v)
}

// Signin func
func Signin(w http.ResponseWriter, r *http.Request) {

}

// Signout func
func Signout(w http.ResponseWriter, r *http.Request) {
	session.Erase(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
