package generator

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/robinhuiser/fca-emu/ent"
	"github.com/robinhuiser/fca-emu/ent/product"
)

func populateProducts(ctx context.Context, client *ent.Client, f *gofakeit.Faker) error {

	products := []struct {
		productName        string
		productType        string
		productTypeName    string
		productSubType     string
		productSubTypeName string
	}{
		{"plan401ka", "INVESTMENT", "Investment Account", "401k", "401k Plan"},
		{"deposit112", "DEPOSIT", "Deposit Account", "DD01", "Direct Deposit Account"},
	}

	// Create products
	for _, p := range products {
		_, err := client.Product.
			Create().
			SetName(p.productName).
			SetType(product.Type(p.productType)).
			SetTypeName(p.productTypeName).
			SetSubType(p.productSubType).
			SetSubTypeName(p.productSubTypeName).
			SetURL("https://my.cdn.com/imgs/product/" + f.DigitN(18) + ".jpg").
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating product: %w", err)
		}
	}
	return nil
}
