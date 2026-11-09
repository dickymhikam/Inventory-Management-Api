package interfaces

import (
	"context"
	"inventory-api/internal/model"

	"gorm.io/gorm"
)

type MovementRepository interface {
	Save(ctx context.Context, tx *gorm.DB, movement *model.Movements) error
	Update(ctx context.Context, tx *gorm.DB, movement *model.Movements) error
	Delete(ctx context.Context, tx *gorm.DB, movement *model.Movements)
	FindById(ctx context.Context, tx *gorm.DB, movementId uint64) (*model.Movements, error)
	FindAll(ctx context.Context, tx *gorm.DB) []*model.Movements
}
