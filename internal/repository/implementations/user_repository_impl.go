package implementations

import (
	"context"
	"inventory-api/internal/model"
	"inventory-api/internal/repository/interfaces"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repository *userRepository) Save(
	ctx context.Context,
	tx *gorm.DB,
	user *model.Users,
) error {
	return tx.WithContext(ctx).Create(user).Error
}

func (repository *userRepository) Update(
	ctx context.Context,
	tx *gorm.DB,
	user *model.Users,
) error {
	return tx.WithContext(ctx).Save(user).Error
}

func (repository *userRepository) FindByID(
	ctx context.Context,
	tx *gorm.DB,
	userID uint64,
)(*model.Users, error){
	var user model.Users

	err := tx.WithContext(ctx).First(&user, userID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *userRepository) FindByEmail(
	ctx context.Context,
	tx *gorm.DB,
	email string,
)(*model.Users, error){
	var user model.Users

	err := tx.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}