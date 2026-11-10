package implementations

import (
	"context"
	"inventory-api/internal/model"
	"inventory-api/internal/repository/interfaces"

	"gorm.io/gorm"
)

type movementRepository struct {
	db *gorm.DB
}

func NewMovementRepository(db *gorm.DB) interfaces.MovementRepository {
	return &movementRepository{
		db: db,
	}
}

func (repository *movementRepository) Save(
	ctx context.Context,
	tx *gorm.DB,
	movement *model.Movements,
) error {
	return tx.WithContext(ctx).Create(movement).Error
}

func (repository *movementRepository) Update(
	ctx context.Context,
	tx *gorm.DB,
	movement *model.Movements,
) error {
	return tx.WithContext(ctx).Save(movement).Error
}

func (repository *movementRepository) Delete(
	ctx context.Context,
	tx *gorm.DB,
	movement *model.Movements,
) error {
	return tx.WithContext(ctx).Delete(movement).Error
}

func (repository *movementRepository) FindByID(
	ctx context.Context,
	tx *gorm.DB,
	movementID uint64,
)(*model.Movements, error){
	var movement model.Movements

	err := tx.WithContext(ctx).First(&movement, movementID).Error
	if err != nil {
		return nil, err
	}
	
	return &movement, nil
}

func (repository *movementRepository) FindAll(
	ctx context.Context,
	tx *gorm.DB,
)([]*model.Movements, error){
	var movement []*model.Movements

	err := tx.WithContext(ctx).Find(&movement).Error
	if err != nil {
		return nil, err
	}

	return movement, nil
}