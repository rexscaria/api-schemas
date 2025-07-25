// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
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

// Fetches a single WARP device. Not supported when
// [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/)
// is enabled for the account.
//
// **Deprecated**: please use one of the following endpoints instead:
//
// - GET /accounts/{account_id}/devices/physical-devices/{device_id}
// - GET /accounts/{account_id}/devices/registrations/{registration_id}
//
// Deprecated: deprecated
func (r *AccountDeviceService) Get(ctx context.Context, accountID string, deviceID string, opts ...option.RequestOption) (res *AccountDeviceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/%s", accountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List WARP devices. Not supported when
// [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/)
// is enabled for the account.
//
// **Deprecated**: please use one of the following endpoints instead:
//
// - GET /accounts/{account_id}/devices/physical-devices
// - GET /accounts/{account_id}/devices/registrations
//
// Deprecated: deprecated
func (r *AccountDeviceService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountDeviceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a one-time use admin override code for a device. This relies on the
// **Admin Override** setting being enabled in your device configuration. Not
// supported when
// [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/)
// is enabled for the account. **Deprecated:** please use GET
// /accounts/{account_id}/devices/registrations/{registration_id}/override_codes
// instead.
//
// Deprecated: deprecated
func (r *AccountDeviceService) GetOverrideCode(ctx context.Context, accountID string, deviceID string, opts ...option.RequestOption) (res *AccountDeviceGetOverrideCodeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/%s/override_codes", accountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a list of the device settings profiles for an account.
func (r *AccountDeviceService) ListPolicies(ctx context.Context, accountID string, opts ...option.RequestOption) (res *DeviceSettingsResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policies", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Revokes a list of devices. Not supported when
// [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/)
// is enabled.
//
// **Deprecated**: please use POST
// /accounts/{account_id}/devices/registrations/revoke instead.
//
// Deprecated: deprecated
func (r *AccountDeviceService) Revoke(ctx context.Context, accountID string, body AccountDeviceRevokeParams, opts ...option.RequestOption) (res *APIResponseSingleTeamsDevices, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/revoke", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unrevokes a list of devices. Not supported when
// [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/)
// is enabled.
//
// **Deprecated**: please use POST
// /accounts/{account_id}/devices/registrations/unrevoke instead.
//
// Deprecated: deprecated
func (r *AccountDeviceService) Unrevoke(ctx context.Context, accountID string, body AccountDeviceUnrevokeParams, opts ...option.RequestOption) (res *APIResponseSingleTeamsDevices, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/unrevoke", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type APIResponseSingleTeamsDevices struct {
	Errors   []APIResponseSingleTeamsDevicesError   `json:"errors,required"`
	Messages []APIResponseSingleTeamsDevicesMessage `json:"messages,required"`
	Result   interface{}                            `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success APIResponseSingleTeamsDevicesSuccess `json:"success,required"`
	JSON    apiResponseSingleTeamsDevicesJSON    `json:"-"`
}

// apiResponseSingleTeamsDevicesJSON contains the JSON metadata for the struct
// [APIResponseSingleTeamsDevices]
type apiResponseSingleTeamsDevicesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleTeamsDevices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleTeamsDevicesJSON) RawJSON() string {
	return r.raw
}

type APIResponseSingleTeamsDevicesError struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           APIResponseSingleTeamsDevicesErrorsSource `json:"source"`
	JSON             apiResponseSingleTeamsDevicesErrorJSON    `json:"-"`
}

// apiResponseSingleTeamsDevicesErrorJSON contains the JSON metadata for the struct
// [APIResponseSingleTeamsDevicesError]
type apiResponseSingleTeamsDevicesErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIResponseSingleTeamsDevicesError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleTeamsDevicesErrorJSON) RawJSON() string {
	return r.raw
}

type APIResponseSingleTeamsDevicesErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    apiResponseSingleTeamsDevicesErrorsSourceJSON `json:"-"`
}

// apiResponseSingleTeamsDevicesErrorsSourceJSON contains the JSON metadata for the
// struct [APIResponseSingleTeamsDevicesErrorsSource]
type apiResponseSingleTeamsDevicesErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleTeamsDevicesErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleTeamsDevicesErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type APIResponseSingleTeamsDevicesMessage struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           APIResponseSingleTeamsDevicesMessagesSource `json:"source"`
	JSON             apiResponseSingleTeamsDevicesMessageJSON    `json:"-"`
}

// apiResponseSingleTeamsDevicesMessageJSON contains the JSON metadata for the
// struct [APIResponseSingleTeamsDevicesMessage]
type apiResponseSingleTeamsDevicesMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIResponseSingleTeamsDevicesMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleTeamsDevicesMessageJSON) RawJSON() string {
	return r.raw
}

type APIResponseSingleTeamsDevicesMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    apiResponseSingleTeamsDevicesMessagesSourceJSON `json:"-"`
}

// apiResponseSingleTeamsDevicesMessagesSourceJSON contains the JSON metadata for
// the struct [APIResponseSingleTeamsDevicesMessagesSource]
type apiResponseSingleTeamsDevicesMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleTeamsDevicesMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleTeamsDevicesMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type APIResponseSingleTeamsDevicesSuccess bool

const (
	APIResponseSingleTeamsDevicesSuccessTrue APIResponseSingleTeamsDevicesSuccess = true
)

func (r APIResponseSingleTeamsDevicesSuccess) IsKnown() bool {
	switch r {
	case APIResponseSingleTeamsDevicesSuccessTrue:
		return true
	}
	return false
}

type DeviceSettingsResponseCollection struct {
	Errors   []DeviceSettingsResponseCollectionError   `json:"errors,required"`
	Messages []DeviceSettingsResponseCollectionMessage `json:"messages,required"`
	Result   []DeviceSettingsPolicy                    `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DeviceSettingsResponseCollectionSuccess    `json:"success,required"`
	ResultInfo DeviceSettingsResponseCollectionResultInfo `json:"result_info"`
	JSON       deviceSettingsResponseCollectionJSON       `json:"-"`
}

// deviceSettingsResponseCollectionJSON contains the JSON metadata for the struct
// [DeviceSettingsResponseCollection]
type deviceSettingsResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingsResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type DeviceSettingsResponseCollectionError struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DeviceSettingsResponseCollectionErrorsSource `json:"source"`
	JSON             deviceSettingsResponseCollectionErrorJSON    `json:"-"`
}

// deviceSettingsResponseCollectionErrorJSON contains the JSON metadata for the
// struct [DeviceSettingsResponseCollectionError]
type deviceSettingsResponseCollectionErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceSettingsResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingsResponseCollectionErrorJSON) RawJSON() string {
	return r.raw
}

type DeviceSettingsResponseCollectionErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    deviceSettingsResponseCollectionErrorsSourceJSON `json:"-"`
}

// deviceSettingsResponseCollectionErrorsSourceJSON contains the JSON metadata for
// the struct [DeviceSettingsResponseCollectionErrorsSource]
type deviceSettingsResponseCollectionErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseCollectionErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingsResponseCollectionErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DeviceSettingsResponseCollectionMessage struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DeviceSettingsResponseCollectionMessagesSource `json:"source"`
	JSON             deviceSettingsResponseCollectionMessageJSON    `json:"-"`
}

// deviceSettingsResponseCollectionMessageJSON contains the JSON metadata for the
// struct [DeviceSettingsResponseCollectionMessage]
type deviceSettingsResponseCollectionMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceSettingsResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingsResponseCollectionMessageJSON) RawJSON() string {
	return r.raw
}

type DeviceSettingsResponseCollectionMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    deviceSettingsResponseCollectionMessagesSourceJSON `json:"-"`
}

// deviceSettingsResponseCollectionMessagesSourceJSON contains the JSON metadata
// for the struct [DeviceSettingsResponseCollectionMessagesSource]
type deviceSettingsResponseCollectionMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseCollectionMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingsResponseCollectionMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceSettingsResponseCollectionSuccess bool

const (
	DeviceSettingsResponseCollectionSuccessTrue DeviceSettingsResponseCollectionSuccess = true
)

func (r DeviceSettingsResponseCollectionSuccess) IsKnown() bool {
	switch r {
	case DeviceSettingsResponseCollectionSuccessTrue:
		return true
	}
	return false
}

type DeviceSettingsResponseCollectionResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                        `json:"total_count"`
	JSON       deviceSettingsResponseCollectionResultInfoJSON `json:"-"`
}

// deviceSettingsResponseCollectionResultInfoJSON contains the JSON metadata for
// the struct [DeviceSettingsResponseCollectionResultInfo]
type deviceSettingsResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingsResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingsResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type User struct {
	// UUID.
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
	Errors   []AccountDeviceGetResponseError   `json:"errors,required"`
	Messages []AccountDeviceGetResponseMessage `json:"messages,required"`
	Result   AccountDeviceGetResponseResult    `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success AccountDeviceGetResponseSuccess `json:"success,required"`
	JSON    accountDeviceGetResponseJSON    `json:"-"`
}

// accountDeviceGetResponseJSON contains the JSON metadata for the struct
// [AccountDeviceGetResponse]
type accountDeviceGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceGetResponseError struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           AccountDeviceGetResponseErrorsSource `json:"source"`
	JSON             accountDeviceGetResponseErrorJSON    `json:"-"`
}

// accountDeviceGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountDeviceGetResponseError]
type accountDeviceGetResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDeviceGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceGetResponseErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    accountDeviceGetResponseErrorsSourceJSON `json:"-"`
}

// accountDeviceGetResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountDeviceGetResponseErrorsSource]
type accountDeviceGetResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceGetResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceGetResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceGetResponseMessage struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           AccountDeviceGetResponseMessagesSource `json:"source"`
	JSON             accountDeviceGetResponseMessageJSON    `json:"-"`
}

// accountDeviceGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountDeviceGetResponseMessage]
type accountDeviceGetResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDeviceGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceGetResponseMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    accountDeviceGetResponseMessagesSourceJSON `json:"-"`
}

// accountDeviceGetResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AccountDeviceGetResponseMessagesSource]
type accountDeviceGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceGetResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceGetResponseResult struct {
	// Registration ID. Equal to Device ID except for accounts which enabled
	// [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/).
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

// Whether the API call was successful.
type AccountDeviceGetResponseSuccess bool

const (
	AccountDeviceGetResponseSuccessTrue AccountDeviceGetResponseSuccess = true
)

func (r AccountDeviceGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDeviceGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDeviceListResponse struct {
	Errors   []AccountDeviceListResponseError   `json:"errors,required"`
	Messages []AccountDeviceListResponseMessage `json:"messages,required"`
	Result   []AccountDeviceListResponseResult  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    AccountDeviceListResponseSuccess    `json:"success,required"`
	ResultInfo AccountDeviceListResponseResultInfo `json:"result_info"`
	JSON       accountDeviceListResponseJSON       `json:"-"`
}

// accountDeviceListResponseJSON contains the JSON metadata for the struct
// [AccountDeviceListResponse]
type accountDeviceListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceListResponseError struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           AccountDeviceListResponseErrorsSource `json:"source"`
	JSON             accountDeviceListResponseErrorJSON    `json:"-"`
}

// accountDeviceListResponseErrorJSON contains the JSON metadata for the struct
// [AccountDeviceListResponseError]
type accountDeviceListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDeviceListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceListResponseErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    accountDeviceListResponseErrorsSourceJSON `json:"-"`
}

// accountDeviceListResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountDeviceListResponseErrorsSource]
type accountDeviceListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceListResponseMessage struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           AccountDeviceListResponseMessagesSource `json:"source"`
	JSON             accountDeviceListResponseMessageJSON    `json:"-"`
}

// accountDeviceListResponseMessageJSON contains the JSON metadata for the struct
// [AccountDeviceListResponseMessage]
type accountDeviceListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDeviceListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceListResponseMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    accountDeviceListResponseMessagesSourceJSON `json:"-"`
}

// accountDeviceListResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AccountDeviceListResponseMessagesSource]
type accountDeviceListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceListResponseResult struct {
	// Registration ID. Equal to Device ID except for accounts which enabled
	// [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/).
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

// Whether the API call was successful.
type AccountDeviceListResponseSuccess bool

const (
	AccountDeviceListResponseSuccessTrue AccountDeviceListResponseSuccess = true
)

func (r AccountDeviceListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDeviceListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDeviceListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                 `json:"total_count"`
	JSON       accountDeviceListResponseResultInfoJSON `json:"-"`
}

// accountDeviceListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountDeviceListResponseResultInfo]
type accountDeviceListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceGetOverrideCodeResponse struct {
	Errors   []MessagesDeviceTestsItems `json:"errors,required"`
	Messages []MessagesDeviceTestsItems `json:"messages,required"`
	Result   []interface{}              `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    AccountDeviceGetOverrideCodeResponseSuccess    `json:"success,required"`
	ResultInfo AccountDeviceGetOverrideCodeResponseResultInfo `json:"result_info"`
	JSON       accountDeviceGetOverrideCodeResponseJSON       `json:"-"`
}

