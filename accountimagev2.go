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
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apiform"
	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountImageV2Service contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountImageV2Service] method instead.
type AccountImageV2Service struct {
	Options []option.RequestOption
}

// NewAccountImageV2Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountImageV2Service(opts ...option.RequestOption) (r *AccountImageV2Service) {
	r = &AccountImageV2Service{}
	r.Options = opts
	return
}

// List up to 10000 images with one request. Use the optional parameters below to
// get a specific range of images. Endpoint returns continuation_token if more
// images are present.
func (r *AccountImageV2Service) List(ctx context.Context, accountID string, query AccountImageV2ListParams, opts ...option.RequestOption) (res *AccountImageV2ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v2", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Direct uploads allow users to upload images without API keys. A common use case
// are web apps, client-side applications, or mobile devices where users upload
// content directly to Cloudflare Images. This method creates a draft record for a
// future image. It returns an upload URL and an image identifier. To verify if the
// image itself has been uploaded, send an image details request
// (accounts/:account_identifier/images/v1/:identifier), and check that the
// `draft: true` property is not present.
func (r *AccountImageV2Service) NewDirectUpload(ctx context.Context, accountID string, body AccountImageV2NewDirectUploadParams, opts ...option.RequestOption) (res *AccountImageV2NewDirectUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v2/direct_upload", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountImageV2ListResponse struct {
	Result AccountImageV2ListResponseResult `json:"result"`
	JSON   accountImageV2ListResponseJSON   `json:"-"`
	APIResponseImages
}

// accountImageV2ListResponseJSON contains the JSON metadata for the struct
// [AccountImageV2ListResponse]
type accountImageV2ListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV2ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV2ListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountImageV2ListResponseResult struct {
	// Continuation token to fetch next page. Passed as a query param when requesting
	// List V2 api endpoint.
	ContinuationToken string                               `json:"continuation_token,nullable"`
	Images            []Image                              `json:"images"`
	JSON              accountImageV2ListResponseResultJSON `json:"-"`
}

// accountImageV2ListResponseResultJSON contains the JSON metadata for the struct
// [AccountImageV2ListResponseResult]
type accountImageV2ListResponseResultJSON struct {
	ContinuationToken apijson.Field
	Images            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountImageV2ListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV2ListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountImageV2NewDirectUploadResponse struct {
	Result AccountImageV2NewDirectUploadResponseResult `json:"result"`
	JSON   accountImageV2NewDirectUploadResponseJSON   `json:"-"`
	APIResponseSingleImages
}

// accountImageV2NewDirectUploadResponseJSON contains the JSON metadata for the
// struct [AccountImageV2NewDirectUploadResponse]
type accountImageV2NewDirectUploadResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV2NewDirectUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV2NewDirectUploadResponseJSON) RawJSON() string {
	return r.raw
}

type AccountImageV2NewDirectUploadResponseResult struct {
	// Image unique identifier.
	ID string `json:"id"`
	// The URL the unauthenticated upload can be performed to using a single HTTP POST
	// (multipart/form-data) request.
	UploadURL string                                          `json:"uploadURL"`
	JSON      accountImageV2NewDirectUploadResponseResultJSON `json:"-"`
}

// accountImageV2NewDirectUploadResponseResultJSON contains the JSON metadata for
// the struct [AccountImageV2NewDirectUploadResponseResult]
type accountImageV2NewDirectUploadResponseResultJSON struct {
	ID          apijson.Field
	UploadURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV2NewDirectUploadResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV2NewDirectUploadResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountImageV2ListParams struct {
	// Continuation token for a next page. List images V2 returns continuation_token
	ContinuationToken param.Field[string] `query:"continuation_token"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Sorting order by upload time.
	SortOrder param.Field[AccountImageV2ListParamsSortOrder] `query:"sort_order"`
}

// URLQuery serializes [AccountImageV2ListParams]'s query parameters as
// `url.Values`.
func (r AccountImageV2ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sorting order by upload time.
type AccountImageV2ListParamsSortOrder string

const (
	AccountImageV2ListParamsSortOrderAsc  AccountImageV2ListParamsSortOrder = "asc"
	AccountImageV2ListParamsSortOrderDesc AccountImageV2ListParamsSortOrder = "desc"
)

func (r AccountImageV2ListParamsSortOrder) IsKnown() bool {
	switch r {
	case AccountImageV2ListParamsSortOrderAsc, AccountImageV2ListParamsSortOrderDesc:
		return true
	}
	return false
}

type AccountImageV2NewDirectUploadParams struct {
	// Optional Image Custom ID. Up to 1024 chars. Can include any number of subpaths,
	// and utf8 characters. Cannot start nor end with a / (forward slash). Cannot be a
	// UUID.
	ID param.Field[string] `json:"id"`
	// The date after which the upload will not be accepted. Minimum: Now + 2 minutes.
	// Maximum: Now + 6 hours.
	Expiry param.Field[time.Time] `json:"expiry" format:"date-time"`
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record, for managing images.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Indicates whether the image requires a signature token to be accessed.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
}

func (r AccountImageV2NewDirectUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
