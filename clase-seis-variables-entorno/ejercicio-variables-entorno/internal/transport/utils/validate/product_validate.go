package validate

import (
	"errors"

	"github.com/D-Sorrow/go-web-meli/clase-seis-variables-entorno/ejercicios-variables-entorno/internal/domain"
)

func Validate(product domain.Product) error {

	errorProduct := ValidatePrice(product.Price)
	if errorProduct != nil {
		return errorProduct
	}
	errorProduct = ValidateName(product.Name)
	if errorProduct != nil {
		return errorProduct
	}
	errorProduct = ValidateQuantity(product.Quantity)
	if errorProduct != nil {
		return errorProduct
	}
	errorProduct = ValidateCodeValue(product.Code_value)
	if errorProduct != nil {
		return errorProduct
	}
	errorProduct = ValidateDateExpirate(product.Expiration)
	if errorProduct != nil {
		return errorProduct
	}

	return nil

}

func ValidateId(id int) error {

	var err error

	if id == 0 {
		err = errors.New("Id is invalid")
	}
	return err
}

func ValidatePrice(price float64) error {
	var err error

	if price == 0.0 {
		err = errors.New("price is invalid")
	}
	return err
}

func ValidateName(name string) error {
	var err error

	if name == "" {
		err = errors.New("name is invalid")
	}
	return err

}

func ValidateQuantity(quantity int) error {
	var err error

	if quantity == 0 {
		err = errors.New("quantity is invalid")
	}
	return err

}

func ValidateCodeValue(codeValue string) error {
	var err error

	if codeValue == "" {
		err = errors.New("code is invalid")
	}
	return err
}

func ValidateDateExpirate(date string) error {
	var err error

	if date == "" {
		err = errors.New("date is invalid")
	}
	return err

}
