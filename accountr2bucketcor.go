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

// AccountR2BucketCorService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountR2BucketCorService] method instead.
type AccountR2BucketCorService struct {
	Options []option.RequestOption
}

// NewAccountR2BucketCorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountR2BucketCorService(opts ...option.RequestOption) (r *AccountR2BucketCorService) {
	r = &AccountR2BucketCorService{}
	r.Options = opts
	return
}

// Get the CORS policy for a bucket
func (r *AccountR2BucketCorService) Get(ctx context.Context, accountID string, bucketName string, query AccountR2BucketCorGetParams, opts ...option.RequestOption) (res *AccountR2BucketCorGetResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/cors", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set the CORS policy for a bucket
func (r *AccountR2BucketCorService) Update(ctx context.Context, accountID string, bucketName string, params AccountR2BucketCorUpdateParams, opts ...option.RequestOption) (res *AccountR2BucketCorUpdateResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/cors", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Delete the CORS policy for a bucket
func (r *AccountR2BucketCorService) Delete(ctx context.Context, accountID string, bucketName string, body AccountR2BucketCorDeleteParams, opts ...option.RequestOption) (res *AccountR2BucketCorDeleteResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/cors", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type R2CorsRule struct {
	// Object specifying allowed origins, methods and headers for this CORS rule.
	Allowed R2CorsRuleAllowed `json:"allowed,required"`
	// Identifier for this rule
	ID string `json:"id"`
	// Specifies the headers that can be exposed back, and accessed by, the JavaScript
	// making the cross-origin request. If you need to access headers beyond the
	// safelisted response headers, such as Content-Encoding or cf-cache-status, you
	// must specify it here.
	ExposeHeaders []string `json:"exposeHeaders"`
	// Specifies the amount of time (in seconds) browsers are allowed to cache CORS
	// preflight responses. Browsers may limit this to 2 hours or less, even if the
	// maximum value (86400) is specified.
	MaxAgeSeconds float64        `json:"maxAgeSeconds"`
	JSON          r2CorsRuleJSON `json:"-"`
}

// r2CorsRuleJSON contains the JSON metadata for the struct [R2CorsRule]
type r2CorsRuleJSON struct {
	Allowed       apijson.Field
	ID            apijson.Field
	ExposeHeaders apijson.Field
	MaxAgeSeconds apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *R2CorsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2CorsRuleJSON) RawJSON() string {
	return r.raw
}

// Object specifying allowed origins, methods and headers for this CORS rule.
type R2CorsRuleAllowed struct {
	// Specifies the value for the Access-Control-Allow-Methods header R2 sets when
	// requesting objects in a bucket from a browser.
	Methods []R2CorsRuleAllowedMethod `json:"methods,required"`
	// Specifies the value for the Access-Control-Allow-Origin header R2 sets when
	// requesting objects in a bucket from a browser.
	Origins []string `json:"origins,required"`
	// Specifies the value for the Access-Control-Allow-Headers header R2 sets when
	// requesting objects in this bucket from a browser. Cross-origin requests that
	// include custom headers (e.g. x-user-id) should specify these headers as
	// AllowedHeaders.
	Headers []string              `json:"headers"`
	JSON    r2CorsRuleAllowedJSON `json:"-"`
}

// r2CorsRuleAllowedJSON contains the JSON metadata for the struct
// [R2CorsRuleAllowed]
type r2CorsRuleAllowedJSON struct {
	Methods     apijson.Field
	Origins     apijson.Field
	Headers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2CorsRuleAllowed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2CorsRuleAllowedJSON) RawJSON() string {
	return r.raw
}

type R2CorsRuleAllowedMethod string

const (
	R2CorsRuleAllowedMethodGet    R2CorsRuleAllowedMethod = "GET"
	R2CorsRuleAllowedMethodPut    R2CorsRuleAllowedMethod = "PUT"
	R2CorsRuleAllowedMethodPost   R2CorsRuleAllowedMethod = "POST"
	R2CorsRuleAllowedMethodDelete R2CorsRuleAllowedMethod = "DELETE"
	R2CorsRuleAllowedMethodHead   R2CorsRuleAllowedMethod = "HEAD"
)

func (r R2CorsRuleAllowedMethod) IsKnown() bool {
	switch r {
	case R2CorsRuleAllowedMethodGet, R2CorsRuleAllowedMethodPut, R2CorsRuleAllowedMethodPost, R2CorsRuleAllowedMethodDelete, R2CorsRuleAllowedMethodHead:
		return true
	}
	return false
}

type R2CorsRuleParam struct {
	// Object specifying allowed origins, methods and headers for this CORS rule.
	Allowed param.Field[R2CorsRuleAllowedParam] `json:"allowed,required"`
	// Identifier for this rule
	ID param.Field[string] `json:"id"`
	// Specifies the headers that can be exposed back, and accessed by, the JavaScript
	// making the cross-origin request. If you need to access headers beyond the
	// safelisted response headers, such as Content-Encoding or cf-cache-status, you
	// must specify it here.
	ExposeHeaders param.Field[[]string] `json:"exposeHeaders"`
	// Specifies the amount of time (in seconds) browsers are allowed to cache CORS
	// preflight responses. Browsers may limit this to 2 hours or less, even if the
	// maximum value (86400) is specified.
	MaxAgeSeconds param.Field[float64] `json:"maxAgeSeconds"`
}

func (r R2CorsRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object specifying allowed origins, methods and headers for this CORS rule.
type R2CorsRuleAllowedParam struct {
	// Specifies the value for the Access-Control-Allow-Methods header R2 sets when
	// requesting objects in a bucket from a browser.
	Methods param.Field[[]R2CorsRuleAllowedMethod] `json:"methods,required"`
	// Specifies the value for the Access-Control-Allow-Origin header R2 sets when
	// requesting objects in a bucket from a browser.
	Origins param.Field[[]string] `json:"origins,required"`
	// Specifies the value for the Access-Control-Allow-Headers header R2 sets when
	// requesting objects in this bucket from a browser. Cross-origin requests that
	// include custom headers (e.g. x-user-id) should specify these headers as
	// AllowedHeaders.
	Headers param.Field[[]string] `json:"headers"`
}

func (r R2CorsRuleAllowedParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountR2BucketCorGetResponse struct {
	Result AccountR2BucketCorGetResponseResult `json:"result"`
	JSON   accountR2BucketCorGetResponseJSON   `json:"-"`
	R2V4Response
}

// accountR2BucketCorGetResponseJSON contains the JSON metadata for the struct
// [AccountR2BucketCorGetResponse]
type accountR2BucketCorGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketCorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketCorGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketCorGetResponseResult struct {
	Rules []R2CorsRule                            `json:"rules"`
	JSON  accountR2BucketCorGetResponseResultJSON `json:"-"`
}

// accountR2BucketCorGetResponseResultJSON contains the JSON metadata for the
// struct [AccountR2BucketCorGetResponseResult]
type accountR2BucketCorGetResponseResultJSON struct {
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketCorGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketCorGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketCorUpdateResponse struct {
	JSON accountR2BucketCorUpdateResponseJSON `json:"-"`
	R2V4Response
}

// accountR2BucketCorUpdateResponseJSON contains the JSON metadata for the struct
// [AccountR2BucketCorUpdateResponse]
type accountR2BucketCorUpdateResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketCorUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketCorUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketCorDeleteResponse struct {
	JSON accountR2BucketCorDeleteResponseJSON `json:"-"`
	R2V4Response
}

// accountR2BucketCorDeleteResponseJSON contains the JSON metadata for the struct
// [AccountR2BucketCorDeleteResponse]
type accountR2BucketCorDeleteResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketCorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketCorDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketCorGetParams struct {
	// The bucket jurisdiction
	Jurisdiction param.Field[AccountR2BucketCorGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction
type AccountR2BucketCorGetParamsCfR2Jurisdiction string

const (
	AccountR2BucketCorGetParamsCfR2JurisdictionDefault AccountR2BucketCorGetParamsCfR2Jurisdiction = "default"
	AccountR2BucketCorGetParamsCfR2JurisdictionEu      AccountR2BucketCorGetParamsCfR2Jurisdiction = "eu"
	AccountR2BucketCorGetParamsCfR2JurisdictionFedramp AccountR2BucketCorGetParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketCorGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketCorGetParamsCfR2JurisdictionDefault, AccountR2BucketCorGetParamsCfR2JurisdictionEu, AccountR2BucketCorGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type AccountR2BucketCorUpdateParams struct {
	Rules param.Field[[]R2CorsRuleParam] `json:"rules"`
	// The bucket jurisdiction
	Jurisdiction param.Field[AccountR2BucketCorUpdateParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r AccountR2BucketCorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The bucket jurisdiction
type AccountR2BucketCorUpdateParamsCfR2Jurisdiction string

const (
	AccountR2BucketCorUpdateParamsCfR2JurisdictionDefault AccountR2BucketCorUpdateParamsCfR2Jurisdiction = "default"
	AccountR2BucketCorUpdateParamsCfR2JurisdictionEu      AccountR2BucketCorUpdateParamsCfR2Jurisdiction = "eu"
	AccountR2BucketCorUpdateParamsCfR2JurisdictionFedramp AccountR2BucketCorUpdateParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketCorUpdateParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketCorUpdateParamsCfR2JurisdictionDefault, AccountR2BucketCorUpdateParamsCfR2JurisdictionEu, AccountR2BucketCorUpdateParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type AccountR2BucketCorDeleteParams struct {
	// The bucket jurisdiction
	Jurisdiction param.Field[AccountR2BucketCorDeleteParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction
type AccountR2BucketCorDeleteParamsCfR2Jurisdiction string

const (
	AccountR2BucketCorDeleteParamsCfR2JurisdictionDefault AccountR2BucketCorDeleteParamsCfR2Jurisdiction = "default"
	AccountR2BucketCorDeleteParamsCfR2JurisdictionEu      AccountR2BucketCorDeleteParamsCfR2Jurisdiction = "eu"
	AccountR2BucketCorDeleteParamsCfR2JurisdictionFedramp AccountR2BucketCorDeleteParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketCorDeleteParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketCorDeleteParamsCfR2JurisdictionDefault, AccountR2BucketCorDeleteParamsCfR2JurisdictionEu, AccountR2BucketCorDeleteParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}
