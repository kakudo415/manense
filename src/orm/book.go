package orm

// NewBook func
func NewBook(name string, ui string) (b *Books) {
	b = new(Books)
	b.ID = UUID()
	b.Name = name
	b.UserID = ui
	b.Balance = 0
	Connect().FirstOrCreate(b)
	return b
}

// GetBook func
func GetBook(id uint64) (b *Books) {
	b = new(Books)
	b.ID = id
	Connect().First(b)
	return b
}

// GetBooks func
func GetBooks(ui string) (bs *[]Books) {
	bs = new([]Books)
	Connect().Find(bs, "user_id=?", ui)
	return bs
}

// UpdateBook func
func UpdateBook(id uint64, name string, balance int) (b *Books) {
	b = new(Books)
	b.ID = id
	b.Name = name
	b.Balance = balance
	Connect().Model(&Books{ID: id}).Update(b)
	return b
}
