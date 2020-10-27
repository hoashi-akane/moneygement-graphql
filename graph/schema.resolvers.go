package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hoashi-akane/moneygement-graphql/graph/generated"
	"github.com/hoashi-akane/moneygement-graphql/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
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
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateExpenseDetail(ctx context.Context, input *model.NewExpenseDetail) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Saving(ctx context.Context) (*model.Saving, error) {
	//itemList := GetPreloads(ctx)
	//for _, item := range itemList {
	//	fmt.Println(item)
	//}
	saving := &model.Saving{nil}
	return saving, nil
}

func (r *savingResolver) SavingsDetails(ctx context.Context, obj *model.Saving, input model.SavingsDetailsFilter) ([]*model.SavingsDetail, error) {
	// 家計簿の利用履歴を取得
	fmt.Println("実行")
	var results []*model.SavingsDetail
	// offset, limit句の代わりにPKに対してBETWEEN句を利用しカラムを指定。速度が早いっぽい
	r.SAVDB.Order("saving_date DESC").Find(&results, "saving_id=? AND id BETWEEN ? AND ?", input.SavingsID, input.First, input.Last)
	return results, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Saving returns generated.SavingResolver implementation.
func (r *Resolver) Saving() generated.SavingResolver { return &savingResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type savingResolver struct{ *Resolver }

