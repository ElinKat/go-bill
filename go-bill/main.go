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

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b

}

func promptOptions(b bill) {

	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("a - add item, s - save bill, t - add tip: ", reader)

	switch option {

	case "a":
		name, _ := getInput("item name: ", reader)
		price, _ := getInput("item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("item added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added - ", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you saved the file ", b.name)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}

}

func main() {

	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill)

	// myBill := newBill("John's Bill")
	// myBill.addItem("Pizza", 12.99)
	// myBill.addItem("Burger", 8.99)
	// myBill.updateTip(2.50)
	// fmt.Println(myBill.format())

}
