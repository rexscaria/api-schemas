// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// ZoneBotManagementService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneBotManagementService] method instead.
type ZoneBotManagementService struct {
	Options []option.RequestOption
}

// NewZoneBotManagementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneBotManagementService(opts ...option.RequestOption) (r *ZoneBotManagementService) {
	r = &ZoneBotManagementService{}
	r.Options = opts
	return
}

// Retrieve a zone's Bot Management Config
func (r *ZoneBotManagementService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ManagementResponseBody, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/bot_management", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Bot Management configuration for a zone.
//
// This API is used to update:
//
// - **Bot Fight Mode**
// - **Super Bot Fight Mode**
// - **Bot Management for Enterprise**
//
// See [Bot Plans](https://developers.cloudflare.com/bots/plans/) for more
// information on the different plans \
// If you recently upgraded or downgraded your plan, refer to the following examples
// to clean up old configurations. Copy and paste the example body to remove old zone
// configurations based on your current plan.
//
// #### Clean up configuration for Bot Fight Mode plan
//
// ```json
//
//	{
//	  "sbfm_likely_automated": "allow",
//	  "sbfm_definitely_automated": "allow",
//	  "sbfm_verified_bots": "allow",
//	  "sbfm_static_resource_protection": false,
//	  "optimize_wordpress": false,
//	  "suppress_session_score": false
//	}
//
// ```
//
// #### Clean up configuration for SBFM Pro plan
//
// ```json
//
//	{
//	  "sbfm_likely_automated": "allow",
//	  "fight_mode": false
//	}
//
// ```
//
// #### Clean up configuration for SBFM Biz plan
//
// ```json
//
//	{
//	  "fight_mode": false
//	}
//
// ```
//
// #### Clean up configuration for BM Enterprise Subscription plan
//
// It is strongly recommended that you ensure you have
// [custom rules](https://developers.cloudflare.com/waf/custom-rules/) in place to
// protect your zone before disabling the SBFM rules. Without these protections,
// your zone is vulnerable to attacks.
//
// ```json
//
//	{
//	  "sbfm_likely_automated": "allow",
//	  "sbfm_definitely_automated": "allow",
//	  "sbfm_verified_bots": "allow",
//	  "sbfm_static_resource_protection": false,
//	  "optimize_wordpress": false,
//	  "fight_mode": false
//	}
//
// ```
func (r *ZoneBotManagementService) Update(ctx context.Context, zoneID string, body ZoneBotManagementUpdateParams, opts ...option.RequestOption) (res *ManagementResponseBody, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/bot_management", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type BmSubscriptionConfig struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection BmSubscriptionConfigAIBotsProtection `json:"ai_bots_protection"`
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel bool `json:"auto_update_model"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection BmSubscriptionConfigCrawlerProtection `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// A read-only field that shows which unauthorized settings are currently active on
	// the zone. These settings typically result from upgrades or downgrades.
	StaleZoneConfiguration BmSubscriptionConfigStaleZoneConfiguration `json:"stale_zone_configuration"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore bool `json:"suppress_session_score"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                     `json:"using_latest_model"`
	JSON             bmSubscriptionConfigJSON `json:"-"`
}

// bmSubscriptionConfigJSON contains the JSON metadata for the struct
// [BmSubscriptionConfig]
type bmSubscriptionConfigJSON struct {
	AIBotsProtection       apijson.Field
	AutoUpdateModel        apijson.Field
	CrawlerProtection      apijson.Field
	EnableJs               apijson.Field
	StaleZoneConfiguration apijson.Field
	SuppressSessionScore   apijson.Field
	UsingLatestModel       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *BmSubscriptionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bmSubscriptionConfigJSON) RawJSON() string {
	return r.raw
}

func (r BmSubscriptionConfig) implementsManagementResponseBodyResult() {}

// Enable rule to block AI Scrapers and Crawlers. Please note the value
// `only_on_ad_pages` is currently not available for Enterprise customers.
type BmSubscriptionConfigAIBotsProtection string

const (
	BmSubscriptionConfigAIBotsProtectionBlock         BmSubscriptionConfigAIBotsProtection = "block"
	BmSubscriptionConfigAIBotsProtectionDisabled      BmSubscriptionConfigAIBotsProtection = "disabled"
	BmSubscriptionConfigAIBotsProtectionOnlyOnAdPages BmSubscriptionConfigAIBotsProtection = "only_on_ad_pages"
)

func (r BmSubscriptionConfigAIBotsProtection) IsKnown() bool {
	switch r {
	case BmSubscriptionConfigAIBotsProtectionBlock, BmSubscriptionConfigAIBotsProtectionDisabled, BmSubscriptionConfigAIBotsProtectionOnlyOnAdPages:
		return true
	}
	return false
}

// Enable rule to punish AI Scrapers and Crawlers via a link maze.
type BmSubscriptionConfigCrawlerProtection string

const (
	BmSubscriptionConfigCrawlerProtectionEnabled  BmSubscriptionConfigCrawlerProtection = "enabled"
	BmSubscriptionConfigCrawlerProtectionDisabled BmSubscriptionConfigCrawlerProtection = "disabled"
)

func (r BmSubscriptionConfigCrawlerProtection) IsKnown() bool {
	switch r {
	case BmSubscriptionConfigCrawlerProtectionEnabled, BmSubscriptionConfigCrawlerProtectionDisabled:
		return true
	}
	return false
}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type BmSubscriptionConfigStaleZoneConfiguration struct {
	// Indicates that the zone's Bot Fight Mode is turned on.
	FightMode bool `json:"fight_mode"`
	// Indicates that the zone's wordpress optimization for SBFM is turned on.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Indicates that the zone's definitely automated requests are being blocked or
	// challenged.
	SbfmDefinitelyAutomated string `json:"sbfm_definitely_automated"`
	// Indicates that the zone's likely automated requests are being blocked or
	// challenged.
	SbfmLikelyAutomated string `json:"sbfm_likely_automated"`
	// Indicates that the zone's static resource protection is turned on.
	SbfmStaticResourceProtection string `json:"sbfm_static_resource_protection"`
	// Indicates that the zone's verified bot requests are being blocked.
	SbfmVerifiedBots string                                         `json:"sbfm_verified_bots"`
	JSON             bmSubscriptionConfigStaleZoneConfigurationJSON `json:"-"`
}

// bmSubscriptionConfigStaleZoneConfigurationJSON contains the JSON metadata for
// the struct [BmSubscriptionConfigStaleZoneConfiguration]
type bmSubscriptionConfigStaleZoneConfigurationJSON struct {
	FightMode                    apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmLikelyAutomated          apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *BmSubscriptionConfigStaleZoneConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bmSubscriptionConfigStaleZoneConfigurationJSON) RawJSON() string {
	return r.raw
}

type BmSubscriptionConfigParam struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection param.Field[BmSubscriptionConfigAIBotsProtection] `json:"ai_bots_protection"`
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel param.Field[bool] `json:"auto_update_model"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection param.Field[BmSubscriptionConfigCrawlerProtection] `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore param.Field[bool] `json:"suppress_session_score"`
}

