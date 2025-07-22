// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/stainless-sdks/cf-rex-go/shared"
	"github.com/tidwall/gjson"
)

// AccountAIRunCfMyshellAIService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfMyshellAIService] method instead.
type AccountAIRunCfMyshellAIService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfMyshellAIService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfMyshellAIService(opts ...option.RequestOption) (r *AccountAIRunCfMyshellAIService) {
	r = &AccountAIRunCfMyshellAIService{}
	r.Options = opts
	return
}

// Execute @cf/myshell-ai/melotts model.
func (r *AccountAIRunCfMyshellAIService) ExecuteMelotts(ctx context.Context, accountID string, params AccountAIRunCfMyshellAIExecuteMelottsParams, opts ...option.RequestOption) (res *AccountAIRunCfMyshellAIExecuteMelottsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/myshell-ai/melotts", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfMyshellAIExecuteMelottsResponse struct {
	// The generated audio in MP3 format
	Result  AccountAIRunCfMyshellAIExecuteMelottsResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                     `json:"success"`
	JSON    accountAIRunCfMyshellAIExecuteMelottsResponseJSON        `json:"-"`
}

// accountAIRunCfMyshellAIExecuteMelottsResponseJSON contains the JSON metadata for
// the struct [AccountAIRunCfMyshellAIExecuteMelottsResponse]
type accountAIRunCfMyshellAIExecuteMelottsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMyshellAIExecuteMelottsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMyshellAIExecuteMelottsResponseJSON) RawJSON() string {
	return r.raw
}

// The generated audio in MP3 format
//
// Union satisfied by [AccountAIRunCfMyshellAIExecuteMelottsResponseResultAudio] or
// [shared.UnionString].
type AccountAIRunCfMyshellAIExecuteMelottsResponseResultUnion interface {
	ImplementsAccountAIRunCfMyshellAIExecuteMelottsResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMyshellAIExecuteMelottsResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMyshellAIExecuteMelottsResponseResultAudio{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMyshellAIExecuteMelottsResponseResultAudio struct {
	// The generated audio in MP3 format, base64-encoded
	Audio string                                                       `json:"audio"`
	JSON  accountAIRunCfMyshellAIExecuteMelottsResponseResultAudioJSON `json:"-"`
}

// accountAIRunCfMyshellAIExecuteMelottsResponseResultAudioJSON contains the JSON
// metadata for the struct
// [AccountAIRunCfMyshellAIExecuteMelottsResponseResultAudio]
type accountAIRunCfMyshellAIExecuteMelottsResponseResultAudioJSON struct {
	Audio       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMyshellAIExecuteMelottsResponseResultAudio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMyshellAIExecuteMelottsResponseResultAudioJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMyshellAIExecuteMelottsResponseResultAudio) ImplementsAccountAIRunCfMyshellAIExecuteMelottsResponseResultUnion() {
}

type AccountAIRunCfMyshellAIExecuteMelottsParams struct {
	// A text description of the audio you want to generate
	Prompt       param.Field[string] `json:"prompt,required"`
	QueueRequest param.Field[string] `query:"queueRequest"`
	// The speech language (e.g., 'en' for English, 'fr' for French). Defaults to 'en'
	// if not specified
	Lang param.Field[string] `json:"lang"`
}

func (r AccountAIRunCfMyshellAIExecuteMelottsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountAIRunCfMyshellAIExecuteMelottsParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfMyshellAIExecuteMelottsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
