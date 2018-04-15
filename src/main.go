package main

import (
	"fmt"
	"net/http"
	"os"

	"./page"
)

func main() {
	go REPL()
	http.HandleFunc("/", page.Home)
	http.HandleFunc("/signin", page.Signin)
	http.HandleFunc("/signout", page.Signout)
	http.ListenAndServe(os.Getenv("MANENSE_ADDRESS"), nil)
}

// REPL func
func REPL() {
	var c string
	sucMessage("SERVER START")
	for {
		fmt.Print(">> ")
		fmt.Scan(&c)
		switch c {
		case "EXIT":
			os.Exit(0)
		default:
			errMessage("UNKNOWN COMMAND")
		}
	}
}

func sucMessage(m string) {
	fmt.Fprintf(os.Stdout, "\033[32m%s\033[0m\n", m)
}

func errMessage(m string) {
	fmt.Fprintf(os.Stderr, "\033[31m%s\033[0m\n", m)
}
