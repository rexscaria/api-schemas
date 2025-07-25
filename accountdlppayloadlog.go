// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountDlpPayloadLogService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDlpPayloadLogService] method instead.
type AccountDlpPayloadLogService struct {
	Options []option.RequestOption
}

// NewAccountDlpPayloadLogService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDlpPayloadLogService(opts ...option.RequestOption) (r *AccountDlpPayloadLogService) {
	r = &AccountDlpPayloadLogService{}
	r.Options = opts
	return
}

// Get payload log settings
func (r *AccountDlpPayloadLogService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountDlpPayloadLogGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set payload log settings
func (r *AccountDlpPayloadLogService) Update(ctx context.Context, accountID string, body AccountDlpPayloadLogUpdateParams, opts ...option.RequestOption) (res *AccountDlpPayloadLogUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type Setting struct {
	UpdatedAt time.Time   `json:"updated_at,required" format:"date-time"`
	PublicKey string      `json:"public_key,nullable"`
	JSON      settingJSON `json:"-"`
}

// settingJSON contains the JSON metadata for the struct [Setting]
type settingJSON struct {
	UpdatedAt   apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Setting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingJSON) RawJSON() string {
	return r.raw
}

type AccountDlpPayloadLogGetResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpPayloadLogGetResponseSuccess `json:"success,required"`
	Result  Setting                                `json:"result"`
	JSON    accountDlpPayloadLogGetResponseJSON    `json:"-"`
}

// accountDlpPayloadLogGetResponseJSON contains the JSON metadata for the struct
// [AccountDlpPayloadLogGetResponse]
type accountDlpPayloadLogGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPayloadLogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpPayloadLogGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDlpPayloadLogGetResponseSuccess bool

const (
	AccountDlpPayloadLogGetResponseSuccessTrue AccountDlpPayloadLogGetResponseSuccess = true
)

func (r AccountDlpPayloadLogGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpPayloadLogGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpPayloadLogUpdateResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpPayloadLogUpdateResponseSuccess `json:"success,required"`
	Result  Setting                                   `json:"result"`
	JSON    accountDlpPayloadLogUpdateResponseJSON    `json:"-"`
}

// accountDlpPayloadLogUpdateResponseJSON contains the JSON metadata for the struct
// [AccountDlpPayloadLogUpdateResponse]
type accountDlpPayloadLogUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPayloadLogUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpPayloadLogUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDlpPayloadLogUpdateResponseSuccess bool

const (
	AccountDlpPayloadLogUpdateResponseSuccessTrue AccountDlpPayloadLogUpdateResponseSuccess = true
)

func (r AccountDlpPayloadLogUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpPayloadLogUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpPayloadLogUpdateParams struct {
	PublicKey param.Field[string] `json:"public_key"`
}

func (r AccountDlpPayloadLogUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
