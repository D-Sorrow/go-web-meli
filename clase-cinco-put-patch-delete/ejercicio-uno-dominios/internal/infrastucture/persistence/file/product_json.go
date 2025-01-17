package file

import (
	"bufio"
	"encoding/json"
	"errors"
	"github.com/D-Sorrow/go-web-meli/clase-cuatro-arquitectura-carpetas/ejercicio-uno-dominios/internal/domain"
	"io"
	"os"
	"time"
)

type ProductJSON struct {
	file *os.File
}

var (
	errMessage = "Error reading product file"
	products   []domain.Product
)

func NewProductJSON(path string) (*ProductJSON, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("error opening product file")
	}

	reader := bufio.NewReader(file)

	byteValue, err := io.ReadAll(reader)

	err = json.Unmarshal(byteValue, &products)

	return &ProductJSON{file}, nil
}

func (product *ProductJSON) GetProducts() ([]domain.Product, error) {

	if len(products) != 0 {
		return products, nil
	}
	return products, nil
}
func (product *ProductJSON) GetProductById(id int) (domain.Product, error) {
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("Product not found")
}
func (product *ProductJSON) GetProductByPriceGt(priceGt float64) ([]domain.Product, error) {
	var productUpThanPriceGt []domain.Product
	for _, product := range products {
		if product.Price > priceGt {
			productUpThanPriceGt = append(productUpThanPriceGt, product)
		}
	}
	return productUpThanPriceGt, nil
}
func (product *ProductJSON) AddProduct(productAdd domain.Product) error {
	if ValidateCodeValue(productAdd.Code_value) {
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
	return nil
}

func ValidateCodeValue(code string) bool {

	for _, product := range products {
		if product.Code_value == code {
			return true
		}
	}
	return false
}

func ValidateDateExpiration(date string) bool {
	dateExp, err := time.Parse("2006/01/02", date)
	if err != nil {
		return true
	}
	if dateExp.Year() < time.Now().Year() {
		return true
	}
	return false
}
