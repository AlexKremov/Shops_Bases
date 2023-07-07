package main

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx"

	"github.com/badimalex/goshop/config"
	"github.com/badimalex/goshop/pkg/database"

	_ "github.com/lib/pq"
)

type Product struct {
	Name         string
	Manufacturer string
	Quantity     float64
}

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

	excelFileName := "TestTovar.xlsx"

	// Открытие файла Excel
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %s\n", err)
		return
	}

	// Перебор всех листов в файле
	for _, sheet := range xlFile.Sheets {
		// Перебор всех строк в листе
		for rowNum, row := range sheet.Rows {
			// Чтение данных из ячеек
			cells := row.Cells

			// Проверка длины среза ячеек
			if len(cells) < 3 {
				log.Printf("Ошибка: недостаточно ячеек в строке %d", rowNum+1)
				continue
			}

			name := cells[0].String()
			manufacturer := cells[1].String()
			quantity, err := cells[2].Float()
			if err != nil {
				log.Printf("Ошибка при чтении цены: %s", err)
				continue // Пропускаем строки с некорректной ценой
			}

			product := &Product{
				Name:         name,
				Manufacturer: manufacturer,
				Quantity:     quantity,
			}

			_, err = db.Exec("INSERT INTO products (name, manufacturer, quantity) VALUES ($1, $2, $3)", product.Name, product.Manufacturer, product.Quantity)
			if err != nil {
				log.Printf("Ошибка при вставке данных в базу данных: %s", err)
			}

		}
	}
	fmt.Println("Данные успешно сохранены в базу данных.")
}
