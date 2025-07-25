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
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountDlpDatasetVersionService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDlpDatasetVersionService] method instead.
type AccountDlpDatasetVersionService struct {
	Options []option.RequestOption
}

// NewAccountDlpDatasetVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDlpDatasetVersionService(opts ...option.RequestOption) (r *AccountDlpDatasetVersionService) {
	r = &AccountDlpDatasetVersionService{}
	r.Options = opts
	return
}

// This is used for multi-column EDMv2 datasets. The EDMv2 format can only be
// created in the Cloudflare dashboard. The columns in the response appear in the
// same order as in the request.
func (r *AccountDlpDatasetVersionService) SetColumnInfo(ctx context.Context, accountID string, datasetID string, version int64, body AccountDlpDatasetVersionSetColumnInfoParams, opts ...option.RequestOption) (res *AccountDlpDatasetVersionSetColumnInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/versions/%v", accountID, datasetID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This is used for multi-column EDMv2 datasets. The EDMv2 format can only be
// created in the Cloudflare dashboard.
func (r *AccountDlpDatasetVersionService) UploadEntry(ctx context.Context, accountID string, datasetID string, version int64, entryID string, body AccountDlpDatasetVersionUploadEntryParams, opts ...option.RequestOption) (res *AccountDlpDatasetVersionUploadEntryResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/versions/%v/entries/%s", accountID, datasetID, version, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DatasetColumn struct {
	EntryID      string            `json:"entry_id,required" format:"uuid"`
	HeaderName   string            `json:"header_name,required"`
	NumCells     int64             `json:"num_cells,required"`
	UploadStatus UploadStatus      `json:"upload_status,required"`
	JSON         datasetColumnJSON `json:"-"`
}

// datasetColumnJSON contains the JSON metadata for the struct [DatasetColumn]
type datasetColumnJSON struct {
	EntryID      apijson.Field
	HeaderName   apijson.Field
	NumCells     apijson.Field
	UploadStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DatasetColumn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetColumnJSON) RawJSON() string {
	return r.raw
}

type AccountDlpDatasetVersionSetColumnInfoResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpDatasetVersionSetColumnInfoResponseSuccess `json:"success,required"`
	Result  []DatasetColumn                                      `json:"result"`
	JSON    accountDlpDatasetVersionSetColumnInfoResponseJSON    `json:"-"`
}

// accountDlpDatasetVersionSetColumnInfoResponseJSON contains the JSON metadata for
// the struct [AccountDlpDatasetVersionSetColumnInfoResponse]
type accountDlpDatasetVersionSetColumnInfoResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpDatasetVersionSetColumnInfoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpDatasetVersionSetColumnInfoResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDlpDatasetVersionSetColumnInfoResponseSuccess bool

const (
	AccountDlpDatasetVersionSetColumnInfoResponseSuccessTrue AccountDlpDatasetVersionSetColumnInfoResponseSuccess = true
)

func (r AccountDlpDatasetVersionSetColumnInfoResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpDatasetVersionSetColumnInfoResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpDatasetVersionUploadEntryResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpDatasetVersionUploadEntryResponseSuccess `json:"success,required"`
	Result  DatasetColumn                                      `json:"result"`
	JSON    accountDlpDatasetVersionUploadEntryResponseJSON    `json:"-"`
}

// accountDlpDatasetVersionUploadEntryResponseJSON contains the JSON metadata for
// the struct [AccountDlpDatasetVersionUploadEntryResponse]
type accountDlpDatasetVersionUploadEntryResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpDatasetVersionUploadEntryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpDatasetVersionUploadEntryResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDlpDatasetVersionUploadEntryResponseSuccess bool

const (
	AccountDlpDatasetVersionUploadEntryResponseSuccessTrue AccountDlpDatasetVersionUploadEntryResponseSuccess = true
)

func (r AccountDlpDatasetVersionUploadEntryResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpDatasetVersionUploadEntryResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpDatasetVersionSetColumnInfoParams struct {
	Body []AccountDlpDatasetVersionSetColumnInfoParamsBodyUnion `json:"body,required"`
}

func (r AccountDlpDatasetVersionSetColumnInfoParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDlpDatasetVersionSetColumnInfoParamsBody struct {
	EntryID    param.Field[string] `json:"entry_id" format:"uuid"`
	EntryName  param.Field[string] `json:"entry_name"`
	HeaderName param.Field[string] `json:"header_name"`
	NumCells   param.Field[int64]  `json:"num_cells"`
}

func (r AccountDlpDatasetVersionSetColumnInfoParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpDatasetVersionSetColumnInfoParamsBody) implementsAccountDlpDatasetVersionSetColumnInfoParamsBodyUnion() {
}

// Satisfied by [AccountDlpDatasetVersionSetColumnInfoParamsBodyExistingColumn],
// [AccountDlpDatasetVersionSetColumnInfoParamsBodyNewColumn],
// [AccountDlpDatasetVersionSetColumnInfoParamsBody].
type AccountDlpDatasetVersionSetColumnInfoParamsBodyUnion interface {
	implementsAccountDlpDatasetVersionSetColumnInfoParamsBodyUnion()
}

type AccountDlpDatasetVersionSetColumnInfoParamsBodyExistingColumn struct {
	EntryID    param.Field[string] `json:"entry_id,required" format:"uuid"`
	HeaderName param.Field[string] `json:"header_name"`
	NumCells   param.Field[int64]  `json:"num_cells"`
}

func (r AccountDlpDatasetVersionSetColumnInfoParamsBodyExistingColumn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpDatasetVersionSetColumnInfoParamsBodyExistingColumn) implementsAccountDlpDatasetVersionSetColumnInfoParamsBodyUnion() {
}

type AccountDlpDatasetVersionSetColumnInfoParamsBodyNewColumn struct {
	EntryName  param.Field[string] `json:"entry_name,required"`
	HeaderName param.Field[string] `json:"header_name"`
	NumCells   param.Field[int64]  `json:"num_cells"`
}

func (r AccountDlpDatasetVersionSetColumnInfoParamsBodyNewColumn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpDatasetVersionSetColumnInfoParamsBodyNewColumn) implementsAccountDlpDatasetVersionSetColumnInfoParamsBodyUnion() {
}

type AccountDlpDatasetVersionUploadEntryParams struct {
	Body io.Reader `json:"body,required" format:"binary"`
}

func (r AccountDlpDatasetVersionUploadEntryParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
