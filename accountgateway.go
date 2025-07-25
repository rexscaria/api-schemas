// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountGatewayService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountGatewayService] method instead.
type AccountGatewayService struct {
	Options          []option.RequestOption
	AuditSSHSettings *AccountGatewayAuditSSHSettingService
	Certificates     *AccountGatewayCertificateService
	Configuration    *AccountGatewayConfigurationService
	Lists            *AccountGatewayListService
	Locations        *AccountGatewayLocationService
	Logging          *AccountGatewayLoggingService
	ProxyEndpoints   *AccountGatewayProxyEndpointService
	Rules            *AccountGatewayRuleService
}

// NewAccountGatewayService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountGatewayService(opts ...option.RequestOption) (r *AccountGatewayService) {
	r = &AccountGatewayService{}
	r.Options = opts
	r.AuditSSHSettings = NewAccountGatewayAuditSSHSettingService(opts...)
	r.Certificates = NewAccountGatewayCertificateService(opts...)
	r.Configuration = NewAccountGatewayConfigurationService(opts...)
	r.Lists = NewAccountGatewayListService(opts...)
	r.Locations = NewAccountGatewayLocationService(opts...)
	r.Logging = NewAccountGatewayLoggingService(opts...)
	r.ProxyEndpoints = NewAccountGatewayProxyEndpointService(opts...)
	r.Rules = NewAccountGatewayRuleService(opts...)
	return
}

