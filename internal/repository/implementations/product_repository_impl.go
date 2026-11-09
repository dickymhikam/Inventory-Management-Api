package implementations

import (
	"context"
	"inventory-api/internal/model"
	"inventory-api/internal/repository/interfaces"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (repository *productRepository) Save(
	ctx context.Context,
	tx *gorm.DB,
	product *model.Products,
) error {
	return tx.WithContext(ctx).Create(product).Error
}

