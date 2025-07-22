// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// ZoneFirewallWafPackageService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneFirewallWafPackageService] method instead.
type ZoneFirewallWafPackageService struct {
	Options []option.RequestOption
	Groups  *ZoneFirewallWafPackageGroupService
	Rules   *ZoneFirewallWafPackageRuleService
}

// NewZoneFirewallWafPackageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneFirewallWafPackageService(opts ...option.RequestOption) (r *ZoneFirewallWafPackageService) {
	r = &ZoneFirewallWafPackageService{}
	r.Options = opts
	r.Groups = NewZoneFirewallWafPackageGroupService(opts...)
	r.Rules = NewZoneFirewallWafPackageRuleService(opts...)
	return
}

// Fetches the details of a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *ZoneFirewallWafPackageService) Get(ctx context.Context, zoneID string, packageID string, opts ...option.RequestOption) (res *FirewallPackageResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s", zoneID, packageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a WAF package. You can update the sensitivity and the action of an
// anomaly detection WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *ZoneFirewallWafPackageService) Update(ctx context.Context, zoneID string, packageID string, body ZoneFirewallWafPackageUpdateParams, opts ...option.RequestOption) (res *ZoneFirewallWafPackageUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s", zoneID, packageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetches WAF packages for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *ZoneFirewallWafPackageService) List(ctx context.Context, zoneID string, query ZoneFirewallWafPackageListParams, opts ...option.RequestOption) (res *ZoneFirewallWafPackageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The default action performed by the rules in the WAF package.
type FirewallActionMode string

const (
	FirewallActionModeSimulate  FirewallActionMode = "simulate"
	FirewallActionModeBlock     FirewallActionMode = "block"
	FirewallActionModeChallenge FirewallActionMode = "challenge"
)

func (r FirewallActionMode) IsKnown() bool {
	switch r {
	case FirewallActionModeSimulate, FirewallActionModeBlock, FirewallActionModeChallenge:
		return true
	}
	return false
}

type FirewallAnomalyPackage struct {
	// The default action performed by the rules in the WAF package.
	ActionMode FirewallActionMode `json:"action_mode"`
	// A summary of the purpose/function of the WAF package.
	Description string `json:"description"`
	// When a WAF package uses anomaly detection, each rule is given a score when
	// triggered. If the total score of all triggered rules exceeds the sensitivity
	// defined on the WAF package, the action defined on the package will be taken.
	DetectionMode string `json:"detection_mode"`
	// The name of the WAF package.
	Name string `json:"name"`
	// The sensitivity of the WAF package.
	Sensitivity FirewallSensitivity        `json:"sensitivity"`
	JSON        firewallAnomalyPackageJSON `json:"-"`
	FirewallPackageDefinition
}

// firewallAnomalyPackageJSON contains the JSON metadata for the struct
// [FirewallAnomalyPackage]
type firewallAnomalyPackageJSON struct {
	ActionMode    apijson.Field
	Description   apijson.Field
	DetectionMode apijson.Field
	Name          apijson.Field
	Sensitivity   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallAnomalyPackage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAnomalyPackageJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAnomalyPackage) implementsZoneFirewallWafPackageListResponseResultResult() {}

