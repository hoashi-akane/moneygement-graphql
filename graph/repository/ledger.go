package repository

import (
	"github.com/hoashi-akane/moneygement-graphql/graph/model"
	"github.com/jinzhu/gorm"
)

func (ledgerDB *LedgerDB) GetLedgerList(userId int, obj *model.LedgerEtc) ([]*model.Ledger, error){
	ledgerDB.DB.Table("ledger").Find(&obj.Ledgers, "id=?", userId)
	return obj.Ledgers, nil
}

func (ledgerDB *LedgerDB) GetLedger(id int) *model.Ledger{
	 var ledger model.Ledger
	 ledgerDB.DB.Table("ledger").Take(&ledger, "id=?", id)
	 return &ledger
}

type LedgerDBInterface interface {
	GetLedgerList(userId int, obj *model.LedgerEtc) ([]*model.Ledger, error)
	GetLedger(id int) *model.Ledger
}

type LedgerDB struct {
	DB *gorm.DB
}
