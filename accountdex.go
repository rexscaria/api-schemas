// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountDexService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDexService] method instead.
type AccountDexService struct {
	Options               []option.RequestOption
	Commands              *AccountDexCommandService
	Devices               *AccountDexDeviceService
	FleetStatus           *AccountDexFleetStatusService
	HTTPTests             *AccountDexHTTPTestService
	Tests                 *AccountDexTestService
	TracerouteTestResults *AccountDexTracerouteTestResultService
	TracerouteTests       *AccountDexTracerouteTestService
}

// NewAccountDexService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountDexService(opts ...option.RequestOption) (r *AccountDexService) {
	r = &AccountDexService{}
	r.Options = opts
	r.Commands = NewAccountDexCommandService(opts...)
	r.Devices = NewAccountDexDeviceService(opts...)
	r.FleetStatus = NewAccountDexFleetStatusService(opts...)
	r.HTTPTests = NewAccountDexHTTPTestService(opts...)
	r.Tests = NewAccountDexTestService(opts...)
	r.TracerouteTestResults = NewAccountDexTracerouteTestResultService(opts...)
	r.TracerouteTests = NewAccountDexTracerouteTestService(opts...)
	return
}

// List Cloudflare colos that account's devices were connected to during a time
// period, sorted by usage starting from the most used colo. Colos without traffic
// are also returned and sorted alphabetically.
func (r *AccountDexService) ListColos(ctx context.Context, accountID string, query AccountDexListColosParams, opts ...option.RequestOption) (res *AccountDexListColosResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/colos", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Item struct {
	Code             int64      `json:"code,required"`
	Message          string     `json:"message,required"`
	DocumentationURL string     `json:"documentation_url"`
	Source           ItemSource `json:"source"`
	JSON             itemJSON   `json:"-"`
}

// itemJSON contains the JSON metadata for the struct [Item]
type itemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Item) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemJSON) RawJSON() string {
	return r.raw
}

type ItemSource struct {
	Pointer string         `json:"pointer"`
	JSON    itemSourceJSON `json:"-"`
}

// itemSourceJSON contains the JSON metadata for the struct [ItemSource]
type itemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemSourceJSON) RawJSON() string {
	return r.raw
}

type AccountDexListColosResponse struct {
	Errors   []Item `json:"errors,required"`
	Messages []Item `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDexListColosResponseSuccess `json:"success,required"`
	// array of colos.
	Result     []interface{}                         `json:"result"`
	ResultInfo AccountDexListColosResponseResultInfo `json:"result_info"`
	JSON       accountDexListColosResponseJSON       `json:"-"`
}

// accountDexListColosResponseJSON contains the JSON metadata for the struct
// [AccountDexListColosResponse]
type accountDexListColosResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexListColosResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexListColosResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDexListColosResponseSuccess bool

const (
	AccountDexListColosResponseSuccessTrue AccountDexListColosResponseSuccess = true
)

func (r AccountDexListColosResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDexListColosResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDexListColosResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                   `json:"total_count"`
	JSON       accountDexListColosResponseResultInfoJSON `json:"-"`
}

// accountDexListColosResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountDexListColosResponseResultInfo]
type accountDexListColosResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexListColosResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexListColosResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountDexListColosParams struct {
	// Start time for connection period in ISO (RFC3339 - ISO 8601) format
	From param.Field[string] `query:"from,required"`
	// End time for connection period in ISO (RFC3339 - ISO 8601) format
	To param.Field[string] `query:"to,required"`
	// Type of usage that colos should be sorted by. If unspecified, returns all
	// Cloudflare colos sorted alphabetically.
	SortBy param.Field[AccountDexListColosParamsSortBy] `query:"sortBy"`
}

// URLQuery serializes [AccountDexListColosParams]'s query parameters as
// `url.Values`.
func (r AccountDexListColosParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Type of usage that colos should be sorted by. If unspecified, returns all
// Cloudflare colos sorted alphabetically.
type AccountDexListColosParamsSortBy string

const (
	AccountDexListColosParamsSortByFleetStatusUsage      AccountDexListColosParamsSortBy = "fleet-status-usage"
	AccountDexListColosParamsSortByApplicationTestsUsage AccountDexListColosParamsSortBy = "application-tests-usage"
)

func (r AccountDexListColosParamsSortBy) IsKnown() bool {
	switch r {
	case AccountDexListColosParamsSortByFleetStatusUsage, AccountDexListColosParamsSortByApplicationTestsUsage:
		return true
	}
	return false
}
