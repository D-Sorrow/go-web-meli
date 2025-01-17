package service

import (
	"errors"
	"github.com/D-Sorrow/go-web-meli/clase-cinco-put-patch-delete/ejercicio-implement-put-patch-delete/internal/domain"
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

func (p *ProductService) UpdateProduct(product domain.Product) error {
	err := p.ProductRepo.UpdateProduct(product)
	if err != nil {
		return errors.New("error updating product")
	}
	return nil
}

func (p *ProductService) PatchProduct(id int, attributes map[string]any) error {
	err := p.ProductRepo.PatchProduct(id, attributes)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductService) DeleteProduct(id int) error {
	err := p.ProductRepo.DeleteProduct(id)
	if err != nil {
		return errors.New("error deleting product")
	}
	return nil
}
func NewProductService(productRepo domain.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: productRepo}
}
