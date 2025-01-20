package domain

type ProductRepository interface {
	GetProducts() ([]Product, error)
	GetProductById(id int) (Product, error)
	GetProductByPriceGt(priceGt float64) ([]Product, error)
	AddProduct(product Product) error
	UpdateProduct(product Product) error
	PatchProduct(id int, attributes map[string]any) error
	DeleteProduct(id int) error
}
