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

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SavingsDetails(ctx context.Context, input *model.SavingsDetailsFilter) ([]*model.SavingsDetail, error) {
	// 家計簿の利用履歴を取得
	var results []*model.SavingsDetail
	// offset, limit句の代わりにPKに対してBETWEEN句を利用しカラムを指定。速度が早いっぽい
	r.SAVDB.Order("saving_date DESC").Find(&results, "saving_id=? AND id BETWEEN ? AND ?", input.SavingsID, input.First, input.Last)
	return results, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
