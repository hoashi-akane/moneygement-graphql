package repository

import (
	"github.com/hoashi-akane/moneygement-graphql/graph/model"
	"github.com/jinzhu/gorm"
)

func (LedgerDB *LedgerDB) GetLedgerList(userId int, obj *model.LedgerEtc) ([]*model.Ledger, error){
	LedgerDB.DB.Table("ledger").Find(&obj.Ledgers, "id=?", userId)
	return obj.Ledgers, nil
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

type LedgerDBInterface interface {
	GetLedgerList(userId int, obj *model.LedgerEtc) ([]*model.Ledger, error)
	GetLedger(id int) *model.Ledger
	GetCategory(id int) *model.Category
	GetIncomes(ledgerId int) []*model.Income
	GetExpenses(ledgerId int) []*model.Expense
}

type LedgerDB struct {
	DB *gorm.DB
}
