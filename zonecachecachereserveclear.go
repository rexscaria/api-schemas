// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneCacheCacheReserveClearService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCacheCacheReserveClearService] method instead.
type ZoneCacheCacheReserveClearService struct {
	Options []option.RequestOption
}

// NewZoneCacheCacheReserveClearService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneCacheCacheReserveClearService(opts ...option.RequestOption) (r *ZoneCacheCacheReserveClearService) {
	r = &ZoneCacheCacheReserveClearService{}
	r.Options = opts
	return
}

// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
// that you cannot undo or cancel this operation.
func (r *ZoneCacheCacheReserveClearService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneCacheCacheReserveClearGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/cache_reserve_clear", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
// that you cannot undo or cancel this operation.
func (r *ZoneCacheCacheReserveClearService) Start(ctx context.Context, zoneID string, body ZoneCacheCacheReserveClearStartParams, opts ...option.RequestOption) (res *ZoneCacheCacheReserveClearStartResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/cache_reserve_clear", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type APIResponseCacheRules struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseCacheRulesSuccess `json:"success,required"`
	JSON    apiResponseCacheRulesJSON    `json:"-"`
}

// apiResponseCacheRulesJSON contains the JSON metadata for the struct
// [APIResponseCacheRules]
type apiResponseCacheRulesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCacheRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCacheRulesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseCacheRulesSuccess bool

const (
	APIResponseCacheRulesSuccessTrue APIResponseCacheRulesSuccess = true
)

func (r APIResponseCacheRulesSuccess) IsKnown() bool {
	switch r {
	case APIResponseCacheRulesSuccessTrue:
		return true
	}
	return false
}

type MessagesCacheRulesItem struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    messagesCacheRulesItemJSON `json:"-"`
}

// messagesCacheRulesItemJSON contains the JSON metadata for the struct
// [MessagesCacheRulesItem]
type messagesCacheRulesItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesCacheRulesItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesCacheRulesItemJSON) RawJSON() string {
	return r.raw
}

type ResponseValueClear struct {
	Result ResponseValueClearResult `json:"result"`
	JSON   responseValueClearJSON   `json:"-"`
}

// responseValueClearJSON contains the JSON metadata for the struct
// [ResponseValueClear]
type responseValueClearJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseValueClear) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseValueClearJSON) RawJSON() string {
	return r.raw
}

type ResponseValueClearResult struct {
	// The time that the latest Cache Reserve Clear operation started.
	StartTs time.Time `json:"start_ts,required" format:"date-time"`
	// The current state of the Cache Reserve Clear operation.
	State ResponseValueClearResultState `json:"state,required"`
	// ID of the zone setting.
	ID ResponseValueClearResultID `json:"id"`
	// The time that the latest Cache Reserve Clear operation completed.
	EndTs time.Time                    `json:"end_ts" format:"date-time"`
	JSON  responseValueClearResultJSON `json:"-"`
	BaseCacheRule
}

// responseValueClearResultJSON contains the JSON metadata for the struct
// [ResponseValueClearResult]
type responseValueClearResultJSON struct {
	StartTs     apijson.Field
	State       apijson.Field
	ID          apijson.Field
	EndTs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseValueClearResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseValueClearResultJSON) RawJSON() string {
	return r.raw
}

// The current state of the Cache Reserve Clear operation.
type ResponseValueClearResultState string

const (
	ResponseValueClearResultStateInProgress ResponseValueClearResultState = "In-progress"
	ResponseValueClearResultStateCompleted  ResponseValueClearResultState = "Completed"
)

func (r ResponseValueClearResultState) IsKnown() bool {
	switch r {
	case ResponseValueClearResultStateInProgress, ResponseValueClearResultStateCompleted:
		return true
	}
	return false
}

// ID of the zone setting.
type ResponseValueClearResultID string

const (
	ResponseValueClearResultIDCacheReserveClear ResponseValueClearResultID = "cache_reserve_clear"
)

func (r ResponseValueClearResultID) IsKnown() bool {
	switch r {
	case ResponseValueClearResultIDCacheReserveClear:
		return true
	}
	return false
}

type ZoneCacheCacheReserveClearGetResponse struct {
	JSON zoneCacheCacheReserveClearGetResponseJSON `json:"-"`
	APIResponseCacheRules
	ResponseValueClear
}

// zoneCacheCacheReserveClearGetResponseJSON contains the JSON metadata for the
// struct [ZoneCacheCacheReserveClearGetResponse]
type zoneCacheCacheReserveClearGetResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheCacheReserveClearGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheCacheReserveClearGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCacheCacheReserveClearStartResponse struct {
	JSON zoneCacheCacheReserveClearStartResponseJSON `json:"-"`
	APIResponseCacheRules
	ResponseValueClear
}

// zoneCacheCacheReserveClearStartResponseJSON contains the JSON metadata for the
// struct [ZoneCacheCacheReserveClearStartResponse]
type zoneCacheCacheReserveClearStartResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheCacheReserveClearStartResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheCacheReserveClearStartResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCacheCacheReserveClearStartParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneCacheCacheReserveClearStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
