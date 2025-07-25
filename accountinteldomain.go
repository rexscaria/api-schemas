// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountIntelDomainService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountIntelDomainService] method instead.
type AccountIntelDomainService struct {
	Options []option.RequestOption
}

// NewAccountIntelDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountIntelDomainService(opts ...option.RequestOption) (r *AccountIntelDomainService) {
	r = &AccountIntelDomainService{}
	r.Options = opts
	return
}

// Gets security details and statistics about a domain.
func (r *AccountIntelDomainService) GetDetails(ctx context.Context, accountID string, query AccountIntelDomainGetDetailsParams, opts ...option.RequestOption) (res *AccountIntelDomainGetDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/domain", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Same as summary.
func (r *AccountIntelDomainService) ListMultipleDetails(ctx context.Context, accountID string, query AccountIntelDomainListMultipleDetailsParams, opts ...option.RequestOption) (res *AccountIntelDomainListMultipleDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/domain/bulk", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Additional information related to the host name.
type AdditionalInformation struct {
	// Suspected DGA malware family.
	SuspectedMalwareFamily string                    `json:"suspected_malware_family"`
	JSON                   additionalInformationJSON `json:"-"`
}

// additionalInformationJSON contains the JSON metadata for the struct
// [AdditionalInformation]
type additionalInformationJSON struct {
	SuspectedMalwareFamily apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AdditionalInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r additionalInformationJSON) RawJSON() string {
	return r.raw
}

// Application that the hostname belongs to.
type Application struct {
	ID   int64           `json:"id"`
	Name string          `json:"name"`
	JSON applicationJSON `json:"-"`
}

// applicationJSON contains the JSON metadata for the struct [Application]
type applicationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Application) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationJSON) RawJSON() string {
	return r.raw
}

type CategoryWithSuperCategory struct {
	ID              int64                         `json:"id"`
	Name            string                        `json:"name"`
	SuperCategoryID int64                         `json:"super_category_id"`
	JSON            categoryWithSuperCategoryJSON `json:"-"`
}

// categoryWithSuperCategoryJSON contains the JSON metadata for the struct
// [CategoryWithSuperCategory]
type categoryWithSuperCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CategoryWithSuperCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r categoryWithSuperCategoryJSON) RawJSON() string {
	return r.raw
}

// Current content categories.
type ContentCategory struct {
	ID              int64               `json:"id"`
	Name            string              `json:"name"`
	SuperCategoryID int64               `json:"super_category_id"`
	JSON            contentCategoryJSON `json:"-"`
}

// contentCategoryJSON contains the JSON metadata for the struct [ContentCategory]
type contentCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ContentCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentCategoryJSON) RawJSON() string {
	return r.raw
}

type AccountIntelDomainGetDetailsResponse struct {
	Errors   []AccountIntelDomainGetDetailsResponseError   `json:"errors,required"`
	Messages []AccountIntelDomainGetDetailsResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountIntelDomainGetDetailsResponseSuccess `json:"success,required"`
	Result  AccountIntelDomainGetDetailsResponseResult  `json:"result"`
	JSON    accountIntelDomainGetDetailsResponseJSON    `json:"-"`
}

// accountIntelDomainGetDetailsResponseJSON contains the JSON metadata for the
// struct [AccountIntelDomainGetDetailsResponse]
type accountIntelDomainGetDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainGetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelDomainGetDetailsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIntelDomainGetDetailsResponseError struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           AccountIntelDomainGetDetailsResponseErrorsSource `json:"source"`
	JSON             accountIntelDomainGetDetailsResponseErrorJSON    `json:"-"`
}

// accountIntelDomainGetDetailsResponseErrorJSON contains the JSON metadata for the
// struct [AccountIntelDomainGetDetailsResponseError]
type accountIntelDomainGetDetailsResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIntelDomainGetDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelDomainGetDetailsResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountIntelDomainGetDetailsResponseErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    accountIntelDomainGetDetailsResponseErrorsSourceJSON `json:"-"`
}

// accountIntelDomainGetDetailsResponseErrorsSourceJSON contains the JSON metadata
// for the struct [AccountIntelDomainGetDetailsResponseErrorsSource]
type accountIntelDomainGetDetailsResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainGetDetailsResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelDomainGetDetailsResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountIntelDomainGetDetailsResponseMessage struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           AccountIntelDomainGetDetailsResponseMessagesSource `json:"source"`
	JSON             accountIntelDomainGetDetailsResponseMessageJSON    `json:"-"`
}

// accountIntelDomainGetDetailsResponseMessageJSON contains the JSON metadata for
// the struct [AccountIntelDomainGetDetailsResponseMessage]
type accountIntelDomainGetDetailsResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIntelDomainGetDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelDomainGetDetailsResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountIntelDomainGetDetailsResponseMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    accountIntelDomainGetDetailsResponseMessagesSourceJSON `json:"-"`
}

