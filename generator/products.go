package generator

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/robinhuiser/fca-emu/ent"
	"github.com/robinhuiser/fca-emu/ent/product"
)

func populateProducts(ctx context.Context, client *ent.Client, f *gofakeit.Faker) (map[string]int, error) {

	products := []struct {
		productName        string
		productType        string
		productTypeName    string
		productSubType     string
		productSubTypeName string
	}{
		{"plan401ka", "INVESTMENT", "Investment Account", "401k", "401k Plan"},
		{"deposit111", "DEPOSIT", "Deposit One Account", "DD01", "Direct 1 Deposit Account"},
		{"deposit112", "DEPOSIT", "Deposit Two Account", "DD02", "Direct 2 Deposit Account"},
		{"deposit113", "DEPOSIT", "Deposit Three Account", "DD03", "Direct 3 Deposit Account"},
		{"deposit114", "DEPOSIT", "Deposit Four Account", "DD04", "Direct 4 Deposit Account"},
	}

	// Create products
	productTypes := map[string]int{}

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
			return nil, fmt.Errorf("failed creating product: %w", err)
		}
		productTypes[p.productType] = productTypes[p.productType] + 1
	}
	return productTypes, nil
}
