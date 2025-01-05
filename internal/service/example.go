package service

import "go-sample/internal/store"

// ExampleService Example 业务层
type ExampleService struct {
	store *store.ExampleStore
}

// NewExampleService 提供者
func NewExampleService(store *store.ExampleStore) *ExampleService {
	return &ExampleService{store: store}
}
