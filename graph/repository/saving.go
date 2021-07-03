package repository

import "github.com/hoashi-akane/moneygement-graphql/graph/model"

func (SavDB *SavDB) GetSavingDetail(userId int) *model.Savings {
	var result model.Savings
	SavDB.DB.Table("savings").Take(&result, "userid=?", userId)
	return &result
}

func (SavDB *SavDB) GetSavingAmount(userId int) *model.SavingAmountList {
	var result model.SavingAmountList
	SavDB.DB.Table("savings s").
		Select("COALESCE(SUM(sd.saving_amount),0) as saving_amount  , COALESCE(SUM(e.expense_amount), 0) as expense_amount").
		Joins("left outer join savings_details sd ON s.id = sd.saving_id").
		Joins("LEFT OUTER JOIN expenses e ON s.id = e.saving_id").
		Find(&result, "userid=?", userId)
	return &result
}

func (SavDB *SavDB) GetSavingDetails(filter *model.SavingsDetailsFilter) []*model.SavingsDetail {
	// 家計簿の利用履歴を取得
	var results []*model.SavingsDetail
	// offset, limit句の代わりにPKに対してBETWEEN句を利用しカラムを指定。速度が早いっぽい
	SavDB.DB.Table("savings_details sd").Order("saving_date DESC").
		Joins("RIGHT OUTER JOIN savings s ON sd.saving_id = s.id").
		Select("sd.id, sd.saving_id, sd.saving_amount, sd.note, sd.saving_date").
		Offset(filter.First).Limit(filter.Last).Find(&results, "s.userid=?", filter.SavingsID)
	return results
}

func (SavDB *SavDB) CreateSavingDetail(detail *model.NewSavingDetail) (int, error) {
	err := SavDB.DB.Table("savings_details").Create(&detail).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (SavDB *SavDB) CreateIncomeDetail(detail *model.NewIncomeDetail) (int, error) {
	err := SavDB.DB.Table("incomes").Create(&detail).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (SavDB *SavDB) CreateExpenseDetail(detail *model.NewExpenseDetail) (int, error) {
	err := SavDB.DB.Table("expenses").Create(&detail).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (SavDB *SavDB) UpdateTargetAmount(data *model.UpdateUser) (int, error) {
	err := SavDB.DB.Table("savings").Where("userid=?", data.ID).
		Update("target_amount", data.TargetAmount).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}

type SavDBInterface interface {
	GetSavingDetail(userId int) *model.Savings
	GetSavingAmount(userId int) *model.SavingAmountList
	GetSavingDetails(filter *model.SavingsDetailsFilter) []*model.SavingsDetail
	CreateSavingDetail(detail *model.NewSavingDetail) (int, error)
	CreateIncomeDetail(detail *model.NewIncomeDetail) (int, error)
	CreateExpenseDetail(detail *model.NewExpenseDetail) (int, error)
	UpdateTargetAmount(data *model.UpdateUser) (int, error)
}
