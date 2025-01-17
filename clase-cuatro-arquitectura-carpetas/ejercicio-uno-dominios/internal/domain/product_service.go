package domain

type ProductService interface {
	GetProducts() ([]Product, error)
	GetProductById(id int) (Product, error)
	GetProductByPriceGt(priceGt float64) ([]Product, error)
	AddProduct(product Product) error
}
