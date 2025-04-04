// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountIntelAsnService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountIntelAsnService] method instead.
type AccountIntelAsnService struct {
	Options []option.RequestOption
}

// NewAccountIntelAsnService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountIntelAsnService(opts ...option.RequestOption) (r *AccountIntelAsnService) {
	r = &AccountIntelAsnService{}
	r.Options = opts
	return
}

// Get ASN Subnets
func (r *AccountIntelAsnService) ListSubnets(ctx context.Context, accountID string, asn int64, opts ...option.RequestOption) (res *AccountIntelAsnListSubnetsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/asn/%v/subnets", accountID, asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets an overview of the Autonomous System Number (ASN) and a list of subnets for
// it.
func (r *AccountIntelAsnService) GetOverview(ctx context.Context, accountID string, asn int64, opts ...option.RequestOption) (res *AccountIntelAsnGetOverviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/asn/%v", accountID, asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountIntelAsnListSubnetsResponse struct {
	Asn int64 `json:"asn"`
	// Total results returned based on your search parameters.
	Count        float64 `json:"count"`
	IPCountTotal int64   `json:"ip_count_total"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64                                `json:"per_page"`
	Subnets []string                               `json:"subnets"`
	JSON    accountIntelAsnListSubnetsResponseJSON `json:"-"`
}

// accountIntelAsnListSubnetsResponseJSON contains the JSON metadata for the struct
// [AccountIntelAsnListSubnetsResponse]
type accountIntelAsnListSubnetsResponseJSON struct {
	Asn          apijson.Field
	Count        apijson.Field
	IPCountTotal apijson.Field
	Page         apijson.Field
	PerPage      apijson.Field
	Subnets      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountIntelAsnListSubnetsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelAsnListSubnetsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIntelAsnGetOverviewResponse struct {
	Result int64                                  `json:"result"`
	JSON   accountIntelAsnGetOverviewResponseJSON `json:"-"`
	SingleResponseIntel
}

// accountIntelAsnGetOverviewResponseJSON contains the JSON metadata for the struct
// [AccountIntelAsnGetOverviewResponse]
type accountIntelAsnGetOverviewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelAsnGetOverviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelAsnGetOverviewResponseJSON) RawJSON() string {
	return r.raw
}
