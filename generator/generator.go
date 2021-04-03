package generator

import (
	"context"
	"fmt"
	"log"

	"github.com/robinhuiser/finite-mock-server/ent"
)

func Accounts(n int, client *ent.Client) error {

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
			a, err := generateRandomAccount(context.Background(), client)
			if err != nil {
				return err
			}
			log.Printf("created account with Id %s", a.ID)
		}
	}
	return nil
}
