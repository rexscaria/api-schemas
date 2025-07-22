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
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// AccountIntelService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountIntelService] method instead.
type AccountIntelService struct {
	Options             []option.RequestOption
	Asn                 *AccountIntelAsnService
	AttackSurfaceReport *AccountIntelAttackSurfaceReportService
	Domain              *AccountIntelDomainService
	IndicatorFeeds      *AccountIntelIndicatorFeedService
}

// NewAccountIntelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountIntelService(opts ...option.RequestOption) (r *AccountIntelService) {
	r = &AccountIntelService{}
	r.Options = opts
	r.Asn = NewAccountIntelAsnService(opts...)
	r.AttackSurfaceReport = NewAccountIntelAttackSurfaceReportService(opts...)
	r.Domain = NewAccountIntelDomainService(opts...)
	r.IndicatorFeeds = NewAccountIntelIndicatorFeedService(opts...)
	return
}

// Allows you to submit requests to change a domain’s category.
func (r *AccountIntelService) NewMiscategorization(ctx context.Context, accountID string, body AccountIntelNewMiscategorizationParams, opts ...option.RequestOption) (res *SingleResponseIntel, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/miscategorization", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets historical security threat and content categories currently and previously
// assigned to a domain.
func (r *AccountIntelService) GetDomainHistory(ctx context.Context, accountID string, query AccountIntelGetDomainHistoryParams, opts ...option.RequestOption) (res *AccountIntelGetDomainHistoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/domain-history", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Gets the geolocation, ASN, infrastructure type of the ASN, and any security
// threat categories of an IP address.
func (r *AccountIntelService) GetIPOverview(ctx context.Context, accountID string, query AccountIntelGetIPOverviewParams, opts ...option.RequestOption) (res *AccountIntelGetIPOverviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/ip", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Gets a list of all the domains that have resolved to a specific IP address.
func (r *AccountIntelService) GetPassiveDNS(ctx context.Context, accountID string, query AccountIntelGetPassiveDNSParams, opts ...option.RequestOption) (res *AccountIntelGetPassiveDNSResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/dns", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get WHOIS Record
func (r *AccountIntelService) GetWhoisRecord(ctx context.Context, accountID string, query AccountIntelGetWhoisRecordParams, opts ...option.RequestOption) (res *AccountIntelGetWhoisRecordResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/whois", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get IP Lists
func (r *AccountIntelService) ListIPLists(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountIntelListIPListsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/ip-list", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List sinkholes owned by this account
func (r *AccountIntelService) ListSinkholes(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountIntelListSinkholesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/sinkholes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CollectionResponse struct {
	Errors   []IntelMessage                `json:"errors,required"`
	Messages []IntelMessage                `json:"messages,required"`
	Result   CollectionResponseResultUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    CollectionResponseSuccess `json:"success,required"`
	ResultInfo ResultInfoIntel           `json:"result_info"`
	JSON       collectionResponseJSON    `json:"-"`
}

// collectionResponseJSON contains the JSON metadata for the struct
// [CollectionResponse]
type collectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r collectionResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [CollectionResponseResultArray] or [shared.UnionString].
type CollectionResponseResultUnion interface {
	ImplementsCollectionResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CollectionResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CollectionResponseResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CollectionResponseResultArray []interface{}

func (r CollectionResponseResultArray) ImplementsCollectionResponseResultUnion() {}

// Whether the API call was successful
type CollectionResponseSuccess bool

const (
	CollectionResponseSuccessTrue CollectionResponseSuccess = true
)

func (r CollectionResponseSuccess) IsKnown() bool {
	switch r {
	case CollectionResponseSuccessTrue:
		return true
	}
	return false
}

type CommonResponseIntel struct {
	Errors   []IntelMessage `json:"errors,required"`
	Messages []IntelMessage `json:"messages,required"`
	// Whether the API call was successful
	Success CommonResponseIntelSuccess `json:"success,required"`
	JSON    commonResponseIntelJSON    `json:"-"`
}

// commonResponseIntelJSON contains the JSON metadata for the struct
// [CommonResponseIntel]
type commonResponseIntelJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommonResponseIntel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commonResponseIntelJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CommonResponseIntelSuccess bool

const (
	CommonResponseIntelSuccessTrue CommonResponseIntelSuccess = true
)

func (r CommonResponseIntelSuccess) IsKnown() bool {
	switch r {
	case CommonResponseIntelSuccessTrue:
		return true
	}
	return false
}

type IntelMessage struct {
	Code    int64            `json:"code,required"`
	Message string           `json:"message,required"`
	JSON    intelMessageJSON `json:"-"`
}

// intelMessageJSON contains the JSON metadata for the struct [IntelMessage]
type intelMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelMessageJSON) RawJSON() string {
	return r.raw
}

type ResultInfoIntel struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64             `json:"total_count"`
	JSON       resultInfoIntelJSON `json:"-"`
}

// resultInfoIntelJSON contains the JSON metadata for the struct [ResultInfoIntel]
type resultInfoIntelJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultInfoIntel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultInfoIntelJSON) RawJSON() string {
	return r.raw
}

type SingleResponseIntel struct {
	Errors   []IntelMessage `json:"errors,required"`
	Messages []IntelMessage `json:"messages,required"`
	// Whether the API call was successful
	Success SingleResponseIntelSuccess `json:"success,required"`
	JSON    singleResponseIntelJSON    `json:"-"`
}

// singleResponseIntelJSON contains the JSON metadata for the struct
// [SingleResponseIntel]
type singleResponseIntelJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseIntel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseIntelJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SingleResponseIntelSuccess bool

const (
	SingleResponseIntelSuccessTrue SingleResponseIntelSuccess = true
)

func (r SingleResponseIntelSuccess) IsKnown() bool {
	switch r {
	case SingleResponseIntelSuccessTrue:
		return true
	}
	return false
}

type SinkholesMessage struct {
	Code    int64                `json:"code,required"`
	Message string               `json:"message,required"`
	JSON    sinkholesMessageJSON `json:"-"`
}

// sinkholesMessageJSON contains the JSON metadata for the struct
// [SinkholesMessage]
type sinkholesMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholesMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholesMessageJSON) RawJSON() string {
	return r.raw
}

type WhoisMessage struct {
	Code    int64            `json:"code,required"`
	Message string           `json:"message,required"`
	JSON    whoisMessageJSON `json:"-"`
}

// whoisMessageJSON contains the JSON metadata for the struct [WhoisMessage]
type whoisMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhoisMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whoisMessageJSON) RawJSON() string {
	return r.raw
}

type AccountIntelGetDomainHistoryResponse struct {
	Result []AccountIntelGetDomainHistoryResponseResult `json:"result"`
	JSON   accountIntelGetDomainHistoryResponseJSON     `json:"-"`
	CollectionResponse
}

// accountIntelGetDomainHistoryResponseJSON contains the JSON metadata for the
// struct [AccountIntelGetDomainHistoryResponse]
type accountIntelGetDomainHistoryResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelGetDomainHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelGetDomainHistoryResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIntelGetDomainHistoryResponseResult struct {
	Categorizations []AccountIntelGetDomainHistoryResponseResultCategorization `json:"categorizations"`
	Domain          string                                                     `json:"domain"`
	JSON            accountIntelGetDomainHistoryResponseResultJSON             `json:"-"`
}

// accountIntelGetDomainHistoryResponseResultJSON contains the JSON metadata for
// the struct [AccountIntelGetDomainHistoryResponseResult]
type accountIntelGetDomainHistoryResponseResultJSON struct {
	Categorizations apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountIntelGetDomainHistoryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelGetDomainHistoryResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountIntelGetDomainHistoryResponseResultCategorization struct {
	Categories interface{}                                                  `json:"categories"`
	End        time.Time                                                    `json:"end" format:"date"`
	Start      time.Time                                                    `json:"start" format:"date"`
	JSON       accountIntelGetDomainHistoryResponseResultCategorizationJSON `json:"-"`
}

// accountIntelGetDomainHistoryResponseResultCategorizationJSON contains the JSON
// metadata for the struct
// [AccountIntelGetDomainHistoryResponseResultCategorization]
type accountIntelGetDomainHistoryResponseResultCategorizationJSON struct {
	Categories  apijson.Field
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelGetDomainHistoryResponseResultCategorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelGetDomainHistoryResponseResultCategorizationJSON) RawJSON() string {
	return r.raw
}

type AccountIntelGetIPOverviewResponse struct {
	Result []AccountIntelGetIPOverviewResponseResult `json:"result"`
	JSON   accountIntelGetIPOverviewResponseJSON     `json:"-"`
	CollectionResponse
}

// accountIntelGetIPOverviewResponseJSON contains the JSON metadata for the struct
// [AccountIntelGetIPOverviewResponse]
type accountIntelGetIPOverviewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelGetIPOverviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelGetIPOverviewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIntelGetIPOverviewResponseResult struct {
	// Specifies a reference to the autonomous systems (AS) that the IP address belongs
	// to.
	BelongsToRef AccountIntelGetIPOverviewResponseResultBelongsToRef `json:"belongs_to_ref"`
	IP           string                                              `json:"ip" format:"ipv4"`
	RiskTypes    []AccountIntelGetIPOverviewResponseResultRiskType   `json:"risk_types"`
	JSON         accountIntelGetIPOverviewResponseResultJSON         `json:"-"`
}

// accountIntelGetIPOverviewResponseResultJSON contains the JSON metadata for the
// struct [AccountIntelGetIPOverviewResponseResult]
type accountIntelGetIPOverviewResponseResultJSON struct {
	BelongsToRef apijson.Field
	IP           apijson.Field
	RiskTypes    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountIntelGetIPOverviewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelGetIPOverviewResponseResultJSON) RawJSON() string {
	return r.raw
}

// Specifies a reference to the autonomous systems (AS) that the IP address belongs
// to.
type AccountIntelGetIPOverviewResponseResultBelongsToRef struct {
	ID          string `json:"id"`
	Country     string `json:"country"`
	Description string `json:"description"`
	// Infrastructure type of this ASN.
	Type  AccountIntelGetIPOverviewResponseResultBelongsToRefType `json:"type"`
	Value string                                                  `json:"value"`
	JSON  accountIntelGetIPOverviewResponseResultBelongsToRefJSON `json:"-"`
}

// accountIntelGetIPOverviewResponseResultBelongsToRefJSON contains the JSON
// metadata for the struct [AccountIntelGetIPOverviewResponseResultBelongsToRef]
type accountIntelGetIPOverviewResponseResultBelongsToRefJSON struct {
	ID          apijson.Field
	Country     apijson.Field
	Description apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelGetIPOverviewResponseResultBelongsToRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelGetIPOverviewResponseResultBelongsToRefJSON) RawJSON() string {
	return r.raw
}

// Infrastructure type of this ASN.
type AccountIntelGetIPOverviewResponseResultBelongsToRefType string

const (
	AccountIntelGetIPOverviewResponseResultBelongsToRefTypeHostingProvider AccountIntelGetIPOverviewResponseResultBelongsToRefType = "hosting_provider"
	AccountIntelGetIPOverviewResponseResultBelongsToRefTypeIsp             AccountIntelGetIPOverviewResponseResultBelongsToRefType = "isp"
	AccountIntelGetIPOverviewResponseResultBelongsToRefTypeOrganization    AccountIntelGetIPOverviewResponseResultBelongsToRefType = "organization"
)

func (r AccountIntelGetIPOverviewResponseResultBelongsToRefType) IsKnown() bool {
	switch r {
	case AccountIntelGetIPOverviewResponseResultBelongsToRefTypeHostingProvider, AccountIntelGetIPOverviewResponseResultBelongsToRefTypeIsp, AccountIntelGetIPOverviewResponseResultBelongsToRefTypeOrganization:
		return true
	}
	return false
}

type AccountIntelGetIPOverviewResponseResultRiskType struct {
	ID              float64                                             `json:"id"`
	Name            string                                              `json:"name"`
	SuperCategoryID float64                                             `json:"super_category_id"`
	JSON            accountIntelGetIPOverviewResponseResultRiskTypeJSON `json:"-"`
}

// accountIntelGetIPOverviewResponseResultRiskTypeJSON contains the JSON metadata
// for the struct [AccountIntelGetIPOverviewResponseResultRiskType]
type accountIntelGetIPOverviewResponseResultRiskTypeJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountIntelGetIPOverviewResponseResultRiskType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelGetIPOverviewResponseResultRiskTypeJSON) RawJSON() string {
	return r.raw
}

type AccountIntelGetPassiveDNSResponse struct {
	Result     AccountIntelGetPassiveDNSResponseResult `json:"result"`
	ResultInfo ResultInfoIntel                         `json:"result_info"`
	JSON       accountIntelGetPassiveDNSResponseJSON   `json:"-"`
	CommonResponseIntel
}

// accountIntelGetPassiveDNSResponseJSON contains the JSON metadata for the struct
// [AccountIntelGetPassiveDNSResponse]
type accountIntelGetPassiveDNSResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelGetPassiveDNSResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelGetPassiveDNSResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIntelGetPassiveDNSResponseResult struct {
	// Total results returned based on your search parameters.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Reverse DNS look-ups observed during the time period.
	ReverseRecords []AccountIntelGetPassiveDNSResponseResultReverseRecord `json:"reverse_records"`
	JSON           accountIntelGetPassiveDNSResponseResultJSON            `json:"-"`
}

// accountIntelGetPassiveDNSResponseResultJSON contains the JSON metadata for the
// struct [AccountIntelGetPassiveDNSResponseResult]
type accountIntelGetPassiveDNSResponseResultJSON struct {
	Count          apijson.Field
	Page           apijson.Field
	PerPage        apijson.Field
	ReverseRecords apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountIntelGetPassiveDNSResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelGetPassiveDNSResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountIntelGetPassiveDNSResponseResultReverseRecord struct {
	// First seen date of the DNS record during the time period.
	FirstSeen time.Time `json:"first_seen" format:"date"`
	// Hostname that the IP was observed resolving to.
	Hostname string `json:"hostname"`
	// Last seen date of the DNS record during the time period.
	LastSeen time.Time                                                `json:"last_seen" format:"date"`
	JSON     accountIntelGetPassiveDNSResponseResultReverseRecordJSON `json:"-"`
}

// accountIntelGetPassiveDNSResponseResultReverseRecordJSON contains the JSON
// metadata for the struct [AccountIntelGetPassiveDNSResponseResultReverseRecord]
type accountIntelGetPassiveDNSResponseResultReverseRecordJSON struct {
	FirstSeen   apijson.Field
	Hostname    apijson.Field
	LastSeen    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelGetPassiveDNSResponseResultReverseRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelGetPassiveDNSResponseResultReverseRecordJSON) RawJSON() string {
	return r.raw
}

type AccountIntelGetWhoisRecordResponse struct {
	Errors   []WhoisMessage `json:"errors,required"`
	Messages []WhoisMessage `json:"messages,required"`
	// Whether the API call was successful
	Success AccountIntelGetWhoisRecordResponseSuccess `json:"success,required"`
	Result  AccountIntelGetWhoisRecordResponseResult  `json:"result"`
	JSON    accountIntelGetWhoisRecordResponseJSON    `json:"-"`
}

// accountIntelGetWhoisRecordResponseJSON contains the JSON metadata for the struct
// [AccountIntelGetWhoisRecordResponse]
type accountIntelGetWhoisRecordResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelGetWhoisRecordResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelGetWhoisRecordResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountIntelGetWhoisRecordResponseSuccess bool

const (
	AccountIntelGetWhoisRecordResponseSuccessTrue AccountIntelGetWhoisRecordResponseSuccess = true
)

func (r AccountIntelGetWhoisRecordResponseSuccess) IsKnown() bool {
	switch r {
	case AccountIntelGetWhoisRecordResponseSuccessTrue:
		return true
	}
	return false
}

type AccountIntelGetWhoisRecordResponseResult struct {
	Domain                    string                                       `json:"domain,required"`
	Extension                 string                                       `json:"extension,required"`
	Found                     bool                                         `json:"found,required"`
	Nameservers               []string                                     `json:"nameservers,required"`
	Punycode                  string                                       `json:"punycode,required"`
	Registrant                string                                       `json:"registrant,required"`
	Registrar                 string                                       `json:"registrar,required"`
	ID                        string                                       `json:"id"`
	AdministrativeCity        string                                       `json:"administrative_city"`
	AdministrativeCountry     string                                       `json:"administrative_country"`
	AdministrativeEmail       string                                       `json:"administrative_email"`
	AdministrativeFax         string                                       `json:"administrative_fax"`
	AdministrativeFaxExt      string                                       `json:"administrative_fax_ext"`
	AdministrativeID          string                                       `json:"administrative_id"`
	AdministrativeName        string                                       `json:"administrative_name"`
	AdministrativeOrg         string                                       `json:"administrative_org"`
	AdministrativePhone       string                                       `json:"administrative_phone"`
	AdministrativePhoneExt    string                                       `json:"administrative_phone_ext"`
	AdministrativePostalCode  string                                       `json:"administrative_postal_code"`
	AdministrativeProvince    string                                       `json:"administrative_province"`
	AdministrativeReferralURL string                                       `json:"administrative_referral_url"`
	AdministrativeStreet      string                                       `json:"administrative_street"`
	BillingCity               string                                       `json:"billing_city"`
	BillingCountry            string                                       `json:"billing_country"`
	BillingEmail              string                                       `json:"billing_email"`
	BillingFax                string                                       `json:"billing_fax"`
	BillingFaxExt             string                                       `json:"billing_fax_ext"`
	BillingID                 string                                       `json:"billing_id"`
	BillingName               string                                       `json:"billing_name"`
	BillingOrg                string                                       `json:"billing_org"`
	BillingPhone              string                                       `json:"billing_phone"`
	BillingPhoneExt           string                                       `json:"billing_phone_ext"`
	BillingPostalCode         string                                       `json:"billing_postal_code"`
	BillingProvince           string                                       `json:"billing_province"`
	BillingReferralURL        string                                       `json:"billing_referral_url"`
	BillingStreet             string                                       `json:"billing_street"`
	CreatedDate               time.Time                                    `json:"created_date" format:"date-time"`
	CreatedDateRaw            string                                       `json:"created_date_raw"`
	Dnssec                    bool                                         `json:"dnssec"`
	ExpirationDate            time.Time                                    `json:"expiration_date" format:"date-time"`
	ExpirationDateRaw         string                                       `json:"expiration_date_raw"`
	RegistrantCity            string                                       `json:"registrant_city"`
	RegistrantCountry         string                                       `json:"registrant_country"`
	RegistrantEmail           string                                       `json:"registrant_email"`
	RegistrantFax             string                                       `json:"registrant_fax"`
	RegistrantFaxExt          string                                       `json:"registrant_fax_ext"`
	RegistrantID              string                                       `json:"registrant_id"`
	RegistrantName            string                                       `json:"registrant_name"`
	RegistrantOrg             string                                       `json:"registrant_org"`
	RegistrantPhone           string                                       `json:"registrant_phone"`
	RegistrantPhoneExt        string                                       `json:"registrant_phone_ext"`
	RegistrantPostalCode      string                                       `json:"registrant_postal_code"`
	RegistrantProvince        string                                       `json:"registrant_province"`
	RegistrantReferralURL     string                                       `json:"registrant_referral_url"`
	RegistrantStreet          string                                       `json:"registrant_street"`
	RegistrarCity             string                                       `json:"registrar_city"`
	RegistrarCountry          string                                       `json:"registrar_country"`
	RegistrarEmail            string                                       `json:"registrar_email"`
	RegistrarFax              string                                       `json:"registrar_fax"`
	RegistrarFaxExt           string                                       `json:"registrar_fax_ext"`
	RegistrarID               string                                       `json:"registrar_id"`
	RegistrarName             string                                       `json:"registrar_name"`
	RegistrarOrg              string                                       `json:"registrar_org"`
	RegistrarPhone            string                                       `json:"registrar_phone"`
	RegistrarPhoneExt         string                                       `json:"registrar_phone_ext"`
	RegistrarPostalCode       string                                       `json:"registrar_postal_code"`
	RegistrarProvince         string                                       `json:"registrar_province"`
	RegistrarReferralURL      string                                       `json:"registrar_referral_url"`
	RegistrarStreet           string                                       `json:"registrar_street"`
	Status                    []string                                     `json:"status"`
	TechnicalCity             string                                       `json:"technical_city"`
	TechnicalCountry          string                                       `json:"technical_country"`
	TechnicalEmail            string                                       `json:"technical_email"`
	TechnicalFax              string                                       `json:"technical_fax"`
	TechnicalFaxExt           string                                       `json:"technical_fax_ext"`
	TechnicalID               string                                       `json:"technical_id"`
	TechnicalName             string                                       `json:"technical_name"`
	TechnicalOrg              string                                       `json:"technical_org"`
	TechnicalPhone            string                                       `json:"technical_phone"`
	TechnicalPhoneExt         string                                       `json:"technical_phone_ext"`
	TechnicalPostalCode       string                                       `json:"technical_postal_code"`
	TechnicalProvince         string                                       `json:"technical_province"`
	TechnicalReferralURL      string                                       `json:"technical_referral_url"`
	TechnicalStreet           string                                       `json:"technical_street"`
	UpdatedDate               time.Time                                    `json:"updated_date" format:"date-time"`
	UpdatedDateRaw            string                                       `json:"updated_date_raw"`
	WhoisServer               string                                       `json:"whois_server"`
	JSON                      accountIntelGetWhoisRecordResponseResultJSON `json:"-"`
}

// accountIntelGetWhoisRecordResponseResultJSON contains the JSON metadata for the
// struct [AccountIntelGetWhoisRecordResponseResult]
type accountIntelGetWhoisRecordResponseResultJSON struct {
	Domain                    apijson.Field
	Extension                 apijson.Field
	Found                     apijson.Field
	Nameservers               apijson.Field
	Punycode                  apijson.Field
	Registrant                apijson.Field
	Registrar                 apijson.Field
	ID                        apijson.Field
	AdministrativeCity        apijson.Field
	AdministrativeCountry     apijson.Field
	AdministrativeEmail       apijson.Field
	AdministrativeFax         apijson.Field
	AdministrativeFaxExt      apijson.Field
	AdministrativeID          apijson.Field
	AdministrativeName        apijson.Field
	AdministrativeOrg         apijson.Field
	AdministrativePhone       apijson.Field
	AdministrativePhoneExt    apijson.Field
	AdministrativePostalCode  apijson.Field
	AdministrativeProvince    apijson.Field
	AdministrativeReferralURL apijson.Field
	AdministrativeStreet      apijson.Field
	BillingCity               apijson.Field
	BillingCountry            apijson.Field
	BillingEmail              apijson.Field
	BillingFax                apijson.Field
	BillingFaxExt             apijson.Field
	BillingID                 apijson.Field
	BillingName               apijson.Field
	BillingOrg                apijson.Field
	BillingPhone              apijson.Field
	BillingPhoneExt           apijson.Field
	BillingPostalCode         apijson.Field
	BillingProvince           apijson.Field
	BillingReferralURL        apijson.Field
	BillingStreet             apijson.Field
	CreatedDate               apijson.Field
	CreatedDateRaw            apijson.Field
	Dnssec                    apijson.Field
	ExpirationDate            apijson.Field
	ExpirationDateRaw         apijson.Field
	RegistrantCity            apijson.Field
	RegistrantCountry         apijson.Field
	RegistrantEmail           apijson.Field
	RegistrantFax             apijson.Field
	RegistrantFaxExt          apijson.Field
	RegistrantID              apijson.Field
	RegistrantName            apijson.Field
	RegistrantOrg             apijson.Field
	RegistrantPhone           apijson.Field
	RegistrantPhoneExt        apijson.Field
	RegistrantPostalCode      apijson.Field
	RegistrantProvince        apijson.Field
	RegistrantReferralURL     apijson.Field
	RegistrantStreet          apijson.Field
	RegistrarCity             apijson.Field
	RegistrarCountry          apijson.Field
	RegistrarEmail            apijson.Field
	RegistrarFax              apijson.Field
	RegistrarFaxExt           apijson.Field
	RegistrarID               apijson.Field
	RegistrarName             apijson.Field
	RegistrarOrg              apijson.Field
	RegistrarPhone            apijson.Field
	RegistrarPhoneExt         apijson.Field
	RegistrarPostalCode       apijson.Field
	RegistrarProvince         apijson.Field
	RegistrarReferralURL      apijson.Field
	RegistrarStreet           apijson.Field
	Status                    apijson.Field
	TechnicalCity             apijson.Field
	TechnicalCountry          apijson.Field
	TechnicalEmail            apijson.Field
	TechnicalFax              apijson.Field
	TechnicalFaxExt           apijson.Field
	TechnicalID               apijson.Field
	TechnicalName             apijson.Field
	TechnicalOrg              apijson.Field
	TechnicalPhone            apijson.Field
	TechnicalPhoneExt         apijson.Field
	TechnicalPostalCode       apijson.Field
	TechnicalProvince         apijson.Field
	TechnicalReferralURL      apijson.Field
	TechnicalStreet           apijson.Field
	UpdatedDate               apijson.Field
	UpdatedDateRaw            apijson.Field
	WhoisServer               apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AccountIntelGetWhoisRecordResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelGetWhoisRecordResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountIntelListIPListsResponse struct {
	Result []AccountIntelListIPListsResponseResult `json:"result"`
	JSON   accountIntelListIPListsResponseJSON     `json:"-"`
	CollectionResponse
}

// accountIntelListIPListsResponseJSON contains the JSON metadata for the struct
// [AccountIntelListIPListsResponse]
type accountIntelListIPListsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelListIPListsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelListIPListsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIntelListIPListsResponseResult struct {
	ID          int64                                     `json:"id"`
	Description string                                    `json:"description"`
	Name        string                                    `json:"name"`
	JSON        accountIntelListIPListsResponseResultJSON `json:"-"`
}

// accountIntelListIPListsResponseResultJSON contains the JSON metadata for the
// struct [AccountIntelListIPListsResponseResult]
type accountIntelListIPListsResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelListIPListsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelListIPListsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountIntelListSinkholesResponse struct {
	Errors   []SinkholesMessage `json:"errors,required"`
	Messages []SinkholesMessage `json:"messages,required"`
	// Whether the API call was successful
	Success AccountIntelListSinkholesResponseSuccess  `json:"success,required"`
	Result  []AccountIntelListSinkholesResponseResult `json:"result"`
	JSON    accountIntelListSinkholesResponseJSON     `json:"-"`
}

// accountIntelListSinkholesResponseJSON contains the JSON metadata for the struct
// [AccountIntelListSinkholesResponse]
type accountIntelListSinkholesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelListSinkholesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelListSinkholesResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountIntelListSinkholesResponseSuccess bool

const (
	AccountIntelListSinkholesResponseSuccessTrue AccountIntelListSinkholesResponseSuccess = true
)

func (r AccountIntelListSinkholesResponseSuccess) IsKnown() bool {
	switch r {
	case AccountIntelListSinkholesResponseSuccessTrue:
		return true
	}
	return false
}

type AccountIntelListSinkholesResponseResult struct {
	// The unique identifier for the sinkhole
	ID int64 `json:"id"`
	// The account tag that owns this sinkhole
	AccountTag string `json:"account_tag"`
	// The date and time when the sinkhole was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The date and time when the sinkhole was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the sinkhole
	Name string `json:"name"`
	// The name of the R2 bucket to store results
	R2Bucket string `json:"r2_bucket"`
	// The id of the R2 instance
	R2ID string                                      `json:"r2_id"`
	JSON accountIntelListSinkholesResponseResultJSON `json:"-"`
}

// accountIntelListSinkholesResponseResultJSON contains the JSON metadata for the
// struct [AccountIntelListSinkholesResponseResult]
type accountIntelListSinkholesResponseResultJSON struct {
	ID          apijson.Field
	AccountTag  apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	R2Bucket    apijson.Field
	R2ID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelListSinkholesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelListSinkholesResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountIntelNewMiscategorizationParams struct {
	// Content category IDs to add.
	ContentAdds param.Field[[]int64] `json:"content_adds"`
	// Content category IDs to remove.
	ContentRemoves param.Field[[]int64]                                             `json:"content_removes"`
	IndicatorType  param.Field[AccountIntelNewMiscategorizationParamsIndicatorType] `json:"indicator_type"`
	// Provide only if indicator_type is `ipv4` or `ipv6`.
	IP param.Field[string] `json:"ip"`
	// Security category IDs to add.
	SecurityAdds param.Field[[]int64] `json:"security_adds"`
	// Security category IDs to remove.
	SecurityRemoves param.Field[[]int64] `json:"security_removes"`
	// Provide only if indicator_type is `domain` or `url`. Example if indicator_type
	// is `domain`: `example.com`. Example if indicator_type is `url`:
	// `https://example.com/news/`.
	URL param.Field[string] `json:"url"`
}

func (r AccountIntelNewMiscategorizationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountIntelNewMiscategorizationParamsIndicatorType string

const (
	AccountIntelNewMiscategorizationParamsIndicatorTypeDomain AccountIntelNewMiscategorizationParamsIndicatorType = "domain"
	AccountIntelNewMiscategorizationParamsIndicatorTypeIpv4   AccountIntelNewMiscategorizationParamsIndicatorType = "ipv4"
	AccountIntelNewMiscategorizationParamsIndicatorTypeIpv6   AccountIntelNewMiscategorizationParamsIndicatorType = "ipv6"
	AccountIntelNewMiscategorizationParamsIndicatorTypeURL    AccountIntelNewMiscategorizationParamsIndicatorType = "url"
)

func (r AccountIntelNewMiscategorizationParamsIndicatorType) IsKnown() bool {
	switch r {
	case AccountIntelNewMiscategorizationParamsIndicatorTypeDomain, AccountIntelNewMiscategorizationParamsIndicatorTypeIpv4, AccountIntelNewMiscategorizationParamsIndicatorTypeIpv6, AccountIntelNewMiscategorizationParamsIndicatorTypeURL:
		return true
	}
	return false
}

type AccountIntelGetDomainHistoryParams struct {
	Domain param.Field[interface{}] `query:"domain"`
}

// URLQuery serializes [AccountIntelGetDomainHistoryParams]'s query parameters as
// `url.Values`.
func (r AccountIntelGetDomainHistoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountIntelGetIPOverviewParams struct {
	Ipv4 param.Field[string] `query:"ipv4"`
	Ipv6 param.Field[string] `query:"ipv6"`
}

// URLQuery serializes [AccountIntelGetIPOverviewParams]'s query parameters as
// `url.Values`.
func (r AccountIntelGetIPOverviewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountIntelGetPassiveDNSParams struct {
	Ipv4 param.Field[string] `query:"ipv4"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage        param.Field[float64]                                       `query:"per_page"`
	StartEndParams param.Field[AccountIntelGetPassiveDNSParamsStartEndParams] `query:"start_end_params"`
}

// URLQuery serializes [AccountIntelGetPassiveDNSParams]'s query parameters as
// `url.Values`.
func (r AccountIntelGetPassiveDNSParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountIntelGetPassiveDNSParamsStartEndParams struct {
	// Defaults to the current date.
	End param.Field[time.Time] `query:"end" format:"date"`
	// Defaults to 30 days before the end parameter value.
	Start param.Field[time.Time] `query:"start" format:"date"`
}

// URLQuery serializes [AccountIntelGetPassiveDNSParamsStartEndParams]'s query
// parameters as `url.Values`.
func (r AccountIntelGetPassiveDNSParamsStartEndParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountIntelGetWhoisRecordParams struct {
	Domain param.Field[string] `query:"domain"`
}

// URLQuery serializes [AccountIntelGetWhoisRecordParams]'s query parameters as
// `url.Values`.
func (r AccountIntelGetWhoisRecordParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
