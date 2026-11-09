package interfaces

import (
	"context"
	"inventory-api/internal/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *gorm.DB, category *model.Categories) error
	Update(ctx context.Context, tx *gorm.DB, category *model.Categories) error
	Delete(ctx context.Context, tx *gorm.DB, category *model.Categories)
	FindByID(ctx context.Context, tx *gorm.DB, categoryID uint64) (*model.Categories, error)
	FindAll(ctx context.Context, tx *gorm.DB) []*model.Categories
}