// Creates a Zero Trust account with an existing Cloudflare account.
func (r *AccountGatewayService) New(ctx context.Context, accountID string, opts ...option.RequestOption) (res *GatewayAccount, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Gets information about the current Zero Trust account.
func (r *AccountGatewayService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *GatewayAccount, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches all application and application type mappings.
func (r *AccountGatewayService) ListAppTypes(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountGatewayListAppTypesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/app_types", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a list of all categories.
func (r *AccountGatewayService) ListCategories(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountGatewayListCategoriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/categories", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Which account types are allowed to create policies based on this category.
// `blocked` categories are blocked unconditionally for all accounts.
// `removalPending` categories can be removed from policies but not added.
// `noBlock` categories cannot be blocked.
type Class string

const (
	ClassFree           Class = "free"
	ClassPremium        Class = "premium"
	ClassBlocked        Class = "blocked"
	ClassRemovalPending Class = "removalPending"
	ClassNoBlock        Class = "noBlock"
)

func (r Class) IsKnown() bool {
	switch r {
	case ClassFree, ClassPremium, ClassBlocked, ClassRemovalPending, ClassNoBlock:
		return true
	}
	return false
}

type GatewayAccount struct {
	Errors   []GatewayAccountError   `json:"errors,required"`
	Messages []GatewayAccountMessage `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayAccountSuccess `json:"success,required"`
	Result  GatewayAccountResult  `json:"result"`
	JSON    gatewayAccountJSON    `json:"-"`
}

// gatewayAccountJSON contains the JSON metadata for the struct [GatewayAccount]
type gatewayAccountJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountJSON) RawJSON() string {
	return r.raw
}

type GatewayAccountError struct {
	Code             int64                      `json:"code,required"`
	Message          string                     `json:"message,required"`
	DocumentationURL string                     `json:"documentation_url"`
	Source           GatewayAccountErrorsSource `json:"source"`
	JSON             gatewayAccountErrorJSON    `json:"-"`
}

// gatewayAccountErrorJSON contains the JSON metadata for the struct
// [GatewayAccountError]
type gatewayAccountErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *GatewayAccountError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountErrorJSON) RawJSON() string {
	return r.raw
}

type GatewayAccountErrorsSource struct {
	Pointer string                         `json:"pointer"`
	JSON    gatewayAccountErrorsSourceJSON `json:"-"`
}

// gatewayAccountErrorsSourceJSON contains the JSON metadata for the struct
// [GatewayAccountErrorsSource]
type gatewayAccountErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type GatewayAccountMessage struct {
	Code             int64                        `json:"code,required"`
	Message          string                       `json:"message,required"`
	DocumentationURL string                       `json:"documentation_url"`
	Source           GatewayAccountMessagesSource `json:"source"`
	JSON             gatewayAccountMessageJSON    `json:"-"`
}

// gatewayAccountMessageJSON contains the JSON metadata for the struct
// [GatewayAccountMessage]
type gatewayAccountMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *GatewayAccountMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountMessageJSON) RawJSON() string {
	return r.raw
}

type GatewayAccountMessagesSource struct {
	Pointer string                           `json:"pointer"`
	JSON    gatewayAccountMessagesSourceJSON `json:"-"`
}

// gatewayAccountMessagesSourceJSON contains the JSON metadata for the struct
// [GatewayAccountMessagesSource]
type gatewayAccountMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayAccountSuccess bool

const (
	GatewayAccountSuccessTrue GatewayAccountSuccess = true
)

func (r GatewayAccountSuccess) IsKnown() bool {
	switch r {
	case GatewayAccountSuccessTrue:
		return true
	}
	return false
}

type GatewayAccountResult struct {
	// Cloudflare account ID.
	ID string `json:"id"`
	// Gateway internal ID.
	GatewayTag string `json:"gateway_tag"`
	// The name of the provider. Usually Cloudflare.
	ProviderName string                   `json:"provider_name"`
	JSON         gatewayAccountResultJSON `json:"-"`
}

// gatewayAccountResultJSON contains the JSON metadata for the struct
// [GatewayAccountResult]
type gatewayAccountResultJSON struct {
	ID           apijson.Field
	GatewayTag   apijson.Field
	ProviderName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GatewayAccountResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountResultJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayMessages struct {
	Code             int64                          `json:"code,required"`
	Message          string                         `json:"message,required"`
	DocumentationURL string                         `json:"documentation_url"`
	Source           ZeroTrustGatewayMessagesSource `json:"source"`
	JSON             zeroTrustGatewayMessagesJSON   `json:"-"`
}

// zeroTrustGatewayMessagesJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayMessages]
type zeroTrustGatewayMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustGatewayMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayMessagesJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayMessagesSource struct {
	Pointer string                             `json:"pointer"`
	JSON    zeroTrustGatewayMessagesSourceJSON `json:"-"`
}

// zeroTrustGatewayMessagesSourceJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayMessagesSource]
type zeroTrustGatewayMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListAppTypesResponse struct {
	Errors   []AccountGatewayListAppTypesResponseError   `json:"errors,required"`
	Messages []AccountGatewayListAppTypesResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success    AccountGatewayListAppTypesResponseSuccess    `json:"success,required"`
	Result     []AccountGatewayListAppTypesResponseResult   `json:"result"`
	ResultInfo AccountGatewayListAppTypesResponseResultInfo `json:"result_info"`
	JSON       accountGatewayListAppTypesResponseJSON       `json:"-"`
}

// accountGatewayListAppTypesResponseJSON contains the JSON metadata for the struct
// [AccountGatewayListAppTypesResponse]
type accountGatewayListAppTypesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListAppTypesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListAppTypesResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListAppTypesResponseError struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           AccountGatewayListAppTypesResponseErrorsSource `json:"source"`
	JSON             accountGatewayListAppTypesResponseErrorJSON    `json:"-"`
}

// accountGatewayListAppTypesResponseErrorJSON contains the JSON metadata for the
// struct [AccountGatewayListAppTypesResponseError]
type accountGatewayListAppTypesResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountGatewayListAppTypesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListAppTypesResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListAppTypesResponseErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    accountGatewayListAppTypesResponseErrorsSourceJSON `json:"-"`
}

// accountGatewayListAppTypesResponseErrorsSourceJSON contains the JSON metadata
// for the struct [AccountGatewayListAppTypesResponseErrorsSource]
type accountGatewayListAppTypesResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListAppTypesResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListAppTypesResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListAppTypesResponseMessage struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           AccountGatewayListAppTypesResponseMessagesSource `json:"source"`
	JSON             accountGatewayListAppTypesResponseMessageJSON    `json:"-"`
}

// accountGatewayListAppTypesResponseMessageJSON contains the JSON metadata for the
// struct [AccountGatewayListAppTypesResponseMessage]
type accountGatewayListAppTypesResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountGatewayListAppTypesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListAppTypesResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListAppTypesResponseMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    accountGatewayListAppTypesResponseMessagesSourceJSON `json:"-"`
}

// accountGatewayListAppTypesResponseMessagesSourceJSON contains the JSON metadata
// for the struct [AccountGatewayListAppTypesResponseMessagesSource]
type accountGatewayListAppTypesResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListAppTypesResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListAppTypesResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountGatewayListAppTypesResponseSuccess bool

const (
	AccountGatewayListAppTypesResponseSuccessTrue AccountGatewayListAppTypesResponseSuccess = true
)

func (r AccountGatewayListAppTypesResponseSuccess) IsKnown() bool {
	switch r {
	case AccountGatewayListAppTypesResponseSuccessTrue:
		return true
	}
	return false
}

type AccountGatewayListAppTypesResponseResult struct {
	// The identifier for this application. There is only one application per ID.
	ID int64 `json:"id"`
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of a returned application type.
	ApplicationTypeID int64     `json:"application_type_id"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// A short summary of applications with this type.
	Description string `json:"description"`
	// The name of the application or application type.
	Name  string                                       `json:"name"`
	JSON  accountGatewayListAppTypesResponseResultJSON `json:"-"`
	union AccountGatewayListAppTypesResponseResultUnion
}

// accountGatewayListAppTypesResponseResultJSON contains the JSON metadata for the
// struct [AccountGatewayListAppTypesResponseResult]
type accountGatewayListAppTypesResponseResultJSON struct {
	ID                apijson.Field
	ApplicationTypeID apijson.Field
	CreatedAt         apijson.Field
	Description       apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r accountGatewayListAppTypesResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *AccountGatewayListAppTypesResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = AccountGatewayListAppTypesResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountGatewayListAppTypesResponseResultUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountGatewayListAppTypesResponseResultZeroTrustGatewayApplication],
// [AccountGatewayListAppTypesResponseResultZeroTrustGatewayApplicationType].
func (r AccountGatewayListAppTypesResponseResult) AsUnion() AccountGatewayListAppTypesResponseResultUnion {
	return r.union
}

// Union satisfied by
// [AccountGatewayListAppTypesResponseResultZeroTrustGatewayApplication] or
// [AccountGatewayListAppTypesResponseResultZeroTrustGatewayApplicationType].
type AccountGatewayListAppTypesResponseResultUnion interface {
	implementsAccountGatewayListAppTypesResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountGatewayListAppTypesResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountGatewayListAppTypesResponseResultZeroTrustGatewayApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountGatewayListAppTypesResponseResultZeroTrustGatewayApplicationType{}),
		},
	)
}

type AccountGatewayListAppTypesResponseResultZeroTrustGatewayApplication struct {
	// The identifier for this application. There is only one application per ID.
	ID int64 `json:"id"`
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of a returned application type.
	ApplicationTypeID int64     `json:"application_type_id"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The name of the application or application type.
	Name string                                                                  `json:"name"`
	JSON accountGatewayListAppTypesResponseResultZeroTrustGatewayApplicationJSON `json:"-"`
}

// accountGatewayListAppTypesResponseResultZeroTrustGatewayApplicationJSON contains
// the JSON metadata for the struct
// [AccountGatewayListAppTypesResponseResultZeroTrustGatewayApplication]
type accountGatewayListAppTypesResponseResultZeroTrustGatewayApplicationJSON struct {
	ID                apijson.Field
	ApplicationTypeID apijson.Field
	CreatedAt         apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountGatewayListAppTypesResponseResultZeroTrustGatewayApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListAppTypesResponseResultZeroTrustGatewayApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccountGatewayListAppTypesResponseResultZeroTrustGatewayApplication) implementsAccountGatewayListAppTypesResponseResult() {
}

type AccountGatewayListAppTypesResponseResultZeroTrustGatewayApplicationType struct {
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of a returned application type.
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A short summary of applications with this type.
	Description string `json:"description"`
	// The name of the application or application type.
	Name string                                                                      `json:"name"`
	JSON accountGatewayListAppTypesResponseResultZeroTrustGatewayApplicationTypeJSON `json:"-"`
}

// accountGatewayListAppTypesResponseResultZeroTrustGatewayApplicationTypeJSON
// contains the JSON metadata for the struct
// [AccountGatewayListAppTypesResponseResultZeroTrustGatewayApplicationType]
type accountGatewayListAppTypesResponseResultZeroTrustGatewayApplicationTypeJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListAppTypesResponseResultZeroTrustGatewayApplicationType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListAppTypesResponseResultZeroTrustGatewayApplicationTypeJSON) RawJSON() string {
	return r.raw
}

func (r AccountGatewayListAppTypesResponseResultZeroTrustGatewayApplicationType) implementsAccountGatewayListAppTypesResponseResult() {
}

type AccountGatewayListAppTypesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       accountGatewayListAppTypesResponseResultInfoJSON `json:"-"`
}

// accountGatewayListAppTypesResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountGatewayListAppTypesResponseResultInfo]
type accountGatewayListAppTypesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListAppTypesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListAppTypesResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListCategoriesResponse struct {
	Errors   []AccountGatewayListCategoriesResponseError   `json:"errors,required"`
	Messages []AccountGatewayListCategoriesResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success    AccountGatewayListCategoriesResponseSuccess    `json:"success,required"`
	Result     []AccountGatewayListCategoriesResponseResult   `json:"result"`
	ResultInfo AccountGatewayListCategoriesResponseResultInfo `json:"result_info"`
	JSON       accountGatewayListCategoriesResponseJSON       `json:"-"`
}

// accountGatewayListCategoriesResponseJSON contains the JSON metadata for the
// struct [AccountGatewayListCategoriesResponse]
type accountGatewayListCategoriesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListCategoriesResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListCategoriesResponseError struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           AccountGatewayListCategoriesResponseErrorsSource `json:"source"`
	JSON             accountGatewayListCategoriesResponseErrorJSON    `json:"-"`
}

// accountGatewayListCategoriesResponseErrorJSON contains the JSON metadata for the
// struct [AccountGatewayListCategoriesResponseError]
type accountGatewayListCategoriesResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountGatewayListCategoriesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListCategoriesResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListCategoriesResponseErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    accountGatewayListCategoriesResponseErrorsSourceJSON `json:"-"`
}

// accountGatewayListCategoriesResponseErrorsSourceJSON contains the JSON metadata
// for the struct [AccountGatewayListCategoriesResponseErrorsSource]
type accountGatewayListCategoriesResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListCategoriesResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListCategoriesResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListCategoriesResponseMessage struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           AccountGatewayListCategoriesResponseMessagesSource `json:"source"`
	JSON             accountGatewayListCategoriesResponseMessageJSON    `json:"-"`
}

// accountGatewayListCategoriesResponseMessageJSON contains the JSON metadata for
// the struct [AccountGatewayListCategoriesResponseMessage]
type accountGatewayListCategoriesResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountGatewayListCategoriesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListCategoriesResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListCategoriesResponseMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    accountGatewayListCategoriesResponseMessagesSourceJSON `json:"-"`
}

// accountGatewayListCategoriesResponseMessagesSourceJSON contains the JSON
// metadata for the struct [AccountGatewayListCategoriesResponseMessagesSource]
type accountGatewayListCategoriesResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListCategoriesResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListCategoriesResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountGatewayListCategoriesResponseSuccess bool

const (
	AccountGatewayListCategoriesResponseSuccessTrue AccountGatewayListCategoriesResponseSuccess = true
)

func (r AccountGatewayListCategoriesResponseSuccess) IsKnown() bool {
	switch r {
	case AccountGatewayListCategoriesResponseSuccessTrue:
		return true
	}
	return false
}

type AccountGatewayListCategoriesResponseResult struct {
	// The identifier for this category. There is only one category per ID.
	ID int64 `json:"id"`
	// True if the category is in beta and subject to change.
	Beta bool `json:"beta"`
	// Which account types are allowed to create policies based on this category.
	// `blocked` categories are blocked unconditionally for all accounts.
	// `removalPending` categories can be removed from policies but not added.
	// `noBlock` categories cannot be blocked.
	Class Class `json:"class"`
	// A short summary of domains in the category.
	Description string `json:"description"`
	// The name of the category.
	Name string `json:"name"`
	// All subcategories for this category.
	Subcategories []AccountGatewayListCategoriesResponseResultSubcategory `json:"subcategories"`
	JSON          accountGatewayListCategoriesResponseResultJSON          `json:"-"`
}

// accountGatewayListCategoriesResponseResultJSON contains the JSON metadata for
// the struct [AccountGatewayListCategoriesResponseResult]
type accountGatewayListCategoriesResponseResultJSON struct {
	ID            apijson.Field
	Beta          apijson.Field
	Class         apijson.Field
	Description   apijson.Field
	Name          apijson.Field
	Subcategories apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountGatewayListCategoriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListCategoriesResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListCategoriesResponseResultSubcategory struct {
	// The identifier for this category. There is only one category per ID.
	ID int64 `json:"id"`
	// True if the category is in beta and subject to change.
	Beta bool `json:"beta"`
	// Which account types are allowed to create policies based on this category.
	// `blocked` categories are blocked unconditionally for all accounts.
	// `removalPending` categories can be removed from policies but not added.
	// `noBlock` categories cannot be blocked.
	Class Class `json:"class"`
	// A short summary of domains in the category.
	Description string `json:"description"`
	// The name of the category.
	Name string                                                    `json:"name"`
	JSON accountGatewayListCategoriesResponseResultSubcategoryJSON `json:"-"`
}

// accountGatewayListCategoriesResponseResultSubcategoryJSON contains the JSON
// metadata for the struct [AccountGatewayListCategoriesResponseResultSubcategory]
type accountGatewayListCategoriesResponseResultSubcategoryJSON struct {
	ID          apijson.Field
	Beta        apijson.Field
	Class       apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListCategoriesResponseResultSubcategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListCategoriesResponseResultSubcategoryJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListCategoriesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       accountGatewayListCategoriesResponseResultInfoJSON `json:"-"`
}

// accountGatewayListCategoriesResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountGatewayListCategoriesResponseResultInfo]
type accountGatewayListCategoriesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListCategoriesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListCategoriesResponseResultInfoJSON) RawJSON() string {
	return r.raw
}
