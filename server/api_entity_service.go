/*
 * Cloud API
 *
 * The public facing API through which connectors are exposed as a single abtract API
 *
 * API version: v1.5
 * Contact: support@trexis.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package finite

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/robinhuiser/fca-emu/ent"
	"github.com/robinhuiser/fca-emu/ent/entity"
)

// EntityApiService is a service that implents the logic for the EntityApiServicer
// This service should implement the business logic for every endpoint for the EntityApi API.
// Include any external packages or services that will be required by this service.
type EntityApiService struct {
}

// NewEntityApiService creates a default api service
func NewEntityApiService() EntityApiServicer {
	return &EntityApiService{}
}

// GetEntity - Return entity by entityId
func (s *EntityApiService) GetEntity(ctx context.Context, entityId string, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {
	// Validate X-Token
	if !isValidSecret(xTOKEN) {
		return Response(401, setErrorResponse(INVALID_TOKEN_MSG)), nil
	}

	// Parse and verify UUID
	entityUUID, err := uuid.Parse(entityId)
	if err != nil {
		return Response(400, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	// Perform the search
	rs, err := clt.Entity.
		Query().
		Where(entity.ID(entityUUID)).
		Only(ctx)
	// Error if none or more than one results are returned
	if err != nil {
		return Response(404, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	entities := mapEntities([]*ent.Entity{rs}, ctx)
	return Response(200, entities[0]), nil
}

// GetEntityProfile - Return entity profile
func (s *EntityApiService) GetEntityProfile(ctx context.Context, entityId string, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {
	// Validate X-Token
	if !isValidSecret(xTOKEN) {
		return Response(401, setErrorResponse(INVALID_TOKEN_MSG)), nil
	}

	// Parse and verify UUID
	entityUUID, err := uuid.Parse(entityId)
	if err != nil {
		return Response(400, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	// Perform the search
	rs, err := clt.Entity.
		Query().
		Where(entity.ID(entityUUID)).
		Only(ctx)
	// Error if none or more than one results are returned
	if err != nil {
		return Response(404, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	// Retrieve the linked entity attributes
	addrs, err := rs.QueryAddresses().All(ctx)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}
	addresses := mapAddresses(addrs)

	// Retrieve the linked entity attributes
	atrs, err := rs.QueryPreferences().All(ctx)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}
	preferences := mapPreferences(atrs)

	// Retrieve the linked entity contact points
	cntpts, err := rs.QueryContactPoints().All(ctx)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}
	contactpoints := mapContactPoints(cntpts)

	// Retrieve the linked entity tax information
	taxspecs, err := rs.QueryTaxSpecifications().All(ctx)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}
	taxspecifications := mapTaxSpecifications(taxspecs)

	entityProfile := EntityProfile{
		Firstname:      rs.Firstname,
		Addresses:      addresses,
		Preferences:    preferences,
		ContactPoints:  contactpoints,
		DateOfBirth:    isValidBankDate(rs.DateOfBirth.Format(API_DATE_LAYOUT)),
		Id:             rs.ID.String(),
		Fullname:       rs.Fullname,
		TaxInformation: taxspecifications,
		Type:           rs.Type.String(),
		URI: FiniteUri{
			URL: rs.URL,
		},
		Lastname: rs.Lastname,
	}
	return Response(200, entityProfile), nil
}

// PostEntityProfile - Create entity
func (s *EntityApiService) PostEntityProfile(ctx context.Context, mask bool, enhance bool, xTRACEID string, xTOKEN string, entityProfile EntityProfile) (ImplResponse, error) {

	return Response(http.StatusNotImplemented, nil), errors.New("PostEntityProfile method not implemented")
}

// PutEntityProfile - Update entity profile
func (s *EntityApiService) PutEntityProfile(ctx context.Context, entityId string, mask bool, enhance bool, xTRACEID string, xTOKEN string, entityProfile EntityProfile) (ImplResponse, error) {

	return Response(http.StatusNotImplemented, nil), errors.New("PutEntityProfile method not implemented")
}

// SearchEntities - Search for entities
func (s *EntityApiService) SearchEntities(ctx context.Context, limit int32, cursor string, mask bool, enhance bool, xTRACEID string, xTOKEN string, searchFilter []SearchFilter) (ImplResponse, error) {

	return Response(http.StatusNotImplemented, nil), errors.New("SearchEntities method not implemented")
}

// Helper functions for mapping
func mapEntities(owners []*ent.Entity, ctx context.Context) []Entity {
	entities := []Entity{}
	taxInformations := []TaxInformation{}
	for _, o := range owners {

		// Get Tax information
		txs, _ := o.QueryTaxSpecifications().All(ctx)
		for _, t := range txs {
			tx := TaxInformation{
				TaxId: t.TaxId,
				Type:  string(t.Type),
			}
			taxInformations = append(taxInformations, tx)
		}

		e := Entity{
			Id:          o.ID.String(),
			Name:        o.Fullname,
			Active:      o.Active,
			DateCreated: isValidBankDate(o.DateCreated.Format(API_DATE_LAYOUT)),
			SecurityInformation: SecurityInformation{
				LastLoginDate: isValidBankDate(o.LastLoginDate.Format(API_DATE_LAYOUT)),
				Username:      o.Username,
				Token:         o.Token,
			},
			URI: FiniteUri{
				URL: o.URL,
			},
			Type:           string(o.Type),
			TaxInformation: taxInformations,
		}
		entities = append(entities, e)
	}
	return entities
}

func mapAddresses(addrs []*ent.EntityAddress) []Address {
	addresses := []Address{}
	for _, addr := range addrs {
		a := Address{
			Country:    addr.Country,
			City:       addr.City,
			PostalCode: addr.PostalCode,
			State:      addr.State,
			Type:       string(addr.Type),
			Line1:      addr.Line1,
			Line2:      addr.Line2,
			Line3:      addr.Line3,
			Primary:    addr.Primary,
		}
		addresses = append(addresses, a)
	}
	return addresses
}

func mapContactPoints(cntpts []*ent.EntityContactPoint) []ContactPoint {
	contactpoints := []ContactPoint{}
	for _, cntpt := range cntpts {
		c := ContactPoint{
			Prefix: cntpt.Prefix,
			Name:   cntpt.Name,
			Type:   cntpt.Type.String(),
			Suffix: cntpt.Suffix,
			Value:  cntpt.Value,
		}
		contactpoints = append(contactpoints, c)
	}
	return contactpoints
}

func mapTaxSpecifications(taxinfs []*ent.EntityTaxInformation) []TaxInformation {
	taxinformations := []TaxInformation{}
	for _, taxinf := range taxinfs {
		c := TaxInformation{
			TaxId: taxinf.TaxId,
			Type:  string(taxinf.Type),
		}
		taxinformations = append(taxinformations, c)
	}
	return taxinformations
}
