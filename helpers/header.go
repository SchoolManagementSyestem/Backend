package helpers

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

// GetTenantId extracts the Tenant ID from the HTTP request headers in GraphQL resolve params
func GetTenantId(p *graphql.ResolveParams) (uuid.UUID, error) {
	// Extract the HTTP request from the context
	req, ok := p.Context.Value("httpRequest").(*http.Request)
	if !ok {
		return uuid.Nil, errors.New("failed to get HTTP request from context")
	}

	// Get Tenant ID from the header
	tenantIDStr := req.Header.Get("X-Tenant-ID")
	if tenantIDStr == "" {
		return uuid.Nil, errors.New("missing X-Tenant-ID header")
	}

	// Parse Tenant ID to UUID
	tenantID, err := uuid.Parse(tenantIDStr)
	if err != nil {
		return uuid.Nil, errors.New("invalid Tenant ID format")
	}

	return tenantID, nil
}

// Authentication extracts the user claims from the HTTP request in GraphQL resolve params
func Authentication(p *graphql.ResolveParams) (CustomClaims, error) {
	// Extract the HTTP request from the context
	claims, ok := p.Context.Value("user").(*CustomClaims)
	if !ok {
		return CustomClaims{}, errors.New("failed to get HTTP request from context")
	}

	return *claims, nil
}

// Authorization checks if the user has the required role to access the resource
func Authorization(claims *CustomClaims, role string) error {
	if claims.Role == "admin" {
		return nil
	} else if claims.Role != role {
		return errors.New("unauthorized access")
	}
	return nil
}

// Authorizations checks if the user has the required role to access the resource
func Authorizations(claims *CustomClaims, roles []string) error {
	// Admins always have access
	if claims.Role == "admin" {
		return nil
	}

	// Check if the user role exists in the allowed roles
	for _, allowedRole := range roles {
		if claims.Role == allowedRole {
			return nil
		}
	}

	return errors.New("unauthorized access")
}
