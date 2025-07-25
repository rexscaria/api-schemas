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

// AccountLogControlCmbConfigService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLogControlCmbConfigService] method instead.
type AccountLogControlCmbConfigService struct {
	Options []option.RequestOption
}

// NewAccountLogControlCmbConfigService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLogControlCmbConfigService(opts ...option.RequestOption) (r *AccountLogControlCmbConfigService) {
	r = &AccountLogControlCmbConfigService{}
	r.Options = opts
	return
}

// Gets CMB config.
func (r *AccountLogControlCmbConfigService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *CmbConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates CMB config.
func (r *AccountLogControlCmbConfigService) Update(ctx context.Context, accountID string, body AccountLogControlCmbConfigUpdateParams, opts ...option.RequestOption) (res *CmbConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes CMB config.
func (r *AccountLogControlCmbConfigService) Delete(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountLogControlCmbConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CmbConfig struct {
	// Allow out of region access
	AllowOutOfRegionAccess bool `json:"allow_out_of_region_access"`
	// Name of the region.
	Regions string        `json:"regions"`
	JSON    cmbConfigJSON `json:"-"`
}

// cmbConfigJSON contains the JSON metadata for the struct [CmbConfig]
type cmbConfigJSON struct {
	AllowOutOfRegionAccess apijson.Field
	Regions                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CmbConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cmbConfigJSON) RawJSON() string {
	return r.raw
}

type CmbConfigParam struct {
	// Allow out of region access
	AllowOutOfRegionAccess param.Field[bool] `json:"allow_out_of_region_access"`
	// Name of the region.
	Regions param.Field[string] `json:"regions"`
}

func (r CmbConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CmbConfigSingleResponse struct {
	Errors   []MessagesLogcontrolItem `json:"errors,required"`
	Messages []MessagesLogcontrolItem `json:"messages,required"`
	// Whether the API call was successful.
	Success CmbConfigSingleResponseSuccess `json:"success,required"`
	Result  CmbConfig                      `json:"result,nullable"`
	JSON    cmbConfigSingleResponseJSON    `json:"-"`
}

// cmbConfigSingleResponseJSON contains the JSON metadata for the struct
// [CmbConfigSingleResponse]
type cmbConfigSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CmbConfigSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cmbConfigSingleResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CmbConfigSingleResponseSuccess bool

const (
	CmbConfigSingleResponseSuccessTrue CmbConfigSingleResponseSuccess = true
)

func (r CmbConfigSingleResponseSuccess) IsKnown() bool {
	switch r {
	case CmbConfigSingleResponseSuccessTrue:
		return true
	}
	return false
}

type MessagesLogcontrolItem struct {
	Code             int64                        `json:"code,required"`
	Message          string                       `json:"message,required"`
	DocumentationURL string                       `json:"documentation_url"`
	Source           MessagesLogcontrolItemSource `json:"source"`
	JSON             messagesLogcontrolItemJSON   `json:"-"`
}

// messagesLogcontrolItemJSON contains the JSON metadata for the struct
// [MessagesLogcontrolItem]
type messagesLogcontrolItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessagesLogcontrolItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesLogcontrolItemJSON) RawJSON() string {
	return r.raw
}

type MessagesLogcontrolItemSource struct {
	Pointer string                           `json:"pointer"`
	JSON    messagesLogcontrolItemSourceJSON `json:"-"`
}

// messagesLogcontrolItemSourceJSON contains the JSON metadata for the struct
// [MessagesLogcontrolItemSource]
type messagesLogcontrolItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesLogcontrolItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesLogcontrolItemSourceJSON) RawJSON() string {
	return r.raw
}

type AccountLogControlCmbConfigDeleteResponse struct {
	Errors   []MessagesLogcontrolItem `json:"errors,required"`
	Messages []MessagesLogcontrolItem `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountLogControlCmbConfigDeleteResponseSuccess `json:"success,required"`
	Result  interface{}                                     `json:"result,nullable"`
	JSON    accountLogControlCmbConfigDeleteResponseJSON    `json:"-"`
}

// accountLogControlCmbConfigDeleteResponseJSON contains the JSON metadata for the
// struct [AccountLogControlCmbConfigDeleteResponse]
type accountLogControlCmbConfigDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogControlCmbConfigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLogControlCmbConfigDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountLogControlCmbConfigDeleteResponseSuccess bool

const (
	AccountLogControlCmbConfigDeleteResponseSuccessTrue AccountLogControlCmbConfigDeleteResponseSuccess = true
)

func (r AccountLogControlCmbConfigDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountLogControlCmbConfigDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountLogControlCmbConfigUpdateParams struct {
	CmbConfig CmbConfigParam `json:"cmb_config,required"`
}

func (r AccountLogControlCmbConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CmbConfig)
}
