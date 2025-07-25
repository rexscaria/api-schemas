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

// AccountChallengeWidgetService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountChallengeWidgetService] method instead.
type AccountChallengeWidgetService struct {
	Options []option.RequestOption
}

// NewAccountChallengeWidgetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountChallengeWidgetService(opts ...option.RequestOption) (r *AccountChallengeWidgetService) {
	r = &AccountChallengeWidgetService{}
	r.Options = opts
	return
}

// Lists challenge widgets.
func (r *AccountChallengeWidgetService) New(ctx context.Context, accountID string, params AccountChallengeWidgetNewParams, opts ...option.RequestOption) (res *AccountChallengeWidgetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/challenges/widgets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Show a single challenge widget configuration.
func (r *AccountChallengeWidgetService) Get(ctx context.Context, accountID string, sitekey string, opts ...option.RequestOption) (res *AccountChallengeWidgetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if sitekey == "" {
		err = errors.New("missing required sitekey parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", accountID, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the configuration of a widget.
func (r *AccountChallengeWidgetService) Update(ctx context.Context, accountID string, sitekey string, body AccountChallengeWidgetUpdateParams, opts ...option.RequestOption) (res *AccountChallengeWidgetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if sitekey == "" {
		err = errors.New("missing required sitekey parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", accountID, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all turnstile widgets of an account.
func (r *AccountChallengeWidgetService) List(ctx context.Context, accountID string, query AccountChallengeWidgetListParams, opts ...option.RequestOption) (res *AccountChallengeWidgetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/challenges/widgets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Destroy a Turnstile Widget.
func (r *AccountChallengeWidgetService) Delete(ctx context.Context, accountID string, sitekey string, opts ...option.RequestOption) (res *AccountChallengeWidgetDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if sitekey == "" {
		err = errors.New("missing required sitekey parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", accountID, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Generate a new secret key for this widget. If `invalidate_immediately` is set to
// `false`, the previous secret remains valid for 2 hours.
//
// Note that secrets cannot be rotated again during the grace period.
func (r *AccountChallengeWidgetService) RotateSecret(ctx context.Context, accountID string, sitekey string, body AccountChallengeWidgetRotateSecretParams, opts ...option.RequestOption) (res *AccountChallengeWidgetRotateSecretResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if sitekey == "" {
		err = errors.New("missing required sitekey parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s/rotate_secret", accountID, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type TurnstileClearanceLevel string

const (
	TurnstileClearanceLevelNoClearance TurnstileClearanceLevel = "no_clearance"
	TurnstileClearanceLevelJschallenge TurnstileClearanceLevel = "jschallenge"
	TurnstileClearanceLevelManaged     TurnstileClearanceLevel = "managed"
	TurnstileClearanceLevelInteractive TurnstileClearanceLevel = "interactive"
)

func (r TurnstileClearanceLevel) IsKnown() bool {
	switch r {
	case TurnstileClearanceLevelNoClearance, TurnstileClearanceLevelJschallenge, TurnstileClearanceLevelManaged, TurnstileClearanceLevelInteractive:
		return true
	}
	return false
}

type TurnstileMessages struct {
	Code             int64                   `json:"code,required"`
	Message          string                  `json:"message,required"`
	DocumentationURL string                  `json:"documentation_url"`
	Source           TurnstileMessagesSource `json:"source"`
	JSON             turnstileMessagesJSON   `json:"-"`
}

// turnstileMessagesJSON contains the JSON metadata for the struct
// [TurnstileMessages]
type turnstileMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TurnstileMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnstileMessagesJSON) RawJSON() string {
	return r.raw
}

type TurnstileMessagesSource struct {
	Pointer string                      `json:"pointer"`
	JSON    turnstileMessagesSourceJSON `json:"-"`
}

// turnstileMessagesSourceJSON contains the JSON metadata for the struct
// [TurnstileMessagesSource]
type turnstileMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TurnstileMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnstileMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Region where this widget can be used. This cannot be changed after creation.
type TurnstileRegion string

const (
	TurnstileRegionWorld TurnstileRegion = "world"
	TurnstileRegionChina TurnstileRegion = "china"
)

func (r TurnstileRegion) IsKnown() bool {
	switch r {
	case TurnstileRegionWorld, TurnstileRegionChina:
		return true
	}
	return false
}

type TurnstileResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count,required"`
	// Current page within paginated list of results
	Page float64 `json:"page,required"`
	// Number of results per page of results
	PerPage float64 `json:"per_page,required"`
	// Total results available without any search parameters
	TotalCount float64                 `json:"total_count,required"`
	JSON       turnstileResultInfoJSON `json:"-"`
}

// turnstileResultInfoJSON contains the JSON metadata for the struct
// [TurnstileResultInfo]
type turnstileResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TurnstileResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnstileResultInfoJSON) RawJSON() string {
	return r.raw
}

// A Turnstile widget's detailed configuration
type TurnstileWidgetDetail struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel TurnstileClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Return the Ephemeral ID in /siteverify (ENT only).
	EphemeralID bool `json:"ephemeral_id,required"`
	// Widget Mode
	Mode TurnstileWidgetMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used. This cannot be changed after creation.
	Region TurnstileRegion `json:"region,required"`
	// Secret key for this widget.
	Secret string `json:"secret,required"`
	// Widget item identifier tag.
	Sitekey string                    `json:"sitekey,required"`
	JSON    turnstileWidgetDetailJSON `json:"-"`
}

// turnstileWidgetDetailJSON contains the JSON metadata for the struct
// [TurnstileWidgetDetail]
type turnstileWidgetDetailJSON struct {
	BotFightMode   apijson.Field
	ClearanceLevel apijson.Field
	CreatedOn      apijson.Field
	Domains        apijson.Field
	EphemeralID    apijson.Field
	Mode           apijson.Field
	ModifiedOn     apijson.Field
	Name           apijson.Field
	Offlabel       apijson.Field
	Region         apijson.Field
	Secret         apijson.Field
	Sitekey        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TurnstileWidgetDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnstileWidgetDetailJSON) RawJSON() string {
	return r.raw
}

// Widget Mode
type TurnstileWidgetMode string

const (
	TurnstileWidgetModeNonInteractive TurnstileWidgetMode = "non-interactive"
	TurnstileWidgetModeInvisible      TurnstileWidgetMode = "invisible"
	TurnstileWidgetModeManaged        TurnstileWidgetMode = "managed"
)

func (r TurnstileWidgetMode) IsKnown() bool {
	switch r {
	case TurnstileWidgetModeNonInteractive, TurnstileWidgetModeInvisible, TurnstileWidgetModeManaged:
		return true
	}
	return false
}

type AccountChallengeWidgetNewResponse struct {
	Errors   []TurnstileMessages `json:"errors,required"`
	Messages []TurnstileMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result     TurnstileWidgetDetail                 `json:"result"`
	ResultInfo TurnstileResultInfo                   `json:"result_info"`
	JSON       accountChallengeWidgetNewResponseJSON `json:"-"`
}

// accountChallengeWidgetNewResponseJSON contains the JSON metadata for the struct
// [AccountChallengeWidgetNewResponse]
type accountChallengeWidgetNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountChallengeWidgetNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountChallengeWidgetGetResponse struct {
	Errors   []TurnstileMessages `json:"errors,required"`
	Messages []TurnstileMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result TurnstileWidgetDetail                 `json:"result"`
	JSON   accountChallengeWidgetGetResponseJSON `json:"-"`
}

// accountChallengeWidgetGetResponseJSON contains the JSON metadata for the struct
// [AccountChallengeWidgetGetResponse]
type accountChallengeWidgetGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountChallengeWidgetGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountChallengeWidgetUpdateResponse struct {
	Errors   []TurnstileMessages `json:"errors,required"`
	Messages []TurnstileMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result TurnstileWidgetDetail                    `json:"result"`
	JSON   accountChallengeWidgetUpdateResponseJSON `json:"-"`
}

// accountChallengeWidgetUpdateResponseJSON contains the JSON metadata for the
// struct [AccountChallengeWidgetUpdateResponse]
type accountChallengeWidgetUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountChallengeWidgetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountChallengeWidgetListResponse struct {
	Errors   []TurnstileMessages `json:"errors,required"`
	Messages []TurnstileMessages `json:"messages,required"`
	// Whether the API call was successful
	Success    bool                                       `json:"success,required"`
	Result     []AccountChallengeWidgetListResponseResult `json:"result"`
	ResultInfo TurnstileResultInfo                        `json:"result_info"`
	JSON       accountChallengeWidgetListResponseJSON     `json:"-"`
}

// accountChallengeWidgetListResponseJSON contains the JSON metadata for the struct
// [AccountChallengeWidgetListResponse]
type accountChallengeWidgetListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountChallengeWidgetListResponseJSON) RawJSON() string {
	return r.raw
}

// A Turnstile Widgets configuration as it appears in listings
type AccountChallengeWidgetListResponseResult struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel TurnstileClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Return the Ephemeral ID in /siteverify (ENT only).
	EphemeralID bool `json:"ephemeral_id,required"`
	// Widget Mode
	Mode TurnstileWidgetMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used. This cannot be changed after creation.
	Region TurnstileRegion `json:"region,required"`
	// Widget item identifier tag.
	Sitekey string                                       `json:"sitekey,required"`
	JSON    accountChallengeWidgetListResponseResultJSON `json:"-"`
}

// accountChallengeWidgetListResponseResultJSON contains the JSON metadata for the
// struct [AccountChallengeWidgetListResponseResult]
type accountChallengeWidgetListResponseResultJSON struct {
	BotFightMode   apijson.Field
	ClearanceLevel apijson.Field
	CreatedOn      apijson.Field
	Domains        apijson.Field
	EphemeralID    apijson.Field
	Mode           apijson.Field
	ModifiedOn     apijson.Field
	Name           apijson.Field
	Offlabel       apijson.Field
	Region         apijson.Field
	Sitekey        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountChallengeWidgetListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountChallengeWidgetListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountChallengeWidgetDeleteResponse struct {
	Errors   []TurnstileMessages `json:"errors,required"`
	Messages []TurnstileMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result TurnstileWidgetDetail                    `json:"result"`
	JSON   accountChallengeWidgetDeleteResponseJSON `json:"-"`
}

// accountChallengeWidgetDeleteResponseJSON contains the JSON metadata for the
// struct [AccountChallengeWidgetDeleteResponse]
type accountChallengeWidgetDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountChallengeWidgetDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountChallengeWidgetRotateSecretResponse struct {
	Errors   []TurnstileMessages `json:"errors,required"`
	Messages []TurnstileMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result TurnstileWidgetDetail                          `json:"result"`
	JSON   accountChallengeWidgetRotateSecretResponseJSON `json:"-"`
}

// accountChallengeWidgetRotateSecretResponseJSON contains the JSON metadata for
// the struct [AccountChallengeWidgetRotateSecretResponse]
type accountChallengeWidgetRotateSecretResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetRotateSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountChallengeWidgetRotateSecretResponseJSON) RawJSON() string {
	return r.raw
}

type AccountChallengeWidgetNewParams struct {
	Domains param.Field[[]string] `json:"domains,required"`
	// Widget Mode
	Mode param.Field[TurnstileWidgetMode] `json:"mode,required"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name param.Field[string] `json:"name,required"`
	// Direction to order widgets.
	Direction param.Field[AccountChallengeWidgetNewParamsDirection] `query:"direction"`
	// Field to order widgets by.
	Order param.Field[AccountChallengeWidgetNewParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode param.Field[bool] `json:"bot_fight_mode"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel param.Field[TurnstileClearanceLevel] `json:"clearance_level"`
	// Return the Ephemeral ID in /siteverify (ENT only).
	EphemeralID param.Field[bool] `json:"ephemeral_id"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel param.Field[bool] `json:"offlabel"`
	// Region where this widget can be used. This cannot be changed after creation.
	Region param.Field[TurnstileRegion] `json:"region"`
}

func (r AccountChallengeWidgetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountChallengeWidgetNewParams]'s query parameters as
// `url.Values`.
func (r AccountChallengeWidgetNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order widgets.
type AccountChallengeWidgetNewParamsDirection string

const (
	AccountChallengeWidgetNewParamsDirectionAsc  AccountChallengeWidgetNewParamsDirection = "asc"
	AccountChallengeWidgetNewParamsDirectionDesc AccountChallengeWidgetNewParamsDirection = "desc"
)

func (r AccountChallengeWidgetNewParamsDirection) IsKnown() bool {
	switch r {
	case AccountChallengeWidgetNewParamsDirectionAsc, AccountChallengeWidgetNewParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order widgets by.
type AccountChallengeWidgetNewParamsOrder string

const (
	AccountChallengeWidgetNewParamsOrderID         AccountChallengeWidgetNewParamsOrder = "id"
	AccountChallengeWidgetNewParamsOrderSitekey    AccountChallengeWidgetNewParamsOrder = "sitekey"
	AccountChallengeWidgetNewParamsOrderName       AccountChallengeWidgetNewParamsOrder = "name"
	AccountChallengeWidgetNewParamsOrderCreatedOn  AccountChallengeWidgetNewParamsOrder = "created_on"
	AccountChallengeWidgetNewParamsOrderModifiedOn AccountChallengeWidgetNewParamsOrder = "modified_on"
)

func (r AccountChallengeWidgetNewParamsOrder) IsKnown() bool {
	switch r {
	case AccountChallengeWidgetNewParamsOrderID, AccountChallengeWidgetNewParamsOrderSitekey, AccountChallengeWidgetNewParamsOrderName, AccountChallengeWidgetNewParamsOrderCreatedOn, AccountChallengeWidgetNewParamsOrderModifiedOn:
		return true
	}
	return false
}

type AccountChallengeWidgetUpdateParams struct {
	Domains param.Field[[]string] `json:"domains,required"`
	// Widget Mode
	Mode param.Field[TurnstileWidgetMode] `json:"mode,required"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name param.Field[string] `json:"name,required"`
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode param.Field[bool] `json:"bot_fight_mode"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel param.Field[TurnstileClearanceLevel] `json:"clearance_level"`
	// Return the Ephemeral ID in /siteverify (ENT only).
	EphemeralID param.Field[bool] `json:"ephemeral_id"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel param.Field[bool] `json:"offlabel"`
	// Region where this widget can be used. This cannot be changed after creation.
	Region param.Field[TurnstileRegion] `json:"region"`
}

func (r AccountChallengeWidgetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountChallengeWidgetListParams struct {
	// Direction to order widgets.
	Direction param.Field[AccountChallengeWidgetListParamsDirection] `query:"direction"`
	// Field to order widgets by.
	Order param.Field[AccountChallengeWidgetListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountChallengeWidgetListParams]'s query parameters as
// `url.Values`.
func (r AccountChallengeWidgetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order widgets.
type AccountChallengeWidgetListParamsDirection string

const (
	AccountChallengeWidgetListParamsDirectionAsc  AccountChallengeWidgetListParamsDirection = "asc"
	AccountChallengeWidgetListParamsDirectionDesc AccountChallengeWidgetListParamsDirection = "desc"
)

func (r AccountChallengeWidgetListParamsDirection) IsKnown() bool {
	switch r {
	case AccountChallengeWidgetListParamsDirectionAsc, AccountChallengeWidgetListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order widgets by.
type AccountChallengeWidgetListParamsOrder string

const (
	AccountChallengeWidgetListParamsOrderID         AccountChallengeWidgetListParamsOrder = "id"
	AccountChallengeWidgetListParamsOrderSitekey    AccountChallengeWidgetListParamsOrder = "sitekey"
	AccountChallengeWidgetListParamsOrderName       AccountChallengeWidgetListParamsOrder = "name"
	AccountChallengeWidgetListParamsOrderCreatedOn  AccountChallengeWidgetListParamsOrder = "created_on"
	AccountChallengeWidgetListParamsOrderModifiedOn AccountChallengeWidgetListParamsOrder = "modified_on"
)

func (r AccountChallengeWidgetListParamsOrder) IsKnown() bool {
	switch r {
	case AccountChallengeWidgetListParamsOrderID, AccountChallengeWidgetListParamsOrderSitekey, AccountChallengeWidgetListParamsOrderName, AccountChallengeWidgetListParamsOrderCreatedOn, AccountChallengeWidgetListParamsOrderModifiedOn:
		return true
	}
	return false
}

type AccountChallengeWidgetRotateSecretParams struct {
	// If `invalidate_immediately` is set to `false`, the previous secret will remain
	// valid for two hours. Otherwise, the secret is immediately invalidated, and
	// requests using it will be rejected.
	InvalidateImmediately param.Field[bool] `json:"invalidate_immediately"`
}

func (r AccountChallengeWidgetRotateSecretParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
