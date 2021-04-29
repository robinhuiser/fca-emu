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
	"github.com/robinhuiser/fca-emu/ent/account"
	"github.com/robinhuiser/fca-emu/ent/product"
	"github.com/robinhuiser/fca-emu/ent/routingnumber"
)

// Account Type to Product map
var atp = map[string]string{
	"DDA": product.TypeCHECKING.String(),
	"NOW": product.TypeSAVING.String(),
	"MMA": product.TypeINVESTMENT.String(),
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

	// Get random account type with strong type preference
	act := randomAccountType(f, "DDA")

	// Retrieve product linked to account type
	pr, err := client.Product.Query().Where(product.TypeEQ(product.Type(atp[act]))).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving product type: %w", err)
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

	if as == account.StatusCLOSED {
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

func randomAccountType(f *gofakeit.Faker, p string) string {
	at := p
	// Return in 50% of the times the preferred account type p
	if f.Bool() {
		at = p
	} else {
		keys := reflect.ValueOf(atp).MapKeys()
		at = keys[rand.Intn(len(keys))].String()
	}
	return at
}

func randomRoutingNumberType(f *gofakeit.Faker) routingnumber.Type {
	rtnt := []string{"WIRE", "ABA"}
	return routingnumber.Type(rtnt[f.Number(0, len(rtnt)-1)])
}

func randomAccountTitle(f *gofakeit.Faker) string {
	at := []string{"401k Plan", "Savings", "Deposit"}
	return at[f.Number(0, len(at)-1)]
}

func randomAccountStatus(f *gofakeit.Faker) account.Status {
	as := []account.Status{account.StatusOPEN, account.StatusCLOSED, account.StatusBLOCKED}
	return as[f.Number(0, len(as)-1)]
}

func randomAccountNumber(f *gofakeit.Faker) string {
	return strconv.Itoa(f.Number(10000, 99999)) +
		strconv.Itoa(f.Number(100000, 999999))
}
