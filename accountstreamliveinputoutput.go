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

// AccountStreamLiveInputOutputService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountStreamLiveInputOutputService] method instead.
type AccountStreamLiveInputOutputService struct {
	Options []option.RequestOption
}

// NewAccountStreamLiveInputOutputService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountStreamLiveInputOutputService(opts ...option.RequestOption) (r *AccountStreamLiveInputOutputService) {
	r = &AccountStreamLiveInputOutputService{}
	r.Options = opts
	return
}

// Creates a new output that can be used to simulcast or restream live video to
// other RTMP or SRT destinations. Outputs are always linked to a specific live
// input — one live input can have many outputs.
func (r *AccountStreamLiveInputOutputService) New(ctx context.Context, accountID string, liveInputIdentifier string, body AccountStreamLiveInputOutputNewParams, opts ...option.RequestOption) (res *OutputResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if liveInputIdentifier == "" {
		err = errors.New("missing required live_input_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs", accountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates the state of an output.
func (r *AccountStreamLiveInputOutputService) Update(ctx context.Context, accountID string, liveInputIdentifier string, outputIdentifier string, body AccountStreamLiveInputOutputUpdateParams, opts ...option.RequestOption) (res *OutputResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if liveInputIdentifier == "" {
		err = errors.New("missing required live_input_identifier parameter")
		return
	}
	if outputIdentifier == "" {
		err = errors.New("missing required output_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs/%s", accountID, liveInputIdentifier, outputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves all outputs associated with a specified live input.
func (r *AccountStreamLiveInputOutputService) List(ctx context.Context, accountID string, liveInputIdentifier string, opts ...option.RequestOption) (res *AccountStreamLiveInputOutputListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if liveInputIdentifier == "" {
		err = errors.New("missing required live_input_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs", accountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an output and removes it from the associated live input.
func (r *AccountStreamLiveInputOutputService) Delete(ctx context.Context, accountID string, liveInputIdentifier string, outputIdentifier string, opts ...option.RequestOption) (err error) {
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
	if outputIdentifier == "" {
		err = errors.New("missing required output_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s/outputs/%s", accountID, liveInputIdentifier, outputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Output struct {
	// When enabled, live video streamed to the associated live input will be sent to
	// the output URL. When disabled, live video will not be sent to the output URL,
	// even when streaming to the associated live input. Use this to control precisely
	// when you start and stop simulcasting to specific destinations like YouTube and
	// Twitch.
	Enabled bool `json:"enabled"`
	// The streamKey used to authenticate against an output's target.
	StreamKey string `json:"streamKey"`
	// A unique identifier for the output.
	Uid string `json:"uid"`
	// The URL an output uses to restream.
	URL  string     `json:"url"`
	JSON outputJSON `json:"-"`
}

// outputJSON contains the JSON metadata for the struct [Output]
type outputJSON struct {
	Enabled     apijson.Field
	StreamKey   apijson.Field
	Uid         apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Output) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outputJSON) RawJSON() string {
	return r.raw
}

type OutputResponseSingle struct {
	Errors   []StreamMessages `json:"errors,required"`
	Messages []StreamMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success OutputResponseSingleSuccess `json:"success,required"`
	Result  Output                      `json:"result"`
	JSON    outputResponseSingleJSON    `json:"-"`
}

// outputResponseSingleJSON contains the JSON metadata for the struct
// [OutputResponseSingle]
type outputResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutputResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outputResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OutputResponseSingleSuccess bool

const (
	OutputResponseSingleSuccessTrue OutputResponseSingleSuccess = true
)

func (r OutputResponseSingleSuccess) IsKnown() bool {
	switch r {
	case OutputResponseSingleSuccessTrue:
		return true
	}
	return false
}

type AccountStreamLiveInputOutputListResponse struct {
	Errors   []StreamMessages `json:"errors,required"`
	Messages []StreamMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountStreamLiveInputOutputListResponseSuccess `json:"success,required"`
	Result  []Output                                        `json:"result"`
	JSON    accountStreamLiveInputOutputListResponseJSON    `json:"-"`
}

// accountStreamLiveInputOutputListResponseJSON contains the JSON metadata for the
// struct [AccountStreamLiveInputOutputListResponse]
type accountStreamLiveInputOutputListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamLiveInputOutputListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamLiveInputOutputListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStreamLiveInputOutputListResponseSuccess bool

const (
	AccountStreamLiveInputOutputListResponseSuccessTrue AccountStreamLiveInputOutputListResponseSuccess = true
)

func (r AccountStreamLiveInputOutputListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStreamLiveInputOutputListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStreamLiveInputOutputNewParams struct {
	// The streamKey used to authenticate against an output's target.
	StreamKey param.Field[string] `json:"streamKey,required"`
	// The URL an output uses to restream.
	URL param.Field[string] `json:"url,required"`
	// When enabled, live video streamed to the associated live input will be sent to
	// the output URL. When disabled, live video will not be sent to the output URL,
	// even when streaming to the associated live input. Use this to control precisely
	// when you start and stop simulcasting to specific destinations like YouTube and
	// Twitch.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountStreamLiveInputOutputNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamLiveInputOutputUpdateParams struct {
	// When enabled, live video streamed to the associated live input will be sent to
	// the output URL. When disabled, live video will not be sent to the output URL,
	// even when streaming to the associated live input. Use this to control precisely
	// when you start and stop simulcasting to specific destinations like YouTube and
	// Twitch.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountStreamLiveInputOutputUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
