// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountEmailSecuritySettingAllowPolicyService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountEmailSecuritySettingAllowPolicyService] method instead.
type AccountEmailSecuritySettingAllowPolicyService struct {
	Options []option.RequestOption
}

// NewAccountEmailSecuritySettingAllowPolicyService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountEmailSecuritySettingAllowPolicyService(opts ...option.RequestOption) (r *AccountEmailSecuritySettingAllowPolicyService) {
	r = &AccountEmailSecuritySettingAllowPolicyService{}
	r.Options = opts
	return
}

// Create an email allow policy
func (r *AccountEmailSecuritySettingAllowPolicyService) New(ctx context.Context, accountID string, body AccountEmailSecuritySettingAllowPolicyNewParams, opts ...option.RequestOption) (res *AccountEmailSecuritySettingAllowPolicyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an email allow policy
func (r *AccountEmailSecuritySettingAllowPolicyService) Get(ctx context.Context, accountID string, policyID int64, opts ...option.RequestOption) (res *AccountEmailSecuritySettingAllowPolicyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies/%v", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an email allow policy
func (r *AccountEmailSecuritySettingAllowPolicyService) Update(ctx context.Context, accountID string, policyID int64, body AccountEmailSecuritySettingAllowPolicyUpdateParams, opts ...option.RequestOption) (res *AccountEmailSecuritySettingAllowPolicyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies/%v", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists, searches, and sorts an account’s email allow policies.
func (r *AccountEmailSecuritySettingAllowPolicyService) List(ctx context.Context, accountID string, query AccountEmailSecuritySettingAllowPolicyListParams, opts ...option.RequestOption) (res *AccountEmailSecuritySettingAllowPolicyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an email allow policy
func (r *AccountEmailSecuritySettingAllowPolicyService) Delete(ctx context.Context, accountID string, policyID int64, opts ...option.RequestOption) (res *AccountEmailSecuritySettingAllowPolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies/%v", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type PatternType string

const (
	PatternTypeEmail   PatternType = "EMAIL"
	PatternTypeDomain  PatternType = "DOMAIN"
	PatternTypeIP      PatternType = "IP"
	PatternTypeUnknown PatternType = "UNKNOWN"
)

func (r PatternType) IsKnown() bool {
	switch r {
	case PatternTypeEmail, PatternTypeDomain, PatternTypeIP, PatternTypeUnknown:
		return true
	}
	return false
}

type SortingDirection string

const (
	SortingDirectionAsc  SortingDirection = "asc"
	SortingDirectionDesc SortingDirection = "desc"
)

func (r SortingDirection) IsKnown() bool {
	switch r {
	case SortingDirectionAsc, SortingDirectionDesc:
		return true
	}
	return false
}

type AccountEmailSecuritySettingAllowPolicyNewResponse struct {
	Result AccountEmailSecuritySettingAllowPolicyNewResponseResult `json:"result,required"`
	JSON   accountEmailSecuritySettingAllowPolicyNewResponseJSON   `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecuritySettingAllowPolicyNewResponseJSON contains the JSON metadata
// for the struct [AccountEmailSecuritySettingAllowPolicyNewResponse]
type accountEmailSecuritySettingAllowPolicyNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingAllowPolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingAllowPolicyNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingAllowPolicyNewResponseResult struct {
	// The unique identifier for the allow policy.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Messages from this sender will be exempted from Spam, Spoof and Bulk
	// dispositions. Note: This will not exempt messages with Malicious or Suspicious
	// dispositions.
	IsAcceptableSender bool `json:"is_acceptable_sender,required"`
	// Messages to this recipient will bypass all detections.
	IsExemptRecipient bool `json:"is_exempt_recipient,required"`
	IsRegex           bool `json:"is_regex,required"`
	// Messages from this sender will bypass all detections and link following.
	IsTrustedSender bool        `json:"is_trusted_sender,required"`
	LastModified    time.Time   `json:"last_modified,required" format:"date-time"`
	Pattern         string      `json:"pattern,required"`
	PatternType     PatternType `json:"pattern_type,required"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender bool   `json:"verify_sender,required"`
	Comments     string `json:"comments,nullable"`
	// Deprecated: deprecated
	IsRecipient bool `json:"is_recipient"`
	// Deprecated: deprecated
	IsSender bool `json:"is_sender"`
	// Deprecated: deprecated
	IsSpoof bool                                                        `json:"is_spoof"`
	JSON    accountEmailSecuritySettingAllowPolicyNewResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingAllowPolicyNewResponseResultJSON contains the JSON
// metadata for the struct
// [AccountEmailSecuritySettingAllowPolicyNewResponseResult]
type accountEmailSecuritySettingAllowPolicyNewResponseResultJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	IsAcceptableSender apijson.Field
	IsExemptRecipient  apijson.Field
	IsRegex            apijson.Field
	IsTrustedSender    apijson.Field
	LastModified       apijson.Field
	Pattern            apijson.Field
	PatternType        apijson.Field
	VerifySender       apijson.Field
	Comments           apijson.Field
	IsRecipient        apijson.Field
	IsSender           apijson.Field
	IsSpoof            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingAllowPolicyNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingAllowPolicyNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingAllowPolicyGetResponse struct {
	Result AccountEmailSecuritySettingAllowPolicyGetResponseResult `json:"result,required"`
	JSON   accountEmailSecuritySettingAllowPolicyGetResponseJSON   `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecuritySettingAllowPolicyGetResponseJSON contains the JSON metadata
// for the struct [AccountEmailSecuritySettingAllowPolicyGetResponse]
type accountEmailSecuritySettingAllowPolicyGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingAllowPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingAllowPolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingAllowPolicyGetResponseResult struct {
	// The unique identifier for the allow policy.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Messages from this sender will be exempted from Spam, Spoof and Bulk
	// dispositions. Note: This will not exempt messages with Malicious or Suspicious
	// dispositions.
	IsAcceptableSender bool `json:"is_acceptable_sender,required"`
	// Messages to this recipient will bypass all detections.
	IsExemptRecipient bool `json:"is_exempt_recipient,required"`
	IsRegex           bool `json:"is_regex,required"`
	// Messages from this sender will bypass all detections and link following.
	IsTrustedSender bool        `json:"is_trusted_sender,required"`
	LastModified    time.Time   `json:"last_modified,required" format:"date-time"`
	Pattern         string      `json:"pattern,required"`
	PatternType     PatternType `json:"pattern_type,required"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender bool   `json:"verify_sender,required"`
	Comments     string `json:"comments,nullable"`
	// Deprecated: deprecated
	IsRecipient bool `json:"is_recipient"`
	// Deprecated: deprecated
	IsSender bool `json:"is_sender"`
	// Deprecated: deprecated
	IsSpoof bool                                                        `json:"is_spoof"`
	JSON    accountEmailSecuritySettingAllowPolicyGetResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingAllowPolicyGetResponseResultJSON contains the JSON
// metadata for the struct
// [AccountEmailSecuritySettingAllowPolicyGetResponseResult]
type accountEmailSecuritySettingAllowPolicyGetResponseResultJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	IsAcceptableSender apijson.Field
	IsExemptRecipient  apijson.Field
	IsRegex            apijson.Field
	IsTrustedSender    apijson.Field
	LastModified       apijson.Field
	Pattern            apijson.Field
	PatternType        apijson.Field
	VerifySender       apijson.Field
	Comments           apijson.Field
	IsRecipient        apijson.Field
	IsSender           apijson.Field
	IsSpoof            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingAllowPolicyGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingAllowPolicyGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingAllowPolicyUpdateResponse struct {
	Result AccountEmailSecuritySettingAllowPolicyUpdateResponseResult `json:"result,required"`
	JSON   accountEmailSecuritySettingAllowPolicyUpdateResponseJSON   `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecuritySettingAllowPolicyUpdateResponseJSON contains the JSON
// metadata for the struct [AccountEmailSecuritySettingAllowPolicyUpdateResponse]
type accountEmailSecuritySettingAllowPolicyUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingAllowPolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingAllowPolicyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingAllowPolicyUpdateResponseResult struct {
	// The unique identifier for the allow policy.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Messages from this sender will be exempted from Spam, Spoof and Bulk
	// dispositions. Note: This will not exempt messages with Malicious or Suspicious
	// dispositions.
	IsAcceptableSender bool `json:"is_acceptable_sender,required"`
	// Messages to this recipient will bypass all detections.
	IsExemptRecipient bool `json:"is_exempt_recipient,required"`
	IsRegex           bool `json:"is_regex,required"`
	// Messages from this sender will bypass all detections and link following.
	IsTrustedSender bool        `json:"is_trusted_sender,required"`
	LastModified    time.Time   `json:"last_modified,required" format:"date-time"`
	Pattern         string      `json:"pattern,required"`
	PatternType     PatternType `json:"pattern_type,required"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender bool   `json:"verify_sender,required"`
	Comments     string `json:"comments,nullable"`
	// Deprecated: deprecated
	IsRecipient bool `json:"is_recipient"`
	// Deprecated: deprecated
	IsSender bool `json:"is_sender"`
	// Deprecated: deprecated
	IsSpoof bool                                                           `json:"is_spoof"`
	JSON    accountEmailSecuritySettingAllowPolicyUpdateResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingAllowPolicyUpdateResponseResultJSON contains the JSON
// metadata for the struct
// [AccountEmailSecuritySettingAllowPolicyUpdateResponseResult]
type accountEmailSecuritySettingAllowPolicyUpdateResponseResultJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	IsAcceptableSender apijson.Field
	IsExemptRecipient  apijson.Field
	IsRegex            apijson.Field
	IsTrustedSender    apijson.Field
	LastModified       apijson.Field
	Pattern            apijson.Field
	PatternType        apijson.Field
	VerifySender       apijson.Field
	Comments           apijson.Field
	IsRecipient        apijson.Field
	IsSender           apijson.Field
	IsSpoof            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingAllowPolicyUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingAllowPolicyUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingAllowPolicyListResponse struct {
	Result     []AccountEmailSecuritySettingAllowPolicyListResponseResult `json:"result,required"`
	ResultInfo ResultInfoEmailSecurity                                    `json:"result_info,required"`
	JSON       accountEmailSecuritySettingAllowPolicyListResponseJSON     `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecuritySettingAllowPolicyListResponseJSON contains the JSON
// metadata for the struct [AccountEmailSecuritySettingAllowPolicyListResponse]
type accountEmailSecuritySettingAllowPolicyListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingAllowPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingAllowPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingAllowPolicyListResponseResult struct {
	// The unique identifier for the allow policy.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Messages from this sender will be exempted from Spam, Spoof and Bulk
	// dispositions. Note: This will not exempt messages with Malicious or Suspicious
	// dispositions.
	IsAcceptableSender bool `json:"is_acceptable_sender,required"`
	// Messages to this recipient will bypass all detections.
	IsExemptRecipient bool `json:"is_exempt_recipient,required"`
	IsRegex           bool `json:"is_regex,required"`
	// Messages from this sender will bypass all detections and link following.
	IsTrustedSender bool        `json:"is_trusted_sender,required"`
	LastModified    time.Time   `json:"last_modified,required" format:"date-time"`
	Pattern         string      `json:"pattern,required"`
	PatternType     PatternType `json:"pattern_type,required"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender bool   `json:"verify_sender,required"`
	Comments     string `json:"comments,nullable"`
	// Deprecated: deprecated
	IsRecipient bool `json:"is_recipient"`
	// Deprecated: deprecated
	IsSender bool `json:"is_sender"`
	// Deprecated: deprecated
	IsSpoof bool                                                         `json:"is_spoof"`
	JSON    accountEmailSecuritySettingAllowPolicyListResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingAllowPolicyListResponseResultJSON contains the JSON
// metadata for the struct
// [AccountEmailSecuritySettingAllowPolicyListResponseResult]
type accountEmailSecuritySettingAllowPolicyListResponseResultJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	IsAcceptableSender apijson.Field
	IsExemptRecipient  apijson.Field
	IsRegex            apijson.Field
	IsTrustedSender    apijson.Field
	LastModified       apijson.Field
	Pattern            apijson.Field
	PatternType        apijson.Field
	VerifySender       apijson.Field
	Comments           apijson.Field
	IsRecipient        apijson.Field
	IsSender           apijson.Field
	IsSpoof            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingAllowPolicyListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingAllowPolicyListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingAllowPolicyDeleteResponse struct {
	Result AccountEmailSecuritySettingAllowPolicyDeleteResponseResult `json:"result,required"`
	JSON   accountEmailSecuritySettingAllowPolicyDeleteResponseJSON   `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecuritySettingAllowPolicyDeleteResponseJSON contains the JSON
// metadata for the struct [AccountEmailSecuritySettingAllowPolicyDeleteResponse]
type accountEmailSecuritySettingAllowPolicyDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingAllowPolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingAllowPolicyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingAllowPolicyDeleteResponseResult struct {
	// The unique identifier for the allow policy.
	ID   int64                                                          `json:"id,required"`
	JSON accountEmailSecuritySettingAllowPolicyDeleteResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingAllowPolicyDeleteResponseResultJSON contains the JSON
// metadata for the struct
// [AccountEmailSecuritySettingAllowPolicyDeleteResponseResult]
type accountEmailSecuritySettingAllowPolicyDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingAllowPolicyDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingAllowPolicyDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingAllowPolicyNewParams struct {
	// Messages from this sender will be exempted from Spam, Spoof and Bulk
	// dispositions. Note: This will not exempt messages with Malicious or Suspicious
	// dispositions.
	IsAcceptableSender param.Field[bool] `json:"is_acceptable_sender,required"`
	// Messages to this recipient will bypass all detections.
	IsExemptRecipient param.Field[bool] `json:"is_exempt_recipient,required"`
	IsRegex           param.Field[bool] `json:"is_regex,required"`
	// Messages from this sender will bypass all detections and link following.
	IsTrustedSender param.Field[bool]        `json:"is_trusted_sender,required"`
	Pattern         param.Field[string]      `json:"pattern,required"`
	PatternType     param.Field[PatternType] `json:"pattern_type,required"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender param.Field[bool]   `json:"verify_sender,required"`
	Comments     param.Field[string] `json:"comments"`
	IsRecipient  param.Field[bool]   `json:"is_recipient"`
	IsSender     param.Field[bool]   `json:"is_sender"`
	IsSpoof      param.Field[bool]   `json:"is_spoof"`
}

func (r AccountEmailSecuritySettingAllowPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountEmailSecuritySettingAllowPolicyUpdateParams struct {
	Comments param.Field[string] `json:"comments"`
	// Messages from this sender will be exempted from Spam, Spoof and Bulk
	// dispositions. Note: This will not exempt messages with Malicious or Suspicious
	// dispositions.
	IsAcceptableSender param.Field[bool] `json:"is_acceptable_sender"`
	// Messages to this recipient will bypass all detections.
	IsExemptRecipient param.Field[bool] `json:"is_exempt_recipient"`
	IsRegex           param.Field[bool] `json:"is_regex"`
	// Messages from this sender will bypass all detections and link following.
	IsTrustedSender param.Field[bool]                                                          `json:"is_trusted_sender"`
	Pattern         param.Field[string]                                                        `json:"pattern"`
	PatternType     param.Field[AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternType] `json:"pattern_type"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender param.Field[bool] `json:"verify_sender"`
}

func (r AccountEmailSecuritySettingAllowPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternType string

const (
	AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternTypeEmail   AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternType = "EMAIL"
	AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternTypeDomain  AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternType = "DOMAIN"
	AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternTypeIP      AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternType = "IP"
	AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternTypeUnknown AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternType = "UNKNOWN"
)

func (r AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternType) IsKnown() bool {
	switch r {
	case AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternTypeEmail, AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternTypeDomain, AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternTypeIP, AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternTypeUnknown:
		return true
	}
	return false
}

type AccountEmailSecuritySettingAllowPolicyListParams struct {
	// The sorting direction.
	Direction          param.Field[SortingDirection] `query:"direction"`
	IsAcceptableSender param.Field[bool]             `query:"is_acceptable_sender"`
	IsExemptRecipient  param.Field[bool]             `query:"is_exempt_recipient"`
	IsRecipient        param.Field[bool]             `query:"is_recipient"`
	IsSender           param.Field[bool]             `query:"is_sender"`
	IsSpoof            param.Field[bool]             `query:"is_spoof"`
	IsTrustedSender    param.Field[bool]             `query:"is_trusted_sender"`
	// The field to sort by.
	Order param.Field[AccountEmailSecuritySettingAllowPolicyListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page        param.Field[int64]       `query:"page"`
	PatternType param.Field[PatternType] `query:"pattern_type"`
	// The number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Allows searching in multiple properties of a record simultaneously. This
	// parameter is intended for human users, not automation. Its exact behavior is
	// intentionally left unspecified and is subject to change in the future.
	Search       param.Field[string] `query:"search"`
	VerifySender param.Field[bool]   `query:"verify_sender"`
}

// URLQuery serializes [AccountEmailSecuritySettingAllowPolicyListParams]'s query
// parameters as `url.Values`.
func (r AccountEmailSecuritySettingAllowPolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The field to sort by.
type AccountEmailSecuritySettingAllowPolicyListParamsOrder string

const (
	AccountEmailSecuritySettingAllowPolicyListParamsOrderPattern   AccountEmailSecuritySettingAllowPolicyListParamsOrder = "pattern"
	AccountEmailSecuritySettingAllowPolicyListParamsOrderCreatedAt AccountEmailSecuritySettingAllowPolicyListParamsOrder = "created_at"
)

func (r AccountEmailSecuritySettingAllowPolicyListParamsOrder) IsKnown() bool {
	switch r {
	case AccountEmailSecuritySettingAllowPolicyListParamsOrderPattern, AccountEmailSecuritySettingAllowPolicyListParamsOrderCreatedAt:
		return true
	}
	return false
}
