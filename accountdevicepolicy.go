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

// AccountDevicePolicyService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDevicePolicyService] method instead.
type AccountDevicePolicyService struct {
	Options         []option.RequestOption
	Exclude         *AccountDevicePolicyExcludeService
	FallbackDomains *AccountDevicePolicyFallbackDomainService
	Include         *AccountDevicePolicyIncludeService
}

// NewAccountDevicePolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDevicePolicyService(opts ...option.RequestOption) (r *AccountDevicePolicyService) {
	r = &AccountDevicePolicyService{}
	r.Options = opts
	r.Exclude = NewAccountDevicePolicyExcludeService(opts...)
	r.FallbackDomains = NewAccountDevicePolicyFallbackDomainService(opts...)
	r.Include = NewAccountDevicePolicyIncludeService(opts...)
	return
}

// Creates a device settings profile to be applied to certain devices matching the
// criteria.
func (r *AccountDevicePolicyService) New(ctx context.Context, accountID string, body AccountDevicePolicyNewParams, opts ...option.RequestOption) (res *DeviceSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the default device settings profile for an account.
func (r *AccountDevicePolicyService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *DefaultDeviceSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the default device settings profile for an account.
func (r *AccountDevicePolicyService) Update(ctx context.Context, accountID string, body AccountDevicePolicyUpdateParams, opts ...option.RequestOption) (res *DefaultDeviceSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes a device settings profile and fetches a list of the remaining profiles
// for an account.
func (r *AccountDevicePolicyService) Delete(ctx context.Context, accountID string, policyID string, opts ...option.RequestOption) (res *DeviceSettingsResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Fetches a device settings profile by ID.
func (r *AccountDevicePolicyService) GetByID(ctx context.Context, accountID string, policyID string, opts ...option.RequestOption) (res *DeviceSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured device settings profile.
func (r *AccountDevicePolicyService) UpdateByID(ctx context.Context, accountID string, policyID string, body AccountDevicePolicyUpdateByIDParams, opts ...option.RequestOption) (res *DeviceSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type DefaultDeviceSettingsResponse struct {
	Errors   []DefaultDeviceSettingsResponseError   `json:"errors,required"`
	Messages []DefaultDeviceSettingsResponseMessage `json:"messages,required"`
	Result   DefaultDeviceSettingsResponseResult    `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DefaultDeviceSettingsResponseSuccess `json:"success,required"`
	JSON    defaultDeviceSettingsResponseJSON    `json:"-"`
}

// defaultDeviceSettingsResponseJSON contains the JSON metadata for the struct
// [DefaultDeviceSettingsResponse]
type defaultDeviceSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefaultDeviceSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defaultDeviceSettingsResponseJSON) RawJSON() string {
	return r.raw
}

type DefaultDeviceSettingsResponseError struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DefaultDeviceSettingsResponseErrorsSource `json:"source"`
	JSON             defaultDeviceSettingsResponseErrorJSON    `json:"-"`
}

// defaultDeviceSettingsResponseErrorJSON contains the JSON metadata for the struct
// [DefaultDeviceSettingsResponseError]
type defaultDeviceSettingsResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DefaultDeviceSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defaultDeviceSettingsResponseErrorJSON) RawJSON() string {
	return r.raw
}

type DefaultDeviceSettingsResponseErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    defaultDeviceSettingsResponseErrorsSourceJSON `json:"-"`
}

// defaultDeviceSettingsResponseErrorsSourceJSON contains the JSON metadata for the
// struct [DefaultDeviceSettingsResponseErrorsSource]
type defaultDeviceSettingsResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefaultDeviceSettingsResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defaultDeviceSettingsResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DefaultDeviceSettingsResponseMessage struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DefaultDeviceSettingsResponseMessagesSource `json:"source"`
	JSON             defaultDeviceSettingsResponseMessageJSON    `json:"-"`
}

// defaultDeviceSettingsResponseMessageJSON contains the JSON metadata for the
// struct [DefaultDeviceSettingsResponseMessage]
type defaultDeviceSettingsResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DefaultDeviceSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defaultDeviceSettingsResponseMessageJSON) RawJSON() string {
	return r.raw
}

type DefaultDeviceSettingsResponseMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    defaultDeviceSettingsResponseMessagesSourceJSON `json:"-"`
}

// defaultDeviceSettingsResponseMessagesSourceJSON contains the JSON metadata for
// the struct [DefaultDeviceSettingsResponseMessagesSource]
type defaultDeviceSettingsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefaultDeviceSettingsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defaultDeviceSettingsResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type DefaultDeviceSettingsResponseResult struct {
	// Whether to allow the user to switch WARP between modes.
	AllowModeSwitch bool `json:"allow_mode_switch"`
	// Whether to receive update notifications when a new version of the client is
	// available.
	AllowUpdates bool `json:"allow_updates"`
	// Whether to allow devices to leave the organization.
	AllowedToLeave bool `json:"allowed_to_leave"`
	// The amount of time in seconds to reconnect after having been disabled.
	AutoConnect float64 `json:"auto_connect"`
	// Turn on the captive portal after the specified amount of time.
	CaptivePortal float64 `json:"captive_portal"`
	// Whether the policy will be applied to matching devices.
	Default bool `json:"default"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback bool `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled bool `json:"enabled"`
	// List of routes excluded in the WARP client's tunnel.
	Exclude []SplitTunnel `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool             `json:"exclude_office_ips"`
	FallbackDomains  []FallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string           `json:"gateway_unique_id"`
	// List of routes included in the WARP client's tunnel.
	Include []SplitTunnelInclude `json:"include"`
	// Determines if the operating system will register WARP's local interface IP with
	// your on-premises DNS server.
	RegisterInterfaceIPWithDNS bool `json:"register_interface_ip_with_dns"`
	// Determines whether the WARP client indicates to SCCM that it is inside a VPN
	// boundary. (Windows only).
	SccmVpnBoundarySupport bool          `json:"sccm_vpn_boundary_support"`
	ServiceModeV2          ServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool `json:"switch_locked"`
	// Determines which tunnel protocol to use.
	TunnelProtocol string                                  `json:"tunnel_protocol"`
	JSON           defaultDeviceSettingsResponseResultJSON `json:"-"`
}

// defaultDeviceSettingsResponseResultJSON contains the JSON metadata for the
// struct [DefaultDeviceSettingsResponseResult]
type defaultDeviceSettingsResponseResultJSON struct {
	AllowModeSwitch            apijson.Field
	AllowUpdates               apijson.Field
	AllowedToLeave             apijson.Field
	AutoConnect                apijson.Field
	CaptivePortal              apijson.Field
	Default                    apijson.Field
	DisableAutoFallback        apijson.Field
	Enabled                    apijson.Field
	Exclude                    apijson.Field
	ExcludeOfficeIPs           apijson.Field
	FallbackDomains            apijson.Field
	GatewayUniqueID            apijson.Field
	Include                    apijson.Field
	RegisterInterfaceIPWithDNS apijson.Field
	SccmVpnBoundarySupport     apijson.Field
	ServiceModeV2              apijson.Field
	SupportURL                 apijson.Field
	SwitchLocked               apijson.Field
	TunnelProtocol             apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *DefaultDeviceSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defaultDeviceSettingsResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DefaultDeviceSettingsResponseSuccess bool

const (
	DefaultDeviceSettingsResponseSuccessTrue DefaultDeviceSettingsResponseSuccess = true
)

func (r DefaultDeviceSettingsResponseSuccess) IsKnown() bool {
	switch r {
	case DefaultDeviceSettingsResponseSuccessTrue:
		return true
	}
	return false
}

type DeviceSettingsPolicy struct {
	// Whether to allow the user to switch WARP between modes.
	AllowModeSwitch bool `json:"allow_mode_switch"`
	// Whether to receive update notifications when a new version of the client is
	// available.
	AllowUpdates bool `json:"allow_updates"`
	// Whether to allow devices to leave the organization.
	AllowedToLeave bool `json:"allowed_to_leave"`
	// The amount of time in seconds to reconnect after having been disabled.
	AutoConnect float64 `json:"auto_connect"`
	// Turn on the captive portal after the specified amount of time.
	CaptivePortal float64 `json:"captive_portal"`
	// Whether the policy is the default policy for an account.
	Default bool `json:"default"`
	// A description of the policy.
	Description string `json:"description"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback bool `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled bool `json:"enabled"`
	// List of routes excluded in the WARP client's tunnel.
	Exclude []SplitTunnel `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool             `json:"exclude_office_ips"`
	FallbackDomains  []FallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string           `json:"gateway_unique_id"`
	// List of routes included in the WARP client's tunnel.
	Include []SplitTunnelInclude `json:"include"`
	// The amount of time in minutes a user is allowed access to their LAN. A value of
	// 0 will allow LAN access until the next WARP reconnection, such as a reboot or a
	// laptop waking from sleep. Note that this field is omitted from the response if
	// null or unset.
	LanAllowMinutes float64 `json:"lan_allow_minutes"`
	// The size of the subnet for the local access network. Note that this field is
	// omitted from the response if null or unset.
	LanAllowSubnetSize float64 `json:"lan_allow_subnet_size"`
	// The wirefilter expression to match devices. Available values: "identity.email",
	// "identity.groups.id", "identity.groups.name", "identity.groups.email",
	// "identity.service_token_uuid", "identity.saml_attributes", "network", "os.name",
	// "os.version".
	Match string `json:"match"`
	// The name of the device settings profile.
	Name     string `json:"name"`
	PolicyID string `json:"policy_id"`
	// The precedence of the policy. Lower values indicate higher precedence. Policies
	// will be evaluated in ascending order of this field.
	Precedence float64 `json:"precedence"`
	// Determines if the operating system will register WARP's local interface IP with
	// your on-premises DNS server.
	RegisterInterfaceIPWithDNS bool `json:"register_interface_ip_with_dns"`
	// Determines whether the WARP client indicates to SCCM that it is inside a VPN
	// boundary. (Windows only).
	SccmVpnBoundarySupport bool          `json:"sccm_vpn_boundary_support"`
	ServiceModeV2          ServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                             `json:"switch_locked"`
	TargetTests  []DeviceSettingsPolicyTargetTest `json:"target_tests"`
	// Determines which tunnel protocol to use.
	TunnelProtocol string                   `json:"tunnel_protocol"`
	JSON           deviceSettingsPolicyJSON `json:"-"`
}

// deviceSettingsPolicyJSON contains the JSON metadata for the struct
// [DeviceSettingsPolicy]
type deviceSettingsPolicyJSON struct {
	AllowModeSwitch            apijson.Field
	AllowUpdates               apijson.Field
	AllowedToLeave             apijson.Field
	AutoConnect                apijson.Field
	CaptivePortal              apijson.Field
	Default                    apijson.Field
	Description                apijson.Field
	DisableAutoFallback        apijson.Field
	Enabled                    apijson.Field
	Exclude                    apijson.Field
	ExcludeOfficeIPs           apijson.Field
	FallbackDomains            apijson.Field
	GatewayUniqueID            apijson.Field
	Include                    apijson.Field
	LanAllowMinutes            apijson.Field
	LanAllowSubnetSize         apijson.Field
	Match                      apijson.Field
	Name                       apijson.Field
	PolicyID                   apijson.Field
	Precedence                 apijson.Field
	RegisterInterfaceIPWithDNS apijson.Field
	SccmVpnBoundarySupport     apijson.Field
	ServiceModeV2              apijson.Field
	SupportURL                 apijson.Field
	SwitchLocked               apijson.Field
	TargetTests                apijson.Field
	TunnelProtocol             apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *DeviceSettingsPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingsPolicyJSON) RawJSON() string {
	return r.raw
}

type DeviceSettingsPolicyTargetTest struct {
	// The id of the DEX test targeting this policy.
	ID string `json:"id"`
	// The name of the DEX test targeting this policy.
	Name string                             `json:"name"`
	JSON deviceSettingsPolicyTargetTestJSON `json:"-"`
}

// deviceSettingsPolicyTargetTestJSON contains the JSON metadata for the struct
// [DeviceSettingsPolicyTargetTest]
type deviceSettingsPolicyTargetTestJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsPolicyTargetTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingsPolicyTargetTestJSON) RawJSON() string {
	return r.raw
}

type DeviceSettingsResponse struct {
	Errors   []DeviceSettingsResponseError   `json:"errors,required"`
	Messages []DeviceSettingsResponseMessage `json:"messages,required"`
	Result   DeviceSettingsPolicy            `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceSettingsResponseSuccess `json:"success,required"`
	JSON    deviceSettingsResponseJSON    `json:"-"`
}

// deviceSettingsResponseJSON contains the JSON metadata for the struct
// [DeviceSettingsResponse]
type deviceSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingsResponseJSON) RawJSON() string {
	return r.raw
}

type DeviceSettingsResponseError struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           DeviceSettingsResponseErrorsSource `json:"source"`
	JSON             deviceSettingsResponseErrorJSON    `json:"-"`
}

// deviceSettingsResponseErrorJSON contains the JSON metadata for the struct
// [DeviceSettingsResponseError]
type deviceSettingsResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingsResponseErrorJSON) RawJSON() string {
	return r.raw
}

type DeviceSettingsResponseErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    deviceSettingsResponseErrorsSourceJSON `json:"-"`
}

// deviceSettingsResponseErrorsSourceJSON contains the JSON metadata for the struct
// [DeviceSettingsResponseErrorsSource]
type deviceSettingsResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingsResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DeviceSettingsResponseMessage struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           DeviceSettingsResponseMessagesSource `json:"source"`
	JSON             deviceSettingsResponseMessageJSON    `json:"-"`
}

// deviceSettingsResponseMessageJSON contains the JSON metadata for the struct
// [DeviceSettingsResponseMessage]
type deviceSettingsResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingsResponseMessageJSON) RawJSON() string {
	return r.raw
}

type DeviceSettingsResponseMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    deviceSettingsResponseMessagesSourceJSON `json:"-"`
}

