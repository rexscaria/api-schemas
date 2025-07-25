// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// ZoneSettingZarazService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSettingZarazService] method instead.
type ZoneSettingZarazService struct {
	Options  []option.RequestOption
	Config   *ZoneSettingZarazConfigService
	History  *ZoneSettingZarazHistoryService
	Workflow *ZoneSettingZarazWorkflowService
}

// NewZoneSettingZarazService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingZarazService(opts ...option.RequestOption) (r *ZoneSettingZarazService) {
	r = &ZoneSettingZarazService{}
	r.Options = opts
	r.Config = NewZoneSettingZarazConfigService(opts...)
	r.History = NewZoneSettingZarazHistoryService(opts...)
	r.Workflow = NewZoneSettingZarazWorkflowService(opts...)
	return
}

// Exports full current published Zaraz configuration for a zone, secret variables
// included.
func (r *ZoneSettingZarazService) Export(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZarazConfigReturn, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/export", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Publish current Zaraz preview configuration for a zone.
func (r *ZoneSettingZarazService) Publish(ctx context.Context, zoneID string, body ZoneSettingZarazPublishParams, opts ...option.RequestOption) (res *ZoneSettingZarazPublishResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/publish", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets default Zaraz configuration for a zone.
func (r *ZoneSettingZarazService) GetDefault(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZarazConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/default", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Zaraz configuration
type ZarazConfigReturn struct {
	// Data layer compatibility mode enabled.
	DataLayer bool `json:"dataLayer,required"`
	// The key for Zaraz debug mode.
	DebugKey string `json:"debugKey,required"`
	// General Zaraz settings.
	Settings ZarazConfigReturnSettings `json:"settings,required"`
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools map[string]ZarazConfigReturnTool `json:"tools,required"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers map[string]ZarazConfigReturnTrigger `json:"triggers,required"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables map[string]ZarazConfigReturnVariable `json:"variables,required"`
	// Zaraz internal version of the config.
	ZarazVersion int64 `json:"zarazVersion,required"`
	// Cloudflare Monitoring settings.
	Analytics ZarazConfigReturnAnalytics `json:"analytics"`
	// Consent management configuration.
	Consent ZarazConfigReturnConsent `json:"consent"`
	// Single Page Application support enabled.
	HistoryChange bool                  `json:"historyChange"`
	JSON          zarazConfigReturnJSON `json:"-"`
}

// zarazConfigReturnJSON contains the JSON metadata for the struct
// [ZarazConfigReturn]
type zarazConfigReturnJSON struct {
	DataLayer     apijson.Field
	DebugKey      apijson.Field
	Settings      apijson.Field
	Tools         apijson.Field
	Triggers      apijson.Field
	Variables     apijson.Field
	ZarazVersion  apijson.Field
	Analytics     apijson.Field
	Consent       apijson.Field
	HistoryChange apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZarazConfigReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnJSON) RawJSON() string {
	return r.raw
}

// General Zaraz settings.
type ZarazConfigReturnSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript bool `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher ZarazConfigReturnSettingsContextEnricher `json:"contextEnricher"`
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
	TrackPath string                        `json:"trackPath"`
	JSON      zarazConfigReturnSettingsJSON `json:"-"`
}

// zarazConfigReturnSettingsJSON contains the JSON metadata for the struct
// [ZarazConfigReturnSettings]
type zarazConfigReturnSettingsJSON struct {
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

func (r *ZarazConfigReturnSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnSettingsJSON) RawJSON() string {
	return r.raw
}

// Details of the worker that receives and edits Zaraz Context object.
type ZarazConfigReturnSettingsContextEnricher struct {
	EscapedWorkerName string                                       `json:"escapedWorkerName,required"`
	WorkerTag         string                                       `json:"workerTag,required"`
	JSON              zarazConfigReturnSettingsContextEnricherJSON `json:"-"`
}

// zarazConfigReturnSettingsContextEnricherJSON contains the JSON metadata for the
// struct [ZarazConfigReturnSettingsContextEnricher]
type zarazConfigReturnSettingsContextEnricherJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazConfigReturnSettingsContextEnricher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnSettingsContextEnricherJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnTool struct {
	// This field can have the runtime type of [[]string].
	BlockingTriggers interface{} `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// This field can have the runtime type of
	// [map[string]ZarazConfigReturnToolsZarazManagedComponentDefaultFieldsUnion],
	// [map[string]ZarazConfigReturnToolsWorkerDefaultFieldsUnion].
	DefaultFields interface{} `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// This field can have the runtime type of [[]string].
	Permissions interface{} `json:"permissions,required"`
	// This field can have the runtime type of
	// [map[string]ZarazConfigReturnToolsZarazManagedComponentSettingsUnion],
	// [map[string]ZarazConfigReturnToolsWorkerSettingsUnion].
	Settings interface{}                `json:"settings,required"`
	Type     ZarazConfigReturnToolsType `json:"type,required"`
	// This field can have the runtime type of
	// [map[string]ZarazConfigReturnToolsZarazManagedComponentAction],
	// [map[string]ZarazConfigReturnToolsWorkerAction].
	Actions interface{} `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// This field can have the runtime type of
	// [[]ZarazConfigReturnToolsZarazManagedComponentNeoEvent],
	// [[]ZarazConfigReturnToolsWorkerNeoEvent].
	NeoEvents interface{} `json:"neoEvents"`
	// Vendor name for TCF compliant consent modal, required for Custom Managed
	// Components and Custom HTML tool with a defaultPurpose assigned
	VendorName string `json:"vendorName"`
	// Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom
	// Managed Components and Custom HTML tool with a defaultPurpose assigned
	VendorPolicyURL string `json:"vendorPolicyUrl"`
	// This field can have the runtime type of [ZarazConfigReturnToolsWorkerWorker].
	Worker interface{}               `json:"worker"`
	JSON   zarazConfigReturnToolJSON `json:"-"`
	union  ZarazConfigReturnToolsUnion
}

// zarazConfigReturnToolJSON contains the JSON metadata for the struct
// [ZarazConfigReturnTool]
type zarazConfigReturnToolJSON struct {
	BlockingTriggers apijson.Field
	Component        apijson.Field
	DefaultFields    apijson.Field
	Enabled          apijson.Field
	Name             apijson.Field
	Permissions      apijson.Field
	Settings         apijson.Field
	Type             apijson.Field
	Actions          apijson.Field
	DefaultPurpose   apijson.Field
	NeoEvents        apijson.Field
	VendorName       apijson.Field
	VendorPolicyURL  apijson.Field
	Worker           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r zarazConfigReturnToolJSON) RawJSON() string {
	return r.raw
}

func (r *ZarazConfigReturnTool) UnmarshalJSON(data []byte) (err error) {
	*r = ZarazConfigReturnTool{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZarazConfigReturnToolsUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [ZarazConfigReturnToolsZarazManagedComponent], [ZarazConfigReturnToolsWorker].
func (r ZarazConfigReturnTool) AsUnion() ZarazConfigReturnToolsUnion {
	return r.union
}

// Union satisfied by [ZarazConfigReturnToolsZarazManagedComponent] or
// [ZarazConfigReturnToolsWorker].
type ZarazConfigReturnToolsUnion interface {
	implementsZarazConfigReturnTool()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigReturnToolsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnToolsZarazManagedComponent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnToolsWorker{}),
		},
	)
}

type ZarazConfigReturnToolsZarazManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazConfigReturnToolsZarazManagedComponentDefaultFieldsUnion `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazConfigReturnToolsZarazManagedComponentSettingsUnion `json:"settings,required"`
	Type     ZarazConfigReturnToolsZarazManagedComponentType                     `json:"type,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazConfigReturnToolsZarazManagedComponentAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazConfigReturnToolsZarazManagedComponentNeoEvent `json:"neoEvents"`
	// Vendor name for TCF compliant consent modal, required for Custom Managed
	// Components and Custom HTML tool with a defaultPurpose assigned
	VendorName string `json:"vendorName"`
	// Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom
	// Managed Components and Custom HTML tool with a defaultPurpose assigned
	VendorPolicyURL string                                          `json:"vendorPolicyUrl"`
	JSON            zarazConfigReturnToolsZarazManagedComponentJSON `json:"-"`
}

