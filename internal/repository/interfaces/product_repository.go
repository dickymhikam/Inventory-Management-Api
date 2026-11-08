package interfaces

import (
	"context"
	"inventory-api/internal/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *gorm.Tx, product *model.Products)error
	Update(ctx context.Context, tx *gorm.Tx, product *model.Products)error
	Delete(ctx context.Context, tx *gorm.Tx, product *model.Products)
	FindById(ctx context.Context, tx *gorm.Tx, productId uint64)(*model.Products, error)
	FindAll(ctx context.Context, tx *gorm.Tx)[]*model.Products
}