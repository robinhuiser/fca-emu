package generator

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/robinhuiser/fca-emu/ent"
	"github.com/robinhuiser/fca-emu/ent/transaction"
)

func populateRandomTransactions(ctx context.Context, client *ent.Client, f *gofakeit.Faker, a *ent.Account) ([]*ent.Transaction, error) {

	trns := []*ent.Transaction{}

	// Define the number of transactions per month per account type
	nrTrans := map[string]int{
		"SAVING":     2,  // Reg-D limit
		"LOAN":       1,  // Once a month downpayment, 2x per year extra payment
		"DEPOSIT":    2,  // Same as savings
		"CHECKING":   72, // Includes debit, credit, or prepaid cards, cash (ATM), paper checks, electronic
		"INVESTMENT": 2,  // Same as savings
	}

	pt, err := a.QueryProduct().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving account product type: %w", err)
	}

	for tr := 0; tr < nrTrans[pt.Type.String()]; tr++ {

		// Some transaction stats
		dir := transaction.DirectionCREDIT
		amnt := f.Price(-3000, 4000)
		if amnt < 0 {
			dir = transaction.DirectionDEBIT
		}

		t, err := client.Transaction.
			Create().
			SetStatus("POSTED").
			SetExecutedAmount(amnt).
			SetExecutedCurrencyCode(a.CurrencyCode).
			SetDirection(dir).
			SetRunningBalance(float64(a.CurrentBalance) + amnt).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed creating transaction: %w", err)
		}

		trns = append(trns, t)
	}

	return trns, nil
}