// zarazConfigReturnToolsZarazManagedComponentJSON contains the JSON metadata for
// the struct [ZarazConfigReturnToolsZarazManagedComponent]
type zarazConfigReturnToolsZarazManagedComponentJSON struct {
	BlockingTriggers apijson.Field
	Component        apijson.Field
	DefaultFields    apijson.Field
	Enabled          apijson.Field
	Name             apijson.Field
	Permissions      apijson.Field
	Settings         apijson.Field
	Type             apijson.Field
	Actions          apijson.Field
	DefaultPurpose   apijson.Field
	NeoEvents        apijson.Field
	VendorName       apijson.Field
	VendorPolicyURL  apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReturnToolsZarazManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnToolsZarazManagedComponentJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnToolsZarazManagedComponent) implementsZarazConfigReturnTool() {}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazConfigReturnToolsZarazManagedComponentDefaultFieldsUnion interface {
	ImplementsZarazConfigReturnToolsZarazManagedComponentDefaultFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigReturnToolsZarazManagedComponentDefaultFieldsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazConfigReturnToolsZarazManagedComponentSettingsUnion interface {
	ImplementsZarazConfigReturnToolsZarazManagedComponentSettingsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigReturnToolsZarazManagedComponentSettingsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ZarazConfigReturnToolsZarazManagedComponentType string

const (
	ZarazConfigReturnToolsZarazManagedComponentTypeComponent ZarazConfigReturnToolsZarazManagedComponentType = "component"
)

func (r ZarazConfigReturnToolsZarazManagedComponentType) IsKnown() bool {
	switch r {
	case ZarazConfigReturnToolsZarazManagedComponentTypeComponent:
		return true
	}
	return false
}

type ZarazConfigReturnToolsZarazManagedComponentAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                              `json:"firingTriggers,required"`
	JSON           zarazConfigReturnToolsZarazManagedComponentActionJSON `json:"-"`
}

// zarazConfigReturnToolsZarazManagedComponentActionJSON contains the JSON metadata
// for the struct [ZarazConfigReturnToolsZarazManagedComponentAction]
type zarazConfigReturnToolsZarazManagedComponentActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReturnToolsZarazManagedComponentAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnToolsZarazManagedComponentActionJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnToolsZarazManagedComponentNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                `json:"firingTriggers,required"`
	JSON           zarazConfigReturnToolsZarazManagedComponentNeoEventJSON `json:"-"`
}

// zarazConfigReturnToolsZarazManagedComponentNeoEventJSON contains the JSON
// metadata for the struct [ZarazConfigReturnToolsZarazManagedComponentNeoEvent]
type zarazConfigReturnToolsZarazManagedComponentNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReturnToolsZarazManagedComponentNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnToolsZarazManagedComponentNeoEventJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnToolsWorker struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazConfigReturnToolsWorkerDefaultFieldsUnion `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazConfigReturnToolsWorkerSettingsUnion `json:"settings,required"`
	Type     ZarazConfigReturnToolsWorkerType                     `json:"type,required"`
	// Cloudflare worker that acts as a managed component
	Worker ZarazConfigReturnToolsWorkerWorker `json:"worker,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazConfigReturnToolsWorkerAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazConfigReturnToolsWorkerNeoEvent `json:"neoEvents"`
	// Vendor name for TCF compliant consent modal, required for Custom Managed
	// Components and Custom HTML tool with a defaultPurpose assigned
	VendorName string `json:"vendorName"`
	// Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom
	// Managed Components and Custom HTML tool with a defaultPurpose assigned
	VendorPolicyURL string                           `json:"vendorPolicyUrl"`
	JSON            zarazConfigReturnToolsWorkerJSON `json:"-"`
}

