package graph

import "github.com/deyr02/crm_mongo_graphql/graph/model"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	tables []*model.Table
}
