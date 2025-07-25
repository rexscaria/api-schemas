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

type IPListResponse struct {
	Errors   []IPListResponseError   `json:"errors,required"`
	Messages []IPListResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success IPListResponseSuccess `json:"success,required"`
	Result  IPListResponseResult  `json:"result"`
	JSON    ipListResponseJSON    `json:"-"`
}

// ipListResponseJSON contains the JSON metadata for the struct [IPListResponse]
type ipListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

type IPListResponseError struct {
	Code             int64                      `json:"code,required"`
	Message          string                     `json:"message,required"`
	DocumentationURL string                     `json:"documentation_url"`
	Source           IPListResponseErrorsSource `json:"source"`
	JSON             ipListResponseErrorJSON    `json:"-"`
}

// ipListResponseErrorJSON contains the JSON metadata for the struct
// [IPListResponseError]
type ipListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IPListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type IPListResponseErrorsSource struct {
	Pointer string                         `json:"pointer"`
	JSON    ipListResponseErrorsSourceJSON `json:"-"`
}

// ipListResponseErrorsSourceJSON contains the JSON metadata for the struct
// [IPListResponseErrorsSource]
type ipListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IPListResponseMessage struct {
	Code             int64                        `json:"code,required"`
	Message          string                       `json:"message,required"`
	DocumentationURL string                       `json:"documentation_url"`
	Source           IPListResponseMessagesSource `json:"source"`
	JSON             ipListResponseMessageJSON    `json:"-"`
}

// ipListResponseMessageJSON contains the JSON metadata for the struct
// [IPListResponseMessage]
type ipListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IPListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type IPListResponseMessagesSource struct {
	Pointer string                           `json:"pointer"`
	JSON    ipListResponseMessagesSourceJSON `json:"-"`
}

// ipListResponseMessagesSourceJSON contains the JSON metadata for the struct
// [IPListResponseMessagesSource]
type ipListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IPListResponseSuccess bool

const (
	IPListResponseSuccessTrue IPListResponseSuccess = true
)

func (r IPListResponseSuccess) IsKnown() bool {
	switch r {
	case IPListResponseSuccessTrue:
		return true
	}
	return false
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
// Possible runtime types of the union are [IPListResponseResultPublicIPIPs],
// [IPListResponseResultPublicIPIPsJdcloud].
func (r IPListResponseResult) AsUnion() IPListResponseResultUnion {
	return r.union
}

// Union satisfied by [IPListResponseResultPublicIPIPs] or
// [IPListResponseResultPublicIPIPsJdcloud].
type IPListResponseResultUnion interface {
	implementsIPListResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IPListResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPListResponseResultPublicIPIPs{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPListResponseResultPublicIPIPsJdcloud{}),
		},
	)
}

type IPListResponseResultPublicIPIPs struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// List of Cloudflare IPv4 CIDR addresses.
	Ipv4Cidrs []string `json:"ipv4_cidrs"`
	// List of Cloudflare IPv6 CIDR addresses.
	Ipv6Cidrs []string                            `json:"ipv6_cidrs"`
	JSON      ipListResponseResultPublicIpiPsJSON `json:"-"`
}

// ipListResponseResultPublicIpiPsJSON contains the JSON metadata for the struct
// [IPListResponseResultPublicIPIPs]
type ipListResponseResultPublicIpiPsJSON struct {
	Etag        apijson.Field
	Ipv4Cidrs   apijson.Field
	Ipv6Cidrs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseResultPublicIPIPs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseResultPublicIpiPsJSON) RawJSON() string {
	return r.raw
}

func (r IPListResponseResultPublicIPIPs) implementsIPListResponseResult() {}

type IPListResponseResultPublicIPIPsJdcloud struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// List of Cloudflare IPv4 CIDR addresses.
	Ipv4Cidrs []string `json:"ipv4_cidrs"`
	// List of Cloudflare IPv6 CIDR addresses.
	Ipv6Cidrs []string `json:"ipv6_cidrs"`
	// List IPv4 and IPv6 CIDRs, only populated if `?networks=jdcloud` is used.
	JdcloudCidrs []string                                   `json:"jdcloud_cidrs"`
	JSON         ipListResponseResultPublicIpiPsJdcloudJSON `json:"-"`
}

// ipListResponseResultPublicIpiPsJdcloudJSON contains the JSON metadata for the
// struct [IPListResponseResultPublicIPIPsJdcloud]
type ipListResponseResultPublicIpiPsJdcloudJSON struct {
	Etag         apijson.Field
	Ipv4Cidrs    apijson.Field
	Ipv6Cidrs    apijson.Field
	JdcloudCidrs apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IPListResponseResultPublicIPIPsJdcloud) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseResultPublicIpiPsJdcloudJSON) RawJSON() string {
	return r.raw
}

func (r IPListResponseResultPublicIPIPsJdcloud) implementsIPListResponseResult() {}

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
