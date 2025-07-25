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

// UserTokenService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserTokenService] method instead.
type UserTokenService struct {
	Options []option.RequestOption
}

// NewUserTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserTokenService(opts ...option.RequestOption) (r *UserTokenService) {
	r = &UserTokenService{}
	r.Options = opts
	return
}

// Create a new access token.
func (r *UserTokenService) New(ctx context.Context, body UserTokenNewParams, opts ...option.RequestOption) (res *IamSingleTokenCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get information about a specific token.
func (r *UserTokenService) Get(ctx context.Context, tokenID string, opts ...option.RequestOption) (res *IamSingleTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("user/tokens/%s", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing token.
func (r *UserTokenService) Update(ctx context.Context, tokenID string, body UserTokenUpdateParams, opts ...option.RequestOption) (res *IamSingleTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("user/tokens/%s", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all access tokens you created.
func (r *UserTokenService) List(ctx context.Context, query UserTokenListParams, opts ...option.RequestOption) (res *IamCollectionTokensResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Destroy a token.
func (r *UserTokenService) Delete(ctx context.Context, tokenID string, opts ...option.RequestOption) (res *IamAPIResponseSingleID, err error) {
	opts = append(r.Options[:], opts...)
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("user/tokens/%s", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Find all available permission groups for API Tokens
func (r *UserTokenService) ListPermissionGroups(ctx context.Context, query UserTokenListPermissionGroupsParams, opts ...option.RequestOption) (res *IamPermissionsGroupResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/tokens/permission_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Roll the token secret.
func (r *UserTokenService) Roll(ctx context.Context, tokenID string, body UserTokenRollParams, opts ...option.RequestOption) (res *IamResponseSingleValue, err error) {
	opts = append(r.Options[:], opts...)
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("user/tokens/%s/value", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Test whether a token works.
func (r *UserTokenService) Verify(ctx context.Context, opts ...option.RequestOption) (res *UserTokenVerifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/tokens/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserTokenVerifyResponse struct {
	Errors   []UserTokenVerifyResponseError   `json:"errors,required"`
	Messages []UserTokenVerifyResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success UserTokenVerifyResponseSuccess `json:"success,required"`
	Result  UserTokenVerifyResponseResult  `json:"result"`
	JSON    userTokenVerifyResponseJSON    `json:"-"`
}

// userTokenVerifyResponseJSON contains the JSON metadata for the struct
// [UserTokenVerifyResponse]
type userTokenVerifyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userTokenVerifyResponseJSON) RawJSON() string {
	return r.raw
}

type UserTokenVerifyResponseError struct {
	Code             int64                               `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           UserTokenVerifyResponseErrorsSource `json:"source"`
	JSON             userTokenVerifyResponseErrorJSON    `json:"-"`
}

// userTokenVerifyResponseErrorJSON contains the JSON metadata for the struct
// [UserTokenVerifyResponseError]
type userTokenVerifyResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserTokenVerifyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userTokenVerifyResponseErrorJSON) RawJSON() string {
	return r.raw
}

type UserTokenVerifyResponseErrorsSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    userTokenVerifyResponseErrorsSourceJSON `json:"-"`
}

// userTokenVerifyResponseErrorsSourceJSON contains the JSON metadata for the
// struct [UserTokenVerifyResponseErrorsSource]
type userTokenVerifyResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userTokenVerifyResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UserTokenVerifyResponseMessage struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           UserTokenVerifyResponseMessagesSource `json:"source"`
	JSON             userTokenVerifyResponseMessageJSON    `json:"-"`
}

// userTokenVerifyResponseMessageJSON contains the JSON metadata for the struct
// [UserTokenVerifyResponseMessage]
type userTokenVerifyResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserTokenVerifyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userTokenVerifyResponseMessageJSON) RawJSON() string {
	return r.raw
}

type UserTokenVerifyResponseMessagesSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    userTokenVerifyResponseMessagesSourceJSON `json:"-"`
}

// userTokenVerifyResponseMessagesSourceJSON contains the JSON metadata for the
// struct [UserTokenVerifyResponseMessagesSource]
type userTokenVerifyResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userTokenVerifyResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserTokenVerifyResponseSuccess bool

const (
	UserTokenVerifyResponseSuccessTrue UserTokenVerifyResponseSuccess = true
)

func (r UserTokenVerifyResponseSuccess) IsKnown() bool {
	switch r {
	case UserTokenVerifyResponseSuccessTrue:
		return true
	}
	return false
}

type UserTokenVerifyResponseResult struct {
	// Token identifier tag.
	ID string `json:"id,required"`
	// Status of the token.
	Status UserTokenVerifyResponseResultStatus `json:"status,required"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore time.Time                         `json:"not_before" format:"date-time"`
	JSON      userTokenVerifyResponseResultJSON `json:"-"`
}

// userTokenVerifyResponseResultJSON contains the JSON metadata for the struct
// [UserTokenVerifyResponseResult]
type userTokenVerifyResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	ExpiresOn   apijson.Field
	NotBefore   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenVerifyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userTokenVerifyResponseResultJSON) RawJSON() string {
	return r.raw
}

// Status of the token.
type UserTokenVerifyResponseResultStatus string

const (
	UserTokenVerifyResponseResultStatusActive   UserTokenVerifyResponseResultStatus = "active"
	UserTokenVerifyResponseResultStatusDisabled UserTokenVerifyResponseResultStatus = "disabled"
	UserTokenVerifyResponseResultStatusExpired  UserTokenVerifyResponseResultStatus = "expired"
)

func (r UserTokenVerifyResponseResultStatus) IsKnown() bool {
	switch r {
	case UserTokenVerifyResponseResultStatusActive, UserTokenVerifyResponseResultStatusDisabled, UserTokenVerifyResponseResultStatusExpired:
		return true
	}
	return false
}

type UserTokenNewParams struct {
	IamCreatePayload IamCreatePayloadParam `json:"iam_create_payload,required"`
}

func (r UserTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.IamCreatePayload)
}

type UserTokenUpdateParams struct {
	IamTokenBody IamTokenBodyParam `json:"iam_token_body,required"`
}

func (r UserTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.IamTokenBody)
}

type UserTokenListParams struct {
	// Direction to order results.
	Direction param.Field[UserTokenListParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [UserTokenListParams]'s query parameters as `url.Values`.
func (r UserTokenListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type UserTokenListParamsDirection string

const (
	UserTokenListParamsDirectionAsc  UserTokenListParamsDirection = "asc"
	UserTokenListParamsDirectionDesc UserTokenListParamsDirection = "desc"
)

func (r UserTokenListParamsDirection) IsKnown() bool {
	switch r {
	case UserTokenListParamsDirectionAsc, UserTokenListParamsDirectionDesc:
		return true
	}
	return false
}

type UserTokenListPermissionGroupsParams struct {
	// Filter by the name of the permission group. The value must be URL-encoded.
	Name param.Field[string] `query:"name"`
	// Filter by the scope of the permission group. The value must be URL-encoded.
	Scope param.Field[string] `query:"scope"`
}

// URLQuery serializes [UserTokenListPermissionGroupsParams]'s query parameters as
// `url.Values`.
func (r UserTokenListPermissionGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserTokenRollParams struct {
	Body interface{} `json:"body,required"`
}

func (r UserTokenRollParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
