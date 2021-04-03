package repository

import (
	"github.com/hoashi-akane/moneygement-graphql/graph/model"
	"github.com/jinzhu/gorm"
)

func GetLedgerList(db *gorm.DB,userId int, obj *model.LedgerEtc) ([]*model.Ledger, error){
	db.Table("ledger").Find(&obj.Ledgers, "id=?", userId)
	return obj.Ledgers, nil
}