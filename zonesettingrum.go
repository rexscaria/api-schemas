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

// ZoneSettingRumService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSettingRumService] method instead.
type ZoneSettingRumService struct {
	Options []option.RequestOption
}

// NewZoneSettingRumService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingRumService(opts ...option.RequestOption) (r *ZoneSettingRumService) {
	r = &ZoneSettingRumService{}
	r.Options = opts
	return
}

// Retrieves RUM status for a zone.
func (r *ZoneSettingRumService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *RumSiteResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/rum", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Toggles RUM on/off for an existing zone
func (r *ZoneSettingRumService) Update(ctx context.Context, zoneID string, body ZoneSettingRumUpdateParams, opts ...option.RequestOption) (res *RumSiteResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/rum", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type RumSiteResponseSingle struct {
	Result RumSiteResponseSingleResult `json:"result"`
	JSON   rumSiteResponseSingleJSON   `json:"-"`
	Common
}

// rumSiteResponseSingleJSON contains the JSON metadata for the struct
// [RumSiteResponseSingle]
type rumSiteResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rumSiteResponseSingleJSON) RawJSON() string {
	return r.raw
}

type RumSiteResponseSingleResult struct {
	ID       string `json:"id"`
	Editable bool   `json:"editable"`
	// Current state of RUM. Returns On, Off, or Manual
	Value string                          `json:"value"`
	JSON  rumSiteResponseSingleResultJSON `json:"-"`
}

// rumSiteResponseSingleResultJSON contains the JSON metadata for the struct
// [RumSiteResponseSingleResult]
type rumSiteResponseSingleResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rumSiteResponseSingleResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingRumUpdateParams struct {
	// Value can either be On or Off
	Value param.Field[string] `json:"value"`
}

func (r ZoneSettingRumUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
