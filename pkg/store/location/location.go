package location

import (
	"fmt"
	"math"
)

type Shop struct {
	ID          int
	Name        string
	Contacts    string
	Description string
	ShopsName   string
	Lat         float64
	Long        float64
}

type Location struct {
	Lat  float64
	Long float64
}

func main() {
	nearestShop := SearchOnLocation()
	if nearestShop != nil {
		fmt.Println(nearestShop.ID, nearestShop.Name, nearestShop.Contacts, nearestShop.Description)
	} else {
		fmt.Println("Магазины не найдены")
	}
}

func SearchOnLocation() *Shop {
	myLocation := Location{
		Lat:  47.766657,
		Long: 29.005876,
	}

	shops := []Shop{
		{
			ID:          1,
			Name:        "Мастерок",
			Contacts:    "контакты 1",
			Description: "описание 1",
			ShopsName:   "masterok",
			Lat:         47.767567,
			Long:        29.004731,
		},
		{
			ID:          2,
			Name:        "Фарба",
			Contacts:    "контакты 2",
			Description: "описание 2",
			ShopsName:   "farba",
			Lat:         47.766095,
			Long:        29.006616,
		},
	}

	var nearestShop *Shop
	minDistance := math.MaxFloat64

	for _, shop := range shops {
		distance := calculateDistance(myLocation, Location{Lat: shop.Lat, Long: shop.Long})
		if distance < minDistance {
			minDistance = distance
			nearestShop = &shop
		}
	}

	return nearestShop
}

func calculateDistance(location1, location2 Location) float64 {
	latDiff := location2.Lat - location1.Lat
	longDiff := location2.Long - location1.Long
	return math.Sqrt(latDiff*latDiff + longDiff*longDiff)
}
