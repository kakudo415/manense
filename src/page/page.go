package page

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"

	"golang.org/x/oauth2"

	"../orm"
	"../session"
)

var config = oauth2.Config{
	Endpoint:     google.Endpoint,
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
}

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("view/home.html", "view/common/common.html"))
	var v = view(w, r)
	if session.Exist(w, r) {
		v.User = orm.GetUser(session.UserID(w, r))
	}
	t.Execute(w, v)
}

// Userinfo JSON
type Userinfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"picture"`
	Err  string `json:"error"`
}

// Signin func
func Signin(w http.ResponseWriter, r *http.Request) {
	t, e := config.Exchange(oauth2.NoContext, r.URL.Query().Get("code"))
	if e != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	c := config.Client(oauth2.NoContext, t)
	i, e := c.Get("https://www.googleapis.com/oauth2/v1/userinfo")
	b, e := ioutil.ReadAll(i.Body)
	var j = new(Userinfo)
	json.Unmarshal(b, j)
	orm.NewUser(j.ID, j.Name)
	orm.UpdateUser(j.ID, j.Name)
	session.Create(w, r, j.ID)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Signout func
func Signout(w http.ResponseWriter, r *http.Request) {
	session.Erase(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

type viewData struct {
	Common map[string]string
	User   *orm.Users
}

func view(w http.ResponseWriter, r *http.Request) (v viewData) {
	v.Common = make(map[string]string)
	v.Common["Name"] = "manense"
	v.Common["URL"] = "http://127.0.0.1:8000"
	v.Common["OAuthURL"] = config.AuthCodeURL(os.Getenv("GOOGLE_REDIRECT_URL"))
	return v
}
