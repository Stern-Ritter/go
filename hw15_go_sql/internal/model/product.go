package model

import "github.com/Stern-Ritter/go/hw15_go_sql/internal/utils"

type Product struct {
	ID    int64
	Name  string
	Price float64
}

type CreateProductDto struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateProductDto struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductDto struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func CreateProductDtoToProduct(dto CreateProductDto) Product {
	return Product{
		Name:  dto.Name,
		Price: dto.Price,
	}
}

func UpdateProductDtoToProduct(dto UpdateProductDto, product Product) Product {
	return Product{
		Name:  utils.Coalesce(dto.Name, product.Name),
		Price: utils.Coalesce(dto.Price, product.Price),
	}
}

func ProductToProductDto(product Product) ProductDto {
	return ProductDto(product)
}

func ProductsToProductsDto(products []Product) []ProductDto {
	productsDto := make([]ProductDto, len(products))
	if products == nil {
		return productsDto
	}

	for i, product := range products {
		productsDto[i] = ProductToProductDto(product)
	}

	return productsDto
}
