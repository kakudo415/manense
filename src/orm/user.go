package orm

// NewUser func
func NewUser(i string, n string) {
	var u = Users{ID: i, Name: n}
	Connect().FirstOrCreate(&u)
}
