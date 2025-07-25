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

// AccountAIRunCfFacebookService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfFacebookService] method instead.
type AccountAIRunCfFacebookService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfFacebookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfFacebookService(opts ...option.RequestOption) (r *AccountAIRunCfFacebookService) {
	r = &AccountAIRunCfFacebookService{}
	r.Options = opts
	return
}

// Execute @cf/facebook/bart-large-cnn model.
func (r *AccountAIRunCfFacebookService) ExecuteBartLargeCnn(ctx context.Context, accountID string, params AccountAIRunCfFacebookExecuteBartLargeCnnParams, opts ...option.RequestOption) (res *AccountAIRunCfFacebookExecuteBartLargeCnnResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/facebook/bart-large-cnn", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/facebook/detr-resnet-50 model.
func (r *AccountAIRunCfFacebookService) ExecuteDetrResnet50(ctx context.Context, accountID string, params AccountAIRunCfFacebookExecuteDetrResnet50Params, opts ...option.RequestOption) (res *AccountAIRunCfFacebookExecuteDetrResnet50Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/facebook/detr-resnet-50", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfFacebookExecuteBartLargeCnnResponse = interface{}

type AccountAIRunCfFacebookExecuteDetrResnet50Response = interface{}

type AccountAIRunCfFacebookExecuteBartLargeCnnParams struct {
	// The text that you want the model to summarize
	InputText    param.Field[string] `json:"input_text,required"`
	QueueRequest param.Field[string] `query:"queueRequest"`
	// The maximum length of the generated summary in tokens
	MaxLength param.Field[int64] `json:"max_length"`
}

func (r AccountAIRunCfFacebookExecuteBartLargeCnnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountAIRunCfFacebookExecuteBartLargeCnnParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfFacebookExecuteBartLargeCnnParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfFacebookExecuteDetrResnet50Params struct {
	QueueRequest param.Field[string] `query:"queueRequest"`
	Body         io.Reader           `json:"body" format:"binary"`
}

func (r AccountAIRunCfFacebookExecuteDetrResnet50Params) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [AccountAIRunCfFacebookExecuteDetrResnet50Params]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfFacebookExecuteDetrResnet50Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
