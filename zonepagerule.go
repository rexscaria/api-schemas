// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/stainless-sdks/cf-rex-go/shared"
	"github.com/tidwall/gjson"
)

// ZonePageruleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePageruleService] method instead.
type ZonePageruleService struct {
	Options []option.RequestOption
}

// NewZonePageruleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZonePageruleService(opts ...option.RequestOption) (r *ZonePageruleService) {
	r = &ZonePageruleService{}
	r.Options = opts
	return
}

// Creates a new Page Rule.
func (r *ZonePageruleService) New(ctx context.Context, zoneID string, body ZonePageruleNewParams, opts ...option.RequestOption) (res *ZonePageruleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the details of a Page Rule.
func (r *ZonePageruleService) Get(ctx context.Context, zoneID string, pageruleID string, opts ...option.RequestOption) (res *ZonePageruleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", zoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Replaces the configuration of an existing Page Rule. The configuration of the
// updated Page Rule will exactly match the data passed in the API request.
func (r *ZonePageruleService) Update(ctx context.Context, zoneID string, pageruleID string, body ZonePageruleUpdateParams, opts ...option.RequestOption) (res *ZonePageruleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", zoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches Page Rules in a zone.
func (r *ZonePageruleService) List(ctx context.Context, zoneID string, query ZonePageruleListParams, opts ...option.RequestOption) (res *ZonePageruleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing Page Rule.
func (r *ZonePageruleService) Delete(ctx context.Context, zoneID string, pageruleID string, opts ...option.RequestOption) (res *ZonePageruleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", zoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Updates one or more fields of an existing Page Rule.
func (r *ZonePageruleService) Edit(ctx context.Context, zoneID string, pageruleID string, body ZonePageruleEditParams, opts ...option.RequestOption) (res *ZonePageruleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", zoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of settings (and their details) that Page Rules can apply to
// matching requests.
func (r *ZonePageruleService) ListSettings(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZonePageruleListSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIResponseSingleZones struct {
	Errors   []MessagesZonesItem `json:"errors,required"`
	Messages []MessagesZonesItem `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseSingleZonesSuccess `json:"success,required"`
	JSON    apiResponseSingleZonesJSON    `json:"-"`
}

// apiResponseSingleZonesJSON contains the JSON metadata for the struct
// [APIResponseSingleZones]
type apiResponseSingleZonesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleZonesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseSingleZonesSuccess bool

const (
	APIResponseSingleZonesSuccessTrue APIResponseSingleZonesSuccess = true
)

func (r APIResponseSingleZonesSuccess) IsKnown() bool {
	switch r {
	case APIResponseSingleZonesSuccessTrue:
		return true
	}
	return false
}

type APIResponseZonesSchemas struct {
	Errors   []MessagesZonesItem `json:"errors,required"`
	Messages []MessagesZonesItem `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseZonesSchemasSuccess `json:"success,required"`
	JSON    apiResponseZonesSchemasJSON    `json:"-"`
}

// apiResponseZonesSchemasJSON contains the JSON metadata for the struct
// [APIResponseZonesSchemas]
type apiResponseZonesSchemasJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseZonesSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseZonesSchemasJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseZonesSchemasSuccess bool

const (
	APIResponseZonesSchemasSuccessTrue APIResponseZonesSchemasSuccess = true
)

func (r APIResponseZonesSchemasSuccess) IsKnown() bool {
	switch r {
	case APIResponseZonesSchemasSuccessTrue:
		return true
	}
	return false
}

type PageRule struct {
	// Identifier
	ID string `json:"id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions []ZoneAction `json:"actions,required"`
	// The timestamp of when the Page Rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The timestamp of when the Page Rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority int64 `json:"priority,required"`
	// The status of the Page Rule.
	Status PageRuleStatus `json:"status,required"`
	// The rule targets to evaluate on each request.
	Targets []RequestConditionTarget `json:"targets,required"`
	JSON    pageRuleJSON             `json:"-"`
}

// pageRuleJSON contains the JSON metadata for the struct [PageRule]
type pageRuleJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Priority    apijson.Field
	Status      apijson.Field
	Targets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleJSON) RawJSON() string {
	return r.raw
}

// The status of the Page Rule.
type PageRuleStatus string

const (
	PageRuleStatusActive   PageRuleStatus = "active"
	PageRuleStatusDisabled PageRuleStatus = "disabled"
)

func (r PageRuleStatus) IsKnown() bool {
	switch r {
	case PageRuleStatusActive, PageRuleStatusDisabled:
		return true
	}
	return false
}

// URL target.
type RequestConditionTarget struct {
	// The constraint of a target.
	Constraint RequestConditionTargetConstraint `json:"constraint,required"`
	// A target based on the URL of the request.
	Target RequestConditionTargetTarget `json:"target,required"`
	JSON   requestConditionTargetJSON   `json:"-"`
}

// requestConditionTargetJSON contains the JSON metadata for the struct
// [RequestConditionTarget]
type requestConditionTargetJSON struct {
	Constraint  apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestConditionTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestConditionTargetJSON) RawJSON() string {
	return r.raw
}

// The constraint of a target.
type RequestConditionTargetConstraint struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator RequestConditionTargetConstraintOperator `json:"operator,required"`
	// The value to apply the operator to.
	Value string                               `json:"value,required"`
	JSON  requestConditionTargetConstraintJSON `json:"-"`
}

// requestConditionTargetConstraintJSON contains the JSON metadata for the struct
// [RequestConditionTargetConstraint]
type requestConditionTargetConstraintJSON struct {
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestConditionTargetConstraint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestConditionTargetConstraintJSON) RawJSON() string {
	return r.raw
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type RequestConditionTargetConstraintOperator string

const (
	RequestConditionTargetConstraintOperatorMatches    RequestConditionTargetConstraintOperator = "matches"
	RequestConditionTargetConstraintOperatorContains   RequestConditionTargetConstraintOperator = "contains"
	RequestConditionTargetConstraintOperatorEquals     RequestConditionTargetConstraintOperator = "equals"
	RequestConditionTargetConstraintOperatorNotEqual   RequestConditionTargetConstraintOperator = "not_equal"
	RequestConditionTargetConstraintOperatorNotContain RequestConditionTargetConstraintOperator = "not_contain"
)

func (r RequestConditionTargetConstraintOperator) IsKnown() bool {
	switch r {
	case RequestConditionTargetConstraintOperatorMatches, RequestConditionTargetConstraintOperatorContains, RequestConditionTargetConstraintOperatorEquals, RequestConditionTargetConstraintOperatorNotEqual, RequestConditionTargetConstraintOperatorNotContain:
		return true
	}
	return false
}

// A target based on the URL of the request.
type RequestConditionTargetTarget string

const (
	RequestConditionTargetTargetURL RequestConditionTargetTarget = "url"
)

func (r RequestConditionTargetTarget) IsKnown() bool {
	switch r {
	case RequestConditionTargetTargetURL:
		return true
	}
	return false
}

// URL target.
type RequestConditionTargetParam struct {
	// The constraint of a target.
	Constraint param.Field[RequestConditionTargetConstraintParam] `json:"constraint,required"`
	// A target based on the URL of the request.
	Target param.Field[RequestConditionTargetTarget] `json:"target,required"`
}

func (r RequestConditionTargetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The constraint of a target.
type RequestConditionTargetConstraintParam struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator param.Field[RequestConditionTargetConstraintOperator] `json:"operator,required"`
	// The value to apply the operator to.
	Value param.Field[string] `json:"value,required"`
}

func (r RequestConditionTargetConstraintParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAction struct {
	// If enabled, any ` http://“ URL is converted to  `https://` through a 301
	// redirect.
	ID ZoneActionID `json:"id"`
	// This field can have the runtime type of [ZoneActionAutomaticHTTPSRewritesValue],
	// [int64], [ZoneActionBrowserCheckValue], [string],
	// [ZoneActionCacheByDeviceTypeValue], [ZoneActionCacheDeceptionArmorValue],
	// [ZoneActionCacheKeyFieldsValue], [ZoneActionCacheLevelValue],
	// [map[string]ZoneActionCacheTtlByStatusValueUnion],
	// [ZoneActionEmailObfuscationValue], [ZoneActionExplicitCacheControlValue],
	// [ZoneActionForwardingURLValue], [ZoneActionIPGeolocationValue],
	// [ZoneActionMirageValue], [ZoneActionOpportunisticEncryptionValue],
	// [ZoneActionOriginErrorPagePassThruValue], [ZoneActionPolishValue],
	// [ZoneActionRespectStrongEtagValue], [ZoneActionResponseBufferingValue],
	// [ZoneActionRocketLoaderValue], [ZoneActionSecurityLevelValue],
	// [ZoneActionSortQueryStringForCacheValue], [ZoneActionSslValue],
	// [ZoneActionTrueClientIPHeaderValue], [ZoneActionWafValue].
	Value interface{}    `json:"value"`
	JSON  zoneActionJSON `json:"-"`
	union ZoneActionUnion
}

// zoneActionJSON contains the JSON metadata for the struct [ZoneAction]
type zoneActionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r zoneActionJSON) RawJSON() string {
	return r.raw
}

func (r *ZoneAction) UnmarshalJSON(data []byte) (err error) {
	*r = ZoneAction{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZoneActionUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [ZoneActionAlwaysUseHTTPS],
// [ZoneActionAutomaticHTTPSRewrites], [ZoneActionBrowserCacheTtl],
// [ZoneActionBrowserCheck], [ZoneActionBypassCacheOnCookie],
// [ZoneActionCacheByDeviceType], [ZoneActionCacheDeceptionArmor],
// [ZoneActionCacheKeyFields], [ZoneActionCacheLevel], [ZoneActionCacheOnCookie],
// [ZoneActionCacheTtlByStatus], [ZoneActionDisableApps],
// [ZoneActionDisablePerformance], [ZoneActionDisableSecurity],
// [ZoneActionDisableZaraz], [ZoneActionEdgeCacheTtl],
// [ZoneActionEmailObfuscation], [ZoneActionExplicitCacheControl],
// [ZoneActionForwardingURL], [ZoneActionHostHeaderOverride],
// [ZoneActionIPGeolocation], [ZoneActionMirage],
// [ZoneActionOpportunisticEncryption], [ZoneActionOriginErrorPagePassThru],
// [ZoneActionPolish], [ZoneActionResolveOverride], [ZoneActionRespectStrongEtag],
// [ZoneActionResponseBuffering], [ZoneActionRocketLoader],
// [ZoneActionSecurityLevel], [ZoneActionSortQueryStringForCache], [ZoneActionSsl],
// [ZoneActionTrueClientIPHeader], [ZoneActionWaf].
func (r ZoneAction) AsUnion() ZoneActionUnion {
	return r.union
}

// Union satisfied by [ZoneActionAlwaysUseHTTPS],
// [ZoneActionAutomaticHTTPSRewrites], [ZoneActionBrowserCacheTtl],
// [ZoneActionBrowserCheck], [ZoneActionBypassCacheOnCookie],
// [ZoneActionCacheByDeviceType], [ZoneActionCacheDeceptionArmor],
// [ZoneActionCacheKeyFields], [ZoneActionCacheLevel], [ZoneActionCacheOnCookie],
// [ZoneActionCacheTtlByStatus], [ZoneActionDisableApps],
// [ZoneActionDisablePerformance], [ZoneActionDisableSecurity],
// [ZoneActionDisableZaraz], [ZoneActionEdgeCacheTtl],
// [ZoneActionEmailObfuscation], [ZoneActionExplicitCacheControl],
// [ZoneActionForwardingURL], [ZoneActionHostHeaderOverride],
// [ZoneActionIPGeolocation], [ZoneActionMirage],
// [ZoneActionOpportunisticEncryption], [ZoneActionOriginErrorPagePassThru],
// [ZoneActionPolish], [ZoneActionResolveOverride], [ZoneActionRespectStrongEtag],
// [ZoneActionResponseBuffering], [ZoneActionRocketLoader],
// [ZoneActionSecurityLevel], [ZoneActionSortQueryStringForCache], [ZoneActionSsl],
// [ZoneActionTrueClientIPHeader] or [ZoneActionWaf].
type ZoneActionUnion interface {
	implementsZoneAction()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneActionUnion)(nil)).Elem(),
		"id",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionAlwaysUseHTTPS{}),
			DiscriminatorValue: "always_use_https",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionAutomaticHTTPSRewrites{}),
			DiscriminatorValue: "automatic_https_rewrites",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionBrowserCacheTtl{}),
			DiscriminatorValue: "browser_cache_ttl",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionBrowserCheck{}),
			DiscriminatorValue: "browser_check",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionBypassCacheOnCookie{}),
			DiscriminatorValue: "bypass_cache_on_cookie",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionCacheByDeviceType{}),
			DiscriminatorValue: "cache_by_device_type",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionCacheDeceptionArmor{}),
			DiscriminatorValue: "cache_deception_armor",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionCacheKeyFields{}),
			DiscriminatorValue: "cache_key_fields",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionCacheLevel{}),
			DiscriminatorValue: "cache_level",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionCacheOnCookie{}),
			DiscriminatorValue: "cache_on_cookie",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionCacheTtlByStatus{}),
			DiscriminatorValue: "cache_ttl_by_status",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionDisableApps{}),
			DiscriminatorValue: "disable_apps",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionDisablePerformance{}),
			DiscriminatorValue: "disable_performance",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionDisableSecurity{}),
			DiscriminatorValue: "disable_security",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionDisableZaraz{}),
			DiscriminatorValue: "disable_zaraz",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionEdgeCacheTtl{}),
			DiscriminatorValue: "edge_cache_ttl",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionEmailObfuscation{}),
			DiscriminatorValue: "email_obfuscation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionExplicitCacheControl{}),
			DiscriminatorValue: "explicit_cache_control",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionForwardingURL{}),
			DiscriminatorValue: "forwarding_url",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionHostHeaderOverride{}),
			DiscriminatorValue: "host_header_override",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionIPGeolocation{}),
			DiscriminatorValue: "ip_geolocation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionMirage{}),
			DiscriminatorValue: "mirage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionOpportunisticEncryption{}),
			DiscriminatorValue: "opportunistic_encryption",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionOriginErrorPagePassThru{}),
			DiscriminatorValue: "origin_error_page_pass_thru",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionPolish{}),
			DiscriminatorValue: "polish",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionResolveOverride{}),
			DiscriminatorValue: "resolve_override",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionRespectStrongEtag{}),
			DiscriminatorValue: "respect_strong_etag",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionResponseBuffering{}),
			DiscriminatorValue: "response_buffering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionRocketLoader{}),
			DiscriminatorValue: "rocket_loader",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionSecurityLevel{}),
			DiscriminatorValue: "security_level",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionSortQueryStringForCache{}),
			DiscriminatorValue: "sort_query_string_for_cache",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionSsl{}),
			DiscriminatorValue: "ssl",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionTrueClientIPHeader{}),
			DiscriminatorValue: "true_client_ip_header",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneActionWaf{}),
			DiscriminatorValue: "waf",
		},
	)
}

