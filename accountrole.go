// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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
	Result IamRole                    `json:"result"`
	JSON   accountRoleGetResponseJSON `json:"-"`
	APIResponseSingleIam
}

// accountRoleGetResponseJSON contains the JSON metadata for the struct
// [AccountRoleGetResponse]
type accountRoleGetResponseJSON struct {
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

type AccountRoleListResponse struct {
	Result []IamRole                   `json:"result"`
	JSON   accountRoleListResponseJSON `json:"-"`
	IamAPIResponseCollection
}

// accountRoleListResponseJSON contains the JSON metadata for the struct
// [AccountRoleListResponse]
type accountRoleListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRoleListResponseJSON) RawJSON() string {
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
