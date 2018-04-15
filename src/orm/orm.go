package orm

import (
	"os"

	"github.com/jinzhu/gorm"
)

// Users M
type Users struct {
	ID   string
	Name string
}

// Connect GORM
func Connect() *gorm.DB {
	var c, e = gorm.Open("mysql", os.Getenv("MANENSE_DATABASE"))
	if e != nil {
		panic(e)
	}
	return c
}
