package orm

import (
	"time"
)

// NewExpense func
func NewExpense(userID string, name string, income int64) (e Expenses) {
	e.UUID = UUID()
	e.Name = name
	e.Income = income
	e.UserID = userID
	e.Time = time.Now()
	Connect().FirstOrCreate(&e).Close()
	return e
}

// GetExpense func
func GetExpense(UUID uint64) (e Expenses) {
	e.UUID = UUID
	Connect().First(&e).Close()
	return e
}

// UpdateExpense func
func UpdateExpense(ex Expenses) {
	Connect().Model(&Expenses{UUID: ex.UUID}).Update(&ex).Close()
}

// EraseExpense func
func EraseExpense(exID uint64) {
	var e = new(Expenses)
	e.UUID = exID
	Connect().Delete(&e).Close()
}

// GetExpenseList func
func GetExpenseList(ui string) (es []Expenses) {
	Connect().Find(&es, "user_id = ?", ui).Close()
	return es
}
