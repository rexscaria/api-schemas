// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountCniCniService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCniCniService] method instead.
type AccountCniCniService struct {
	Options []option.RequestOption
}

// NewAccountCniCniService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountCniCniService(opts ...option.RequestOption) (r *AccountCniCniService) {
	r = &AccountCniCniService{}
	r.Options = opts
	return
}

// Create a new CNI object
func (r *AccountCniCniService) New(ctx context.Context, accountID string, body AccountCniCniNewParams, opts ...option.RequestOption) (res *NscCni, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/cnis", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get information about a CNI object
func (r *AccountCniCniService) Get(ctx context.Context, accountID string, cni string, opts ...option.RequestOption) (res *NscCni, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if cni == "" {
		err = errors.New("missing required cni parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/cnis/%s", accountID, cni)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify stored information about a CNI object
func (r *AccountCniCniService) Update(ctx context.Context, accountID string, cni string, body AccountCniCniUpdateParams, opts ...option.RequestOption) (res *NscCni, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if cni == "" {
		err = errors.New("missing required cni parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/cnis/%s", accountID, cni)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List existing CNI objects
func (r *AccountCniCniService) List(ctx context.Context, accountID string, query AccountCniCniListParams, opts ...option.RequestOption) (res *AccountCniCniListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/cnis", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a specified CNI object
func (r *AccountCniCniService) Delete(ctx context.Context, accountID string, cni string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if cni == "" {
		err = errors.New("missing required cni parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/cnis/%s", accountID, cni)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type NscBgpControl struct {
	// ASN used on the customer end of the BGP session
	CustomerAsn int64 `json:"customer_asn,required"`
	// Extra set of static prefixes to advertise to the customer's end of the session
	ExtraPrefixes []string `json:"extra_prefixes,required" format:"A.B.C.D/N"`
	// MD5 key to use for session authentication.
	//
	// Note that _this is not a security measure_. MD5 is not a valid security
	// mechanism, and the key is not treated as a secret value. This is _only_
	// supported for preventing misconfiguration, not for defending against malicious
	// attacks.
	//
	// The MD5 key, if set, must be of non-zero length and consist only of the
	// following types of character:
	//
	// - ASCII alphanumerics: `[a-zA-Z0-9]`
	// - Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \|`
	//
	// In other words, MD5 keys may contain any printable ASCII character aside from
	// newline (0x0A), quotation mark (`"`), vertical tab (0x0B), carriage return
	// (0x0D), tab (0x09), form feed (0x0C), and the question mark (`?`). Requests
	// specifying an MD5 key with one or more of these disallowed characters will be
	// rejected.
	Md5Key string            `json:"md5_key,nullable"`
	JSON   nscBgpControlJSON `json:"-"`
}

// nscBgpControlJSON contains the JSON metadata for the struct [NscBgpControl]
type nscBgpControlJSON struct {
	CustomerAsn   apijson.Field
	ExtraPrefixes apijson.Field
	Md5Key        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *NscBgpControl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nscBgpControlJSON) RawJSON() string {
	return r.raw
}

type NscBgpControlParam struct {
	// ASN used on the customer end of the BGP session
	CustomerAsn param.Field[int64] `json:"customer_asn,required"`
	// Extra set of static prefixes to advertise to the customer's end of the session
	ExtraPrefixes param.Field[[]string] `json:"extra_prefixes,required" format:"A.B.C.D/N"`
	// MD5 key to use for session authentication.
	//
	// Note that _this is not a security measure_. MD5 is not a valid security
	// mechanism, and the key is not treated as a secret value. This is _only_
	// supported for preventing misconfiguration, not for defending against malicious
	// attacks.
	//
	// The MD5 key, if set, must be of non-zero length and consist only of the
	// following types of character:
	//
	// - ASCII alphanumerics: `[a-zA-Z0-9]`
	// - Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \|`
	//
	// In other words, MD5 keys may contain any printable ASCII character aside from
	// newline (0x0A), quotation mark (`"`), vertical tab (0x0B), carriage return
	// (0x0D), tab (0x09), form feed (0x0C), and the question mark (`?`). Requests
	// specifying an MD5 key with one or more of these disallowed characters will be
	// rejected.
	Md5Key param.Field[string] `json:"md5_key"`
}

func (r NscBgpControlParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NscCni struct {
	ID string `json:"id,required" format:"uuid"`
	// Customer account tag
	Account string `json:"account,required"`
	// Customer end of the point-to-point link
	//
	// This should always be inside the same prefix as `p2p_ip`.
	CustIP string `json:"cust_ip,required" format:"A.B.C.D/N"`
	// Interconnect identifier hosting this CNI
	Interconnect string           `json:"interconnect,required"`
	Magic        NscMagicSettings `json:"magic,required"`
	// Cloudflare end of the point-to-point link
	P2pIP string        `json:"p2p_ip,required" format:"A.B.C.D/N"`
	Bgp   NscBgpControl `json:"bgp"`
	JSON  nscCniJSON    `json:"-"`
}

// nscCniJSON contains the JSON metadata for the struct [NscCni]
type nscCniJSON struct {
	ID           apijson.Field
	Account      apijson.Field
	CustIP       apijson.Field
	Interconnect apijson.Field
	Magic        apijson.Field
	P2pIP        apijson.Field
	Bgp          apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NscCni) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nscCniJSON) RawJSON() string {
	return r.raw
}

type NscCniParam struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
	// Customer account tag
	Account param.Field[string] `json:"account,required"`
	// Customer end of the point-to-point link
	//
	// This should always be inside the same prefix as `p2p_ip`.
	CustIP param.Field[string] `json:"cust_ip,required" format:"A.B.C.D/N"`
	// Interconnect identifier hosting this CNI
	Interconnect param.Field[string]                `json:"interconnect,required"`
	Magic        param.Field[NscMagicSettingsParam] `json:"magic,required"`
	// Cloudflare end of the point-to-point link
	P2pIP param.Field[string]             `json:"p2p_ip,required" format:"A.B.C.D/N"`
	Bgp   param.Field[NscBgpControlParam] `json:"bgp"`
}

func (r NscCniParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NscMagicSettings struct {
	ConduitName string               `json:"conduit_name,required"`
	Description string               `json:"description,required"`
	Mtu         int64                `json:"mtu,required"`
	JSON        nscMagicSettingsJSON `json:"-"`
}

// nscMagicSettingsJSON contains the JSON metadata for the struct
// [NscMagicSettings]
type nscMagicSettingsJSON struct {
	ConduitName apijson.Field
	Description apijson.Field
	Mtu         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NscMagicSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nscMagicSettingsJSON) RawJSON() string {
	return r.raw
}

type NscMagicSettingsParam struct {
	ConduitName param.Field[string] `json:"conduit_name,required"`
	Description param.Field[string] `json:"description,required"`
	Mtu         param.Field[int64]  `json:"mtu,required"`
}

func (r NscMagicSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCniCniListResponse struct {
	Items []NscCni                      `json:"items,required"`
	Next  int64                         `json:"next,nullable"`
	JSON  accountCniCniListResponseJSON `json:"-"`
}

// accountCniCniListResponseJSON contains the JSON metadata for the struct
// [AccountCniCniListResponse]
type accountCniCniListResponseJSON struct {
	Items       apijson.Field
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCniCniListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCniCniListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCniCniNewParams struct {
	// Customer account tag
	Account      param.Field[string]                `json:"account,required"`
	Interconnect param.Field[string]                `json:"interconnect,required"`
	Magic        param.Field[NscMagicSettingsParam] `json:"magic,required"`
	Bgp          param.Field[NscBgpControlParam]    `json:"bgp"`
}

func (r AccountCniCniNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCniCniUpdateParams struct {
	NscCni NscCniParam `json:"nsc_cni,required"`
}

func (r AccountCniCniUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.NscCni)
}

type AccountCniCniListParams struct {
	Cursor param.Field[int64] `query:"cursor"`
	Limit  param.Field[int64] `query:"limit"`
	// If specified, only show CNIs associated with the specified slot
	Slot param.Field[string] `query:"slot"`
	// If specified, only show cnis associated with the specified tunnel id
	TunnelID param.Field[string] `query:"tunnel_id"`
}

// URLQuery serializes [AccountCniCniListParams]'s query parameters as
// `url.Values`.
func (r AccountCniCniListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
