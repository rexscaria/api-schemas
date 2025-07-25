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

// AccountDlpProfileService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDlpProfileService] method instead.
type AccountDlpProfileService struct {
	Options    []option.RequestOption
	Custom     *AccountDlpProfileCustomService
	Predefined *AccountDlpProfilePredefinedService
}

// NewAccountDlpProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDlpProfileService(opts ...option.RequestOption) (r *AccountDlpProfileService) {
	r = &AccountDlpProfileService{}
	r.Options = opts
	r.Custom = NewAccountDlpProfileCustomService(opts...)
	r.Predefined = NewAccountDlpProfilePredefinedService(opts...)
	return
}

// Fetches a DLP profile by ID.
func (r *AccountDlpProfileService) Get(ctx context.Context, accountID string, profileID string, opts ...option.RequestOption) (res *AccountDlpProfileGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/%s", accountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all DLP profiles in an account.
func (r *AccountDlpProfileService) List(ctx context.Context, accountID string, query AccountDlpProfileListParams, opts ...option.RequestOption) (res *AccountDlpProfileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Confidence string

const (
	ConfidenceLow      Confidence = "low"
	ConfidenceMedium   Confidence = "medium"
	ConfidenceHigh     Confidence = "high"
	ConfidenceVeryHigh Confidence = "very_high"
)

func (r Confidence) IsKnown() bool {
	switch r {
	case ConfidenceLow, ConfidenceMedium, ConfidenceHigh, ConfidenceVeryHigh:
		return true
	}
	return false
}

type Profile struct {
	// The id of the profile (uuid).
	ID string `json:"id,required" format:"uuid"`
	// This field can have the runtime type of [[]DlpEntry].
	Entries interface{} `json:"entries,required"`
	// The name of the profile.
	Name             string      `json:"name,required"`
	Type             ProfileType `json:"type,required"`
	AIContextEnabled bool        `json:"ai_context_enabled"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   int64      `json:"allowed_match_count"`
	ConfidenceThreshold Confidence `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// When the profile was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description,nullable"`
	OcrEnabled  bool   `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool `json:"open_access"`
	// When the profile was lasted updated.
	UpdatedAt time.Time   `json:"updated_at" format:"date-time"`
	JSON      profileJSON `json:"-"`
	union     ProfileUnion
}

// profileJSON contains the JSON metadata for the struct [Profile]
type profileJSON struct {
	ID                  apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	Type                apijson.Field
	AIContextEnabled    apijson.Field
	AllowedMatchCount   apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	CreatedAt           apijson.Field
	Description         apijson.Field
	OcrEnabled          apijson.Field
	OpenAccess          apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r profileJSON) RawJSON() string {
	return r.raw
}

func (r *Profile) UnmarshalJSON(data []byte) (err error) {
	*r = Profile{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProfileUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [ProfileCustomProfile],
// [ProfilePredefinedProfile], [ProfileIntegrationProfile].
func (r Profile) AsUnion() ProfileUnion {
	return r.union
}

// Union satisfied by [ProfileCustomProfile], [ProfilePredefinedProfile] or
// [ProfileIntegrationProfile].
type ProfileUnion interface {
	implementsProfile()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfileUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfilePredefinedProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfile{}),
		},
	)
}

type ProfileCustomProfile struct {
	// The id of the profile (uuid).
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// When the profile was created.
	CreatedAt time.Time  `json:"created_at,required" format:"date-time"`
	Entries   []DlpEntry `json:"entries,required"`
	// The name of the profile.
	Name       string                   `json:"name,required"`
	OcrEnabled bool                     `json:"ocr_enabled,required"`
	Type       ProfileCustomProfileType `json:"type,required"`
	// When the profile was lasted updated.
	UpdatedAt           time.Time  `json:"updated_at,required" format:"date-time"`
	AIContextEnabled    bool       `json:"ai_context_enabled"`
	ConfidenceThreshold Confidence `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// The description of the profile.
	Description string                   `json:"description,nullable"`
	JSON        profileCustomProfileJSON `json:"-"`
}

// profileCustomProfileJSON contains the JSON metadata for the struct
// [ProfileCustomProfile]
type profileCustomProfileJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	CreatedAt           apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	OcrEnabled          apijson.Field
	Type                apijson.Field
	UpdatedAt           apijson.Field
	AIContextEnabled    apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	Description         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProfileCustomProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfile) implementsProfile() {}

type ProfileCustomProfileType string

const (
	ProfileCustomProfileTypeCustom ProfileCustomProfileType = "custom"
)

func (r ProfileCustomProfileType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileTypeCustom:
		return true
	}
	return false
}

