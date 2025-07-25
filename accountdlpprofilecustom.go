// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
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
//
// Deprecated: deprecated
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
//
// Deprecated: deprecated
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
	//
	// Deprecated: deprecated
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	OcrEnabled  param.Field[bool]   `json:"ocr_enabled"`
	// Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your
	// Microsoft Information Protection profiles).
	SharedEntries param.Field[[]NewCustomProfileSharedEntriesUnionParam] `json:"shared_entries"`
}

func (r NewCustomProfileParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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

// Satisfied by [NewCustomProfileSharedEntriesCustomParam],
// [NewCustomProfileSharedEntriesPredefinedParam],
// [NewCustomProfileSharedEntriesIntegrationParam],
// [NewCustomProfileSharedEntriesExactDataParam],
// [NewCustomProfileSharedEntriesObjectParam], [NewCustomProfileSharedEntryParam].
type NewCustomProfileSharedEntriesUnionParam interface {
	implementsNewCustomProfileSharedEntriesUnionParam()
}

type NewCustomProfileSharedEntriesCustomParam struct {
	Enabled   param.Field[bool]                                         `json:"enabled,required"`
	EntryID   param.Field[string]                                       `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[NewCustomProfileSharedEntriesCustomEntryType] `json:"entry_type,required"`
}

func (r NewCustomProfileSharedEntriesCustomParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewCustomProfileSharedEntriesCustomParam) implementsNewCustomProfileSharedEntriesUnionParam() {
}

type NewCustomProfileSharedEntriesCustomEntryType string

const (
	NewCustomProfileSharedEntriesCustomEntryTypeCustom NewCustomProfileSharedEntriesCustomEntryType = "custom"
)

func (r NewCustomProfileSharedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case NewCustomProfileSharedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type NewCustomProfileSharedEntriesPredefinedParam struct {
	Enabled   param.Field[bool]                                             `json:"enabled,required"`
	EntryID   param.Field[string]                                           `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[NewCustomProfileSharedEntriesPredefinedEntryType] `json:"entry_type,required"`
}

func (r NewCustomProfileSharedEntriesPredefinedParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewCustomProfileSharedEntriesPredefinedParam) implementsNewCustomProfileSharedEntriesUnionParam() {
}

type NewCustomProfileSharedEntriesPredefinedEntryType string

const (
	NewCustomProfileSharedEntriesPredefinedEntryTypePredefined NewCustomProfileSharedEntriesPredefinedEntryType = "predefined"
)

func (r NewCustomProfileSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case NewCustomProfileSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type NewCustomProfileSharedEntriesIntegrationParam struct {
	Enabled   param.Field[bool]                                              `json:"enabled,required"`
	EntryID   param.Field[string]                                            `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[NewCustomProfileSharedEntriesIntegrationEntryType] `json:"entry_type,required"`
}

func (r NewCustomProfileSharedEntriesIntegrationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewCustomProfileSharedEntriesIntegrationParam) implementsNewCustomProfileSharedEntriesUnionParam() {
}

type NewCustomProfileSharedEntriesIntegrationEntryType string

const (
	NewCustomProfileSharedEntriesIntegrationEntryTypeIntegration NewCustomProfileSharedEntriesIntegrationEntryType = "integration"
)

func (r NewCustomProfileSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case NewCustomProfileSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type NewCustomProfileSharedEntriesExactDataParam struct {
	Enabled   param.Field[bool]                                            `json:"enabled,required"`
	EntryID   param.Field[string]                                          `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[NewCustomProfileSharedEntriesExactDataEntryType] `json:"entry_type,required"`
}

func (r NewCustomProfileSharedEntriesExactDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewCustomProfileSharedEntriesExactDataParam) implementsNewCustomProfileSharedEntriesUnionParam() {
}

type NewCustomProfileSharedEntriesExactDataEntryType string

const (
	NewCustomProfileSharedEntriesExactDataEntryTypeExactData NewCustomProfileSharedEntriesExactDataEntryType = "exact_data"
)

func (r NewCustomProfileSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case NewCustomProfileSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
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
	NewCustomProfileSharedEntriesObjectEntryTypeDocumentFingerprint NewCustomProfileSharedEntriesObjectEntryType = "document_fingerprint"
)

func (r NewCustomProfileSharedEntriesObjectEntryType) IsKnown() bool {
	switch r {
	case NewCustomProfileSharedEntriesObjectEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type NewCustomProfileSharedEntriesEntryType string

const (
	NewCustomProfileSharedEntriesEntryTypeCustom              NewCustomProfileSharedEntriesEntryType = "custom"
	NewCustomProfileSharedEntriesEntryTypePredefined          NewCustomProfileSharedEntriesEntryType = "predefined"
	NewCustomProfileSharedEntriesEntryTypeIntegration         NewCustomProfileSharedEntriesEntryType = "integration"
	NewCustomProfileSharedEntriesEntryTypeExactData           NewCustomProfileSharedEntriesEntryType = "exact_data"
	NewCustomProfileSharedEntriesEntryTypeDocumentFingerprint NewCustomProfileSharedEntriesEntryType = "document_fingerprint"
)

func (r NewCustomProfileSharedEntriesEntryType) IsKnown() bool {
	switch r {
	case NewCustomProfileSharedEntriesEntryTypeCustom, NewCustomProfileSharedEntriesEntryTypePredefined, NewCustomProfileSharedEntriesEntryTypeIntegration, NewCustomProfileSharedEntriesEntryTypeExactData, NewCustomProfileSharedEntriesEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type AccountDlpProfileCustomNewResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpProfileCustomNewResponseSuccess `json:"success,required"`
	Result  Profile                                   `json:"result"`
	JSON    accountDlpProfileCustomNewResponseJSON    `json:"-"`
}

// accountDlpProfileCustomNewResponseJSON contains the JSON metadata for the struct
// [AccountDlpProfileCustomNewResponse]
type accountDlpProfileCustomNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful.
type AccountDlpProfileCustomNewResponseSuccess bool

const (
	AccountDlpProfileCustomNewResponseSuccessTrue AccountDlpProfileCustomNewResponseSuccess = true
)

func (r AccountDlpProfileCustomNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpProfileCustomNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpProfileCustomGetResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpProfileCustomGetResponseSuccess `json:"success,required"`
	Result  Profile                                   `json:"result"`
	JSON    accountDlpProfileCustomGetResponseJSON    `json:"-"`
}

// accountDlpProfileCustomGetResponseJSON contains the JSON metadata for the struct
// [AccountDlpProfileCustomGetResponse]
type accountDlpProfileCustomGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful.
type AccountDlpProfileCustomGetResponseSuccess bool

const (
	AccountDlpProfileCustomGetResponseSuccessTrue AccountDlpProfileCustomGetResponseSuccess = true
)

func (r AccountDlpProfileCustomGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpProfileCustomGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpProfileCustomUpdateResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpProfileCustomUpdateResponseSuccess `json:"success,required"`
	Result  Profile                                      `json:"result"`
	JSON    accountDlpProfileCustomUpdateResponseJSON    `json:"-"`
}

// accountDlpProfileCustomUpdateResponseJSON contains the JSON metadata for the
// struct [AccountDlpProfileCustomUpdateResponse]
type accountDlpProfileCustomUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful.
type AccountDlpProfileCustomUpdateResponseSuccess bool

const (
	AccountDlpProfileCustomUpdateResponseSuccessTrue AccountDlpProfileCustomUpdateResponseSuccess = true
)

func (r AccountDlpProfileCustomUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpProfileCustomUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpProfileCustomDeleteResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpProfileCustomDeleteResponseSuccess `json:"success,required"`
	Result  interface{}                                  `json:"result,nullable"`
	JSON    accountDlpProfileCustomDeleteResponseJSON    `json:"-"`
}

// accountDlpProfileCustomDeleteResponseJSON contains the JSON metadata for the
// struct [AccountDlpProfileCustomDeleteResponse]
type accountDlpProfileCustomDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful.
type AccountDlpProfileCustomDeleteResponseSuccess bool

const (
	AccountDlpProfileCustomDeleteResponseSuccessTrue AccountDlpProfileCustomDeleteResponseSuccess = true
)

func (r AccountDlpProfileCustomDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpProfileCustomDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpProfileCustomNewParams struct {
	NewCustomProfile NewCustomProfileParam `json:"new_custom_profile,required"`
}

func (r AccountDlpProfileCustomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.NewCustomProfile)
}

type AccountDlpProfileCustomUpdateParams struct {
	Name                param.Field[string] `json:"name,required"`
	AIContextEnabled    param.Field[bool]   `json:"ai_context_enabled"`
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile.
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

type AccountDlpProfileCustomUpdateParamsEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
	EntryID param.Field[string]       `json:"entry_id" format:"uuid"`
}

func (r AccountDlpProfileCustomUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpProfileCustomUpdateParamsEntry) implementsAccountDlpProfileCustomUpdateParamsEntryUnion() {
}

// Satisfied by
// [AccountDlpProfileCustomUpdateParamsEntriesDlpNewCustomEntryWithID],
// [NewCustomEntryParam], [AccountDlpProfileCustomUpdateParamsEntry].
type AccountDlpProfileCustomUpdateParamsEntryUnion interface {
	implementsAccountDlpProfileCustomUpdateParamsEntryUnion()
}

type AccountDlpProfileCustomUpdateParamsEntriesDlpNewCustomEntryWithID struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	EntryID param.Field[string]       `json:"entry_id,required" format:"uuid"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
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

// Satisfied by [AccountDlpProfileCustomUpdateParamsSharedEntriesPredefined],
// [AccountDlpProfileCustomUpdateParamsSharedEntriesIntegration],
// [AccountDlpProfileCustomUpdateParamsSharedEntriesExactData],
// [AccountDlpProfileCustomUpdateParamsSharedEntriesObject],
// [AccountDlpProfileCustomUpdateParamsSharedEntry].
type AccountDlpProfileCustomUpdateParamsSharedEntryUnion interface {
	implementsAccountDlpProfileCustomUpdateParamsSharedEntryUnion()
}

type AccountDlpProfileCustomUpdateParamsSharedEntriesPredefined struct {
	Enabled   param.Field[bool]                                                                `json:"enabled,required"`
	EntryID   param.Field[string]                                                              `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[AccountDlpProfileCustomUpdateParamsSharedEntriesPredefinedEntryType] `json:"entry_type,required"`
}

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesPredefined) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesPredefined) implementsAccountDlpProfileCustomUpdateParamsSharedEntryUnion() {
}

type AccountDlpProfileCustomUpdateParamsSharedEntriesPredefinedEntryType string

const (
	AccountDlpProfileCustomUpdateParamsSharedEntriesPredefinedEntryTypePredefined AccountDlpProfileCustomUpdateParamsSharedEntriesPredefinedEntryType = "predefined"
)

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case AccountDlpProfileCustomUpdateParamsSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type AccountDlpProfileCustomUpdateParamsSharedEntriesIntegration struct {
	Enabled   param.Field[bool]                                                                 `json:"enabled,required"`
	EntryID   param.Field[string]                                                               `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[AccountDlpProfileCustomUpdateParamsSharedEntriesIntegrationEntryType] `json:"entry_type,required"`
}

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesIntegration) implementsAccountDlpProfileCustomUpdateParamsSharedEntryUnion() {
}

type AccountDlpProfileCustomUpdateParamsSharedEntriesIntegrationEntryType string

const (
	AccountDlpProfileCustomUpdateParamsSharedEntriesIntegrationEntryTypeIntegration AccountDlpProfileCustomUpdateParamsSharedEntriesIntegrationEntryType = "integration"
)

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case AccountDlpProfileCustomUpdateParamsSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type AccountDlpProfileCustomUpdateParamsSharedEntriesExactData struct {
	Enabled   param.Field[bool]                                                               `json:"enabled,required"`
	EntryID   param.Field[string]                                                             `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[AccountDlpProfileCustomUpdateParamsSharedEntriesExactDataEntryType] `json:"entry_type,required"`
}

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesExactData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesExactData) implementsAccountDlpProfileCustomUpdateParamsSharedEntryUnion() {
}

type AccountDlpProfileCustomUpdateParamsSharedEntriesExactDataEntryType string

const (
	AccountDlpProfileCustomUpdateParamsSharedEntriesExactDataEntryTypeExactData AccountDlpProfileCustomUpdateParamsSharedEntriesExactDataEntryType = "exact_data"
)

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case AccountDlpProfileCustomUpdateParamsSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
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
	AccountDlpProfileCustomUpdateParamsSharedEntriesObjectEntryTypeDocumentFingerprint AccountDlpProfileCustomUpdateParamsSharedEntriesObjectEntryType = "document_fingerprint"
)

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesObjectEntryType) IsKnown() bool {
	switch r {
	case AccountDlpProfileCustomUpdateParamsSharedEntriesObjectEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type AccountDlpProfileCustomUpdateParamsSharedEntriesEntryType string

const (
	AccountDlpProfileCustomUpdateParamsSharedEntriesEntryTypePredefined          AccountDlpProfileCustomUpdateParamsSharedEntriesEntryType = "predefined"
	AccountDlpProfileCustomUpdateParamsSharedEntriesEntryTypeIntegration         AccountDlpProfileCustomUpdateParamsSharedEntriesEntryType = "integration"
	AccountDlpProfileCustomUpdateParamsSharedEntriesEntryTypeExactData           AccountDlpProfileCustomUpdateParamsSharedEntriesEntryType = "exact_data"
	AccountDlpProfileCustomUpdateParamsSharedEntriesEntryTypeDocumentFingerprint AccountDlpProfileCustomUpdateParamsSharedEntriesEntryType = "document_fingerprint"
)

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesEntryType) IsKnown() bool {
	switch r {
	case AccountDlpProfileCustomUpdateParamsSharedEntriesEntryTypePredefined, AccountDlpProfileCustomUpdateParamsSharedEntriesEntryTypeIntegration, AccountDlpProfileCustomUpdateParamsSharedEntriesEntryTypeExactData, AccountDlpProfileCustomUpdateParamsSharedEntriesEntryTypeDocumentFingerprint:
		return true
	}
	return false
}