func (r BmSubscriptionConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BmSubscriptionConfigParam) implementsZoneBotManagementUpdateParamsBodyUnion() {}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type BmSubscriptionConfigStaleZoneConfigurationParam struct {
	// Indicates that the zone's Bot Fight Mode is turned on.
	FightMode param.Field[bool] `json:"fight_mode"`
	// Indicates that the zone's wordpress optimization for SBFM is turned on.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Indicates that the zone's definitely automated requests are being blocked or
	// challenged.
	SbfmDefinitelyAutomated param.Field[string] `json:"sbfm_definitely_automated"`
	// Indicates that the zone's likely automated requests are being blocked or
	// challenged.
	SbfmLikelyAutomated param.Field[string] `json:"sbfm_likely_automated"`
	// Indicates that the zone's static resource protection is turned on.
	SbfmStaticResourceProtection param.Field[string] `json:"sbfm_static_resource_protection"`
	// Indicates that the zone's verified bot requests are being blocked.
	SbfmVerifiedBots param.Field[string] `json:"sbfm_verified_bots"`
}

func (r BmSubscriptionConfigStaleZoneConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BotFightModeConfig struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection BotFightModeConfigAIBotsProtection `json:"ai_bots_protection"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection BotFightModeConfigCrawlerProtection `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// A read-only field that shows which unauthorized settings are currently active on
	// the zone. These settings typically result from upgrades or downgrades.
	StaleZoneConfiguration BotFightModeConfigStaleZoneConfiguration `json:"stale_zone_configuration"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                   `json:"using_latest_model"`
	JSON             botFightModeConfigJSON `json:"-"`
}

// botFightModeConfigJSON contains the JSON metadata for the struct
// [BotFightModeConfig]
type botFightModeConfigJSON struct {
	AIBotsProtection       apijson.Field
	CrawlerProtection      apijson.Field
	EnableJs               apijson.Field
	FightMode              apijson.Field
	StaleZoneConfiguration apijson.Field
	UsingLatestModel       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *BotFightModeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botFightModeConfigJSON) RawJSON() string {
	return r.raw
}

func (r BotFightModeConfig) implementsManagementResponseBodyResult() {}

// Enable rule to block AI Scrapers and Crawlers. Please note the value
// `only_on_ad_pages` is currently not available for Enterprise customers.
type BotFightModeConfigAIBotsProtection string

const (
	BotFightModeConfigAIBotsProtectionBlock         BotFightModeConfigAIBotsProtection = "block"
	BotFightModeConfigAIBotsProtectionDisabled      BotFightModeConfigAIBotsProtection = "disabled"
	BotFightModeConfigAIBotsProtectionOnlyOnAdPages BotFightModeConfigAIBotsProtection = "only_on_ad_pages"
)

func (r BotFightModeConfigAIBotsProtection) IsKnown() bool {
	switch r {
	case BotFightModeConfigAIBotsProtectionBlock, BotFightModeConfigAIBotsProtectionDisabled, BotFightModeConfigAIBotsProtectionOnlyOnAdPages:
		return true
	}
	return false
}

// Enable rule to punish AI Scrapers and Crawlers via a link maze.
type BotFightModeConfigCrawlerProtection string

const (
	BotFightModeConfigCrawlerProtectionEnabled  BotFightModeConfigCrawlerProtection = "enabled"
	BotFightModeConfigCrawlerProtectionDisabled BotFightModeConfigCrawlerProtection = "disabled"
)

func (r BotFightModeConfigCrawlerProtection) IsKnown() bool {
	switch r {
	case BotFightModeConfigCrawlerProtectionEnabled, BotFightModeConfigCrawlerProtectionDisabled:
		return true
	}
	return false
}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type BotFightModeConfigStaleZoneConfiguration struct {
	// Indicates that the zone's wordpress optimization for SBFM is turned on.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Indicates that the zone's definitely automated requests are being blocked or
	// challenged.
	SbfmDefinitelyAutomated string `json:"sbfm_definitely_automated"`
	// Indicates that the zone's likely automated requests are being blocked or
	// challenged.
	SbfmLikelyAutomated string `json:"sbfm_likely_automated"`
	// Indicates that the zone's static resource protection is turned on.
	SbfmStaticResourceProtection string `json:"sbfm_static_resource_protection"`
	// Indicates that the zone's verified bot requests are being blocked.
	SbfmVerifiedBots string `json:"sbfm_verified_bots"`
	// Indicates that the zone's session score tracking is disabled.
	SuppressSessionScore bool                                         `json:"suppress_session_score"`
	JSON                 botFightModeConfigStaleZoneConfigurationJSON `json:"-"`
}

// botFightModeConfigStaleZoneConfigurationJSON contains the JSON metadata for the
// struct [BotFightModeConfigStaleZoneConfiguration]
type botFightModeConfigStaleZoneConfigurationJSON struct {
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmLikelyAutomated          apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	SuppressSessionScore         apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *BotFightModeConfigStaleZoneConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botFightModeConfigStaleZoneConfigurationJSON) RawJSON() string {
	return r.raw
}

type BotFightModeConfigParam struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection param.Field[BotFightModeConfigAIBotsProtection] `json:"ai_bots_protection"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection param.Field[BotFightModeConfigCrawlerProtection] `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode param.Field[bool] `json:"fight_mode"`
}

func (r BotFightModeConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BotFightModeConfigParam) implementsZoneBotManagementUpdateParamsBodyUnion() {}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type BotFightModeConfigStaleZoneConfigurationParam struct {
	// Indicates that the zone's wordpress optimization for SBFM is turned on.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Indicates that the zone's definitely automated requests are being blocked or
	// challenged.
	SbfmDefinitelyAutomated param.Field[string] `json:"sbfm_definitely_automated"`
	// Indicates that the zone's likely automated requests are being blocked or
	// challenged.
	SbfmLikelyAutomated param.Field[string] `json:"sbfm_likely_automated"`
	// Indicates that the zone's static resource protection is turned on.
	SbfmStaticResourceProtection param.Field[string] `json:"sbfm_static_resource_protection"`
	// Indicates that the zone's verified bot requests are being blocked.
	SbfmVerifiedBots param.Field[string] `json:"sbfm_verified_bots"`
	// Indicates that the zone's session score tracking is disabled.
	SuppressSessionScore param.Field[bool] `json:"suppress_session_score"`
}

func (r BotFightModeConfigStaleZoneConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ManagementResponseBody struct {
	Errors   []MessagesBotManagementItem `json:"errors,required"`
	Messages []MessagesBotManagementItem `json:"messages,required"`
	// Whether the API call was successful.
	Success ManagementResponseBodySuccess `json:"success,required"`
	Result  ManagementResponseBodyResult  `json:"result"`
	JSON    managementResponseBodyJSON    `json:"-"`
}

// managementResponseBodyJSON contains the JSON metadata for the struct
// [ManagementResponseBody]
type managementResponseBodyJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagementResponseBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managementResponseBodyJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ManagementResponseBodySuccess bool

const (
	ManagementResponseBodySuccessTrue ManagementResponseBodySuccess = true
)

func (r ManagementResponseBodySuccess) IsKnown() bool {
	switch r {
	case ManagementResponseBodySuccessTrue:
		return true
	}
	return false
}

type ManagementResponseBodyResult struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection ManagementResponseBodyResultAIBotsProtection `json:"ai_bots_protection"`
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel bool `json:"auto_update_model"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection ManagementResponseBodyResultCrawlerProtection `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated SbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated ManagementResponseBodyResultSbfmLikelyAutomated `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots SbfmVerifiedBots `json:"sbfm_verified_bots"`
	// This field can have the runtime type of
	// [BotFightModeConfigStaleZoneConfiguration],
	// [SbfmDefinitelyConfigStaleZoneConfiguration],
	// [SbfmLikelyConfigStaleZoneConfiguration],
	// [BmSubscriptionConfigStaleZoneConfiguration].
	StaleZoneConfiguration interface{} `json:"stale_zone_configuration"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore bool `json:"suppress_session_score"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                             `json:"using_latest_model"`
	JSON             managementResponseBodyResultJSON `json:"-"`
	union            ManagementResponseBodyResultUnion
}

