// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/tidwall/gjson"
)

// ZoneSettingZarazConfigService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSettingZarazConfigService] method instead.
type ZoneSettingZarazConfigService struct {
	Options []option.RequestOption
}

// NewZoneSettingZarazConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingZarazConfigService(opts ...option.RequestOption) (r *ZoneSettingZarazConfigService) {
	r = &ZoneSettingZarazConfigService{}
	r.Options = opts
	return
}

// Gets latest Zaraz configuration for a zone. It can be preview or published
// configuration, whichever was the last updated. Secret variables values will not
// be included.
func (r *ZoneSettingZarazConfigService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZarazConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/config", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates Zaraz configuration for a zone.
func (r *ZoneSettingZarazConfigService) Update(ctx context.Context, zoneID string, body ZoneSettingZarazConfigUpdateParams, opts ...option.RequestOption) (res *ZarazConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/config", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Zaraz configuration
type ZarazConfigBase struct {
	// Data layer compatibility mode enabled.
	DataLayer bool `json:"dataLayer,required"`
	// The key for Zaraz debug mode.
	DebugKey string `json:"debugKey,required"`
	// General Zaraz settings.
	Settings ZarazConfigBaseSettings `json:"settings,required"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers map[string]ZarazConfigBaseTrigger `json:"triggers,required"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables map[string]ZarazConfigBaseVariable `json:"variables,required"`
	// Zaraz internal version of the config.
	ZarazVersion int64 `json:"zarazVersion,required"`
	// Cloudflare Monitoring settings.
	Analytics ZarazConfigBaseAnalytics `json:"analytics"`
	// Consent management configuration.
	Consent ZarazConfigBaseConsent `json:"consent"`
	// Single Page Application support enabled.
	HistoryChange bool                `json:"historyChange"`
	JSON          zarazConfigBaseJSON `json:"-"`
}

// zarazConfigBaseJSON contains the JSON metadata for the struct [ZarazConfigBase]
type zarazConfigBaseJSON struct {
	DataLayer     apijson.Field
	DebugKey      apijson.Field
	Settings      apijson.Field
	Triggers      apijson.Field
	Variables     apijson.Field
	ZarazVersion  apijson.Field
	Analytics     apijson.Field
	Consent       apijson.Field
	HistoryChange apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZarazConfigBase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseJSON) RawJSON() string {
	return r.raw
}

// General Zaraz settings.
type ZarazConfigBaseSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript bool `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher ZarazConfigBaseSettingsContextEnricher `json:"contextEnricher"`
	// The domain Zaraz will use for writing and reading its cookies.
	CookieDomain string `json:"cookieDomain"`
	// Ecommerce API enabled.
	Ecommerce bool `json:"ecommerce"`
	// Custom endpoint for server-side track events.
	EventsAPIPath string `json:"eventsApiPath"`
	// Hiding external referrer URL enabled.
	HideExternalReferer bool `json:"hideExternalReferer"`
	// Trimming IP address enabled.
	HideIPAddress bool `json:"hideIPAddress"`
	// Removing URL query params enabled.
	HideQueryParams bool `json:"hideQueryParams"`
	// Removing sensitive data from User Aagent string enabled.
	HideUserAgent bool `json:"hideUserAgent"`
	// Custom endpoint for Zaraz init script.
	InitPath string `json:"initPath"`
	// Injection of Zaraz scripts into iframes enabled.
	InjectIframes bool `json:"injectIframes"`
	// Custom path for Managed Components server functionalities.
	McRootPath string `json:"mcRootPath"`
	// Custom endpoint for Zaraz main script.
	ScriptPath string `json:"scriptPath"`
	// Custom endpoint for Zaraz tracking requests.
	TrackPath string                      `json:"trackPath"`
	JSON      zarazConfigBaseSettingsJSON `json:"-"`
}

// zarazConfigBaseSettingsJSON contains the JSON metadata for the struct
// [ZarazConfigBaseSettings]
type zarazConfigBaseSettingsJSON struct {
	AutoInjectScript    apijson.Field
	ContextEnricher     apijson.Field
	CookieDomain        apijson.Field
	Ecommerce           apijson.Field
	EventsAPIPath       apijson.Field
	HideExternalReferer apijson.Field
	HideIPAddress       apijson.Field
	HideQueryParams     apijson.Field
	HideUserAgent       apijson.Field
	InitPath            apijson.Field
	InjectIframes       apijson.Field
	McRootPath          apijson.Field
	ScriptPath          apijson.Field
	TrackPath           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZarazConfigBaseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseSettingsJSON) RawJSON() string {
	return r.raw
}

// Details of the worker that receives and edits Zaraz Context object.
type ZarazConfigBaseSettingsContextEnricher struct {
	EscapedWorkerName string                                     `json:"escapedWorkerName,required"`
	WorkerTag         string                                     `json:"workerTag,required"`
	JSON              zarazConfigBaseSettingsContextEnricherJSON `json:"-"`
}

// zarazConfigBaseSettingsContextEnricherJSON contains the JSON metadata for the
// struct [ZarazConfigBaseSettingsContextEnricher]
type zarazConfigBaseSettingsContextEnricherJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazConfigBaseSettingsContextEnricher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseSettingsContextEnricherJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseTrigger struct {
	// Rules defining when the trigger is not fired.
	ExcludeRules []ZarazConfigBaseTriggersExcludeRule `json:"excludeRules,required"`
	// Rules defining when the trigger is fired.
	LoadRules []ZarazConfigBaseTriggersLoadRule `json:"loadRules,required"`
	// Trigger name.
	Name string `json:"name,required"`
	// Trigger description.
	Description string                        `json:"description"`
	System      ZarazConfigBaseTriggersSystem `json:"system"`
	JSON        zarazConfigBaseTriggerJSON    `json:"-"`
}

// zarazConfigBaseTriggerJSON contains the JSON metadata for the struct
// [ZarazConfigBaseTrigger]
type zarazConfigBaseTriggerJSON struct {
	ExcludeRules apijson.Field
	LoadRules    apijson.Field
	Name         apijson.Field
	Description  apijson.Field
	System       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZarazConfigBaseTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggerJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseTriggersExcludeRule struct {
	ID     string                                    `json:"id,required"`
	Action ZarazConfigBaseTriggersExcludeRulesAction `json:"action"`
	Match  string                                    `json:"match"`
	Op     ZarazConfigBaseTriggersExcludeRulesOp     `json:"op"`
	// This field can have the runtime type of
	// [ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettings],
	// [ZarazConfigBaseTriggersExcludeRulesZarazTimerRuleSettings],
	// [ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleSettings],
	// [ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleSettings],
	// [ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleSettings],
	// [ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleSettings].
	Settings interface{}                            `json:"settings"`
	Value    string                                 `json:"value"`
	JSON     zarazConfigBaseTriggersExcludeRuleJSON `json:"-"`
	union    ZarazConfigBaseTriggersExcludeRulesUnion
}

// zarazConfigBaseTriggersExcludeRuleJSON contains the JSON metadata for the struct
// [ZarazConfigBaseTriggersExcludeRule]
type zarazConfigBaseTriggersExcludeRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Settings    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r zarazConfigBaseTriggersExcludeRuleJSON) RawJSON() string {
	return r.raw
}

func (r *ZarazConfigBaseTriggersExcludeRule) UnmarshalJSON(data []byte) (err error) {
	*r = ZarazConfigBaseTriggersExcludeRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZarazConfigBaseTriggersExcludeRulesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ZarazConfigBaseTriggersExcludeRulesZarazLoadRule],
// [ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRule],
// [ZarazConfigBaseTriggersExcludeRulesZarazTimerRule],
// [ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRule],
// [ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRule],
// [ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRule],
// [ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRule].
func (r ZarazConfigBaseTriggersExcludeRule) AsUnion() ZarazConfigBaseTriggersExcludeRulesUnion {
	return r.union
}

// Union satisfied by [ZarazConfigBaseTriggersExcludeRulesZarazLoadRule],
// [ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRule],
// [ZarazConfigBaseTriggersExcludeRulesZarazTimerRule],
// [ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRule],
// [ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRule],
// [ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRule] or
// [ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRule].
type ZarazConfigBaseTriggersExcludeRulesUnion interface {
	implementsZarazConfigBaseTriggersExcludeRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigBaseTriggersExcludeRulesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigBaseTriggersExcludeRulesZarazLoadRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigBaseTriggersExcludeRulesZarazTimerRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRule{}),
		},
	)
}

type ZarazConfigBaseTriggersExcludeRulesZarazLoadRule struct {
	ID    string                                               `json:"id,required"`
	Match string                                               `json:"match,required"`
	Op    ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                               `json:"value,required"`
	JSON  zarazConfigBaseTriggersExcludeRulesZarazLoadRuleJSON `json:"-"`
}

// zarazConfigBaseTriggersExcludeRulesZarazLoadRuleJSON contains the JSON metadata
// for the struct [ZarazConfigBaseTriggersExcludeRulesZarazLoadRule]
type zarazConfigBaseTriggersExcludeRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersExcludeRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersExcludeRulesZarazLoadRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseTriggersExcludeRulesZarazLoadRule) implementsZarazConfigBaseTriggersExcludeRule() {
}

type ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOp string

const (
	ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpContains           ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOp = "CONTAINS"
	ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpEquals             ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOp = "EQUALS"
	ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpStartsWith         ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpEndsWith           ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpMatchRegex         ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex      ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpGreaterThan        ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpLessThan           ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual    ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

func (r ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOp) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpContains, ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpEquals, ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpStartsWith, ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpEndsWith, ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpMatchRegex, ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex, ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpGreaterThan, ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual, ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpLessThan, ZarazConfigBaseTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRule struct {
	ID       string                                                            `json:"id,required"`
	Action   ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRule]
type zarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRule) implementsZarazConfigBaseTriggersExcludeRule() {
}

type ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleAction string

const (
	ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleActionClickListener ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleAction = "clickListener"
)

func (r ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleActionClickListener:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                `json:"selector,required"`
	Type        ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                 `json:"waitForTags,required"`
	JSON        zarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettings]
type zarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettingsType string

const (
	ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCss   ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "css"
)

func (r ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettingsType) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath, ZarazConfigBaseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCss:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersExcludeRulesZarazTimerRule struct {
	ID       string                                                    `json:"id,required"`
	Action   ZarazConfigBaseTriggersExcludeRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazConfigBaseTriggersExcludeRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazConfigBaseTriggersExcludeRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazConfigBaseTriggersExcludeRulesZarazTimerRuleJSON contains the JSON metadata
// for the struct [ZarazConfigBaseTriggersExcludeRulesZarazTimerRule]
type zarazConfigBaseTriggersExcludeRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersExcludeRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersExcludeRulesZarazTimerRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseTriggersExcludeRulesZarazTimerRule) implementsZarazConfigBaseTriggersExcludeRule() {
}

type ZarazConfigBaseTriggersExcludeRulesZarazTimerRuleAction string

const (
	ZarazConfigBaseTriggersExcludeRulesZarazTimerRuleActionTimer ZarazConfigBaseTriggersExcludeRulesZarazTimerRuleAction = "timer"
)

func (r ZarazConfigBaseTriggersExcludeRulesZarazTimerRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersExcludeRulesZarazTimerRuleActionTimer:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersExcludeRulesZarazTimerRuleSettings struct {
	Interval int64                                                         `json:"interval,required"`
	Limit    int64                                                         `json:"limit,required"`
	JSON     zarazConfigBaseTriggersExcludeRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazConfigBaseTriggersExcludeRulesZarazTimerRuleSettingsJSON contains the JSON
// metadata for the struct
// [ZarazConfigBaseTriggersExcludeRulesZarazTimerRuleSettings]
type zarazConfigBaseTriggersExcludeRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersExcludeRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersExcludeRulesZarazTimerRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRule struct {
	ID       string                                                             `json:"id,required"`
	Action   ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRule]
type zarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRule) implementsZarazConfigBaseTriggersExcludeRule() {
}

type ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleAction string

const (
	ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleAction = "formSubmission"
)

func (r ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                 `json:"selector,required"`
	Validate bool                                                                   `json:"validate,required"`
	JSON     zarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleSettings]
type zarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRule struct {
	ID       string                                                            `json:"id,required"`
	Action   ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRule]
type zarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRule) implementsZarazConfigBaseTriggersExcludeRule() {
}

type ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleAction string

const (
	ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleAction = "variableMatch"
)

func (r ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                `json:"match,required"`
	Variable string                                                                `json:"variable,required"`
	JSON     zarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleSettings]
type zarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRule struct {
	ID       string                                                          `json:"id,required"`
	Action   ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRule]
type zarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRule) implementsZarazConfigBaseTriggersExcludeRule() {
}

type ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleAction string

const (
	ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleAction = "scrollDepth"
)

func (r ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                              `json:"positions,required"`
	JSON      zarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON contains the
// JSON metadata for the struct
// [ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleSettings]
type zarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRule struct {
	ID       string                                                                `json:"id,required"`
	Action   ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRule]
type zarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRule) implementsZarazConfigBaseTriggersExcludeRule() {
}

type ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleAction string

const (
	ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

func (r ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                    `json:"selector,required"`
	JSON     zarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleSettings]
type zarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseTriggersExcludeRulesAction string

const (
	ZarazConfigBaseTriggersExcludeRulesActionClickListener     ZarazConfigBaseTriggersExcludeRulesAction = "clickListener"
	ZarazConfigBaseTriggersExcludeRulesActionTimer             ZarazConfigBaseTriggersExcludeRulesAction = "timer"
	ZarazConfigBaseTriggersExcludeRulesActionFormSubmission    ZarazConfigBaseTriggersExcludeRulesAction = "formSubmission"
	ZarazConfigBaseTriggersExcludeRulesActionVariableMatch     ZarazConfigBaseTriggersExcludeRulesAction = "variableMatch"
	ZarazConfigBaseTriggersExcludeRulesActionScrollDepth       ZarazConfigBaseTriggersExcludeRulesAction = "scrollDepth"
	ZarazConfigBaseTriggersExcludeRulesActionElementVisibility ZarazConfigBaseTriggersExcludeRulesAction = "elementVisibility"
)

func (r ZarazConfigBaseTriggersExcludeRulesAction) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersExcludeRulesActionClickListener, ZarazConfigBaseTriggersExcludeRulesActionTimer, ZarazConfigBaseTriggersExcludeRulesActionFormSubmission, ZarazConfigBaseTriggersExcludeRulesActionVariableMatch, ZarazConfigBaseTriggersExcludeRulesActionScrollDepth, ZarazConfigBaseTriggersExcludeRulesActionElementVisibility:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersExcludeRulesOp string

const (
	ZarazConfigBaseTriggersExcludeRulesOpContains           ZarazConfigBaseTriggersExcludeRulesOp = "CONTAINS"
	ZarazConfigBaseTriggersExcludeRulesOpEquals             ZarazConfigBaseTriggersExcludeRulesOp = "EQUALS"
	ZarazConfigBaseTriggersExcludeRulesOpStartsWith         ZarazConfigBaseTriggersExcludeRulesOp = "STARTS_WITH"
	ZarazConfigBaseTriggersExcludeRulesOpEndsWith           ZarazConfigBaseTriggersExcludeRulesOp = "ENDS_WITH"
	ZarazConfigBaseTriggersExcludeRulesOpMatchRegex         ZarazConfigBaseTriggersExcludeRulesOp = "MATCH_REGEX"
	ZarazConfigBaseTriggersExcludeRulesOpNotMatchRegex      ZarazConfigBaseTriggersExcludeRulesOp = "NOT_MATCH_REGEX"
	ZarazConfigBaseTriggersExcludeRulesOpGreaterThan        ZarazConfigBaseTriggersExcludeRulesOp = "GREATER_THAN"
	ZarazConfigBaseTriggersExcludeRulesOpGreaterThanOrEqual ZarazConfigBaseTriggersExcludeRulesOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigBaseTriggersExcludeRulesOpLessThan           ZarazConfigBaseTriggersExcludeRulesOp = "LESS_THAN"
	ZarazConfigBaseTriggersExcludeRulesOpLessThanOrEqual    ZarazConfigBaseTriggersExcludeRulesOp = "LESS_THAN_OR_EQUAL"
)

func (r ZarazConfigBaseTriggersExcludeRulesOp) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersExcludeRulesOpContains, ZarazConfigBaseTriggersExcludeRulesOpEquals, ZarazConfigBaseTriggersExcludeRulesOpStartsWith, ZarazConfigBaseTriggersExcludeRulesOpEndsWith, ZarazConfigBaseTriggersExcludeRulesOpMatchRegex, ZarazConfigBaseTriggersExcludeRulesOpNotMatchRegex, ZarazConfigBaseTriggersExcludeRulesOpGreaterThan, ZarazConfigBaseTriggersExcludeRulesOpGreaterThanOrEqual, ZarazConfigBaseTriggersExcludeRulesOpLessThan, ZarazConfigBaseTriggersExcludeRulesOpLessThanOrEqual:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersLoadRule struct {
	ID     string                                 `json:"id,required"`
	Action ZarazConfigBaseTriggersLoadRulesAction `json:"action"`
	Match  string                                 `json:"match"`
	Op     ZarazConfigBaseTriggersLoadRulesOp     `json:"op"`
	// This field can have the runtime type of
	// [ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettings],
	// [ZarazConfigBaseTriggersLoadRulesZarazTimerRuleSettings],
	// [ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleSettings],
	// [ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleSettings],
	// [ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleSettings],
	// [ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleSettings].
	Settings interface{}                         `json:"settings"`
	Value    string                              `json:"value"`
	JSON     zarazConfigBaseTriggersLoadRuleJSON `json:"-"`
	union    ZarazConfigBaseTriggersLoadRulesUnion
}

// zarazConfigBaseTriggersLoadRuleJSON contains the JSON metadata for the struct
// [ZarazConfigBaseTriggersLoadRule]
type zarazConfigBaseTriggersLoadRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Settings    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r zarazConfigBaseTriggersLoadRuleJSON) RawJSON() string {
	return r.raw
}

func (r *ZarazConfigBaseTriggersLoadRule) UnmarshalJSON(data []byte) (err error) {
	*r = ZarazConfigBaseTriggersLoadRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZarazConfigBaseTriggersLoadRulesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ZarazConfigBaseTriggersLoadRulesZarazLoadRule],
// [ZarazConfigBaseTriggersLoadRulesZarazClickListenerRule],
// [ZarazConfigBaseTriggersLoadRulesZarazTimerRule],
// [ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRule],
// [ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRule],
// [ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRule],
// [ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRule].
func (r ZarazConfigBaseTriggersLoadRule) AsUnion() ZarazConfigBaseTriggersLoadRulesUnion {
	return r.union
}

// Union satisfied by [ZarazConfigBaseTriggersLoadRulesZarazLoadRule],
// [ZarazConfigBaseTriggersLoadRulesZarazClickListenerRule],
// [ZarazConfigBaseTriggersLoadRulesZarazTimerRule],
// [ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRule],
// [ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRule],
// [ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRule] or
// [ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRule].
type ZarazConfigBaseTriggersLoadRulesUnion interface {
	implementsZarazConfigBaseTriggersLoadRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigBaseTriggersLoadRulesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigBaseTriggersLoadRulesZarazLoadRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigBaseTriggersLoadRulesZarazClickListenerRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigBaseTriggersLoadRulesZarazTimerRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRule{}),
		},
	)
}

type ZarazConfigBaseTriggersLoadRulesZarazLoadRule struct {
	ID    string                                            `json:"id,required"`
	Match string                                            `json:"match,required"`
	Op    ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                            `json:"value,required"`
	JSON  zarazConfigBaseTriggersLoadRulesZarazLoadRuleJSON `json:"-"`
}

// zarazConfigBaseTriggersLoadRulesZarazLoadRuleJSON contains the JSON metadata for
// the struct [ZarazConfigBaseTriggersLoadRulesZarazLoadRule]
type zarazConfigBaseTriggersLoadRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersLoadRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersLoadRulesZarazLoadRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseTriggersLoadRulesZarazLoadRule) implementsZarazConfigBaseTriggersLoadRule() {}

type ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOp string

const (
	ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpContains           ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOp = "CONTAINS"
	ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpEquals             ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOp = "EQUALS"
	ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpStartsWith         ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpEndsWith           ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpMatchRegex         ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpNotMatchRegex      ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpGreaterThan        ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpLessThan           ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual    ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

func (r ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOp) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpContains, ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpEquals, ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpStartsWith, ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpEndsWith, ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpMatchRegex, ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpNotMatchRegex, ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpGreaterThan, ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual, ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpLessThan, ZarazConfigBaseTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersLoadRulesZarazClickListenerRule struct {
	ID       string                                                         `json:"id,required"`
	Action   ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazConfigBaseTriggersLoadRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazConfigBaseTriggersLoadRulesZarazClickListenerRuleJSON contains the JSON
// metadata for the struct [ZarazConfigBaseTriggersLoadRulesZarazClickListenerRule]
type zarazConfigBaseTriggersLoadRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersLoadRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersLoadRulesZarazClickListenerRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseTriggersLoadRulesZarazClickListenerRule) implementsZarazConfigBaseTriggersLoadRule() {
}

type ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleAction string

const (
	ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleActionClickListener ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleAction = "clickListener"
)

func (r ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleActionClickListener:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                             `json:"selector,required"`
	Type        ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                              `json:"waitForTags,required"`
	JSON        zarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettingsJSON contains the
// JSON metadata for the struct
// [ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettings]
type zarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettingsType string

const (
	ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettingsTypeCss   ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettingsType = "css"
)

func (r ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettingsType) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath, ZarazConfigBaseTriggersLoadRulesZarazClickListenerRuleSettingsTypeCss:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersLoadRulesZarazTimerRule struct {
	ID       string                                                 `json:"id,required"`
	Action   ZarazConfigBaseTriggersLoadRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazConfigBaseTriggersLoadRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazConfigBaseTriggersLoadRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazConfigBaseTriggersLoadRulesZarazTimerRuleJSON contains the JSON metadata
// for the struct [ZarazConfigBaseTriggersLoadRulesZarazTimerRule]
type zarazConfigBaseTriggersLoadRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersLoadRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersLoadRulesZarazTimerRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseTriggersLoadRulesZarazTimerRule) implementsZarazConfigBaseTriggersLoadRule() {}

type ZarazConfigBaseTriggersLoadRulesZarazTimerRuleAction string

const (
	ZarazConfigBaseTriggersLoadRulesZarazTimerRuleActionTimer ZarazConfigBaseTriggersLoadRulesZarazTimerRuleAction = "timer"
)

func (r ZarazConfigBaseTriggersLoadRulesZarazTimerRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersLoadRulesZarazTimerRuleActionTimer:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersLoadRulesZarazTimerRuleSettings struct {
	Interval int64                                                      `json:"interval,required"`
	Limit    int64                                                      `json:"limit,required"`
	JSON     zarazConfigBaseTriggersLoadRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazConfigBaseTriggersLoadRulesZarazTimerRuleSettingsJSON contains the JSON
// metadata for the struct [ZarazConfigBaseTriggersLoadRulesZarazTimerRuleSettings]
type zarazConfigBaseTriggersLoadRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersLoadRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersLoadRulesZarazTimerRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRule struct {
	ID       string                                                          `json:"id,required"`
	Action   ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRule]
type zarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRule) implementsZarazConfigBaseTriggersLoadRule() {
}

type ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleAction string

const (
	ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleAction = "formSubmission"
)

func (r ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                              `json:"selector,required"`
	Validate bool                                                                `json:"validate,required"`
	JSON     zarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON contains the
// JSON metadata for the struct
// [ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleSettings]
type zarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRule struct {
	ID       string                                                         `json:"id,required"`
	Action   ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleJSON contains the JSON
// metadata for the struct [ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRule]
type zarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRule) implementsZarazConfigBaseTriggersLoadRule() {
}

type ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleAction string

const (
	ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleAction = "variableMatch"
)

func (r ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                             `json:"match,required"`
	Variable string                                                             `json:"variable,required"`
	JSON     zarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON contains the
// JSON metadata for the struct
// [ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleSettings]
type zarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRule struct {
	ID       string                                                       `json:"id,required"`
	Action   ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleJSON contains the JSON
// metadata for the struct [ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRule]
type zarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRule) implementsZarazConfigBaseTriggersLoadRule() {
}

type ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleAction string

const (
	ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleAction = "scrollDepth"
)

func (r ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                           `json:"positions,required"`
	JSON      zarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON contains the
// JSON metadata for the struct
// [ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleSettings]
type zarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRule struct {
	ID       string                                                             `json:"id,required"`
	Action   ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRule]
type zarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRule) implementsZarazConfigBaseTriggersLoadRule() {
}

type ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleAction string

const (
	ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

func (r ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                 `json:"selector,required"`
	JSON     zarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleSettings]
type zarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseTriggersLoadRulesAction string

const (
	ZarazConfigBaseTriggersLoadRulesActionClickListener     ZarazConfigBaseTriggersLoadRulesAction = "clickListener"
	ZarazConfigBaseTriggersLoadRulesActionTimer             ZarazConfigBaseTriggersLoadRulesAction = "timer"
	ZarazConfigBaseTriggersLoadRulesActionFormSubmission    ZarazConfigBaseTriggersLoadRulesAction = "formSubmission"
	ZarazConfigBaseTriggersLoadRulesActionVariableMatch     ZarazConfigBaseTriggersLoadRulesAction = "variableMatch"
	ZarazConfigBaseTriggersLoadRulesActionScrollDepth       ZarazConfigBaseTriggersLoadRulesAction = "scrollDepth"
	ZarazConfigBaseTriggersLoadRulesActionElementVisibility ZarazConfigBaseTriggersLoadRulesAction = "elementVisibility"
)

func (r ZarazConfigBaseTriggersLoadRulesAction) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersLoadRulesActionClickListener, ZarazConfigBaseTriggersLoadRulesActionTimer, ZarazConfigBaseTriggersLoadRulesActionFormSubmission, ZarazConfigBaseTriggersLoadRulesActionVariableMatch, ZarazConfigBaseTriggersLoadRulesActionScrollDepth, ZarazConfigBaseTriggersLoadRulesActionElementVisibility:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersLoadRulesOp string

const (
	ZarazConfigBaseTriggersLoadRulesOpContains           ZarazConfigBaseTriggersLoadRulesOp = "CONTAINS"
	ZarazConfigBaseTriggersLoadRulesOpEquals             ZarazConfigBaseTriggersLoadRulesOp = "EQUALS"
	ZarazConfigBaseTriggersLoadRulesOpStartsWith         ZarazConfigBaseTriggersLoadRulesOp = "STARTS_WITH"
	ZarazConfigBaseTriggersLoadRulesOpEndsWith           ZarazConfigBaseTriggersLoadRulesOp = "ENDS_WITH"
	ZarazConfigBaseTriggersLoadRulesOpMatchRegex         ZarazConfigBaseTriggersLoadRulesOp = "MATCH_REGEX"
	ZarazConfigBaseTriggersLoadRulesOpNotMatchRegex      ZarazConfigBaseTriggersLoadRulesOp = "NOT_MATCH_REGEX"
	ZarazConfigBaseTriggersLoadRulesOpGreaterThan        ZarazConfigBaseTriggersLoadRulesOp = "GREATER_THAN"
	ZarazConfigBaseTriggersLoadRulesOpGreaterThanOrEqual ZarazConfigBaseTriggersLoadRulesOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigBaseTriggersLoadRulesOpLessThan           ZarazConfigBaseTriggersLoadRulesOp = "LESS_THAN"
	ZarazConfigBaseTriggersLoadRulesOpLessThanOrEqual    ZarazConfigBaseTriggersLoadRulesOp = "LESS_THAN_OR_EQUAL"
)

func (r ZarazConfigBaseTriggersLoadRulesOp) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersLoadRulesOpContains, ZarazConfigBaseTriggersLoadRulesOpEquals, ZarazConfigBaseTriggersLoadRulesOpStartsWith, ZarazConfigBaseTriggersLoadRulesOpEndsWith, ZarazConfigBaseTriggersLoadRulesOpMatchRegex, ZarazConfigBaseTriggersLoadRulesOpNotMatchRegex, ZarazConfigBaseTriggersLoadRulesOpGreaterThan, ZarazConfigBaseTriggersLoadRulesOpGreaterThanOrEqual, ZarazConfigBaseTriggersLoadRulesOpLessThan, ZarazConfigBaseTriggersLoadRulesOpLessThanOrEqual:
		return true
	}
	return false
}

type ZarazConfigBaseTriggersSystem string

const (
	ZarazConfigBaseTriggersSystemPageload ZarazConfigBaseTriggersSystem = "pageload"
)

func (r ZarazConfigBaseTriggersSystem) IsKnown() bool {
	switch r {
	case ZarazConfigBaseTriggersSystemPageload:
		return true
	}
	return false
}

type ZarazConfigBaseVariable struct {
	Name string                       `json:"name,required"`
	Type ZarazConfigBaseVariablesType `json:"type,required"`
	// This field can have the runtime type of [string],
	// [ZarazConfigBaseVariablesZarazWorkerVariableValue].
	Value interface{}                 `json:"value,required"`
	JSON  zarazConfigBaseVariableJSON `json:"-"`
	union ZarazConfigBaseVariablesUnion
}

// zarazConfigBaseVariableJSON contains the JSON metadata for the struct
// [ZarazConfigBaseVariable]
type zarazConfigBaseVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r zarazConfigBaseVariableJSON) RawJSON() string {
	return r.raw
}

func (r *ZarazConfigBaseVariable) UnmarshalJSON(data []byte) (err error) {
	*r = ZarazConfigBaseVariable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZarazConfigBaseVariablesUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ZarazConfigBaseVariablesZarazStringVariable],
// [ZarazConfigBaseVariablesZarazSecretVariable],
// [ZarazConfigBaseVariablesZarazWorkerVariable].
func (r ZarazConfigBaseVariable) AsUnion() ZarazConfigBaseVariablesUnion {
	return r.union
}

// Union satisfied by [ZarazConfigBaseVariablesZarazStringVariable],
// [ZarazConfigBaseVariablesZarazSecretVariable] or
// [ZarazConfigBaseVariablesZarazWorkerVariable].
type ZarazConfigBaseVariablesUnion interface {
	implementsZarazConfigBaseVariable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigBaseVariablesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZarazConfigBaseVariablesZarazStringVariable{}),
			DiscriminatorValue: "string",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZarazConfigBaseVariablesZarazSecretVariable{}),
			DiscriminatorValue: "secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZarazConfigBaseVariablesZarazWorkerVariable{}),
			DiscriminatorValue: "worker",
		},
	)
}

type ZarazConfigBaseVariablesZarazStringVariable struct {
	Name  string                                          `json:"name,required"`
	Type  ZarazConfigBaseVariablesZarazStringVariableType `json:"type,required"`
	Value string                                          `json:"value,required"`
	JSON  zarazConfigBaseVariablesZarazStringVariableJSON `json:"-"`
}

// zarazConfigBaseVariablesZarazStringVariableJSON contains the JSON metadata for
// the struct [ZarazConfigBaseVariablesZarazStringVariable]
type zarazConfigBaseVariablesZarazStringVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseVariablesZarazStringVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseVariablesZarazStringVariableJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseVariablesZarazStringVariable) implementsZarazConfigBaseVariable() {}

type ZarazConfigBaseVariablesZarazStringVariableType string

const (
	ZarazConfigBaseVariablesZarazStringVariableTypeString ZarazConfigBaseVariablesZarazStringVariableType = "string"
)

func (r ZarazConfigBaseVariablesZarazStringVariableType) IsKnown() bool {
	switch r {
	case ZarazConfigBaseVariablesZarazStringVariableTypeString:
		return true
	}
	return false
}

type ZarazConfigBaseVariablesZarazSecretVariable struct {
	Name  string                                          `json:"name,required"`
	Type  ZarazConfigBaseVariablesZarazSecretVariableType `json:"type,required"`
	Value string                                          `json:"value,required"`
	JSON  zarazConfigBaseVariablesZarazSecretVariableJSON `json:"-"`
}

// zarazConfigBaseVariablesZarazSecretVariableJSON contains the JSON metadata for
// the struct [ZarazConfigBaseVariablesZarazSecretVariable]
type zarazConfigBaseVariablesZarazSecretVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseVariablesZarazSecretVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseVariablesZarazSecretVariableJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseVariablesZarazSecretVariable) implementsZarazConfigBaseVariable() {}

type ZarazConfigBaseVariablesZarazSecretVariableType string

const (
	ZarazConfigBaseVariablesZarazSecretVariableTypeSecret ZarazConfigBaseVariablesZarazSecretVariableType = "secret"
)

func (r ZarazConfigBaseVariablesZarazSecretVariableType) IsKnown() bool {
	switch r {
	case ZarazConfigBaseVariablesZarazSecretVariableTypeSecret:
		return true
	}
	return false
}

type ZarazConfigBaseVariablesZarazWorkerVariable struct {
	Name  string                                           `json:"name,required"`
	Type  ZarazConfigBaseVariablesZarazWorkerVariableType  `json:"type,required"`
	Value ZarazConfigBaseVariablesZarazWorkerVariableValue `json:"value,required"`
	JSON  zarazConfigBaseVariablesZarazWorkerVariableJSON  `json:"-"`
}

// zarazConfigBaseVariablesZarazWorkerVariableJSON contains the JSON metadata for
// the struct [ZarazConfigBaseVariablesZarazWorkerVariable]
type zarazConfigBaseVariablesZarazWorkerVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseVariablesZarazWorkerVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseVariablesZarazWorkerVariableJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigBaseVariablesZarazWorkerVariable) implementsZarazConfigBaseVariable() {}

type ZarazConfigBaseVariablesZarazWorkerVariableType string

const (
	ZarazConfigBaseVariablesZarazWorkerVariableTypeWorker ZarazConfigBaseVariablesZarazWorkerVariableType = "worker"
)

func (r ZarazConfigBaseVariablesZarazWorkerVariableType) IsKnown() bool {
	switch r {
	case ZarazConfigBaseVariablesZarazWorkerVariableTypeWorker:
		return true
	}
	return false
}

type ZarazConfigBaseVariablesZarazWorkerVariableValue struct {
	EscapedWorkerName string                                               `json:"escapedWorkerName,required"`
	WorkerTag         string                                               `json:"workerTag,required"`
	JSON              zarazConfigBaseVariablesZarazWorkerVariableValueJSON `json:"-"`
}

// zarazConfigBaseVariablesZarazWorkerVariableValueJSON contains the JSON metadata
// for the struct [ZarazConfigBaseVariablesZarazWorkerVariableValue]
type zarazConfigBaseVariablesZarazWorkerVariableValueJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazConfigBaseVariablesZarazWorkerVariableValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseVariablesZarazWorkerVariableValueJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseVariablesType string

const (
	ZarazConfigBaseVariablesTypeString ZarazConfigBaseVariablesType = "string"
	ZarazConfigBaseVariablesTypeSecret ZarazConfigBaseVariablesType = "secret"
	ZarazConfigBaseVariablesTypeWorker ZarazConfigBaseVariablesType = "worker"
)

func (r ZarazConfigBaseVariablesType) IsKnown() bool {
	switch r {
	case ZarazConfigBaseVariablesTypeString, ZarazConfigBaseVariablesTypeSecret, ZarazConfigBaseVariablesTypeWorker:
		return true
	}
	return false
}

// Cloudflare Monitoring settings.
type ZarazConfigBaseAnalytics struct {
	// Consent purpose assigned to Monitoring.
	DefaultPurpose string `json:"defaultPurpose"`
	// Whether Advanced Monitoring reports are enabled.
	Enabled bool `json:"enabled"`
	// Session expiration time (seconds).
	SessionExpTime int64                        `json:"sessionExpTime"`
	JSON           zarazConfigBaseAnalyticsJSON `json:"-"`
}

// zarazConfigBaseAnalyticsJSON contains the JSON metadata for the struct
// [ZarazConfigBaseAnalytics]
type zarazConfigBaseAnalyticsJSON struct {
	DefaultPurpose apijson.Field
	Enabled        apijson.Field
	SessionExpTime apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZarazConfigBaseAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseAnalyticsJSON) RawJSON() string {
	return r.raw
}

// Consent management configuration.
type ZarazConfigBaseConsent struct {
	Enabled                bool                                         `json:"enabled,required"`
	ButtonTextTranslations ZarazConfigBaseConsentButtonTextTranslations `json:"buttonTextTranslations"`
	CompanyEmail           string                                       `json:"companyEmail"`
	CompanyName            string                                       `json:"companyName"`
	CompanyStreetAddress   string                                       `json:"companyStreetAddress"`
	ConsentModalIntroHTML  string                                       `json:"consentModalIntroHTML"`
	// Object where keys are language codes
	ConsentModalIntroHTMLWithTranslations map[string]string `json:"consentModalIntroHTMLWithTranslations"`
	CookieName                            string            `json:"cookieName"`
	CustomCss                             string            `json:"customCSS"`
	CustomIntroDisclaimerDismissed        bool              `json:"customIntroDisclaimerDismissed"`
	DefaultLanguage                       string            `json:"defaultLanguage"`
	HideModal                             bool              `json:"hideModal"`
	// Object where keys are purpose alpha-numeric IDs
	Purposes map[string]ZarazConfigBaseConsentPurpose `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations map[string]ZarazConfigBaseConsentPurposesWithTranslation `json:"purposesWithTranslations"`
	TcfCompliant             bool                                                     `json:"tcfCompliant"`
	JSON                     zarazConfigBaseConsentJSON                               `json:"-"`
}

// zarazConfigBaseConsentJSON contains the JSON metadata for the struct
// [ZarazConfigBaseConsent]
type zarazConfigBaseConsentJSON struct {
	Enabled                               apijson.Field
	ButtonTextTranslations                apijson.Field
	CompanyEmail                          apijson.Field
	CompanyName                           apijson.Field
	CompanyStreetAddress                  apijson.Field
	ConsentModalIntroHTML                 apijson.Field
	ConsentModalIntroHTMLWithTranslations apijson.Field
	CookieName                            apijson.Field
	CustomCss                             apijson.Field
	CustomIntroDisclaimerDismissed        apijson.Field
	DefaultLanguage                       apijson.Field
	HideModal                             apijson.Field
	Purposes                              apijson.Field
	PurposesWithTranslations              apijson.Field
	TcfCompliant                          apijson.Field
	raw                                   string
	ExtraFields                           map[string]apijson.Field
}

func (r *ZarazConfigBaseConsent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseConsentJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll map[string]string `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices map[string]string `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll map[string]string                                `json:"reject_all,required"`
	JSON      zarazConfigBaseConsentButtonTextTranslationsJSON `json:"-"`
}

// zarazConfigBaseConsentButtonTextTranslationsJSON contains the JSON metadata for
// the struct [ZarazConfigBaseConsentButtonTextTranslations]
type zarazConfigBaseConsentButtonTextTranslationsJSON struct {
	AcceptAll        apijson.Field
	ConfirmMyChoices apijson.Field
	RejectAll        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigBaseConsentButtonTextTranslations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseConsentButtonTextTranslationsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseConsentPurpose struct {
	Description string                            `json:"description,required"`
	Name        string                            `json:"name,required"`
	JSON        zarazConfigBaseConsentPurposeJSON `json:"-"`
}

// zarazConfigBaseConsentPurposeJSON contains the JSON metadata for the struct
// [ZarazConfigBaseConsentPurpose]
type zarazConfigBaseConsentPurposeJSON struct {
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseConsentPurpose) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseConsentPurposeJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigBaseConsentPurposesWithTranslation struct {
	// Object where keys are language codes
	Description map[string]string `json:"description,required"`
	// Object where keys are language codes
	Name  map[string]string                                 `json:"name,required"`
	Order int64                                             `json:"order,required"`
	JSON  zarazConfigBaseConsentPurposesWithTranslationJSON `json:"-"`
}

// zarazConfigBaseConsentPurposesWithTranslationJSON contains the JSON metadata for
// the struct [ZarazConfigBaseConsentPurposesWithTranslation]
type zarazConfigBaseConsentPurposesWithTranslationJSON struct {
	Description apijson.Field
	Name        apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigBaseConsentPurposesWithTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigBaseConsentPurposesWithTranslationJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigResponse struct {
	// Zaraz configuration
	Result ZarazConfigReturn       `json:"result"`
	JSON   zarazConfigResponseJSON `json:"-"`
	ZarazAPIResponseCommon
}

// zarazConfigResponseJSON contains the JSON metadata for the struct
// [ZarazConfigResponse]
type zarazConfigResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingZarazConfigUpdateParams struct {
	// Data layer compatibility mode enabled.
	DataLayer param.Field[bool] `json:"dataLayer,required"`
	// The key for Zaraz debug mode.
	DebugKey param.Field[string] `json:"debugKey,required"`
	// General Zaraz settings.
	Settings param.Field[ZoneSettingZarazConfigUpdateParamsSettings] `json:"settings,required"`
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools param.Field[map[string]ZoneSettingZarazConfigUpdateParamsToolsUnion] `json:"tools,required"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers param.Field[map[string]ZoneSettingZarazConfigUpdateParamsTriggers] `json:"triggers,required"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables param.Field[map[string]ZoneSettingZarazConfigUpdateParamsVariablesUnion] `json:"variables,required"`
	// Zaraz internal version of the config.
	ZarazVersion param.Field[int64] `json:"zarazVersion,required"`
	// Cloudflare Monitoring settings.
	Analytics param.Field[ZoneSettingZarazConfigUpdateParamsAnalytics] `json:"analytics"`
	// Consent management configuration.
	Consent param.Field[ZoneSettingZarazConfigUpdateParamsConsent] `json:"consent"`
	// Single Page Application support enabled.
	HistoryChange param.Field[bool] `json:"historyChange"`
}

func (r ZoneSettingZarazConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// General Zaraz settings.
type ZoneSettingZarazConfigUpdateParamsSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript param.Field[bool] `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher param.Field[ZoneSettingZarazConfigUpdateParamsSettingsContextEnricher] `json:"contextEnricher"`
	// The domain Zaraz will use for writing and reading its cookies.
	CookieDomain param.Field[string] `json:"cookieDomain"`
	// Ecommerce API enabled.
	Ecommerce param.Field[bool] `json:"ecommerce"`
	// Custom endpoint for server-side track events.
	EventsAPIPath param.Field[string] `json:"eventsApiPath"`
	// Hiding external referrer URL enabled.
	HideExternalReferer param.Field[bool] `json:"hideExternalReferer"`
	// Trimming IP address enabled.
	HideIPAddress param.Field[bool] `json:"hideIPAddress"`
	// Removing URL query params enabled.
	HideQueryParams param.Field[bool] `json:"hideQueryParams"`
	// Removing sensitive data from User Aagent string enabled.
	HideUserAgent param.Field[bool] `json:"hideUserAgent"`
	// Custom endpoint for Zaraz init script.
	InitPath param.Field[string] `json:"initPath"`
	// Injection of Zaraz scripts into iframes enabled.
	InjectIframes param.Field[bool] `json:"injectIframes"`
	// Custom path for Managed Components server functionalities.
	McRootPath param.Field[string] `json:"mcRootPath"`
	// Custom endpoint for Zaraz main script.
	ScriptPath param.Field[string] `json:"scriptPath"`
	// Custom endpoint for Zaraz tracking requests.
	TrackPath param.Field[string] `json:"trackPath"`
}

func (r ZoneSettingZarazConfigUpdateParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ZoneSettingZarazConfigUpdateParamsSettingsContextEnricher struct {
	EscapedWorkerName param.Field[string] `json:"escapedWorkerName,required"`
	WorkerTag         param.Field[string] `json:"workerTag,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsSettingsContextEnricher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsTools struct {
	BlockingTriggers param.Field[interface{}] `json:"blockingTriggers,required"`
	// Tool's internal name
	Component     param.Field[string]      `json:"component,required"`
	DefaultFields param.Field[interface{}] `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Tool's name defined by the user
	Name        param.Field[string]                                      `json:"name,required"`
	Permissions param.Field[interface{}]                                 `json:"permissions,required"`
	Settings    param.Field[interface{}]                                 `json:"settings,required"`
	Type        param.Field[ZoneSettingZarazConfigUpdateParamsToolsType] `json:"type,required"`
	Actions     param.Field[interface{}]                                 `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose param.Field[string]      `json:"defaultPurpose"`
	NeoEvents      param.Field[interface{}] `json:"neoEvents"`
	// Vendor name for TCF compliant consent modal, required for Custom Managed
	// Components and Custom HTML tool with a defaultPurpose assigned
	VendorName param.Field[string] `json:"vendorName"`
	// Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom
	// Managed Components and Custom HTML tool with a defaultPurpose assigned
	VendorPolicyURL param.Field[string]      `json:"vendorPolicyUrl"`
	Worker          param.Field[interface{}] `json:"worker"`
}

func (r ZoneSettingZarazConfigUpdateParamsTools) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTools) implementsZoneSettingZarazConfigUpdateParamsToolsUnion() {
}

// Satisfied by [ZoneSettingZarazConfigUpdateParamsToolsObject],
// [ZoneSettingZarazConfigUpdateParamsToolsObject],
// [ZoneSettingZarazConfigUpdateParamsTools].
type ZoneSettingZarazConfigUpdateParamsToolsUnion interface {
	implementsZoneSettingZarazConfigUpdateParamsToolsUnion()
}

type ZoneSettingZarazConfigUpdateParamsToolsObject struct {
	// List of blocking trigger IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Tool's internal name
	Component param.Field[string] `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields param.Field[map[string]ZoneSettingZarazConfigUpdateParamsToolsObjectDefaultFieldsUnion] `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Tool's name defined by the user
	Name param.Field[string] `json:"name,required"`
	// List of permissions granted to the component
	Permissions param.Field[[]string] `json:"permissions,required"`
	// Tool's settings
	Settings param.Field[map[string]ZoneSettingZarazConfigUpdateParamsToolsObjectSettingsUnion] `json:"settings,required"`
	Type     param.Field[ZoneSettingZarazConfigUpdateParamsToolsObjectType]                     `json:"type,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions param.Field[map[string]ZoneSettingZarazConfigUpdateParamsToolsObjectActions] `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose param.Field[string] `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents param.Field[[]ZoneSettingZarazConfigUpdateParamsToolsObjectNeoEvent] `json:"neoEvents"`
	// Vendor name for TCF compliant consent modal, required for Custom Managed
	// Components and Custom HTML tool with a defaultPurpose assigned
	VendorName param.Field[string] `json:"vendorName"`
	// Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom
	// Managed Components and Custom HTML tool with a defaultPurpose assigned
	VendorPolicyURL param.Field[string] `json:"vendorPolicyUrl"`
}

func (r ZoneSettingZarazConfigUpdateParamsToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsToolsObject) implementsZoneSettingZarazConfigUpdateParamsToolsUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZoneSettingZarazConfigUpdateParamsToolsObjectDefaultFieldsUnion interface {
	ImplementsZoneSettingZarazConfigUpdateParamsToolsObjectDefaultFieldsUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZoneSettingZarazConfigUpdateParamsToolsObjectSettingsUnion interface {
	ImplementsZoneSettingZarazConfigUpdateParamsToolsObjectSettingsUnion()
}

type ZoneSettingZarazConfigUpdateParamsToolsObjectType string

const (
	ZoneSettingZarazConfigUpdateParamsToolsObjectTypeComponent ZoneSettingZarazConfigUpdateParamsToolsObjectType = "component"
)

func (r ZoneSettingZarazConfigUpdateParamsToolsObjectType) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsToolsObjectTypeComponent:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsToolsObjectActions struct {
	// Tool event type
	ActionType param.Field[string] `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsToolsObjectActions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsToolsObjectNeoEvent struct {
	// Tool event type
	ActionType param.Field[string] `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsToolsObjectNeoEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsToolsType string

const (
	ZoneSettingZarazConfigUpdateParamsToolsTypeComponent ZoneSettingZarazConfigUpdateParamsToolsType = "component"
	ZoneSettingZarazConfigUpdateParamsToolsTypeCustomMc  ZoneSettingZarazConfigUpdateParamsToolsType = "custom-mc"
)

func (r ZoneSettingZarazConfigUpdateParamsToolsType) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsToolsTypeComponent, ZoneSettingZarazConfigUpdateParamsToolsTypeCustomMc:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggers struct {
	// Rules defining when the trigger is not fired.
	ExcludeRules param.Field[[]ZoneSettingZarazConfigUpdateParamsTriggersExcludeRuleUnion] `json:"excludeRules,required"`
	// Rules defining when the trigger is fired.
	LoadRules param.Field[[]ZoneSettingZarazConfigUpdateParamsTriggersLoadRuleUnion] `json:"loadRules,required"`
	// Trigger name.
	Name param.Field[string] `json:"name,required"`
	// Trigger description.
	Description param.Field[string]                                           `json:"description"`
	System      param.Field[ZoneSettingZarazConfigUpdateParamsTriggersSystem] `json:"system"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRule struct {
	ID       param.Field[string]                                                       `json:"id,required"`
	Action   param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesAction] `json:"action"`
	Match    param.Field[string]                                                       `json:"match"`
	Op       param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOp]     `json:"op"`
	Settings param.Field[interface{}]                                                  `json:"settings"`
	Value    param.Field[string]                                                       `json:"value"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRule) implementsZoneSettingZarazConfigUpdateParamsTriggersExcludeRuleUnion() {
}

// Satisfied by
// [ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRule],
// [ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRule],
// [ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRule],
// [ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRule],
// [ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRule],
// [ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRule],
// [ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRule],
// [ZoneSettingZarazConfigUpdateParamsTriggersExcludeRule].
type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRuleUnion interface {
	implementsZoneSettingZarazConfigUpdateParamsTriggersExcludeRuleUnion()
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRule struct {
	ID    param.Field[string]                                                                `json:"id,required"`
	Match param.Field[string]                                                                `json:"match,required"`
	Op    param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp] `json:"op,required"`
	Value param.Field[string]                                                                `json:"value,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRule) implementsZoneSettingZarazConfigUpdateParamsTriggersExcludeRuleUnion() {
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpContains           ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "CONTAINS"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpEquals             ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "EQUALS"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpStartsWith         ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "STARTS_WITH"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpEndsWith           ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "ENDS_WITH"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpMatchRegex         ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex      ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpGreaterThan        ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpLessThan           ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual    ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpContains, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpEquals, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpStartsWith, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpEndsWith, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpMatchRegex, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpGreaterThan, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpLessThan, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRule struct {
	ID       param.Field[string]                                                                               `json:"id,required"`
	Action   param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleAction]   `json:"action,required"`
	Settings param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettings] `json:"settings,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRule) implementsZoneSettingZarazConfigUpdateParamsTriggersExcludeRuleUnion() {
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleAction string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleActionClickListener ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleAction = "clickListener"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleAction) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleActionClickListener:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettings struct {
	Selector    param.Field[string]                                                                                   `json:"selector,required"`
	Type        param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType] `json:"type,required"`
	WaitForTags param.Field[int64]                                                                                    `json:"waitForTags,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType = "xpath"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCss   ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType = "css"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCss:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRule struct {
	ID       param.Field[string]                                                                       `json:"id,required"`
	Action   param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleAction]   `json:"action,required"`
	Settings param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleSettings] `json:"settings,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRule) implementsZoneSettingZarazConfigUpdateParamsTriggersExcludeRuleUnion() {
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleAction string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleActionTimer ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleAction = "timer"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleAction) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleActionTimer:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleSettings struct {
	Interval param.Field[int64] `json:"interval,required"`
	Limit    param.Field[int64] `json:"limit,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRule struct {
	ID       param.Field[string]                                                                                `json:"id,required"`
	Action   param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleAction]   `json:"action,required"`
	Settings param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleSettings] `json:"settings,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRule) implementsZoneSettingZarazConfigUpdateParamsTriggersExcludeRuleUnion() {
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleAction string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleAction = "formSubmission"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleAction) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
	Validate param.Field[bool]   `json:"validate,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRule struct {
	ID       param.Field[string]                                                                               `json:"id,required"`
	Action   param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleAction]   `json:"action,required"`
	Settings param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleSettings] `json:"settings,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRule) implementsZoneSettingZarazConfigUpdateParamsTriggersExcludeRuleUnion() {
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleAction string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleAction = "variableMatch"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleAction) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleSettings struct {
	Match    param.Field[string] `json:"match,required"`
	Variable param.Field[string] `json:"variable,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRule struct {
	ID       param.Field[string]                                                                             `json:"id,required"`
	Action   param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleAction]   `json:"action,required"`
	Settings param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleSettings] `json:"settings,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRule) implementsZoneSettingZarazConfigUpdateParamsTriggersExcludeRuleUnion() {
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleAction string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleAction = "scrollDepth"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleAction) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleSettings struct {
	Positions param.Field[string] `json:"positions,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRule struct {
	ID       param.Field[string]                                                                                   `json:"id,required"`
	Action   param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleAction]   `json:"action,required"`
	Settings param.Field[ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleSettings] `json:"settings,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRule) implementsZoneSettingZarazConfigUpdateParamsTriggersExcludeRuleUnion() {
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleAction string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleAction) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesAction string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesActionClickListener     ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesAction = "clickListener"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesActionTimer             ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesAction = "timer"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesActionFormSubmission    ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesAction = "formSubmission"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesActionVariableMatch     ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesAction = "variableMatch"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesActionScrollDepth       ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesAction = "scrollDepth"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesActionElementVisibility ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesAction = "elementVisibility"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesAction) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesActionClickListener, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesActionTimer, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesActionFormSubmission, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesActionVariableMatch, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesActionScrollDepth, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesActionElementVisibility:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOp string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpContains           ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOp = "CONTAINS"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpEquals             ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOp = "EQUALS"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpStartsWith         ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOp = "STARTS_WITH"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpEndsWith           ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOp = "ENDS_WITH"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpMatchRegex         ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOp = "MATCH_REGEX"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpNotMatchRegex      ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOp = "NOT_MATCH_REGEX"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpGreaterThan        ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOp = "GREATER_THAN"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpGreaterThanOrEqual ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOp = "GREATER_THAN_OR_EQUAL"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpLessThan           ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOp = "LESS_THAN"
	ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpLessThanOrEqual    ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOp = "LESS_THAN_OR_EQUAL"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOp) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpContains, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpEquals, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpStartsWith, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpEndsWith, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpMatchRegex, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpNotMatchRegex, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpGreaterThan, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpGreaterThanOrEqual, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpLessThan, ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesOpLessThanOrEqual:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRule struct {
	ID       param.Field[string]                                                    `json:"id,required"`
	Action   param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesAction] `json:"action"`
	Match    param.Field[string]                                                    `json:"match"`
	Op       param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOp]     `json:"op"`
	Settings param.Field[interface{}]                                               `json:"settings"`
	Value    param.Field[string]                                                    `json:"value"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRule) implementsZoneSettingZarazConfigUpdateParamsTriggersLoadRuleUnion() {
}

