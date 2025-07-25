// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountWorkerAssetService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerAssetService] method instead.
type AccountWorkerAssetService struct {
	Options []option.RequestOption
}

// NewAccountWorkerAssetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountWorkerAssetService(opts ...option.RequestOption) (r *AccountWorkerAssetService) {
	r = &AccountWorkerAssetService{}
	r.Options = opts
	return
}

// Upload assets ahead of creating a Worker version. To learn more about the direct
// uploads of assets, see
// https://developers.cloudflare.com/workers/static-assets/direct-upload/.
func (r *AccountWorkerAssetService) Upload(ctx context.Context, accountID string, params AccountWorkerAssetUploadParams, opts ...option.RequestOption) (res *AccountWorkerAssetUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/assets/upload", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountWorkerAssetUploadResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountWorkerAssetUploadResponseSuccess `json:"success,required"`
	Result  AccountWorkerAssetUploadResponseResult  `json:"result"`
	JSON    accountWorkerAssetUploadResponseJSON    `json:"-"`
}

// accountWorkerAssetUploadResponseJSON contains the JSON metadata for the struct
// [AccountWorkerAssetUploadResponse]
type accountWorkerAssetUploadResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerAssetUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerAssetUploadResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerAssetUploadResponseSuccess bool

const (
	AccountWorkerAssetUploadResponseSuccessTrue AccountWorkerAssetUploadResponseSuccess = true
)

func (r AccountWorkerAssetUploadResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerAssetUploadResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerAssetUploadResponseResult struct {
	// A "completion" JWT which can be redeemed when creating a Worker version.
	Jwt  string                                     `json:"jwt"`
	JSON accountWorkerAssetUploadResponseResultJSON `json:"-"`
}

// accountWorkerAssetUploadResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkerAssetUploadResponseResult]
type accountWorkerAssetUploadResponseResultJSON struct {
	Jwt         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerAssetUploadResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerAssetUploadResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerAssetUploadParams struct {
	// Whether the file contents are base64-encoded. Must be `true`.
	Base64 param.Field[AccountWorkerAssetUploadParamsBase64] `query:"base64,required"`
	Body   map[string]string                                 `json:"body,required"`
}

func (r AccountWorkerAssetUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [AccountWorkerAssetUploadParams]'s query parameters as
// `url.Values`.
func (r AccountWorkerAssetUploadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether the file contents are base64-encoded. Must be `true`.
type AccountWorkerAssetUploadParamsBase64 bool

const (
	AccountWorkerAssetUploadParamsBase64True AccountWorkerAssetUploadParamsBase64 = true
)

func (r AccountWorkerAssetUploadParamsBase64) IsKnown() bool {
	switch r {
	case AccountWorkerAssetUploadParamsBase64True:
		return true
	}
	return false
}