// managementResponseBodyResultJSON contains the JSON metadata for the struct
// [ManagementResponseBodyResult]
type managementResponseBodyResultJSON struct {
	AIBotsProtection             apijson.Field
	AutoUpdateModel              apijson.Field
	CrawlerProtection            apijson.Field
	EnableJs                     apijson.Field
	FightMode                    apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmLikelyAutomated          apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	StaleZoneConfiguration       apijson.Field
	SuppressSessionScore         apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r managementResponseBodyResultJSON) RawJSON() string {
	return r.raw
}

func (r *ManagementResponseBodyResult) UnmarshalJSON(data []byte) (err error) {
	*r = ManagementResponseBodyResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ManagementResponseBodyResultUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [BotFightModeConfig],
// [SbfmDefinitelyConfig], [SbfmLikelyConfig], [BmSubscriptionConfig].
func (r ManagementResponseBodyResult) AsUnion() ManagementResponseBodyResultUnion {
	return r.union
}

// Union satisfied by [BotFightModeConfig], [SbfmDefinitelyConfig],
// [SbfmLikelyConfig] or [BmSubscriptionConfig].
type ManagementResponseBodyResultUnion interface {
	implementsManagementResponseBodyResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ManagementResponseBodyResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BotFightModeConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SbfmDefinitelyConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SbfmLikelyConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BmSubscriptionConfig{}),
		},
	)
}

