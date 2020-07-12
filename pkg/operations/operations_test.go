package operations

import (
	"github.com/lozovoya/gohomework6_1/pkg/card"
	"reflect"
	"testing"
)

func TestSortTransactions(t *testing.T) {
	type args struct {
		service    *card.Service
		cardnumber string
	}

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
		{
			name: "Wrong Card",
			args: args{
				service:    &bankserv,
				cardnumber: "5106 2111 1111 116",
			},
			wantErr: card.ErrorCardNotFound,
			wantCard: []*card.Card{
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
		},
	}
	for _, tt := range tests {
		if got := SortTransactions(tt.args.service, tt.args.cardnumber); (!reflect.DeepEqual(tt.args.service.Cards[0], tt.wantCard[0])) && (got == nil) {
			t.Errorf("Error = %v, want %v", got, tt.wantCard)
		}
	}
}
