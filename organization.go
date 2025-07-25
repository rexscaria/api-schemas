// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// OrganizationService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	return
}

// Lists all organization shares.
func (r *OrganizationService) ListShares(ctx context.Context, organizationID string, query OrganizationListSharesParams, opts ...option.RequestOption) (res *ShareResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/shares", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OrganizationListSharesParams struct {
	// Direction to sort objects.
	Direction param.Field[OrganizationListSharesParamsDirection] `query:"direction"`
	// Filter shares by kind.
	Kind param.Field[OrganizationListSharesParamsKind] `query:"kind"`
	// Order shares by values in the given field.
	Order param.Field[OrganizationListSharesParamsOrder] `query:"order"`
	// Page number.
	Page param.Field[int64] `query:"page"`
	// Number of objects to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter shares by status.
	Status param.Field[OrganizationListSharesParamsStatus] `query:"status"`
	// Filter shares by target_type.
	TargetType param.Field[OrganizationListSharesParamsTargetType] `query:"target_type"`
}

// URLQuery serializes [OrganizationListSharesParams]'s query parameters as
// `url.Values`.
func (r OrganizationListSharesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to sort objects.
type OrganizationListSharesParamsDirection string

const (
	OrganizationListSharesParamsDirectionAsc  OrganizationListSharesParamsDirection = "asc"
	OrganizationListSharesParamsDirectionDesc OrganizationListSharesParamsDirection = "desc"
)

func (r OrganizationListSharesParamsDirection) IsKnown() bool {
	switch r {
	case OrganizationListSharesParamsDirectionAsc, OrganizationListSharesParamsDirectionDesc:
		return true
	}
	return false
}

// Filter shares by kind.
type OrganizationListSharesParamsKind string

const (
	OrganizationListSharesParamsKindSent     OrganizationListSharesParamsKind = "sent"
	OrganizationListSharesParamsKindReceived OrganizationListSharesParamsKind = "received"
)

func (r OrganizationListSharesParamsKind) IsKnown() bool {
	switch r {
	case OrganizationListSharesParamsKindSent, OrganizationListSharesParamsKindReceived:
		return true
	}
	return false
}

// Order shares by values in the given field.
type OrganizationListSharesParamsOrder string

const (
	OrganizationListSharesParamsOrderName    OrganizationListSharesParamsOrder = "name"
	OrganizationListSharesParamsOrderCreated OrganizationListSharesParamsOrder = "created"
)

func (r OrganizationListSharesParamsOrder) IsKnown() bool {
	switch r {
	case OrganizationListSharesParamsOrderName, OrganizationListSharesParamsOrderCreated:
		return true
	}
	return false
}

// Filter shares by status.
type OrganizationListSharesParamsStatus string

const (
	OrganizationListSharesParamsStatusActive   OrganizationListSharesParamsStatus = "active"
	OrganizationListSharesParamsStatusDeleting OrganizationListSharesParamsStatus = "deleting"
	OrganizationListSharesParamsStatusDeleted  OrganizationListSharesParamsStatus = "deleted"
)

func (r OrganizationListSharesParamsStatus) IsKnown() bool {
	switch r {
	case OrganizationListSharesParamsStatusActive, OrganizationListSharesParamsStatusDeleting, OrganizationListSharesParamsStatusDeleted:
		return true
	}
	return false
}

// Filter shares by target_type.
type OrganizationListSharesParamsTargetType string

const (
	OrganizationListSharesParamsTargetTypeAccount      OrganizationListSharesParamsTargetType = "account"
	OrganizationListSharesParamsTargetTypeOrganization OrganizationListSharesParamsTargetType = "organization"
)

func (r OrganizationListSharesParamsTargetType) IsKnown() bool {
	switch r {
	case OrganizationListSharesParamsTargetTypeAccount, OrganizationListSharesParamsTargetTypeOrganization:
		return true
	}
	return false
}
