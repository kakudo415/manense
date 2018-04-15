package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	go REPL()
	http.ListenAndServe(os.Getenv("MANENSE_ADDRESS"), nil)
}

// REPL func
func REPL() {
	var c string
	for {
		fmt.Print(">> ")
		fmt.Scan(&c)
		switch c {
		case "exit":
			os.Exit(0)
		default:
			errMessage("UNKNOWN COMMAND")
		}
	}
}

func errMessage(m string) {
	fmt.Fprintf(os.Stderr, "\033[31m%s\033[0m\n", m)
}
