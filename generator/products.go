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
		{"SAV01", product.TypeSAVING.String(), "Savings Account", "SAV-BASIC", "Basic Savings"},
		{"SAV02", product.TypeSAVING.String(), "Savings Account", "SAV-COL", "College Savings"},
		{"SAV03", product.TypeSAVING.String(), "Savings Account", "SAV-TAX", "Tax Savings"},
		{"CHK01", product.TypeCHECKING.String(), "Checking Account", "CHK-BASIC", "Basic Checking"},
		{"CHK02", product.TypeCHECKING.String(), "Checking Account", "CHK-BRONZE", "Bronze Checking"},
		{"CHK03", product.TypeCHECKING.String(), "Checking Account", "CHK-SILVER", "Silver Checking"},
		{"CHK04", product.TypeCHECKING.String(), "Checking Account", "CHK-GOLD", "Gold Checking"},
		{"INV01", product.TypeINVESTMENT.String(), "Investment Account", "INV-BASIC", "Basic Investments"},
		{"INV02", product.TypeINVESTMENT.String(), "Investment Account", "INV-BUFFET", "Warren's Investments"},
		{"INV03", product.TypeINVESTMENT.String(), "Investment Account", "INV-ACKMAN", "Bill's Investments"},
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
