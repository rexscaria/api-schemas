// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAccessAppService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessAppService] method instead.
type AccountAccessAppService struct {
	Options  []option.RequestOption
	Ca       *AccountAccessAppCaService
	Policies *AccountAccessAppPolicyService
}

// NewAccountAccessAppService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessAppService(opts ...option.RequestOption) (r *AccountAccessAppService) {
	r = &AccountAccessAppService{}
	r.Options = opts
	r.Ca = NewAccountAccessAppCaService(opts...)
	r.Policies = NewAccountAccessAppPolicyService(opts...)
	return
}

type IDResponseApps struct {
	Errors   []MessagesAccessItem `json:"errors,required"`
	Messages []MessagesAccessItem `json:"messages,required"`
	// Whether the API call was successful.
	Success IDResponseAppsSuccess `json:"success,required"`
	Result  IDResponseAppsResult  `json:"result"`
	JSON    idResponseAppsJSON    `json:"-"`
}

// idResponseAppsJSON contains the JSON metadata for the struct [IDResponseApps]
type idResponseAppsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseAppsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IDResponseAppsSuccess bool

const (
	IDResponseAppsSuccessTrue IDResponseAppsSuccess = true
)

func (r IDResponseAppsSuccess) IsKnown() bool {
	switch r {
	case IDResponseAppsSuccessTrue:
		return true
	}
	return false
}

type IDResponseAppsResult struct {
	// UUID.
	ID   string                   `json:"id"`
	JSON idResponseAppsResultJSON `json:"-"`
}

// idResponseAppsResultJSON contains the JSON metadata for the struct
// [IDResponseAppsResult]
type idResponseAppsResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseAppsResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseAppsResultJSON) RawJSON() string {
	return r.raw
}