// accountIntelDomainGetDetailsResponseMessagesSourceJSON contains the JSON
// metadata for the struct [AccountIntelDomainGetDetailsResponseMessagesSource]
type accountIntelDomainGetDetailsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainGetDetailsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelDomainGetDetailsResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountIntelDomainGetDetailsResponseSuccess bool

const (
	AccountIntelDomainGetDetailsResponseSuccessTrue AccountIntelDomainGetDetailsResponseSuccess = true
)

func (r AccountIntelDomainGetDetailsResponseSuccess) IsKnown() bool {
	switch r {
	case AccountIntelDomainGetDetailsResponseSuccessTrue:
		return true
	}
	return false
}

type AccountIntelDomainGetDetailsResponseResult struct {
	// Additional information related to the host name.
	AdditionalInformation AdditionalInformation `json:"additional_information"`
	// Application that the hostname belongs to.
	Application                Application                 `json:"application"`
	ContentCategories          []ContentCategory           `json:"content_categories"`
	Domain                     string                      `json:"domain"`
	InheritedContentCategories []CategoryWithSuperCategory `json:"inherited_content_categories"`
	// Domain from which `inherited_content_categories` and `inherited_risk_types` are
	// inherited, if applicable.
	InheritedFrom      string                      `json:"inherited_from"`
	InheritedRiskTypes []CategoryWithSuperCategory `json:"inherited_risk_types"`
	// Global Cloudflare 100k ranking for the last 30 days, if available for the
	// hostname. The top ranked domain is 1, the lowest ranked domain is 100,000.
	PopularityRank int64 `json:"popularity_rank"`
	// Specifies a list of references to one or more IP addresses or domain names that
	// the domain name currently resolves to.
	ResolvesToRefs []AccountIntelDomainGetDetailsResponseResultResolvesToRef `json:"resolves_to_refs"`
	// Hostname risk score, which is a value between 0 (lowest risk) to 1 (highest
	// risk).
	RiskScore float64                                        `json:"risk_score"`
	RiskTypes []CategoryWithSuperCategory                    `json:"risk_types"`
	JSON      accountIntelDomainGetDetailsResponseResultJSON `json:"-"`
}

// accountIntelDomainGetDetailsResponseResultJSON contains the JSON metadata for
// the struct [AccountIntelDomainGetDetailsResponseResult]
type accountIntelDomainGetDetailsResponseResultJSON struct {
	AdditionalInformation      apijson.Field
	Application                apijson.Field
	ContentCategories          apijson.Field
	Domain                     apijson.Field
	InheritedContentCategories apijson.Field
	InheritedFrom              apijson.Field
	InheritedRiskTypes         apijson.Field
	PopularityRank             apijson.Field
	ResolvesToRefs             apijson.Field
	RiskScore                  apijson.Field
	RiskTypes                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountIntelDomainGetDetailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelDomainGetDetailsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountIntelDomainGetDetailsResponseResultResolvesToRef struct {
	// STIX 2.1 identifier:
	// https://docs.oasis-open.org/cti/stix/v2.1/cs02/stix-v2.1-cs02.html#_64yvzeku5a5c.
	ID string `json:"id"`
	// IP address or domain name.
	Value string                                                      `json:"value"`
	JSON  accountIntelDomainGetDetailsResponseResultResolvesToRefJSON `json:"-"`
}

// accountIntelDomainGetDetailsResponseResultResolvesToRefJSON contains the JSON
// metadata for the struct
// [AccountIntelDomainGetDetailsResponseResultResolvesToRef]
type accountIntelDomainGetDetailsResponseResultResolvesToRefJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainGetDetailsResponseResultResolvesToRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelDomainGetDetailsResponseResultResolvesToRefJSON) RawJSON() string {
	return r.raw
}

type AccountIntelDomainListMultipleDetailsResponse struct {
	Errors   []AccountIntelDomainListMultipleDetailsResponseError   `json:"errors,required"`
	Messages []AccountIntelDomainListMultipleDetailsResponseMessage `json:"messages,required"`
	Result   []AccountIntelDomainListMultipleDetailsResponseResult  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    AccountIntelDomainListMultipleDetailsResponseSuccess `json:"success,required"`
	ResultInfo ResultInfoIntel                                      `json:"result_info"`
	JSON       accountIntelDomainListMultipleDetailsResponseJSON    `json:"-"`
}

// accountIntelDomainListMultipleDetailsResponseJSON contains the JSON metadata for
// the struct [AccountIntelDomainListMultipleDetailsResponse]
type accountIntelDomainListMultipleDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainListMultipleDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelDomainListMultipleDetailsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIntelDomainListMultipleDetailsResponseError struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AccountIntelDomainListMultipleDetailsResponseErrorsSource `json:"source"`
	JSON             accountIntelDomainListMultipleDetailsResponseErrorJSON    `json:"-"`
}

