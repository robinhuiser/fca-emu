package generator

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/robinhuiser/fca-emu/ent"
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

		// Generate static data: card networks, banks with branches & products
		if err := populateCardNetworks(context.Background(), c); err != nil {
			return err
		}
		if err := populateBanks(branches, context.Background(), c, f); err != nil {
			return err
		}
		pTypes, err := populateProducts(context.Background(), c, f)
		if err != nil {
			return err
		}
		for pt, n := range pTypes {
			log.Printf("created %d product(s) with type %s", n, pt)
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
			case "BUSINESS":
				accts = f.Number(3, 10)
			case "SYSTEM":
				accts = f.Number(2, 5)
			default:
				accts = f.Number(1, 3)
			}

			// Generate for each entity x number of accounts
			for account := 0; account < accts; account++ {
				a, err := populateRandomAccount(context.Background(), c, f, e)
				if err != nil {
					return err
				}

				// Generate for each account x number of transactions
				t, err := populateRandomTransactions(context.Background(), c, f, a)
				if err != nil {
					return err
				}

				// Generate for each DDA account x number of cards
				cards := []*ent.Card{}
				if a.Type == "DDA" {
					cards, err = populateRandomCards(context.Background(), c, f, a)
					if err != nil {
						return err
					}
				}
				log.Printf("  > added account (%s) with Id %s - %d transactions, %d cards", a.Type, a.ID, len(t), len(cards))
			}
		}
	}

	return nil
}
