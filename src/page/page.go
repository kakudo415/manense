package page

import (
	"net/http"
	"text/template"

	"../common"
	"../session"
)

const indexPath = "view/index.html"
const commonPath = "view/common.html"

// Index page
func Index(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles(indexPath, commonPath))
	var d = common.Common()
	if session.Exist(w, r) == nil {
		d["OAuthURL"] = session.OAuthURL()
	} else {
		d["UserName"] = session.Exist(w, r).Name
	}
	t.Execute(w, d)
}
