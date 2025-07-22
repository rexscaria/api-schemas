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

// AccountAccessLogScimService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessLogScimService] method instead.
type AccountAccessLogScimService struct {
	Options []option.RequestOption
}

// NewAccountAccessLogScimService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessLogScimService(opts ...option.RequestOption) (r *AccountAccessLogScimService) {
	r = &AccountAccessLogScimService{}
	r.Options = opts
	return
}

// Lists Access SCIM update logs that maintain a record of updates made to User and
// Group resources synced to Cloudflare via the System for Cross-domain Identity
// Management (SCIM).
func (r *AccountAccessLogScimService) Updates(ctx context.Context, accountID string, query AccountAccessLogScimUpdatesParams, opts ...option.RequestOption) (res *AccountAccessLogScimUpdatesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/logs/scim/updates", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountAccessLogScimUpdatesResponse struct {
	Result []AccountAccessLogScimUpdatesResponseResult `json:"result"`
	JSON   accountAccessLogScimUpdatesResponseJSON     `json:"-"`
	APIResponseCollectionAccess
}

// accountAccessLogScimUpdatesResponseJSON contains the JSON metadata for the
// struct [AccountAccessLogScimUpdatesResponse]
type accountAccessLogScimUpdatesResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessLogScimUpdatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessLogScimUpdatesResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessLogScimUpdatesResponseResult struct {
	// The unique Cloudflare-generated Id of the SCIM resource.
	CfResourceID string `json:"cf_resource_id"`
	// The error message which is generated when the status of the SCIM request is
	// 'FAILURE'.
	ErrorDescription string `json:"error_description"`
	// The unique Id of the IdP that has SCIM enabled.
	IdpID string `json:"idp_id"`
	// The IdP-generated Id of the SCIM resource.
	IdpResourceID string    `json:"idp_resource_id"`
	LoggedAt      time.Time `json:"logged_at" format:"date-time"`
	// The JSON-encoded string body of the SCIM request.
	RequestBody string `json:"request_body"`
	// The request method of the SCIM request.
	RequestMethod string `json:"request_method"`
	// The display name of the SCIM Group resource if it exists.
	ResourceGroupName string `json:"resource_group_name"`
	// The resource type of the SCIM request.
	ResourceType string `json:"resource_type"`
	// The email address of the SCIM User resource if it exists.
	ResourceUserEmail string `json:"resource_user_email" format:"email"`
	// The status of the SCIM request.
	Status string                                        `json:"status"`
	JSON   accountAccessLogScimUpdatesResponseResultJSON `json:"-"`
}

// accountAccessLogScimUpdatesResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessLogScimUpdatesResponseResult]
type accountAccessLogScimUpdatesResponseResultJSON struct {
	CfResourceID      apijson.Field
	ErrorDescription  apijson.Field
	IdpID             apijson.Field
	IdpResourceID     apijson.Field
	LoggedAt          apijson.Field
	RequestBody       apijson.Field
	RequestMethod     apijson.Field
	ResourceGroupName apijson.Field
	ResourceType      apijson.Field
	ResourceUserEmail apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountAccessLogScimUpdatesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessLogScimUpdatesResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessLogScimUpdatesParams struct {
	// The unique Id of the IdP that has SCIM enabled.
	IdpID param.Field[[]string] `query:"idp_id,required"`
	// The unique Cloudflare-generated Id of the SCIM resource.
	CfResourceID param.Field[string] `query:"cf_resource_id"`
	// The chronological order used to sort the logs.
	Direction param.Field[AccountAccessLogScimUpdatesParamsDirection] `query:"direction"`
	// The IdP-generated Id of the SCIM resource.
	IdpResourceID param.Field[string] `query:"idp_resource_id"`
	// The maximum number of update logs to retrieve.
	Limit param.Field[int64] `query:"limit"`
	// The request method of the SCIM request.
	RequestMethod param.Field[[]AccountAccessLogScimUpdatesParamsRequestMethod] `query:"request_method"`
	// The display name of the SCIM Group resource.
	ResourceGroupName param.Field[string] `query:"resource_group_name"`
	// The resource type of the SCIM request.
	ResourceType param.Field[[]AccountAccessLogScimUpdatesParamsResourceType] `query:"resource_type"`
	// The email address of the SCIM User resource.
	ResourceUserEmail param.Field[string] `query:"resource_user_email" format:"email"`
	// the timestamp of the earliest update log.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// The status of the SCIM request.
	Status param.Field[[]AccountAccessLogScimUpdatesParamsStatus] `query:"status"`
	// the timestamp of the most-recent update log.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AccountAccessLogScimUpdatesParams]'s query parameters as
// `url.Values`.
func (r AccountAccessLogScimUpdatesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The chronological order used to sort the logs.
type AccountAccessLogScimUpdatesParamsDirection string

const (
	AccountAccessLogScimUpdatesParamsDirectionDesc AccountAccessLogScimUpdatesParamsDirection = "desc"
	AccountAccessLogScimUpdatesParamsDirectionAsc  AccountAccessLogScimUpdatesParamsDirection = "asc"
)

func (r AccountAccessLogScimUpdatesParamsDirection) IsKnown() bool {
	switch r {
	case AccountAccessLogScimUpdatesParamsDirectionDesc, AccountAccessLogScimUpdatesParamsDirectionAsc:
		return true
	}
	return false
}

type AccountAccessLogScimUpdatesParamsRequestMethod string

const (
	AccountAccessLogScimUpdatesParamsRequestMethodDelete AccountAccessLogScimUpdatesParamsRequestMethod = "DELETE"
	AccountAccessLogScimUpdatesParamsRequestMethodPatch  AccountAccessLogScimUpdatesParamsRequestMethod = "PATCH"
	AccountAccessLogScimUpdatesParamsRequestMethodPost   AccountAccessLogScimUpdatesParamsRequestMethod = "POST"
	AccountAccessLogScimUpdatesParamsRequestMethodPut    AccountAccessLogScimUpdatesParamsRequestMethod = "PUT"
)

func (r AccountAccessLogScimUpdatesParamsRequestMethod) IsKnown() bool {
	switch r {
	case AccountAccessLogScimUpdatesParamsRequestMethodDelete, AccountAccessLogScimUpdatesParamsRequestMethodPatch, AccountAccessLogScimUpdatesParamsRequestMethodPost, AccountAccessLogScimUpdatesParamsRequestMethodPut:
		return true
	}
	return false
}

type AccountAccessLogScimUpdatesParamsResourceType string

const (
	AccountAccessLogScimUpdatesParamsResourceTypeUser  AccountAccessLogScimUpdatesParamsResourceType = "USER"
	AccountAccessLogScimUpdatesParamsResourceTypeGroup AccountAccessLogScimUpdatesParamsResourceType = "GROUP"
)

func (r AccountAccessLogScimUpdatesParamsResourceType) IsKnown() bool {
	switch r {
	case AccountAccessLogScimUpdatesParamsResourceTypeUser, AccountAccessLogScimUpdatesParamsResourceTypeGroup:
		return true
	}
	return false
}

type AccountAccessLogScimUpdatesParamsStatus string

const (
	AccountAccessLogScimUpdatesParamsStatusFailure AccountAccessLogScimUpdatesParamsStatus = "FAILURE"
	AccountAccessLogScimUpdatesParamsStatusSuccess AccountAccessLogScimUpdatesParamsStatus = "SUCCESS"
)

func (r AccountAccessLogScimUpdatesParamsStatus) IsKnown() bool {
	switch r {
	case AccountAccessLogScimUpdatesParamsStatusFailure, AccountAccessLogScimUpdatesParamsStatusSuccess:
		return true
	}
	return false
}
