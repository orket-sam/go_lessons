package main

import "fmt"

type Bill struct {
	name  string
	items map[string]int
}

func NewBill(name string, items map[string]int) Bill {

	return Bill{name, items}
}

func (b Bill) printBill() {
	fs := fmt.Sprintf("%s '\n", b.name)
	var total int = 0
	for k, v := range b.items {
		fs += fmt.Sprintf("%-20s ...... %d \n", k, v)
		total += v
	}
	fs += fmt.Sprintf("Total %20d", total)

	fmt.Println(fs)

}
