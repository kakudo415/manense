package page

import (
	"fmt"
	"net/http"
)

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HOME")
}