// deviceSettingsResponseMessagesSourceJSON contains the JSON metadata for the
// struct [DeviceSettingsResponseMessagesSource]
type deviceSettingsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingsResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceSettingsResponseSuccess bool

const (
	DeviceSettingsResponseSuccessTrue DeviceSettingsResponseSuccess = true
)

func (r DeviceSettingsResponseSuccess) IsKnown() bool {
	switch r {
	case DeviceSettingsResponseSuccessTrue:
		return true
	}
	return false
}

type ServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64           `json:"port"`
	JSON serviceModeV2JSON `json:"-"`
}

// serviceModeV2JSON contains the JSON metadata for the struct [ServiceModeV2]
type serviceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceModeV2JSON) RawJSON() string {
	return r.raw
}

type ServiceModeV2Param struct {
	// The mode to run the WARP client under.
	Mode param.Field[string] `json:"mode"`
	// The port number when used with proxy mode.
	Port param.Field[float64] `json:"port"`
}

func (r ServiceModeV2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePolicyNewParams struct {
	// The wirefilter expression to match devices. Available values: "identity.email",
	// "identity.groups.id", "identity.groups.name", "identity.groups.email",
	// "identity.service_token_uuid", "identity.saml_attributes", "network", "os.name",
	// "os.version".
	Match param.Field[string] `json:"match,required"`
	// The name of the device settings profile.
	Name param.Field[string] `json:"name,required"`
	// The precedence of the policy. Lower values indicate higher precedence. Policies
	// will be evaluated in ascending order of this field.
	Precedence param.Field[float64] `json:"precedence,required"`
	// Whether to allow the user to switch WARP between modes.
	AllowModeSwitch param.Field[bool] `json:"allow_mode_switch"`
	// Whether to receive update notifications when a new version of the client is
	// available.
	AllowUpdates param.Field[bool] `json:"allow_updates"`
	// Whether to allow devices to leave the organization.
	AllowedToLeave param.Field[bool] `json:"allowed_to_leave"`
	// The amount of time in seconds to reconnect after having been disabled.
	AutoConnect param.Field[float64] `json:"auto_connect"`
	// Turn on the captive portal after the specified amount of time.
	CaptivePortal param.Field[float64] `json:"captive_portal"`
	// A description of the policy.
	Description param.Field[string] `json:"description"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback param.Field[bool] `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled param.Field[bool] `json:"enabled"`
	// List of routes excluded in the WARP client's tunnel. Both 'exclude' and
	// 'include' cannot be set in the same request.
	Exclude param.Field[[]SplitTunnelUnionParam] `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs param.Field[bool] `json:"exclude_office_ips"`
	// List of routes included in the WARP client's tunnel. Both 'exclude' and
	// 'include' cannot be set in the same request.
	Include param.Field[[]SplitTunnelIncludeUnionParam] `json:"include"`
	// The amount of time in minutes a user is allowed access to their LAN. A value of
	// 0 will allow LAN access until the next WARP reconnection, such as a reboot or a
	// laptop waking from sleep. Note that this field is omitted from the response if
	// null or unset.
	LanAllowMinutes param.Field[float64] `json:"lan_allow_minutes"`
	// The size of the subnet for the local access network. Note that this field is
	// omitted from the response if null or unset.
	LanAllowSubnetSize param.Field[float64] `json:"lan_allow_subnet_size"`
	// Determines if the operating system will register WARP's local interface IP with
	// your on-premises DNS server.
	RegisterInterfaceIPWithDNS param.Field[bool] `json:"register_interface_ip_with_dns"`
	// Determines whether the WARP client indicates to SCCM that it is inside a VPN
	// boundary. (Windows only).
	SccmVpnBoundarySupport param.Field[bool]               `json:"sccm_vpn_boundary_support"`
	ServiceModeV2          param.Field[ServiceModeV2Param] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
	// Determines which tunnel protocol to use.
	TunnelProtocol param.Field[string] `json:"tunnel_protocol"`
}

func (r AccountDevicePolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePolicyUpdateParams struct {
	// Whether to allow the user to switch WARP between modes.
	AllowModeSwitch param.Field[bool] `json:"allow_mode_switch"`
	// Whether to receive update notifications when a new version of the client is
	// available.
	AllowUpdates param.Field[bool] `json:"allow_updates"`
	// Whether to allow devices to leave the organization.
	AllowedToLeave param.Field[bool] `json:"allowed_to_leave"`
	// The amount of time in seconds to reconnect after having been disabled.
	AutoConnect param.Field[float64] `json:"auto_connect"`
	// Turn on the captive portal after the specified amount of time.
	CaptivePortal param.Field[float64] `json:"captive_portal"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback param.Field[bool] `json:"disable_auto_fallback"`
	// List of routes excluded in the WARP client's tunnel. Both 'exclude' and
	// 'include' cannot be set in the same request.
	Exclude param.Field[[]SplitTunnelUnionParam] `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs param.Field[bool] `json:"exclude_office_ips"`
	// List of routes included in the WARP client's tunnel. Both 'exclude' and
	// 'include' cannot be set in the same request.
	Include param.Field[[]SplitTunnelIncludeUnionParam] `json:"include"`
	// The amount of time in minutes a user is allowed access to their LAN. A value of
	// 0 will allow LAN access until the next WARP reconnection, such as a reboot or a
	// laptop waking from sleep. Note that this field is omitted from the response if
	// null or unset.
	LanAllowMinutes param.Field[float64] `json:"lan_allow_minutes"`
	// The size of the subnet for the local access network. Note that this field is
	// omitted from the response if null or unset.
	LanAllowSubnetSize param.Field[float64] `json:"lan_allow_subnet_size"`
	// Determines if the operating system will register WARP's local interface IP with
	// your on-premises DNS server.
	RegisterInterfaceIPWithDNS param.Field[bool] `json:"register_interface_ip_with_dns"`
	// Determines whether the WARP client indicates to SCCM that it is inside a VPN
	// boundary. (Windows only).
	SccmVpnBoundarySupport param.Field[bool]               `json:"sccm_vpn_boundary_support"`
	ServiceModeV2          param.Field[ServiceModeV2Param] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
	// Determines which tunnel protocol to use.
	TunnelProtocol param.Field[string] `json:"tunnel_protocol"`
}

func (r AccountDevicePolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePolicyUpdateByIDParams struct {
	// Whether to allow the user to switch WARP between modes.
	AllowModeSwitch param.Field[bool] `json:"allow_mode_switch"`
	// Whether to receive update notifications when a new version of the client is
	// available.
	AllowUpdates param.Field[bool] `json:"allow_updates"`
	// Whether to allow devices to leave the organization.
	AllowedToLeave param.Field[bool] `json:"allowed_to_leave"`
	// The amount of time in seconds to reconnect after having been disabled.
	AutoConnect param.Field[float64] `json:"auto_connect"`
	// Turn on the captive portal after the specified amount of time.
	CaptivePortal param.Field[float64] `json:"captive_portal"`
	// A description of the policy.
	Description param.Field[string] `json:"description"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback param.Field[bool] `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled param.Field[bool] `json:"enabled"`
	// List of routes excluded in the WARP client's tunnel. Both 'exclude' and
	// 'include' cannot be set in the same request.
	Exclude param.Field[[]SplitTunnelUnionParam] `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs param.Field[bool] `json:"exclude_office_ips"`
	// List of routes included in the WARP client's tunnel. Both 'exclude' and
	// 'include' cannot be set in the same request.
	Include param.Field[[]SplitTunnelIncludeUnionParam] `json:"include"`
	// The amount of time in minutes a user is allowed access to their LAN. A value of
	// 0 will allow LAN access until the next WARP reconnection, such as a reboot or a
	// laptop waking from sleep. Note that this field is omitted from the response if
	// null or unset.
	LanAllowMinutes param.Field[float64] `json:"lan_allow_minutes"`
	// The size of the subnet for the local access network. Note that this field is
	// omitted from the response if null or unset.
	LanAllowSubnetSize param.Field[float64] `json:"lan_allow_subnet_size"`
	// The wirefilter expression to match devices. Available values: "identity.email",
	// "identity.groups.id", "identity.groups.name", "identity.groups.email",
	// "identity.service_token_uuid", "identity.saml_attributes", "network", "os.name",
	// "os.version".
	Match param.Field[string] `json:"match"`
	// The name of the device settings profile.
	Name param.Field[string] `json:"name"`
	// The precedence of the policy. Lower values indicate higher precedence. Policies
	// will be evaluated in ascending order of this field.
	Precedence param.Field[float64] `json:"precedence"`
	// Determines if the operating system will register WARP's local interface IP with
	// your on-premises DNS server.
	RegisterInterfaceIPWithDNS param.Field[bool] `json:"register_interface_ip_with_dns"`
	// Determines whether the WARP client indicates to SCCM that it is inside a VPN
	// boundary. (Windows only).
	SccmVpnBoundarySupport param.Field[bool]               `json:"sccm_vpn_boundary_support"`
	ServiceModeV2          param.Field[ServiceModeV2Param] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
	// Determines which tunnel protocol to use.
	TunnelProtocol param.Field[string] `json:"tunnel_protocol"`
}

func (r AccountDevicePolicyUpdateByIDParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
