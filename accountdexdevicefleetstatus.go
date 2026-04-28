// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountDexDeviceFleetStatusService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDexDeviceFleetStatusService] method instead.
type AccountDexDeviceFleetStatusService struct {
	Options []option.RequestOption
}

// NewAccountDexDeviceFleetStatusService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDexDeviceFleetStatusService(opts ...option.RequestOption) (r *AccountDexDeviceFleetStatusService) {
	r = &AccountDexDeviceFleetStatusService{}
	r.Options = opts
	return
}

// Get the live status of a latest device given device_id from the device_state
// table
func (r *AccountDexDeviceFleetStatusService) GetLiveStatus(ctx context.Context, accountID string, deviceID string, query AccountDexDeviceFleetStatusGetLiveStatusParams, opts ...option.RequestOption) (res *Device, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dex/devices/%s/fleet-status/live", accountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type Device struct {
	// Cloudflare colo
	Colo string `json:"colo" api:"required"`
	// Device identifier (UUID v4)
	DeviceID string `json:"deviceId" api:"required"`
	// The mode under which the WARP client is run
	Mode string `json:"mode" api:"required"`
	// Operating system
	Platform string `json:"platform" api:"required"`
	// Network status
	Status string `json:"status" api:"required"`
	// Timestamp in ISO format
	Timestamp string `json:"timestamp" api:"required"`
	// WARP client version
	Version         string                `json:"version" api:"required"`
	AlwaysOn        bool                  `json:"alwaysOn" api:"nullable"`
	BatteryCharging bool                  `json:"batteryCharging" api:"nullable"`
	BatteryCycles   int64                 `json:"batteryCycles" api:"nullable"`
	BatteryPct      float64               `json:"batteryPct" api:"nullable"`
	ConnectionType  string                `json:"connectionType" api:"nullable"`
	CPUPct          float64               `json:"cpuPct" api:"nullable"`
	CPUPctByApp     [][]DeviceCPUPctByApp `json:"cpuPctByApp" api:"nullable"`
	DeviceIpv4      IPInfoFleetStatus     `json:"deviceIpv4"`
	DeviceIpv6      IPInfoFleetStatus     `json:"deviceIpv6"`
	// Device identifier (human readable)
	DeviceName         string            `json:"deviceName"`
	DiskReadBps        int64             `json:"diskReadBps" api:"nullable"`
	DiskUsagePct       float64           `json:"diskUsagePct" api:"nullable"`
	DiskWriteBps       int64             `json:"diskWriteBps" api:"nullable"`
	DohSubdomain       string            `json:"dohSubdomain" api:"nullable"`
	EstimatedLossPct   float64           `json:"estimatedLossPct" api:"nullable"`
	FirewallEnabled    bool              `json:"firewallEnabled" api:"nullable"`
	GatewayIpv4        IPInfoFleetStatus `json:"gatewayIpv4"`
	GatewayIpv6        IPInfoFleetStatus `json:"gatewayIpv6"`
	HandshakeLatencyMs float64           `json:"handshakeLatencyMs" api:"nullable"`
	IspIpv4            IPInfoFleetStatus `json:"ispIpv4"`
	IspIpv6            IPInfoFleetStatus `json:"ispIpv6"`
	Metal              string            `json:"metal" api:"nullable"`
	NetworkRcvdBps     int64             `json:"networkRcvdBps" api:"nullable"`
	NetworkSentBps     int64             `json:"networkSentBps" api:"nullable"`
	NetworkSsid        string            `json:"networkSsid" api:"nullable"`
	// User contact email address
	PersonEmail     string                    `json:"personEmail"`
	RamAvailableKB  int64                     `json:"ramAvailableKb" api:"nullable"`
	RamUsedPct      float64                   `json:"ramUsedPct" api:"nullable"`
	RamUsedPctByApp [][]DeviceRamUsedPctByApp `json:"ramUsedPctByApp" api:"nullable"`
	SwitchLocked    bool                      `json:"switchLocked" api:"nullable"`
	WifiStrengthDbm int64                     `json:"wifiStrengthDbm" api:"nullable"`
	JSON            deviceJSON                `json:"-"`
}

