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

// AccountDeviceSettingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDeviceSettingService] method instead.
type AccountDeviceSettingService struct {
	Options []option.RequestOption
}

// NewAccountDeviceSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDeviceSettingService(opts ...option.RequestOption) (r *AccountDeviceSettingService) {
	r = &AccountDeviceSettingService{}
	r.Options = opts
	return
}

// Describes the current device settings for a Zero Trust account.
func (r *AccountDeviceSettingService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *ZeroTrustAccountDeviceSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the current device settings for a Zero Trust account.
func (r *AccountDeviceSettingService) Update(ctx context.Context, accountID string, body AccountDeviceSettingUpdateParams, opts ...option.RequestOption) (res *ZeroTrustAccountDeviceSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Patches the current device settings for a Zero Trust account.
func (r *AccountDeviceSettingService) Patch(ctx context.Context, accountID string, body AccountDeviceSettingPatchParams, opts ...option.RequestOption) (res *ZeroTrustAccountDeviceSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZeroTrustAccountDeviceSettings struct {
	// Sets the time limit, in seconds, that a user can use an override code to bypass
	// WARP.
	DisableForTime float64 `json:"disable_for_time"`
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled bool `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled bool `json:"gateway_udp_proxy_enabled"`
	// Enable installation of cloudflare managed root certificate.
	RootCertificateInstallationEnabled bool `json:"root_certificate_installation_enabled"`
	// Enable using CGNAT virtual IPv4.
	UseZtVirtualIP bool                               `json:"use_zt_virtual_ip"`
	JSON           zeroTrustAccountDeviceSettingsJSON `json:"-"`
}

// zeroTrustAccountDeviceSettingsJSON contains the JSON metadata for the struct
// [ZeroTrustAccountDeviceSettings]
type zeroTrustAccountDeviceSettingsJSON struct {
	DisableForTime                     apijson.Field
	GatewayProxyEnabled                apijson.Field
	GatewayUdpProxyEnabled             apijson.Field
	RootCertificateInstallationEnabled apijson.Field
	UseZtVirtualIP                     apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *ZeroTrustAccountDeviceSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccountDeviceSettingsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccountDeviceSettingsParam struct {
	// Sets the time limit, in seconds, that a user can use an override code to bypass
	// WARP.
	DisableForTime param.Field[float64] `json:"disable_for_time"`
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled param.Field[bool] `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled param.Field[bool] `json:"gateway_udp_proxy_enabled"`
	// Enable installation of cloudflare managed root certificate.
	RootCertificateInstallationEnabled param.Field[bool] `json:"root_certificate_installation_enabled"`
	// Enable using CGNAT virtual IPv4.
	UseZtVirtualIP param.Field[bool] `json:"use_zt_virtual_ip"`
}

func (r ZeroTrustAccountDeviceSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccountDeviceSettingsResponse struct {
	Errors   []MessagesDeviceTestsItems     `json:"errors,required"`
	Messages []MessagesDeviceTestsItems     `json:"messages,required"`
	Result   ZeroTrustAccountDeviceSettings `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustAccountDeviceSettingsResponseSuccess `json:"success,required"`
	JSON    zeroTrustAccountDeviceSettingsResponseJSON    `json:"-"`
}

// zeroTrustAccountDeviceSettingsResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccountDeviceSettingsResponse]
type zeroTrustAccountDeviceSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccountDeviceSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccountDeviceSettingsResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustAccountDeviceSettingsResponseSuccess bool

const (
	ZeroTrustAccountDeviceSettingsResponseSuccessTrue ZeroTrustAccountDeviceSettingsResponseSuccess = true
)

func (r ZeroTrustAccountDeviceSettingsResponseSuccess) IsKnown() bool {
	switch r {
	case ZeroTrustAccountDeviceSettingsResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDeviceSettingUpdateParams struct {
	ZeroTrustAccountDeviceSettings ZeroTrustAccountDeviceSettingsParam `json:"zero_trust_account_device_settings,required"`
}

func (r AccountDeviceSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ZeroTrustAccountDeviceSettings)
}

type AccountDeviceSettingPatchParams struct {
	ZeroTrustAccountDeviceSettings ZeroTrustAccountDeviceSettingsParam `json:"zero_trust_account_device_settings,required"`
}

func (r AccountDeviceSettingPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ZeroTrustAccountDeviceSettings)
}
