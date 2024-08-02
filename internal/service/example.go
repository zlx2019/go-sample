package service

import "go-sample/internal/data"

type ExampleService struct {
	repo data.ExampleRepo
}

func ProvideExampleService(repo data.ExampleRepo) ExampleService {
	return ExampleService{repo: repo}
}