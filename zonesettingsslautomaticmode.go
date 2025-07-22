// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneSettingSslAutomaticModeService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSettingSslAutomaticModeService] method instead.
type ZoneSettingSslAutomaticModeService struct {
	Options []option.RequestOption
}

// NewZoneSettingSslAutomaticModeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingSslAutomaticModeService(opts ...option.RequestOption) (r *ZoneSettingSslAutomaticModeService) {
	r = &ZoneSettingSslAutomaticModeService{}
	r.Options = opts
	return
}

// If the system is enabled, the response will include next_scheduled_scan,
// representing the next time this zone will be scanned and the zone's ssl/tls
// encryption mode is potentially upgraded by the system. If the system is
// disabled, next_scheduled_scan will not be present in the response body.
func (r *ZoneSettingSslAutomaticModeService) Get(ctx context.Context, zoneTag interface{}, opts ...option.RequestOption) (res *CacheAutomaticUpgraderResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/settings/ssl_automatic_mode", zoneTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// The automatic system is enabled when this endpoint is hit with value in the
// request body is set to "auto", and disabled when the request body value is set
// to "custom".
func (r *ZoneSettingSslAutomaticModeService) Update(ctx context.Context, zoneTag interface{}, body ZoneSettingSslAutomaticModeUpdateParams, opts ...option.RequestOption) (res *CacheAutomaticUpgraderResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/settings/ssl_automatic_mode", zoneTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type CacheAutomaticUpgraderResponse struct {
	Result CacheAutomaticUpgraderResponseResult `json:"result"`
	JSON   cacheAutomaticUpgraderResponseJSON   `json:"-"`
	CacheAPIResponseSingle
}

// cacheAutomaticUpgraderResponseJSON contains the JSON metadata for the struct
// [CacheAutomaticUpgraderResponse]
type cacheAutomaticUpgraderResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheAutomaticUpgraderResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheAutomaticUpgraderResponseJSON) RawJSON() string {
	return r.raw
}

type CacheAutomaticUpgraderResponseResult struct {
	ID string `json:"id,required"`
	// Whether this setting can be updated or not.
	Editable bool `json:"editable,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Current setting of the automatic SSL/TLS.
	Value CacheAutomaticUpgraderResponseResultValue `json:"value,required"`
	// next time this zone will be scanned by the Automatic SSL/TLS.
	NextScheduledScan time.Time                                `json:"next_scheduled_scan,nullable" format:"date-time"`
	JSON              cacheAutomaticUpgraderResponseResultJSON `json:"-"`
}

// cacheAutomaticUpgraderResponseResultJSON contains the JSON metadata for the
// struct [CacheAutomaticUpgraderResponseResult]
type cacheAutomaticUpgraderResponseResultJSON struct {
	ID                apijson.Field
	Editable          apijson.Field
	ModifiedOn        apijson.Field
	Value             apijson.Field
	NextScheduledScan apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CacheAutomaticUpgraderResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheAutomaticUpgraderResponseResultJSON) RawJSON() string {
	return r.raw
}

// Current setting of the automatic SSL/TLS.
type CacheAutomaticUpgraderResponseResultValue string

const (
	CacheAutomaticUpgraderResponseResultValueAuto   CacheAutomaticUpgraderResponseResultValue = "auto"
	CacheAutomaticUpgraderResponseResultValueCustom CacheAutomaticUpgraderResponseResultValue = "custom"
)

func (r CacheAutomaticUpgraderResponseResultValue) IsKnown() bool {
	switch r {
	case CacheAutomaticUpgraderResponseResultValueAuto, CacheAutomaticUpgraderResponseResultValueCustom:
		return true
	}
	return false
}

type ZoneSettingSslAutomaticModeUpdateParams struct {
	// Controls enablement of Automatic SSL/TLS.
	Value param.Field[ZoneSettingSslAutomaticModeUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingSslAutomaticModeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls enablement of Automatic SSL/TLS.
type ZoneSettingSslAutomaticModeUpdateParamsValue string

const (
	ZoneSettingSslAutomaticModeUpdateParamsValueAuto   ZoneSettingSslAutomaticModeUpdateParamsValue = "auto"
	ZoneSettingSslAutomaticModeUpdateParamsValueCustom ZoneSettingSslAutomaticModeUpdateParamsValue = "custom"
)

func (r ZoneSettingSslAutomaticModeUpdateParamsValue) IsKnown() bool {
	switch r {
	case ZoneSettingSslAutomaticModeUpdateParamsValueAuto, ZoneSettingSslAutomaticModeUpdateParamsValueCustom:
		return true
	}
	return false
}
