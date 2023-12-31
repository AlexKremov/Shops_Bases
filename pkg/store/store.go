package store

import (
	"database/sql"
	"fmt"
	"log"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

type Product struct {
	Name         string
	Manufacturer string
	Quantity     float64
	Price        float64
	Shops_name   string
}

type AllShops struct {
	Name        string
	Contacts    string
	Description string
	ShopsName   string
	Lat         float64
	Long        float64
}

func (s *Store) AddShops(name, contacts, description, shopsName string, lat, long float64) {
	_, err := s.db.Exec("INSERT INTO stores (name, contacts, description, shops_name, lat, long) VALUES ($1, $2, $3, $4, $5, $6)", name, contacts, description, shopsName, lat, long)
	if err != nil {
		log.Printf("Ошибка при вставке данных в базу stores: %s", err)
	}
}

func (s *Store) AddShopsSeeds() {
	s.AddShops("Мастерок", "контакты 1", "описание 1", "masterok", 47.767567, 29.004731)
	s.AddShops("Фарба", "контакты 2", "описание 2", "farba", 47.766095, 29.006616)
}

func (s *Store) SearchByName(keyword string) ([]*Product, error) {
	query := `
		SELECT name, manufacturer, quantity, price, shops_name
		FROM products
		WHERE name ILIKE '%' || $1 || '%'
	`

	rows, err := s.db.Query(query, keyword)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении SQL-запроса: %s", err)
	}
	defer rows.Close()

	var products []*Product

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Name, &product.Manufacturer, &product.Quantity, &product.Price, &product.Shops_name)
		if err != nil {
			return nil, fmt.Errorf("ошибка при сканировании результатов запроса: %s", err)
		}
		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при получении результатов запроса: %s", err)
	}

	return products, nil
}
