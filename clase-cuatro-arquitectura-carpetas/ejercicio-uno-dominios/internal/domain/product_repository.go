package domain

type ProductRepository interface {
	GetProducts() ([]Product, error)
	GetProductById(id int) (Product, error)
	GetProductByPriceGt(priceGt float64) ([]Product, error)
	AddProduct(product Product) error
	UpdateProduct(product Product) error
}
