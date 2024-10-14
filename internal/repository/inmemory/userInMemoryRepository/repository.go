package userInMemoryRepository

import (
	"auth-service-test/internal/domain"
	"auth-service-test/pkg/consts"
	"golang.org/x/sync/syncmap"
	"sync"
)

type UserRepository struct {
	mu    sync.Mutex
	users syncmap.Map //map[domain.Username]domain.User

}

func New() *UserRepository {
	repo := &UserRepository{
		users: syncmap.Map{},
	}

	defaultAdmin := domain.User{
		Firstname:   "Иван",
		Lastname:    "Иванов",
		Email:       "ivan@mail.ru",
		PhoneNumber: "+79998887766",
		Role:        consts.AdminRole,

		Username: "ivan123",
		Password: "$2a$10$eczUQF30cqdMDqUtpeKuJ.HhkqgxHa4O2iREOA24.3wqBimdlhUgO",
	}
	repo.users.Store(defaultAdmin.Username, defaultAdmin)

	return repo
}
