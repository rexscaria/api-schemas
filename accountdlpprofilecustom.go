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
	"github.com/tidwall/gjson"
)

// AccountDlpProfileCustomService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDlpProfileCustomService] method instead.
type AccountDlpProfileCustomService struct {
	Options []option.RequestOption
}

// NewAccountDlpProfileCustomService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDlpProfileCustomService(opts ...option.RequestOption) (r *AccountDlpProfileCustomService) {
	r = &AccountDlpProfileCustomService{}
	r.Options = opts
	return
}

// Creates a DLP custom profile.
func (r *AccountDlpProfileCustomService) New(ctx context.Context, accountID string, body AccountDlpProfileCustomNewParams, opts ...option.RequestOption) (res *AccountDlpProfileCustomNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a custom DLP profile by id.
func (r *AccountDlpProfileCustomService) Get(ctx context.Context, accountID string, profileID string, opts ...option.RequestOption) (res *AccountDlpProfileCustomGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", accountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a DLP custom profile.
func (r *AccountDlpProfileCustomService) Update(ctx context.Context, accountID string, profileID string, body AccountDlpProfileCustomUpdateParams, opts ...option.RequestOption) (res *AccountDlpProfileCustomUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", accountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a DLP custom profile.
func (r *AccountDlpProfileCustomService) Delete(ctx context.Context, accountID string, profileID string, opts ...option.RequestOption) (res *AccountDlpProfileCustomDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", accountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type ContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip ContextAwarenessSkip `json:"skip,required"`
	JSON contextAwarenessJSON `json:"-"`
}

// contextAwarenessJSON contains the JSON metadata for the struct
// [ContextAwareness]
type contextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type ContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                     `json:"files,required"`
	JSON  contextAwarenessSkipJSON `json:"-"`
}

// contextAwarenessSkipJSON contains the JSON metadata for the struct
// [ContextAwarenessSkip]
type contextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contextAwarenessSkipJSON) RawJSON() string {
	return r.raw
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type ContextAwarenessParam struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip param.Field[ContextAwarenessSkipParam] `json:"skip,required"`
}

func (r ContextAwarenessParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Content types to exclude from context analysis and return all matches.
type ContextAwarenessSkipParam struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files param.Field[bool] `json:"files,required"`
}

func (r ContextAwarenessSkipParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewCustomEntryParam struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
}

func (r NewCustomEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewCustomEntryParam) implementsNewCustomProfileEntriesUnionParam() {}

func (r NewCustomEntryParam) implementsAccountDlpProfileCustomUpdateParamsEntryUnion() {}

type NewCustomProfileParam struct {
	Entries          param.Field[[]NewCustomProfileEntriesUnionParam] `json:"entries,required"`
	Name             param.Field[string]                              `json:"name,required"`
	AIContextEnabled param.Field[bool]                                `json:"ai_context_enabled"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile
	Description param.Field[string] `json:"description"`
	OcrEnabled  param.Field[bool]   `json:"ocr_enabled"`
	// Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your
	// Microsoft Information Protection profiles).
	SharedEntries param.Field[[]NewCustomProfileSharedEntriesUnionParam] `json:"shared_entries"`
}

func (r NewCustomProfileParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewCustomProfileParam) implementsAccountDlpProfileCustomNewParamsBodyUnion() {}

type NewCustomProfileEntryParam struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern"`
	Words   param.Field[interface{}]  `json:"words"`
}

func (r NewCustomProfileEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewCustomProfileEntryParam) implementsNewCustomProfileEntriesUnionParam() {}

// Satisfied by [NewCustomEntryParam],
// [NewCustomProfileEntriesDlpNewWordListEntryParam], [NewCustomProfileEntryParam].
type NewCustomProfileEntriesUnionParam interface {
	implementsNewCustomProfileEntriesUnionParam()
}

type NewCustomProfileEntriesDlpNewWordListEntryParam struct {
	Enabled param.Field[bool]     `json:"enabled,required"`
	Name    param.Field[string]   `json:"name,required"`
	Words   param.Field[[]string] `json:"words,required"`
}

func (r NewCustomProfileEntriesDlpNewWordListEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewCustomProfileEntriesDlpNewWordListEntryParam) implementsNewCustomProfileEntriesUnionParam() {
}

type NewCustomProfileSharedEntryParam struct {
	Enabled   param.Field[bool]                                   `json:"enabled,required"`
	EntryID   param.Field[string]                                 `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[NewCustomProfileSharedEntriesEntryType] `json:"entry_type,required"`
}

func (r NewCustomProfileSharedEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewCustomProfileSharedEntryParam) implementsNewCustomProfileSharedEntriesUnionParam() {}

// Satisfied by [NewCustomProfileSharedEntriesObjectParam],
// [NewCustomProfileSharedEntriesObjectParam],
// [NewCustomProfileSharedEntriesObjectParam],
// [NewCustomProfileSharedEntriesObjectParam], [NewCustomProfileSharedEntryParam].
type NewCustomProfileSharedEntriesUnionParam interface {
	implementsNewCustomProfileSharedEntriesUnionParam()
}

type NewCustomProfileSharedEntriesObjectParam struct {
	Enabled   param.Field[bool]                                         `json:"enabled,required"`
	EntryID   param.Field[string]                                       `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[NewCustomProfileSharedEntriesObjectEntryType] `json:"entry_type,required"`
}

func (r NewCustomProfileSharedEntriesObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewCustomProfileSharedEntriesObjectParam) implementsNewCustomProfileSharedEntriesUnionParam() {
}

type NewCustomProfileSharedEntriesObjectEntryType string

const (
	NewCustomProfileSharedEntriesObjectEntryTypeCustom NewCustomProfileSharedEntriesObjectEntryType = "custom"
)

func (r NewCustomProfileSharedEntriesObjectEntryType) IsKnown() bool {
	switch r {
	case NewCustomProfileSharedEntriesObjectEntryTypeCustom:
		return true
	}
	return false
}

type NewCustomProfileSharedEntriesEntryType string

const (
	NewCustomProfileSharedEntriesEntryTypeCustom      NewCustomProfileSharedEntriesEntryType = "custom"
	NewCustomProfileSharedEntriesEntryTypePredefined  NewCustomProfileSharedEntriesEntryType = "predefined"
	NewCustomProfileSharedEntriesEntryTypeIntegration NewCustomProfileSharedEntriesEntryType = "integration"
	NewCustomProfileSharedEntriesEntryTypeExactData   NewCustomProfileSharedEntriesEntryType = "exact_data"
)

func (r NewCustomProfileSharedEntriesEntryType) IsKnown() bool {
	switch r {
	case NewCustomProfileSharedEntriesEntryTypeCustom, NewCustomProfileSharedEntriesEntryTypePredefined, NewCustomProfileSharedEntriesEntryTypeIntegration, NewCustomProfileSharedEntriesEntryTypeExactData:
		return true
	}
	return false
}

type AccountDlpProfileCustomNewResponse struct {
	Result AccountDlpProfileCustomNewResponseResultUnion `json:"result"`
	JSON   accountDlpProfileCustomNewResponseJSON        `json:"-"`
	APIResponseSingleDlp
}

// accountDlpProfileCustomNewResponseJSON contains the JSON metadata for the struct
// [AccountDlpProfileCustomNewResponse]
type accountDlpProfileCustomNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpProfileCustomNewResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [AccountDlpProfileCustomNewResponseResultCustomProfile],
// [AccountDlpProfileCustomNewResponseResultPredefinedProfile],
// [AccountDlpProfileCustomNewResponseResultIntegrationProfile] or
// [AccountDlpProfileCustomNewResponseResultArray].
type AccountDlpProfileCustomNewResponseResultUnion interface {
	implementsAccountDlpProfileCustomNewResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountDlpProfileCustomNewResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountDlpProfileCustomNewResponseResultCustomProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountDlpProfileCustomNewResponseResultPredefinedProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountDlpProfileCustomNewResponseResultIntegrationProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountDlpProfileCustomNewResponseResultArray{}),
		},
	)
}

