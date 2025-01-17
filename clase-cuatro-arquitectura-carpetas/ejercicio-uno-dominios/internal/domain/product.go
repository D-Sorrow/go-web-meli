package domain

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

func NewProduct(id int, name string, quantity int, codeValue string, isPublished bool, dateExpiration string, price float64) *Product {
	return &Product{
		Id:           id,
		Name:         name,
		Quantity:     quantity,
		Code_value:   codeValue,
		Is_published: isPublished,
		Expiration:   dateExpiration,
		Price:        price,
	}
}
