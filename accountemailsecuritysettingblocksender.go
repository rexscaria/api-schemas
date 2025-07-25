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

// AccountEmailSecuritySettingBlockSenderService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountEmailSecuritySettingBlockSenderService] method instead.
type AccountEmailSecuritySettingBlockSenderService struct {
	Options []option.RequestOption
}

// NewAccountEmailSecuritySettingBlockSenderService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountEmailSecuritySettingBlockSenderService(opts ...option.RequestOption) (r *AccountEmailSecuritySettingBlockSenderService) {
	r = &AccountEmailSecuritySettingBlockSenderService{}
	r.Options = opts
	return
}

// Create a blocked email sender
func (r *AccountEmailSecuritySettingBlockSenderService) New(ctx context.Context, accountID string, body AccountEmailSecuritySettingBlockSenderNewParams, opts ...option.RequestOption) (res *AccountEmailSecuritySettingBlockSenderNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/block_senders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a blocked email sender
func (r *AccountEmailSecuritySettingBlockSenderService) Get(ctx context.Context, accountID string, patternID int64, opts ...option.RequestOption) (res *AccountEmailSecuritySettingBlockSenderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/block_senders/%v", accountID, patternID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a blocked email sender
func (r *AccountEmailSecuritySettingBlockSenderService) Update(ctx context.Context, accountID string, patternID int64, body AccountEmailSecuritySettingBlockSenderUpdateParams, opts ...option.RequestOption) (res *AccountEmailSecuritySettingBlockSenderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/block_senders/%v", accountID, patternID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List blocked email senders
func (r *AccountEmailSecuritySettingBlockSenderService) List(ctx context.Context, accountID string, query AccountEmailSecuritySettingBlockSenderListParams, opts ...option.RequestOption) (res *AccountEmailSecuritySettingBlockSenderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/block_senders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a blocked email sender
func (r *AccountEmailSecuritySettingBlockSenderService) Delete(ctx context.Context, accountID string, patternID int64, opts ...option.RequestOption) (res *AccountEmailSecuritySettingBlockSenderDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/block_senders/%v", accountID, patternID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountEmailSecuritySettingBlockSenderNewResponse struct {
	Errors   []EmailSecurityMessage                                  `json:"errors,required"`
	Messages []EmailSecurityMessage                                  `json:"messages,required"`
	Result   AccountEmailSecuritySettingBlockSenderNewResponseResult `json:"result,required"`
	Success  bool                                                    `json:"success,required"`
	JSON     accountEmailSecuritySettingBlockSenderNewResponseJSON   `json:"-"`
}

// accountEmailSecuritySettingBlockSenderNewResponseJSON contains the JSON metadata
// for the struct [AccountEmailSecuritySettingBlockSenderNewResponse]
type accountEmailSecuritySettingBlockSenderNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingBlockSenderNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingBlockSenderNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingBlockSenderNewResponseResult struct {
	// The unique identifier for the allow policy.
	ID           int64                                                       `json:"id,required"`
	CreatedAt    time.Time                                                   `json:"created_at,required" format:"date-time"`
	IsRegex      bool                                                        `json:"is_regex,required"`
	LastModified time.Time                                                   `json:"last_modified,required" format:"date-time"`
	Pattern      string                                                      `json:"pattern,required"`
	PatternType  PatternType                                                 `json:"pattern_type,required"`
	Comments     string                                                      `json:"comments,nullable"`
	JSON         accountEmailSecuritySettingBlockSenderNewResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingBlockSenderNewResponseResultJSON contains the JSON
// metadata for the struct
// [AccountEmailSecuritySettingBlockSenderNewResponseResult]
type accountEmailSecuritySettingBlockSenderNewResponseResultJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRegex      apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingBlockSenderNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingBlockSenderNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingBlockSenderGetResponse struct {
	Errors   []EmailSecurityMessage                                  `json:"errors,required"`
	Messages []EmailSecurityMessage                                  `json:"messages,required"`
	Result   AccountEmailSecuritySettingBlockSenderGetResponseResult `json:"result,required"`
	Success  bool                                                    `json:"success,required"`
	JSON     accountEmailSecuritySettingBlockSenderGetResponseJSON   `json:"-"`
}

// accountEmailSecuritySettingBlockSenderGetResponseJSON contains the JSON metadata
// for the struct [AccountEmailSecuritySettingBlockSenderGetResponse]
type accountEmailSecuritySettingBlockSenderGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingBlockSenderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingBlockSenderGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingBlockSenderGetResponseResult struct {
	// The unique identifier for the allow policy.
	ID           int64                                                       `json:"id,required"`
	CreatedAt    time.Time                                                   `json:"created_at,required" format:"date-time"`
	IsRegex      bool                                                        `json:"is_regex,required"`
	LastModified time.Time                                                   `json:"last_modified,required" format:"date-time"`
	Pattern      string                                                      `json:"pattern,required"`
	PatternType  PatternType                                                 `json:"pattern_type,required"`
	Comments     string                                                      `json:"comments,nullable"`
	JSON         accountEmailSecuritySettingBlockSenderGetResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingBlockSenderGetResponseResultJSON contains the JSON
// metadata for the struct
// [AccountEmailSecuritySettingBlockSenderGetResponseResult]
type accountEmailSecuritySettingBlockSenderGetResponseResultJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRegex      apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingBlockSenderGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingBlockSenderGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingBlockSenderUpdateResponse struct {
	Errors   []EmailSecurityMessage                                     `json:"errors,required"`
	Messages []EmailSecurityMessage                                     `json:"messages,required"`
	Result   AccountEmailSecuritySettingBlockSenderUpdateResponseResult `json:"result,required"`
	Success  bool                                                       `json:"success,required"`
	JSON     accountEmailSecuritySettingBlockSenderUpdateResponseJSON   `json:"-"`
}

// accountEmailSecuritySettingBlockSenderUpdateResponseJSON contains the JSON
// metadata for the struct [AccountEmailSecuritySettingBlockSenderUpdateResponse]
type accountEmailSecuritySettingBlockSenderUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingBlockSenderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingBlockSenderUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingBlockSenderUpdateResponseResult struct {
	// The unique identifier for the allow policy.
	ID           int64                                                          `json:"id,required"`
	CreatedAt    time.Time                                                      `json:"created_at,required" format:"date-time"`
	IsRegex      bool                                                           `json:"is_regex,required"`
	LastModified time.Time                                                      `json:"last_modified,required" format:"date-time"`
	Pattern      string                                                         `json:"pattern,required"`
	PatternType  PatternType                                                    `json:"pattern_type,required"`
	Comments     string                                                         `json:"comments,nullable"`
	JSON         accountEmailSecuritySettingBlockSenderUpdateResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingBlockSenderUpdateResponseResultJSON contains the JSON
// metadata for the struct
// [AccountEmailSecuritySettingBlockSenderUpdateResponseResult]
type accountEmailSecuritySettingBlockSenderUpdateResponseResultJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRegex      apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingBlockSenderUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingBlockSenderUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingBlockSenderListResponse struct {
	Errors     []EmailSecurityMessage                                     `json:"errors,required"`
	Messages   []EmailSecurityMessage                                     `json:"messages,required"`
	Result     []AccountEmailSecuritySettingBlockSenderListResponseResult `json:"result,required"`
	ResultInfo ResultInfoEmailSecurity                                    `json:"result_info,required"`
	Success    bool                                                       `json:"success,required"`
	JSON       accountEmailSecuritySettingBlockSenderListResponseJSON     `json:"-"`
}

// accountEmailSecuritySettingBlockSenderListResponseJSON contains the JSON
// metadata for the struct [AccountEmailSecuritySettingBlockSenderListResponse]
type accountEmailSecuritySettingBlockSenderListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingBlockSenderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingBlockSenderListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingBlockSenderListResponseResult struct {
	// The unique identifier for the allow policy.
	ID           int64                                                        `json:"id,required"`
	CreatedAt    time.Time                                                    `json:"created_at,required" format:"date-time"`
	IsRegex      bool                                                         `json:"is_regex,required"`
	LastModified time.Time                                                    `json:"last_modified,required" format:"date-time"`
	Pattern      string                                                       `json:"pattern,required"`
	PatternType  PatternType                                                  `json:"pattern_type,required"`
	Comments     string                                                       `json:"comments,nullable"`
	JSON         accountEmailSecuritySettingBlockSenderListResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingBlockSenderListResponseResultJSON contains the JSON
// metadata for the struct
// [AccountEmailSecuritySettingBlockSenderListResponseResult]
type accountEmailSecuritySettingBlockSenderListResponseResultJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRegex      apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingBlockSenderListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingBlockSenderListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingBlockSenderDeleteResponse struct {
	Errors   []EmailSecurityMessage                                     `json:"errors,required"`
	Messages []EmailSecurityMessage                                     `json:"messages,required"`
	Result   AccountEmailSecuritySettingBlockSenderDeleteResponseResult `json:"result,required"`
	Success  bool                                                       `json:"success,required"`
	JSON     accountEmailSecuritySettingBlockSenderDeleteResponseJSON   `json:"-"`
}

// accountEmailSecuritySettingBlockSenderDeleteResponseJSON contains the JSON
// metadata for the struct [AccountEmailSecuritySettingBlockSenderDeleteResponse]
type accountEmailSecuritySettingBlockSenderDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingBlockSenderDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingBlockSenderDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingBlockSenderDeleteResponseResult struct {
	// The unique identifier for the allow policy.
	ID   int64                                                          `json:"id,required"`
	JSON accountEmailSecuritySettingBlockSenderDeleteResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingBlockSenderDeleteResponseResultJSON contains the JSON
// metadata for the struct
// [AccountEmailSecuritySettingBlockSenderDeleteResponseResult]
type accountEmailSecuritySettingBlockSenderDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingBlockSenderDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingBlockSenderDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingBlockSenderNewParams struct {
	IsRegex     param.Field[bool]        `json:"is_regex,required"`
	Pattern     param.Field[string]      `json:"pattern,required"`
	PatternType param.Field[PatternType] `json:"pattern_type,required"`
	Comments    param.Field[string]      `json:"comments"`
}

func (r AccountEmailSecuritySettingBlockSenderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountEmailSecuritySettingBlockSenderUpdateParams struct {
	Comments    param.Field[string]      `json:"comments"`
	IsRegex     param.Field[bool]        `json:"is_regex"`
	Pattern     param.Field[string]      `json:"pattern"`
	PatternType param.Field[PatternType] `json:"pattern_type"`
}

func (r AccountEmailSecuritySettingBlockSenderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountEmailSecuritySettingBlockSenderListParams struct {
	// The sorting direction.
	Direction param.Field[SortingDirection] `query:"direction"`
	// The field to sort by.
	Order param.Field[AccountEmailSecuritySettingBlockSenderListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page        param.Field[int64]       `query:"page"`
	Pattern     param.Field[string]      `query:"pattern"`
	PatternType param.Field[PatternType] `query:"pattern_type"`
	// The number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Allows searching in multiple properties of a record simultaneously. This
	// parameter is intended for human users, not automation. Its exact behavior is
	// intentionally left unspecified and is subject to change in the future.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccountEmailSecuritySettingBlockSenderListParams]'s query
// parameters as `url.Values`.
func (r AccountEmailSecuritySettingBlockSenderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The field to sort by.
type AccountEmailSecuritySettingBlockSenderListParamsOrder string

const (
	AccountEmailSecuritySettingBlockSenderListParamsOrderPattern   AccountEmailSecuritySettingBlockSenderListParamsOrder = "pattern"
	AccountEmailSecuritySettingBlockSenderListParamsOrderCreatedAt AccountEmailSecuritySettingBlockSenderListParamsOrder = "created_at"
)

func (r AccountEmailSecuritySettingBlockSenderListParamsOrder) IsKnown() bool {
	switch r {
	case AccountEmailSecuritySettingBlockSenderListParamsOrderPattern, AccountEmailSecuritySettingBlockSenderListParamsOrderCreatedAt:
		return true
	}
	return false
}
