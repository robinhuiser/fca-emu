package generator

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/robinhuiser/finite-mock-server/ent"
)

func populateBanks(ctx context.Context, client *ent.Client, f *gofakeit.Faker) error {

	b := map[string]string{
		"BOFAUS3NXXX": "Bank of America",
		"NFBKUS33XXX": "Capital One",
		"CHASUSU3XXX": "Chase Bank (Jp Morgan Chase)",
		"CITIUS33XXX": "Citibank",
		"FTBCUS3CXXX": "Fifth Third Bank",
		"MRMDUS33XXX": "HSBC",
		"PNCCUS33XXX": "PNC Bank",
		"SNTRUS3AXXX": "Truist Bank",
		"USBKUS44IMT": "U.S. Bancorp",
		"WFBIUS6SXXX": "Wells Fargo Bank",
	}

	// Create banks with branches
	for swiftcode, bankname := range b {
		bank, err := client.Bank.
			Create().
			SetBankName(bankname).
			SetBankCode(swiftcode[0:4]).
			SetSwift(swiftcode).
			SetURL("https://my.cdn.com/imgs/bank/" + f.DigitN(18) + ".jpg").
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating bank: %w", err)
		}

		// Generate x number of branches per bank
		for i := 0; i < f.Number(3, 50); i++ {
			branch, err := client.Branch.
				Create().
				SetBranchCode(fmt.Sprintf("%s-%03d", swiftcode[0:4], i)).
				SetStreetNumber(f.StreetNumber()).
				SetStreetName(f.StreetName()).
				SetCity(f.City()).
				SetZip(f.Zip()).
				SetState(f.StateAbr()).
				SetLatitude(f.Latitude()).
				SetLongitude(f.Longitude()).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed creating branch: %w", err)
			}
			bank.Update().
				AddBranches(branch).
				Save(ctx)
		}
	}
	return nil
}
