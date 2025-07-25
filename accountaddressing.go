// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAddressingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAddressingService] method instead.
type AccountAddressingService struct {
	Options           []option.RequestOption
	AddressMaps       *AccountAddressingAddressMapService
	LoaDocuments      *AccountAddressingLoaDocumentService
	Prefixes          *AccountAddressingPrefixService
	RegionalHostnames *AccountAddressingRegionalHostnameService
}

// NewAccountAddressingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAddressingService(opts ...option.RequestOption) (r *AccountAddressingService) {
	r = &AccountAddressingService{}
	r.Options = opts
	r.AddressMaps = NewAccountAddressingAddressMapService(opts...)
	r.LoaDocuments = NewAccountAddressingLoaDocumentService(opts...)
	r.Prefixes = NewAccountAddressingPrefixService(opts...)
	r.RegionalHostnames = NewAccountAddressingRegionalHostnameService(opts...)
	return
}

// Bring-Your-Own IP (BYOIP) prefixes onboarded to Cloudflare must be bound to a
// service running on the Cloudflare network to enable a Cloudflare product on the
// IP addresses. This endpoint can be used as a reference of available services on
// the Cloudflare network, and their service IDs.
func (r *AccountAddressingService) ListServices(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAddressingListServicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/services", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAddressingListServicesResponse struct {
	Errors   []AddressingMessages `json:"errors,required"`
	Messages []AddressingMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountAddressingListServicesResponseSuccess  `json:"success,required"`
	Result  []AccountAddressingListServicesResponseResult `json:"result"`
	JSON    accountAddressingListServicesResponseJSON     `json:"-"`
}

// accountAddressingListServicesResponseJSON contains the JSON metadata for the
// struct [AccountAddressingListServicesResponse]
type accountAddressingListServicesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingListServicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingListServicesResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountAddressingListServicesResponseSuccess bool

const (
	AccountAddressingListServicesResponseSuccessTrue AccountAddressingListServicesResponseSuccess = true
)

func (r AccountAddressingListServicesResponseSuccess) IsKnown() bool {
	switch r {
	case AccountAddressingListServicesResponseSuccessTrue:
		return true
	}
	return false
}

type AccountAddressingListServicesResponseResult struct {
	// Identifier of a Service on the Cloudflare network. Available services and their
	// IDs may be found in the **List Services** endpoint.
	ID string `json:"id"`
	// Name of a service running on the Cloudflare network
	Name string                                          `json:"name"`
	JSON accountAddressingListServicesResponseResultJSON `json:"-"`
}

// accountAddressingListServicesResponseResultJSON contains the JSON metadata for
// the struct [AccountAddressingListServicesResponseResult]
type accountAddressingListServicesResponseResultJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingListServicesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingListServicesResponseResultJSON) RawJSON() string {
	return r.raw
}
