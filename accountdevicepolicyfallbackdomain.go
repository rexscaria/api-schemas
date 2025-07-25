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

// AccountDevicePolicyFallbackDomainService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDevicePolicyFallbackDomainService] method instead.
type AccountDevicePolicyFallbackDomainService struct {
	Options []option.RequestOption
}

// NewAccountDevicePolicyFallbackDomainService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDevicePolicyFallbackDomainService(opts ...option.RequestOption) (r *AccountDevicePolicyFallbackDomainService) {
	r = &AccountDevicePolicyFallbackDomainService{}
	r.Options = opts
	return
}

// Fetches the list of domains to bypass Gateway DNS resolution from a specified
// device settings profile. These domains will use the specified local DNS resolver
// instead.
func (r *AccountDevicePolicyFallbackDomainService) List(ctx context.Context, accountID string, policyID string, opts ...option.RequestOption) (res *FallbackDomainResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/fallback_domains", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead.
func (r *AccountDevicePolicyFallbackDomainService) GlobalList(ctx context.Context, accountID string, opts ...option.RequestOption) (res *FallbackDomainResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/fallback_domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead.
func (r *AccountDevicePolicyFallbackDomainService) GlobalSet(ctx context.Context, accountID string, body AccountDevicePolicyFallbackDomainGlobalSetParams, opts ...option.RequestOption) (res *FallbackDomainResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/fallback_domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead. This will only apply to the
// specified device settings profile.
func (r *AccountDevicePolicyFallbackDomainService) Set(ctx context.Context, accountID string, policyID string, body AccountDevicePolicyFallbackDomainSetParams, opts ...option.RequestOption) (res *FallbackDomainResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/fallback_domains", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type FallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []string           `json:"dns_server"`
	JSON      fallbackDomainJSON `json:"-"`
}

// fallbackDomainJSON contains the JSON metadata for the struct [FallbackDomain]
type fallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackDomainJSON) RawJSON() string {
	return r.raw
}

type FallbackDomainParam struct {
	// The domain suffix to match when resolving locally.
	Suffix param.Field[string] `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer param.Field[[]string] `json:"dns_server"`
}

func (r FallbackDomainParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FallbackDomainResponseCollection struct {
	Errors   []FallbackDomainResponseCollectionError   `json:"errors,required"`
	Messages []FallbackDomainResponseCollectionMessage `json:"messages,required"`
	Result   []FallbackDomain                          `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    FallbackDomainResponseCollectionSuccess    `json:"success,required"`
	ResultInfo FallbackDomainResponseCollectionResultInfo `json:"result_info"`
	JSON       fallbackDomainResponseCollectionJSON       `json:"-"`
}

// fallbackDomainResponseCollectionJSON contains the JSON metadata for the struct
// [FallbackDomainResponseCollection]
type fallbackDomainResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackDomainResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackDomainResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type FallbackDomainResponseCollectionError struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           FallbackDomainResponseCollectionErrorsSource `json:"source"`
	JSON             fallbackDomainResponseCollectionErrorJSON    `json:"-"`
}

// fallbackDomainResponseCollectionErrorJSON contains the JSON metadata for the
// struct [FallbackDomainResponseCollectionError]
type fallbackDomainResponseCollectionErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FallbackDomainResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackDomainResponseCollectionErrorJSON) RawJSON() string {
	return r.raw
}

type FallbackDomainResponseCollectionErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    fallbackDomainResponseCollectionErrorsSourceJSON `json:"-"`
}

// fallbackDomainResponseCollectionErrorsSourceJSON contains the JSON metadata for
// the struct [FallbackDomainResponseCollectionErrorsSource]
type fallbackDomainResponseCollectionErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackDomainResponseCollectionErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackDomainResponseCollectionErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FallbackDomainResponseCollectionMessage struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           FallbackDomainResponseCollectionMessagesSource `json:"source"`
	JSON             fallbackDomainResponseCollectionMessageJSON    `json:"-"`
}

// fallbackDomainResponseCollectionMessageJSON contains the JSON metadata for the
// struct [FallbackDomainResponseCollectionMessage]
type fallbackDomainResponseCollectionMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FallbackDomainResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackDomainResponseCollectionMessageJSON) RawJSON() string {
	return r.raw
}

type FallbackDomainResponseCollectionMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    fallbackDomainResponseCollectionMessagesSourceJSON `json:"-"`
}

// fallbackDomainResponseCollectionMessagesSourceJSON contains the JSON metadata
// for the struct [FallbackDomainResponseCollectionMessagesSource]
type fallbackDomainResponseCollectionMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackDomainResponseCollectionMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackDomainResponseCollectionMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type FallbackDomainResponseCollectionSuccess bool

const (
	FallbackDomainResponseCollectionSuccessTrue FallbackDomainResponseCollectionSuccess = true
)

func (r FallbackDomainResponseCollectionSuccess) IsKnown() bool {
	switch r {
	case FallbackDomainResponseCollectionSuccessTrue:
		return true
	}
	return false
}

type FallbackDomainResponseCollectionResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                        `json:"total_count"`
	JSON       fallbackDomainResponseCollectionResultInfoJSON `json:"-"`
}

// fallbackDomainResponseCollectionResultInfoJSON contains the JSON metadata for
// the struct [FallbackDomainResponseCollectionResultInfo]
type fallbackDomainResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackDomainResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackDomainResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountDevicePolicyFallbackDomainGlobalSetParams struct {
	Body []FallbackDomainParam `json:"body,required"`
}

func (r AccountDevicePolicyFallbackDomainGlobalSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDevicePolicyFallbackDomainSetParams struct {
	Body []FallbackDomainParam `json:"body,required"`
}

func (r AccountDevicePolicyFallbackDomainSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