// Satisfied by [ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRule],
// [ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRule],
// [ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRule],
// [ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRule],
// [ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRule],
// [ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRule],
// [ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRule],
// [ZoneSettingZarazConfigUpdateParamsTriggersLoadRule].
type ZoneSettingZarazConfigUpdateParamsTriggersLoadRuleUnion interface {
	implementsZoneSettingZarazConfigUpdateParamsTriggersLoadRuleUnion()
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRule struct {
	ID    param.Field[string]                                                             `json:"id,required"`
	Match param.Field[string]                                                             `json:"match,required"`
	Op    param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp] `json:"op,required"`
	Value param.Field[string]                                                             `json:"value,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRule) implementsZoneSettingZarazConfigUpdateParamsTriggersLoadRuleUnion() {
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpContains           ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "CONTAINS"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpEquals             ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "EQUALS"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpStartsWith         ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "STARTS_WITH"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpEndsWith           ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "ENDS_WITH"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpMatchRegex         ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpNotMatchRegex      ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpGreaterThan        ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpLessThan           ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual    ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpContains, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpEquals, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpStartsWith, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpEndsWith, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpMatchRegex, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpNotMatchRegex, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpGreaterThan, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpLessThan, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRule struct {
	ID       param.Field[string]                                                                            `json:"id,required"`
	Action   param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleAction]   `json:"action,required"`
	Settings param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettings] `json:"settings,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRule) implementsZoneSettingZarazConfigUpdateParamsTriggersLoadRuleUnion() {
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleAction string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleActionClickListener ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleAction = "clickListener"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleAction) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleActionClickListener:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettings struct {
	Selector    param.Field[string]                                                                                `json:"selector,required"`
	Type        param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsType] `json:"type,required"`
	WaitForTags param.Field[int64]                                                                                 `json:"waitForTags,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsType string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsType = "xpath"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsTypeCss   ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsType = "css"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsType) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsTypeCss:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRule struct {
	ID       param.Field[string]                                                                    `json:"id,required"`
	Action   param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleAction]   `json:"action,required"`
	Settings param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleSettings] `json:"settings,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRule) implementsZoneSettingZarazConfigUpdateParamsTriggersLoadRuleUnion() {
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleAction string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleActionTimer ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleAction = "timer"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleAction) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleActionTimer:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleSettings struct {
	Interval param.Field[int64] `json:"interval,required"`
	Limit    param.Field[int64] `json:"limit,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRule struct {
	ID       param.Field[string]                                                                             `json:"id,required"`
	Action   param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleAction]   `json:"action,required"`
	Settings param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleSettings] `json:"settings,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRule) implementsZoneSettingZarazConfigUpdateParamsTriggersLoadRuleUnion() {
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleAction string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleAction = "formSubmission"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleAction) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
	Validate param.Field[bool]   `json:"validate,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRule struct {
	ID       param.Field[string]                                                                            `json:"id,required"`
	Action   param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleAction]   `json:"action,required"`
	Settings param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleSettings] `json:"settings,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRule) implementsZoneSettingZarazConfigUpdateParamsTriggersLoadRuleUnion() {
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleAction string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleAction = "variableMatch"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleAction) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleSettings struct {
	Match    param.Field[string] `json:"match,required"`
	Variable param.Field[string] `json:"variable,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRule struct {
	ID       param.Field[string]                                                                          `json:"id,required"`
	Action   param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleAction]   `json:"action,required"`
	Settings param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleSettings] `json:"settings,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRule) implementsZoneSettingZarazConfigUpdateParamsTriggersLoadRuleUnion() {
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleAction string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleAction = "scrollDepth"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleAction) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleSettings struct {
	Positions param.Field[string] `json:"positions,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRule struct {
	ID       param.Field[string]                                                                                `json:"id,required"`
	Action   param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleAction]   `json:"action,required"`
	Settings param.Field[ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleSettings] `json:"settings,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRule) implementsZoneSettingZarazConfigUpdateParamsTriggersLoadRuleUnion() {
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleAction string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleAction) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesAction string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesActionClickListener     ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesAction = "clickListener"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesActionTimer             ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesAction = "timer"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesActionFormSubmission    ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesAction = "formSubmission"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesActionVariableMatch     ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesAction = "variableMatch"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesActionScrollDepth       ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesAction = "scrollDepth"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesActionElementVisibility ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesAction = "elementVisibility"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesAction) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesActionClickListener, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesActionTimer, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesActionFormSubmission, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesActionVariableMatch, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesActionScrollDepth, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesActionElementVisibility:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOp string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpContains           ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOp = "CONTAINS"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpEquals             ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOp = "EQUALS"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpStartsWith         ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOp = "STARTS_WITH"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpEndsWith           ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOp = "ENDS_WITH"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpMatchRegex         ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOp = "MATCH_REGEX"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpNotMatchRegex      ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOp = "NOT_MATCH_REGEX"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpGreaterThan        ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOp = "GREATER_THAN"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpGreaterThanOrEqual ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOp = "GREATER_THAN_OR_EQUAL"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpLessThan           ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOp = "LESS_THAN"
	ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpLessThanOrEqual    ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOp = "LESS_THAN_OR_EQUAL"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOp) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpContains, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpEquals, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpStartsWith, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpEndsWith, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpMatchRegex, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpNotMatchRegex, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpGreaterThan, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpGreaterThanOrEqual, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpLessThan, ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesOpLessThanOrEqual:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsTriggersSystem string

