package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sort"
)

// Given a max budget and a list of sales items implement a function that outputs the sorted list of deals within budget.

const path = "items.json"

// SaleItem represent the item part of the big sale
type SaleItem struct {
	Name           string  `json:"name"`
	OriginalPrice  float64 `json:"originalPrice"`
	ReducedPrice   float64 `json:"reducedPrice"`
	SalePercentage float64
}

// matchSales adds the sales procentage of the item
// and sorts the array accordingly.
func matchSales(budget float64, items []SaleItem) []SaleItem {
	var sales []SaleItem
	for _, item := range items {
		if budget >= item.ReducedPrice {
			item.SalePercentage = (item.OriginalPrice - item.ReducedPrice) / item.OriginalPrice * 100
			sales = append(sales, item)
		}
	}
	sort.Slice(sales, func(i, j int) bool {
		return sales[i].SalePercentage > sales[j].SalePercentage
	})

	return sales
}

func importData() []SaleItem {
	file, err := os.ReadFile("items.json")
	if err != nil {
		log.Fatal(err)
	}

	var data []SaleItem
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func printItems(items []SaleItem) {
	log.Println("The BIG sale has started with our amazing offers!")
	if len(items) == 0 {
		log.Println("No items found.:( Try increasing your budget.")
	}

	for i, r := range items {
		log.Printf("[%d]: %s is %.2f%% OFF! Get it now for JUST %.2f!\n", i, r.Name, r.SalePercentage, r.ReducedPrice)
	}
}

func main() {
	budget := flag.Float64("budget", 0.0, "The max budget you want to shop with.")
	flag.Parse()
	items := importData()
	matchedItems := matchSales(*budget, items)
	printItems(matchedItems)
}