type FirewallPackageDefinition struct {
	// Identifier
	ID string `json:"id,required"`
	// A summary of the purpose/function of the WAF package.
	Description string `json:"description,required"`
	// The mode that defines how rules within the package are evaluated during the
	// course of a request. When a package uses anomaly detection mode (`anomaly`
	// value), each rule is given a score when triggered. If the total score of all
	// triggered rules exceeds the sensitivity defined in the WAF package, the action
	// configured in the package will be performed. Traditional detection mode
	// (`traditional` value) will decide the action to take when it is triggered by the
	// request. If multiple rules are triggered, the action providing the highest
	// protection will be applied (for example, a 'block' action will win over a
	// 'challenge' action).
	DetectionMode FirewallPackageDefinitionDetectionMode `json:"detection_mode,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status FirewallPackageDefinitionStatus `json:"status"`
	JSON   firewallPackageDefinitionJSON   `json:"-"`
}

// firewallPackageDefinitionJSON contains the JSON metadata for the struct
// [FirewallPackageDefinition]
type firewallPackageDefinitionJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	DetectionMode apijson.Field
	Name          apijson.Field
	ZoneID        apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallPackageDefinition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallPackageDefinitionJSON) RawJSON() string {
	return r.raw
}

func (r FirewallPackageDefinition) implementsZoneFirewallWafPackageListResponseResultResult() {}

// The mode that defines how rules within the package are evaluated during the
// course of a request. When a package uses anomaly detection mode (`anomaly`
// value), each rule is given a score when triggered. If the total score of all
// triggered rules exceeds the sensitivity defined in the WAF package, the action
// configured in the package will be performed. Traditional detection mode
// (`traditional` value) will decide the action to take when it is triggered by the
// request. If multiple rules are triggered, the action providing the highest
// protection will be applied (for example, a 'block' action will win over a
// 'challenge' action).
type FirewallPackageDefinitionDetectionMode string

const (
	FirewallPackageDefinitionDetectionModeAnomaly     FirewallPackageDefinitionDetectionMode = "anomaly"
	FirewallPackageDefinitionDetectionModeTraditional FirewallPackageDefinitionDetectionMode = "traditional"
)

func (r FirewallPackageDefinitionDetectionMode) IsKnown() bool {
	switch r {
	case FirewallPackageDefinitionDetectionModeAnomaly, FirewallPackageDefinitionDetectionModeTraditional:
		return true
	}
	return false
}

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type FirewallPackageDefinitionStatus string

const (
	FirewallPackageDefinitionStatusActive FirewallPackageDefinitionStatus = "active"
)

func (r FirewallPackageDefinitionStatus) IsKnown() bool {
	switch r {
	case FirewallPackageDefinitionStatusActive:
		return true
	}
	return false
}

type FirewallPackageResponseSingle struct {
	// This field can have the runtime type of [[]FirewallMessagesItem].
	Errors interface{} `json:"errors"`
	// This field can have the runtime type of [[]FirewallMessagesItem].
	Messages interface{} `json:"messages"`
	// This field can have the runtime type of [FirewallAPIResponseCommonResultUnion],
	// [interface{}].
	Result interface{} `json:"result"`
	// Whether the API call was successful
	Success FirewallPackageResponseSingleSuccess `json:"success"`
	JSON    firewallPackageResponseSingleJSON    `json:"-"`
	union   FirewallPackageResponseSingleUnion
}

// firewallPackageResponseSingleJSON contains the JSON metadata for the struct
// [FirewallPackageResponseSingle]
type firewallPackageResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r firewallPackageResponseSingleJSON) RawJSON() string {
	return r.raw
}

func (r *FirewallPackageResponseSingle) UnmarshalJSON(data []byte) (err error) {
	*r = FirewallPackageResponseSingle{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FirewallPackageResponseSingleUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [FirewallAPIResponseSingle],
// [FirewallPackageResponseSingleResult].
func (r FirewallPackageResponseSingle) AsUnion() FirewallPackageResponseSingleUnion {
	return r.union
}

// Union satisfied by [FirewallAPIResponseSingle] or
// [FirewallPackageResponseSingleResult].
type FirewallPackageResponseSingleUnion interface {
	implementsFirewallPackageResponseSingle()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallPackageResponseSingleUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAPIResponseSingle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallPackageResponseSingleResult{}),
		},
	)
}

type FirewallPackageResponseSingleResult struct {
	Result interface{}                             `json:"result"`
	JSON   firewallPackageResponseSingleResultJSON `json:"-"`
}

// firewallPackageResponseSingleResultJSON contains the JSON metadata for the
// struct [FirewallPackageResponseSingleResult]
type firewallPackageResponseSingleResultJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallPackageResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallPackageResponseSingleResultJSON) RawJSON() string {
	return r.raw
}

func (r FirewallPackageResponseSingleResult) implementsFirewallPackageResponseSingle() {}

// Whether the API call was successful
type FirewallPackageResponseSingleSuccess bool

const (
	FirewallPackageResponseSingleSuccessTrue FirewallPackageResponseSingleSuccess = true
)

func (r FirewallPackageResponseSingleSuccess) IsKnown() bool {
	switch r {
	case FirewallPackageResponseSingleSuccessTrue:
		return true
	}
	return false
}

// The sensitivity of the WAF package.
type FirewallSensitivity string

const (
	FirewallSensitivityHigh   FirewallSensitivity = "high"
	FirewallSensitivityMedium FirewallSensitivity = "medium"
	FirewallSensitivityLow    FirewallSensitivity = "low"
	FirewallSensitivityOff    FirewallSensitivity = "off"
)

func (r FirewallSensitivity) IsKnown() bool {
	switch r {
	case FirewallSensitivityHigh, FirewallSensitivityMedium, FirewallSensitivityLow, FirewallSensitivityOff:
		return true
	}
	return false
}

type ZoneFirewallWafPackageUpdateResponse struct {
	// This field can have the runtime type of [[]FirewallMessagesItem].
	Errors interface{} `json:"errors"`
	// This field can have the runtime type of [[]FirewallMessagesItem].
	Messages interface{} `json:"messages"`
	// This field can have the runtime type of [FirewallAPIResponseCommonResultUnion],
	// [FirewallAnomalyPackage].
	Result interface{} `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallWafPackageUpdateResponseSuccess `json:"success"`
	JSON    zoneFirewallWafPackageUpdateResponseJSON    `json:"-"`
	union   ZoneFirewallWafPackageUpdateResponseUnion
}

// zoneFirewallWafPackageUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallWafPackageUpdateResponse]
type zoneFirewallWafPackageUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r zoneFirewallWafPackageUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ZoneFirewallWafPackageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ZoneFirewallWafPackageUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZoneFirewallWafPackageUpdateResponseUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ZoneFirewallWafPackageUpdateResponseAccountsFirewallAccessRulesRulesFirewallAPIResponseSingle],
// [ZoneFirewallWafPackageUpdateResponseResult].
func (r ZoneFirewallWafPackageUpdateResponse) AsUnion() ZoneFirewallWafPackageUpdateResponseUnion {
	return r.union
}

// Union satisfied by
// [ZoneFirewallWafPackageUpdateResponseAccountsFirewallAccessRulesRulesFirewallAPIResponseSingle]
// or [ZoneFirewallWafPackageUpdateResponseResult].
type ZoneFirewallWafPackageUpdateResponseUnion interface {
	implementsZoneFirewallWafPackageUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneFirewallWafPackageUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneFirewallWafPackageUpdateResponseAccountsFirewallAccessRulesRulesFirewallAPIResponseSingle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneFirewallWafPackageUpdateResponseResult{}),
		},
	)
}

type ZoneFirewallWafPackageUpdateResponseAccountsFirewallAccessRulesRulesFirewallAPIResponseSingle struct {
	Result FirewallAnomalyPackage                                                                            `json:"result"`
	JSON   zoneFirewallWafPackageUpdateResponseAccountsFirewallAccessRulesRulesFirewallAPIResponseSingleJSON `json:"-"`
	FirewallAPIResponseSingle
}

// zoneFirewallWafPackageUpdateResponseAccountsFirewallAccessRulesRulesFirewallAPIResponseSingleJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageUpdateResponseAccountsFirewallAccessRulesRulesFirewallAPIResponseSingle]
type zoneFirewallWafPackageUpdateResponseAccountsFirewallAccessRulesRulesFirewallAPIResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageUpdateResponseAccountsFirewallAccessRulesRulesFirewallAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallWafPackageUpdateResponseAccountsFirewallAccessRulesRulesFirewallAPIResponseSingleJSON) RawJSON() string {
	return r.raw
}

func (r ZoneFirewallWafPackageUpdateResponseAccountsFirewallAccessRulesRulesFirewallAPIResponseSingle) implementsZoneFirewallWafPackageUpdateResponse() {
}

type ZoneFirewallWafPackageUpdateResponseResult struct {
	Result FirewallAnomalyPackage                         `json:"result"`
	JSON   zoneFirewallWafPackageUpdateResponseResultJSON `json:"-"`
}

// zoneFirewallWafPackageUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneFirewallWafPackageUpdateResponseResult]
type zoneFirewallWafPackageUpdateResponseResultJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallWafPackageUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r ZoneFirewallWafPackageUpdateResponseResult) implementsZoneFirewallWafPackageUpdateResponse() {
}

