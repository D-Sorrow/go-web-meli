package file

import (
	"errors"
	"fmt"
	"github.com/D-Sorrow/go-web-meli/clase-seis-variables-entorno/ejercicios-variables-entorno/internal/domain"
	"github.com/D-Sorrow/go-web-meli/clase-seis-variables-entorno/ejercicios-variables-entorno/internal/infrastucture/persistence/storage"
	"time"
)

type ProductJSON struct {
	storage storage.StorageJSON
}

func NewProductJSON(storageJSON storage.StorageJSON) (*ProductJSON, error) {
	return &ProductJSON{storageJSON}, nil
}

func (product *ProductJSON) GetProducts() ([]domain.Product, error) {
	return product.storage.ReadAll()
}
func (product *ProductJSON) GetProductById(id int) (domain.Product, error) {
	products, err := product.storage.ReadAll()
	if err != nil {
		return domain.Product{}, err
	}
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("Product not found")
}
func (product *ProductJSON) GetProductByPriceGt(priceGt float64) ([]domain.Product, error) {
	var productUpThanPriceGt []domain.Product
	products, err := product.storage.ReadAll()
	if err != nil {
		return productUpThanPriceGt, err
	}
	for _, product := range products {
		if product.Price > priceGt {
			productUpThanPriceGt = append(productUpThanPriceGt, product)
		}
	}
	return productUpThanPriceGt, nil
}
func (product *ProductJSON) AddProduct(productAdd domain.Product) error {
	products, err := product.storage.ReadAll()
	if err != nil {
		return err
	}

	if ValidateCodeValue(productAdd.Code_value, products) {
		return errors.New("Product code is already in use")
	}
	if ValidateDateExpiration(productAdd.Expiration) {
		return errors.New("Product date expiration is not valid")
	}

	if len(products) > 0 {
		productAdd.Id = (products)[len(products)-1].Id + 1
	} else {
		productAdd.Id = 1
	}
	products = append(products, productAdd)
	return product.storage.WriteAll(products)
}

func (product *ProductJSON) UpdateProduct(productUpdate domain.Product) error {

	products, err := product.storage.ReadAll()
	if err != nil {
		return err
	}
	for _, productToUpdate := range products {
		if productToUpdate.Id == productUpdate.Id {
			products[productToUpdate.Id-1] = productUpdate
			return product.storage.WriteAll(products)
		}
	}
	if ValidateCodeValue(productUpdate.Code_value, products) {
		return errors.New("Product code is already in use")
	}
	if ValidateDateExpiration(productUpdate.Expiration) {
		return errors.New("Product date expiration is not valid")
	}
	productUpdate.Id = products[len(products)-1].Id + 1
	products = append(products, productUpdate)
	return product.storage.WriteAll(products)

}

func (product *ProductJSON) PatchProduct(id int, attributes map[string]any) error {
	products, err := product.storage.ReadAll()
	if err != nil {
		return err
	}
	for _, product := range products {
		if product.Id == id {
			for key, value := range attributes {
				switch key {
				case "name":
					product.Name = value.(string)
				case "quantity":
					product.Quantity = int(value.(float64))
				case "code_value":
					fmt.Println(ValidateCodeValue(value.(string), products))
					if ValidateCodeValue(value.(string), products) {
						return errors.New("Product code is already in use")
					}
					product.Code_value = value.(string)
				case "is_published":
					product.Is_published = value.(bool)
				case "expiration":
					product.Expiration = value.(string)
				case "price":
					product.Price = value.(float64)
				default:
					return errors.New("Error patching product")
				}
			}
			products[id-1] = product
			break
		}
	}
	return product.storage.WriteAll(products)
}

func (p *ProductJSON) DeleteProduct(id int) error {
	products, err := p.storage.ReadAll()
	if err != nil {
		return err
	}
	for i, product := range products {
		if product.Id == id {
			products = append(products[:i], products[i+1:]...)
		}
	}
	return p.storage.WriteAll(products)
}

func ValidateCodeValue(code string, products []domain.Product) bool {

	for _, product := range products {
		if product.Code_value == code {
			fmt.Println("" + product.Code_value)
			return true
		}
	}
	return false
}

func ValidateDateExpiration(date string) bool {
	dateExp, err := time.Parse("02/01/2006", date)
	if err != nil {
		return true
	}
	if dateExp.Year() < time.Now().Year() {
		return true
	}
	return false
}
