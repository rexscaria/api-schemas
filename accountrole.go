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

// AccountRoleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRoleService] method instead.
type AccountRoleService struct {
	Options []option.RequestOption
}

// NewAccountRoleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountRoleService(opts ...option.RequestOption) (r *AccountRoleService) {
	r = &AccountRoleService{}
	r.Options = opts
	return
}

// Get information about a specific role for an account.
func (r *AccountRoleService) Get(ctx context.Context, accountID string, roleID string, opts ...option.RequestOption) (res *AccountRoleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if roleID == "" {
		err = errors.New("missing required role_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/roles/%s", accountID, roleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all available roles for an account.
func (r *AccountRoleService) List(ctx context.Context, accountID string, query AccountRoleListParams, opts ...option.RequestOption) (res *AccountRoleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/roles", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type IamGrants struct {
	Read  bool          `json:"read"`
	Write bool          `json:"write"`
	JSON  iamGrantsJSON `json:"-"`
}

// iamGrantsJSON contains the JSON metadata for the struct [IamGrants]
type iamGrantsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamGrants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamGrantsJSON) RawJSON() string {
	return r.raw
}

type IamGrantsParam struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r IamGrantsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IamPermissions struct {
	Analytics    IamGrants          `json:"analytics"`
	Billing      IamGrants          `json:"billing"`
	CachePurge   IamGrants          `json:"cache_purge"`
	DNS          IamGrants          `json:"dns"`
	DNSRecords   IamGrants          `json:"dns_records"`
	Lb           IamGrants          `json:"lb"`
	Logs         IamGrants          `json:"logs"`
	Organization IamGrants          `json:"organization"`
	Ssl          IamGrants          `json:"ssl"`
	Waf          IamGrants          `json:"waf"`
	ZoneSettings IamGrants          `json:"zone_settings"`
	Zones        IamGrants          `json:"zones"`
	JSON         iamPermissionsJSON `json:"-"`
}

// iamPermissionsJSON contains the JSON metadata for the struct [IamPermissions]
type iamPermissionsJSON struct {
	Analytics    apijson.Field
	Billing      apijson.Field
	CachePurge   apijson.Field
	DNS          apijson.Field
	DNSRecords   apijson.Field
	Lb           apijson.Field
	Logs         apijson.Field
	Organization apijson.Field
	Ssl          apijson.Field
	Waf          apijson.Field
	ZoneSettings apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IamPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamPermissionsJSON) RawJSON() string {
	return r.raw
}

type IamPermissionsParam struct {
	Analytics    param.Field[IamGrantsParam] `json:"analytics"`
	Billing      param.Field[IamGrantsParam] `json:"billing"`
	CachePurge   param.Field[IamGrantsParam] `json:"cache_purge"`
	DNS          param.Field[IamGrantsParam] `json:"dns"`
	DNSRecords   param.Field[IamGrantsParam] `json:"dns_records"`
	Lb           param.Field[IamGrantsParam] `json:"lb"`
	Logs         param.Field[IamGrantsParam] `json:"logs"`
	Organization param.Field[IamGrantsParam] `json:"organization"`
	Ssl          param.Field[IamGrantsParam] `json:"ssl"`
	Waf          param.Field[IamGrantsParam] `json:"waf"`
	ZoneSettings param.Field[IamGrantsParam] `json:"zone_settings"`
	Zones        param.Field[IamGrantsParam] `json:"zones"`
}

func (r IamPermissionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IamRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string         `json:"name,required"`
	Permissions IamPermissions `json:"permissions,required"`
	JSON        iamRoleJSON    `json:"-"`
}

// iamRoleJSON contains the JSON metadata for the struct [IamRole]
type iamRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamRoleJSON) RawJSON() string {
	return r.raw
}

type IamRoleParam struct {
	// Role identifier tag.
	ID param.Field[string] `json:"id,required"`
}

func (r IamRoleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRoleGetResponse struct {
	Errors   []AccountRoleGetResponseError   `json:"errors,required"`
	Messages []AccountRoleGetResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountRoleGetResponseSuccess `json:"success,required"`
	Result  IamRole                       `json:"result"`
	JSON    accountRoleGetResponseJSON    `json:"-"`
}

// accountRoleGetResponseJSON contains the JSON metadata for the struct
// [AccountRoleGetResponse]
type accountRoleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRoleGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRoleGetResponseError struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           AccountRoleGetResponseErrorsSource `json:"source"`
	JSON             accountRoleGetResponseErrorJSON    `json:"-"`
}

// accountRoleGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountRoleGetResponseError]
type accountRoleGetResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountRoleGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRoleGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountRoleGetResponseErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    accountRoleGetResponseErrorsSourceJSON `json:"-"`
}

