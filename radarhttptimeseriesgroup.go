// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// RadarHTTPTimeseriesGroupService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarHTTPTimeseriesGroupService] method instead.
type RadarHTTPTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTimeseriesGroupService(opts ...option.RequestOption) (r *RadarHTTPTimeseriesGroupService) {
	r = &RadarHTTPTimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of HTTP requests classified as automated or human
// over time. Visit https://developers.cloudflare.com/radar/concepts/bot-classes/
// for more information.
func (r *RadarHTTPTimeseriesGroupService) GetByBotClass(ctx context.Context, query RadarHTTPTimeseriesGroupGetByBotClassParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupGetByBotClassResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/bot_class"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of HTTP requests by user agent over time.
func (r *RadarHTTPTimeseriesGroupService) GetByBrowser(ctx context.Context, query RadarHTTPTimeseriesGroupGetByBrowserParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupGetByBrowserResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/browser"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of HTTP requests by user agent family over time.
func (r *RadarHTTPTimeseriesGroupService) GetByBrowserFamily(ctx context.Context, query RadarHTTPTimeseriesGroupGetByBrowserFamilyParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupGetByBrowserFamilyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/browser_family"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of HTTP requests by device type over time.
func (r *RadarHTTPTimeseriesGroupService) GetByDeviceType(ctx context.Context, query RadarHTTPTimeseriesGroupGetByDeviceTypeParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupGetByDeviceTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/device_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of HTTP requests by HTTP protocol (HTTP vs. HTTPS)
// over time.
func (r *RadarHTTPTimeseriesGroupService) GetByHTTPProtocol(ctx context.Context, query RadarHTTPTimeseriesGroupGetByHTTPProtocolParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupGetByHTTPProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/http_protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of HTTP requests by HTTP version over time.
func (r *RadarHTTPTimeseriesGroupService) GetByHTTPVersion(ctx context.Context, query RadarHTTPTimeseriesGroupGetByHTTPVersionParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupGetByHTTPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of HTTP requests by IP version over time.
func (r *RadarHTTPTimeseriesGroupService) GetByIPVersion(ctx context.Context, query RadarHTTPTimeseriesGroupGetByIPVersionParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupGetByIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of HTTP requests by operating system over time.
func (r *RadarHTTPTimeseriesGroupService) GetByOs(ctx context.Context, query RadarHTTPTimeseriesGroupGetByOsParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupGetByOsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/os"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of HTTP requests by post-quantum support over time.
func (r *RadarHTTPTimeseriesGroupService) GetByPostQuantum(ctx context.Context, query RadarHTTPTimeseriesGroupGetByPostQuantumParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupGetByPostQuantumResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/post_quantum"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of HTTP requests by TLS version over time.
func (r *RadarHTTPTimeseriesGroupService) GetByTlsVersion(ctx context.Context, query RadarHTTPTimeseriesGroupGetByTlsVersionParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupGetByTlsVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/tls_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseriesGroupGetByBotClassResponse struct {
	Result  RadarHTTPTimeseriesGroupGetByBotClassResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupGetByBotClassResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByBotClassResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseriesGroupGetByBotClassResponse]
type radarHTTPTimeseriesGroupGetByBotClassResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByBotClassResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByBotClassResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByBotClassResponseResult struct {
	Meta   interface{}                                               `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupGetByBotClassResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupGetByBotClassResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByBotClassResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupGetByBotClassResponseResult]
type radarHTTPTimeseriesGroupGetByBotClassResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByBotClassResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByBotClassResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByBotClassResponseResultSerie0 struct {
	Bot        []string                                                      `json:"bot,required"`
	Human      []string                                                      `json:"human,required"`
	Timestamps []string                                                      `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupGetByBotClassResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupGetByBotClassResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarHTTPTimeseriesGroupGetByBotClassResponseResultSerie0]
type radarHTTPTimeseriesGroupGetByBotClassResponseResultSerie0JSON struct {
	Bot         apijson.Field
	Human       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByBotClassResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByBotClassResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByBrowserResponse struct {
	Result  RadarHTTPTimeseriesGroupGetByBrowserResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupGetByBrowserResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByBrowserResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseriesGroupGetByBrowserResponse]
type radarHTTPTimeseriesGroupGetByBrowserResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByBrowserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByBrowserResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByBrowserResponseResult struct {
	Meta   interface{}                                              `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupGetByBrowserResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupGetByBrowserResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByBrowserResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupGetByBrowserResponseResult]
type radarHTTPTimeseriesGroupGetByBrowserResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByBrowserResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByBrowserResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByBrowserResponseResultSerie0 struct {
	Timestamps  []string                                                     `json:"timestamps,required"`
	ExtraFields map[string][]string                                          `json:"-,extras"`
	JSON        radarHTTPTimeseriesGroupGetByBrowserResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupGetByBrowserResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarHTTPTimeseriesGroupGetByBrowserResponseResultSerie0]
type radarHTTPTimeseriesGroupGetByBrowserResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByBrowserResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByBrowserResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByBrowserFamilyResponse struct {
	Result  RadarHTTPTimeseriesGroupGetByBrowserFamilyResponseResult `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupGetByBrowserFamilyResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByBrowserFamilyResponseJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupGetByBrowserFamilyResponse]
type radarHTTPTimeseriesGroupGetByBrowserFamilyResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByBrowserFamilyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByBrowserFamilyResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByBrowserFamilyResponseResult struct {
	Meta   interface{}                                                    `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupGetByBrowserFamilyResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupGetByBrowserFamilyResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByBrowserFamilyResponseResultJSON contains the JSON
// metadata for the struct
// [RadarHTTPTimeseriesGroupGetByBrowserFamilyResponseResult]
type radarHTTPTimeseriesGroupGetByBrowserFamilyResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByBrowserFamilyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByBrowserFamilyResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByBrowserFamilyResponseResultSerie0 struct {
	Timestamps  []string                                                           `json:"timestamps,required"`
	ExtraFields map[string][]string                                                `json:"-,extras"`
	JSON        radarHTTPTimeseriesGroupGetByBrowserFamilyResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupGetByBrowserFamilyResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarHTTPTimeseriesGroupGetByBrowserFamilyResponseResultSerie0]
type radarHTTPTimeseriesGroupGetByBrowserFamilyResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByBrowserFamilyResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByBrowserFamilyResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByDeviceTypeResponse struct {
	Result  RadarHTTPTimeseriesGroupGetByDeviceTypeResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupGetByDeviceTypeResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByDeviceTypeResponseJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupGetByDeviceTypeResponse]
type radarHTTPTimeseriesGroupGetByDeviceTypeResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByDeviceTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByDeviceTypeResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByDeviceTypeResponseResult struct {
	Meta   interface{}                                                 `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupGetByDeviceTypeResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupGetByDeviceTypeResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByDeviceTypeResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupGetByDeviceTypeResponseResult]
type radarHTTPTimeseriesGroupGetByDeviceTypeResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByDeviceTypeResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByDeviceTypeResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByDeviceTypeResponseResultSerie0 struct {
	Desktop    []string                                                        `json:"desktop,required"`
	Mobile     []string                                                        `json:"mobile,required"`
	Other      []string                                                        `json:"other,required"`
	Timestamps []string                                                        `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupGetByDeviceTypeResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupGetByDeviceTypeResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarHTTPTimeseriesGroupGetByDeviceTypeResponseResultSerie0]
type radarHTTPTimeseriesGroupGetByDeviceTypeResponseResultSerie0JSON struct {
	Desktop     apijson.Field
	Mobile      apijson.Field
	Other       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByDeviceTypeResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByDeviceTypeResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByHTTPProtocolResponse struct {
	Result  RadarHTTPTimeseriesGroupGetByHTTPProtocolResponseResult `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupGetByHTTPProtocolResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByHTTPProtocolResponseJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupGetByHTTPProtocolResponse]
type radarHTTPTimeseriesGroupGetByHTTPProtocolResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByHTTPProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByHTTPProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByHTTPProtocolResponseResult struct {
	Meta   interface{}                                                   `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupGetByHTTPProtocolResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupGetByHTTPProtocolResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByHTTPProtocolResponseResultJSON contains the JSON
// metadata for the struct
// [RadarHTTPTimeseriesGroupGetByHTTPProtocolResponseResult]
type radarHTTPTimeseriesGroupGetByHTTPProtocolResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByHTTPProtocolResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByHTTPProtocolResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByHTTPProtocolResponseResultSerie0 struct {
	HTTP       []string                                                          `json:"http,required"`
	HTTPS      []string                                                          `json:"https,required"`
	Timestamps []string                                                          `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupGetByHTTPProtocolResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupGetByHTTPProtocolResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarHTTPTimeseriesGroupGetByHTTPProtocolResponseResultSerie0]
type radarHTTPTimeseriesGroupGetByHTTPProtocolResponseResultSerie0JSON struct {
	HTTP        apijson.Field
	HTTPS       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByHTTPProtocolResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByHTTPProtocolResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByHTTPVersionResponse struct {
	Result  RadarHTTPTimeseriesGroupGetByHTTPVersionResponseResult `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupGetByHTTPVersionResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByHTTPVersionResponseJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupGetByHTTPVersionResponse]
type radarHTTPTimeseriesGroupGetByHTTPVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByHTTPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByHTTPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByHTTPVersionResponseResult struct {
	Meta   interface{}                                                  `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupGetByHTTPVersionResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupGetByHTTPVersionResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByHTTPVersionResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupGetByHTTPVersionResponseResult]
type radarHTTPTimeseriesGroupGetByHTTPVersionResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByHTTPVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByHTTPVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByHTTPVersionResponseResultSerie0 struct {
	HTTP1X     []string                                                         `json:"HTTP/1.x,required"`
	HTTP2      []string                                                         `json:"HTTP/2,required"`
	HTTP3      []string                                                         `json:"HTTP/3,required"`
	Timestamps []string                                                         `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupGetByHTTPVersionResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupGetByHTTPVersionResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarHTTPTimeseriesGroupGetByHTTPVersionResponseResultSerie0]
type radarHTTPTimeseriesGroupGetByHTTPVersionResponseResultSerie0JSON struct {
	HTTP1X      apijson.Field
	HTTP2       apijson.Field
	HTTP3       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByHTTPVersionResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByHTTPVersionResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByIPVersionResponse struct {
	Result  RadarHTTPTimeseriesGroupGetByIPVersionResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupGetByIPVersionResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByIPVersionResponseJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupGetByIPVersionResponse]
type radarHTTPTimeseriesGroupGetByIPVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByIPVersionResponseResult struct {
	Meta   interface{}                                                `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupGetByIPVersionResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupGetByIPVersionResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByIPVersionResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupGetByIPVersionResponseResult]
type radarHTTPTimeseriesGroupGetByIPVersionResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByIPVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByIPVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByIPVersionResponseResultSerie0 struct {
	IPv4       []string                                                       `json:"IPv4,required"`
	IPv6       []string                                                       `json:"IPv6,required"`
	Timestamps []string                                                       `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupGetByIPVersionResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupGetByIPVersionResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarHTTPTimeseriesGroupGetByIPVersionResponseResultSerie0]
type radarHTTPTimeseriesGroupGetByIPVersionResponseResultSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByIPVersionResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByIPVersionResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByOsResponse struct {
	Result  RadarHTTPTimeseriesGroupGetByOsResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupGetByOsResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByOsResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTimeseriesGroupGetByOsResponse]
type radarHTTPTimeseriesGroupGetByOsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByOsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByOsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByOsResponseResult struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupGetByOsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupGetByOsResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByOsResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseriesGroupGetByOsResponseResult]
type radarHTTPTimeseriesGroupGetByOsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByOsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByOsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByOsResponseResultSerie0 struct {
	Timestamps  []string                                                `json:"timestamps,required"`
	ExtraFields map[string][]string                                     `json:"-,extras"`
	JSON        radarHTTPTimeseriesGroupGetByOsResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupGetByOsResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupGetByOsResponseResultSerie0]
type radarHTTPTimeseriesGroupGetByOsResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByOsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByOsResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByPostQuantumResponse struct {
	Result  RadarHTTPTimeseriesGroupGetByPostQuantumResponseResult `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupGetByPostQuantumResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByPostQuantumResponseJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupGetByPostQuantumResponse]
type radarHTTPTimeseriesGroupGetByPostQuantumResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByPostQuantumResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByPostQuantumResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByPostQuantumResponseResult struct {
	Meta   interface{}                                                  `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupGetByPostQuantumResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupGetByPostQuantumResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByPostQuantumResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupGetByPostQuantumResponseResult]
type radarHTTPTimeseriesGroupGetByPostQuantumResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByPostQuantumResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByPostQuantumResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByPostQuantumResponseResultSerie0 struct {
	NotSupported []string                                                         `json:"NOT_SUPPORTED,required"`
	Supported    []string                                                         `json:"SUPPORTED,required"`
	Timestamps   []string                                                         `json:"timestamps,required"`
	JSON         radarHTTPTimeseriesGroupGetByPostQuantumResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupGetByPostQuantumResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarHTTPTimeseriesGroupGetByPostQuantumResponseResultSerie0]
type radarHTTPTimeseriesGroupGetByPostQuantumResponseResultSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	Timestamps   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByPostQuantumResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByPostQuantumResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByTlsVersionResponse struct {
	Result  RadarHTTPTimeseriesGroupGetByTlsVersionResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupGetByTlsVersionResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByTlsVersionResponseJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupGetByTlsVersionResponse]
type radarHTTPTimeseriesGroupGetByTlsVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByTlsVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByTlsVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByTlsVersionResponseResult struct {
	Meta   interface{}                                                 `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupGetByTlsVersionResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupGetByTlsVersionResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupGetByTlsVersionResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupGetByTlsVersionResponseResult]
type radarHTTPTimeseriesGroupGetByTlsVersionResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByTlsVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByTlsVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByTlsVersionResponseResultSerie0 struct {
	Timestamps []string                                                        `json:"timestamps,required"`
	Tls1_0     []string                                                        `json:"TLS 1.0,required"`
	Tls1_1     []string                                                        `json:"TLS 1.1,required"`
	Tls1_2     []string                                                        `json:"TLS 1.2,required"`
	Tls1_3     []string                                                        `json:"TLS 1.3,required"`
	TlsQuic    []string                                                        `json:"TLS QUIC,required"`
	JSON       radarHTTPTimeseriesGroupGetByTlsVersionResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupGetByTlsVersionResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarHTTPTimeseriesGroupGetByTlsVersionResponseResultSerie0]
type radarHTTPTimeseriesGroupGetByTlsVersionResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Tls1_0      apijson.Field
	Tls1_1      apijson.Field
	Tls1_2      apijson.Field
	Tls1_3      apijson.Field
	TlsQuic     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupGetByTlsVersionResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTimeseriesGroupGetByTlsVersionResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTimeseriesGroupGetByBotClassParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamily] `query:"browserFamily"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupGetByBotClassParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTimeseriesGroupGetByBotClassParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByBotClassParamsIPVersion] `query:"ipVersion"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTimeseriesGroupGetByBotClassParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupGetByBotClassParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupGetByBotClassParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval15m RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval1h  RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval1d  RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval1w  RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval = "1w"
)

func (r RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval15m, RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval1h, RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval1d, RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval1w:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamily string

const (
	RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamilyChrome  RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamily = "CHROME"
	RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamilyEdge    RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamily = "EDGE"
	RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamilyFirefox RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamily = "FIREFOX"
	RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamilySafari  RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamilyChrome, RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamilyEdge, RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamilyFirefox, RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBotClassParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupGetByBotClassParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupGetByBotClassParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupGetByBotClassParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupGetByBotClassParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupGetByBotClassParamsDeviceTypeOther   RadarHTTPTimeseriesGroupGetByBotClassParamsDeviceType = "OTHER"
)

func (r RadarHTTPTimeseriesGroupGetByBotClassParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBotClassParamsDeviceTypeDesktop, RadarHTTPTimeseriesGroupGetByBotClassParamsDeviceTypeMobile, RadarHTTPTimeseriesGroupGetByBotClassParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTimeseriesGroupGetByBotClassParamsFormat string

const (
	RadarHTTPTimeseriesGroupGetByBotClassParamsFormatJson RadarHTTPTimeseriesGroupGetByBotClassParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupGetByBotClassParamsFormatCsv  RadarHTTPTimeseriesGroupGetByBotClassParamsFormat = "CSV"
)

func (r RadarHTTPTimeseriesGroupGetByBotClassParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBotClassParamsFormatJson, RadarHTTPTimeseriesGroupGetByBotClassParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPProtocolHTTP, RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPVersionHttPv1, RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPVersionHttPv2, RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBotClassParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupGetByBotClassParamsIPVersionIPv4 RadarHTTPTimeseriesGroupGetByBotClassParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupGetByBotClassParamsIPVersionIPv6 RadarHTTPTimeseriesGroupGetByBotClassParamsIPVersion = "IPv6"
)

func (r RadarHTTPTimeseriesGroupGetByBotClassParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBotClassParamsIPVersionIPv4, RadarHTTPTimeseriesGroupGetByBotClassParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBotClassParamsO string

const (
	RadarHTTPTimeseriesGroupGetByBotClassParamsOWindows  RadarHTTPTimeseriesGroupGetByBotClassParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupGetByBotClassParamsOMacosx   RadarHTTPTimeseriesGroupGetByBotClassParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupGetByBotClassParamsOIos      RadarHTTPTimeseriesGroupGetByBotClassParamsO = "IOS"
	RadarHTTPTimeseriesGroupGetByBotClassParamsOAndroid  RadarHTTPTimeseriesGroupGetByBotClassParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupGetByBotClassParamsOChromeos RadarHTTPTimeseriesGroupGetByBotClassParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupGetByBotClassParamsOLinux    RadarHTTPTimeseriesGroupGetByBotClassParamsO = "LINUX"
	RadarHTTPTimeseriesGroupGetByBotClassParamsOSmartTv  RadarHTTPTimeseriesGroupGetByBotClassParamsO = "SMART_TV"
)

func (r RadarHTTPTimeseriesGroupGetByBotClassParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBotClassParamsOWindows, RadarHTTPTimeseriesGroupGetByBotClassParamsOMacosx, RadarHTTPTimeseriesGroupGetByBotClassParamsOIos, RadarHTTPTimeseriesGroupGetByBotClassParamsOAndroid, RadarHTTPTimeseriesGroupGetByBotClassParamsOChromeos, RadarHTTPTimeseriesGroupGetByBotClassParamsOLinux, RadarHTTPTimeseriesGroupGetByBotClassParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersionTlSv1_0, RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersionTlSv1_1, RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersionTlSv1_2, RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersionTlSv1_3, RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupGetByBrowserParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamily] `query:"browserFamily"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupGetByBrowserParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTimeseriesGroupGetByBrowserParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByBrowserParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTimeseriesGroupGetByBrowserParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupGetByBrowserParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupGetByBrowserParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval15m RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval1h  RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval1d  RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval1w  RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval = "1w"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval15m, RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval1h, RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval1d, RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval1w:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserParamsBotClass string

const (
	RadarHTTPTimeseriesGroupGetByBrowserParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupGetByBrowserParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupGetByBrowserParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupGetByBrowserParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserParamsBotClassLikelyAutomated, RadarHTTPTimeseriesGroupGetByBrowserParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamily string

const (
	RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamilyChrome  RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamily = "CHROME"
	RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamilyEdge    RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamily = "EDGE"
	RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamilyFirefox RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamily = "FIREFOX"
	RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamilySafari  RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamilyChrome, RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamilyEdge, RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamilyFirefox, RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupGetByBrowserParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupGetByBrowserParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupGetByBrowserParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupGetByBrowserParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupGetByBrowserParamsDeviceTypeOther   RadarHTTPTimeseriesGroupGetByBrowserParamsDeviceType = "OTHER"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserParamsDeviceTypeDesktop, RadarHTTPTimeseriesGroupGetByBrowserParamsDeviceTypeMobile, RadarHTTPTimeseriesGroupGetByBrowserParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTimeseriesGroupGetByBrowserParamsFormat string

const (
	RadarHTTPTimeseriesGroupGetByBrowserParamsFormatJson RadarHTTPTimeseriesGroupGetByBrowserParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupGetByBrowserParamsFormatCsv  RadarHTTPTimeseriesGroupGetByBrowserParamsFormat = "CSV"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserParamsFormatJson, RadarHTTPTimeseriesGroupGetByBrowserParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPProtocolHTTP, RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPVersionHttPv1, RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPVersionHttPv2, RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupGetByBrowserParamsIPVersionIPv4 RadarHTTPTimeseriesGroupGetByBrowserParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupGetByBrowserParamsIPVersionIPv6 RadarHTTPTimeseriesGroupGetByBrowserParamsIPVersion = "IPv6"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserParamsIPVersionIPv4, RadarHTTPTimeseriesGroupGetByBrowserParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserParamsO string

const (
	RadarHTTPTimeseriesGroupGetByBrowserParamsOWindows  RadarHTTPTimeseriesGroupGetByBrowserParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupGetByBrowserParamsOMacosx   RadarHTTPTimeseriesGroupGetByBrowserParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupGetByBrowserParamsOIos      RadarHTTPTimeseriesGroupGetByBrowserParamsO = "IOS"
	RadarHTTPTimeseriesGroupGetByBrowserParamsOAndroid  RadarHTTPTimeseriesGroupGetByBrowserParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupGetByBrowserParamsOChromeos RadarHTTPTimeseriesGroupGetByBrowserParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupGetByBrowserParamsOLinux    RadarHTTPTimeseriesGroupGetByBrowserParamsO = "LINUX"
	RadarHTTPTimeseriesGroupGetByBrowserParamsOSmartTv  RadarHTTPTimeseriesGroupGetByBrowserParamsO = "SMART_TV"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserParamsOWindows, RadarHTTPTimeseriesGroupGetByBrowserParamsOMacosx, RadarHTTPTimeseriesGroupGetByBrowserParamsOIos, RadarHTTPTimeseriesGroupGetByBrowserParamsOAndroid, RadarHTTPTimeseriesGroupGetByBrowserParamsOChromeos, RadarHTTPTimeseriesGroupGetByBrowserParamsOLinux, RadarHTTPTimeseriesGroupGetByBrowserParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersionTlSv1_0, RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersionTlSv1_1, RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersionTlSv1_2, RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersionTlSv1_3, RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserFamilyParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsBotClass] `query:"botClass"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupGetByBrowserFamilyParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupGetByBrowserFamilyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval15m RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval1h  RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval1d  RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval1w  RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval = "1w"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval15m, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval1h, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval1d, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval1w:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsBotClass string

const (
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsBotClassLikelyAutomated, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsDeviceTypeOther   RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsDeviceType = "OTHER"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsDeviceTypeDesktop, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsDeviceTypeMobile, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsFormat string

const (
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsFormatJson RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsFormatCsv  RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsFormat = "CSV"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsFormatJson, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPProtocolHTTP, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPVersionHttPv1, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPVersionHttPv2, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsIPVersionIPv4 RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsIPVersionIPv6 RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsIPVersion = "IPv6"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsIPVersionIPv4, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsO string

const (
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsOWindows  RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsOMacosx   RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsOIos      RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsO = "IOS"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsOAndroid  RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsOChromeos RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsOLinux    RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsO = "LINUX"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsOSmartTv  RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsO = "SMART_TV"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsOWindows, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsOMacosx, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsOIos, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsOAndroid, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsOChromeos, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsOLinux, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersionTlSv1_0, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersionTlSv1_1, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersionTlSv1_2, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersionTlSv1_3, RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByDeviceTypeParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamily] `query:"browserFamily"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTimeseriesGroupGetByDeviceTypeParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByDeviceTypeParamsIPVersion] `query:"ipVersion"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTimeseriesGroupGetByDeviceTypeParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupGetByDeviceTypeParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupGetByDeviceTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval15m RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval1h  RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval1d  RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval1w  RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval = "1w"
)

func (r RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval15m, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval1h, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval1d, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval1w:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBotClass string

const (
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBotClassLikelyAutomated, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamily string

const (
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamilyChrome  RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamily = "CHROME"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamilyEdge    RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamily = "EDGE"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamilyFirefox RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamily = "FIREFOX"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamilySafari  RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamilyChrome, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamilyEdge, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamilyFirefox, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamilySafari:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTimeseriesGroupGetByDeviceTypeParamsFormat string

const (
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsFormatJson RadarHTTPTimeseriesGroupGetByDeviceTypeParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsFormatCsv  RadarHTTPTimeseriesGroupGetByDeviceTypeParamsFormat = "CSV"
)

func (r RadarHTTPTimeseriesGroupGetByDeviceTypeParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByDeviceTypeParamsFormatJson, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPProtocolHTTP, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPVersionHttPv1, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPVersionHttPv2, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByDeviceTypeParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsIPVersionIPv4 RadarHTTPTimeseriesGroupGetByDeviceTypeParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsIPVersionIPv6 RadarHTTPTimeseriesGroupGetByDeviceTypeParamsIPVersion = "IPv6"
)

func (r RadarHTTPTimeseriesGroupGetByDeviceTypeParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByDeviceTypeParamsIPVersionIPv4, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByDeviceTypeParamsO string

const (
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsOWindows  RadarHTTPTimeseriesGroupGetByDeviceTypeParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsOMacosx   RadarHTTPTimeseriesGroupGetByDeviceTypeParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsOIos      RadarHTTPTimeseriesGroupGetByDeviceTypeParamsO = "IOS"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsOAndroid  RadarHTTPTimeseriesGroupGetByDeviceTypeParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsOChromeos RadarHTTPTimeseriesGroupGetByDeviceTypeParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsOLinux    RadarHTTPTimeseriesGroupGetByDeviceTypeParamsO = "LINUX"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsOSmartTv  RadarHTTPTimeseriesGroupGetByDeviceTypeParamsO = "SMART_TV"
)

func (r RadarHTTPTimeseriesGroupGetByDeviceTypeParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByDeviceTypeParamsOWindows, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsOMacosx, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsOIos, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsOAndroid, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsOChromeos, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsOLinux, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersionTlSv1_0, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersionTlSv1_1, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersionTlSv1_2, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersionTlSv1_3, RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPProtocolParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamily] `query:"browserFamily"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsFormat] `query:"format"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsIPVersion] `query:"ipVersion"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupGetByHTTPProtocolParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupGetByHTTPProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval15m RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval1h  RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval1d  RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval1w  RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval = "1w"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval15m, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval1h, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval1d, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval1w:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBotClass string

const (
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBotClassLikelyAutomated, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamily string

const (
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamilyChrome  RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamily = "CHROME"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamilyEdge    RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamily = "EDGE"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamilyFirefox RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamily = "FIREFOX"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamilySafari  RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamilyChrome, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamilyEdge, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamilyFirefox, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsDeviceTypeOther   RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsDeviceType = "OTHER"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsDeviceTypeDesktop, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsDeviceTypeMobile, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsFormat string

const (
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsFormatJson RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsFormatCsv  RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsFormat = "CSV"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsFormatJson, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsHTTPVersionHttPv1, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsHTTPVersionHttPv2, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsIPVersionIPv4 RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsIPVersionIPv6 RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsIPVersion = "IPv6"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsIPVersionIPv4, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsO string

const (
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsOWindows  RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsOMacosx   RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsOIos      RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsO = "IOS"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsOAndroid  RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsOChromeos RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsOLinux    RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsO = "LINUX"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsOSmartTv  RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsO = "SMART_TV"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsOWindows, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsOMacosx, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsOIos, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsOAndroid, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsOChromeos, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsOLinux, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersionTlSv1_0, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersionTlSv1_1, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersionTlSv1_2, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersionTlSv1_3, RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamily] `query:"browserFamily"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupGetByHTTPVersionParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTimeseriesGroupGetByHTTPVersionParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupGetByHTTPVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByHTTPVersionParamsIPVersion] `query:"ipVersion"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTimeseriesGroupGetByHTTPVersionParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupGetByHTTPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupGetByHTTPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval15m RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval1h  RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval1d  RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval1w  RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval = "1w"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval15m, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval1h, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval1d, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval1w:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBotClass string

const (
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBotClassLikelyAutomated, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamily string

const (
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamilyChrome  RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamily = "CHROME"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamilyEdge    RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamily = "EDGE"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamilyFirefox RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamily = "FIREFOX"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamilySafari  RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamilyChrome, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamilyEdge, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamilyFirefox, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPVersionParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupGetByHTTPVersionParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupGetByHTTPVersionParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsDeviceTypeOther   RadarHTTPTimeseriesGroupGetByHTTPVersionParamsDeviceType = "OTHER"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPVersionParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPVersionParamsDeviceTypeDesktop, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsDeviceTypeMobile, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTimeseriesGroupGetByHTTPVersionParamsFormat string

const (
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsFormatJson RadarHTTPTimeseriesGroupGetByHTTPVersionParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsFormatCsv  RadarHTTPTimeseriesGroupGetByHTTPVersionParamsFormat = "CSV"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPVersionParamsFormatJson, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPVersionParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupGetByHTTPVersionParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupGetByHTTPVersionParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPVersionParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPVersionParamsHTTPProtocolHTTP, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPVersionParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsIPVersionIPv4 RadarHTTPTimeseriesGroupGetByHTTPVersionParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsIPVersionIPv6 RadarHTTPTimeseriesGroupGetByHTTPVersionParamsIPVersion = "IPv6"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPVersionParamsIPVersionIPv4, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPVersionParamsO string

const (
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsOWindows  RadarHTTPTimeseriesGroupGetByHTTPVersionParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsOMacosx   RadarHTTPTimeseriesGroupGetByHTTPVersionParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsOIos      RadarHTTPTimeseriesGroupGetByHTTPVersionParamsO = "IOS"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsOAndroid  RadarHTTPTimeseriesGroupGetByHTTPVersionParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsOChromeos RadarHTTPTimeseriesGroupGetByHTTPVersionParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsOLinux    RadarHTTPTimeseriesGroupGetByHTTPVersionParamsO = "LINUX"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsOSmartTv  RadarHTTPTimeseriesGroupGetByHTTPVersionParamsO = "SMART_TV"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPVersionParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPVersionParamsOWindows, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsOMacosx, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsOIos, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsOAndroid, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsOChromeos, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsOLinux, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersionTlSv1_0, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersionTlSv1_1, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersionTlSv1_2, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersionTlSv1_3, RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupGetByIPVersionParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamily] `query:"browserFamily"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupGetByIPVersionParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTimeseriesGroupGetByIPVersionParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPVersion] `query:"httpVersion"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTimeseriesGroupGetByIPVersionParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupGetByIPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupGetByIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval15m RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval1h  RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval1d  RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval1w  RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval = "1w"
)

func (r RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval15m, RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval1h, RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval1d, RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval1w:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByIPVersionParamsBotClass string

const (
	RadarHTTPTimeseriesGroupGetByIPVersionParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupGetByIPVersionParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupGetByIPVersionParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTimeseriesGroupGetByIPVersionParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByIPVersionParamsBotClassLikelyAutomated, RadarHTTPTimeseriesGroupGetByIPVersionParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamily string

const (
	RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamilyChrome  RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamily = "CHROME"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamilyEdge    RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamily = "EDGE"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamilyFirefox RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamily = "FIREFOX"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamilySafari  RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamilyChrome, RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamilyEdge, RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamilyFirefox, RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByIPVersionParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupGetByIPVersionParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupGetByIPVersionParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupGetByIPVersionParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsDeviceTypeOther   RadarHTTPTimeseriesGroupGetByIPVersionParamsDeviceType = "OTHER"
)

func (r RadarHTTPTimeseriesGroupGetByIPVersionParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByIPVersionParamsDeviceTypeDesktop, RadarHTTPTimeseriesGroupGetByIPVersionParamsDeviceTypeMobile, RadarHTTPTimeseriesGroupGetByIPVersionParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTimeseriesGroupGetByIPVersionParamsFormat string

const (
	RadarHTTPTimeseriesGroupGetByIPVersionParamsFormatJson RadarHTTPTimeseriesGroupGetByIPVersionParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsFormatCsv  RadarHTTPTimeseriesGroupGetByIPVersionParamsFormat = "CSV"
)

func (r RadarHTTPTimeseriesGroupGetByIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByIPVersionParamsFormatJson, RadarHTTPTimeseriesGroupGetByIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPProtocolHTTP, RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPVersionHttPv1, RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPVersionHttPv2, RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByIPVersionParamsO string

const (
	RadarHTTPTimeseriesGroupGetByIPVersionParamsOWindows  RadarHTTPTimeseriesGroupGetByIPVersionParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsOMacosx   RadarHTTPTimeseriesGroupGetByIPVersionParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsOIos      RadarHTTPTimeseriesGroupGetByIPVersionParamsO = "IOS"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsOAndroid  RadarHTTPTimeseriesGroupGetByIPVersionParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsOChromeos RadarHTTPTimeseriesGroupGetByIPVersionParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsOLinux    RadarHTTPTimeseriesGroupGetByIPVersionParamsO = "LINUX"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsOSmartTv  RadarHTTPTimeseriesGroupGetByIPVersionParamsO = "SMART_TV"
)

func (r RadarHTTPTimeseriesGroupGetByIPVersionParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByIPVersionParamsOWindows, RadarHTTPTimeseriesGroupGetByIPVersionParamsOMacosx, RadarHTTPTimeseriesGroupGetByIPVersionParamsOIos, RadarHTTPTimeseriesGroupGetByIPVersionParamsOAndroid, RadarHTTPTimeseriesGroupGetByIPVersionParamsOChromeos, RadarHTTPTimeseriesGroupGetByIPVersionParamsOLinux, RadarHTTPTimeseriesGroupGetByIPVersionParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersionTlSv1_0, RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersionTlSv1_1, RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersionTlSv1_2, RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersionTlSv1_3, RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByOsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupGetByOsParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupGetByOsParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamily] `query:"browserFamily"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupGetByOsParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTimeseriesGroupGetByOsParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupGetByOsParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByOsParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByOsParamsIPVersion] `query:"ipVersion"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupGetByOsParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupGetByOsParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTimeseriesGroupGetByOsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupGetByOsParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupGetByOsParamsAggInterval15m RadarHTTPTimeseriesGroupGetByOsParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupGetByOsParamsAggInterval1h  RadarHTTPTimeseriesGroupGetByOsParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupGetByOsParamsAggInterval1d  RadarHTTPTimeseriesGroupGetByOsParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupGetByOsParamsAggInterval1w  RadarHTTPTimeseriesGroupGetByOsParamsAggInterval = "1w"
)

func (r RadarHTTPTimeseriesGroupGetByOsParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByOsParamsAggInterval15m, RadarHTTPTimeseriesGroupGetByOsParamsAggInterval1h, RadarHTTPTimeseriesGroupGetByOsParamsAggInterval1d, RadarHTTPTimeseriesGroupGetByOsParamsAggInterval1w:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByOsParamsBotClass string

const (
	RadarHTTPTimeseriesGroupGetByOsParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupGetByOsParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupGetByOsParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupGetByOsParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTimeseriesGroupGetByOsParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByOsParamsBotClassLikelyAutomated, RadarHTTPTimeseriesGroupGetByOsParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamily string

const (
	RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamilyChrome  RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamily = "CHROME"
	RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamilyEdge    RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamily = "EDGE"
	RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamilyFirefox RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamily = "FIREFOX"
	RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamilySafari  RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamilyChrome, RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamilyEdge, RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamilyFirefox, RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByOsParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupGetByOsParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupGetByOsParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupGetByOsParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupGetByOsParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupGetByOsParamsDeviceTypeOther   RadarHTTPTimeseriesGroupGetByOsParamsDeviceType = "OTHER"
)

func (r RadarHTTPTimeseriesGroupGetByOsParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByOsParamsDeviceTypeDesktop, RadarHTTPTimeseriesGroupGetByOsParamsDeviceTypeMobile, RadarHTTPTimeseriesGroupGetByOsParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTimeseriesGroupGetByOsParamsFormat string

const (
	RadarHTTPTimeseriesGroupGetByOsParamsFormatJson RadarHTTPTimeseriesGroupGetByOsParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupGetByOsParamsFormatCsv  RadarHTTPTimeseriesGroupGetByOsParamsFormat = "CSV"
)

func (r RadarHTTPTimeseriesGroupGetByOsParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByOsParamsFormatJson, RadarHTTPTimeseriesGroupGetByOsParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByOsParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupGetByOsParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupGetByOsParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupGetByOsParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupGetByOsParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTimeseriesGroupGetByOsParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByOsParamsHTTPProtocolHTTP, RadarHTTPTimeseriesGroupGetByOsParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByOsParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupGetByOsParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupGetByOsParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupGetByOsParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupGetByOsParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupGetByOsParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupGetByOsParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTimeseriesGroupGetByOsParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByOsParamsHTTPVersionHttPv1, RadarHTTPTimeseriesGroupGetByOsParamsHTTPVersionHttPv2, RadarHTTPTimeseriesGroupGetByOsParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByOsParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupGetByOsParamsIPVersionIPv4 RadarHTTPTimeseriesGroupGetByOsParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupGetByOsParamsIPVersionIPv6 RadarHTTPTimeseriesGroupGetByOsParamsIPVersion = "IPv6"
)

func (r RadarHTTPTimeseriesGroupGetByOsParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByOsParamsIPVersionIPv4, RadarHTTPTimeseriesGroupGetByOsParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByOsParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupGetByOsParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupGetByOsParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupGetByOsParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupGetByOsParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupGetByOsParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupGetByOsParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupGetByOsParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupGetByOsParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupGetByOsParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupGetByOsParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTimeseriesGroupGetByOsParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByOsParamsTlsVersionTlSv1_0, RadarHTTPTimeseriesGroupGetByOsParamsTlsVersionTlSv1_1, RadarHTTPTimeseriesGroupGetByOsParamsTlsVersionTlSv1_2, RadarHTTPTimeseriesGroupGetByOsParamsTlsVersionTlSv1_3, RadarHTTPTimeseriesGroupGetByOsParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByPostQuantumParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupGetByPostQuantumParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamily] `query:"browserFamily"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupGetByPostQuantumParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTimeseriesGroupGetByPostQuantumParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByPostQuantumParamsIPVersion] `query:"ipVersion"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTimeseriesGroupGetByPostQuantumParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupGetByPostQuantumParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupGetByPostQuantumParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval15m RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval1h  RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval1d  RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval1w  RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval = "1w"
)

func (r RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval15m, RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval1h, RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval1d, RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval1w:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByPostQuantumParamsBotClass string

const (
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupGetByPostQuantumParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupGetByPostQuantumParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTimeseriesGroupGetByPostQuantumParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByPostQuantumParamsBotClassLikelyAutomated, RadarHTTPTimeseriesGroupGetByPostQuantumParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamily string

const (
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamilyChrome  RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamily = "CHROME"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamilyEdge    RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamily = "EDGE"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamilyFirefox RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamily = "FIREFOX"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamilySafari  RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamilyChrome, RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamilyEdge, RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamilyFirefox, RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByPostQuantumParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupGetByPostQuantumParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupGetByPostQuantumParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsDeviceTypeOther   RadarHTTPTimeseriesGroupGetByPostQuantumParamsDeviceType = "OTHER"
)

func (r RadarHTTPTimeseriesGroupGetByPostQuantumParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByPostQuantumParamsDeviceTypeDesktop, RadarHTTPTimeseriesGroupGetByPostQuantumParamsDeviceTypeMobile, RadarHTTPTimeseriesGroupGetByPostQuantumParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTimeseriesGroupGetByPostQuantumParamsFormat string

const (
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsFormatJson RadarHTTPTimeseriesGroupGetByPostQuantumParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsFormatCsv  RadarHTTPTimeseriesGroupGetByPostQuantumParamsFormat = "CSV"
)

func (r RadarHTTPTimeseriesGroupGetByPostQuantumParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByPostQuantumParamsFormatJson, RadarHTTPTimeseriesGroupGetByPostQuantumParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPProtocolHTTP, RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPVersionHttPv1, RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPVersionHttPv2, RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByPostQuantumParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsIPVersionIPv4 RadarHTTPTimeseriesGroupGetByPostQuantumParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsIPVersionIPv6 RadarHTTPTimeseriesGroupGetByPostQuantumParamsIPVersion = "IPv6"
)

func (r RadarHTTPTimeseriesGroupGetByPostQuantumParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByPostQuantumParamsIPVersionIPv4, RadarHTTPTimeseriesGroupGetByPostQuantumParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByPostQuantumParamsO string

const (
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsOWindows  RadarHTTPTimeseriesGroupGetByPostQuantumParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsOMacosx   RadarHTTPTimeseriesGroupGetByPostQuantumParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsOIos      RadarHTTPTimeseriesGroupGetByPostQuantumParamsO = "IOS"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsOAndroid  RadarHTTPTimeseriesGroupGetByPostQuantumParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsOChromeos RadarHTTPTimeseriesGroupGetByPostQuantumParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsOLinux    RadarHTTPTimeseriesGroupGetByPostQuantumParamsO = "LINUX"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsOSmartTv  RadarHTTPTimeseriesGroupGetByPostQuantumParamsO = "SMART_TV"
)

func (r RadarHTTPTimeseriesGroupGetByPostQuantumParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByPostQuantumParamsOWindows, RadarHTTPTimeseriesGroupGetByPostQuantumParamsOMacosx, RadarHTTPTimeseriesGroupGetByPostQuantumParamsOIos, RadarHTTPTimeseriesGroupGetByPostQuantumParamsOAndroid, RadarHTTPTimeseriesGroupGetByPostQuantumParamsOChromeos, RadarHTTPTimeseriesGroupGetByPostQuantumParamsOLinux, RadarHTTPTimeseriesGroupGetByPostQuantumParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersionTlSv1_0, RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersionTlSv1_1, RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersionTlSv1_2, RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersionTlSv1_3, RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByTlsVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupGetByTlsVersionParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamily] `query:"browserFamily"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupGetByTlsVersionParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTimeseriesGroupGetByTlsVersionParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupGetByTlsVersionParamsIPVersion] `query:"ipVersion"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTimeseriesGroupGetByTlsVersionParamsO] `query:"os"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupGetByTlsVersionParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupGetByTlsVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval15m RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval1h  RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval1d  RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval1w  RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval = "1w"
)

func (r RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval15m, RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval1h, RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval1d, RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval1w:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByTlsVersionParamsBotClass string

const (
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupGetByTlsVersionParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupGetByTlsVersionParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTimeseriesGroupGetByTlsVersionParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByTlsVersionParamsBotClassLikelyAutomated, RadarHTTPTimeseriesGroupGetByTlsVersionParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamily string

const (
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamilyChrome  RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamily = "CHROME"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamilyEdge    RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamily = "EDGE"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamilyFirefox RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamily = "FIREFOX"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamilySafari  RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamilyChrome, RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamilyEdge, RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamilyFirefox, RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByTlsVersionParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupGetByTlsVersionParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupGetByTlsVersionParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsDeviceTypeOther   RadarHTTPTimeseriesGroupGetByTlsVersionParamsDeviceType = "OTHER"
)

func (r RadarHTTPTimeseriesGroupGetByTlsVersionParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByTlsVersionParamsDeviceTypeDesktop, RadarHTTPTimeseriesGroupGetByTlsVersionParamsDeviceTypeMobile, RadarHTTPTimeseriesGroupGetByTlsVersionParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTimeseriesGroupGetByTlsVersionParamsFormat string

const (
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsFormatJson RadarHTTPTimeseriesGroupGetByTlsVersionParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsFormatCsv  RadarHTTPTimeseriesGroupGetByTlsVersionParamsFormat = "CSV"
)

func (r RadarHTTPTimeseriesGroupGetByTlsVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByTlsVersionParamsFormatJson, RadarHTTPTimeseriesGroupGetByTlsVersionParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPProtocolHTTPS RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPProtocolHTTP, RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPVersionHttPv1, RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPVersionHttPv2, RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByTlsVersionParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsIPVersionIPv4 RadarHTTPTimeseriesGroupGetByTlsVersionParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsIPVersionIPv6 RadarHTTPTimeseriesGroupGetByTlsVersionParamsIPVersion = "IPv6"
)

func (r RadarHTTPTimeseriesGroupGetByTlsVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByTlsVersionParamsIPVersionIPv4, RadarHTTPTimeseriesGroupGetByTlsVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTimeseriesGroupGetByTlsVersionParamsO string

const (
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsOWindows  RadarHTTPTimeseriesGroupGetByTlsVersionParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsOMacosx   RadarHTTPTimeseriesGroupGetByTlsVersionParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsOIos      RadarHTTPTimeseriesGroupGetByTlsVersionParamsO = "IOS"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsOAndroid  RadarHTTPTimeseriesGroupGetByTlsVersionParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsOChromeos RadarHTTPTimeseriesGroupGetByTlsVersionParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsOLinux    RadarHTTPTimeseriesGroupGetByTlsVersionParamsO = "LINUX"
	RadarHTTPTimeseriesGroupGetByTlsVersionParamsOSmartTv  RadarHTTPTimeseriesGroupGetByTlsVersionParamsO = "SMART_TV"
)

func (r RadarHTTPTimeseriesGroupGetByTlsVersionParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTimeseriesGroupGetByTlsVersionParamsOWindows, RadarHTTPTimeseriesGroupGetByTlsVersionParamsOMacosx, RadarHTTPTimeseriesGroupGetByTlsVersionParamsOIos, RadarHTTPTimeseriesGroupGetByTlsVersionParamsOAndroid, RadarHTTPTimeseriesGroupGetByTlsVersionParamsOChromeos, RadarHTTPTimeseriesGroupGetByTlsVersionParamsOLinux, RadarHTTPTimeseriesGroupGetByTlsVersionParamsOSmartTv:
		return true
	}
	return false
}
