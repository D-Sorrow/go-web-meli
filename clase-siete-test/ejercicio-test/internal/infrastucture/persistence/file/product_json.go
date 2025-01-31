package file

import (
	"errors"
	"fmt"
	"github.com/D-Sorrow/go-web-meli/clase-siete-test/ejercicio-test/internal/domain"
	"time"
)

type ProductJSON struct {
	db []domain.Product
}

func NewProductJSON(storageJSON []domain.Product) (*ProductJSON, error) {
	return &ProductJSON{storageJSON}, nil
}

func (product *ProductJSON) GetProducts() ([]domain.Product, error) {
	return product.db, nil
}
func (product *ProductJSON) GetProductById(id int) (domain.Product, error) {
	for _, productTo := range product.db {
		if productTo.Id == id {
			return productTo, nil
		}
	}
	return domain.Product{}, errors.New("Product not found")
}
func (product *ProductJSON) GetProductByPriceGt(priceGt float64) ([]domain.Product, error) {
	var productUpThanPriceGt []domain.Product
	for _, productTo := range product.db {
		if productTo.Price > priceGt {
			productUpThanPriceGt = append(productUpThanPriceGt, productTo)
		}
	}
	return productUpThanPriceGt, nil
}
func (product *ProductJSON) AddProduct(productAdd domain.Product) error {

	if ValidateCodeValue(productAdd.Code_value, product.db) {
		return errors.New("Product code is already in use")
	}
	if ValidateDateExpiration(productAdd.Expiration) {
		return errors.New("Product date expiration is not valid")
	}

	if len(product.db) > 0 {
		productAdd.Id = (product.db)[len(product.db)-1].Id + 1
	} else {
		productAdd.Id = 1
	}
	product.db = append(product.db, productAdd)
	return nil
}

func (product *ProductJSON) UpdateProduct(productUpdate domain.Product) error {
	for _, productToUpdate := range product.db {
		if productToUpdate.Id == productUpdate.Id {
			product.db[productToUpdate.Id-1] = productUpdate
			return nil
		}
	}
	if ValidateCodeValue(productUpdate.Code_value, product.db) {
		return errors.New("Product code is already in use")
	}
	if ValidateDateExpiration(productUpdate.Expiration) {
		return errors.New("Product date expiration is not valid")
	}
	productUpdate.Id = product.db[len(product.db)-1].Id + 1
	product.db = append(product.db, productUpdate)
	return nil

}

func (product *ProductJSON) PatchProduct(id int, attributes map[string]any) error {

	for _, productTo := range product.db {
		if productTo.Id == id {
			for key, value := range attributes {
				switch key {
				case "name":
					productTo.Name = value.(string)
				case "quantity":
					productTo.Quantity = int(value.(float64))
				case "code_value":
					fmt.Println(ValidateCodeValue(value.(string), product.db))
					if ValidateCodeValue(value.(string), product.db) {
						return errors.New("Product code is already in use")
					}
					productTo.Code_value = value.(string)
				case "is_published":
					productTo.Is_published = value.(bool)
				case "expiration":
					productTo.Expiration = value.(string)
				case "price":
					productTo.Price = value.(float64)
				default:
					return errors.New("Error patching product")
				}
			}
			product.db[id-1] = productTo
			break
		}
	}
	return errors.New("Product not found")
}

func (product *ProductJSON) DeleteProduct(id int) error {
	for i, productTo := range product.db {
		if productTo.Id == id {
			product.db = append(product.db[:i], product.db[i+1:]...)
			return nil
		}
	}
	return errors.New("Product not found")
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
