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

type ZarazConfigResponse struct {
	Errors   []ZarazMessagesItems `json:"errors,required"`
	Messages []ZarazMessagesItems `json:"messages,required"`
	// Zaraz configuration
	Result ZarazConfigReturn `json:"result,required"`
	// Whether the API call was successful
	Success bool                    `json:"success,required"`
	JSON    zarazConfigResponseJSON `json:"-"`
}

// zarazConfigResponseJSON contains the JSON metadata for the struct
// [ZarazConfigResponse]
type zarazConfigResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Satisfied by [ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponent],
// [ZoneSettingZarazConfigUpdateParamsToolsWorker],
// [ZoneSettingZarazConfigUpdateParamsTools].
type ZoneSettingZarazConfigUpdateParamsToolsUnion interface {
	implementsZoneSettingZarazConfigUpdateParamsToolsUnion()
}

type ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Tool's internal name
	Component param.Field[string] `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields param.Field[map[string]ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentDefaultFieldsUnion] `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Tool's name defined by the user
	Name param.Field[string] `json:"name,required"`
	// List of permissions granted to the component
	Permissions param.Field[[]string] `json:"permissions,required"`
	// Tool's settings
	Settings param.Field[map[string]ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentSettingsUnion] `json:"settings,required"`
	Type     param.Field[ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentType]                     `json:"type,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions param.Field[map[string]ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentActions] `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose param.Field[string] `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents param.Field[[]ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentNeoEvent] `json:"neoEvents"`
	// Vendor name for TCF compliant consent modal, required for Custom Managed
	// Components and Custom HTML tool with a defaultPurpose assigned
	VendorName param.Field[string] `json:"vendorName"`
	// Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom
	// Managed Components and Custom HTML tool with a defaultPurpose assigned
	VendorPolicyURL param.Field[string] `json:"vendorPolicyUrl"`
}

func (r ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponent) implementsZoneSettingZarazConfigUpdateParamsToolsUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentDefaultFieldsUnion interface {
	ImplementsZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentDefaultFieldsUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentSettingsUnion interface {
	ImplementsZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentSettingsUnion()
}

type ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentType string

const (
	ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentTypeComponent ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentType = "component"
)

func (r ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentType) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentTypeComponent:
		return true
	}
	return false
}

type ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentActions struct {
	// Tool event type
	ActionType param.Field[string] `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentActions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentNeoEvent struct {
	// Tool event type
	ActionType param.Field[string] `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentNeoEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsToolsWorker struct {
	// List of blocking trigger IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Tool's internal name
	Component param.Field[string] `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields param.Field[map[string]ZoneSettingZarazConfigUpdateParamsToolsWorkerDefaultFieldsUnion] `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Tool's name defined by the user
	Name param.Field[string] `json:"name,required"`
	// List of permissions granted to the component
	Permissions param.Field[[]string] `json:"permissions,required"`
	// Tool's settings
	Settings param.Field[map[string]ZoneSettingZarazConfigUpdateParamsToolsWorkerSettingsUnion] `json:"settings,required"`
	Type     param.Field[ZoneSettingZarazConfigUpdateParamsToolsWorkerType]                     `json:"type,required"`
	// Cloudflare worker that acts as a managed component
	Worker param.Field[ZoneSettingZarazConfigUpdateParamsToolsWorkerWorker] `json:"worker,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions param.Field[map[string]ZoneSettingZarazConfigUpdateParamsToolsWorkerActions] `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose param.Field[string] `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents param.Field[[]ZoneSettingZarazConfigUpdateParamsToolsWorkerNeoEvent] `json:"neoEvents"`
	// Vendor name for TCF compliant consent modal, required for Custom Managed
	// Components and Custom HTML tool with a defaultPurpose assigned
	VendorName param.Field[string] `json:"vendorName"`
	// Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom
	// Managed Components and Custom HTML tool with a defaultPurpose assigned
	VendorPolicyURL param.Field[string] `json:"vendorPolicyUrl"`
}

func (r ZoneSettingZarazConfigUpdateParamsToolsWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingZarazConfigUpdateParamsToolsWorker) implementsZoneSettingZarazConfigUpdateParamsToolsUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZoneSettingZarazConfigUpdateParamsToolsWorkerDefaultFieldsUnion interface {
	ImplementsZoneSettingZarazConfigUpdateParamsToolsWorkerDefaultFieldsUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZoneSettingZarazConfigUpdateParamsToolsWorkerSettingsUnion interface {
	ImplementsZoneSettingZarazConfigUpdateParamsToolsWorkerSettingsUnion()
}

type ZoneSettingZarazConfigUpdateParamsToolsWorkerType string

const (
	ZoneSettingZarazConfigUpdateParamsToolsWorkerTypeCustomMc ZoneSettingZarazConfigUpdateParamsToolsWorkerType = "custom-mc"
)

func (r ZoneSettingZarazConfigUpdateParamsToolsWorkerType) IsKnown() bool {
	switch r {
	case ZoneSettingZarazConfigUpdateParamsToolsWorkerTypeCustomMc:
		return true
	}
	return false
}

// Cloudflare worker that acts as a managed component
type ZoneSettingZarazConfigUpdateParamsToolsWorkerWorker struct {
	EscapedWorkerName param.Field[string] `json:"escapedWorkerName,required"`
	WorkerTag         param.Field[string] `json:"workerTag,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsToolsWorkerWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsToolsWorkerActions struct {
	// Tool event type
	ActionType param.Field[string] `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsToolsWorkerActions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsToolsWorkerNeoEvent struct {
	// Tool event type
	ActionType param.Field[string] `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsToolsWorkerNeoEvent) MarshalJSON() (data []byte, err error) {
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