type ProfilePredefinedProfile struct {
	// The id of the predefined profile (uuid).
	ID                string     `json:"id,required" format:"uuid"`
	AllowedMatchCount int64      `json:"allowed_match_count,required"`
	Entries           []DlpEntry `json:"entries,required"`
	// The name of the predefined profile.
	Name                string                       `json:"name,required"`
	Type                ProfilePredefinedProfileType `json:"type,required"`
	AIContextEnabled    bool                         `json:"ai_context_enabled"`
	ConfidenceThreshold Confidence                   `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	OcrEnabled       bool             `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool                         `json:"open_access"`
	JSON       profilePredefinedProfileJSON `json:"-"`
}

// profilePredefinedProfileJSON contains the JSON metadata for the struct
// [ProfilePredefinedProfile]
type profilePredefinedProfileJSON struct {
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

func (r *ProfilePredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedProfile) implementsProfile() {}

type ProfilePredefinedProfileType string

const (
	ProfilePredefinedProfileTypePredefined ProfilePredefinedProfileType = "predefined"
)

func (r ProfilePredefinedProfileType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileTypePredefined:
		return true
	}
	return false
}

type ProfileIntegrationProfile struct {
	ID        string                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                     `json:"created_at,required" format:"date-time"`
	Entries   []DlpEntry                    `json:"entries,required"`
	Name      string                        `json:"name,required"`
	Type      ProfileIntegrationProfileType `json:"type,required"`
	UpdatedAt time.Time                     `json:"updated_at,required" format:"date-time"`
	// The description of the profile.
	Description string                        `json:"description,nullable"`
	JSON        profileIntegrationProfileJSON `json:"-"`
}

// profileIntegrationProfileJSON contains the JSON metadata for the struct
// [ProfileIntegrationProfile]
type profileIntegrationProfileJSON struct {
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

func (r *ProfileIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfile) implementsProfile() {}

type ProfileIntegrationProfileType string

const (
	ProfileIntegrationProfileTypeIntegration ProfileIntegrationProfileType = "integration"
)

func (r ProfileIntegrationProfileType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileTypeIntegration:
		return true
	}
	return false
}

type ProfileType string

const (
	ProfileTypeCustom      ProfileType = "custom"
	ProfileTypePredefined  ProfileType = "predefined"
	ProfileTypeIntegration ProfileType = "integration"
)

func (r ProfileType) IsKnown() bool {
	switch r {
	case ProfileTypeCustom, ProfileTypePredefined, ProfileTypeIntegration:
		return true
	}
	return false
}

type AccountDlpProfileGetResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpProfileGetResponseSuccess `json:"success,required"`
	Result  Profile                             `json:"result"`
	JSON    accountDlpProfileGetResponseJSON    `json:"-"`
}

// accountDlpProfileGetResponseJSON contains the JSON metadata for the struct
// [AccountDlpProfileGetResponse]
type accountDlpProfileGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpProfileGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDlpProfileGetResponseSuccess bool

const (
	AccountDlpProfileGetResponseSuccessTrue AccountDlpProfileGetResponseSuccess = true
)

func (r AccountDlpProfileGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpProfileGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpProfileListResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpProfileListResponseSuccess `json:"success,required"`
	Result  []Profile                            `json:"result"`
	JSON    accountDlpProfileListResponseJSON    `json:"-"`
}

// accountDlpProfileListResponseJSON contains the JSON metadata for the struct
// [AccountDlpProfileListResponse]
type accountDlpProfileListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpProfileListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDlpProfileListResponseSuccess bool

const (
	AccountDlpProfileListResponseSuccessTrue AccountDlpProfileListResponseSuccess = true
)

func (r AccountDlpProfileListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpProfileListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpProfileListParams struct {
	// Return all profiles, including those that current account does not have access
	// to.
	All param.Field[bool] `query:"all"`
}

// URLQuery serializes [AccountDlpProfileListParams]'s query parameters as
// `url.Values`.
func (r AccountDlpProfileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
