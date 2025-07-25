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

type UsageModelResponse struct {
	Errors   []WorkersMessages        `json:"errors,required"`
	Messages []WorkersMessages        `json:"messages,required"`
	Result   UsageModelResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success UsageModelResponseSuccess `json:"success,required"`
	JSON    usageModelResponseJSON    `json:"-"`
}

// usageModelResponseJSON contains the JSON metadata for the struct
// [UsageModelResponse]
type usageModelResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageModelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageModelResponseJSON) RawJSON() string {
	return r.raw
}

type UsageModelResponseResult struct {
	// Usage model for the Worker invocations.
	UsageModel UsageModel                   `json:"usage_model"`
	JSON       usageModelResponseResultJSON `json:"-"`
}

// usageModelResponseResultJSON contains the JSON metadata for the struct
// [UsageModelResponseResult]
type usageModelResponseResultJSON struct {
	UsageModel  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageModelResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageModelResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UsageModelResponseSuccess bool

const (
	UsageModelResponseSuccessTrue UsageModelResponseSuccess = true
)

func (r UsageModelResponseSuccess) IsKnown() bool {
	switch r {
	case UsageModelResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptUsageModelUpdateParams struct {
	// Usage model for the Worker invocations.
	UsageModel param.Field[UsageModel] `json:"usage_model"`
}

func (r AccountWorkerScriptUsageModelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