// zarazConfigReturnToolsWorkerJSON contains the JSON metadata for the struct
// [ZarazConfigReturnToolsWorker]
type zarazConfigReturnToolsWorkerJSON struct {
	BlockingTriggers apijson.Field
	Component        apijson.Field
	DefaultFields    apijson.Field
	Enabled          apijson.Field
	Name             apijson.Field
	Permissions      apijson.Field
	Settings         apijson.Field
	Type             apijson.Field
	Worker           apijson.Field
	Actions          apijson.Field
	DefaultPurpose   apijson.Field
	NeoEvents        apijson.Field
	VendorName       apijson.Field
	VendorPolicyURL  apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReturnToolsWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnToolsWorkerJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnToolsWorker) implementsZarazConfigReturnTool() {}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazConfigReturnToolsWorkerDefaultFieldsUnion interface {
	ImplementsZarazConfigReturnToolsWorkerDefaultFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigReturnToolsWorkerDefaultFieldsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazConfigReturnToolsWorkerSettingsUnion interface {
	ImplementsZarazConfigReturnToolsWorkerSettingsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigReturnToolsWorkerSettingsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ZarazConfigReturnToolsWorkerType string

const (
	ZarazConfigReturnToolsWorkerTypeCustomMc ZarazConfigReturnToolsWorkerType = "custom-mc"
)

func (r ZarazConfigReturnToolsWorkerType) IsKnown() bool {
	switch r {
	case ZarazConfigReturnToolsWorkerTypeCustomMc:
		return true
	}
	return false
}

// Cloudflare worker that acts as a managed component
type ZarazConfigReturnToolsWorkerWorker struct {
	EscapedWorkerName string                                 `json:"escapedWorkerName,required"`
	WorkerTag         string                                 `json:"workerTag,required"`
	JSON              zarazConfigReturnToolsWorkerWorkerJSON `json:"-"`
}

// zarazConfigReturnToolsWorkerWorkerJSON contains the JSON metadata for the struct
// [ZarazConfigReturnToolsWorkerWorker]
type zarazConfigReturnToolsWorkerWorkerJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazConfigReturnToolsWorkerWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnToolsWorkerWorkerJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnToolsWorkerAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                               `json:"firingTriggers,required"`
	JSON           zarazConfigReturnToolsWorkerActionJSON `json:"-"`
}

// zarazConfigReturnToolsWorkerActionJSON contains the JSON metadata for the struct
// [ZarazConfigReturnToolsWorkerAction]
type zarazConfigReturnToolsWorkerActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReturnToolsWorkerAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnToolsWorkerActionJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnToolsWorkerNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                 `json:"firingTriggers,required"`
	JSON           zarazConfigReturnToolsWorkerNeoEventJSON `json:"-"`
}

// zarazConfigReturnToolsWorkerNeoEventJSON contains the JSON metadata for the
// struct [ZarazConfigReturnToolsWorkerNeoEvent]
type zarazConfigReturnToolsWorkerNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReturnToolsWorkerNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnToolsWorkerNeoEventJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnToolsType string

const (
	ZarazConfigReturnToolsTypeComponent ZarazConfigReturnToolsType = "component"
	ZarazConfigReturnToolsTypeCustomMc  ZarazConfigReturnToolsType = "custom-mc"
)

func (r ZarazConfigReturnToolsType) IsKnown() bool {
	switch r {
	case ZarazConfigReturnToolsTypeComponent, ZarazConfigReturnToolsTypeCustomMc:
		return true
	}
	return false
}

type ZarazConfigReturnTrigger struct {
	// Rules defining when the trigger is not fired.
	ExcludeRules []ZarazConfigReturnTriggersExcludeRule `json:"excludeRules,required"`
	// Rules defining when the trigger is fired.
	LoadRules []ZarazConfigReturnTriggersLoadRule `json:"loadRules,required"`
	// Trigger name.
	Name string `json:"name,required"`
	// Trigger description.
	Description string                          `json:"description"`
	System      ZarazConfigReturnTriggersSystem `json:"system"`
	JSON        zarazConfigReturnTriggerJSON    `json:"-"`
}

// zarazConfigReturnTriggerJSON contains the JSON metadata for the struct
// [ZarazConfigReturnTrigger]
type zarazConfigReturnTriggerJSON struct {
	ExcludeRules apijson.Field
	LoadRules    apijson.Field
	Name         apijson.Field
	Description  apijson.Field
	System       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZarazConfigReturnTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggerJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnTriggersExcludeRule struct {
	ID     string                                      `json:"id,required"`
	Action ZarazConfigReturnTriggersExcludeRulesAction `json:"action"`
	Match  string                                      `json:"match"`
	Op     ZarazConfigReturnTriggersExcludeRulesOp     `json:"op"`
	// This field can have the runtime type of
	// [ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettings],
	// [ZarazConfigReturnTriggersExcludeRulesZarazTimerRuleSettings],
	// [ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleSettings],
	// [ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleSettings],
	// [ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleSettings],
	// [ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleSettings].
	Settings interface{}                              `json:"settings"`
	Value    string                                   `json:"value"`
	JSON     zarazConfigReturnTriggersExcludeRuleJSON `json:"-"`
	union    ZarazConfigReturnTriggersExcludeRulesUnion
}

// zarazConfigReturnTriggersExcludeRuleJSON contains the JSON metadata for the
// struct [ZarazConfigReturnTriggersExcludeRule]
type zarazConfigReturnTriggersExcludeRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Settings    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r zarazConfigReturnTriggersExcludeRuleJSON) RawJSON() string {
	return r.raw
}

func (r *ZarazConfigReturnTriggersExcludeRule) UnmarshalJSON(data []byte) (err error) {
	*r = ZarazConfigReturnTriggersExcludeRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZarazConfigReturnTriggersExcludeRulesUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ZarazConfigReturnTriggersExcludeRulesZarazLoadRule],
// [ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRule],
// [ZarazConfigReturnTriggersExcludeRulesZarazTimerRule],
// [ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRule],
// [ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRule],
// [ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRule],
// [ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRule].
func (r ZarazConfigReturnTriggersExcludeRule) AsUnion() ZarazConfigReturnTriggersExcludeRulesUnion {
	return r.union
}

// Union satisfied by [ZarazConfigReturnTriggersExcludeRulesZarazLoadRule],
// [ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRule],
// [ZarazConfigReturnTriggersExcludeRulesZarazTimerRule],
// [ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRule],
// [ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRule],
// [ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRule] or
// [ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRule].
type ZarazConfigReturnTriggersExcludeRulesUnion interface {
	implementsZarazConfigReturnTriggersExcludeRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigReturnTriggersExcludeRulesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnTriggersExcludeRulesZarazLoadRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnTriggersExcludeRulesZarazTimerRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRule{}),
		},
	)
}

type ZarazConfigReturnTriggersExcludeRulesZarazLoadRule struct {
	ID    string                                                 `json:"id,required"`
	Match string                                                 `json:"match,required"`
	Op    ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                                 `json:"value,required"`
	JSON  zarazConfigReturnTriggersExcludeRulesZarazLoadRuleJSON `json:"-"`
}

// zarazConfigReturnTriggersExcludeRulesZarazLoadRuleJSON contains the JSON
// metadata for the struct [ZarazConfigReturnTriggersExcludeRulesZarazLoadRule]
type zarazConfigReturnTriggersExcludeRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersExcludeRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersExcludeRulesZarazLoadRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnTriggersExcludeRulesZarazLoadRule) implementsZarazConfigReturnTriggersExcludeRule() {
}

type ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOp string

