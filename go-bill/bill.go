package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

func (b bill) format() string {
	formattedString := "Bill Breakdown: \n"
	var total float64 = 0

	// list items

	for key, value := range b.items {
		formattedString += fmt.Sprintf("%-25v ...$%.2f \n", key+":", value)
		total += value
	}

	// add tip

	formattedString += fmt.Sprintf("%-25v ...$%.2f \n", "tip:", b.tip)

	// total

	formattedString += fmt.Sprintf("%-25v ...$%.2f \n", "Total:", total+b.tip)

	return formattedString
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}
