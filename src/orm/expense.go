package orm

// NewExpense func
func NewExpense(name string, bi uint64, in int) (e *Expenses) {
	e = new(Expenses)
	e.Name = name
	e.BookID = bi
	e.Income = in
	Connect().FirstOrCreate(e)
	return e
}

// GetExpense func
func GetExpense(id uint64) (e *Expenses) {
	e = new(Expenses)
	e.ID = id
	Connect().First(e)
	return e
}

// GetExpenses func
func GetExpenses(bi uint64) (es *[]Expenses) {
	es = new([]Expenses)
	Connect().Find(es, "book_id=?", bi)
	return es
}

// UpdateExpense func
func UpdateExpense(id uint64, name string, in int) (e *Expenses) {
	e = new(Expenses)
	e.ID = id
	e.Name = name
	e.Income = in
	Connect().Model(&Expenses{ID: id}).Update(e)
	return e
}
