package repositories

import "service/internal/boundedcontexts/hello/domain/entities"

type IHelloRepository interface {
	GetUser(userId int64) (*entities.User, error)
}
