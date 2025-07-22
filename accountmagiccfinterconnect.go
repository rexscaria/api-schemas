// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// AccountMagicCfInterconnectService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicCfInterconnectService] method instead.
type AccountMagicCfInterconnectService struct {
	Options []option.RequestOption
}

// NewAccountMagicCfInterconnectService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountMagicCfInterconnectService(opts ...option.RequestOption) (r *AccountMagicCfInterconnectService) {
	r = &AccountMagicCfInterconnectService{}
	r.Options = opts
	return
}

// Lists details for a specific interconnect.
func (r *AccountMagicCfInterconnectService) Get(ctx context.Context, accountID string, cfInterconnectID string, query AccountMagicCfInterconnectGetParams, opts ...option.RequestOption) (res *AccountMagicCfInterconnectGetResponse, err error) {
	if query.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", query.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if cfInterconnectID == "" {
		err = errors.New("missing required cf_interconnect_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", accountID, cfInterconnectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a specific interconnect associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *AccountMagicCfInterconnectService) Update(ctx context.Context, accountID string, cfInterconnectID string, params AccountMagicCfInterconnectUpdateParams, opts ...option.RequestOption) (res *AccountMagicCfInterconnectUpdateResponse, err error) {
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if cfInterconnectID == "" {
		err = errors.New("missing required cf_interconnect_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", accountID, cfInterconnectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Lists interconnects associated with an account.
func (r *AccountMagicCfInterconnectService) List(ctx context.Context, accountID string, query AccountMagicCfInterconnectListParams, opts ...option.RequestOption) (res *AccountMagicCfInterconnectListResponse, err error) {
	if query.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", query.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type MagicAPIResponseSingle struct {
	Errors   []MagicMessageItem                `json:"errors,required"`
	Messages []MagicMessageItem                `json:"messages,required"`
	Result   MagicAPIResponseSingleResultUnion `json:"result,required"`
	// Whether the API call was successful
	Success MagicAPIResponseSingleSuccess `json:"success,required"`
	JSON    magicAPIResponseSingleJSON    `json:"-"`
}

// magicAPIResponseSingleJSON contains the JSON metadata for the struct
// [MagicAPIResponseSingle]
type magicAPIResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicAPIResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [MagicAPIResponseSingleResultArray] or [shared.UnionString].
type MagicAPIResponseSingleResultUnion interface {
	ImplementsMagicAPIResponseSingleResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MagicAPIResponseSingleResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MagicAPIResponseSingleResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type MagicAPIResponseSingleResultArray []interface{}

func (r MagicAPIResponseSingleResultArray) ImplementsMagicAPIResponseSingleResultUnion() {}

// Whether the API call was successful
type MagicAPIResponseSingleSuccess bool

const (
	MagicAPIResponseSingleSuccessTrue MagicAPIResponseSingleSuccess = true
)

func (r MagicAPIResponseSingleSuccess) IsKnown() bool {
	switch r {
	case MagicAPIResponseSingleSuccessTrue:
		return true
	}
	return false
}

// The configuration specific to GRE interconnects.
type MagicGre struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint string       `json:"cloudflare_endpoint"`
	JSON               magicGreJSON `json:"-"`
}

// magicGreJSON contains the JSON metadata for the struct [MagicGre]
type magicGreJSON struct {
	CloudflareEndpoint apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MagicGre) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicGreJSON) RawJSON() string {
	return r.raw
}

// The configuration specific to GRE interconnects.
type MagicGreParam struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint param.Field[string] `json:"cloudflare_endpoint"`
}

func (r MagicGreParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicHealthCheckBase struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicHealthCheckBaseRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target MagicHealthCheckBaseTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicHealthCheckBaseType `json:"type"`
	JSON magicHealthCheckBaseJSON `json:"-"`
}

// magicHealthCheckBaseJSON contains the JSON metadata for the struct
// [MagicHealthCheckBase]
type magicHealthCheckBaseJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicHealthCheckBase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicHealthCheckBaseJSON) RawJSON() string {
	return r.raw
}

// How frequent the health check is run. The default value is `mid`.
type MagicHealthCheckBaseRate string

const (
	MagicHealthCheckBaseRateLow  MagicHealthCheckBaseRate = "low"
	MagicHealthCheckBaseRateMid  MagicHealthCheckBaseRate = "mid"
	MagicHealthCheckBaseRateHigh MagicHealthCheckBaseRate = "high"
)

func (r MagicHealthCheckBaseRate) IsKnown() bool {
	switch r {
	case MagicHealthCheckBaseRateLow, MagicHealthCheckBaseRateMid, MagicHealthCheckBaseRateHigh:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by [MagicHealthCheckBaseTargetMagicHealthCheckTarget] or
// [shared.UnionString].
type MagicHealthCheckBaseTargetUnion interface {
	ImplementsMagicHealthCheckBaseTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MagicHealthCheckBaseTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MagicHealthCheckBaseTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type MagicHealthCheckBaseTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                               `json:"saved"`
	JSON  magicHealthCheckBaseTargetMagicHealthCheckTargetJSON `json:"-"`
}

// magicHealthCheckBaseTargetMagicHealthCheckTargetJSON contains the JSON metadata
// for the struct [MagicHealthCheckBaseTargetMagicHealthCheckTarget]
type magicHealthCheckBaseTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicHealthCheckBaseTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicHealthCheckBaseTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r MagicHealthCheckBaseTargetMagicHealthCheckTarget) ImplementsMagicHealthCheckBaseTargetUnion() {
}

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicHealthCheckBaseType string

const (
	MagicHealthCheckBaseTypeReply   MagicHealthCheckBaseType = "reply"
	MagicHealthCheckBaseTypeRequest MagicHealthCheckBaseType = "request"
)

func (r MagicHealthCheckBaseType) IsKnown() bool {
	switch r {
	case MagicHealthCheckBaseTypeReply, MagicHealthCheckBaseTypeRequest:
		return true
	}
	return false
}

type MagicHealthCheckBaseParam struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[MagicHealthCheckBaseRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target param.Field[MagicHealthCheckBaseTargetUnionParam] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[MagicHealthCheckBaseType] `json:"type"`
}

func (r MagicHealthCheckBaseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Satisfied by [MagicHealthCheckBaseTargetMagicHealthCheckTargetParam],
// [shared.UnionString].
type MagicHealthCheckBaseTargetUnionParam interface {
	ImplementsMagicHealthCheckBaseTargetUnionParam()
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type MagicHealthCheckBaseTargetMagicHealthCheckTargetParam struct {
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved param.Field[string] `json:"saved"`
}

func (r MagicHealthCheckBaseTargetMagicHealthCheckTargetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MagicHealthCheckBaseTargetMagicHealthCheckTargetParam) ImplementsMagicHealthCheckBaseTargetUnionParam() {
}

type MagicInterconnect struct {
	// Tunnel identifier tag.
	ID string `json:"id"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// The configuration specific to GRE interconnects.
	Gre         MagicGre             `json:"gre"`
	HealthCheck MagicHealthCheckBase `json:"health_check"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu int64 `json:"mtu"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	Name string                `json:"name"`
	JSON magicInterconnectJSON `json:"-"`
}

// magicInterconnectJSON contains the JSON metadata for the struct
// [MagicInterconnect]
type magicInterconnectJSON struct {
	ID               apijson.Field
	ColoName         apijson.Field
	CreatedOn        apijson.Field
	Description      apijson.Field
	Gre              apijson.Field
	HealthCheck      apijson.Field
	InterfaceAddress apijson.Field
	ModifiedOn       apijson.Field
	Mtu              apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicInterconnectJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCfInterconnectGetResponse struct {
	Result AccountMagicCfInterconnectGetResponseResult `json:"result"`
	JSON   accountMagicCfInterconnectGetResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicCfInterconnectGetResponseJSON contains the JSON metadata for the
// struct [AccountMagicCfInterconnectGetResponse]
type accountMagicCfInterconnectGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCfInterconnectGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCfInterconnectGetResponseResult struct {
	Interconnect MagicInterconnect                               `json:"interconnect"`
	JSON         accountMagicCfInterconnectGetResponseResultJSON `json:"-"`
}

// accountMagicCfInterconnectGetResponseResultJSON contains the JSON metadata for
// the struct [AccountMagicCfInterconnectGetResponseResult]
type accountMagicCfInterconnectGetResponseResultJSON struct {
	Interconnect apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCfInterconnectGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCfInterconnectUpdateResponse struct {
	Result AccountMagicCfInterconnectUpdateResponseResult `json:"result"`
	JSON   accountMagicCfInterconnectUpdateResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicCfInterconnectUpdateResponseJSON contains the JSON metadata for the
// struct [AccountMagicCfInterconnectUpdateResponse]
type accountMagicCfInterconnectUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCfInterconnectUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCfInterconnectUpdateResponseResult struct {
	Modified             bool                                               `json:"modified"`
	ModifiedInterconnect MagicInterconnect                                  `json:"modified_interconnect"`
	JSON                 accountMagicCfInterconnectUpdateResponseResultJSON `json:"-"`
}

// accountMagicCfInterconnectUpdateResponseResultJSON contains the JSON metadata
// for the struct [AccountMagicCfInterconnectUpdateResponseResult]
type accountMagicCfInterconnectUpdateResponseResultJSON struct {
	Modified             apijson.Field
	ModifiedInterconnect apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCfInterconnectUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCfInterconnectListResponse struct {
	Result AccountMagicCfInterconnectListResponseResult `json:"result"`
	JSON   accountMagicCfInterconnectListResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicCfInterconnectListResponseJSON contains the JSON metadata for the
// struct [AccountMagicCfInterconnectListResponse]
type accountMagicCfInterconnectListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCfInterconnectListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCfInterconnectListResponseResult struct {
	Interconnects []MagicInterconnect                              `json:"interconnects"`
	JSON          accountMagicCfInterconnectListResponseResultJSON `json:"-"`
}

// accountMagicCfInterconnectListResponseResultJSON contains the JSON metadata for
// the struct [AccountMagicCfInterconnectListResponseResult]
type accountMagicCfInterconnectListResponseResultJSON struct {
	Interconnects apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountMagicCfInterconnectListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCfInterconnectListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCfInterconnectGetParams struct {
	XMagicNewHcTarget param.Field[bool] `header:"x-magic-new-hc-target"`
}

type AccountMagicCfInterconnectUpdateParams struct {
	// An optional description of the interconnect.
	Description param.Field[string] `json:"description"`
	// The configuration specific to GRE interconnects.
	Gre         param.Field[MagicGreParam]             `json:"gre"`
	HealthCheck param.Field[MagicHealthCheckBaseParam] `json:"health_check"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu               param.Field[int64] `json:"mtu"`
	XMagicNewHcTarget param.Field[bool]  `header:"x-magic-new-hc-target"`
}

func (r AccountMagicCfInterconnectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicCfInterconnectListParams struct {
	XMagicNewHcTarget param.Field[bool] `header:"x-magic-new-hc-target"`
}
