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

func (repository *productRepository) Update(
	ctx context.Context,
	tx *gorm.DB,
	product *model.Products,
)error {
	return tx.WithContext(ctx).Save(product).Error
}

func (repository *productRepository) Delete(
	ctx context.Context,
	tx *gorm.DB,
	product *model.Products,
) error {
	return tx.WithContext(ctx).Delete(product).Error
}

func (repository *productRepository) FindByID(
	ctx context.Context,
	tx *gorm.DB,
	productID uint64,
)(*model.Products, error){
	var product model.Products

	err := tx.WithContext(ctx).First(&product, productID).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (repository *productRepository) FindAll (
	ctx context.Context,
	tx *gorm.DB,
) ([]*model.Products, error) {
	var product []*model.Products

	err := tx.WithContext(ctx).Find(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}