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
func (r *AccountLogControlCmbConfigService) Delete(ctx context.Context, accountID string, body AccountLogControlCmbConfigDeleteParams, opts ...option.RequestOption) (res *AccountLogControlCmbConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type CmbConfig struct {
	// Comma-separated list of regions.
	Regions string        `json:"regions"`
	JSON    cmbConfigJSON `json:"-"`
}

// cmbConfigJSON contains the JSON metadata for the struct [CmbConfig]
type cmbConfigJSON struct {
	Regions     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CmbConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cmbConfigJSON) RawJSON() string {
	return r.raw
}

type CmbConfigParam struct {
	// Comma-separated list of regions.
	Regions param.Field[string] `json:"regions"`
}

func (r CmbConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CmbConfigSingleResponse struct {
	Result CmbConfig                   `json:"result,nullable"`
	JSON   cmbConfigSingleResponseJSON `json:"-"`
	SingleResponseLogControl
}

// cmbConfigSingleResponseJSON contains the JSON metadata for the struct
// [CmbConfigSingleResponse]
type cmbConfigSingleResponseJSON struct {
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

type CommonResponseLogControl struct {
	Errors   []MessagesLogcontrolItem `json:"errors,required"`
	Messages []MessagesLogcontrolItem `json:"messages,required"`
	// Whether the API call was successful
	Success CommonResponseLogControlSuccess `json:"success,required"`
	JSON    commonResponseLogControlJSON    `json:"-"`
}

// commonResponseLogControlJSON contains the JSON metadata for the struct
// [CommonResponseLogControl]
type commonResponseLogControlJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommonResponseLogControl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commonResponseLogControlJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CommonResponseLogControlSuccess bool

const (
	CommonResponseLogControlSuccessTrue CommonResponseLogControlSuccess = true
)

func (r CommonResponseLogControlSuccess) IsKnown() bool {
	switch r {
	case CommonResponseLogControlSuccessTrue:
		return true
	}
	return false
}

type MessagesLogcontrolItem struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    messagesLogcontrolItemJSON `json:"-"`
}

// messagesLogcontrolItemJSON contains the JSON metadata for the struct
// [MessagesLogcontrolItem]
type messagesLogcontrolItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesLogcontrolItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesLogcontrolItemJSON) RawJSON() string {
	return r.raw
}

type SingleResponseLogControl struct {
	Errors   []MessagesLogcontrolItem `json:"errors,required"`
	Messages []MessagesLogcontrolItem `json:"messages,required"`
	// Whether the API call was successful
	Success SingleResponseLogControlSuccess `json:"success,required"`
	JSON    singleResponseLogControlJSON    `json:"-"`
}

// singleResponseLogControlJSON contains the JSON metadata for the struct
// [SingleResponseLogControl]
type singleResponseLogControlJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLogControl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseLogControlJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SingleResponseLogControlSuccess bool

const (
	SingleResponseLogControlSuccessTrue SingleResponseLogControlSuccess = true
)

func (r SingleResponseLogControlSuccess) IsKnown() bool {
	switch r {
	case SingleResponseLogControlSuccessTrue:
		return true
	}
	return false
}

type AccountLogControlCmbConfigDeleteResponse struct {
	Result interface{}                                  `json:"result,nullable"`
	JSON   accountLogControlCmbConfigDeleteResponseJSON `json:"-"`
	CommonResponseLogControl
}

// accountLogControlCmbConfigDeleteResponseJSON contains the JSON metadata for the
// struct [AccountLogControlCmbConfigDeleteResponse]
type accountLogControlCmbConfigDeleteResponseJSON struct {
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

type AccountLogControlCmbConfigUpdateParams struct {
	CmbConfig CmbConfigParam `json:"cmb_config,required"`
}

func (r AccountLogControlCmbConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CmbConfig)
}

type AccountLogControlCmbConfigDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountLogControlCmbConfigDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
