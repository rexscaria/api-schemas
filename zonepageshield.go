// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZonePageShieldService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePageShieldService] method instead.
type ZonePageShieldService struct {
	Options     []option.RequestOption
	Connections *ZonePageShieldConnectionService
	Cookies     *ZonePageShieldCookieService
	Policies    *ZonePageShieldPolicyService
	Scripts     *ZonePageShieldScriptService
}

// NewZonePageShieldService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZonePageShieldService(opts ...option.RequestOption) (r *ZonePageShieldService) {
	r = &ZonePageShieldService{}
	r.Options = opts
	r.Connections = NewZonePageShieldConnectionService(opts...)
	r.Cookies = NewZonePageShieldCookieService(opts...)
	r.Policies = NewZonePageShieldPolicyService(opts...)
	r.Scripts = NewZonePageShieldScriptService(opts...)
	return
}

// Fetches the Page Shield settings.
func (r *ZonePageShieldService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZonePageShieldGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates Page Shield settings.
func (r *ZonePageShieldService) Update(ctx context.Context, zoneID string, body ZonePageShieldUpdateParams, opts ...option.RequestOption) (res *ZonePageShieldUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type GetResponseCollection struct {
	Result interface{}               `json:"result,nullable"`
	JSON   getResponseCollectionJSON `json:"-"`
	ResponseCommonShield
}

// getResponseCollectionJSON contains the JSON metadata for the struct
// [GetResponseCollection]
type getResponseCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type MessagesPageShieldItem struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    messagesPageShieldItemJSON `json:"-"`
}

// messagesPageShieldItemJSON contains the JSON metadata for the struct
// [MessagesPageShieldItem]
type messagesPageShieldItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesPageShieldItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesPageShieldItemJSON) RawJSON() string {
	return r.raw
}

type ResponseCommonShield struct {
	// Whether the API call was successful
	Success  ResponseCommonShieldSuccess `json:"success,required"`
	Errors   []MessagesPageShieldItem    `json:"errors"`
	Messages []MessagesPageShieldItem    `json:"messages"`
	JSON     responseCommonShieldJSON    `json:"-"`
}

// responseCommonShieldJSON contains the JSON metadata for the struct
// [ResponseCommonShield]
type responseCommonShieldJSON struct {
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCommonShield) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCommonShieldJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ResponseCommonShieldSuccess bool

const (
	ResponseCommonShieldSuccessTrue ResponseCommonShieldSuccess = true
)

func (r ResponseCommonShieldSuccess) IsKnown() bool {
	switch r {
	case ResponseCommonShieldSuccessTrue:
		return true
	}
	return false
}

type ZonePageShieldGetResponse struct {
	Result ZonePageShieldGetResponseResult `json:"result"`
	JSON   zonePageShieldGetResponseJSON   `json:"-"`
	GetResponseCollection
}

// zonePageShieldGetResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldGetResponse]
type zonePageShieldGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldGetResponseResult struct {
	// When true, indicates that Page Shield is enabled.
	Enabled bool `json:"enabled,required"`
	// The timestamp of when Page Shield was last updated.
	UpdatedAt string `json:"updated_at,required"`
	// When true, CSP reports will be sent to
	// https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report
	UseCloudflareReportingEndpoint bool `json:"use_cloudflare_reporting_endpoint,required"`
	// When true, the paths associated with connections URLs will also be analyzed.
	UseConnectionURLPath bool                                `json:"use_connection_url_path,required"`
	JSON                 zonePageShieldGetResponseResultJSON `json:"-"`
}

// zonePageShieldGetResponseResultJSON contains the JSON metadata for the struct
// [ZonePageShieldGetResponseResult]
type zonePageShieldGetResponseResultJSON struct {
	Enabled                        apijson.Field
	UpdatedAt                      apijson.Field
	UseCloudflareReportingEndpoint apijson.Field
	UseConnectionURLPath           apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *ZonePageShieldGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldUpdateResponse struct {
	Result interface{}                      `json:"result"`
	JSON   zonePageShieldUpdateResponseJSON `json:"-"`
	ResponseCommonShield
}

// zonePageShieldUpdateResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldUpdateResponse]
type zonePageShieldUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldUpdateParams struct {
	// When true, indicates that Page Shield is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// When true, CSP reports will be sent to
	// https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report
	UseCloudflareReportingEndpoint param.Field[bool] `json:"use_cloudflare_reporting_endpoint"`
	// When true, the paths associated with connections URLs will also be analyzed.
	UseConnectionURLPath param.Field[bool] `json:"use_connection_url_path"`
}

func (r ZonePageShieldUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