const (
	ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpContains           ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOp = "CONTAINS"
	ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpEquals             ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOp = "EQUALS"
	ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpStartsWith         ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpEndsWith           ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpMatchRegex         ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex      ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpGreaterThan        ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpLessThan           ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual    ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

func (r ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOp) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpContains, ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpEquals, ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpStartsWith, ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpEndsWith, ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpMatchRegex, ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex, ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpGreaterThan, ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual, ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpLessThan, ZarazConfigReturnTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRule struct {
	ID       string                                                              `json:"id,required"`
	Action   ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRule]
type zarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRule) implementsZarazConfigReturnTriggersExcludeRule() {
}

type ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleAction string

const (
	ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleActionClickListener ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleAction = "clickListener"
)

func (r ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleActionClickListener:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                  `json:"selector,required"`
	Type        ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                   `json:"waitForTags,required"`
	JSON        zarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettings]
type zarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettingsType string

const (
	ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCss   ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettingsType = "css"
)

func (r ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettingsType) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath, ZarazConfigReturnTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCss:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersExcludeRulesZarazTimerRule struct {
	ID       string                                                      `json:"id,required"`
	Action   ZarazConfigReturnTriggersExcludeRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazConfigReturnTriggersExcludeRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazConfigReturnTriggersExcludeRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazConfigReturnTriggersExcludeRulesZarazTimerRuleJSON contains the JSON
// metadata for the struct [ZarazConfigReturnTriggersExcludeRulesZarazTimerRule]
type zarazConfigReturnTriggersExcludeRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersExcludeRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersExcludeRulesZarazTimerRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnTriggersExcludeRulesZarazTimerRule) implementsZarazConfigReturnTriggersExcludeRule() {
}

type ZarazConfigReturnTriggersExcludeRulesZarazTimerRuleAction string

const (
	ZarazConfigReturnTriggersExcludeRulesZarazTimerRuleActionTimer ZarazConfigReturnTriggersExcludeRulesZarazTimerRuleAction = "timer"
)

func (r ZarazConfigReturnTriggersExcludeRulesZarazTimerRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersExcludeRulesZarazTimerRuleActionTimer:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersExcludeRulesZarazTimerRuleSettings struct {
	Interval int64                                                           `json:"interval,required"`
	Limit    int64                                                           `json:"limit,required"`
	JSON     zarazConfigReturnTriggersExcludeRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazConfigReturnTriggersExcludeRulesZarazTimerRuleSettingsJSON contains the
// JSON metadata for the struct
// [ZarazConfigReturnTriggersExcludeRulesZarazTimerRuleSettings]
type zarazConfigReturnTriggersExcludeRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersExcludeRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersExcludeRulesZarazTimerRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRule struct {
	ID       string                                                               `json:"id,required"`
	Action   ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRule]
type zarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRule) implementsZarazConfigReturnTriggersExcludeRule() {
}

type ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleAction string

const (
	ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleAction = "formSubmission"
)

func (r ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                   `json:"selector,required"`
	Validate bool                                                                     `json:"validate,required"`
	JSON     zarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleSettings]
type zarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRule struct {
	ID       string                                                              `json:"id,required"`
	Action   ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRule]
type zarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRule) implementsZarazConfigReturnTriggersExcludeRule() {
}

type ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleAction string

const (
	ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleAction = "variableMatch"
)

func (r ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                  `json:"match,required"`
	Variable string                                                                  `json:"variable,required"`
	JSON     zarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleSettings]
type zarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRule struct {
	ID       string                                                            `json:"id,required"`
	Action   ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRule]
type zarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRule) implementsZarazConfigReturnTriggersExcludeRule() {
}

type ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleAction string

const (
	ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleAction = "scrollDepth"
)

func (r ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                                `json:"positions,required"`
	JSON      zarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleSettings]
type zarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRule struct {
	ID       string                                                                  `json:"id,required"`
	Action   ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRule]
type zarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRule) implementsZarazConfigReturnTriggersExcludeRule() {
}

type ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleAction string

const (
	ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

func (r ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                      `json:"selector,required"`
	JSON     zarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleSettings]
type zarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnTriggersExcludeRulesAction string

const (
	ZarazConfigReturnTriggersExcludeRulesActionClickListener     ZarazConfigReturnTriggersExcludeRulesAction = "clickListener"
	ZarazConfigReturnTriggersExcludeRulesActionTimer             ZarazConfigReturnTriggersExcludeRulesAction = "timer"
	ZarazConfigReturnTriggersExcludeRulesActionFormSubmission    ZarazConfigReturnTriggersExcludeRulesAction = "formSubmission"
	ZarazConfigReturnTriggersExcludeRulesActionVariableMatch     ZarazConfigReturnTriggersExcludeRulesAction = "variableMatch"
	ZarazConfigReturnTriggersExcludeRulesActionScrollDepth       ZarazConfigReturnTriggersExcludeRulesAction = "scrollDepth"
	ZarazConfigReturnTriggersExcludeRulesActionElementVisibility ZarazConfigReturnTriggersExcludeRulesAction = "elementVisibility"
)

func (r ZarazConfigReturnTriggersExcludeRulesAction) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersExcludeRulesActionClickListener, ZarazConfigReturnTriggersExcludeRulesActionTimer, ZarazConfigReturnTriggersExcludeRulesActionFormSubmission, ZarazConfigReturnTriggersExcludeRulesActionVariableMatch, ZarazConfigReturnTriggersExcludeRulesActionScrollDepth, ZarazConfigReturnTriggersExcludeRulesActionElementVisibility:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersExcludeRulesOp string

const (
	ZarazConfigReturnTriggersExcludeRulesOpContains           ZarazConfigReturnTriggersExcludeRulesOp = "CONTAINS"
	ZarazConfigReturnTriggersExcludeRulesOpEquals             ZarazConfigReturnTriggersExcludeRulesOp = "EQUALS"
	ZarazConfigReturnTriggersExcludeRulesOpStartsWith         ZarazConfigReturnTriggersExcludeRulesOp = "STARTS_WITH"
	ZarazConfigReturnTriggersExcludeRulesOpEndsWith           ZarazConfigReturnTriggersExcludeRulesOp = "ENDS_WITH"
	ZarazConfigReturnTriggersExcludeRulesOpMatchRegex         ZarazConfigReturnTriggersExcludeRulesOp = "MATCH_REGEX"
	ZarazConfigReturnTriggersExcludeRulesOpNotMatchRegex      ZarazConfigReturnTriggersExcludeRulesOp = "NOT_MATCH_REGEX"
	ZarazConfigReturnTriggersExcludeRulesOpGreaterThan        ZarazConfigReturnTriggersExcludeRulesOp = "GREATER_THAN"
	ZarazConfigReturnTriggersExcludeRulesOpGreaterThanOrEqual ZarazConfigReturnTriggersExcludeRulesOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigReturnTriggersExcludeRulesOpLessThan           ZarazConfigReturnTriggersExcludeRulesOp = "LESS_THAN"
	ZarazConfigReturnTriggersExcludeRulesOpLessThanOrEqual    ZarazConfigReturnTriggersExcludeRulesOp = "LESS_THAN_OR_EQUAL"
)

func (r ZarazConfigReturnTriggersExcludeRulesOp) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersExcludeRulesOpContains, ZarazConfigReturnTriggersExcludeRulesOpEquals, ZarazConfigReturnTriggersExcludeRulesOpStartsWith, ZarazConfigReturnTriggersExcludeRulesOpEndsWith, ZarazConfigReturnTriggersExcludeRulesOpMatchRegex, ZarazConfigReturnTriggersExcludeRulesOpNotMatchRegex, ZarazConfigReturnTriggersExcludeRulesOpGreaterThan, ZarazConfigReturnTriggersExcludeRulesOpGreaterThanOrEqual, ZarazConfigReturnTriggersExcludeRulesOpLessThan, ZarazConfigReturnTriggersExcludeRulesOpLessThanOrEqual:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersLoadRule struct {
	ID     string                                   `json:"id,required"`
	Action ZarazConfigReturnTriggersLoadRulesAction `json:"action"`
	Match  string                                   `json:"match"`
	Op     ZarazConfigReturnTriggersLoadRulesOp     `json:"op"`
	// This field can have the runtime type of
	// [ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettings],
	// [ZarazConfigReturnTriggersLoadRulesZarazTimerRuleSettings],
	// [ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleSettings],
	// [ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleSettings],
	// [ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleSettings],
	// [ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleSettings].
	Settings interface{}                           `json:"settings"`
	Value    string                                `json:"value"`
	JSON     zarazConfigReturnTriggersLoadRuleJSON `json:"-"`
	union    ZarazConfigReturnTriggersLoadRulesUnion
}

// zarazConfigReturnTriggersLoadRuleJSON contains the JSON metadata for the struct
// [ZarazConfigReturnTriggersLoadRule]
type zarazConfigReturnTriggersLoadRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Settings    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r zarazConfigReturnTriggersLoadRuleJSON) RawJSON() string {
	return r.raw
}

