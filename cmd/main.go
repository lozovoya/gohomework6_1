package main

import (
	"fmt"
	"github.com/lozovoya/gohomework6_1/pkg/card"
	"github.com/lozovoya/gohomework6_1/pkg/operations"
	"math/rand"
)

func main() {
	//amounts := []int64{1_000_00, 3_000_00, 7_000_00, 2_000_00, 11_000_00, 1_000_00}

	amounts := make([]int64, 100)
	penBank := *card.NewService("PengiunBank")

	penBank.AddCard("visa", 1_000_00, "rub", "http://...", "5106 2111 1111 1116")
	penBank.AddCard("visa", 1_000_000_00, "rub", "http://...", "5106 2100 0000 0003")

	for index := range amounts {
		amounts[index] = int64(rand.Intn(1_000_000))
	}

	for _, amount := range amounts {
		err := penBank.AddTransaction("5106 2111 1111 1116", amount, "1234")
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("перед сортировкой:")
	penBank.PrintTransactions("5106 2111 1111 1116")

	err := operations.SortTransactions(&penBank, "5106 2111 1111 1116")
	if err == nil {
		fmt.Println("после сортировки:")
		penBank.PrintTransactions("5106 2111 1111 1116")
	} else {
		fmt.Println(err)
	}

}
