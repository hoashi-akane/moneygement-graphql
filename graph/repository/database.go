package repository

import "github.com/jinzhu/gorm"

var savDB *gorm.DB
var usrDB *gorm.DB
var baseDB *gorm.DB

func DbStart() {
	var err error
	savDB, err = gorm.Open("mysql", dataSourceSavings)
	usrDB, err = gorm.Open("mysql", dataSourceUser)
	baseDB, err = gorm.Open("mysql", dataSourceLedger)

	if err != nil{
		panic(err)
	}
	if savDB == nil || usrDB == nil || baseDB == nil{
		panic(err)
	}
	savDB.LogMode(true)
	usrDB.LogMode(true)
	baseDB.LogMode(true)
}

func DbClose() {
	if savDB != nil{
		if err := savDB.Close(); err != nil{
			panic(err)
		}
	}
	if usrDB != nil{
		if err := usrDB.Close(); err != nil{
			panic(err)
		}
	}
	if baseDB != nil{
		if err := baseDB.Close(); err != nil{
			panic(err)
		}
	}
}

type LedgerDB struct {
	DB *gorm.DB
}

func NewLedgerDB() *LedgerDB {
	return &LedgerDB{DB: baseDB}
}

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB() *UserDB {
	return &UserDB{DB: usrDB}
}

type SavDB struct {
	DB *gorm.DB
}

func NewSavDB() *SavDB {
	return &SavDB{DB: baseDB}
}
