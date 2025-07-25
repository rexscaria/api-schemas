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

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// ZoneFirewallLockdownService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneFirewallLockdownService] method instead.
type ZoneFirewallLockdownService struct {
	Options []option.RequestOption
}

// NewZoneFirewallLockdownService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneFirewallLockdownService(opts ...option.RequestOption) (r *ZoneFirewallLockdownService) {
	r = &ZoneFirewallLockdownService{}
	r.Options = opts
	return
}

// Creates a new Zone Lockdown rule.
func (r *ZoneFirewallLockdownService) New(ctx context.Context, zoneID string, body ZoneFirewallLockdownNewParams, opts ...option.RequestOption) (res *FirewallZonelockdownResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the details of a Zone Lockdown rule.
func (r *ZoneFirewallLockdownService) Get(ctx context.Context, zoneID string, lockDownsID string, opts ...option.RequestOption) (res *FirewallZonelockdownResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if lockDownsID == "" {
		err = errors.New("missing required lock_downs_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneID, lockDownsID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing Zone Lockdown rule.
func (r *ZoneFirewallLockdownService) Update(ctx context.Context, zoneID string, lockDownsID string, body ZoneFirewallLockdownUpdateParams, opts ...option.RequestOption) (res *FirewallZonelockdownResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if lockDownsID == "" {
		err = errors.New("missing required lock_downs_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneID, lockDownsID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches Zone Lockdown rules. You can filter the results using several optional
// parameters.
func (r *ZoneFirewallLockdownService) List(ctx context.Context, zoneID string, query ZoneFirewallLockdownListParams, opts ...option.RequestOption) (res *ZoneFirewallLockdownListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing Zone Lockdown rule.
func (r *ZoneFirewallLockdownService) Delete(ctx context.Context, zoneID string, lockDownsID string, opts ...option.RequestOption) (res *ZoneFirewallLockdownDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if lockDownsID == "" {
		err = errors.New("missing required lock_downs_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneID, lockDownsID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type FirewallLockdownConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target FirewallLockdownConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                            `json:"value"`
	JSON  firewallLockdownConfigurationJSON `json:"-"`
	union FirewallLockdownConfigurationUnion
}

// firewallLockdownConfigurationJSON contains the JSON metadata for the struct
// [FirewallLockdownConfiguration]
type firewallLockdownConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r firewallLockdownConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r *FirewallLockdownConfiguration) UnmarshalJSON(data []byte) (err error) {
	*r = FirewallLockdownConfiguration{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FirewallLockdownConfigurationUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [FirewallLockdownConfigurationFirewallSchemasIPConfiguration],
// [FirewallLockdownConfigurationFirewallSchemasCidrConfiguration].
func (r FirewallLockdownConfiguration) AsUnion() FirewallLockdownConfigurationUnion {
	return r.union
}

// Union satisfied by [FirewallLockdownConfigurationFirewallSchemasIPConfiguration]
// or [FirewallLockdownConfigurationFirewallSchemasCidrConfiguration].
type FirewallLockdownConfigurationUnion interface {
	implementsFirewallLockdownConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallLockdownConfigurationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallLockdownConfigurationFirewallSchemasIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallLockdownConfigurationFirewallSchemasCidrConfiguration{}),
		},
	)
}

type FirewallLockdownConfigurationFirewallSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target FirewallLockdownConfigurationFirewallSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                          `json:"value"`
	JSON  firewallLockdownConfigurationFirewallSchemasIPConfigurationJSON `json:"-"`
}

// firewallLockdownConfigurationFirewallSchemasIPConfigurationJSON contains the
// JSON metadata for the struct
// [FirewallLockdownConfigurationFirewallSchemasIPConfiguration]
type firewallLockdownConfigurationFirewallSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownConfigurationFirewallSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallLockdownConfigurationFirewallSchemasIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallLockdownConfigurationFirewallSchemasIPConfiguration) implementsFirewallLockdownConfiguration() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type FirewallLockdownConfigurationFirewallSchemasIPConfigurationTarget string

const (
	FirewallLockdownConfigurationFirewallSchemasIPConfigurationTargetIP FirewallLockdownConfigurationFirewallSchemasIPConfigurationTarget = "ip"
)

func (r FirewallLockdownConfigurationFirewallSchemasIPConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallLockdownConfigurationFirewallSchemasIPConfigurationTargetIP:
		return true
	}
	return false
}

type FirewallLockdownConfigurationFirewallSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target FirewallLockdownConfigurationFirewallSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                            `json:"value"`
	JSON  firewallLockdownConfigurationFirewallSchemasCidrConfigurationJSON `json:"-"`
}

// firewallLockdownConfigurationFirewallSchemasCidrConfigurationJSON contains the
// JSON metadata for the struct
// [FirewallLockdownConfigurationFirewallSchemasCidrConfiguration]
type firewallLockdownConfigurationFirewallSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownConfigurationFirewallSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallLockdownConfigurationFirewallSchemasCidrConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallLockdownConfigurationFirewallSchemasCidrConfiguration) implementsFirewallLockdownConfiguration() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type FirewallLockdownConfigurationFirewallSchemasCidrConfigurationTarget string

const (
	FirewallLockdownConfigurationFirewallSchemasCidrConfigurationTargetIPRange FirewallLockdownConfigurationFirewallSchemasCidrConfigurationTarget = "ip_range"
)

func (r FirewallLockdownConfigurationFirewallSchemasCidrConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallLockdownConfigurationFirewallSchemasCidrConfigurationTargetIPRange:
		return true
	}
	return false
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type FirewallLockdownConfigurationTarget string

const (
	FirewallLockdownConfigurationTargetIP      FirewallLockdownConfigurationTarget = "ip"
	FirewallLockdownConfigurationTargetIPRange FirewallLockdownConfigurationTarget = "ip_range"
)

func (r FirewallLockdownConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallLockdownConfigurationTargetIP, FirewallLockdownConfigurationTargetIPRange:
		return true
	}
	return false
}

type FirewallLockdownConfigurationParam struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target param.Field[FirewallLockdownConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r FirewallLockdownConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallLockdownConfigurationParam) implementsFirewallLockdownConfigurationUnionParam() {}

// Satisfied by [FirewallLockdownConfigurationFirewallSchemasIPConfigurationParam],
// [FirewallLockdownConfigurationFirewallSchemasCidrConfigurationParam],
// [FirewallLockdownConfigurationParam].
type FirewallLockdownConfigurationUnionParam interface {
	implementsFirewallLockdownConfigurationUnionParam()
}

type FirewallLockdownConfigurationFirewallSchemasIPConfigurationParam struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target param.Field[FirewallLockdownConfigurationFirewallSchemasIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r FirewallLockdownConfigurationFirewallSchemasIPConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallLockdownConfigurationFirewallSchemasIPConfigurationParam) implementsFirewallLockdownConfigurationUnionParam() {
}

type FirewallLockdownConfigurationFirewallSchemasCidrConfigurationParam struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target param.Field[FirewallLockdownConfigurationFirewallSchemasCidrConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value param.Field[string] `json:"value"`
}

func (r FirewallLockdownConfigurationFirewallSchemasCidrConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallLockdownConfigurationFirewallSchemasCidrConfigurationParam) implementsFirewallLockdownConfigurationUnionParam() {
}

type FirewallZonelockdown struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations []FirewallLockdownConfiguration `json:"configurations,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// An informative summary of the rule.
	Description string `json:"description,required"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// When true, indicates that the rule is currently paused.
	Paused bool `json:"paused,required"`
	// The URLs to include in the rule definition. You can use wildcards. Each entered
	// URL will be escaped before use, which means you can only use simple wildcard
	// patterns.
	URLs []string                 `json:"urls,required"`
	JSON firewallZonelockdownJSON `json:"-"`
}

// firewallZonelockdownJSON contains the JSON metadata for the struct
// [FirewallZonelockdown]
type firewallZonelockdownJSON struct {
	ID             apijson.Field
	Configurations apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	ModifiedOn     apijson.Field
	Paused         apijson.Field
	URLs           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FirewallZonelockdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallZonelockdownJSON) RawJSON() string {
	return r.raw
}

type FirewallZonelockdownResponseSingle struct {
	Errors   []FirewallMessagesItem `json:"errors,required"`
	Messages []FirewallMessagesItem `json:"messages,required"`
	Result   FirewallZonelockdown   `json:"result,required"`
	// Defines whether the API call was successful.
	Success FirewallZonelockdownResponseSingleSuccess `json:"success,required"`
	JSON    firewallZonelockdownResponseSingleJSON    `json:"-"`
}

// firewallZonelockdownResponseSingleJSON contains the JSON metadata for the struct
// [FirewallZonelockdownResponseSingle]
type firewallZonelockdownResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallZonelockdownResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallZonelockdownResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type FirewallZonelockdownResponseSingleSuccess bool

const (
	FirewallZonelockdownResponseSingleSuccessTrue FirewallZonelockdownResponseSingleSuccess = true
)

func (r FirewallZonelockdownResponseSingleSuccess) IsKnown() bool {
	switch r {
	case FirewallZonelockdownResponseSingleSuccessTrue:
		return true
	}
	return false
}

type ZoneFirewallLockdownListResponse struct {
	Errors   []FirewallMessagesItem `json:"errors,required"`
	Messages []FirewallMessagesItem `json:"messages,required"`
	Result   []FirewallZonelockdown `json:"result,required,nullable"`
	// Defines whether the API call was successful.
	Success    ZoneFirewallLockdownListResponseSuccess    `json:"success,required"`
	ResultInfo ZoneFirewallLockdownListResponseResultInfo `json:"result_info"`
	JSON       zoneFirewallLockdownListResponseJSON       `json:"-"`
}

// zoneFirewallLockdownListResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallLockdownListResponse]
type zoneFirewallLockdownListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallLockdownListResponseJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ZoneFirewallLockdownListResponseSuccess bool

const (
	ZoneFirewallLockdownListResponseSuccessTrue ZoneFirewallLockdownListResponseSuccess = true
)

func (r ZoneFirewallLockdownListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneFirewallLockdownListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneFirewallLockdownListResponseResultInfo struct {
	// Defines the total number of results for the requested service.
	Count float64 `json:"count"`
	// Defines the current page within paginated list of results.
	Page float64 `json:"page"`
	// Defines the number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Defines the total results available without any search parameters.
	TotalCount float64                                        `json:"total_count"`
	JSON       zoneFirewallLockdownListResponseResultInfoJSON `json:"-"`
}

// zoneFirewallLockdownListResponseResultInfoJSON contains the JSON metadata for
// the struct [ZoneFirewallLockdownListResponseResultInfo]
type zoneFirewallLockdownListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallLockdownListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallLockdownDeleteResponse struct {
	Result ZoneFirewallLockdownDeleteResponseResult `json:"result"`
	JSON   zoneFirewallLockdownDeleteResponseJSON   `json:"-"`
}

// zoneFirewallLockdownDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallLockdownDeleteResponse]
type zoneFirewallLockdownDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallLockdownDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallLockdownDeleteResponseResult struct {
	// The unique identifier of the Zone Lockdown rule.
	ID   string                                       `json:"id"`
	JSON zoneFirewallLockdownDeleteResponseResultJSON `json:"-"`
}

// zoneFirewallLockdownDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneFirewallLockdownDeleteResponseResult]
type zoneFirewallLockdownDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallLockdownDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallLockdownNewParams struct {
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations param.Field[[]FirewallLockdownConfigurationUnionParam] `json:"configurations,required"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs param.Field[[]string] `json:"urls,required"`
	// An informative summary of the rule. This value is sanitized and any tags will be
	// removed.
	Description param.Field[string] `json:"description"`
	// When true, indicates that the rule is currently paused.
	Paused param.Field[bool] `json:"paused"`
	// The priority of the rule to control the processing order. A lower number
	// indicates higher priority. If not provided, any rules with a configured priority
	// will be processed before rules without a priority.
	Priority param.Field[float64] `json:"priority"`
}

func (r ZoneFirewallLockdownNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFirewallLockdownUpdateParams struct {
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations param.Field[[]FirewallLockdownConfigurationUnionParam] `json:"configurations,required"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs param.Field[[]string] `json:"urls,required"`
}

func (r ZoneFirewallLockdownUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFirewallLockdownListParams struct {
	// The timestamp of when the rule was created.
	CreatedOn param.Field[time.Time] `query:"created_on" format:"date-time"`
	// A string to search for in the description of existing rules.
	Description param.Field[string] `query:"description"`
	// A string to search for in the description of existing rules.
	DescriptionSearch param.Field[string] `query:"description_search"`
	// A single IP address to search for in existing rules.
	IP param.Field[string] `query:"ip"`
	// A single IP address range to search for in existing rules.
	IPRangeSearch param.Field[string] `query:"ip_range_search"`
	// A single IP address to search for in existing rules.
	IPSearch param.Field[string] `query:"ip_search"`
	// The timestamp of when the rule was last modified.
	ModifiedOn param.Field[time.Time] `query:"modified_on" format:"date-time"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
	// The priority of the rule to control the processing order. A lower number
	// indicates higher priority. If not provided, any rules with a configured priority
	// will be processed before rules without a priority.
	Priority param.Field[float64] `query:"priority"`
	// A single URI to search for in the list of URLs of existing rules.
	UriSearch param.Field[string] `query:"uri_search"`
}

// URLQuery serializes [ZoneFirewallLockdownListParams]'s query parameters as
// `url.Values`.
func (r ZoneFirewallLockdownListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