type AccountDlpProfileCustomNewResponseResultCustomProfile struct {
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness,required"`
	// When the profile was created
	CreatedAt time.Time  `json:"created_at,required" format:"date-time"`
	Entries   []DlpEntry `json:"entries,required"`
	// The name of the profile
	Name       string                                                    `json:"name,required"`
	OcrEnabled bool                                                      `json:"ocr_enabled,required"`
	Type       AccountDlpProfileCustomNewResponseResultCustomProfileType `json:"type,required"`
	// When the profile was lasted updated
	UpdatedAt           time.Time  `json:"updated_at,required" format:"date-time"`
	AIContextEnabled    bool       `json:"ai_context_enabled"`
	ConfidenceThreshold Confidence `json:"confidence_threshold"`
	// The description of the profile
	Description string                                                    `json:"description,nullable"`
	JSON        accountDlpProfileCustomNewResponseResultCustomProfileJSON `json:"-"`
}

// accountDlpProfileCustomNewResponseResultCustomProfileJSON contains the JSON
// metadata for the struct [AccountDlpProfileCustomNewResponseResultCustomProfile]
type accountDlpProfileCustomNewResponseResultCustomProfileJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	ContextAwareness    apijson.Field
	CreatedAt           apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	OcrEnabled          apijson.Field
	Type                apijson.Field
	UpdatedAt           apijson.Field
	AIContextEnabled    apijson.Field
	ConfidenceThreshold apijson.Field
	Description         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountDlpProfileCustomNewResponseResultCustomProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpProfileCustomNewResponseResultCustomProfileJSON) RawJSON() string {
	return r.raw
}