// deviceJSON contains the JSON metadata for the struct [Device]
type deviceJSON struct {
	Colo               apijson.Field
	DeviceID           apijson.Field
	Mode               apijson.Field
	Platform           apijson.Field
	Status             apijson.Field
	Timestamp          apijson.Field
	Version            apijson.Field
	AlwaysOn           apijson.Field
	BatteryCharging    apijson.Field
	BatteryCycles      apijson.Field
	BatteryPct         apijson.Field
	ConnectionType     apijson.Field
	CPUPct             apijson.Field
	CPUPctByApp        apijson.Field
	DeviceIpv4         apijson.Field
	DeviceIpv6         apijson.Field
	DeviceName         apijson.Field
	DiskReadBps        apijson.Field
	DiskUsagePct       apijson.Field
	DiskWriteBps       apijson.Field
	DohSubdomain       apijson.Field
	EstimatedLossPct   apijson.Field
	FirewallEnabled    apijson.Field
	GatewayIpv4        apijson.Field
	GatewayIpv6        apijson.Field
	HandshakeLatencyMs apijson.Field
	IspIpv4            apijson.Field
	IspIpv6            apijson.Field
	Metal              apijson.Field
	NetworkRcvdBps     apijson.Field
	NetworkSentBps     apijson.Field
	NetworkSsid        apijson.Field
	PersonEmail        apijson.Field
	RamAvailableKB     apijson.Field
	RamUsedPct         apijson.Field
	RamUsedPctByApp    apijson.Field
	SwitchLocked       apijson.Field
	WifiStrengthDbm    apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Device) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceJSON) RawJSON() string {
	return r.raw
}

type DeviceCPUPctByApp struct {
	CPUPct float64               `json:"cpu_pct"`
	Name   string                `json:"name"`
	JSON   deviceCPUPctByAppJSON `json:"-"`
}

// deviceCPUPctByAppJSON contains the JSON metadata for the struct
// [DeviceCPUPctByApp]
type deviceCPUPctByAppJSON struct {
	CPUPct      apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceCPUPctByApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceCPUPctByAppJSON) RawJSON() string {
	return r.raw
}

type DeviceRamUsedPctByApp struct {
	Name       string                    `json:"name"`
	RamUsedPct float64                   `json:"ram_used_pct"`
	JSON       deviceRamUsedPctByAppJSON `json:"-"`
}

// deviceRamUsedPctByAppJSON contains the JSON metadata for the struct
// [DeviceRamUsedPctByApp]
type deviceRamUsedPctByAppJSON struct {
	Name        apijson.Field
	RamUsedPct  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRamUsedPctByApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRamUsedPctByAppJSON) RawJSON() string {
	return r.raw
}

type IPInfoFleetStatus struct {
	Address  string                    `json:"address" api:"nullable"`
	Asn      int64                     `json:"asn" api:"nullable"`
	Aso      string                    `json:"aso" api:"nullable"`
	Location IPInfoFleetStatusLocation `json:"location"`
	Netmask  string                    `json:"netmask" api:"nullable"`
	Version  string                    `json:"version" api:"nullable"`
	JSON     ipInfoFleetStatusJSON     `json:"-"`
}

// ipInfoFleetStatusJSON contains the JSON metadata for the struct
// [IPInfoFleetStatus]
type ipInfoFleetStatusJSON struct {
	Address     apijson.Field
	Asn         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPInfoFleetStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipInfoFleetStatusJSON) RawJSON() string {
	return r.raw
}

type IPInfoFleetStatusLocation struct {
	City       string                        `json:"city" api:"nullable"`
	CountryISO string                        `json:"country_iso" api:"nullable"`
	StateISO   string                        `json:"state_iso" api:"nullable"`
	Zip        string                        `json:"zip" api:"nullable"`
	JSON       ipInfoFleetStatusLocationJSON `json:"-"`
}

// ipInfoFleetStatusLocationJSON contains the JSON metadata for the struct
// [IPInfoFleetStatusLocation]
type ipInfoFleetStatusLocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPInfoFleetStatusLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipInfoFleetStatusLocationJSON) RawJSON() string {
	return r.raw
}

type AccountDexDeviceFleetStatusGetLiveStatusParams struct {
	// Number of minutes before current time
	SinceMinutes param.Field[float64] `query:"since_minutes" api:"required"`
	// List of data centers to filter results
	Colo param.Field[string] `query:"colo"`
	// Number of minutes before current time
	TimeNow param.Field[string] `query:"time_now"`
}

// URLQuery serializes [AccountDexDeviceFleetStatusGetLiveStatusParams]'s query
// parameters as `url.Values`.
func (r AccountDexDeviceFleetStatusGetLiveStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
