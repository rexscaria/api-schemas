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

// AccountPcapOwnershipService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPcapOwnershipService] method instead.
type AccountPcapOwnershipService struct {
	Options []option.RequestOption
}

// NewAccountPcapOwnershipService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountPcapOwnershipService(opts ...option.RequestOption) (r *AccountPcapOwnershipService) {
	r = &AccountPcapOwnershipService{}
	r.Options = opts
	return
}

// Adds an AWS or GCP bucket to use with full packet captures.
func (r *AccountPcapOwnershipService) New(ctx context.Context, accountID string, body AccountPcapOwnershipNewParams, opts ...option.RequestOption) (res *OwnershipSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all buckets configured for use with PCAPs API.
func (r *AccountPcapOwnershipService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountPcapOwnershipListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes buckets added to the packet captures API.
func (r *AccountPcapOwnershipService) Delete(ctx context.Context, accountID string, ownershipID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ownershipID == "" {
		err = errors.New("missing required ownership_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/%s", accountID, ownershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Validates buckets added to the packet captures API.
func (r *AccountPcapOwnershipService) Validate(ctx context.Context, accountID string, body AccountPcapOwnershipValidateParams, opts ...option.RequestOption) (res *OwnershipSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/validate", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MessagesMagicVisibilityPcapsItem struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           MessagesMagicVisibilityPcapsItemSource `json:"source"`
	JSON             messagesMagicVisibilityPcapsItemJSON   `json:"-"`
}

// messagesMagicVisibilityPcapsItemJSON contains the JSON metadata for the struct
// [MessagesMagicVisibilityPcapsItem]
type messagesMagicVisibilityPcapsItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessagesMagicVisibilityPcapsItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesMagicVisibilityPcapsItemJSON) RawJSON() string {
	return r.raw
}

type MessagesMagicVisibilityPcapsItemSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    messagesMagicVisibilityPcapsItemSourceJSON `json:"-"`
}

// messagesMagicVisibilityPcapsItemSourceJSON contains the JSON metadata for the
// struct [MessagesMagicVisibilityPcapsItemSource]
type messagesMagicVisibilityPcapsItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesMagicVisibilityPcapsItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesMagicVisibilityPcapsItemSourceJSON) RawJSON() string {
	return r.raw
}

type OwnershipResponse struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status OwnershipResponseStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                `json:"validated"`
	JSON      ownershipResponseJSON `json:"-"`
}

// ownershipResponseJSON contains the JSON metadata for the struct
// [OwnershipResponse]
type ownershipResponseJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *OwnershipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the ownership challenge. Can be pending, success or failed.
type OwnershipResponseStatus string

const (
	OwnershipResponseStatusPending OwnershipResponseStatus = "pending"
	OwnershipResponseStatusSuccess OwnershipResponseStatus = "success"
	OwnershipResponseStatusFailed  OwnershipResponseStatus = "failed"
)

func (r OwnershipResponseStatus) IsKnown() bool {
	switch r {
	case OwnershipResponseStatusPending, OwnershipResponseStatusSuccess, OwnershipResponseStatusFailed:
		return true
	}
	return false
}

type OwnershipSingleResponse struct {
	Errors   []MessagesMagicVisibilityPcapsItem `json:"errors,required"`
	Messages []MessagesMagicVisibilityPcapsItem `json:"messages,required"`
	Result   OwnershipResponse                  `json:"result,required"`
	// Whether the API call was successful
	Success OwnershipSingleResponseSuccess `json:"success,required"`
	JSON    ownershipSingleResponseJSON    `json:"-"`
}

// ownershipSingleResponseJSON contains the JSON metadata for the struct
// [OwnershipSingleResponse]
type ownershipSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipSingleResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OwnershipSingleResponseSuccess bool

const (
	OwnershipSingleResponseSuccessTrue OwnershipSingleResponseSuccess = true
)

func (r OwnershipSingleResponseSuccess) IsKnown() bool {
	switch r {
	case OwnershipSingleResponseSuccessTrue:
		return true
	}
	return false
}

type AccountPcapOwnershipListResponse struct {
	Errors   []MessagesMagicVisibilityPcapsItem `json:"errors,required"`
	Messages []MessagesMagicVisibilityPcapsItem `json:"messages,required"`
	Result   []OwnershipResponse                `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccountPcapOwnershipListResponseSuccess    `json:"success,required"`
	ResultInfo AccountPcapOwnershipListResponseResultInfo `json:"result_info"`
	JSON       accountPcapOwnershipListResponseJSON       `json:"-"`
}

// accountPcapOwnershipListResponseJSON contains the JSON metadata for the struct
// [AccountPcapOwnershipListResponse]
type accountPcapOwnershipListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapOwnershipListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPcapOwnershipListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountPcapOwnershipListResponseSuccess bool

const (
	AccountPcapOwnershipListResponseSuccessTrue AccountPcapOwnershipListResponseSuccess = true
)

func (r AccountPcapOwnershipListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountPcapOwnershipListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountPcapOwnershipListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       accountPcapOwnershipListResponseResultInfoJSON `json:"-"`
}

// accountPcapOwnershipListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountPcapOwnershipListResponseResultInfo]
type accountPcapOwnershipListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapOwnershipListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPcapOwnershipListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountPcapOwnershipNewParams struct {
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
}

func (r AccountPcapOwnershipNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountPcapOwnershipValidateParams struct {
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
}

func (r AccountPcapOwnershipValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
