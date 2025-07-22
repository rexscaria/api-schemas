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
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAIFinetuneService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIFinetuneService] method instead.
type AccountAIFinetuneService struct {
	Options []option.RequestOption
}

// NewAccountAIFinetuneService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIFinetuneService(opts ...option.RequestOption) (r *AccountAIFinetuneService) {
	r = &AccountAIFinetuneService{}
	r.Options = opts
	return
}

// Create a new Finetune
func (r *AccountAIFinetuneService) New(ctx context.Context, accountID string, body AccountAIFinetuneNewParams, opts ...option.RequestOption) (res *AccountAIFinetuneNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/finetunes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Finetunes
func (r *AccountAIFinetuneService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAIFinetuneListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/finetunes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Public Finetunes
func (r *AccountAIFinetuneService) ListPublic(ctx context.Context, accountID string, query AccountAIFinetuneListPublicParams, opts ...option.RequestOption) (res *AccountAIFinetuneListPublicResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/finetunes/public", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Upload a Finetune Asset
func (r *AccountAIFinetuneService) UploadAsset(ctx context.Context, accountID string, finetuneID string, body AccountAIFinetuneUploadAssetParams, opts ...option.RequestOption) (res *AccountAIFinetuneUploadAssetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if finetuneID == "" {
		err = errors.New("missing required finetune_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/finetunes/%s/finetune-assets", accountID, finetuneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountAIFinetuneNewResponse struct {
	Result  AccountAIFinetuneNewResponseResult `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    accountAIFinetuneNewResponseJSON   `json:"-"`
}

// accountAIFinetuneNewResponseJSON contains the JSON metadata for the struct
// [AccountAIFinetuneNewResponse]
type accountAIFinetuneNewResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIFinetuneNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIFinetuneNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIFinetuneNewResponseResult struct {
	ID          string                                 `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                              `json:"created_at,required" format:"date-time"`
	Model       string                                 `json:"model,required"`
	ModifiedAt  time.Time                              `json:"modified_at,required" format:"date-time"`
	Name        string                                 `json:"name,required"`
	Public      bool                                   `json:"public,required"`
	Description string                                 `json:"description"`
	JSON        accountAIFinetuneNewResponseResultJSON `json:"-"`
}

// accountAIFinetuneNewResponseResultJSON contains the JSON metadata for the struct
// [AccountAIFinetuneNewResponseResult]
type accountAIFinetuneNewResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Model       apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Public      apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIFinetuneNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIFinetuneNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIFinetuneListResponse struct {
	Result  AccountAIFinetuneListResponseResult `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    accountAIFinetuneListResponseJSON   `json:"-"`
}

// accountAIFinetuneListResponseJSON contains the JSON metadata for the struct
// [AccountAIFinetuneListResponse]
type accountAIFinetuneListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIFinetuneListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIFinetuneListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIFinetuneListResponseResult struct {
	ID          string                                  `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                               `json:"created_at,required" format:"date-time"`
	Model       string                                  `json:"model,required"`
	ModifiedAt  time.Time                               `json:"modified_at,required" format:"date-time"`
	Name        string                                  `json:"name,required"`
	Description string                                  `json:"description"`
	JSON        accountAIFinetuneListResponseResultJSON `json:"-"`
}

// accountAIFinetuneListResponseResultJSON contains the JSON metadata for the
// struct [AccountAIFinetuneListResponseResult]
type accountAIFinetuneListResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Model       apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIFinetuneListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIFinetuneListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIFinetuneListPublicResponse struct {
	Result  []AccountAIFinetuneListPublicResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    accountAIFinetuneListPublicResponseJSON     `json:"-"`
}

// accountAIFinetuneListPublicResponseJSON contains the JSON metadata for the
// struct [AccountAIFinetuneListPublicResponse]
type accountAIFinetuneListPublicResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIFinetuneListPublicResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIFinetuneListPublicResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIFinetuneListPublicResponseResult struct {
	ID          string                                        `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                                     `json:"created_at,required" format:"date-time"`
	Model       string                                        `json:"model,required"`
	ModifiedAt  time.Time                                     `json:"modified_at,required" format:"date-time"`
	Name        string                                        `json:"name,required"`
	Public      bool                                          `json:"public,required"`
	Description string                                        `json:"description"`
	JSON        accountAIFinetuneListPublicResponseResultJSON `json:"-"`
}

// accountAIFinetuneListPublicResponseResultJSON contains the JSON metadata for the
// struct [AccountAIFinetuneListPublicResponseResult]
type accountAIFinetuneListPublicResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Model       apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Public      apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIFinetuneListPublicResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIFinetuneListPublicResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIFinetuneUploadAssetResponse struct {
	Result  AccountAIFinetuneUploadAssetResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    accountAIFinetuneUploadAssetResponseJSON   `json:"-"`
}

// accountAIFinetuneUploadAssetResponseJSON contains the JSON metadata for the
// struct [AccountAIFinetuneUploadAssetResponse]
type accountAIFinetuneUploadAssetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIFinetuneUploadAssetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIFinetuneUploadAssetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIFinetuneUploadAssetResponseResult struct {
	ID         string                                         `json:"id,required" format:"uuid"`
	BucketName string                                         `json:"bucket_name,required"`
	CreatedAt  time.Time                                      `json:"created_at,required" format:"date-time"`
	FileName   string                                         `json:"file_name,required"`
	FinetuneID string                                         `json:"finetune_id,required"`
	ModifiedAt time.Time                                      `json:"modified_at,required" format:"date-time"`
	JSON       accountAIFinetuneUploadAssetResponseResultJSON `json:"-"`
}

// accountAIFinetuneUploadAssetResponseResultJSON contains the JSON metadata for
// the struct [AccountAIFinetuneUploadAssetResponseResult]
type accountAIFinetuneUploadAssetResponseResultJSON struct {
	ID          apijson.Field
	BucketName  apijson.Field
	CreatedAt   apijson.Field
	FileName    apijson.Field
	FinetuneID  apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIFinetuneUploadAssetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIFinetuneUploadAssetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIFinetuneNewParams struct {
	Model       param.Field[string] `json:"model,required"`
	Name        param.Field[string] `json:"name,required"`
	Description param.Field[string] `json:"description"`
	Public      param.Field[bool]   `json:"public"`
}

func (r AccountAIFinetuneNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIFinetuneListPublicParams struct {
	// Pagination Limit
	Limit param.Field[float64] `query:"limit"`
	// Pagination Offset
	Offset param.Field[float64] `query:"offset"`
	// Order By Column Name
	OrderBy param.Field[string] `query:"orderBy"`
}

// URLQuery serializes [AccountAIFinetuneListPublicParams]'s query parameters as
// `url.Values`.
func (r AccountAIFinetuneListPublicParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIFinetuneUploadAssetParams struct {
	File     param.Field[io.Reader] `json:"file" format:"binary"`
	FileName param.Field[string]    `json:"file_name"`
}

func (r AccountAIFinetuneUploadAssetParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
