package operations

import (
	"github.com/lozovoya/gohomework6_1/pkg/card"
	"sort"
)

func SortTransactions(service *card.Service, cardnumber string) error {
	for _, card := range service.Cards {
		if card.Number == cardnumber {
			sort.SliceStable(card.Transactions, func(i, j int) bool { return card.Transactions[i].Amount > card.Transactions[j].Amount })
			return nil
		}
	}
	return card.ErrorCardNotFound
}
