package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)
	return b
}

func promptOptions(b Bill) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch input {
	case "a", "A":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		parsedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be valid!")
			promptOptions(b)
		}

		b.addItem(name, parsedPrice)

		fmt.Println("item added -", name, price)
		promptOptions(b)
	case "s", "S":
		b.save()
		fmt.Println("Bill saved to file -", b.name+".txt")
	case "t", "T":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		parsedTip, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Tip must be valid!")
			promptOptions(b)
		}
		b.updateTip(parsedTip)

		fmt.Println("Tip updated to -", parsedTip)
		promptOptions(b)
	default:
		fmt.Println("Invalid Option!")
		promptOptions(b)
	}

}

func main() {
	bill := createBill()
	promptOptions(bill)
}