func (r AccountDlpProfileCustomNewResponseResultCustomProfile) implementsAccountDlpProfileCustomNewResponseResultUnion() {
}

type AccountDlpProfileCustomNewResponseResultCustomProfileType string

const (
	AccountDlpProfileCustomNewResponseResultCustomProfileTypeCustom AccountDlpProfileCustomNewResponseResultCustomProfileType = "custom"
)

func (r AccountDlpProfileCustomNewResponseResultCustomProfileType) IsKnown() bool {
	switch r {
	case AccountDlpProfileCustomNewResponseResultCustomProfileTypeCustom:
		return true
	}
	return false
}

type AccountDlpProfileCustomNewResponseResultPredefinedProfile struct {
	// The id of the predefined profile (uuid)
	ID                string     `json:"id,required" format:"uuid"`
	AllowedMatchCount int64      `json:"allowed_match_count,required"`
	Entries           []DlpEntry `json:"entries,required"`
	// The name of the predefined profile
	Name                string                                                        `json:"name,required"`
	Type                AccountDlpProfileCustomNewResponseResultPredefinedProfileType `json:"type,required"`
	AIContextEnabled    bool                                                          `json:"ai_context_enabled"`
	ConfidenceThreshold Confidence                                                    `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness"`
	OcrEnabled       bool             `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone
	OpenAccess bool                                                          `json:"open_access"`
	JSON       accountDlpProfileCustomNewResponseResultPredefinedProfileJSON `json:"-"`
}

// accountDlpProfileCustomNewResponseResultPredefinedProfileJSON contains the JSON
// metadata for the struct
// [AccountDlpProfileCustomNewResponseResultPredefinedProfile]
type accountDlpProfileCustomNewResponseResultPredefinedProfileJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	Type                apijson.Field
	AIContextEnabled    apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	OcrEnabled          apijson.Field
	OpenAccess          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountDlpProfileCustomNewResponseResultPredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpProfileCustomNewResponseResultPredefinedProfileJSON) RawJSON() string {
	return r.raw
}

func (r AccountDlpProfileCustomNewResponseResultPredefinedProfile) implementsAccountDlpProfileCustomNewResponseResultUnion() {
}

type AccountDlpProfileCustomNewResponseResultPredefinedProfileType string

const (
	AccountDlpProfileCustomNewResponseResultPredefinedProfileTypePredefined AccountDlpProfileCustomNewResponseResultPredefinedProfileType = "predefined"
)

func (r AccountDlpProfileCustomNewResponseResultPredefinedProfileType) IsKnown() bool {
	switch r {
	case AccountDlpProfileCustomNewResponseResultPredefinedProfileTypePredefined:
		return true
	}
	return false
}

type AccountDlpProfileCustomNewResponseResultIntegrationProfile struct {
	ID        string                                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                      `json:"created_at,required" format:"date-time"`
	Entries   []DlpEntry                                                     `json:"entries,required"`
	Name      string                                                         `json:"name,required"`
	Type      AccountDlpProfileCustomNewResponseResultIntegrationProfileType `json:"type,required"`
	UpdatedAt time.Time                                                      `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                                                         `json:"description,nullable"`
	JSON        accountDlpProfileCustomNewResponseResultIntegrationProfileJSON `json:"-"`
}

