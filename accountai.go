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

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAIService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIService] method instead.
type AccountAIService struct {
	Options   []option.RequestOption
	Authors   *AccountAIAuthorService
	Finetunes *AccountAIFinetuneService
	Models    *AccountAIModelService
	Run       *AccountAIRunService
	Tasks     *AccountAITaskService
}

// NewAccountAIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountAIService(opts ...option.RequestOption) (r *AccountAIService) {
	r = &AccountAIService{}
	r.Options = opts
	r.Authors = NewAccountAIAuthorService(opts...)
	r.Finetunes = NewAccountAIFinetuneService(opts...)
	r.Models = NewAccountAIModelService(opts...)
	r.Run = NewAccountAIRunService(opts...)
	r.Tasks = NewAccountAITaskService(opts...)
	return
}

// Convert Files into Markdown
func (r *AccountAIService) ConvertToMarkdown(ctx context.Context, accountID string, body AccountAIConvertToMarkdownParams, opts ...option.RequestOption) (res *AccountAIConvertToMarkdownResponse, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "applications/json")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/tomarkdown", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountAIConvertToMarkdownResponse struct {
	Result  []AccountAIConvertToMarkdownResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    accountAIConvertToMarkdownResponseJSON     `json:"-"`
}

// accountAIConvertToMarkdownResponseJSON contains the JSON metadata for the struct
// [AccountAIConvertToMarkdownResponse]
type accountAIConvertToMarkdownResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIConvertToMarkdownResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIConvertToMarkdownResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIConvertToMarkdownResponseResult struct {
	Data     string                                       `json:"data,required"`
	Format   string                                       `json:"format,required"`
	MimeType string                                       `json:"mimeType,required"`
	Name     string                                       `json:"name,required"`
	Tokens   string                                       `json:"tokens,required"`
	JSON     accountAIConvertToMarkdownResponseResultJSON `json:"-"`
}

// accountAIConvertToMarkdownResponseResultJSON contains the JSON metadata for the
// struct [AccountAIConvertToMarkdownResponseResult]
type accountAIConvertToMarkdownResponseResultJSON struct {
	Data        apijson.Field
	Format      apijson.Field
	MimeType    apijson.Field
	Name        apijson.Field
	Tokens      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIConvertToMarkdownResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIConvertToMarkdownResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIConvertToMarkdownParams struct {
	Body io.Reader `json:"body" format:"binary"`
}

func (r AccountAIConvertToMarkdownParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
