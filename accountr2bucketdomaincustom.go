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

// AccountR2BucketDomainCustomService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountR2BucketDomainCustomService] method instead.
type AccountR2BucketDomainCustomService struct {
	Options []option.RequestOption
}

// NewAccountR2BucketDomainCustomService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountR2BucketDomainCustomService(opts ...option.RequestOption) (r *AccountR2BucketDomainCustomService) {
	r = &AccountR2BucketDomainCustomService{}
	r.Options = opts
	return
}

// Get the configuration for a custom domain on an existing R2 bucket.
func (r *AccountR2BucketDomainCustomService) Get(ctx context.Context, accountID string, bucketName string, domain string, query AccountR2BucketDomainCustomGetParams, opts ...option.RequestOption) (res *AccountR2BucketDomainCustomGetResponse, err error) {
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
	if domain == "" {
		err = errors.New("missing required domain parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/custom/%s", accountID, bucketName, domain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit the configuration for a custom domain on an existing R2 bucket.
func (r *AccountR2BucketDomainCustomService) Update(ctx context.Context, accountID string, bucketName string, domain string, params AccountR2BucketDomainCustomUpdateParams, opts ...option.RequestOption) (res *AccountR2BucketDomainCustomUpdateResponse, err error) {
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
	if domain == "" {
		err = errors.New("missing required domain parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/custom/%s", accountID, bucketName, domain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Gets a list of all custom domains registered with an existing R2 bucket.
func (r *AccountR2BucketDomainCustomService) List(ctx context.Context, accountID string, bucketName string, query AccountR2BucketDomainCustomListParams, opts ...option.RequestOption) (res *AccountR2BucketDomainCustomListResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/custom", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Register a new custom domain for an existing R2 bucket.
func (r *AccountR2BucketDomainCustomService) Attach(ctx context.Context, accountID string, bucketName string, params AccountR2BucketDomainCustomAttachParams, opts ...option.RequestOption) (res *AccountR2BucketDomainCustomAttachResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/custom", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Remove custom domain registration from an existing R2 bucket.
func (r *AccountR2BucketDomainCustomService) Remove(ctx context.Context, accountID string, bucketName string, domain string, body AccountR2BucketDomainCustomRemoveParams, opts ...option.RequestOption) (res *AccountR2BucketDomainCustomRemoveResponse, err error) {
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
	if domain == "" {
		err = errors.New("missing required domain parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/custom/%s", accountID, bucketName, domain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountR2BucketDomainCustomGetResponse struct {
	Errors   []AccountR2BucketDomainCustomGetResponseError `json:"errors,required"`
	Messages []string                                      `json:"messages,required"`
	Result   AccountR2BucketDomainCustomGetResponseResult  `json:"result,required"`
	// Whether the API call was successful.
	Success AccountR2BucketDomainCustomGetResponseSuccess `json:"success,required"`
	JSON    accountR2BucketDomainCustomGetResponseJSON    `json:"-"`
}

// accountR2BucketDomainCustomGetResponseJSON contains the JSON metadata for the
// struct [AccountR2BucketDomainCustomGetResponse]
type accountR2BucketDomainCustomGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomGetResponseError struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           AccountR2BucketDomainCustomGetResponseErrorsSource `json:"source"`
	JSON             accountR2BucketDomainCustomGetResponseErrorJSON    `json:"-"`
}

// accountR2BucketDomainCustomGetResponseErrorJSON contains the JSON metadata for
// the struct [AccountR2BucketDomainCustomGetResponseError]
type accountR2BucketDomainCustomGetResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomGetResponseErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    accountR2BucketDomainCustomGetResponseErrorsSourceJSON `json:"-"`
}

// accountR2BucketDomainCustomGetResponseErrorsSourceJSON contains the JSON
// metadata for the struct [AccountR2BucketDomainCustomGetResponseErrorsSource]
type accountR2BucketDomainCustomGetResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomGetResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomGetResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomGetResponseResult struct {
	// Domain name of the custom domain to be added.
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the specified custom domain.
	Enabled bool                                               `json:"enabled,required"`
	Status  AccountR2BucketDomainCustomGetResponseResultStatus `json:"status,required"`
	// Minimum TLS Version the custom domain will accept for incoming connections. If
	// not set, defaults to 1.0.
	MinTls AccountR2BucketDomainCustomGetResponseResultMinTls `json:"minTLS"`
	// Zone ID of the custom domain resides in.
	ZoneID string `json:"zoneId"`
	// Zone that the custom domain resides in.
	ZoneName string                                           `json:"zoneName"`
	JSON     accountR2BucketDomainCustomGetResponseResultJSON `json:"-"`
}

// accountR2BucketDomainCustomGetResponseResultJSON contains the JSON metadata for
// the struct [AccountR2BucketDomainCustomGetResponseResult]
type accountR2BucketDomainCustomGetResponseResultJSON struct {
	Domain      apijson.Field
	Enabled     apijson.Field
	Status      apijson.Field
	MinTls      apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomGetResponseResultStatus struct {
	// Ownership status of the domain.
	Ownership AccountR2BucketDomainCustomGetResponseResultStatusOwnership `json:"ownership,required"`
	// SSL certificate status.
	Ssl  AccountR2BucketDomainCustomGetResponseResultStatusSsl  `json:"ssl,required"`
	JSON accountR2BucketDomainCustomGetResponseResultStatusJSON `json:"-"`
}

// accountR2BucketDomainCustomGetResponseResultStatusJSON contains the JSON
// metadata for the struct [AccountR2BucketDomainCustomGetResponseResultStatus]
type accountR2BucketDomainCustomGetResponseResultStatusJSON struct {
	Ownership   apijson.Field
	Ssl         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomGetResponseResultStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomGetResponseResultStatusJSON) RawJSON() string {
	return r.raw
}

// Ownership status of the domain.
type AccountR2BucketDomainCustomGetResponseResultStatusOwnership string

const (
	AccountR2BucketDomainCustomGetResponseResultStatusOwnershipPending     AccountR2BucketDomainCustomGetResponseResultStatusOwnership = "pending"
	AccountR2BucketDomainCustomGetResponseResultStatusOwnershipActive      AccountR2BucketDomainCustomGetResponseResultStatusOwnership = "active"
	AccountR2BucketDomainCustomGetResponseResultStatusOwnershipDeactivated AccountR2BucketDomainCustomGetResponseResultStatusOwnership = "deactivated"
	AccountR2BucketDomainCustomGetResponseResultStatusOwnershipBlocked     AccountR2BucketDomainCustomGetResponseResultStatusOwnership = "blocked"
	AccountR2BucketDomainCustomGetResponseResultStatusOwnershipError       AccountR2BucketDomainCustomGetResponseResultStatusOwnership = "error"
	AccountR2BucketDomainCustomGetResponseResultStatusOwnershipUnknown     AccountR2BucketDomainCustomGetResponseResultStatusOwnership = "unknown"
)

func (r AccountR2BucketDomainCustomGetResponseResultStatusOwnership) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomGetResponseResultStatusOwnershipPending, AccountR2BucketDomainCustomGetResponseResultStatusOwnershipActive, AccountR2BucketDomainCustomGetResponseResultStatusOwnershipDeactivated, AccountR2BucketDomainCustomGetResponseResultStatusOwnershipBlocked, AccountR2BucketDomainCustomGetResponseResultStatusOwnershipError, AccountR2BucketDomainCustomGetResponseResultStatusOwnershipUnknown:
		return true
	}
	return false
}

// SSL certificate status.
type AccountR2BucketDomainCustomGetResponseResultStatusSsl string

const (
	AccountR2BucketDomainCustomGetResponseResultStatusSslInitializing AccountR2BucketDomainCustomGetResponseResultStatusSsl = "initializing"
	AccountR2BucketDomainCustomGetResponseResultStatusSslPending      AccountR2BucketDomainCustomGetResponseResultStatusSsl = "pending"
	AccountR2BucketDomainCustomGetResponseResultStatusSslActive       AccountR2BucketDomainCustomGetResponseResultStatusSsl = "active"
	AccountR2BucketDomainCustomGetResponseResultStatusSslDeactivated  AccountR2BucketDomainCustomGetResponseResultStatusSsl = "deactivated"
	AccountR2BucketDomainCustomGetResponseResultStatusSslError        AccountR2BucketDomainCustomGetResponseResultStatusSsl = "error"
	AccountR2BucketDomainCustomGetResponseResultStatusSslUnknown      AccountR2BucketDomainCustomGetResponseResultStatusSsl = "unknown"
)

func (r AccountR2BucketDomainCustomGetResponseResultStatusSsl) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomGetResponseResultStatusSslInitializing, AccountR2BucketDomainCustomGetResponseResultStatusSslPending, AccountR2BucketDomainCustomGetResponseResultStatusSslActive, AccountR2BucketDomainCustomGetResponseResultStatusSslDeactivated, AccountR2BucketDomainCustomGetResponseResultStatusSslError, AccountR2BucketDomainCustomGetResponseResultStatusSslUnknown:
		return true
	}
	return false
}

// Minimum TLS Version the custom domain will accept for incoming connections. If
// not set, defaults to 1.0.
type AccountR2BucketDomainCustomGetResponseResultMinTls string

const (
	AccountR2BucketDomainCustomGetResponseResultMinTls1_0 AccountR2BucketDomainCustomGetResponseResultMinTls = "1.0"
	AccountR2BucketDomainCustomGetResponseResultMinTls1_1 AccountR2BucketDomainCustomGetResponseResultMinTls = "1.1"
	AccountR2BucketDomainCustomGetResponseResultMinTls1_2 AccountR2BucketDomainCustomGetResponseResultMinTls = "1.2"
	AccountR2BucketDomainCustomGetResponseResultMinTls1_3 AccountR2BucketDomainCustomGetResponseResultMinTls = "1.3"
)

func (r AccountR2BucketDomainCustomGetResponseResultMinTls) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomGetResponseResultMinTls1_0, AccountR2BucketDomainCustomGetResponseResultMinTls1_1, AccountR2BucketDomainCustomGetResponseResultMinTls1_2, AccountR2BucketDomainCustomGetResponseResultMinTls1_3:
		return true
	}
	return false
}

// Whether the API call was successful.
type AccountR2BucketDomainCustomGetResponseSuccess bool

const (
	AccountR2BucketDomainCustomGetResponseSuccessTrue AccountR2BucketDomainCustomGetResponseSuccess = true
)

func (r AccountR2BucketDomainCustomGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountR2BucketDomainCustomUpdateResponse struct {
	Errors   []AccountR2BucketDomainCustomUpdateResponseError `json:"errors,required"`
	Messages []string                                         `json:"messages,required"`
	Result   AccountR2BucketDomainCustomUpdateResponseResult  `json:"result,required"`
	// Whether the API call was successful.
	Success AccountR2BucketDomainCustomUpdateResponseSuccess `json:"success,required"`
	JSON    accountR2BucketDomainCustomUpdateResponseJSON    `json:"-"`
}

// accountR2BucketDomainCustomUpdateResponseJSON contains the JSON metadata for the
// struct [AccountR2BucketDomainCustomUpdateResponse]
type accountR2BucketDomainCustomUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomUpdateResponseError struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           AccountR2BucketDomainCustomUpdateResponseErrorsSource `json:"source"`
	JSON             accountR2BucketDomainCustomUpdateResponseErrorJSON    `json:"-"`
}

// accountR2BucketDomainCustomUpdateResponseErrorJSON contains the JSON metadata
// for the struct [AccountR2BucketDomainCustomUpdateResponseError]
type accountR2BucketDomainCustomUpdateResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomUpdateResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomUpdateResponseErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    accountR2BucketDomainCustomUpdateResponseErrorsSourceJSON `json:"-"`
}

// accountR2BucketDomainCustomUpdateResponseErrorsSourceJSON contains the JSON
// metadata for the struct [AccountR2BucketDomainCustomUpdateResponseErrorsSource]
type accountR2BucketDomainCustomUpdateResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomUpdateResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomUpdateResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomUpdateResponseResult struct {
	// Domain name of the affected custom domain.
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the specified custom domain.
	Enabled bool `json:"enabled"`
	// Minimum TLS Version the custom domain will accept for incoming connections. If
	// not set, defaults to 1.0.
	MinTls AccountR2BucketDomainCustomUpdateResponseResultMinTls `json:"minTLS"`
	JSON   accountR2BucketDomainCustomUpdateResponseResultJSON   `json:"-"`
}

// accountR2BucketDomainCustomUpdateResponseResultJSON contains the JSON metadata
// for the struct [AccountR2BucketDomainCustomUpdateResponseResult]
type accountR2BucketDomainCustomUpdateResponseResultJSON struct {
	Domain      apijson.Field
	Enabled     apijson.Field
	MinTls      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// Minimum TLS Version the custom domain will accept for incoming connections. If
// not set, defaults to 1.0.
type AccountR2BucketDomainCustomUpdateResponseResultMinTls string

const (
	AccountR2BucketDomainCustomUpdateResponseResultMinTls1_0 AccountR2BucketDomainCustomUpdateResponseResultMinTls = "1.0"
	AccountR2BucketDomainCustomUpdateResponseResultMinTls1_1 AccountR2BucketDomainCustomUpdateResponseResultMinTls = "1.1"
	AccountR2BucketDomainCustomUpdateResponseResultMinTls1_2 AccountR2BucketDomainCustomUpdateResponseResultMinTls = "1.2"
	AccountR2BucketDomainCustomUpdateResponseResultMinTls1_3 AccountR2BucketDomainCustomUpdateResponseResultMinTls = "1.3"
)

func (r AccountR2BucketDomainCustomUpdateResponseResultMinTls) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomUpdateResponseResultMinTls1_0, AccountR2BucketDomainCustomUpdateResponseResultMinTls1_1, AccountR2BucketDomainCustomUpdateResponseResultMinTls1_2, AccountR2BucketDomainCustomUpdateResponseResultMinTls1_3:
		return true
	}
	return false
}

// Whether the API call was successful.
type AccountR2BucketDomainCustomUpdateResponseSuccess bool

const (
	AccountR2BucketDomainCustomUpdateResponseSuccessTrue AccountR2BucketDomainCustomUpdateResponseSuccess = true
)

func (r AccountR2BucketDomainCustomUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountR2BucketDomainCustomListResponse struct {
	Errors   []AccountR2BucketDomainCustomListResponseError `json:"errors,required"`
	Messages []string                                       `json:"messages,required"`
	Result   AccountR2BucketDomainCustomListResponseResult  `json:"result,required"`
	// Whether the API call was successful.
	Success AccountR2BucketDomainCustomListResponseSuccess `json:"success,required"`
	JSON    accountR2BucketDomainCustomListResponseJSON    `json:"-"`
}

// accountR2BucketDomainCustomListResponseJSON contains the JSON metadata for the
// struct [AccountR2BucketDomainCustomListResponse]
type accountR2BucketDomainCustomListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomListResponseError struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           AccountR2BucketDomainCustomListResponseErrorsSource `json:"source"`
	JSON             accountR2BucketDomainCustomListResponseErrorJSON    `json:"-"`
}

// accountR2BucketDomainCustomListResponseErrorJSON contains the JSON metadata for
// the struct [AccountR2BucketDomainCustomListResponseError]
type accountR2BucketDomainCustomListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomListResponseErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    accountR2BucketDomainCustomListResponseErrorsSourceJSON `json:"-"`
}

// accountR2BucketDomainCustomListResponseErrorsSourceJSON contains the JSON
// metadata for the struct [AccountR2BucketDomainCustomListResponseErrorsSource]
type accountR2BucketDomainCustomListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomListResponseResult struct {
	Domains []AccountR2BucketDomainCustomListResponseResultDomain `json:"domains,required"`
	JSON    accountR2BucketDomainCustomListResponseResultJSON     `json:"-"`
}

// accountR2BucketDomainCustomListResponseResultJSON contains the JSON metadata for
// the struct [AccountR2BucketDomainCustomListResponseResult]
type accountR2BucketDomainCustomListResponseResultJSON struct {
	Domains     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomListResponseResultDomain struct {
	// Domain name of the custom domain to be added.
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the specified custom domain.
	Enabled bool                                                       `json:"enabled,required"`
	Status  AccountR2BucketDomainCustomListResponseResultDomainsStatus `json:"status,required"`
	// Minimum TLS Version the custom domain will accept for incoming connections. If
	// not set, defaults to 1.0.
	MinTls AccountR2BucketDomainCustomListResponseResultDomainsMinTls `json:"minTLS"`
	// Zone ID of the custom domain resides in.
	ZoneID string `json:"zoneId"`
	// Zone that the custom domain resides in.
	ZoneName string                                                  `json:"zoneName"`
	JSON     accountR2BucketDomainCustomListResponseResultDomainJSON `json:"-"`
}

// accountR2BucketDomainCustomListResponseResultDomainJSON contains the JSON
// metadata for the struct [AccountR2BucketDomainCustomListResponseResultDomain]
type accountR2BucketDomainCustomListResponseResultDomainJSON struct {
	Domain      apijson.Field
	Enabled     apijson.Field
	Status      apijson.Field
	MinTls      apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomListResponseResultDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomListResponseResultDomainJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomListResponseResultDomainsStatus struct {
	// Ownership status of the domain.
	Ownership AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnership `json:"ownership,required"`
	// SSL certificate status.
	Ssl  AccountR2BucketDomainCustomListResponseResultDomainsStatusSsl  `json:"ssl,required"`
	JSON accountR2BucketDomainCustomListResponseResultDomainsStatusJSON `json:"-"`
}

// accountR2BucketDomainCustomListResponseResultDomainsStatusJSON contains the JSON
// metadata for the struct
// [AccountR2BucketDomainCustomListResponseResultDomainsStatus]
type accountR2BucketDomainCustomListResponseResultDomainsStatusJSON struct {
	Ownership   apijson.Field
	Ssl         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomListResponseResultDomainsStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomListResponseResultDomainsStatusJSON) RawJSON() string {
	return r.raw
}

// Ownership status of the domain.
type AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnership string

const (
	AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnershipPending     AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnership = "pending"
	AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnershipActive      AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnership = "active"
	AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnershipDeactivated AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnership = "deactivated"
	AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnershipBlocked     AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnership = "blocked"
	AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnershipError       AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnership = "error"
	AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnershipUnknown     AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnership = "unknown"
)

func (r AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnership) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnershipPending, AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnershipActive, AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnershipDeactivated, AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnershipBlocked, AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnershipError, AccountR2BucketDomainCustomListResponseResultDomainsStatusOwnershipUnknown:
		return true
	}
	return false
}

// SSL certificate status.
type AccountR2BucketDomainCustomListResponseResultDomainsStatusSsl string

const (
	AccountR2BucketDomainCustomListResponseResultDomainsStatusSslInitializing AccountR2BucketDomainCustomListResponseResultDomainsStatusSsl = "initializing"
	AccountR2BucketDomainCustomListResponseResultDomainsStatusSslPending      AccountR2BucketDomainCustomListResponseResultDomainsStatusSsl = "pending"
	AccountR2BucketDomainCustomListResponseResultDomainsStatusSslActive       AccountR2BucketDomainCustomListResponseResultDomainsStatusSsl = "active"
	AccountR2BucketDomainCustomListResponseResultDomainsStatusSslDeactivated  AccountR2BucketDomainCustomListResponseResultDomainsStatusSsl = "deactivated"
	AccountR2BucketDomainCustomListResponseResultDomainsStatusSslError        AccountR2BucketDomainCustomListResponseResultDomainsStatusSsl = "error"
	AccountR2BucketDomainCustomListResponseResultDomainsStatusSslUnknown      AccountR2BucketDomainCustomListResponseResultDomainsStatusSsl = "unknown"
)

func (r AccountR2BucketDomainCustomListResponseResultDomainsStatusSsl) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomListResponseResultDomainsStatusSslInitializing, AccountR2BucketDomainCustomListResponseResultDomainsStatusSslPending, AccountR2BucketDomainCustomListResponseResultDomainsStatusSslActive, AccountR2BucketDomainCustomListResponseResultDomainsStatusSslDeactivated, AccountR2BucketDomainCustomListResponseResultDomainsStatusSslError, AccountR2BucketDomainCustomListResponseResultDomainsStatusSslUnknown:
		return true
	}
	return false
}

