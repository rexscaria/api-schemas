// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAddressingRegionalHostnameService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAddressingRegionalHostnameService] method instead.
type AccountAddressingRegionalHostnameService struct {
	Options []option.RequestOption
}

// NewAccountAddressingRegionalHostnameService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressingRegionalHostnameService(opts ...option.RequestOption) (r *AccountAddressingRegionalHostnameService) {
	r = &AccountAddressingRegionalHostnameService{}
	r.Options = opts
	return
}

// List all Regional Services regions available for use by this account.
func (r *AccountAddressingRegionalHostnameService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAddressingRegionalHostnameListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/regional_hostnames/regions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAddressingRegionalHostnameListResponse struct {
	Errors   []MessagesDlsItem `json:"errors,required"`
	Messages []MessagesDlsItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountAddressingRegionalHostnameListResponseSuccess    `json:"success,required"`
	Result     []AccountAddressingRegionalHostnameListResponseResult   `json:"result"`
	ResultInfo AccountAddressingRegionalHostnameListResponseResultInfo `json:"result_info"`
	JSON       accountAddressingRegionalHostnameListResponseJSON       `json:"-"`
}

// accountAddressingRegionalHostnameListResponseJSON contains the JSON metadata for
// the struct [AccountAddressingRegionalHostnameListResponse]
type accountAddressingRegionalHostnameListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingRegionalHostnameListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingRegionalHostnameListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountAddressingRegionalHostnameListResponseSuccess bool

const (
	AccountAddressingRegionalHostnameListResponseSuccessTrue AccountAddressingRegionalHostnameListResponseSuccess = true
)

func (r AccountAddressingRegionalHostnameListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountAddressingRegionalHostnameListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountAddressingRegionalHostnameListResponseResult struct {
	// Identifying key for the region
	Key string `json:"key"`
	// Human-readable text label for the region
	Label string                                                  `json:"label"`
	JSON  accountAddressingRegionalHostnameListResponseResultJSON `json:"-"`
}

// accountAddressingRegionalHostnameListResponseResultJSON contains the JSON
// metadata for the struct [AccountAddressingRegionalHostnameListResponseResult]
type accountAddressingRegionalHostnameListResponseResultJSON struct {
	Key         apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingRegionalHostnameListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingRegionalHostnameListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAddressingRegionalHostnameListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                     `json:"total_count"`
	JSON       accountAddressingRegionalHostnameListResponseResultInfoJSON `json:"-"`
}

// accountAddressingRegionalHostnameListResponseResultInfoJSON contains the JSON
// metadata for the struct
// [AccountAddressingRegionalHostnameListResponseResultInfo]
type accountAddressingRegionalHostnameListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingRegionalHostnameListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingRegionalHostnameListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}
