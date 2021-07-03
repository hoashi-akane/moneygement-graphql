package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/hoashi-akane/moneygement-graphql/graph/repository"
	"log"

	"github.com/hoashi-akane/moneygement-graphql/graph/generated"
	"github.com/hoashi-akane/moneygement-graphql/graph/model"
)

func (r *expenseResolver) Category(ctx context.Context, obj *model.Expense) (*model.Category, error) {
	return repository.NewLedgerDB().GetCategory(obj.CategoryID), nil
}

func (r *incomeResolver) Category(ctx context.Context, obj *model.Income) (*model.Category, error) {
	return repository.NewLedgerDB().GetCategory(obj.CategoryID), nil
}

func (r *ledgerResolver) Incomes(ctx context.Context, obj *model.Ledger) ([]*model.Income, error) {
	return repository.NewLedgerDB().GetIncomes(obj.ID), nil
}

func (r *ledgerResolver) Expenses(ctx context.Context, obj *model.Ledger) ([]*model.Expense, error) {
	return repository.NewLedgerDB().GetExpenses(obj.ID), nil
}

func (r *ledgerEtcResolver) Ledgers(ctx context.Context, obj *model.LedgerEtc, userID int) ([]*model.Ledger, error) {
	return repository.NewLedgerDB().GetLedgerList(userID), nil
}

func (r *ledgerEtcResolver) Ledger(ctx context.Context, obj *model.LedgerEtc, id int) (*model.Ledger, error) {
	return repository.NewLedgerDB().GetLedger(id), nil
}

func (r *ledgerEtcResolver) ShareLedgers(ctx context.Context, obj *model.LedgerEtc, userID int) ([]*model.Ledger, error) {
	db := repository.NewLedgerDB()
	enrollments := db.GetEnrollments(userID)
	var groupSlice []int
	for _, value := range enrollments {
		groupSlice = append(groupSlice, value.GroupID)
	}
	return db.GetShareLedgerList(groupSlice), nil
}

func (r *ledgerEtcResolver) AdviserLedgers(ctx context.Context, obj *model.LedgerEtc, adviserID int) ([]*model.Ledger, error) {
	return repository.NewLedgerDB().GetAdviserLedgers(adviserID), nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	var user = model.User{Nickname: input.NickName, Name: input.Name, Email: input.Email}
	newUser := repository.NewUserDB().CreateUser(&user, input.Password)
	if newUser == nil {
		return nil, nil
	}
	return newUser, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error) {
	var user model.User
	if input == nil || input.ID == 0 {
		return nil, nil
	}
	// ユーザテーブル更新
	db := repository.NewUserDB()
	_, err := db.UpdateUser(input)
	if err != nil {
		return nil, err
	}
	// 貯金テーブル更新 ( 目標貯金額
	_, err = repository.NewSavDB().UpdateTargetAmount(input)
	if err != nil {
		return nil, err
	}
	// ユーザ情報取得
	db.GetIdUser(input.ID)
	return &user, nil
}

func (r *mutationResolver) CreateAdviser(ctx context.Context, input *model.NewAdviser) (*int, error) {
	if input.ID == 0 {
		return nil, nil
	}
	result, err := repository.NewUserDB().CreateAdviser(input)
	if err != nil {
		return nil, nil
	}
	return &result, nil
}

func (r *mutationResolver) AddGroupUser(ctx context.Context, input *model.NewGroupUser) (*int, error) {
	result, err := repository.NewUserDB().InsertGroupUser(input)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *mutationResolver) AddUseAdviser(ctx context.Context, input *model.AddAdviser) (*int, error) {
	if input == nil || input.AdviserID == 0 || input.LedgerID == 0 || input.UserID == 0 {
		return nil, nil
	}
	rows, err := repository.NewLedgerDB().AddUseAdviser(input)
	if err != nil {
		return nil, err
	}
	return &rows, nil
}

