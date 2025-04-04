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

type AccountAIRunCfFacebookExecuteBartLargeCnnResponse struct {
	Result  AccountAIRunCfFacebookExecuteBartLargeCnnResponseResult `json:"result"`
	Success bool                                                    `json:"success"`
	JSON    accountAIRunCfFacebookExecuteBartLargeCnnResponseJSON   `json:"-"`
}

// accountAIRunCfFacebookExecuteBartLargeCnnResponseJSON contains the JSON metadata
// for the struct [AccountAIRunCfFacebookExecuteBartLargeCnnResponse]
type accountAIRunCfFacebookExecuteBartLargeCnnResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfFacebookExecuteBartLargeCnnResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfFacebookExecuteBartLargeCnnResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfFacebookExecuteBartLargeCnnResponseResult struct {
	// The summarized version of the input text
	Summary string                                                      `json:"summary"`
	JSON    accountAIRunCfFacebookExecuteBartLargeCnnResponseResultJSON `json:"-"`
}

// accountAIRunCfFacebookExecuteBartLargeCnnResponseResultJSON contains the JSON
// metadata for the struct
// [AccountAIRunCfFacebookExecuteBartLargeCnnResponseResult]
type accountAIRunCfFacebookExecuteBartLargeCnnResponseResultJSON struct {
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfFacebookExecuteBartLargeCnnResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfFacebookExecuteBartLargeCnnResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfFacebookExecuteDetrResnet50Response struct {
	// An array of detected objects within the input image
	Result  []AccountAIRunCfFacebookExecuteDetrResnet50ResponseResult `json:"result"`
	Success bool                                                      `json:"success"`
	JSON    accountAIRunCfFacebookExecuteDetrResnet50ResponseJSON     `json:"-"`
}

// accountAIRunCfFacebookExecuteDetrResnet50ResponseJSON contains the JSON metadata
// for the struct [AccountAIRunCfFacebookExecuteDetrResnet50Response]
type accountAIRunCfFacebookExecuteDetrResnet50ResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfFacebookExecuteDetrResnet50Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfFacebookExecuteDetrResnet50ResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfFacebookExecuteDetrResnet50ResponseResult struct {
	// Coordinates defining the bounding box around the detected object
	Box AccountAIRunCfFacebookExecuteDetrResnet50ResponseResultBox `json:"box"`
	// The class label or name of the detected object
	Label string `json:"label"`
	// Confidence score indicating the likelihood that the detection is correct
	Score float64                                                     `json:"score"`
	JSON  accountAIRunCfFacebookExecuteDetrResnet50ResponseResultJSON `json:"-"`
}

// accountAIRunCfFacebookExecuteDetrResnet50ResponseResultJSON contains the JSON
// metadata for the struct
// [AccountAIRunCfFacebookExecuteDetrResnet50ResponseResult]
type accountAIRunCfFacebookExecuteDetrResnet50ResponseResultJSON struct {
	Box         apijson.Field
	Label       apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfFacebookExecuteDetrResnet50ResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfFacebookExecuteDetrResnet50ResponseResultJSON) RawJSON() string {
	return r.raw
}

// Coordinates defining the bounding box around the detected object
type AccountAIRunCfFacebookExecuteDetrResnet50ResponseResultBox struct {
	// The x-coordinate of the bottom-right corner of the bounding box
	Xmax float64 `json:"xmax"`
	// The x-coordinate of the top-left corner of the bounding box
	Xmin float64 `json:"xmin"`
	// The y-coordinate of the bottom-right corner of the bounding box
	Ymax float64 `json:"ymax"`
	// The y-coordinate of the top-left corner of the bounding box
	Ymin float64                                                        `json:"ymin"`
	JSON accountAIRunCfFacebookExecuteDetrResnet50ResponseResultBoxJSON `json:"-"`
}

// accountAIRunCfFacebookExecuteDetrResnet50ResponseResultBoxJSON contains the JSON
// metadata for the struct
// [AccountAIRunCfFacebookExecuteDetrResnet50ResponseResultBox]
type accountAIRunCfFacebookExecuteDetrResnet50ResponseResultBoxJSON struct {
	Xmax        apijson.Field
	Xmin        apijson.Field
	Ymax        apijson.Field
	Ymin        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfFacebookExecuteDetrResnet50ResponseResultBox) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfFacebookExecuteDetrResnet50ResponseResultBoxJSON) RawJSON() string {
	return r.raw
}

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
