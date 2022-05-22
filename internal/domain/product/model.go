package product

type Product struct {
	Uuid        string `json:"uuid"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Currency    string `json:"currency"`
	LeftInStock string `json:"left_in_stock"`
}
