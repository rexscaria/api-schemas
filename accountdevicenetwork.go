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
func (r *AccountDeviceNetworkService) New(ctx context.Context, accountID string, body AccountDeviceNetworkNewParams, opts ...option.RequestOption) (res *SingleResponseNetwork, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/networks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches details for a single managed network.
func (r *AccountDeviceNetworkService) Get(ctx context.Context, accountID string, networkID string, opts ...option.RequestOption) (res *SingleResponseNetwork, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/networks/%s", accountID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured device managed network.
func (r *AccountDeviceNetworkService) Update(ctx context.Context, accountID string, networkID string, body AccountDeviceNetworkUpdateParams, opts ...option.RequestOption) (res *SingleResponseNetwork, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/networks/%s", accountID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches a list of managed networks for an account.
func (r *AccountDeviceNetworkService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *ResponseCollectionDevices, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/networks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a device managed network and fetches a list of the remaining device
// managed networks for an account.
func (r *AccountDeviceNetworkService) Delete(ctx context.Context, accountID string, networkID string, opts ...option.RequestOption) (res *ResponseCollectionDevices, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/networks/%s", accountID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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
	// The configuration object containing information for the WARP client to detect
	// the managed network.
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

// The configuration object containing information for the WARP client to detect
// the managed network.
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
	Errors   []ResponseCollectionDevicesError   `json:"errors,required"`
	Messages []ResponseCollectionDevicesMessage `json:"messages,required"`
	Result   []DeviceManagedNetworks            `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ResponseCollectionDevicesSuccess    `json:"success,required"`
	ResultInfo ResponseCollectionDevicesResultInfo `json:"result_info"`
	JSON       responseCollectionDevicesJSON       `json:"-"`
}

// responseCollectionDevicesJSON contains the JSON metadata for the struct
// [ResponseCollectionDevices]
type responseCollectionDevicesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionDevices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionDevicesJSON) RawJSON() string {
	return r.raw
}

type ResponseCollectionDevicesError struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           ResponseCollectionDevicesErrorsSource `json:"source"`
	JSON             responseCollectionDevicesErrorJSON    `json:"-"`
}

// responseCollectionDevicesErrorJSON contains the JSON metadata for the struct
// [ResponseCollectionDevicesError]
type responseCollectionDevicesErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResponseCollectionDevicesError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionDevicesErrorJSON) RawJSON() string {
	return r.raw
}

type ResponseCollectionDevicesErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    responseCollectionDevicesErrorsSourceJSON `json:"-"`
}

// responseCollectionDevicesErrorsSourceJSON contains the JSON metadata for the
// struct [ResponseCollectionDevicesErrorsSource]
type responseCollectionDevicesErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionDevicesErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionDevicesErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ResponseCollectionDevicesMessage struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           ResponseCollectionDevicesMessagesSource `json:"source"`
	JSON             responseCollectionDevicesMessageJSON    `json:"-"`
}

// responseCollectionDevicesMessageJSON contains the JSON metadata for the struct
// [ResponseCollectionDevicesMessage]
type responseCollectionDevicesMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResponseCollectionDevicesMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionDevicesMessageJSON) RawJSON() string {
	return r.raw
}

type ResponseCollectionDevicesMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    responseCollectionDevicesMessagesSourceJSON `json:"-"`
}

// responseCollectionDevicesMessagesSourceJSON contains the JSON metadata for the
// struct [ResponseCollectionDevicesMessagesSource]
type responseCollectionDevicesMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionDevicesMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionDevicesMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ResponseCollectionDevicesSuccess bool

const (
	ResponseCollectionDevicesSuccessTrue ResponseCollectionDevicesSuccess = true
)

func (r ResponseCollectionDevicesSuccess) IsKnown() bool {
	switch r {
	case ResponseCollectionDevicesSuccessTrue:
		return true
	}
	return false
}

type ResponseCollectionDevicesResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                 `json:"total_count"`
	JSON       responseCollectionDevicesResultInfoJSON `json:"-"`
}

// responseCollectionDevicesResultInfoJSON contains the JSON metadata for the
// struct [ResponseCollectionDevicesResultInfo]
type responseCollectionDevicesResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionDevicesResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionDevicesResultInfoJSON) RawJSON() string {
	return r.raw
}

type SingleResponseNetwork struct {
	Errors   []SingleResponseNetworkError   `json:"errors,required"`
	Messages []SingleResponseNetworkMessage `json:"messages,required"`
	Result   DeviceManagedNetworks          `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success SingleResponseNetworkSuccess `json:"success,required"`
	JSON    singleResponseNetworkJSON    `json:"-"`
}

// singleResponseNetworkJSON contains the JSON metadata for the struct
// [SingleResponseNetwork]
type singleResponseNetworkJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseNetworkJSON) RawJSON() string {
	return r.raw
}

type SingleResponseNetworkError struct {
	Code             int64                             `json:"code,required"`
	Message          string                            `json:"message,required"`
	DocumentationURL string                            `json:"documentation_url"`
	Source           SingleResponseNetworkErrorsSource `json:"source"`
	JSON             singleResponseNetworkErrorJSON    `json:"-"`
}

// singleResponseNetworkErrorJSON contains the JSON metadata for the struct
// [SingleResponseNetworkError]
type singleResponseNetworkErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SingleResponseNetworkError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseNetworkErrorJSON) RawJSON() string {
	return r.raw
}

type SingleResponseNetworkErrorsSource struct {
	Pointer string                                `json:"pointer"`
	JSON    singleResponseNetworkErrorsSourceJSON `json:"-"`
}

// singleResponseNetworkErrorsSourceJSON contains the JSON metadata for the struct
// [SingleResponseNetworkErrorsSource]
type singleResponseNetworkErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseNetworkErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseNetworkErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SingleResponseNetworkMessage struct {
	Code             int64                               `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           SingleResponseNetworkMessagesSource `json:"source"`
	JSON             singleResponseNetworkMessageJSON    `json:"-"`
}

// singleResponseNetworkMessageJSON contains the JSON metadata for the struct
// [SingleResponseNetworkMessage]
type singleResponseNetworkMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SingleResponseNetworkMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseNetworkMessageJSON) RawJSON() string {
	return r.raw
}

type SingleResponseNetworkMessagesSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    singleResponseNetworkMessagesSourceJSON `json:"-"`
}

// singleResponseNetworkMessagesSourceJSON contains the JSON metadata for the
// struct [SingleResponseNetworkMessagesSource]
type singleResponseNetworkMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseNetworkMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseNetworkMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SingleResponseNetworkSuccess bool

const (
	SingleResponseNetworkSuccessTrue SingleResponseNetworkSuccess = true
)

func (r SingleResponseNetworkSuccess) IsKnown() bool {
	switch r {
	case SingleResponseNetworkSuccessTrue:
		return true
	}
	return false
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