const (
	ZoneSettingZarazConfigUpdateParamsTriggersSystemPageload ZoneSettingZarazConfigUpdateParamsTriggersSystem = "pageload"
)

func (r ZoneSettingZarazConfigUpdateParamsTriggersSystem) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsTriggersSystemPageload:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsVariables struct {
	Name  param.Field[string]                                          `json:"name,required"`
	Type  param.Field[ZoneSettingZarazConfigUpdateParamsVariablesType] `json:"type,required"`
	Value param.Field[interface{}]                                     `json:"value,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsVariables) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsVariables) implementsZoneSettingZarazConfigUpdateParamsVariablesUnion() {
}

// Satisfied by [ZoneSettingZarazConfigUpdateParamsVariablesZarazStringVariable],
// [ZoneSettingZarazConfigUpdateParamsVariablesZarazSecretVariable],
// [ZoneSettingZarazConfigUpdateParamsVariablesZarazWorkerVariable],
// [ZoneSettingZarazConfigUpdateParamsVariables].
type ZoneSettingZarazConfigUpdateParamsVariablesUnion interface {
	implementsZoneSettingZarazConfigUpdateParamsVariablesUnion()
}

type ZoneSettingZarazConfigUpdateParamsVariablesZarazStringVariable struct {
	Name  param.Field[string]                                                             `json:"name,required"`
	Type  param.Field[ZoneSettingZarazConfigUpdateParamsVariablesZarazStringVariableType] `json:"type,required"`
	Value param.Field[string]                                                             `json:"value,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsVariablesZarazStringVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsVariablesZarazStringVariable) implementsZoneSettingZarazConfigUpdateParamsVariablesUnion() {
}

