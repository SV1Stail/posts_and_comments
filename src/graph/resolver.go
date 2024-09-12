package graph

import "github.com/SV1Stail/test_ozon/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	posts    []*model.Post
	comments []*model.Comment
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Subscription() SubscriptionResolver {
	return nil
}

// Optionally, if you implement subscriptions, you can add the Subscription resolver as well
// func (r *Resolver) Subscription() SubscriptionResolver {
// 	return &subscriptionResolver{r}
// }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
