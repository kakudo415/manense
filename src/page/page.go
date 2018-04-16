package page

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"../orm"
	"../session"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// HomeView data
type HomeView struct {
	Common struct {
		Name string
		AURL string // OAuth URL
	}
	User orm.Users
}

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("view/home.html", "view/common.html"))
	var v = new(HomeView)
	v.Common.Name = "manense"
	v.Common.AURL = oauth.AuthCodeURL(oauth.RedirectURL)
	if session.Exist(w, r) {
		v.User = orm.GetUser(session.Get(w, r))
	}
	t.Execute(w, v)
}

var oauth = oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Endpoint:     google.Endpoint,
	RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
}

// Signin func
func Signin(w http.ResponseWriter, r *http.Request) {
	var c = r.URL.Query().Get("code")
	if len(c) != 0 {
		var u orm.Users
		token, err := oauth.Exchange(oauth2.NoContext, c)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		client := oauth.Client(oauth2.NoContext, token)
		res, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
		value, err := ioutil.ReadAll(res.Body)
		err = json.Unmarshal(value, &u)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		u.New()
		session.New(w, r, u.ID)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// Signout func
func Signout(w http.ResponseWriter, r *http.Request) {
	session.Erase(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
