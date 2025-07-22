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

// ZoneLogControlRetentionFlagService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneLogControlRetentionFlagService] method instead.
type ZoneLogControlRetentionFlagService struct {
	Options []option.RequestOption
}

// NewZoneLogControlRetentionFlagService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneLogControlRetentionFlagService(opts ...option.RequestOption) (r *ZoneLogControlRetentionFlagService) {
	r = &ZoneLogControlRetentionFlagService{}
	r.Options = opts
	return
}

// Updates log retention flag for Logpull API.
func (r *ZoneLogControlRetentionFlagService) Update(ctx context.Context, zoneID string, body ZoneLogControlRetentionFlagUpdateParams, opts ...option.RequestOption) (res *FlagResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logs/control/retention/flag", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets log retention flag for Logpull API.
func (r *ZoneLogControlRetentionFlagService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *FlagResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logs/control/retention/flag", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Flag struct {
	// The log retention flag for Logpull API.
	Flag bool     `json:"flag"`
	JSON flagJSON `json:"-"`
}

// flagJSON contains the JSON metadata for the struct [Flag]
type flagJSON struct {
	Flag        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Flag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r flagJSON) RawJSON() string {
	return r.raw
}

type FlagParam struct {
	// The log retention flag for Logpull API.
	Flag param.Field[bool] `json:"flag"`
}

func (r FlagParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FlagResponseSingle struct {
	Result Flag                   `json:"result,nullable"`
	JSON   flagResponseSingleJSON `json:"-"`
	SingleResponseLogControl
}

// flagResponseSingleJSON contains the JSON metadata for the struct
// [FlagResponseSingle]
type flagResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FlagResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r flagResponseSingleJSON) RawJSON() string {
	return r.raw
}

type ZoneLogControlRetentionFlagUpdateParams struct {
	Flag FlagParam `json:"flag,required"`
}

func (r ZoneLogControlRetentionFlagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Flag)
}
