main: package
	go build -o bin/manense src/main.go
package:
	go get -u github.com/jinzhu/gorm
	go get -u github.com/gorilla/sessions
	go get -u golang.org/x/oauth2
	go get -u golang.org/x/oauth2/google