// Enable rule to block AI Scrapers and Crawlers. Please note the value
// `only_on_ad_pages` is currently not available for Enterprise customers.
type ManagementResponseBodyResultAIBotsProtection string

const (
	ManagementResponseBodyResultAIBotsProtectionBlock         ManagementResponseBodyResultAIBotsProtection = "block"
	ManagementResponseBodyResultAIBotsProtectionDisabled      ManagementResponseBodyResultAIBotsProtection = "disabled"
	ManagementResponseBodyResultAIBotsProtectionOnlyOnAdPages ManagementResponseBodyResultAIBotsProtection = "only_on_ad_pages"
)

func (r ManagementResponseBodyResultAIBotsProtection) IsKnown() bool {
	switch r {
	case ManagementResponseBodyResultAIBotsProtectionBlock, ManagementResponseBodyResultAIBotsProtectionDisabled, ManagementResponseBodyResultAIBotsProtectionOnlyOnAdPages:
		return true
	}
	return false
}

// Enable rule to punish AI Scrapers and Crawlers via a link maze.
type ManagementResponseBodyResultCrawlerProtection string

const (
	ManagementResponseBodyResultCrawlerProtectionEnabled  ManagementResponseBodyResultCrawlerProtection = "enabled"
	ManagementResponseBodyResultCrawlerProtectionDisabled ManagementResponseBodyResultCrawlerProtection = "disabled"
)

