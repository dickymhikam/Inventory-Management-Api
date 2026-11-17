package service

import (
	"context"
	"inventory-api/internal/model"
	"inventory-api/internal/repository/interfaces"

	"gorm.io/gorm"
)

type UserService struct {
	userRepository interfaces.UserRepository
	db *gorm.DB
}

func NewUserService(userRepository interfaces.UserRepository, db *gorm.DB) *UserService{
	return &UserService{
		userRepository: userRepository,
		db: db,
	}
}

func (service *UserService) SaveUser(
	ctx context.Context,
	user *model.Users,
)error{
	return service.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return service.userRepository.Save(ctx, tx, user)
	})
}

func (service *UserService) UpdateUser(
	ctx context.Context,
	user *model.Users,
) error {
	return service.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return service.userRepository.Update(ctx, tx, user)
	})
}

func (service *UserService) FindByIDUser(
	ctx context.Context,
	userID uint64,
) (*model.Users, error) {
	user, err := service.userRepository.FindByID(
		ctx,
		service.db,
		userID,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) FindByEmailUser(
	ctx context.Context,
	email string,
) (*model.Users, error) {
	user, err := service.userRepository.FindByEmail(
		ctx,
		service.db,
		email,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}