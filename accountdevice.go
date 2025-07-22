// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountDeviceService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDeviceService] method instead.
type AccountDeviceService struct {
	Options    []option.RequestOption
	DexTests   *AccountDeviceDexTestService
	Networks   *AccountDeviceNetworkService
	Policy     *AccountDevicePolicyService
	Posture    *AccountDevicePostureService
	Resilience *AccountDeviceResilienceService
	Settings   *AccountDeviceSettingService
}

// NewAccountDeviceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountDeviceService(opts ...option.RequestOption) (r *AccountDeviceService) {
	r = &AccountDeviceService{}
	r.Options = opts
	r.DexTests = NewAccountDeviceDexTestService(opts...)
	r.Networks = NewAccountDeviceNetworkService(opts...)
	r.Policy = NewAccountDevicePolicyService(opts...)
	r.Posture = NewAccountDevicePostureService(opts...)
	r.Resilience = NewAccountDeviceResilienceService(opts...)
	r.Settings = NewAccountDeviceSettingService(opts...)
	return
}

// Fetches details for a single device.
func (r *AccountDeviceService) Get(ctx context.Context, accountID interface{}, deviceID string, opts ...option.RequestOption) (res *AccountDeviceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/devices/%s", accountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a list of enrolled devices.
func (r *AccountDeviceService) List(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *AccountDeviceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a one-time use admin override code for a device. This relies on the
// **Admin Override** setting being enabled in your device configuration.
func (r *AccountDeviceService) GetOverrideCode(ctx context.Context, accountID interface{}, deviceID string, opts ...option.RequestOption) (res *AccountDeviceGetOverrideCodeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/devices/%s/override_codes", accountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a list of the device settings profiles for an account.
func (r *AccountDeviceService) ListPolicies(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *DeviceSettingsResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policies", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Revokes a list of devices.
func (r *AccountDeviceService) Revoke(ctx context.Context, accountID interface{}, body AccountDeviceRevokeParams, opts ...option.RequestOption) (res *APIResponseSingleTeamsDevices, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/revoke", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unrevokes a list of devices.
func (r *AccountDeviceService) Unrevoke(ctx context.Context, accountID interface{}, body AccountDeviceUnrevokeParams, opts ...option.RequestOption) (res *APIResponseSingleTeamsDevices, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/unrevoke", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type APIResponseCollectionTeamsDevices struct {
	Result     []interface{}                               `json:"result,nullable"`
	ResultInfo APIResponseCollectionTeamsDevicesResultInfo `json:"result_info"`
	JSON       apiResponseCollectionTeamsDevicesJSON       `json:"-"`
	APIResponseTeamsDevices
}

// apiResponseCollectionTeamsDevicesJSON contains the JSON metadata for the struct
// [APIResponseCollectionTeamsDevices]
type apiResponseCollectionTeamsDevicesJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionTeamsDevices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionTeamsDevicesJSON) RawJSON() string {
	return r.raw
}

type APIResponseCollectionTeamsDevicesResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       apiResponseCollectionTeamsDevicesResultInfoJSON `json:"-"`
}

// apiResponseCollectionTeamsDevicesResultInfoJSON contains the JSON metadata for
// the struct [APIResponseCollectionTeamsDevicesResultInfo]
type apiResponseCollectionTeamsDevicesResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionTeamsDevicesResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionTeamsDevicesResultInfoJSON) RawJSON() string {
	return r.raw
}

type APIResponseSingleTeamsDevices struct {
	Result interface{}                       `json:"result,nullable"`
	JSON   apiResponseSingleTeamsDevicesJSON `json:"-"`
	APIResponseTeamsDevices
}

// apiResponseSingleTeamsDevicesJSON contains the JSON metadata for the struct
// [APIResponseSingleTeamsDevices]
type apiResponseSingleTeamsDevicesJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleTeamsDevices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleTeamsDevicesJSON) RawJSON() string {
	return r.raw
}

type DeviceSettingsResponseCollection struct {
	Result []DeviceSettingsPolicy               `json:"result"`
	JSON   deviceSettingsResponseCollectionJSON `json:"-"`
	APIResponseCollectionTeamsDevices
}

// deviceSettingsResponseCollectionJSON contains the JSON metadata for the struct
// [DeviceSettingsResponseCollection]
type deviceSettingsResponseCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingsResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type User struct {
	// UUID
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The enrolled device user's name.
	Name string   `json:"name"`
	JSON userJSON `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *User) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceGetResponse struct {
	Result AccountDeviceGetResponseResult `json:"result"`
	JSON   accountDeviceGetResponseJSON   `json:"-"`
	APIResponseSingleTeamsDevices
}

// accountDeviceGetResponseJSON contains the JSON metadata for the struct
// [AccountDeviceGetResponse]
type accountDeviceGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceGetResponseResult struct {
	// Device ID.
	ID      string                                `json:"id"`
	Account AccountDeviceGetResponseResultAccount `json:"account"`
	// When the device was created.
	Created time.Time `json:"created" format:"date-time"`
	// True if the device was deleted.
	Deleted    bool   `json:"deleted"`
	DeviceType string `json:"device_type"`
	// Deprecated: deprecated
	GatewayDeviceID string `json:"gateway_device_id"`
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// The device's public key.
	Key string `json:"key"`
	// Type of the key.
	KeyType string `json:"key_type"`
	// When the device last connected to Cloudflare services.
	LastSeen time.Time `json:"last_seen" format:"date-time"`
	// The device mac address.
	MacAddress string `json:"mac_address"`
	// The device model name.
	Model string `json:"model"`
	// The device name.
	Name string `json:"name"`
	// The operating system version.
	OsVersion string `json:"os_version"`
	// The device serial number.
	SerialNumber string `json:"serial_number"`
	// Type of the tunnel connection used.
	TunnelType string `json:"tunnel_type"`
	// When the device was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	User    User      `json:"user"`
	// The WARP client version.
	Version string                             `json:"version"`
	JSON    accountDeviceGetResponseResultJSON `json:"-"`
}

// accountDeviceGetResponseResultJSON contains the JSON metadata for the struct
// [AccountDeviceGetResponseResult]
type accountDeviceGetResponseResultJSON struct {
	ID              apijson.Field
	Account         apijson.Field
	Created         apijson.Field
	Deleted         apijson.Field
	DeviceType      apijson.Field
	GatewayDeviceID apijson.Field
	IP              apijson.Field
	Key             apijson.Field
	KeyType         apijson.Field
	LastSeen        apijson.Field
	MacAddress      apijson.Field
	Model           apijson.Field
	Name            apijson.Field
	OsVersion       apijson.Field
	SerialNumber    apijson.Field
	TunnelType      apijson.Field
	Updated         apijson.Field
	User            apijson.Field
	Version         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDeviceGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceGetResponseResultAccount struct {
	// Deprecated: deprecated
	ID string `json:"id"`
	// Deprecated: deprecated
	AccountType string `json:"account_type"`
	// The name of the enrolled account.
	Name string                                    `json:"name"`
	JSON accountDeviceGetResponseResultAccountJSON `json:"-"`
}

// accountDeviceGetResponseResultAccountJSON contains the JSON metadata for the
// struct [AccountDeviceGetResponseResultAccount]
type accountDeviceGetResponseResultAccountJSON struct {
	ID          apijson.Field
	AccountType apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceGetResponseResultAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceGetResponseResultAccountJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceListResponse struct {
	Result []AccountDeviceListResponseResult `json:"result"`
	JSON   accountDeviceListResponseJSON     `json:"-"`
	APIResponseCollectionTeamsDevices
}

// accountDeviceListResponseJSON contains the JSON metadata for the struct
// [AccountDeviceListResponse]
type accountDeviceListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceListResponseResult struct {
	// Device ID.
	ID string `json:"id"`
	// When the device was created.
	Created time.Time `json:"created" format:"date-time"`
	// True if the device was deleted.
	Deleted    bool     `json:"deleted"`
	DeviceType Platform `json:"device_type"`
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// The device's public key.
	Key string `json:"key"`
	// When the device last connected to Cloudflare services.
	LastSeen time.Time `json:"last_seen" format:"date-time"`
	// The device mac address.
	MacAddress string `json:"mac_address"`
	// The device manufacturer name.
	Manufacturer string `json:"manufacturer"`
	// The device model name.
	Model string `json:"model"`
	// The device name.
	Name string `json:"name"`
	// The Linux distro name.
	OsDistroName string `json:"os_distro_name"`
	// The Linux distro revision.
	OsDistroRevision string `json:"os_distro_revision"`
	// The operating system version.
	OsVersion string `json:"os_version"`
	// The operating system version extra parameter.
	OsVersionExtra string `json:"os_version_extra"`
	// When the device was revoked.
	RevokedAt time.Time `json:"revoked_at" format:"date-time"`
	// The device serial number.
	SerialNumber string `json:"serial_number"`
	// When the device was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	User    User      `json:"user"`
	// The WARP client version.
	Version string                              `json:"version"`
	JSON    accountDeviceListResponseResultJSON `json:"-"`
}

// accountDeviceListResponseResultJSON contains the JSON metadata for the struct
// [AccountDeviceListResponseResult]
type accountDeviceListResponseResultJSON struct {
	ID               apijson.Field
	Created          apijson.Field
	Deleted          apijson.Field
	DeviceType       apijson.Field
	IP               apijson.Field
	Key              apijson.Field
	LastSeen         apijson.Field
	MacAddress       apijson.Field
	Manufacturer     apijson.Field
	Model            apijson.Field
	Name             apijson.Field
	OsDistroName     apijson.Field
	OsDistroRevision apijson.Field
	OsVersion        apijson.Field
	OsVersionExtra   apijson.Field
	RevokedAt        apijson.Field
	SerialNumber     apijson.Field
	Updated          apijson.Field
	User             apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDeviceListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceGetOverrideCodeResponse struct {
	Result AccountDeviceGetOverrideCodeResponseResult `json:"result"`
	JSON   accountDeviceGetOverrideCodeResponseJSON   `json:"-"`
	APIResponseCollectionTeamsDevices
}

// accountDeviceGetOverrideCodeResponseJSON contains the JSON metadata for the
// struct [AccountDeviceGetOverrideCodeResponse]
type accountDeviceGetOverrideCodeResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceGetOverrideCodeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceGetOverrideCodeResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceGetOverrideCodeResponseResult struct {
	DisableForTime AccountDeviceGetOverrideCodeResponseResultDisableForTime `json:"disable_for_time"`
	JSON           accountDeviceGetOverrideCodeResponseResultJSON           `json:"-"`
}

// accountDeviceGetOverrideCodeResponseResultJSON contains the JSON metadata for
// the struct [AccountDeviceGetOverrideCodeResponseResult]
type accountDeviceGetOverrideCodeResponseResultJSON struct {
	DisableForTime apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountDeviceGetOverrideCodeResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceGetOverrideCodeResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceGetOverrideCodeResponseResultDisableForTime struct {
	// Override code that is valid for 1 hour.
	Number1 interface{} `json:"1"`
	// Override code that is valid for 12 hour2.
	Number12 interface{} `json:"12"`
	// Override code that is valid for 24 hour.2.
	Number24 interface{} `json:"24"`
	// Override code that is valid for 3 hours.
	Number3 interface{} `json:"3"`
	// Override code that is valid for 6 hours.
	Number6 interface{}                                                  `json:"6"`
	JSON    accountDeviceGetOverrideCodeResponseResultDisableForTimeJSON `json:"-"`
}

// accountDeviceGetOverrideCodeResponseResultDisableForTimeJSON contains the JSON
// metadata for the struct
// [AccountDeviceGetOverrideCodeResponseResultDisableForTime]
type accountDeviceGetOverrideCodeResponseResultDisableForTimeJSON struct {
	Number1     apijson.Field
	Number12    apijson.Field
	Number24    apijson.Field
	Number3     apijson.Field
	Number6     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceGetOverrideCodeResponseResultDisableForTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceGetOverrideCodeResponseResultDisableForTimeJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceRevokeParams struct {
	// A list of device ids to revoke.
	Body []string `json:"body,required"`
}

func (r AccountDeviceRevokeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDeviceUnrevokeParams struct {
	// A list of device ids to unrevoke.
	Body []string `json:"body,required"`
}

func (r AccountDeviceUnrevokeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
