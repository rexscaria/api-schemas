// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountStreamService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountStreamService] method instead.
type AccountStreamService struct {
	Options    []option.RequestOption
	Audio      *AccountStreamAudioService
	Captions   *AccountStreamCaptionService
	Downloads  *AccountStreamDownloadService
	Keys       *AccountStreamKeyService
	LiveInputs *AccountStreamLiveInputService
	Watermarks *AccountStreamWatermarkService
	Webhook    *AccountStreamWebhookService
}

// NewAccountStreamService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountStreamService(opts ...option.RequestOption) (r *AccountStreamService) {
	r = &AccountStreamService{}
	r.Options = opts
	r.Audio = NewAccountStreamAudioService(opts...)
	r.Captions = NewAccountStreamCaptionService(opts...)
	r.Downloads = NewAccountStreamDownloadService(opts...)
	r.Keys = NewAccountStreamKeyService(opts...)
	r.LiveInputs = NewAccountStreamLiveInputService(opts...)
	r.Watermarks = NewAccountStreamWatermarkService(opts...)
	r.Webhook = NewAccountStreamWebhookService(opts...)
	return
}

// Initiates a video upload using the TUS protocol. On success, the server responds
// with a status code 201 (created) and includes a `location` header to indicate
// where the content should be uploaded. Refer to https://tus.io for protocol
// details.
func (r *AccountStreamService) New(ctx context.Context, accountID string, params AccountStreamNewParams, opts ...option.RequestOption) (err error) {
	if params.TusResumable.Present {
		opts = append(opts, option.WithHeader("Tus-Resumable", fmt.Sprintf("%s", params.TusResumable)))
	}
	if params.UploadLength.Present {
		opts = append(opts, option.WithHeader("Upload-Length", fmt.Sprintf("%s", params.UploadLength)))
	}
	if params.UploadCreator.Present {
		opts = append(opts, option.WithHeader("Upload-Creator", fmt.Sprintf("%s", params.UploadCreator)))
	}
	if params.UploadMetadata.Present {
		opts = append(opts, option.WithHeader("Upload-Metadata", fmt.Sprintf("%s", params.UploadMetadata)))
	}
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Fetches details for a single video.
func (r *AccountStreamService) Get(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *VideoResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit details for a single video.
func (r *AccountStreamService) Update(ctx context.Context, accountID string, identifier string, body AccountStreamUpdateParams, opts ...option.RequestOption) (res *VideoResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists up to 1000 videos from a single request. For a specific range, refer to
// the optional parameters.
func (r *AccountStreamService) List(ctx context.Context, accountID string, query AccountStreamListParams, opts ...option.RequestOption) (res *AccountStreamListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a video and its copies from Cloudflare Stream.
func (r *AccountStreamService) Delete(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Clips a video based on the specified start and end times provided in seconds.
func (r *AccountStreamService) Clip(ctx context.Context, accountID string, body AccountStreamClipParams, opts ...option.RequestOption) (res *AccountStreamClipResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/clip", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Uploads a video to Stream from a provided URL.
func (r *AccountStreamService) Copy(ctx context.Context, accountID string, params AccountStreamCopyParams, opts ...option.RequestOption) (res *VideoResponseSingle, err error) {
	if params.UploadCreator.Present {
		opts = append(opts, option.WithHeader("Upload-Creator", fmt.Sprintf("%s", params.UploadCreator)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/copy", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Creates a signed URL token for a video. If a body is not provided in the
// request, a token is created with default values.
func (r *AccountStreamService) NewSignedURL(ctx context.Context, accountID string, identifier string, body AccountStreamNewSignedURLParams, opts ...option.RequestOption) (res *AccountStreamNewSignedURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/token", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates a direct upload that allows video uploads without an API key.
func (r *AccountStreamService) DirectUpload(ctx context.Context, accountID string, params AccountStreamDirectUploadParams, opts ...option.RequestOption) (res *AccountStreamDirectUploadResponse, err error) {
	if params.UploadCreator.Present {
		opts = append(opts, option.WithHeader("Upload-Creator", fmt.Sprintf("%s", params.UploadCreator)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/direct_upload", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Fetches an HTML code snippet to embed a video in a web page delivered through
// Cloudflare. On success, returns an HTML fragment for use on web pages to display
// a video. On failure, returns a JSON response body.
func (r *AccountStreamService) GetEmbedCode(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/embed", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns information about an account's storage use.
func (r *AccountStreamService) GetStorageUsage(ctx context.Context, accountID string, query AccountStreamGetStorageUsageParams, opts ...option.RequestOption) (res *AccountStreamGetStorageUsageResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/storage-usage", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Specifies the processing status for all quality levels for a video.
type MediaState string

const (
	MediaStatePendingupload  MediaState = "pendingupload"
	MediaStateDownloading    MediaState = "downloading"
	MediaStateQueued         MediaState = "queued"
	MediaStateInprogress     MediaState = "inprogress"
	MediaStateReady          MediaState = "ready"
	MediaStateError          MediaState = "error"
	MediaStateLiveInprogress MediaState = "live-inprogress"
)

func (r MediaState) IsKnown() bool {
	switch r {
	case MediaStatePendingupload, MediaStateDownloading, MediaStateQueued, MediaStateInprogress, MediaStateReady, MediaStateError, MediaStateLiveInprogress:
		return true
	}
	return false
}

type Playback struct {
	// DASH Media Presentation Description for the video.
	Dash string `json:"dash"`
	// The HLS manifest for the video.
	Hls  string       `json:"hls"`
	JSON playbackJSON `json:"-"`
}

// playbackJSON contains the JSON metadata for the struct [Playback]
type playbackJSON struct {
	Dash        apijson.Field
	Hls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Playback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r playbackJSON) RawJSON() string {
	return r.raw
}

type VideoResponseSingle struct {
	Errors   []StreamMessages `json:"errors,required"`
	Messages []StreamMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success VideoResponseSingleSuccess `json:"success,required"`
	Result  Videos                     `json:"result"`
	JSON    videoResponseSingleJSON    `json:"-"`
}

// videoResponseSingleJSON contains the JSON metadata for the struct
// [VideoResponseSingle]
type videoResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videoResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type VideoResponseSingleSuccess bool

const (
	VideoResponseSingleSuccessTrue VideoResponseSingleSuccess = true
)

func (r VideoResponseSingleSuccess) IsKnown() bool {
	switch r {
	case VideoResponseSingleSuccessTrue:
		return true
	}
	return false
}

type Videos struct {
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// The date and time the media item was created.
	Created time.Time `json:"created" format:"date-time"`
	// A user-defined identifier for the media creator.
	Creator string `json:"creator"`
	// The duration of the video in seconds. A value of `-1` means the duration is
	// unknown. The duration becomes available after the upload and before the video is
	// ready.
	Duration float64     `json:"duration"`
	Input    VideosInput `json:"input"`
	// The live input ID used to upload a video with Stream Live.
	LiveInput string `json:"liveInput"`
	// The maximum duration in seconds for a video upload. Can be set for a video that
	// is not yet uploaded to limit its duration. Uploads that exceed the specified
	// duration will fail during processing. A value of `-1` means the value is
	// unknown.
	MaxDurationSeconds int64 `json:"maxDurationSeconds"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing videos.
	Meta interface{} `json:"meta"`
	// The date and time the media item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	Playback Playback  `json:"playback"`
	// The video's preview page URI. This field is omitted until encoding is complete.
	Preview string `json:"preview" format:"uri"`
	// Indicates whether the video is playable. The field is empty if the video is not
	// ready for viewing or the live stream is still in progress.
	ReadyToStream bool `json:"readyToStream"`
	// Indicates the time at which the video became playable. The field is empty if the
	// video is not ready for viewing or the live stream is still in progress.
	ReadyToStreamAt time.Time `json:"readyToStreamAt" format:"date-time"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Indicates the date and time at which the video will be deleted. Omit the field
	// to indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion. If specified, must be at least 30 days from upload time.
	ScheduledDeletion time.Time `json:"scheduledDeletion" format:"date-time"`
	// The size of the media item in bytes.
	Size float64 `json:"size"`
	// Specifies a detailed status for a video. If the `state` is `inprogress` or
	// `error`, the `step` field returns `encoding` or `manifest`. If the `state` is
	// `inprogress`, `pctComplete` returns a number between 0 and 100 to indicate the
	// approximate percent of completion. If the `state` is `error`, `errorReasonCode`
	// and `errorReasonText` provide additional details.
	Status VideosStatus `json:"status"`
	// The media item's thumbnail URI. This field is omitted until encoding is
	// complete.
	Thumbnail string `json:"thumbnail" format:"uri"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct float64 `json:"thumbnailTimestampPct"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid string `json:"uid"`
	// The date and time the media item was uploaded.
	Uploaded time.Time `json:"uploaded" format:"date-time"`
	// The date and time when the video upload URL is no longer valid for direct user
	// uploads.
	UploadExpiry time.Time  `json:"uploadExpiry" format:"date-time"`
	Watermark    Watermarks `json:"watermark"`
	JSON         videosJSON `json:"-"`
}

// videosJSON contains the JSON metadata for the struct [Videos]
type videosJSON struct {
	AllowedOrigins        apijson.Field
	Created               apijson.Field
	Creator               apijson.Field
	Duration              apijson.Field
	Input                 apijson.Field
	LiveInput             apijson.Field
	MaxDurationSeconds    apijson.Field
	Meta                  apijson.Field
	Modified              apijson.Field
	Playback              apijson.Field
	Preview               apijson.Field
	ReadyToStream         apijson.Field
	ReadyToStreamAt       apijson.Field
	RequireSignedURLs     apijson.Field
	ScheduledDeletion     apijson.Field
	Size                  apijson.Field
	Status                apijson.Field
	Thumbnail             apijson.Field
	ThumbnailTimestampPct apijson.Field
	Uid                   apijson.Field
	Uploaded              apijson.Field
	UploadExpiry          apijson.Field
	Watermark             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Videos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videosJSON) RawJSON() string {
	return r.raw
}

type VideosInput struct {
	// The video height in pixels. A value of `-1` means the height is unknown. The
	// value becomes available after the upload and before the video is ready.
	Height int64 `json:"height"`
	// The video width in pixels. A value of `-1` means the width is unknown. The value
	// becomes available after the upload and before the video is ready.
	Width int64           `json:"width"`
	JSON  videosInputJSON `json:"-"`
}

// videosInputJSON contains the JSON metadata for the struct [VideosInput]
type videosInputJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideosInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videosInputJSON) RawJSON() string {
	return r.raw
}

// Specifies a detailed status for a video. If the `state` is `inprogress` or
// `error`, the `step` field returns `encoding` or `manifest`. If the `state` is
// `inprogress`, `pctComplete` returns a number between 0 and 100 to indicate the
// approximate percent of completion. If the `state` is `error`, `errorReasonCode`
// and `errorReasonText` provide additional details.
type VideosStatus struct {
	// Specifies why the video failed to encode. This field is empty if the video is
	// not in an `error` state. Preferred for programmatic use.
	ErrorReasonCode string `json:"errorReasonCode"`
	// Specifies why the video failed to encode using a human readable error message in
	// English. This field is empty if the video is not in an `error` state.
	ErrorReasonText string `json:"errorReasonText"`
	// Indicates the size of the entire upload in bytes. The value must be a
	// non-negative integer.
	PctComplete string `json:"pctComplete"`
	// Specifies the processing status for all quality levels for a video.
	State MediaState       `json:"state"`
	JSON  videosStatusJSON `json:"-"`
}

// videosStatusJSON contains the JSON metadata for the struct [VideosStatus]
type videosStatusJSON struct {
	ErrorReasonCode apijson.Field
	ErrorReasonText apijson.Field
	PctComplete     apijson.Field
	State           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *VideosStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videosStatusJSON) RawJSON() string {
	return r.raw
}

type WatermarkAtUploadEvent struct {
	// The unique identifier for the watermark profile.
	Uid  string                     `json:"uid"`
	JSON watermarkAtUploadEventJSON `json:"-"`
}

// watermarkAtUploadEventJSON contains the JSON metadata for the struct
// [WatermarkAtUploadEvent]
type watermarkAtUploadEventJSON struct {
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatermarkAtUploadEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watermarkAtUploadEventJSON) RawJSON() string {
	return r.raw
}

type WatermarkAtUploadEventParam struct {
	// The unique identifier for the watermark profile.
	Uid param.Field[string] `json:"uid"`
}

func (r WatermarkAtUploadEventParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WatermarkAtUploadStreamParam struct {
	// The unique identifier for the watermark profile.
	Uid param.Field[string] `json:"uid"`
}

func (r WatermarkAtUploadStreamParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamListResponse struct {
	Errors   []StreamMessages `json:"errors,required"`
	Messages []StreamMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountStreamListResponseSuccess `json:"success,required"`
	// The total number of remaining videos based on cursor position.
	Range  int64    `json:"range"`
	Result []Videos `json:"result"`
	// The total number of videos that match the provided filters.
	Total int64                         `json:"total"`
	JSON  accountStreamListResponseJSON `json:"-"`
}

// accountStreamListResponseJSON contains the JSON metadata for the struct
// [AccountStreamListResponse]
type accountStreamListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Range       apijson.Field
	Result      apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStreamListResponseSuccess bool

const (
	AccountStreamListResponseSuccessTrue AccountStreamListResponseSuccess = true
)

func (r AccountStreamListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStreamListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStreamClipResponse struct {
	Errors   []StreamMessages `json:"errors,required"`
	Messages []StreamMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountStreamClipResponseSuccess `json:"success,required"`
	Result  AccountStreamClipResponseResult  `json:"result"`
	JSON    accountStreamClipResponseJSON    `json:"-"`
}

// accountStreamClipResponseJSON contains the JSON metadata for the struct
// [AccountStreamClipResponse]
type accountStreamClipResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamClipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamClipResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStreamClipResponseSuccess bool

const (
	AccountStreamClipResponseSuccessTrue AccountStreamClipResponseSuccess = true
)

func (r AccountStreamClipResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStreamClipResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStreamClipResponseResult struct {
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// The unique video identifier (UID).
	ClippedFromVideoUid string `json:"clippedFromVideoUID"`
	// The date and time the clip was created.
	Created time.Time `json:"created" format:"date-time"`
	// A user-defined identifier for the media creator.
	Creator string `json:"creator"`
	// Specifies the end time for the video clip in seconds.
	EndTimeSeconds int64 `json:"endTimeSeconds"`
	// The maximum duration in seconds for a video upload. Can be set for a video that
	// is not yet uploaded to limit its duration. Uploads that exceed the specified
	// duration will fail during processing. A value of `-1` means the value is
	// unknown.
	MaxDurationSeconds int64 `json:"maxDurationSeconds"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing videos.
	Meta interface{} `json:"meta"`
	// The date and time the live input was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	Playback Playback  `json:"playback"`
	// The video's preview page URI. This field is omitted until encoding is complete.
	Preview string `json:"preview" format:"uri"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Specifies the start time for the video clip in seconds.
	StartTimeSeconds int64 `json:"startTimeSeconds"`
	// Specifies the processing status for all quality levels for a video.
	Status MediaState `json:"status"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct float64                             `json:"thumbnailTimestampPct"`
	Watermark             WatermarkAtUploadEvent              `json:"watermark"`
	JSON                  accountStreamClipResponseResultJSON `json:"-"`
}

// accountStreamClipResponseResultJSON contains the JSON metadata for the struct
// [AccountStreamClipResponseResult]
type accountStreamClipResponseResultJSON struct {
	AllowedOrigins        apijson.Field
	ClippedFromVideoUid   apijson.Field
	Created               apijson.Field
	Creator               apijson.Field
	EndTimeSeconds        apijson.Field
	MaxDurationSeconds    apijson.Field
	Meta                  apijson.Field
	Modified              apijson.Field
	Playback              apijson.Field
	Preview               apijson.Field
	RequireSignedURLs     apijson.Field
	StartTimeSeconds      apijson.Field
	Status                apijson.Field
	ThumbnailTimestampPct apijson.Field
	Watermark             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountStreamClipResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamClipResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountStreamNewSignedURLResponse struct {
	Errors   []StreamMessages `json:"errors,required"`
	Messages []StreamMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountStreamNewSignedURLResponseSuccess `json:"success,required"`
	Result  AccountStreamNewSignedURLResponseResult  `json:"result"`
	JSON    accountStreamNewSignedURLResponseJSON    `json:"-"`
}

// accountStreamNewSignedURLResponseJSON contains the JSON metadata for the struct
// [AccountStreamNewSignedURLResponse]
type accountStreamNewSignedURLResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamNewSignedURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamNewSignedURLResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStreamNewSignedURLResponseSuccess bool

const (
	AccountStreamNewSignedURLResponseSuccessTrue AccountStreamNewSignedURLResponseSuccess = true
)

func (r AccountStreamNewSignedURLResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStreamNewSignedURLResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStreamNewSignedURLResponseResult struct {
	// The signed token used with the signed URLs feature.
	Token string                                      `json:"token"`
	JSON  accountStreamNewSignedURLResponseResultJSON `json:"-"`
}

// accountStreamNewSignedURLResponseResultJSON contains the JSON metadata for the
// struct [AccountStreamNewSignedURLResponseResult]
type accountStreamNewSignedURLResponseResultJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamNewSignedURLResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamNewSignedURLResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountStreamDirectUploadResponse struct {
	Errors   []StreamMessages `json:"errors,required"`
	Messages []StreamMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountStreamDirectUploadResponseSuccess `json:"success,required"`
	Result  AccountStreamDirectUploadResponseResult  `json:"result"`
	JSON    accountStreamDirectUploadResponseJSON    `json:"-"`
}

// accountStreamDirectUploadResponseJSON contains the JSON metadata for the struct
// [AccountStreamDirectUploadResponse]
type accountStreamDirectUploadResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamDirectUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamDirectUploadResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStreamDirectUploadResponseSuccess bool

const (
	AccountStreamDirectUploadResponseSuccessTrue AccountStreamDirectUploadResponseSuccess = true
)

func (r AccountStreamDirectUploadResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStreamDirectUploadResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStreamDirectUploadResponseResult struct {
	// Indicates the date and time at which the video will be deleted. Omit the field
	// to indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion. If specified, must be at least 30 days from upload time.
	ScheduledDeletion time.Time `json:"scheduledDeletion" format:"date-time"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid string `json:"uid"`
	// The URL an unauthenticated upload can use for a single
	// `HTTP POST multipart/form-data` request.
	UploadURL string                                      `json:"uploadURL"`
	Watermark Watermarks                                  `json:"watermark"`
	JSON      accountStreamDirectUploadResponseResultJSON `json:"-"`
}

// accountStreamDirectUploadResponseResultJSON contains the JSON metadata for the
// struct [AccountStreamDirectUploadResponseResult]
type accountStreamDirectUploadResponseResultJSON struct {
	ScheduledDeletion apijson.Field
	Uid               apijson.Field
	UploadURL         apijson.Field
	Watermark         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountStreamDirectUploadResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamDirectUploadResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountStreamGetStorageUsageResponse struct {
	Errors   []StreamMessages `json:"errors,required"`
	Messages []StreamMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountStreamGetStorageUsageResponseSuccess `json:"success,required"`
	Result  AccountStreamGetStorageUsageResponseResult  `json:"result"`
	JSON    accountStreamGetStorageUsageResponseJSON    `json:"-"`
}

// accountStreamGetStorageUsageResponseJSON contains the JSON metadata for the
// struct [AccountStreamGetStorageUsageResponse]
type accountStreamGetStorageUsageResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamGetStorageUsageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamGetStorageUsageResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStreamGetStorageUsageResponseSuccess bool

const (
	AccountStreamGetStorageUsageResponseSuccessTrue AccountStreamGetStorageUsageResponseSuccess = true
)

func (r AccountStreamGetStorageUsageResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStreamGetStorageUsageResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStreamGetStorageUsageResponseResult struct {
	// A user-defined identifier for the media creator.
	Creator string `json:"creator"`
	// The total minutes of video content stored in the account.
	TotalStorageMinutes int64 `json:"totalStorageMinutes"`
	// The storage capacity alloted for the account.
	TotalStorageMinutesLimit int64 `json:"totalStorageMinutesLimit"`
	// The total count of videos associated with the account.
	VideoCount int64                                          `json:"videoCount"`
	JSON       accountStreamGetStorageUsageResponseResultJSON `json:"-"`
}

// accountStreamGetStorageUsageResponseResultJSON contains the JSON metadata for
// the struct [AccountStreamGetStorageUsageResponseResult]
type accountStreamGetStorageUsageResponseResultJSON struct {
	Creator                  apijson.Field
	TotalStorageMinutes      apijson.Field
	TotalStorageMinutesLimit apijson.Field
	VideoCount               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccountStreamGetStorageUsageResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamGetStorageUsageResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountStreamNewParams struct {
	Body interface{} `json:"body,required"`
	// Specifies the TUS protocol version. This value must be included in every upload
	// request. Notes: The only supported version of TUS protocol is 1.0.0.
	TusResumable param.Field[AccountStreamNewParamsTusResumable] `header:"Tus-Resumable,required"`
	// Indicates the size of the entire upload in bytes. The value must be a
	// non-negative integer.
	UploadLength param.Field[int64] `header:"Upload-Length,required"`
	// Provisions a URL to let your end users upload videos directly to Cloudflare
	// Stream without exposing your API token to clients.
	DirectUser param.Field[bool] `query:"direct_user"`
	// A user-defined identifier for the media creator.
	UploadCreator param.Field[string] `header:"Upload-Creator"`
	// Comma-separated key-value pairs following the TUS protocol specification. Values
	// are Base-64 encoded. Supported keys: `name`, `requiresignedurls`,
	// `allowedorigins`, `thumbnailtimestamppct`, `watermark`, `scheduleddeletion`,
	// `maxdurationseconds`.
	UploadMetadata param.Field[string] `header:"Upload-Metadata"`
}

func (r AccountStreamNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountStreamNewParams]'s query parameters as `url.Values`.
func (r AccountStreamNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the TUS protocol version. This value must be included in every upload
// request. Notes: The only supported version of TUS protocol is 1.0.0.
type AccountStreamNewParamsTusResumable string

const (
	AccountStreamNewParamsTusResumable1_0_0 AccountStreamNewParamsTusResumable = "1.0.0"
)

func (r AccountStreamNewParamsTusResumable) IsKnown() bool {
	switch r {
	case AccountStreamNewParamsTusResumable1_0_0:
		return true
	}
	return false
}

type AccountStreamUpdateParams struct {
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `json:"creator"`
	// The maximum duration in seconds for a video upload. Can be set for a video that
	// is not yet uploaded to limit its duration. Uploads that exceed the specified
	// duration will fail during processing. A value of `-1` means the value is
	// unknown.
	MaxDurationSeconds param.Field[int64] `json:"maxDurationSeconds"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing videos.
	Meta param.Field[interface{}] `json:"meta"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Indicates the date and time at which the video will be deleted. Omit the field
	// to indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion. If specified, must be at least 30 days from upload time.
	ScheduledDeletion param.Field[time.Time] `json:"scheduledDeletion" format:"date-time"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct param.Field[float64] `json:"thumbnailTimestampPct"`
	// The date and time when the video upload URL is no longer valid for direct user
	// uploads.
	UploadExpiry param.Field[time.Time] `json:"uploadExpiry" format:"date-time"`
}

func (r AccountStreamUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamListParams struct {
	// Lists videos in ascending order of creation.
	Asc param.Field[bool] `query:"asc"`
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `query:"creator"`
	// Lists videos created before the specified date.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// Includes the total number of videos associated with the submitted query
	// parameters.
	IncludeCounts param.Field[bool] `query:"include_counts"`
	// Searches over the `name` key in the `meta` field. This field can be set with or
	// after the upload request.
	Search param.Field[string] `query:"search"`
	// Lists videos created after the specified date.
	Start param.Field[time.Time] `query:"start" format:"date-time"`
	// Specifies the processing status for all quality levels for a video.
	Status param.Field[MediaState] `query:"status"`
	// Specifies whether the video is `vod` or `live`.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [AccountStreamListParams]'s query parameters as
// `url.Values`.
func (r AccountStreamListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountStreamClipParams struct {
	// The unique video identifier (UID).
	ClippedFromVideoUid param.Field[string] `json:"clippedFromVideoUID,required"`
	// Specifies the end time for the video clip in seconds.
	EndTimeSeconds param.Field[int64] `json:"endTimeSeconds,required"`
	// Specifies the start time for the video clip in seconds.
	StartTimeSeconds param.Field[int64] `json:"startTimeSeconds,required"`
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `json:"creator"`
	// The maximum duration in seconds for a video upload. Can be set for a video that
	// is not yet uploaded to limit its duration. Uploads that exceed the specified
	// duration will fail during processing. A value of `-1` means the value is
	// unknown.
	MaxDurationSeconds param.Field[int64] `json:"maxDurationSeconds"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct param.Field[float64]                     `json:"thumbnailTimestampPct"`
	Watermark             param.Field[WatermarkAtUploadEventParam] `json:"watermark"`
}

func (r AccountStreamClipParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamCopyParams struct {
	// A video's URL. The server must be publicly routable and support `HTTP HEAD`
	// requests and `HTTP GET` range requests. The server should respond to `HTTP HEAD`
	// requests with a `content-range` header that includes the size of the file.
	URL param.Field[string] `json:"url,required" format:"uri"`
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `json:"creator"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing videos.
	Meta param.Field[interface{}] `json:"meta"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Indicates the date and time at which the video will be deleted. Omit the field
	// to indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion. If specified, must be at least 30 days from upload time.
	ScheduledDeletion param.Field[time.Time] `json:"scheduledDeletion" format:"date-time"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct param.Field[float64]                      `json:"thumbnailTimestampPct"`
	Watermark             param.Field[WatermarkAtUploadStreamParam] `json:"watermark"`
	// A user-defined identifier for the media creator.
	UploadCreator param.Field[string] `header:"Upload-Creator"`
}

func (r AccountStreamCopyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamNewSignedURLParams struct {
	// The optional ID of a Stream signing key. If present, the `pem` field is also
	// required.
	ID param.Field[string] `json:"id"`
	// The optional list of access rule constraints on the token. Access can be blocked
	// or allowed based on an IP, IP range, or by country. Access rules are evaluated
	// from first to last. If a rule matches, the associated action is applied and no
	// further rules are evaluated.
	AccessRules param.Field[[]AccountStreamNewSignedURLParamsAccessRule] `json:"accessRules"`
	// The optional boolean value that enables using signed tokens to access MP4
	// download links for a video.
	Downloadable param.Field[bool] `json:"downloadable"`
	// The optional unix epoch timestamp that specficies the time after a token is not
	// accepted. The maximum time specification is 24 hours from issuing time. If this
	// field is not set, the default is one hour after issuing.
	Exp param.Field[int64] `json:"exp"`
	// The optional unix epoch timestamp that specifies the time before a the token is
	// not accepted. If this field is not set, the default is one hour before issuing.
	Nbf param.Field[int64] `json:"nbf"`
	// The optional base64 encoded private key in PEM format associated with a Stream
	// signing key. If present, the `id` field is also required.
	Pem param.Field[string] `json:"pem"`
}

func (r AccountStreamNewSignedURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines rules for fine-grained control over content than signed URL tokens
// alone. Access rules primarily make tokens conditionally valid based on user
// information. Access Rules are specified on token payloads as the `accessRules`
// property containing an array of Rule objects.
type AccountStreamNewSignedURLParamsAccessRule struct {
	// The action to take when a request matches a rule. If the action is `block`, the
	// signed token blocks views for viewers matching the rule.
	Action param.Field[AccountStreamNewSignedURLParamsAccessRulesAction] `json:"action"`
	// An array of 2-letter country codes in ISO 3166-1 Alpha-2 format used to match
	// requests.
	Country param.Field[[]string] `json:"country"`
	// An array of IPv4 or IPV6 addresses or CIDRs used to match requests.
	IP param.Field[[]string] `json:"ip"`
	// Lists available rule types to match for requests. An `any` type matches all
	// requests and can be used as a wildcard to apply default actions after other
	// rules.
	Type param.Field[AccountStreamNewSignedURLParamsAccessRulesType] `json:"type"`
}

func (r AccountStreamNewSignedURLParamsAccessRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when a request matches a rule. If the action is `block`, the
// signed token blocks views for viewers matching the rule.
type AccountStreamNewSignedURLParamsAccessRulesAction string

const (
	AccountStreamNewSignedURLParamsAccessRulesActionAllow AccountStreamNewSignedURLParamsAccessRulesAction = "allow"
	AccountStreamNewSignedURLParamsAccessRulesActionBlock AccountStreamNewSignedURLParamsAccessRulesAction = "block"
)

func (r AccountStreamNewSignedURLParamsAccessRulesAction) IsKnown() bool {
	switch r {
	case AccountStreamNewSignedURLParamsAccessRulesActionAllow, AccountStreamNewSignedURLParamsAccessRulesActionBlock:
		return true
	}
	return false
}

// Lists available rule types to match for requests. An `any` type matches all
// requests and can be used as a wildcard to apply default actions after other
// rules.
type AccountStreamNewSignedURLParamsAccessRulesType string

const (
	AccountStreamNewSignedURLParamsAccessRulesTypeAny            AccountStreamNewSignedURLParamsAccessRulesType = "any"
	AccountStreamNewSignedURLParamsAccessRulesTypeIPSrc          AccountStreamNewSignedURLParamsAccessRulesType = "ip.src"
	AccountStreamNewSignedURLParamsAccessRulesTypeIPGeoipCountry AccountStreamNewSignedURLParamsAccessRulesType = "ip.geoip.country"
)

func (r AccountStreamNewSignedURLParamsAccessRulesType) IsKnown() bool {
	switch r {
	case AccountStreamNewSignedURLParamsAccessRulesTypeAny, AccountStreamNewSignedURLParamsAccessRulesTypeIPSrc, AccountStreamNewSignedURLParamsAccessRulesTypeIPGeoipCountry:
		return true
	}
	return false
}

type AccountStreamDirectUploadParams struct {
	// The maximum duration in seconds for a video upload. Can be set for a video that
	// is not yet uploaded to limit its duration. Uploads that exceed the specified
	// duration will fail during processing. A value of `-1` means the value is
	// unknown.
	MaxDurationSeconds param.Field[int64] `json:"maxDurationSeconds,required"`
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `json:"creator"`
	// The date and time after upload when videos will not be accepted.
	Expiry param.Field[time.Time] `json:"expiry" format:"date-time"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing videos.
	Meta param.Field[interface{}] `json:"meta"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Indicates the date and time at which the video will be deleted. Omit the field
	// to indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion. If specified, must be at least 30 days from upload time.
	ScheduledDeletion param.Field[time.Time] `json:"scheduledDeletion" format:"date-time"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct param.Field[float64]                      `json:"thumbnailTimestampPct"`
	Watermark             param.Field[WatermarkAtUploadStreamParam] `json:"watermark"`
	// A user-defined identifier for the media creator.
	UploadCreator param.Field[string] `header:"Upload-Creator"`
}

func (r AccountStreamDirectUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamGetStorageUsageParams struct {
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `query:"creator"`
}

// URLQuery serializes [AccountStreamGetStorageUsageParams]'s query parameters as
// `url.Values`.
func (r AccountStreamGetStorageUsageParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
