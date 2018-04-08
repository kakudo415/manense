package session

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/gorilla/sessions"

	"../database"
)

var store = sessions.NewCookieStore([]byte("SECRET"))
var name = "SESSION_NAME"
var config = oauth2.Config{
	Endpoint:     google.Endpoint,
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
}

type oauthJSON struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	Locale     string `json:"locale"`
	Err        string `json:"error"`
}

func init() {
	store.Options = &sessions.Options{MaxAge: 60 * 60 * 24 * 7, Path: "/"}
}

// Get Session
func Get(w http.ResponseWriter, r *http.Request) *sessions.Session {
	var s, e = store.Get(r, name)
	if e != nil {
		s, e = store.New(r, name)
	}
	return s
}

// Exist session
func Exist(w http.ResponseWriter, r *http.Request) (u *database.Users) {
	var s = Get(w, r)
	if s.Values["ID"] == nil {
		return nil
	}
	return database.User(s.Values["ID"].(string), "")
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

// Signin OAuth callback
func Signin(w http.ResponseWriter, r *http.Request) {
	if Exist(w, r) != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		t, e := config.Exchange(oauth2.NoContext, r.URL.Query().Get("code"))
		if e != nil {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			c := config.Client(oauth2.NoContext, t)
			i, e := c.Get("https://www.googleapis.com/oauth2/v1/userinfo")
			b, e := ioutil.ReadAll(i.Body)
			var j oauthJSON
			if e = json.Unmarshal(b, &j); e != nil {
				http.Redirect(w, r, "/", http.StatusFound)
			}
			Create(w, r, j.ID)
			database.User(j.ID, "")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

// OAuthURL func
func OAuthURL() string {
	return config.AuthCodeURL(os.Getenv("GOOGLE_REDIRECT_URL"))
}

// Signout erase session
func Signout(w http.ResponseWriter, r *http.Request) {
	Erase(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