// Whether the API call was successful
type ZoneFirewallWafPackageUpdateResponseSuccess bool

const (
	ZoneFirewallWafPackageUpdateResponseSuccessTrue ZoneFirewallWafPackageUpdateResponseSuccess = true
)

func (r ZoneFirewallWafPackageUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneFirewallWafPackageListResponse struct {
	// This field can have the runtime type of [[]FirewallMessagesItem].
	Errors interface{} `json:"errors"`
	// This field can have the runtime type of [[]FirewallMessagesItem].
	Messages interface{} `json:"messages"`
	// This field can have the runtime type of [FirewallAPIResponseCommonResultUnion],
	// [[]ZoneFirewallWafPackageListResponseResultResult].
	Result interface{} `json:"result"`
	// This field can have the runtime type of
	// [FirewallAPIResponseCollectionResultInfo].
	ResultInfo interface{} `json:"result_info"`
	// Whether the API call was successful
	Success ZoneFirewallWafPackageListResponseSuccess `json:"success"`
	JSON    zoneFirewallWafPackageListResponseJSON    `json:"-"`
	union   ZoneFirewallWafPackageListResponseUnion
}

// zoneFirewallWafPackageListResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallWafPackageListResponse]
type zoneFirewallWafPackageListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r zoneFirewallWafPackageListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ZoneFirewallWafPackageListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ZoneFirewallWafPackageListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZoneFirewallWafPackageListResponseUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [FirewallAPIResponseCollection],
// [ZoneFirewallWafPackageListResponseResult].
func (r ZoneFirewallWafPackageListResponse) AsUnion() ZoneFirewallWafPackageListResponseUnion {
	return r.union
}

// Union satisfied by [FirewallAPIResponseCollection] or
// [ZoneFirewallWafPackageListResponseResult].
type ZoneFirewallWafPackageListResponseUnion interface {
	implementsZoneFirewallWafPackageListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneFirewallWafPackageListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAPIResponseCollection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneFirewallWafPackageListResponseResult{}),
		},
	)
}

type ZoneFirewallWafPackageListResponseResult struct {
	Result []ZoneFirewallWafPackageListResponseResultResult `json:"result"`
	JSON   zoneFirewallWafPackageListResponseResultJSON     `json:"-"`
}

// zoneFirewallWafPackageListResponseResultJSON contains the JSON metadata for the
// struct [ZoneFirewallWafPackageListResponseResult]
type zoneFirewallWafPackageListResponseResultJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallWafPackageListResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r ZoneFirewallWafPackageListResponseResult) implementsZoneFirewallWafPackageListResponse() {}

