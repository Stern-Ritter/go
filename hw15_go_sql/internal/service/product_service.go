package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	er "github.com/Stern-Ritter/go/hw15_go_sql/internal/errors"
	"github.com/Stern-Ritter/go/hw15_go_sql/internal/model"
	"github.com/Stern-Ritter/go/hw15_go_sql/internal/storage"
	"github.com/sirupsen/logrus"
)

type ProductService interface {
	CreateProduct(ctx context.Context, productDto model.CreateProductDto) (model.ProductDto, error)
	UpdateProduct(ctx context.Context, productDto model.UpdateProductDto, productID int64) (model.ProductDto, error)
	DeleteProduct(ctx context.Context, productID int64) (model.ProductDto, error)
	GetAllProductsByPrice(ctx context.Context, min *float64, max *float64) ([]model.ProductDto, error)
	GetProduct(ctx context.Context, productID int64) (model.Product, error)
	GetProductsByIDs(ctx context.Context, productsIDs []int64) ([]model.Product, error)
}

type ProductServiceImpl struct {
	productStorage storage.ProductStorage
	Logger         *logrus.Logger
}

func NewProductService(productStorage storage.ProductStorage, logger *logrus.Logger) ProductService {
	return &ProductServiceImpl{
		productStorage: productStorage,
		Logger:         logger,
	}
}

func (p *ProductServiceImpl) CreateProduct(ctx context.Context,
	createProductDto model.CreateProductDto,
) (model.ProductDto, error) {
	product := model.CreateProductDtoToProduct(createProductDto)
	productID, err := p.productStorage.CreateProduct(ctx, product)
	if err != nil {
		p.Logger.WithError(err).Errorf("Failed to create product: %v", createProductDto)
		return model.ProductDto{}, err
	}

	product.ID = productID
	productDto := model.ProductToProductDto(product)

	p.Logger.Infof("Product created successfully: %v", productDto)
	return productDto, nil
}

//nolint:dupl
func (p *ProductServiceImpl) UpdateProduct(ctx context.Context,
	updateProductDto model.UpdateProductDto, productID int64,
) (model.ProductDto, error) {
	product, err := p.GetProduct(ctx, productID)
	if err != nil {
		return model.ProductDto{}, err
	}

	updatedProduct := model.UpdateProductDtoToProduct(updateProductDto, product)
	updatedProduct.ID = productID

	err = p.productStorage.UpdateProduct(ctx, updatedProduct)
	if err != nil {
		p.Logger.WithError(err).Errorf("Failed to update product: %v", updatedProduct)
		return model.ProductDto{}, err
	}

	savedProduct, err := p.productStorage.GetProductByID(ctx, productID)
	if err != nil {
		p.Logger.WithError(err).Errorf("Failed to get product by ID: %d", productID)
		return model.ProductDto{}, err
	}
	productDto := model.ProductToProductDto(savedProduct)

	p.Logger.Infof("Product updated successfully: %v", productDto)
	return productDto, nil
}

func (p *ProductServiceImpl) DeleteProduct(ctx context.Context, productID int64) (model.ProductDto, error) {
	product, err := p.GetProduct(ctx, productID)
	if err != nil {
		return model.ProductDto{}, err
	}

	err = p.productStorage.DeleteProduct(ctx, productID)
	if err != nil {
		p.Logger.WithError(err).Errorf("Failed to delete product: %v", productID)
		return model.ProductDto{}, err
	}
	productDto := model.ProductToProductDto(product)

	p.Logger.Infof("Product deleted successfully: %v", productDto)
	return productDto, nil
}

func (p *ProductServiceImpl) GetAllProductsByPrice(ctx context.Context, min *float64,
	max *float64,
) ([]model.ProductDto, error) {
	products, err := p.productStorage.GetAllProductsByPrice(ctx, min, max)
	if err != nil {
		p.Logger.WithError(err).Errorf("Failed to get products by price range")
		return nil, err
	}
	productsDto := model.ProductsToProductsDto(products)

	p.Logger.Infof("Products by price found successfully: %v", productsDto)
	return productsDto, nil
}

func (p *ProductServiceImpl) GetProduct(ctx context.Context, productID int64) (model.Product, error) {
	product, err := p.productStorage.GetProductByID(ctx, productID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.Logger.WithError(err).Infof("Product with ID: %d doesn't exist", productID)
			return model.Product{}, er.NewNotFoundError(fmt.Sprintf("Product with ID: %d doesn`t exist", productID), err)
		}
		p.Logger.WithError(err).Errorf("Failed to get product by ID: %d", productID)
		return model.Product{}, err
	}

	return product, nil
}

func (p *ProductServiceImpl) GetProductsByIDs(ctx context.Context, productsIDs []int64) ([]model.Product, error) {
	products, err := p.productStorage.GetProductsByIDs(ctx, productsIDs)
	if err != nil {
		p.Logger.WithError(err).Errorf("Failed to get products by IDs: %v", productsIDs)
		return make([]model.Product, 0), err
	}

	return products, nil
}
