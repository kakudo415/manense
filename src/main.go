package main

import (
	"net/http"
	"os"

	"./page"
)

func main() {
	http.HandleFunc("/", page.Home)
	http.HandleFunc("/signin", page.Signin)
	http.HandleFunc("/signout", page.Signout)
	http.Handle("/common/", http.StripPrefix("/common/", http.FileServer(http.Dir("view/common"))))
	http.ListenAndServe(os.Getenv("MANENSE_ADDRESS"), nil)
}
