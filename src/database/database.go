package database

import (
	"os"

	"github.com/jinzhu/gorm"
	// MySQL
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Users table
type Users struct {
	ID   string `gorm:"primary_key"`
	Name string `gorm:"not null"`
}

func init() {
	Connect().AutoMigrate(&Users{})
}

// Connect ORM
func Connect() *gorm.DB {
	c, e := gorm.Open("mysql", os.Getenv("DATABASE_URL"))
	if e != nil {
		panic(e)
	}
	return c
}

// User func
func User(id string, name string) (u *Users) {
	u = &Users{ID: id}
	if len(name) == 0 {
		u.Name = "Anonymous"
		Connect().FirstOrCreate(u)
	} else {
		u.Name = name
		Connect().Model(&Users{ID: id}).Update(u)
	}
	return u
}
