// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// ZoneSpectrumAppService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSpectrumAppService] method instead.
type ZoneSpectrumAppService struct {
	Options []option.RequestOption
}

// NewZoneSpectrumAppService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSpectrumAppService(opts ...option.RequestOption) (r *ZoneSpectrumAppService) {
	r = &ZoneSpectrumAppService{}
	r.Options = opts
	return
}

// Creates a new Spectrum application from a configuration using a name for the
// origin.
func (r *ZoneSpectrumAppService) New(ctx context.Context, zoneID ZoneIdentifierParam, body ZoneSpectrumAppNewParams, opts ...option.RequestOption) (res *AppConfigSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/apps", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets the application configuration of a specific application inside a zone.
func (r *ZoneSpectrumAppService) Get(ctx context.Context, zoneID ZoneIdentifierParam, appID AppIdentifierParam, opts ...option.RequestOption) (res *AppConfigSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", zoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a previously existing application's configuration that uses a name for
// the origin.
func (r *ZoneSpectrumAppService) Update(ctx context.Context, zoneID ZoneIdentifierParam, appID AppIdentifierParam, body ZoneSpectrumAppUpdateParams, opts ...option.RequestOption) (res *AppConfigSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", zoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a list of currently existing Spectrum applications inside a zone.
func (r *ZoneSpectrumAppService) List(ctx context.Context, zoneID ZoneIdentifierParam, query ZoneSpectrumAppListParams, opts ...option.RequestOption) (res *ZoneSpectrumAppListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/apps", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a previously existing application.
func (r *ZoneSpectrumAppService) Delete(ctx context.Context, zoneID ZoneIdentifierParam, appID AppIdentifierParam, opts ...option.RequestOption) (res *ZoneSpectrumAppDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", zoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AppConfig struct {
	// App identifier.
	ID AppIdentifier `json:"id,required"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS DNS `json:"dns,required"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall,required"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol,required"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol AppConfigProxyProtocol `json:"proxy_protocol,required"`
	// The type of TLS termination associated with the application.
	Tls AppConfigTls `json:"tls,required"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType AppConfigTrafficType `json:"traffic_type,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs AppConfigEdgeIPs `json:"edge_ips"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load
	// balancing.
	OriginDirect []string `json:"origin_direct" format:"URI"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS AppConfigOriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort AppConfigOriginPortUnion `json:"origin_port"`
	JSON       appConfigJSON            `json:"-"`
}

// appConfigJSON contains the JSON metadata for the struct [AppConfig]
type appConfigJSON struct {
	ID               apijson.Field
	CreatedOn        apijson.Field
	DNS              apijson.Field
	IPFirewall       apijson.Field
	ModifiedOn       apijson.Field
	Protocol         apijson.Field
	ProxyProtocol    apijson.Field
	Tls              apijson.Field
	TrafficType      apijson.Field
	ArgoSmartRouting apijson.Field
	EdgeIPs          apijson.Field
	OriginDirect     apijson.Field
	OriginDNS        apijson.Field
	OriginPort       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appConfigJSON) RawJSON() string {
	return r.raw
}

func (r AppConfig) implementsAppConfigSingleResult() {}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppConfigProxyProtocol string

const (
	AppConfigProxyProtocolOff    AppConfigProxyProtocol = "off"
	AppConfigProxyProtocolV1     AppConfigProxyProtocol = "v1"
	AppConfigProxyProtocolV2     AppConfigProxyProtocol = "v2"
	AppConfigProxyProtocolSimple AppConfigProxyProtocol = "simple"
)

func (r AppConfigProxyProtocol) IsKnown() bool {
	switch r {
	case AppConfigProxyProtocolOff, AppConfigProxyProtocolV1, AppConfigProxyProtocolV2, AppConfigProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppConfigTls string

const (
	AppConfigTlsOff      AppConfigTls = "off"
	AppConfigTlsFlexible AppConfigTls = "flexible"
	AppConfigTlsFull     AppConfigTls = "full"
	AppConfigTlsStrict   AppConfigTls = "strict"
)

func (r AppConfigTls) IsKnown() bool {
	switch r {
	case AppConfigTlsOff, AppConfigTlsFlexible, AppConfigTlsFull, AppConfigTlsStrict:
		return true
	}
	return false
}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppConfigTrafficType string

const (
	AppConfigTrafficTypeDirect AppConfigTrafficType = "direct"
	AppConfigTrafficTypeHTTP   AppConfigTrafficType = "http"
	AppConfigTrafficTypeHTTPS  AppConfigTrafficType = "https"
)

func (r AppConfigTrafficType) IsKnown() bool {
	switch r {
	case AppConfigTrafficTypeDirect, AppConfigTrafficTypeHTTP, AppConfigTrafficTypeHTTPS:
		return true
	}
	return false
}

// The anycast edge IP configuration for the hostname of this application.
type AppConfigEdgeIPs struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity AppConfigEdgeIPsConnectivity `json:"connectivity"`
	// This field can have the runtime type of [[]string].
	IPs interface{} `json:"ips"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type  AppConfigEdgeIPsType `json:"type"`
	JSON  appConfigEdgeIPsJSON `json:"-"`
	union AppConfigEdgeIPsUnion
}

// appConfigEdgeIPsJSON contains the JSON metadata for the struct
// [AppConfigEdgeIPs]
type appConfigEdgeIPsJSON struct {
	Connectivity apijson.Field
	IPs          apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r appConfigEdgeIPsJSON) RawJSON() string {
	return r.raw
}

func (r *AppConfigEdgeIPs) UnmarshalJSON(data []byte) (err error) {
	*r = AppConfigEdgeIPs{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AppConfigEdgeIPsUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [AppConfigEdgeIPsDynamic],
// [AppConfigEdgeIPsStatic].
func (r AppConfigEdgeIPs) AsUnion() AppConfigEdgeIPsUnion {
	return r.union
}

// The anycast edge IP configuration for the hostname of this application.
//
// Union satisfied by [AppConfigEdgeIPsDynamic] or [AppConfigEdgeIPsStatic].
type AppConfigEdgeIPsUnion interface {
	implementsAppConfigEdgeIPs()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppConfigEdgeIPsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppConfigEdgeIPsDynamic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppConfigEdgeIPsStatic{}),
		},
	)
}

type AppConfigEdgeIPsDynamic struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity AppConfigEdgeIPsDynamicConnectivity `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type AppConfigEdgeIPsDynamicType `json:"type"`
	JSON appConfigEdgeIPsDynamicJSON `json:"-"`
}

// appConfigEdgeIPsDynamicJSON contains the JSON metadata for the struct
// [AppConfigEdgeIPsDynamic]
type appConfigEdgeIPsDynamicJSON struct {
	Connectivity apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AppConfigEdgeIPsDynamic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appConfigEdgeIPsDynamicJSON) RawJSON() string {
	return r.raw
}

func (r AppConfigEdgeIPsDynamic) implementsAppConfigEdgeIPs() {}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type AppConfigEdgeIPsDynamicConnectivity string

const (
	AppConfigEdgeIPsDynamicConnectivityAll  AppConfigEdgeIPsDynamicConnectivity = "all"
	AppConfigEdgeIPsDynamicConnectivityIpv4 AppConfigEdgeIPsDynamicConnectivity = "ipv4"
	AppConfigEdgeIPsDynamicConnectivityIpv6 AppConfigEdgeIPsDynamicConnectivity = "ipv6"
)

func (r AppConfigEdgeIPsDynamicConnectivity) IsKnown() bool {
	switch r {
	case AppConfigEdgeIPsDynamicConnectivityAll, AppConfigEdgeIPsDynamicConnectivityIpv4, AppConfigEdgeIPsDynamicConnectivityIpv6:
		return true
	}
	return false
}

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type AppConfigEdgeIPsDynamicType string

const (
	AppConfigEdgeIPsDynamicTypeDynamic AppConfigEdgeIPsDynamicType = "dynamic"
)

func (r AppConfigEdgeIPsDynamicType) IsKnown() bool {
	switch r {
	case AppConfigEdgeIPsDynamicTypeDynamic:
		return true
	}
	return false
}

type AppConfigEdgeIPsStatic struct {
	// The array of customer owned IPs we broadcast via anycast for this hostname and
	// application.
	IPs []string `json:"ips"`
	// The type of edge IP configuration specified. Statically allocated edge IPs use
	// customer IPs in accordance with the ips array you specify. Only valid with
	// ADDRESS DNS names.
	Type AppConfigEdgeIPsStaticType `json:"type"`
	JSON appConfigEdgeIPsStaticJSON `json:"-"`
}

// appConfigEdgeIPsStaticJSON contains the JSON metadata for the struct
// [AppConfigEdgeIPsStatic]
type appConfigEdgeIPsStaticJSON struct {
	IPs         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppConfigEdgeIPsStatic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appConfigEdgeIPsStaticJSON) RawJSON() string {
	return r.raw
}

func (r AppConfigEdgeIPsStatic) implementsAppConfigEdgeIPs() {}

// The type of edge IP configuration specified. Statically allocated edge IPs use
// customer IPs in accordance with the ips array you specify. Only valid with
// ADDRESS DNS names.
type AppConfigEdgeIPsStaticType string

const (
	AppConfigEdgeIPsStaticTypeStatic AppConfigEdgeIPsStaticType = "static"
)

func (r AppConfigEdgeIPsStaticType) IsKnown() bool {
	switch r {
	case AppConfigEdgeIPsStaticTypeStatic:
		return true
	}
	return false
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type AppConfigEdgeIPsConnectivity string

const (
	AppConfigEdgeIPsConnectivityAll  AppConfigEdgeIPsConnectivity = "all"
	AppConfigEdgeIPsConnectivityIpv4 AppConfigEdgeIPsConnectivity = "ipv4"
	AppConfigEdgeIPsConnectivityIpv6 AppConfigEdgeIPsConnectivity = "ipv6"
)

func (r AppConfigEdgeIPsConnectivity) IsKnown() bool {
	switch r {
	case AppConfigEdgeIPsConnectivityAll, AppConfigEdgeIPsConnectivityIpv4, AppConfigEdgeIPsConnectivityIpv6:
		return true
	}
	return false
}

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type AppConfigEdgeIPsType string

const (
	AppConfigEdgeIPsTypeDynamic AppConfigEdgeIPsType = "dynamic"
	AppConfigEdgeIPsTypeStatic  AppConfigEdgeIPsType = "static"
)

func (r AppConfigEdgeIPsType) IsKnown() bool {
	switch r {
	case AppConfigEdgeIPsTypeDynamic, AppConfigEdgeIPsTypeStatic:
		return true
	}
	return false
}

// The name and type of DNS record for the Spectrum application.
type AppConfigOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name string `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	Ttl int64 `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type AppConfigOriginDNSType `json:"type"`
	JSON appConfigOriginDNSJSON `json:"-"`
}

// appConfigOriginDNSJSON contains the JSON metadata for the struct
// [AppConfigOriginDNS]
type appConfigOriginDNSJSON struct {
	Name        apijson.Field
	Ttl         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppConfigOriginDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appConfigOriginDNSJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type AppConfigOriginDNSType string

const (
	AppConfigOriginDNSTypeEmpty AppConfigOriginDNSType = ""
	AppConfigOriginDNSTypeA     AppConfigOriginDNSType = "A"
	AppConfigOriginDNSTypeAaaa  AppConfigOriginDNSType = "AAAA"
	AppConfigOriginDNSTypeSrv   AppConfigOriginDNSType = "SRV"
)

func (r AppConfigOriginDNSType) IsKnown() bool {
	switch r {
	case AppConfigOriginDNSTypeEmpty, AppConfigOriginDNSTypeA, AppConfigOriginDNSTypeAaaa, AppConfigOriginDNSTypeSrv:
		return true
	}
	return false
}

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Union satisfied by [shared.UnionInt] or [shared.UnionString].
type AppConfigOriginPortUnion interface {
	ImplementsAppConfigOriginPortUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppConfigOriginPortUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AppConfigParam struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[DNSParam] `json:"dns,required"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall param.Field[bool] `json:"ip_firewall,required"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol param.Field[AppConfigProxyProtocol] `json:"proxy_protocol,required"`
	// The type of TLS termination associated with the application.
	Tls param.Field[AppConfigTls] `json:"tls,required"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType param.Field[AppConfigTrafficType] `json:"traffic_type,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting param.Field[bool] `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs param.Field[AppConfigEdgeIPsUnionParam] `json:"edge_ips"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load
	// balancing.
	OriginDirect param.Field[[]string] `json:"origin_direct" format:"URI"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS param.Field[AppConfigOriginDNSParam] `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort param.Field[AppConfigOriginPortUnionParam] `json:"origin_port"`
}

func (r AppConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppConfigParam) implementsUpdateAppConfigUnionParam() {}

// The anycast edge IP configuration for the hostname of this application.
type AppConfigEdgeIPsParam struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[AppConfigEdgeIPsConnectivity] `json:"connectivity"`
	IPs          param.Field[interface{}]                  `json:"ips"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[AppConfigEdgeIPsType] `json:"type"`
}

func (r AppConfigEdgeIPsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppConfigEdgeIPsParam) implementsAppConfigEdgeIPsUnionParam() {}

// The anycast edge IP configuration for the hostname of this application.
//
// Satisfied by [AppConfigEdgeIPsDynamicParam], [AppConfigEdgeIPsStaticParam],
// [AppConfigEdgeIPsParam].
type AppConfigEdgeIPsUnionParam interface {
	implementsAppConfigEdgeIPsUnionParam()
}

type AppConfigEdgeIPsDynamicParam struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[AppConfigEdgeIPsDynamicConnectivity] `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[AppConfigEdgeIPsDynamicType] `json:"type"`
}

func (r AppConfigEdgeIPsDynamicParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppConfigEdgeIPsDynamicParam) implementsAppConfigEdgeIPsUnionParam() {}

type AppConfigEdgeIPsStaticParam struct {
	// The array of customer owned IPs we broadcast via anycast for this hostname and
	// application.
	IPs param.Field[[]string] `json:"ips"`
	// The type of edge IP configuration specified. Statically allocated edge IPs use
	// customer IPs in accordance with the ips array you specify. Only valid with
	// ADDRESS DNS names.
	Type param.Field[AppConfigEdgeIPsStaticType] `json:"type"`
}

func (r AppConfigEdgeIPsStaticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppConfigEdgeIPsStaticParam) implementsAppConfigEdgeIPsUnionParam() {}

// The name and type of DNS record for the Spectrum application.
type AppConfigOriginDNSParam struct {
	// The name of the DNS record associated with the origin.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	Ttl param.Field[int64] `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type param.Field[AppConfigOriginDNSType] `json:"type"`
}

func (r AppConfigOriginDNSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type AppConfigOriginPortUnionParam interface {
	ImplementsAppConfigOriginPortUnionParam()
}

type AppConfigSingle struct {
	Errors   []SpectrumConfigMessageItem `json:"errors,required"`
	Messages []SpectrumConfigMessageItem `json:"messages,required"`
	// Whether the API call was successful.
	Success AppConfigSingleSuccess `json:"success,required"`
	Result  AppConfigSingleResult  `json:"result"`
	JSON    appConfigSingleJSON    `json:"-"`
}

// appConfigSingleJSON contains the JSON metadata for the struct [AppConfigSingle]
type appConfigSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppConfigSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appConfigSingleJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AppConfigSingleSuccess bool

const (
	AppConfigSingleSuccessTrue AppConfigSingleSuccess = true
)

func (r AppConfigSingleSuccess) IsKnown() bool {
	switch r {
	case AppConfigSingleSuccessTrue:
		return true
	}
	return false
}

type AppConfigSingleResult struct {
	// App identifier.
	ID shared.UnionString `json:"id,required"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS DNS `json:"dns,required"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// This field can have the runtime type of [AppConfigEdgeIPs].
	EdgeIPs interface{} `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// This field can have the runtime type of [[]string].
	OriginDirect interface{} `json:"origin_direct"`
	// This field can have the runtime type of [AppConfigOriginDNS].
	OriginDNS interface{} `json:"origin_dns"`
	// This field can have the runtime type of [AppConfigOriginPortUnion].
	OriginPort interface{} `json:"origin_port"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol AppConfigSingleResultProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	Tls AppConfigSingleResultTls `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType AppConfigSingleResultTrafficType `json:"traffic_type"`
	JSON        appConfigSingleResultJSON        `json:"-"`
	union       AppConfigSingleResultUnion
}

// appConfigSingleResultJSON contains the JSON metadata for the struct
// [AppConfigSingleResult]
type appConfigSingleResultJSON struct {
	ID               apijson.Field
	CreatedOn        apijson.Field
	DNS              apijson.Field
	ModifiedOn       apijson.Field
	Protocol         apijson.Field
	ArgoSmartRouting apijson.Field
	EdgeIPs          apijson.Field
	IPFirewall       apijson.Field
	OriginDirect     apijson.Field
	OriginDNS        apijson.Field
	OriginPort       apijson.Field
	ProxyProtocol    apijson.Field
	Tls              apijson.Field
	TrafficType      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r appConfigSingleResultJSON) RawJSON() string {
	return r.raw
}

func (r *AppConfigSingleResult) UnmarshalJSON(data []byte) (err error) {
	*r = AppConfigSingleResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AppConfigSingleResultUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [AppConfig], [PaygoAppConfig].
func (r AppConfigSingleResult) AsUnion() AppConfigSingleResultUnion {
	return r.union
}

// Union satisfied by [AppConfig] or [PaygoAppConfig].
type AppConfigSingleResultUnion interface {
	implementsAppConfigSingleResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppConfigSingleResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PaygoAppConfig{}),
		},
	)
}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppConfigSingleResultProxyProtocol string

const (
	AppConfigSingleResultProxyProtocolOff    AppConfigSingleResultProxyProtocol = "off"
	AppConfigSingleResultProxyProtocolV1     AppConfigSingleResultProxyProtocol = "v1"
	AppConfigSingleResultProxyProtocolV2     AppConfigSingleResultProxyProtocol = "v2"
	AppConfigSingleResultProxyProtocolSimple AppConfigSingleResultProxyProtocol = "simple"
)

func (r AppConfigSingleResultProxyProtocol) IsKnown() bool {
	switch r {
	case AppConfigSingleResultProxyProtocolOff, AppConfigSingleResultProxyProtocolV1, AppConfigSingleResultProxyProtocolV2, AppConfigSingleResultProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppConfigSingleResultTls string

const (
	AppConfigSingleResultTlsOff      AppConfigSingleResultTls = "off"
	AppConfigSingleResultTlsFlexible AppConfigSingleResultTls = "flexible"
	AppConfigSingleResultTlsFull     AppConfigSingleResultTls = "full"
	AppConfigSingleResultTlsStrict   AppConfigSingleResultTls = "strict"
)

func (r AppConfigSingleResultTls) IsKnown() bool {
	switch r {
	case AppConfigSingleResultTlsOff, AppConfigSingleResultTlsFlexible, AppConfigSingleResultTlsFull, AppConfigSingleResultTlsStrict:
		return true
	}
	return false
}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppConfigSingleResultTrafficType string

const (
	AppConfigSingleResultTrafficTypeDirect AppConfigSingleResultTrafficType = "direct"
	AppConfigSingleResultTrafficTypeHTTP   AppConfigSingleResultTrafficType = "http"
	AppConfigSingleResultTrafficTypeHTTPS  AppConfigSingleResultTrafficType = "https"
)

func (r AppConfigSingleResultTrafficType) IsKnown() bool {
	switch r {
	case AppConfigSingleResultTrafficTypeDirect, AppConfigSingleResultTrafficTypeHTTP, AppConfigSingleResultTrafficTypeHTTPS:
		return true
	}
	return false
}

type AppIdentifier = string

type AppIdentifierParam = string

// The name and type of DNS record for the Spectrum application.
type DNS struct {
	// The name of the DNS record associated with the application.
	Name string `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type DNSType `json:"type"`
	JSON dnsJSON `json:"-"`
}

// dnsJSON contains the JSON metadata for the struct [DNS]
type dnsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the application.
type DNSType string

const (
	DNSTypeCname   DNSType = "CNAME"
	DNSTypeAddress DNSType = "ADDRESS"
)

func (r DNSType) IsKnown() bool {
	switch r {
	case DNSTypeCname, DNSTypeAddress:
		return true
	}
	return false
}

// The name and type of DNS record for the Spectrum application.
type DNSParam struct {
	// The name of the DNS record associated with the application.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type param.Field[DNSType] `json:"type"`
}

func (r DNSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaygoAppConfig struct {
	// App identifier.
	ID AppIdentifier `json:"id,required"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS DNS `json:"dns,required"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol,required"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load
	// balancing.
	OriginDirect []string           `json:"origin_direct" format:"URI"`
	JSON         paygoAppConfigJSON `json:"-"`
}

// paygoAppConfigJSON contains the JSON metadata for the struct [PaygoAppConfig]
type paygoAppConfigJSON struct {
	ID           apijson.Field
	CreatedOn    apijson.Field
	DNS          apijson.Field
	ModifiedOn   apijson.Field
	Protocol     apijson.Field
	OriginDirect apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PaygoAppConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paygoAppConfigJSON) RawJSON() string {
	return r.raw
}

func (r PaygoAppConfig) implementsAppConfigSingleResult() {}

type PaygoAppConfigParam struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[DNSParam] `json:"dns,required"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load
	// balancing.
	OriginDirect param.Field[[]string] `json:"origin_direct" format:"URI"`
}

func (r PaygoAppConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PaygoAppConfigParam) implementsUpdateAppConfigUnionParam() {}

type SpectrumConfigMessageItem struct {
	Code             int64                           `json:"code,required"`
	Message          string                          `json:"message,required"`
	DocumentationURL string                          `json:"documentation_url"`
	Source           SpectrumConfigMessageItemSource `json:"source"`
	JSON             spectrumConfigMessageItemJSON   `json:"-"`
}

// spectrumConfigMessageItemJSON contains the JSON metadata for the struct
// [SpectrumConfigMessageItem]
type spectrumConfigMessageItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SpectrumConfigMessageItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumConfigMessageItemJSON) RawJSON() string {
	return r.raw
}

type SpectrumConfigMessageItemSource struct {
	Pointer string                              `json:"pointer"`
	JSON    spectrumConfigMessageItemSourceJSON `json:"-"`
}

// spectrumConfigMessageItemSourceJSON contains the JSON metadata for the struct
// [SpectrumConfigMessageItemSource]
type spectrumConfigMessageItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumConfigMessageItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumConfigMessageItemSourceJSON) RawJSON() string {
	return r.raw
}

type UpdateAppConfigParam struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[DNSParam] `json:"dns,required"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting param.Field[bool]        `json:"argo_smart_routing"`
	EdgeIPs          param.Field[interface{}] `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall   param.Field[bool]        `json:"ip_firewall"`
	OriginDirect param.Field[interface{}] `json:"origin_direct"`
	OriginDNS    param.Field[interface{}] `json:"origin_dns"`
	OriginPort   param.Field[interface{}] `json:"origin_port"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol param.Field[UpdateAppConfigProxyProtocol] `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	Tls param.Field[UpdateAppConfigTls] `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType param.Field[UpdateAppConfigTrafficType] `json:"traffic_type"`
}

func (r UpdateAppConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UpdateAppConfigParam) implementsUpdateAppConfigUnionParam() {}

// Satisfied by [AppConfigParam], [PaygoAppConfigParam], [UpdateAppConfigParam].
type UpdateAppConfigUnionParam interface {
	implementsUpdateAppConfigUnionParam()
}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type UpdateAppConfigProxyProtocol string

const (
	UpdateAppConfigProxyProtocolOff    UpdateAppConfigProxyProtocol = "off"
	UpdateAppConfigProxyProtocolV1     UpdateAppConfigProxyProtocol = "v1"
	UpdateAppConfigProxyProtocolV2     UpdateAppConfigProxyProtocol = "v2"
	UpdateAppConfigProxyProtocolSimple UpdateAppConfigProxyProtocol = "simple"
)

func (r UpdateAppConfigProxyProtocol) IsKnown() bool {
	switch r {
	case UpdateAppConfigProxyProtocolOff, UpdateAppConfigProxyProtocolV1, UpdateAppConfigProxyProtocolV2, UpdateAppConfigProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type UpdateAppConfigTls string

const (
	UpdateAppConfigTlsOff      UpdateAppConfigTls = "off"
	UpdateAppConfigTlsFlexible UpdateAppConfigTls = "flexible"
	UpdateAppConfigTlsFull     UpdateAppConfigTls = "full"
	UpdateAppConfigTlsStrict   UpdateAppConfigTls = "strict"
)

func (r UpdateAppConfigTls) IsKnown() bool {
	switch r {
	case UpdateAppConfigTlsOff, UpdateAppConfigTlsFlexible, UpdateAppConfigTlsFull, UpdateAppConfigTlsStrict:
		return true
	}
	return false
}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type UpdateAppConfigTrafficType string

const (
	UpdateAppConfigTrafficTypeDirect UpdateAppConfigTrafficType = "direct"
	UpdateAppConfigTrafficTypeHTTP   UpdateAppConfigTrafficType = "http"
	UpdateAppConfigTrafficTypeHTTPS  UpdateAppConfigTrafficType = "https"
)

func (r UpdateAppConfigTrafficType) IsKnown() bool {
	switch r {
	case UpdateAppConfigTrafficTypeDirect, UpdateAppConfigTrafficTypeHTTP, UpdateAppConfigTrafficTypeHTTPS:
		return true
	}
	return false
}

type ZoneIdentifierParam = string

type ZoneSpectrumAppListResponse struct {
	Errors   []SpectrumConfigMessageItem `json:"errors,required"`
	Messages []SpectrumConfigMessageItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    ZoneSpectrumAppListResponseSuccess     `json:"success,required"`
	Result     ZoneSpectrumAppListResponseResultUnion `json:"result"`
	ResultInfo ZoneSpectrumAppListResponseResultInfo  `json:"result_info"`
	JSON       zoneSpectrumAppListResponseJSON        `json:"-"`
}

// zoneSpectrumAppListResponseJSON contains the JSON metadata for the struct
// [ZoneSpectrumAppListResponse]
type zoneSpectrumAppListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpectrumAppListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpectrumAppListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneSpectrumAppListResponseSuccess bool

const (
	ZoneSpectrumAppListResponseSuccessTrue ZoneSpectrumAppListResponseSuccess = true
)

func (r ZoneSpectrumAppListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSpectrumAppListResponseSuccessTrue:
		return true
	}
	return false
}

// Union satisfied by [ZoneSpectrumAppListResponseResultArray] or
// [ZoneSpectrumAppListResponseResultArray].
type ZoneSpectrumAppListResponseResultUnion interface {
	implementsZoneSpectrumAppListResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneSpectrumAppListResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSpectrumAppListResponseResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSpectrumAppListResponseResultArray{}),
		},
	)
}

