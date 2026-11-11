package service

import (
	"context"
	"inventory-api/internal/model"
	"inventory-api/internal/repository/interfaces"

	"gorm.io/gorm"
)

type ProductService struct {
	productRepository interfaces.ProductRepository
	db                *gorm.DB
}

func NewProductServices(db *gorm.DB, productRepository interfaces.ProductRepository) *ProductService{
	return &ProductService{
		db: db,
		productRepository: productRepository,
	}
}

func (service *ProductService) CreateProduct(
	ctx context.Context,
	product *model.Products,
) error {
	return service.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return service.productRepository.Save(ctx, tx, product)
	})
}


func (service *ProductService) UpdateProduct(
	ctx context.Context,
	product *model.Products,
) error {
	return service.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return service.productRepository.Update(ctx, tx, product)
	})
}

func (service *ProductService) DeleteProduct(
	ctx context.Context,
	product *model.Products,
) error {
	return service.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return service.productRepository.Delete(ctx, tx, product)
	})
}

func (service *ProductService) FindByIDProduct(
	ctx context.Context,
	productID uint64,
)(*model.Products, error){
	product, err := service.productRepository.FindByID(
		ctx,
		service.db,
		productID,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductService) FindAllProduct(
	ctx context.Context, 
)([]*model.Products, error){
	product , err := service.productRepository.FindAll(
		ctx,
		service.db,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}