// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAccessOrganizationDohService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessOrganizationDohService] method instead.
type AccountAccessOrganizationDohService struct {
	Options []option.RequestOption
}

// NewAccountAccessOrganizationDohService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessOrganizationDohService(opts ...option.RequestOption) (r *AccountAccessOrganizationDohService) {
	r = &AccountAccessOrganizationDohService{}
	r.Options = opts
	return
}

// Returns the DoH settings for your Zero Trust organization.
func (r *AccountAccessOrganizationDohService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAccessOrganizationDohGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/organizations/doh", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the DoH settings for your Zero Trust organization.
func (r *AccountAccessOrganizationDohService) Update(ctx context.Context, accountID string, body AccountAccessOrganizationDohUpdateParams, opts ...option.RequestOption) (res *AccountAccessOrganizationDohUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/organizations/doh", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountAccessOrganizationDohGetResponse struct {
	Result AccountAccessOrganizationDohGetResponseResult `json:"result"`
	JSON   accountAccessOrganizationDohGetResponseJSON   `json:"-"`
	SchemasAccessSingleResponse
}

// accountAccessOrganizationDohGetResponseJSON contains the JSON metadata for the
// struct [AccountAccessOrganizationDohGetResponse]
type accountAccessOrganizationDohGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessOrganizationDohGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessOrganizationDohGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessOrganizationDohGetResponseResult struct {
	// The duration the DoH JWT is valid for. Must be in the format `300ms` or `2h45m`.
	// Valid time units are: ns, us (or µs), ms, s, m, h. Note that the maximum
	// duration for this setting is the same as the key rotation period on the account.
	DohJwtDuration string                                            `json:"doh_jwt_duration"`
	JSON           accountAccessOrganizationDohGetResponseResultJSON `json:"-"`
}

// accountAccessOrganizationDohGetResponseResultJSON contains the JSON metadata for
// the struct [AccountAccessOrganizationDohGetResponseResult]
type accountAccessOrganizationDohGetResponseResultJSON struct {
	DohJwtDuration apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessOrganizationDohGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessOrganizationDohGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessOrganizationDohUpdateResponse struct {
	Result AccountAccessOrganizationDohUpdateResponseResult `json:"result"`
	JSON   accountAccessOrganizationDohUpdateResponseJSON   `json:"-"`
	SchemasAccessSingleResponse
}

// accountAccessOrganizationDohUpdateResponseJSON contains the JSON metadata for
// the struct [AccountAccessOrganizationDohUpdateResponse]
type accountAccessOrganizationDohUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessOrganizationDohUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessOrganizationDohUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessOrganizationDohUpdateResponseResult struct {
	// The duration the DoH JWT is valid for. Must be in the format `300ms` or `2h45m`.
	// Valid time units are: ns, us (or µs), ms, s, m, h. Note that the maximum
	// duration for this setting is the same as the key rotation period on the account.
	// Default expiration is 24h
	DohJwtDuration string                                               `json:"doh_jwt_duration"`
	JSON           accountAccessOrganizationDohUpdateResponseResultJSON `json:"-"`
}

// accountAccessOrganizationDohUpdateResponseResultJSON contains the JSON metadata
// for the struct [AccountAccessOrganizationDohUpdateResponseResult]
type accountAccessOrganizationDohUpdateResponseResultJSON struct {
	DohJwtDuration apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessOrganizationDohUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessOrganizationDohUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessOrganizationDohUpdateParams struct {
	// The duration the DoH JWT is valid for. Must be in the format `300ms` or `2h45m`.
	// Valid time units are: ns, us (or µs), ms, s, m, h. Note that the maximum
	// duration for this setting is the same as the key rotation period on the account.
	// Default expiration is 24h
	DohJwtDuration param.Field[string] `json:"doh_jwt_duration"`
	// The uuid of the service token you want to use for DoH authentication
	ServiceTokenID param.Field[string] `json:"service_token_id"`
}

func (r AccountAccessOrganizationDohUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
