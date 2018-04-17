package orm

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"
	// MySQL
	_ "github.com/go-sql-driver/mysql"
)

// Users M
type Users struct {
	ID     string `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	GName  string `json:"given_name"`
	FName  string `json:"family_name"`
	Icon   string `json:"picture"`
	Locale string `json:"locale"`

	Balance int64
}

// Expenses M
type Expenses struct {
	UUID uint64 `gorm:"primary_key"`
	Name string `gorm:"not null"`

	Income int64  `gorm:"not null"`
	UserID string `gorm:"not null"`

	Time time.Time `gorm:"not null"`
}

func init() {
	Connect().AutoMigrate(new(Users)).Close()
	Connect().AutoMigrate(new(Expenses)).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Close()
}

// Connect GORM
func Connect() *gorm.DB {
	var c, e = gorm.Open("mysql", os.Getenv("MANENSE_DATABASE"))
	if e != nil {
		panic(e)
	}
	return c
}

// UUID func
func UUID() (u uint64) {
	Connect().Raw("SELECT UUID_SHORT()").Row().Scan(&u)
	return u
}