type ZoneSettingZarazConfigUpdateParamsVariablesZarazStringVariableType string

const (
	ZoneSettingZarazConfigUpdateParamsVariablesZarazStringVariableTypeString ZoneSettingZarazConfigUpdateParamsVariablesZarazStringVariableType = "string"
)

func (r ZoneSettingZarazConfigUpdateParamsVariablesZarazStringVariableType) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsVariablesZarazStringVariableTypeString:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsVariablesZarazSecretVariable struct {
	Name  param.Field[string]                                                             `json:"name,required"`
	Type  param.Field[ZoneSettingZarazConfigUpdateParamsVariablesZarazSecretVariableType] `json:"type,required"`
	Value param.Field[string]                                                             `json:"value,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsVariablesZarazSecretVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsVariablesZarazSecretVariable) implementsZoneSettingZarazConfigUpdateParamsVariablesUnion() {
}

type ZoneSettingZarazConfigUpdateParamsVariablesZarazSecretVariableType string

const (
	ZoneSettingZarazConfigUpdateParamsVariablesZarazSecretVariableTypeSecret ZoneSettingZarazConfigUpdateParamsVariablesZarazSecretVariableType = "secret"
)

func (r ZoneSettingZarazConfigUpdateParamsVariablesZarazSecretVariableType) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsVariablesZarazSecretVariableTypeSecret:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsVariablesZarazWorkerVariable struct {
	Name  param.Field[string]                                                              `json:"name,required"`
	Type  param.Field[ZoneSettingZarazConfigUpdateParamsVariablesZarazWorkerVariableType]  `json:"type,required"`
	Value param.Field[ZoneSettingZarazConfigUpdateParamsVariablesZarazWorkerVariableValue] `json:"value,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsVariablesZarazWorkerVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsVariablesZarazWorkerVariable) implementsZoneSettingZarazConfigUpdateParamsVariablesUnion() {
}

