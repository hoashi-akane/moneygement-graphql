package repository

import (
	"github.com/hoashi-akane/moneygement-graphql/graph/model"
)

func (LedgerDB *LedgerDB) GetLedgerList(userId int) []*model.Ledger {
	var ledgers []*model.Ledger
	LedgerDB.DB.Table("ledger").Find(&ledgers, "id=?", userId)
	return ledgers
}

func (LedgerDB *LedgerDB) GetShareLedgerList(groupIds []int) []*model.Ledger {
	var ledgers []*model.Ledger
	LedgerDB.DB.Table("ledger").Find(&ledgers, "group_id IN (?)", groupIds)
	return ledgers
}

func (LedgerDB *LedgerDB) GetLedger(id int) *model.Ledger {
	var ledger model.Ledger
	LedgerDB.DB.Table("ledger").Take(&ledger, "id=?", id)
	return &ledger
}

func (LedgerDB *LedgerDB) GetCategory(id int) *model.Category {
	var category model.Category
	LedgerDB.DB.Table("category").Take(&category, "id=?", id)
	return &category
}

func (LedgerDB *LedgerDB) GetIncomes(ledgerId int) []*model.Income {
	var incomes []*model.Income
	LedgerDB.DB.Table("incomes").Find(&incomes, "ledger_id=?", ledgerId)
	return incomes
}

func (LedgerDB *LedgerDB) GetExpenses(ledgerId int) []*model.Expense {
	var expenses []*model.Expense
	LedgerDB.DB.Table("expenses").Find(&expenses, "ledger_id=?", ledgerId)
	return expenses
}

func (LedgerDB *LedgerDB) GetEnrollments(userID int) []*model.Enrollment {
	var enrollment []*model.Enrollment
	LedgerDB.DB.Table("enrollment").Find(&enrollment, "user_id=?", userID)
	return enrollment
}

type LedgerDBInterface interface {
	GetLedgerList(userId int) []*model.Ledger
	GetShareLedgerList(groupIds []int) []*model.Ledger
	GetLedger(id int) *model.Ledger
	GetCategory(id int) *model.Category
	GetIncomes(ledgerId int) []*model.Income
	GetExpenses(ledgerId int) []*model.Expense
	GetEnrollments(userID int) []*model.Enrollment
}
