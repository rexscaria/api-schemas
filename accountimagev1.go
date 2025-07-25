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
func (r *AccountImageV1Service) Delete(ctx context.Context, accountID string, imageID string, opts ...option.RequestOption) (res *DeletedImagesResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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

type DeletedImagesResponse struct {
	Errors   []DeletedImagesResponseError   `json:"errors,required"`
	Messages []DeletedImagesResponseMessage `json:"messages,required"`
	Result   interface{}                    `json:"result,required"`
	// Whether the API call was successful
	Success DeletedImagesResponseSuccess `json:"success,required"`
	JSON    deletedImagesResponseJSON    `json:"-"`
}

// deletedImagesResponseJSON contains the JSON metadata for the struct
// [DeletedImagesResponse]
type deletedImagesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeletedImagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deletedImagesResponseJSON) RawJSON() string {
	return r.raw
}

type DeletedImagesResponseError struct {
	Code             int64                             `json:"code,required"`
	Message          string                            `json:"message,required"`
	DocumentationURL string                            `json:"documentation_url"`
	Source           DeletedImagesResponseErrorsSource `json:"source"`
	JSON             deletedImagesResponseErrorJSON    `json:"-"`
}

// deletedImagesResponseErrorJSON contains the JSON metadata for the struct
// [DeletedImagesResponseError]
type deletedImagesResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeletedImagesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deletedImagesResponseErrorJSON) RawJSON() string {
	return r.raw
}

type DeletedImagesResponseErrorsSource struct {
	Pointer string                                `json:"pointer"`
	JSON    deletedImagesResponseErrorsSourceJSON `json:"-"`
}

