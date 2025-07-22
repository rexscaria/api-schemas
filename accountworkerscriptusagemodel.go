// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountWorkerScriptUsageModelService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerScriptUsageModelService] method instead.
type AccountWorkerScriptUsageModelService struct {
	Options []option.RequestOption
}

// NewAccountWorkerScriptUsageModelService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptUsageModelService(opts ...option.RequestOption) (r *AccountWorkerScriptUsageModelService) {
	r = &AccountWorkerScriptUsageModelService{}
	r.Options = opts
	return
}

// Updates the Usage Model for a given Worker. Requires a Workers Paid
// subscription.
func (r *AccountWorkerScriptUsageModelService) Update(ctx context.Context, accountID string, scriptName string, body AccountWorkerScriptUsageModelUpdateParams, opts ...option.RequestOption) (res *UsageModelResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/usage-model", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches the Usage Model for a given Worker.
func (r *AccountWorkerScriptUsageModelService) Get(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *UsageModelResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/usage-model", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UsageModelObject struct {
	UsageModel string               `json:"usage_model"`
	JSON       usageModelObjectJSON `json:"-"`
}

// usageModelObjectJSON contains the JSON metadata for the struct
// [UsageModelObject]
type usageModelObjectJSON struct {
	UsageModel  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageModelObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageModelObjectJSON) RawJSON() string {
	return r.raw
}

type UsageModelObjectParam struct {
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r UsageModelObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UsageModelResponse struct {
	Result UsageModelObject       `json:"result"`
	JSON   usageModelResponseJSON `json:"-"`
	CommonResponseWorkers
}

// usageModelResponseJSON contains the JSON metadata for the struct
// [UsageModelResponse]
type usageModelResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageModelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageModelResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptUsageModelUpdateParams struct {
	UsageModelObject UsageModelObjectParam `json:"usage_model_object,required"`
}

func (r AccountWorkerScriptUsageModelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UsageModelObject)
}
