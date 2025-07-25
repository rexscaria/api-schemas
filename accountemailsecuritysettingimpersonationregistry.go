// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountEmailSecuritySettingImpersonationRegistryService contains methods and
// other services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountEmailSecuritySettingImpersonationRegistryService] method instead.
type AccountEmailSecuritySettingImpersonationRegistryService struct {
	Options []option.RequestOption
}

// NewAccountEmailSecuritySettingImpersonationRegistryService generates a new
// service that applies the given options to each request. These options are
// applied after the parent client's options (if there is one), and before any
// request-specific options.
func NewAccountEmailSecuritySettingImpersonationRegistryService(opts ...option.RequestOption) (r *AccountEmailSecuritySettingImpersonationRegistryService) {
	r = &AccountEmailSecuritySettingImpersonationRegistryService{}
	r.Options = opts
	return
}

// Create an entry in impersonation registry
func (r *AccountEmailSecuritySettingImpersonationRegistryService) New(ctx context.Context, accountID string, body AccountEmailSecuritySettingImpersonationRegistryNewParams, opts ...option.RequestOption) (res *AccountEmailSecuritySettingImpersonationRegistryNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/impersonation_registry", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an entry in impersonation registry
func (r *AccountEmailSecuritySettingImpersonationRegistryService) Get(ctx context.Context, accountID string, displayNameID int64, opts ...option.RequestOption) (res *AccountEmailSecuritySettingImpersonationRegistryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/impersonation_registry/%v", accountID, displayNameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an entry in impersonation registry
func (r *AccountEmailSecuritySettingImpersonationRegistryService) Update(ctx context.Context, accountID string, displayNameID int64, body AccountEmailSecuritySettingImpersonationRegistryUpdateParams, opts ...option.RequestOption) (res *AccountEmailSecuritySettingImpersonationRegistryUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/impersonation_registry/%v", accountID, displayNameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists, searches, and sorts entries in the impersonation registry.
func (r *AccountEmailSecuritySettingImpersonationRegistryService) List(ctx context.Context, accountID string, query AccountEmailSecuritySettingImpersonationRegistryListParams, opts ...option.RequestOption) (res *AccountEmailSecuritySettingImpersonationRegistryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/impersonation_registry", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an entry from impersonation registry
func (r *AccountEmailSecuritySettingImpersonationRegistryService) Delete(ctx context.Context, accountID string, displayNameID int64, opts ...option.RequestOption) (res *AccountEmailSecuritySettingImpersonationRegistryDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/impersonation_registry/%v", accountID, displayNameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountEmailSecuritySettingImpersonationRegistryNewResponse struct {
	Errors   []EmailSecurityMessage                                            `json:"errors,required"`
	Messages []EmailSecurityMessage                                            `json:"messages,required"`
	Result   AccountEmailSecuritySettingImpersonationRegistryNewResponseResult `json:"result,required"`
	Success  bool                                                              `json:"success,required"`
	JSON     accountEmailSecuritySettingImpersonationRegistryNewResponseJSON   `json:"-"`
}

// accountEmailSecuritySettingImpersonationRegistryNewResponseJSON contains the
// JSON metadata for the struct
// [AccountEmailSecuritySettingImpersonationRegistryNewResponse]
type accountEmailSecuritySettingImpersonationRegistryNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingImpersonationRegistryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingImpersonationRegistryNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingImpersonationRegistryNewResponseResult struct {
	ID              int64     `json:"id,required"`
	CreatedAt       time.Time `json:"created_at,required" format:"date-time"`
	Email           string    `json:"email,required"`
	IsEmailRegex    bool      `json:"is_email_regex,required"`
	LastModified    time.Time `json:"last_modified,required" format:"date-time"`
	Name            string    `json:"name,required"`
	Comments        string    `json:"comments,nullable"`
	DirectoryID     int64     `json:"directory_id,nullable"`
	DirectoryNodeID int64     `json:"directory_node_id,nullable"`
	// Deprecated: deprecated
	ExternalDirectoryNodeID string                                                                `json:"external_directory_node_id,nullable"`
	Provenance              string                                                                `json:"provenance,nullable"`
	JSON                    accountEmailSecuritySettingImpersonationRegistryNewResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingImpersonationRegistryNewResponseResultJSON contains
// the JSON metadata for the struct
// [AccountEmailSecuritySettingImpersonationRegistryNewResponseResult]
type accountEmailSecuritySettingImpersonationRegistryNewResponseResultJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	Email                   apijson.Field
	IsEmailRegex            apijson.Field
	LastModified            apijson.Field
	Name                    apijson.Field
	Comments                apijson.Field
	DirectoryID             apijson.Field
	DirectoryNodeID         apijson.Field
	ExternalDirectoryNodeID apijson.Field
	Provenance              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingImpersonationRegistryNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingImpersonationRegistryNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingImpersonationRegistryGetResponse struct {
	Errors   []EmailSecurityMessage                                            `json:"errors,required"`
	Messages []EmailSecurityMessage                                            `json:"messages,required"`
	Result   AccountEmailSecuritySettingImpersonationRegistryGetResponseResult `json:"result,required"`
	Success  bool                                                              `json:"success,required"`
	JSON     accountEmailSecuritySettingImpersonationRegistryGetResponseJSON   `json:"-"`
}

// accountEmailSecuritySettingImpersonationRegistryGetResponseJSON contains the
// JSON metadata for the struct
// [AccountEmailSecuritySettingImpersonationRegistryGetResponse]
type accountEmailSecuritySettingImpersonationRegistryGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingImpersonationRegistryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingImpersonationRegistryGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingImpersonationRegistryGetResponseResult struct {
	ID              int64     `json:"id,required"`
	CreatedAt       time.Time `json:"created_at,required" format:"date-time"`
	Email           string    `json:"email,required"`
	IsEmailRegex    bool      `json:"is_email_regex,required"`
	LastModified    time.Time `json:"last_modified,required" format:"date-time"`
	Name            string    `json:"name,required"`
	Comments        string    `json:"comments,nullable"`
	DirectoryID     int64     `json:"directory_id,nullable"`
	DirectoryNodeID int64     `json:"directory_node_id,nullable"`
	// Deprecated: deprecated
	ExternalDirectoryNodeID string                                                                `json:"external_directory_node_id,nullable"`
	Provenance              string                                                                `json:"provenance,nullable"`
	JSON                    accountEmailSecuritySettingImpersonationRegistryGetResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingImpersonationRegistryGetResponseResultJSON contains
// the JSON metadata for the struct
// [AccountEmailSecuritySettingImpersonationRegistryGetResponseResult]
type accountEmailSecuritySettingImpersonationRegistryGetResponseResultJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	Email                   apijson.Field
	IsEmailRegex            apijson.Field
	LastModified            apijson.Field
	Name                    apijson.Field
	Comments                apijson.Field
	DirectoryID             apijson.Field
	DirectoryNodeID         apijson.Field
	ExternalDirectoryNodeID apijson.Field
	Provenance              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingImpersonationRegistryGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingImpersonationRegistryGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingImpersonationRegistryUpdateResponse struct {
	Errors   []EmailSecurityMessage                                               `json:"errors,required"`
	Messages []EmailSecurityMessage                                               `json:"messages,required"`
	Result   AccountEmailSecuritySettingImpersonationRegistryUpdateResponseResult `json:"result,required"`
	Success  bool                                                                 `json:"success,required"`
	JSON     accountEmailSecuritySettingImpersonationRegistryUpdateResponseJSON   `json:"-"`
}

// accountEmailSecuritySettingImpersonationRegistryUpdateResponseJSON contains the
// JSON metadata for the struct
// [AccountEmailSecuritySettingImpersonationRegistryUpdateResponse]
type accountEmailSecuritySettingImpersonationRegistryUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingImpersonationRegistryUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingImpersonationRegistryUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingImpersonationRegistryUpdateResponseResult struct {
	ID              int64     `json:"id,required"`
	CreatedAt       time.Time `json:"created_at,required" format:"date-time"`
	Email           string    `json:"email,required"`
	IsEmailRegex    bool      `json:"is_email_regex,required"`
	LastModified    time.Time `json:"last_modified,required" format:"date-time"`
	Name            string    `json:"name,required"`
	Comments        string    `json:"comments,nullable"`
	DirectoryID     int64     `json:"directory_id,nullable"`
	DirectoryNodeID int64     `json:"directory_node_id,nullable"`
	// Deprecated: deprecated
	ExternalDirectoryNodeID string                                                                   `json:"external_directory_node_id,nullable"`
	Provenance              string                                                                   `json:"provenance,nullable"`
	JSON                    accountEmailSecuritySettingImpersonationRegistryUpdateResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingImpersonationRegistryUpdateResponseResultJSON
// contains the JSON metadata for the struct
// [AccountEmailSecuritySettingImpersonationRegistryUpdateResponseResult]
type accountEmailSecuritySettingImpersonationRegistryUpdateResponseResultJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	Email                   apijson.Field
	IsEmailRegex            apijson.Field
	LastModified            apijson.Field
	Name                    apijson.Field
	Comments                apijson.Field
	DirectoryID             apijson.Field
	DirectoryNodeID         apijson.Field
	ExternalDirectoryNodeID apijson.Field
	Provenance              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingImpersonationRegistryUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingImpersonationRegistryUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingImpersonationRegistryListResponse struct {
	Errors     []EmailSecurityMessage                                               `json:"errors,required"`
	Messages   []EmailSecurityMessage                                               `json:"messages,required"`
	Result     []AccountEmailSecuritySettingImpersonationRegistryListResponseResult `json:"result,required"`
	ResultInfo ResultInfoEmailSecurity                                              `json:"result_info,required"`
	Success    bool                                                                 `json:"success,required"`
	JSON       accountEmailSecuritySettingImpersonationRegistryListResponseJSON     `json:"-"`
}

// accountEmailSecuritySettingImpersonationRegistryListResponseJSON contains the
// JSON metadata for the struct
// [AccountEmailSecuritySettingImpersonationRegistryListResponse]
type accountEmailSecuritySettingImpersonationRegistryListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingImpersonationRegistryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingImpersonationRegistryListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingImpersonationRegistryListResponseResult struct {
	ID              int64     `json:"id,required"`
	CreatedAt       time.Time `json:"created_at,required" format:"date-time"`
	Email           string    `json:"email,required"`
	IsEmailRegex    bool      `json:"is_email_regex,required"`
	LastModified    time.Time `json:"last_modified,required" format:"date-time"`
	Name            string    `json:"name,required"`
	Comments        string    `json:"comments,nullable"`
	DirectoryID     int64     `json:"directory_id,nullable"`
	DirectoryNodeID int64     `json:"directory_node_id,nullable"`
	// Deprecated: deprecated
	ExternalDirectoryNodeID string                                                                 `json:"external_directory_node_id,nullable"`
	Provenance              string                                                                 `json:"provenance,nullable"`
	JSON                    accountEmailSecuritySettingImpersonationRegistryListResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingImpersonationRegistryListResponseResultJSON contains
// the JSON metadata for the struct
// [AccountEmailSecuritySettingImpersonationRegistryListResponseResult]
type accountEmailSecuritySettingImpersonationRegistryListResponseResultJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	Email                   apijson.Field
	IsEmailRegex            apijson.Field
	LastModified            apijson.Field
	Name                    apijson.Field
	Comments                apijson.Field
	DirectoryID             apijson.Field
	DirectoryNodeID         apijson.Field
	ExternalDirectoryNodeID apijson.Field
	Provenance              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingImpersonationRegistryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingImpersonationRegistryListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingImpersonationRegistryDeleteResponse struct {
	Errors   []EmailSecurityMessage                                               `json:"errors,required"`
	Messages []EmailSecurityMessage                                               `json:"messages,required"`
	Result   AccountEmailSecuritySettingImpersonationRegistryDeleteResponseResult `json:"result,required"`
	Success  bool                                                                 `json:"success,required"`
	JSON     accountEmailSecuritySettingImpersonationRegistryDeleteResponseJSON   `json:"-"`
}

// accountEmailSecuritySettingImpersonationRegistryDeleteResponseJSON contains the
// JSON metadata for the struct
// [AccountEmailSecuritySettingImpersonationRegistryDeleteResponse]
type accountEmailSecuritySettingImpersonationRegistryDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingImpersonationRegistryDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingImpersonationRegistryDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingImpersonationRegistryDeleteResponseResult struct {
	ID   int64                                                                    `json:"id,required"`
	JSON accountEmailSecuritySettingImpersonationRegistryDeleteResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingImpersonationRegistryDeleteResponseResultJSON
// contains the JSON metadata for the struct
// [AccountEmailSecuritySettingImpersonationRegistryDeleteResponseResult]
type accountEmailSecuritySettingImpersonationRegistryDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingImpersonationRegistryDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingImpersonationRegistryDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingImpersonationRegistryNewParams struct {
	Email        param.Field[string] `json:"email,required"`
	IsEmailRegex param.Field[bool]   `json:"is_email_regex,required"`
	Name         param.Field[string] `json:"name,required"`
}

func (r AccountEmailSecuritySettingImpersonationRegistryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountEmailSecuritySettingImpersonationRegistryUpdateParams struct {
	Email        param.Field[string] `json:"email"`
	IsEmailRegex param.Field[bool]   `json:"is_email_regex"`
	Name         param.Field[string] `json:"name"`
}

func (r AccountEmailSecuritySettingImpersonationRegistryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountEmailSecuritySettingImpersonationRegistryListParams struct {
	// The sorting direction.
	Direction param.Field[SortingDirection] `query:"direction"`
	// The field to sort by.
	Order param.Field[AccountEmailSecuritySettingImpersonationRegistryListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page.
	PerPage    param.Field[int64]                                                                `query:"per_page"`
	Provenance param.Field[AccountEmailSecuritySettingImpersonationRegistryListParamsProvenance] `query:"provenance"`
	// Allows searching in multiple properties of a record simultaneously. This
	// parameter is intended for human users, not automation. Its exact behavior is
	// intentionally left unspecified and is subject to change in the future.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes
// [AccountEmailSecuritySettingImpersonationRegistryListParams]'s query parameters
// as `url.Values`.
func (r AccountEmailSecuritySettingImpersonationRegistryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The field to sort by.
type AccountEmailSecuritySettingImpersonationRegistryListParamsOrder string

const (
	AccountEmailSecuritySettingImpersonationRegistryListParamsOrderName      AccountEmailSecuritySettingImpersonationRegistryListParamsOrder = "name"
	AccountEmailSecuritySettingImpersonationRegistryListParamsOrderEmail     AccountEmailSecuritySettingImpersonationRegistryListParamsOrder = "email"
	AccountEmailSecuritySettingImpersonationRegistryListParamsOrderCreatedAt AccountEmailSecuritySettingImpersonationRegistryListParamsOrder = "created_at"
)

func (r AccountEmailSecuritySettingImpersonationRegistryListParamsOrder) IsKnown() bool {
	switch r {
	case AccountEmailSecuritySettingImpersonationRegistryListParamsOrderName, AccountEmailSecuritySettingImpersonationRegistryListParamsOrderEmail, AccountEmailSecuritySettingImpersonationRegistryListParamsOrderCreatedAt:
		return true
	}
	return false
}

type AccountEmailSecuritySettingImpersonationRegistryListParamsProvenance string

const (
	AccountEmailSecuritySettingImpersonationRegistryListParamsProvenanceA1SInternal           AccountEmailSecuritySettingImpersonationRegistryListParamsProvenance = "A1S_INTERNAL"
	AccountEmailSecuritySettingImpersonationRegistryListParamsProvenanceSnoopyCasbOffice365   AccountEmailSecuritySettingImpersonationRegistryListParamsProvenance = "SNOOPY-CASB_OFFICE_365"
	AccountEmailSecuritySettingImpersonationRegistryListParamsProvenanceSnoopyOffice365       AccountEmailSecuritySettingImpersonationRegistryListParamsProvenance = "SNOOPY-OFFICE_365"
	AccountEmailSecuritySettingImpersonationRegistryListParamsProvenanceSnoopyGoogleDirectory AccountEmailSecuritySettingImpersonationRegistryListParamsProvenance = "SNOOPY-GOOGLE_DIRECTORY"
)

func (r AccountEmailSecuritySettingImpersonationRegistryListParamsProvenance) IsKnown() bool {
	switch r {
	case AccountEmailSecuritySettingImpersonationRegistryListParamsProvenanceA1SInternal, AccountEmailSecuritySettingImpersonationRegistryListParamsProvenanceSnoopyCasbOffice365, AccountEmailSecuritySettingImpersonationRegistryListParamsProvenanceSnoopyOffice365, AccountEmailSecuritySettingImpersonationRegistryListParamsProvenanceSnoopyGoogleDirectory:
		return true
	}
	return false
}
