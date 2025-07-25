// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountDlpDatasetUploadService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDlpDatasetUploadService] method instead.
type AccountDlpDatasetUploadService struct {
	Options []option.RequestOption
}

// NewAccountDlpDatasetUploadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDlpDatasetUploadService(opts ...option.RequestOption) (r *AccountDlpDatasetUploadService) {
	r = &AccountDlpDatasetUploadService{}
	r.Options = opts
	return
}

// Prepare to upload a new version of a dataset
func (r *AccountDlpDatasetUploadService) Prepare(ctx context.Context, accountID string, datasetID string, opts ...option.RequestOption) (res *AccountDlpDatasetUploadPrepareResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This is used for single-column EDMv1 and Custom Word Lists. The EDM format can
// only be created in the Cloudflare dashboard. For other clients, this operation
// can only be used for non-secret Custom Word Lists. The body must be a UTF-8
// encoded, newline (NL or CRNL) separated list of words to be matched.
func (r *AccountDlpDatasetUploadService) Version(ctx context.Context, accountID string, datasetID string, version int64, body AccountDlpDatasetUploadVersionParams, opts ...option.RequestOption) (res *AccountDlpDatasetUploadVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload/%v", accountID, datasetID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountDlpDatasetUploadPrepareResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpDatasetUploadPrepareResponseSuccess `json:"success,required"`
	Result  AccountDlpDatasetUploadPrepareResponseResult  `json:"result"`
	JSON    accountDlpDatasetUploadPrepareResponseJSON    `json:"-"`
}

// accountDlpDatasetUploadPrepareResponseJSON contains the JSON metadata for the
// struct [AccountDlpDatasetUploadPrepareResponse]
type accountDlpDatasetUploadPrepareResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpDatasetUploadPrepareResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpDatasetUploadPrepareResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDlpDatasetUploadPrepareResponseSuccess bool

const (
	AccountDlpDatasetUploadPrepareResponseSuccessTrue AccountDlpDatasetUploadPrepareResponseSuccess = true
)

func (r AccountDlpDatasetUploadPrepareResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpDatasetUploadPrepareResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpDatasetUploadPrepareResponseResult struct {
	EncodingVersion int64                                            `json:"encoding_version,required"`
	MaxCells        int64                                            `json:"max_cells,required"`
	Version         int64                                            `json:"version,required"`
	CaseSensitive   bool                                             `json:"case_sensitive"`
	Columns         []DatasetColumn                                  `json:"columns"`
	Secret          string                                           `json:"secret" format:"password"`
	JSON            accountDlpDatasetUploadPrepareResponseResultJSON `json:"-"`
}

// accountDlpDatasetUploadPrepareResponseResultJSON contains the JSON metadata for
// the struct [AccountDlpDatasetUploadPrepareResponseResult]
type accountDlpDatasetUploadPrepareResponseResultJSON struct {
	EncodingVersion apijson.Field
	MaxCells        apijson.Field
	Version         apijson.Field
	CaseSensitive   apijson.Field
	Columns         apijson.Field
	Secret          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDlpDatasetUploadPrepareResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpDatasetUploadPrepareResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDlpDatasetUploadVersionResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpDatasetUploadVersionResponseSuccess `json:"success,required"`
	Result  Dataset                                       `json:"result"`
	JSON    accountDlpDatasetUploadVersionResponseJSON    `json:"-"`
}

// accountDlpDatasetUploadVersionResponseJSON contains the JSON metadata for the
// struct [AccountDlpDatasetUploadVersionResponse]
type accountDlpDatasetUploadVersionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpDatasetUploadVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpDatasetUploadVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDlpDatasetUploadVersionResponseSuccess bool

const (
	AccountDlpDatasetUploadVersionResponseSuccessTrue AccountDlpDatasetUploadVersionResponseSuccess = true
)

func (r AccountDlpDatasetUploadVersionResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpDatasetUploadVersionResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpDatasetUploadVersionParams struct {
	Body io.Reader `json:"body,required" format:"binary"`
}

func (r AccountDlpDatasetUploadVersionParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
