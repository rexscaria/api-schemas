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

// ZoneEmailRoutingDNSService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneEmailRoutingDNSService] method instead.
type ZoneEmailRoutingDNSService struct {
	Options []option.RequestOption
}

// NewZoneEmailRoutingDNSService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneEmailRoutingDNSService(opts ...option.RequestOption) (r *ZoneEmailRoutingDNSService) {
	r = &ZoneEmailRoutingDNSService{}
	r.Options = opts
	return
}

// Show the DNS records needed to configure your Email Routing zone.
func (r *ZoneEmailRoutingDNSService) Get(ctx context.Context, zoneID string, query ZoneEmailRoutingDNSGetParams, opts ...option.RequestOption) (res *ZoneEmailRoutingDNSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Disable your Email Routing zone. Also removes additional MX records previously
// required for Email Routing to work.
func (r *ZoneEmailRoutingDNSService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneEmailRoutingDNSDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Enable you Email Routing zone. Add and lock the necessary MX and SPF records.
func (r *ZoneEmailRoutingDNSService) Enable(ctx context.Context, zoneID string, body ZoneEmailRoutingDNSEnableParams, opts ...option.RequestOption) (res *EmailEmailSettingsResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unlock MX Records previously locked by Email Routing.
func (r *ZoneEmailRoutingDNSService) Unlock(ctx context.Context, zoneID string, body ZoneEmailRoutingDNSUnlockParams, opts ...option.RequestOption) (res *EmailEmailSettingsResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type EmailAPIResponseSingle struct {
	Errors   []EmailMessagesItem `json:"errors,required"`
	Messages []EmailMessagesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success EmailAPIResponseSingleSuccess `json:"success,required"`
	JSON    emailAPIResponseSingleJSON    `json:"-"`
}

// emailAPIResponseSingleJSON contains the JSON metadata for the struct
// [EmailAPIResponseSingle]
type emailAPIResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailAPIResponseSingleJSON) RawJSON() string {
	return r.raw
}

func (r EmailAPIResponseSingle) implementsZoneEmailRoutingDNSDeleteResponse() {}

// Whether the API call was successful.
type EmailAPIResponseSingleSuccess bool

const (
	EmailAPIResponseSingleSuccessTrue EmailAPIResponseSingleSuccess = true
)

func (r EmailAPIResponseSingleSuccess) IsKnown() bool {
	switch r {
	case EmailAPIResponseSingleSuccessTrue:
		return true
	}
	return false
}

// List of records needed to enable an Email Routing zone.
type EmailDNSRecord struct {
	// DNS record content.
	Content string `json:"content"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name"`
	// Required for MX, SRV and URI records. Unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
	// for 'automatic'.
	Ttl EmailDNSRecordTtl `json:"ttl"`
	// DNS record type.
	Type EmailDNSRecordType `json:"type"`
	JSON emailDNSRecordJSON `json:"-"`
}

// emailDNSRecordJSON contains the JSON metadata for the struct [EmailDNSRecord]
type emailDNSRecordJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Ttl         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailDNSRecordJSON) RawJSON() string {
	return r.raw
}

// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
// for 'automatic'.
type EmailDNSRecordTtl float64

const (
	EmailDNSRecordTtl1 EmailDNSRecordTtl = 1
)

func (r EmailDNSRecordTtl) IsKnown() bool {
	switch r {
	case EmailDNSRecordTtl1:
		return true
	}
	return false
}

// DNS record type.
type EmailDNSRecordType string

const (
	EmailDNSRecordTypeA      EmailDNSRecordType = "A"
	EmailDNSRecordTypeAaaa   EmailDNSRecordType = "AAAA"
	EmailDNSRecordTypeCname  EmailDNSRecordType = "CNAME"
	EmailDNSRecordTypeHTTPS  EmailDNSRecordType = "HTTPS"
	EmailDNSRecordTypeTxt    EmailDNSRecordType = "TXT"
	EmailDNSRecordTypeSrv    EmailDNSRecordType = "SRV"
	EmailDNSRecordTypeLoc    EmailDNSRecordType = "LOC"
	EmailDNSRecordTypeMx     EmailDNSRecordType = "MX"
	EmailDNSRecordTypeNs     EmailDNSRecordType = "NS"
	EmailDNSRecordTypeCert   EmailDNSRecordType = "CERT"
	EmailDNSRecordTypeDnskey EmailDNSRecordType = "DNSKEY"
	EmailDNSRecordTypeDs     EmailDNSRecordType = "DS"
	EmailDNSRecordTypeNaptr  EmailDNSRecordType = "NAPTR"
	EmailDNSRecordTypeSmimea EmailDNSRecordType = "SMIMEA"
	EmailDNSRecordTypeSshfp  EmailDNSRecordType = "SSHFP"
	EmailDNSRecordTypeSvcb   EmailDNSRecordType = "SVCB"
	EmailDNSRecordTypeTlsa   EmailDNSRecordType = "TLSA"
	EmailDNSRecordTypeUri    EmailDNSRecordType = "URI"
)

func (r EmailDNSRecordType) IsKnown() bool {
	switch r {
	case EmailDNSRecordTypeA, EmailDNSRecordTypeAaaa, EmailDNSRecordTypeCname, EmailDNSRecordTypeHTTPS, EmailDNSRecordTypeTxt, EmailDNSRecordTypeSrv, EmailDNSRecordTypeLoc, EmailDNSRecordTypeMx, EmailDNSRecordTypeNs, EmailDNSRecordTypeCert, EmailDNSRecordTypeDnskey, EmailDNSRecordTypeDs, EmailDNSRecordTypeNaptr, EmailDNSRecordTypeSmimea, EmailDNSRecordTypeSshfp, EmailDNSRecordTypeSvcb, EmailDNSRecordTypeTlsa, EmailDNSRecordTypeUri:
		return true
	}
	return false
}

type EmailDNSSettingsResponseCollection struct {
	Errors   []EmailMessagesItem `json:"errors,required"`
	Messages []EmailMessagesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    EmailDNSSettingsResponseCollectionSuccess    `json:"success,required"`
	Result     []EmailDNSRecord                             `json:"result"`
	ResultInfo EmailDNSSettingsResponseCollectionResultInfo `json:"result_info"`
	JSON       emailDNSSettingsResponseCollectionJSON       `json:"-"`
}

// emailDNSSettingsResponseCollectionJSON contains the JSON metadata for the struct
// [EmailDNSSettingsResponseCollection]
type emailDNSSettingsResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailDNSSettingsResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailDNSSettingsResponseCollectionJSON) RawJSON() string {
	return r.raw
}

func (r EmailDNSSettingsResponseCollection) implementsZoneEmailRoutingDNSGetResponse() {}

func (r EmailDNSSettingsResponseCollection) implementsZoneEmailRoutingDNSDeleteResponse() {}

// Whether the API call was successful.
type EmailDNSSettingsResponseCollectionSuccess bool

const (
	EmailDNSSettingsResponseCollectionSuccessTrue EmailDNSSettingsResponseCollectionSuccess = true
)

func (r EmailDNSSettingsResponseCollectionSuccess) IsKnown() bool {
	switch r {
	case EmailDNSSettingsResponseCollectionSuccessTrue:
		return true
	}
	return false
}

type EmailDNSSettingsResponseCollectionResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                          `json:"total_count"`
	JSON       emailDNSSettingsResponseCollectionResultInfoJSON `json:"-"`
}

