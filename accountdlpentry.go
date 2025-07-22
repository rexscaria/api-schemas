// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountDlpEntryService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDlpEntryService] method instead.
type AccountDlpEntryService struct {
	Options []option.RequestOption
}

// NewAccountDlpEntryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountDlpEntryService(opts ...option.RequestOption) (r *AccountDlpEntryService) {
	r = &AccountDlpEntryService{}
	r.Options = opts
	return
}

// Creates a DLP custom entry.
func (r *AccountDlpEntryService) New(ctx context.Context, accountID string, body AccountDlpEntryNewParams, opts ...option.RequestOption) (res *AccountDlpEntryNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a DLP entry by ID
func (r *AccountDlpEntryService) Get(ctx context.Context, accountID string, entryID string, opts ...option.RequestOption) (res *AccountDlpEntryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/%s", accountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a DLP entry.
func (r *AccountDlpEntryService) Update(ctx context.Context, accountID string, entryID string, body AccountDlpEntryUpdateParams, opts ...option.RequestOption) (res *AccountDlpEntryUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/%s", accountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all DLP entries in an account.
func (r *AccountDlpEntryService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountDlpEntryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a DLP custom entry.
func (r *AccountDlpEntryService) Delete(ctx context.Context, accountID string, entryID string, opts ...option.RequestOption) (res *AccountDlpEntryDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/%s", accountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CustomEntry struct {
	ID        string          `json:"id,required" format:"uuid"`
	CreatedAt time.Time       `json:"created_at,required" format:"date-time"`
	Enabled   bool            `json:"enabled,required"`
	Name      string          `json:"name,required"`
	Pattern   Pattern         `json:"pattern,required"`
	UpdatedAt time.Time       `json:"updated_at,required" format:"date-time"`
	ProfileID string          `json:"profile_id,nullable" format:"uuid"`
	JSON      customEntryJSON `json:"-"`
}

// customEntryJSON contains the JSON metadata for the struct [CustomEntry]
type customEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customEntryJSON) RawJSON() string {
	return r.raw
}

type DlpEntry struct {
	ID string `json:"id" format:"uuid"`
	// This field can have the runtime type of [DlpEntryPredefinedEntryConfidence].
	Confidence interface{}  `json:"confidence"`
	CreatedAt  time.Time    `json:"created_at" format:"date-time"`
	Enabled    bool         `json:"enabled"`
	Name       string       `json:"name"`
	Pattern    Pattern      `json:"pattern"`
	ProfileID  string       `json:"profile_id,nullable" format:"uuid"`
	Secret     bool         `json:"secret"`
	Type       DlpEntryType `json:"type"`
	UpdatedAt  time.Time    `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}  `json:"word_list"`
	JSON     dlpEntryJSON `json:"-"`
	union    DlpEntryUnion
}

// dlpEntryJSON contains the JSON metadata for the struct [DlpEntry]
type dlpEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	Secret      apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dlpEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DlpEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DlpEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DlpEntryUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [DlpEntryCustomEntry],
// [DlpEntryPredefinedEntry], [DlpEntryIntegrationEntry], [DlpEntryExactDataEntry],
// [DlpEntryWordListEntry].
func (r DlpEntry) AsUnion() DlpEntryUnion {
	return r.union
}

// Union satisfied by [DlpEntryCustomEntry], [DlpEntryPredefinedEntry],
// [DlpEntryIntegrationEntry], [DlpEntryExactDataEntry] or [DlpEntryWordListEntry].
type DlpEntryUnion interface {
	implementsDlpEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DlpEntryUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DlpEntryCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DlpEntryPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DlpEntryIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DlpEntryExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DlpEntryWordListEntry{}),
		},
	)
}

type DlpEntryCustomEntry struct {
	Type DlpEntryCustomEntryType `json:"type,required"`
	JSON dlpEntryCustomEntryJSON `json:"-"`
	CustomEntry
}

// dlpEntryCustomEntryJSON contains the JSON metadata for the struct
// [DlpEntryCustomEntry]
type dlpEntryCustomEntryJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpEntryCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DlpEntryCustomEntry) implementsDlpEntry() {}

type DlpEntryCustomEntryType string

const (
	DlpEntryCustomEntryTypeCustom DlpEntryCustomEntryType = "custom"
)

func (r DlpEntryCustomEntryType) IsKnown() bool {
	switch r {
	case DlpEntryCustomEntryTypeCustom:
		return true
	}
	return false
}

type DlpEntryPredefinedEntry struct {
	ID         string                            `json:"id,required" format:"uuid"`
	Confidence DlpEntryPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                              `json:"enabled,required"`
	Name       string                            `json:"name,required"`
	Type       DlpEntryPredefinedEntryType       `json:"type,required"`
	ProfileID  string                            `json:"profile_id,nullable" format:"uuid"`
	JSON       dlpEntryPredefinedEntryJSON       `json:"-"`
}

// dlpEntryPredefinedEntryJSON contains the JSON metadata for the struct
// [DlpEntryPredefinedEntry]
type dlpEntryPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpEntryPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DlpEntryPredefinedEntry) implementsDlpEntry() {}

type DlpEntryPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service
	Available bool                                  `json:"available,required"`
	JSON      dlpEntryPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpEntryPredefinedEntryConfidenceJSON contains the JSON metadata for the struct
// [DlpEntryPredefinedEntryConfidence]
type dlpEntryPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DlpEntryPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DlpEntryPredefinedEntryType string

const (
	DlpEntryPredefinedEntryTypePredefined DlpEntryPredefinedEntryType = "predefined"
)

func (r DlpEntryPredefinedEntryType) IsKnown() bool {
	switch r {
	case DlpEntryPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DlpEntryIntegrationEntry struct {
	ID        string                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                         `json:"enabled,required"`
	Name      string                       `json:"name,required"`
	Type      DlpEntryIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                    `json:"updated_at,required" format:"date-time"`
	ProfileID string                       `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryIntegrationEntryJSON `json:"-"`
}

// dlpEntryIntegrationEntryJSON contains the JSON metadata for the struct
// [DlpEntryIntegrationEntry]
type dlpEntryIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpEntryIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DlpEntryIntegrationEntry) implementsDlpEntry() {}

type DlpEntryIntegrationEntryType string

const (
	DlpEntryIntegrationEntryTypeIntegration DlpEntryIntegrationEntryType = "integration"
)

func (r DlpEntryIntegrationEntryType) IsKnown() bool {
	switch r {
	case DlpEntryIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DlpEntryExactDataEntry struct {
	ID        string                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                       `json:"enabled,required"`
	Name      string                     `json:"name,required"`
	Secret    bool                       `json:"secret,required"`
	Type      DlpEntryExactDataEntryType `json:"type,required"`
	UpdatedAt time.Time                  `json:"updated_at,required" format:"date-time"`
	JSON      dlpEntryExactDataEntryJSON `json:"-"`
}

// dlpEntryExactDataEntryJSON contains the JSON metadata for the struct
// [DlpEntryExactDataEntry]
type dlpEntryExactDataEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpEntryExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DlpEntryExactDataEntry) implementsDlpEntry() {}

type DlpEntryExactDataEntryType string

const (
	DlpEntryExactDataEntryTypeExactData DlpEntryExactDataEntryType = "exact_data"
)

func (r DlpEntryExactDataEntryType) IsKnown() bool {
	switch r {
	case DlpEntryExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DlpEntryWordListEntry struct {
	ID        string                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                      `json:"enabled,required"`
	Name      string                    `json:"name,required"`
	Type      DlpEntryWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                 `json:"updated_at,required" format:"date-time"`
	WordList  interface{}               `json:"word_list,required"`
	ProfileID string                    `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryWordListEntryJSON `json:"-"`
}

// dlpEntryWordListEntryJSON contains the JSON metadata for the struct
// [DlpEntryWordListEntry]
type dlpEntryWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DlpEntryWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DlpEntryWordListEntry) implementsDlpEntry() {}

type DlpEntryWordListEntryType string

const (
	DlpEntryWordListEntryTypeWordList DlpEntryWordListEntryType = "word_list"
)

func (r DlpEntryWordListEntryType) IsKnown() bool {
	switch r {
	case DlpEntryWordListEntryTypeWordList:
		return true
	}
	return false
}

type DlpEntryType string

const (
	DlpEntryTypeCustom      DlpEntryType = "custom"
	DlpEntryTypePredefined  DlpEntryType = "predefined"
	DlpEntryTypeIntegration DlpEntryType = "integration"
	DlpEntryTypeExactData   DlpEntryType = "exact_data"
	DlpEntryTypeWordList    DlpEntryType = "word_list"
)

func (r DlpEntryType) IsKnown() bool {
	switch r {
	case DlpEntryTypeCustom, DlpEntryTypePredefined, DlpEntryTypeIntegration, DlpEntryTypeExactData, DlpEntryTypeWordList:
		return true
	}
	return false
}

type Pattern struct {
	Regex string `json:"regex,required"`
	// Deprecated: deprecated
	Validation PatternValidation `json:"validation"`
	JSON       patternJSON       `json:"-"`
}

// patternJSON contains the JSON metadata for the struct [Pattern]
type patternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Pattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r patternJSON) RawJSON() string {
	return r.raw
}

type PatternValidation string

const (
	PatternValidationLuhn PatternValidation = "luhn"
)

func (r PatternValidation) IsKnown() bool {
	switch r {
	case PatternValidationLuhn:
		return true
	}
	return false
}

type PatternParam struct {
	Regex param.Field[string] `json:"regex,required"`
	// Deprecated: deprecated
	Validation param.Field[PatternValidation] `json:"validation"`
}

func (r PatternParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDlpEntryNewResponse struct {
	Result CustomEntry                    `json:"result"`
	JSON   accountDlpEntryNewResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpEntryNewResponseJSON contains the JSON metadata for the struct
// [AccountDlpEntryNewResponse]
type accountDlpEntryNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpEntryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpEntryNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpEntryGetResponse struct {
	Result DlpEntry                       `json:"result"`
	JSON   accountDlpEntryGetResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpEntryGetResponseJSON contains the JSON metadata for the struct
// [AccountDlpEntryGetResponse]
type accountDlpEntryGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpEntryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpEntryGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpEntryUpdateResponse struct {
	Result DlpEntry                          `json:"result"`
	JSON   accountDlpEntryUpdateResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpEntryUpdateResponseJSON contains the JSON metadata for the struct
// [AccountDlpEntryUpdateResponse]
type accountDlpEntryUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpEntryUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpEntryUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpEntryListResponse struct {
	Result []DlpEntry                      `json:"result"`
	JSON   accountDlpEntryListResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpEntryListResponseJSON contains the JSON metadata for the struct
// [AccountDlpEntryListResponse]
type accountDlpEntryListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpEntryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpEntryListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpEntryDeleteResponse struct {
	Result interface{}                       `json:"result,nullable"`
	JSON   accountDlpEntryDeleteResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpEntryDeleteResponseJSON contains the JSON metadata for the struct
// [AccountDlpEntryDeleteResponse]
type accountDlpEntryDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpEntryDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpEntryDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpEntryNewParams struct {
	Enabled   param.Field[bool]         `json:"enabled,required"`
	Name      param.Field[string]       `json:"name,required"`
	Pattern   param.Field[PatternParam] `json:"pattern,required"`
	ProfileID param.Field[string]       `json:"profile_id,required" format:"uuid"`
}

func (r AccountDlpEntryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDlpEntryUpdateParams struct {
	Body AccountDlpEntryUpdateParamsBodyUnion `json:"body,required"`
}

func (r AccountDlpEntryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDlpEntryUpdateParamsBody struct {
	Type    param.Field[AccountDlpEntryUpdateParamsBodyType] `json:"type,required"`
	Enabled param.Field[bool]                                `json:"enabled"`
	Name    param.Field[string]                              `json:"name"`
	Pattern param.Field[PatternParam]                        `json:"pattern"`
}

func (r AccountDlpEntryUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpEntryUpdateParamsBody) implementsAccountDlpEntryUpdateParamsBodyUnion() {}

// Satisfied by [AccountDlpEntryUpdateParamsBodyObject],
// [AccountDlpEntryUpdateParamsBodyObject],
// [AccountDlpEntryUpdateParamsBodyObject], [AccountDlpEntryUpdateParamsBody].
type AccountDlpEntryUpdateParamsBodyUnion interface {
	implementsAccountDlpEntryUpdateParamsBodyUnion()
}

type AccountDlpEntryUpdateParamsBodyObject struct {
	Name    param.Field[string]                                    `json:"name,required"`
	Pattern param.Field[PatternParam]                              `json:"pattern,required"`
	Type    param.Field[AccountDlpEntryUpdateParamsBodyObjectType] `json:"type,required"`
	Enabled param.Field[bool]                                      `json:"enabled"`
}

func (r AccountDlpEntryUpdateParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpEntryUpdateParamsBodyObject) implementsAccountDlpEntryUpdateParamsBodyUnion() {}

type AccountDlpEntryUpdateParamsBodyObjectType string

const (
	AccountDlpEntryUpdateParamsBodyObjectTypeCustom AccountDlpEntryUpdateParamsBodyObjectType = "custom"
)

func (r AccountDlpEntryUpdateParamsBodyObjectType) IsKnown() bool {
	switch r {
	case AccountDlpEntryUpdateParamsBodyObjectTypeCustom:
		return true
	}
	return false
}

type AccountDlpEntryUpdateParamsBodyType string

const (
	AccountDlpEntryUpdateParamsBodyTypeCustom      AccountDlpEntryUpdateParamsBodyType = "custom"
	AccountDlpEntryUpdateParamsBodyTypePredefined  AccountDlpEntryUpdateParamsBodyType = "predefined"
	AccountDlpEntryUpdateParamsBodyTypeIntegration AccountDlpEntryUpdateParamsBodyType = "integration"
)

func (r AccountDlpEntryUpdateParamsBodyType) IsKnown() bool {
	switch r {
	case AccountDlpEntryUpdateParamsBodyTypeCustom, AccountDlpEntryUpdateParamsBodyTypePredefined, AccountDlpEntryUpdateParamsBodyTypeIntegration:
		return true
	}
	return false
}
