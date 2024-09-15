package graph

import (
	"sync"

	"github.com/SV1Stail/test_ozon/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	posts       []*model.Post
	comments    []*model.Comment
	subscribers map[string][]chan *model.Comment // Храним каналы для подписок
	mu          sync.Mutex
}

// init Resolver struct
func NewResolver() *Resolver {
	return &Resolver{
		posts:       make([]*model.Post, 0),
		comments:    make([]*model.Comment, 0),
		subscribers: make(map[string][]chan *model.Comment),
	}
}
