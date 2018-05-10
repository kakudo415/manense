package orm

// New User
func (u Users) New() {
	Connect().FirstOrCreate(&u).Close()
}

// Update User
func (u Users) Update() {
	Connect().Model(&Users{ID: u.ID}).Update(&u).Close()
}

// Balance func
func Balance(i string) (b int64) {
	Connect().Model(new(Expenses)).Select("SUM(income)").Where("user_id = ?", i).Row().Scan(&b)
	return b
}

// GetUser func
func GetUser(i string) (u Users) {
	u.ID = i
	Connect().First(&u).Close()
	return u
}
