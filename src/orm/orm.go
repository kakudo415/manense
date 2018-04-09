package orm

import (
	"os"

	"github.com/jinzhu/gorm"
	// MySQL Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Users struct
type Users struct {
	ID   string `gorm:"primary_key"`
	Name string `gorm:"not null"`
}

// Books struct
type Books struct {
	ID      uint64 `sql:"type:BIGINT UNSIGNED" gorm:"PRIMARY_KEY"`
	Name    string `gorm:"not null"`
	UserID  string `gorm:"not null"`
	Balance int    `gorm:"not null"`
}

// Expenses struct
type Expenses struct {
	ID     uint64 `sql:"type:BIGINT UNSIGNED" gorm:"PRIMARY_KEY"`
	Name   string `gorm:"not null"`
	BookID uint64 `gorm:"not null"`
	Income int    `gorm:"not null"`
}

func init() {
	Connect().AutoMigrate(new(Users))
	Connect().AutoMigrate(new(Books)).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	Connect().AutoMigrate(new(Expenses)).AddForeignKey("book_id", "books(id)", "CASCADE", "CASCADE")
}

// Connect gorm
func Connect() *gorm.DB {
	var c, e = gorm.Open("mysql", os.Getenv("MANENSE_DATABASE"))
	if e != nil {
		panic(e)
	}
	return c
}

// UUID func
func UUID() (u uint64) {
	Connect().Unscoped().Raw("SELECT UUID_SHORT()").Row().Scan(&u)
	return u
}
