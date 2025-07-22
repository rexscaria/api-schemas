// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountR2BucketSippyService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountR2BucketSippyService] method instead.
type AccountR2BucketSippyService struct {
	Options []option.RequestOption
}

// NewAccountR2BucketSippyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountR2BucketSippyService(opts ...option.RequestOption) (r *AccountR2BucketSippyService) {
	r = &AccountR2BucketSippyService{}
	r.Options = opts
	return
}

// Gets configuration for Sippy for an existing R2 bucket.
func (r *AccountR2BucketSippyService) Get(ctx context.Context, accountID string, bucketName string, query AccountR2BucketSippyGetParams, opts ...option.RequestOption) (res *AccountR2BucketSippyGetResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Disables Sippy on this bucket
func (r *AccountR2BucketSippyService) Disable(ctx context.Context, accountID string, bucketName string, body AccountR2BucketSippyDisableParams, opts ...option.RequestOption) (res *AccountR2BucketSippyDisableResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Sets configuration for Sippy for an existing R2 bucket.
func (r *AccountR2BucketSippyService) Enable(ctx context.Context, accountID string, bucketName string, params AccountR2BucketSippyEnableParams, opts ...option.RequestOption) (res *AccountR2BucketSippyEnableResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type R2Sippy struct {
	// Details about the configured destination bucket
	Destination R2SippyDestination `json:"destination"`
	// State of Sippy for this bucket
	Enabled bool `json:"enabled"`
	// Details about the configured source bucket
	Source R2SippySource `json:"source"`
	JSON   r2SippyJSON   `json:"-"`
}

// r2SippyJSON contains the JSON metadata for the struct [R2Sippy]
type r2SippyJSON struct {
	Destination apijson.Field
	Enabled     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2Sippy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2SippyJSON) RawJSON() string {
	return r.raw
}

// Details about the configured destination bucket
type R2SippyDestination struct {
	// ID of the Cloudflare API token used when writing objects to this bucket
	AccessKeyID string `json:"accessKeyId"`
	Account     string `json:"account"`
	// Name of the bucket on the provider
	Bucket   string                     `json:"bucket"`
	Provider R2SippyDestinationProvider `json:"provider"`
	JSON     r2SippyDestinationJSON     `json:"-"`
}

// r2SippyDestinationJSON contains the JSON metadata for the struct
// [R2SippyDestination]
type r2SippyDestinationJSON struct {
	AccessKeyID apijson.Field
	Account     apijson.Field
	Bucket      apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippyDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2SippyDestinationJSON) RawJSON() string {
	return r.raw
}

type R2SippyDestinationProvider string

const (
	R2SippyDestinationProviderR2 R2SippyDestinationProvider = "r2"
)

func (r R2SippyDestinationProvider) IsKnown() bool {
	switch r {
	case R2SippyDestinationProviderR2:
		return true
	}
	return false
}

// Details about the configured source bucket
type R2SippySource struct {
	// Name of the bucket on the provider
	Bucket   string                `json:"bucket"`
	Provider R2SippySourceProvider `json:"provider"`
	// Region where the bucket resides (AWS only)
	Region string            `json:"region,nullable"`
	JSON   r2SippySourceJSON `json:"-"`
}

// r2SippySourceJSON contains the JSON metadata for the struct [R2SippySource]
type r2SippySourceJSON struct {
	Bucket      apijson.Field
	Provider    apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippySource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2SippySourceJSON) RawJSON() string {
	return r.raw
}

type R2SippySourceProvider string

const (
	R2SippySourceProviderAws R2SippySourceProvider = "aws"
	R2SippySourceProviderGcs R2SippySourceProvider = "gcs"
)

func (r R2SippySourceProvider) IsKnown() bool {
	switch r {
	case R2SippySourceProviderAws, R2SippySourceProviderGcs:
		return true
	}
	return false
}

type AccountR2BucketSippyGetResponse struct {
	Result R2Sippy                             `json:"result"`
	JSON   accountR2BucketSippyGetResponseJSON `json:"-"`
	R2V4Response
}

// accountR2BucketSippyGetResponseJSON contains the JSON metadata for the struct
// [AccountR2BucketSippyGetResponse]
type accountR2BucketSippyGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketSippyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketSippyGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketSippyDisableResponse struct {
	Result AccountR2BucketSippyDisableResponseResult `json:"result"`
	JSON   accountR2BucketSippyDisableResponseJSON   `json:"-"`
	R2V4Response
}

// accountR2BucketSippyDisableResponseJSON contains the JSON metadata for the
// struct [AccountR2BucketSippyDisableResponse]
type accountR2BucketSippyDisableResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketSippyDisableResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketSippyDisableResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketSippyDisableResponseResult struct {
	Enabled AccountR2BucketSippyDisableResponseResultEnabled `json:"enabled"`
	JSON    accountR2BucketSippyDisableResponseResultJSON    `json:"-"`
}

// accountR2BucketSippyDisableResponseResultJSON contains the JSON metadata for the
// struct [AccountR2BucketSippyDisableResponseResult]
type accountR2BucketSippyDisableResponseResultJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketSippyDisableResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketSippyDisableResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketSippyDisableResponseResultEnabled bool

const (
	AccountR2BucketSippyDisableResponseResultEnabledFalse AccountR2BucketSippyDisableResponseResultEnabled = false
)

func (r AccountR2BucketSippyDisableResponseResultEnabled) IsKnown() bool {
	switch r {
	case AccountR2BucketSippyDisableResponseResultEnabledFalse:
		return true
	}
	return false
}

type AccountR2BucketSippyEnableResponse struct {
	Result R2Sippy                                `json:"result"`
	JSON   accountR2BucketSippyEnableResponseJSON `json:"-"`
	R2V4Response
}

// accountR2BucketSippyEnableResponseJSON contains the JSON metadata for the struct
// [AccountR2BucketSippyEnableResponse]
type accountR2BucketSippyEnableResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketSippyEnableResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketSippyEnableResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketSippyGetParams struct {
	// The bucket jurisdiction
	Jurisdiction param.Field[AccountR2BucketSippyGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction
type AccountR2BucketSippyGetParamsCfR2Jurisdiction string

const (
	AccountR2BucketSippyGetParamsCfR2JurisdictionDefault AccountR2BucketSippyGetParamsCfR2Jurisdiction = "default"
	AccountR2BucketSippyGetParamsCfR2JurisdictionEu      AccountR2BucketSippyGetParamsCfR2Jurisdiction = "eu"
	AccountR2BucketSippyGetParamsCfR2JurisdictionFedramp AccountR2BucketSippyGetParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketSippyGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketSippyGetParamsCfR2JurisdictionDefault, AccountR2BucketSippyGetParamsCfR2JurisdictionEu, AccountR2BucketSippyGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type AccountR2BucketSippyDisableParams struct {
	// The bucket jurisdiction
	Jurisdiction param.Field[AccountR2BucketSippyDisableParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction
type AccountR2BucketSippyDisableParamsCfR2Jurisdiction string

const (
	AccountR2BucketSippyDisableParamsCfR2JurisdictionDefault AccountR2BucketSippyDisableParamsCfR2Jurisdiction = "default"
	AccountR2BucketSippyDisableParamsCfR2JurisdictionEu      AccountR2BucketSippyDisableParamsCfR2Jurisdiction = "eu"
	AccountR2BucketSippyDisableParamsCfR2JurisdictionFedramp AccountR2BucketSippyDisableParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketSippyDisableParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketSippyDisableParamsCfR2JurisdictionDefault, AccountR2BucketSippyDisableParamsCfR2JurisdictionEu, AccountR2BucketSippyDisableParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type AccountR2BucketSippyEnableParams struct {
	Body AccountR2BucketSippyEnableParamsBodyUnion `json:"body,required"`
	// The bucket jurisdiction
	Jurisdiction param.Field[AccountR2BucketSippyEnableParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r AccountR2BucketSippyEnableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountR2BucketSippyEnableParamsBody struct {
	Destination param.Field[interface{}] `json:"destination"`
	Source      param.Field[interface{}] `json:"source"`
}

func (r AccountR2BucketSippyEnableParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountR2BucketSippyEnableParamsBody) implementsAccountR2BucketSippyEnableParamsBodyUnion() {}

// Satisfied by [AccountR2BucketSippyEnableParamsBodyR2EnableSippyAws],
// [AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcs],
// [AccountR2BucketSippyEnableParamsBody].
type AccountR2BucketSippyEnableParamsBodyUnion interface {
	implementsAccountR2BucketSippyEnableParamsBodyUnion()
}

type AccountR2BucketSippyEnableParamsBodyR2EnableSippyAws struct {
	// R2 bucket to copy objects to
	Destination param.Field[AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsDestination] `json:"destination"`
	// AWS S3 bucket to copy objects from
	Source param.Field[AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsSource] `json:"source"`
}

func (r AccountR2BucketSippyEnableParamsBodyR2EnableSippyAws) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountR2BucketSippyEnableParamsBodyR2EnableSippyAws) implementsAccountR2BucketSippyEnableParamsBodyUnion() {
}

// R2 bucket to copy objects to
type AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsDestination struct {
	// ID of a Cloudflare API token. This is the value labelled "Access Key ID" when
	// creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	AccessKeyID param.Field[string]                                                                  `json:"accessKeyId"`
	Provider    param.Field[AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsDestinationProvider] `json:"provider"`
	// Value of a Cloudflare API token. This is the value labelled "Secret Access Key"
	// when creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsDestinationProvider string

const (
	AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsDestinationProviderR2 AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsDestinationProvider = "r2"
)

func (r AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsDestinationProvider) IsKnown() bool {
	switch r {
	case AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsDestinationProviderR2:
		return true
	}
	return false
}

// AWS S3 bucket to copy objects from
type AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsSource struct {
	// Access Key ID of an IAM credential (ideally scoped to a single S3 bucket)
	AccessKeyID param.Field[string] `json:"accessKeyId"`
	// Name of the AWS S3 bucket
	Bucket   param.Field[string]                                                             `json:"bucket"`
	Provider param.Field[AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsSourceProvider] `json:"provider"`
	// Name of the AWS availability zone
	Region param.Field[string] `json:"region"`
	// Secret Access Key of an IAM credential (ideally scoped to a single S3 bucket)
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsSourceProvider string

const (
	AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsSourceProviderAws AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsSourceProvider = "aws"
)

func (r AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsSourceProvider) IsKnown() bool {
	switch r {
	case AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsSourceProviderAws:
		return true
	}
	return false
}

type AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcs struct {
	// R2 bucket to copy objects to
	Destination param.Field[AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsDestination] `json:"destination"`
	// GCS bucket to copy objects from
	Source param.Field[AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsSource] `json:"source"`
}

func (r AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcs) implementsAccountR2BucketSippyEnableParamsBodyUnion() {
}

// R2 bucket to copy objects to
type AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsDestination struct {
	// ID of a Cloudflare API token. This is the value labelled "Access Key ID" when
	// creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	AccessKeyID param.Field[string]                                                                  `json:"accessKeyId"`
	Provider    param.Field[AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsDestinationProvider] `json:"provider"`
	// Value of a Cloudflare API token. This is the value labelled "Secret Access Key"
	// when creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsDestinationProvider string

const (
	AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsDestinationProviderR2 AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsDestinationProvider = "r2"
)

func (r AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsDestinationProvider) IsKnown() bool {
	switch r {
	case AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsDestinationProviderR2:
		return true
	}
	return false
}

// GCS bucket to copy objects from
type AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsSource struct {
	// Name of the GCS bucket
	Bucket param.Field[string] `json:"bucket"`
	// Client email of an IAM credential (ideally scoped to a single GCS bucket)
	ClientEmail param.Field[string] `json:"clientEmail"`
	// Private Key of an IAM credential (ideally scoped to a single GCS bucket)
	PrivateKey param.Field[string]                                                             `json:"privateKey"`
	Provider   param.Field[AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsSourceProvider] `json:"provider"`
}

func (r AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsSourceProvider string

const (
	AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsSourceProviderGcs AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsSourceProvider = "gcs"
)

func (r AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsSourceProvider) IsKnown() bool {
	switch r {
	case AccountR2BucketSippyEnableParamsBodyR2EnableSippyGcsSourceProviderGcs:
		return true
	}
	return false
}

// The bucket jurisdiction
type AccountR2BucketSippyEnableParamsCfR2Jurisdiction string

const (
	AccountR2BucketSippyEnableParamsCfR2JurisdictionDefault AccountR2BucketSippyEnableParamsCfR2Jurisdiction = "default"
	AccountR2BucketSippyEnableParamsCfR2JurisdictionEu      AccountR2BucketSippyEnableParamsCfR2Jurisdiction = "eu"
	AccountR2BucketSippyEnableParamsCfR2JurisdictionFedramp AccountR2BucketSippyEnableParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketSippyEnableParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketSippyEnableParamsCfR2JurisdictionDefault, AccountR2BucketSippyEnableParamsCfR2JurisdictionEu, AccountR2BucketSippyEnableParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}
