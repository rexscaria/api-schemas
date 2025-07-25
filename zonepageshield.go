// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
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

type MessagesPageShieldItem struct {
	Code             int64                        `json:"code,required"`
	Message          string                       `json:"message,required"`
	DocumentationURL string                       `json:"documentation_url"`
	Source           MessagesPageShieldItemSource `json:"source"`
	JSON             messagesPageShieldItemJSON   `json:"-"`
}

// messagesPageShieldItemJSON contains the JSON metadata for the struct
// [MessagesPageShieldItem]
type messagesPageShieldItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessagesPageShieldItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesPageShieldItemJSON) RawJSON() string {
	return r.raw
}

type MessagesPageShieldItemSource struct {
	Pointer string                           `json:"pointer"`
	JSON    messagesPageShieldItemSourceJSON `json:"-"`
}

// messagesPageShieldItemSourceJSON contains the JSON metadata for the struct
// [MessagesPageShieldItemSource]
type messagesPageShieldItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesPageShieldItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesPageShieldItemSourceJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldGetResponse struct {
	// Whether the API call was successful
	Success  ZonePageShieldGetResponseSuccess `json:"success,required"`
	Errors   []MessagesPageShieldItem         `json:"errors"`
	Messages []MessagesPageShieldItem         `json:"messages"`
	Result   ZonePageShieldGetResponseResult  `json:"result,nullable"`
	JSON     zonePageShieldGetResponseJSON    `json:"-"`
}

// zonePageShieldGetResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldGetResponse]
type zonePageShieldGetResponseJSON struct {
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
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

// Whether the API call was successful
type ZonePageShieldGetResponseSuccess bool

const (
	ZonePageShieldGetResponseSuccessTrue ZonePageShieldGetResponseSuccess = true
)

func (r ZonePageShieldGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZonePageShieldGetResponseSuccessTrue:
		return true
	}
	return false
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
	// Whether the API call was successful
	Success  ZonePageShieldUpdateResponseSuccess `json:"success,required"`
	Errors   []MessagesPageShieldItem            `json:"errors"`
	Messages []MessagesPageShieldItem            `json:"messages"`
	Result   ZonePageShieldUpdateResponseResult  `json:"result"`
	JSON     zonePageShieldUpdateResponseJSON    `json:"-"`
}

// zonePageShieldUpdateResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldUpdateResponse]
type zonePageShieldUpdateResponseJSON struct {
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
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

// Whether the API call was successful
type ZonePageShieldUpdateResponseSuccess bool

const (
	ZonePageShieldUpdateResponseSuccessTrue ZonePageShieldUpdateResponseSuccess = true
)

func (r ZonePageShieldUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZonePageShieldUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZonePageShieldUpdateResponseResult struct {
	// When true, indicates that Page Shield is enabled.
	Enabled bool `json:"enabled,required"`
	// The timestamp of when Page Shield was last updated.
	UpdatedAt string `json:"updated_at,required"`
	// When true, CSP reports will be sent to
	// https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report
	UseCloudflareReportingEndpoint bool `json:"use_cloudflare_reporting_endpoint,required"`
	// When true, the paths associated with connections URLs will also be analyzed.
	UseConnectionURLPath bool                                   `json:"use_connection_url_path,required"`
	JSON                 zonePageShieldUpdateResponseResultJSON `json:"-"`
}

// zonePageShieldUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZonePageShieldUpdateResponseResult]
type zonePageShieldUpdateResponseResultJSON struct {
	Enabled                        apijson.Field
	UpdatedAt                      apijson.Field
	UseCloudflareReportingEndpoint apijson.Field
	UseConnectionURLPath           apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *ZonePageShieldUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldUpdateResponseResultJSON) RawJSON() string {
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