// accountRoleGetResponseErrorsSourceJSON contains the JSON metadata for the struct
// [AccountRoleGetResponseErrorsSource]
type accountRoleGetResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleGetResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRoleGetResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountRoleGetResponseMessage struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           AccountRoleGetResponseMessagesSource `json:"source"`
	JSON             accountRoleGetResponseMessageJSON    `json:"-"`
}

// accountRoleGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountRoleGetResponseMessage]
type accountRoleGetResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountRoleGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRoleGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountRoleGetResponseMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    accountRoleGetResponseMessagesSourceJSON `json:"-"`
}

// accountRoleGetResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AccountRoleGetResponseMessagesSource]
type accountRoleGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRoleGetResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRoleGetResponseSuccess bool

const (
	AccountRoleGetResponseSuccessTrue AccountRoleGetResponseSuccess = true
)

func (r AccountRoleGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRoleGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountRoleListResponse struct {
	Errors   []AccountRoleListResponseError   `json:"errors,required"`
	Messages []AccountRoleListResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountRoleListResponseSuccess    `json:"success,required"`
	Result     []IamRole                         `json:"result"`
	ResultInfo AccountRoleListResponseResultInfo `json:"result_info"`
	JSON       accountRoleListResponseJSON       `json:"-"`
}

// accountRoleListResponseJSON contains the JSON metadata for the struct
// [AccountRoleListResponse]
type accountRoleListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRoleListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRoleListResponseError struct {
	Code             int64                               `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           AccountRoleListResponseErrorsSource `json:"source"`
	JSON             accountRoleListResponseErrorJSON    `json:"-"`
}

// accountRoleListResponseErrorJSON contains the JSON metadata for the struct
// [AccountRoleListResponseError]
type accountRoleListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountRoleListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRoleListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountRoleListResponseErrorsSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    accountRoleListResponseErrorsSourceJSON `json:"-"`
}

// accountRoleListResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountRoleListResponseErrorsSource]
type accountRoleListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRoleListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountRoleListResponseMessage struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           AccountRoleListResponseMessagesSource `json:"source"`
	JSON             accountRoleListResponseMessageJSON    `json:"-"`
}

// accountRoleListResponseMessageJSON contains the JSON metadata for the struct
// [AccountRoleListResponseMessage]
type accountRoleListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountRoleListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRoleListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountRoleListResponseMessagesSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    accountRoleListResponseMessagesSourceJSON `json:"-"`
}

// accountRoleListResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AccountRoleListResponseMessagesSource]
type accountRoleListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRoleListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRoleListResponseSuccess bool

const (
	AccountRoleListResponseSuccessTrue AccountRoleListResponseSuccess = true
)

func (r AccountRoleListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRoleListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountRoleListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                               `json:"total_count"`
	JSON       accountRoleListResponseResultInfoJSON `json:"-"`
}

// accountRoleListResponseResultInfoJSON contains the JSON metadata for the struct
// [AccountRoleListResponseResultInfo]
type accountRoleListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRoleListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountRoleListParams struct {
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of roles per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountRoleListParams]'s query parameters as `url.Values`.
func (r AccountRoleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
