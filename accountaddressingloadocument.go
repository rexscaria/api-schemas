// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAddressingLoaDocumentService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAddressingLoaDocumentService] method instead.
type AccountAddressingLoaDocumentService struct {
	Options []option.RequestOption
}

// NewAccountAddressingLoaDocumentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressingLoaDocumentService(opts ...option.RequestOption) (r *AccountAddressingLoaDocumentService) {
	r = &AccountAddressingLoaDocumentService{}
	r.Options = opts
	return
}

// Download specified LOA document under the account.
func (r *AccountAddressingLoaDocumentService) Download(ctx context.Context, accountID string, loaDocumentID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if loaDocumentID == "" {
		err = errors.New("missing required loa_document_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/loa_documents/%s/download", accountID, loaDocumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Submit LOA document (pdf format) under the account.
func (r *AccountAddressingLoaDocumentService) Upload(ctx context.Context, accountID string, body AccountAddressingLoaDocumentUploadParams, opts ...option.RequestOption) (res *AccountAddressingLoaDocumentUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/loa_documents", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountAddressingLoaDocumentUploadResponse struct {
	Result AccountAddressingLoaDocumentUploadResponseResult `json:"result"`
	JSON   accountAddressingLoaDocumentUploadResponseJSON   `json:"-"`
	AddressingAPIResponseSingle
}

// accountAddressingLoaDocumentUploadResponseJSON contains the JSON metadata for
// the struct [AccountAddressingLoaDocumentUploadResponse]
type accountAddressingLoaDocumentUploadResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingLoaDocumentUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingLoaDocumentUploadResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAddressingLoaDocumentUploadResponseResult struct {
	// Identifier for the uploaded LOA document.
	ID string `json:"id,nullable"`
	// Identifier of a Cloudflare account.
	AccountID string    `json:"account_id"`
	Created   time.Time `json:"created" format:"date-time"`
	// Name of LOA document. Max file size 10MB, and supported filetype is pdf.
	Filename string `json:"filename"`
	// File size of the uploaded LOA document.
	SizeBytes int64 `json:"size_bytes"`
	// Whether the LOA has been verified by Cloudflare staff.
	Verified bool `json:"verified"`
	// Timestamp of the moment the LOA was marked as validated.
	VerifiedAt time.Time                                            `json:"verified_at,nullable" format:"date-time"`
	JSON       accountAddressingLoaDocumentUploadResponseResultJSON `json:"-"`
}

// accountAddressingLoaDocumentUploadResponseResultJSON contains the JSON metadata
// for the struct [AccountAddressingLoaDocumentUploadResponseResult]
type accountAddressingLoaDocumentUploadResponseResultJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Created     apijson.Field
	Filename    apijson.Field
	SizeBytes   apijson.Field
	Verified    apijson.Field
	VerifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingLoaDocumentUploadResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingLoaDocumentUploadResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAddressingLoaDocumentUploadParams struct {
	// LOA document to upload.
	LoaDocument param.Field[string] `json:"loa_document,required"`
}

func (r AccountAddressingLoaDocumentUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
