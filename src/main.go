package main

import (
	"net/http"
	"os"

	"./page"
	"./session"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("view/static/"))))
	http.HandleFunc("/", page.Index)
	http.HandleFunc("/signin", session.Signin)
	http.HandleFunc("/signout", session.Signout)
	http.ListenAndServe(":"+os.Getenv("MANENSE_PORT"), nil)
}
