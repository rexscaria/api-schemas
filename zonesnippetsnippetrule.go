// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneSnippetSnippetRuleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSnippetSnippetRuleService] method instead.
type ZoneSnippetSnippetRuleService struct {
	Options []option.RequestOption
}

// NewZoneSnippetSnippetRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSnippetSnippetRuleService(opts ...option.RequestOption) (r *ZoneSnippetSnippetRuleService) {
	r = &ZoneSnippetSnippetRuleService{}
	r.Options = opts
	return
}

// Updates all snippet rules belonging to the zone.
func (r *ZoneSnippetSnippetRuleService) Update(ctx context.Context, zoneID string, body ZoneSnippetSnippetRuleUpdateParams, opts ...option.RequestOption) (res *ZoneSnippetSnippetRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches all snippet rules belonging to the zone.
func (r *ZoneSnippetSnippetRuleService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSnippetSnippetRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes all snippet rules belonging to the zone.
func (r *ZoneSnippetSnippetRuleService) DeleteAll(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSnippetSnippetRuleDeleteAllResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// A response object.
type ZoneSnippetSnippetRuleUpdateResponse struct {
	// A list of error messages.
	Errors []ZoneSnippetSnippetRuleUpdateResponseError `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneSnippetSnippetRuleUpdateResponseMessage `json:"messages,required"`
	// A result.
	Result []ZoneSnippetSnippetRuleUpdateResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneSnippetSnippetRuleUpdateResponseSuccess `json:"success,required"`
	JSON    zoneSnippetSnippetRuleUpdateResponseJSON    `json:"-"`
}

// zoneSnippetSnippetRuleUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSnippetSnippetRuleUpdateResponse]
type zoneSnippetSnippetRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetSnippetRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetSnippetRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneSnippetSnippetRuleUpdateResponseError struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                                         `json:"code"`
	JSON zoneSnippetSnippetRuleUpdateResponseErrorJSON `json:"-"`
}

// zoneSnippetSnippetRuleUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSnippetSnippetRuleUpdateResponseError]
type zoneSnippetSnippetRuleUpdateResponseErrorJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetSnippetRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetSnippetRuleUpdateResponseErrorJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneSnippetSnippetRuleUpdateResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                                           `json:"code"`
	JSON zoneSnippetSnippetRuleUpdateResponseMessageJSON `json:"-"`
}

// zoneSnippetSnippetRuleUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSnippetSnippetRuleUpdateResponseMessage]
type zoneSnippetSnippetRuleUpdateResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetSnippetRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetSnippetRuleUpdateResponseMessageJSON) RawJSON() string {
	return r.raw
}

// A snippet rule.
type ZoneSnippetSnippetRuleUpdateResponseResult struct {
	// The unique ID of the rule.
	ID string `json:"id,required"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression,required"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The identifying name of the snippet.
	SnippetName string `json:"snippet_name,required"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool                                           `json:"enabled"`
	JSON    zoneSnippetSnippetRuleUpdateResponseResultJSON `json:"-"`
}

// zoneSnippetSnippetRuleUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSnippetSnippetRuleUpdateResponseResult]
type zoneSnippetSnippetRuleUpdateResponseResultJSON struct {
	ID          apijson.Field
	Expression  apijson.Field
	LastUpdated apijson.Field
	SnippetName apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetSnippetRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetSnippetRuleUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneSnippetSnippetRuleUpdateResponseSuccess bool

const (
	ZoneSnippetSnippetRuleUpdateResponseSuccessTrue ZoneSnippetSnippetRuleUpdateResponseSuccess = true
)

func (r ZoneSnippetSnippetRuleUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSnippetSnippetRuleUpdateResponseSuccessTrue:
		return true
	}
	return false
}

// A response object.
type ZoneSnippetSnippetRuleListResponse struct {
	// A list of error messages.
	Errors []ZoneSnippetSnippetRuleListResponseError `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneSnippetSnippetRuleListResponseMessage `json:"messages,required"`
	// A result.
	Result []ZoneSnippetSnippetRuleListResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneSnippetSnippetRuleListResponseSuccess `json:"success,required"`
	JSON    zoneSnippetSnippetRuleListResponseJSON    `json:"-"`
}

// zoneSnippetSnippetRuleListResponseJSON contains the JSON metadata for the struct
// [ZoneSnippetSnippetRuleListResponse]
type zoneSnippetSnippetRuleListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetSnippetRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetSnippetRuleListResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneSnippetSnippetRuleListResponseError struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                                       `json:"code"`
	JSON zoneSnippetSnippetRuleListResponseErrorJSON `json:"-"`
}

// zoneSnippetSnippetRuleListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSnippetSnippetRuleListResponseError]
type zoneSnippetSnippetRuleListResponseErrorJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetSnippetRuleListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetSnippetRuleListResponseErrorJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneSnippetSnippetRuleListResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                                         `json:"code"`
	JSON zoneSnippetSnippetRuleListResponseMessageJSON `json:"-"`
}

// zoneSnippetSnippetRuleListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSnippetSnippetRuleListResponseMessage]
type zoneSnippetSnippetRuleListResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetSnippetRuleListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetSnippetRuleListResponseMessageJSON) RawJSON() string {
	return r.raw
}