// Minimum TLS Version the custom domain will accept for incoming connections. If
// not set, defaults to 1.0.
type AccountR2BucketDomainCustomListResponseResultDomainsMinTls string

const (
	AccountR2BucketDomainCustomListResponseResultDomainsMinTls1_0 AccountR2BucketDomainCustomListResponseResultDomainsMinTls = "1.0"
	AccountR2BucketDomainCustomListResponseResultDomainsMinTls1_1 AccountR2BucketDomainCustomListResponseResultDomainsMinTls = "1.1"
	AccountR2BucketDomainCustomListResponseResultDomainsMinTls1_2 AccountR2BucketDomainCustomListResponseResultDomainsMinTls = "1.2"
	AccountR2BucketDomainCustomListResponseResultDomainsMinTls1_3 AccountR2BucketDomainCustomListResponseResultDomainsMinTls = "1.3"
)

func (r AccountR2BucketDomainCustomListResponseResultDomainsMinTls) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomListResponseResultDomainsMinTls1_0, AccountR2BucketDomainCustomListResponseResultDomainsMinTls1_1, AccountR2BucketDomainCustomListResponseResultDomainsMinTls1_2, AccountR2BucketDomainCustomListResponseResultDomainsMinTls1_3:
		return true
	}
	return false
}

// Whether the API call was successful.
type AccountR2BucketDomainCustomListResponseSuccess bool