type ZoneActionAlwaysUseHTTPS struct {
	// If enabled, any ` http://“ URL is converted to  `https://` through a 301
	// redirect.
	ID   ZoneActionAlwaysUseHTTPSID   `json:"id"`
	JSON zoneActionAlwaysUseHTTPSJSON `json:"-"`
}

// zoneActionAlwaysUseHTTPSJSON contains the JSON metadata for the struct
// [ZoneActionAlwaysUseHTTPS]
type zoneActionAlwaysUseHTTPSJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionAlwaysUseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionAlwaysUseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionAlwaysUseHTTPS) implementsZoneAction() {}

// If enabled, any ` http://“ URL is converted to  `https://` through a 301
// redirect.
type ZoneActionAlwaysUseHTTPSID string

const (
	ZoneActionAlwaysUseHTTPSIDAlwaysUseHTTPS ZoneActionAlwaysUseHTTPSID = "always_use_https"
)

func (r ZoneActionAlwaysUseHTTPSID) IsKnown() bool {
	switch r {
	case ZoneActionAlwaysUseHTTPSIDAlwaysUseHTTPS:
		return true
	}
	return false
}

type ZoneActionAutomaticHTTPSRewrites struct {
	// Turn on or off Automatic HTTPS Rewrites.
	ID ZoneActionAutomaticHTTPSRewritesID `json:"id"`
	// The status of Automatic HTTPS Rewrites.
	Value ZoneActionAutomaticHTTPSRewritesValue `json:"value"`
	JSON  zoneActionAutomaticHTTPSRewritesJSON  `json:"-"`
}

// zoneActionAutomaticHTTPSRewritesJSON contains the JSON metadata for the struct
// [ZoneActionAutomaticHTTPSRewrites]
type zoneActionAutomaticHTTPSRewritesJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionAutomaticHTTPSRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionAutomaticHTTPSRewritesJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionAutomaticHTTPSRewrites) implementsZoneAction() {}

// Turn on or off Automatic HTTPS Rewrites.
type ZoneActionAutomaticHTTPSRewritesID string

const (
	ZoneActionAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites ZoneActionAutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

func (r ZoneActionAutomaticHTTPSRewritesID) IsKnown() bool {
	switch r {
	case ZoneActionAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites:
		return true
	}
	return false
}

// The status of Automatic HTTPS Rewrites.
type ZoneActionAutomaticHTTPSRewritesValue string

const (
	ZoneActionAutomaticHTTPSRewritesValueOn  ZoneActionAutomaticHTTPSRewritesValue = "on"
	ZoneActionAutomaticHTTPSRewritesValueOff ZoneActionAutomaticHTTPSRewritesValue = "off"
)

func (r ZoneActionAutomaticHTTPSRewritesValue) IsKnown() bool {
	switch r {
	case ZoneActionAutomaticHTTPSRewritesValueOn, ZoneActionAutomaticHTTPSRewritesValueOff:
		return true
	}
	return false
}

type ZoneActionBrowserCacheTtl struct {
	// Control how long resources cached by client browsers remain valid.
	ID ZoneActionBrowserCacheTtlID `json:"id"`
	// The number of seconds to cache resources for. The API prohibits setting this to
	// 0 for non-Enterprise domains.
	Value int64                         `json:"value"`
	JSON  zoneActionBrowserCacheTtlJSON `json:"-"`
}

// zoneActionBrowserCacheTtlJSON contains the JSON metadata for the struct
// [ZoneActionBrowserCacheTtl]
type zoneActionBrowserCacheTtlJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionBrowserCacheTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionBrowserCacheTtlJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionBrowserCacheTtl) implementsZoneAction() {}

// Control how long resources cached by client browsers remain valid.
type ZoneActionBrowserCacheTtlID string

const (
	ZoneActionBrowserCacheTtlIDBrowserCacheTtl ZoneActionBrowserCacheTtlID = "browser_cache_ttl"
)

func (r ZoneActionBrowserCacheTtlID) IsKnown() bool {
	switch r {
	case ZoneActionBrowserCacheTtlIDBrowserCacheTtl:
		return true
	}
	return false
}

type ZoneActionBrowserCheck struct {
	// Inspect the visitor's browser for headers commonly associated with spammers and
	// certain bots.
	ID ZoneActionBrowserCheckID `json:"id"`
	// The status of Browser Integrity Check.
	Value ZoneActionBrowserCheckValue `json:"value"`
	JSON  zoneActionBrowserCheckJSON  `json:"-"`
}

// zoneActionBrowserCheckJSON contains the JSON metadata for the struct
// [ZoneActionBrowserCheck]
type zoneActionBrowserCheckJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionBrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionBrowserCheckJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionBrowserCheck) implementsZoneAction() {}

// Inspect the visitor's browser for headers commonly associated with spammers and
// certain bots.
type ZoneActionBrowserCheckID string

const (
	ZoneActionBrowserCheckIDBrowserCheck ZoneActionBrowserCheckID = "browser_check"
)

func (r ZoneActionBrowserCheckID) IsKnown() bool {
	switch r {
	case ZoneActionBrowserCheckIDBrowserCheck:
		return true
	}
	return false
}

// The status of Browser Integrity Check.
type ZoneActionBrowserCheckValue string

const (
	ZoneActionBrowserCheckValueOn  ZoneActionBrowserCheckValue = "on"
	ZoneActionBrowserCheckValueOff ZoneActionBrowserCheckValue = "off"
)

func (r ZoneActionBrowserCheckValue) IsKnown() bool {
	switch r {
	case ZoneActionBrowserCheckValueOn, ZoneActionBrowserCheckValueOff:
		return true
	}
	return false
}

type ZoneActionBypassCacheOnCookie struct {
	// Bypass cache and fetch resources from the origin server if a regular expression
	// matches against a cookie name present in the request.
	ID ZoneActionBypassCacheOnCookieID `json:"id"`
	// The regular expression to use for matching cookie names in the request. Refer to
	// [Bypass Cache on Cookie setting](https://developers.cloudflare.com/rules/page-rules/reference/additional-reference/#bypass-cache-on-cookie-setting)
	// to learn about limited regular expression support.
	Value string                            `json:"value"`
	JSON  zoneActionBypassCacheOnCookieJSON `json:"-"`
}

// zoneActionBypassCacheOnCookieJSON contains the JSON metadata for the struct
// [ZoneActionBypassCacheOnCookie]
type zoneActionBypassCacheOnCookieJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionBypassCacheOnCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionBypassCacheOnCookieJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionBypassCacheOnCookie) implementsZoneAction() {}

// Bypass cache and fetch resources from the origin server if a regular expression
// matches against a cookie name present in the request.
type ZoneActionBypassCacheOnCookieID string

const (
	ZoneActionBypassCacheOnCookieIDBypassCacheOnCookie ZoneActionBypassCacheOnCookieID = "bypass_cache_on_cookie"
)

func (r ZoneActionBypassCacheOnCookieID) IsKnown() bool {
	switch r {
	case ZoneActionBypassCacheOnCookieIDBypassCacheOnCookie:
		return true
	}
	return false
}

type ZoneActionCacheByDeviceType struct {
	// Separate cached content based on the visitor's device type.
	ID ZoneActionCacheByDeviceTypeID `json:"id"`
	// The status of Cache By Device Type.
	Value ZoneActionCacheByDeviceTypeValue `json:"value"`
	JSON  zoneActionCacheByDeviceTypeJSON  `json:"-"`
}

// zoneActionCacheByDeviceTypeJSON contains the JSON metadata for the struct
// [ZoneActionCacheByDeviceType]
type zoneActionCacheByDeviceTypeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionCacheByDeviceType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionCacheByDeviceTypeJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionCacheByDeviceType) implementsZoneAction() {}

// Separate cached content based on the visitor's device type.
type ZoneActionCacheByDeviceTypeID string

const (
	ZoneActionCacheByDeviceTypeIDCacheByDeviceType ZoneActionCacheByDeviceTypeID = "cache_by_device_type"
)

func (r ZoneActionCacheByDeviceTypeID) IsKnown() bool {
	switch r {
	case ZoneActionCacheByDeviceTypeIDCacheByDeviceType:
		return true
	}
	return false
}

// The status of Cache By Device Type.
type ZoneActionCacheByDeviceTypeValue string

const (
	ZoneActionCacheByDeviceTypeValueOn  ZoneActionCacheByDeviceTypeValue = "on"
	ZoneActionCacheByDeviceTypeValueOff ZoneActionCacheByDeviceTypeValue = "off"
)

func (r ZoneActionCacheByDeviceTypeValue) IsKnown() bool {
	switch r {
	case ZoneActionCacheByDeviceTypeValueOn, ZoneActionCacheByDeviceTypeValueOff:
		return true
	}
	return false
}

type ZoneActionCacheDeceptionArmor struct {
	// Protect from web cache deception attacks while still allowing static assets to
	// be cached. This setting verifies that the URL's extension matches the returned
	// `Content-Type`.
	ID ZoneActionCacheDeceptionArmorID `json:"id"`
	// The status of Cache Deception Armor.
	Value ZoneActionCacheDeceptionArmorValue `json:"value"`
	JSON  zoneActionCacheDeceptionArmorJSON  `json:"-"`
}

// zoneActionCacheDeceptionArmorJSON contains the JSON metadata for the struct
// [ZoneActionCacheDeceptionArmor]
type zoneActionCacheDeceptionArmorJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionCacheDeceptionArmor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionCacheDeceptionArmorJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionCacheDeceptionArmor) implementsZoneAction() {}

// Protect from web cache deception attacks while still allowing static assets to
// be cached. This setting verifies that the URL's extension matches the returned
// `Content-Type`.
type ZoneActionCacheDeceptionArmorID string

const (
	ZoneActionCacheDeceptionArmorIDCacheDeceptionArmor ZoneActionCacheDeceptionArmorID = "cache_deception_armor"
)

func (r ZoneActionCacheDeceptionArmorID) IsKnown() bool {
	switch r {
	case ZoneActionCacheDeceptionArmorIDCacheDeceptionArmor:
		return true
	}
	return false
}

// The status of Cache Deception Armor.
type ZoneActionCacheDeceptionArmorValue string

const (
	ZoneActionCacheDeceptionArmorValueOn  ZoneActionCacheDeceptionArmorValue = "on"
	ZoneActionCacheDeceptionArmorValueOff ZoneActionCacheDeceptionArmorValue = "off"
)

func (r ZoneActionCacheDeceptionArmorValue) IsKnown() bool {
	switch r {
	case ZoneActionCacheDeceptionArmorValueOn, ZoneActionCacheDeceptionArmorValueOff:
		return true
	}
	return false
}

type ZoneActionCacheKeyFields struct {
	// Control specifically what variables to include when deciding which resources to
	// cache. This allows customers to determine what to cache based on something other
	// than just the URL.
	ID    ZoneActionCacheKeyFieldsID    `json:"id"`
	Value ZoneActionCacheKeyFieldsValue `json:"value"`
	JSON  zoneActionCacheKeyFieldsJSON  `json:"-"`
}

// zoneActionCacheKeyFieldsJSON contains the JSON metadata for the struct
// [ZoneActionCacheKeyFields]
type zoneActionCacheKeyFieldsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionCacheKeyFields) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionCacheKeyFieldsJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionCacheKeyFields) implementsZoneAction() {}

// Control specifically what variables to include when deciding which resources to
// cache. This allows customers to determine what to cache based on something other
// than just the URL.
type ZoneActionCacheKeyFieldsID string

const (
	ZoneActionCacheKeyFieldsIDCacheKeyFields ZoneActionCacheKeyFieldsID = "cache_key_fields"
)

func (r ZoneActionCacheKeyFieldsID) IsKnown() bool {
	switch r {
	case ZoneActionCacheKeyFieldsIDCacheKeyFields:
		return true
	}
	return false
}

type ZoneActionCacheKeyFieldsValue struct {
	// Controls which cookies appear in the Cache Key.
	Cookie ZoneActionCacheKeyFieldsValueCookie `json:"cookie"`
	// Controls which headers go into the Cache Key. Exactly one of `include` or
	// `exclude` is expected.
	Header ZoneActionCacheKeyFieldsValueHeader `json:"header"`
	// Determines which host header to include in the Cache Key.
	Host ZoneActionCacheKeyFieldsValueHost `json:"host"`
	// Controls which URL query string parameters go into the Cache Key. Exactly one of
	// `include` or `exclude` is expected.
	QueryString ZoneActionCacheKeyFieldsValueQueryString `json:"query_string"`
	// Feature fields to add features about the end-user (client) into the Cache Key.
	User ZoneActionCacheKeyFieldsValueUser `json:"user"`
	JSON zoneActionCacheKeyFieldsValueJSON `json:"-"`
}