func (r *ZarazConfigReturnTriggersLoadRule) UnmarshalJSON(data []byte) (err error) {
	*r = ZarazConfigReturnTriggersLoadRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZarazConfigReturnTriggersLoadRulesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ZarazConfigReturnTriggersLoadRulesZarazLoadRule],
// [ZarazConfigReturnTriggersLoadRulesZarazClickListenerRule],
// [ZarazConfigReturnTriggersLoadRulesZarazTimerRule],
// [ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRule],
// [ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRule],
// [ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRule],
// [ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRule].
func (r ZarazConfigReturnTriggersLoadRule) AsUnion() ZarazConfigReturnTriggersLoadRulesUnion {
	return r.union
}

// Union satisfied by [ZarazConfigReturnTriggersLoadRulesZarazLoadRule],
// [ZarazConfigReturnTriggersLoadRulesZarazClickListenerRule],
// [ZarazConfigReturnTriggersLoadRulesZarazTimerRule],
// [ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRule],
// [ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRule],
// [ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRule] or
// [ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRule].
type ZarazConfigReturnTriggersLoadRulesUnion interface {
	implementsZarazConfigReturnTriggersLoadRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigReturnTriggersLoadRulesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnTriggersLoadRulesZarazLoadRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnTriggersLoadRulesZarazClickListenerRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnTriggersLoadRulesZarazTimerRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRule{}),
		},
	)
}

type ZarazConfigReturnTriggersLoadRulesZarazLoadRule struct {
	ID    string                                              `json:"id,required"`
	Match string                                              `json:"match,required"`
	Op    ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                              `json:"value,required"`
	JSON  zarazConfigReturnTriggersLoadRulesZarazLoadRuleJSON `json:"-"`
}

// zarazConfigReturnTriggersLoadRulesZarazLoadRuleJSON contains the JSON metadata
// for the struct [ZarazConfigReturnTriggersLoadRulesZarazLoadRule]
type zarazConfigReturnTriggersLoadRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersLoadRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersLoadRulesZarazLoadRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnTriggersLoadRulesZarazLoadRule) implementsZarazConfigReturnTriggersLoadRule() {
}

type ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOp string

const (
	ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpContains           ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOp = "CONTAINS"
	ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpEquals             ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOp = "EQUALS"
	ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpStartsWith         ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpEndsWith           ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpMatchRegex         ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpNotMatchRegex      ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpGreaterThan        ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpLessThan           ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual    ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

func (r ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOp) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpContains, ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpEquals, ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpStartsWith, ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpEndsWith, ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpMatchRegex, ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpNotMatchRegex, ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpGreaterThan, ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual, ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpLessThan, ZarazConfigReturnTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersLoadRulesZarazClickListenerRule struct {
	ID       string                                                           `json:"id,required"`
	Action   ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazConfigReturnTriggersLoadRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazConfigReturnTriggersLoadRulesZarazClickListenerRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigReturnTriggersLoadRulesZarazClickListenerRule]
type zarazConfigReturnTriggersLoadRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersLoadRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersLoadRulesZarazClickListenerRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnTriggersLoadRulesZarazClickListenerRule) implementsZarazConfigReturnTriggersLoadRule() {
}

type ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleAction string

const (
	ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleActionClickListener ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleAction = "clickListener"
)

func (r ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleActionClickListener:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                               `json:"selector,required"`
	Type        ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                `json:"waitForTags,required"`
	JSON        zarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettings]
type zarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettingsType string

const (
	ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettingsTypeCss   ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettingsType = "css"
)

func (r ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettingsType) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath, ZarazConfigReturnTriggersLoadRulesZarazClickListenerRuleSettingsTypeCss:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersLoadRulesZarazTimerRule struct {
	ID       string                                                   `json:"id,required"`
	Action   ZarazConfigReturnTriggersLoadRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazConfigReturnTriggersLoadRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazConfigReturnTriggersLoadRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazConfigReturnTriggersLoadRulesZarazTimerRuleJSON contains the JSON metadata
// for the struct [ZarazConfigReturnTriggersLoadRulesZarazTimerRule]
type zarazConfigReturnTriggersLoadRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersLoadRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersLoadRulesZarazTimerRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnTriggersLoadRulesZarazTimerRule) implementsZarazConfigReturnTriggersLoadRule() {
}

