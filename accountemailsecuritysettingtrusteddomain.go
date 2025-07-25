// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountEmailSecuritySettingTrustedDomainService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountEmailSecuritySettingTrustedDomainService] method instead.
type AccountEmailSecuritySettingTrustedDomainService struct {
	Options []option.RequestOption
}

// NewAccountEmailSecuritySettingTrustedDomainService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountEmailSecuritySettingTrustedDomainService(opts ...option.RequestOption) (r *AccountEmailSecuritySettingTrustedDomainService) {
	r = &AccountEmailSecuritySettingTrustedDomainService{}
	r.Options = opts
	return
}

// Create a trusted email domain
func (r *AccountEmailSecuritySettingTrustedDomainService) New(ctx context.Context, accountID string, body AccountEmailSecuritySettingTrustedDomainNewParams, opts ...option.RequestOption) (res *AccountEmailSecuritySettingTrustedDomainNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/trusted_domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a trusted email domain
func (r *AccountEmailSecuritySettingTrustedDomainService) Get(ctx context.Context, accountID string, trustedDomainID int64, opts ...option.RequestOption) (res *AccountEmailSecuritySettingTrustedDomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/trusted_domains/%v", accountID, trustedDomainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a trusted email domain
func (r *AccountEmailSecuritySettingTrustedDomainService) Update(ctx context.Context, accountID string, trustedDomainID int64, body AccountEmailSecuritySettingTrustedDomainUpdateParams, opts ...option.RequestOption) (res *AccountEmailSecuritySettingTrustedDomainUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/trusted_domains/%v", accountID, trustedDomainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists, searches, and sorts an account’s trusted email domains.
func (r *AccountEmailSecuritySettingTrustedDomainService) List(ctx context.Context, accountID string, query AccountEmailSecuritySettingTrustedDomainListParams, opts ...option.RequestOption) (res *AccountEmailSecuritySettingTrustedDomainListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/trusted_domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a trusted email domain
func (r *AccountEmailSecuritySettingTrustedDomainService) Delete(ctx context.Context, accountID string, trustedDomainID int64, opts ...option.RequestOption) (res *AccountEmailSecuritySettingTrustedDomainDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/trusted_domains/%v", accountID, trustedDomainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CreateTrustedDomainParam struct {
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent param.Field[bool] `json:"is_recent,required"`
	IsRegex  param.Field[bool] `json:"is_regex,required"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity param.Field[bool]   `json:"is_similarity,required"`
	Pattern      param.Field[string] `json:"pattern,required"`
	Comments     param.Field[string] `json:"comments"`
}

func (r CreateTrustedDomainParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateTrustedDomainParam) implementsAccountEmailSecuritySettingTrustedDomainNewParamsBodyUnion() {
}

type TrustedDomain struct {
	// The unique identifier for the trusted domain.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent bool `json:"is_recent,required"`
	IsRegex  bool `json:"is_regex,required"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity bool              `json:"is_similarity,required"`
	LastModified time.Time         `json:"last_modified,required" format:"date-time"`
	Pattern      string            `json:"pattern,required"`
	Comments     string            `json:"comments,nullable"`
	JSON         trustedDomainJSON `json:"-"`
}

// trustedDomainJSON contains the JSON metadata for the struct [TrustedDomain]
type trustedDomainJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRecent     apijson.Field
	IsRegex      apijson.Field
	IsSimilarity apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TrustedDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trustedDomainJSON) RawJSON() string {
	return r.raw
}

func (r TrustedDomain) implementsAccountEmailSecuritySettingTrustedDomainNewResponseResultUnion() {}

type AccountEmailSecuritySettingTrustedDomainNewResponse struct {
	Errors   []EmailSecurityMessage                                         `json:"errors,required"`
	Messages []EmailSecurityMessage                                         `json:"messages,required"`
	Result   AccountEmailSecuritySettingTrustedDomainNewResponseResultUnion `json:"result,required"`
	Success  bool                                                           `json:"success,required"`
	JSON     accountEmailSecuritySettingTrustedDomainNewResponseJSON        `json:"-"`
}

// accountEmailSecuritySettingTrustedDomainNewResponseJSON contains the JSON
// metadata for the struct [AccountEmailSecuritySettingTrustedDomainNewResponse]
type accountEmailSecuritySettingTrustedDomainNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingTrustedDomainNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingTrustedDomainNewResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [TrustedDomain] or
// [AccountEmailSecuritySettingTrustedDomainNewResponseResultArray].
type AccountEmailSecuritySettingTrustedDomainNewResponseResultUnion interface {
	implementsAccountEmailSecuritySettingTrustedDomainNewResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountEmailSecuritySettingTrustedDomainNewResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TrustedDomain{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountEmailSecuritySettingTrustedDomainNewResponseResultArray{}),
		},
	)
}

type AccountEmailSecuritySettingTrustedDomainNewResponseResultArray []TrustedDomain

func (r AccountEmailSecuritySettingTrustedDomainNewResponseResultArray) implementsAccountEmailSecuritySettingTrustedDomainNewResponseResultUnion() {
}

type AccountEmailSecuritySettingTrustedDomainGetResponse struct {
	Errors   []EmailSecurityMessage                                    `json:"errors,required"`
	Messages []EmailSecurityMessage                                    `json:"messages,required"`
	Result   AccountEmailSecuritySettingTrustedDomainGetResponseResult `json:"result,required"`
	Success  bool                                                      `json:"success,required"`
	JSON     accountEmailSecuritySettingTrustedDomainGetResponseJSON   `json:"-"`
}

// accountEmailSecuritySettingTrustedDomainGetResponseJSON contains the JSON
// metadata for the struct [AccountEmailSecuritySettingTrustedDomainGetResponse]
type accountEmailSecuritySettingTrustedDomainGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingTrustedDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingTrustedDomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingTrustedDomainGetResponseResult struct {
	// The unique identifier for the trusted domain.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent bool `json:"is_recent,required"`
	IsRegex  bool `json:"is_regex,required"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity bool                                                          `json:"is_similarity,required"`
	LastModified time.Time                                                     `json:"last_modified,required" format:"date-time"`
	Pattern      string                                                        `json:"pattern,required"`
	Comments     string                                                        `json:"comments,nullable"`
	JSON         accountEmailSecuritySettingTrustedDomainGetResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingTrustedDomainGetResponseResultJSON contains the JSON
// metadata for the struct
// [AccountEmailSecuritySettingTrustedDomainGetResponseResult]
type accountEmailSecuritySettingTrustedDomainGetResponseResultJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRecent     apijson.Field
	IsRegex      apijson.Field
	IsSimilarity apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingTrustedDomainGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingTrustedDomainGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingTrustedDomainUpdateResponse struct {
	Errors   []EmailSecurityMessage                                       `json:"errors,required"`
	Messages []EmailSecurityMessage                                       `json:"messages,required"`
	Result   AccountEmailSecuritySettingTrustedDomainUpdateResponseResult `json:"result,required"`
	Success  bool                                                         `json:"success,required"`
	JSON     accountEmailSecuritySettingTrustedDomainUpdateResponseJSON   `json:"-"`
}

// accountEmailSecuritySettingTrustedDomainUpdateResponseJSON contains the JSON
// metadata for the struct [AccountEmailSecuritySettingTrustedDomainUpdateResponse]
type accountEmailSecuritySettingTrustedDomainUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingTrustedDomainUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingTrustedDomainUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingTrustedDomainUpdateResponseResult struct {
	// The unique identifier for the trusted domain.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent bool `json:"is_recent,required"`
	IsRegex  bool `json:"is_regex,required"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity bool                                                             `json:"is_similarity,required"`
	LastModified time.Time                                                        `json:"last_modified,required" format:"date-time"`
	Pattern      string                                                           `json:"pattern,required"`
	Comments     string                                                           `json:"comments,nullable"`
	JSON         accountEmailSecuritySettingTrustedDomainUpdateResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingTrustedDomainUpdateResponseResultJSON contains the
// JSON metadata for the struct
// [AccountEmailSecuritySettingTrustedDomainUpdateResponseResult]
type accountEmailSecuritySettingTrustedDomainUpdateResponseResultJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRecent     apijson.Field
	IsRegex      apijson.Field
	IsSimilarity apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingTrustedDomainUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingTrustedDomainUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingTrustedDomainListResponse struct {
	Errors     []EmailSecurityMessage                                   `json:"errors,required"`
	Messages   []EmailSecurityMessage                                   `json:"messages,required"`
	Result     []TrustedDomain                                          `json:"result,required"`
	ResultInfo ResultInfoEmailSecurity                                  `json:"result_info,required"`
	Success    bool                                                     `json:"success,required"`
	JSON       accountEmailSecuritySettingTrustedDomainListResponseJSON `json:"-"`
}

// accountEmailSecuritySettingTrustedDomainListResponseJSON contains the JSON
// metadata for the struct [AccountEmailSecuritySettingTrustedDomainListResponse]
type accountEmailSecuritySettingTrustedDomainListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingTrustedDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingTrustedDomainListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingTrustedDomainDeleteResponse struct {
	Errors   []EmailSecurityMessage                                       `json:"errors,required"`
	Messages []EmailSecurityMessage                                       `json:"messages,required"`
	Result   AccountEmailSecuritySettingTrustedDomainDeleteResponseResult `json:"result,required"`
	Success  bool                                                         `json:"success,required"`
	JSON     accountEmailSecuritySettingTrustedDomainDeleteResponseJSON   `json:"-"`
}

// accountEmailSecuritySettingTrustedDomainDeleteResponseJSON contains the JSON
// metadata for the struct [AccountEmailSecuritySettingTrustedDomainDeleteResponse]
type accountEmailSecuritySettingTrustedDomainDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingTrustedDomainDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingTrustedDomainDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingTrustedDomainDeleteResponseResult struct {
	// The unique identifier for the trusted domain.
	ID   int64                                                            `json:"id,required"`
	JSON accountEmailSecuritySettingTrustedDomainDeleteResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingTrustedDomainDeleteResponseResultJSON contains the
// JSON metadata for the struct
// [AccountEmailSecuritySettingTrustedDomainDeleteResponseResult]
type accountEmailSecuritySettingTrustedDomainDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingTrustedDomainDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingTrustedDomainDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingTrustedDomainNewParams struct {
	Body AccountEmailSecuritySettingTrustedDomainNewParamsBodyUnion `json:"body,required"`
}

func (r AccountEmailSecuritySettingTrustedDomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Satisfied by [CreateTrustedDomainParam],
// [AccountEmailSecuritySettingTrustedDomainNewParamsBodyArray].
type AccountEmailSecuritySettingTrustedDomainNewParamsBodyUnion interface {
	implementsAccountEmailSecuritySettingTrustedDomainNewParamsBodyUnion()
}

type AccountEmailSecuritySettingTrustedDomainNewParamsBodyArray []CreateTrustedDomainParam

func (r AccountEmailSecuritySettingTrustedDomainNewParamsBodyArray) implementsAccountEmailSecuritySettingTrustedDomainNewParamsBodyUnion() {
}

type AccountEmailSecuritySettingTrustedDomainUpdateParams struct {
	Comments param.Field[string] `json:"comments"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent param.Field[bool] `json:"is_recent"`
	IsRegex  param.Field[bool] `json:"is_regex"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity param.Field[bool]   `json:"is_similarity"`
	Pattern      param.Field[string] `json:"pattern"`
}

func (r AccountEmailSecuritySettingTrustedDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountEmailSecuritySettingTrustedDomainListParams struct {
	// The sorting direction.
	Direction    param.Field[SortingDirection] `query:"direction"`
	IsRecent     param.Field[bool]             `query:"is_recent"`
	IsSimilarity param.Field[bool]             `query:"is_similarity"`
	// The field to sort by.
	Order param.Field[AccountEmailSecuritySettingTrustedDomainListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page    param.Field[int64]  `query:"page"`
	Pattern param.Field[string] `query:"pattern"`
	// The number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Allows searching in multiple properties of a record simultaneously. This
	// parameter is intended for human users, not automation. Its exact behavior is
	// intentionally left unspecified and is subject to change in the future.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccountEmailSecuritySettingTrustedDomainListParams]'s query
// parameters as `url.Values`.
func (r AccountEmailSecuritySettingTrustedDomainListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The field to sort by.
type AccountEmailSecuritySettingTrustedDomainListParamsOrder string

const (
	AccountEmailSecuritySettingTrustedDomainListParamsOrderPattern   AccountEmailSecuritySettingTrustedDomainListParamsOrder = "pattern"
	AccountEmailSecuritySettingTrustedDomainListParamsOrderCreatedAt AccountEmailSecuritySettingTrustedDomainListParamsOrder = "created_at"
)

func (r AccountEmailSecuritySettingTrustedDomainListParamsOrder) IsKnown() bool {
	switch r {
	case AccountEmailSecuritySettingTrustedDomainListParamsOrderPattern, AccountEmailSecuritySettingTrustedDomainListParamsOrderCreatedAt:
		return true
	}
	return false
}
