package database

var productList []Product

type Product struct {
	ID          int     `json:"id"` // it converted into small leter id
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(id int) *Product {
	for _, product := range productList {
		if product.ID == id {
			return &product
		}
	}

	return nil
}

func init() {

	// some demo product

	product1 := Product{
		ID:          1,
		Name:        "Laptop",
		Price:       45000.00,
		Stock:       5,
		Description: "This is a gaming laptop",
	}

	product2 := Product{
		ID:          2,
		Name:        "Mobile Phone",
		Price:       25000.00,
		Stock:       10,
		Description: "This is a smartphone",
	}

	product3 := Product{
		ID:          3,
		Name:        "Headphones",
		Price:       5000.00,
		Stock:       15,
		Description: "This is a wireless headphone",
	}

	productList = append(productList, product1, product2, product3)

}
