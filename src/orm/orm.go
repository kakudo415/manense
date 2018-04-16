package orm

import (
	"os"

	"github.com/jinzhu/gorm"
	// MySQL
	_ "github.com/go-sql-driver/mysql"
)

// Users M
type Users struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	GName  string `json:"given_name"`
	FName  string `json:"family_name"`
	Icon   string `json:"picture"`
	Locale string `json:"locale"`
}

func init() {
	Connect().AutoMigrate(new(Users))
}

// Connect GORM
func Connect() *gorm.DB {
	var c, e = gorm.Open("mysql", os.Getenv("MANENSE_DATABASE"))
	if e != nil {
		panic(e)
	}
	return c
}
