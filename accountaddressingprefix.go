// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAddressingPrefixService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAddressingPrefixService] method instead.
type AccountAddressingPrefixService struct {
	Options     []option.RequestOption
	Bgp         *AccountAddressingPrefixBgpService
	Bindings    *AccountAddressingPrefixBindingService
	Delegations *AccountAddressingPrefixDelegationService
}

// NewAccountAddressingPrefixService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAddressingPrefixService(opts ...option.RequestOption) (r *AccountAddressingPrefixService) {
	r = &AccountAddressingPrefixService{}
	r.Options = opts
	r.Bgp = NewAccountAddressingPrefixBgpService(opts...)
	r.Bindings = NewAccountAddressingPrefixBindingService(opts...)
	r.Delegations = NewAccountAddressingPrefixDelegationService(opts...)
	return
}

// List a particular prefix owned by the account.
func (r *AccountAddressingPrefixService) Get(ctx context.Context, accountID string, prefixID string, opts ...option.RequestOption) (res *SingleResponsePrefix, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify the description for a prefix owned by the account.
func (r *AccountAddressingPrefixService) Update(ctx context.Context, accountID string, prefixID string, body AccountAddressingPrefixUpdateParams, opts ...option.RequestOption) (res *SingleResponsePrefix, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all prefixes owned by the account.
func (r *AccountAddressingPrefixService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAddressingPrefixListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an unapproved prefix owned by the account.
func (r *AccountAddressingPrefixService) Delete(ctx context.Context, accountID string, prefixID string, opts ...option.RequestOption) (res *APIResponseCollectionAddressing, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add a new prefix under the account.
func (r *AccountAddressingPrefixService) Add(ctx context.Context, accountID string, body AccountAddressingPrefixAddParams, opts ...option.RequestOption) (res *SingleResponsePrefix, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type IpamPrefixes struct {
	// Identifier of an IP Prefix.
	ID string `json:"id"`
	// Identifier of a Cloudflare account.
	AccountID string `json:"account_id"`
	// Prefix advertisement status to the Internet. This field is only not 'null' if on
	// demand is enabled.
	Advertised bool `json:"advertised,nullable"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time `json:"advertised_modified_at,nullable" format:"date-time"`
	// Approval state of the prefix (P = pending, V = active).
	Approved string `json:"approved"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	Asn int64 `json:"asn,nullable"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Description of the prefix.
	Description string `json:"description"`
	// Identifier for the uploaded LOA document.
	LoaDocumentID string    `json:"loa_document_id,nullable"`
	ModifiedAt    time.Time `json:"modified_at" format:"date-time"`
	// Whether advertisement of the prefix to the Internet may be dynamically enabled
	// or disabled.
	OnDemandEnabled bool `json:"on_demand_enabled"`
	// Whether advertisement status of the prefix is locked, meaning it cannot be
	// changed.
	OnDemandLocked bool             `json:"on_demand_locked"`
	JSON           ipamPrefixesJSON `json:"-"`
}

// ipamPrefixesJSON contains the JSON metadata for the struct [IpamPrefixes]
type ipamPrefixesJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	Approved             apijson.Field
	Asn                  apijson.Field
	Cidr                 apijson.Field
	CreatedAt            apijson.Field
	Description          apijson.Field
	LoaDocumentID        apijson.Field
	ModifiedAt           apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *IpamPrefixes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipamPrefixesJSON) RawJSON() string {
	return r.raw
}

type SingleResponsePrefix struct {
	Errors   []AddressingMessages `json:"errors,required"`
	Messages []AddressingMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SingleResponsePrefixSuccess `json:"success,required"`
	Result  IpamPrefixes                `json:"result"`
	JSON    singleResponsePrefixJSON    `json:"-"`
}

// singleResponsePrefixJSON contains the JSON metadata for the struct
// [SingleResponsePrefix]
type singleResponsePrefixJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponsePrefix) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponsePrefixJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SingleResponsePrefixSuccess bool

const (
	SingleResponsePrefixSuccessTrue SingleResponsePrefixSuccess = true
)

func (r SingleResponsePrefixSuccess) IsKnown() bool {
	switch r {
	case SingleResponsePrefixSuccessTrue:
		return true
	}
	return false
}

type AccountAddressingPrefixListResponse struct {
	Errors   []AddressingMessages `json:"errors,required"`
	Messages []AddressingMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountAddressingPrefixListResponseSuccess    `json:"success,required"`
	Result     []IpamPrefixes                                `json:"result"`
	ResultInfo AccountAddressingPrefixListResponseResultInfo `json:"result_info"`
	JSON       accountAddressingPrefixListResponseJSON       `json:"-"`
}

// accountAddressingPrefixListResponseJSON contains the JSON metadata for the
// struct [AccountAddressingPrefixListResponse]
type accountAddressingPrefixListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingPrefixListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountAddressingPrefixListResponseSuccess bool

const (
	AccountAddressingPrefixListResponseSuccessTrue AccountAddressingPrefixListResponseSuccess = true
)

func (r AccountAddressingPrefixListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountAddressingPrefixListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountAddressingPrefixListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                           `json:"total_count"`
	JSON       accountAddressingPrefixListResponseResultInfoJSON `json:"-"`
}

// accountAddressingPrefixListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountAddressingPrefixListResponseResultInfo]
type accountAddressingPrefixListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingPrefixListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountAddressingPrefixUpdateParams struct {
	// Description of the prefix.
	Description param.Field[string] `json:"description,required"`
}

func (r AccountAddressingPrefixUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAddressingPrefixAddParams struct {
	// Autonomous System Number (ASN) the prefix will be advertised under.
	Asn param.Field[int64] `json:"asn,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr,required"`
	// Identifier for the uploaded LOA document.
	LoaDocumentID param.Field[string] `json:"loa_document_id,required"`
}

func (r AccountAddressingPrefixAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
