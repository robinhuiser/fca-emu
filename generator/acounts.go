package generator

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/robinhuiser/finite-mock-server/ent"
)

func populateRandomAccount(ctx context.Context, client *ent.Client, f *gofakeit.Faker, e *ent.Entity) (*ent.Account, error) {

	dco := f.DateRange(time.Now().AddDate(-5, 0, -1), time.Now()) // random date create open
	dlc := f.DateRange(dco, time.Now())                           // random date lastupdated, closed
	ra := randomAccountNumber(f)                                  // random account number
	as := randomAccountStatus(f)                                  // account status
	ab := float32(f.Price(0, 10000))                              // account balance

	// Retrieve available branches
	b, err := client.Branch.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving available branches: %w", err)
	}

	a, err := client.Account.
		Create().
		SetType(randomAccountType(f)).
		SetNumber(ra).
		SetName(f.FirstName() + " " + f.LastName()).
		SetTitle(randomAccountTitle(f)).
		SetDateCreated(dco).
		SetDateOpened(dco).
		SetDateLastUpdated(dlc).
		SetCurrencyCode("USD").
		SetStatus(as).
		SetSource("core").
		SetInterestReporting(true).
		SetAvailableBalance(ab).
		SetCurrentBalance(ab).
		SetURL("https://my.cdn.com/imgs/account/" + f.DigitN(18) + ".jpg").
		SetBranch(b[f.Number(0, len(b))]).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating account: %w", err)
	}

	if as == "CLOSED" {
		a.Update().SetDateClosed(dlc).Save(ctx)
	}

	a.Update().AddOwner(e).Save(ctx)

	return a, nil
}

func randomAccountType(f *gofakeit.Faker) string {
	at := []string{"DDA", "NOW", "MMA"}
	return at[f.Number(0, len(at)-1)]
}

func randomAccountTitle(f *gofakeit.Faker) string {
	at := []string{"401k Plan", "Savings", "Deposit"}
	return at[f.Number(0, len(at)-1)]
}

func randomAccountStatus(f *gofakeit.Faker) string {
	at := []string{"OPEN", "CLOSED", "BLOCKED"}
	return at[f.Number(0, len(at)-1)]
}

func randomAccountNumber(f *gofakeit.Faker) string {
	return strconv.Itoa(f.Number(10000, 99999)) +
		strconv.Itoa(f.Number(100000, 999999))
}