// accountIntelDomainListMultipleDetailsResponseErrorJSON contains the JSON
// metadata for the struct [AccountIntelDomainListMultipleDetailsResponseError]
type accountIntelDomainListMultipleDetailsResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIntelDomainListMultipleDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelDomainListMultipleDetailsResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountIntelDomainListMultipleDetailsResponseErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    accountIntelDomainListMultipleDetailsResponseErrorsSourceJSON `json:"-"`
}

// accountIntelDomainListMultipleDetailsResponseErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccountIntelDomainListMultipleDetailsResponseErrorsSource]
type accountIntelDomainListMultipleDetailsResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainListMultipleDetailsResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelDomainListMultipleDetailsResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountIntelDomainListMultipleDetailsResponseMessage struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           AccountIntelDomainListMultipleDetailsResponseMessagesSource `json:"source"`
	JSON             accountIntelDomainListMultipleDetailsResponseMessageJSON    `json:"-"`
}

// accountIntelDomainListMultipleDetailsResponseMessageJSON contains the JSON
// metadata for the struct [AccountIntelDomainListMultipleDetailsResponseMessage]
type accountIntelDomainListMultipleDetailsResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIntelDomainListMultipleDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelDomainListMultipleDetailsResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountIntelDomainListMultipleDetailsResponseMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    accountIntelDomainListMultipleDetailsResponseMessagesSourceJSON `json:"-"`
}

// accountIntelDomainListMultipleDetailsResponseMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccountIntelDomainListMultipleDetailsResponseMessagesSource]
type accountIntelDomainListMultipleDetailsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainListMultipleDetailsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelDomainListMultipleDetailsResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type AccountIntelDomainListMultipleDetailsResponseResult struct {
	// Additional information related to the host name.
	AdditionalInformation AdditionalInformation `json:"additional_information"`
	// Application that the hostname belongs to.
	Application                Application                 `json:"application"`
	ContentCategories          []ContentCategory           `json:"content_categories"`
	Domain                     string                      `json:"domain"`
	InheritedContentCategories []CategoryWithSuperCategory `json:"inherited_content_categories"`
	// Domain from which `inherited_content_categories` and `inherited_risk_types` are
	// inherited, if applicable.
	InheritedFrom      string                      `json:"inherited_from"`
	InheritedRiskTypes []CategoryWithSuperCategory `json:"inherited_risk_types"`
	// Global Cloudflare 100k ranking for the last 30 days, if available for the
	// hostname. The top ranked domain is 1, the lowest ranked domain is 100,000.
	PopularityRank int64 `json:"popularity_rank"`
	// Hostname risk score, which is a value between 0 (lowest risk) to 1 (highest
	// risk).
	RiskScore float64                                                 `json:"risk_score"`
	RiskTypes []CategoryWithSuperCategory                             `json:"risk_types"`
	JSON      accountIntelDomainListMultipleDetailsResponseResultJSON `json:"-"`
}

// accountIntelDomainListMultipleDetailsResponseResultJSON contains the JSON
// metadata for the struct [AccountIntelDomainListMultipleDetailsResponseResult]
type accountIntelDomainListMultipleDetailsResponseResultJSON struct {
	AdditionalInformation      apijson.Field
	Application                apijson.Field
	ContentCategories          apijson.Field
	Domain                     apijson.Field
	InheritedContentCategories apijson.Field
	InheritedFrom              apijson.Field
	InheritedRiskTypes         apijson.Field
	PopularityRank             apijson.Field
	RiskScore                  apijson.Field
	RiskTypes                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountIntelDomainListMultipleDetailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelDomainListMultipleDetailsResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountIntelDomainListMultipleDetailsResponseSuccess bool

const (
	AccountIntelDomainListMultipleDetailsResponseSuccessTrue AccountIntelDomainListMultipleDetailsResponseSuccess = true
)

func (r AccountIntelDomainListMultipleDetailsResponseSuccess) IsKnown() bool {
	switch r {
	case AccountIntelDomainListMultipleDetailsResponseSuccessTrue:
		return true
	}
	return false
}

type AccountIntelDomainGetDetailsParams struct {
	Domain param.Field[string] `query:"domain"`
}

// URLQuery serializes [AccountIntelDomainGetDetailsParams]'s query parameters as
// `url.Values`.
func (r AccountIntelDomainGetDetailsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountIntelDomainListMultipleDetailsParams struct {
	// Accepts multiple values like `?domain=cloudflare.com&domain=example.com`.
	Domain param.Field[[]string] `query:"domain"`
}

// URLQuery serializes [AccountIntelDomainListMultipleDetailsParams]'s query
// parameters as `url.Values`.
func (r AccountIntelDomainListMultipleDetailsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
