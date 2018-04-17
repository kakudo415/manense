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
	var c uint
	Connect().Model(new(Expenses)).Select("COUNT(*)").Where("user_id = ?", u.ID).Row().Scan(&c)
	if c == 0 {
		u.Balance = 0
	} else {
		Connect().Model(new(Expenses)).Select("SUM(income)").Where("user_id = ?", u.ID).Row().Scan(&u.Balance)
	}
	println(u.ID)
	println(u.Balance)
	Connect().Model(&Users{ID: i}).Update(u).Close()
}

// GetUser func
func GetUser(i string) (u Users) {
	u.ID = i
	Connect().First(&u).Close()
	return u
}
