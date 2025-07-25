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

// AccountAddressingPrefixDelegationService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAddressingPrefixDelegationService] method instead.
type AccountAddressingPrefixDelegationService struct {
	Options []option.RequestOption
}

// NewAccountAddressingPrefixDelegationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressingPrefixDelegationService(opts ...option.RequestOption) (r *AccountAddressingPrefixDelegationService) {
	r = &AccountAddressingPrefixDelegationService{}
	r.Options = opts
	return
}

// Create a new account delegation for a given IP prefix.
func (r *AccountAddressingPrefixDelegationService) New(ctx context.Context, accountID string, prefixID string, body AccountAddressingPrefixDelegationNewParams, opts ...option.RequestOption) (res *AccountAddressingPrefixDelegationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all delegations for a given account IP prefix.
func (r *AccountAddressingPrefixDelegationService) List(ctx context.Context, accountID string, prefixID string, opts ...option.RequestOption) (res *AccountAddressingPrefixDelegationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an account delegation for a given IP prefix.
func (r *AccountAddressingPrefixDelegationService) Delete(ctx context.Context, accountID string, prefixID string, delegationID string, opts ...option.RequestOption) (res *AccountAddressingPrefixDelegationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	if delegationID == "" {
		err = errors.New("missing required delegation_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations/%s", accountID, prefixID, delegationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Delegation struct {
	// Identifier of a Delegation.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID string    `json:"delegated_account_id"`
	ModifiedAt         time.Time `json:"modified_at" format:"date-time"`
	// Identifier of an IP Prefix.
	ParentPrefixID string         `json:"parent_prefix_id"`
	JSON           delegationJSON `json:"-"`
}

// delegationJSON contains the JSON metadata for the struct [Delegation]
type delegationJSON struct {
	ID                 apijson.Field
	Cidr               apijson.Field
	CreatedAt          apijson.Field
	DelegatedAccountID apijson.Field
	ModifiedAt         apijson.Field
	ParentPrefixID     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Delegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r delegationJSON) RawJSON() string {
	return r.raw
}

type AccountAddressingPrefixDelegationNewResponse struct {
	Errors   []AddressingMessages `json:"errors,required"`
	Messages []AddressingMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountAddressingPrefixDelegationNewResponseSuccess `json:"success,required"`
	Result  Delegation                                          `json:"result"`
	JSON    accountAddressingPrefixDelegationNewResponseJSON    `json:"-"`
}

// accountAddressingPrefixDelegationNewResponseJSON contains the JSON metadata for
// the struct [AccountAddressingPrefixDelegationNewResponse]
type accountAddressingPrefixDelegationNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixDelegationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingPrefixDelegationNewResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountAddressingPrefixDelegationNewResponseSuccess bool

const (
	AccountAddressingPrefixDelegationNewResponseSuccessTrue AccountAddressingPrefixDelegationNewResponseSuccess = true
)

func (r AccountAddressingPrefixDelegationNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountAddressingPrefixDelegationNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountAddressingPrefixDelegationListResponse struct {
	Errors   []AddressingMessages `json:"errors,required"`
	Messages []AddressingMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountAddressingPrefixDelegationListResponseSuccess    `json:"success,required"`
	Result     []Delegation                                            `json:"result"`
	ResultInfo AccountAddressingPrefixDelegationListResponseResultInfo `json:"result_info"`
	JSON       accountAddressingPrefixDelegationListResponseJSON       `json:"-"`
}

// accountAddressingPrefixDelegationListResponseJSON contains the JSON metadata for
// the struct [AccountAddressingPrefixDelegationListResponse]
type accountAddressingPrefixDelegationListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixDelegationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingPrefixDelegationListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountAddressingPrefixDelegationListResponseSuccess bool

const (
	AccountAddressingPrefixDelegationListResponseSuccessTrue AccountAddressingPrefixDelegationListResponseSuccess = true
)

func (r AccountAddressingPrefixDelegationListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountAddressingPrefixDelegationListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountAddressingPrefixDelegationListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                     `json:"total_count"`
	JSON       accountAddressingPrefixDelegationListResponseResultInfoJSON `json:"-"`
}

// accountAddressingPrefixDelegationListResponseResultInfoJSON contains the JSON
// metadata for the struct
// [AccountAddressingPrefixDelegationListResponseResultInfo]
type accountAddressingPrefixDelegationListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixDelegationListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingPrefixDelegationListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountAddressingPrefixDelegationDeleteResponse struct {
	Errors   []AddressingMessages `json:"errors,required"`
	Messages []AddressingMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountAddressingPrefixDelegationDeleteResponseSuccess `json:"success,required"`
	Result  AccountAddressingPrefixDelegationDeleteResponseResult  `json:"result"`
	JSON    accountAddressingPrefixDelegationDeleteResponseJSON    `json:"-"`
}

// accountAddressingPrefixDelegationDeleteResponseJSON contains the JSON metadata
// for the struct [AccountAddressingPrefixDelegationDeleteResponse]
type accountAddressingPrefixDelegationDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixDelegationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingPrefixDelegationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountAddressingPrefixDelegationDeleteResponseSuccess bool

const (
	AccountAddressingPrefixDelegationDeleteResponseSuccessTrue AccountAddressingPrefixDelegationDeleteResponseSuccess = true
)

func (r AccountAddressingPrefixDelegationDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountAddressingPrefixDelegationDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountAddressingPrefixDelegationDeleteResponseResult struct {
	// Identifier of a Delegation.
	ID   string                                                    `json:"id"`
	JSON accountAddressingPrefixDelegationDeleteResponseResultJSON `json:"-"`
}

// accountAddressingPrefixDelegationDeleteResponseResultJSON contains the JSON
// metadata for the struct [AccountAddressingPrefixDelegationDeleteResponseResult]
type accountAddressingPrefixDelegationDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixDelegationDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingPrefixDelegationDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAddressingPrefixDelegationNewParams struct {
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr,required"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID param.Field[string] `json:"delegated_account_id,required"`
}

func (r AccountAddressingPrefixDelegationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