// accountDlpProfileCustomNewResponseResultIntegrationProfileJSON contains the JSON
// metadata for the struct
// [AccountDlpProfileCustomNewResponseResultIntegrationProfile]
type accountDlpProfileCustomNewResponseResultIntegrationProfileJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Entries     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomNewResponseResultIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpProfileCustomNewResponseResultIntegrationProfileJSON) RawJSON() string {
	return r.raw
}

func (r AccountDlpProfileCustomNewResponseResultIntegrationProfile) implementsAccountDlpProfileCustomNewResponseResultUnion() {
}

type AccountDlpProfileCustomNewResponseResultIntegrationProfileType string

const (
	AccountDlpProfileCustomNewResponseResultIntegrationProfileTypeIntegration AccountDlpProfileCustomNewResponseResultIntegrationProfileType = "integration"
)

func (r AccountDlpProfileCustomNewResponseResultIntegrationProfileType) IsKnown() bool {
	switch r {
	case AccountDlpProfileCustomNewResponseResultIntegrationProfileTypeIntegration:
		return true
	}
	return false
}

type AccountDlpProfileCustomNewResponseResultArray []Profile

func (r AccountDlpProfileCustomNewResponseResultArray) implementsAccountDlpProfileCustomNewResponseResultUnion() {
}

