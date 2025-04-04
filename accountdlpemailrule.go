// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/stainless-sdks/cf-rex-go/shared"
	"github.com/tidwall/gjson"
)

// AccountDlpEmailRuleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDlpEmailRuleService] method instead.
type AccountDlpEmailRuleService struct {
	Options []option.RequestOption
}

// NewAccountDlpEmailRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDlpEmailRuleService(opts ...option.RequestOption) (r *AccountDlpEmailRuleService) {
	r = &AccountDlpEmailRuleService{}
	r.Options = opts
	return
}

// Create email scanner rule
func (r *AccountDlpEmailRuleService) New(ctx context.Context, accountID string, body AccountDlpEmailRuleNewParams, opts ...option.RequestOption) (res *AccountDlpEmailRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/rules", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an email scanner rule
func (r *AccountDlpEmailRuleService) Get(ctx context.Context, accountID string, ruleID string, opts ...option.RequestOption) (res *AccountDlpEmailRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/rules/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update email scanner rule
func (r *AccountDlpEmailRuleService) Update(ctx context.Context, accountID string, ruleID string, body AccountDlpEmailRuleUpdateParams, opts ...option.RequestOption) (res *AccountDlpEmailRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/rules/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all email scanner rules for an account.
func (r *AccountDlpEmailRuleService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountDlpEmailRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/rules", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete email scanner rule
func (r *AccountDlpEmailRuleService) Delete(ctx context.Context, accountID string, ruleID string, opts ...option.RequestOption) (res *AccountDlpEmailRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/rules/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Update email scanner rule priorities
func (r *AccountDlpEmailRuleService) UpdatePriorities(ctx context.Context, accountID string, body AccountDlpEmailRuleUpdatePrioritiesParams, opts ...option.RequestOption) (res *AccountDlpEmailRuleUpdatePrioritiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/rules", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type CreateEmailRuleParam struct {
	Action param.Field[EmailRuleActionRuleParam] `json:"action,required"`
	// Rule is triggered if all conditions match
	Conditions  param.Field[[]EmailRuleConditionParam] `json:"conditions,required"`
	Enabled     param.Field[bool]                      `json:"enabled,required"`
	Name        param.Field[string]                    `json:"name,required"`
	Description param.Field[string]                    `json:"description"`
}

func (r CreateEmailRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmailRule struct {
	Action EmailRuleActionRule `json:"action,required"`
	// Rule is triggered if all conditions match
	Conditions  []EmailRuleCondition `json:"conditions,required"`
	CreatedAt   time.Time            `json:"created_at,required" format:"date-time"`
	Enabled     bool                 `json:"enabled,required"`
	Name        string               `json:"name,required"`
	Priority    int64                `json:"priority,required"`
	RuleID      string               `json:"rule_id,required" format:"uuid"`
	UpdatedAt   time.Time            `json:"updated_at,required" format:"date-time"`
	Description string               `json:"description,nullable"`
	JSON        emailRuleJSON        `json:"-"`
}

// emailRuleJSON contains the JSON metadata for the struct [EmailRule]
type emailRuleJSON struct {
	Action      apijson.Field
	Conditions  apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	RuleID      apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRuleJSON) RawJSON() string {
	return r.raw
}

type EmailRuleActionRule struct {
	Action  EmailRuleActionRuleAction `json:"action,required"`
	Message string                    `json:"message,nullable"`
	JSON    emailRuleActionRuleJSON   `json:"-"`
}

// emailRuleActionRuleJSON contains the JSON metadata for the struct
// [EmailRuleActionRule]
type emailRuleActionRuleJSON struct {
	Action      apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRuleActionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRuleActionRuleJSON) RawJSON() string {
	return r.raw
}

type EmailRuleActionRuleAction string

const (
	EmailRuleActionRuleActionBlock EmailRuleActionRuleAction = "Block"
)

func (r EmailRuleActionRuleAction) IsKnown() bool {
	switch r {
	case EmailRuleActionRuleActionBlock:
		return true
	}
	return false
}

type EmailRuleActionRuleParam struct {
	Action  param.Field[EmailRuleActionRuleAction] `json:"action,required"`
	Message param.Field[string]                    `json:"message"`
}

func (r EmailRuleActionRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmailRuleCondition struct {
	Operator EmailRuleConditionOperator   `json:"operator,required"`
	Selector EmailRuleConditionSelector   `json:"selector,required"`
	Value    EmailRuleConditionValueUnion `json:"value,required"`
	JSON     emailRuleConditionJSON       `json:"-"`
}

// emailRuleConditionJSON contains the JSON metadata for the struct
// [EmailRuleCondition]
type emailRuleConditionJSON struct {
	Operator    apijson.Field
	Selector    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRuleCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRuleConditionJSON) RawJSON() string {
	return r.raw
}

type EmailRuleConditionOperator string

const (
	EmailRuleConditionOperatorInList        EmailRuleConditionOperator = "InList"
	EmailRuleConditionOperatorNotInList     EmailRuleConditionOperator = "NotInList"
	EmailRuleConditionOperatorMatchRegex    EmailRuleConditionOperator = "MatchRegex"
	EmailRuleConditionOperatorNotMatchRegex EmailRuleConditionOperator = "NotMatchRegex"
)

func (r EmailRuleConditionOperator) IsKnown() bool {
	switch r {
	case EmailRuleConditionOperatorInList, EmailRuleConditionOperatorNotInList, EmailRuleConditionOperatorMatchRegex, EmailRuleConditionOperatorNotMatchRegex:
		return true
	}
	return false
}

type EmailRuleConditionSelector string

const (
	EmailRuleConditionSelectorRecipients  EmailRuleConditionSelector = "Recipients"
	EmailRuleConditionSelectorSender      EmailRuleConditionSelector = "Sender"
	EmailRuleConditionSelectorDlpProfiles EmailRuleConditionSelector = "DLPProfiles"
)

func (r EmailRuleConditionSelector) IsKnown() bool {
	switch r {
	case EmailRuleConditionSelectorRecipients, EmailRuleConditionSelectorSender, EmailRuleConditionSelectorDlpProfiles:
		return true
	}
	return false
}

// Union satisfied by [EmailRuleConditionValueArray] or [shared.UnionString].
type EmailRuleConditionValueUnion interface {
	ImplementsEmailRuleConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EmailRuleConditionValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmailRuleConditionValueArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type EmailRuleConditionValueArray []string

func (r EmailRuleConditionValueArray) ImplementsEmailRuleConditionValueUnion() {}

type EmailRuleConditionParam struct {
	Operator param.Field[EmailRuleConditionOperator]        `json:"operator,required"`
	Selector param.Field[EmailRuleConditionSelector]        `json:"selector,required"`
	Value    param.Field[EmailRuleConditionValueUnionParam] `json:"value,required"`
}

func (r EmailRuleConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [EmailRuleConditionValueArrayParam], [shared.UnionString].
type EmailRuleConditionValueUnionParam interface {
	ImplementsEmailRuleConditionValueUnionParam()
}

type EmailRuleConditionValueArrayParam []string

func (r EmailRuleConditionValueArrayParam) ImplementsEmailRuleConditionValueUnionParam() {}

type AccountDlpEmailRuleNewResponse struct {
	Result EmailRule                          `json:"result"`
	JSON   accountDlpEmailRuleNewResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpEmailRuleNewResponseJSON contains the JSON metadata for the struct
// [AccountDlpEmailRuleNewResponse]
type accountDlpEmailRuleNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpEmailRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpEmailRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpEmailRuleGetResponse struct {
	Result EmailRule                          `json:"result"`
	JSON   accountDlpEmailRuleGetResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpEmailRuleGetResponseJSON contains the JSON metadata for the struct
// [AccountDlpEmailRuleGetResponse]
type accountDlpEmailRuleGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpEmailRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpEmailRuleGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpEmailRuleUpdateResponse struct {
	Result EmailRule                             `json:"result"`
	JSON   accountDlpEmailRuleUpdateResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpEmailRuleUpdateResponseJSON contains the JSON metadata for the struct
// [AccountDlpEmailRuleUpdateResponse]
type accountDlpEmailRuleUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpEmailRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpEmailRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpEmailRuleListResponse struct {
	Result []EmailRule                         `json:"result"`
	JSON   accountDlpEmailRuleListResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpEmailRuleListResponseJSON contains the JSON metadata for the struct
// [AccountDlpEmailRuleListResponse]
type accountDlpEmailRuleListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpEmailRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpEmailRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpEmailRuleDeleteResponse struct {
	Result EmailRule                             `json:"result"`
	JSON   accountDlpEmailRuleDeleteResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpEmailRuleDeleteResponseJSON contains the JSON metadata for the struct
// [AccountDlpEmailRuleDeleteResponse]
type accountDlpEmailRuleDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpEmailRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpEmailRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpEmailRuleUpdatePrioritiesResponse struct {
	Result EmailRule                                       `json:"result"`
	JSON   accountDlpEmailRuleUpdatePrioritiesResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpEmailRuleUpdatePrioritiesResponseJSON contains the JSON metadata for
// the struct [AccountDlpEmailRuleUpdatePrioritiesResponse]
type accountDlpEmailRuleUpdatePrioritiesResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpEmailRuleUpdatePrioritiesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpEmailRuleUpdatePrioritiesResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpEmailRuleNewParams struct {
	CreateEmailRule CreateEmailRuleParam `json:"create_email_rule,required"`
}

func (r AccountDlpEmailRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateEmailRule)
}

type AccountDlpEmailRuleUpdateParams struct {
	CreateEmailRule CreateEmailRuleParam `json:"create_email_rule,required"`
}

func (r AccountDlpEmailRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateEmailRule)
}

type AccountDlpEmailRuleUpdatePrioritiesParams struct {
	NewPriorities param.Field[map[string]int64] `json:"new_priorities,required"`
}

func (r AccountDlpEmailRuleUpdatePrioritiesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