type ZarazConfigReturnTriggersLoadRulesZarazTimerRuleAction string

const (
	ZarazConfigReturnTriggersLoadRulesZarazTimerRuleActionTimer ZarazConfigReturnTriggersLoadRulesZarazTimerRuleAction = "timer"
)

func (r ZarazConfigReturnTriggersLoadRulesZarazTimerRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersLoadRulesZarazTimerRuleActionTimer:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersLoadRulesZarazTimerRuleSettings struct {
	Interval int64                                                        `json:"interval,required"`
	Limit    int64                                                        `json:"limit,required"`
	JSON     zarazConfigReturnTriggersLoadRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazConfigReturnTriggersLoadRulesZarazTimerRuleSettingsJSON contains the JSON
// metadata for the struct
// [ZarazConfigReturnTriggersLoadRulesZarazTimerRuleSettings]
type zarazConfigReturnTriggersLoadRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersLoadRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersLoadRulesZarazTimerRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRule struct {
	ID       string                                                            `json:"id,required"`
	Action   ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRule]
type zarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRule) implementsZarazConfigReturnTriggersLoadRule() {
}

type ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleAction string

const (
	ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleAction = "formSubmission"
)

func (r ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                `json:"selector,required"`
	Validate bool                                                                  `json:"validate,required"`
	JSON     zarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleSettings]
type zarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRule struct {
	ID       string                                                           `json:"id,required"`
	Action   ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRule]
type zarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRule) implementsZarazConfigReturnTriggersLoadRule() {
}

type ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleAction string

const (
	ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleAction = "variableMatch"
)

func (r ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                               `json:"match,required"`
	Variable string                                                               `json:"variable,required"`
	JSON     zarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleSettings]
type zarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersLoadRulesZarazVariableMatchRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRule struct {
	ID       string                                                         `json:"id,required"`
	Action   ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleJSON contains the JSON
// metadata for the struct [ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRule]
type zarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRule) implementsZarazConfigReturnTriggersLoadRule() {
}

type ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleAction string

const (
	ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleAction = "scrollDepth"
)

func (r ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                             `json:"positions,required"`
	JSON      zarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleSettingsJSON contains the
// JSON metadata for the struct
// [ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleSettings]
type zarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersLoadRulesZarazScrollDepthRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRule struct {
	ID       string                                                               `json:"id,required"`
	Action   ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRule]
type zarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRule) implementsZarazConfigReturnTriggersLoadRule() {
}

type ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleAction string

const (
	ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

func (r ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleAction) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                   `json:"selector,required"`
	JSON     zarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleSettings]
type zarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnTriggersLoadRulesAction string

const (
	ZarazConfigReturnTriggersLoadRulesActionClickListener     ZarazConfigReturnTriggersLoadRulesAction = "clickListener"
	ZarazConfigReturnTriggersLoadRulesActionTimer             ZarazConfigReturnTriggersLoadRulesAction = "timer"
	ZarazConfigReturnTriggersLoadRulesActionFormSubmission    ZarazConfigReturnTriggersLoadRulesAction = "formSubmission"
	ZarazConfigReturnTriggersLoadRulesActionVariableMatch     ZarazConfigReturnTriggersLoadRulesAction = "variableMatch"
	ZarazConfigReturnTriggersLoadRulesActionScrollDepth       ZarazConfigReturnTriggersLoadRulesAction = "scrollDepth"
	ZarazConfigReturnTriggersLoadRulesActionElementVisibility ZarazConfigReturnTriggersLoadRulesAction = "elementVisibility"
)

func (r ZarazConfigReturnTriggersLoadRulesAction) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersLoadRulesActionClickListener, ZarazConfigReturnTriggersLoadRulesActionTimer, ZarazConfigReturnTriggersLoadRulesActionFormSubmission, ZarazConfigReturnTriggersLoadRulesActionVariableMatch, ZarazConfigReturnTriggersLoadRulesActionScrollDepth, ZarazConfigReturnTriggersLoadRulesActionElementVisibility:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersLoadRulesOp string

const (
	ZarazConfigReturnTriggersLoadRulesOpContains           ZarazConfigReturnTriggersLoadRulesOp = "CONTAINS"
	ZarazConfigReturnTriggersLoadRulesOpEquals             ZarazConfigReturnTriggersLoadRulesOp = "EQUALS"
	ZarazConfigReturnTriggersLoadRulesOpStartsWith         ZarazConfigReturnTriggersLoadRulesOp = "STARTS_WITH"
	ZarazConfigReturnTriggersLoadRulesOpEndsWith           ZarazConfigReturnTriggersLoadRulesOp = "ENDS_WITH"
	ZarazConfigReturnTriggersLoadRulesOpMatchRegex         ZarazConfigReturnTriggersLoadRulesOp = "MATCH_REGEX"
	ZarazConfigReturnTriggersLoadRulesOpNotMatchRegex      ZarazConfigReturnTriggersLoadRulesOp = "NOT_MATCH_REGEX"
	ZarazConfigReturnTriggersLoadRulesOpGreaterThan        ZarazConfigReturnTriggersLoadRulesOp = "GREATER_THAN"
	ZarazConfigReturnTriggersLoadRulesOpGreaterThanOrEqual ZarazConfigReturnTriggersLoadRulesOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigReturnTriggersLoadRulesOpLessThan           ZarazConfigReturnTriggersLoadRulesOp = "LESS_THAN"
	ZarazConfigReturnTriggersLoadRulesOpLessThanOrEqual    ZarazConfigReturnTriggersLoadRulesOp = "LESS_THAN_OR_EQUAL"
)

func (r ZarazConfigReturnTriggersLoadRulesOp) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersLoadRulesOpContains, ZarazConfigReturnTriggersLoadRulesOpEquals, ZarazConfigReturnTriggersLoadRulesOpStartsWith, ZarazConfigReturnTriggersLoadRulesOpEndsWith, ZarazConfigReturnTriggersLoadRulesOpMatchRegex, ZarazConfigReturnTriggersLoadRulesOpNotMatchRegex, ZarazConfigReturnTriggersLoadRulesOpGreaterThan, ZarazConfigReturnTriggersLoadRulesOpGreaterThanOrEqual, ZarazConfigReturnTriggersLoadRulesOpLessThan, ZarazConfigReturnTriggersLoadRulesOpLessThanOrEqual:
		return true
	}
	return false
}

