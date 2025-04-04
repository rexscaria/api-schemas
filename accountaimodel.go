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

// AccountAIModelService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIModelService] method instead.
type AccountAIModelService struct {
	Options []option.RequestOption
}

// NewAccountAIModelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountAIModelService(opts ...option.RequestOption) (r *AccountAIModelService) {
	r = &AccountAIModelService{}
	r.Options = opts
	return
}

// Get Model Schema
func (r *AccountAIModelService) GetSchema(ctx context.Context, accountID string, query AccountAIModelGetSchemaParams, opts ...option.RequestOption) (res *AccountAIModelGetSchemaResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/models/schema", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Model Search
func (r *AccountAIModelService) Search(ctx context.Context, accountID string, query AccountAIModelSearchParams, opts ...option.RequestOption) (res *AccountAIModelSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/models/search", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountAIModelGetSchemaResponse struct {
	Result  interface{}                         `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    accountAIModelGetSchemaResponseJSON `json:"-"`
}

// accountAIModelGetSchemaResponseJSON contains the JSON metadata for the struct
// [AccountAIModelGetSchemaResponse]
type accountAIModelGetSchemaResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIModelGetSchemaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIModelGetSchemaResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIModelSearchResponse struct {
	Errors   []interface{}                    `json:"errors,required"`
	Messages []string                         `json:"messages,required"`
	Result   []interface{}                    `json:"result,required"`
	Success  bool                             `json:"success,required"`
	JSON     accountAIModelSearchResponseJSON `json:"-"`
}

// accountAIModelSearchResponseJSON contains the JSON metadata for the struct
// [AccountAIModelSearchResponse]
type accountAIModelSearchResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIModelSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIModelSearchResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIModelGetSchemaParams struct {
	// Model Name
	Model param.Field[string] `query:"model,required"`
}

// URLQuery serializes [AccountAIModelGetSchemaParams]'s query parameters as
// `url.Values`.
func (r AccountAIModelGetSchemaParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIModelSearchParams struct {
	// Filter by Author
	Author param.Field[string] `query:"author"`
	// Filter to hide experimental models
	HideExperimental param.Field[bool]  `query:"hide_experimental"`
	Page             param.Field[int64] `query:"page"`
	PerPage          param.Field[int64] `query:"per_page"`
	// Search
	Search param.Field[string] `query:"search"`
	// Filter by Source Id
	Source param.Field[float64] `query:"source"`
	// Filter by Task Name
	Task param.Field[string] `query:"task"`
}

// URLQuery serializes [AccountAIModelSearchParams]'s query parameters as
// `url.Values`.
func (r AccountAIModelSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
