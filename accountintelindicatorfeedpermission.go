// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountIntelIndicatorFeedPermissionService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountIntelIndicatorFeedPermissionService] method instead.
type AccountIntelIndicatorFeedPermissionService struct {
	Options []option.RequestOption
}

// NewAccountIntelIndicatorFeedPermissionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountIntelIndicatorFeedPermissionService(opts ...option.RequestOption) (r *AccountIntelIndicatorFeedPermissionService) {
	r = &AccountIntelIndicatorFeedPermissionService{}
	r.Options = opts
	return
}

// Grant permission to indicator feed
func (r *AccountIntelIndicatorFeedPermissionService) AddPermission(ctx context.Context, accountID string, body AccountIntelIndicatorFeedPermissionAddPermissionParams, opts ...option.RequestOption) (res *PermissionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/add", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List indicator feed permissions
func (r *AccountIntelIndicatorFeedPermissionService) ListPermissions(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountIntelIndicatorFeedPermissionListPermissionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/view", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Revoke permission to indicator feed
func (r *AccountIntelIndicatorFeedPermissionService) RemovePermission(ctx context.Context, accountID string, body AccountIntelIndicatorFeedPermissionRemovePermissionParams, opts ...option.RequestOption) (res *PermissionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/remove", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type PermissionsResponse struct {
	Result PermissionsResponseResult `json:"result"`
	JSON   permissionsResponseJSON   `json:"-"`
	SingleResponseFeed
}

// permissionsResponseJSON contains the JSON metadata for the struct
// [PermissionsResponse]
type permissionsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionsResponseJSON) RawJSON() string {
	return r.raw
}

type PermissionsResponseResult struct {
	// Whether the update succeeded or not
	Success bool                          `json:"success"`
	JSON    permissionsResponseResultJSON `json:"-"`
}

// permissionsResponseResultJSON contains the JSON metadata for the struct
// [PermissionsResponseResult]
type permissionsResponseResultJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RequestParam struct {
	// The Cloudflare account tag of the account to change permissions on
	AccountTag param.Field[string] `json:"account_tag"`
	// The ID of the feed to add/remove permissions on
	FeedID param.Field[int64] `json:"feed_id"`
}

func (r RequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountIntelIndicatorFeedPermissionListPermissionsResponse struct {
	Result []AccountIntelIndicatorFeedPermissionListPermissionsResponseResult `json:"result"`
	JSON   accountIntelIndicatorFeedPermissionListPermissionsResponseJSON     `json:"-"`
	CommonResponseIndicatorFeeds
}

// accountIntelIndicatorFeedPermissionListPermissionsResponseJSON contains the JSON
// metadata for the struct
// [AccountIntelIndicatorFeedPermissionListPermissionsResponse]
type accountIntelIndicatorFeedPermissionListPermissionsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedPermissionListPermissionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedPermissionListPermissionsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedPermissionListPermissionsResponseResult struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The description of the example test
	Description string `json:"description"`
	// Whether the indicator feed can be attributed to a provider
	IsAttributable bool `json:"is_attributable"`
	// Whether the indicator feed can be downloaded
	IsDownloadable bool `json:"is_downloadable"`
	// Whether the indicator feed is exposed to customers
	IsPublic bool `json:"is_public"`
	// The name of the indicator feed
	Name string                                                               `json:"name"`
	JSON accountIntelIndicatorFeedPermissionListPermissionsResponseResultJSON `json:"-"`
}

// accountIntelIndicatorFeedPermissionListPermissionsResponseResultJSON contains
// the JSON metadata for the struct
// [AccountIntelIndicatorFeedPermissionListPermissionsResponseResult]
type accountIntelIndicatorFeedPermissionListPermissionsResponseResultJSON struct {
	ID             apijson.Field
	Description    apijson.Field
	IsAttributable apijson.Field
	IsDownloadable apijson.Field
	IsPublic       apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedPermissionListPermissionsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedPermissionListPermissionsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedPermissionAddPermissionParams struct {
	Request RequestParam `json:"request,required"`
}

func (r AccountIntelIndicatorFeedPermissionAddPermissionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Request)
}

type AccountIntelIndicatorFeedPermissionRemovePermissionParams struct {
	Request RequestParam `json:"request,required"`
}

func (r AccountIntelIndicatorFeedPermissionRemovePermissionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Request)
}
