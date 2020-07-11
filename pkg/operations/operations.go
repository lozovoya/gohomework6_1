package operations

import (
	"github.com/lozovoya/gohomework6_1/pkg/card"
	"time"
)

func (s *card.Service)(cardnumber string) ([]*card.Transaction, error) {
	for _, cards := range s.Cards {
		if cards.Number == cardnumber {
			transaction := &Transaction{
				Id:     0,
				Amount: amount,
				Mcc:    mcc,
				TrTime: time.Now().Unix(),
			}
			cards.Transactions = append(cards.Transactions, transaction)
			return nil
		}
	}
}
