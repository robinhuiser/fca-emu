package generator

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/robinhuiser/finite-mock-server/ent"
)

func generateRandomAccount(ctx context.Context, client *ent.Client) (*ent.Account, error) {
	a, err := client.Account.
		Create().
		SetType("DDA").
		SetNumber("xxxxxxx1213").
		SetParentId(uuid.New()).
		SetName("Matt Dollar").
		SetTitle("401k Plan").
		SetIban("GB29NWBK60161331926819").
		SetDateCreated(time.Now()).
		SetDateOpened(time.Now()).
		SetDateLastUpdated(time.Now()).
		SetCurrencyCode("USD").
		SetStatus("OPEN").
		SetSource("core").
		SetInterestReporting(true).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating account: %w", err)
	}
	return a, nil
}