// zoneActionCacheKeyFieldsValueJSON contains the JSON metadata for the struct
// [ZoneActionCacheKeyFieldsValue]
type zoneActionCacheKeyFieldsValueJSON struct {
	Cookie      apijson.Field
	Header      apijson.Field
	Host        apijson.Field
	QueryString apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionCacheKeyFieldsValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionCacheKeyFieldsValueJSON) RawJSON() string {
	return r.raw
}

// Controls which cookies appear in the Cache Key.
type ZoneActionCacheKeyFieldsValueCookie struct {
	// A list of cookies to check for the presence of, without including their actual
	// values.
	CheckPresence []string `json:"check_presence"`
	// A list of cookies to include.
	Include []string                                `json:"include"`
	JSON    zoneActionCacheKeyFieldsValueCookieJSON `json:"-"`
}

// zoneActionCacheKeyFieldsValueCookieJSON contains the JSON metadata for the
// struct [ZoneActionCacheKeyFieldsValueCookie]
type zoneActionCacheKeyFieldsValueCookieJSON struct {
	CheckPresence apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneActionCacheKeyFieldsValueCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionCacheKeyFieldsValueCookieJSON) RawJSON() string {
	return r.raw
}

// Controls which headers go into the Cache Key. Exactly one of `include` or
// `exclude` is expected.
type ZoneActionCacheKeyFieldsValueHeader struct {
	// A list of headers to check for the presence of, without including their actual
	// values.
	CheckPresence []string `json:"check_presence"`
	// A list of headers to ignore.
	Exclude []string `json:"exclude"`
	// A list of headers to include.
	Include []string                                `json:"include"`
	JSON    zoneActionCacheKeyFieldsValueHeaderJSON `json:"-"`
}

// zoneActionCacheKeyFieldsValueHeaderJSON contains the JSON metadata for the
// struct [ZoneActionCacheKeyFieldsValueHeader]
type zoneActionCacheKeyFieldsValueHeaderJSON struct {
	CheckPresence apijson.Field
	Exclude       apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneActionCacheKeyFieldsValueHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionCacheKeyFieldsValueHeaderJSON) RawJSON() string {
	return r.raw
}

// Determines which host header to include in the Cache Key.
type ZoneActionCacheKeyFieldsValueHost struct {
	// Whether to include the Host header in the HTTP request sent to the origin.
	Resolved bool                                  `json:"resolved"`
	JSON     zoneActionCacheKeyFieldsValueHostJSON `json:"-"`
}

// zoneActionCacheKeyFieldsValueHostJSON contains the JSON metadata for the struct
// [ZoneActionCacheKeyFieldsValueHost]
type zoneActionCacheKeyFieldsValueHostJSON struct {
	Resolved    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionCacheKeyFieldsValueHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionCacheKeyFieldsValueHostJSON) RawJSON() string {
	return r.raw
}

// Controls which URL query string parameters go into the Cache Key. Exactly one of
// `include` or `exclude` is expected.
type ZoneActionCacheKeyFieldsValueQueryString struct {
	// Ignore all query string parameters.
	Exclude ZoneActionCacheKeyFieldsValueQueryStringExcludeUnion `json:"exclude"`
	// Include all query string parameters.
	Include ZoneActionCacheKeyFieldsValueQueryStringIncludeUnion `json:"include"`
	JSON    zoneActionCacheKeyFieldsValueQueryStringJSON         `json:"-"`
}

// zoneActionCacheKeyFieldsValueQueryStringJSON contains the JSON metadata for the
// struct [ZoneActionCacheKeyFieldsValueQueryString]
type zoneActionCacheKeyFieldsValueQueryStringJSON struct {
	Exclude     apijson.Field
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionCacheKeyFieldsValueQueryString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionCacheKeyFieldsValueQueryStringJSON) RawJSON() string {
	return r.raw
}

// Ignore all query string parameters.
//
// Union satisfied by [ZoneActionCacheKeyFieldsValueQueryStringExcludeString] or
// [ZoneActionCacheKeyFieldsValueQueryStringExcludeArray].
type ZoneActionCacheKeyFieldsValueQueryStringExcludeUnion interface {
	implementsZoneActionCacheKeyFieldsValueQueryStringExcludeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneActionCacheKeyFieldsValueQueryStringExcludeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ZoneActionCacheKeyFieldsValueQueryStringExcludeString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneActionCacheKeyFieldsValueQueryStringExcludeArray{}),
		},
	)
}

// Ignore all query string parameters.
type ZoneActionCacheKeyFieldsValueQueryStringExcludeString string

const (
	ZoneActionCacheKeyFieldsValueQueryStringExcludeStringStar ZoneActionCacheKeyFieldsValueQueryStringExcludeString = "*"
)

func (r ZoneActionCacheKeyFieldsValueQueryStringExcludeString) IsKnown() bool {
	switch r {
	case ZoneActionCacheKeyFieldsValueQueryStringExcludeStringStar:
		return true
	}
	return false
}

func (r ZoneActionCacheKeyFieldsValueQueryStringExcludeString) implementsZoneActionCacheKeyFieldsValueQueryStringExcludeUnion() {
}

func (r ZoneActionCacheKeyFieldsValueQueryStringExcludeString) implementsZoneActionCacheKeyFieldsValueQueryStringExcludeUnionParam() {
}

type ZoneActionCacheKeyFieldsValueQueryStringExcludeArray []string

func (r ZoneActionCacheKeyFieldsValueQueryStringExcludeArray) implementsZoneActionCacheKeyFieldsValueQueryStringExcludeUnion() {
}

// Include all query string parameters.
//
// Union satisfied by [ZoneActionCacheKeyFieldsValueQueryStringIncludeString] or
// [ZoneActionCacheKeyFieldsValueQueryStringIncludeArray].
type ZoneActionCacheKeyFieldsValueQueryStringIncludeUnion interface {
	implementsZoneActionCacheKeyFieldsValueQueryStringIncludeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneActionCacheKeyFieldsValueQueryStringIncludeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ZoneActionCacheKeyFieldsValueQueryStringIncludeString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneActionCacheKeyFieldsValueQueryStringIncludeArray{}),
		},
	)
}

// Include all query string parameters.
type ZoneActionCacheKeyFieldsValueQueryStringIncludeString string

const (
	ZoneActionCacheKeyFieldsValueQueryStringIncludeStringStar ZoneActionCacheKeyFieldsValueQueryStringIncludeString = "*"
)

func (r ZoneActionCacheKeyFieldsValueQueryStringIncludeString) IsKnown() bool {
	switch r {
	case ZoneActionCacheKeyFieldsValueQueryStringIncludeStringStar:
		return true
	}
	return false
}

func (r ZoneActionCacheKeyFieldsValueQueryStringIncludeString) implementsZoneActionCacheKeyFieldsValueQueryStringIncludeUnion() {
}

func (r ZoneActionCacheKeyFieldsValueQueryStringIncludeString) implementsZoneActionCacheKeyFieldsValueQueryStringIncludeUnionParam() {
}

type ZoneActionCacheKeyFieldsValueQueryStringIncludeArray []string

func (r ZoneActionCacheKeyFieldsValueQueryStringIncludeArray) implementsZoneActionCacheKeyFieldsValueQueryStringIncludeUnion() {
}

// Feature fields to add features about the end-user (client) into the Cache Key.
type ZoneActionCacheKeyFieldsValueUser struct {
	// Classifies a request as `mobile`, `desktop`, or `tablet` based on the User
	// Agent.
	DeviceType bool `json:"device_type"`
	// Includes the client's country, derived from the IP address.
	Geo bool `json:"geo"`
	// Includes the first language code contained in the `Accept-Language` header sent
	// by the client.
	Lang bool                                  `json:"lang"`
	JSON zoneActionCacheKeyFieldsValueUserJSON `json:"-"`
}

// zoneActionCacheKeyFieldsValueUserJSON contains the JSON metadata for the struct
// [ZoneActionCacheKeyFieldsValueUser]
type zoneActionCacheKeyFieldsValueUserJSON struct {
	DeviceType  apijson.Field
	Geo         apijson.Field
	Lang        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionCacheKeyFieldsValueUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionCacheKeyFieldsValueUserJSON) RawJSON() string {
	return r.raw
}

type ZoneActionCacheLevel struct {
	// Apply custom caching based on the option selected.
	ID ZoneActionCacheLevelID `json:"id"`
	//   - `bypass`: Cloudflare does not cache.
	//   - `basic`: Delivers resources from cache when there is no query string.
	//   - `simplified`: Delivers the same resource to everyone independent of the query
	//     string.
	//   - `aggressive`: Caches all static content that has a query string.
	//   - `cache_everything`: Treats all content as static and caches all file types
	//     beyond the
	//     [Cloudflare default cached content](https://developers.cloudflare.com/cache/concepts/default-cache-behavior/#default-cached-file-extensions).
	Value ZoneActionCacheLevelValue `json:"value"`
	JSON  zoneActionCacheLevelJSON  `json:"-"`
}

// zoneActionCacheLevelJSON contains the JSON metadata for the struct
// [ZoneActionCacheLevel]
type zoneActionCacheLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionCacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionCacheLevelJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionCacheLevel) implementsZoneAction() {}

// Apply custom caching based on the option selected.
type ZoneActionCacheLevelID string

const (
	ZoneActionCacheLevelIDCacheLevel ZoneActionCacheLevelID = "cache_level"
)

func (r ZoneActionCacheLevelID) IsKnown() bool {
	switch r {
	case ZoneActionCacheLevelIDCacheLevel:
		return true
	}
	return false
}

//   - `bypass`: Cloudflare does not cache.
//   - `basic`: Delivers resources from cache when there is no query string.
//   - `simplified`: Delivers the same resource to everyone independent of the query
//     string.
//   - `aggressive`: Caches all static content that has a query string.
//   - `cache_everything`: Treats all content as static and caches all file types
//     beyond the
//     [Cloudflare default cached content](https://developers.cloudflare.com/cache/concepts/default-cache-behavior/#default-cached-file-extensions).
type ZoneActionCacheLevelValue string

const (
	ZoneActionCacheLevelValueBypass          ZoneActionCacheLevelValue = "bypass"
	ZoneActionCacheLevelValueBasic           ZoneActionCacheLevelValue = "basic"
	ZoneActionCacheLevelValueSimplified      ZoneActionCacheLevelValue = "simplified"
	ZoneActionCacheLevelValueAggressive      ZoneActionCacheLevelValue = "aggressive"
	ZoneActionCacheLevelValueCacheEverything ZoneActionCacheLevelValue = "cache_everything"
)

func (r ZoneActionCacheLevelValue) IsKnown() bool {
	switch r {
	case ZoneActionCacheLevelValueBypass, ZoneActionCacheLevelValueBasic, ZoneActionCacheLevelValueSimplified, ZoneActionCacheLevelValueAggressive, ZoneActionCacheLevelValueCacheEverything:
		return true
	}
	return false
}

type ZoneActionCacheOnCookie struct {
	// Apply the Cache Everything option (Cache Level setting) based on a regular
	// expression match against a cookie name.
	ID ZoneActionCacheOnCookieID `json:"id"`
	// The regular expression to use for matching cookie names in the request.
	Value string                      `json:"value"`
	JSON  zoneActionCacheOnCookieJSON `json:"-"`
}

// zoneActionCacheOnCookieJSON contains the JSON metadata for the struct
// [ZoneActionCacheOnCookie]
type zoneActionCacheOnCookieJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionCacheOnCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionCacheOnCookieJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionCacheOnCookie) implementsZoneAction() {}

// Apply the Cache Everything option (Cache Level setting) based on a regular
// expression match against a cookie name.
type ZoneActionCacheOnCookieID string

const (
	ZoneActionCacheOnCookieIDCacheOnCookie ZoneActionCacheOnCookieID = "cache_on_cookie"
)

func (r ZoneActionCacheOnCookieID) IsKnown() bool {
	switch r {
	case ZoneActionCacheOnCookieIDCacheOnCookie:
		return true
	}
	return false
}

type ZoneActionCacheTtlByStatus struct {
	// Enterprise customers can set cache time-to-live (TTL) based on the response
	// status from the origin web server. Cache TTL refers to the duration of a
	// resource in the Cloudflare network before being marked as stale or discarded
	// from cache. Status codes are returned by a resource's origin. Setting cache TTL
	// based on response status overrides the default cache behavior (standard caching)
	// for static files and overrides cache instructions sent by the origin web server.
	// To cache non-static assets, set a Cache Level of Cache Everything using a Page
	// Rule. Setting no-store Cache-Control or a low TTL (using `max-age`/`s-maxage`)
	// increases requests to origin web servers and decreases performance.
	ID ZoneActionCacheTtlByStatusID `json:"id"`
	// A JSON object containing status codes and their corresponding TTLs. Each
	// key-value pair in the cache TTL by status cache rule has the following syntax
	//
	//   - `status_code`: An integer value such as 200 or 500. status_code matches the
	//     exact status code from the origin web server. Valid status codes are between
	//     100-999.
	//   - `status_code_range`: Integer values for from and to. status_code_range matches
	//     any status code from the origin web server within the specified range.
	//   - `value`: An integer value that defines the duration an asset is valid in
	//     seconds or one of the following strings: no-store (equivalent to -1), no-cache
	//     (equivalent to 0).
	Value map[string]ZoneActionCacheTtlByStatusValueUnion `json:"value"`
	JSON  zoneActionCacheTtlByStatusJSON                  `json:"-"`
}

// zoneActionCacheTtlByStatusJSON contains the JSON metadata for the struct
// [ZoneActionCacheTtlByStatus]
type zoneActionCacheTtlByStatusJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionCacheTtlByStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionCacheTtlByStatusJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionCacheTtlByStatus) implementsZoneAction() {}

