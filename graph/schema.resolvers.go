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
	db := repository.LedgerDB{DB: r.BASEDB}
	return db.GetCategory(obj.CategoryID), nil
}

func (r *incomeResolver) Category(ctx context.Context, obj *model.Income) (*model.Category, error) {
	db := repository.LedgerDB{DB: r.BASEDB}
	return db.GetCategory(obj.CategoryID), nil
}

func (r *ledgerResolver) Incomes(ctx context.Context, obj *model.Ledger) ([]*model.Income, error) {
	db := repository.LedgerDB{DB: r.BASEDB}
	return db.GetIncomes(obj.ID), nil
}

func (r *ledgerResolver) Expenses(ctx context.Context, obj *model.Ledger) ([]*model.Expense, error) {
	db := repository.LedgerDB{DB: r.BASEDB}
	return db.GetExpenses(obj.ID), nil
}

func (r *ledgerEtcResolver) Ledgers(ctx context.Context, obj *model.LedgerEtc, userID int) ([]*model.Ledger, error) {
	db := repository.LedgerDB{DB: r.BASEDB}
	obj.Ledgers, _ = db.GetLedgerList(userID, obj)
	return obj.Ledgers, nil
}

func (r *ledgerEtcResolver) Ledger(ctx context.Context, obj *model.LedgerEtc, id int) (*model.Ledger, error) {
	db := repository.LedgerDB{DB: r.BASEDB}
	return db.GetLedger(id), nil
}

func (r *ledgerEtcResolver) ShareLedgers(ctx context.Context, obj *model.LedgerEtc, userID int) ([]*model.Ledger, error) {
	var enrollment []*model.Enrollment
	r.USRDB.Table("enrollment").Find(&enrollment, "user_id=?", userID)
	var groupSlice []int
	for _, value := range enrollment {
		groupSlice = append(groupSlice, value.GroupID)
	}
	r.BASEDB.Table("ledger").Find(&obj.Ledgers, "group_id IN (?)", groupSlice)
	return obj.Ledgers, nil
}

func (r *ledgerEtcResolver) AdviserLedgers(ctx context.Context, obj *model.LedgerEtc, adviserID int) ([]*model.Ledger, error) {
	r.BASEDB.Table("ledger").Find(&obj.Ledgers, "adviser_id=?", adviserID)
	return obj.Ledgers, nil
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

func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error) {
	var user model.User
	if input == nil || input.ID == 0 {
		return nil, nil
	}
	// ユーザテーブル更新
	err := r.USRDB.Table("users").Where("id = ?", input.ID).Updates(map[string]interface{}{"name": input.Name, "nickname": input.Nickname, "email": input.Email, "introduction": input.Introduction, "adviser_name": input.AdviserName}).Error
	if err != nil {
		return nil, nil
	}
	// 貯金テーブル更新 ( 目標貯金額
	err = r.SAVDB.Table("savings").Where("userid=?", input.ID).Update("target_amount", input.TargetAmount).Error
	if err != nil {
		return nil, nil
	}
	// ユーザ情報取得
	r.USRDB.Table("users").Where("id=?", input.ID).Find(&user)
	return &user, nil
}

func (r *mutationResolver) CreateAdviser(ctx context.Context, input *model.NewAdviser) (*int, error) {
	if input.ID == 0 {
		return nil, nil
	}
	err := r.USRDB.Table("users").Where("id=?", input.ID).Updates(map[string]interface{}{"name": input.Name, "introduction": input.Introduction, "adviser_name": input.AdviserName, "is_adviser": true}).Error
	if err != nil {
		return nil, nil
	}
	var result = 1
	return &result, nil
}

func (r *mutationResolver) AddGroupUser(ctx context.Context, input *model.NewGroupUser) (*int, error) {
	var user model.User
	r.USRDB.Table("users").Select("id").Find(&user, "email=? AND nickname=?", input.Email, input.NickName)
	if user.ID == 0 {
		return nil, nil
	}
	var enrollment = model.Enrollment{UserID: user.ID, GroupID: input.GroupID}
	r.USRDB.Table("enrollment").Create(&enrollment)
	var i = 1
	return &i, nil
}

func (r *mutationResolver) AddUseAdviser(ctx context.Context, input *model.AddAdviser) (*int, error) {
	if input == nil || input.AdviserID == 0 || input.LedgerID == 0 || input.UserID == 0 {
		return nil, nil
	}
	err := r.BASEDB.Table("ledger").Where("id=? AND user_id=?", input.LedgerID, input.UserID).Updates(map[string]interface{}{"adviser_id": input.AdviserID}).Error
	if err != nil {
		return nil, nil
	}
	var i = 1
	return &i, nil
}

