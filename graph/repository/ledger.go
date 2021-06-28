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

func (LedgerDB *LedgerDB) GetAdviserLedgers(adviserId int) []*model.Ledger {
	var ledgers []*model.Ledger
	LedgerDB.DB.Table("ledger").Find(&ledgers, "adviser_id=?", adviserId)
	return ledgers
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

func (LedgerDB *LedgerDB) CreateLedger(newLedger *model.NewLedger) (int, error) {
	err := LedgerDB.DB.Table("ledger").Create(&newLedger).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (LedgerDB *LedgerDB) DeleteLedger(id int) (int, error) {
	var ledger = model.Ledger{ID: id}
	err := LedgerDB.DB.Table("ledger").Where("user_id=? AND group_id=0", id).Delete(&ledger).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}

type LedgerDBInterface interface {
	GetLedgerList(userId int) []*model.Ledger
	GetShareLedgerList(groupIds []int) []*model.Ledger
	GetLedger(id int) *model.Ledger
	GetAdviserLedgers(adviserId int) []*model.Ledger
	GetCategory(id int) *model.Category
	GetIncomes(ledgerId int) []*model.Income
	GetExpenses(ledgerId int) []*model.Expense
	GetEnrollments(userID int) []*model.Enrollment
	CreateLedger(newLedger *model.NewLedger) (int, error)
	DeleteLedger(id int) (int, error)
}
