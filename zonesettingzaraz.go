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

type ZarazAPIResponseCommon struct {
	Errors   []ZarazMessagesItems `json:"errors,required"`
	Messages []ZarazMessagesItems `json:"messages,required"`
	// Whether the API call was successful
	Success bool                       `json:"success,required"`
	JSON    zarazAPIResponseCommonJSON `json:"-"`
}

// zarazAPIResponseCommonJSON contains the JSON metadata for the struct
// [ZarazAPIResponseCommon]
type zarazAPIResponseCommonJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazAPIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazAPIResponseCommonJSON) RawJSON() string {
	return r.raw
}

// Zaraz configuration
type ZarazConfigReturn struct {
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools map[string]ZarazConfigReturnTool `json:"tools"`
	JSON  zarazConfigReturnJSON            `json:"-"`
	ZarazConfigBase
}

// zarazConfigReturnJSON contains the JSON metadata for the struct
// [ZarazConfigReturn]
type zarazConfigReturnJSON struct {
	Tools       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnTool struct {
	// This field can have the runtime type of [[]string].
	BlockingTriggers interface{} `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// This field can have the runtime type of
	// [map[string]ZarazConfigReturnToolsObjectDefaultFieldsUnion].
	DefaultFields interface{} `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// This field can have the runtime type of [[]string].
	Permissions interface{} `json:"permissions,required"`
	// This field can have the runtime type of
	// [map[string]ZarazConfigReturnToolsObjectSettingsUnion].
	Settings interface{}                `json:"settings,required"`
	Type     ZarazConfigReturnToolsType `json:"type,required"`
	// This field can have the runtime type of
	// [map[string]ZarazConfigReturnToolsObjectAction].
	Actions interface{} `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// This field can have the runtime type of
	// [[]ZarazConfigReturnToolsObjectNeoEvent].
	NeoEvents interface{} `json:"neoEvents"`
	// Vendor name for TCF compliant consent modal, required for Custom Managed
	// Components and Custom HTML tool with a defaultPurpose assigned
	VendorName string `json:"vendorName"`
	// Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom
	// Managed Components and Custom HTML tool with a defaultPurpose assigned
	VendorPolicyURL string `json:"vendorPolicyUrl"`
	// This field can have the runtime type of [ZarazConfigReturnToolsObjectWorker].
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
// Possible runtime types of the union are [ZarazConfigReturnToolsObject],
// [ZarazConfigReturnToolsObject].
func (r ZarazConfigReturnTool) AsUnion() ZarazConfigReturnToolsUnion {
	return r.union
}

// Union satisfied by [ZarazConfigReturnToolsObject] or
// [ZarazConfigReturnToolsObject].
type ZarazConfigReturnToolsUnion interface {
	implementsZarazConfigReturnTool()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigReturnToolsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnToolsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZarazConfigReturnToolsObject{}),
		},
	)
}

type ZarazConfigReturnToolsObject struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazConfigReturnToolsObjectDefaultFieldsUnion `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazConfigReturnToolsObjectSettingsUnion `json:"settings,required"`
	Type     ZarazConfigReturnToolsObjectType                     `json:"type,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazConfigReturnToolsObjectAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazConfigReturnToolsObjectNeoEvent `json:"neoEvents"`
	// Vendor name for TCF compliant consent modal, required for Custom Managed
	// Components and Custom HTML tool with a defaultPurpose assigned
	VendorName string `json:"vendorName"`
	// Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom
	// Managed Components and Custom HTML tool with a defaultPurpose assigned
	VendorPolicyURL string                           `json:"vendorPolicyUrl"`
	JSON            zarazConfigReturnToolsObjectJSON `json:"-"`
}

// zarazConfigReturnToolsObjectJSON contains the JSON metadata for the struct
// [ZarazConfigReturnToolsObject]
type zarazConfigReturnToolsObjectJSON struct {
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

func (r *ZarazConfigReturnToolsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnToolsObjectJSON) RawJSON() string {
	return r.raw
}

func (r ZarazConfigReturnToolsObject) implementsZarazConfigReturnTool() {}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazConfigReturnToolsObjectDefaultFieldsUnion interface {
	ImplementsZarazConfigReturnToolsObjectDefaultFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigReturnToolsObjectDefaultFieldsUnion)(nil)).Elem(),
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
type ZarazConfigReturnToolsObjectSettingsUnion interface {
	ImplementsZarazConfigReturnToolsObjectSettingsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigReturnToolsObjectSettingsUnion)(nil)).Elem(),
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

type ZarazConfigReturnToolsObjectType string

const (
	ZarazConfigReturnToolsObjectTypeComponent ZarazConfigReturnToolsObjectType = "component"
)

func (r ZarazConfigReturnToolsObjectType) IsKnown() bool {
	switch r {
	case ZarazConfigReturnToolsObjectTypeComponent:
		return true
	}
	return false
}

type ZarazConfigReturnToolsObjectAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                               `json:"firingTriggers,required"`
	JSON           zarazConfigReturnToolsObjectActionJSON `json:"-"`
}

// zarazConfigReturnToolsObjectActionJSON contains the JSON metadata for the struct
// [ZarazConfigReturnToolsObjectAction]
type zarazConfigReturnToolsObjectActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReturnToolsObjectAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnToolsObjectActionJSON) RawJSON() string {
	return r.raw
}

type ZarazConfigReturnToolsObjectNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                 `json:"firingTriggers,required"`
	JSON           zarazConfigReturnToolsObjectNeoEventJSON `json:"-"`
}

// zarazConfigReturnToolsObjectNeoEventJSON contains the JSON metadata for the
// struct [ZarazConfigReturnToolsObjectNeoEvent]
type zarazConfigReturnToolsObjectNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReturnToolsObjectNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazConfigReturnToolsObjectNeoEventJSON) RawJSON() string {
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

type ZarazMessagesItems struct {
	Code    int64                  `json:"code,required"`
	Message string                 `json:"message,required"`
	JSON    zarazMessagesItemsJSON `json:"-"`
}

// zarazMessagesItemsJSON contains the JSON metadata for the struct
// [ZarazMessagesItems]
type zarazMessagesItemsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazMessagesItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazMessagesItemsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingZarazPublishResponse struct {
	Result string                              `json:"result"`
	JSON   zoneSettingZarazPublishResponseJSON `json:"-"`
	ZarazAPIResponseCommon
}

// zoneSettingZarazPublishResponseJSON contains the JSON metadata for the struct
// [ZoneSettingZarazPublishResponse]
type zoneSettingZarazPublishResponseJSON struct {
	Result      apijson.Field
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
