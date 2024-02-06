package main

func main() {

	bill := NewBill("sam's basket", map[string]int{"sunrice": 1454, "brookside": 657, "ramtons fridge": 24800})
	bill.printBill()
}
