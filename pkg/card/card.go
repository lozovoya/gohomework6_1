package card

import (
	"errors"
	"fmt"
	"time"
)

type Service struct {
	BankName string
	Cards    []*Card
}

type Transaction struct {
	Id     int64
	Amount int64
	Type   string
	Mcc    string
	Status string
	TrTime int64
}

type Card struct {
	id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Icon         string
	Number       string
	Transactions []*Transaction
}

func NewService(bankname string) *Service {
	return &Service{BankName: bankname}
}

var ErrorGeneral = errors.New("General Error")
var ErrorCardNotFound = errors.New("Wrong card number")

func (s *Service) AddCard(issuer string, balance int64, currency string, icon string, number string) (*Card, error) {
	card := &Card{
		id:       int64(len(s.Cards)),
		Issuer:   issuer,
		Balance:  balance,
		Currency: currency,
		Icon:     icon,
		Number:   number,
	}
	s.Cards = append(s.Cards, card)
	return card, nil
}

func (s *Service) AddTransaction(cardnumber string, amount int64, mcc string) error {
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
	return ErrorCardNotFound
}

func (s *Service) PrintTransactions(cardnumber string) error {
	for _, cards := range s.Cards {
		if cards.Number == cardnumber {
			for _, i := range cards.Transactions {
				fmt.Println(i.Amount)
			}
			return nil
		}
	}
	return ErrorCardNotFound
}