// deletedImagesResponseErrorsSourceJSON contains the JSON metadata for the struct
// [DeletedImagesResponseErrorsSource]
type deletedImagesResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeletedImagesResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deletedImagesResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DeletedImagesResponseMessage struct {
	Code             int64                               `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           DeletedImagesResponseMessagesSource `json:"source"`
	JSON             deletedImagesResponseMessageJSON    `json:"-"`
}

// deletedImagesResponseMessageJSON contains the JSON metadata for the struct
// [DeletedImagesResponseMessage]
type deletedImagesResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeletedImagesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deletedImagesResponseMessageJSON) RawJSON() string {
	return r.raw
}

type DeletedImagesResponseMessagesSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    deletedImagesResponseMessagesSourceJSON `json:"-"`
}

// deletedImagesResponseMessagesSourceJSON contains the JSON metadata for the
// struct [DeletedImagesResponseMessagesSource]
type deletedImagesResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeletedImagesResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deletedImagesResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DeletedImagesResponseSuccess bool

const (
	DeletedImagesResponseSuccessTrue DeletedImagesResponseSuccess = true
)

func (r DeletedImagesResponseSuccess) IsKnown() bool {
	switch r {
	case DeletedImagesResponseSuccessTrue:
		return true
	}
	return false
}

type Image struct {
	// Image unique identifier.
	ID string `json:"id"`
	// Can set the creator field with an internal user ID.
	Creator string `json:"creator,nullable"`
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
	Creator           apijson.Field
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
	Errors   []ImageResponseSingleError   `json:"errors,required"`
	Messages []ImageResponseSingleMessage `json:"messages,required"`
	Result   Image                        `json:"result,required"`
	// Whether the API call was successful
	Success ImageResponseSingleSuccess `json:"success,required"`
	JSON    imageResponseSingleJSON    `json:"-"`
}

// imageResponseSingleJSON contains the JSON metadata for the struct
// [ImageResponseSingle]
type imageResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageResponseSingleJSON) RawJSON() string {
	return r.raw
}

type ImageResponseSingleError struct {
	Code             int64                           `json:"code,required"`
	Message          string                          `json:"message,required"`
	DocumentationURL string                          `json:"documentation_url"`
	Source           ImageResponseSingleErrorsSource `json:"source"`
	JSON             imageResponseSingleErrorJSON    `json:"-"`
}

// imageResponseSingleErrorJSON contains the JSON metadata for the struct
// [ImageResponseSingleError]
type imageResponseSingleErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ImageResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageResponseSingleErrorJSON) RawJSON() string {
	return r.raw
}

type ImageResponseSingleErrorsSource struct {
	Pointer string                              `json:"pointer"`
	JSON    imageResponseSingleErrorsSourceJSON `json:"-"`
}

// imageResponseSingleErrorsSourceJSON contains the JSON metadata for the struct
// [ImageResponseSingleErrorsSource]
type imageResponseSingleErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageResponseSingleErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageResponseSingleErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ImageResponseSingleMessage struct {
	Code             int64                             `json:"code,required"`
	Message          string                            `json:"message,required"`
	DocumentationURL string                            `json:"documentation_url"`
	Source           ImageResponseSingleMessagesSource `json:"source"`
	JSON             imageResponseSingleMessageJSON    `json:"-"`
}

// imageResponseSingleMessageJSON contains the JSON metadata for the struct
// [ImageResponseSingleMessage]
type imageResponseSingleMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ImageResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageResponseSingleMessageJSON) RawJSON() string {
	return r.raw
}

type ImageResponseSingleMessagesSource struct {
	Pointer string                                `json:"pointer"`
	JSON    imageResponseSingleMessagesSourceJSON `json:"-"`
}

// imageResponseSingleMessagesSourceJSON contains the JSON metadata for the struct
// [ImageResponseSingleMessagesSource]
type imageResponseSingleMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageResponseSingleMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageResponseSingleMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ImageResponseSingleSuccess bool

const (
	ImageResponseSingleSuccessTrue ImageResponseSingleSuccess = true
)

func (r ImageResponseSingleSuccess) IsKnown() bool {
	switch r {
	case ImageResponseSingleSuccessTrue:
		return true
	}
	return false
}

type AccountImageV1ListResponse struct {
	Errors   []AccountImageV1ListResponseError   `json:"errors,required"`
	Messages []AccountImageV1ListResponseMessage `json:"messages,required"`
	Result   AccountImageV1ListResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success AccountImageV1ListResponseSuccess `json:"success,required"`
	JSON    accountImageV1ListResponseJSON    `json:"-"`
}

// accountImageV1ListResponseJSON contains the JSON metadata for the struct
// [AccountImageV1ListResponse]
type accountImageV1ListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1ListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1ListResponseError struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           AccountImageV1ListResponseErrorsSource `json:"source"`
	JSON             accountImageV1ListResponseErrorJSON    `json:"-"`
}

// accountImageV1ListResponseErrorJSON contains the JSON metadata for the struct
// [AccountImageV1ListResponseError]
type accountImageV1ListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountImageV1ListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1ListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1ListResponseErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    accountImageV1ListResponseErrorsSourceJSON `json:"-"`
}

// accountImageV1ListResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountImageV1ListResponseErrorsSource]
type accountImageV1ListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1ListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1ListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1ListResponseMessage struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           AccountImageV1ListResponseMessagesSource `json:"source"`
	JSON             accountImageV1ListResponseMessageJSON    `json:"-"`
}

// accountImageV1ListResponseMessageJSON contains the JSON metadata for the struct
// [AccountImageV1ListResponseMessage]
type accountImageV1ListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountImageV1ListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1ListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1ListResponseMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    accountImageV1ListResponseMessagesSourceJSON `json:"-"`
}

// accountImageV1ListResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AccountImageV1ListResponseMessagesSource]
type accountImageV1ListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1ListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1ListResponseMessagesSourceJSON) RawJSON() string {
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

// Whether the API call was successful
type AccountImageV1ListResponseSuccess bool

const (
	AccountImageV1ListResponseSuccessTrue AccountImageV1ListResponseSuccess = true
)

func (r AccountImageV1ListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountImageV1ListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountImageV1StatsResponse struct {
	Errors   []AccountImageV1StatsResponseError   `json:"errors,required"`
	Messages []AccountImageV1StatsResponseMessage `json:"messages,required"`
	Result   AccountImageV1StatsResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success AccountImageV1StatsResponseSuccess `json:"success,required"`
	JSON    accountImageV1StatsResponseJSON    `json:"-"`
}

// accountImageV1StatsResponseJSON contains the JSON metadata for the struct
// [AccountImageV1StatsResponse]
type accountImageV1StatsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1StatsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1StatsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1StatsResponseError struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           AccountImageV1StatsResponseErrorsSource `json:"source"`
	JSON             accountImageV1StatsResponseErrorJSON    `json:"-"`
}

// accountImageV1StatsResponseErrorJSON contains the JSON metadata for the struct
// [AccountImageV1StatsResponseError]
type accountImageV1StatsResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountImageV1StatsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1StatsResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1StatsResponseErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    accountImageV1StatsResponseErrorsSourceJSON `json:"-"`
}

// accountImageV1StatsResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountImageV1StatsResponseErrorsSource]
type accountImageV1StatsResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1StatsResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1StatsResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1StatsResponseMessage struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           AccountImageV1StatsResponseMessagesSource `json:"source"`
	JSON             accountImageV1StatsResponseMessageJSON    `json:"-"`
}

// accountImageV1StatsResponseMessageJSON contains the JSON metadata for the struct
// [AccountImageV1StatsResponseMessage]
type accountImageV1StatsResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountImageV1StatsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1StatsResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1StatsResponseMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    accountImageV1StatsResponseMessagesSourceJSON `json:"-"`
}

// accountImageV1StatsResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AccountImageV1StatsResponseMessagesSource]
type accountImageV1StatsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1StatsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1StatsResponseMessagesSourceJSON) RawJSON() string {
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

// Whether the API call was successful
type AccountImageV1StatsResponseSuccess bool

const (
	AccountImageV1StatsResponseSuccessTrue AccountImageV1StatsResponseSuccess = true
)

func (r AccountImageV1StatsResponseSuccess) IsKnown() bool {
	switch r {
	case AccountImageV1StatsResponseSuccessTrue:
		return true
	}
	return false
}

type AccountImageV1UpdateParams struct {
	// Can set the creator field with an internal user ID.
	Creator param.Field[string] `json:"creator"`
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
	// Internal user ID set within the creator field. Setting to empty string "" will
	// return images where creator field is not set
	Creator param.Field[string] `query:"creator"`
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

type AccountImageV1UploadParams struct {
	// An optional custom unique identifier for your image.
	ID param.Field[string] `json:"id"`
	// Can set the creator field with an internal user ID.
	Creator param.Field[string] `json:"creator"`
	// An image binary data. Only needed when type is uploading a file.
	File param.Field[io.Reader] `json:"file" format:"binary"`
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
