// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneAddressingRegionalHostnameService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAddressingRegionalHostnameService] method instead.
type ZoneAddressingRegionalHostnameService struct {
	Options []option.RequestOption
}

// NewZoneAddressingRegionalHostnameService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneAddressingRegionalHostnameService(opts ...option.RequestOption) (r *ZoneAddressingRegionalHostnameService) {
	r = &ZoneAddressingRegionalHostnameService{}
	r.Options = opts
	return
}

// Create a new Regional Hostname entry. Cloudflare will only use data centers that
// are physically located within the chosen region to decrypt and service HTTPS
// traffic. Learn more about
// [Regional Services](https://developers.cloudflare.com/data-localization/regional-services/get-started/).
func (r *ZoneAddressingRegionalHostnameService) New(ctx context.Context, zoneID string, body ZoneAddressingRegionalHostnameNewParams, opts ...option.RequestOption) (res *ZoneAddressingRegionalHostnameNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/addressing/regional_hostnames", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch the configuration for a specific Regional Hostname, within a zone.
func (r *ZoneAddressingRegionalHostnameService) Get(ctx context.Context, zoneID string, hostname string, opts ...option.RequestOption) (res *ZoneAddressingRegionalHostnameGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if hostname == "" {
		err = errors.New("missing required hostname parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/addressing/regional_hostnames/%s", zoneID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the configuration for a specific Regional Hostname. Only the region_key
// of a hostname is mutable.
func (r *ZoneAddressingRegionalHostnameService) Update(ctx context.Context, zoneID string, hostname string, body ZoneAddressingRegionalHostnameUpdateParams, opts ...option.RequestOption) (res *ZoneAddressingRegionalHostnameUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if hostname == "" {
		err = errors.New("missing required hostname parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/addressing/regional_hostnames/%s", zoneID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all Regional Hostnames within a zone.
func (r *ZoneAddressingRegionalHostnameService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneAddressingRegionalHostnameListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/addressing/regional_hostnames", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the region configuration for a specific Regional Hostname.
func (r *ZoneAddressingRegionalHostnameService) Delete(ctx context.Context, zoneID string, hostname string, opts ...option.RequestOption) (res *APIResponseDls, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if hostname == "" {
		err = errors.New("missing required hostname parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/addressing/regional_hostnames/%s", zoneID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type APIResponseDls struct {
	Errors   []MessagesDlsItem `json:"errors,required"`
	Messages []MessagesDlsItem `json:"messages,required"`
	// Whether the API call was successful.
	Success APIResponseDlsSuccess `json:"success,required"`
	JSON    apiResponseDlsJSON    `json:"-"`
}

// apiResponseDlsJSON contains the JSON metadata for the struct [APIResponseDls]
type apiResponseDlsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseDls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseDlsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type APIResponseDlsSuccess bool

const (
	APIResponseDlsSuccessTrue APIResponseDlsSuccess = true
)

func (r APIResponseDlsSuccess) IsKnown() bool {
	switch r {
	case APIResponseDlsSuccessTrue:
		return true
	}
	return false
}

type MessagesDlsItem struct {
	Code             int64                 `json:"code,required"`
	Message          string                `json:"message,required"`
	DocumentationURL string                `json:"documentation_url"`
	Source           MessagesDlsItemSource `json:"source"`
	JSON             messagesDlsItemJSON   `json:"-"`
}

// messagesDlsItemJSON contains the JSON metadata for the struct [MessagesDlsItem]
type messagesDlsItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessagesDlsItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesDlsItemJSON) RawJSON() string {
	return r.raw
}

type MessagesDlsItemSource struct {
	Pointer string                    `json:"pointer"`
	JSON    messagesDlsItemSourceJSON `json:"-"`
}

// messagesDlsItemSourceJSON contains the JSON metadata for the struct
// [MessagesDlsItemSource]
type messagesDlsItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesDlsItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesDlsItemSourceJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameResponse struct {
	// When the regional hostname was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are
	// supported for one level, e.g `*.example.com`
	Hostname string `json:"hostname,required"`
	// Identifying key for the region
	RegionKey string `json:"region_key,required"`
	// Configure which routing method to use for the regional hostname
	Routing string                       `json:"routing"`
	JSON    regionalHostnameResponseJSON `json:"-"`
}

// regionalHostnameResponseJSON contains the JSON metadata for the struct
// [RegionalHostnameResponse]
type regionalHostnameResponseJSON struct {
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	RegionKey   apijson.Field
	Routing     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAddressingRegionalHostnameNewResponse struct {
	Errors   []MessagesDlsItem `json:"errors,required"`
	Messages []MessagesDlsItem `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneAddressingRegionalHostnameNewResponseSuccess `json:"success,required"`
	Result  RegionalHostnameResponse                         `json:"result"`
	JSON    zoneAddressingRegionalHostnameNewResponseJSON    `json:"-"`
}

// zoneAddressingRegionalHostnameNewResponseJSON contains the JSON metadata for the
// struct [ZoneAddressingRegionalHostnameNewResponse]
type zoneAddressingRegionalHostnameNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAddressingRegionalHostnameNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAddressingRegionalHostnameNewResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneAddressingRegionalHostnameNewResponseSuccess bool

const (
	ZoneAddressingRegionalHostnameNewResponseSuccessTrue ZoneAddressingRegionalHostnameNewResponseSuccess = true
)

func (r ZoneAddressingRegionalHostnameNewResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAddressingRegionalHostnameNewResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAddressingRegionalHostnameGetResponse struct {
	Errors   []MessagesDlsItem `json:"errors,required"`
	Messages []MessagesDlsItem `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneAddressingRegionalHostnameGetResponseSuccess `json:"success,required"`
	Result  RegionalHostnameResponse                         `json:"result"`
	JSON    zoneAddressingRegionalHostnameGetResponseJSON    `json:"-"`
}

// zoneAddressingRegionalHostnameGetResponseJSON contains the JSON metadata for the
// struct [ZoneAddressingRegionalHostnameGetResponse]
type zoneAddressingRegionalHostnameGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAddressingRegionalHostnameGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAddressingRegionalHostnameGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneAddressingRegionalHostnameGetResponseSuccess bool

const (
	ZoneAddressingRegionalHostnameGetResponseSuccessTrue ZoneAddressingRegionalHostnameGetResponseSuccess = true
)

func (r ZoneAddressingRegionalHostnameGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAddressingRegionalHostnameGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAddressingRegionalHostnameUpdateResponse struct {
	Errors   []MessagesDlsItem `json:"errors,required"`
	Messages []MessagesDlsItem `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneAddressingRegionalHostnameUpdateResponseSuccess `json:"success,required"`
	Result  RegionalHostnameResponse                            `json:"result"`
	JSON    zoneAddressingRegionalHostnameUpdateResponseJSON    `json:"-"`
}

// zoneAddressingRegionalHostnameUpdateResponseJSON contains the JSON metadata for
// the struct [ZoneAddressingRegionalHostnameUpdateResponse]
type zoneAddressingRegionalHostnameUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAddressingRegionalHostnameUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAddressingRegionalHostnameUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneAddressingRegionalHostnameUpdateResponseSuccess bool

const (
	ZoneAddressingRegionalHostnameUpdateResponseSuccessTrue ZoneAddressingRegionalHostnameUpdateResponseSuccess = true
)

func (r ZoneAddressingRegionalHostnameUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAddressingRegionalHostnameUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAddressingRegionalHostnameListResponse struct {
	Errors   []MessagesDlsItem `json:"errors,required"`
	Messages []MessagesDlsItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    ZoneAddressingRegionalHostnameListResponseSuccess    `json:"success,required"`
	Result     []RegionalHostnameResponse                           `json:"result"`
	ResultInfo ZoneAddressingRegionalHostnameListResponseResultInfo `json:"result_info"`
	JSON       zoneAddressingRegionalHostnameListResponseJSON       `json:"-"`
}

// zoneAddressingRegionalHostnameListResponseJSON contains the JSON metadata for
// the struct [ZoneAddressingRegionalHostnameListResponse]
type zoneAddressingRegionalHostnameListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAddressingRegionalHostnameListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAddressingRegionalHostnameListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneAddressingRegionalHostnameListResponseSuccess bool

const (
	ZoneAddressingRegionalHostnameListResponseSuccessTrue ZoneAddressingRegionalHostnameListResponseSuccess = true
)

func (r ZoneAddressingRegionalHostnameListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAddressingRegionalHostnameListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAddressingRegionalHostnameListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                  `json:"total_count"`
	JSON       zoneAddressingRegionalHostnameListResponseResultInfoJSON `json:"-"`
}

// zoneAddressingRegionalHostnameListResponseResultInfoJSON contains the JSON
// metadata for the struct [ZoneAddressingRegionalHostnameListResponseResultInfo]
type zoneAddressingRegionalHostnameListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAddressingRegionalHostnameListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAddressingRegionalHostnameListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneAddressingRegionalHostnameNewParams struct {
	// DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are
	// supported for one level, e.g `*.example.com`
	Hostname param.Field[string] `json:"hostname,required"`
	// Identifying key for the region
	RegionKey param.Field[string] `json:"region_key,required"`
	// Configure which routing method to use for the regional hostname
	Routing param.Field[string] `json:"routing"`
}

func (r ZoneAddressingRegionalHostnameNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAddressingRegionalHostnameUpdateParams struct {
	// Identifying key for the region
	RegionKey param.Field[string] `json:"region_key,required"`
}

func (r ZoneAddressingRegionalHostnameUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
