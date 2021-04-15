package generator

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/robinhuiser/fca-emu/ent"
	"github.com/robinhuiser/fca-emu/ent/entity"
	"github.com/robinhuiser/fca-emu/ent/entityaddress"
)

func populateEntity(ctx context.Context, client *ent.Client, f *gofakeit.Faker) (*ent.Entity, error) {

	dco := f.DateRange(time.Now().AddDate(-5, 0, -1), time.Now()) // random date create entity
	var firstname, lastname, fullname string
	var entity_type entity.Type
	var address_type entityaddress.Type

	// Person 70%, organization 10% or corporate 20%
	switch a := f.Number(0, 100); {
	case a >= 70 && a <= 90:
		fullname = f.Company() + ", " + f.CompanySuffix()
		entity_type = "CORPORATE"
	case a >= 90:
		fullname = f.Company()
		entity_type = "ORGANIZATION"
	default:
		firstname = f.FirstName()
		lastname = f.LastName()
		fullname = firstname + " " + lastname
		entity_type = "PERSON"
	}

	// Create entity
	e, err := client.Entity.
		Create().
		SetDateCreated(dco).
		SetFirstname(firstname).
		SetLastname(lastname).
		SetFullname(fullname).
		SetType(entity_type).
		SetDateOfBirth(f.DateRange(
			time.Now().AddDate(-60, 0, -1),
			time.Now().AddDate(-18, 0, -1))).
		SetLastLoginDate(f.DateRange(time.Now().AddDate(0, -30, -1), time.Now())).
		SetUsername(f.Username()).
		SetToken(f.LetterN(24) + "==").
		SetURL("https://my.cdn.com/imgs/entity/" + f.DigitN(18) + ".jpg").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating entity: %w", err)
	}

	// Private business or person
	switch et := entity_type.String(); et {
	case "CORPORATE":
		address_type = "BUSINESS"
	case "ORGANIZATION":
		address_type = "MAILBOX"
	default:
		address_type = "PRIVATE"
	}

	// Add one or more addresses
	primary_address := true
	for i := 0; i < f.Number(1, 2); i++ {
		a, err := client.EntityAddress.
			Create().
			SetCountry("Unites States of America").
			SetCity(f.City()).
			SetPostalCode(f.Zip()).
			SetState(f.State()).
			SetLine1(f.StreetNumber()).
			SetLine2(f.StreetName()).
			SetLine3(f.StreetSuffix()).
			SetType(address_type).
			SetPrimary(primary_address).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed creating entity address: %w", err)
		}
		e.Update().AddEntityAddresses(a).Save(ctx)
		primary_address = false
	}

	// Add one or more tax information records
	for i := 0; i < f.Number(1, 2); i++ {
		t, err := client.EntityTaxInformation.
			Create().
			SetType("SSN").
			SetTaxId(f.SSN()).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed creating entity tax info: %w", err)
		}
		e.Update().AddEntityTaxInformation(t).Save(ctx)
	}

	// Add one or more preferences
	for i := 0; i < f.Number(1, 8); i++ {
		p, err := client.Preference.
			Create().
			SetName(f.Noun()).
			SetValue(f.Word()).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed creating preferences: %w", err)
		}
		e.Update().AddEntityPreferences(p).Save(ctx)
	}

	// Add one or more contact points
	// "SMS", "EMAIL", "PHONE"
	for i := 0; i < f.Number(1, 2); i++ {
		switch a := f.Number(0, 2); {
		case a == 0:
			c, err := client.EntityContactPoint.
				Create().
				SetName("Work mobile").
				SetType("SMS").
				SetValue(f.Phone()).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed creating entity contact points: %w", err)
			}
			e.Update().AddEntityContactPoints(c).Save(ctx)
		case a == 1:
			c, err := client.EntityContactPoint.
				Create().
				SetName("Private email").
				SetType("EMAIL").
				SetValue(f.Email()).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed creating entity contact points: %w", err)
			}
			e.Update().AddEntityContactPoints(c).Save(ctx)
		case a == 2:
			c, err := client.EntityContactPoint.
				Create().
				SetPrefix(strconv.Itoa(f.Number(1, 9))).
				SetName("Work phone").
				SetType("PHONE").
				SetSuffix(strconv.Itoa(f.Number(111, 999))).
				SetValue(f.Phone()).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed creating entity contact points: %w", err)
			}
			e.Update().AddEntityContactPoints(c).Save(ctx)
		}

	}

	return e, nil
}
