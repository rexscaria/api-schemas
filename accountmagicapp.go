// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountMagicAppService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicAppService] method instead.
type AccountMagicAppService struct {
	Options []option.RequestOption
}

// NewAccountMagicAppService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountMagicAppService(opts ...option.RequestOption) (r *AccountMagicAppService) {
	r = &AccountMagicAppService{}
	r.Options = opts
	return
}

// Creates a new App for an account
func (r *AccountMagicAppService) New(ctx context.Context, accountID string, body AccountMagicAppNewParams, opts ...option.RequestOption) (res *MagicAppSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/apps", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an Account App
func (r *AccountMagicAppService) Update(ctx context.Context, accountID string, accountAppID string, body AccountMagicAppUpdateParams, opts ...option.RequestOption) (res *MagicAppSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if accountAppID == "" {
		err = errors.New("missing required account_app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/apps/%s", accountID, accountAppID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists Apps associated with an account.
func (r *AccountMagicAppService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountMagicAppListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/apps", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes specific Account App.
func (r *AccountMagicAppService) Delete(ctx context.Context, accountID string, accountAppID string, opts ...option.RequestOption) (res *MagicAppSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if accountAppID == "" {
		err = errors.New("missing required account_app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/apps/%s", accountID, accountAppID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Custom app defined for an account.
type MagicAccountApp struct {
	// Magic account app ID.
	AccountAppID string `json:"account_app_id,required"`
	// FQDNs to associate with traffic decisions.
	Hostnames []string `json:"hostnames"`
	// CIDRs to associate with traffic decisions.
	IPSubnets []MagicAppSubnetItem `json:"ip_subnets"`
	// Display name for the app.
	Name string `json:"name"`
	// Category of the app.
	Type string              `json:"type"`
	JSON magicAccountAppJSON `json:"-"`
}

// magicAccountAppJSON contains the JSON metadata for the struct [MagicAccountApp]
type magicAccountAppJSON struct {
	AccountAppID apijson.Field
	Hostnames    apijson.Field
	IPSubnets    apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicAccountApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicAccountAppJSON) RawJSON() string {
	return r.raw
}

func (r MagicAccountApp) implementsAccountMagicAppListResponseResult() {}

type MagicAppSingleResponse struct {
	// Custom app defined for an account.
	Result MagicAccountApp            `json:"result"`
	JSON   magicAppSingleResponseJSON `json:"-"`
	MagicAppsResponseObject
}

// magicAppSingleResponseJSON contains the JSON metadata for the struct
// [MagicAppSingleResponse]
type magicAppSingleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicAppSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicAppSingleResponseJSON) RawJSON() string {
	return r.raw
}

type MagicAppSubnetItem = string

type MagicAppSubnetItemParam = string

type MagicAppsResponseObject struct {
	Errors   []MagicMessageItem `json:"errors,required"`
	Messages []MagicMessageItem `json:"messages,required"`
	Result   interface{}        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MagicAppsResponseObjectSuccess `json:"success,required"`
	JSON    magicAppsResponseObjectJSON    `json:"-"`
}

// magicAppsResponseObjectJSON contains the JSON metadata for the struct
// [MagicAppsResponseObject]
type magicAppsResponseObjectJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicAppsResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicAppsResponseObjectJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicAppsResponseObjectSuccess bool

const (
	MagicAppsResponseObjectSuccessTrue MagicAppsResponseObjectSuccess = true
)

func (r MagicAppsResponseObjectSuccess) IsKnown() bool {
	switch r {
	case MagicAppsResponseObjectSuccessTrue:
		return true
	}
	return false
}

type MagicMessageItem struct {
	Code    int64                `json:"code,required"`
	Message string               `json:"message,required"`
	JSON    magicMessageItemJSON `json:"-"`
}

// magicMessageItemJSON contains the JSON metadata for the struct
// [MagicMessageItem]
type magicMessageItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicMessageItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicMessageItemJSON) RawJSON() string {
	return r.raw
}

type AccountMagicAppListResponse struct {
	Errors   []MagicMessageItem                  `json:"errors,required"`
	Messages []MagicMessageItem                  `json:"messages,required"`
	Result   []AccountMagicAppListResponseResult `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AccountMagicAppListResponseSuccess `json:"success,required"`
	JSON    accountMagicAppListResponseJSON    `json:"-"`
}

// accountMagicAppListResponseJSON contains the JSON metadata for the struct
// [AccountMagicAppListResponse]
type accountMagicAppListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicAppListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicAppListResponseJSON) RawJSON() string {
	return r.raw
}

// Custom app defined for an account.
type AccountMagicAppListResponseResult struct {
	// Magic account app ID.
	AccountAppID string `json:"account_app_id"`
	// This field can have the runtime type of [[]string].
	Hostnames interface{} `json:"hostnames"`
	// This field can have the runtime type of [[]MagicAppSubnetItem].
	IPSubnets interface{} `json:"ip_subnets"`
	// Managed app ID.
	ManagedAppID string `json:"managed_app_id"`
	// Display name for the app.
	Name string `json:"name"`
	// Category of the app.
	Type  string                                `json:"type"`
	JSON  accountMagicAppListResponseResultJSON `json:"-"`
	union AccountMagicAppListResponseResultUnion
}

// accountMagicAppListResponseResultJSON contains the JSON metadata for the struct
// [AccountMagicAppListResponseResult]
type accountMagicAppListResponseResultJSON struct {
	AccountAppID apijson.Field
	Hostnames    apijson.Field
	IPSubnets    apijson.Field
	ManagedAppID apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r accountMagicAppListResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *AccountMagicAppListResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = AccountMagicAppListResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountMagicAppListResponseResultUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [MagicAccountApp],
// [AccountMagicAppListResponseResultMagicManagedApp].
func (r AccountMagicAppListResponseResult) AsUnion() AccountMagicAppListResponseResultUnion {
	return r.union
}

// Custom app defined for an account.
//
// Union satisfied by [MagicAccountApp] or
// [AccountMagicAppListResponseResultMagicManagedApp].
type AccountMagicAppListResponseResultUnion interface {
	implementsAccountMagicAppListResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountMagicAppListResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MagicAccountApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicAppListResponseResultMagicManagedApp{}),
		},
	)
}

// Managed app defined by Cloudflare.
type AccountMagicAppListResponseResultMagicManagedApp struct {
	// Managed app ID.
	ManagedAppID string `json:"managed_app_id,required"`
	// FQDNs to associate with traffic decisions.
	Hostnames []string `json:"hostnames"`
	// CIDRs to associate with traffic decisions.
	IPSubnets []MagicAppSubnetItem `json:"ip_subnets"`
	// Display name for the app.
	Name string `json:"name"`
	// Category of the app.
	Type string                                               `json:"type"`
	JSON accountMagicAppListResponseResultMagicManagedAppJSON `json:"-"`
}

// accountMagicAppListResponseResultMagicManagedAppJSON contains the JSON metadata
// for the struct [AccountMagicAppListResponseResultMagicManagedApp]
type accountMagicAppListResponseResultMagicManagedAppJSON struct {
	ManagedAppID apijson.Field
	Hostnames    apijson.Field
	IPSubnets    apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountMagicAppListResponseResultMagicManagedApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicAppListResponseResultMagicManagedAppJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicAppListResponseResultMagicManagedApp) implementsAccountMagicAppListResponseResult() {
}

// Whether the API call was successful
type AccountMagicAppListResponseSuccess bool

const (
	AccountMagicAppListResponseSuccessTrue AccountMagicAppListResponseSuccess = true
)

func (r AccountMagicAppListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicAppListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicAppNewParams struct {
	// Display name for the app.
	Name param.Field[string] `json:"name,required"`
	// Category of the app.
	Type param.Field[string] `json:"type,required"`
	// FQDNs to associate with traffic decisions.
	Hostnames param.Field[[]string] `json:"hostnames"`
	// CIDRs to associate with traffic decisions.
	IPSubnets param.Field[[]MagicAppSubnetItemParam] `json:"ip_subnets"`
}

func (r AccountMagicAppNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicAppUpdateParams struct {
	// FQDNs to associate with traffic decisions.
	Hostnames param.Field[[]string] `json:"hostnames"`
	// CIDRs to associate with traffic decisions.
	IPSubnets param.Field[[]MagicAppSubnetItemParam] `json:"ip_subnets"`
	// Display name for the app.
	Name param.Field[string] `json:"name"`
	// Category of the app.
	Type param.Field[string] `json:"type"`
}

func (r AccountMagicAppUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
