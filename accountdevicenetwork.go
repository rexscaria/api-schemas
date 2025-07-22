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

// AccountDeviceNetworkService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDeviceNetworkService] method instead.
type AccountDeviceNetworkService struct {
	Options []option.RequestOption
}

// NewAccountDeviceNetworkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDeviceNetworkService(opts ...option.RequestOption) (r *AccountDeviceNetworkService) {
	r = &AccountDeviceNetworkService{}
	r.Options = opts
	return
}

// Creates a new device managed network.
func (r *AccountDeviceNetworkService) New(ctx context.Context, accountID interface{}, body AccountDeviceNetworkNewParams, opts ...option.RequestOption) (res *SingleResponseNetwork, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/networks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches details for a single managed network.
func (r *AccountDeviceNetworkService) Get(ctx context.Context, accountID interface{}, networkID string, opts ...option.RequestOption) (res *SingleResponseNetwork, err error) {
	opts = append(r.Options[:], opts...)
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/devices/networks/%s", accountID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured device managed network.
func (r *AccountDeviceNetworkService) Update(ctx context.Context, accountID interface{}, networkID string, body AccountDeviceNetworkUpdateParams, opts ...option.RequestOption) (res *SingleResponseNetwork, err error) {
	opts = append(r.Options[:], opts...)
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/devices/networks/%s", accountID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches a list of managed networks for an account.
func (r *AccountDeviceNetworkService) List(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *ResponseCollectionDevices, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/networks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a device managed network and fetches a list of the remaining device
// managed networks for an account.
func (r *AccountDeviceNetworkService) Delete(ctx context.Context, accountID interface{}, networkID string, body AccountDeviceNetworkDeleteParams, opts ...option.RequestOption) (res *ResponseCollectionDevices, err error) {
	opts = append(r.Options[:], opts...)
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/devices/networks/%s", accountID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type ConfigRequestNetworkParam struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TlsSockaddr param.Field[string] `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 param.Field[string] `json:"sha256"`
}

func (r ConfigRequestNetworkParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceManagedNetworks struct {
	// The Managed Network TLS Config Response.
	Config DeviceManagedNetworksConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type DeviceTypeManagedNetwork  `json:"type"`
	JSON deviceManagedNetworksJSON `json:"-"`
}

// deviceManagedNetworksJSON contains the JSON metadata for the struct
// [DeviceManagedNetworks]
type deviceManagedNetworksJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceManagedNetworks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceManagedNetworksJSON) RawJSON() string {
	return r.raw
}

// The Managed Network TLS Config Response.
type DeviceManagedNetworksConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TlsSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                          `json:"sha256"`
	JSON   deviceManagedNetworksConfigJSON `json:"-"`
}

// deviceManagedNetworksConfigJSON contains the JSON metadata for the struct
// [DeviceManagedNetworksConfig]
type deviceManagedNetworksConfigJSON struct {
	TlsSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceManagedNetworksConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceManagedNetworksConfigJSON) RawJSON() string {
	return r.raw
}

// The type of device managed network.
type DeviceTypeManagedNetwork string

const (
	DeviceTypeManagedNetworkTls DeviceTypeManagedNetwork = "tls"
)

func (r DeviceTypeManagedNetwork) IsKnown() bool {
	switch r {
	case DeviceTypeManagedNetworkTls:
		return true
	}
	return false
}

type ResponseCollectionDevices struct {
	Result []DeviceManagedNetworks       `json:"result"`
	JSON   responseCollectionDevicesJSON `json:"-"`
	APIResponseCollectionTeamsDevices
}

// responseCollectionDevicesJSON contains the JSON metadata for the struct
// [ResponseCollectionDevices]
type responseCollectionDevicesJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionDevices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionDevicesJSON) RawJSON() string {
	return r.raw
}

type SingleResponseNetwork struct {
	Result DeviceManagedNetworks     `json:"result"`
	JSON   singleResponseNetworkJSON `json:"-"`
	APIResponseSingleTeamsDevices
}

// singleResponseNetworkJSON contains the JSON metadata for the struct
// [SingleResponseNetwork]
type singleResponseNetworkJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseNetworkJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceNetworkNewParams struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config param.Field[ConfigRequestNetworkParam] `json:"config,required"`
	// The name of the device managed network. This name must be unique.
	Name param.Field[string] `json:"name,required"`
	// The type of device managed network.
	Type param.Field[DeviceTypeManagedNetwork] `json:"type,required"`
}

func (r AccountDeviceNetworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDeviceNetworkUpdateParams struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config param.Field[ConfigRequestNetworkParam] `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name param.Field[string] `json:"name"`
	// The type of device managed network.
	Type param.Field[DeviceTypeManagedNetwork] `json:"type"`
}

func (r AccountDeviceNetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDeviceNetworkDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountDeviceNetworkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
