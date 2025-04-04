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

// AccountAddressingPrefixBindingService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAddressingPrefixBindingService] method instead.
type AccountAddressingPrefixBindingService struct {
	Options []option.RequestOption
}

// NewAccountAddressingPrefixBindingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressingPrefixBindingService(opts ...option.RequestOption) (r *AccountAddressingPrefixBindingService) {
	r = &AccountAddressingPrefixBindingService{}
	r.Options = opts
	return
}

// Creates a new Service Binding, routing traffic to IPs within the given CIDR to a
// service running on Cloudflare's network. **Note:** This API may only be used on
// prefixes currently configured with a Magic Transit/Cloudflare CDN/Cloudflare
// Spectrum service binding, and only allows creating upgrade service bindings for
// the Cloudflare CDN or Cloudflare Spectrum.
func (r *AccountAddressingPrefixBindingService) New(ctx context.Context, accountID string, prefixID string, body AccountAddressingPrefixBindingNewParams, opts ...option.RequestOption) (res *AccountAddressingPrefixBindingNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a single Service Binding
func (r *AccountAddressingPrefixBindingService) Get(ctx context.Context, accountID string, prefixID string, bindingID string, opts ...option.RequestOption) (res *AccountAddressingPrefixBindingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	if bindingID == "" {
		err = errors.New("missing required binding_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", accountID, prefixID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the Cloudflare services this prefix is currently bound to. Traffic sent to
// an address within an IP prefix will be routed to the Cloudflare service of the
// most-specific Service Binding matching the address. **Example:** binding
// `192.0.2.0/24` to Cloudflare Magic Transit and `192.0.2.1/32` to the Cloudflare
// CDN would route traffic for `192.0.2.1` to the CDN, and traffic for all other
// IPs in the prefix to Cloudflare Magic Transit.
func (r *AccountAddressingPrefixBindingService) List(ctx context.Context, accountID string, prefixID string, opts ...option.RequestOption) (res *AccountAddressingPrefixBindingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Service Binding
func (r *AccountAddressingPrefixBindingService) Delete(ctx context.Context, accountID string, prefixID string, bindingID string, opts ...option.RequestOption) (res *APIResponseAddressing, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	if bindingID == "" {
		err = errors.New("missing required binding_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", accountID, prefixID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AddressingMessages struct {
	Code    int64                  `json:"code,required"`
	Message string                 `json:"message,required"`
	JSON    addressingMessagesJSON `json:"-"`
}

// addressingMessagesJSON contains the JSON metadata for the struct
// [AddressingMessages]
type addressingMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingMessagesJSON) RawJSON() string {
	return r.raw
}

type APIResponseAddressing struct {
	Errors   []AddressingMessages `json:"errors,required"`
	Messages []AddressingMessages `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseAddressingSuccess `json:"success,required"`
	JSON    apiResponseAddressingJSON    `json:"-"`
}

// apiResponseAddressingJSON contains the JSON metadata for the struct
// [APIResponseAddressing]
type apiResponseAddressingJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseAddressingJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseAddressingSuccess bool

const (
	APIResponseAddressingSuccessTrue APIResponseAddressingSuccess = true
)

func (r APIResponseAddressingSuccess) IsKnown() bool {
	switch r {
	case APIResponseAddressingSuccessTrue:
		return true
	}
	return false
}

type ServiceBinding struct {
	// Identifier of a Service Binding.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning ServiceBindingProvisioning `json:"provisioning"`
	// Identifier of a Service on the Cloudflare network. Available services and their
	// IDs may be found in the **List Services** endpoint.
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string             `json:"service_name"`
	JSON        serviceBindingJSON `json:"-"`
}

// serviceBindingJSON contains the JSON metadata for the struct [ServiceBinding]
type serviceBindingJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceBindingJSON) RawJSON() string {
	return r.raw
}

// Status of a Service Binding's deployment to the Cloudflare network
type ServiceBindingProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State ServiceBindingProvisioningState `json:"state"`
	JSON  serviceBindingProvisioningJSON  `json:"-"`
}

// serviceBindingProvisioningJSON contains the JSON metadata for the struct
// [ServiceBindingProvisioning]
type serviceBindingProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceBindingProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceBindingProvisioningJSON) RawJSON() string {
	return r.raw
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type ServiceBindingProvisioningState string

const (
	ServiceBindingProvisioningStateProvisioning ServiceBindingProvisioningState = "provisioning"
	ServiceBindingProvisioningStateActive       ServiceBindingProvisioningState = "active"
)

func (r ServiceBindingProvisioningState) IsKnown() bool {
	switch r {
	case ServiceBindingProvisioningStateProvisioning, ServiceBindingProvisioningStateActive:
		return true
	}
	return false
}

type AccountAddressingPrefixBindingNewResponse struct {
	Result ServiceBinding                                `json:"result"`
	JSON   accountAddressingPrefixBindingNewResponseJSON `json:"-"`
	APIResponseAddressing
}

// accountAddressingPrefixBindingNewResponseJSON contains the JSON metadata for the
// struct [AccountAddressingPrefixBindingNewResponse]
type accountAddressingPrefixBindingNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingPrefixBindingNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAddressingPrefixBindingGetResponse struct {
	Result ServiceBinding                                `json:"result"`
	JSON   accountAddressingPrefixBindingGetResponseJSON `json:"-"`
	APIResponseAddressing
}

// accountAddressingPrefixBindingGetResponseJSON contains the JSON metadata for the
// struct [AccountAddressingPrefixBindingGetResponse]
type accountAddressingPrefixBindingGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingPrefixBindingGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAddressingPrefixBindingListResponse struct {
	Result []ServiceBinding                               `json:"result"`
	JSON   accountAddressingPrefixBindingListResponseJSON `json:"-"`
	APIResponseAddressing
}

// accountAddressingPrefixBindingListResponseJSON contains the JSON metadata for
// the struct [AccountAddressingPrefixBindingListResponse]
type accountAddressingPrefixBindingListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingPrefixBindingListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAddressingPrefixBindingNewParams struct {
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr"`
	// Identifier of a Service on the Cloudflare network. Available services and their
	// IDs may be found in the **List Services** endpoint.
	ServiceID param.Field[string] `json:"service_id"`
}

func (r AccountAddressingPrefixBindingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
