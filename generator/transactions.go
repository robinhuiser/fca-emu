package generator

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/robinhuiser/fca-emu/ent"
	"github.com/robinhuiser/fca-emu/ent/product"
	"github.com/robinhuiser/fca-emu/ent/transaction"
	"github.com/robinhuiser/fca-emu/util"
)

const (
	CHECK_IMAGE_CREATE_ERROR         = "failed creating check image"
	TRANSACTION_CREATE_ERROR         = "failed creating transaction"
	ACCOUNT_OWNER_ADDRESS_ERROR      = "failed retrieving account owners address"
	ACCOUNT_OWNER_TYPE_ERROR         = "failed retrieving account owners type"
	ACCOUNT_OWNER_PRODUCT_TYPE_ERROR = "failed retrieving account product type"
)

func populateRandomTransactions(ctx context.Context, client *ent.Client, f *gofakeit.Faker, a *ent.Account) ([]*ent.Transaction, error) {

	trns := []*ent.Transaction{}

	// Define the AVERAGE number of TRANSACTIONS per MONTH per account TYPE
	nrTrans := map[string]int{
		product.TypeSAVING.String():     2, // Reg-D limit = 6
		product.TypeLOAN.String():       14,
		product.TypeDEPOSIT.String():    2,  // Reg-D limit = 6
		product.TypeCHECKING.String():   72, // Includes debit, credit, or prepaid cards, cash (ATM), paper checks, electronic
		product.TypeINVESTMENT.String(): 2,  // Same as savings
	}

	pt, err := a.QueryProduct().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ACCOUNT_OWNER_PRODUCT_TYPE_ERROR, err)
	}

	acco, err := a.QueryOwners().First(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ACCOUNT_OWNER_TYPE_ERROR, err)
	}

	acca, err := acco.QueryAddresses().First(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ACCOUNT_OWNER_ADDRESS_ERROR, err)
	}

	// Calculate the number of transactions to generate
	n := int(math.Round(time.Since(a.DateOpened).Hours() / 24))
	it := (n / 30) * nrTrans[pt.Type.String()]
	tranDates := generateTransactionDates(f, it, a.DateOpened)
	accBalance := float64(a.CurrentBalance)

	for tr := 0; tr < it; tr++ {

		// Some transaction stats
		dir := transaction.DirectionCREDIT
		trdesc := f.Sentence(4)
		amnt := f.Price(-3000, 4000)
		if amnt < 0 {
			dir = transaction.DirectionDEBIT
			trdesc = f.Company() + " " + f.CompanySuffix()
		}

		tranCreatedDate := tranDates[tr]

		var hasCheck bool = false
		if f.Number(1, 12) == 12 && pt.Type == product.TypeCHECKING {
			hasCheck = true
		}

		accBalance += amnt

		t, err := client.Transaction.
			Create().
			SetAccount(a).
			SetStatus(transaction.StatusPOSTED).
			SetExecutedAmount(math.Abs(amnt)).
			SetExecutedCurrencyCode(a.CurrencyCode).
			SetDirection(dir).
			SetRunningBalance(math.Round(accBalance*100) / 100).
			SetCreatedDate(tranCreatedDate).
			SetPostedDate(tranCreatedDate).
			SetExecutedDate(tranCreatedDate).
			SetDescription(trdesc).
			// SetLatitude(f.Latitude()).
			// SetLongitude(f.Longitude()).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", TRANSACTION_CREATE_ERROR, err)
		}

		if hasCheck {
			checkImage, checkNumber, err := renderCheck(f, dir, acco, acca, tranCreatedDate, amnt, trdesc)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", CHECK_IMAGE_CREATE_ERROR, err)
			}

			_, err = client.BinaryItem.
				Create().
				SetContent(checkImage.Bytes()).
				SetTransaction(t).
				SetFormat("application/pdf").
				SetLength(len(checkImage.Bytes())).
				SetURL("https://my.cdn.com/imgs/check/" + f.DigitN(18) + ".pdf").
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", CHECK_IMAGE_CREATE_ERROR, err)
			}

			t.Update().SetCheckNumber(checkNumber).Save(ctx)
		}
		trns = append(trns, t)
	}
	return trns, nil
}

func generateTransactionDates(f *gofakeit.Faker, days int, ao time.Time) []time.Time {
	a := util.TimeSlice{}
	for d := 0; d < days; d++ {
		rd := f.DateRange(ao, time.Now())

		// Shift date to next workable day
		switch int(rd.Weekday()) {
		case 0:
			rd = rd.AddDate(0, 0, 1)
		case 6:
			rd = rd.AddDate(0, 0, 2)
		}
		a = append(a, rd)
	}
	sort.Sort(a)
	return a
}

func renderCheck(f *gofakeit.Faker, dir transaction.Direction, acco *ent.Entity, acca *ent.EntityAddress, tranCreatedDate time.Time, amnt float64, trdesc string) (bytes.Buffer, string, error) {
	var checkImage bytes.Buffer
	var checkNumber string
	var err error

	if dir == transaction.DirectionDEBIT {
		checkImage, checkNumber, err = renderCheckAsPDF(f, tranCreatedDate, acco.Fullname, acca.Line1, acca.City, acca.PostalCode, acca.State, amnt, trdesc, f.Sentence(4))
		if err != nil {
			return bytes.Buffer{}, "", fmt.Errorf("%s: %w", CHECK_IMAGE_CREATE_ERROR, err)
		}
	} else {
		checkImage, checkNumber, err = renderCheckAsPDF(f, tranCreatedDate, f.FirstName()+" "+f.LastName(), f.StreetNumber()+" "+f.Street(), f.City(), f.Zip(), f.StateAbr(), amnt, acco.Fullname, trdesc)
		if err != nil {
			return bytes.Buffer{}, "", fmt.Errorf("%s: %w", CHECK_IMAGE_CREATE_ERROR, err)
		}
	}
	return checkImage, checkNumber, nil
}
