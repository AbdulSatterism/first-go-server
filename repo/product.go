package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Price       float64 `json:"price" db:"price"`
	Stock       int     `json:"stock" db:"stock"`
	Description string  `json:"description" db:"description"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	// Update(productId int) (*Product, error)
	// Delete(productId int) error
}

type productRepo struct {
	db *sqlx.DB
}

// constructor function
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}

}

func (r *productRepo) Create(p Product) (*Product, error) {

	query := `Insert into products (name,price,stock,description) values($1,$2,$3,$4) returning id`

	row := r.db.QueryRow(query, p.Name, p.Price, p.Stock, p.Description)

	err := row.Scan(&p.ID)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &p, nil

}

func (r *productRepo) Get(productId int) (*Product, error) {
	fmt.Println("Getting product with ID:", productId) // Debugging line

	return &Product{
		ID:          productId,
		Name:        "Sample Product",
		Price:       0.0,
		Stock:       0,
		Description: "This is a sample product",
	}, nil
}

// func (r *productRepo) Update(productId int) (*Product, error) {

// }

// func (r *productRepo) Delete(productId int) error {

// }

func (r *productRepo) List() ([]*Product, error) {
	return nil, nil
}
