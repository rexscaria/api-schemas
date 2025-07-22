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
	"github.com/tidwall/gjson"
)

// AccountDevicePolicyExcludeService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDevicePolicyExcludeService] method instead.
type AccountDevicePolicyExcludeService struct {
	Options []option.RequestOption
}

// NewAccountDevicePolicyExcludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDevicePolicyExcludeService(opts ...option.RequestOption) (r *AccountDevicePolicyExcludeService) {
	r = &AccountDevicePolicyExcludeService{}
	r.Options = opts
	return
}

// Fetches the list of routes excluded from the WARP client's tunnel for a specific
// device settings profile.
func (r *AccountDevicePolicyExcludeService) List(ctx context.Context, accountID interface{}, policyID string, opts ...option.RequestOption) (res *SplitTunnelResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/exclude", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches the list of routes excluded from the WARP client's tunnel.
func (r *AccountDevicePolicyExcludeService) GlobalList(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *SplitTunnelResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/exclude", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sets the list of routes excluded from the WARP client's tunnel.
func (r *AccountDevicePolicyExcludeService) GlobalSet(ctx context.Context, accountID interface{}, body AccountDevicePolicyExcludeGlobalSetParams, opts ...option.RequestOption) (res *SplitTunnelResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/exclude", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Sets the list of routes excluded from the WARP client's tunnel for a specific
// device settings profile.
func (r *AccountDevicePolicyExcludeService) Set(ctx context.Context, accountID interface{}, policyID string, body AccountDevicePolicyExcludeSetParams, opts ...option.RequestOption) (res *SplitTunnelResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/exclude", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SplitTunnel struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host  string          `json:"host"`
	JSON  splitTunnelJSON `json:"-"`
	union SplitTunnelUnion
}

// splitTunnelJSON contains the JSON metadata for the struct [SplitTunnel]
type splitTunnelJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r splitTunnelJSON) RawJSON() string {
	return r.raw
}

func (r *SplitTunnel) UnmarshalJSON(data []byte) (err error) {
	*r = SplitTunnel{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SplitTunnelUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [SplitTunnelTeamsDevicesExcludeSplitTunnelWithAddress],
// [SplitTunnelTeamsDevicesExcludeSplitTunnelWithHost].
func (r SplitTunnel) AsUnion() SplitTunnelUnion {
	return r.union
}

// Union satisfied by [SplitTunnelTeamsDevicesExcludeSplitTunnelWithAddress] or
// [SplitTunnelTeamsDevicesExcludeSplitTunnelWithHost].
type SplitTunnelUnion interface {
	implementsSplitTunnel()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SplitTunnelUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplitTunnelTeamsDevicesExcludeSplitTunnelWithAddress{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplitTunnelTeamsDevicesExcludeSplitTunnelWithHost{}),
		},
	)
}

type SplitTunnelTeamsDevicesExcludeSplitTunnelWithAddress struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string                                                   `json:"description"`
	JSON        splitTunnelTeamsDevicesExcludeSplitTunnelWithAddressJSON `json:"-"`
}

// splitTunnelTeamsDevicesExcludeSplitTunnelWithAddressJSON contains the JSON
// metadata for the struct [SplitTunnelTeamsDevicesExcludeSplitTunnelWithAddress]
type splitTunnelTeamsDevicesExcludeSplitTunnelWithAddressJSON struct {
	Address     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelTeamsDevicesExcludeSplitTunnelWithAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitTunnelTeamsDevicesExcludeSplitTunnelWithAddressJSON) RawJSON() string {
	return r.raw
}

func (r SplitTunnelTeamsDevicesExcludeSplitTunnelWithAddress) implementsSplitTunnel() {}

type SplitTunnelTeamsDevicesExcludeSplitTunnelWithHost struct {
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string `json:"host,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string                                                `json:"description"`
	JSON        splitTunnelTeamsDevicesExcludeSplitTunnelWithHostJSON `json:"-"`
}

// splitTunnelTeamsDevicesExcludeSplitTunnelWithHostJSON contains the JSON metadata
// for the struct [SplitTunnelTeamsDevicesExcludeSplitTunnelWithHost]
type splitTunnelTeamsDevicesExcludeSplitTunnelWithHostJSON struct {
	Host        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelTeamsDevicesExcludeSplitTunnelWithHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitTunnelTeamsDevicesExcludeSplitTunnelWithHostJSON) RawJSON() string {
	return r.raw
}

func (r SplitTunnelTeamsDevicesExcludeSplitTunnelWithHost) implementsSplitTunnel() {}

type SplitTunnelParam struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host"`
}

func (r SplitTunnelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitTunnelParam) implementsSplitTunnelUnionParam() {}

// Satisfied by [SplitTunnelTeamsDevicesExcludeSplitTunnelWithAddressParam],
// [SplitTunnelTeamsDevicesExcludeSplitTunnelWithHostParam], [SplitTunnelParam].
type SplitTunnelUnionParam interface {
	implementsSplitTunnelUnionParam()
}

type SplitTunnelTeamsDevicesExcludeSplitTunnelWithAddressParam struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description"`
}

func (r SplitTunnelTeamsDevicesExcludeSplitTunnelWithAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitTunnelTeamsDevicesExcludeSplitTunnelWithAddressParam) implementsSplitTunnelUnionParam() {
}

type SplitTunnelTeamsDevicesExcludeSplitTunnelWithHostParam struct {
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description"`
}

func (r SplitTunnelTeamsDevicesExcludeSplitTunnelWithHostParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitTunnelTeamsDevicesExcludeSplitTunnelWithHostParam) implementsSplitTunnelUnionParam() {}

type SplitTunnelResponseCollection struct {
	Result []SplitTunnel                     `json:"result"`
	JSON   splitTunnelResponseCollectionJSON `json:"-"`
	APIResponseCollectionTeamsDevices
}

// splitTunnelResponseCollectionJSON contains the JSON metadata for the struct
// [SplitTunnelResponseCollection]
type splitTunnelResponseCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitTunnelResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type AccountDevicePolicyExcludeGlobalSetParams struct {
	Body []SplitTunnelUnionParam `json:"body,required"`
}

func (r AccountDevicePolicyExcludeGlobalSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDevicePolicyExcludeSetParams struct {
	Body []SplitTunnelUnionParam `json:"body,required"`
}

func (r AccountDevicePolicyExcludeSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
