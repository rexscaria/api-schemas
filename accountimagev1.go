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
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// AccountImageV1Service contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountImageV1Service] method instead.
type AccountImageV1Service struct {
	Options  []option.RequestOption
	Keys     *AccountImageV1KeyService
	Variants *AccountImageV1VariantService
}

// NewAccountImageV1Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountImageV1Service(opts ...option.RequestOption) (r *AccountImageV1Service) {
	r = &AccountImageV1Service{}
	r.Options = opts
	r.Keys = NewAccountImageV1KeyService(opts...)
	r.Variants = NewAccountImageV1VariantService(opts...)
	return
}

// Fetch details for a single image.
func (r *AccountImageV1Service) Get(ctx context.Context, accountID string, imageID string, opts ...option.RequestOption) (res *ImageResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/%s", accountID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update image access control. On access control change, all copies of the image
// are purged from cache.
func (r *AccountImageV1Service) Update(ctx context.Context, accountID string, imageID string, body AccountImageV1UpdateParams, opts ...option.RequestOption) (res *ImageResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/%s", accountID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List up to 100 images with one request. Use the optional parameters below to get
// a specific range of images.
//
// Deprecated: deprecated
func (r *AccountImageV1Service) List(ctx context.Context, accountID string, query AccountImageV1ListParams, opts ...option.RequestOption) (res *AccountImageV1ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an image on Cloudflare Images. On success, all copies of the image are
// deleted and purged from cache.
func (r *AccountImageV1Service) Delete(ctx context.Context, accountID string, imageID string, body AccountImageV1DeleteParams, opts ...option.RequestOption) (res *DeletedImagesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/%s", accountID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Fetch base image. For most images this will be the originally uploaded file. For
// larger images it can be a near-lossless version of the original.
func (r *AccountImageV1Service) FetchBase(ctx context.Context, accountID string, imageID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/*")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/%s/blob", accountID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetch usage statistics details for Cloudflare Images.
func (r *AccountImageV1Service) Stats(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountImageV1StatsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/stats", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload an image with up to 10 Megabytes using a single HTTP POST
// (multipart/form-data) request. An image can be uploaded by sending an image file
// or passing an accessible to an API url.
func (r *AccountImageV1Service) Upload(ctx context.Context, accountID string, body AccountImageV1UploadParams, opts ...option.RequestOption) (res *ImageResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type APIResponseImages struct {
	Errors   []MessagesImageItem          `json:"errors,required"`
	Messages []MessagesImageItem          `json:"messages,required"`
	Result   APIResponseImagesResultUnion `json:"result,required"`
	// Whether the API call was successful
	Success APIResponseImagesSuccess `json:"success,required"`
	JSON    apiResponseImagesJSON    `json:"-"`
}

// apiResponseImagesJSON contains the JSON metadata for the struct
// [APIResponseImages]
type apiResponseImagesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseImages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseImagesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [APIResponseImagesResultArray] or [shared.UnionString].
type APIResponseImagesResultUnion interface {
	ImplementsAPIResponseImagesResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIResponseImagesResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIResponseImagesResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type APIResponseImagesResultArray []interface{}

func (r APIResponseImagesResultArray) ImplementsAPIResponseImagesResultUnion() {}

// Whether the API call was successful
type APIResponseImagesSuccess bool

const (
	APIResponseImagesSuccessTrue APIResponseImagesSuccess = true
)

func (r APIResponseImagesSuccess) IsKnown() bool {
	switch r {
	case APIResponseImagesSuccessTrue:
		return true
	}
	return false
}

type APIResponseSingleImages struct {
	Result interface{}                 `json:"result"`
	JSON   apiResponseSingleImagesJSON `json:"-"`
	APIResponseImages
}

// apiResponseSingleImagesJSON contains the JSON metadata for the struct
// [APIResponseSingleImages]
type apiResponseSingleImagesJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleImages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleImagesJSON) RawJSON() string {
	return r.raw
}

type DeletedImagesResponse struct {
	Result interface{}               `json:"result"`
	JSON   deletedImagesResponseJSON `json:"-"`
	APIResponseSingleImages
}

// deletedImagesResponseJSON contains the JSON metadata for the struct
// [DeletedImagesResponse]
type deletedImagesResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeletedImagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deletedImagesResponseJSON) RawJSON() string {
	return r.raw
}

type Image struct {
	// Image unique identifier.
	ID string `json:"id"`
	// Image file name.
	Filename string `json:"filename"`
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record for managing images. Metadata must not exceed 1024 bytes.
	Meta interface{} `json:"meta"`
	// Indicates whether the image can be a accessed only using it's UID. If set to
	// true, a signed token needs to be generated with a signing key to view the image.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// When the media item was uploaded.
	Uploaded time.Time `json:"uploaded" format:"date-time"`
	// Object specifying available variants for an image.
	Variants []string  `json:"variants" format:"uri"`
	JSON     imageJSON `json:"-"`
}

// imageJSON contains the JSON metadata for the struct [Image]
type imageJSON struct {
	ID                apijson.Field
	Filename          apijson.Field
	Meta              apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Image) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageJSON) RawJSON() string {
	return r.raw
}

type ImageResponseSingle struct {
	Result Image                   `json:"result"`
	JSON   imageResponseSingleJSON `json:"-"`
	APIResponseSingleImages
}

// imageResponseSingleJSON contains the JSON metadata for the struct
// [ImageResponseSingle]
type imageResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageResponseSingleJSON) RawJSON() string {
	return r.raw
}

type MessagesImageItem struct {
	Code    int64                 `json:"code,required"`
	Message string                `json:"message,required"`
	JSON    messagesImageItemJSON `json:"-"`
}

// messagesImageItemJSON contains the JSON metadata for the struct
// [MessagesImageItem]
type messagesImageItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesImageItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesImageItemJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1ListResponse struct {
	Result AccountImageV1ListResponseResult `json:"result"`
	JSON   accountImageV1ListResponseJSON   `json:"-"`
	APIResponseImages
}

// accountImageV1ListResponseJSON contains the JSON metadata for the struct
// [AccountImageV1ListResponse]
type accountImageV1ListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1ListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1ListResponseResult struct {
	Images []Image                              `json:"images"`
	JSON   accountImageV1ListResponseResultJSON `json:"-"`
}

// accountImageV1ListResponseResultJSON contains the JSON metadata for the struct
// [AccountImageV1ListResponseResult]
type accountImageV1ListResponseResultJSON struct {
	Images      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1ListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1ListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1StatsResponse struct {
	Result AccountImageV1StatsResponseResult `json:"result"`
	JSON   accountImageV1StatsResponseJSON   `json:"-"`
	APIResponseSingleImages
}

// accountImageV1StatsResponseJSON contains the JSON metadata for the struct
// [AccountImageV1StatsResponse]
type accountImageV1StatsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1StatsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1StatsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1StatsResponseResult struct {
	Count AccountImageV1StatsResponseResultCount `json:"count"`
	JSON  accountImageV1StatsResponseResultJSON  `json:"-"`
}

// accountImageV1StatsResponseResultJSON contains the JSON metadata for the struct
// [AccountImageV1StatsResponseResult]
type accountImageV1StatsResponseResultJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1StatsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1StatsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1StatsResponseResultCount struct {
	// Cloudflare Images allowed usage.
	Allowed float64 `json:"allowed"`
	// Cloudflare Images current usage.
	Current float64                                    `json:"current"`
	JSON    accountImageV1StatsResponseResultCountJSON `json:"-"`
}

// accountImageV1StatsResponseResultCountJSON contains the JSON metadata for the
// struct [AccountImageV1StatsResponseResultCount]
type accountImageV1StatsResponseResultCountJSON struct {
	Allowed     apijson.Field
	Current     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1StatsResponseResultCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1StatsResponseResultCountJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1UpdateParams struct {
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record for managing images. No change if not specified.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Indicates whether the image can be accessed using only its UID. If set to
	// `true`, a signed token needs to be generated with a signing key to view the
	// image. Returns a new UID on a change. No change if not specified.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
}

func (r AccountImageV1UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountImageV1ListParams struct {
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountImageV1ListParams]'s query parameters as
// `url.Values`.
func (r AccountImageV1ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountImageV1DeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountImageV1DeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountImageV1UploadParams struct {
	// An image binary data. Only needed when type is uploading a file.
	File param.Field[interface{}] `json:"file"`
	// User modifiable key-value store. Can use used for keeping references to another
	// system of record for managing images.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Indicates whether the image requires a signature token for the access.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// A URL to fetch an image from origin. Only needed when type is uploading from a
	// URL.
	URL param.Field[string] `json:"url"`
}

func (r AccountImageV1UploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
