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

// AccountStreamLiveInputService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountStreamLiveInputService] method instead.
type AccountStreamLiveInputService struct {
	Options []option.RequestOption
	Outputs *AccountStreamLiveInputOutputService
}

// NewAccountStreamLiveInputService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamLiveInputService(opts ...option.RequestOption) (r *AccountStreamLiveInputService) {
	r = &AccountStreamLiveInputService{}
	r.Options = opts
	r.Outputs = NewAccountStreamLiveInputOutputService(opts...)
	return
}

// Creates a live input, and returns credentials that you or your users can use to
// stream live video to Cloudflare Stream.
func (r *AccountStreamLiveInputService) New(ctx context.Context, accountID string, body AccountStreamLiveInputNewParams, opts ...option.RequestOption) (res *LiveInputResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves details of an existing live input.
func (r *AccountStreamLiveInputService) Get(ctx context.Context, accountID string, liveInputIdentifier string, opts ...option.RequestOption) (res *LiveInputResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if liveInputIdentifier == "" {
		err = errors.New("missing required live_input_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", accountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a specified live input.
func (r *AccountStreamLiveInputService) Update(ctx context.Context, accountID string, liveInputIdentifier string, body AccountStreamLiveInputUpdateParams, opts ...option.RequestOption) (res *LiveInputResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if liveInputIdentifier == "" {
		err = errors.New("missing required live_input_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", accountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists the live inputs created for an account. To get the credentials needed to
// stream to a specific live input, request a single live input.
func (r *AccountStreamLiveInputService) List(ctx context.Context, accountID string, query AccountStreamLiveInputListParams, opts ...option.RequestOption) (res *AccountStreamLiveInputListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Prevents a live input from being streamed to and makes the live input
// inaccessible to any future API calls.
func (r *AccountStreamLiveInputService) Delete(ctx context.Context, accountID string, liveInputIdentifier string, body AccountStreamLiveInputDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if liveInputIdentifier == "" {
		err = errors.New("missing required live_input_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", accountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type LiveInputResponseSingle struct {
	// Details about a live input.
	Result LiveInputResponseSingleResult `json:"result"`
	JSON   liveInputResponseSingleJSON   `json:"-"`
	APIResponseSingleStream
}

// liveInputResponseSingleJSON contains the JSON metadata for the struct
// [LiveInputResponseSingle]
type liveInputResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Details about a live input.
type LiveInputResponseSingleResult struct {
	// The date and time the live input was created.
	Created time.Time `json:"created" format:"date-time"`
	// Indicates the number of days after which the live inputs recordings will be
	// deleted. When a stream completes and the recording is ready, the value is used
	// to calculate a scheduled deletion date for that recording. Omit the field to
	// indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion.
	DeleteRecordingAfterDays float64 `json:"deleteRecordingAfterDays"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta interface{} `json:"meta"`
	// The date and time the live input was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
	// most cases, the video will initially be viewable as a live video and transition
	// to on-demand after a condition is satisfied.
	Recording RecordingSettings `json:"recording"`
	// Details for streaming to an live input using RTMPS.
	Rtmps LiveInputResponseSingleResultRtmps `json:"rtmps"`
	// Details for playback from an live input using RTMPS.
	RtmpsPlayback LiveInputResponseSingleResultRtmpsPlayback `json:"rtmpsPlayback"`
	// Details for streaming to a live input using SRT.
	Srt LiveInputResponseSingleResultSrt `json:"srt"`
	// Details for playback from an live input using SRT.
	SrtPlayback LiveInputResponseSingleResultSrtPlayback `json:"srtPlayback"`
	// The connection status of a live input.
	Status LiveInputResponseSingleResultStatus `json:"status,nullable"`
	// A unique identifier for a live input.
	Uid string `json:"uid"`
	// Details for streaming to a live input using WebRTC.
	WebRtc LiveInputResponseSingleResultWebRtc `json:"webRTC"`
	// Details for playback from a live input using WebRTC.
	WebRtcPlayback LiveInputResponseSingleResultWebRtcPlayback `json:"webRTCPlayback"`
	JSON           liveInputResponseSingleResultJSON           `json:"-"`
}

// liveInputResponseSingleResultJSON contains the JSON metadata for the struct
// [LiveInputResponseSingleResult]
type liveInputResponseSingleResultJSON struct {
	Created                  apijson.Field
	DeleteRecordingAfterDays apijson.Field
	Meta                     apijson.Field
	Modified                 apijson.Field
	Recording                apijson.Field
	Rtmps                    apijson.Field
	RtmpsPlayback            apijson.Field
	Srt                      apijson.Field
	SrtPlayback              apijson.Field
	Status                   apijson.Field
	Uid                      apijson.Field
	WebRtc                   apijson.Field
	WebRtcPlayback           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *LiveInputResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputResponseSingleResultJSON) RawJSON() string {
	return r.raw
}

// Details for streaming to an live input using RTMPS.
type LiveInputResponseSingleResultRtmps struct {
	// The secret key to use when streaming via RTMPS to a live input.
	StreamKey string `json:"streamKey"`
	// The RTMPS URL you provide to the broadcaster, which they stream live video to.
	URL  string                                 `json:"url"`
	JSON liveInputResponseSingleResultRtmpsJSON `json:"-"`
}

// liveInputResponseSingleResultRtmpsJSON contains the JSON metadata for the struct
// [LiveInputResponseSingleResultRtmps]
type liveInputResponseSingleResultRtmpsJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingleResultRtmps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputResponseSingleResultRtmpsJSON) RawJSON() string {
	return r.raw
}

// Details for playback from an live input using RTMPS.
type LiveInputResponseSingleResultRtmpsPlayback struct {
	// The secret key to use for playback via RTMPS.
	StreamKey string `json:"streamKey"`
	// The URL used to play live video over RTMPS.
	URL  string                                         `json:"url"`
	JSON liveInputResponseSingleResultRtmpsPlaybackJSON `json:"-"`
}

// liveInputResponseSingleResultRtmpsPlaybackJSON contains the JSON metadata for
// the struct [LiveInputResponseSingleResultRtmpsPlayback]
type liveInputResponseSingleResultRtmpsPlaybackJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingleResultRtmpsPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputResponseSingleResultRtmpsPlaybackJSON) RawJSON() string {
	return r.raw
}

// Details for streaming to a live input using SRT.
type LiveInputResponseSingleResultSrt struct {
	// The secret key to use when streaming via SRT to a live input.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use when streaming via SRT.
	StreamID string `json:"streamId"`
	// The SRT URL you provide to the broadcaster, which they stream live video to.
	URL  string                               `json:"url"`
	JSON liveInputResponseSingleResultSrtJSON `json:"-"`
}

// liveInputResponseSingleResultSrtJSON contains the JSON metadata for the struct
// [LiveInputResponseSingleResultSrt]
type liveInputResponseSingleResultSrtJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingleResultSrt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputResponseSingleResultSrtJSON) RawJSON() string {
	return r.raw
}

// Details for playback from an live input using SRT.
type LiveInputResponseSingleResultSrtPlayback struct {
	// The secret key to use for playback via SRT.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use for playback via SRT.
	StreamID string `json:"streamId"`
	// The URL used to play live video over SRT.
	URL  string                                       `json:"url"`
	JSON liveInputResponseSingleResultSrtPlaybackJSON `json:"-"`
}

// liveInputResponseSingleResultSrtPlaybackJSON contains the JSON metadata for the
// struct [LiveInputResponseSingleResultSrtPlayback]
type liveInputResponseSingleResultSrtPlaybackJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingleResultSrtPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputResponseSingleResultSrtPlaybackJSON) RawJSON() string {
	return r.raw
}

// The connection status of a live input.
type LiveInputResponseSingleResultStatus string

const (
	LiveInputResponseSingleResultStatusConnected                LiveInputResponseSingleResultStatus = "connected"
	LiveInputResponseSingleResultStatusReconnected              LiveInputResponseSingleResultStatus = "reconnected"
	LiveInputResponseSingleResultStatusReconnecting             LiveInputResponseSingleResultStatus = "reconnecting"
	LiveInputResponseSingleResultStatusClientDisconnect         LiveInputResponseSingleResultStatus = "client_disconnect"
	LiveInputResponseSingleResultStatusTtlExceeded              LiveInputResponseSingleResultStatus = "ttl_exceeded"
	LiveInputResponseSingleResultStatusFailedToConnect          LiveInputResponseSingleResultStatus = "failed_to_connect"
	LiveInputResponseSingleResultStatusFailedToReconnect        LiveInputResponseSingleResultStatus = "failed_to_reconnect"
	LiveInputResponseSingleResultStatusNewConfigurationAccepted LiveInputResponseSingleResultStatus = "new_configuration_accepted"
)

func (r LiveInputResponseSingleResultStatus) IsKnown() bool {
	switch r {
	case LiveInputResponseSingleResultStatusConnected, LiveInputResponseSingleResultStatusReconnected, LiveInputResponseSingleResultStatusReconnecting, LiveInputResponseSingleResultStatusClientDisconnect, LiveInputResponseSingleResultStatusTtlExceeded, LiveInputResponseSingleResultStatusFailedToConnect, LiveInputResponseSingleResultStatusFailedToReconnect, LiveInputResponseSingleResultStatusNewConfigurationAccepted:
		return true
	}
	return false
}

// Details for streaming to a live input using WebRTC.
type LiveInputResponseSingleResultWebRtc struct {
	// The WebRTC URL you provide to the broadcaster, which they stream live video to.
	URL  string                                  `json:"url"`
	JSON liveInputResponseSingleResultWebRtcJSON `json:"-"`
}

// liveInputResponseSingleResultWebRtcJSON contains the JSON metadata for the
// struct [LiveInputResponseSingleResultWebRtc]
type liveInputResponseSingleResultWebRtcJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingleResultWebRtc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputResponseSingleResultWebRtcJSON) RawJSON() string {
	return r.raw
}

// Details for playback from a live input using WebRTC.
type LiveInputResponseSingleResultWebRtcPlayback struct {
	// The URL used to play live video over WebRTC.
	URL  string                                          `json:"url"`
	JSON liveInputResponseSingleResultWebRtcPlaybackJSON `json:"-"`
}

// liveInputResponseSingleResultWebRtcPlaybackJSON contains the JSON metadata for
// the struct [LiveInputResponseSingleResultWebRtcPlayback]
type liveInputResponseSingleResultWebRtcPlaybackJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputResponseSingleResultWebRtcPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputResponseSingleResultWebRtcPlaybackJSON) RawJSON() string {
	return r.raw
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type RecordingSettings struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Disables reporting the number of live viewers when this property is set to
	// `true`.
	HideLiveViewerCount bool `json:"hideLiveViewerCount"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode RecordingSettingsMode `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds int64                 `json:"timeoutSeconds"`
	JSON           recordingSettingsJSON `json:"-"`
}

// recordingSettingsJSON contains the JSON metadata for the struct
// [RecordingSettings]
type recordingSettingsJSON struct {
	AllowedOrigins      apijson.Field
	HideLiveViewerCount apijson.Field
	Mode                apijson.Field
	RequireSignedURLs   apijson.Field
	TimeoutSeconds      apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RecordingSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingSettingsJSON) RawJSON() string {
	return r.raw
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type RecordingSettingsMode string

const (
	RecordingSettingsModeOff       RecordingSettingsMode = "off"
	RecordingSettingsModeAutomatic RecordingSettingsMode = "automatic"
)

func (r RecordingSettingsMode) IsKnown() bool {
	switch r {
	case RecordingSettingsModeOff, RecordingSettingsModeAutomatic:
		return true
	}
	return false
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type RecordingSettingsParam struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// Disables reporting the number of live viewers when this property is set to
	// `true`.
	HideLiveViewerCount param.Field[bool] `json:"hideLiveViewerCount"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode param.Field[RecordingSettingsMode] `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds param.Field[int64] `json:"timeoutSeconds"`
}

func (r RecordingSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamLiveInputListResponse struct {
	Result AccountStreamLiveInputListResponseResult `json:"result"`
	JSON   accountStreamLiveInputListResponseJSON   `json:"-"`
	APIResponseStream
}

// accountStreamLiveInputListResponseJSON contains the JSON metadata for the struct
// [AccountStreamLiveInputListResponse]
type accountStreamLiveInputListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamLiveInputListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountStreamLiveInputListResponseResult struct {
	LiveInputs []AccountStreamLiveInputListResponseResultLiveInput `json:"liveInputs"`
	// The total number of remaining live inputs based on cursor position.
	Range int64 `json:"range"`
	// The total number of live inputs that match the provided filters.
	Total int64                                        `json:"total"`
	JSON  accountStreamLiveInputListResponseResultJSON `json:"-"`
}

// accountStreamLiveInputListResponseResultJSON contains the JSON metadata for the
// struct [AccountStreamLiveInputListResponseResult]
type accountStreamLiveInputListResponseResultJSON struct {
	LiveInputs  apijson.Field
	Range       apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamLiveInputListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountStreamLiveInputListResponseResultLiveInput struct {
	// The date and time the live input was created.
	Created time.Time `json:"created" format:"date-time"`
	// Indicates the number of days after which the live inputs recordings will be
	// deleted. When a stream completes and the recording is ready, the value is used
	// to calculate a scheduled deletion date for that recording. Omit the field to
	// indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion.
	DeleteRecordingAfterDays float64 `json:"deleteRecordingAfterDays"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta interface{} `json:"meta"`
	// The date and time the live input was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A unique identifier for a live input.
	Uid  string                                                `json:"uid"`
	JSON accountStreamLiveInputListResponseResultLiveInputJSON `json:"-"`
}

// accountStreamLiveInputListResponseResultLiveInputJSON contains the JSON metadata
// for the struct [AccountStreamLiveInputListResponseResultLiveInput]
type accountStreamLiveInputListResponseResultLiveInputJSON struct {
	Created                  apijson.Field
	DeleteRecordingAfterDays apijson.Field
	Meta                     apijson.Field
	Modified                 apijson.Field
	Uid                      apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccountStreamLiveInputListResponseResultLiveInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamLiveInputListResponseResultLiveInputJSON) RawJSON() string {
	return r.raw
}

type AccountStreamLiveInputNewParams struct {
	// Sets the creator ID asssociated with this live input.
	DefaultCreator param.Field[string] `json:"defaultCreator"`
	// Indicates the number of days after which the live inputs recordings will be
	// deleted. When a stream completes and the recording is ready, the value is used
	// to calculate a scheduled deletion date for that recording. Omit the field to
	// indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion.
	DeleteRecordingAfterDays param.Field[float64] `json:"deleteRecordingAfterDays"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta param.Field[interface{}] `json:"meta"`
	// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
	// most cases, the video will initially be viewable as a live video and transition
	// to on-demand after a condition is satisfied.
	Recording param.Field[RecordingSettingsParam] `json:"recording"`
}

func (r AccountStreamLiveInputNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamLiveInputUpdateParams struct {
	// Sets the creator ID asssociated with this live input.
	DefaultCreator param.Field[string] `json:"defaultCreator"`
	// Indicates the number of days after which the live inputs recordings will be
	// deleted. When a stream completes and the recording is ready, the value is used
	// to calculate a scheduled deletion date for that recording. Omit the field to
	// indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion.
	DeleteRecordingAfterDays param.Field[float64] `json:"deleteRecordingAfterDays"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta param.Field[interface{}] `json:"meta"`
	// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
	// most cases, the video will initially be viewable as a live video and transition
	// to on-demand after a condition is satisfied.
	Recording param.Field[RecordingSettingsParam] `json:"recording"`
}

func (r AccountStreamLiveInputUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamLiveInputListParams struct {
	// Includes the total number of videos associated with the submitted query
	// parameters.
	IncludeCounts param.Field[bool] `query:"include_counts"`
}

// URLQuery serializes [AccountStreamLiveInputListParams]'s query parameters as
// `url.Values`.
func (r AccountStreamLiveInputListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountStreamLiveInputDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountStreamLiveInputDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
