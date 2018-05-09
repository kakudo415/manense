package orm

// New User
func (u Users) New() {
	Connect().FirstOrCreate(&u).Close()
}

// Update User
func (u Users) Update() {
	Connect().Model(&Users{ID: u.ID}).Update(&u).Close()
}

// BalanceInquiry func
func BalanceInquiry(i string) {
	var u = GetUser(i)
	var b int
	Connect().Model(new(Expenses)).Select("SUM(income)").Where("user_id = ?", u.ID).Row().Scan(&b)
	u.Balance = int64(b)
	Connect().Model(&Users{ID: i}).Update(&u).Close()
}

// GetUser func
func GetUser(i string) (u Users) {
	u.ID = i
	Connect().First(&u).Close()
	return u
}
