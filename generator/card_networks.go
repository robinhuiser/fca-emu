package generator

import (
	"context"
	"fmt"

	"github.com/robinhuiser/fca-emu/ent"
)

func populateCardNetworks(ctx context.Context, client *ent.Client) error {

	c := map[string]string{
		"VISA":       "VISA",
		"MASTERCARD": "Mastercard",
		"AMEX":       "American Express",
		"DC":         "Diners Club",
		"DISCOVER":   "Discover",
		"ENROUTE":    "EnRoute",
		"JCB":        "JCB",
		"VOYAGER":    "Voyager",
	}

	for code, company := range c {
		_, err := client.CardNetwork.Create().
			SetName(company).
			SetCode(code).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating card network: %w", err)
		}
	}

	return nil
}
