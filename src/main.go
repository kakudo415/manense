package main

import (
	"net/http"
	"os"

	"./page"
)

func main() {
	http.HandleFunc("/", page.Home)
	http.HandleFunc("/new", page.New)
	http.HandleFunc("/update", page.Update)
	http.HandleFunc("/erase", page.Erase)
	http.HandleFunc("/info", page.Info)
	http.HandleFunc("/o/", page.Other)
	http.HandleFunc("/signin", page.Signin)
	http.HandleFunc("/signout", page.Signout)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
