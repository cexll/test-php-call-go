package repositories

import (
	"service/internal/boundedcontexts/hello/domain/entities"
	"service/internal/svc"
)

type HelloRepository struct {
	Svc *svc.ServiceContext
}

func (repo *HelloRepository) GetUser(userId int64) (*entities.User, error) {
	user := entities.NewUser(userId, "mix-plus")
	return user, nil
}
