package service

import (
    "context"
    "sync"
)

var (
    once         sync.Once
    ExampleSvc   *ExampleService
    UserSvc      *UserService
    SimulatorSvc *SimulatorService
)

func Init(ctx context.Context) {
    once.Do(func() {
        ExampleSvc = NewExampleService(ctx)
        UserSvc = NewUserService(ctx)
        SimulatorSvc = NewSimulatorService(ctx)
    })
}
