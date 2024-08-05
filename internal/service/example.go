package service

import "go-sample/internal/dao"

// ExampleService Example 业务层
type ExampleService struct {
	repo *dao.ExampleRepo
}

// NewExampleService 提供者
func NewExampleService(repo *dao.ExampleRepo) *ExampleService {
	return &ExampleService{repo: repo}
}