package main

import (
	"fmt"
	"os"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
	total float64
}

func newBill(name string) Bill {
	b := Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
		total: 0,
	}
	return b
}

func (b *Bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *Bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}
	total += b.tip
	fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)
	fs += fmt.Sprintf("%-25v ...$%v \n", "total:", total)
	return fs
}

func (b *Bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *Bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Bill saved successfully!")
	}
}