type ZoneSpectrumAppListResponseResultArray []AppConfig

func (r ZoneSpectrumAppListResponseResultArray) implementsZoneSpectrumAppListResponseResultUnion() {}

type ZoneSpectrumAppListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                   `json:"total_count"`
	JSON       zoneSpectrumAppListResponseResultInfoJSON `json:"-"`
}

// zoneSpectrumAppListResponseResultInfoJSON contains the JSON metadata for the
// struct [ZoneSpectrumAppListResponseResultInfo]
type zoneSpectrumAppListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpectrumAppListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpectrumAppListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneSpectrumAppDeleteResponse struct {
	Errors   []SpectrumConfigMessageItem `json:"errors,required"`
	Messages []SpectrumConfigMessageItem `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneSpectrumAppDeleteResponseSuccess `json:"success,required"`
	Result  ZoneSpectrumAppDeleteResponseResult  `json:"result,nullable"`
	JSON    zoneSpectrumAppDeleteResponseJSON    `json:"-"`
}

// zoneSpectrumAppDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneSpectrumAppDeleteResponse]
type zoneSpectrumAppDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpectrumAppDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpectrumAppDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneSpectrumAppDeleteResponseSuccess bool

const (
	ZoneSpectrumAppDeleteResponseSuccessTrue ZoneSpectrumAppDeleteResponseSuccess = true
)

func (r ZoneSpectrumAppDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSpectrumAppDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneSpectrumAppDeleteResponseResult struct {
	// Identifier.
	ID   string                                  `json:"id,required"`
	JSON zoneSpectrumAppDeleteResponseResultJSON `json:"-"`
}

// zoneSpectrumAppDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneSpectrumAppDeleteResponseResult]
type zoneSpectrumAppDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpectrumAppDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpectrumAppDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSpectrumAppNewParams struct {
	UpdateAppConfig UpdateAppConfigUnionParam `json:"update_app_config,required"`
}

func (r ZoneSpectrumAppNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdateAppConfig)
}

type ZoneSpectrumAppUpdateParams struct {
	UpdateAppConfig UpdateAppConfigUnionParam `json:"update_app_config,required"`
}

func (r ZoneSpectrumAppUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdateAppConfig)
}

type ZoneSpectrumAppListParams struct {
	// Sets the direction by which results are ordered.
	Direction param.Field[ZoneSpectrumAppListParamsDirection] `query:"direction"`
	// Application field by which results are ordered.
	Order param.Field[ZoneSpectrumAppListParamsOrder] `query:"order"`
	// Page number of paginated results. This parameter is required in order to use
	// other pagination parameters. If included in the query, `result_info` will be
	// present in the response.
	Page param.Field[float64] `query:"page"`
	// Sets the maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ZoneSpectrumAppListParams]'s query parameters as
// `url.Values`.
func (r ZoneSpectrumAppListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sets the direction by which results are ordered.
type ZoneSpectrumAppListParamsDirection string

const (
	ZoneSpectrumAppListParamsDirectionAsc  ZoneSpectrumAppListParamsDirection = "asc"
	ZoneSpectrumAppListParamsDirectionDesc ZoneSpectrumAppListParamsDirection = "desc"
)

func (r ZoneSpectrumAppListParamsDirection) IsKnown() bool {
	switch r {
	case ZoneSpectrumAppListParamsDirectionAsc, ZoneSpectrumAppListParamsDirectionDesc:
		return true
	}
	return false
}

// Application field by which results are ordered.
type ZoneSpectrumAppListParamsOrder string

const (
	ZoneSpectrumAppListParamsOrderProtocol   ZoneSpectrumAppListParamsOrder = "protocol"
	ZoneSpectrumAppListParamsOrderAppID      ZoneSpectrumAppListParamsOrder = "app_id"
	ZoneSpectrumAppListParamsOrderCreatedOn  ZoneSpectrumAppListParamsOrder = "created_on"
	ZoneSpectrumAppListParamsOrderModifiedOn ZoneSpectrumAppListParamsOrder = "modified_on"
	ZoneSpectrumAppListParamsOrderDNS        ZoneSpectrumAppListParamsOrder = "dns"
)

func (r ZoneSpectrumAppListParamsOrder) IsKnown() bool {
	switch r {
	case ZoneSpectrumAppListParamsOrderProtocol, ZoneSpectrumAppListParamsOrderAppID, ZoneSpectrumAppListParamsOrderCreatedOn, ZoneSpectrumAppListParamsOrderModifiedOn, ZoneSpectrumAppListParamsOrderDNS:
		return true
	}
	return false
}