// Enterprise customers can set cache time-to-live (TTL) based on the response
// status from the origin web server. Cache TTL refers to the duration of a
// resource in the Cloudflare network before being marked as stale or discarded
// from cache. Status codes are returned by a resource's origin. Setting cache TTL
// based on response status overrides the default cache behavior (standard caching)
// for static files and overrides cache instructions sent by the origin web server.
// To cache non-static assets, set a Cache Level of Cache Everything using a Page
// Rule. Setting no-store Cache-Control or a low TTL (using `max-age`/`s-maxage`)
// increases requests to origin web servers and decreases performance.
type ZoneActionCacheTtlByStatusID string

const (
	ZoneActionCacheTtlByStatusIDCacheTtlByStatus ZoneActionCacheTtlByStatusID = "cache_ttl_by_status"
)

func (r ZoneActionCacheTtlByStatusID) IsKnown() bool {
	switch r {
	case ZoneActionCacheTtlByStatusIDCacheTtlByStatus:
		return true
	}
	return false
}

// `no-store` (equivalent to -1), `no-cache` (equivalent to 0)
//
// Union satisfied by [ZoneActionCacheTtlByStatusValueString] or [shared.UnionInt].
type ZoneActionCacheTtlByStatusValueUnion interface {
	ImplementsZoneActionCacheTtlByStatusValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneActionCacheTtlByStatusValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ZoneActionCacheTtlByStatusValueString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// `no-store` (equivalent to -1), `no-cache` (equivalent to 0)
type ZoneActionCacheTtlByStatusValueString string

const (
	ZoneActionCacheTtlByStatusValueStringNoCache ZoneActionCacheTtlByStatusValueString = "no-cache"
	ZoneActionCacheTtlByStatusValueStringNoStore ZoneActionCacheTtlByStatusValueString = "no-store"
)

func (r ZoneActionCacheTtlByStatusValueString) IsKnown() bool {
	switch r {
	case ZoneActionCacheTtlByStatusValueStringNoCache, ZoneActionCacheTtlByStatusValueStringNoStore:
		return true
	}
	return false
}

func (r ZoneActionCacheTtlByStatusValueString) ImplementsZoneActionCacheTtlByStatusValueUnion() {}

func (r ZoneActionCacheTtlByStatusValueString) ImplementsZoneActionCacheTtlByStatusValueUnionParam() {
}

type ZoneActionDisableApps struct {
	// Turn off all active
	// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
	// (deprecated).
	ID   ZoneActionDisableAppsID   `json:"id"`
	JSON zoneActionDisableAppsJSON `json:"-"`
}

// zoneActionDisableAppsJSON contains the JSON metadata for the struct
// [ZoneActionDisableApps]
type zoneActionDisableAppsJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionDisableApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionDisableAppsJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionDisableApps) implementsZoneAction() {}

// Turn off all active
// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
// (deprecated).
type ZoneActionDisableAppsID string

const (
	ZoneActionDisableAppsIDDisableApps ZoneActionDisableAppsID = "disable_apps"
)

func (r ZoneActionDisableAppsID) IsKnown() bool {
	switch r {
	case ZoneActionDisableAppsIDDisableApps:
		return true
	}
	return false
}

type ZoneActionDisablePerformance struct {
	// Turn off
	// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
	// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
	// and [Polish](https://developers.cloudflare.com/images/polish/).
	ID   ZoneActionDisablePerformanceID   `json:"id"`
	JSON zoneActionDisablePerformanceJSON `json:"-"`
}

// zoneActionDisablePerformanceJSON contains the JSON metadata for the struct
// [ZoneActionDisablePerformance]
type zoneActionDisablePerformanceJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionDisablePerformance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionDisablePerformanceJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionDisablePerformance) implementsZoneAction() {}

// Turn off
// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
// and [Polish](https://developers.cloudflare.com/images/polish/).
type ZoneActionDisablePerformanceID string

const (
	ZoneActionDisablePerformanceIDDisablePerformance ZoneActionDisablePerformanceID = "disable_performance"
)

func (r ZoneActionDisablePerformanceID) IsKnown() bool {
	switch r {
	case ZoneActionDisablePerformanceIDDisablePerformance:
		return true
	}
	return false
}

type ZoneActionDisableSecurity struct {
	// Turn off
	// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
	// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
	// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
	// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
	// and
	// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
	ID   ZoneActionDisableSecurityID   `json:"id"`
	JSON zoneActionDisableSecurityJSON `json:"-"`
}

// zoneActionDisableSecurityJSON contains the JSON metadata for the struct
// [ZoneActionDisableSecurity]
type zoneActionDisableSecurityJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionDisableSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionDisableSecurityJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionDisableSecurity) implementsZoneAction() {}

// Turn off
// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
// and
// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
type ZoneActionDisableSecurityID string

const (
	ZoneActionDisableSecurityIDDisableSecurity ZoneActionDisableSecurityID = "disable_security"
)

func (r ZoneActionDisableSecurityID) IsKnown() bool {
	switch r {
	case ZoneActionDisableSecurityIDDisableSecurity:
		return true
	}
	return false
}

type ZoneActionDisableZaraz struct {
	// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
	ID   ZoneActionDisableZarazID   `json:"id"`
	JSON zoneActionDisableZarazJSON `json:"-"`
}

// zoneActionDisableZarazJSON contains the JSON metadata for the struct
// [ZoneActionDisableZaraz]
type zoneActionDisableZarazJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionDisableZaraz) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionDisableZarazJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionDisableZaraz) implementsZoneAction() {}

// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
type ZoneActionDisableZarazID string

const (
	ZoneActionDisableZarazIDDisableZaraz ZoneActionDisableZarazID = "disable_zaraz"
)

func (r ZoneActionDisableZarazID) IsKnown() bool {
	switch r {
	case ZoneActionDisableZarazIDDisableZaraz:
		return true
	}
	return false
}

type ZoneActionEdgeCacheTtl struct {
	// Specify how long to cache a resource in the Cloudflare global network. _Edge
	// Cache TTL_ is not visible in response headers.
	ID    ZoneActionEdgeCacheTtlID   `json:"id"`
	Value int64                      `json:"value"`
	JSON  zoneActionEdgeCacheTtlJSON `json:"-"`
}

// zoneActionEdgeCacheTtlJSON contains the JSON metadata for the struct
// [ZoneActionEdgeCacheTtl]
type zoneActionEdgeCacheTtlJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionEdgeCacheTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionEdgeCacheTtlJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionEdgeCacheTtl) implementsZoneAction() {}

// Specify how long to cache a resource in the Cloudflare global network. _Edge
// Cache TTL_ is not visible in response headers.
type ZoneActionEdgeCacheTtlID string

const (
	ZoneActionEdgeCacheTtlIDEdgeCacheTtl ZoneActionEdgeCacheTtlID = "edge_cache_ttl"
)

func (r ZoneActionEdgeCacheTtlID) IsKnown() bool {
	switch r {
	case ZoneActionEdgeCacheTtlIDEdgeCacheTtl:
		return true
	}
	return false
}

type ZoneActionEmailObfuscation struct {
	// Turn on or off **Email Obfuscation**.
	ID ZoneActionEmailObfuscationID `json:"id"`
	// The status of Email Obfuscation.
	Value ZoneActionEmailObfuscationValue `json:"value"`
	JSON  zoneActionEmailObfuscationJSON  `json:"-"`
}

// zoneActionEmailObfuscationJSON contains the JSON metadata for the struct
// [ZoneActionEmailObfuscation]
type zoneActionEmailObfuscationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionEmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionEmailObfuscationJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionEmailObfuscation) implementsZoneAction() {}

// Turn on or off **Email Obfuscation**.
type ZoneActionEmailObfuscationID string

const (
	ZoneActionEmailObfuscationIDEmailObfuscation ZoneActionEmailObfuscationID = "email_obfuscation"
)

func (r ZoneActionEmailObfuscationID) IsKnown() bool {
	switch r {
	case ZoneActionEmailObfuscationIDEmailObfuscation:
		return true
	}
	return false
}

// The status of Email Obfuscation.
type ZoneActionEmailObfuscationValue string

const (
	ZoneActionEmailObfuscationValueOn  ZoneActionEmailObfuscationValue = "on"
	ZoneActionEmailObfuscationValueOff ZoneActionEmailObfuscationValue = "off"
)

func (r ZoneActionEmailObfuscationValue) IsKnown() bool {
	switch r {
	case ZoneActionEmailObfuscationValueOn, ZoneActionEmailObfuscationValueOff:
		return true
	}
	return false
}

type ZoneActionExplicitCacheControl struct {
	// Origin Cache Control is enabled by default for Free, Pro, and Business domains
	// and disabled by default for Enterprise domains.
	ID ZoneActionExplicitCacheControlID `json:"id"`
	// The status of Origin Cache Control.
	Value ZoneActionExplicitCacheControlValue `json:"value"`
	JSON  zoneActionExplicitCacheControlJSON  `json:"-"`
}

// zoneActionExplicitCacheControlJSON contains the JSON metadata for the struct
// [ZoneActionExplicitCacheControl]
type zoneActionExplicitCacheControlJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionExplicitCacheControl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionExplicitCacheControlJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionExplicitCacheControl) implementsZoneAction() {}

// Origin Cache Control is enabled by default for Free, Pro, and Business domains
// and disabled by default for Enterprise domains.
type ZoneActionExplicitCacheControlID string

const (
	ZoneActionExplicitCacheControlIDExplicitCacheControl ZoneActionExplicitCacheControlID = "explicit_cache_control"
)

func (r ZoneActionExplicitCacheControlID) IsKnown() bool {
	switch r {
	case ZoneActionExplicitCacheControlIDExplicitCacheControl:
		return true
	}
	return false
}

// The status of Origin Cache Control.
type ZoneActionExplicitCacheControlValue string

const (
	ZoneActionExplicitCacheControlValueOn  ZoneActionExplicitCacheControlValue = "on"
	ZoneActionExplicitCacheControlValueOff ZoneActionExplicitCacheControlValue = "off"
)

func (r ZoneActionExplicitCacheControlValue) IsKnown() bool {
	switch r {
	case ZoneActionExplicitCacheControlValueOn, ZoneActionExplicitCacheControlValueOff:
		return true
	}
	return false
}

type ZoneActionForwardingURL struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    ZoneActionForwardingURLID    `json:"id"`
	Value ZoneActionForwardingURLValue `json:"value"`
	JSON  zoneActionForwardingURLJSON  `json:"-"`
}

// zoneActionForwardingURLJSON contains the JSON metadata for the struct
// [ZoneActionForwardingURL]
type zoneActionForwardingURLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionForwardingURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionForwardingURLJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionForwardingURL) implementsZoneAction() {}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type ZoneActionForwardingURLID string

const (
	ZoneActionForwardingURLIDForwardingURL ZoneActionForwardingURLID = "forwarding_url"
)

func (r ZoneActionForwardingURLID) IsKnown() bool {
	switch r {
	case ZoneActionForwardingURLIDForwardingURL:
		return true
	}
	return false
}

type ZoneActionForwardingURLValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode ZoneActionForwardingURLValueStatusCode `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL  string                           `json:"url"`
	JSON zoneActionForwardingURLValueJSON `json:"-"`
}

// zoneActionForwardingURLValueJSON contains the JSON metadata for the struct
// [ZoneActionForwardingURLValue]
type zoneActionForwardingURLValueJSON struct {
	StatusCode  apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionForwardingURLValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionForwardingURLValueJSON) RawJSON() string {
	return r.raw
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type ZoneActionForwardingURLValueStatusCode int64

const (
	ZoneActionForwardingURLValueStatusCode301 ZoneActionForwardingURLValueStatusCode = 301
	ZoneActionForwardingURLValueStatusCode302 ZoneActionForwardingURLValueStatusCode = 302
)

func (r ZoneActionForwardingURLValueStatusCode) IsKnown() bool {
	switch r {
	case ZoneActionForwardingURLValueStatusCode301, ZoneActionForwardingURLValueStatusCode302:
		return true
	}
	return false
}

type ZoneActionHostHeaderOverride struct {
	// Apply a specific host header.
	ID ZoneActionHostHeaderOverrideID `json:"id"`
	// The hostname to use in the `Host` header
	Value string                           `json:"value"`
	JSON  zoneActionHostHeaderOverrideJSON `json:"-"`
}

// zoneActionHostHeaderOverrideJSON contains the JSON metadata for the struct
// [ZoneActionHostHeaderOverride]
type zoneActionHostHeaderOverrideJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionHostHeaderOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionHostHeaderOverrideJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionHostHeaderOverride) implementsZoneAction() {}

// Apply a specific host header.
type ZoneActionHostHeaderOverrideID string

const (
	ZoneActionHostHeaderOverrideIDHostHeaderOverride ZoneActionHostHeaderOverrideID = "host_header_override"
)

func (r ZoneActionHostHeaderOverrideID) IsKnown() bool {
	switch r {
	case ZoneActionHostHeaderOverrideIDHostHeaderOverride:
		return true
	}
	return false
}

type ZoneActionIPGeolocation struct {
	// Cloudflare adds a CF-IPCountry HTTP header containing the country code that
	// corresponds to the visitor.
	ID ZoneActionIPGeolocationID `json:"id"`
	// The status of adding the IP Geolocation Header.
	Value ZoneActionIPGeolocationValue `json:"value"`
	JSON  zoneActionIPGeolocationJSON  `json:"-"`
}

// zoneActionIPGeolocationJSON contains the JSON metadata for the struct
// [ZoneActionIPGeolocation]
type zoneActionIPGeolocationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionIPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionIPGeolocationJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionIPGeolocation) implementsZoneAction() {}

// Cloudflare adds a CF-IPCountry HTTP header containing the country code that
// corresponds to the visitor.
type ZoneActionIPGeolocationID string

const (
	ZoneActionIPGeolocationIDIPGeolocation ZoneActionIPGeolocationID = "ip_geolocation"
)

func (r ZoneActionIPGeolocationID) IsKnown() bool {
	switch r {
	case ZoneActionIPGeolocationIDIPGeolocation:
		return true
	}
	return false
}

// The status of adding the IP Geolocation Header.
type ZoneActionIPGeolocationValue string

const (
	ZoneActionIPGeolocationValueOn  ZoneActionIPGeolocationValue = "on"
	ZoneActionIPGeolocationValueOff ZoneActionIPGeolocationValue = "off"
)

func (r ZoneActionIPGeolocationValue) IsKnown() bool {
	switch r {
	case ZoneActionIPGeolocationValueOn, ZoneActionIPGeolocationValueOff:
		return true
	}
	return false
}

type ZoneActionMirage struct {
	// Cloudflare Mirage reduces bandwidth used by images in mobile browsers. It can
	// accelerate loading of image-heavy websites on very slow mobile connections and
	// HTTP/1.
	ID ZoneActionMirageID `json:"id"`
	// The status of Mirage.
	Value ZoneActionMirageValue `json:"value"`
	JSON  zoneActionMirageJSON  `json:"-"`
}

// zoneActionMirageJSON contains the JSON metadata for the struct
// [ZoneActionMirage]
type zoneActionMirageJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionMirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionMirageJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionMirage) implementsZoneAction() {}

// Cloudflare Mirage reduces bandwidth used by images in mobile browsers. It can
// accelerate loading of image-heavy websites on very slow mobile connections and
// HTTP/1.
type ZoneActionMirageID string

const (
	ZoneActionMirageIDMirage ZoneActionMirageID = "mirage"
)

func (r ZoneActionMirageID) IsKnown() bool {
	switch r {
	case ZoneActionMirageIDMirage:
		return true
	}
	return false
}

// The status of Mirage.
type ZoneActionMirageValue string

const (
	ZoneActionMirageValueOn  ZoneActionMirageValue = "on"
	ZoneActionMirageValueOff ZoneActionMirageValue = "off"
)

func (r ZoneActionMirageValue) IsKnown() bool {
	switch r {
	case ZoneActionMirageValueOn, ZoneActionMirageValueOff:
		return true
	}
	return false
}

type ZoneActionOpportunisticEncryption struct {
	// Opportunistic Encryption allows browsers to access HTTP URIs over an encrypted
	// TLS channel. It's not a substitute for HTTPS, but provides additional security
	// for otherwise vulnerable requests.
	ID ZoneActionOpportunisticEncryptionID `json:"id"`
	// The status of Opportunistic Encryption.
	Value ZoneActionOpportunisticEncryptionValue `json:"value"`
	JSON  zoneActionOpportunisticEncryptionJSON  `json:"-"`
}

// zoneActionOpportunisticEncryptionJSON contains the JSON metadata for the struct
// [ZoneActionOpportunisticEncryption]
type zoneActionOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionOpportunisticEncryptionJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionOpportunisticEncryption) implementsZoneAction() {}

// Opportunistic Encryption allows browsers to access HTTP URIs over an encrypted
// TLS channel. It's not a substitute for HTTPS, but provides additional security
// for otherwise vulnerable requests.
type ZoneActionOpportunisticEncryptionID string

const (
	ZoneActionOpportunisticEncryptionIDOpportunisticEncryption ZoneActionOpportunisticEncryptionID = "opportunistic_encryption"
)

func (r ZoneActionOpportunisticEncryptionID) IsKnown() bool {
	switch r {
	case ZoneActionOpportunisticEncryptionIDOpportunisticEncryption:
		return true
	}
	return false
}

// The status of Opportunistic Encryption.
type ZoneActionOpportunisticEncryptionValue string

const (
	ZoneActionOpportunisticEncryptionValueOn  ZoneActionOpportunisticEncryptionValue = "on"
	ZoneActionOpportunisticEncryptionValueOff ZoneActionOpportunisticEncryptionValue = "off"
)

func (r ZoneActionOpportunisticEncryptionValue) IsKnown() bool {
	switch r {
	case ZoneActionOpportunisticEncryptionValueOn, ZoneActionOpportunisticEncryptionValueOff:
		return true
	}
	return false
}

type ZoneActionOriginErrorPagePassThru struct {
	// Turn on or off Cloudflare error pages generated from issues sent from the origin
	// server. If enabled, this setting triggers error pages issued by the origin.
	ID ZoneActionOriginErrorPagePassThruID `json:"id"`
	// The status of Origin Error Page Passthru.
	Value ZoneActionOriginErrorPagePassThruValue `json:"value"`
	JSON  zoneActionOriginErrorPagePassThruJSON  `json:"-"`
}

// zoneActionOriginErrorPagePassThruJSON contains the JSON metadata for the struct
// [ZoneActionOriginErrorPagePassThru]
type zoneActionOriginErrorPagePassThruJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionOriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionOriginErrorPagePassThruJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionOriginErrorPagePassThru) implementsZoneAction() {}

// Turn on or off Cloudflare error pages generated from issues sent from the origin
// server. If enabled, this setting triggers error pages issued by the origin.
type ZoneActionOriginErrorPagePassThruID string

const (
	ZoneActionOriginErrorPagePassThruIDOriginErrorPagePassThru ZoneActionOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

func (r ZoneActionOriginErrorPagePassThruID) IsKnown() bool {
	switch r {
	case ZoneActionOriginErrorPagePassThruIDOriginErrorPagePassThru:
		return true
	}
	return false
}

// The status of Origin Error Page Passthru.
type ZoneActionOriginErrorPagePassThruValue string

const (
	ZoneActionOriginErrorPagePassThruValueOn  ZoneActionOriginErrorPagePassThruValue = "on"
	ZoneActionOriginErrorPagePassThruValueOff ZoneActionOriginErrorPagePassThruValue = "off"
)

func (r ZoneActionOriginErrorPagePassThruValue) IsKnown() bool {
	switch r {
	case ZoneActionOriginErrorPagePassThruValueOn, ZoneActionOriginErrorPagePassThruValueOff:
		return true
	}
	return false
}

type ZoneActionPolish struct {
	// Apply options from the Polish feature of the Cloudflare Speed app.
	ID ZoneActionPolishID `json:"id"`
	// The level of Polish you want applied to your origin.
	Value ZoneActionPolishValue `json:"value"`
	JSON  zoneActionPolishJSON  `json:"-"`
}

// zoneActionPolishJSON contains the JSON metadata for the struct
// [ZoneActionPolish]
type zoneActionPolishJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionPolish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionPolishJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionPolish) implementsZoneAction() {}

// Apply options from the Polish feature of the Cloudflare Speed app.
type ZoneActionPolishID string

const (
	ZoneActionPolishIDPolish ZoneActionPolishID = "polish"
)

func (r ZoneActionPolishID) IsKnown() bool {
	switch r {
	case ZoneActionPolishIDPolish:
		return true
	}
	return false
}

// The level of Polish you want applied to your origin.
type ZoneActionPolishValue string

const (
	ZoneActionPolishValueOff      ZoneActionPolishValue = "off"
	ZoneActionPolishValueLossless ZoneActionPolishValue = "lossless"
	ZoneActionPolishValueLossy    ZoneActionPolishValue = "lossy"
)

func (r ZoneActionPolishValue) IsKnown() bool {
	switch r {
	case ZoneActionPolishValueOff, ZoneActionPolishValueLossless, ZoneActionPolishValueLossy:
		return true
	}
	return false
}

type ZoneActionResolveOverride struct {
	// Change the origin address to the value specified in this setting.
	ID ZoneActionResolveOverrideID `json:"id"`
	// The origin address you want to override with.
	Value string                        `json:"value"`
	JSON  zoneActionResolveOverrideJSON `json:"-"`
}

// zoneActionResolveOverrideJSON contains the JSON metadata for the struct
// [ZoneActionResolveOverride]
type zoneActionResolveOverrideJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionResolveOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionResolveOverrideJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionResolveOverride) implementsZoneAction() {}

// Change the origin address to the value specified in this setting.
type ZoneActionResolveOverrideID string

const (
	ZoneActionResolveOverrideIDResolveOverride ZoneActionResolveOverrideID = "resolve_override"
)

func (r ZoneActionResolveOverrideID) IsKnown() bool {
	switch r {
	case ZoneActionResolveOverrideIDResolveOverride:
		return true
	}
	return false
}

type ZoneActionRespectStrongEtag struct {
	// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
	// the origin server.
	ID ZoneActionRespectStrongEtagID `json:"id"`
	// The status of Respect Strong ETags
	Value ZoneActionRespectStrongEtagValue `json:"value"`
	JSON  zoneActionRespectStrongEtagJSON  `json:"-"`
}

// zoneActionRespectStrongEtagJSON contains the JSON metadata for the struct
// [ZoneActionRespectStrongEtag]
type zoneActionRespectStrongEtagJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionRespectStrongEtag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionRespectStrongEtagJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionRespectStrongEtag) implementsZoneAction() {}

// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
// the origin server.
type ZoneActionRespectStrongEtagID string

const (
	ZoneActionRespectStrongEtagIDRespectStrongEtag ZoneActionRespectStrongEtagID = "respect_strong_etag"
)

func (r ZoneActionRespectStrongEtagID) IsKnown() bool {
	switch r {
	case ZoneActionRespectStrongEtagIDRespectStrongEtag:
		return true
	}
	return false
}

// The status of Respect Strong ETags
type ZoneActionRespectStrongEtagValue string

const (
	ZoneActionRespectStrongEtagValueOn  ZoneActionRespectStrongEtagValue = "on"
	ZoneActionRespectStrongEtagValueOff ZoneActionRespectStrongEtagValue = "off"
)

func (r ZoneActionRespectStrongEtagValue) IsKnown() bool {
	switch r {
	case ZoneActionRespectStrongEtagValueOn, ZoneActionRespectStrongEtagValueOff:
		return true
	}
	return false
}

type ZoneActionResponseBuffering struct {
	// Turn on or off whether Cloudflare should wait for an entire file from the origin
	// server before forwarding it to the site visitor. By default, Cloudflare sends
	// packets to the client as they arrive from the origin server.
	ID ZoneActionResponseBufferingID `json:"id"`
	// The status of Response Buffering
	Value ZoneActionResponseBufferingValue `json:"value"`
	JSON  zoneActionResponseBufferingJSON  `json:"-"`
}

// zoneActionResponseBufferingJSON contains the JSON metadata for the struct
// [ZoneActionResponseBuffering]
type zoneActionResponseBufferingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionResponseBufferingJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionResponseBuffering) implementsZoneAction() {}

// Turn on or off whether Cloudflare should wait for an entire file from the origin
// server before forwarding it to the site visitor. By default, Cloudflare sends
// packets to the client as they arrive from the origin server.
type ZoneActionResponseBufferingID string

const (
	ZoneActionResponseBufferingIDResponseBuffering ZoneActionResponseBufferingID = "response_buffering"
)

func (r ZoneActionResponseBufferingID) IsKnown() bool {
	switch r {
	case ZoneActionResponseBufferingIDResponseBuffering:
		return true
	}
	return false
}

// The status of Response Buffering
type ZoneActionResponseBufferingValue string

const (
	ZoneActionResponseBufferingValueOn  ZoneActionResponseBufferingValue = "on"
	ZoneActionResponseBufferingValueOff ZoneActionResponseBufferingValue = "off"
)

func (r ZoneActionResponseBufferingValue) IsKnown() bool {
	switch r {
	case ZoneActionResponseBufferingValueOn, ZoneActionResponseBufferingValueOff:
		return true
	}
	return false
}

type ZoneActionRocketLoader struct {
	// Turn on or off Rocket Loader in the Cloudflare Speed app.
	ID ZoneActionRocketLoaderID `json:"id"`
	// The status of Rocket Loader
	Value ZoneActionRocketLoaderValue `json:"value"`
	JSON  zoneActionRocketLoaderJSON  `json:"-"`
}

// zoneActionRocketLoaderJSON contains the JSON metadata for the struct
// [ZoneActionRocketLoader]
type zoneActionRocketLoaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionRocketLoader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionRocketLoaderJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionRocketLoader) implementsZoneAction() {}

// Turn on or off Rocket Loader in the Cloudflare Speed app.
type ZoneActionRocketLoaderID string

const (
	ZoneActionRocketLoaderIDRocketLoader ZoneActionRocketLoaderID = "rocket_loader"
)

func (r ZoneActionRocketLoaderID) IsKnown() bool {
	switch r {
	case ZoneActionRocketLoaderIDRocketLoader:
		return true
	}
	return false
}

// The status of Rocket Loader
type ZoneActionRocketLoaderValue string

const (
	ZoneActionRocketLoaderValueOn  ZoneActionRocketLoaderValue = "on"
	ZoneActionRocketLoaderValueOff ZoneActionRocketLoaderValue = "off"
)

func (r ZoneActionRocketLoaderValue) IsKnown() bool {
	switch r {
	case ZoneActionRocketLoaderValueOn, ZoneActionRocketLoaderValueOff:
		return true
	}
	return false
}

type ZoneActionSecurityLevel struct {
	// Control options for the **Security Level** feature from the **Security** app.
	ID    ZoneActionSecurityLevelID    `json:"id"`
	Value ZoneActionSecurityLevelValue `json:"value"`
	JSON  zoneActionSecurityLevelJSON  `json:"-"`
}

// zoneActionSecurityLevelJSON contains the JSON metadata for the struct
// [ZoneActionSecurityLevel]
type zoneActionSecurityLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionSecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionSecurityLevelJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionSecurityLevel) implementsZoneAction() {}

// Control options for the **Security Level** feature from the **Security** app.
type ZoneActionSecurityLevelID string

const (
	ZoneActionSecurityLevelIDSecurityLevel ZoneActionSecurityLevelID = "security_level"
)

func (r ZoneActionSecurityLevelID) IsKnown() bool {
	switch r {
	case ZoneActionSecurityLevelIDSecurityLevel:
		return true
	}
	return false
}

type ZoneActionSecurityLevelValue string

const (
	ZoneActionSecurityLevelValueOff            ZoneActionSecurityLevelValue = "off"
	ZoneActionSecurityLevelValueEssentiallyOff ZoneActionSecurityLevelValue = "essentially_off"
	ZoneActionSecurityLevelValueLow            ZoneActionSecurityLevelValue = "low"
	ZoneActionSecurityLevelValueMedium         ZoneActionSecurityLevelValue = "medium"
	ZoneActionSecurityLevelValueHigh           ZoneActionSecurityLevelValue = "high"
	ZoneActionSecurityLevelValueUnderAttack    ZoneActionSecurityLevelValue = "under_attack"
)

func (r ZoneActionSecurityLevelValue) IsKnown() bool {
	switch r {
	case ZoneActionSecurityLevelValueOff, ZoneActionSecurityLevelValueEssentiallyOff, ZoneActionSecurityLevelValueLow, ZoneActionSecurityLevelValueMedium, ZoneActionSecurityLevelValueHigh, ZoneActionSecurityLevelValueUnderAttack:
		return true
	}
	return false
}

type ZoneActionSortQueryStringForCache struct {
	// Turn on or off the reordering of query strings. When query strings have the same
	// structure, caching improves.
	ID ZoneActionSortQueryStringForCacheID `json:"id"`
	// The status of Query String Sort
	Value ZoneActionSortQueryStringForCacheValue `json:"value"`
	JSON  zoneActionSortQueryStringForCacheJSON  `json:"-"`
}

// zoneActionSortQueryStringForCacheJSON contains the JSON metadata for the struct
// [ZoneActionSortQueryStringForCache]
type zoneActionSortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionSortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionSortQueryStringForCacheJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionSortQueryStringForCache) implementsZoneAction() {}

// Turn on or off the reordering of query strings. When query strings have the same
// structure, caching improves.
type ZoneActionSortQueryStringForCacheID string

const (
	ZoneActionSortQueryStringForCacheIDSortQueryStringForCache ZoneActionSortQueryStringForCacheID = "sort_query_string_for_cache"
)

func (r ZoneActionSortQueryStringForCacheID) IsKnown() bool {
	switch r {
	case ZoneActionSortQueryStringForCacheIDSortQueryStringForCache:
		return true
	}
	return false
}

// The status of Query String Sort
type ZoneActionSortQueryStringForCacheValue string

const (
	ZoneActionSortQueryStringForCacheValueOn  ZoneActionSortQueryStringForCacheValue = "on"
	ZoneActionSortQueryStringForCacheValueOff ZoneActionSortQueryStringForCacheValue = "off"
)

func (r ZoneActionSortQueryStringForCacheValue) IsKnown() bool {
	switch r {
	case ZoneActionSortQueryStringForCacheValueOn, ZoneActionSortQueryStringForCacheValueOff:
		return true
	}
	return false
}

type ZoneActionSsl struct {
	// Control options for the SSL feature of the Edge Certificates tab in the
	// Cloudflare SSL/TLS app.
	ID ZoneActionSslID `json:"id"`
	// The encryption mode that Cloudflare uses to connect to your origin server.
	Value ZoneActionSslValue `json:"value"`
	JSON  zoneActionSslJSON  `json:"-"`
}

// zoneActionSslJSON contains the JSON metadata for the struct [ZoneActionSsl]
type zoneActionSslJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionSslJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionSsl) implementsZoneAction() {}

// Control options for the SSL feature of the Edge Certificates tab in the
// Cloudflare SSL/TLS app.
type ZoneActionSslID string

const (
	ZoneActionSslIDSsl ZoneActionSslID = "ssl"
)

func (r ZoneActionSslID) IsKnown() bool {
	switch r {
	case ZoneActionSslIDSsl:
		return true
	}
	return false
}

// The encryption mode that Cloudflare uses to connect to your origin server.
type ZoneActionSslValue string

const (
	ZoneActionSslValueOff        ZoneActionSslValue = "off"
	ZoneActionSslValueFlexible   ZoneActionSslValue = "flexible"
	ZoneActionSslValueFull       ZoneActionSslValue = "full"
	ZoneActionSslValueStrict     ZoneActionSslValue = "strict"
	ZoneActionSslValueOriginPull ZoneActionSslValue = "origin_pull"
)

func (r ZoneActionSslValue) IsKnown() bool {
	switch r {
	case ZoneActionSslValueOff, ZoneActionSslValueFlexible, ZoneActionSslValueFull, ZoneActionSslValueStrict, ZoneActionSslValueOriginPull:
		return true
	}
	return false
}

type ZoneActionTrueClientIPHeader struct {
	// Turn on or off the True-Client-IP Header feature of the Cloudflare Network app.
	ID ZoneActionTrueClientIPHeaderID `json:"id"`
	// The status of True Client IP Header.
	Value ZoneActionTrueClientIPHeaderValue `json:"value"`
	JSON  zoneActionTrueClientIPHeaderJSON  `json:"-"`
}

// zoneActionTrueClientIPHeaderJSON contains the JSON metadata for the struct
// [ZoneActionTrueClientIPHeader]
type zoneActionTrueClientIPHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionTrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionTrueClientIPHeaderJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionTrueClientIPHeader) implementsZoneAction() {}

// Turn on or off the True-Client-IP Header feature of the Cloudflare Network app.
type ZoneActionTrueClientIPHeaderID string

const (
	ZoneActionTrueClientIPHeaderIDTrueClientIPHeader ZoneActionTrueClientIPHeaderID = "true_client_ip_header"
)

func (r ZoneActionTrueClientIPHeaderID) IsKnown() bool {
	switch r {
	case ZoneActionTrueClientIPHeaderIDTrueClientIPHeader:
		return true
	}
	return false
}

// The status of True Client IP Header.
type ZoneActionTrueClientIPHeaderValue string

const (
	ZoneActionTrueClientIPHeaderValueOn  ZoneActionTrueClientIPHeaderValue = "on"
	ZoneActionTrueClientIPHeaderValueOff ZoneActionTrueClientIPHeaderValue = "off"
)

func (r ZoneActionTrueClientIPHeaderValue) IsKnown() bool {
	switch r {
	case ZoneActionTrueClientIPHeaderValueOn, ZoneActionTrueClientIPHeaderValueOff:
		return true
	}
	return false
}

type ZoneActionWaf struct {
	// Turn on or off
	// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
	// You cannot enable or disable individual WAF managed rules via Page Rules.
	ID ZoneActionWafID `json:"id"`
	// The status of WAF managed rules (previous version).
	Value ZoneActionWafValue `json:"value"`
	JSON  zoneActionWafJSON  `json:"-"`
}

// zoneActionWafJSON contains the JSON metadata for the struct [ZoneActionWaf]
type zoneActionWafJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActionWaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActionWafJSON) RawJSON() string {
	return r.raw
}

func (r ZoneActionWaf) implementsZoneAction() {}

// Turn on or off
// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
// You cannot enable or disable individual WAF managed rules via Page Rules.
type ZoneActionWafID string

const (
	ZoneActionWafIDWaf ZoneActionWafID = "waf"
)

func (r ZoneActionWafID) IsKnown() bool {
	switch r {
	case ZoneActionWafIDWaf:
		return true
	}
	return false
}

// The status of WAF managed rules (previous version).
type ZoneActionWafValue string

const (
	ZoneActionWafValueOn  ZoneActionWafValue = "on"
	ZoneActionWafValueOff ZoneActionWafValue = "off"
)

func (r ZoneActionWafValue) IsKnown() bool {
	switch r {
	case ZoneActionWafValueOn, ZoneActionWafValueOff:
		return true
	}
	return false
}

// If enabled, any ` http://“ URL is converted to  `https://` through a 301
// redirect.
type ZoneActionID string

