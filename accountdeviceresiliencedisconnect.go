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

// AccountDeviceResilienceDisconnectService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDeviceResilienceDisconnectService] method instead.
type AccountDeviceResilienceDisconnectService struct {
	Options []option.RequestOption
}

// NewAccountDeviceResilienceDisconnectService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDeviceResilienceDisconnectService(opts ...option.RequestOption) (r *AccountDeviceResilienceDisconnectService) {
	r = &AccountDeviceResilienceDisconnectService{}
	r.Options = opts
	return
}

// Fetch the Global WARP override state.
func (r *AccountDeviceResilienceDisconnectService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *GlobalWarpOverrideResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/resilience/disconnect", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sets the Global WARP override state.
func (r *AccountDeviceResilienceDisconnectService) Set(ctx context.Context, accountID string, body AccountDeviceResilienceDisconnectSetParams, opts ...option.RequestOption) (res *GlobalWarpOverrideResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/resilience/disconnect", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GlobalWarpOverrideResponse struct {
	Errors   []GlobalWarpOverrideResponseError   `json:"errors,required"`
	Messages []GlobalWarpOverrideResponseMessage `json:"messages,required"`
	Result   GlobalWarpOverrideResponseResult    `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success GlobalWarpOverrideResponseSuccess `json:"success,required"`
	JSON    globalWarpOverrideResponseJSON    `json:"-"`
}

// globalWarpOverrideResponseJSON contains the JSON metadata for the struct
// [GlobalWarpOverrideResponse]
type globalWarpOverrideResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalWarpOverrideResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalWarpOverrideResponseJSON) RawJSON() string {
	return r.raw
}

type GlobalWarpOverrideResponseError struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           GlobalWarpOverrideResponseErrorsSource `json:"source"`
	JSON             globalWarpOverrideResponseErrorJSON    `json:"-"`
}

// globalWarpOverrideResponseErrorJSON contains the JSON metadata for the struct
// [GlobalWarpOverrideResponseError]
type globalWarpOverrideResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *GlobalWarpOverrideResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalWarpOverrideResponseErrorJSON) RawJSON() string {
	return r.raw
}

type GlobalWarpOverrideResponseErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    globalWarpOverrideResponseErrorsSourceJSON `json:"-"`
}

// globalWarpOverrideResponseErrorsSourceJSON contains the JSON metadata for the
// struct [GlobalWarpOverrideResponseErrorsSource]
type globalWarpOverrideResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalWarpOverrideResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalWarpOverrideResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type GlobalWarpOverrideResponseMessage struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           GlobalWarpOverrideResponseMessagesSource `json:"source"`
	JSON             globalWarpOverrideResponseMessageJSON    `json:"-"`
}

// globalWarpOverrideResponseMessageJSON contains the JSON metadata for the struct
// [GlobalWarpOverrideResponseMessage]
type globalWarpOverrideResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *GlobalWarpOverrideResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalWarpOverrideResponseMessageJSON) RawJSON() string {
	return r.raw
}

type GlobalWarpOverrideResponseMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    globalWarpOverrideResponseMessagesSourceJSON `json:"-"`
}

// globalWarpOverrideResponseMessagesSourceJSON contains the JSON metadata for the
// struct [GlobalWarpOverrideResponseMessagesSource]
type globalWarpOverrideResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalWarpOverrideResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalWarpOverrideResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type GlobalWarpOverrideResponseResult struct {
	// Disconnects all devices on the account using Global WARP override.
	Disconnect bool `json:"disconnect"`
	// When the Global WARP override state was updated.
	Timestamp time.Time                            `json:"timestamp" format:"date-time"`
	JSON      globalWarpOverrideResponseResultJSON `json:"-"`
}

// globalWarpOverrideResponseResultJSON contains the JSON metadata for the struct
// [GlobalWarpOverrideResponseResult]
type globalWarpOverrideResponseResultJSON struct {
	Disconnect  apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalWarpOverrideResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalWarpOverrideResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type GlobalWarpOverrideResponseSuccess bool

const (
	GlobalWarpOverrideResponseSuccessTrue GlobalWarpOverrideResponseSuccess = true
)

func (r GlobalWarpOverrideResponseSuccess) IsKnown() bool {
	switch r {
	case GlobalWarpOverrideResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDeviceResilienceDisconnectSetParams struct {
	// Disconnects all devices on the account using Global WARP override.
	Disconnect param.Field[bool] `json:"disconnect,required"`
	// Reasoning for setting the Global WARP override state. This will be surfaced in
	// the audit log.
	Justification param.Field[string] `json:"justification"`
}

func (r AccountDeviceResilienceDisconnectSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
