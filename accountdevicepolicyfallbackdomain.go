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
func (r *AccountDevicePolicyFallbackDomainService) List(ctx context.Context, accountID interface{}, policyID string, opts ...option.RequestOption) (res *FallbackDomainResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/fallback_domains", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead.
func (r *AccountDevicePolicyFallbackDomainService) GlobalList(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *FallbackDomainResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/fallback_domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead.
func (r *AccountDevicePolicyFallbackDomainService) GlobalSet(ctx context.Context, accountID interface{}, body AccountDevicePolicyFallbackDomainGlobalSetParams, opts ...option.RequestOption) (res *FallbackDomainResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/fallback_domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead. This will only apply to the
// specified device settings profile.
func (r *AccountDevicePolicyFallbackDomainService) Set(ctx context.Context, accountID interface{}, policyID string, body AccountDevicePolicyFallbackDomainSetParams, opts ...option.RequestOption) (res *FallbackDomainResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/fallback_domains", accountID, policyID)
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
	Result []FallbackDomain                     `json:"result"`
	JSON   fallbackDomainResponseCollectionJSON `json:"-"`
	APIResponseCollectionTeamsDevices
}

// fallbackDomainResponseCollectionJSON contains the JSON metadata for the struct
// [FallbackDomainResponseCollection]
type fallbackDomainResponseCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackDomainResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackDomainResponseCollectionJSON) RawJSON() string {
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
