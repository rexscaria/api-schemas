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

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAIRunCfOpenAIService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfOpenAIService] method instead.
type AccountAIRunCfOpenAIService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfOpenAIService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfOpenAIService(opts ...option.RequestOption) (r *AccountAIRunCfOpenAIService) {
	r = &AccountAIRunCfOpenAIService{}
	r.Options = opts
	return
}

// Execute @cf/openai/whisper model.
func (r *AccountAIRunCfOpenAIService) ExecuteWhisper(ctx context.Context, accountID string, params AccountAIRunCfOpenAIExecuteWhisperParams, opts ...option.RequestOption) (res *AccountAIRunCfOpenAIExecuteWhisperResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/openai/whisper", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/openai/whisper-large-v3-turbo model.
func (r *AccountAIRunCfOpenAIService) ExecuteWhisperLargeV3Turbo(ctx context.Context, accountID string, params AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboParams, opts ...option.RequestOption) (res *AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/openai/whisper-large-v3-turbo", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/openai/whisper-tiny-en model.
func (r *AccountAIRunCfOpenAIService) ExecuteWhisperTinyEn(ctx context.Context, accountID string, params AccountAIRunCfOpenAIExecuteWhisperTinyEnParams, opts ...option.RequestOption) (res *AccountAIRunCfOpenAIExecuteWhisperTinyEnResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/openai/whisper-tiny-en", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfOpenAIExecuteWhisperResponse = interface{}

type AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponse = interface{}

type AccountAIRunCfOpenAIExecuteWhisperTinyEnResponse = interface{}

type AccountAIRunCfOpenAIExecuteWhisperParams struct {
	QueueRequest param.Field[string] `query:"queueRequest"`
	Body         io.Reader           `json:"body" format:"binary"`
}

func (r AccountAIRunCfOpenAIExecuteWhisperParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [AccountAIRunCfOpenAIExecuteWhisperParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfOpenAIExecuteWhisperParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboParams struct {
	// Base64 encoded value of the audio data.
	Audio        param.Field[string] `json:"audio,required"`
	QueueRequest param.Field[string] `query:"queueRequest"`
	// A text prompt to help provide context to the model on the contents of the audio.
	InitialPrompt param.Field[string] `json:"initial_prompt"`
	// The language of the audio being transcribed or translated.
	Language param.Field[string] `json:"language"`
	// The prefix it appended the the beginning of the output of the transcription and
	// can guide the transcription result.
	Prefix param.Field[string] `json:"prefix"`
	// Supported tasks are 'translate' or 'transcribe'.
	Task param.Field[string] `json:"task"`
	// Preprocess the audio with a voice activity detection model.
	VadFilter param.Field[bool] `json:"vad_filter"`
}

func (r AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboParams]'s
// query parameters as `url.Values`.
func (r AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfOpenAIExecuteWhisperTinyEnParams struct {
	QueueRequest param.Field[string] `query:"queueRequest"`
	Body         io.Reader           `json:"body" format:"binary"`
}

func (r AccountAIRunCfOpenAIExecuteWhisperTinyEnParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [AccountAIRunCfOpenAIExecuteWhisperTinyEnParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfOpenAIExecuteWhisperTinyEnParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