// emailDNSSettingsResponseCollectionResultInfoJSON contains the JSON metadata for
// the struct [EmailDNSSettingsResponseCollectionResultInfo]
type emailDNSSettingsResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailDNSSettingsResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailDNSSettingsResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type EmailEmailSettingDNSRequestBodyParam struct {
	// Domain of your zone.
	Name param.Field[string] `json:"name,required"`
}

func (r EmailEmailSettingDNSRequestBodyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmailMessagesItem struct {
	Code             int64                   `json:"code,required"`
	Message          string                  `json:"message,required"`
	DocumentationURL string                  `json:"documentation_url"`
	Source           EmailMessagesItemSource `json:"source"`
	JSON             emailMessagesItemJSON   `json:"-"`
}

// emailMessagesItemJSON contains the JSON metadata for the struct
// [EmailMessagesItem]
type emailMessagesItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EmailMessagesItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailMessagesItemJSON) RawJSON() string {
	return r.raw
}

type EmailMessagesItemSource struct {
	Pointer string                      `json:"pointer"`
	JSON    emailMessagesItemSourceJSON `json:"-"`
}

// emailMessagesItemSourceJSON contains the JSON metadata for the struct
// [EmailMessagesItemSource]
type emailMessagesItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailMessagesItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailMessagesItemSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneEmailRoutingDNSGetResponse struct {
	// This field can have the runtime type of [[]EmailMessagesItem].
	Errors interface{} `json:"errors,required"`
	// This field can have the runtime type of [[]EmailMessagesItem].
	Messages interface{} `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneEmailRoutingDNSGetResponseSuccess `json:"success,required"`
	// This field can have the runtime type of
	// [ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResult],
	// [[]EmailDNSRecord].
	Result interface{} `json:"result"`
	// This field can have the runtime type of
	// [ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultInfo],
	// [EmailDNSSettingsResponseCollectionResultInfo].
	ResultInfo interface{}                        `json:"result_info"`
	JSON       zoneEmailRoutingDNSGetResponseJSON `json:"-"`
	union      ZoneEmailRoutingDNSGetResponseUnion
}

// zoneEmailRoutingDNSGetResponseJSON contains the JSON metadata for the struct
// [ZoneEmailRoutingDNSGetResponse]
type zoneEmailRoutingDNSGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r zoneEmailRoutingDNSGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ZoneEmailRoutingDNSGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ZoneEmailRoutingDNSGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZoneEmailRoutingDNSGetResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponse],
// [EmailDNSSettingsResponseCollection].
func (r ZoneEmailRoutingDNSGetResponse) AsUnion() ZoneEmailRoutingDNSGetResponseUnion {
	return r.union
}

// Union satisfied by
// [ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponse] or
// [EmailDNSSettingsResponseCollection].
type ZoneEmailRoutingDNSGetResponseUnion interface {
	implementsZoneEmailRoutingDNSGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneEmailRoutingDNSGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmailDNSSettingsResponseCollection{}),
		},
	)
}

type ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponse struct {
	Errors   []EmailMessagesItem `json:"errors,required"`
	Messages []EmailMessagesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseSuccess    `json:"success,required"`
	Result     ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResult     `json:"result"`
	ResultInfo ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultInfo `json:"result_info"`
	JSON       zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseJSON       `json:"-"`
}

// zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseJSON contains the
// JSON metadata for the struct
// [ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponse]
type zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseJSON) RawJSON() string {
	return r.raw
}

func (r ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponse) implementsZoneEmailRoutingDNSGetResponse() {
}

// Whether the API call was successful.
type ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseSuccess bool

const (
	ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseSuccessTrue ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseSuccess = true
)

func (r ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResult struct {
	Errors []ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultError `json:"errors"`
	Record []EmailDNSRecord                                                             `json:"record"`
	JSON   zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultJSON    `json:"-"`
}

// zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResult]
type zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultJSON struct {
	Errors      apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultError struct {
	Code string `json:"code"`
	// List of records needed to enable an Email Routing zone.
	Missing EmailDNSRecord                                                                 `json:"missing"`
	JSON    zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultErrorJSON `json:"-"`
}

// zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultErrorJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultError]
type zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultErrorJSON struct {
	Code        apijson.Field
	Missing     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultErrorJSON) RawJSON() string {
	return r.raw
}

type ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                                       `json:"total_count"`
	JSON       zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultInfoJSON `json:"-"`
}

// zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultInfo]
type zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEmailRoutingDNSGetResponseEmailEmailRoutingDNSQueryResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneEmailRoutingDNSGetResponseSuccess bool

const (
	ZoneEmailRoutingDNSGetResponseSuccessTrue ZoneEmailRoutingDNSGetResponseSuccess = true
)

func (r ZoneEmailRoutingDNSGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneEmailRoutingDNSGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneEmailRoutingDNSDeleteResponse struct {
	// This field can have the runtime type of [[]EmailMessagesItem].
	Errors interface{} `json:"errors,required"`
	// This field can have the runtime type of [[]EmailMessagesItem].
	Messages interface{} `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneEmailRoutingDNSDeleteResponseSuccess `json:"success,required"`
	// This field can have the runtime type of [[]EmailDNSRecord].
	Result interface{} `json:"result"`
	// This field can have the runtime type of
	// [EmailDNSSettingsResponseCollectionResultInfo].
	ResultInfo interface{}                           `json:"result_info"`
	JSON       zoneEmailRoutingDNSDeleteResponseJSON `json:"-"`
	union      ZoneEmailRoutingDNSDeleteResponseUnion
}

// zoneEmailRoutingDNSDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneEmailRoutingDNSDeleteResponse]
type zoneEmailRoutingDNSDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r zoneEmailRoutingDNSDeleteResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ZoneEmailRoutingDNSDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ZoneEmailRoutingDNSDeleteResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZoneEmailRoutingDNSDeleteResponseUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [EmailAPIResponseSingle],
// [EmailDNSSettingsResponseCollection].
func (r ZoneEmailRoutingDNSDeleteResponse) AsUnion() ZoneEmailRoutingDNSDeleteResponseUnion {
	return r.union
}

// Union satisfied by [EmailAPIResponseSingle] or
// [EmailDNSSettingsResponseCollection].
type ZoneEmailRoutingDNSDeleteResponseUnion interface {
	implementsZoneEmailRoutingDNSDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneEmailRoutingDNSDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmailAPIResponseSingle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmailDNSSettingsResponseCollection{}),
		},
	)
}

// Whether the API call was successful.
type ZoneEmailRoutingDNSDeleteResponseSuccess bool

const (
	ZoneEmailRoutingDNSDeleteResponseSuccessTrue ZoneEmailRoutingDNSDeleteResponseSuccess = true
)

func (r ZoneEmailRoutingDNSDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneEmailRoutingDNSDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneEmailRoutingDNSGetParams struct {
	// Domain of your zone.
	Subdomain param.Field[string] `query:"subdomain"`
}

// URLQuery serializes [ZoneEmailRoutingDNSGetParams]'s query parameters as
// `url.Values`.
func (r ZoneEmailRoutingDNSGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneEmailRoutingDNSEnableParams struct {
	EmailEmailSettingDNSRequestBody EmailEmailSettingDNSRequestBodyParam `json:"email_email_setting_dns_request_body,required"`
}

func (r ZoneEmailRoutingDNSEnableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EmailEmailSettingDNSRequestBody)
}

type ZoneEmailRoutingDNSUnlockParams struct {
	EmailEmailSettingDNSRequestBody EmailEmailSettingDNSRequestBodyParam `json:"email_email_setting_dns_request_body,required"`
}

func (r ZoneEmailRoutingDNSUnlockParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EmailEmailSettingDNSRequestBody)
}