// accountDeviceGetOverrideCodeResponseJSON contains the JSON metadata for the
// struct [AccountDeviceGetOverrideCodeResponse]
type accountDeviceGetOverrideCodeResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceGetOverrideCodeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceGetOverrideCodeResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDeviceGetOverrideCodeResponseSuccess bool

const (
	AccountDeviceGetOverrideCodeResponseSuccessTrue AccountDeviceGetOverrideCodeResponseSuccess = true
)

func (r AccountDeviceGetOverrideCodeResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDeviceGetOverrideCodeResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDeviceGetOverrideCodeResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                            `json:"total_count"`
	JSON       accountDeviceGetOverrideCodeResponseResultInfoJSON `json:"-"`
}

// accountDeviceGetOverrideCodeResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountDeviceGetOverrideCodeResponseResultInfo]
type accountDeviceGetOverrideCodeResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceGetOverrideCodeResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceGetOverrideCodeResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceRevokeParams struct {
	// A list of Registration IDs to revoke.
	Body []string `json:"body,required"`
}

func (r AccountDeviceRevokeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDeviceUnrevokeParams struct {
	// A list of Registration IDs to unrevoke.
	Body []string `json:"body,required"`
}

func (r AccountDeviceUnrevokeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
