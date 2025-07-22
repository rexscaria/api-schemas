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

// AccountR2BucketService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountR2BucketService] method instead.
type AccountR2BucketService struct {
	Options   []option.RequestOption
	Cors      *AccountR2BucketCorService
	Domains   *AccountR2BucketDomainService
	Lifecycle *AccountR2BucketLifecycleService
	Lock      *AccountR2BucketLockService
	Sippy     *AccountR2BucketSippyService
}

// NewAccountR2BucketService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountR2BucketService(opts ...option.RequestOption) (r *AccountR2BucketService) {
	r = &AccountR2BucketService{}
	r.Options = opts
	r.Cors = NewAccountR2BucketCorService(opts...)
	r.Domains = NewAccountR2BucketDomainService(opts...)
	r.Lifecycle = NewAccountR2BucketLifecycleService(opts...)
	r.Lock = NewAccountR2BucketLockService(opts...)
	r.Sippy = NewAccountR2BucketSippyService(opts...)
	return
}

// Creates a new R2 bucket.
func (r *AccountR2BucketService) New(ctx context.Context, accountID string, params AccountR2BucketNewParams, opts ...option.RequestOption) (res *AccountR2BucketNewResponse, err error) {
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Gets metadata for an existing R2 bucket.
func (r *AccountR2BucketService) Get(ctx context.Context, accountID string, bucketName string, query AccountR2BucketGetParams, opts ...option.RequestOption) (res *AccountR2BucketGetResponse, err error) {
	if query.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", query.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all R2 buckets on your account
func (r *AccountR2BucketService) List(ctx context.Context, accountID string, params AccountR2BucketListParams, opts ...option.RequestOption) (res *AccountR2BucketListResponse, err error) {
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Deletes an existing R2 bucket.
func (r *AccountR2BucketService) Delete(ctx context.Context, accountID string, bucketName string, body AccountR2BucketDeleteParams, opts ...option.RequestOption) (res *R2V4Response, err error) {
	if body.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", body.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// A single R2 bucket
type R2Bucket struct {
	// Creation timestamp
	CreationDate string `json:"creation_date"`
	// Location of the bucket
	Location R2BucketLocation `json:"location"`
	// Name of the bucket
	Name string `json:"name"`
	// Storage class for newly uploaded objects, unless specified otherwise.
	StorageClass R2StorageClass `json:"storage_class"`
	JSON         r2BucketJSON   `json:"-"`
}

// r2BucketJSON contains the JSON metadata for the struct [R2Bucket]
type r2BucketJSON struct {
	CreationDate apijson.Field
	Location     apijson.Field
	Name         apijson.Field
	StorageClass apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *R2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2BucketJSON) RawJSON() string {
	return r.raw
}

// Location of the bucket
type R2BucketLocation string

const (
	R2BucketLocationApac R2BucketLocation = "apac"
	R2BucketLocationEeur R2BucketLocation = "eeur"
	R2BucketLocationEnam R2BucketLocation = "enam"
	R2BucketLocationWeur R2BucketLocation = "weur"
	R2BucketLocationWnam R2BucketLocation = "wnam"
	R2BucketLocationOc   R2BucketLocation = "oc"
)

func (r R2BucketLocation) IsKnown() bool {
	switch r {
	case R2BucketLocationApac, R2BucketLocationEeur, R2BucketLocationEnam, R2BucketLocationWeur, R2BucketLocationWnam, R2BucketLocationOc:
		return true
	}
	return false
}

// Storage class for newly uploaded objects, unless specified otherwise.
type R2StorageClass string

const (
	R2StorageClassStandard         R2StorageClass = "Standard"
	R2StorageClassInfrequentAccess R2StorageClass = "InfrequentAccess"
)

func (r R2StorageClass) IsKnown() bool {
	switch r {
	case R2StorageClassStandard, R2StorageClassInfrequentAccess:
		return true
	}
	return false
}

type R2V4Response struct {
	Errors   []R2V4ResponseError `json:"errors,required"`
	Messages []string            `json:"messages,required"`
	Result   interface{}         `json:"result,required"`
	// Whether the API call was successful
	Success R2V4ResponseSuccess `json:"success,required"`
	JSON    r2V4ResponseJSON    `json:"-"`
}

// r2V4ResponseJSON contains the JSON metadata for the struct [R2V4Response]
type r2V4ResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2V4Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2V4ResponseJSON) RawJSON() string {
	return r.raw
}

type R2V4ResponseError struct {
	Code    int64                 `json:"code,required"`
	Message string                `json:"message,required"`
	JSON    r2V4ResponseErrorJSON `json:"-"`
}

// r2V4ResponseErrorJSON contains the JSON metadata for the struct
// [R2V4ResponseError]
type r2V4ResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2V4ResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2V4ResponseErrorJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type R2V4ResponseSuccess bool

const (
	R2V4ResponseSuccessTrue R2V4ResponseSuccess = true
)

func (r R2V4ResponseSuccess) IsKnown() bool {
	switch r {
	case R2V4ResponseSuccessTrue:
		return true
	}
	return false
}

type AccountR2BucketNewResponse struct {
	// A single R2 bucket
	Result R2Bucket                       `json:"result"`
	JSON   accountR2BucketNewResponseJSON `json:"-"`
	R2V4Response
}

// accountR2BucketNewResponseJSON contains the JSON metadata for the struct
// [AccountR2BucketNewResponse]
type accountR2BucketNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketGetResponse struct {
	// A single R2 bucket
	Result R2Bucket                       `json:"result"`
	JSON   accountR2BucketGetResponseJSON `json:"-"`
	R2V4Response
}

// accountR2BucketGetResponseJSON contains the JSON metadata for the struct
// [AccountR2BucketGetResponse]
type accountR2BucketGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketListResponse struct {
	Result     AccountR2BucketListResponseResult     `json:"result"`
	ResultInfo AccountR2BucketListResponseResultInfo `json:"result_info"`
	JSON       accountR2BucketListResponseJSON       `json:"-"`
	R2V4Response
}

// accountR2BucketListResponseJSON contains the JSON metadata for the struct
// [AccountR2BucketListResponse]
type accountR2BucketListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketListResponseResult struct {
	Buckets []R2Bucket                            `json:"buckets"`
	JSON    accountR2BucketListResponseResultJSON `json:"-"`
}

// accountR2BucketListResponseResultJSON contains the JSON metadata for the struct
// [AccountR2BucketListResponseResult]
type accountR2BucketListResponseResultJSON struct {
	Buckets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketListResponseResultInfo struct {
	// A continuation token that should be used to fetch the next page of results
	Cursor string `json:"cursor"`
	// Maximum number of results on this page
	PerPage float64                                   `json:"per_page"`
	JSON    accountR2BucketListResponseResultInfoJSON `json:"-"`
}

// accountR2BucketListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountR2BucketListResponseResultInfo]
type accountR2BucketListResponseResultInfoJSON struct {
	Cursor      apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketNewParams struct {
	// Name of the bucket
	Name param.Field[string] `json:"name,required"`
	// Location of the bucket
	LocationHint param.Field[R2BucketLocation] `json:"locationHint"`
	// Storage class for newly uploaded objects, unless specified otherwise.
	StorageClass param.Field[R2StorageClass] `json:"storageClass"`
	// Creates the bucket in the provided jurisdiction
	Jurisdiction param.Field[AccountR2BucketNewParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r AccountR2BucketNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Creates the bucket in the provided jurisdiction
type AccountR2BucketNewParamsCfR2Jurisdiction string

const (
	AccountR2BucketNewParamsCfR2JurisdictionDefault AccountR2BucketNewParamsCfR2Jurisdiction = "default"
	AccountR2BucketNewParamsCfR2JurisdictionEu      AccountR2BucketNewParamsCfR2Jurisdiction = "eu"
	AccountR2BucketNewParamsCfR2JurisdictionFedramp AccountR2BucketNewParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketNewParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketNewParamsCfR2JurisdictionDefault, AccountR2BucketNewParamsCfR2JurisdictionEu, AccountR2BucketNewParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type AccountR2BucketGetParams struct {
	// The bucket jurisdiction
	Jurisdiction param.Field[AccountR2BucketGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction
type AccountR2BucketGetParamsCfR2Jurisdiction string

const (
	AccountR2BucketGetParamsCfR2JurisdictionDefault AccountR2BucketGetParamsCfR2Jurisdiction = "default"
	AccountR2BucketGetParamsCfR2JurisdictionEu      AccountR2BucketGetParamsCfR2Jurisdiction = "eu"
	AccountR2BucketGetParamsCfR2JurisdictionFedramp AccountR2BucketGetParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketGetParamsCfR2JurisdictionDefault, AccountR2BucketGetParamsCfR2JurisdictionEu, AccountR2BucketGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type AccountR2BucketListParams struct {
	// Pagination cursor received during the last List Buckets call. R2 buckets are
	// paginated using cursors instead of page numbers.
	Cursor param.Field[string] `query:"cursor"`
	// Direction to order buckets
	Direction param.Field[AccountR2BucketListParamsDirection] `query:"direction"`
	// Bucket names to filter by. Only buckets with this phrase in their name will be
	// returned.
	NameContains param.Field[string] `query:"name_contains"`
	// Field to order buckets by
	Order param.Field[AccountR2BucketListParamsOrder] `query:"order"`
	// Maximum number of buckets to return in a single call
	PerPage param.Field[float64] `query:"per_page"`
	// Bucket name to start searching after. Buckets are ordered lexicographically.
	StartAfter param.Field[string] `query:"start_after"`
	// Lists buckets in the provided jurisdiction
	Jurisdiction param.Field[AccountR2BucketListParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// URLQuery serializes [AccountR2BucketListParams]'s query parameters as
// `url.Values`.
func (r AccountR2BucketListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order buckets
type AccountR2BucketListParamsDirection string

const (
	AccountR2BucketListParamsDirectionAsc  AccountR2BucketListParamsDirection = "asc"
	AccountR2BucketListParamsDirectionDesc AccountR2BucketListParamsDirection = "desc"
)

func (r AccountR2BucketListParamsDirection) IsKnown() bool {
	switch r {
	case AccountR2BucketListParamsDirectionAsc, AccountR2BucketListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order buckets by
type AccountR2BucketListParamsOrder string

const (
	AccountR2BucketListParamsOrderName AccountR2BucketListParamsOrder = "name"
)

func (r AccountR2BucketListParamsOrder) IsKnown() bool {
	switch r {
	case AccountR2BucketListParamsOrderName:
		return true
	}
	return false
}

// Lists buckets in the provided jurisdiction
type AccountR2BucketListParamsCfR2Jurisdiction string

const (
	AccountR2BucketListParamsCfR2JurisdictionDefault AccountR2BucketListParamsCfR2Jurisdiction = "default"
	AccountR2BucketListParamsCfR2JurisdictionEu      AccountR2BucketListParamsCfR2Jurisdiction = "eu"
	AccountR2BucketListParamsCfR2JurisdictionFedramp AccountR2BucketListParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketListParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketListParamsCfR2JurisdictionDefault, AccountR2BucketListParamsCfR2JurisdictionEu, AccountR2BucketListParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type AccountR2BucketDeleteParams struct {
	// The bucket jurisdiction
	Jurisdiction param.Field[AccountR2BucketDeleteParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction
type AccountR2BucketDeleteParamsCfR2Jurisdiction string

const (
	AccountR2BucketDeleteParamsCfR2JurisdictionDefault AccountR2BucketDeleteParamsCfR2Jurisdiction = "default"
	AccountR2BucketDeleteParamsCfR2JurisdictionEu      AccountR2BucketDeleteParamsCfR2Jurisdiction = "eu"
	AccountR2BucketDeleteParamsCfR2JurisdictionFedramp AccountR2BucketDeleteParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketDeleteParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketDeleteParamsCfR2JurisdictionDefault, AccountR2BucketDeleteParamsCfR2JurisdictionEu, AccountR2BucketDeleteParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}
