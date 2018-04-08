package orm

// NewUser func
func NewUser(id string, name string) (u *Users) {
	u = new(Users)
	u.ID = id
	u.Name = name
	Connect().FirstOrCreate(u)
	return u
}

// GetUser func
func GetUser(id string) (u *Users) {
	u = new(Users)
	u.ID = id
	Connect().First(u)
	return u
}

// UpdateUser func
func UpdateUser(id string, name string) (u *Users) {
	u = new(Users)
	u.ID = id
	u.Name = name
	Connect().Model(&Users{ID: id}).Update(u)
	return u
}
