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

var customerRepo repository.CustomerRepository = repository.New()

func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.NewCustomer) (*model.Customer, error) {
	cus := &model.Customer{
		ID:           strconv.Itoa(rand.Int()),
		Cid:          input.Cid,
		CustomerCode: input.CustomerCode,
		Address1:     input.Address1,
		Address2:     input.Address2,
		Address3:     input.Address3,
		Address4:     input.Address4,
		CountryCode:  input.CountryCode,
		PostCode:     input.PostCode,
		WebAddress:   input.WebAddress,
		EmailAddress: input.EmailAddress,
		PhoneNo1:     input.PhoneNo1,
		PhoneNo2:     input.PhoneNo2,
	}
	customerRepo.Save(cus)
	return cus, nil

}

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	return customerRepo.FindAll(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
