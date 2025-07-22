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
	// Whether the API call was successful
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

// Whether the API call was successful
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
	Code    int64               `json:"code,required"`
	Message string              `json:"message,required"`
	JSON    messagesDlsItemJSON `json:"-"`
}

// messagesDlsItemJSON contains the JSON metadata for the struct [MessagesDlsItem]
type messagesDlsItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesDlsItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesDlsItemJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameResponse struct {
	// When the regional hostname was created
	CreatedOn RegionalHostnameResponseCreatedOn `json:"created_on,required"`
	// DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are
	// supported for one level, e.g `*.example.com`
	Hostname string `json:"hostname,required"`
	// Identifying key for the region
	RegionKey string                       `json:"region_key,required"`
	JSON      regionalHostnameResponseJSON `json:"-"`
}

// regionalHostnameResponseJSON contains the JSON metadata for the struct
// [RegionalHostnameResponse]
type regionalHostnameResponseJSON struct {
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	RegionKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameResponseJSON) RawJSON() string {
	return r.raw
}

// When the regional hostname was created
type RegionalHostnameResponseCreatedOn struct {
	JSON regionalHostnameResponseCreatedOnJSON `json:"-"`
}

// regionalHostnameResponseCreatedOnJSON contains the JSON metadata for the struct
// [RegionalHostnameResponseCreatedOn]
type regionalHostnameResponseCreatedOnJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameResponseCreatedOn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameResponseCreatedOnJSON) RawJSON() string {
	return r.raw
}

type ZoneAddressingRegionalHostnameNewResponse struct {
	Result RegionalHostnameResponse                      `json:"result"`
	JSON   zoneAddressingRegionalHostnameNewResponseJSON `json:"-"`
	APIResponseDls
}

// zoneAddressingRegionalHostnameNewResponseJSON contains the JSON metadata for the
// struct [ZoneAddressingRegionalHostnameNewResponse]
type zoneAddressingRegionalHostnameNewResponseJSON struct {
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

type ZoneAddressingRegionalHostnameGetResponse struct {
	Result RegionalHostnameResponse                      `json:"result"`
	JSON   zoneAddressingRegionalHostnameGetResponseJSON `json:"-"`
	APIResponseDls
}

// zoneAddressingRegionalHostnameGetResponseJSON contains the JSON metadata for the
// struct [ZoneAddressingRegionalHostnameGetResponse]
type zoneAddressingRegionalHostnameGetResponseJSON struct {
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

type ZoneAddressingRegionalHostnameUpdateResponse struct {
	Result RegionalHostnameResponse                         `json:"result"`
	JSON   zoneAddressingRegionalHostnameUpdateResponseJSON `json:"-"`
	APIResponseDls
}

// zoneAddressingRegionalHostnameUpdateResponseJSON contains the JSON metadata for
// the struct [ZoneAddressingRegionalHostnameUpdateResponse]
type zoneAddressingRegionalHostnameUpdateResponseJSON struct {
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

type ZoneAddressingRegionalHostnameListResponse struct {
	Result []RegionalHostnameResponse                     `json:"result"`
	JSON   zoneAddressingRegionalHostnameListResponseJSON `json:"-"`
	APIResponseCollectionDls
}

// zoneAddressingRegionalHostnameListResponseJSON contains the JSON metadata for
// the struct [ZoneAddressingRegionalHostnameListResponse]
type zoneAddressingRegionalHostnameListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAddressingRegionalHostnameListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAddressingRegionalHostnameListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAddressingRegionalHostnameNewParams struct {
	// DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are
	// supported for one level, e.g `*.example.com`
	Hostname param.Field[string] `json:"hostname,required"`
	// Identifying key for the region
	RegionKey param.Field[string] `json:"region_key,required"`
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