func (r *mutationResolver) CreateGroup(ctx context.Context, input *model.NewGroup) (*int, error) {
	// グループ作成
	userDb := repository.NewUserDB()
	var group = model.Group{Author: input.UserID, Name: input.GroupName}
	_, err := userDb.CreateGroup(&group)
	if err != nil {
		fmt.Println("エラー")
		return nil, err
	}
	var enrollment = model.Enrollment{UserID: group.Author, GroupID: group.ID}
	var ledger = model.Ledger{GroupID: group.ID, Name: input.LedgerName, UserID: input.UserID}
	_, err = repository.NewLedgerDB().CreateGroupLedger(&ledger)
	if err != nil {
		return nil, err
	}
	_, err = userDb.CreateEnrollment(&enrollment)
	if err != nil {
		fmt.Println("エラー")
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) CreateChat(ctx context.Context, input *model.NewChat) (*int, error) {
	result, err := repository.NewUserDB().CreateChat(input)
	if err != nil {
		fmt.Println("エラー")
		return nil, err
	}
	return &result, nil
}

func (r *mutationResolver) CreateSavingDetail(ctx context.Context, input *model.NewSavingDetail) (*int, error) {
	// 登録処理
	result, err := repository.NewSavDB().CreateSavingDetail(input)

	if err != nil {
		panic(fmt.Errorf("構文エラーもしくは制約に引っかかっている"))
		return nil, err
	}

	return &result, nil
}

func (r *mutationResolver) CreateIncomeDetail(ctx context.Context, input *model.NewIncomeDetail) (*int, error) {
	result, err := repository.NewSavDB().CreateIncomeDetail(input)

	if err != nil {
		panic(fmt.Errorf("構文エラーもしくは制約に引っかかっている"))
		return nil, err
	}

	return &result, nil
}

func (r *mutationResolver) CreateExpenseDetail(ctx context.Context, input *model.NewExpenseDetail) (*int, error) {
	result, err := repository.NewSavDB().CreateExpenseDetail(input)

	if err != nil {
		panic(fmt.Errorf("構文エラーもしくは制約に引っかかっている"))
		return nil, err
	}

	return &result, nil
}

func (r *mutationResolver) CreateLedger(ctx context.Context, input *model.NewLedger) (*int, error) {
	rows, err := repository.NewLedgerDB().CreateLedger(input)
	if err != nil {
		panic(fmt.Errorf("構文エラーもしくは制約に引っかかっている"))
	}
	return &rows, nil
}

func (r *mutationResolver) DeleteLedger(ctx context.Context, input *model.DeleteLedger) (*int, error) {
	if input == nil || input.ID == 0 {
		return nil, nil
	}
	rows, err := repository.NewLedgerDB().DeleteLedger(input.ID)
	if err != nil {
		panic(fmt.Errorf("構文エラーもしくは制約に引っかかっている"))
	}
	return &rows, nil
}

func (r *mutationResolver) DeleteGroup(ctx context.Context, groupID int) (*int, error) {
	//	グループ削除、　一緒に家計簿も削除する必要がある
	//	0またはnilだと全カラム削除されてしまうので注意
	if groupID == 0 {
		return nil, nil
	}
	// グループ削除
	_, err := repository.NewUserDB().DeleteGroup(groupID)
	if err != nil {
		log.Fatal("エラー")
		return nil, err
	}

	// 家計簿削除
	result, err := repository.NewLedgerDB().DeleteGroupLedger(groupID)
	if err != nil {
		log.Fatal("エラー")
		return nil, err
	}
	return &result, nil
}

func (r *queryResolver) Login(ctx context.Context, input model.LoginInfo) (*model.User, error) {
	user, _ := repository.NewUserDB().Login(&input)
	return user, nil
}

func (r *queryResolver) AdviserList(ctx context.Context, input model.AdviserListFilter) ([]*model.User, error) {
	userList, _ := repository.NewUserDB().GetAdviserList(&input)
	return userList, nil
}

func (r *queryResolver) UseAdviserMemberList(ctx context.Context, input model.UseAdviserMemberFilter) ([]*model.AdviserMember, error) {
	adviserList, _ := repository.NewUserDB().GetAdviserMemberList(&input)
	return adviserList, nil
}

func (r *queryResolver) ChatList(ctx context.Context, input model.ChatFilter) ([]*model.Chat, error) {
	chatList, _ := repository.NewUserDB().GetChatList(&input)
	return chatList, nil
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
	result := repository.NewSavDB().GetSavingDetail(userID)
	return result, nil
}

func (r *savingResolver) SavingAmount(ctx context.Context, obj *model.Saving, userID int) (*model.SavingAmountList, error) {
	result := repository.NewSavDB().GetSavingAmount(userID)
	return result, nil
}

func (r *savingResolver) SavingsDetails(ctx context.Context, obj *model.Saving, input model.SavingsDetailsFilter) ([]*model.SavingsDetail, error) {
	// 家計簿の利用履歴を取得
	results := repository.NewSavDB().GetSavingDetails(&input)
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
