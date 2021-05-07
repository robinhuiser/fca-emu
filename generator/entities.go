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
	"github.com/robinhuiser/fca-emu/ent/entitycontactpoint"
	"github.com/robinhuiser/fca-emu/ent/entitytaxinformation"
)

const (
	ENTITY_CREATE_ERROR            = "failed creating entity"
	ENTITY_CONTACT_POINT_ERROR     = "failed creating entity contact point"
	ENTITY_PREFERENCE_ERROR        = "failed creating preference"
	ENTITY_TAX_SPECIFICATION_ERROR = "failed creating entity tax specification"
	ENTITY_ADDRESS_ERROR           = "failed creating entity address"
)

func populateEntity(ctx context.Context, client *ent.Client, f *gofakeit.Faker) (*ent.Entity, error) {

	dco := f.DateRange(time.Now().AddDate(-5, 0, -1), time.Now()) // random date create entity
	var firstname, lastname, fullname string
	var entityType entity.Type
	var addressType entityaddress.Type

	// Person 70%, organization 10% or corporate 20%
	switch a := f.Number(0, 100); {
	case a >= 70 && a <= 90:
		fullname = f.Company() + ", " + f.CompanySuffix()
		entityType = entity.TypeBUSINESS
	case a >= 90:
		fullname = f.Company()
		entityType = entity.TypeSYSTEM
	default:
		firstname = f.FirstName()
		lastname = f.LastName()
		fullname = firstname + " " + lastname
		entityType = entity.TypePERSON
	}

	// Create entity
	e, err := client.Entity.
		Create().
		SetDateCreated(dco).
		SetFirstname(firstname).
		SetLastname(lastname).
		SetFullname(fullname).
		SetType(entityType).
		SetDateOfBirth(f.DateRange(
			time.Now().AddDate(-60, 0, -1),
			time.Now().AddDate(-18, 0, -1))).
		SetLastLoginDate(f.DateRange(time.Now().AddDate(0, -30, -1), time.Now())).
		SetUsername(f.Username()).
		SetToken(f.LetterN(24) + "==").
		SetURL("https://my.cdn.com/imgs/entity/" + f.DigitN(18) + ".jpg").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ENTITY_CREATE_ERROR, err)
	}

	// Private business or person
	switch et := entityType; et {
	case entity.TypeBUSINESS:
		addressType = entityaddress.TypeBUSINESS
	case entity.TypeSYSTEM:
		addressType = entityaddress.TypePOBOX
	default:
		addressType = entityaddress.TypeRESIDENTIAL
	}

	// Add one or more addresses
	primaryAddress := true
	for i := 0; i < f.Number(1, 2); i++ {
		a, err := client.EntityAddress.
			Create().
			SetCountry("US").
			SetCity(f.City()).
			SetPostalCode(f.Zip()).
			SetState(f.StateAbr()).
			SetLine1((f.StreetNumber()) + " " + f.StreetName() + " " + f.StreetSuffix()).
			SetType(addressType).
			SetPrimary(primaryAddress).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", ENTITY_ADDRESS_ERROR, err)
		}
		e.Update().AddAddresses(a).Save(ctx)
		primaryAddress = false
	}

	// Add one or more tax information records
	for i := 0; i < f.Number(1, 2); i++ {
		t, err := client.EntityTaxInformation.
			Create().
			SetType(entitytaxinformation.TypeSSN).
			SetTaxId(f.SSN()).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", ENTITY_TAX_SPECIFICATION_ERROR, err)
		}
		e.Update().AddTaxSpecifications(t).Save(ctx)
	}

	// Add one or more preferences
	for i := 0; i < f.Number(1, 8); i++ {
		p, err := client.Preference.
			Create().
			SetName(f.Noun()).
			SetValue(f.Word()).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", ENTITY_PREFERENCE_ERROR, err)
		}
		e.Update().AddPreferences(p).Save(ctx)
	}

	// Add one or more contact points
	for i := 0; i < f.Number(1, 2); i++ {
		switch a := f.Number(0, 2); {
		case a == 0:
			c, err := client.EntityContactPoint.
				Create().
				SetName("Work mobile").
				SetType(entitycontactpoint.TypeSMS).
				SetValue(f.Phone()).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", ENTITY_CONTACT_POINT_ERROR, err)
			}
			e.Update().AddContactPoints(c).Save(ctx)
		case a == 1:
			c, err := client.EntityContactPoint.
				Create().
				SetName("Private email").
				SetType(entitycontactpoint.TypeEMAIL).
				SetValue(f.Email()).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", ENTITY_CONTACT_POINT_ERROR, err)
			}
			e.Update().AddContactPoints(c).Save(ctx)
		case a == 2:
			c, err := client.EntityContactPoint.
				Create().
				SetPrefix(strconv.Itoa(f.Number(1, 9))).
				SetName("Work phone").
				SetType(entitycontactpoint.TypeVOICE).
				SetSuffix(strconv.Itoa(f.Number(111, 999))).
				SetValue(f.Phone()).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", ENTITY_CONTACT_POINT_ERROR, err)
			}
			e.Update().AddContactPoints(c).Save(ctx)
		}

	}

	return e, nil
}
