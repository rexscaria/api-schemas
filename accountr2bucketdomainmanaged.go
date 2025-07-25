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

// AccountR2BucketDomainManagedService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountR2BucketDomainManagedService] method instead.
type AccountR2BucketDomainManagedService struct {
	Options []option.RequestOption
}

// NewAccountR2BucketDomainManagedService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountR2BucketDomainManagedService(opts ...option.RequestOption) (r *AccountR2BucketDomainManagedService) {
	r = &AccountR2BucketDomainManagedService{}
	r.Options = opts
	return
}

// Gets state of public access over the bucket's R2-managed (r2.dev) domain.
func (r *AccountR2BucketDomainManagedService) Get(ctx context.Context, accountID string, bucketName string, query AccountR2BucketDomainManagedGetParams, opts ...option.RequestOption) (res *AccountR2BucketDomainManagedGetResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/managed", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates state of public access over the bucket's R2-managed (r2.dev) domain.
func (r *AccountR2BucketDomainManagedService) Update(ctx context.Context, accountID string, bucketName string, params AccountR2BucketDomainManagedUpdateParams, opts ...option.RequestOption) (res *AccountR2BucketDomainManagedUpdateResponse, err error) {
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/managed", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type R2ManagedDomainResponse struct {
	// Bucket ID.
	BucketID string `json:"bucketId,required"`
	// Domain name of the bucket's r2.dev domain.
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the r2.dev domain.
	Enabled bool                        `json:"enabled,required"`
	JSON    r2ManagedDomainResponseJSON `json:"-"`
}

// r2ManagedDomainResponseJSON contains the JSON metadata for the struct
// [R2ManagedDomainResponse]
type r2ManagedDomainResponseJSON struct {
	BucketID    apijson.Field
	Domain      apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2ManagedDomainResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2ManagedDomainResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainManagedGetResponse struct {
	Errors   []AccountR2BucketDomainManagedGetResponseError `json:"errors,required"`
	Messages []string                                       `json:"messages,required"`
	Result   R2ManagedDomainResponse                        `json:"result,required"`
	// Whether the API call was successful.
	Success AccountR2BucketDomainManagedGetResponseSuccess `json:"success,required"`
	JSON    accountR2BucketDomainManagedGetResponseJSON    `json:"-"`
}

// accountR2BucketDomainManagedGetResponseJSON contains the JSON metadata for the
// struct [AccountR2BucketDomainManagedGetResponse]
type accountR2BucketDomainManagedGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainManagedGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainManagedGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainManagedGetResponseError struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           AccountR2BucketDomainManagedGetResponseErrorsSource `json:"source"`
	JSON             accountR2BucketDomainManagedGetResponseErrorJSON    `json:"-"`
}

// accountR2BucketDomainManagedGetResponseErrorJSON contains the JSON metadata for
// the struct [AccountR2BucketDomainManagedGetResponseError]
type accountR2BucketDomainManagedGetResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountR2BucketDomainManagedGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainManagedGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainManagedGetResponseErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    accountR2BucketDomainManagedGetResponseErrorsSourceJSON `json:"-"`
}

// accountR2BucketDomainManagedGetResponseErrorsSourceJSON contains the JSON
// metadata for the struct [AccountR2BucketDomainManagedGetResponseErrorsSource]
type accountR2BucketDomainManagedGetResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainManagedGetResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainManagedGetResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountR2BucketDomainManagedGetResponseSuccess bool

const (
	AccountR2BucketDomainManagedGetResponseSuccessTrue AccountR2BucketDomainManagedGetResponseSuccess = true
)

func (r AccountR2BucketDomainManagedGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainManagedGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountR2BucketDomainManagedUpdateResponse struct {
	Errors   []AccountR2BucketDomainManagedUpdateResponseError `json:"errors,required"`
	Messages []string                                          `json:"messages,required"`
	Result   R2ManagedDomainResponse                           `json:"result,required"`
	// Whether the API call was successful.
	Success AccountR2BucketDomainManagedUpdateResponseSuccess `json:"success,required"`
	JSON    accountR2BucketDomainManagedUpdateResponseJSON    `json:"-"`
}

// accountR2BucketDomainManagedUpdateResponseJSON contains the JSON metadata for
// the struct [AccountR2BucketDomainManagedUpdateResponse]
type accountR2BucketDomainManagedUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainManagedUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainManagedUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainManagedUpdateResponseError struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AccountR2BucketDomainManagedUpdateResponseErrorsSource `json:"source"`
	JSON             accountR2BucketDomainManagedUpdateResponseErrorJSON    `json:"-"`
}

// accountR2BucketDomainManagedUpdateResponseErrorJSON contains the JSON metadata
// for the struct [AccountR2BucketDomainManagedUpdateResponseError]
type accountR2BucketDomainManagedUpdateResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountR2BucketDomainManagedUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainManagedUpdateResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainManagedUpdateResponseErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    accountR2BucketDomainManagedUpdateResponseErrorsSourceJSON `json:"-"`
}

// accountR2BucketDomainManagedUpdateResponseErrorsSourceJSON contains the JSON
// metadata for the struct [AccountR2BucketDomainManagedUpdateResponseErrorsSource]
type accountR2BucketDomainManagedUpdateResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainManagedUpdateResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainManagedUpdateResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountR2BucketDomainManagedUpdateResponseSuccess bool

const (
	AccountR2BucketDomainManagedUpdateResponseSuccessTrue AccountR2BucketDomainManagedUpdateResponseSuccess = true
)

func (r AccountR2BucketDomainManagedUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainManagedUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountR2BucketDomainManagedGetParams struct {
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[AccountR2BucketDomainManagedGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type AccountR2BucketDomainManagedGetParamsCfR2Jurisdiction string

const (
	AccountR2BucketDomainManagedGetParamsCfR2JurisdictionDefault AccountR2BucketDomainManagedGetParamsCfR2Jurisdiction = "default"
	AccountR2BucketDomainManagedGetParamsCfR2JurisdictionEu      AccountR2BucketDomainManagedGetParamsCfR2Jurisdiction = "eu"
	AccountR2BucketDomainManagedGetParamsCfR2JurisdictionFedramp AccountR2BucketDomainManagedGetParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketDomainManagedGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainManagedGetParamsCfR2JurisdictionDefault, AccountR2BucketDomainManagedGetParamsCfR2JurisdictionEu, AccountR2BucketDomainManagedGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type AccountR2BucketDomainManagedUpdateParams struct {
	// Whether to enable public bucket access at the r2.dev domain.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[AccountR2BucketDomainManagedUpdateParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r AccountR2BucketDomainManagedUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type AccountR2BucketDomainManagedUpdateParamsCfR2Jurisdiction string

const (
	AccountR2BucketDomainManagedUpdateParamsCfR2JurisdictionDefault AccountR2BucketDomainManagedUpdateParamsCfR2Jurisdiction = "default"
	AccountR2BucketDomainManagedUpdateParamsCfR2JurisdictionEu      AccountR2BucketDomainManagedUpdateParamsCfR2Jurisdiction = "eu"
	AccountR2BucketDomainManagedUpdateParamsCfR2JurisdictionFedramp AccountR2BucketDomainManagedUpdateParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketDomainManagedUpdateParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainManagedUpdateParamsCfR2JurisdictionDefault, AccountR2BucketDomainManagedUpdateParamsCfR2JurisdictionEu, AccountR2BucketDomainManagedUpdateParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}
