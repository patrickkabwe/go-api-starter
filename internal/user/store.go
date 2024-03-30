package user

import (
	"github.com/patrickkabwe/go-api-starter/types"
	"gorm.io/gorm"
)

type store struct {
	db *gorm.DB
}

type Store interface {
	CreateUser(user *types.User) error
	GetUser(id string) (*types.User, error)
}

func NewStore(db *gorm.DB) Store {
	return &store{db: db}
}

func (s *store) CreateUser(user *types.User) error {
	// TODO: add validation
	return s.db.Create(user).Error
}

func (s *store) GetUser(id string) (*types.User, error) {
	// TODO: add validation
	user := &types.User{}
	err := s.db.First(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
