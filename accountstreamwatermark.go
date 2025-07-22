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

// AccountStreamWatermarkService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountStreamWatermarkService] method instead.
type AccountStreamWatermarkService struct {
	Options []option.RequestOption
}

// NewAccountStreamWatermarkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamWatermarkService(opts ...option.RequestOption) (r *AccountStreamWatermarkService) {
	r = &AccountStreamWatermarkService{}
	r.Options = opts
	return
}

// Creates watermark profiles using a single `HTTP POST multipart/form-data`
// request.
func (r *AccountStreamWatermarkService) New(ctx context.Context, accountID string, body AccountStreamWatermarkNewParams, opts ...option.RequestOption) (res *WatermarkResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/watermarks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves details for a single watermark profile.
func (r *AccountStreamWatermarkService) Get(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *WatermarkResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/watermarks/%s", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all watermark profiles for an account.
func (r *AccountStreamWatermarkService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountStreamWatermarkListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/watermarks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a watermark profile.
func (r *AccountStreamWatermarkService) Delete(ctx context.Context, accountID string, identifier string, body AccountStreamWatermarkDeleteParams, opts ...option.RequestOption) (res *AccountStreamWatermarkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/watermarks/%s", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type APIResponseSingleStream struct {
	Errors   []StreamMessages `json:"errors,required"`
	Messages []StreamMessages `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseSingleStreamSuccess `json:"success,required"`
	JSON    apiResponseSingleStreamJSON    `json:"-"`
}

// apiResponseSingleStreamJSON contains the JSON metadata for the struct
// [APIResponseSingleStream]
type apiResponseSingleStreamJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleStream) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleStreamJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseSingleStreamSuccess bool

const (
	APIResponseSingleStreamSuccessTrue APIResponseSingleStreamSuccess = true
)

func (r APIResponseSingleStreamSuccess) IsKnown() bool {
	switch r {
	case APIResponseSingleStreamSuccessTrue:
		return true
	}
	return false
}

type WatermarkResponseSingle struct {
	Result Watermarks                  `json:"result"`
	JSON   watermarkResponseSingleJSON `json:"-"`
	APIResponseSingleStream
}

// watermarkResponseSingleJSON contains the JSON metadata for the struct
// [WatermarkResponseSingle]
type watermarkResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatermarkResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watermarkResponseSingleJSON) RawJSON() string {
	return r.raw
}

type Watermarks struct {
	// The date and a time a watermark profile was created.
	Created time.Time `json:"created" format:"date-time"`
	// The source URL for a downloaded image. If the watermark profile was created via
	// direct upload, this field is null.
	DownloadedFrom string `json:"downloadedFrom"`
	// The height of the image in pixels.
	Height int64 `json:"height"`
	// A short description of the watermark profile.
	Name string `json:"name"`
	// The translucency of the image. A value of `0.0` makes the image completely
	// transparent, and `1.0` makes the image completely opaque. Note that if the image
	// is already semi-transparent, setting this to `1.0` will not make the image
	// completely opaque.
	Opacity float64 `json:"opacity"`
	// The whitespace between the adjacent edges (determined by position) of the video
	// and the image. `0.0` indicates no padding, and `1.0` indicates a fully padded
	// video width or length, as determined by the algorithm.
	Padding float64 `json:"padding"`
	// The location of the image. Valid positions are: `upperRight`, `upperLeft`,
	// `lowerLeft`, `lowerRight`, and `center`. Note that `center` ignores the
	// `padding` parameter.
	Position string `json:"position"`
	// The size of the image relative to the overall size of the video. This parameter
	// will adapt to horizontal and vertical videos automatically. `0.0` indicates no
	// scaling (use the size of the image as-is), and `1.0 `fills the entire video.
	Scale float64 `json:"scale"`
	// The size of the image in bytes.
	Size float64 `json:"size"`
	// The unique identifier for a watermark profile.
	Uid string `json:"uid"`
	// The width of the image in pixels.
	Width int64          `json:"width"`
	JSON  watermarksJSON `json:"-"`
}

// watermarksJSON contains the JSON metadata for the struct [Watermarks]
type watermarksJSON struct {
	Created        apijson.Field
	DownloadedFrom apijson.Field
	Height         apijson.Field
	Name           apijson.Field
	Opacity        apijson.Field
	Padding        apijson.Field
	Position       apijson.Field
	Scale          apijson.Field
	Size           apijson.Field
	Uid            apijson.Field
	Width          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Watermarks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watermarksJSON) RawJSON() string {
	return r.raw
}

type AccountStreamWatermarkListResponse struct {
	Result []Watermarks                           `json:"result"`
	JSON   accountStreamWatermarkListResponseJSON `json:"-"`
	APIResponseStream
}

// accountStreamWatermarkListResponseJSON contains the JSON metadata for the struct
// [AccountStreamWatermarkListResponse]
type accountStreamWatermarkListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWatermarkListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamWatermarkListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountStreamWatermarkDeleteResponse struct {
	Result string                                   `json:"result"`
	JSON   accountStreamWatermarkDeleteResponseJSON `json:"-"`
	APIResponseSingleStream
}

// accountStreamWatermarkDeleteResponseJSON contains the JSON metadata for the
// struct [AccountStreamWatermarkDeleteResponse]
type accountStreamWatermarkDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWatermarkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamWatermarkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountStreamWatermarkNewParams struct {
	// The image file to upload.
	File param.Field[string] `json:"file,required"`
	// A short description of the watermark profile.
	Name param.Field[string] `json:"name"`
	// The translucency of the image. A value of `0.0` makes the image completely
	// transparent, and `1.0` makes the image completely opaque. Note that if the image
	// is already semi-transparent, setting this to `1.0` will not make the image
	// completely opaque.
	Opacity param.Field[float64] `json:"opacity"`
	// The whitespace between the adjacent edges (determined by position) of the video
	// and the image. `0.0` indicates no padding, and `1.0` indicates a fully padded
	// video width or length, as determined by the algorithm.
	Padding param.Field[float64] `json:"padding"`
	// The location of the image. Valid positions are: `upperRight`, `upperLeft`,
	// `lowerLeft`, `lowerRight`, and `center`. Note that `center` ignores the
	// `padding` parameter.
	Position param.Field[string] `json:"position"`
	// The size of the image relative to the overall size of the video. This parameter
	// will adapt to horizontal and vertical videos automatically. `0.0` indicates no
	// scaling (use the size of the image as-is), and `1.0 `fills the entire video.
	Scale param.Field[float64] `json:"scale"`
}

func (r AccountStreamWatermarkNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type AccountStreamWatermarkDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountStreamWatermarkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
