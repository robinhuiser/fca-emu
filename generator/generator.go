package generator

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/robinhuiser/finite-mock-server/ent"
)

func Generate(ents int, branches int, c *ent.Client) error {

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
		var accts = 0
		for entity := 0; entity < ents; entity++ {
			e, err := populateEntity(context.Background(), c, f)
			if err != nil {
				return err
			}
			log.Printf("created entity (%s) with id %s", strings.ToLower(e.Type.String()), e.ID)

			switch a := e.Type.String(); a {
			case "ORGANIZATION":
				accts = f.Number(2, 5)
			case "CORPORATE":
				accts = f.Number(3, 10)
			default:
				accts = f.Number(1, 3)
			}

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
