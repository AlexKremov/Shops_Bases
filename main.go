package main

import (
	"Dima-project-ryb/pkg/store"
	"fmt"
	"log"

	"github.com/badimalex/goshop/config"
	"github.com/badimalex/goshop/pkg/database"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	s := store.New(db)

	s.AddShops("Мастерок", "контакты 1", "описание 1", "masterok", 47.767567, 29.004731)
	s.AddShops("Фарба", "контакты 2", "описание 2", "farba", 47.766095, 29.006616)

	// s.AddProducts()
	masterok := "masterok.xlsx"
	farba := "farba.xlsx"

	err = s.InsertDataFromXLSX(masterok, "masterok")
	if err != nil {
		log.Fatalf("ошибка при вставке данных из файла: %s", err)
	}

	err = s.InsertDataFromXLSX(farba, "farba")
	if err != nil {
		log.Fatalf("ошибка при вставке данных из файла: %s", err)
	}

	products, err := s.SearchByName("ремень")
	if err != nil {
		log.Printf("ошибка при выполнении поиска: %s", err)
	}

	for _, product := range products {
		fmt.Printf("Название: %s\n", product.Name)
		fmt.Printf("Производитель: %s\n", product.Manufacturer)
		fmt.Printf("Количество: %.2f\n", product.Quantity)
		fmt.Printf("Цена: %.2f\n", product.Price)
		fmt.Printf("Магазин: %s\n", product.Shops_name)
		fmt.Println()
	}
}
