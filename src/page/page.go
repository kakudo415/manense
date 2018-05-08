package page

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"../orm"
	"../session"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// HomeView data
type HomeView struct {
	Common struct {
		Name string
		URL  string
		AURL string // OAuth URL
	}
	User     orm.Users
	Expenses []orm.Expenses
}

// OtherView data
type OtherView struct {
	Common struct {
		Name string
		URL  string
		AURL string
	}
	User     orm.Users
	Other    orm.Users
	Expenses []orm.Expenses
}

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("view/home.html", "view/common.html"))
	var v = new(HomeView)
	v.Common.Name = "manense"
	v.Common.URL = os.Getenv("MANENSE_URL")
	v.Common.AURL = oauth.AuthCodeURL(oauth.RedirectURL)
	if session.Exist(w, r) {
		var userID = session.Get(w, r)
		v.User = orm.GetUser(userID)
		v.Expenses = orm.GetExpenseList(userID)
		sort.Slice(v.Expenses, func(i, j int) bool { return v.Expenses[i].Time.After(v.Expenses[j].Time) })
	}
	t.Execute(w, v)
}

// Other Expense page
func Other(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("view/other.html", "view/common.html"))
	var v = new(OtherView)
	v.Common.Name = "manense"
	v.Common.URL = os.Getenv("MANENSE_URL")
	v.Common.AURL = oauth.AuthCodeURL(oauth.RedirectURL)
	if session.Exist(w, r) {
		var userID = session.Get(w, r)
		v.User = orm.GetUser(userID)
		v.Other = orm.GetUser(strings.TrimPrefix(r.URL.Path, "/other/"))
		v.Expenses = orm.GetExpenseList(v.Other.ID)
	}
	t.Execute(w, v)
}

// New Expense
func New(w http.ResponseWriter, r *http.Request) {
	if session.Exist(w, r) && r.Method == "POST" {
		r.ParseForm()
		var i, e = strconv.ParseInt(r.Form["expense-income"][0], 10, 64)
		if e != nil {
			w.WriteHeader(400)
			w.Write([]byte(""))
			return
		}
		var ne = orm.NewExpense(session.Get(w, r), r.Form["expense-name"][0], i, r.Form["expense-time"][0])
		var u = orm.GetUser(session.Get(w, r))
		u.Balance += i
		u.Update()
		print(ne.UUID)
		println(u.Balance)
		w.Write([]byte(fmt.Sprintf("{ \"uuid\": \"%d\", \"balance\": \"%d\" }", ne.UUID, u.Balance)))
	}
}

// Update Expense
func Update(w http.ResponseWriter, r *http.Request) {
	if session.Exist(w, r) && r.Method == "POST" {
		r.ParseForm()
		var uuid, e = strconv.ParseUint(r.Form["expense-uuid"][0], 10, 64)
		if e != nil {
			w.WriteHeader(400)
			w.Write([]byte(""))
			return
		}
		var ex = orm.GetExpense(uuid)
		ex.Name = r.Form["expense-name"][0]
		ex.Income, e = strconv.ParseInt(r.Form["expense-income"][0], 10, 64)
		if e != nil {
			w.WriteHeader(400)
			w.Write([]byte(""))
			return
		}
		ex.Time, e = time.Parse("2006-01-02", r.Form["expense-time"][0])
		if e != nil {
			w.WriteHeader(400)
			w.Write([]byte(""))
			return
		}
		orm.UpdateExpense(ex)
	}
}

// Erase Expense
func Erase(w http.ResponseWriter, r *http.Request) {
	if session.Exist(w, r) && r.Method == "POST" {
		r.ParseForm()
		var i, e = strconv.ParseUint(r.Form["expense-uuid"][0], 10, 64)
		if e != nil {
			w.WriteHeader(400)
			w.Write([]byte(""))
			return
		}
		var eraseExpense = orm.GetExpense(i)
		orm.EraseExpense(i)
		var u = orm.GetUser(session.Get(w, r))
		u.Balance -= eraseExpense.Income
		u.Update()
		w.Write([]byte(strconv.FormatInt(u.Balance, 10)))
	}
}

// Info Expense
func Info(w http.ResponseWriter, r *http.Request) {
	if session.Exist(w, r) && r.Method == "POST" {
		r.ParseForm()
		var i, e = strconv.ParseUint(r.Form["expense-uuid"][0], 10, 64)
		if e != nil {
			w.WriteHeader(400)
			w.Write([]byte(""))
			return
		}
		var b, _ = json.Marshal(orm.GetExpense(i))
		w.Write(b)
	}
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
		u.Update()
		orm.BalanceInquiry(u.ID)
		session.New(w, r, u.ID)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// Signout func
func Signout(w http.ResponseWriter, r *http.Request) {
	session.Erase(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
