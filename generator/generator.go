package generator

import (
	"context"
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/robinhuiser/finite-mock-server/ent"
)

func Accounts(n int, client *ent.Client) error {

	faker := gofakeit.New(0)

	// Skip bank and branches generation if db already contains records
	b, err := client.Bank.
		Query().
		Limit(1).
		All(context.Background())
	if err != nil {
		return fmt.Errorf("could not get banks: %w", err)
	}

	if len(b) == 0 {
		if err := populateBanks(context.Background(), client, faker); err != nil {
			return err
		}
	}

	// Skip account generation if db already contains records
	a, err := client.Account.
		Query().
		Limit(1).
		All(context.Background())
	if err != nil {
		return fmt.Errorf("could not get accounts: %w", err)
	}

	if len(a) == 0 {
		for i := 0; i < n; i++ {
			a, err := generateRandomAccount(context.Background(), client, faker)
			if err != nil {
				return err
			}
			log.Printf("created account with Id %s", a.ID)
		}
	}

	return nil
}
