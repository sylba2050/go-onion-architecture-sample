package interactor

import (
	"github.com/jinzhu/gorm"
	"github.com/sylba2050/go-onion-architecture-sample/domain/repository"
	"github.com/sylba2050/go-onion-architecture-sample/domain/service"
	"github.com/sylba2050/go-onion-architecture-sample/infrastructure/persistence/datastore"
	"github.com/sylba2050/go-onion-architecture-sample/presenter/http/handler"
	"github.com/sylba2050/go-onion-architecture-sample/usecase"
)

// Interactor interfase Intractorは安易DIコンテナとしての役割を持つ.
type Interactor interface {
	NewUserRepository() repository.UserRepository
	NewUserService() service.UserService
	NewUserUseCase() usecase.UserUseCase
	NewUserHandler() handler.UserHandler
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	Conn *gorm.DB
}

// NewInteractor intractorを取得します.
func NewInteractor(Conn *gorm.DB) Interactor {
	return &interactor{Conn}
}

type appHandler struct {
	handler.UserHandler
	// embed all handler interfaces
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.UserHandler = i.NewUserHandler()
	return appHandler
}

func (i *interactor) NewUserRepository() repository.UserRepository {
	return datastore.NewUserRepository(i.Conn)
}

func (i *interactor) NewUserService() service.UserService {
	return service.NewUserService(i.NewUserRepository())
}

func (i *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(i.NewUserRepository())
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserUseCase())
}
