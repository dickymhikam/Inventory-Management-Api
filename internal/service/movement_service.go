package service

import (
	"context"
	"inventory-api/internal/model"
	"inventory-api/internal/repository/interfaces"

	"gorm.io/gorm"
)

type MovementService struct {
	movementRepository interfaces.MovementRepository
	db                 *gorm.DB
}

func NewMovementService(db *gorm.DB, movementRepository interfaces.MovementRepository) *MovementService {
	return &MovementService{
		db:                 db,
		movementRepository: movementRepository,
	}
}

func (service *MovementService) CreateMovement(
	ctx context.Context,
	movement *model.Movements,
) error {
	return service.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return service.movementRepository.Save(ctx, tx, movement)
	})
}

func (service *MovementService) UpdateMovement(
	ctx context.Context,
	movement *model.Movements,
) error {
	return service.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return service.movementRepository.Update(ctx, tx, movement)
	})
}

func (service *MovementService) DeleteMovement(
	ctx context.Context,
	movement *model.Movements,
) error {
	return service.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return service.movementRepository.Delete(ctx, tx, movement)
	})
}

func (service *MovementService) FindByIDMovement(
	ctx context.Context,
	movementID uint64,
) (*model.Movements, error) {
	movement, err := service.movementRepository.FindByID(
		ctx,
		service.db,
		movementID,
	)

	if err != nil {
		return nil, err
	}

	return movement, nil
}

func (service *MovementService) FindAllMovement(
	ctx context.Context,
) ([]*model.Movements, error) {

	movement, err := service.movementRepository.FindAll(
		ctx,
		service.db,
	)

	if err != nil {
		return nil, err
	}

	return movement, nil
}