const (
	AccountR2BucketDomainCustomListResponseSuccessTrue AccountR2BucketDomainCustomListResponseSuccess = true
)

func (r AccountR2BucketDomainCustomListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountR2BucketDomainCustomAttachResponse struct {
	Errors   []AccountR2BucketDomainCustomAttachResponseError `json:"errors,required"`
	Messages []string                                         `json:"messages,required"`
	Result   AccountR2BucketDomainCustomAttachResponseResult  `json:"result,required"`
	// Whether the API call was successful.
	Success AccountR2BucketDomainCustomAttachResponseSuccess `json:"success,required"`
	JSON    accountR2BucketDomainCustomAttachResponseJSON    `json:"-"`
}

// accountR2BucketDomainCustomAttachResponseJSON contains the JSON metadata for the
// struct [AccountR2BucketDomainCustomAttachResponse]
type accountR2BucketDomainCustomAttachResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomAttachResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomAttachResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomAttachResponseError struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           AccountR2BucketDomainCustomAttachResponseErrorsSource `json:"source"`
	JSON             accountR2BucketDomainCustomAttachResponseErrorJSON    `json:"-"`
}

// accountR2BucketDomainCustomAttachResponseErrorJSON contains the JSON metadata
// for the struct [AccountR2BucketDomainCustomAttachResponseError]
type accountR2BucketDomainCustomAttachResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomAttachResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomAttachResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomAttachResponseErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    accountR2BucketDomainCustomAttachResponseErrorsSourceJSON `json:"-"`
}

