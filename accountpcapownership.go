// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
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

type APIResponseMagicVisibilityPcaps struct {
	Errors   []MessagesMagicVisibilityPcapsItem         `json:"errors,required"`
	Messages []MessagesMagicVisibilityPcapsItem         `json:"messages,required"`
	Result   APIResponseMagicVisibilityPcapsResultUnion `json:"result,required"`
	// Whether the API call was successful
	Success APIResponseMagicVisibilityPcapsSuccess `json:"success,required"`
	JSON    apiResponseMagicVisibilityPcapsJSON    `json:"-"`
}

// apiResponseMagicVisibilityPcapsJSON contains the JSON metadata for the struct
// [APIResponseMagicVisibilityPcaps]
type apiResponseMagicVisibilityPcapsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseMagicVisibilityPcaps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseMagicVisibilityPcapsJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [APIResponseMagicVisibilityPcapsResultArray] or
// [shared.UnionString].
type APIResponseMagicVisibilityPcapsResultUnion interface {
	ImplementsAPIResponseMagicVisibilityPcapsResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIResponseMagicVisibilityPcapsResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIResponseMagicVisibilityPcapsResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type APIResponseMagicVisibilityPcapsResultArray []interface{}

func (r APIResponseMagicVisibilityPcapsResultArray) ImplementsAPIResponseMagicVisibilityPcapsResultUnion() {
}

// Whether the API call was successful
type APIResponseMagicVisibilityPcapsSuccess bool

const (
	APIResponseMagicVisibilityPcapsSuccessTrue APIResponseMagicVisibilityPcapsSuccess = true
)

func (r APIResponseMagicVisibilityPcapsSuccess) IsKnown() bool {
	switch r {
	case APIResponseMagicVisibilityPcapsSuccessTrue:
		return true
	}
	return false
}

type MessagesMagicVisibilityPcapsItem struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    messagesMagicVisibilityPcapsItemJSON `json:"-"`
}

// messagesMagicVisibilityPcapsItemJSON contains the JSON metadata for the struct
// [MessagesMagicVisibilityPcapsItem]
type messagesMagicVisibilityPcapsItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesMagicVisibilityPcapsItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesMagicVisibilityPcapsItemJSON) RawJSON() string {
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
	Result OwnershipResponse           `json:"result"`
	JSON   ownershipSingleResponseJSON `json:"-"`
	APIResponseMagicVisibilityPcaps
}

// ownershipSingleResponseJSON contains the JSON metadata for the struct
// [OwnershipSingleResponse]
type ownershipSingleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipSingleResponseJSON) RawJSON() string {
	return r.raw
}

type AccountPcapOwnershipListResponse struct {
	Result []OwnershipResponse                  `json:"result,nullable"`
	JSON   accountPcapOwnershipListResponseJSON `json:"-"`
	APIResponseCollectionMagicVisibilityPcaps
}

// accountPcapOwnershipListResponseJSON contains the JSON metadata for the struct
// [AccountPcapOwnershipListResponse]
type accountPcapOwnershipListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapOwnershipListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPcapOwnershipListResponseJSON) RawJSON() string {
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
