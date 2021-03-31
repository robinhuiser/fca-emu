/*
 * Cloud API
 *
 * The public facing API through which connectors are exposed as a single abtract API
 *
 * API version: v1.5
 * Contact: support@trexis.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
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
	// TODO - update GetEntity with the required logic for this service method.
	// Add api_entity_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Entity{}) or use other options such as http.Ok ...
	//return Response(200, Entity{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetEntity method not implemented")
}

// GetEntityProfile - Return entity profile
func (s *EntityApiService) GetEntityProfile(ctx context.Context, entityId string, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {
	// TODO - update GetEntityProfile with the required logic for this service method.
	// Add api_entity_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, EntityProfile{}) or use other options such as http.Ok ...
	//return Response(200, EntityProfile{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetEntityProfile method not implemented")
}

// PostEntityProfile - Create entity
func (s *EntityApiService) PostEntityProfile(ctx context.Context, mask bool, enhance bool, xTRACEID string, xTOKEN string, entityProfile EntityProfile) (ImplResponse, error) {
	// TODO - update PostEntityProfile with the required logic for this service method.
	// Add api_entity_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Entity{}) or use other options such as http.Ok ...
	//return Response(200, Entity{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PostEntityProfile method not implemented")
}

// PutEntityProfile - Update entity profile
func (s *EntityApiService) PutEntityProfile(ctx context.Context, entityId string, mask bool, enhance bool, xTRACEID string, xTOKEN string, entityProfile EntityProfile) (ImplResponse, error) {
	// TODO - update PutEntityProfile with the required logic for this service method.
	// Add api_entity_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, EntityProfile{}) or use other options such as http.Ok ...
	//return Response(200, EntityProfile{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PutEntityProfile method not implemented")
}

// SearchEntities - Search for entities
func (s *EntityApiService) SearchEntities(ctx context.Context, limit int32, cursor string, mask bool, enhance bool, xTRACEID string, xTOKEN string, searchFilter []SearchFilter) (ImplResponse, error) {
	// TODO - update SearchEntities with the required logic for this service method.
	// Add api_entity_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, EntityList{}) or use other options such as http.Ok ...
	//return Response(200, EntityList{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SearchEntities method not implemented")
}