// accountR2BucketDomainCustomAttachResponseErrorsSourceJSON contains the JSON
// metadata for the struct [AccountR2BucketDomainCustomAttachResponseErrorsSource]
type accountR2BucketDomainCustomAttachResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomAttachResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomAttachResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomAttachResponseResult struct {
	// Domain name of the affected custom domain.
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the specified custom domain.
	Enabled bool `json:"enabled,required"`
	// Minimum TLS Version the custom domain will accept for incoming connections. If
	// not set, defaults to 1.0.
	MinTls AccountR2BucketDomainCustomAttachResponseResultMinTls `json:"minTLS"`
	JSON   accountR2BucketDomainCustomAttachResponseResultJSON   `json:"-"`
}

// accountR2BucketDomainCustomAttachResponseResultJSON contains the JSON metadata
// for the struct [AccountR2BucketDomainCustomAttachResponseResult]
type accountR2BucketDomainCustomAttachResponseResultJSON struct {
	Domain      apijson.Field
	Enabled     apijson.Field
	MinTls      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomAttachResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomAttachResponseResultJSON) RawJSON() string {
	return r.raw
}

// Minimum TLS Version the custom domain will accept for incoming connections. If
// not set, defaults to 1.0.
type AccountR2BucketDomainCustomAttachResponseResultMinTls string

