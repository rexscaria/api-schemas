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

// AccountDevicePolicyIncludeService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDevicePolicyIncludeService] method instead.
type AccountDevicePolicyIncludeService struct {
	Options []option.RequestOption
}

// NewAccountDevicePolicyIncludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDevicePolicyIncludeService(opts ...option.RequestOption) (r *AccountDevicePolicyIncludeService) {
	r = &AccountDevicePolicyIncludeService{}
	r.Options = opts
	return
}

// Fetches the list of routes included in the WARP client's tunnel for a specific
// device settings profile.
func (r *AccountDevicePolicyIncludeService) List(ctx context.Context, accountID string, policyID string, opts ...option.RequestOption) (res *SplitTunnelIncludeResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/include", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches the list of routes included in the WARP client's tunnel.
func (r *AccountDevicePolicyIncludeService) GlobalList(ctx context.Context, accountID string, opts ...option.RequestOption) (res *SplitTunnelIncludeResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/include", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sets the list of routes included in the WARP client's tunnel.
func (r *AccountDevicePolicyIncludeService) GlobalSet(ctx context.Context, accountID string, body AccountDevicePolicyIncludeGlobalSetParams, opts ...option.RequestOption) (res *SplitTunnelIncludeResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/include", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Sets the list of routes included in the WARP client's tunnel for a specific
// device settings profile.
func (r *AccountDevicePolicyIncludeService) Set(ctx context.Context, accountID string, policyID string, body AccountDevicePolicyIncludeSetParams, opts ...option.RequestOption) (res *SplitTunnelIncludeResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/include", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SplitTunnelInclude struct {
	// The address in CIDR format to include in the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description"`
	// The domain name to include in the tunnel. If `host` is present, `address` must
	// not be present.
	Host  string                 `json:"host"`
	JSON  splitTunnelIncludeJSON `json:"-"`
	union SplitTunnelIncludeUnion
}

// splitTunnelIncludeJSON contains the JSON metadata for the struct
// [SplitTunnelInclude]
type splitTunnelIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r splitTunnelIncludeJSON) RawJSON() string {
	return r.raw
}

func (r *SplitTunnelInclude) UnmarshalJSON(data []byte) (err error) {
	*r = SplitTunnelInclude{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SplitTunnelIncludeUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddress],
// [SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHost].
func (r SplitTunnelInclude) AsUnion() SplitTunnelIncludeUnion {
	return r.union
}

// Union satisfied by [SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddress]
// or [SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHost].
type SplitTunnelIncludeUnion interface {
	implementsSplitTunnelInclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SplitTunnelIncludeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddress{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHost{}),
		},
	)
}

type SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddress struct {
	// The address in CIDR format to include in the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string                                                          `json:"description"`
	JSON        splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressJSON `json:"-"`
}

// splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressJSON contains the
// JSON metadata for the struct
// [SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddress]
type splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressJSON struct {
	Address     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressJSON) RawJSON() string {
	return r.raw
}

func (r SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddress) implementsSplitTunnelInclude() {}

type SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHost struct {
	// The domain name to include in the tunnel. If `host` is present, `address` must
	// not be present.
	Host string `json:"host,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string                                                       `json:"description"`
	JSON        splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostJSON `json:"-"`
}

// splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostJSON contains the JSON
// metadata for the struct
// [SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHost]
type splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostJSON struct {
	Host        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostJSON) RawJSON() string {
	return r.raw
}

func (r SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHost) implementsSplitTunnelInclude() {}

type SplitTunnelIncludeParam struct {
	// The address in CIDR format to include in the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// The domain name to include in the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host"`
}

func (r SplitTunnelIncludeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitTunnelIncludeParam) implementsSplitTunnelIncludeUnionParam() {}

// Satisfied by [SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressParam],
// [SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostParam],
// [SplitTunnelIncludeParam].
type SplitTunnelIncludeUnionParam interface {
	implementsSplitTunnelIncludeUnionParam()
}

type SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressParam struct {
	// The address in CIDR format to include in the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description"`
}

func (r SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressParam) implementsSplitTunnelIncludeUnionParam() {
}

type SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostParam struct {
	// The domain name to include in the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description"`
}

func (r SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostParam) implementsSplitTunnelIncludeUnionParam() {
}

type SplitTunnelIncludeResponseCollection struct {
	Errors   []MessagesDeviceTestsItems `json:"errors,required"`
	Messages []MessagesDeviceTestsItems `json:"messages,required"`
	Result   []SplitTunnelInclude       `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    SplitTunnelIncludeResponseCollectionSuccess    `json:"success,required"`
	ResultInfo SplitTunnelIncludeResponseCollectionResultInfo `json:"result_info"`
	JSON       splitTunnelIncludeResponseCollectionJSON       `json:"-"`
}

// splitTunnelIncludeResponseCollectionJSON contains the JSON metadata for the
// struct [SplitTunnelIncludeResponseCollection]
type splitTunnelIncludeResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelIncludeResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitTunnelIncludeResponseCollectionJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SplitTunnelIncludeResponseCollectionSuccess bool

const (
	SplitTunnelIncludeResponseCollectionSuccessTrue SplitTunnelIncludeResponseCollectionSuccess = true
)

func (r SplitTunnelIncludeResponseCollectionSuccess) IsKnown() bool {
	switch r {
	case SplitTunnelIncludeResponseCollectionSuccessTrue:
		return true
	}
	return false
}

type SplitTunnelIncludeResponseCollectionResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                            `json:"total_count"`
	JSON       splitTunnelIncludeResponseCollectionResultInfoJSON `json:"-"`
}

// splitTunnelIncludeResponseCollectionResultInfoJSON contains the JSON metadata
// for the struct [SplitTunnelIncludeResponseCollectionResultInfo]
type splitTunnelIncludeResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelIncludeResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitTunnelIncludeResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountDevicePolicyIncludeGlobalSetParams struct {
	Body []SplitTunnelIncludeUnionParam `json:"body,required"`
}

func (r AccountDevicePolicyIncludeGlobalSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDevicePolicyIncludeSetParams struct {
	Body []SplitTunnelIncludeUnionParam `json:"body,required"`
}

func (r AccountDevicePolicyIncludeSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
