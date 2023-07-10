package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/badimalex/goshop/config"
	"github.com/badimalex/goshop/pkg/database"
	"github.com/tealeg/xlsx"
)

type Product struct {
	Name         string
	Manufacturer string
	Quantity     float64
	Price        float64
	Shops_name   string
}

func main() {
	fileName := flag.String("f", "", "Имя файла XLSX")
	flag.Parse()

	if *fileName == "" {
		log.Fatal("Необходимо указать имя файла с помощью флага -f")
	}

	// Извлекаем имя файла без расширения
	shop := removeExtension(*fileName)

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("TRUNCATE TABLE products") // Очистить таблицу перед вставкой
	if err != nil {
		log.Printf("ошибка при очистке таблицы: %s", err)
	}

	xlFile, err := xlsx.OpenFile(*fileName)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}

	for _, sheet := range xlFile.Sheets {
		for rowNum, row := range sheet.Rows {
			cells := row.Cells

			if len(cells) < 4 {
				log.Printf("ошибка: недостаточно ячеек в строке %d", rowNum+1)
				continue
			}

			name := cells[0].String()
			manufacturer := cells[1].String()
			quantity, err := cells[2].Float()
			if err != nil {
				log.Printf("ошибка при чтении количества: %s", err)
				continue
			}
			price, err := cells[3].Float()
			if err != nil {
				log.Printf("ошибка при чтении цены: %s", err)
				continue
			}

			product := &Product{
				Name:         name,
				Manufacturer: manufacturer,
				Quantity:     quantity,
				Price:        price,
				Shops_name:   shop,
			}

			_, err = db.Exec("INSERT INTO products (name, manufacturer, quantity, price, shops_name) VALUES ($1, $2, $3, $4, $5)", product.Name, product.Manufacturer, product.Quantity, product.Price, product.Shops_name)
			if err != nil {
				log.Printf("ошибка при вставке данных в базу данных: %s", err)
			}
		}
	}

	fmt.Println("Данные успешно сохранены в базу данных!!!.")
}

func removeExtension(fileName string) string {
	return filepath.Base(fileName[:len(fileName)-len(filepath.Ext(fileName))])
}
