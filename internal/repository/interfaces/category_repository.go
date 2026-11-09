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
	FindById(ctx context.Context, tx *gorm.DB, categoryId uint64) (*model.Categories, error)
	FindAll(ctx context.Context, tx *gorm.DB) []*model.Categories
}