// A snippet rule.
type ZoneSnippetSnippetRuleListResponseResult struct {
	// The unique ID of the rule.
	ID string `json:"id,required"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression,required"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The identifying name of the snippet.
	SnippetName string `json:"snippet_name,required"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool                                         `json:"enabled"`
	JSON    zoneSnippetSnippetRuleListResponseResultJSON `json:"-"`
}

// zoneSnippetSnippetRuleListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSnippetSnippetRuleListResponseResult]
type zoneSnippetSnippetRuleListResponseResultJSON struct {
	ID          apijson.Field
	Expression  apijson.Field
	LastUpdated apijson.Field
	SnippetName apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetSnippetRuleListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetSnippetRuleListResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneSnippetSnippetRuleListResponseSuccess bool

const (
	ZoneSnippetSnippetRuleListResponseSuccessTrue ZoneSnippetSnippetRuleListResponseSuccess = true
)

func (r ZoneSnippetSnippetRuleListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSnippetSnippetRuleListResponseSuccessTrue:
		return true
	}
	return false
}

// A response object.
type ZoneSnippetSnippetRuleDeleteAllResponse struct {
	// A list of error messages.
	Errors []ZoneSnippetSnippetRuleDeleteAllResponseError `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneSnippetSnippetRuleDeleteAllResponseMessage `json:"messages,required"`
	// A result.
	Result []ZoneSnippetSnippetRuleDeleteAllResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneSnippetSnippetRuleDeleteAllResponseSuccess `json:"success,required"`
	JSON    zoneSnippetSnippetRuleDeleteAllResponseJSON    `json:"-"`
}

// zoneSnippetSnippetRuleDeleteAllResponseJSON contains the JSON metadata for the
// struct [ZoneSnippetSnippetRuleDeleteAllResponse]
type zoneSnippetSnippetRuleDeleteAllResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetSnippetRuleDeleteAllResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetSnippetRuleDeleteAllResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneSnippetSnippetRuleDeleteAllResponseError struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                                            `json:"code"`
	JSON zoneSnippetSnippetRuleDeleteAllResponseErrorJSON `json:"-"`
}

// zoneSnippetSnippetRuleDeleteAllResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSnippetSnippetRuleDeleteAllResponseError]
type zoneSnippetSnippetRuleDeleteAllResponseErrorJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetSnippetRuleDeleteAllResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetSnippetRuleDeleteAllResponseErrorJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneSnippetSnippetRuleDeleteAllResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                                              `json:"code"`
	JSON zoneSnippetSnippetRuleDeleteAllResponseMessageJSON `json:"-"`
}

// zoneSnippetSnippetRuleDeleteAllResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSnippetSnippetRuleDeleteAllResponseMessage]
type zoneSnippetSnippetRuleDeleteAllResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetSnippetRuleDeleteAllResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetSnippetRuleDeleteAllResponseMessageJSON) RawJSON() string {
	return r.raw
}

// A snippet rule.
type ZoneSnippetSnippetRuleDeleteAllResponseResult struct {
	// The unique ID of the rule.
	ID string `json:"id,required"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression,required"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The identifying name of the snippet.
	SnippetName string `json:"snippet_name,required"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool                                              `json:"enabled"`
	JSON    zoneSnippetSnippetRuleDeleteAllResponseResultJSON `json:"-"`
}

// zoneSnippetSnippetRuleDeleteAllResponseResultJSON contains the JSON metadata for
// the struct [ZoneSnippetSnippetRuleDeleteAllResponseResult]
type zoneSnippetSnippetRuleDeleteAllResponseResultJSON struct {
	ID          apijson.Field
	Expression  apijson.Field
	LastUpdated apijson.Field
	SnippetName apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetSnippetRuleDeleteAllResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetSnippetRuleDeleteAllResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneSnippetSnippetRuleDeleteAllResponseSuccess bool

const (
	ZoneSnippetSnippetRuleDeleteAllResponseSuccessTrue ZoneSnippetSnippetRuleDeleteAllResponseSuccess = true
)

func (r ZoneSnippetSnippetRuleDeleteAllResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSnippetSnippetRuleDeleteAllResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneSnippetSnippetRuleUpdateParams struct {
	// A list of snippet rules.
	Body []ZoneSnippetSnippetRuleUpdateParamsBody `json:"body,required"`
}

func (r ZoneSnippetSnippetRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// A snippet rule.
type ZoneSnippetSnippetRuleUpdateParamsBody struct {
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression,required"`
	// The identifying name of the snippet.
	SnippetName param.Field[string] `json:"snippet_name,required"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSnippetSnippetRuleUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
