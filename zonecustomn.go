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

// ZoneCustomNService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCustomNService] method instead.
type ZoneCustomNService struct {
	Options []option.RequestOption
}

// NewZoneCustomNService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneCustomNService(opts ...option.RequestOption) (r *ZoneCustomNService) {
	r = &ZoneCustomNService{}
	r.Options = opts
	return
}

// Get metadata for account-level custom nameservers on a zone.
//
// Deprecated in favor of
// [Show DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-a-zone-list-dns-settings).
//
// Deprecated: deprecated
func (r *ZoneCustomNService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneCustomNGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_ns", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set metadata for account-level custom nameservers on a zone.
//
// If you would like new zones in the account to use account custom nameservers by
// default, use PUT /accounts/:identifier to set the account setting
// use_account_custom_ns_by_default to true.
//
// Deprecated in favor of
// [Update DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-a-zone-update-dns-settings).
//
// Deprecated: deprecated
func (r *ZoneCustomNService) Update(ctx context.Context, zoneID string, body ZoneCustomNUpdateParams, opts ...option.RequestOption) (res *ZoneCustomNUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_ns", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneMetadata struct {
	// Whether zone uses account-level custom nameservers.
	Enabled bool `json:"enabled"`
	// The number of the name server set to assign to the zone.
	NsSet float64          `json:"ns_set"`
	JSON  zoneMetadataJSON `json:"-"`
}

// zoneMetadataJSON contains the JSON metadata for the struct [ZoneMetadata]
type zoneMetadataJSON struct {
	Enabled     apijson.Field
	NsSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneMetadataJSON) RawJSON() string {
	return r.raw
}

type ZoneMetadataParam struct {
	// Whether zone uses account-level custom nameservers.
	Enabled param.Field[bool] `json:"enabled"`
	// The number of the name server set to assign to the zone.
	NsSet param.Field[float64] `json:"ns_set"`
}

func (r ZoneMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCustomNGetResponse struct {
	JSON zoneCustomNGetResponseJSON `json:"-"`
	CustomNsCollection
	ZoneMetadata
}

// zoneCustomNGetResponseJSON contains the JSON metadata for the struct
// [ZoneCustomNGetResponse]
type zoneCustomNGetResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomNUpdateResponse struct {
	Result []string                      `json:"result"`
	JSON   zoneCustomNUpdateResponseJSON `json:"-"`
	CustomNsCollection
}

// zoneCustomNUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneCustomNUpdateResponse]
type zoneCustomNUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomNUpdateParams struct {
	ZoneMetadata ZoneMetadataParam `json:"zone_metadata,required"`
}

func (r ZoneCustomNUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ZoneMetadata)
}
