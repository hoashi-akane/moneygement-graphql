package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hoashi-akane/moneygement-graphql/graph/generated"
	"github.com/hoashi-akane/moneygement-graphql/graph/model"
)

func (r *expenseResolver) Category(ctx context.Context, obj *model.Expense) (*model.Category, error) {
	var category model.Category
	r.BASEDB.Table("category").Take(&category, "id=?", obj.CategoryID)
	return &category, nil
}

func (r *incomeResolver) Category(ctx context.Context, obj *model.Income) (*model.Category, error) {
	var category model.Category
	r.BASEDB.Table("category").Take(&category, "id=?", obj.CategoryID)
	return &category, nil
}

func (r *ledgerResolver) Incomes(ctx context.Context, obj *model.Ledger) ([]*model.Income, error) {
	r.BASEDB.Table("incomes").Find(&obj.Incomes, "ledger_id=?", &obj.ID)
	return obj.Incomes, nil
}

func (r *ledgerResolver) Expenses(ctx context.Context, obj *model.Ledger) ([]*model.Expense, error) {
	r.BASEDB.Table("expenses").Find(&obj.Expenses, "ledger_id=?", &obj.ID)
	return obj.Expenses, nil
}

func (r *ledgerEtcResolver) Ledgers(ctx context.Context, obj *model.LedgerEtc, userID int) ([]*model.Ledger, error) {
	r.BASEDB.Table("ledger").Find(&obj.Ledgers, "user_id=?", userID)
	return obj.Ledgers, nil
}

func (r *ledgerEtcResolver) Ledger(ctx context.Context, obj *model.LedgerEtc, id int) (*model.Ledger, error) {
	var ledger model.Ledger
	r.BASEDB.Table("ledger").Take(&ledger, "id=?", id)
	return &ledger, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	var user = model.User{Nickname: input.NickName, Name: input.Name, Email: input.Email}

	err := r.USRDB.Table("users").Create(&user).Error
	if err != nil {
		panic("エラ-")
		return nil, nil
	}
	r.USRDB.Table("users").Find(&user, "email = ?", user.Email)

	var auth = model.UserAuth{UserID: user.ID, Password: input.Password}
	err = r.USRDB.Table("user_auth").Create(&auth).Error
	if err != nil {
		return nil, nil
	}
	return &user, nil
}

func (r *mutationResolver) CreateGroup(ctx context.Context, input *model.NewGroup) (*int, error) {
	// グループ作成
	var group = model.Group{Author: input.UserID, Name: input.GroupName}
	err := r.USRDB.Table("groups").Create(&group).Error
	if err != nil {
		return nil, nil
	}
	// グループに自身を追加
	var enrollment = model.Enrollment{UserID: group.Author, GroupID: group.ID}
	err = r.USRDB.Table("enrollment").Create(&enrollment).Error
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (r *mutationResolver) CreateSavingDetail(ctx context.Context, input *model.NewSavingDetail) (*int, error) {
	// 登録処理
	err := r.SAVDB.Table("savings_details").Create(&input).Error
	//
	if err != nil {
		panic(fmt.Errorf("構文エラーもしくは制約に引っかかっている"))
	}

	return nil, nil
}

func (r *mutationResolver) CreateIncomeDetail(ctx context.Context, input *model.NewIncomeDetail) (*int, error) {
	err := r.BASEDB.Table("incomes").Create(&input).Error
	if err != nil {
		panic(fmt.Errorf("構文エラーもしくは制約に引っかかっている"))
	}
	return nil, nil
}

func (r *mutationResolver) CreateExpenseDetail(ctx context.Context, input *model.NewExpenseDetail) (*int, error) {
	err := r.BASEDB.Table("expenses").Create(&input).Error
	if err != nil {
		panic(fmt.Errorf("構文エラーもしくは制約に引っかかっている"))
	}
	return nil, nil
}

func (r *mutationResolver) CreateLedger(ctx context.Context, input *model.NewLedger) (*int, error) {
	err := r.BASEDB.Table("ledger").Create(input).Error
	if err != nil {
		panic(fmt.Errorf("構文エラーもしくは制約に引っかかっている"))
	}
	return nil, nil
}

func (r *queryResolver) Login(ctx context.Context, input model.LoginInfo) (*model.User, error) {
	var user model.User
	r.USRDB.Table("users").Select("*").Joins("left join user_auth on users.id = user_auth.user_id").Find(&user, "email = ? and user_auth.password = ?", input.Email, input.Password)
	return &user, nil
}

func (r *queryResolver) Saving(ctx context.Context) (*model.Saving, error) {
	saving := &model.Saving{}
	return saving, nil
}

func (r *queryResolver) Ledger(ctx context.Context) (*model.LedgerEtc, error) {
	ledger := &model.LedgerEtc{}
	return ledger, nil
}

func (r *savingResolver) SavingDetail(ctx context.Context, obj *model.Saving, userID int) (*model.Savings, error) {
	var result model.Savings
	r.SAVDB.Table("savings").Take(&result, "userid=?", userID)
	return &result, nil
}

func (r *savingResolver) SavingAmount(ctx context.Context, obj *model.Saving, userID int) (*model.SavingAmountList, error) {
	var result model.SavingAmountList
	r.SAVDB.Table("savings s").Select("COALESCE(SUM(sd.saving_amount),0) as saving_amount  , COALESCE(SUM(e.expense_amount), 0) as expense_amount").Joins("left outer join savings_details sd ON s.id = sd.saving_id").Joins("LEFT OUTER JOIN expenses e ON s.id = e.saving_id").Find(&result, "userid=?", userID)
	return &result, nil
}

func (r *savingResolver) SavingsDetails(ctx context.Context, obj *model.Saving, input model.SavingsDetailsFilter) ([]*model.SavingsDetail, error) {
	// 家計簿の利用履歴を取得
	var results []*model.SavingsDetail
	// offset, limit句の代わりにPKに対してBETWEEN句を利用しカラムを指定。速度が早いっぽい
	r.SAVDB.Order("saving_date DESC").Find(&results, "saving_id=? AND id BETWEEN ? AND ?", input.SavingsID, input.First, input.Last)
	return results, nil
}

// Expense returns generated.ExpenseResolver implementation.
func (r *Resolver) Expense() generated.ExpenseResolver { return &expenseResolver{r} }

// Income returns generated.IncomeResolver implementation.
func (r *Resolver) Income() generated.IncomeResolver { return &incomeResolver{r} }

// Ledger returns generated.LedgerResolver implementation.
func (r *Resolver) Ledger() generated.LedgerResolver { return &ledgerResolver{r} }

// LedgerEtc returns generated.LedgerEtcResolver implementation.
func (r *Resolver) LedgerEtc() generated.LedgerEtcResolver { return &ledgerEtcResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Saving returns generated.SavingResolver implementation.
func (r *Resolver) Saving() generated.SavingResolver { return &savingResolver{r} }

type expenseResolver struct{ *Resolver }
type incomeResolver struct{ *Resolver }
type ledgerResolver struct{ *Resolver }
type ledgerEtcResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type savingResolver struct{ *Resolver }
