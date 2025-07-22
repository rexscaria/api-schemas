// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountRulesetPhaseEntrypointVersionService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRulesetPhaseEntrypointVersionService] method instead.
type AccountRulesetPhaseEntrypointVersionService struct {
	Options []option.RequestOption
}

// NewAccountRulesetPhaseEntrypointVersionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountRulesetPhaseEntrypointVersionService(opts ...option.RequestOption) (r *AccountRulesetPhaseEntrypointVersionService) {
	r = &AccountRulesetPhaseEntrypointVersionService{}
	r.Options = opts
	return
}

// Fetches a specific version of an account entry point ruleset.
func (r *AccountRulesetPhaseEntrypointVersionService) Get(ctx context.Context, accountID string, rulesetPhase RulesetPhase, rulesetVersion string, opts ...option.RequestOption) (res *AccountRulesetPhaseEntrypointVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetVersion == "" {
		err = errors.New("missing required ruleset_version parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets/phases/%v/entrypoint/versions/%s", accountID, rulesetPhase, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches the versions of an account entry point ruleset.
func (r *AccountRulesetPhaseEntrypointVersionService) List(ctx context.Context, accountID string, rulesetPhase RulesetPhase, opts ...option.RequestOption) (res *AccountRulesetPhaseEntrypointVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets/phases/%v/entrypoint/versions", accountID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountRulesetPhaseEntrypointVersionGetResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []AccountRulesetPhaseEntrypointVersionGetResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success AccountRulesetPhaseEntrypointVersionGetResponseSuccess `json:"success,required"`
	JSON    accountRulesetPhaseEntrypointVersionGetResponseJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointVersionGetResponseJSON contains the JSON metadata
// for the struct [AccountRulesetPhaseEntrypointVersionGetResponse]
type accountRulesetPhaseEntrypointVersionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetPhaseEntrypointVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type AccountRulesetPhaseEntrypointVersionGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetPhaseEntrypointVersionGetResponseMessagesSource `json:"source"`
	JSON   accountRulesetPhaseEntrypointVersionGetResponseMessageJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointVersionGetResponseMessageJSON contains the JSON
// metadata for the struct [AccountRulesetPhaseEntrypointVersionGetResponseMessage]
type accountRulesetPhaseEntrypointVersionGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointVersionGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetPhaseEntrypointVersionGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type AccountRulesetPhaseEntrypointVersionGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                            `json:"pointer,required"`
	JSON    accountRulesetPhaseEntrypointVersionGetResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetPhaseEntrypointVersionGetResponseMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccountRulesetPhaseEntrypointVersionGetResponseMessagesSource]
type accountRulesetPhaseEntrypointVersionGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointVersionGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetPhaseEntrypointVersionGetResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRulesetPhaseEntrypointVersionGetResponseSuccess bool

const (
	AccountRulesetPhaseEntrypointVersionGetResponseSuccessTrue AccountRulesetPhaseEntrypointVersionGetResponseSuccess = true
)

func (r AccountRulesetPhaseEntrypointVersionGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRulesetPhaseEntrypointVersionGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountRulesetPhaseEntrypointVersionListResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []AccountRulesetPhaseEntrypointVersionListResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success AccountRulesetPhaseEntrypointVersionListResponseSuccess `json:"success,required"`
	// Cursors information to navigate the results.
	ResultInfo AccountRulesetPhaseEntrypointVersionListResponseResultInfo `json:"result_info"`
	JSON       accountRulesetPhaseEntrypointVersionListResponseJSON       `json:"-"`
}

// accountRulesetPhaseEntrypointVersionListResponseJSON contains the JSON metadata
// for the struct [AccountRulesetPhaseEntrypointVersionListResponse]
type accountRulesetPhaseEntrypointVersionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetPhaseEntrypointVersionListResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type AccountRulesetPhaseEntrypointVersionListResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetPhaseEntrypointVersionListResponseMessagesSource `json:"source"`
	JSON   accountRulesetPhaseEntrypointVersionListResponseMessageJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointVersionListResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountRulesetPhaseEntrypointVersionListResponseMessage]
type accountRulesetPhaseEntrypointVersionListResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointVersionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetPhaseEntrypointVersionListResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type AccountRulesetPhaseEntrypointVersionListResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                             `json:"pointer,required"`
	JSON    accountRulesetPhaseEntrypointVersionListResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetPhaseEntrypointVersionListResponseMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccountRulesetPhaseEntrypointVersionListResponseMessagesSource]
type accountRulesetPhaseEntrypointVersionListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointVersionListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetPhaseEntrypointVersionListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRulesetPhaseEntrypointVersionListResponseSuccess bool

const (
	AccountRulesetPhaseEntrypointVersionListResponseSuccessTrue AccountRulesetPhaseEntrypointVersionListResponseSuccess = true
)

func (r AccountRulesetPhaseEntrypointVersionListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRulesetPhaseEntrypointVersionListResponseSuccessTrue:
		return true
	}
	return false
}

// Cursors information to navigate the results.
type AccountRulesetPhaseEntrypointVersionListResponseResultInfo struct {
	// Set of cursors.
	Cursors AccountRulesetPhaseEntrypointVersionListResponseResultInfoCursors `json:"cursors"`
	JSON    accountRulesetPhaseEntrypointVersionListResponseResultInfoJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointVersionListResponseResultInfoJSON contains the JSON
// metadata for the struct
// [AccountRulesetPhaseEntrypointVersionListResponseResultInfo]
type accountRulesetPhaseEntrypointVersionListResponseResultInfoJSON struct {
	Cursors     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointVersionListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetPhaseEntrypointVersionListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

// Set of cursors.
type AccountRulesetPhaseEntrypointVersionListResponseResultInfoCursors struct {
	// Cursor to use for the next page.
	After string                                                                `json:"after"`
	JSON  accountRulesetPhaseEntrypointVersionListResponseResultInfoCursorsJSON `json:"-"`
}

// accountRulesetPhaseEntrypointVersionListResponseResultInfoCursorsJSON contains
// the JSON metadata for the struct
// [AccountRulesetPhaseEntrypointVersionListResponseResultInfoCursors]
type accountRulesetPhaseEntrypointVersionListResponseResultInfoCursorsJSON struct {
	After       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointVersionListResponseResultInfoCursors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetPhaseEntrypointVersionListResponseResultInfoCursorsJSON) RawJSON() string {
	return r.raw
}
