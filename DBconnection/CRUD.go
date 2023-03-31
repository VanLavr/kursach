package DBconnection

import "log"

type Product struct {
	ID       int
	Name     string
	Category string
	Price    float64
}

func GetAllProducts() ([]Product, error) {
	var AllProducts []Product

	Rows, readError := DB.Query("SELECT * FROM products")
	if readError != nil {
		log.Printf("[CRUD/GetAllProducts] cannot read rows from database (%v)", readError)
		return nil, readError
	}

	for Rows.Next() {
		var (
			id        int
			name, cat string
			price     float64
		)

		scanError := Rows.Scan(&id, &name, &cat, &price)
		if scanError != nil {
			log.Printf("[CRUD/GetAllProducts] cannot write data into variables (%v)", scanError)
			return nil, scanError
		}

		product := Product{
			ID:       id,
			Name:     name,
			Category: cat,
			Price:    price,
		}

		AllProducts = append(AllProducts, product)
	}

	return AllProducts, nil
}
