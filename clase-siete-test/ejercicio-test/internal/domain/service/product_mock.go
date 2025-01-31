package service

import "github.com/D-Sorrow/go-web-meli/clase-siete-test/ejercicio-test/internal/domain"

type ProductMock struct {
	GetProductsFunc    func() ([]domain.Product, error)
	GetProductByIDFunc func(id int) (domain.Product, error)
	AddProductFunc     func(domain.Product) error
	DeleteProductFunc  func(id int) error
	UpdateProductFun   func(product domain.Product) error
	PatchProductFunc   func(id int, attributes map[string]any) error
}

func (p *ProductMock) GetProducts() ([]domain.Product, error) {
	return p.GetProductsFunc()
}

func (p *ProductMock) GetProductById(id int) (domain.Product, error) {
	return p.GetProductByIDFunc(id)
}

func (p *ProductMock) GetProductByPriceGt(priceGt float64) ([]domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductMock) AddProduct(product domain.Product) error {
	return p.AddProductFunc(product)
}

func (p *ProductMock) UpdateProduct(product domain.Product) error {
	return p.UpdateProductFun(product)
}

func (p *ProductMock) PatchProduct(id int, attributes map[string]any) error {
	return p.PatchProductFunc(id, attributes)
}

func (p *ProductMock) DeleteProduct(id int) error {
	return p.DeleteProductFunc(id)
}
