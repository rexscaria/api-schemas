// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
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

// IPService contains methods and other services that help with interacting with
// the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPService] method instead.
type IPService struct {
	Options []option.RequestOption
}

// NewIPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPService(opts ...option.RequestOption) (r *IPService) {
	r = &IPService{}
	r.Options = opts
	return
}

// Get IPs used on the Cloudflare/JD Cloud network, see
// https://www.cloudflare.com/ips for Cloudflare IPs or
// https://developers.cloudflare.com/china-network/reference/infrastructure/ for JD
// Cloud IPs.
func (r *IPService) List(ctx context.Context, query IPListParams, opts ...option.RequestOption) (res *IPListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AddressingAPIResponseSingle struct {
	Errors   []AddressingMessages `json:"errors,required"`
	Messages []AddressingMessages `json:"messages,required"`
	// Whether the API call was successful
	Success AddressingAPIResponseSingleSuccess `json:"success,required"`
	JSON    addressingAPIResponseSingleJSON    `json:"-"`
}

// addressingAPIResponseSingleJSON contains the JSON metadata for the struct
// [AddressingAPIResponseSingle]
type addressingAPIResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAPIResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingAPIResponseSingleSuccess bool

const (
	AddressingAPIResponseSingleSuccessTrue AddressingAPIResponseSingleSuccess = true
)

func (r AddressingAPIResponseSingleSuccess) IsKnown() bool {
	switch r {
	case AddressingAPIResponseSingleSuccessTrue:
		return true
	}
	return false
}

type IPListResponse struct {
	Result IPListResponseResult `json:"result"`
	JSON   ipListResponseJSON   `json:"-"`
	AddressingAPIResponseSingle
}

// ipListResponseJSON contains the JSON metadata for the struct [IPListResponse]
type ipListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseJSON) RawJSON() string {
	return r.raw
}

type IPListResponseResult struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// This field can have the runtime type of [[]string].
	Ipv4Cidrs interface{} `json:"ipv4_cidrs"`
	// This field can have the runtime type of [[]string].
	Ipv6Cidrs interface{} `json:"ipv6_cidrs"`
	// This field can have the runtime type of [[]string].
	JdcloudCidrs interface{}              `json:"jdcloud_cidrs"`
	JSON         ipListResponseResultJSON `json:"-"`
	union        IPListResponseResultUnion
}

// ipListResponseResultJSON contains the JSON metadata for the struct
// [IPListResponseResult]
type ipListResponseResultJSON struct {
	Etag         apijson.Field
	Ipv4Cidrs    apijson.Field
	Ipv6Cidrs    apijson.Field
	JdcloudCidrs apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r ipListResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *IPListResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = IPListResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [IPListResponseResultUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [IPListResponseResultAddressingIPs],
// [IPListResponseResultAddressingIPsJdcloud].
func (r IPListResponseResult) AsUnion() IPListResponseResultUnion {
	return r.union
}

// Union satisfied by [IPListResponseResultAddressingIPs] or
// [IPListResponseResultAddressingIPsJdcloud].
type IPListResponseResultUnion interface {
	implementsIPListResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IPListResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPListResponseResultAddressingIPs{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPListResponseResultAddressingIPsJdcloud{}),
		},
	)
}

type IPListResponseResultAddressingIPs struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// List of Cloudflare IPv4 CIDR addresses.
	Ipv4Cidrs []string `json:"ipv4_cidrs"`
	// List of Cloudflare IPv6 CIDR addresses.
	Ipv6Cidrs []string                              `json:"ipv6_cidrs"`
	JSON      ipListResponseResultAddressingIPsJSON `json:"-"`
}

// ipListResponseResultAddressingIPsJSON contains the JSON metadata for the struct
// [IPListResponseResultAddressingIPs]
type ipListResponseResultAddressingIPsJSON struct {
	Etag        apijson.Field
	Ipv4Cidrs   apijson.Field
	Ipv6Cidrs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseResultAddressingIPs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseResultAddressingIPsJSON) RawJSON() string {
	return r.raw
}

func (r IPListResponseResultAddressingIPs) implementsIPListResponseResult() {}

type IPListResponseResultAddressingIPsJdcloud struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// List of Cloudflare IPv4 CIDR addresses.
	Ipv4Cidrs []string `json:"ipv4_cidrs"`
	// List of Cloudflare IPv6 CIDR addresses.
	Ipv6Cidrs []string `json:"ipv6_cidrs"`
	// List IPv4 and IPv6 CIDRs, only populated if `?networks=jdcloud` is used.
	JdcloudCidrs []string                                     `json:"jdcloud_cidrs"`
	JSON         ipListResponseResultAddressingIPsJdcloudJSON `json:"-"`
}

// ipListResponseResultAddressingIPsJdcloudJSON contains the JSON metadata for the
// struct [IPListResponseResultAddressingIPsJdcloud]
type ipListResponseResultAddressingIPsJdcloudJSON struct {
	Etag         apijson.Field
	Ipv4Cidrs    apijson.Field
	Ipv6Cidrs    apijson.Field
	JdcloudCidrs apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IPListResponseResultAddressingIPsJdcloud) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseResultAddressingIPsJdcloudJSON) RawJSON() string {
	return r.raw
}

func (r IPListResponseResultAddressingIPsJdcloud) implementsIPListResponseResult() {}

type IPListParams struct {
	// Specified as `jdcloud` to list IPs used by JD Cloud data centers.
	Networks param.Field[string] `query:"networks"`
}

// URLQuery serializes [IPListParams]'s query parameters as `url.Values`.
func (r IPListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
