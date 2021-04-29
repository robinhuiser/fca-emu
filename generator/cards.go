package generator

import (
	"context"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/robinhuiser/fca-emu/ent"
	"github.com/robinhuiser/fca-emu/ent/card"
)

func populateRandomCards(ctx context.Context, client *ent.Client, f *gofakeit.Faker, a *ent.Account) ([]*ent.Card, error) {
	crds := []*ent.Card{}
	cardType := [3]string{card.TypeCREDIT.String(), card.TypeDEBIT.String(), card.TypeLOYALTY.String()}

	cn, err := client.CardNetwork.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving card networks: %w", err)
	}

	acco, err := a.QueryOwners().First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving account owners: %w", err)
	}

	// Create between 0 and 3 cards for each account
	for crd := 0; crd < f.Number(0, 3); crd++ {
		cardStartDate := f.DateRange(a.DateOpened, time.Now())
		cardExpiryDate := f.DateRange(cardStartDate, time.Now().AddDate(2, 0, 0))
		cardStatus := card.StatusOPERATIONAL

		if cardExpiryDate.Before(time.Now()) {
			cardStatus = card.StatusLOCKED
		}

		c, err := client.Card.
			Create().
			SetAccount(a).
			SetType(card.Type(cardType[f.Number(0, len(cardType)-1)])).
			SetNetwork(cn[f.Number(0, len(cn)-1)]).
			SetNumber(f.CreditCardNumber(nil)).
			SetStartDate(cardStartDate).
			SetExpiryDate(cardExpiryDate).
			SetHolderName(acco.Fullname).
			SetStatus(cardStatus).
			SetURL("https://my.cdn.com/imgs/card/" + f.DigitN(18) + ".jpg").
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed creating card: %w", err)
		}
		crds = append(crds, c)
	}

	return crds, nil
}
