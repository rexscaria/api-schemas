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

// AccountRulesetVersionService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRulesetVersionService] method instead.
type AccountRulesetVersionService struct {
	Options []option.RequestOption
}

// NewAccountRulesetVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRulesetVersionService(opts ...option.RequestOption) (r *AccountRulesetVersionService) {
	r = &AccountRulesetVersionService{}
	r.Options = opts
	return
}

// Fetches a specific version of an account ruleset.
func (r *AccountRulesetVersionService) Get(ctx context.Context, accountID string, rulesetID string, rulesetVersion string, opts ...option.RequestOption) (res *AccountRulesetVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if rulesetVersion == "" {
		err = errors.New("missing required ruleset_version parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s", accountID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches the versions of an account ruleset.
func (r *AccountRulesetVersionService) List(ctx context.Context, accountID string, rulesetID string, opts ...option.RequestOption) (res *AccountRulesetVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an existing version of an account ruleset.
func (r *AccountRulesetVersionService) Delete(ctx context.Context, accountID string, rulesetID string, rulesetVersion string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if rulesetVersion == "" {
		err = errors.New("missing required ruleset_version parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s", accountID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches the rules of a managed account ruleset version for a given tag.
func (r *AccountRulesetVersionService) ListByTag(ctx context.Context, accountID string, rulesetID string, rulesetVersion string, ruleTag string, opts ...option.RequestOption) (res *AccountRulesetVersionListByTagResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if rulesetVersion == "" {
		err = errors.New("missing required ruleset_version parameter")
		return
	}
	if ruleTag == "" {
		err = errors.New("missing required rule_tag parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s/by_tag/%s", accountID, rulesetID, rulesetVersion, ruleTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountRulesetVersionGetResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []AccountRulesetVersionGetResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success AccountRulesetVersionGetResponseSuccess `json:"success,required"`
	JSON    accountRulesetVersionGetResponseJSON    `json:"-"`
}

// accountRulesetVersionGetResponseJSON contains the JSON metadata for the struct
// [AccountRulesetVersionGetResponse]
type accountRulesetVersionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type AccountRulesetVersionGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetVersionGetResponseMessagesSource `json:"source"`
	JSON   accountRulesetVersionGetResponseMessageJSON    `json:"-"`
}

// accountRulesetVersionGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountRulesetVersionGetResponseMessage]
type accountRulesetVersionGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetVersionGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type AccountRulesetVersionGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                             `json:"pointer,required"`
	JSON    accountRulesetVersionGetResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetVersionGetResponseMessagesSourceJSON contains the JSON metadata
// for the struct [AccountRulesetVersionGetResponseMessagesSource]
type accountRulesetVersionGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetVersionGetResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRulesetVersionGetResponseSuccess bool

const (
	AccountRulesetVersionGetResponseSuccessTrue AccountRulesetVersionGetResponseSuccess = true
)

func (r AccountRulesetVersionGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRulesetVersionGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountRulesetVersionListResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []AccountRulesetVersionListResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success AccountRulesetVersionListResponseSuccess `json:"success,required"`
	// Cursors information to navigate the results.
	ResultInfo AccountRulesetVersionListResponseResultInfo `json:"result_info"`
	JSON       accountRulesetVersionListResponseJSON       `json:"-"`
}

// accountRulesetVersionListResponseJSON contains the JSON metadata for the struct
// [AccountRulesetVersionListResponse]
type accountRulesetVersionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetVersionListResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type AccountRulesetVersionListResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetVersionListResponseMessagesSource `json:"source"`
	JSON   accountRulesetVersionListResponseMessageJSON    `json:"-"`
}

// accountRulesetVersionListResponseMessageJSON contains the JSON metadata for the
// struct [AccountRulesetVersionListResponseMessage]
type accountRulesetVersionListResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetVersionListResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type AccountRulesetVersionListResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                              `json:"pointer,required"`
	JSON    accountRulesetVersionListResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetVersionListResponseMessagesSourceJSON contains the JSON metadata
// for the struct [AccountRulesetVersionListResponseMessagesSource]
type accountRulesetVersionListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetVersionListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRulesetVersionListResponseSuccess bool

const (
	AccountRulesetVersionListResponseSuccessTrue AccountRulesetVersionListResponseSuccess = true
)

func (r AccountRulesetVersionListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRulesetVersionListResponseSuccessTrue:
		return true
	}
	return false
}

// Cursors information to navigate the results.
type AccountRulesetVersionListResponseResultInfo struct {
	// Set of cursors.
	Cursors AccountRulesetVersionListResponseResultInfoCursors `json:"cursors"`
	JSON    accountRulesetVersionListResponseResultInfoJSON    `json:"-"`
}

// accountRulesetVersionListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountRulesetVersionListResponseResultInfo]
type accountRulesetVersionListResponseResultInfoJSON struct {
	Cursors     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetVersionListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

// Set of cursors.
type AccountRulesetVersionListResponseResultInfoCursors struct {
	// Cursor to use for the next page.
	After string                                                 `json:"after"`
	JSON  accountRulesetVersionListResponseResultInfoCursorsJSON `json:"-"`
}

// accountRulesetVersionListResponseResultInfoCursorsJSON contains the JSON
// metadata for the struct [AccountRulesetVersionListResponseResultInfoCursors]
type accountRulesetVersionListResponseResultInfoCursorsJSON struct {
	After       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionListResponseResultInfoCursors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetVersionListResponseResultInfoCursorsJSON) RawJSON() string {
	return r.raw
}

type AccountRulesetVersionListByTagResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []AccountRulesetVersionListByTagResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success AccountRulesetVersionListByTagResponseSuccess `json:"success,required"`
	JSON    accountRulesetVersionListByTagResponseJSON    `json:"-"`
}

// accountRulesetVersionListByTagResponseJSON contains the JSON metadata for the
// struct [AccountRulesetVersionListByTagResponse]
type accountRulesetVersionListByTagResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionListByTagResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetVersionListByTagResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type AccountRulesetVersionListByTagResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetVersionListByTagResponseMessagesSource `json:"source"`
	JSON   accountRulesetVersionListByTagResponseMessageJSON    `json:"-"`
}

// accountRulesetVersionListByTagResponseMessageJSON contains the JSON metadata for
// the struct [AccountRulesetVersionListByTagResponseMessage]
type accountRulesetVersionListByTagResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionListByTagResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetVersionListByTagResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type AccountRulesetVersionListByTagResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                   `json:"pointer,required"`
	JSON    accountRulesetVersionListByTagResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetVersionListByTagResponseMessagesSourceJSON contains the JSON
// metadata for the struct [AccountRulesetVersionListByTagResponseMessagesSource]
type accountRulesetVersionListByTagResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionListByTagResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetVersionListByTagResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRulesetVersionListByTagResponseSuccess bool

const (
	AccountRulesetVersionListByTagResponseSuccessTrue AccountRulesetVersionListByTagResponseSuccess = true
)

func (r AccountRulesetVersionListByTagResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRulesetVersionListByTagResponseSuccessTrue:
		return true
	}
	return false
}