type ZoneFirewallWafPackageListResponseResultResult struct {
	// Identifier
	ID string `json:"id"`
	// The default action performed by the rules in the WAF package.
	ActionMode FirewallActionMode `json:"action_mode"`
	// A summary of the purpose/function of the WAF package.
	Description string `json:"description"`
	// The mode that defines how rules within the package are evaluated during the
	// course of a request. When a package uses anomaly detection mode (`anomaly`
	// value), each rule is given a score when triggered. If the total score of all
	// triggered rules exceeds the sensitivity defined in the WAF package, the action
	// configured in the package will be performed. Traditional detection mode
	// (`traditional` value) will decide the action to take when it is triggered by the
	// request. If multiple rules are triggered, the action providing the highest
	// protection will be applied (for example, a 'block' action will win over a
	// 'challenge' action).
	DetectionMode string `json:"detection_mode"`
	// The name of the WAF package.
	Name string `json:"name"`
	// The sensitivity of the WAF package.
	Sensitivity FirewallSensitivity `json:"sensitivity"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status ZoneFirewallWafPackageListResponseResultResultStatus `json:"status"`
	// Identifier
	ZoneID string                                             `json:"zone_id"`
	JSON   zoneFirewallWafPackageListResponseResultResultJSON `json:"-"`
	union  ZoneFirewallWafPackageListResponseResultResultUnion
}

// zoneFirewallWafPackageListResponseResultResultJSON contains the JSON metadata
// for the struct [ZoneFirewallWafPackageListResponseResultResult]
type zoneFirewallWafPackageListResponseResultResultJSON struct {
	ID            apijson.Field
	ActionMode    apijson.Field
	Description   apijson.Field
	DetectionMode apijson.Field
	Name          apijson.Field
	Sensitivity   apijson.Field
	Status        apijson.Field
	ZoneID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r zoneFirewallWafPackageListResponseResultResultJSON) RawJSON() string {
	return r.raw
}

func (r *ZoneFirewallWafPackageListResponseResultResult) UnmarshalJSON(data []byte) (err error) {
	*r = ZoneFirewallWafPackageListResponseResultResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZoneFirewallWafPackageListResponseResultResultUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [FirewallPackageDefinition],
// [FirewallAnomalyPackage].
func (r ZoneFirewallWafPackageListResponseResultResult) AsUnion() ZoneFirewallWafPackageListResponseResultResultUnion {
	return r.union
}

// Union satisfied by [FirewallPackageDefinition] or [FirewallAnomalyPackage].
type ZoneFirewallWafPackageListResponseResultResultUnion interface {
	implementsZoneFirewallWafPackageListResponseResultResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneFirewallWafPackageListResponseResultResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallPackageDefinition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAnomalyPackage{}),
		},
	)
}

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type ZoneFirewallWafPackageListResponseResultResultStatus string

const (
	ZoneFirewallWafPackageListResponseResultResultStatusActive ZoneFirewallWafPackageListResponseResultResultStatus = "active"
)

func (r ZoneFirewallWafPackageListResponseResultResultStatus) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageListResponseResultResultStatusActive:
		return true
	}
	return false
}

// Whether the API call was successful
type ZoneFirewallWafPackageListResponseSuccess bool

const (
	ZoneFirewallWafPackageListResponseSuccessTrue ZoneFirewallWafPackageListResponseSuccess = true
)

func (r ZoneFirewallWafPackageListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneFirewallWafPackageUpdateParams struct {
	// The default action performed by the rules in the WAF package.
	ActionMode param.Field[FirewallActionMode] `json:"action_mode"`
	// The sensitivity of the WAF package.
	Sensitivity param.Field[FirewallSensitivity] `json:"sensitivity"`
}

func (r ZoneFirewallWafPackageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFirewallWafPackageListParams struct {
	// The direction used to sort returned packages.
	Direction param.Field[ZoneFirewallWafPackageListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[ZoneFirewallWafPackageListParamsMatch] `query:"match"`
	// The name of the WAF package.
	Name param.Field[string] `query:"name"`
	// The field used to sort returned packages.
	Order param.Field[ZoneFirewallWafPackageListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of packages per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ZoneFirewallWafPackageListParams]'s query parameters as
// `url.Values`.
func (r ZoneFirewallWafPackageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned packages.
type ZoneFirewallWafPackageListParamsDirection string

const (
	ZoneFirewallWafPackageListParamsDirectionAsc  ZoneFirewallWafPackageListParamsDirection = "asc"
	ZoneFirewallWafPackageListParamsDirectionDesc ZoneFirewallWafPackageListParamsDirection = "desc"
)

func (r ZoneFirewallWafPackageListParamsDirection) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageListParamsDirectionAsc, ZoneFirewallWafPackageListParamsDirectionDesc:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type ZoneFirewallWafPackageListParamsMatch string

const (
	ZoneFirewallWafPackageListParamsMatchAny ZoneFirewallWafPackageListParamsMatch = "any"
	ZoneFirewallWafPackageListParamsMatchAll ZoneFirewallWafPackageListParamsMatch = "all"
)

func (r ZoneFirewallWafPackageListParamsMatch) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageListParamsMatchAny, ZoneFirewallWafPackageListParamsMatchAll:
		return true
	}
	return false
}

// The field used to sort returned packages.
type ZoneFirewallWafPackageListParamsOrder string

const (
	ZoneFirewallWafPackageListParamsOrderName ZoneFirewallWafPackageListParamsOrder = "name"
)

func (r ZoneFirewallWafPackageListParamsOrder) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageListParamsOrderName:
		return true
	}
	return false
}