type ZoneSettingZarazConfigUpdateParamsVariablesZarazWorkerVariableType string

const (
	ZoneSettingZarazConfigUpdateParamsVariablesZarazWorkerVariableTypeWorker ZoneSettingZarazConfigUpdateParamsVariablesZarazWorkerVariableType = "worker"
)

func (r ZoneSettingZarazConfigUpdateParamsVariablesZarazWorkerVariableType) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsVariablesZarazWorkerVariableTypeWorker:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsVariablesZarazWorkerVariableValue struct {
	EscapedWorkerName param.Field[string] `json:"escapedWorkerName,required"`
	WorkerTag         param.Field[string] `json:"workerTag,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsVariablesZarazWorkerVariableValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsVariablesType string

const (
	ZoneSettingZarazConfigUpdateParamsVariablesTypeString ZoneSettingZarazConfigUpdateParamsVariablesType = "string"
	ZoneSettingZarazConfigUpdateParamsVariablesTypeSecret ZoneSettingZarazConfigUpdateParamsVariablesType = "secret"
	ZoneSettingZarazConfigUpdateParamsVariablesTypeWorker ZoneSettingZarazConfigUpdateParamsVariablesType = "worker"
)

func (r ZoneSettingZarazConfigUpdateParamsVariablesType) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsVariablesTypeString, ZoneSettingZarazConfigUpdateParamsVariablesTypeSecret, ZoneSettingZarazConfigUpdateParamsVariablesTypeWorker:
		return true
	}
	return false
}

// Cloudflare Monitoring settings.
type ZoneSettingZarazConfigUpdateParamsAnalytics struct {
	// Consent purpose assigned to Monitoring.
	DefaultPurpose param.Field[string] `json:"defaultPurpose"`
	// Whether Advanced Monitoring reports are enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Session expiration time (seconds).
	SessionExpTime param.Field[int64] `json:"sessionExpTime"`
}

func (r ZoneSettingZarazConfigUpdateParamsAnalytics) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Consent management configuration.
type ZoneSettingZarazConfigUpdateParamsConsent struct {
	Enabled                param.Field[bool]                                                            `json:"enabled,required"`
	ButtonTextTranslations param.Field[ZoneSettingZarazConfigUpdateParamsConsentButtonTextTranslations] `json:"buttonTextTranslations"`
	CompanyEmail           param.Field[string]                                                          `json:"companyEmail"`
	CompanyName            param.Field[string]                                                          `json:"companyName"`
	CompanyStreetAddress   param.Field[string]                                                          `json:"companyStreetAddress"`
	ConsentModalIntroHTML  param.Field[string]                                                          `json:"consentModalIntroHTML"`
	// Object where keys are language codes
	ConsentModalIntroHTMLWithTranslations param.Field[map[string]string] `json:"consentModalIntroHTMLWithTranslations"`
	CookieName                            param.Field[string]            `json:"cookieName"`
	CustomCss                             param.Field[string]            `json:"customCSS"`
	CustomIntroDisclaimerDismissed        param.Field[bool]              `json:"customIntroDisclaimerDismissed"`
	DefaultLanguage                       param.Field[string]            `json:"defaultLanguage"`
	HideModal                             param.Field[bool]              `json:"hideModal"`
	// Object where keys are purpose alpha-numeric IDs
	Purposes param.Field[map[string]ZoneSettingZarazConfigUpdateParamsConsentPurposes] `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations param.Field[map[string]ZoneSettingZarazConfigUpdateParamsConsentPurposesWithTranslations] `json:"purposesWithTranslations"`
	TcfCompliant             param.Field[bool]                                                                         `json:"tcfCompliant"`
}

func (r ZoneSettingZarazConfigUpdateParamsConsent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll param.Field[map[string]string] `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices param.Field[map[string]string] `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll param.Field[map[string]string] `json:"reject_all,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsConsentButtonTextTranslations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsConsentPurposes struct {
	Description param.Field[string] `json:"description,required"`
	Name        param.Field[string] `json:"name,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsConsentPurposes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsConsentPurposesWithTranslations struct {
	// Object where keys are language codes
	Description param.Field[map[string]string] `json:"description,required"`
	// Object where keys are language codes
	Name  param.Field[map[string]string] `json:"name,required"`
	Order param.Field[int64]             `json:"order,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsConsentPurposesWithTranslations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
