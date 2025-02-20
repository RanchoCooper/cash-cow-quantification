package service

import (
	"context"

	"quants/adapter/repository"
	"quants/domain/entity"
	"quants/domain/repo"
	"quants/util/logger"
)

type ExampleService struct {
	Repository repo.IExampleRepo
}

func NewExampleService(ctx context.Context) *ExampleService {
	srv := &ExampleService{Repository: repository.Example}
	logger.Log.Info(ctx, "example service init successfully")
	return srv
}

func (e *ExampleService) Create(ctx context.Context, example *entity.Example) (*entity.Example, error) {
	example, err := e.Repository.Create(ctx, nil, example)
	if err != nil {
		return nil, err
	}
	return example, nil
}

func (e *ExampleService) Delete(ctx context.Context, id int) error {
	err := e.Repository.Delete(ctx, nil, id)
	if err != nil {
		return err
	}
	return nil
}

func (e *ExampleService) Update(ctx context.Context, example *entity.Example) error {
	err := e.Repository.Update(ctx, nil, example)
	if err != nil {
		return err
	}
	return nil
}

func (e *ExampleService) Get(ctx context.Context, id int) (*entity.Example, error) {
	example, err := e.Repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return example, nil
}
