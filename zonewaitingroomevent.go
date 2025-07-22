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

// ZoneWaitingRoomEventService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneWaitingRoomEventService] method instead.
type ZoneWaitingRoomEventService struct {
	Options []option.RequestOption
}

// NewZoneWaitingRoomEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneWaitingRoomEventService(opts ...option.RequestOption) (r *ZoneWaitingRoomEventService) {
	r = &ZoneWaitingRoomEventService{}
	r.Options = opts
	return
}

// Only available for the Waiting Room Advanced subscription. Creates an event for
// a waiting room. An event takes place during a specified period of time,
// temporarily changing the behavior of a waiting room. While the event is active,
// some of the properties in the event's configuration may either override or
// inherit from the waiting room's configuration. Note that events cannot overlap
// with each other, so only one event can be active at a time.
func (r *ZoneWaitingRoomEventService) New(ctx context.Context, zoneID string, waitingRoomID string, body ZoneWaitingRoomEventNewParams, opts ...option.RequestOption) (res *EventResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/events", zoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single configured event for a waiting room.
func (r *ZoneWaitingRoomEventService) Get(ctx context.Context, zoneID string, waitingRoomID string, eventID string, opts ...option.RequestOption) (res *EventResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/events/%s", zoneID, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured event for a waiting room.
func (r *ZoneWaitingRoomEventService) Update(ctx context.Context, zoneID string, waitingRoomID string, eventID string, body ZoneWaitingRoomEventUpdateParams, opts ...option.RequestOption) (res *EventResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/events/%s", zoneID, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists events for a waiting room.
func (r *ZoneWaitingRoomEventService) List(ctx context.Context, zoneID string, waitingRoomID string, query ZoneWaitingRoomEventListParams, opts ...option.RequestOption) (res *ZoneWaitingRoomEventListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/events", zoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an event for a waiting room.
func (r *ZoneWaitingRoomEventService) Delete(ctx context.Context, zoneID string, waitingRoomID string, eventID string, body ZoneWaitingRoomEventDeleteParams, opts ...option.RequestOption) (res *ZoneWaitingRoomEventDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/events/%s", zoneID, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Patches a configured event for a waiting room.
func (r *ZoneWaitingRoomEventService) Patch(ctx context.Context, zoneID string, waitingRoomID string, eventID string, body ZoneWaitingRoomEventPatchParams, opts ...option.RequestOption) (res *EventResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/events/%s", zoneID, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Previews an event's configuration as if it was active. Inherited fields from the
// waiting room will be displayed with their current values.
func (r *ZoneWaitingRoomEventService) Preview(ctx context.Context, zoneID string, waitingRoomID string, eventID string, opts ...option.RequestOption) (res *ZoneWaitingRoomEventPreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/events/%s/details", zoneID, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EventResponse struct {
	Result EventResult       `json:"result"`
	JSON   eventResponseJSON `json:"-"`
	APIResponseSingle
}

// eventResponseJSON contains the JSON metadata for the struct [EventResponse]
type eventResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventResponseJSON) RawJSON() string {
	return r.raw
}

type EventResult struct {
	ID        string    `json:"id"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// If set, the event will override the waiting room's `custom_page_html` property
	// while it is active. If null, the event will inherit it.
	CustomPageHTML string `json:"custom_page_html,nullable"`
	// A note that you can use to add more details about the event.
	Description string `json:"description"`
	// If set, the event will override the waiting room's `disable_session_renewal`
	// property while it is active. If null, the event will inherit it.
	DisableSessionRenewal bool `json:"disable_session_renewal,nullable"`
	// An ISO 8601 timestamp that marks the end of the event.
	EventEndTime string `json:"event_end_time"`
	// An ISO 8601 timestamp that marks the start of the event. At this time, queued
	// users will be processed with the event's configuration. The start time must be
	// at least one minute before `event_end_time`.
	EventStartTime string    `json:"event_start_time"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	// A unique name to identify the event. Only alphanumeric characters, hyphens and
	// underscores are allowed.
	Name string `json:"name"`
	// If set, the event will override the waiting room's `new_users_per_minute`
	// property while it is active. If null, the event will inherit it. This can only
	// be set if the event's `total_active_users` property is also set.
	NewUsersPerMinute int64 `json:"new_users_per_minute,nullable"`
	// An ISO 8601 timestamp that marks when to begin queueing all users before the
	// event starts. The prequeue must start at least five minutes before
	// `event_start_time`.
	PrequeueStartTime string `json:"prequeue_start_time,nullable"`
	// If set, the event will override the waiting room's `queueing_method` property
	// while it is active. If null, the event will inherit it.
	QueueingMethod string `json:"queueing_method,nullable"`
	// If set, the event will override the waiting room's `session_duration` property
	// while it is active. If null, the event will inherit it.
	SessionDuration int64 `json:"session_duration,nullable"`
	// If enabled, users in the prequeue will be shuffled randomly at the
	// `event_start_time`. Requires that `prequeue_start_time` is not null. This is
	// useful for situations when many users will join the event prequeue at the same
	// time and you want to shuffle them to ensure fairness. Naturally, it makes the
	// most sense to enable this feature when the `queueing_method` during the event
	// respects ordering such as **fifo**, or else the shuffling may be unnecessary.
	ShuffleAtEventStart bool `json:"shuffle_at_event_start"`
	// Suspends or allows an event. If set to `true`, the event is ignored and traffic
	// will be handled based on the waiting room configuration.
	Suspended bool `json:"suspended"`
	// If set, the event will override the waiting room's `total_active_users` property
	// while it is active. If null, the event will inherit it. This can only be set if
	// the event's `new_users_per_minute` property is also set.
	TotalActiveUsers int64 `json:"total_active_users,nullable"`
	// If set, the event will override the waiting room's `turnstile_action` property
	// while it is active. If null, the event will inherit it.
	TurnstileAction TurnstileEventAction `json:"turnstile_action,nullable"`
	// If set, the event will override the waiting room's `turnstile_mode` property
	// while it is active. If null, the event will inherit it.
	TurnstileMode TurnstileEventMode `json:"turnstile_mode,nullable"`
	JSON          eventResultJSON    `json:"-"`
}

// eventResultJSON contains the JSON metadata for the struct [EventResult]
type eventResultJSON struct {
	ID                    apijson.Field
	CreatedOn             apijson.Field
	CustomPageHTML        apijson.Field
	Description           apijson.Field
	DisableSessionRenewal apijson.Field
	EventEndTime          apijson.Field
	EventStartTime        apijson.Field
	ModifiedOn            apijson.Field
	Name                  apijson.Field
	NewUsersPerMinute     apijson.Field
	PrequeueStartTime     apijson.Field
	QueueingMethod        apijson.Field
	SessionDuration       apijson.Field
	ShuffleAtEventStart   apijson.Field
	Suspended             apijson.Field
	TotalActiveUsers      apijson.Field
	TurnstileAction       apijson.Field
	TurnstileMode         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *EventResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventResultJSON) RawJSON() string {
	return r.raw
}

type QueryEventParam struct {
	// An ISO 8601 timestamp that marks the end of the event.
	EventEndTime param.Field[string] `json:"event_end_time,required"`
	// An ISO 8601 timestamp that marks the start of the event. At this time, queued
	// users will be processed with the event's configuration. The start time must be
	// at least one minute before `event_end_time`.
	EventStartTime param.Field[string] `json:"event_start_time,required"`
	// A unique name to identify the event. Only alphanumeric characters, hyphens and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// If set, the event will override the waiting room's `custom_page_html` property
	// while it is active. If null, the event will inherit it.
	CustomPageHTML param.Field[string] `json:"custom_page_html"`
	// A note that you can use to add more details about the event.
	Description param.Field[string] `json:"description"`
	// If set, the event will override the waiting room's `disable_session_renewal`
	// property while it is active. If null, the event will inherit it.
	DisableSessionRenewal param.Field[bool] `json:"disable_session_renewal"`
	// If set, the event will override the waiting room's `new_users_per_minute`
	// property while it is active. If null, the event will inherit it. This can only
	// be set if the event's `total_active_users` property is also set.
	NewUsersPerMinute param.Field[int64] `json:"new_users_per_minute"`
	// An ISO 8601 timestamp that marks when to begin queueing all users before the
	// event starts. The prequeue must start at least five minutes before
	// `event_start_time`.
	PrequeueStartTime param.Field[string] `json:"prequeue_start_time"`
	// If set, the event will override the waiting room's `queueing_method` property
	// while it is active. If null, the event will inherit it.
	QueueingMethod param.Field[string] `json:"queueing_method"`
	// If set, the event will override the waiting room's `session_duration` property
	// while it is active. If null, the event will inherit it.
	SessionDuration param.Field[int64] `json:"session_duration"`
	// If enabled, users in the prequeue will be shuffled randomly at the
	// `event_start_time`. Requires that `prequeue_start_time` is not null. This is
	// useful for situations when many users will join the event prequeue at the same
	// time and you want to shuffle them to ensure fairness. Naturally, it makes the
	// most sense to enable this feature when the `queueing_method` during the event
	// respects ordering such as **fifo**, or else the shuffling may be unnecessary.
	ShuffleAtEventStart param.Field[bool] `json:"shuffle_at_event_start"`
	// Suspends or allows an event. If set to `true`, the event is ignored and traffic
	// will be handled based on the waiting room configuration.
	Suspended param.Field[bool] `json:"suspended"`
	// If set, the event will override the waiting room's `total_active_users` property
	// while it is active. If null, the event will inherit it. This can only be set if
	// the event's `new_users_per_minute` property is also set.
	TotalActiveUsers param.Field[int64] `json:"total_active_users"`
	// If set, the event will override the waiting room's `turnstile_action` property
	// while it is active. If null, the event will inherit it.
	TurnstileAction param.Field[TurnstileEventAction] `json:"turnstile_action"`
	// If set, the event will override the waiting room's `turnstile_mode` property
	// while it is active. If null, the event will inherit it.
	TurnstileMode param.Field[TurnstileEventMode] `json:"turnstile_mode"`
}

func (r QueryEventParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the event will override the waiting room's `turnstile_action` property
// while it is active. If null, the event will inherit it.
type TurnstileEventAction string

const (
	TurnstileEventActionLog           TurnstileEventAction = "log"
	TurnstileEventActionInfiniteQueue TurnstileEventAction = "infinite_queue"
)

func (r TurnstileEventAction) IsKnown() bool {
	switch r {
	case TurnstileEventActionLog, TurnstileEventActionInfiniteQueue:
		return true
	}
	return false
}

// If set, the event will override the waiting room's `turnstile_mode` property
// while it is active. If null, the event will inherit it.
type TurnstileEventMode string

const (
	TurnstileEventModeOff                   TurnstileEventMode = "off"
	TurnstileEventModeInvisible             TurnstileEventMode = "invisible"
	TurnstileEventModeVisibleNonInteractive TurnstileEventMode = "visible_non_interactive"
	TurnstileEventModeVisibleManaged        TurnstileEventMode = "visible_managed"
)

func (r TurnstileEventMode) IsKnown() bool {
	switch r {
	case TurnstileEventModeOff, TurnstileEventModeInvisible, TurnstileEventModeVisibleNonInteractive, TurnstileEventModeVisibleManaged:
		return true
	}
	return false
}

type ZoneWaitingRoomEventListResponse struct {
	Result []EventResult                        `json:"result"`
	JSON   zoneWaitingRoomEventListResponseJSON `json:"-"`
	WaitingroomAPIResponseCollection
}

// zoneWaitingRoomEventListResponseJSON contains the JSON metadata for the struct
// [ZoneWaitingRoomEventListResponse]
type zoneWaitingRoomEventListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWaitingRoomEventListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWaitingRoomEventDeleteResponse struct {
	Result ZoneWaitingRoomEventDeleteResponseResult `json:"result"`
	JSON   zoneWaitingRoomEventDeleteResponseJSON   `json:"-"`
	APIResponseSingle
}

// zoneWaitingRoomEventDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneWaitingRoomEventDeleteResponse]
type zoneWaitingRoomEventDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomEventDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWaitingRoomEventDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWaitingRoomEventDeleteResponseResult struct {
	ID   string                                       `json:"id"`
	JSON zoneWaitingRoomEventDeleteResponseResultJSON `json:"-"`
}

// zoneWaitingRoomEventDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneWaitingRoomEventDeleteResponseResult]
type zoneWaitingRoomEventDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomEventDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWaitingRoomEventDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneWaitingRoomEventPreviewResponse struct {
	Result ZoneWaitingRoomEventPreviewResponseResult `json:"result"`
	JSON   zoneWaitingRoomEventPreviewResponseJSON   `json:"-"`
	APIResponseSingle
}

// zoneWaitingRoomEventPreviewResponseJSON contains the JSON metadata for the
// struct [ZoneWaitingRoomEventPreviewResponse]
type zoneWaitingRoomEventPreviewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomEventPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWaitingRoomEventPreviewResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWaitingRoomEventPreviewResponseResult struct {
	ID             string    `json:"id"`
	CreatedOn      time.Time `json:"created_on" format:"date-time"`
	CustomPageHTML string    `json:"custom_page_html"`
	// A note that you can use to add more details about the event.
	Description           string `json:"description"`
	DisableSessionRenewal bool   `json:"disable_session_renewal"`
	// An ISO 8601 timestamp that marks the end of the event.
	EventEndTime string `json:"event_end_time"`
	// An ISO 8601 timestamp that marks the start of the event. At this time, queued
	// users will be processed with the event's configuration. The start time must be
	// at least one minute before `event_end_time`.
	EventStartTime string    `json:"event_start_time"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	// A unique name to identify the event. Only alphanumeric characters, hyphens and
	// underscores are allowed.
	Name              string `json:"name"`
	NewUsersPerMinute int64  `json:"new_users_per_minute"`
	// An ISO 8601 timestamp that marks when to begin queueing all users before the
	// event starts. The prequeue must start at least five minutes before
	// `event_start_time`.
	PrequeueStartTime string `json:"prequeue_start_time,nullable"`
	QueueingMethod    string `json:"queueing_method"`
	SessionDuration   int64  `json:"session_duration"`
	// If enabled, users in the prequeue will be shuffled randomly at the
	// `event_start_time`. Requires that `prequeue_start_time` is not null. This is
	// useful for situations when many users will join the event prequeue at the same
	// time and you want to shuffle them to ensure fairness. Naturally, it makes the
	// most sense to enable this feature when the `queueing_method` during the event
	// respects ordering such as **fifo**, or else the shuffling may be unnecessary.
	ShuffleAtEventStart bool `json:"shuffle_at_event_start"`
	// Suspends or allows an event. If set to `true`, the event is ignored and traffic
	// will be handled based on the waiting room configuration.
	Suspended        bool                                          `json:"suspended"`
	TotalActiveUsers int64                                         `json:"total_active_users"`
	JSON             zoneWaitingRoomEventPreviewResponseResultJSON `json:"-"`
}

// zoneWaitingRoomEventPreviewResponseResultJSON contains the JSON metadata for the
// struct [ZoneWaitingRoomEventPreviewResponseResult]
type zoneWaitingRoomEventPreviewResponseResultJSON struct {
	ID                    apijson.Field
	CreatedOn             apijson.Field
	CustomPageHTML        apijson.Field
	Description           apijson.Field
	DisableSessionRenewal apijson.Field
	EventEndTime          apijson.Field
	EventStartTime        apijson.Field
	ModifiedOn            apijson.Field
	Name                  apijson.Field
	NewUsersPerMinute     apijson.Field
	PrequeueStartTime     apijson.Field
	QueueingMethod        apijson.Field
	SessionDuration       apijson.Field
	ShuffleAtEventStart   apijson.Field
	Suspended             apijson.Field
	TotalActiveUsers      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ZoneWaitingRoomEventPreviewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWaitingRoomEventPreviewResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneWaitingRoomEventNewParams struct {
	QueryEvent QueryEventParam `json:"query_event,required"`
}

func (r ZoneWaitingRoomEventNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.QueryEvent)
}

type ZoneWaitingRoomEventUpdateParams struct {
	QueryEvent QueryEventParam `json:"query_event,required"`
}

func (r ZoneWaitingRoomEventUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.QueryEvent)
}

type ZoneWaitingRoomEventListParams struct {
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page. Must be a multiple of 5.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ZoneWaitingRoomEventListParams]'s query parameters as
// `url.Values`.
func (r ZoneWaitingRoomEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneWaitingRoomEventDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneWaitingRoomEventDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneWaitingRoomEventPatchParams struct {
	QueryEvent QueryEventParam `json:"query_event,required"`
}

func (r ZoneWaitingRoomEventPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.QueryEvent)
}