const (
	AccountR2BucketDomainCustomAttachResponseResultMinTls1_0 AccountR2BucketDomainCustomAttachResponseResultMinTls = "1.0"
	AccountR2BucketDomainCustomAttachResponseResultMinTls1_1 AccountR2BucketDomainCustomAttachResponseResultMinTls = "1.1"
	AccountR2BucketDomainCustomAttachResponseResultMinTls1_2 AccountR2BucketDomainCustomAttachResponseResultMinTls = "1.2"
	AccountR2BucketDomainCustomAttachResponseResultMinTls1_3 AccountR2BucketDomainCustomAttachResponseResultMinTls = "1.3"
)

func (r AccountR2BucketDomainCustomAttachResponseResultMinTls) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomAttachResponseResultMinTls1_0, AccountR2BucketDomainCustomAttachResponseResultMinTls1_1, AccountR2BucketDomainCustomAttachResponseResultMinTls1_2, AccountR2BucketDomainCustomAttachResponseResultMinTls1_3:
		return true
	}
	return false
}

// Whether the API call was successful.
type AccountR2BucketDomainCustomAttachResponseSuccess bool

const (
	AccountR2BucketDomainCustomAttachResponseSuccessTrue AccountR2BucketDomainCustomAttachResponseSuccess = true
)

