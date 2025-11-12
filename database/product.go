package database

var ProductList []IProduct

type IProduct struct {
	ID          int     `json:"id"` // it converted into small leter id
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

func init() {

	// some demo product

	product1 := IProduct{
		ID:          1,
		Name:        "Laptop",
		Price:       45000.00,
		Stock:       5,
		Description: "This is a gaming laptop",
	}

	product2 := IProduct{
		ID:          2,
		Name:        "Mobile Phone",
		Price:       25000.00,
		Stock:       10,
		Description: "This is a smartphone",
	}

	product3 := IProduct{
		ID:          3,
		Name:        "Headphones",
		Price:       5000.00,
		Stock:       15,
		Description: "This is a wireless headphone",
	}

	ProductList = append(ProductList, product1, product2, product3)

}
