package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tealeg/xlsx"

	_ "github.com/lib/pq"
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
}

func (c *Store) AddProducts() {
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

			_, err = c.db.Exec("INSERT INTO products (name, manufacturer, quantity) VALUES ($1, $2, $3)", product.Name, product.Manufacturer, product.Quantity)
			if err != nil {
				log.Printf("Ошибка при вставке данных в базу данных: %s", err)
			}

		}
	}
	fmt.Println("Данные успешно сохранены в базу данных.")
}
