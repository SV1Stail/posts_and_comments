package graph

import "github.com/SV1Stail/test_ozon/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	posts    []*model.Post
	comments []*model.Comment
}
