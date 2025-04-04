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
func (r *UserTokenService) Delete(ctx context.Context, tokenID string, body UserTokenDeleteParams, opts ...option.RequestOption) (res *IamAPIResponseSingleID, err error) {
	opts = append(r.Options[:], opts...)
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("user/tokens/%s", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Find all available permission groups for API Tokens
func (r *UserTokenService) ListPermissionGroups(ctx context.Context, opts ...option.RequestOption) (res *IamPermissionsGroupResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/tokens/permission_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *UserTokenService) Verify(ctx context.Context, opts ...option.RequestOption) (res *IamResponseSingleSegment, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/tokens/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
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

type UserTokenDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r UserTokenDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type UserTokenRollParams struct {
	Body interface{} `json:"body,required"`
}

func (r UserTokenRollParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