func (r AccountR2BucketDomainCustomAttachResponseSuccess) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomAttachResponseSuccessTrue:
		return true
	}
	return false
}

type AccountR2BucketDomainCustomRemoveResponse struct {
	Errors   []AccountR2BucketDomainCustomRemoveResponseError `json:"errors,required"`
	Messages []string                                         `json:"messages,required"`
	Result   AccountR2BucketDomainCustomRemoveResponseResult  `json:"result,required"`
	// Whether the API call was successful.
	Success AccountR2BucketDomainCustomRemoveResponseSuccess `json:"success,required"`
	JSON    accountR2BucketDomainCustomRemoveResponseJSON    `json:"-"`
}

// accountR2BucketDomainCustomRemoveResponseJSON contains the JSON metadata for the
// struct [AccountR2BucketDomainCustomRemoveResponse]
type accountR2BucketDomainCustomRemoveResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomRemoveResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomRemoveResponseError struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           AccountR2BucketDomainCustomRemoveResponseErrorsSource `json:"source"`
	JSON             accountR2BucketDomainCustomRemoveResponseErrorJSON    `json:"-"`
}

// accountR2BucketDomainCustomRemoveResponseErrorJSON contains the JSON metadata
// for the struct [AccountR2BucketDomainCustomRemoveResponseError]
type accountR2BucketDomainCustomRemoveResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomRemoveResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomRemoveResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomRemoveResponseErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    accountR2BucketDomainCustomRemoveResponseErrorsSourceJSON `json:"-"`
}

