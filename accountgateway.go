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

type APIResponseCollectionZeroTrustGateway struct {
	ResultInfo APIResponseCollectionZeroTrustGatewayResultInfo `json:"result_info"`
	JSON       apiResponseCollectionZeroTrustGatewayJSON       `json:"-"`
	APIResponseZeroTrustGateway
}

// apiResponseCollectionZeroTrustGatewayJSON contains the JSON metadata for the
// struct [APIResponseCollectionZeroTrustGateway]
type apiResponseCollectionZeroTrustGatewayJSON struct {
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionZeroTrustGateway) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionZeroTrustGatewayJSON) RawJSON() string {
	return r.raw
}

type APIResponseCollectionZeroTrustGatewayResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       apiResponseCollectionZeroTrustGatewayResultInfoJSON `json:"-"`
}

// apiResponseCollectionZeroTrustGatewayResultInfoJSON contains the JSON metadata
// for the struct [APIResponseCollectionZeroTrustGatewayResultInfo]
type apiResponseCollectionZeroTrustGatewayResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionZeroTrustGatewayResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionZeroTrustGatewayResultInfoJSON) RawJSON() string {
	return r.raw
}

type APIResponseSingleZeroTrustGateway struct {
	Errors   []ZeroTrustGatewayMessages `json:"errors,required"`
	Messages []ZeroTrustGatewayMessages `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseSingleZeroTrustGatewaySuccess `json:"success,required"`
	JSON    apiResponseSingleZeroTrustGatewayJSON    `json:"-"`
}

// apiResponseSingleZeroTrustGatewayJSON contains the JSON metadata for the struct
// [APIResponseSingleZeroTrustGateway]
type apiResponseSingleZeroTrustGatewayJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleZeroTrustGateway) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleZeroTrustGatewayJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseSingleZeroTrustGatewaySuccess bool

const (
	APIResponseSingleZeroTrustGatewaySuccessTrue APIResponseSingleZeroTrustGatewaySuccess = true
)

func (r APIResponseSingleZeroTrustGatewaySuccess) IsKnown() bool {
	switch r {
	case APIResponseSingleZeroTrustGatewaySuccessTrue:
		return true
	}
	return false
}

type APIResponseZeroTrustGateway struct {
	Errors   []ZeroTrustGatewayMessages `json:"errors,required"`
	Messages []ZeroTrustGatewayMessages `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseZeroTrustGatewaySuccess `json:"success,required"`
	JSON    apiResponseZeroTrustGatewayJSON    `json:"-"`
}

// apiResponseZeroTrustGatewayJSON contains the JSON metadata for the struct
// [APIResponseZeroTrustGateway]
type apiResponseZeroTrustGatewayJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseZeroTrustGateway) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseZeroTrustGatewayJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseZeroTrustGatewaySuccess bool

const (
	APIResponseZeroTrustGatewaySuccessTrue APIResponseZeroTrustGatewaySuccess = true
)

func (r APIResponseZeroTrustGatewaySuccess) IsKnown() bool {
	switch r {
	case APIResponseZeroTrustGatewaySuccessTrue:
		return true
	}
	return false
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
	Result GatewayAccountResult `json:"result"`
	JSON   gatewayAccountJSON   `json:"-"`
	APIResponseSingleZeroTrustGateway
}

// gatewayAccountJSON contains the JSON metadata for the struct [GatewayAccount]
type gatewayAccountJSON struct {
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
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    zeroTrustGatewayMessagesJSON `json:"-"`
}

// zeroTrustGatewayMessagesJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayMessages]
type zeroTrustGatewayMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayMessagesJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListAppTypesResponse struct {
	Result []AccountGatewayListAppTypesResponseResult `json:"result"`
	JSON   accountGatewayListAppTypesResponseJSON     `json:"-"`
	APIResponseCollectionZeroTrustGateway
}

// accountGatewayListAppTypesResponseJSON contains the JSON metadata for the struct
// [AccountGatewayListAppTypesResponse]
type accountGatewayListAppTypesResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListAppTypesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListAppTypesResponseJSON) RawJSON() string {
	return r.raw
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

type AccountGatewayListCategoriesResponse struct {
	Result []AccountGatewayListCategoriesResponseResult `json:"result"`
	JSON   accountGatewayListCategoriesResponseJSON     `json:"-"`
	APIResponseCollectionZeroTrustGateway
}

// accountGatewayListCategoriesResponseJSON contains the JSON metadata for the
// struct [AccountGatewayListCategoriesResponse]
type accountGatewayListCategoriesResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListCategoriesResponseJSON) RawJSON() string {
	return r.raw
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
