package orm

// New User
func (u Users) New() {
	Connect().FirstOrCreate(&u).Close()
}

// GetUser func
func GetUser(i string) (u Users) {
	u.ID = i
	Connect().First(&u).Close()
	return u
}
