package service

import (
	"errors"
	"github.com/D-Sorrow/go-web-meli/clase-cuatro-arquitectura-carpetas/ejercicio-uno-dominios/internal/domain"
)

type ProductService struct {
	ProductRepo domain.ProductRepository
}

func (p *ProductService) GetProducts() ([]domain.Product, error) {
	products, err := p.ProductRepo.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductService) GetProductById(id int) (domain.Product, error) {
	if id == 0 {
		return domain.Product{}, errors.New("id is required")
	}
	product, err := p.ProductRepo.GetProductById(id)

	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (p *ProductService) GetProductByPriceGt(priceGt float64) ([]domain.Product, error) {

	products, err := p.ProductRepo.GetProductByPriceGt(priceGt)

	if err != nil {
		return nil, err
	}

	return products, nil
}
func (p *ProductService) AddProduct(product domain.Product) error {
	err := p.ProductRepo.AddProduct(product)
	if err != nil {
		return errors.New("error adding product")
	}
	return nil
}

func NewProductService(productRepo domain.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: productRepo}
}
