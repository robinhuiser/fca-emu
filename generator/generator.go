package generator

import (
	"context"
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/robinhuiser/finite-mock-server/ent"
)

func Generate(ents int, accts int, branches int, c *ent.Client) error {

	f := gofakeit.New(0)

	// Skip data generation if db already contains records
	ec, err := c.Entity.
		Query().
		Limit(1).
		All(context.Background())
	if err != nil {
		return fmt.Errorf("could not get entities: %w", err)
	}

	// If there is not test data...
	if len(ec) == 0 {

		// Generate card networks
		if err := populateCardNetworks(context.Background(), c, f); err != nil {
			return err
		}

		// Generate banks & branches
		if err := populateBanks(branches, context.Background(), c, f); err != nil {
			return err
		}

		// Generate x number of entities
		for entity := 0; entity < ents; entity++ {
			e, err := populateEntity(context.Background(), c, f)
			if err != nil {
				return err
			}
			log.Printf("created entity with Id %s", e.ID)

			// Generate for each entity x number of accounts
			for account := 0; account < accts; account++ {
				a, err := populateRandomAccount(context.Background(), c, f, e)
				if err != nil {
					return err
				}
				log.Printf("  > added account with Id %s", a.ID)
			}
		}
	}

	return nil
}
