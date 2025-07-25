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

// AccountSlurperTargetService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountSlurperTargetService] method instead.
type AccountSlurperTargetService struct {
	Options []option.RequestOption
}

// NewAccountSlurperTargetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountSlurperTargetService(opts ...option.RequestOption) (r *AccountSlurperTargetService) {
	r = &AccountSlurperTargetService{}
	r.Options = opts
	return
}

// Check whether tokens are valid against the target bucket
func (r *AccountSlurperTargetService) CheckConnectivity(ctx context.Context, accountID string, body AccountSlurperTargetCheckConnectivityParams, opts ...option.RequestOption) (res *AccountSlurperTargetCheckConnectivityResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/target/connectivity-precheck", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type R2TargetSchemaParam struct {
	Bucket       param.Field[string]                 `json:"bucket"`
	Jurisdiction param.Field[Jurisdiction]           `json:"jurisdiction"`
	Secret       param.Field[S3LikeCredsSchemaParam] `json:"secret"`
	Vendor       param.Field[R2TargetSchemaVendor]   `json:"vendor"`
}

func (r R2TargetSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type R2TargetSchemaVendor string

const (
	R2TargetSchemaVendorR2 R2TargetSchemaVendor = "r2"
)

func (r R2TargetSchemaVendor) IsKnown() bool {
	switch r {
	case R2TargetSchemaVendorR2:
		return true
	}
	return false
}

type AccountSlurperTargetCheckConnectivityResponse struct {
	Errors   []AccountSlurperTargetCheckConnectivityResponseError `json:"errors"`
	Messages []string                                             `json:"messages"`
	Result   ConnectivityResponse                                 `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountSlurperTargetCheckConnectivityResponseSuccess `json:"success"`
	JSON    accountSlurperTargetCheckConnectivityResponseJSON    `json:"-"`
}

// accountSlurperTargetCheckConnectivityResponseJSON contains the JSON metadata for
// the struct [AccountSlurperTargetCheckConnectivityResponse]
type accountSlurperTargetCheckConnectivityResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperTargetCheckConnectivityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperTargetCheckConnectivityResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperTargetCheckConnectivityResponseError struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AccountSlurperTargetCheckConnectivityResponseErrorsSource `json:"source"`
	JSON             accountSlurperTargetCheckConnectivityResponseErrorJSON    `json:"-"`
}

// accountSlurperTargetCheckConnectivityResponseErrorJSON contains the JSON
// metadata for the struct [AccountSlurperTargetCheckConnectivityResponseError]
type accountSlurperTargetCheckConnectivityResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountSlurperTargetCheckConnectivityResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperTargetCheckConnectivityResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperTargetCheckConnectivityResponseErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    accountSlurperTargetCheckConnectivityResponseErrorsSourceJSON `json:"-"`
}

// accountSlurperTargetCheckConnectivityResponseErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccountSlurperTargetCheckConnectivityResponseErrorsSource]
type accountSlurperTargetCheckConnectivityResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperTargetCheckConnectivityResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperTargetCheckConnectivityResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountSlurperTargetCheckConnectivityResponseSuccess bool

const (
	AccountSlurperTargetCheckConnectivityResponseSuccessTrue AccountSlurperTargetCheckConnectivityResponseSuccess = true
)

func (r AccountSlurperTargetCheckConnectivityResponseSuccess) IsKnown() bool {
	switch r {
	case AccountSlurperTargetCheckConnectivityResponseSuccessTrue:
		return true
	}
	return false
}

type AccountSlurperTargetCheckConnectivityParams struct {
	R2TargetSchema R2TargetSchemaParam `json:"r2_target_schema,required"`
}

func (r AccountSlurperTargetCheckConnectivityParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.R2TargetSchema)
}
