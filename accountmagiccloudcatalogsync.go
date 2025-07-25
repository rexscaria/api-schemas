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

// AccountMagicCloudCatalogSyncService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicCloudCatalogSyncService] method instead.
type AccountMagicCloudCatalogSyncService struct {
	Options []option.RequestOption
}

// NewAccountMagicCloudCatalogSyncService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountMagicCloudCatalogSyncService(opts ...option.RequestOption) (r *AccountMagicCloudCatalogSyncService) {
	r = &AccountMagicCloudCatalogSyncService{}
	r.Options = opts
	return
}

// Create a new Catalog Sync (Closed Beta).
func (r *AccountMagicCloudCatalogSyncService) New(ctx context.Context, accountID string, params AccountMagicCloudCatalogSyncNewParams, opts ...option.RequestOption) (res *AccountMagicCloudCatalogSyncNewResponse, err error) {
	if params.Forwarded.Present {
		opts = append(opts, option.WithHeader("forwarded", fmt.Sprintf("%s", params.Forwarded)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Read a Catalog Sync (Closed Beta).
func (r *AccountMagicCloudCatalogSyncService) Get(ctx context.Context, accountID string, syncID string, opts ...option.RequestOption) (res *AccountMagicCloudCatalogSyncGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if syncID == "" {
		err = errors.New("missing required sync_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs/%s", accountID, syncID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Catalog Sync (Closed Beta).
func (r *AccountMagicCloudCatalogSyncService) Update(ctx context.Context, accountID string, syncID string, body AccountMagicCloudCatalogSyncUpdateParams, opts ...option.RequestOption) (res *McnUpdateCatalogSyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if syncID == "" {
		err = errors.New("missing required sync_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs/%s", accountID, syncID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List Catalog Syncs (Closed Beta).
func (r *AccountMagicCloudCatalogSyncService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountMagicCloudCatalogSyncListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Catalog Sync (Closed Beta).
func (r *AccountMagicCloudCatalogSyncService) Delete(ctx context.Context, accountID string, syncID string, body AccountMagicCloudCatalogSyncDeleteParams, opts ...option.RequestOption) (res *AccountMagicCloudCatalogSyncDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if syncID == "" {
		err = errors.New("missing required sync_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs/%s", accountID, syncID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// List prebuilt catalog sync policies (Closed Beta).
func (r *AccountMagicCloudCatalogSyncService) ListPolicies(ctx context.Context, accountID string, query AccountMagicCloudCatalogSyncListPoliciesParams, opts ...option.RequestOption) (res *AccountMagicCloudCatalogSyncListPoliciesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs/prebuilt-policies", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update a Catalog Sync (Closed Beta).
func (r *AccountMagicCloudCatalogSyncService) Patch(ctx context.Context, accountID string, syncID string, body AccountMagicCloudCatalogSyncPatchParams, opts ...option.RequestOption) (res *McnUpdateCatalogSyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if syncID == "" {
		err = errors.New("missing required sync_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs/%s", accountID, syncID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Refresh a Catalog Sync's destination by running the sync policy against latest
// resource catalog (Closed Beta).
func (r *AccountMagicCloudCatalogSyncService) Run(ctx context.Context, accountID string, syncID string, opts ...option.RequestOption) (res *AccountMagicCloudCatalogSyncRunResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if syncID == "" {
		err = errors.New("missing required sync_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs/%s/refresh", accountID, syncID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type McnCatalogSync struct {
	ID                       string                        `json:"id,required" format:"uuid"`
	Description              string                        `json:"description,required"`
	DestinationID            string                        `json:"destination_id,required" format:"uuid"`
	DestinationType          McnCatalogSyncDestinationType `json:"destination_type,required"`
	LastUserUpdateAt         string                        `json:"last_user_update_at,required"`
	Name                     string                        `json:"name,required"`
	Policy                   string                        `json:"policy,required"`
	UpdateMode               McnCatalogSyncUpdateMode      `json:"update_mode,required"`
	Errors                   map[string]McnError           `json:"errors"`
	IncludesDiscoveriesUntil string                        `json:"includes_discoveries_until"`
	LastAttemptedUpdateAt    string                        `json:"last_attempted_update_at"`
	LastSuccessfulUpdateAt   string                        `json:"last_successful_update_at"`
	JSON                     mcnCatalogSyncJSON            `json:"-"`
}

// mcnCatalogSyncJSON contains the JSON metadata for the struct [McnCatalogSync]
type mcnCatalogSyncJSON struct {
	ID                       apijson.Field
	Description              apijson.Field
	DestinationID            apijson.Field
	DestinationType          apijson.Field
	LastUserUpdateAt         apijson.Field
	Name                     apijson.Field
	Policy                   apijson.Field
	UpdateMode               apijson.Field
	Errors                   apijson.Field
	IncludesDiscoveriesUntil apijson.Field
	LastAttemptedUpdateAt    apijson.Field
	LastSuccessfulUpdateAt   apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *McnCatalogSync) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnCatalogSyncJSON) RawJSON() string {
	return r.raw
}

type McnCatalogSyncDestinationType string

const (
	McnCatalogSyncDestinationTypeNone          McnCatalogSyncDestinationType = "NONE"
	McnCatalogSyncDestinationTypeZeroTrustList McnCatalogSyncDestinationType = "ZERO_TRUST_LIST"
)

func (r McnCatalogSyncDestinationType) IsKnown() bool {
	switch r {
	case McnCatalogSyncDestinationTypeNone, McnCatalogSyncDestinationTypeZeroTrustList:
		return true
	}
	return false
}

type McnCatalogSyncUpdateMode string

const (
	McnCatalogSyncUpdateModeAuto   McnCatalogSyncUpdateMode = "AUTO"
	McnCatalogSyncUpdateModeManual McnCatalogSyncUpdateMode = "MANUAL"
)

func (r McnCatalogSyncUpdateMode) IsKnown() bool {
	switch r {
	case McnCatalogSyncUpdateModeAuto, McnCatalogSyncUpdateModeManual:
		return true
	}
	return false
}

type McnUpdateCatalogSyncRequestParam struct {
	Description param.Field[string]                   `json:"description"`
	Name        param.Field[string]                   `json:"name"`
	Policy      param.Field[string]                   `json:"policy"`
	UpdateMode  param.Field[McnCatalogSyncUpdateMode] `json:"update_mode"`
}

func (r McnUpdateCatalogSyncRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type McnUpdateCatalogSyncResponse struct {
	Errors   []McnError                       `json:"errors,required"`
	Messages []McnError                       `json:"messages,required"`
	Result   McnCatalogSync                   `json:"result,required"`
	Success  bool                             `json:"success,required"`
	JSON     mcnUpdateCatalogSyncResponseJSON `json:"-"`
}

// mcnUpdateCatalogSyncResponseJSON contains the JSON metadata for the struct
// [McnUpdateCatalogSyncResponse]
type mcnUpdateCatalogSyncResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnUpdateCatalogSyncResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnUpdateCatalogSyncResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudCatalogSyncNewResponse struct {
	Errors   []McnError                                  `json:"errors,required"`
	Messages []McnError                                  `json:"messages,required"`
	Result   McnCatalogSync                              `json:"result,required"`
	Success  bool                                        `json:"success,required"`
	JSON     accountMagicCloudCatalogSyncNewResponseJSON `json:"-"`
}

// accountMagicCloudCatalogSyncNewResponseJSON contains the JSON metadata for the
// struct [AccountMagicCloudCatalogSyncNewResponse]
type accountMagicCloudCatalogSyncNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudCatalogSyncNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudCatalogSyncNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudCatalogSyncGetResponse struct {
	Errors   []McnError                                  `json:"errors,required"`
	Messages []McnError                                  `json:"messages,required"`
	Result   McnCatalogSync                              `json:"result,required"`
	Success  bool                                        `json:"success,required"`
	JSON     accountMagicCloudCatalogSyncGetResponseJSON `json:"-"`
}

// accountMagicCloudCatalogSyncGetResponseJSON contains the JSON metadata for the
// struct [AccountMagicCloudCatalogSyncGetResponse]
type accountMagicCloudCatalogSyncGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudCatalogSyncGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudCatalogSyncGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudCatalogSyncListResponse struct {
	Errors   []McnError                                   `json:"errors,required"`
	Messages []McnError                                   `json:"messages,required"`
	Result   []McnCatalogSync                             `json:"result,required"`
	Success  bool                                         `json:"success,required"`
	JSON     accountMagicCloudCatalogSyncListResponseJSON `json:"-"`
}

// accountMagicCloudCatalogSyncListResponseJSON contains the JSON metadata for the
// struct [AccountMagicCloudCatalogSyncListResponse]
type accountMagicCloudCatalogSyncListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudCatalogSyncListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudCatalogSyncListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudCatalogSyncDeleteResponse struct {
	Errors   []McnError                                       `json:"errors,required"`
	Messages []McnError                                       `json:"messages,required"`
	Result   AccountMagicCloudCatalogSyncDeleteResponseResult `json:"result,required"`
	Success  bool                                             `json:"success,required"`
	JSON     accountMagicCloudCatalogSyncDeleteResponseJSON   `json:"-"`
}

// accountMagicCloudCatalogSyncDeleteResponseJSON contains the JSON metadata for
// the struct [AccountMagicCloudCatalogSyncDeleteResponse]
type accountMagicCloudCatalogSyncDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudCatalogSyncDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudCatalogSyncDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudCatalogSyncDeleteResponseResult struct {
	ID   string                                               `json:"id,required" format:"uuid"`
	JSON accountMagicCloudCatalogSyncDeleteResponseResultJSON `json:"-"`
}

// accountMagicCloudCatalogSyncDeleteResponseResultJSON contains the JSON metadata
// for the struct [AccountMagicCloudCatalogSyncDeleteResponseResult]
type accountMagicCloudCatalogSyncDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudCatalogSyncDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudCatalogSyncDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudCatalogSyncListPoliciesResponse struct {
	Errors   []McnError                                               `json:"errors,required"`
	Messages []McnError                                               `json:"messages,required"`
	Result   []AccountMagicCloudCatalogSyncListPoliciesResponseResult `json:"result,required"`
	Success  bool                                                     `json:"success,required"`
	JSON     accountMagicCloudCatalogSyncListPoliciesResponseJSON     `json:"-"`
}

// accountMagicCloudCatalogSyncListPoliciesResponseJSON contains the JSON metadata
// for the struct [AccountMagicCloudCatalogSyncListPoliciesResponse]
type accountMagicCloudCatalogSyncListPoliciesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudCatalogSyncListPoliciesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudCatalogSyncListPoliciesResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudCatalogSyncListPoliciesResponseResult struct {
	ApplicableDestinations []McnCatalogSyncDestinationType                            `json:"applicable_destinations,required"`
	PolicyDescription      string                                                     `json:"policy_description,required"`
	PolicyName             string                                                     `json:"policy_name,required"`
	PolicyString           string                                                     `json:"policy_string,required"`
	JSON                   accountMagicCloudCatalogSyncListPoliciesResponseResultJSON `json:"-"`
}

// accountMagicCloudCatalogSyncListPoliciesResponseResultJSON contains the JSON
// metadata for the struct [AccountMagicCloudCatalogSyncListPoliciesResponseResult]
type accountMagicCloudCatalogSyncListPoliciesResponseResultJSON struct {
	ApplicableDestinations apijson.Field
	PolicyDescription      apijson.Field
	PolicyName             apijson.Field
	PolicyString           apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountMagicCloudCatalogSyncListPoliciesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudCatalogSyncListPoliciesResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudCatalogSyncRunResponse struct {
	Errors   []McnError                                  `json:"errors,required"`
	Messages []McnError                                  `json:"messages,required"`
	Result   string                                      `json:"result,required"`
	Success  bool                                        `json:"success,required"`
	JSON     accountMagicCloudCatalogSyncRunResponseJSON `json:"-"`
}

// accountMagicCloudCatalogSyncRunResponseJSON contains the JSON metadata for the
// struct [AccountMagicCloudCatalogSyncRunResponse]
type accountMagicCloudCatalogSyncRunResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudCatalogSyncRunResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudCatalogSyncRunResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudCatalogSyncNewParams struct {
	DestinationType param.Field[McnCatalogSyncDestinationType] `json:"destination_type,required"`
	Name            param.Field[string]                        `json:"name,required"`
	UpdateMode      param.Field[McnCatalogSyncUpdateMode]      `json:"update_mode,required"`
	Description     param.Field[string]                        `json:"description"`
	Policy          param.Field[string]                        `json:"policy"`
	Forwarded       param.Field[string]                        `header:"forwarded"`
}

func (r AccountMagicCloudCatalogSyncNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicCloudCatalogSyncUpdateParams struct {
	McnUpdateCatalogSyncRequest McnUpdateCatalogSyncRequestParam `json:"mcn_update_catalog_sync_request,required"`
}

func (r AccountMagicCloudCatalogSyncUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.McnUpdateCatalogSyncRequest)
}

type AccountMagicCloudCatalogSyncDeleteParams struct {
	DeleteDestination param.Field[bool] `query:"delete_destination"`
}

// URLQuery serializes [AccountMagicCloudCatalogSyncDeleteParams]'s query
// parameters as `url.Values`.
func (r AccountMagicCloudCatalogSyncDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountMagicCloudCatalogSyncListPoliciesParams struct {
	// Specify type of destination, omit to return all.
	DestinationType param.Field[McnCatalogSyncDestinationType] `query:"destination_type"`
}

// URLQuery serializes [AccountMagicCloudCatalogSyncListPoliciesParams]'s query
// parameters as `url.Values`.
func (r AccountMagicCloudCatalogSyncListPoliciesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountMagicCloudCatalogSyncPatchParams struct {
	McnUpdateCatalogSyncRequest McnUpdateCatalogSyncRequestParam `json:"mcn_update_catalog_sync_request,required"`
}

func (r AccountMagicCloudCatalogSyncPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.McnUpdateCatalogSyncRequest)
}
