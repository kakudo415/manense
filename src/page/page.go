package page

import (
	"html/template"
	"net/http"

	"../orm"
)

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("view/home.html", "view/common/common.html"))
	var v = view(w, r)
	t.Execute(w, v)
}

// Signin func
func Signin(w http.ResponseWriter, r *http.Request) {

}

// Signout func
func Signout(w http.ResponseWriter, r *http.Request) {

}

type viewData struct {
	Common map[string]string
	Books  orm.Books
}

func view(w http.ResponseWriter, r *http.Request) (v viewData) {
	v.Common = make(map[string]string)
	v.Common["Name"] = "manense"
	v.Common["URL"] = "http://127.0.0.1:8000"
	return v
}
