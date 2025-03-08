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
