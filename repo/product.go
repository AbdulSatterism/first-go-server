package repo

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	// Update(productId int) (*Product, error)
	// Delete(productId int) error
}

type productRepo struct {
	productList []*Product
}

// constructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productId int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productId {
			return product, nil
		}
	}

	return nil, nil
}

// func (r *productRepo) Update(productId int) (*Product, error) {

// }

// func (r *productRepo) Delete(productId int) error {

// }

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func generateInitialProducts(r *productRepo) {

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

	r.productList = append(r.productList, &product1, &product2, &product3)

}
