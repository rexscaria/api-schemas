// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountCniInterconnectService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCniInterconnectService] method instead.
type AccountCniInterconnectService struct {
	Options []option.RequestOption
}

// NewAccountCniInterconnectService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountCniInterconnectService(opts ...option.RequestOption) (r *AccountCniInterconnectService) {
	r = &AccountCniInterconnectService{}
	r.Options = opts
	return
}

// Create a new interconnect
func (r *AccountCniInterconnectService) New(ctx context.Context, accountID string, body AccountCniInterconnectNewParams, opts ...option.RequestOption) (res *NscInterconnect, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/interconnects", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get information about an interconnect object
func (r *AccountCniInterconnectService) Get(ctx context.Context, accountID string, icon string, opts ...option.RequestOption) (res *NscInterconnect, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if icon == "" {
		err = errors.New("missing required icon parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/interconnects/%s", accountID, icon)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List existing interconnects
func (r *AccountCniInterconnectService) List(ctx context.Context, accountID string, query AccountCniInterconnectListParams, opts ...option.RequestOption) (res *AccountCniInterconnectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/interconnects", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an interconnect object
func (r *AccountCniInterconnectService) Delete(ctx context.Context, accountID string, icon string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if icon == "" {
		err = errors.New("missing required icon parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/interconnects/%s", accountID, icon)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Generate the Letter of Authorization (LOA) for a given interconnect
func (r *AccountCniInterconnectService) GenerateLoa(ctx context.Context, accountID string, icon string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if icon == "" {
		err = errors.New("missing required icon parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/interconnects/%s/loa", accountID, icon)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Get the current status of an interconnect object
func (r *AccountCniInterconnectService) GetStatus(ctx context.Context, accountID string, icon string, opts ...option.RequestOption) (res *AccountCniInterconnectGetStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if icon == "" {
		err = errors.New("missing required icon parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/interconnects/%s/status", accountID, icon)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type NscInterconnect struct {
	Account  string          `json:"account,required"`
	Name     string          `json:"name,required"`
	Type     string          `json:"type,required"`
	Facility NscFacilityInfo `json:"facility"`
	Owner    string          `json:"owner"`
	Region   string          `json:"region"`
	// A Cloudflare site name.
	Site   string              `json:"site"`
	SlotID string              `json:"slot_id" format:"uuid"`
	Speed  string              `json:"speed"`
	JSON   nscInterconnectJSON `json:"-"`
	union  NscInterconnectUnion
}

// nscInterconnectJSON contains the JSON metadata for the struct [NscInterconnect]
type nscInterconnectJSON struct {
	Account     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Facility    apijson.Field
	Owner       apijson.Field
	Region      apijson.Field
	Site        apijson.Field
	SlotID      apijson.Field
	Speed       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r nscInterconnectJSON) RawJSON() string {
	return r.raw
}

func (r *NscInterconnect) UnmarshalJSON(data []byte) (err error) {
	*r = NscInterconnect{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [NscInterconnectUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [NscInterconnectNscInterconnectPhysicalBody],
// [NscInterconnectNscInterconnectGcpPartnerBody].
func (r NscInterconnect) AsUnion() NscInterconnectUnion {
	return r.union
}

// Union satisfied by [NscInterconnectNscInterconnectPhysicalBody] or
// [NscInterconnectNscInterconnectGcpPartnerBody].
type NscInterconnectUnion interface {
	implementsNscInterconnect()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NscInterconnectUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NscInterconnectNscInterconnectPhysicalBody{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NscInterconnectNscInterconnectGcpPartnerBody{}),
		},
	)
}

type NscInterconnectNscInterconnectPhysicalBody struct {
	Account  string          `json:"account,required"`
	Facility NscFacilityInfo `json:"facility,required"`
	Name     string          `json:"name,required"`
	// A Cloudflare site name.
	Site   string                                         `json:"site,required"`
	SlotID string                                         `json:"slot_id,required" format:"uuid"`
	Speed  string                                         `json:"speed,required"`
	Type   string                                         `json:"type,required"`
	Owner  string                                         `json:"owner"`
	JSON   nscInterconnectNscInterconnectPhysicalBodyJSON `json:"-"`
}

// nscInterconnectNscInterconnectPhysicalBodyJSON contains the JSON metadata for
// the struct [NscInterconnectNscInterconnectPhysicalBody]
type nscInterconnectNscInterconnectPhysicalBodyJSON struct {
	Account     apijson.Field
	Facility    apijson.Field
	Name        apijson.Field
	Site        apijson.Field
	SlotID      apijson.Field
	Speed       apijson.Field
	Type        apijson.Field
	Owner       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NscInterconnectNscInterconnectPhysicalBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nscInterconnectNscInterconnectPhysicalBodyJSON) RawJSON() string {
	return r.raw
}

func (r NscInterconnectNscInterconnectPhysicalBody) implementsNscInterconnect() {}

type NscInterconnectNscInterconnectGcpPartnerBody struct {
	Account string `json:"account,required"`
	Name    string `json:"name,required"`
	Region  string `json:"region,required"`
	Type    string `json:"type,required"`
	Owner   string `json:"owner"`
	// Bandwidth structure as visible through the customer-facing API.
	Speed NscInterconnectNscInterconnectGcpPartnerBodySpeed `json:"speed"`
	JSON  nscInterconnectNscInterconnectGcpPartnerBodyJSON  `json:"-"`
}

// nscInterconnectNscInterconnectGcpPartnerBodyJSON contains the JSON metadata for
// the struct [NscInterconnectNscInterconnectGcpPartnerBody]
type nscInterconnectNscInterconnectGcpPartnerBodyJSON struct {
	Account     apijson.Field
	Name        apijson.Field
	Region      apijson.Field
	Type        apijson.Field
	Owner       apijson.Field
	Speed       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NscInterconnectNscInterconnectGcpPartnerBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nscInterconnectNscInterconnectGcpPartnerBodyJSON) RawJSON() string {
	return r.raw
}

func (r NscInterconnectNscInterconnectGcpPartnerBody) implementsNscInterconnect() {}

// Bandwidth structure as visible through the customer-facing API.
type NscInterconnectNscInterconnectGcpPartnerBodySpeed string

const (
	NscInterconnectNscInterconnectGcpPartnerBodySpeed50M  NscInterconnectNscInterconnectGcpPartnerBodySpeed = "50M"
	NscInterconnectNscInterconnectGcpPartnerBodySpeed100M NscInterconnectNscInterconnectGcpPartnerBodySpeed = "100M"
	NscInterconnectNscInterconnectGcpPartnerBodySpeed200M NscInterconnectNscInterconnectGcpPartnerBodySpeed = "200M"
	NscInterconnectNscInterconnectGcpPartnerBodySpeed300M NscInterconnectNscInterconnectGcpPartnerBodySpeed = "300M"
	NscInterconnectNscInterconnectGcpPartnerBodySpeed400M NscInterconnectNscInterconnectGcpPartnerBodySpeed = "400M"
	NscInterconnectNscInterconnectGcpPartnerBodySpeed500M NscInterconnectNscInterconnectGcpPartnerBodySpeed = "500M"
	NscInterconnectNscInterconnectGcpPartnerBodySpeed1G   NscInterconnectNscInterconnectGcpPartnerBodySpeed = "1G"
	NscInterconnectNscInterconnectGcpPartnerBodySpeed2G   NscInterconnectNscInterconnectGcpPartnerBodySpeed = "2G"
	NscInterconnectNscInterconnectGcpPartnerBodySpeed5G   NscInterconnectNscInterconnectGcpPartnerBodySpeed = "5G"
	NscInterconnectNscInterconnectGcpPartnerBodySpeed10G  NscInterconnectNscInterconnectGcpPartnerBodySpeed = "10G"
	NscInterconnectNscInterconnectGcpPartnerBodySpeed20G  NscInterconnectNscInterconnectGcpPartnerBodySpeed = "20G"
	NscInterconnectNscInterconnectGcpPartnerBodySpeed50G  NscInterconnectNscInterconnectGcpPartnerBodySpeed = "50G"
)

func (r NscInterconnectNscInterconnectGcpPartnerBodySpeed) IsKnown() bool {
	switch r {
	case NscInterconnectNscInterconnectGcpPartnerBodySpeed50M, NscInterconnectNscInterconnectGcpPartnerBodySpeed100M, NscInterconnectNscInterconnectGcpPartnerBodySpeed200M, NscInterconnectNscInterconnectGcpPartnerBodySpeed300M, NscInterconnectNscInterconnectGcpPartnerBodySpeed400M, NscInterconnectNscInterconnectGcpPartnerBodySpeed500M, NscInterconnectNscInterconnectGcpPartnerBodySpeed1G, NscInterconnectNscInterconnectGcpPartnerBodySpeed2G, NscInterconnectNscInterconnectGcpPartnerBodySpeed5G, NscInterconnectNscInterconnectGcpPartnerBodySpeed10G, NscInterconnectNscInterconnectGcpPartnerBodySpeed20G, NscInterconnectNscInterconnectGcpPartnerBodySpeed50G:
		return true
	}
	return false
}

type AccountCniInterconnectListResponse struct {
	Items []NscInterconnect                      `json:"items,required"`
	Next  int64                                  `json:"next,nullable"`
	JSON  accountCniInterconnectListResponseJSON `json:"-"`
}

// accountCniInterconnectListResponseJSON contains the JSON metadata for the struct
// [AccountCniInterconnectListResponse]
type accountCniInterconnectListResponseJSON struct {
	Items       apijson.Field
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCniInterconnectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCniInterconnectListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCniInterconnectGetStatusResponse struct {
	State AccountCniInterconnectGetStatusResponseState `json:"state,required"`
	// Diagnostic information, if available
	Reason string                                      `json:"reason,nullable"`
	JSON   accountCniInterconnectGetStatusResponseJSON `json:"-"`
	union  AccountCniInterconnectGetStatusResponseUnion
}

// accountCniInterconnectGetStatusResponseJSON contains the JSON metadata for the
// struct [AccountCniInterconnectGetStatusResponse]
type accountCniInterconnectGetStatusResponseJSON struct {
	State       apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r accountCniInterconnectGetStatusResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AccountCniInterconnectGetStatusResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AccountCniInterconnectGetStatusResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountCniInterconnectGetStatusResponseUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountCniInterconnectGetStatusResponsePending],
// [AccountCniInterconnectGetStatusResponseDown],
// [AccountCniInterconnectGetStatusResponseUnhealthy],
// [AccountCniInterconnectGetStatusResponseHealthy].
func (r AccountCniInterconnectGetStatusResponse) AsUnion() AccountCniInterconnectGetStatusResponseUnion {
	return r.union
}

// Union satisfied by [AccountCniInterconnectGetStatusResponsePending],
// [AccountCniInterconnectGetStatusResponseDown],
// [AccountCniInterconnectGetStatusResponseUnhealthy] or
// [AccountCniInterconnectGetStatusResponseHealthy].
type AccountCniInterconnectGetStatusResponseUnion interface {
	implementsAccountCniInterconnectGetStatusResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountCniInterconnectGetStatusResponseUnion)(nil)).Elem(),
		"state",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountCniInterconnectGetStatusResponsePending{}),
			DiscriminatorValue: "Pending",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountCniInterconnectGetStatusResponseDown{}),
			DiscriminatorValue: "Down",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountCniInterconnectGetStatusResponseUnhealthy{}),
			DiscriminatorValue: "Unhealthy",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountCniInterconnectGetStatusResponseHealthy{}),
			DiscriminatorValue: "Healthy",
		},
	)
}

type AccountCniInterconnectGetStatusResponsePending struct {
	State AccountCniInterconnectGetStatusResponsePendingState `json:"state,required"`
	JSON  accountCniInterconnectGetStatusResponsePendingJSON  `json:"-"`
}

// accountCniInterconnectGetStatusResponsePendingJSON contains the JSON metadata
// for the struct [AccountCniInterconnectGetStatusResponsePending]
type accountCniInterconnectGetStatusResponsePendingJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCniInterconnectGetStatusResponsePending) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCniInterconnectGetStatusResponsePendingJSON) RawJSON() string {
	return r.raw
}

func (r AccountCniInterconnectGetStatusResponsePending) implementsAccountCniInterconnectGetStatusResponse() {
}

type AccountCniInterconnectGetStatusResponsePendingState string

const (
	AccountCniInterconnectGetStatusResponsePendingStatePending AccountCniInterconnectGetStatusResponsePendingState = "Pending"
)

func (r AccountCniInterconnectGetStatusResponsePendingState) IsKnown() bool {
	switch r {
	case AccountCniInterconnectGetStatusResponsePendingStatePending:
		return true
	}
	return false
}

type AccountCniInterconnectGetStatusResponseDown struct {
	State AccountCniInterconnectGetStatusResponseDownState `json:"state,required"`
	// Diagnostic information, if available
	Reason string                                          `json:"reason,nullable"`
	JSON   accountCniInterconnectGetStatusResponseDownJSON `json:"-"`
}

// accountCniInterconnectGetStatusResponseDownJSON contains the JSON metadata for
// the struct [AccountCniInterconnectGetStatusResponseDown]
type accountCniInterconnectGetStatusResponseDownJSON struct {
	State       apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCniInterconnectGetStatusResponseDown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCniInterconnectGetStatusResponseDownJSON) RawJSON() string {
	return r.raw
}

func (r AccountCniInterconnectGetStatusResponseDown) implementsAccountCniInterconnectGetStatusResponse() {
}

type AccountCniInterconnectGetStatusResponseDownState string

const (
	AccountCniInterconnectGetStatusResponseDownStateDown AccountCniInterconnectGetStatusResponseDownState = "Down"
)

func (r AccountCniInterconnectGetStatusResponseDownState) IsKnown() bool {
	switch r {
	case AccountCniInterconnectGetStatusResponseDownStateDown:
		return true
	}
	return false
}

type AccountCniInterconnectGetStatusResponseUnhealthy struct {
	State AccountCniInterconnectGetStatusResponseUnhealthyState `json:"state,required"`
	// Diagnostic information, if available
	Reason string                                               `json:"reason,nullable"`
	JSON   accountCniInterconnectGetStatusResponseUnhealthyJSON `json:"-"`
}

// accountCniInterconnectGetStatusResponseUnhealthyJSON contains the JSON metadata
// for the struct [AccountCniInterconnectGetStatusResponseUnhealthy]
type accountCniInterconnectGetStatusResponseUnhealthyJSON struct {
	State       apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCniInterconnectGetStatusResponseUnhealthy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCniInterconnectGetStatusResponseUnhealthyJSON) RawJSON() string {
	return r.raw
}

func (r AccountCniInterconnectGetStatusResponseUnhealthy) implementsAccountCniInterconnectGetStatusResponse() {
}

type AccountCniInterconnectGetStatusResponseUnhealthyState string

const (
	AccountCniInterconnectGetStatusResponseUnhealthyStateUnhealthy AccountCniInterconnectGetStatusResponseUnhealthyState = "Unhealthy"
)

func (r AccountCniInterconnectGetStatusResponseUnhealthyState) IsKnown() bool {
	switch r {
	case AccountCniInterconnectGetStatusResponseUnhealthyStateUnhealthy:
		return true
	}
	return false
}

type AccountCniInterconnectGetStatusResponseHealthy struct {
	State AccountCniInterconnectGetStatusResponseHealthyState `json:"state,required"`
	JSON  accountCniInterconnectGetStatusResponseHealthyJSON  `json:"-"`
}

// accountCniInterconnectGetStatusResponseHealthyJSON contains the JSON metadata
// for the struct [AccountCniInterconnectGetStatusResponseHealthy]
type accountCniInterconnectGetStatusResponseHealthyJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCniInterconnectGetStatusResponseHealthy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCniInterconnectGetStatusResponseHealthyJSON) RawJSON() string {
	return r.raw
}

func (r AccountCniInterconnectGetStatusResponseHealthy) implementsAccountCniInterconnectGetStatusResponse() {
}

type AccountCniInterconnectGetStatusResponseHealthyState string

const (
	AccountCniInterconnectGetStatusResponseHealthyStateHealthy AccountCniInterconnectGetStatusResponseHealthyState = "Healthy"
)

func (r AccountCniInterconnectGetStatusResponseHealthyState) IsKnown() bool {
	switch r {
	case AccountCniInterconnectGetStatusResponseHealthyStateHealthy:
		return true
	}
	return false
}

type AccountCniInterconnectGetStatusResponseState string

const (
	AccountCniInterconnectGetStatusResponseStatePending   AccountCniInterconnectGetStatusResponseState = "Pending"
	AccountCniInterconnectGetStatusResponseStateDown      AccountCniInterconnectGetStatusResponseState = "Down"
	AccountCniInterconnectGetStatusResponseStateUnhealthy AccountCniInterconnectGetStatusResponseState = "Unhealthy"
	AccountCniInterconnectGetStatusResponseStateHealthy   AccountCniInterconnectGetStatusResponseState = "Healthy"
)

func (r AccountCniInterconnectGetStatusResponseState) IsKnown() bool {
	switch r {
	case AccountCniInterconnectGetStatusResponseStatePending, AccountCniInterconnectGetStatusResponseStateDown, AccountCniInterconnectGetStatusResponseStateUnhealthy, AccountCniInterconnectGetStatusResponseStateHealthy:
		return true
	}
	return false
}

type AccountCniInterconnectNewParams struct {
	Body AccountCniInterconnectNewParamsBodyUnion `json:"body,required"`
}

func (r AccountCniInterconnectNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountCniInterconnectNewParamsBody struct {
	Account param.Field[string] `json:"account,required"`
	Type    param.Field[string] `json:"type,required"`
	// Bandwidth structure as visible through the customer-facing API.
	Bandwidth param.Field[AccountCniInterconnectNewParamsBodyBandwidth] `json:"bandwidth"`
	// Pairing key provided by GCP
	PairingKey param.Field[string] `json:"pairing_key"`
	SlotID     param.Field[string] `json:"slot_id" format:"uuid"`
	Speed      param.Field[string] `json:"speed"`
}

func (r AccountCniInterconnectNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountCniInterconnectNewParamsBody) implementsAccountCniInterconnectNewParamsBodyUnion() {}

// Satisfied by
// [AccountCniInterconnectNewParamsBodyNscInterconnectCreatePhysicalBody],
// [AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBody],
// [AccountCniInterconnectNewParamsBody].
type AccountCniInterconnectNewParamsBodyUnion interface {
	implementsAccountCniInterconnectNewParamsBodyUnion()
}

type AccountCniInterconnectNewParamsBodyNscInterconnectCreatePhysicalBody struct {
	Account param.Field[string] `json:"account,required"`
	SlotID  param.Field[string] `json:"slot_id,required" format:"uuid"`
	Type    param.Field[string] `json:"type,required"`
	Speed   param.Field[string] `json:"speed"`
}

func (r AccountCniInterconnectNewParamsBodyNscInterconnectCreatePhysicalBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountCniInterconnectNewParamsBodyNscInterconnectCreatePhysicalBody) implementsAccountCniInterconnectNewParamsBodyUnion() {
}

type AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBody struct {
	Account param.Field[string] `json:"account,required"`
	// Bandwidth structure as visible through the customer-facing API.
	Bandwidth param.Field[AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth] `json:"bandwidth,required"`
	// Pairing key provided by GCP
	PairingKey param.Field[string] `json:"pairing_key,required"`
	Type       param.Field[string] `json:"type,required"`
}

func (r AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBody) implementsAccountCniInterconnectNewParamsBodyUnion() {
}

// Bandwidth structure as visible through the customer-facing API.
type AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth string

const (
	AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth50M  AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "50M"
	AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth100M AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "100M"
	AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth200M AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "200M"
	AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth300M AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "300M"
	AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth400M AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "400M"
	AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth500M AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "500M"
	AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth1G   AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "1G"
	AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth2G   AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "2G"
	AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth5G   AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "5G"
	AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth10G  AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "10G"
	AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth20G  AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "20G"
	AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth50G  AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "50G"
)

func (r AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth) IsKnown() bool {
	switch r {
	case AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth50M, AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth100M, AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth200M, AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth300M, AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth400M, AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth500M, AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth1G, AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth2G, AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth5G, AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth10G, AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth20G, AccountCniInterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth50G:
		return true
	}
	return false
}

// Bandwidth structure as visible through the customer-facing API.
type AccountCniInterconnectNewParamsBodyBandwidth string

const (
	AccountCniInterconnectNewParamsBodyBandwidth50M  AccountCniInterconnectNewParamsBodyBandwidth = "50M"
	AccountCniInterconnectNewParamsBodyBandwidth100M AccountCniInterconnectNewParamsBodyBandwidth = "100M"
	AccountCniInterconnectNewParamsBodyBandwidth200M AccountCniInterconnectNewParamsBodyBandwidth = "200M"
	AccountCniInterconnectNewParamsBodyBandwidth300M AccountCniInterconnectNewParamsBodyBandwidth = "300M"
	AccountCniInterconnectNewParamsBodyBandwidth400M AccountCniInterconnectNewParamsBodyBandwidth = "400M"
	AccountCniInterconnectNewParamsBodyBandwidth500M AccountCniInterconnectNewParamsBodyBandwidth = "500M"
	AccountCniInterconnectNewParamsBodyBandwidth1G   AccountCniInterconnectNewParamsBodyBandwidth = "1G"
	AccountCniInterconnectNewParamsBodyBandwidth2G   AccountCniInterconnectNewParamsBodyBandwidth = "2G"
	AccountCniInterconnectNewParamsBodyBandwidth5G   AccountCniInterconnectNewParamsBodyBandwidth = "5G"
	AccountCniInterconnectNewParamsBodyBandwidth10G  AccountCniInterconnectNewParamsBodyBandwidth = "10G"
	AccountCniInterconnectNewParamsBodyBandwidth20G  AccountCniInterconnectNewParamsBodyBandwidth = "20G"
	AccountCniInterconnectNewParamsBodyBandwidth50G  AccountCniInterconnectNewParamsBodyBandwidth = "50G"
)

func (r AccountCniInterconnectNewParamsBodyBandwidth) IsKnown() bool {
	switch r {
	case AccountCniInterconnectNewParamsBodyBandwidth50M, AccountCniInterconnectNewParamsBodyBandwidth100M, AccountCniInterconnectNewParamsBodyBandwidth200M, AccountCniInterconnectNewParamsBodyBandwidth300M, AccountCniInterconnectNewParamsBodyBandwidth400M, AccountCniInterconnectNewParamsBodyBandwidth500M, AccountCniInterconnectNewParamsBodyBandwidth1G, AccountCniInterconnectNewParamsBodyBandwidth2G, AccountCniInterconnectNewParamsBodyBandwidth5G, AccountCniInterconnectNewParamsBodyBandwidth10G, AccountCniInterconnectNewParamsBodyBandwidth20G, AccountCniInterconnectNewParamsBodyBandwidth50G:
		return true
	}
	return false
}

type AccountCniInterconnectListParams struct {
	Cursor param.Field[int64] `query:"cursor"`
	Limit  param.Field[int64] `query:"limit"`
	// If specified, only show interconnects located at the given site
	Site param.Field[string] `query:"site"`
	// If specified, only show interconnects of the given type
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [AccountCniInterconnectListParams]'s query parameters as
// `url.Values`.
func (r AccountCniInterconnectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
