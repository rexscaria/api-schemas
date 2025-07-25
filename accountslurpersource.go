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

// AccountSlurperSourceService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountSlurperSourceService] method instead.
type AccountSlurperSourceService struct {
	Options []option.RequestOption
}

// NewAccountSlurperSourceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountSlurperSourceService(opts ...option.RequestOption) (r *AccountSlurperSourceService) {
	r = &AccountSlurperSourceService{}
	r.Options = opts
	return
}

// Check whether tokens are valid against the source bucket
func (r *AccountSlurperSourceService) CheckConnectivity(ctx context.Context, accountID string, body AccountSlurperSourceCheckConnectivityParams, opts ...option.RequestOption) (res *AccountSlurperSourceCheckConnectivityResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/source/connectivity-precheck", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ConnectivityResponse struct {
	ConnectivityStatus ConnectivityResponseConnectivityStatus `json:"connectivityStatus"`
	JSON               connectivityResponseJSON               `json:"-"`
}

// connectivityResponseJSON contains the JSON metadata for the struct
// [ConnectivityResponse]
type connectivityResponseJSON struct {
	ConnectivityStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ConnectivityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectivityResponseConnectivityStatus string

const (
	ConnectivityResponseConnectivityStatusSuccess ConnectivityResponseConnectivityStatus = "success"
	ConnectivityResponseConnectivityStatusError   ConnectivityResponseConnectivityStatus = "error"
)

func (r ConnectivityResponseConnectivityStatus) IsKnown() bool {
	switch r {
	case ConnectivityResponseConnectivityStatusSuccess, ConnectivityResponseConnectivityStatusError:
		return true
	}
	return false
}

type R2SourceSchemaParam struct {
	Bucket       param.Field[string]                 `json:"bucket"`
	Jurisdiction param.Field[Jurisdiction]           `json:"jurisdiction"`
	Secret       param.Field[S3LikeCredsSchemaParam] `json:"secret"`
	Vendor       param.Field[R2SourceSchemaVendor]   `json:"vendor"`
}

func (r R2SourceSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r R2SourceSchemaParam) implementsSourceJobSchemaUnionParam() {}

type R2SourceSchemaVendor string

const (
	R2SourceSchemaVendorR2 R2SourceSchemaVendor = "r2"
)

func (r R2SourceSchemaVendor) IsKnown() bool {
	switch r {
	case R2SourceSchemaVendorR2:
		return true
	}
	return false
}

type S3LikeCredsSchemaParam struct {
	AccessKeyID     param.Field[string] `json:"accessKeyId"`
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r S3LikeCredsSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SourceJobSchemaParam struct {
	Bucket       param.Field[string]                `json:"bucket"`
	Endpoint     param.Field[string]                `json:"endpoint"`
	Jurisdiction param.Field[Jurisdiction]          `json:"jurisdiction"`
	Secret       param.Field[interface{}]           `json:"secret"`
	Vendor       param.Field[SourceJobSchemaVendor] `json:"vendor"`
}

func (r SourceJobSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SourceJobSchemaParam) implementsSourceJobSchemaUnionParam() {}

// Satisfied by [SourceJobSchemaR2SlurperS3SourceSchemaParam],
// [SourceJobSchemaR2SlurperGcsSourceSchemaParam], [R2SourceSchemaParam],
// [SourceJobSchemaParam].
type SourceJobSchemaUnionParam interface {
	implementsSourceJobSchemaUnionParam()
}

type SourceJobSchemaR2SlurperS3SourceSchemaParam struct {
	Bucket   param.Field[string]                                       `json:"bucket"`
	Endpoint param.Field[string]                                       `json:"endpoint"`
	Secret   param.Field[S3LikeCredsSchemaParam]                       `json:"secret"`
	Vendor   param.Field[SourceJobSchemaR2SlurperS3SourceSchemaVendor] `json:"vendor"`
}

func (r SourceJobSchemaR2SlurperS3SourceSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SourceJobSchemaR2SlurperS3SourceSchemaParam) implementsSourceJobSchemaUnionParam() {}

type SourceJobSchemaR2SlurperS3SourceSchemaVendor string

const (
	SourceJobSchemaR2SlurperS3SourceSchemaVendorS3 SourceJobSchemaR2SlurperS3SourceSchemaVendor = "s3"
)

func (r SourceJobSchemaR2SlurperS3SourceSchemaVendor) IsKnown() bool {
	switch r {
	case SourceJobSchemaR2SlurperS3SourceSchemaVendorS3:
		return true
	}
	return false
}

type SourceJobSchemaR2SlurperGcsSourceSchemaParam struct {
	Bucket param.Field[string]                                             `json:"bucket"`
	Secret param.Field[SourceJobSchemaR2SlurperGcsSourceSchemaSecretParam] `json:"secret"`
	Vendor param.Field[SourceJobSchemaR2SlurperGcsSourceSchemaVendor]      `json:"vendor"`
}

func (r SourceJobSchemaR2SlurperGcsSourceSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SourceJobSchemaR2SlurperGcsSourceSchemaParam) implementsSourceJobSchemaUnionParam() {}

type SourceJobSchemaR2SlurperGcsSourceSchemaSecretParam struct {
	ClientEmail param.Field[string] `json:"clientEmail"`
	PrivateKey  param.Field[string] `json:"privateKey"`
}

func (r SourceJobSchemaR2SlurperGcsSourceSchemaSecretParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SourceJobSchemaR2SlurperGcsSourceSchemaVendor string

const (
	SourceJobSchemaR2SlurperGcsSourceSchemaVendorGcs SourceJobSchemaR2SlurperGcsSourceSchemaVendor = "gcs"
)

func (r SourceJobSchemaR2SlurperGcsSourceSchemaVendor) IsKnown() bool {
	switch r {
	case SourceJobSchemaR2SlurperGcsSourceSchemaVendorGcs:
		return true
	}
	return false
}

type SourceJobSchemaVendor string

const (
	SourceJobSchemaVendorS3  SourceJobSchemaVendor = "s3"
	SourceJobSchemaVendorGcs SourceJobSchemaVendor = "gcs"
	SourceJobSchemaVendorR2  SourceJobSchemaVendor = "r2"
)

func (r SourceJobSchemaVendor) IsKnown() bool {
	switch r {
	case SourceJobSchemaVendorS3, SourceJobSchemaVendorGcs, SourceJobSchemaVendorR2:
		return true
	}
	return false
}

type AccountSlurperSourceCheckConnectivityResponse struct {
	Errors   []AccountSlurperSourceCheckConnectivityResponseError `json:"errors"`
	Messages []string                                             `json:"messages"`
	Result   ConnectivityResponse                                 `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountSlurperSourceCheckConnectivityResponseSuccess `json:"success"`
	JSON    accountSlurperSourceCheckConnectivityResponseJSON    `json:"-"`
}

// accountSlurperSourceCheckConnectivityResponseJSON contains the JSON metadata for
// the struct [AccountSlurperSourceCheckConnectivityResponse]
type accountSlurperSourceCheckConnectivityResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperSourceCheckConnectivityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperSourceCheckConnectivityResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperSourceCheckConnectivityResponseError struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AccountSlurperSourceCheckConnectivityResponseErrorsSource `json:"source"`
	JSON             accountSlurperSourceCheckConnectivityResponseErrorJSON    `json:"-"`
}

// accountSlurperSourceCheckConnectivityResponseErrorJSON contains the JSON
// metadata for the struct [AccountSlurperSourceCheckConnectivityResponseError]
type accountSlurperSourceCheckConnectivityResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountSlurperSourceCheckConnectivityResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperSourceCheckConnectivityResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperSourceCheckConnectivityResponseErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    accountSlurperSourceCheckConnectivityResponseErrorsSourceJSON `json:"-"`
}

// accountSlurperSourceCheckConnectivityResponseErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccountSlurperSourceCheckConnectivityResponseErrorsSource]
type accountSlurperSourceCheckConnectivityResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperSourceCheckConnectivityResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperSourceCheckConnectivityResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountSlurperSourceCheckConnectivityResponseSuccess bool

const (
	AccountSlurperSourceCheckConnectivityResponseSuccessTrue AccountSlurperSourceCheckConnectivityResponseSuccess = true
)

func (r AccountSlurperSourceCheckConnectivityResponseSuccess) IsKnown() bool {
	switch r {
	case AccountSlurperSourceCheckConnectivityResponseSuccessTrue:
		return true
	}
	return false
}

type AccountSlurperSourceCheckConnectivityParams struct {
	SourceJobSchema SourceJobSchemaUnionParam `json:"source_job_schema,required"`
}

func (r AccountSlurperSourceCheckConnectivityParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SourceJobSchema)
}
