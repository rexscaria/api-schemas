// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// AccountMnmConfigService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMnmConfigService] method instead.
type AccountMnmConfigService struct {
	Options []option.RequestOption
}

// NewAccountMnmConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMnmConfigService(opts ...option.RequestOption) (r *AccountMnmConfigService) {
	r = &AccountMnmConfigService{}
	r.Options = opts
	return
}

// Create a new network monitoring configuration.
func (r *AccountMnmConfigService) New(ctx context.Context, accountID string, body AccountMnmConfigNewParams, opts ...option.RequestOption) (res *ConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/config", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an existing network monitoring configuration, requires the entire
// configuration to be updated at once.
func (r *AccountMnmConfigService) Update(ctx context.Context, accountID string, body AccountMnmConfigUpdateParams, opts ...option.RequestOption) (res *ConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/config", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists default sampling, router IPs and warp devices for account.
func (r *AccountMnmConfigService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *ConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/config", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an existing network monitoring configuration.
func (r *AccountMnmConfigService) Delete(ctx context.Context, accountID string, body AccountMnmConfigDeleteParams, opts ...option.RequestOption) (res *ConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/config", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Lists default sampling, router IPs, warp devices, and rules for account.
func (r *AccountMnmConfigService) ListFull(ctx context.Context, accountID string, opts ...option.RequestOption) (res *ConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/config/full", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update fields in an existing network monitoring configuration.
func (r *AccountMnmConfigService) UpdateFields(ctx context.Context, accountID string, body AccountMnmConfigUpdateFieldsParams, opts ...option.RequestOption) (res *ConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/config", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type APIResponseMagicVisibilityMnm struct {
	Errors   []MessagesMagicVisibilityMnmItem         `json:"errors,required"`
	Messages []MessagesMagicVisibilityMnmItem         `json:"messages,required"`
	Result   APIResponseMagicVisibilityMnmResultUnion `json:"result,required"`
	// Whether the API call was successful
	Success APIResponseMagicVisibilityMnmSuccess `json:"success,required"`
	JSON    apiResponseMagicVisibilityMnmJSON    `json:"-"`
}

// apiResponseMagicVisibilityMnmJSON contains the JSON metadata for the struct
// [APIResponseMagicVisibilityMnm]
type apiResponseMagicVisibilityMnmJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseMagicVisibilityMnm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseMagicVisibilityMnmJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [APIResponseMagicVisibilityMnmResultArray] or
// [shared.UnionString].
type APIResponseMagicVisibilityMnmResultUnion interface {
	ImplementsAPIResponseMagicVisibilityMnmResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIResponseMagicVisibilityMnmResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIResponseMagicVisibilityMnmResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type APIResponseMagicVisibilityMnmResultArray []interface{}

func (r APIResponseMagicVisibilityMnmResultArray) ImplementsAPIResponseMagicVisibilityMnmResultUnion() {
}

// Whether the API call was successful
type APIResponseMagicVisibilityMnmSuccess bool

const (
	APIResponseMagicVisibilityMnmSuccessTrue APIResponseMagicVisibilityMnmSuccess = true
)

func (r APIResponseMagicVisibilityMnmSuccess) IsKnown() bool {
	switch r {
	case APIResponseMagicVisibilityMnmSuccessTrue:
		return true
	}
	return false
}

type APIResponseSingleMagicVisibility struct {
	Result interface{}                          `json:"result"`
	JSON   apiResponseSingleMagicVisibilityJSON `json:"-"`
	APIResponseMagicVisibilityMnm
}

// apiResponseSingleMagicVisibilityJSON contains the JSON metadata for the struct
// [APIResponseSingleMagicVisibility]
type apiResponseSingleMagicVisibilityJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleMagicVisibility) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleMagicVisibilityJSON) RawJSON() string {
	return r.raw
}

type ConfigSingleResponse struct {
	Result ConfigSingleResponseResult `json:"result"`
	JSON   configSingleResponseJSON   `json:"-"`
	APIResponseSingleMagicVisibility
}

// configSingleResponseJSON contains the JSON metadata for the struct
// [ConfigSingleResponse]
type configSingleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configSingleResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigSingleResponseResult struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name        string                         `json:"name,required"`
	RouterIPs   []string                       `json:"router_ips,required"`
	WarpDevices []WarpDevice                   `json:"warp_devices,required"`
	JSON        configSingleResponseResultJSON `json:"-"`
}

// configSingleResponseResultJSON contains the JSON metadata for the struct
// [ConfigSingleResponseResult]
type configSingleResponseResultJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	WarpDevices     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConfigSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configSingleResponseResultJSON) RawJSON() string {
	return r.raw
}

type MessagesMagicVisibilityMnmItem struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    messagesMagicVisibilityMnmItemJSON `json:"-"`
}

// messagesMagicVisibilityMnmItemJSON contains the JSON metadata for the struct
// [MessagesMagicVisibilityMnmItem]
type messagesMagicVisibilityMnmItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesMagicVisibilityMnmItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesMagicVisibilityMnmItemJSON) RawJSON() string {
	return r.raw
}

// Object representing a warp device with an ID and name.
type WarpDevice struct {
	// Unique identifier for the warp device.
	ID string `json:"id,required"`
	// Name of the warp device.
	Name string `json:"name,required"`
	// IPv4 CIDR of the router sourcing flow data associated with this warp device.
	// Only /32 addresses are currently supported.
	RouterIP string         `json:"router_ip,required"`
	JSON     warpDeviceJSON `json:"-"`
}

// warpDeviceJSON contains the JSON metadata for the struct [WarpDevice]
type warpDeviceJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	RouterIP    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpDeviceJSON) RawJSON() string {
	return r.raw
}

// Object representing a warp device with an ID and name.
type WarpDeviceParam struct {
	// Unique identifier for the warp device.
	ID param.Field[string] `json:"id,required"`
	// Name of the warp device.
	Name param.Field[string] `json:"name,required"`
	// IPv4 CIDR of the router sourcing flow data associated with this warp device.
	// Only /32 addresses are currently supported.
	RouterIP param.Field[string] `json:"router_ip,required"`
}

func (r WarpDeviceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMnmConfigNewParams struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling param.Field[float64] `json:"default_sampling,required"`
	// The account name.
	Name        param.Field[string]            `json:"name,required"`
	RouterIPs   param.Field[[]string]          `json:"router_ips"`
	WarpDevices param.Field[[]WarpDeviceParam] `json:"warp_devices"`
}

func (r AccountMnmConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMnmConfigUpdateParams struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling param.Field[float64] `json:"default_sampling,required"`
	// The account name.
	Name        param.Field[string]            `json:"name,required"`
	RouterIPs   param.Field[[]string]          `json:"router_ips"`
	WarpDevices param.Field[[]WarpDeviceParam] `json:"warp_devices"`
}

func (r AccountMnmConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMnmConfigDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountMnmConfigDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMnmConfigUpdateFieldsParams struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling param.Field[float64] `json:"default_sampling"`
	// The account name.
	Name        param.Field[string]            `json:"name"`
	RouterIPs   param.Field[[]string]          `json:"router_ips"`
	WarpDevices param.Field[[]WarpDeviceParam] `json:"warp_devices"`
}

func (r AccountMnmConfigUpdateFieldsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
