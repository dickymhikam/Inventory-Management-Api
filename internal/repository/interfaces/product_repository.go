package interfaces

import (
	"context"
	"inventory-api/internal/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *gorm.DB, product *model.Products) error
	Update(ctx context.Context, tx *gorm.DB, product *model.Products) error
	Delete(ctx context.Context, tx *gorm.DB, product *model.Products) error
	FindByID(ctx context.Context, tx *gorm.DB, productID uint64) (*model.Products, error)
	FindAll(ctx context.Context, tx *gorm.DB) ([]*model.Products, error)
}
