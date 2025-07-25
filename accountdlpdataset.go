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

// AccountDlpDatasetService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDlpDatasetService] method instead.
type AccountDlpDatasetService struct {
	Options  []option.RequestOption
	Upload   *AccountDlpDatasetUploadService
	Versions *AccountDlpDatasetVersionService
}

// NewAccountDlpDatasetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDlpDatasetService(opts ...option.RequestOption) (r *AccountDlpDatasetService) {
	r = &AccountDlpDatasetService{}
	r.Options = opts
	r.Upload = NewAccountDlpDatasetUploadService(opts...)
	r.Versions = NewAccountDlpDatasetVersionService(opts...)
	return
}

// Create a new dataset
func (r *AccountDlpDatasetService) New(ctx context.Context, accountID string, body AccountDlpDatasetNewParams, opts ...option.RequestOption) (res *AccountDlpDatasetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a specific dataset
func (r *AccountDlpDatasetService) Get(ctx context.Context, accountID string, datasetID string, opts ...option.RequestOption) (res *AccountDlpDatasetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update details about a dataset
func (r *AccountDlpDatasetService) Update(ctx context.Context, accountID string, datasetID string, body AccountDlpDatasetUpdateParams, opts ...option.RequestOption) (res *AccountDlpDatasetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetch all datasets
func (r *AccountDlpDatasetService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountDlpDatasetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This deletes all versions of the dataset.
func (r *AccountDlpDatasetService) Delete(ctx context.Context, accountID string, datasetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Dataset struct {
	ID              string          `json:"id,required" format:"uuid"`
	Columns         []DatasetColumn `json:"columns,required"`
	CreatedAt       time.Time       `json:"created_at,required" format:"date-time"`
	EncodingVersion int64           `json:"encoding_version,required"`
	Name            string          `json:"name,required"`
	NumCells        int64           `json:"num_cells,required"`
	Secret          bool            `json:"secret,required"`
	Status          UploadStatus    `json:"status,required"`
	// When the dataset was last updated.
	//
	// This includes name or description changes as well as uploads.
	UpdatedAt     time.Time       `json:"updated_at,required" format:"date-time"`
	Uploads       []DatasetUpload `json:"uploads,required"`
	CaseSensitive bool            `json:"case_sensitive"`
	// The description of the dataset.
	Description string      `json:"description,nullable"`
	JSON        datasetJSON `json:"-"`
}

// datasetJSON contains the JSON metadata for the struct [Dataset]
type datasetJSON struct {
	ID              apijson.Field
	Columns         apijson.Field
	CreatedAt       apijson.Field
	EncodingVersion apijson.Field
	Name            apijson.Field
	NumCells        apijson.Field
	Secret          apijson.Field
	Status          apijson.Field
	UpdatedAt       apijson.Field
	Uploads         apijson.Field
	CaseSensitive   apijson.Field
	Description     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Dataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetJSON) RawJSON() string {
	return r.raw
}

type DatasetUpload struct {
	NumCells int64             `json:"num_cells,required"`
	Status   UploadStatus      `json:"status,required"`
	Version  int64             `json:"version,required"`
	JSON     datasetUploadJSON `json:"-"`
}

// datasetUploadJSON contains the JSON metadata for the struct [DatasetUpload]
type datasetUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetUploadJSON) RawJSON() string {
	return r.raw
}

type MessagesDlpItems struct {
	Code             int64                  `json:"code,required"`
	Message          string                 `json:"message,required"`
	DocumentationURL string                 `json:"documentation_url"`
	Source           MessagesDlpItemsSource `json:"source"`
	JSON             messagesDlpItemsJSON   `json:"-"`
}

// messagesDlpItemsJSON contains the JSON metadata for the struct
// [MessagesDlpItems]
type messagesDlpItemsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessagesDlpItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesDlpItemsJSON) RawJSON() string {
	return r.raw
}

type MessagesDlpItemsSource struct {
	Pointer string                     `json:"pointer"`
	JSON    messagesDlpItemsSourceJSON `json:"-"`
}

// messagesDlpItemsSourceJSON contains the JSON metadata for the struct
// [MessagesDlpItemsSource]
type messagesDlpItemsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesDlpItemsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesDlpItemsSourceJSON) RawJSON() string {
	return r.raw
}

type UploadStatus string

const (
	UploadStatusEmpty      UploadStatus = "empty"
	UploadStatusUploading  UploadStatus = "uploading"
	UploadStatusPending    UploadStatus = "pending"
	UploadStatusProcessing UploadStatus = "processing"
	UploadStatusFailed     UploadStatus = "failed"
	UploadStatusComplete   UploadStatus = "complete"
)

func (r UploadStatus) IsKnown() bool {
	switch r {
	case UploadStatusEmpty, UploadStatusUploading, UploadStatusPending, UploadStatusProcessing, UploadStatusFailed, UploadStatusComplete:
		return true
	}
	return false
}

type AccountDlpDatasetNewResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpDatasetNewResponseSuccess `json:"success,required"`
	Result  AccountDlpDatasetNewResponseResult  `json:"result"`
	JSON    accountDlpDatasetNewResponseJSON    `json:"-"`
}

// accountDlpDatasetNewResponseJSON contains the JSON metadata for the struct
// [AccountDlpDatasetNewResponse]
type accountDlpDatasetNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpDatasetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpDatasetNewResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDlpDatasetNewResponseSuccess bool

const (
	AccountDlpDatasetNewResponseSuccessTrue AccountDlpDatasetNewResponseSuccess = true
)

func (r AccountDlpDatasetNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpDatasetNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpDatasetNewResponseResult struct {
	Dataset Dataset `json:"dataset,required"`
	// Encoding version to use for dataset.
	EncodingVersion int64 `json:"encoding_version,required"`
	MaxCells        int64 `json:"max_cells,required"`
	// The version to use when uploading the dataset.
	Version int64 `json:"version,required"`
	// The secret to use for Exact Data Match datasets. This is not present in Custom
	// Wordlists.
	Secret string                                 `json:"secret" format:"password"`
	JSON   accountDlpDatasetNewResponseResultJSON `json:"-"`
}

// accountDlpDatasetNewResponseResultJSON contains the JSON metadata for the struct
// [AccountDlpDatasetNewResponseResult]
type accountDlpDatasetNewResponseResultJSON struct {
	Dataset         apijson.Field
	EncodingVersion apijson.Field
	MaxCells        apijson.Field
	Version         apijson.Field
	Secret          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDlpDatasetNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpDatasetNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDlpDatasetGetResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpDatasetGetResponseSuccess `json:"success,required"`
	Result  Dataset                             `json:"result"`
	JSON    accountDlpDatasetGetResponseJSON    `json:"-"`
}

// accountDlpDatasetGetResponseJSON contains the JSON metadata for the struct
// [AccountDlpDatasetGetResponse]
type accountDlpDatasetGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpDatasetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpDatasetGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDlpDatasetGetResponseSuccess bool

const (
	AccountDlpDatasetGetResponseSuccessTrue AccountDlpDatasetGetResponseSuccess = true
)

func (r AccountDlpDatasetGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpDatasetGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpDatasetUpdateResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpDatasetUpdateResponseSuccess `json:"success,required"`
	Result  Dataset                                `json:"result"`
	JSON    accountDlpDatasetUpdateResponseJSON    `json:"-"`
}

// accountDlpDatasetUpdateResponseJSON contains the JSON metadata for the struct
// [AccountDlpDatasetUpdateResponse]
type accountDlpDatasetUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpDatasetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpDatasetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDlpDatasetUpdateResponseSuccess bool

const (
	AccountDlpDatasetUpdateResponseSuccessTrue AccountDlpDatasetUpdateResponseSuccess = true
)

func (r AccountDlpDatasetUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpDatasetUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpDatasetListResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpDatasetListResponseSuccess `json:"success,required"`
	Result  []Dataset                            `json:"result"`
	JSON    accountDlpDatasetListResponseJSON    `json:"-"`
}

// accountDlpDatasetListResponseJSON contains the JSON metadata for the struct
// [AccountDlpDatasetListResponse]
type accountDlpDatasetListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpDatasetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpDatasetListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDlpDatasetListResponseSuccess bool

const (
	AccountDlpDatasetListResponseSuccessTrue AccountDlpDatasetListResponseSuccess = true
)

func (r AccountDlpDatasetListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpDatasetListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpDatasetNewParams struct {
	Name param.Field[string] `json:"name,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if `secret` is true or undefined
	CaseSensitive param.Field[bool] `json:"case_sensitive"`
	// The description of the dataset.
	Description param.Field[string] `json:"description"`
	// Dataset encoding version
	//
	// Non-secret custom word lists with no header are always version 1. Secret EDM
	// lists with no header are version 1. Multicolumn CSV with headers are version 2.
	// Omitting this field provides the default value 0, which is interpreted the same
	// as 1.
	EncodingVersion param.Field[int64] `json:"encoding_version"`
	// Generate a secret dataset.
	//
	// If true, the response will include a secret to use with the EDM encoder. If
	// false, the response has no secret and the dataset is uploaded in plaintext.
	Secret param.Field[bool] `json:"secret"`
}

func (r AccountDlpDatasetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDlpDatasetUpdateParams struct {
	// Determines if the words should be matched in a case-sensitive manner.
	//
	// Only required for custom word lists.
	CaseSensitive param.Field[bool] `json:"case_sensitive"`
	// The description of the dataset.
	Description param.Field[string] `json:"description"`
	// The name of the dataset, must be unique.
	Name param.Field[string] `json:"name"`
}

func (r AccountDlpDatasetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