const (
	ZoneActionIDAlwaysUseHTTPS          ZoneActionID = "always_use_https"
	ZoneActionIDAutomaticHTTPSRewrites  ZoneActionID = "automatic_https_rewrites"
	ZoneActionIDBrowserCacheTtl         ZoneActionID = "browser_cache_ttl"
	ZoneActionIDBrowserCheck            ZoneActionID = "browser_check"
	ZoneActionIDBypassCacheOnCookie     ZoneActionID = "bypass_cache_on_cookie"
	ZoneActionIDCacheByDeviceType       ZoneActionID = "cache_by_device_type"
	ZoneActionIDCacheDeceptionArmor     ZoneActionID = "cache_deception_armor"
	ZoneActionIDCacheKeyFields          ZoneActionID = "cache_key_fields"
	ZoneActionIDCacheLevel              ZoneActionID = "cache_level"
	ZoneActionIDCacheOnCookie           ZoneActionID = "cache_on_cookie"
	ZoneActionIDCacheTtlByStatus        ZoneActionID = "cache_ttl_by_status"
	ZoneActionIDDisableApps             ZoneActionID = "disable_apps"
	ZoneActionIDDisablePerformance      ZoneActionID = "disable_performance"
	ZoneActionIDDisableSecurity         ZoneActionID = "disable_security"
	ZoneActionIDDisableZaraz            ZoneActionID = "disable_zaraz"
	ZoneActionIDEdgeCacheTtl            ZoneActionID = "edge_cache_ttl"
	ZoneActionIDEmailObfuscation        ZoneActionID = "email_obfuscation"
	ZoneActionIDExplicitCacheControl    ZoneActionID = "explicit_cache_control"
	ZoneActionIDForwardingURL           ZoneActionID = "forwarding_url"
	ZoneActionIDHostHeaderOverride      ZoneActionID = "host_header_override"
	ZoneActionIDIPGeolocation           ZoneActionID = "ip_geolocation"
	ZoneActionIDMirage                  ZoneActionID = "mirage"
	ZoneActionIDOpportunisticEncryption ZoneActionID = "opportunistic_encryption"
	ZoneActionIDOriginErrorPagePassThru ZoneActionID = "origin_error_page_pass_thru"
	ZoneActionIDPolish                  ZoneActionID = "polish"
	ZoneActionIDResolveOverride         ZoneActionID = "resolve_override"
	ZoneActionIDRespectStrongEtag       ZoneActionID = "respect_strong_etag"
	ZoneActionIDResponseBuffering       ZoneActionID = "response_buffering"
	ZoneActionIDRocketLoader            ZoneActionID = "rocket_loader"
	ZoneActionIDSecurityLevel           ZoneActionID = "security_level"
	ZoneActionIDSortQueryStringForCache ZoneActionID = "sort_query_string_for_cache"
	ZoneActionIDSsl                     ZoneActionID = "ssl"
	ZoneActionIDTrueClientIPHeader      ZoneActionID = "true_client_ip_header"
	ZoneActionIDWaf                     ZoneActionID = "waf"
)

func (r ZoneActionID) IsKnown() bool {
	switch r {
	case ZoneActionIDAlwaysUseHTTPS, ZoneActionIDAutomaticHTTPSRewrites, ZoneActionIDBrowserCacheTtl, ZoneActionIDBrowserCheck, ZoneActionIDBypassCacheOnCookie, ZoneActionIDCacheByDeviceType, ZoneActionIDCacheDeceptionArmor, ZoneActionIDCacheKeyFields, ZoneActionIDCacheLevel, ZoneActionIDCacheOnCookie, ZoneActionIDCacheTtlByStatus, ZoneActionIDDisableApps, ZoneActionIDDisablePerformance, ZoneActionIDDisableSecurity, ZoneActionIDDisableZaraz, ZoneActionIDEdgeCacheTtl, ZoneActionIDEmailObfuscation, ZoneActionIDExplicitCacheControl, ZoneActionIDForwardingURL, ZoneActionIDHostHeaderOverride, ZoneActionIDIPGeolocation, ZoneActionIDMirage, ZoneActionIDOpportunisticEncryption, ZoneActionIDOriginErrorPagePassThru, ZoneActionIDPolish, ZoneActionIDResolveOverride, ZoneActionIDRespectStrongEtag, ZoneActionIDResponseBuffering, ZoneActionIDRocketLoader, ZoneActionIDSecurityLevel, ZoneActionIDSortQueryStringForCache, ZoneActionIDSsl, ZoneActionIDTrueClientIPHeader, ZoneActionIDWaf:
		return true
	}
	return false
}

type ZoneActionParam struct {
	// If enabled, any ` http://“ URL is converted to  `https://` through a 301
	// redirect.
	ID    param.Field[ZoneActionID] `json:"id"`
	Value param.Field[interface{}]  `json:"value"`
}

func (r ZoneActionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionParam) implementsZoneActionUnionParam() {}

// Satisfied by [ZoneActionAlwaysUseHTTPSParam],
// [ZoneActionAutomaticHTTPSRewritesParam], [ZoneActionBrowserCacheTtlParam],
// [ZoneActionBrowserCheckParam], [ZoneActionBypassCacheOnCookieParam],
// [ZoneActionCacheByDeviceTypeParam], [ZoneActionCacheDeceptionArmorParam],
// [ZoneActionCacheKeyFieldsParam], [ZoneActionCacheLevelParam],
// [ZoneActionCacheOnCookieParam], [ZoneActionCacheTtlByStatusParam],
// [ZoneActionDisableAppsParam], [ZoneActionDisablePerformanceParam],
// [ZoneActionDisableSecurityParam], [ZoneActionDisableZarazParam],
// [ZoneActionEdgeCacheTtlParam], [ZoneActionEmailObfuscationParam],
// [ZoneActionExplicitCacheControlParam], [ZoneActionForwardingURLParam],
// [ZoneActionHostHeaderOverrideParam], [ZoneActionIPGeolocationParam],
// [ZoneActionMirageParam], [ZoneActionOpportunisticEncryptionParam],
// [ZoneActionOriginErrorPagePassThruParam], [ZoneActionPolishParam],
// [ZoneActionResolveOverrideParam], [ZoneActionRespectStrongEtagParam],
// [ZoneActionResponseBufferingParam], [ZoneActionRocketLoaderParam],
// [ZoneActionSecurityLevelParam], [ZoneActionSortQueryStringForCacheParam],
// [ZoneActionSslParam], [ZoneActionTrueClientIPHeaderParam], [ZoneActionWafParam],
// [ZoneActionParam].
type ZoneActionUnionParam interface {
	implementsZoneActionUnionParam()
}

