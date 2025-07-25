// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountStreamCaptionService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountStreamCaptionService] method instead.
type AccountStreamCaptionService struct {
	Options []option.RequestOption
}

// NewAccountStreamCaptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamCaptionService(opts ...option.RequestOption) (r *AccountStreamCaptionService) {
	r = &AccountStreamCaptionService{}
	r.Options = opts
	return
}

// Lists the captions or subtitles for provided language.
func (r *AccountStreamCaptionService) Get(ctx context.Context, accountID string, identifier string, language string, opts ...option.RequestOption) (res *LanguageResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if language == "" {
		err = errors.New("missing required language parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s", accountID, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists the available captions or subtitles for a specific video.
func (r *AccountStreamCaptionService) List(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *AccountStreamCaptionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/captions", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Removes the captions or subtitles from a video.
func (r *AccountStreamCaptionService) Delete(ctx context.Context, accountID string, identifier string, language string, opts ...option.RequestOption) (res *AccountStreamCaptionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if language == "" {
		err = errors.New("missing required language parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s", accountID, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Generate captions or subtitles for provided language via AI.
func (r *AccountStreamCaptionService) Generate(ctx context.Context, accountID string, identifier string, language string, opts ...option.RequestOption) (res *LanguageResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if language == "" {
		err = errors.New("missing required language parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s/generate", accountID, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Return WebVTT captions for a provided language.
func (r *AccountStreamCaptionService) GetVtt(ctx context.Context, accountID string, identifier string, language string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/vtt")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if language == "" {
		err = errors.New("missing required language parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s/vtt", accountID, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Uploads the caption or subtitle file to the endpoint for a specific BCP47
// language. One caption or subtitle file per language is allowed.
func (r *AccountStreamCaptionService) Upload(ctx context.Context, accountID string, identifier string, language string, body AccountStreamCaptionUploadParams, opts ...option.RequestOption) (res *LanguageResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if language == "" {
		err = errors.New("missing required language parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s", accountID, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type Captions struct {
	// Whether the caption was generated via AI.
	Generated bool `json:"generated"`
	// The language label displayed in the native language to users.
	Label string `json:"label"`
	// The language tag in BCP 47 format.
	Language string `json:"language"`
	// The status of a generated caption.
	Status CaptionsStatus `json:"status"`
	JSON   captionsJSON   `json:"-"`
}

// captionsJSON contains the JSON metadata for the struct [Captions]
type captionsJSON struct {
	Generated   apijson.Field
	Label       apijson.Field
	Language    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Captions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionsJSON) RawJSON() string {
	return r.raw
}

// The status of a generated caption.
type CaptionsStatus string

const (
	CaptionsStatusReady      CaptionsStatus = "ready"
	CaptionsStatusInprogress CaptionsStatus = "inprogress"
	CaptionsStatusError      CaptionsStatus = "error"
)

func (r CaptionsStatus) IsKnown() bool {
	switch r {
	case CaptionsStatusReady, CaptionsStatusInprogress, CaptionsStatusError:
		return true
	}
	return false
}

type LanguageResponseSingle struct {
	Errors   []StreamMessages `json:"errors,required"`
	Messages []StreamMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success LanguageResponseSingleSuccess `json:"success,required"`
	Result  Captions                      `json:"result"`
	JSON    languageResponseSingleJSON    `json:"-"`
}

// languageResponseSingleJSON contains the JSON metadata for the struct
// [LanguageResponseSingle]
type languageResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LanguageResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r languageResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LanguageResponseSingleSuccess bool

const (
	LanguageResponseSingleSuccessTrue LanguageResponseSingleSuccess = true
)

func (r LanguageResponseSingleSuccess) IsKnown() bool {
	switch r {
	case LanguageResponseSingleSuccessTrue:
		return true
	}
	return false
}

type StreamMessages struct {
	Code             int64                `json:"code,required"`
	Message          string               `json:"message,required"`
	DocumentationURL string               `json:"documentation_url"`
	Source           StreamMessagesSource `json:"source"`
	JSON             streamMessagesJSON   `json:"-"`
}

// streamMessagesJSON contains the JSON metadata for the struct [StreamMessages]
type streamMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StreamMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamMessagesJSON) RawJSON() string {
	return r.raw
}

type StreamMessagesSource struct {
	Pointer string                   `json:"pointer"`
	JSON    streamMessagesSourceJSON `json:"-"`
}

// streamMessagesSourceJSON contains the JSON metadata for the struct
// [StreamMessagesSource]
type streamMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type AccountStreamCaptionListResponse struct {
	Errors   []StreamMessages `json:"errors,required"`
	Messages []StreamMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountStreamCaptionListResponseSuccess `json:"success,required"`
	Result  []Captions                              `json:"result"`
	JSON    accountStreamCaptionListResponseJSON    `json:"-"`
}

// accountStreamCaptionListResponseJSON contains the JSON metadata for the struct
// [AccountStreamCaptionListResponse]
type accountStreamCaptionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCaptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamCaptionListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStreamCaptionListResponseSuccess bool

const (
	AccountStreamCaptionListResponseSuccessTrue AccountStreamCaptionListResponseSuccess = true
)

func (r AccountStreamCaptionListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStreamCaptionListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStreamCaptionDeleteResponse struct {
	Errors   []StreamMessages `json:"errors,required"`
	Messages []StreamMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountStreamCaptionDeleteResponseSuccess `json:"success,required"`
	Result  string                                    `json:"result"`
	JSON    accountStreamCaptionDeleteResponseJSON    `json:"-"`
}

// accountStreamCaptionDeleteResponseJSON contains the JSON metadata for the struct
// [AccountStreamCaptionDeleteResponse]
type accountStreamCaptionDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCaptionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamCaptionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStreamCaptionDeleteResponseSuccess bool

const (
	AccountStreamCaptionDeleteResponseSuccessTrue AccountStreamCaptionDeleteResponseSuccess = true
)

func (r AccountStreamCaptionDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStreamCaptionDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStreamCaptionUploadParams struct {
	// The WebVTT file containing the caption or subtitle content.
	File param.Field[string] `json:"file,required"`
}

func (r AccountStreamCaptionUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
