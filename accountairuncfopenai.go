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

	"github.com/stainless-sdks/cf-rex-go/internal/apiform"
	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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

type AccountAIRunCfOpenAIExecuteWhisperResponse struct {
	Result  AccountAIRunCfOpenAIExecuteWhisperResponseResult `json:"result"`
	Success bool                                             `json:"success"`
	JSON    accountAIRunCfOpenAIExecuteWhisperResponseJSON   `json:"-"`
}

// accountAIRunCfOpenAIExecuteWhisperResponseJSON contains the JSON metadata for
// the struct [AccountAIRunCfOpenAIExecuteWhisperResponse]
type accountAIRunCfOpenAIExecuteWhisperResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfOpenAIExecuteWhisperResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfOpenAIExecuteWhisperResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfOpenAIExecuteWhisperResponseResult struct {
	// The transcription
	Text      string                                                 `json:"text,required"`
	Vtt       string                                                 `json:"vtt"`
	WordCount float64                                                `json:"word_count"`
	Words     []AccountAIRunCfOpenAIExecuteWhisperResponseResultWord `json:"words"`
	JSON      accountAIRunCfOpenAIExecuteWhisperResponseResultJSON   `json:"-"`
}

// accountAIRunCfOpenAIExecuteWhisperResponseResultJSON contains the JSON metadata
// for the struct [AccountAIRunCfOpenAIExecuteWhisperResponseResult]
type accountAIRunCfOpenAIExecuteWhisperResponseResultJSON struct {
	Text        apijson.Field
	Vtt         apijson.Field
	WordCount   apijson.Field
	Words       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfOpenAIExecuteWhisperResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfOpenAIExecuteWhisperResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfOpenAIExecuteWhisperResponseResultWord struct {
	// The ending second when the word completes
	End float64 `json:"end"`
	// The second this word begins in the recording
	Start float64                                                  `json:"start"`
	Word  string                                                   `json:"word"`
	JSON  accountAIRunCfOpenAIExecuteWhisperResponseResultWordJSON `json:"-"`
}

// accountAIRunCfOpenAIExecuteWhisperResponseResultWordJSON contains the JSON
// metadata for the struct [AccountAIRunCfOpenAIExecuteWhisperResponseResultWord]
type accountAIRunCfOpenAIExecuteWhisperResponseResultWordJSON struct {
	End         apijson.Field
	Start       apijson.Field
	Word        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfOpenAIExecuteWhisperResponseResultWord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfOpenAIExecuteWhisperResponseResultWordJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponse struct {
	Result  AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResult `json:"result"`
	Success bool                                                         `json:"success"`
	JSON    accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseJSON   `json:"-"`
}

// accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseJSON contains the JSON
// metadata for the struct [AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponse]
type accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResult struct {
	// The complete transcription of the audio.
	Text              string                                                                        `json:"text,required"`
	Segments          []AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegment         `json:"segments"`
	TranscriptionInfo AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultTranscriptionInfo `json:"transcription_info"`
	// The transcription in WebVTT format, which includes timing and text information
	// for use in subtitles.
	Vtt string `json:"vtt"`
	// The total number of words in the transcription.
	WordCount float64                                                          `json:"word_count"`
	JSON      accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultJSON `json:"-"`
}

// accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResult]
type accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultJSON struct {
	Text              apijson.Field
	Segments          apijson.Field
	TranscriptionInfo apijson.Field
	Vtt               apijson.Field
	WordCount         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegment struct {
	// The average log probability of the predictions for the words in this segment,
	// indicating overall confidence.
	AvgLogprob float64 `json:"avg_logprob"`
	// The compression ratio of the input to the output, measuring how much the text
	// was compressed during the transcription process.
	CompressionRatio float64 `json:"compression_ratio"`
	// The ending time of the segment within the audio, in seconds.
	End float64 `json:"end"`
	// The probability that the segment contains no speech, represented as a decimal
	// between 0 and 1.
	NoSpeechProb float64 `json:"no_speech_prob"`
	// The starting time of the segment within the audio, in seconds.
	Start float64 `json:"start"`
	// The temperature used in the decoding process, controlling randomness in
	// predictions. Lower values result in more deterministic outputs.
	Temperature float64 `json:"temperature"`
	// The transcription of the segment.
	Text  string                                                                     `json:"text"`
	Words []AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegmentsWord `json:"words"`
	JSON  accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegmentJSON    `json:"-"`
}

// accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegmentJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegment]
type accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegmentJSON struct {
	AvgLogprob       apijson.Field
	CompressionRatio apijson.Field
	End              apijson.Field
	NoSpeechProb     apijson.Field
	Start            apijson.Field
	Temperature      apijson.Field
	Text             apijson.Field
	Words            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegmentJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegmentsWord struct {
	// The ending time of the word within the audio, in seconds.
	End float64 `json:"end"`
	// The starting time of the word within the audio, in seconds.
	Start float64 `json:"start"`
	// The individual word transcribed from the audio.
	Word string                                                                       `json:"word"`
	JSON accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegmentsWordJSON `json:"-"`
}

// accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegmentsWordJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegmentsWord]
type accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegmentsWordJSON struct {
	End         apijson.Field
	Start       apijson.Field
	Word        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegmentsWord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultSegmentsWordJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultTranscriptionInfo struct {
	// The total duration of the original audio file, in seconds.
	Duration float64 `json:"duration"`
	// The duration of the audio after applying Voice Activity Detection (VAD) to
	// remove silent or irrelevant sections, in seconds.
	DurationAfterVad float64 `json:"duration_after_vad"`
	// The language of the audio being transcribed or translated.
	Language string `json:"language"`
	// The confidence level or probability of the detected language being accurate,
	// represented as a decimal between 0 and 1.
	LanguageProbability float64                                                                           `json:"language_probability"`
	JSON                accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultTranscriptionInfoJSON `json:"-"`
}

// accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultTranscriptionInfoJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultTranscriptionInfo]
type accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultTranscriptionInfoJSON struct {
	Duration            apijson.Field
	DurationAfterVad    apijson.Field
	Language            apijson.Field
	LanguageProbability apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultTranscriptionInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfOpenAIExecuteWhisperLargeV3TurboResponseResultTranscriptionInfoJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfOpenAIExecuteWhisperTinyEnResponse struct {
	Result  AccountAIRunCfOpenAIExecuteWhisperTinyEnResponseResult `json:"result"`
	Success bool                                                   `json:"success"`
	JSON    accountAIRunCfOpenAIExecuteWhisperTinyEnResponseJSON   `json:"-"`
}

// accountAIRunCfOpenAIExecuteWhisperTinyEnResponseJSON contains the JSON metadata
// for the struct [AccountAIRunCfOpenAIExecuteWhisperTinyEnResponse]
type accountAIRunCfOpenAIExecuteWhisperTinyEnResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfOpenAIExecuteWhisperTinyEnResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfOpenAIExecuteWhisperTinyEnResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfOpenAIExecuteWhisperTinyEnResponseResult struct {
	// The transcription
	Text      string                                                       `json:"text,required"`
	Vtt       string                                                       `json:"vtt"`
	WordCount float64                                                      `json:"word_count"`
	Words     []AccountAIRunCfOpenAIExecuteWhisperTinyEnResponseResultWord `json:"words"`
	JSON      accountAIRunCfOpenAIExecuteWhisperTinyEnResponseResultJSON   `json:"-"`
}

// accountAIRunCfOpenAIExecuteWhisperTinyEnResponseResultJSON contains the JSON
// metadata for the struct [AccountAIRunCfOpenAIExecuteWhisperTinyEnResponseResult]
type accountAIRunCfOpenAIExecuteWhisperTinyEnResponseResultJSON struct {
	Text        apijson.Field
	Vtt         apijson.Field
	WordCount   apijson.Field
	Words       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfOpenAIExecuteWhisperTinyEnResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfOpenAIExecuteWhisperTinyEnResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfOpenAIExecuteWhisperTinyEnResponseResultWord struct {
	// The ending second when the word completes
	End float64 `json:"end"`
	// The second this word begins in the recording
	Start float64                                                        `json:"start"`
	Word  string                                                         `json:"word"`
	JSON  accountAIRunCfOpenAIExecuteWhisperTinyEnResponseResultWordJSON `json:"-"`
}

// accountAIRunCfOpenAIExecuteWhisperTinyEnResponseResultWordJSON contains the JSON
// metadata for the struct
// [AccountAIRunCfOpenAIExecuteWhisperTinyEnResponseResultWord]
type accountAIRunCfOpenAIExecuteWhisperTinyEnResponseResultWordJSON struct {
	End         apijson.Field
	Start       apijson.Field
	Word        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfOpenAIExecuteWhisperTinyEnResponseResultWord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfOpenAIExecuteWhisperTinyEnResponseResultWordJSON) RawJSON() string {
	return r.raw
}

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
	VadFilter param.Field[string] `json:"vad_filter"`
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
