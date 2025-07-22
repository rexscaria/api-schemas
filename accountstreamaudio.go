// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountStreamAudioService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountStreamAudioService] method instead.
type AccountStreamAudioService struct {
	Options []option.RequestOption
}

// NewAccountStreamAudioService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamAudioService(opts ...option.RequestOption) (r *AccountStreamAudioService) {
	r = &AccountStreamAudioService{}
	r.Options = opts
	return
}

// Adds an additional audio track to a video using the provided audio track URL.
func (r *AccountStreamAudioService) New(ctx context.Context, accountID string, identifier string, body AccountStreamAudioNewParams, opts ...option.RequestOption) (res *AddAudioTrack, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/copy", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Edits additional audio tracks on a video. Editing the default status of an audio
// track to `true` will mark all other audio tracks on the video default status to
// `false`.
func (r *AccountStreamAudioService) Update(ctx context.Context, accountID string, identifier string, audioIdentifier string, body AccountStreamAudioUpdateParams, opts ...option.RequestOption) (res *AddAudioTrack, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if audioIdentifier == "" {
		err = errors.New("missing required audio_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/%s", accountID, identifier, audioIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists additional audio tracks on a video. Note this API will not return
// information for audio attached to the video upload.
func (r *AccountStreamAudioService) List(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *AccountStreamAudioListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/audio", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes additional audio tracks on a video. Deleting a default audio track is
// not allowed. You must assign another audio track as default prior to deletion.
func (r *AccountStreamAudioService) Delete(ctx context.Context, accountID string, identifier string, audioIdentifier string, opts ...option.RequestOption) (res *DeletedStreamResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if audioIdentifier == "" {
		err = errors.New("missing required audio_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/%s", accountID, identifier, audioIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AddAudioTrack struct {
	Result AdditionalAudio   `json:"result"`
	JSON   addAudioTrackJSON `json:"-"`
	APIResponseStream
}

// addAudioTrackJSON contains the JSON metadata for the struct [AddAudioTrack]
type addAudioTrackJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddAudioTrack) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addAudioTrackJSON) RawJSON() string {
	return r.raw
}

type AdditionalAudio struct {
	// Denotes whether the audio track will be played by default in a player.
	Default bool `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label string `json:"label"`
	// Specifies the processing status of the video.
	Status AdditionalAudioStatus `json:"status"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid  string              `json:"uid"`
	JSON additionalAudioJSON `json:"-"`
}

// additionalAudioJSON contains the JSON metadata for the struct [AdditionalAudio]
type additionalAudioJSON struct {
	Default     apijson.Field
	Label       apijson.Field
	Status      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdditionalAudio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r additionalAudioJSON) RawJSON() string {
	return r.raw
}

// Specifies the processing status of the video.
type AdditionalAudioStatus string

const (
	AdditionalAudioStatusQueued AdditionalAudioStatus = "queued"
	AdditionalAudioStatusReady  AdditionalAudioStatus = "ready"
	AdditionalAudioStatusError  AdditionalAudioStatus = "error"
)

func (r AdditionalAudioStatus) IsKnown() bool {
	switch r {
	case AdditionalAudioStatusQueued, AdditionalAudioStatusReady, AdditionalAudioStatusError:
		return true
	}
	return false
}

type DeletedStreamResponse struct {
	Result string                    `json:"result"`
	JSON   deletedStreamResponseJSON `json:"-"`
	APIResponseSingleStream
}

// deletedStreamResponseJSON contains the JSON metadata for the struct
// [DeletedStreamResponse]
type deletedStreamResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeletedStreamResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deletedStreamResponseJSON) RawJSON() string {
	return r.raw
}

type AccountStreamAudioListResponse struct {
	Result []AdditionalAudio                  `json:"result"`
	JSON   accountStreamAudioListResponseJSON `json:"-"`
	APIResponseStream
}

// accountStreamAudioListResponseJSON contains the JSON metadata for the struct
// [AccountStreamAudioListResponse]
type accountStreamAudioListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamAudioListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamAudioListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountStreamAudioNewParams struct {
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label param.Field[string] `json:"label,required"`
	// An audio track URL. The server must be publicly routable and support `HTTP HEAD`
	// requests and `HTTP GET` range requests. The server should respond to `HTTP HEAD`
	// requests with a `content-range` header that includes the size of the file.
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r AccountStreamAudioNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamAudioUpdateParams struct {
	// Denotes whether the audio track will be played by default in a player.
	Default param.Field[bool] `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label param.Field[string] `json:"label"`
}

func (r AccountStreamAudioUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