type ZoneActionAlwaysUseHTTPSParam struct {
	// If enabled, any ` http://“ URL is converted to  `https://` through a 301
	// redirect.
	ID param.Field[ZoneActionAlwaysUseHTTPSID] `json:"id"`
}

func (r ZoneActionAlwaysUseHTTPSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionAlwaysUseHTTPSParam) implementsZoneActionUnionParam() {}

type ZoneActionAutomaticHTTPSRewritesParam struct {
	// Turn on or off Automatic HTTPS Rewrites.
	ID param.Field[ZoneActionAutomaticHTTPSRewritesID] `json:"id"`
	// The status of Automatic HTTPS Rewrites.
	Value param.Field[ZoneActionAutomaticHTTPSRewritesValue] `json:"value"`
}

func (r ZoneActionAutomaticHTTPSRewritesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionAutomaticHTTPSRewritesParam) implementsZoneActionUnionParam() {}

type ZoneActionBrowserCacheTtlParam struct {
	// Control how long resources cached by client browsers remain valid.
	ID param.Field[ZoneActionBrowserCacheTtlID] `json:"id"`
	// The number of seconds to cache resources for. The API prohibits setting this to
	// 0 for non-Enterprise domains.
	Value param.Field[int64] `json:"value"`
}

func (r ZoneActionBrowserCacheTtlParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionBrowserCacheTtlParam) implementsZoneActionUnionParam() {}

type ZoneActionBrowserCheckParam struct {
	// Inspect the visitor's browser for headers commonly associated with spammers and
	// certain bots.
	ID param.Field[ZoneActionBrowserCheckID] `json:"id"`
	// The status of Browser Integrity Check.
	Value param.Field[ZoneActionBrowserCheckValue] `json:"value"`
}

func (r ZoneActionBrowserCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionBrowserCheckParam) implementsZoneActionUnionParam() {}

type ZoneActionBypassCacheOnCookieParam struct {
	// Bypass cache and fetch resources from the origin server if a regular expression
	// matches against a cookie name present in the request.
	ID param.Field[ZoneActionBypassCacheOnCookieID] `json:"id"`
	// The regular expression to use for matching cookie names in the request. Refer to
	// [Bypass Cache on Cookie setting](https://developers.cloudflare.com/rules/page-rules/reference/additional-reference/#bypass-cache-on-cookie-setting)
	// to learn about limited regular expression support.
	Value param.Field[string] `json:"value"`
}

func (r ZoneActionBypassCacheOnCookieParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionBypassCacheOnCookieParam) implementsZoneActionUnionParam() {}

type ZoneActionCacheByDeviceTypeParam struct {
	// Separate cached content based on the visitor's device type.
	ID param.Field[ZoneActionCacheByDeviceTypeID] `json:"id"`
	// The status of Cache By Device Type.
	Value param.Field[ZoneActionCacheByDeviceTypeValue] `json:"value"`
}

func (r ZoneActionCacheByDeviceTypeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionCacheByDeviceTypeParam) implementsZoneActionUnionParam() {}

type ZoneActionCacheDeceptionArmorParam struct {
	// Protect from web cache deception attacks while still allowing static assets to
	// be cached. This setting verifies that the URL's extension matches the returned
	// `Content-Type`.
	ID param.Field[ZoneActionCacheDeceptionArmorID] `json:"id"`
	// The status of Cache Deception Armor.
	Value param.Field[ZoneActionCacheDeceptionArmorValue] `json:"value"`
}

func (r ZoneActionCacheDeceptionArmorParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionCacheDeceptionArmorParam) implementsZoneActionUnionParam() {}

type ZoneActionCacheKeyFieldsParam struct {
	// Control specifically what variables to include when deciding which resources to
	// cache. This allows customers to determine what to cache based on something other
	// than just the URL.
	ID    param.Field[ZoneActionCacheKeyFieldsID]         `json:"id"`
	Value param.Field[ZoneActionCacheKeyFieldsValueParam] `json:"value"`
}

func (r ZoneActionCacheKeyFieldsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionCacheKeyFieldsParam) implementsZoneActionUnionParam() {}

type ZoneActionCacheKeyFieldsValueParam struct {
	// Controls which cookies appear in the Cache Key.
	Cookie param.Field[ZoneActionCacheKeyFieldsValueCookieParam] `json:"cookie"`
	// Controls which headers go into the Cache Key. Exactly one of `include` or
	// `exclude` is expected.
	Header param.Field[ZoneActionCacheKeyFieldsValueHeaderParam] `json:"header"`
	// Determines which host header to include in the Cache Key.
	Host param.Field[ZoneActionCacheKeyFieldsValueHostParam] `json:"host"`
	// Controls which URL query string parameters go into the Cache Key. Exactly one of
	// `include` or `exclude` is expected.
	QueryString param.Field[ZoneActionCacheKeyFieldsValueQueryStringParam] `json:"query_string"`
	// Feature fields to add features about the end-user (client) into the Cache Key.
	User param.Field[ZoneActionCacheKeyFieldsValueUserParam] `json:"user"`
}

func (r ZoneActionCacheKeyFieldsValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which cookies appear in the Cache Key.
type ZoneActionCacheKeyFieldsValueCookieParam struct {
	// A list of cookies to check for the presence of, without including their actual
	// values.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// A list of cookies to include.
	Include param.Field[[]string] `json:"include"`
}

func (r ZoneActionCacheKeyFieldsValueCookieParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which headers go into the Cache Key. Exactly one of `include` or
// `exclude` is expected.
type ZoneActionCacheKeyFieldsValueHeaderParam struct {
	// A list of headers to check for the presence of, without including their actual
	// values.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// A list of headers to ignore.
	Exclude param.Field[[]string] `json:"exclude"`
	// A list of headers to include.
	Include param.Field[[]string] `json:"include"`
}

func (r ZoneActionCacheKeyFieldsValueHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which host header to include in the Cache Key.
type ZoneActionCacheKeyFieldsValueHostParam struct {
	// Whether to include the Host header in the HTTP request sent to the origin.
	Resolved param.Field[bool] `json:"resolved"`
}

func (r ZoneActionCacheKeyFieldsValueHostParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which URL query string parameters go into the Cache Key. Exactly one of
// `include` or `exclude` is expected.
type ZoneActionCacheKeyFieldsValueQueryStringParam struct {
	// Ignore all query string parameters.
	Exclude param.Field[ZoneActionCacheKeyFieldsValueQueryStringExcludeUnionParam] `json:"exclude"`
	// Include all query string parameters.
	Include param.Field[ZoneActionCacheKeyFieldsValueQueryStringIncludeUnionParam] `json:"include"`
}

func (r ZoneActionCacheKeyFieldsValueQueryStringParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Ignore all query string parameters.
//
// Satisfied by [ZoneActionCacheKeyFieldsValueQueryStringExcludeString],
// [ZoneActionCacheKeyFieldsValueQueryStringExcludeArrayParam].
type ZoneActionCacheKeyFieldsValueQueryStringExcludeUnionParam interface {
	implementsZoneActionCacheKeyFieldsValueQueryStringExcludeUnionParam()
}

type ZoneActionCacheKeyFieldsValueQueryStringExcludeArrayParam []string

func (r ZoneActionCacheKeyFieldsValueQueryStringExcludeArrayParam) implementsZoneActionCacheKeyFieldsValueQueryStringExcludeUnionParam() {
}

// Include all query string parameters.
//
// Satisfied by [ZoneActionCacheKeyFieldsValueQueryStringIncludeString],
// [ZoneActionCacheKeyFieldsValueQueryStringIncludeArrayParam].
type ZoneActionCacheKeyFieldsValueQueryStringIncludeUnionParam interface {
	implementsZoneActionCacheKeyFieldsValueQueryStringIncludeUnionParam()
}

type ZoneActionCacheKeyFieldsValueQueryStringIncludeArrayParam []string

func (r ZoneActionCacheKeyFieldsValueQueryStringIncludeArrayParam) implementsZoneActionCacheKeyFieldsValueQueryStringIncludeUnionParam() {
}

// Feature fields to add features about the end-user (client) into the Cache Key.
type ZoneActionCacheKeyFieldsValueUserParam struct {
	// Classifies a request as `mobile`, `desktop`, or `tablet` based on the User
	// Agent.
	DeviceType param.Field[bool] `json:"device_type"`
	// Includes the client's country, derived from the IP address.
	Geo param.Field[bool] `json:"geo"`
	// Includes the first language code contained in the `Accept-Language` header sent
	// by the client.
	Lang param.Field[bool] `json:"lang"`
}

func (r ZoneActionCacheKeyFieldsValueUserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneActionCacheLevelParam struct {
	// Apply custom caching based on the option selected.
	ID param.Field[ZoneActionCacheLevelID] `json:"id"`
	//   - `bypass`: Cloudflare does not cache.
	//   - `basic`: Delivers resources from cache when there is no query string.
	//   - `simplified`: Delivers the same resource to everyone independent of the query
	//     string.
	//   - `aggressive`: Caches all static content that has a query string.
	//   - `cache_everything`: Treats all content as static and caches all file types
	//     beyond the
	//     [Cloudflare default cached content](https://developers.cloudflare.com/cache/concepts/default-cache-behavior/#default-cached-file-extensions).
	Value param.Field[ZoneActionCacheLevelValue] `json:"value"`
}

func (r ZoneActionCacheLevelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionCacheLevelParam) implementsZoneActionUnionParam() {}

type ZoneActionCacheOnCookieParam struct {
	// Apply the Cache Everything option (Cache Level setting) based on a regular
	// expression match against a cookie name.
	ID param.Field[ZoneActionCacheOnCookieID] `json:"id"`
	// The regular expression to use for matching cookie names in the request.
	Value param.Field[string] `json:"value"`
}

func (r ZoneActionCacheOnCookieParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionCacheOnCookieParam) implementsZoneActionUnionParam() {}

type ZoneActionCacheTtlByStatusParam struct {
	// Enterprise customers can set cache time-to-live (TTL) based on the response
	// status from the origin web server. Cache TTL refers to the duration of a
	// resource in the Cloudflare network before being marked as stale or discarded
	// from cache. Status codes are returned by a resource's origin. Setting cache TTL
	// based on response status overrides the default cache behavior (standard caching)
	// for static files and overrides cache instructions sent by the origin web server.
	// To cache non-static assets, set a Cache Level of Cache Everything using a Page
	// Rule. Setting no-store Cache-Control or a low TTL (using `max-age`/`s-maxage`)
	// increases requests to origin web servers and decreases performance.
	ID param.Field[ZoneActionCacheTtlByStatusID] `json:"id"`
	// A JSON object containing status codes and their corresponding TTLs. Each
	// key-value pair in the cache TTL by status cache rule has the following syntax
	//
	//   - `status_code`: An integer value such as 200 or 500. status_code matches the
	//     exact status code from the origin web server. Valid status codes are between
	//     100-999.
	//   - `status_code_range`: Integer values for from and to. status_code_range matches
	//     any status code from the origin web server within the specified range.
	//   - `value`: An integer value that defines the duration an asset is valid in
	//     seconds or one of the following strings: no-store (equivalent to -1), no-cache
	//     (equivalent to 0).
	Value param.Field[map[string]ZoneActionCacheTtlByStatusValueUnionParam] `json:"value"`
}

func (r ZoneActionCacheTtlByStatusParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionCacheTtlByStatusParam) implementsZoneActionUnionParam() {}

// `no-store` (equivalent to -1), `no-cache` (equivalent to 0)
//
// Satisfied by [ZoneActionCacheTtlByStatusValueString], [shared.UnionInt].
type ZoneActionCacheTtlByStatusValueUnionParam interface {
	ImplementsZoneActionCacheTtlByStatusValueUnionParam()
}

type ZoneActionDisableAppsParam struct {
	// Turn off all active
	// [Cloudflare Apps](https://developers.cloudflare.com/support/more-dashboard-apps/cloudflare-apps/)
	// (deprecated).
	ID param.Field[ZoneActionDisableAppsID] `json:"id"`
}

func (r ZoneActionDisableAppsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionDisableAppsParam) implementsZoneActionUnionParam() {}

type ZoneActionDisablePerformanceParam struct {
	// Turn off
	// [Rocket Loader](https://developers.cloudflare.com/speed/optimization/content/rocket-loader/),
	// [Mirage](https://developers.cloudflare.com/speed/optimization/images/mirage/),
	// and [Polish](https://developers.cloudflare.com/images/polish/).
	ID param.Field[ZoneActionDisablePerformanceID] `json:"id"`
}

func (r ZoneActionDisablePerformanceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionDisablePerformanceParam) implementsZoneActionUnionParam() {}

type ZoneActionDisableSecurityParam struct {
	// Turn off
	// [Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),
	// [Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),
	// [Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),
	// [URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/),
	// and
	// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
	ID param.Field[ZoneActionDisableSecurityID] `json:"id"`
}

func (r ZoneActionDisableSecurityParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionDisableSecurityParam) implementsZoneActionUnionParam() {}

type ZoneActionDisableZarazParam struct {
	// Turn off [Zaraz](https://developers.cloudflare.com/zaraz/).
	ID param.Field[ZoneActionDisableZarazID] `json:"id"`
}

func (r ZoneActionDisableZarazParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionDisableZarazParam) implementsZoneActionUnionParam() {}

type ZoneActionEdgeCacheTtlParam struct {
	// Specify how long to cache a resource in the Cloudflare global network. _Edge
	// Cache TTL_ is not visible in response headers.
	ID    param.Field[ZoneActionEdgeCacheTtlID] `json:"id"`
	Value param.Field[int64]                    `json:"value"`
}

func (r ZoneActionEdgeCacheTtlParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionEdgeCacheTtlParam) implementsZoneActionUnionParam() {}

type ZoneActionEmailObfuscationParam struct {
	// Turn on or off **Email Obfuscation**.
	ID param.Field[ZoneActionEmailObfuscationID] `json:"id"`
	// The status of Email Obfuscation.
	Value param.Field[ZoneActionEmailObfuscationValue] `json:"value"`
}

func (r ZoneActionEmailObfuscationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionEmailObfuscationParam) implementsZoneActionUnionParam() {}

type ZoneActionExplicitCacheControlParam struct {
	// Origin Cache Control is enabled by default for Free, Pro, and Business domains
	// and disabled by default for Enterprise domains.
	ID param.Field[ZoneActionExplicitCacheControlID] `json:"id"`
	// The status of Origin Cache Control.
	Value param.Field[ZoneActionExplicitCacheControlValue] `json:"value"`
}

func (r ZoneActionExplicitCacheControlParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionExplicitCacheControlParam) implementsZoneActionUnionParam() {}

type ZoneActionForwardingURLParam struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    param.Field[ZoneActionForwardingURLID]         `json:"id"`
	Value param.Field[ZoneActionForwardingURLValueParam] `json:"value"`
}

func (r ZoneActionForwardingURLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionForwardingURLParam) implementsZoneActionUnionParam() {}

type ZoneActionForwardingURLValueParam struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode param.Field[ZoneActionForwardingURLValueStatusCode] `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r ZoneActionForwardingURLValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneActionHostHeaderOverrideParam struct {
	// Apply a specific host header.
	ID param.Field[ZoneActionHostHeaderOverrideID] `json:"id"`
	// The hostname to use in the `Host` header
	Value param.Field[string] `json:"value"`
}

func (r ZoneActionHostHeaderOverrideParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionHostHeaderOverrideParam) implementsZoneActionUnionParam() {}

type ZoneActionIPGeolocationParam struct {
	// Cloudflare adds a CF-IPCountry HTTP header containing the country code that
	// corresponds to the visitor.
	ID param.Field[ZoneActionIPGeolocationID] `json:"id"`
	// The status of adding the IP Geolocation Header.
	Value param.Field[ZoneActionIPGeolocationValue] `json:"value"`
}

func (r ZoneActionIPGeolocationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionIPGeolocationParam) implementsZoneActionUnionParam() {}

type ZoneActionMirageParam struct {
	// Cloudflare Mirage reduces bandwidth used by images in mobile browsers. It can
	// accelerate loading of image-heavy websites on very slow mobile connections and
	// HTTP/1.
	ID param.Field[ZoneActionMirageID] `json:"id"`
	// The status of Mirage.
	Value param.Field[ZoneActionMirageValue] `json:"value"`
}

func (r ZoneActionMirageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionMirageParam) implementsZoneActionUnionParam() {}

type ZoneActionOpportunisticEncryptionParam struct {
	// Opportunistic Encryption allows browsers to access HTTP URIs over an encrypted
	// TLS channel. It's not a substitute for HTTPS, but provides additional security
	// for otherwise vulnerable requests.
	ID param.Field[ZoneActionOpportunisticEncryptionID] `json:"id"`
	// The status of Opportunistic Encryption.
	Value param.Field[ZoneActionOpportunisticEncryptionValue] `json:"value"`
}

func (r ZoneActionOpportunisticEncryptionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionOpportunisticEncryptionParam) implementsZoneActionUnionParam() {}

type ZoneActionOriginErrorPagePassThruParam struct {
	// Turn on or off Cloudflare error pages generated from issues sent from the origin
	// server. If enabled, this setting triggers error pages issued by the origin.
	ID param.Field[ZoneActionOriginErrorPagePassThruID] `json:"id"`
	// The status of Origin Error Page Passthru.
	Value param.Field[ZoneActionOriginErrorPagePassThruValue] `json:"value"`
}

func (r ZoneActionOriginErrorPagePassThruParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionOriginErrorPagePassThruParam) implementsZoneActionUnionParam() {}

type ZoneActionPolishParam struct {
	// Apply options from the Polish feature of the Cloudflare Speed app.
	ID param.Field[ZoneActionPolishID] `json:"id"`
	// The level of Polish you want applied to your origin.
	Value param.Field[ZoneActionPolishValue] `json:"value"`
}

func (r ZoneActionPolishParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionPolishParam) implementsZoneActionUnionParam() {}

type ZoneActionResolveOverrideParam struct {
	// Change the origin address to the value specified in this setting.
	ID param.Field[ZoneActionResolveOverrideID] `json:"id"`
	// The origin address you want to override with.
	Value param.Field[string] `json:"value"`
}

func (r ZoneActionResolveOverrideParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionResolveOverrideParam) implementsZoneActionUnionParam() {}

type ZoneActionRespectStrongEtagParam struct {
	// Turn on or off byte-for-byte equivalency checks between the Cloudflare cache and
	// the origin server.
	ID param.Field[ZoneActionRespectStrongEtagID] `json:"id"`
	// The status of Respect Strong ETags
	Value param.Field[ZoneActionRespectStrongEtagValue] `json:"value"`
}

func (r ZoneActionRespectStrongEtagParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionRespectStrongEtagParam) implementsZoneActionUnionParam() {}

type ZoneActionResponseBufferingParam struct {
	// Turn on or off whether Cloudflare should wait for an entire file from the origin
	// server before forwarding it to the site visitor. By default, Cloudflare sends
	// packets to the client as they arrive from the origin server.
	ID param.Field[ZoneActionResponseBufferingID] `json:"id"`
	// The status of Response Buffering
	Value param.Field[ZoneActionResponseBufferingValue] `json:"value"`
}

func (r ZoneActionResponseBufferingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionResponseBufferingParam) implementsZoneActionUnionParam() {}

type ZoneActionRocketLoaderParam struct {
	// Turn on or off Rocket Loader in the Cloudflare Speed app.
	ID param.Field[ZoneActionRocketLoaderID] `json:"id"`
	// The status of Rocket Loader
	Value param.Field[ZoneActionRocketLoaderValue] `json:"value"`
}

func (r ZoneActionRocketLoaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionRocketLoaderParam) implementsZoneActionUnionParam() {}

type ZoneActionSecurityLevelParam struct {
	// Control options for the **Security Level** feature from the **Security** app.
	ID    param.Field[ZoneActionSecurityLevelID]    `json:"id"`
	Value param.Field[ZoneActionSecurityLevelValue] `json:"value"`
}

func (r ZoneActionSecurityLevelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionSecurityLevelParam) implementsZoneActionUnionParam() {}

type ZoneActionSortQueryStringForCacheParam struct {
	// Turn on or off the reordering of query strings. When query strings have the same
	// structure, caching improves.
	ID param.Field[ZoneActionSortQueryStringForCacheID] `json:"id"`
	// The status of Query String Sort
	Value param.Field[ZoneActionSortQueryStringForCacheValue] `json:"value"`
}

func (r ZoneActionSortQueryStringForCacheParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionSortQueryStringForCacheParam) implementsZoneActionUnionParam() {}

type ZoneActionSslParam struct {
	// Control options for the SSL feature of the Edge Certificates tab in the
	// Cloudflare SSL/TLS app.
	ID param.Field[ZoneActionSslID] `json:"id"`
	// The encryption mode that Cloudflare uses to connect to your origin server.
	Value param.Field[ZoneActionSslValue] `json:"value"`
}

func (r ZoneActionSslParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionSslParam) implementsZoneActionUnionParam() {}

type ZoneActionTrueClientIPHeaderParam struct {
	// Turn on or off the True-Client-IP Header feature of the Cloudflare Network app.
	ID param.Field[ZoneActionTrueClientIPHeaderID] `json:"id"`
	// The status of True Client IP Header.
	Value param.Field[ZoneActionTrueClientIPHeaderValue] `json:"value"`
}

func (r ZoneActionTrueClientIPHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionTrueClientIPHeaderParam) implementsZoneActionUnionParam() {}

type ZoneActionWafParam struct {
	// Turn on or off
	// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
	// You cannot enable or disable individual WAF managed rules via Page Rules.
	ID param.Field[ZoneActionWafID] `json:"id"`
	// The status of WAF managed rules (previous version).
	Value param.Field[ZoneActionWafValue] `json:"value"`
}

func (r ZoneActionWafParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneActionWafParam) implementsZoneActionUnionParam() {}

type ZonePageruleNewResponse struct {
	Result PageRule                    `json:"result"`
	JSON   zonePageruleNewResponseJSON `json:"-"`
	APIResponseSingleZones
}

// zonePageruleNewResponseJSON contains the JSON metadata for the struct
// [ZonePageruleNewResponse]
type zonePageruleNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageruleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageruleNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageruleGetResponse struct {
	Result PageRule                    `json:"result"`
	JSON   zonePageruleGetResponseJSON `json:"-"`
	APIResponseSingleZones
}

// zonePageruleGetResponseJSON contains the JSON metadata for the struct
// [ZonePageruleGetResponse]
type zonePageruleGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageruleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageruleGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageruleUpdateResponse struct {
	Result PageRule                       `json:"result"`
	JSON   zonePageruleUpdateResponseJSON `json:"-"`
	APIResponseSingleZones
}

// zonePageruleUpdateResponseJSON contains the JSON metadata for the struct
// [ZonePageruleUpdateResponse]
type zonePageruleUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageruleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageruleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageruleListResponse struct {
	Result []PageRule                   `json:"result"`
	JSON   zonePageruleListResponseJSON `json:"-"`
	APIResponseZonesSchemas
}

// zonePageruleListResponseJSON contains the JSON metadata for the struct
// [ZonePageruleListResponse]
type zonePageruleListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageruleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageruleListResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageruleDeleteResponse struct {
	Result ZonePageruleDeleteResponseResult `json:"result,nullable"`
	JSON   zonePageruleDeleteResponseJSON   `json:"-"`
	APIResponseZonesSchemas
}

// zonePageruleDeleteResponseJSON contains the JSON metadata for the struct
// [ZonePageruleDeleteResponse]
type zonePageruleDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageruleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageruleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageruleDeleteResponseResult struct {
	// Identifier
	ID   string                               `json:"id,required"`
	JSON zonePageruleDeleteResponseResultJSON `json:"-"`
}

// zonePageruleDeleteResponseResultJSON contains the JSON metadata for the struct
// [ZonePageruleDeleteResponseResult]
type zonePageruleDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageruleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageruleDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZonePageruleEditResponse struct {
	Result PageRule                     `json:"result"`
	JSON   zonePageruleEditResponseJSON `json:"-"`
	APIResponseSingleZones
}

// zonePageruleEditResponseJSON contains the JSON metadata for the struct
// [ZonePageruleEditResponse]
type zonePageruleEditResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageruleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageruleEditResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageruleListSettingsResponse struct {
	// Settings available for the zone.
	Result []interface{}                        `json:"result"`
	JSON   zonePageruleListSettingsResponseJSON `json:"-"`
	APIResponseZonesSchemas
}

// zonePageruleListSettingsResponseJSON contains the JSON metadata for the struct
// [ZonePageruleListSettingsResponse]
type zonePageruleListSettingsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageruleListSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageruleListSettingsResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageruleNewParams struct {
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]ZoneActionUnionParam] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]RequestConditionTargetParam] `json:"targets,required"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageRuleStatus] `json:"status"`
}

func (r ZonePageruleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZonePageruleUpdateParams struct {
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]ZoneActionUnionParam] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]RequestConditionTargetParam] `json:"targets,required"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageRuleStatus] `json:"status"`
}

func (r ZonePageruleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZonePageruleListParams struct {
	// The direction used to sort returned Page Rules.
	Direction param.Field[ZonePageruleListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[ZonePageruleListParamsMatch] `query:"match"`
	// The field used to sort returned Page Rules.
	Order param.Field[ZonePageruleListParamsOrder] `query:"order"`
	// The status of the Page Rule.
	Status param.Field[ZonePageruleListParamsStatus] `query:"status"`
}

// URLQuery serializes [ZonePageruleListParams]'s query parameters as `url.Values`.
func (r ZonePageruleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned Page Rules.
type ZonePageruleListParamsDirection string

const (
	ZonePageruleListParamsDirectionAsc  ZonePageruleListParamsDirection = "asc"
	ZonePageruleListParamsDirectionDesc ZonePageruleListParamsDirection = "desc"
)

func (r ZonePageruleListParamsDirection) IsKnown() bool {
	switch r {
	case ZonePageruleListParamsDirectionAsc, ZonePageruleListParamsDirectionDesc:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type ZonePageruleListParamsMatch string

const (
	ZonePageruleListParamsMatchAny ZonePageruleListParamsMatch = "any"
	ZonePageruleListParamsMatchAll ZonePageruleListParamsMatch = "all"
)

func (r ZonePageruleListParamsMatch) IsKnown() bool {
	switch r {
	case ZonePageruleListParamsMatchAny, ZonePageruleListParamsMatchAll:
		return true
	}
	return false
}

// The field used to sort returned Page Rules.
type ZonePageruleListParamsOrder string

const (
	ZonePageruleListParamsOrderStatus   ZonePageruleListParamsOrder = "status"
	ZonePageruleListParamsOrderPriority ZonePageruleListParamsOrder = "priority"
)

func (r ZonePageruleListParamsOrder) IsKnown() bool {
	switch r {
	case ZonePageruleListParamsOrderStatus, ZonePageruleListParamsOrderPriority:
		return true
	}
	return false
}

// The status of the Page Rule.
type ZonePageruleListParamsStatus string

const (
	ZonePageruleListParamsStatusActive   ZonePageruleListParamsStatus = "active"
	ZonePageruleListParamsStatusDisabled ZonePageruleListParamsStatus = "disabled"
)

func (r ZonePageruleListParamsStatus) IsKnown() bool {
	switch r {
	case ZonePageruleListParamsStatusActive, ZonePageruleListParamsStatusDisabled:
		return true
	}
	return false
}

type ZonePageruleEditParams struct {
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]ZoneActionUnionParam] `json:"actions"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageRuleStatus] `json:"status"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]RequestConditionTargetParam] `json:"targets"`
}

func (r ZonePageruleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
