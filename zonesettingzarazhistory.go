// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneSettingZarazHistoryService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSettingZarazHistoryService] method instead.
type ZoneSettingZarazHistoryService struct {
	Options []option.RequestOption
}

// NewZoneSettingZarazHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingZarazHistoryService(opts ...option.RequestOption) (r *ZoneSettingZarazHistoryService) {
	r = &ZoneSettingZarazHistoryService{}
	r.Options = opts
	return
}

// Lists a history of published Zaraz configuration records for a zone.
func (r *ZoneSettingZarazHistoryService) List(ctx context.Context, zoneID string, query ZoneSettingZarazHistoryListParams, opts ...option.RequestOption) (res *ZoneSettingZarazHistoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/history", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Restores a historical published Zaraz configuration by ID for a zone.
func (r *ZoneSettingZarazHistoryService) Restore(ctx context.Context, zoneID string, body ZoneSettingZarazHistoryRestoreParams, opts ...option.RequestOption) (res *ZarazConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/history", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Gets a history of published Zaraz configurations by ID(s) for a zone.
func (r *ZoneSettingZarazHistoryService) GetConfigs(ctx context.Context, zoneID string, query ZoneSettingZarazHistoryGetConfigsParams, opts ...option.RequestOption) (res *ZoneSettingZarazHistoryGetConfigsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/history/configs", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneSettingZarazHistoryListResponse struct {
	Result []ZoneSettingZarazHistoryListResponseResult `json:"result"`
	JSON   zoneSettingZarazHistoryListResponseJSON     `json:"-"`
	ZarazAPIResponseCommon
}

// zoneSettingZarazHistoryListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingZarazHistoryListResponse]
type zoneSettingZarazHistoryListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZarazHistoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingZarazHistoryListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingZarazHistoryListResponseResult struct {
	// ID of the configuration
	ID int64 `json:"id,required"`
	// Date and time the configuration was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Configuration description provided by the user who published this configuration
	Description string `json:"description,required"`
	// Date and time the configuration was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// Alpha-numeric ID of the account user who published the configuration
	UserID string                                        `json:"userId,required"`
	JSON   zoneSettingZarazHistoryListResponseResultJSON `json:"-"`
}

// zoneSettingZarazHistoryListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingZarazHistoryListResponseResult]
type zoneSettingZarazHistoryListResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	UpdatedAt   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZarazHistoryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingZarazHistoryListResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingZarazHistoryGetConfigsResponse struct {
	// Object where keys are numericc onfiguration IDs
	Result map[string]ZoneSettingZarazHistoryGetConfigsResponseResult `json:"result"`
	JSON   zoneSettingZarazHistoryGetConfigsResponseJSON              `json:"-"`
	ZarazAPIResponseCommon
}

// zoneSettingZarazHistoryGetConfigsResponseJSON contains the JSON metadata for the
// struct [ZoneSettingZarazHistoryGetConfigsResponse]
type zoneSettingZarazHistoryGetConfigsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZarazHistoryGetConfigsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingZarazHistoryGetConfigsResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingZarazHistoryGetConfigsResponseResult struct {
	// ID of the configuration
	ID int64 `json:"id,required"`
	// Zaraz configuration
	Config ZarazConfigReturn `json:"config,required"`
	// Date and time the configuration was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Date and time the configuration was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// Alpha-numeric ID of the account user who published the configuration
	UserID string                                              `json:"userId,required"`
	JSON   zoneSettingZarazHistoryGetConfigsResponseResultJSON `json:"-"`
}

// zoneSettingZarazHistoryGetConfigsResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingZarazHistoryGetConfigsResponseResult]
type zoneSettingZarazHistoryGetConfigsResponseResultJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZarazHistoryGetConfigsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingZarazHistoryGetConfigsResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingZarazHistoryListParams struct {
	// Maximum amount of results to list. Default value is 10.
	Limit param.Field[int64] `query:"limit"`
	// Ordinal number to start listing the results with. Default value is 0.
	Offset param.Field[int64] `query:"offset"`
	// The field to sort by. Default is updated_at.
	SortField param.Field[ZoneSettingZarazHistoryListParamsSortField] `query:"sortField"`
	// Sorting order. Default is DESC.
	SortOrder param.Field[ZoneSettingZarazHistoryListParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [ZoneSettingZarazHistoryListParams]'s query parameters as
// `url.Values`.
func (r ZoneSettingZarazHistoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The field to sort by. Default is updated_at.
type ZoneSettingZarazHistoryListParamsSortField string

const (
	ZoneSettingZarazHistoryListParamsSortFieldID          ZoneSettingZarazHistoryListParamsSortField = "id"
	ZoneSettingZarazHistoryListParamsSortFieldUserID      ZoneSettingZarazHistoryListParamsSortField = "user_id"
	ZoneSettingZarazHistoryListParamsSortFieldDescription ZoneSettingZarazHistoryListParamsSortField = "description"
	ZoneSettingZarazHistoryListParamsSortFieldCreatedAt   ZoneSettingZarazHistoryListParamsSortField = "created_at"
	ZoneSettingZarazHistoryListParamsSortFieldUpdatedAt   ZoneSettingZarazHistoryListParamsSortField = "updated_at"
)

func (r ZoneSettingZarazHistoryListParamsSortField) IsKnown() bool {
	switch r {
	case ZoneSettingZarazHistoryListParamsSortFieldID, ZoneSettingZarazHistoryListParamsSortFieldUserID, ZoneSettingZarazHistoryListParamsSortFieldDescription, ZoneSettingZarazHistoryListParamsSortFieldCreatedAt, ZoneSettingZarazHistoryListParamsSortFieldUpdatedAt:
		return true
	}
	return false
}

// Sorting order. Default is DESC.
type ZoneSettingZarazHistoryListParamsSortOrder string

const (
	ZoneSettingZarazHistoryListParamsSortOrderDesc ZoneSettingZarazHistoryListParamsSortOrder = "DESC"
	ZoneSettingZarazHistoryListParamsSortOrderAsc  ZoneSettingZarazHistoryListParamsSortOrder = "ASC"
)

func (r ZoneSettingZarazHistoryListParamsSortOrder) IsKnown() bool {
	switch r {
	case ZoneSettingZarazHistoryListParamsSortOrderDesc, ZoneSettingZarazHistoryListParamsSortOrderAsc:
		return true
	}
	return false
}

type ZoneSettingZarazHistoryRestoreParams struct {
	// ID of the Zaraz configuration to restore.
	Body int64 `json:"body,required"`
}

func (r ZoneSettingZarazHistoryRestoreParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneSettingZarazHistoryGetConfigsParams struct {
	// Comma separated list of Zaraz configuration IDs
	IDs param.Field[[]int64] `query:"ids,required"`
}

// URLQuery serializes [ZoneSettingZarazHistoryGetConfigsParams]'s query parameters
// as `url.Values`.
func (r ZoneSettingZarazHistoryGetConfigsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
