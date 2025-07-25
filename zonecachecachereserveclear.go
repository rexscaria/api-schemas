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

type MessagesCacheRulesItem struct {
	Code             int64                        `json:"code,required"`
	Message          string                       `json:"message,required"`
	DocumentationURL string                       `json:"documentation_url"`
	Source           MessagesCacheRulesItemSource `json:"source"`
	JSON             messagesCacheRulesItemJSON   `json:"-"`
}

// messagesCacheRulesItemJSON contains the JSON metadata for the struct
// [MessagesCacheRulesItem]
type messagesCacheRulesItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessagesCacheRulesItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesCacheRulesItemJSON) RawJSON() string {
	return r.raw
}

type MessagesCacheRulesItemSource struct {
	Pointer string                           `json:"pointer"`
	JSON    messagesCacheRulesItemSourceJSON `json:"-"`
}

// messagesCacheRulesItemSourceJSON contains the JSON metadata for the struct
// [MessagesCacheRulesItemSource]
type messagesCacheRulesItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesCacheRulesItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesCacheRulesItemSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneCacheCacheReserveClearGetResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneCacheCacheReserveClearGetResponseSuccess `json:"success,required"`
	// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
	// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
	// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
	// that you cannot undo or cancel this operation.
	Result ZoneCacheCacheReserveClearGetResponseResult `json:"result"`
	JSON   zoneCacheCacheReserveClearGetResponseJSON   `json:"-"`
}

// zoneCacheCacheReserveClearGetResponseJSON contains the JSON metadata for the
// struct [ZoneCacheCacheReserveClearGetResponse]
type zoneCacheCacheReserveClearGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheCacheReserveClearGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheCacheReserveClearGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCacheCacheReserveClearGetResponseSuccess bool

const (
	ZoneCacheCacheReserveClearGetResponseSuccessTrue ZoneCacheCacheReserveClearGetResponseSuccess = true
)

func (r ZoneCacheCacheReserveClearGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCacheCacheReserveClearGetResponseSuccessTrue:
		return true
	}
	return false
}

// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
// that you cannot undo or cancel this operation.
type ZoneCacheCacheReserveClearGetResponseResult struct {
	// ID of the zone setting.
	ID ZoneCacheCacheReserveClearGetResponseResultID `json:"id,required"`
	// The time that the latest Cache Reserve Clear operation started.
	StartTs time.Time `json:"start_ts,required" format:"date-time"`
	// The current state of the Cache Reserve Clear operation.
	State ZoneCacheCacheReserveClearGetResponseResultState `json:"state,required"`
	// The time that the latest Cache Reserve Clear operation completed.
	EndTs time.Time `json:"end_ts" format:"date-time"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCacheCacheReserveClearGetResponseResultJSON `json:"-"`
}

// zoneCacheCacheReserveClearGetResponseResultJSON contains the JSON metadata for
// the struct [ZoneCacheCacheReserveClearGetResponseResult]
type zoneCacheCacheReserveClearGetResponseResultJSON struct {
	ID          apijson.Field
	StartTs     apijson.Field
	State       apijson.Field
	EndTs       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheCacheReserveClearGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheCacheReserveClearGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneCacheCacheReserveClearGetResponseResultID string

const (
	ZoneCacheCacheReserveClearGetResponseResultIDCacheReserveClear ZoneCacheCacheReserveClearGetResponseResultID = "cache_reserve_clear"
)

func (r ZoneCacheCacheReserveClearGetResponseResultID) IsKnown() bool {
	switch r {
	case ZoneCacheCacheReserveClearGetResponseResultIDCacheReserveClear:
		return true
	}
	return false
}

// The current state of the Cache Reserve Clear operation.
type ZoneCacheCacheReserveClearGetResponseResultState string

const (
	ZoneCacheCacheReserveClearGetResponseResultStateInProgress ZoneCacheCacheReserveClearGetResponseResultState = "In-progress"
	ZoneCacheCacheReserveClearGetResponseResultStateCompleted  ZoneCacheCacheReserveClearGetResponseResultState = "Completed"
)

func (r ZoneCacheCacheReserveClearGetResponseResultState) IsKnown() bool {
	switch r {
	case ZoneCacheCacheReserveClearGetResponseResultStateInProgress, ZoneCacheCacheReserveClearGetResponseResultStateCompleted:
		return true
	}
	return false
}

type ZoneCacheCacheReserveClearStartResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneCacheCacheReserveClearStartResponseSuccess `json:"success,required"`
	// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
	// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
	// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
	// that you cannot undo or cancel this operation.
	Result ZoneCacheCacheReserveClearStartResponseResult `json:"result"`
	JSON   zoneCacheCacheReserveClearStartResponseJSON   `json:"-"`
}

// zoneCacheCacheReserveClearStartResponseJSON contains the JSON metadata for the
// struct [ZoneCacheCacheReserveClearStartResponse]
type zoneCacheCacheReserveClearStartResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheCacheReserveClearStartResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheCacheReserveClearStartResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCacheCacheReserveClearStartResponseSuccess bool

const (
	ZoneCacheCacheReserveClearStartResponseSuccessTrue ZoneCacheCacheReserveClearStartResponseSuccess = true
)

func (r ZoneCacheCacheReserveClearStartResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCacheCacheReserveClearStartResponseSuccessTrue:
		return true
	}
	return false
}

// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
// that you cannot undo or cancel this operation.
type ZoneCacheCacheReserveClearStartResponseResult struct {
	// ID of the zone setting.
	ID ZoneCacheCacheReserveClearStartResponseResultID `json:"id,required"`
	// The time that the latest Cache Reserve Clear operation started.
	StartTs time.Time `json:"start_ts,required" format:"date-time"`
	// The current state of the Cache Reserve Clear operation.
	State ZoneCacheCacheReserveClearStartResponseResultState `json:"state,required"`
	// The time that the latest Cache Reserve Clear operation completed.
	EndTs time.Time `json:"end_ts" format:"date-time"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCacheCacheReserveClearStartResponseResultJSON `json:"-"`
}

// zoneCacheCacheReserveClearStartResponseResultJSON contains the JSON metadata for
// the struct [ZoneCacheCacheReserveClearStartResponseResult]
type zoneCacheCacheReserveClearStartResponseResultJSON struct {
	ID          apijson.Field
	StartTs     apijson.Field
	State       apijson.Field
	EndTs       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheCacheReserveClearStartResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheCacheReserveClearStartResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneCacheCacheReserveClearStartResponseResultID string

const (
	ZoneCacheCacheReserveClearStartResponseResultIDCacheReserveClear ZoneCacheCacheReserveClearStartResponseResultID = "cache_reserve_clear"
)

func (r ZoneCacheCacheReserveClearStartResponseResultID) IsKnown() bool {
	switch r {
	case ZoneCacheCacheReserveClearStartResponseResultIDCacheReserveClear:
		return true
	}
	return false
}

// The current state of the Cache Reserve Clear operation.
type ZoneCacheCacheReserveClearStartResponseResultState string

const (
	ZoneCacheCacheReserveClearStartResponseResultStateInProgress ZoneCacheCacheReserveClearStartResponseResultState = "In-progress"
	ZoneCacheCacheReserveClearStartResponseResultStateCompleted  ZoneCacheCacheReserveClearStartResponseResultState = "Completed"
)

func (r ZoneCacheCacheReserveClearStartResponseResultState) IsKnown() bool {
	switch r {
	case ZoneCacheCacheReserveClearStartResponseResultStateInProgress, ZoneCacheCacheReserveClearStartResponseResultStateCompleted:
		return true
	}
	return false
}

type ZoneCacheCacheReserveClearStartParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneCacheCacheReserveClearStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
