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

func (r *mutationResolver) AddColumn(ctx context.Context, id string, input model.NewCustomField) (*model.Table, error) {
	column := &model.CustomField{
		FieldID:        strconv.Itoa(rand.Int()),
		FieldName:      input.FieldName,
		DataType:       input.DataType,
		FieldType:      input.FieldType,
		MaxValue:       input.MaxValue,
		MinValue:       input.MinValue,
		DefaultValue:   input.DefaultValue,
		IsRequired:     input.IsRequired,
		Visibility:     input.Visibility,
		PossibleValues: input.PossibleValues,
	}
	return tableRepo.AddColumn(id, column), nil

}

func (r *mutationResolver) ModifyColumn(ctx context.Context, tableid string, columnid string, input model.NewCustomField) (*model.Table, error) {
	return tableRepo.ModifyTableColumn(tableid, columnid, &input), nil

}

func (r *mutationResolver) DeleteColumn(ctx context.Context, tableid string, columnid string) (*model.Table, error) {
	return tableRepo.DeleteTableColumn(tableid, columnid), nil
}

func (r *mutationResolver) CreateTable(ctx context.Context, input model.NewTable) (*model.Table, error) {
	var customFields []*model.CustomField

	for _, element := range input.Fields {
		ele := &model.CustomField{
			FieldID:        strconv.Itoa(rand.Int()),
			FieldName:      element.FieldName,
			DataType:       element.DataType,
			FieldType:      element.FieldType,
			MaxValue:       element.MaxValue,
			MinValue:       element.MinValue,
			DefaultValue:   element.DefaultValue,
			IsRequired:     element.IsRequired,
			Visibility:     element.Visibility,
			PossibleValues: element.PossibleValues,
		}

		customFields = append(customFields, ele)
	}

	tab := &model.Table{
		TableID:   strconv.Itoa(rand.Int()),
		TableName: input.TableName,
		Fields:    customFields,
	}

	tableRepo.Save(tab)
	return tab, nil
}

func (r *mutationResolver) DeleteTable(ctx context.Context, id string) (*model.Table, error) {
	return tableRepo.DeleteTable(id), nil
}

func (r *queryResolver) Table(ctx context.Context, id string) (*model.Table, error) {
	return tableRepo.FindTableById(id), nil
}

func (r *queryResolver) Tables(ctx context.Context) ([]*model.Table, error) {
	return tableRepo.FindAll(), nil
}

func (r *mutationResolver) AddData(ctx context.Context, collectionName string, data string) (*string, error) {
	return tableRepo.AddData(collectionName, data), nil
}

func (r *queryResolver) GetAllData(ctx context.Context, collectionName string) ([]*string, error) {
	return tableRepo.GetAllData(collectionName), nil
}

func (r *queryResolver) GetData(ctx context.Context, collectionName string, query string) ([]*string, error) {
	return tableRepo.GetData(collectionName, query), nil
}

func (r *mutationResolver) SaveData(ctx context.Context, collectionName string, input *model.NewRecord) (*model.Record, error) {
	var fieldValues []*model.FieldValue

	for _, element := range input.Data {
		ele := &model.FieldValue{
			Key:      element.Key,
			DataType: element.DataType,
			Value:    element.Value,
		}

		fieldValues = append(fieldValues, ele)
	}

	rec := &model.Record{
		RecordID: strconv.Itoa(rand.Int()),
		Data:     fieldValues,
	}

	tableRepo.SaveData(collectionName, rec)
	return rec, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
