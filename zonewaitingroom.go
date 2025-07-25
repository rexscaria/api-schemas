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

// ZoneWaitingRoomService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneWaitingRoomService] method instead.
type ZoneWaitingRoomService struct {
	Options  []option.RequestOption
	Events   *ZoneWaitingRoomEventService
	Rules    *ZoneWaitingRoomRuleService
	Settings *ZoneWaitingRoomSettingService
}

// NewZoneWaitingRoomService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneWaitingRoomService(opts ...option.RequestOption) (r *ZoneWaitingRoomService) {
	r = &ZoneWaitingRoomService{}
	r.Options = opts
	r.Events = NewZoneWaitingRoomEventService(opts...)
	r.Rules = NewZoneWaitingRoomRuleService(opts...)
	r.Settings = NewZoneWaitingRoomSettingService(opts...)
	return
}

// Creates a new waiting room.
func (r *ZoneWaitingRoomService) New(ctx context.Context, zoneID string, body ZoneWaitingRoomNewParams, opts ...option.RequestOption) (res *SingleResponseWaitingRoom, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single configured waiting room.
func (r *ZoneWaitingRoomService) Get(ctx context.Context, zoneID string, waitingRoomID string, opts ...option.RequestOption) (res *SingleResponseWaitingRoom, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s", zoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured waiting room.
func (r *ZoneWaitingRoomService) Update(ctx context.Context, zoneID string, waitingRoomID string, body ZoneWaitingRoomUpdateParams, opts ...option.RequestOption) (res *SingleResponseWaitingRoom, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s", zoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a waiting room.
func (r *ZoneWaitingRoomService) Delete(ctx context.Context, zoneID string, waitingRoomID string, opts ...option.RequestOption) (res *ZoneWaitingRoomDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s", zoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Patches a configured waiting room.
func (r *ZoneWaitingRoomService) Patch(ctx context.Context, zoneID string, waitingRoomID string, body ZoneWaitingRoomPatchParams, opts ...option.RequestOption) (res *SingleResponseWaitingRoom, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s", zoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Creates a waiting room page preview. Upload a custom waiting room page for
// preview. You will receive a preview URL in the form
// `http://waitingrooms.dev/preview/<uuid>`. You can use the following query
// parameters to change the state of the preview:
//
//  1. `force_queue`: Boolean indicating if all users will be queued in the waiting
//     room and no one will be let into the origin website (also known as queueAll).
//  2. `queue_is_full`: Boolean indicating if the waiting room's queue is currently
//     full and not accepting new users at the moment.
//  3. `queueing_method`: The queueing method currently used by the waiting room.
//     - **fifo** indicates a FIFO queue.
//     - **random** indicates a Random queue.
//     - **passthrough** indicates a Passthrough queue. Keep in mind that the
//     waiting room page will only be displayed if `force_queue=true` or
//     `event=prequeueing` — for other cases the request will pass through to the
//     origin. For our preview, this will be a fake origin website returning
//     \"Welcome\".
//     - **reject** indicates a Reject queue.
//  4. `event`: Used to preview a waiting room event.
//     - **none** indicates no event is occurring.
//     - **prequeueing** indicates that an event is prequeueing (between
//     `prequeue_start_time` and `event_start_time`).
//     - **started** indicates that an event has started (between `event_start_time`
//     and `event_end_time`).
//  5. `shuffle_at_event_start`: Boolean indicating if the event will shuffle users
//     in the prequeue when it starts. This can only be set to **true** if an event
//     is active (`event` is not **none**).
//
// For example, you can make a request to
// `http://waitingrooms.dev/preview/<uuid>?force_queue=false&queue_is_full=false&queueing_method=random&event=started&shuffle_at_event_start=true` 6.
// `waitTime`: Non-zero, positive integer indicating the estimated wait time in
// minutes. The default value is 10 minutes.
//
// For example, you can make a request to
// `http://waitingrooms.dev/preview/<uuid>?waitTime=50` to configure the estimated
// wait time as 50 minutes.
func (r *ZoneWaitingRoomService) Preview(ctx context.Context, zoneID string, body ZoneWaitingRoomPreviewParams, opts ...option.RequestOption) (res *ZoneWaitingRoomPreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/preview", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the status of a configured waiting room. Response fields include:
//
//  1. `status`: String indicating the status of the waiting room. The possible
//     status are:
//     - **not_queueing** indicates that the configured thresholds have not been met
//     and all users are going through to the origin.
//     - **queueing** indicates that the thresholds have been met and some users are
//     held in the waiting room.
//     - **event_prequeueing** indicates that an event is active and is currently
//     prequeueing users before it starts.
//     - **suspended** indicates that the room is suspended.
//  2. `event_id`: String of the current event's `id` if an event is active,
//     otherwise an empty string.
//  3. `estimated_queued_users`: Integer of the estimated number of users currently
//     waiting in the queue.
//  4. `estimated_total_active_users`: Integer of the estimated number of users
//     currently active on the origin.
//  5. `max_estimated_time_minutes`: Integer of the maximum estimated time currently
//     presented to the users.
func (r *ZoneWaitingRoomService) Status(ctx context.Context, zoneID string, waitingRoomID string, opts ...option.RequestOption) (res *ZoneWaitingRoomStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/status", zoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AdditionalRoute struct {
	// The hostname to which this waiting room will be applied (no wildcards). The
	// hostname must be the primary domain, subdomain, or custom hostname (if using SSL
	// for SaaS) of this zone. Please do not include the scheme (http:// or https://).
	Host string `json:"host"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path string              `json:"path"`
	JSON additionalRouteJSON `json:"-"`
}

// additionalRouteJSON contains the JSON metadata for the struct [AdditionalRoute]
type additionalRouteJSON struct {
	Host        apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdditionalRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r additionalRouteJSON) RawJSON() string {
	return r.raw
}

type AdditionalRouteParam struct {
	// The hostname to which this waiting room will be applied (no wildcards). The
	// hostname must be the primary domain, subdomain, or custom hostname (if using SSL
	// for SaaS) of this zone. Please do not include the scheme (http:// or https://).
	Host param.Field[string] `json:"host"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path param.Field[string] `json:"path"`
}

func (r AdditionalRouteParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures cookie attributes for the waiting room cookie. This encrypted cookie
// stores a user's status in the waiting room, such as queue position.
type CookieAttributes struct {
	// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
	// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
	// Note that when using value `none`, the secure attribute cannot be set to
	// `never`.
	Samesite CookieAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on the waiting room cookie. Value `always`
	// indicates that the Secure attribute will be set in the Set-Cookie header,
	// `never` indicates that the Secure attribute will not be set, and `auto` will set
	// the Secure attribute depending if **Always Use HTTPS** is enabled.
	Secure CookieAttributesSecure `json:"secure"`
	JSON   cookieAttributesJSON   `json:"-"`
}

// cookieAttributesJSON contains the JSON metadata for the struct
// [CookieAttributes]
type cookieAttributesJSON struct {
	Samesite    apijson.Field
	Secure      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CookieAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cookieAttributesJSON) RawJSON() string {
	return r.raw
}

// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
// Note that when using value `none`, the secure attribute cannot be set to
// `never`.
type CookieAttributesSamesite string

const (
	CookieAttributesSamesiteAuto   CookieAttributesSamesite = "auto"
	CookieAttributesSamesiteLax    CookieAttributesSamesite = "lax"
	CookieAttributesSamesiteNone   CookieAttributesSamesite = "none"
	CookieAttributesSamesiteStrict CookieAttributesSamesite = "strict"
)

func (r CookieAttributesSamesite) IsKnown() bool {
	switch r {
	case CookieAttributesSamesiteAuto, CookieAttributesSamesiteLax, CookieAttributesSamesiteNone, CookieAttributesSamesiteStrict:
		return true
	}
	return false
}

// Configures the Secure attribute on the waiting room cookie. Value `always`
// indicates that the Secure attribute will be set in the Set-Cookie header,
// `never` indicates that the Secure attribute will not be set, and `auto` will set
// the Secure attribute depending if **Always Use HTTPS** is enabled.
type CookieAttributesSecure string

const (
	CookieAttributesSecureAuto   CookieAttributesSecure = "auto"
	CookieAttributesSecureAlways CookieAttributesSecure = "always"
	CookieAttributesSecureNever  CookieAttributesSecure = "never"
)

func (r CookieAttributesSecure) IsKnown() bool {
	switch r {
	case CookieAttributesSecureAuto, CookieAttributesSecureAlways, CookieAttributesSecureNever:
		return true
	}
	return false
}

// Configures cookie attributes for the waiting room cookie. This encrypted cookie
// stores a user's status in the waiting room, such as queue position.
type CookieAttributesParam struct {
	// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
	// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
	// Note that when using value `none`, the secure attribute cannot be set to
	// `never`.
	Samesite param.Field[CookieAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on the waiting room cookie. Value `always`
	// indicates that the Secure attribute will be set in the Set-Cookie header,
	// `never` indicates that the Secure attribute will not be set, and `auto` will set
	// the Secure attribute depending if **Always Use HTTPS** is enabled.
	Secure param.Field[CookieAttributesSecure] `json:"secure"`
}

func (r CookieAttributesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The language of the default page template. If no default_template_language is
// provided, then `en-US` (English) will be used.
type DefaultTemplateLanguage string

const (
	DefaultTemplateLanguageEnUs DefaultTemplateLanguage = "en-US"
	DefaultTemplateLanguageEsEs DefaultTemplateLanguage = "es-ES"
	DefaultTemplateLanguageDeDe DefaultTemplateLanguage = "de-DE"
	DefaultTemplateLanguageFrFr DefaultTemplateLanguage = "fr-FR"
	DefaultTemplateLanguageItIt DefaultTemplateLanguage = "it-IT"
	DefaultTemplateLanguageJaJp DefaultTemplateLanguage = "ja-JP"
	DefaultTemplateLanguageKoKr DefaultTemplateLanguage = "ko-KR"
	DefaultTemplateLanguagePtBr DefaultTemplateLanguage = "pt-BR"
	DefaultTemplateLanguageZhCn DefaultTemplateLanguage = "zh-CN"
	DefaultTemplateLanguageZhTw DefaultTemplateLanguage = "zh-TW"
	DefaultTemplateLanguageNlNl DefaultTemplateLanguage = "nl-NL"
	DefaultTemplateLanguagePlPl DefaultTemplateLanguage = "pl-PL"
	DefaultTemplateLanguageIDID DefaultTemplateLanguage = "id-ID"
	DefaultTemplateLanguageTrTr DefaultTemplateLanguage = "tr-TR"
	DefaultTemplateLanguageArEg DefaultTemplateLanguage = "ar-EG"
	DefaultTemplateLanguageRuRu DefaultTemplateLanguage = "ru-RU"
	DefaultTemplateLanguageFaIr DefaultTemplateLanguage = "fa-IR"
	DefaultTemplateLanguageBgBg DefaultTemplateLanguage = "bg-BG"
	DefaultTemplateLanguageHrHr DefaultTemplateLanguage = "hr-HR"
	DefaultTemplateLanguageCsCz DefaultTemplateLanguage = "cs-CZ"
	DefaultTemplateLanguageDaDk DefaultTemplateLanguage = "da-DK"
	DefaultTemplateLanguageFiFi DefaultTemplateLanguage = "fi-FI"
	DefaultTemplateLanguageLtLt DefaultTemplateLanguage = "lt-LT"
	DefaultTemplateLanguageMsMy DefaultTemplateLanguage = "ms-MY"
	DefaultTemplateLanguageNbNo DefaultTemplateLanguage = "nb-NO"
	DefaultTemplateLanguageRoRo DefaultTemplateLanguage = "ro-RO"
	DefaultTemplateLanguageElGr DefaultTemplateLanguage = "el-GR"
	DefaultTemplateLanguageHeIl DefaultTemplateLanguage = "he-IL"
	DefaultTemplateLanguageHiIn DefaultTemplateLanguage = "hi-IN"
	DefaultTemplateLanguageHuHu DefaultTemplateLanguage = "hu-HU"
	DefaultTemplateLanguageSrBa DefaultTemplateLanguage = "sr-BA"
	DefaultTemplateLanguageSkSk DefaultTemplateLanguage = "sk-SK"
	DefaultTemplateLanguageSlSi DefaultTemplateLanguage = "sl-SI"
	DefaultTemplateLanguageSvSe DefaultTemplateLanguage = "sv-SE"
	DefaultTemplateLanguageTlPh DefaultTemplateLanguage = "tl-PH"
	DefaultTemplateLanguageThTh DefaultTemplateLanguage = "th-TH"
	DefaultTemplateLanguageUkUa DefaultTemplateLanguage = "uk-UA"
	DefaultTemplateLanguageViVn DefaultTemplateLanguage = "vi-VN"
)

func (r DefaultTemplateLanguage) IsKnown() bool {
	switch r {
	case DefaultTemplateLanguageEnUs, DefaultTemplateLanguageEsEs, DefaultTemplateLanguageDeDe, DefaultTemplateLanguageFrFr, DefaultTemplateLanguageItIt, DefaultTemplateLanguageJaJp, DefaultTemplateLanguageKoKr, DefaultTemplateLanguagePtBr, DefaultTemplateLanguageZhCn, DefaultTemplateLanguageZhTw, DefaultTemplateLanguageNlNl, DefaultTemplateLanguagePlPl, DefaultTemplateLanguageIDID, DefaultTemplateLanguageTrTr, DefaultTemplateLanguageArEg, DefaultTemplateLanguageRuRu, DefaultTemplateLanguageFaIr, DefaultTemplateLanguageBgBg, DefaultTemplateLanguageHrHr, DefaultTemplateLanguageCsCz, DefaultTemplateLanguageDaDk, DefaultTemplateLanguageFiFi, DefaultTemplateLanguageLtLt, DefaultTemplateLanguageMsMy, DefaultTemplateLanguageNbNo, DefaultTemplateLanguageRoRo, DefaultTemplateLanguageElGr, DefaultTemplateLanguageHeIl, DefaultTemplateLanguageHiIn, DefaultTemplateLanguageHuHu, DefaultTemplateLanguageSrBa, DefaultTemplateLanguageSkSk, DefaultTemplateLanguageSlSi, DefaultTemplateLanguageSvSe, DefaultTemplateLanguageTlPh, DefaultTemplateLanguageThTh, DefaultTemplateLanguageUkUa, DefaultTemplateLanguageViVn:
		return true
	}
	return false
}

type EnabledOriginCommand string

const (
	EnabledOriginCommandRevoke EnabledOriginCommand = "revoke"
)

func (r EnabledOriginCommand) IsKnown() bool {
	switch r {
	case EnabledOriginCommandRevoke:
		return true
	}
	return false
}

type QueryWaitingroomParam struct {
	// The host name to which the waiting room will be applied (no wildcards). Please
	// do not include the scheme (http:// or https://). The host and path combination
	// must be unique.
	Host param.Field[string] `json:"host,required"`
	// A unique name to identify the waiting room. Only alphanumeric characters,
	// hyphens and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// Sets the number of new users that will be let into the route every minute. This
	// value is used as baseline for the number of users that are let in per minute. So
	// it is possible that there is a little more or little less traffic coming to the
	// route based on the traffic patterns at that time around the world.
	NewUsersPerMinute param.Field[int64] `json:"new_users_per_minute,required"`
	// Sets the total number of active user sessions on the route at a point in time. A
	// route is a combination of host and path on which a waiting room is available.
	// This value is used as a baseline for the total number of active user sessions on
	// the route. It is possible to have a situation where there are more or less
	// active users sessions on the route based on the traffic patterns at that time
	// around the world.
	TotalActiveUsers param.Field[int64] `json:"total_active_users,required"`
	// Only available for the Waiting Room Advanced subscription. Additional hostname
	// and path combinations to which this waiting room will be applied. There is an
	// implied wildcard at the end of the path. The hostname and path combination must
	// be unique to this and all other waiting rooms.
	AdditionalRoutes param.Field[[]AdditionalRouteParam] `json:"additional_routes"`
	// Configures cookie attributes for the waiting room cookie. This encrypted cookie
	// stores a user's status in the waiting room, such as queue position.
	CookieAttributes param.Field[CookieAttributesParam] `json:"cookie_attributes"`
	// Appends a '\_' + a custom suffix to the end of Cloudflare Waiting Room's cookie
	// name(**cf_waitingroom). If `cookie_suffix` is "abcd", the cookie name will be
	// `**cf_waitingroom_abcd`. This field is required if using `additional_routes`.
	CookieSuffix param.Field[string] `json:"cookie_suffix"`
	// Only available for the Waiting Room Advanced subscription. This is a template
	// html file that will be rendered at the edge. If no custom_page_html is provided,
	// the default waiting room will be used. The template is based on mustache (
	// https://mustache.github.io/ ). There are several variables that are evaluated by
	// the Cloudflare edge:
	//
	//  1. {{`waitTimeKnown`}} Acts like a boolean value that indicates the behavior to
	//     take when wait time is not available, for instance when queue_all is
	//     **true**.
	//  2. {{`waitTimeFormatted`}} Estimated wait time for the user. For example, five
	//     minutes. Alternatively, you can use:
	//  3. {{`waitTime`}} Number of minutes of estimated wait for a user.
	//  4. {{`waitTimeHours`}} Number of hours of estimated wait for a user
	//     (`Math.floor(waitTime/60)`).
	//  5. {{`waitTimeHourMinutes`}} Number of minutes above the `waitTimeHours` value
	//     (`waitTime%60`).
	//  6. {{`queueIsFull`}} Changes to **true** when no more people can be added to the
	//     queue.
	//
	// To view the full list of variables, look at the `cfWaitingRoom` object described
	// under the `json_response_enabled` property in other Waiting Room API calls.
	CustomPageHTML param.Field[string] `json:"custom_page_html"`
	// The language of the default page template. If no default_template_language is
	// provided, then `en-US` (English) will be used.
	DefaultTemplateLanguage param.Field[DefaultTemplateLanguage] `json:"default_template_language"`
	// A note that you can use to add more details about the waiting room.
	Description param.Field[string] `json:"description"`
	// Only available for the Waiting Room Advanced subscription. Disables automatic
	// renewal of session cookies. If `true`, an accepted user will have
	// session_duration minutes to browse the site. After that, they will have to go
	// through the waiting room again. If `false`, a user's session cookie will be
	// automatically renewed on every request.
	DisableSessionRenewal param.Field[bool] `json:"disable_session_renewal"`
	// A list of enabled origin commands.
	EnabledOriginCommands param.Field[[]EnabledOriginCommand] `json:"enabled_origin_commands"`
	// Only available for the Waiting Room Advanced subscription. If `true`, requests
	// to the waiting room with the header `Accept: application/json` will receive a
	// JSON response object with information on the user's status in the waiting room
	// as opposed to the configured static HTML page. This JSON response object has one
	// property `cfWaitingRoom` which is an object containing the following fields:
	//
	//  1. `inWaitingRoom`: Boolean indicating if the user is in the waiting room
	//     (always **true**).
	//  2. `waitTimeKnown`: Boolean indicating if the current estimated wait times are
	//     accurate. If **false**, they are not available.
	//  3. `waitTime`: Valid only when `waitTimeKnown` is **true**. Integer indicating
	//     the current estimated time in minutes the user will wait in the waiting room.
	//     When `queueingMethod` is **random**, this is set to `waitTime50Percentile`.
	//  4. `waitTime25Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 25% of users that gain entry the fastest (25th percentile).
	//  5. `waitTime50Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 50% of users that gain entry the fastest (50th percentile).
	//     In other words, half of the queued users are expected to let into the origin
	//     website before `waitTime50Percentile` and half are expected to be let in
	//     after it.
	//  6. `waitTime75Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 75% of users that gain entry the fastest (75th percentile).
	//  7. `waitTimeFormatted`: String displaying the `waitTime` formatted in English
	//     for users. If `waitTimeKnown` is **false**, `waitTimeFormatted` will display
	//     **unavailable**.
	//  8. `queueIsFull`: Boolean indicating if the waiting room's queue is currently
	//     full and not accepting new users at the moment.
	//  9. `queueAll`: Boolean indicating if all users will be queued in the waiting
	//     room and no one will be let into the origin website.
	//  10. `lastUpdated`: String displaying the timestamp as an ISO 8601 string of the
	//     user's last attempt to leave the waiting room and be let into the origin
	//     website. The user is able to make another attempt after
	//     `refreshIntervalSeconds` past this time. If the user makes a request too
	//     soon, it will be ignored and `lastUpdated` will not change.
	//  11. `refreshIntervalSeconds`: Integer indicating the number of seconds after
	//     `lastUpdated` until the user is able to make another attempt to leave the
	//     waiting room and be let into the origin website. When the `queueingMethod`
	//     is `reject`, there is no specified refresh time —\_it will always be
	//     **zero**.
	//  12. `queueingMethod`: The queueing method currently used by the waiting room. It
	//     is either **fifo**, **random**, **passthrough**, or **reject**.
	//  13. `isFIFOQueue`: Boolean indicating if the waiting room uses a FIFO
	//     (First-In-First-Out) queue.
	//  14. `isRandomQueue`: Boolean indicating if the waiting room uses a Random queue
	//     where users gain access randomly.
	//  15. `isPassthroughQueue`: Boolean indicating if the waiting room uses a
	//     passthrough queue. Keep in mind that when passthrough is enabled, this JSON
	//     response will only exist when `queueAll` is **true** or `isEventPrequeueing`
	//     is **true** because in all other cases requests will go directly to the
	//     origin.
	//  16. `isRejectQueue`: Boolean indicating if the waiting room uses a reject queue.
	//  17. `isEventActive`: Boolean indicating if an event is currently occurring.
	//     Events are able to change a waiting room's behavior during a specified
	//     period of time. For additional information, look at the event properties
	//     `prequeue_start_time`, `event_start_time`, and `event_end_time` in the
	//     documentation for creating waiting room events. Events are considered active
	//     between these start and end times, as well as during the prequeueing period
	//     if it exists.
	//  18. `isEventPrequeueing`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if an event is currently prequeueing users before it starts.
	//  19. `timeUntilEventStart`: Valid only when `isEventPrequeueing` is **true**.
	//     Integer indicating the number of minutes until the event starts.
	//  20. `timeUntilEventStartFormatted`: String displaying the `timeUntilEventStart`
	//     formatted in English for users. If `isEventPrequeueing` is **false**,
	//     `timeUntilEventStartFormatted` will display **unavailable**.
	//  21. `timeUntilEventEnd`: Valid only when `isEventActive` is **true**. Integer
	//     indicating the number of minutes until the event ends.
	//  22. `timeUntilEventEndFormatted`: String displaying the `timeUntilEventEnd`
	//     formatted in English for users. If `isEventActive` is **false**,
	//     `timeUntilEventEndFormatted` will display **unavailable**.
	//  23. `shuffleAtEventStart`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if the users in the prequeue are shuffled randomly when the event
	//     starts.
	//  24. `turnstile`: Empty when turnstile isn't enabled. String displaying an html
	//     tag to display the Turnstile widget. Please add the `{{{turnstile}}}` tag to
	//     the `custom_html` template to ensure the Turnstile widget appears.
	//  25. `infiniteQueue`: Boolean indicating whether the response is for a user in
	//     the infinite queue.
	//
	// An example cURL to a waiting room could be:
	//
	//	curl -X GET "https://example.com/waitingroom" \
	//		-H "Accept: application/json"
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **fifo** and no event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 0,
	//			"waitTime50Percentile": 0,
	//			"waitTime75Percentile": 0,
	//			"waitTimeFormatted": "10 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "fifo",
	//			"isFIFOQueue": true,
	//			"isRandomQueue": false,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": false,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 0,
	//			"timeUntilEventEndFormatted": "unavailable",
	//			"shuffleAtEventStart": false
	//		}
	//	}
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **random** and an event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 5,
	//			"waitTime50Percentile": 10,
	//			"waitTime75Percentile": 15,
	//			"waitTimeFormatted": "5 minutes to 15 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "random",
	//			"isFIFOQueue": false,
	//			"isRandomQueue": true,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": true,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 15,
	//			"timeUntilEventEndFormatted": "15 minutes",
	//			"shuffleAtEventStart": true
	//		}
	//	}
	JsonResponseEnabled param.Field[bool] `json:"json_response_enabled"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path param.Field[string] `json:"path"`
	// If queue_all is `true`, all the traffic that is coming to a route will be sent
	// to the waiting room. No new traffic can get to the route once this field is set
	// and estimated time will become unavailable.
	QueueAll param.Field[bool] `json:"queue_all"`
	// Sets the queueing method used by the waiting room. Changing this parameter from
	// the **default** queueing method is only available for the Waiting Room Advanced
	// subscription. Regardless of the queueing method, if `queue_all` is enabled or an
	// event is prequeueing, users in the waiting room will not be accepted to the
	// origin. These users will always see a waiting room page that refreshes
	// automatically. The valid queueing methods are:
	//
	//  1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in
	//     the order they arrived.
	//  2. `random`: Random queue where customers gain access randomly, regardless of
	//     arrival time.
	//  3. `passthrough`: Users will pass directly through the waiting room and into the
	//     origin website. As a result, any configured limits will not be respected
	//     while this is enabled. This method can be used as an alternative to disabling
	//     a waiting room (with `suspended`) so that analytics are still reported. This
	//     can be used if you wish to allow all traffic normally, but want to restrict
	//     traffic during a waiting room event, or vice versa.
	//  4. `reject`: Users will be immediately rejected from the waiting room. As a
	//     result, no users will reach the origin website while this is enabled. This
	//     can be used if you wish to reject all traffic while performing maintenance,
	//     block traffic during a specified period of time (an event), or block traffic
	//     while events are not occurring. Consider a waiting room used for vaccine
	//     distribution that only allows traffic during sign-up events, and otherwise
	//     blocks all traffic. For this case, the waiting room uses `reject`, and its
	//     events override this with `fifo`, `random`, or `passthrough`. When this
	//     queueing method is enabled and neither `queueAll` is enabled nor an event is
	//     prequeueing, the waiting room page **will not refresh automatically**.
	QueueingMethod param.Field[QueueingMethod] `json:"queueing_method"`
	// HTTP status code returned to a user while in the queue.
	QueueingStatusCode param.Field[QueueingStatusCode] `json:"queueing_status_code"`
	// Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to
	// the route. If a user is not seen by Cloudflare again in that time period, they
	// will be treated as a new user that visits the route.
	SessionDuration param.Field[int64] `json:"session_duration"`
	// Suspends or allows traffic going to the waiting room. If set to `true`, the
	// traffic will not go to the waiting room.
	Suspended param.Field[bool] `json:"suspended"`
	// Which action to take when a bot is detected using Turnstile. `log` will have no
	// impact on queueing behavior, simply keeping track of how many bots are detected
	// in Waiting Room Analytics. `infinite_queue` will send bots to a false queueing
	// state, where they will never reach your origin. `infinite_queue` requires
	// Advanced Waiting Room.
	TurnstileAction param.Field[TurnstileDetectionAction] `json:"turnstile_action"`
	// Which Turnstile widget type to use for detecting bot traffic. See
	// [the Turnstile documentation](https://developers.cloudflare.com/turnstile/concepts/widget/#widget-types)
	// for the definitions of these widget types. Set to `off` to disable the Turnstile
	// integration entirely. Setting this to anything other than `off` or `invisible`
	// requires Advanced Waiting Room.
	TurnstileMode param.Field[TurnstileWidgetMode] `json:"turnstile_mode"`
}

func (r QueryWaitingroomParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sets the queueing method used by the waiting room. Changing this parameter from
// the **default** queueing method is only available for the Waiting Room Advanced
// subscription. Regardless of the queueing method, if `queue_all` is enabled or an
// event is prequeueing, users in the waiting room will not be accepted to the
// origin. These users will always see a waiting room page that refreshes
// automatically. The valid queueing methods are:
//
//  1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in
//     the order they arrived.
//  2. `random`: Random queue where customers gain access randomly, regardless of
//     arrival time.
//  3. `passthrough`: Users will pass directly through the waiting room and into the
//     origin website. As a result, any configured limits will not be respected
//     while this is enabled. This method can be used as an alternative to disabling
//     a waiting room (with `suspended`) so that analytics are still reported. This
//     can be used if you wish to allow all traffic normally, but want to restrict
//     traffic during a waiting room event, or vice versa.
//  4. `reject`: Users will be immediately rejected from the waiting room. As a
//     result, no users will reach the origin website while this is enabled. This
//     can be used if you wish to reject all traffic while performing maintenance,
//     block traffic during a specified period of time (an event), or block traffic
//     while events are not occurring. Consider a waiting room used for vaccine
//     distribution that only allows traffic during sign-up events, and otherwise
//     blocks all traffic. For this case, the waiting room uses `reject`, and its
//     events override this with `fifo`, `random`, or `passthrough`. When this
//     queueing method is enabled and neither `queueAll` is enabled nor an event is
//     prequeueing, the waiting room page **will not refresh automatically**.
type QueueingMethod string

const (
	QueueingMethodFifo        QueueingMethod = "fifo"
	QueueingMethodRandom      QueueingMethod = "random"
	QueueingMethodPassthrough QueueingMethod = "passthrough"
	QueueingMethodReject      QueueingMethod = "reject"
)

func (r QueueingMethod) IsKnown() bool {
	switch r {
	case QueueingMethodFifo, QueueingMethodRandom, QueueingMethodPassthrough, QueueingMethodReject:
		return true
	}
	return false
}

// HTTP status code returned to a user while in the queue.
type QueueingStatusCode int64

const (
	QueueingStatusCode200 QueueingStatusCode = 200
	QueueingStatusCode202 QueueingStatusCode = 202
	QueueingStatusCode429 QueueingStatusCode = 429
)

func (r QueueingStatusCode) IsKnown() bool {
	switch r {
	case QueueingStatusCode200, QueueingStatusCode202, QueueingStatusCode429:
		return true
	}
	return false
}

type SingleResponseWaitingRoom struct {
	Result Waitingroom                   `json:"result,required"`
	JSON   singleResponseWaitingRoomJSON `json:"-"`
}

// singleResponseWaitingRoomJSON contains the JSON metadata for the struct
// [SingleResponseWaitingRoom]
type singleResponseWaitingRoomJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseWaitingRoom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseWaitingRoomJSON) RawJSON() string {
	return r.raw
}

// Which action to take when a bot is detected using Turnstile. `log` will have no
// impact on queueing behavior, simply keeping track of how many bots are detected
// in Waiting Room Analytics. `infinite_queue` will send bots to a false queueing
// state, where they will never reach your origin. `infinite_queue` requires
// Advanced Waiting Room.
type TurnstileDetectionAction string

const (
	TurnstileDetectionActionLog           TurnstileDetectionAction = "log"
	TurnstileDetectionActionInfiniteQueue TurnstileDetectionAction = "infinite_queue"
)

func (r TurnstileDetectionAction) IsKnown() bool {
	switch r {
	case TurnstileDetectionActionLog, TurnstileDetectionActionInfiniteQueue:
		return true
	}
	return false
}

// Which Turnstile widget type to use for detecting bot traffic. See
// [the Turnstile documentation](https://developers.cloudflare.com/turnstile/concepts/widget/#widget-types)
// for the definitions of these widget types. Set to `off` to disable the Turnstile
// integration entirely. Setting this to anything other than `off` or `invisible`
// requires Advanced Waiting Room.
type TurnstileWidgetMode string

const (
	TurnstileWidgetModeOff                   TurnstileWidgetMode = "off"
	TurnstileWidgetModeInvisible             TurnstileWidgetMode = "invisible"
	TurnstileWidgetModeVisibleNonInteractive TurnstileWidgetMode = "visible_non_interactive"
	TurnstileWidgetModeVisibleManaged        TurnstileWidgetMode = "visible_managed"
)

func (r TurnstileWidgetMode) IsKnown() bool {
	switch r {
	case TurnstileWidgetModeOff, TurnstileWidgetModeInvisible, TurnstileWidgetModeVisibleNonInteractive, TurnstileWidgetModeVisibleManaged:
		return true
	}
	return false
}

type Waitingroom struct {
	ID string `json:"id"`
	// Only available for the Waiting Room Advanced subscription. Additional hostname
	// and path combinations to which this waiting room will be applied. There is an
	// implied wildcard at the end of the path. The hostname and path combination must
	// be unique to this and all other waiting rooms.
	AdditionalRoutes []AdditionalRoute `json:"additional_routes"`
	// Configures cookie attributes for the waiting room cookie. This encrypted cookie
	// stores a user's status in the waiting room, such as queue position.
	CookieAttributes CookieAttributes `json:"cookie_attributes"`
	// Appends a '\_' + a custom suffix to the end of Cloudflare Waiting Room's cookie
	// name(**cf_waitingroom). If `cookie_suffix` is "abcd", the cookie name will be
	// `**cf_waitingroom_abcd`. This field is required if using `additional_routes`.
	CookieSuffix string    `json:"cookie_suffix"`
	CreatedOn    time.Time `json:"created_on" format:"date-time"`
	// Only available for the Waiting Room Advanced subscription. This is a template
	// html file that will be rendered at the edge. If no custom_page_html is provided,
	// the default waiting room will be used. The template is based on mustache (
	// https://mustache.github.io/ ). There are several variables that are evaluated by
	// the Cloudflare edge:
	//
	//  1. {{`waitTimeKnown`}} Acts like a boolean value that indicates the behavior to
	//     take when wait time is not available, for instance when queue_all is
	//     **true**.
	//  2. {{`waitTimeFormatted`}} Estimated wait time for the user. For example, five
	//     minutes. Alternatively, you can use:
	//  3. {{`waitTime`}} Number of minutes of estimated wait for a user.
	//  4. {{`waitTimeHours`}} Number of hours of estimated wait for a user
	//     (`Math.floor(waitTime/60)`).
	//  5. {{`waitTimeHourMinutes`}} Number of minutes above the `waitTimeHours` value
	//     (`waitTime%60`).
	//  6. {{`queueIsFull`}} Changes to **true** when no more people can be added to the
	//     queue.
	//
	// To view the full list of variables, look at the `cfWaitingRoom` object described
	// under the `json_response_enabled` property in other Waiting Room API calls.
	CustomPageHTML string `json:"custom_page_html"`
	// The language of the default page template. If no default_template_language is
	// provided, then `en-US` (English) will be used.
	DefaultTemplateLanguage DefaultTemplateLanguage `json:"default_template_language"`
	// A note that you can use to add more details about the waiting room.
	Description string `json:"description"`
	// Only available for the Waiting Room Advanced subscription. Disables automatic
	// renewal of session cookies. If `true`, an accepted user will have
	// session_duration minutes to browse the site. After that, they will have to go
	// through the waiting room again. If `false`, a user's session cookie will be
	// automatically renewed on every request.
	DisableSessionRenewal bool `json:"disable_session_renewal"`
	// A list of enabled origin commands.
	EnabledOriginCommands []EnabledOriginCommand `json:"enabled_origin_commands"`
	// The host name to which the waiting room will be applied (no wildcards). Please
	// do not include the scheme (http:// or https://). The host and path combination
	// must be unique.
	Host string `json:"host"`
	// Only available for the Waiting Room Advanced subscription. If `true`, requests
	// to the waiting room with the header `Accept: application/json` will receive a
	// JSON response object with information on the user's status in the waiting room
	// as opposed to the configured static HTML page. This JSON response object has one
	// property `cfWaitingRoom` which is an object containing the following fields:
	//
	//  1. `inWaitingRoom`: Boolean indicating if the user is in the waiting room
	//     (always **true**).
	//  2. `waitTimeKnown`: Boolean indicating if the current estimated wait times are
	//     accurate. If **false**, they are not available.
	//  3. `waitTime`: Valid only when `waitTimeKnown` is **true**. Integer indicating
	//     the current estimated time in minutes the user will wait in the waiting room.
	//     When `queueingMethod` is **random**, this is set to `waitTime50Percentile`.
	//  4. `waitTime25Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 25% of users that gain entry the fastest (25th percentile).
	//  5. `waitTime50Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 50% of users that gain entry the fastest (50th percentile).
	//     In other words, half of the queued users are expected to let into the origin
	//     website before `waitTime50Percentile` and half are expected to be let in
	//     after it.
	//  6. `waitTime75Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 75% of users that gain entry the fastest (75th percentile).
	//  7. `waitTimeFormatted`: String displaying the `waitTime` formatted in English
	//     for users. If `waitTimeKnown` is **false**, `waitTimeFormatted` will display
	//     **unavailable**.
	//  8. `queueIsFull`: Boolean indicating if the waiting room's queue is currently
	//     full and not accepting new users at the moment.
	//  9. `queueAll`: Boolean indicating if all users will be queued in the waiting
	//     room and no one will be let into the origin website.
	//  10. `lastUpdated`: String displaying the timestamp as an ISO 8601 string of the
	//     user's last attempt to leave the waiting room and be let into the origin
	//     website. The user is able to make another attempt after
	//     `refreshIntervalSeconds` past this time. If the user makes a request too
	//     soon, it will be ignored and `lastUpdated` will not change.
	//  11. `refreshIntervalSeconds`: Integer indicating the number of seconds after
	//     `lastUpdated` until the user is able to make another attempt to leave the
	//     waiting room and be let into the origin website. When the `queueingMethod`
	//     is `reject`, there is no specified refresh time —\_it will always be
	//     **zero**.
	//  12. `queueingMethod`: The queueing method currently used by the waiting room. It
	//     is either **fifo**, **random**, **passthrough**, or **reject**.
	//  13. `isFIFOQueue`: Boolean indicating if the waiting room uses a FIFO
	//     (First-In-First-Out) queue.
	//  14. `isRandomQueue`: Boolean indicating if the waiting room uses a Random queue
	//     where users gain access randomly.
	//  15. `isPassthroughQueue`: Boolean indicating if the waiting room uses a
	//     passthrough queue. Keep in mind that when passthrough is enabled, this JSON
	//     response will only exist when `queueAll` is **true** or `isEventPrequeueing`
	//     is **true** because in all other cases requests will go directly to the
	//     origin.
	//  16. `isRejectQueue`: Boolean indicating if the waiting room uses a reject queue.
	//  17. `isEventActive`: Boolean indicating if an event is currently occurring.
	//     Events are able to change a waiting room's behavior during a specified
	//     period of time. For additional information, look at the event properties
	//     `prequeue_start_time`, `event_start_time`, and `event_end_time` in the
	//     documentation for creating waiting room events. Events are considered active
	//     between these start and end times, as well as during the prequeueing period
	//     if it exists.
	//  18. `isEventPrequeueing`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if an event is currently prequeueing users before it starts.
	//  19. `timeUntilEventStart`: Valid only when `isEventPrequeueing` is **true**.
	//     Integer indicating the number of minutes until the event starts.
	//  20. `timeUntilEventStartFormatted`: String displaying the `timeUntilEventStart`
	//     formatted in English for users. If `isEventPrequeueing` is **false**,
	//     `timeUntilEventStartFormatted` will display **unavailable**.
	//  21. `timeUntilEventEnd`: Valid only when `isEventActive` is **true**. Integer
	//     indicating the number of minutes until the event ends.
	//  22. `timeUntilEventEndFormatted`: String displaying the `timeUntilEventEnd`
	//     formatted in English for users. If `isEventActive` is **false**,
	//     `timeUntilEventEndFormatted` will display **unavailable**.
	//  23. `shuffleAtEventStart`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if the users in the prequeue are shuffled randomly when the event
	//     starts.
	//  24. `turnstile`: Empty when turnstile isn't enabled. String displaying an html
	//     tag to display the Turnstile widget. Please add the `{{{turnstile}}}` tag to
	//     the `custom_html` template to ensure the Turnstile widget appears.
	//  25. `infiniteQueue`: Boolean indicating whether the response is for a user in
	//     the infinite queue.
	//
	// An example cURL to a waiting room could be:
	//
	//	curl -X GET "https://example.com/waitingroom" \
	//		-H "Accept: application/json"
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **fifo** and no event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 0,
	//			"waitTime50Percentile": 0,
	//			"waitTime75Percentile": 0,
	//			"waitTimeFormatted": "10 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "fifo",
	//			"isFIFOQueue": true,
	//			"isRandomQueue": false,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": false,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 0,
	//			"timeUntilEventEndFormatted": "unavailable",
	//			"shuffleAtEventStart": false
	//		}
	//	}
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **random** and an event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 5,
	//			"waitTime50Percentile": 10,
	//			"waitTime75Percentile": 15,
	//			"waitTimeFormatted": "5 minutes to 15 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "random",
	//			"isFIFOQueue": false,
	//			"isRandomQueue": true,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": true,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 15,
	//			"timeUntilEventEndFormatted": "15 minutes",
	//			"shuffleAtEventStart": true
	//		}
	//	}
	JsonResponseEnabled bool      `json:"json_response_enabled"`
	ModifiedOn          time.Time `json:"modified_on" format:"date-time"`
	// A unique name to identify the waiting room. Only alphanumeric characters,
	// hyphens and underscores are allowed.
	Name string `json:"name"`
	// Sets the number of new users that will be let into the route every minute. This
	// value is used as baseline for the number of users that are let in per minute. So
	// it is possible that there is a little more or little less traffic coming to the
	// route based on the traffic patterns at that time around the world.
	NewUsersPerMinute int64 `json:"new_users_per_minute"`
	// An ISO 8601 timestamp that marks when the next event will begin queueing.
	NextEventPrequeueStartTime string `json:"next_event_prequeue_start_time,nullable"`
	// An ISO 8601 timestamp that marks when the next event will start.
	NextEventStartTime string `json:"next_event_start_time,nullable"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path string `json:"path"`
	// If queue_all is `true`, all the traffic that is coming to a route will be sent
	// to the waiting room. No new traffic can get to the route once this field is set
	// and estimated time will become unavailable.
	QueueAll bool `json:"queue_all"`
	// Sets the queueing method used by the waiting room. Changing this parameter from
	// the **default** queueing method is only available for the Waiting Room Advanced
	// subscription. Regardless of the queueing method, if `queue_all` is enabled or an
	// event is prequeueing, users in the waiting room will not be accepted to the
	// origin. These users will always see a waiting room page that refreshes
	// automatically. The valid queueing methods are:
	//
	//  1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in
	//     the order they arrived.
	//  2. `random`: Random queue where customers gain access randomly, regardless of
	//     arrival time.
	//  3. `passthrough`: Users will pass directly through the waiting room and into the
	//     origin website. As a result, any configured limits will not be respected
	//     while this is enabled. This method can be used as an alternative to disabling
	//     a waiting room (with `suspended`) so that analytics are still reported. This
	//     can be used if you wish to allow all traffic normally, but want to restrict
	//     traffic during a waiting room event, or vice versa.
	//  4. `reject`: Users will be immediately rejected from the waiting room. As a
	//     result, no users will reach the origin website while this is enabled. This
	//     can be used if you wish to reject all traffic while performing maintenance,
	//     block traffic during a specified period of time (an event), or block traffic
	//     while events are not occurring. Consider a waiting room used for vaccine
	//     distribution that only allows traffic during sign-up events, and otherwise
	//     blocks all traffic. For this case, the waiting room uses `reject`, and its
	//     events override this with `fifo`, `random`, or `passthrough`. When this
	//     queueing method is enabled and neither `queueAll` is enabled nor an event is
	//     prequeueing, the waiting room page **will not refresh automatically**.
	QueueingMethod QueueingMethod `json:"queueing_method"`
	// HTTP status code returned to a user while in the queue.
	QueueingStatusCode QueueingStatusCode `json:"queueing_status_code"`
	// Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to
	// the route. If a user is not seen by Cloudflare again in that time period, they
	// will be treated as a new user that visits the route.
	SessionDuration int64 `json:"session_duration"`
	// Suspends or allows traffic going to the waiting room. If set to `true`, the
	// traffic will not go to the waiting room.
	Suspended bool `json:"suspended"`
	// Sets the total number of active user sessions on the route at a point in time. A
	// route is a combination of host and path on which a waiting room is available.
	// This value is used as a baseline for the total number of active user sessions on
	// the route. It is possible to have a situation where there are more or less
	// active users sessions on the route based on the traffic patterns at that time
	// around the world.
	TotalActiveUsers int64 `json:"total_active_users"`
	// Which action to take when a bot is detected using Turnstile. `log` will have no
	// impact on queueing behavior, simply keeping track of how many bots are detected
	// in Waiting Room Analytics. `infinite_queue` will send bots to a false queueing
	// state, where they will never reach your origin. `infinite_queue` requires
	// Advanced Waiting Room.
	TurnstileAction TurnstileDetectionAction `json:"turnstile_action"`
	// Which Turnstile widget type to use for detecting bot traffic. See
	// [the Turnstile documentation](https://developers.cloudflare.com/turnstile/concepts/widget/#widget-types)
	// for the definitions of these widget types. Set to `off` to disable the Turnstile
	// integration entirely. Setting this to anything other than `off` or `invisible`
	// requires Advanced Waiting Room.
	TurnstileMode TurnstileWidgetMode `json:"turnstile_mode"`
	JSON          waitingroomJSON     `json:"-"`
}

// waitingroomJSON contains the JSON metadata for the struct [Waitingroom]
type waitingroomJSON struct {
	ID                         apijson.Field
	AdditionalRoutes           apijson.Field
	CookieAttributes           apijson.Field
	CookieSuffix               apijson.Field
	CreatedOn                  apijson.Field
	CustomPageHTML             apijson.Field
	DefaultTemplateLanguage    apijson.Field
	Description                apijson.Field
	DisableSessionRenewal      apijson.Field
	EnabledOriginCommands      apijson.Field
	Host                       apijson.Field
	JsonResponseEnabled        apijson.Field
	ModifiedOn                 apijson.Field
	Name                       apijson.Field
	NewUsersPerMinute          apijson.Field
	NextEventPrequeueStartTime apijson.Field
	NextEventStartTime         apijson.Field
	Path                       apijson.Field
	QueueAll                   apijson.Field
	QueueingMethod             apijson.Field
	QueueingStatusCode         apijson.Field
	SessionDuration            apijson.Field
	Suspended                  apijson.Field
	TotalActiveUsers           apijson.Field
	TurnstileAction            apijson.Field
	TurnstileMode              apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *Waitingroom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingroomJSON) RawJSON() string {
	return r.raw
}

type ZoneWaitingRoomDeleteResponse struct {
	Result ZoneWaitingRoomDeleteResponseResult `json:"result,required"`
	JSON   zoneWaitingRoomDeleteResponseJSON   `json:"-"`
}

// zoneWaitingRoomDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneWaitingRoomDeleteResponse]
type zoneWaitingRoomDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWaitingRoomDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWaitingRoomDeleteResponseResult struct {
	ID   string                                  `json:"id"`
	JSON zoneWaitingRoomDeleteResponseResultJSON `json:"-"`
}

// zoneWaitingRoomDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneWaitingRoomDeleteResponseResult]
type zoneWaitingRoomDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWaitingRoomDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneWaitingRoomPreviewResponse struct {
	Result ZoneWaitingRoomPreviewResponseResult `json:"result,required"`
	JSON   zoneWaitingRoomPreviewResponseJSON   `json:"-"`
}

// zoneWaitingRoomPreviewResponseJSON contains the JSON metadata for the struct
// [ZoneWaitingRoomPreviewResponse]
type zoneWaitingRoomPreviewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWaitingRoomPreviewResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWaitingRoomPreviewResponseResult struct {
	// URL where the custom waiting room page can temporarily be previewed.
	PreviewURL string                                   `json:"preview_url"`
	JSON       zoneWaitingRoomPreviewResponseResultJSON `json:"-"`
}

// zoneWaitingRoomPreviewResponseResultJSON contains the JSON metadata for the
// struct [ZoneWaitingRoomPreviewResponseResult]
type zoneWaitingRoomPreviewResponseResultJSON struct {
	PreviewURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomPreviewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWaitingRoomPreviewResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneWaitingRoomStatusResponse struct {
	Result ZoneWaitingRoomStatusResponseResult `json:"result,required"`
	JSON   zoneWaitingRoomStatusResponseJSON   `json:"-"`
}

// zoneWaitingRoomStatusResponseJSON contains the JSON metadata for the struct
// [ZoneWaitingRoomStatusResponse]
type zoneWaitingRoomStatusResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWaitingRoomStatusResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWaitingRoomStatusResponseResult struct {
	EstimatedQueuedUsers      int64                                     `json:"estimated_queued_users"`
	EstimatedTotalActiveUsers int64                                     `json:"estimated_total_active_users"`
	EventID                   string                                    `json:"event_id"`
	MaxEstimatedTimeMinutes   int64                                     `json:"max_estimated_time_minutes"`
	Status                    ZoneWaitingRoomStatusResponseResultStatus `json:"status"`
	JSON                      zoneWaitingRoomStatusResponseResultJSON   `json:"-"`
}

// zoneWaitingRoomStatusResponseResultJSON contains the JSON metadata for the
// struct [ZoneWaitingRoomStatusResponseResult]
type zoneWaitingRoomStatusResponseResultJSON struct {
	EstimatedQueuedUsers      apijson.Field
	EstimatedTotalActiveUsers apijson.Field
	EventID                   apijson.Field
	MaxEstimatedTimeMinutes   apijson.Field
	Status                    apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ZoneWaitingRoomStatusResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWaitingRoomStatusResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneWaitingRoomStatusResponseResultStatus string

const (
	ZoneWaitingRoomStatusResponseResultStatusEventPrequeueing ZoneWaitingRoomStatusResponseResultStatus = "event_prequeueing"
	ZoneWaitingRoomStatusResponseResultStatusNotQueueing      ZoneWaitingRoomStatusResponseResultStatus = "not_queueing"
	ZoneWaitingRoomStatusResponseResultStatusQueueing         ZoneWaitingRoomStatusResponseResultStatus = "queueing"
	ZoneWaitingRoomStatusResponseResultStatusSuspended        ZoneWaitingRoomStatusResponseResultStatus = "suspended"
)

func (r ZoneWaitingRoomStatusResponseResultStatus) IsKnown() bool {
	switch r {
	case ZoneWaitingRoomStatusResponseResultStatusEventPrequeueing, ZoneWaitingRoomStatusResponseResultStatusNotQueueing, ZoneWaitingRoomStatusResponseResultStatusQueueing, ZoneWaitingRoomStatusResponseResultStatusSuspended:
		return true
	}
	return false
}

type ZoneWaitingRoomNewParams struct {
	QueryWaitingroom QueryWaitingroomParam `json:"query_waitingroom,required"`
}

func (r ZoneWaitingRoomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.QueryWaitingroom)
}

type ZoneWaitingRoomUpdateParams struct {
	QueryWaitingroom QueryWaitingroomParam `json:"query_waitingroom,required"`
}

func (r ZoneWaitingRoomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.QueryWaitingroom)
}

type ZoneWaitingRoomPatchParams struct {
	QueryWaitingroom QueryWaitingroomParam `json:"query_waitingroom,required"`
}

func (r ZoneWaitingRoomPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.QueryWaitingroom)
}

type ZoneWaitingRoomPreviewParams struct {
	// Only available for the Waiting Room Advanced subscription. This is a template
	// html file that will be rendered at the edge. If no custom_page_html is provided,
	// the default waiting room will be used. The template is based on mustache (
	// https://mustache.github.io/ ). There are several variables that are evaluated by
	// the Cloudflare edge:
	//
	//  1. {{`waitTimeKnown`}} Acts like a boolean value that indicates the behavior to
	//     take when wait time is not available, for instance when queue_all is
	//     **true**.
	//  2. {{`waitTimeFormatted`}} Estimated wait time for the user. For example, five
	//     minutes. Alternatively, you can use:
	//  3. {{`waitTime`}} Number of minutes of estimated wait for a user.
	//  4. {{`waitTimeHours`}} Number of hours of estimated wait for a user
	//     (`Math.floor(waitTime/60)`).
	//  5. {{`waitTimeHourMinutes`}} Number of minutes above the `waitTimeHours` value
	//     (`waitTime%60`).
	//  6. {{`queueIsFull`}} Changes to **true** when no more people can be added to the
	//     queue.
	//
	// To view the full list of variables, look at the `cfWaitingRoom` object described
	// under the `json_response_enabled` property in other Waiting Room API calls.
	CustomHTML param.Field[string] `json:"custom_html,required"`
}

func (r ZoneWaitingRoomPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
