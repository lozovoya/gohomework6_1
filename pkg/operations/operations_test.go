package operations

import (
	"fmt"
	"github.com/lozovoya/gohomework6_1/pkg/card"
	"reflect"
	"testing"
)

func TestSortTransactions(t *testing.T) {
	type args struct {
		service    *card.Service
		cardnumber string
	}

	//testCard := []card.Card{
	//	{
	//		Number: "5106 2111 1111 1116",
	//		Transactions: []*card.Transaction{
	//			{
	//				Id:     1,
	//				Amount: 1_000,
	//			},
	//			{
	//				Id:     2,
	//				Amount: 2_000,
	//			},
	//			{
	//				Id:     3,
	//				Amount: 3_000,
	//			},
	//		},
	//	},
	//}

	bankserv := card.Service{
		BankName: "test bank",
		Cards: []*card.Card{
			{
				Number: "5106 2111 1111 1116",
				Transactions: []*card.Transaction{
					{
						Id:     1,
						Amount: 1_000,
					},
					{
						Id:     2,
						Amount: 2_000,
					},
					{
						Id:     3,
						Amount: 3_000,
					},
				},
			},
		},
	}
	//bankserv.AddCard("visa", 1_000_00, "rub", "http://...", "5106 2111 1111 1116")
	//bankserv.AddTransaction("5106 2111 1111 1116", 1_00)

	tests := []struct {
		name     string
		args     args
		wantErr  error
		wantCard []*card.Card
	}{
		{
			name: "Sorted well",
			args: args{
				service:    &bankserv,
				cardnumber: "5106 2111 1111 1116",
			},
			wantErr: nil,
			wantCard: []*card.Card{
				{
					Number: "5106 2111 1111 1116",
					Transactions: []*card.Transaction{
						{
							Id:     3,
							Amount: 3_000,
						},
						{
							Id:     2,
							Amount: 2_000,
						},
						{
							Id:     1,
							Amount: 1_000,
						},
					},
				},
			},
		},
		//{
		//	name:    "Wrong Card",
		//	args:    args{
		//		service:    &bankserv,
		//		cardnumber: "5106 2111 1111 1116",
		//	},
		//	wantErr: card.ErrorCardNotFound,
		//	wantCard: []*card.Card{
		//		{
		//			Number: "5106 2111 1111 1116",
		//			Transactions: []*card.Transaction{
		//				{
		//					Id:     1,
		//					Amount: 1_000,
		//				},
		//				{
		//					Id:     2,
		//					Amount: 2_000,
		//				},
		//				{
		//					Id:     3,
		//					Amount: 3_000,
		//				},
		//			},
		//		},
		//	},
		//},
	}

	for _, tt := range tests {
		//s:= card.PrintTransactions(tt.args.service.Cards)
		fmt.Println(reflect.DeepEqual(tt.args.service.Cards[0].Transactions, tt.wantCard[0].Transactions))
		if got := SortTransactions(tt.args.service, tt.args.cardnumber); !reflect.DeepEqual(tt.args.service.Cards, tt.wantCard) {
			fmt.Println(tt.args.service.Cards[0].Transactions, &tt.wantCard[0].Transactions)
			t.Errorf("Sum() = %v, want %v", got, tt.wantCard)
		}

		//t.Run(tt.name, func(t *testing.T) {
		//if err := SortTransactions(tt.args.service, tt.args.cardnumber); (err != nil) != tt.wantErr {
		//	t.Errorf("SortTransactions() error = %v, wantErr %v", err, tt.wantErr)
		//}

	}
}