type AccountDlpProfileCustomGetResponse struct {
	Result Profile                                `json:"result"`
	JSON   accountDlpProfileCustomGetResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpProfileCustomGetResponseJSON contains the JSON metadata for the struct
// [AccountDlpProfileCustomGetResponse]
type accountDlpProfileCustomGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpProfileCustomGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpProfileCustomUpdateResponse struct {
	Result Profile                                   `json:"result"`
	JSON   accountDlpProfileCustomUpdateResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpProfileCustomUpdateResponseJSON contains the JSON metadata for the
// struct [AccountDlpProfileCustomUpdateResponse]
type accountDlpProfileCustomUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpProfileCustomUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpProfileCustomDeleteResponse struct {
	Result interface{}                               `json:"result,nullable"`
	JSON   accountDlpProfileCustomDeleteResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpProfileCustomDeleteResponseJSON contains the JSON metadata for the
// struct [AccountDlpProfileCustomDeleteResponse]
type accountDlpProfileCustomDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpProfileCustomDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpProfileCustomNewParams struct {
	Body AccountDlpProfileCustomNewParamsBodyUnion `json:"body,required"`
}

func (r AccountDlpProfileCustomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDlpProfileCustomNewParamsBody struct {
	AIContextEnabled param.Field[bool] `json:"ai_context_enabled"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile
	Description   param.Field[string]      `json:"description"`
	Entries       param.Field[interface{}] `json:"entries"`
	Name          param.Field[string]      `json:"name"`
	OcrEnabled    param.Field[bool]        `json:"ocr_enabled"`
	Profiles      param.Field[interface{}] `json:"profiles"`
	SharedEntries param.Field[interface{}] `json:"shared_entries"`
}

func (r AccountDlpProfileCustomNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpProfileCustomNewParamsBody) implementsAccountDlpProfileCustomNewParamsBodyUnion() {}

// Satisfied by [AccountDlpProfileCustomNewParamsBodyProfiles],
// [NewCustomProfileParam], [AccountDlpProfileCustomNewParamsBody].
type AccountDlpProfileCustomNewParamsBodyUnion interface {
	implementsAccountDlpProfileCustomNewParamsBodyUnion()
}

// Deprecated: deprecated
type AccountDlpProfileCustomNewParamsBodyProfiles struct {
	Profiles param.Field[[]NewCustomProfileParam] `json:"profiles,required"`
}

func (r AccountDlpProfileCustomNewParamsBodyProfiles) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpProfileCustomNewParamsBodyProfiles) implementsAccountDlpProfileCustomNewParamsBodyUnion() {
}

type AccountDlpProfileCustomUpdateParams struct {
	Name                param.Field[string] `json:"name,required"`
	AIContextEnabled    param.Field[bool]   `json:"ai_context_enabled"`
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile
	Description param.Field[string] `json:"description"`
	// Custom entries from this profile. If this field is omitted, entries owned by
	// this profile will not be changed.
	Entries    param.Field[[]AccountDlpProfileCustomUpdateParamsEntryUnion] `json:"entries"`
	OcrEnabled param.Field[bool]                                            `json:"ocr_enabled"`
	// Other entries, e.g. predefined or integration.
	SharedEntries param.Field[[]AccountDlpProfileCustomUpdateParamsSharedEntryUnion] `json:"shared_entries"`
}

func (r AccountDlpProfileCustomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [AccountDlpProfileCustomUpdateParamsEntriesDlpNewCustomEntryWithID],
// [NewCustomEntryParam].
type AccountDlpProfileCustomUpdateParamsEntryUnion interface {
	implementsAccountDlpProfileCustomUpdateParamsEntryUnion()
}

type AccountDlpProfileCustomUpdateParamsEntriesDlpNewCustomEntryWithID struct {
	EntryID param.Field[string] `json:"entry_id,required" format:"uuid"`
	NewCustomEntryParam
}

func (r AccountDlpProfileCustomUpdateParamsEntriesDlpNewCustomEntryWithID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpProfileCustomUpdateParamsEntriesDlpNewCustomEntryWithID) implementsAccountDlpProfileCustomUpdateParamsEntryUnion() {
}

type AccountDlpProfileCustomUpdateParamsSharedEntry struct {
	Enabled   param.Field[bool]                                                      `json:"enabled,required"`
	EntryID   param.Field[string]                                                    `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[AccountDlpProfileCustomUpdateParamsSharedEntriesEntryType] `json:"entry_type,required"`
}

func (r AccountDlpProfileCustomUpdateParamsSharedEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpProfileCustomUpdateParamsSharedEntry) implementsAccountDlpProfileCustomUpdateParamsSharedEntryUnion() {
}

// Satisfied by [AccountDlpProfileCustomUpdateParamsSharedEntriesObject],
// [AccountDlpProfileCustomUpdateParamsSharedEntriesObject],
// [AccountDlpProfileCustomUpdateParamsSharedEntriesObject],
// [AccountDlpProfileCustomUpdateParamsSharedEntry].
type AccountDlpProfileCustomUpdateParamsSharedEntryUnion interface {
	implementsAccountDlpProfileCustomUpdateParamsSharedEntryUnion()
}

type AccountDlpProfileCustomUpdateParamsSharedEntriesObject struct {
	Enabled   param.Field[bool]                                                            `json:"enabled,required"`
	EntryID   param.Field[string]                                                          `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[AccountDlpProfileCustomUpdateParamsSharedEntriesObjectEntryType] `json:"entry_type,required"`
}

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesObject) implementsAccountDlpProfileCustomUpdateParamsSharedEntryUnion() {
}

type AccountDlpProfileCustomUpdateParamsSharedEntriesObjectEntryType string

const (
	AccountDlpProfileCustomUpdateParamsSharedEntriesObjectEntryTypePredefined AccountDlpProfileCustomUpdateParamsSharedEntriesObjectEntryType = "predefined"
)

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesObjectEntryType) IsKnown() bool {
	switch r {
	case AccountDlpProfileCustomUpdateParamsSharedEntriesObjectEntryTypePredefined:
		return true
	}
	return false
}

type AccountDlpProfileCustomUpdateParamsSharedEntriesEntryType string

const (
	AccountDlpProfileCustomUpdateParamsSharedEntriesEntryTypePredefined  AccountDlpProfileCustomUpdateParamsSharedEntriesEntryType = "predefined"
	AccountDlpProfileCustomUpdateParamsSharedEntriesEntryTypeIntegration AccountDlpProfileCustomUpdateParamsSharedEntriesEntryType = "integration"
	AccountDlpProfileCustomUpdateParamsSharedEntriesEntryTypeExactData   AccountDlpProfileCustomUpdateParamsSharedEntriesEntryType = "exact_data"
)

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesEntryType) IsKnown() bool {
	switch r {
	case AccountDlpProfileCustomUpdateParamsSharedEntriesEntryTypePredefined, AccountDlpProfileCustomUpdateParamsSharedEntriesEntryTypeIntegration, AccountDlpProfileCustomUpdateParamsSharedEntriesEntryTypeExactData:
		return true
	}
	return false
}
