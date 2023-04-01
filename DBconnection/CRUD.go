package DBconnection

import (
	"log"
	"errors"
)

type Product struct {
	ID       int
	Name     string
	Category string
	Price    float64
}

func GetAllIDs() ([]int, error) {
	var AllIDs []int

	Rows, readError := DB.Query("SELECT id FROM products")
	if readError != nil {
		log.Printf("[CRUD/GetAllIDs] cannot read rows from database (%v)", readError)
		return nil, readError
	}

	for Rows.Next() {
		var idFromdB int

		scanError := Rows.Scan(&idFromdB)
		if scanError != nil {
			log.Printf("[CRUD/GetAllIDs] cannot write data into variables (%v)", scanError)
			return nil, scanError
		}

		AllIDs = append(AllIDs, idFromdB)
	}

	return AllIDs, nil
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

func GetProductById(idOfProd int) (Product, error) {
	var prod []Product
	emptyProduct := new(Product)

	Rows, readErr := DB.Query("SELECT * FROM products WHERE id = ?", idOfProd)
	if readErr != nil {
		log.Printf("[CRUD/GetProductById] cannot read rows from database (%v)", readErr)
		return *emptyProduct, readErr
	}
	defer Rows.Close()

	for Rows.Next() {
		var (
			id        int
			name, cat string
			price     float64
		)

		scanError := Rows.Scan(&id, &name, &cat, &price)
		if scanError != nil {
			log.Printf("[CRUD/GetProductById] cannot read rows from database (%v)", scanError)
			return *emptyProduct, scanError
		}

		product := Product{
			ID:       id,
			Name:     name,
			Category: cat,
			Price:    price,
		}

		prod = append(prod, product)
	}

	return prod[0], nil
}

func CreateItem(item Product) error {
	lastIDarr, getErr := GetAllIDs()
	if getErr != nil {
		log.Printf("[CRUD/CreateItem] cannot get all ids (%v)", getErr)
		return getErr
	}
	lastID := lastIDarr[len(lastIDarr)-1]
	lastID += 1

	_, insertionError := DB.Query("INSERT INTO products(id, NameOfProduct, Category, Price) VALUES(?, ?, ?, ?)", lastID, item.Name, item.Category, item.Price)
	if insertionError != nil {
		log.Printf("[CRUD/CreateItem] cannot insert data (%v)", insertionError)
		return insertionError
	}

	return nil
}

func DeleteItem(itemID int) error {
	ids, idErr := GetAllIDs()
	if idErr != nil {
		log.Printf("[CRUD/DeleteItem] cannot read IDs (%v)", idErr)
		return idErr
	}

	flag := true
	for _, v := range ids {
		if v == itemID {
			flag = false
		}
	}
	if flag {
		return errors.New("There is no such id")
	}

	_, deletionError := DB.Query("DELETE FROM products WHERE ID = ?", itemID)
	if deletionError != nil {
		log.Printf("[CRUD/DeleteItem] cannot insert data (%v)", deletionError)
		return deletionError
	}

	return nil
}