func (r *mutationResolver) CreateGroup(ctx context.Context, input *model.NewGroup) (*int, error) {
	// グループ作成
	var group = model.Group{Author: input.UserID, Name: input.GroupName}
	err := r.USRDB.Table("groups").Create(&group).Error
	if err != nil {
		return nil, nil
	}
	var enrollment = model.Enrollment{UserID: group.Author, GroupID: group.ID}
	var ledger = model.Ledger{GroupID: group.ID, Name: input.LedgerName, UserID: input.UserID}
	err = r.BASEDB.Table("ledger").Select("group_id", "name", "user_id").Create(&ledger).Error
	if err != nil {
		return nil, nil
	}
	err = r.USRDB.Table("enrollment").Create(&enrollment).Error
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (r *mutationResolver) CreateChat(ctx context.Context, input *model.NewChat) (*int, error) {
	var err = r.USRDB.Table("chat").Create(&input).Error
	if err != nil {
		fmt.Println("エラー")
		return nil, nil
	}
	var result = 1
	return &result, nil
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

func (r *mutationResolver) DeleteLedger(ctx context.Context, input *model.DeleteLedger) (*int, error) {
	if input == nil || input.ID == 0 {
		return nil, nil
	}
	var ledger = model.Ledger{ID: input.ID}
	err := r.BASEDB.Table("ledger").Where("user_id=? AND group_id=0", input.UserID).Delete(&ledger).Error
	if err != nil {
		return nil, nil
	}
	var i = 1
	return &i, nil
}

func (r *mutationResolver) DeleteGroup(ctx context.Context, groupID int) (*int, error) {
	//	グループ削除、　一緒に家計簿も削除する必要がある
	//	0またはnilだと全カラム削除されてしまうので注意
	if groupID == 0 {
		return nil, nil
	}
	var group = model.Group{ID: groupID}
	// グループ削除
	err := r.USRDB.Table("groups").Delete(&group).Error
	if err != nil {
		log.Fatal("エラー")
	}
	var ledger = model.Ledger{GroupID: groupID}

	// 家計簿削除
	err = r.BASEDB.Table("ledger").Where("group_id =?", groupID).Delete(ledger).Error
	if err != nil {
		log.Fatal("エラー")
	}
	return nil, nil
}

func (r *queryResolver) Login(ctx context.Context, input model.LoginInfo) (*model.User, error) {
	var user model.User
	r.USRDB.Table("users").Select("users.id, users.name, users.nickname, users.email, users.adviser_name, users.introduction, user_auth.token").Joins("left join user_auth on users.id = user_auth.user_id").Find(&user, "email = ? and user_auth.password = ?", input.Email, input.Password)

	if &user.ID != nil && user.ID != 0 && user.Token != input.Token {
		r.USRDB.Table("user_auth").Where("user_id=? AND password=?", user.ID, input.Password).Update("token", input.Token)
	}
	return &user, nil
}

func (r *queryResolver) AdviserList(ctx context.Context, input model.AdviserListFilter) ([]*model.User, error) {
	var userList []*model.User
	r.USRDB.Table("users").Select("id, adviser_name, introduction").Offset(input.First).Limit(input.Last).Find(&userList, "is_adviser=true")
	return userList, nil
}

func (r *queryResolver) UseAdviserMemberList(ctx context.Context, input model.UseAdviserMemberFilter) ([]*model.AdviserMember, error) {
	var adviserList []*model.AdviserMember
	var userList []*model.User
	var ledgerList []*model.Ledger
	var userIdList []int
	r.BASEDB.Table("ledger").Select("id, user_id, name").Find(&ledgerList, "adviser_id = ?", input.UserID)
	for _, adviser := range ledgerList {
		if userIdList == nil {
			userIdList = append(userIdList, adviser.UserID)
		} else {
			for _, v := range userIdList {
				if adviser.UserID != v {
					userIdList = append(userIdList, v)
				}
			}
		}
	}

	r.USRDB.Table("users").Find(&userList, "id IN (?)", userIdList)
	for _, ledger := range ledgerList {
		for _, user := range userList {
			if ledger.UserID == user.ID {
				var adviser = model.AdviserMember{ID: ledger.ID, UserID: user.ID, NickName: user.Nickname, LedgerName: ledger.Name}
				adviserList = append(adviserList, &adviser)
			}
		}
	}

	return adviserList, nil
}

func (r *queryResolver) ChatList(ctx context.Context, input model.ChatFilter) ([]*model.Chat, error) {
	var chatList []*model.Chat
	r.USRDB.Order("created_at desc").Table("chat").Select("chat.id, chat.ledger_id, chat.user_id, chat.comment, chat.created_at, users.nickname").Joins("left join users on chat.user_id = users.id").Offset(input.First).Limit(input.Last).Find(&chatList, "ledger_id = ?", input.LedgerID)
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
	r.SAVDB.Table("savings_details sd").Order("saving_date DESC").Joins("RIGHT OUTER JOIN savings s ON sd.saving_id = s.id").Select("sd.id, sd.saving_id, sd.saving_amount, sd.note, sd.saving_date").Offset(input.First).Limit(input.Last).Find(&results, "s.userid=?", input.SavingsID)
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