type ZarazConfigReturnTriggersSystem string

const (
	ZarazConfigReturnTriggersSystemPageload ZarazConfigReturnTriggersSystem = "pageload"
)

func (r ZarazConfigReturnTriggersSystem) IsKnown() bool {
	switch r {
	case ZarazConfigReturnTriggersSystemPageload:
		return true
	}
	return false
}

type ZarazConfigReturnVariable struct {
	Name string                         `json:"name,required"`
	Type ZarazConfigReturnVariablesType `json:"type,required"`
	// This field can have the runtime type of [string],
	// [ZarazConfigReturnVariablesZarazWorkerVariableValue].
	Value interface{}                   `json:"value,required"`
	JSON  zarazConfigReturnVariableJSON `json:"-"`
	union ZarazConfigReturnVariablesUnion
}

// zarazConfigReturnVariableJSON contains the JSON metadata for the struct
// [ZarazConfigReturnVariable]
type zarazConfigReturnVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r zarazConfigReturnVariableJSON) RawJSON() string {
	return r.raw
}

func (r *ZarazConfigReturnVariable) UnmarshalJSON(data []byte) (err error) {
	*r = ZarazConfigReturnVariable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZarazConfigReturnVariablesUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ZarazConfigReturnVariablesZarazStringVariable],
// [ZarazConfigReturnVariablesZarazSecretVariable],
// [ZarazConfigReturnVariablesZarazWorkerVariable].
func (r ZarazConfigReturnVariable) AsUnion() ZarazConfigReturnVariablesUnion {
	return r.union
}

// Union satisfied by [ZarazConfigReturnVariablesZarazStringVariable],
// [ZarazConfigReturnVariablesZarazSecretVariable] or
// [ZarazConfigReturnVariablesZarazWorkerVariable].
type ZarazConfigReturnVariablesUnion interface {
	implementsZarazConfigReturnVariable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigReturnVariablesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZarazConfigReturnVariablesZarazStringVariable{}),
			DiscriminatorValue: "string",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZarazConfigReturnVariablesZarazSecretVariable{}),
			DiscriminatorValue: "secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZarazConfigReturnVariablesZarazWorkerVariable{}),
			DiscriminatorValue: "worker",
		},
	)
}

type ZarazConfigReturnVariablesZarazStringVariable struct {
	Name  string                                            `json:"name,required"`
	Type  ZarazConfigReturnVariablesZarazStringVariableType `json:"type,required"`
	Value string                                            `json:"value,required"`
	JSON  zarazConfigReturnVariablesZarazStringVariableJSON `json:"-"`
}

// zarazConfigReturnVariablesZarazStringVariableJSON contains the JSON metadata for
// the struct [ZarazConfigReturnVariablesZarazStringVariable]
type zarazConfigReturnVariablesZarazStringVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnVariablesZarazStringVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnVariablesZarazStringVariableJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnVariablesZarazStringVariable) implementsZarazConfigReturnVariable() {}

type ZarazConfigReturnVariablesZarazStringVariableType string

const (
	ZarazConfigReturnVariablesZarazStringVariableTypeString ZarazConfigReturnVariablesZarazStringVariableType = "string"
)

func (r ZarazConfigReturnVariablesZarazStringVariableType) IsKnown() bool {
	switch r {
	case ZarazConfigReturnVariablesZarazStringVariableTypeString:
		return true
	}
	return false
}

type ZarazConfigReturnVariablesZarazSecretVariable struct {
	Name  string                                            `json:"name,required"`
	Type  ZarazConfigReturnVariablesZarazSecretVariableType `json:"type,required"`
	Value string                                            `json:"value,required"`
	JSON  zarazConfigReturnVariablesZarazSecretVariableJSON `json:"-"`
}

// zarazConfigReturnVariablesZarazSecretVariableJSON contains the JSON metadata for
// the struct [ZarazConfigReturnVariablesZarazSecretVariable]
type zarazConfigReturnVariablesZarazSecretVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnVariablesZarazSecretVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnVariablesZarazSecretVariableJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnVariablesZarazSecretVariable) implementsZarazConfigReturnVariable() {}

type ZarazConfigReturnVariablesZarazSecretVariableType string

const (
	ZarazConfigReturnVariablesZarazSecretVariableTypeSecret ZarazConfigReturnVariablesZarazSecretVariableType = "secret"
)

func (r ZarazConfigReturnVariablesZarazSecretVariableType) IsKnown() bool {
	switch r {
	case ZarazConfigReturnVariablesZarazSecretVariableTypeSecret:
		return true
	}
	return false
}

type ZarazConfigReturnVariablesZarazWorkerVariable struct {
	Name  string                                             `json:"name,required"`
	Type  ZarazConfigReturnVariablesZarazWorkerVariableType  `json:"type,required"`
	Value ZarazConfigReturnVariablesZarazWorkerVariableValue `json:"value,required"`
	JSON  zarazConfigReturnVariablesZarazWorkerVariableJSON  `json:"-"`
}

// zarazConfigReturnVariablesZarazWorkerVariableJSON contains the JSON metadata for
// the struct [ZarazConfigReturnVariablesZarazWorkerVariable]
type zarazConfigReturnVariablesZarazWorkerVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnVariablesZarazWorkerVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnVariablesZarazWorkerVariableJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnVariablesZarazWorkerVariable) implementsZarazConfigReturnVariable() {}

type ZarazConfigReturnVariablesZarazWorkerVariableType string

const (
	ZarazConfigReturnVariablesZarazWorkerVariableTypeWorker ZarazConfigReturnVariablesZarazWorkerVariableType = "worker"
)

func (r ZarazConfigReturnVariablesZarazWorkerVariableType) IsKnown() bool {
	switch r {
	case ZarazConfigReturnVariablesZarazWorkerVariableTypeWorker:
		return true
	}
	return false
}

type ZarazConfigReturnVariablesZarazWorkerVariableValue struct {
	EscapedWorkerName string                                                 `json:"escapedWorkerName,required"`
	WorkerTag         string                                                 `json:"workerTag,required"`
	JSON              zarazConfigReturnVariablesZarazWorkerVariableValueJSON `json:"-"`
}

// zarazConfigReturnVariablesZarazWorkerVariableValueJSON contains the JSON
// metadata for the struct [ZarazConfigReturnVariablesZarazWorkerVariableValue]
type zarazConfigReturnVariablesZarazWorkerVariableValueJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazConfigReturnVariablesZarazWorkerVariableValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnVariablesZarazWorkerVariableValueJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnVariablesType string

