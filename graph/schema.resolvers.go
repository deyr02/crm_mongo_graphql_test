package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/deyr02/crm_mongo_graphql/graph/generated"
	"github.com/deyr02/crm_mongo_graphql/graph/model"
	"github.com/deyr02/crm_mongo_graphql/repository"
)

var tableRepo repository.TableRepository = repository.New()

func (r *mutationResolver) CreateTable(ctx context.Context, input model.NewTable) (*model.Table, error) {
	var customFields []*model.CustomField

	for _, element := range input.Fields {
		ele := &model.CustomField{
			FieldID:      strconv.Itoa(rand.Int()),
			FieldName:    element.FieldName,
			DataType:     element.DataType,
			Value:        element.Value,
			MaxValue:     element.MaxValue,
			MinValue:     element.MinValue,
			DefaultValue: element.DefaultValue,
			IsRequired:   element.IsRequired,
			Visibility:   element.Visibility,
		}

		customFields = append(customFields, ele)
	}

	tab := &model.Table{
		TableID: strconv.Itoa(rand.Int()),
		Fields:  customFields,
	}

	tableRepo.Save(tab)
	return tab, nil
}

func (r *queryResolver) Table(ctx context.Context, id string) (*model.Table, error) {
	return tableRepo.FindTableById(id), nil
}

func (r *queryResolver) Tables(ctx context.Context) ([]*model.Table, error) {
	return tableRepo.FindAll(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