// accountR2BucketDomainCustomRemoveResponseErrorsSourceJSON contains the JSON
// metadata for the struct [AccountR2BucketDomainCustomRemoveResponseErrorsSource]
type accountR2BucketDomainCustomRemoveResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomRemoveResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomRemoveResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketDomainCustomRemoveResponseResult struct {
	// Name of the removed custom domain.
	Domain string                                              `json:"domain,required"`
	JSON   accountR2BucketDomainCustomRemoveResponseResultJSON `json:"-"`
}

// accountR2BucketDomainCustomRemoveResponseResultJSON contains the JSON metadata
// for the struct [AccountR2BucketDomainCustomRemoveResponseResult]
type accountR2BucketDomainCustomRemoveResponseResultJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDomainCustomRemoveResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketDomainCustomRemoveResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountR2BucketDomainCustomRemoveResponseSuccess bool

const (
	AccountR2BucketDomainCustomRemoveResponseSuccessTrue AccountR2BucketDomainCustomRemoveResponseSuccess = true
)

func (r AccountR2BucketDomainCustomRemoveResponseSuccess) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomRemoveResponseSuccessTrue:
		return true
	}
	return false
}

type AccountR2BucketDomainCustomGetParams struct {
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[AccountR2BucketDomainCustomGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type AccountR2BucketDomainCustomGetParamsCfR2Jurisdiction string

const (
	AccountR2BucketDomainCustomGetParamsCfR2JurisdictionDefault AccountR2BucketDomainCustomGetParamsCfR2Jurisdiction = "default"
	AccountR2BucketDomainCustomGetParamsCfR2JurisdictionEu      AccountR2BucketDomainCustomGetParamsCfR2Jurisdiction = "eu"
	AccountR2BucketDomainCustomGetParamsCfR2JurisdictionFedramp AccountR2BucketDomainCustomGetParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketDomainCustomGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomGetParamsCfR2JurisdictionDefault, AccountR2BucketDomainCustomGetParamsCfR2JurisdictionEu, AccountR2BucketDomainCustomGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type AccountR2BucketDomainCustomUpdateParams struct {
	// Whether to enable public bucket access at the specified custom domain.
	Enabled param.Field[bool] `json:"enabled"`
	// Minimum TLS Version the custom domain will accept for incoming connections. If
	// not set, defaults to previous value.
	MinTls param.Field[AccountR2BucketDomainCustomUpdateParamsMinTls] `json:"minTLS"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[AccountR2BucketDomainCustomUpdateParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r AccountR2BucketDomainCustomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Minimum TLS Version the custom domain will accept for incoming connections. If
// not set, defaults to previous value.
type AccountR2BucketDomainCustomUpdateParamsMinTls string

const (
	AccountR2BucketDomainCustomUpdateParamsMinTls1_0 AccountR2BucketDomainCustomUpdateParamsMinTls = "1.0"
	AccountR2BucketDomainCustomUpdateParamsMinTls1_1 AccountR2BucketDomainCustomUpdateParamsMinTls = "1.1"
	AccountR2BucketDomainCustomUpdateParamsMinTls1_2 AccountR2BucketDomainCustomUpdateParamsMinTls = "1.2"
	AccountR2BucketDomainCustomUpdateParamsMinTls1_3 AccountR2BucketDomainCustomUpdateParamsMinTls = "1.3"
)

func (r AccountR2BucketDomainCustomUpdateParamsMinTls) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomUpdateParamsMinTls1_0, AccountR2BucketDomainCustomUpdateParamsMinTls1_1, AccountR2BucketDomainCustomUpdateParamsMinTls1_2, AccountR2BucketDomainCustomUpdateParamsMinTls1_3:
		return true
	}
	return false
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type AccountR2BucketDomainCustomUpdateParamsCfR2Jurisdiction string

const (
	AccountR2BucketDomainCustomUpdateParamsCfR2JurisdictionDefault AccountR2BucketDomainCustomUpdateParamsCfR2Jurisdiction = "default"
	AccountR2BucketDomainCustomUpdateParamsCfR2JurisdictionEu      AccountR2BucketDomainCustomUpdateParamsCfR2Jurisdiction = "eu"
	AccountR2BucketDomainCustomUpdateParamsCfR2JurisdictionFedramp AccountR2BucketDomainCustomUpdateParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketDomainCustomUpdateParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomUpdateParamsCfR2JurisdictionDefault, AccountR2BucketDomainCustomUpdateParamsCfR2JurisdictionEu, AccountR2BucketDomainCustomUpdateParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type AccountR2BucketDomainCustomListParams struct {
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[AccountR2BucketDomainCustomListParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type AccountR2BucketDomainCustomListParamsCfR2Jurisdiction string

const (
	AccountR2BucketDomainCustomListParamsCfR2JurisdictionDefault AccountR2BucketDomainCustomListParamsCfR2Jurisdiction = "default"
	AccountR2BucketDomainCustomListParamsCfR2JurisdictionEu      AccountR2BucketDomainCustomListParamsCfR2Jurisdiction = "eu"
	AccountR2BucketDomainCustomListParamsCfR2JurisdictionFedramp AccountR2BucketDomainCustomListParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketDomainCustomListParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomListParamsCfR2JurisdictionDefault, AccountR2BucketDomainCustomListParamsCfR2JurisdictionEu, AccountR2BucketDomainCustomListParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type AccountR2BucketDomainCustomAttachParams struct {
	// Name of the custom domain to be added.
	Domain param.Field[string] `json:"domain,required"`
	// Whether to enable public bucket access at the custom domain. If undefined, the
	// domain will be enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Zone ID of the custom domain.
	ZoneID param.Field[string] `json:"zoneId,required"`
	// Minimum TLS Version the custom domain will accept for incoming connections. If
	// not set, defaults to 1.0.
	MinTls param.Field[AccountR2BucketDomainCustomAttachParamsMinTls] `json:"minTLS"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[AccountR2BucketDomainCustomAttachParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r AccountR2BucketDomainCustomAttachParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Minimum TLS Version the custom domain will accept for incoming connections. If
// not set, defaults to 1.0.
type AccountR2BucketDomainCustomAttachParamsMinTls string

const (
	AccountR2BucketDomainCustomAttachParamsMinTls1_0 AccountR2BucketDomainCustomAttachParamsMinTls = "1.0"
	AccountR2BucketDomainCustomAttachParamsMinTls1_1 AccountR2BucketDomainCustomAttachParamsMinTls = "1.1"
	AccountR2BucketDomainCustomAttachParamsMinTls1_2 AccountR2BucketDomainCustomAttachParamsMinTls = "1.2"
	AccountR2BucketDomainCustomAttachParamsMinTls1_3 AccountR2BucketDomainCustomAttachParamsMinTls = "1.3"
)

func (r AccountR2BucketDomainCustomAttachParamsMinTls) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomAttachParamsMinTls1_0, AccountR2BucketDomainCustomAttachParamsMinTls1_1, AccountR2BucketDomainCustomAttachParamsMinTls1_2, AccountR2BucketDomainCustomAttachParamsMinTls1_3:
		return true
	}
	return false
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type AccountR2BucketDomainCustomAttachParamsCfR2Jurisdiction string

const (
	AccountR2BucketDomainCustomAttachParamsCfR2JurisdictionDefault AccountR2BucketDomainCustomAttachParamsCfR2Jurisdiction = "default"
	AccountR2BucketDomainCustomAttachParamsCfR2JurisdictionEu      AccountR2BucketDomainCustomAttachParamsCfR2Jurisdiction = "eu"
	AccountR2BucketDomainCustomAttachParamsCfR2JurisdictionFedramp AccountR2BucketDomainCustomAttachParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketDomainCustomAttachParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomAttachParamsCfR2JurisdictionDefault, AccountR2BucketDomainCustomAttachParamsCfR2JurisdictionEu, AccountR2BucketDomainCustomAttachParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type AccountR2BucketDomainCustomRemoveParams struct {
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[AccountR2BucketDomainCustomRemoveParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type AccountR2BucketDomainCustomRemoveParamsCfR2Jurisdiction string

const (
	AccountR2BucketDomainCustomRemoveParamsCfR2JurisdictionDefault AccountR2BucketDomainCustomRemoveParamsCfR2Jurisdiction = "default"
	AccountR2BucketDomainCustomRemoveParamsCfR2JurisdictionEu      AccountR2BucketDomainCustomRemoveParamsCfR2Jurisdiction = "eu"
	AccountR2BucketDomainCustomRemoveParamsCfR2JurisdictionFedramp AccountR2BucketDomainCustomRemoveParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketDomainCustomRemoveParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketDomainCustomRemoveParamsCfR2JurisdictionDefault, AccountR2BucketDomainCustomRemoveParamsCfR2JurisdictionEu, AccountR2BucketDomainCustomRemoveParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}