const (
	ZarazConfigReturnVariablesTypeString ZarazConfigReturnVariablesType = "string"
	ZarazConfigReturnVariablesTypeSecret ZarazConfigReturnVariablesType = "secret"
	ZarazConfigReturnVariablesTypeWorker ZarazConfigReturnVariablesType = "worker"
)

func (r ZarazConfigReturnVariablesType) IsKnown() bool {
	switch r {
	case ZarazConfigReturnVariablesTypeString, ZarazConfigReturnVariablesTypeSecret, ZarazConfigReturnVariablesTypeWorker:
		return true
	}
	return false
}

// Cloudflare Monitoring settings.
type ZarazConfigReturnAnalytics struct {
	// Consent purpose assigned to Monitoring.
	DefaultPurpose string `json:"defaultPurpose"`
	// Whether Advanced Monitoring reports are enabled.
	Enabled bool `json:"enabled"`
	// Session expiration time (seconds).
	SessionExpTime int64                          `json:"sessionExpTime"`
	JSON           zarazConfigReturnAnalyticsJSON `json:"-"`
}

// zarazConfigReturnAnalyticsJSON contains the JSON metadata for the struct
// [ZarazConfigReturnAnalytics]
type zarazConfigReturnAnalyticsJSON struct {
	DefaultPurpose apijson.Field
	Enabled        apijson.Field
	SessionExpTime apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZarazConfigReturnAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnAnalyticsJSON) RawJSON() string {
	return r.raw
}

// Consent management configuration.
type ZarazConfigReturnConsent struct {
	Enabled                bool                                           `json:"enabled,required"`
	ButtonTextTranslations ZarazConfigReturnConsentButtonTextTranslations `json:"buttonTextTranslations"`
	CompanyEmail           string                                         `json:"companyEmail"`
	CompanyName            string                                         `json:"companyName"`
	CompanyStreetAddress   string                                         `json:"companyStreetAddress"`
	ConsentModalIntroHTML  string                                         `json:"consentModalIntroHTML"`
	// Object where keys are language codes
	ConsentModalIntroHTMLWithTranslations map[string]string `json:"consentModalIntroHTMLWithTranslations"`
	CookieName                            string            `json:"cookieName"`
	CustomCss                             string            `json:"customCSS"`
	CustomIntroDisclaimerDismissed        bool              `json:"customIntroDisclaimerDismissed"`
	DefaultLanguage                       string            `json:"defaultLanguage"`
	HideModal                             bool              `json:"hideModal"`
	// Object where keys are purpose alpha-numeric IDs
	Purposes map[string]ZarazConfigReturnConsentPurpose `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations map[string]ZarazConfigReturnConsentPurposesWithTranslation `json:"purposesWithTranslations"`
	TcfCompliant             bool                                                       `json:"tcfCompliant"`
	JSON                     zarazConfigReturnConsentJSON                               `json:"-"`
}

// zarazConfigReturnConsentJSON contains the JSON metadata for the struct
// [ZarazConfigReturnConsent]
type zarazConfigReturnConsentJSON struct {
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

func (r *ZarazConfigReturnConsent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnConsentJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll map[string]string `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices map[string]string `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll map[string]string                                  `json:"reject_all,required"`
	JSON      zarazConfigReturnConsentButtonTextTranslationsJSON `json:"-"`
}

// zarazConfigReturnConsentButtonTextTranslationsJSON contains the JSON metadata
// for the struct [ZarazConfigReturnConsentButtonTextTranslations]
type zarazConfigReturnConsentButtonTextTranslationsJSON struct {
	AcceptAll        apijson.Field
	ConfirmMyChoices apijson.Field
	RejectAll        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReturnConsentButtonTextTranslations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnConsentButtonTextTranslationsJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnConsentPurpose struct {
	Description string                              `json:"description,required"`
	Name        string                              `json:"name,required"`
	JSON        zarazConfigReturnConsentPurposeJSON `json:"-"`
}

// zarazConfigReturnConsentPurposeJSON contains the JSON metadata for the struct
// [ZarazConfigReturnConsentPurpose]
type zarazConfigReturnConsentPurposeJSON struct {
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnConsentPurpose) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnConsentPurposeJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnConsentPurposesWithTranslation struct {
	// Object where keys are language codes
	Description map[string]string `json:"description,required"`
	// Object where keys are language codes
	Name  map[string]string                                   `json:"name,required"`
	Order int64                                               `json:"order,required"`
	JSON  zarazConfigReturnConsentPurposesWithTranslationJSON `json:"-"`
}

// zarazConfigReturnConsentPurposesWithTranslationJSON contains the JSON metadata
// for the struct [ZarazConfigReturnConsentPurposesWithTranslation]
type zarazConfigReturnConsentPurposesWithTranslationJSON struct {
	Description apijson.Field
	Name        apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturnConsentPurposesWithTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnConsentPurposesWithTranslationJSON) RawJSON() string {
	return r.raw
}

type ZarazMessagesItems struct {
	Code             int64                    `json:"code,required"`
	Message          string                   `json:"message,required"`
	DocumentationURL string                   `json:"documentation_url"`
	Source           ZarazMessagesItemsSource `json:"source"`
	JSON             zarazMessagesItemsJSON   `json:"-"`
}

// zarazMessagesItemsJSON contains the JSON metadata for the struct
// [ZarazMessagesItems]
type zarazMessagesItemsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazMessagesItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazMessagesItemsJSON) RawJSON() string {
	return r.raw
}

type ZarazMessagesItemsSource struct {
	Pointer string                       `json:"pointer"`
	JSON    zarazMessagesItemsSourceJSON `json:"-"`
}

// zarazMessagesItemsSourceJSON contains the JSON metadata for the struct
// [ZarazMessagesItemsSource]
type zarazMessagesItemsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazMessagesItemsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazMessagesItemsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingZarazPublishResponse struct {
	Errors   []ZarazMessagesItems `json:"errors,required"`
	Messages []ZarazMessagesItems `json:"messages,required"`
	Result   string               `json:"result,required"`
	// Whether the API call was successful
	Success bool                                `json:"success,required"`
	JSON    zoneSettingZarazPublishResponseJSON `json:"-"`
}

// zoneSettingZarazPublishResponseJSON contains the JSON metadata for the struct
// [ZoneSettingZarazPublishResponse]
type zoneSettingZarazPublishResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZarazPublishResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingZarazPublishResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingZarazPublishParams struct {
	// Zaraz configuration description.
	Body string `json:"body"`
}

func (r ZoneSettingZarazPublishParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
