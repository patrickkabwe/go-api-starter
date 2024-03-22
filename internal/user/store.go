package user

import (
	"github.com/patrickkabwe/go-api-starter/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) Store {
	return Store{db: db}
}

func (s *Store) CreateUser(user *types.User) error {
	return s.db.Create(user).Error
}

func (s *Store) GetUser(id uint) (*types.User, error) {
	user := &types.User{}
	err := s.db.First(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