func (r ManagementResponseBodyResultCrawlerProtection) IsKnown() bool {
	switch r {
	case ManagementResponseBodyResultCrawlerProtectionEnabled, ManagementResponseBodyResultCrawlerProtectionDisabled:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type ManagementResponseBodyResultSbfmLikelyAutomated string

const (
	ManagementResponseBodyResultSbfmLikelyAutomatedAllow            ManagementResponseBodyResultSbfmLikelyAutomated = "allow"
	ManagementResponseBodyResultSbfmLikelyAutomatedBlock            ManagementResponseBodyResultSbfmLikelyAutomated = "block"
	ManagementResponseBodyResultSbfmLikelyAutomatedManagedChallenge ManagementResponseBodyResultSbfmLikelyAutomated = "managed_challenge"
)

func (r ManagementResponseBodyResultSbfmLikelyAutomated) IsKnown() bool {
	switch r {
	case ManagementResponseBodyResultSbfmLikelyAutomatedAllow, ManagementResponseBodyResultSbfmLikelyAutomatedBlock, ManagementResponseBodyResultSbfmLikelyAutomatedManagedChallenge:
		return true
	}
	return false
}

type MessagesBotManagementItem struct {
	Code             int64                           `json:"code,required"`
	Message          string                          `json:"message,required"`
	DocumentationURL string                          `json:"documentation_url"`
	Source           MessagesBotManagementItemSource `json:"source"`
	JSON             messagesBotManagementItemJSON   `json:"-"`
}

// messagesBotManagementItemJSON contains the JSON metadata for the struct
// [MessagesBotManagementItem]
type messagesBotManagementItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessagesBotManagementItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesBotManagementItemJSON) RawJSON() string {
	return r.raw
}

type MessagesBotManagementItemSource struct {
	Pointer string                              `json:"pointer"`
	JSON    messagesBotManagementItemSourceJSON `json:"-"`
}

// messagesBotManagementItemSourceJSON contains the JSON metadata for the struct
// [MessagesBotManagementItemSource]
type messagesBotManagementItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesBotManagementItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesBotManagementItemSourceJSON) RawJSON() string {
	return r.raw
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type SbfmDefinitelyAutomated string

const (
	SbfmDefinitelyAutomatedAllow            SbfmDefinitelyAutomated = "allow"
	SbfmDefinitelyAutomatedBlock            SbfmDefinitelyAutomated = "block"
	SbfmDefinitelyAutomatedManagedChallenge SbfmDefinitelyAutomated = "managed_challenge"
)

func (r SbfmDefinitelyAutomated) IsKnown() bool {
	switch r {
	case SbfmDefinitelyAutomatedAllow, SbfmDefinitelyAutomatedBlock, SbfmDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

type SbfmDefinitelyConfig struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection SbfmDefinitelyConfigAIBotsProtection `json:"ai_bots_protection"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection SbfmDefinitelyConfigCrawlerProtection `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated SbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots SbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that shows which unauthorized settings are currently active on
	// the zone. These settings typically result from upgrades or downgrades.
	StaleZoneConfiguration SbfmDefinitelyConfigStaleZoneConfiguration `json:"stale_zone_configuration"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                     `json:"using_latest_model"`
	JSON             sbfmDefinitelyConfigJSON `json:"-"`
}

// sbfmDefinitelyConfigJSON contains the JSON metadata for the struct
// [SbfmDefinitelyConfig]
type sbfmDefinitelyConfigJSON struct {
	AIBotsProtection             apijson.Field
	CrawlerProtection            apijson.Field
	EnableJs                     apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	StaleZoneConfiguration       apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SbfmDefinitelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sbfmDefinitelyConfigJSON) RawJSON() string {
	return r.raw
}

func (r SbfmDefinitelyConfig) implementsManagementResponseBodyResult() {}

// Enable rule to block AI Scrapers and Crawlers. Please note the value
// `only_on_ad_pages` is currently not available for Enterprise customers.
type SbfmDefinitelyConfigAIBotsProtection string

const (
	SbfmDefinitelyConfigAIBotsProtectionBlock         SbfmDefinitelyConfigAIBotsProtection = "block"
	SbfmDefinitelyConfigAIBotsProtectionDisabled      SbfmDefinitelyConfigAIBotsProtection = "disabled"
	SbfmDefinitelyConfigAIBotsProtectionOnlyOnAdPages SbfmDefinitelyConfigAIBotsProtection = "only_on_ad_pages"
)

func (r SbfmDefinitelyConfigAIBotsProtection) IsKnown() bool {
	switch r {
	case SbfmDefinitelyConfigAIBotsProtectionBlock, SbfmDefinitelyConfigAIBotsProtectionDisabled, SbfmDefinitelyConfigAIBotsProtectionOnlyOnAdPages:
		return true
	}
	return false
}

// Enable rule to punish AI Scrapers and Crawlers via a link maze.
type SbfmDefinitelyConfigCrawlerProtection string

const (
	SbfmDefinitelyConfigCrawlerProtectionEnabled  SbfmDefinitelyConfigCrawlerProtection = "enabled"
	SbfmDefinitelyConfigCrawlerProtectionDisabled SbfmDefinitelyConfigCrawlerProtection = "disabled"
)

func (r SbfmDefinitelyConfigCrawlerProtection) IsKnown() bool {
	switch r {
	case SbfmDefinitelyConfigCrawlerProtectionEnabled, SbfmDefinitelyConfigCrawlerProtectionDisabled:
		return true
	}
	return false
}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type SbfmDefinitelyConfigStaleZoneConfiguration struct {
	// Indicates that the zone's Bot Fight Mode is turned on.
	FightMode bool `json:"fight_mode"`
	// Indicates that the zone's likely automated requests are being blocked or
	// challenged.
	SbfmLikelyAutomated string                                         `json:"sbfm_likely_automated"`
	JSON                sbfmDefinitelyConfigStaleZoneConfigurationJSON `json:"-"`
}

// sbfmDefinitelyConfigStaleZoneConfigurationJSON contains the JSON metadata for
// the struct [SbfmDefinitelyConfigStaleZoneConfiguration]
type sbfmDefinitelyConfigStaleZoneConfigurationJSON struct {
	FightMode           apijson.Field
	SbfmLikelyAutomated apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SbfmDefinitelyConfigStaleZoneConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sbfmDefinitelyConfigStaleZoneConfigurationJSON) RawJSON() string {
	return r.raw
}

type SbfmDefinitelyConfigParam struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection param.Field[SbfmDefinitelyConfigAIBotsProtection] `json:"ai_bots_protection"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection param.Field[SbfmDefinitelyConfigCrawlerProtection] `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated param.Field[SbfmDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots param.Field[SbfmVerifiedBots] `json:"sbfm_verified_bots"`
}

func (r SbfmDefinitelyConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SbfmDefinitelyConfigParam) implementsZoneBotManagementUpdateParamsBodyUnion() {}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type SbfmDefinitelyConfigStaleZoneConfigurationParam struct {
	// Indicates that the zone's Bot Fight Mode is turned on.
	FightMode param.Field[bool] `json:"fight_mode"`
	// Indicates that the zone's likely automated requests are being blocked or
	// challenged.
	SbfmLikelyAutomated param.Field[string] `json:"sbfm_likely_automated"`
}

func (r SbfmDefinitelyConfigStaleZoneConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SbfmLikelyConfig struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection SbfmLikelyConfigAIBotsProtection `json:"ai_bots_protection"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection SbfmLikelyConfigCrawlerProtection `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated SbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated SbfmLikelyConfigSbfmLikelyAutomated `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots SbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that shows which unauthorized settings are currently active on
	// the zone. These settings typically result from upgrades or downgrades.
	StaleZoneConfiguration SbfmLikelyConfigStaleZoneConfiguration `json:"stale_zone_configuration"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                 `json:"using_latest_model"`
	JSON             sbfmLikelyConfigJSON `json:"-"`
}

// sbfmLikelyConfigJSON contains the JSON metadata for the struct
// [SbfmLikelyConfig]
type sbfmLikelyConfigJSON struct {
	AIBotsProtection             apijson.Field
	CrawlerProtection            apijson.Field
	EnableJs                     apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmLikelyAutomated          apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	StaleZoneConfiguration       apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SbfmLikelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sbfmLikelyConfigJSON) RawJSON() string {
	return r.raw
}

func (r SbfmLikelyConfig) implementsManagementResponseBodyResult() {}

// Enable rule to block AI Scrapers and Crawlers. Please note the value
// `only_on_ad_pages` is currently not available for Enterprise customers.
type SbfmLikelyConfigAIBotsProtection string

const (
	SbfmLikelyConfigAIBotsProtectionBlock         SbfmLikelyConfigAIBotsProtection = "block"
	SbfmLikelyConfigAIBotsProtectionDisabled      SbfmLikelyConfigAIBotsProtection = "disabled"
	SbfmLikelyConfigAIBotsProtectionOnlyOnAdPages SbfmLikelyConfigAIBotsProtection = "only_on_ad_pages"
)

func (r SbfmLikelyConfigAIBotsProtection) IsKnown() bool {
	switch r {
	case SbfmLikelyConfigAIBotsProtectionBlock, SbfmLikelyConfigAIBotsProtectionDisabled, SbfmLikelyConfigAIBotsProtectionOnlyOnAdPages:
		return true
	}
	return false
}

// Enable rule to punish AI Scrapers and Crawlers via a link maze.
type SbfmLikelyConfigCrawlerProtection string

const (
	SbfmLikelyConfigCrawlerProtectionEnabled  SbfmLikelyConfigCrawlerProtection = "enabled"
	SbfmLikelyConfigCrawlerProtectionDisabled SbfmLikelyConfigCrawlerProtection = "disabled"
)

func (r SbfmLikelyConfigCrawlerProtection) IsKnown() bool {
	switch r {
	case SbfmLikelyConfigCrawlerProtectionEnabled, SbfmLikelyConfigCrawlerProtectionDisabled:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type SbfmLikelyConfigSbfmLikelyAutomated string

const (
	SbfmLikelyConfigSbfmLikelyAutomatedAllow            SbfmLikelyConfigSbfmLikelyAutomated = "allow"
	SbfmLikelyConfigSbfmLikelyAutomatedBlock            SbfmLikelyConfigSbfmLikelyAutomated = "block"
	SbfmLikelyConfigSbfmLikelyAutomatedManagedChallenge SbfmLikelyConfigSbfmLikelyAutomated = "managed_challenge"
)

func (r SbfmLikelyConfigSbfmLikelyAutomated) IsKnown() bool {
	switch r {
	case SbfmLikelyConfigSbfmLikelyAutomatedAllow, SbfmLikelyConfigSbfmLikelyAutomatedBlock, SbfmLikelyConfigSbfmLikelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type SbfmLikelyConfigStaleZoneConfiguration struct {
	// Indicates that the zone's Bot Fight Mode is turned on.
	FightMode bool                                       `json:"fight_mode"`
	JSON      sbfmLikelyConfigStaleZoneConfigurationJSON `json:"-"`
}

// sbfmLikelyConfigStaleZoneConfigurationJSON contains the JSON metadata for the
// struct [SbfmLikelyConfigStaleZoneConfiguration]
type sbfmLikelyConfigStaleZoneConfigurationJSON struct {
	FightMode   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SbfmLikelyConfigStaleZoneConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sbfmLikelyConfigStaleZoneConfigurationJSON) RawJSON() string {
	return r.raw
}

type SbfmLikelyConfigParam struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection param.Field[SbfmLikelyConfigAIBotsProtection] `json:"ai_bots_protection"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection param.Field[SbfmLikelyConfigCrawlerProtection] `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated param.Field[SbfmDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated param.Field[SbfmLikelyConfigSbfmLikelyAutomated] `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots param.Field[SbfmVerifiedBots] `json:"sbfm_verified_bots"`
}

func (r SbfmLikelyConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SbfmLikelyConfigParam) implementsZoneBotManagementUpdateParamsBodyUnion() {}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type SbfmLikelyConfigStaleZoneConfigurationParam struct {
	// Indicates that the zone's Bot Fight Mode is turned on.
	FightMode param.Field[bool] `json:"fight_mode"`
}

func (r SbfmLikelyConfigStaleZoneConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type SbfmVerifiedBots string

const (
	SbfmVerifiedBotsAllow SbfmVerifiedBots = "allow"
	SbfmVerifiedBotsBlock SbfmVerifiedBots = "block"
)

func (r SbfmVerifiedBots) IsKnown() bool {
	switch r {
	case SbfmVerifiedBotsAllow, SbfmVerifiedBotsBlock:
		return true
	}
	return false
}

type ZoneBotManagementUpdateParams struct {
	Body ZoneBotManagementUpdateParamsBodyUnion `json:"body,required"`
}

func (r ZoneBotManagementUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneBotManagementUpdateParamsBody struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection param.Field[ZoneBotManagementUpdateParamsBodyAIBotsProtection] `json:"ai_bots_protection"`
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel param.Field[bool] `json:"auto_update_model"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection param.Field[ZoneBotManagementUpdateParamsBodyCrawlerProtection] `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode param.Field[bool] `json:"fight_mode"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated param.Field[SbfmDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated param.Field[ZoneBotManagementUpdateParamsBodySbfmLikelyAutomated] `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots       param.Field[SbfmVerifiedBots] `json:"sbfm_verified_bots"`
	StaleZoneConfiguration param.Field[interface{}]      `json:"stale_zone_configuration"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore param.Field[bool] `json:"suppress_session_score"`
}

func (r ZoneBotManagementUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneBotManagementUpdateParamsBody) implementsZoneBotManagementUpdateParamsBodyUnion() {}

// Satisfied by [BotFightModeConfigParam], [SbfmDefinitelyConfigParam],
// [SbfmLikelyConfigParam], [BmSubscriptionConfigParam],
// [ZoneBotManagementUpdateParamsBody].
type ZoneBotManagementUpdateParamsBodyUnion interface {
	implementsZoneBotManagementUpdateParamsBodyUnion()
}

// Enable rule to block AI Scrapers and Crawlers. Please note the value
// `only_on_ad_pages` is currently not available for Enterprise customers.
type ZoneBotManagementUpdateParamsBodyAIBotsProtection string

const (
	ZoneBotManagementUpdateParamsBodyAIBotsProtectionBlock         ZoneBotManagementUpdateParamsBodyAIBotsProtection = "block"
	ZoneBotManagementUpdateParamsBodyAIBotsProtectionDisabled      ZoneBotManagementUpdateParamsBodyAIBotsProtection = "disabled"
	ZoneBotManagementUpdateParamsBodyAIBotsProtectionOnlyOnAdPages ZoneBotManagementUpdateParamsBodyAIBotsProtection = "only_on_ad_pages"
)

func (r ZoneBotManagementUpdateParamsBodyAIBotsProtection) IsKnown() bool {
	switch r {
	case ZoneBotManagementUpdateParamsBodyAIBotsProtectionBlock, ZoneBotManagementUpdateParamsBodyAIBotsProtectionDisabled, ZoneBotManagementUpdateParamsBodyAIBotsProtectionOnlyOnAdPages:
		return true
	}
	return false
}

// Enable rule to punish AI Scrapers and Crawlers via a link maze.
type ZoneBotManagementUpdateParamsBodyCrawlerProtection string

const (
	ZoneBotManagementUpdateParamsBodyCrawlerProtectionEnabled  ZoneBotManagementUpdateParamsBodyCrawlerProtection = "enabled"
	ZoneBotManagementUpdateParamsBodyCrawlerProtectionDisabled ZoneBotManagementUpdateParamsBodyCrawlerProtection = "disabled"
)

func (r ZoneBotManagementUpdateParamsBodyCrawlerProtection) IsKnown() bool {
	switch r {
	case ZoneBotManagementUpdateParamsBodyCrawlerProtectionEnabled, ZoneBotManagementUpdateParamsBodyCrawlerProtectionDisabled:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type ZoneBotManagementUpdateParamsBodySbfmLikelyAutomated string

const (
	ZoneBotManagementUpdateParamsBodySbfmLikelyAutomatedAllow            ZoneBotManagementUpdateParamsBodySbfmLikelyAutomated = "allow"
	ZoneBotManagementUpdateParamsBodySbfmLikelyAutomatedBlock            ZoneBotManagementUpdateParamsBodySbfmLikelyAutomated = "block"
	ZoneBotManagementUpdateParamsBodySbfmLikelyAutomatedManagedChallenge ZoneBotManagementUpdateParamsBodySbfmLikelyAutomated = "managed_challenge"
)

func (r ZoneBotManagementUpdateParamsBodySbfmLikelyAutomated) IsKnown() bool {
	switch r {
	case ZoneBotManagementUpdateParamsBodySbfmLikelyAutomatedAllow, ZoneBotManagementUpdateParamsBodySbfmLikelyAutomatedBlock, ZoneBotManagementUpdateParamsBodySbfmLikelyAutomatedManagedChallenge:
		return true
	}
	return false
}
