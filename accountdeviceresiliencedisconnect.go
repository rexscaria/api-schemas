// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
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

// Fetch the Global WARP override state
func (r *AccountDeviceResilienceDisconnectService) Get(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *GlobalWarpOverrideResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/resilience/disconnect", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sets the Global WARP override state.
func (r *AccountDeviceResilienceDisconnectService) Set(ctx context.Context, accountID interface{}, body AccountDeviceResilienceDisconnectSetParams, opts ...option.RequestOption) (res *GlobalWarpOverrideResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/resilience/disconnect", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GlobalWarpOverrideResponse struct {
	Result GlobalWarpOverrideResponseResult `json:"result"`
	JSON   globalWarpOverrideResponseJSON   `json:"-"`
	APIResponseSingleTeamsDevices
}

// globalWarpOverrideResponseJSON contains the JSON metadata for the struct
// [GlobalWarpOverrideResponse]
type globalWarpOverrideResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalWarpOverrideResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalWarpOverrideResponseJSON) RawJSON() string {
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
