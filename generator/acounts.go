package generator

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/robinhuiser/fca-emu/ent"
	"github.com/robinhuiser/fca-emu/ent/product"
	"github.com/robinhuiser/fca-emu/ent/routingnumber"
)

// Account Type to Product map
var atp = map[string]string{
	"DDA": "CHECKING",
	"NOW": "SAVING",
	"MMA": "INVESTMENT",
}

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

	// Retrieve available products
	act := randomAccountType(f)
	pr, err := client.Product.Query().Where(product.TypeEQ(product.Type(atp[act]))).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving available products: %w", err)
	}

	a, err := client.Account.
		Create().
		SetType(act).
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
		SetBranch(b[f.Number(0, len(b)-1)]).
		SetProduct(pr[f.Number(0, len(pr)-1)]).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating account: %w", err)
	}

	if as == "CLOSED" {
		a.Update().SetDateClosed(dlc).Save(ctx)
	}

	// Add the owner
	a.Update().AddOwners(e).Save(ctx)

	// Add one or more preferences
	for i := 0; i < f.Number(1, 8); i++ {
		pref, err := client.Preference.
			Create().
			SetName(f.Noun()).
			SetValue(f.Word()).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed creating preference: %w", err)
		}
		a.Update().AddPreferences(pref).Save(ctx)
	}

	// Add one or more routing numbers
	for i := 0; i < f.Number(1, 2); i++ {
		rtn, err := client.RoutingNumber.
			Create().
			SetNumber(f.AchRouting()).
			SetType(randomRoutingNumberType(f)).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed creating routing number: %w", err)
		}
		a.Update().AddRoutingnumbers(rtn).Save(ctx)
	}

	return a, nil
}

func randomAccountType(f *gofakeit.Faker) string {
	keys := reflect.ValueOf(atp).MapKeys()
	return keys[rand.Intn(len(keys))].String()
}

func randomRoutingNumberType(f *gofakeit.Faker) routingnumber.Type {
	rtnt := []string{"WIRE", "ABA"}
	return routingnumber.Type(rtnt[f.Number(0, len(rtnt)-1)])
}

func randomAccountTitle(f *gofakeit.Faker) string {
	at := []string{"401k Plan", "Savings", "Deposit"}
	return at[f.Number(0, len(at)-1)]
}

func randomAccountStatus(f *gofakeit.Faker) string {
	as := []string{"OPEN", "CLOSED", "BLOCKED"}
	return as[f.Number(0, len(as)-1)]
}

func randomAccountNumber(f *gofakeit.Faker) string {
	return strconv.Itoa(f.Number(10000, 99999)) +
		strconv.Itoa(f.Number(100000, 999999))
}

// func randomAccountProduct(at string) string {

// }
