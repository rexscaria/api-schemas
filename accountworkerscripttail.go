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

// AccountWorkerScriptTailService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerScriptTailService] method instead.
type AccountWorkerScriptTailService struct {
	Options []option.RequestOption
}

// NewAccountWorkerScriptTailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptTailService(opts ...option.RequestOption) (r *AccountWorkerScriptTailService) {
	r = &AccountWorkerScriptTailService{}
	r.Options = opts
	return
}

// Get list of tails currently deployed on a Worker.
func (r *AccountWorkerScriptTailService) List(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *TailResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a tail from a Worker.
func (r *AccountWorkerScriptTailService) Delete(ctx context.Context, accountID string, scriptName string, id string, body AccountWorkerScriptTailDeleteParams, opts ...option.RequestOption) (res *CommonResponseWorkers, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails/%s", accountID, scriptName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Starts a tail that receives logs and exception from a Worker.
func (r *AccountWorkerScriptTailService) Start(ctx context.Context, accountID string, scriptName string, body AccountWorkerScriptTailStartParams, opts ...option.RequestOption) (res *TailResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CommonResponseWorkers struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// Whether the API call was successful
	Success CommonResponseWorkersSuccess `json:"success,required"`
	JSON    commonResponseWorkersJSON    `json:"-"`
}

// commonResponseWorkersJSON contains the JSON metadata for the struct
// [CommonResponseWorkers]
type commonResponseWorkersJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommonResponseWorkers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commonResponseWorkersJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CommonResponseWorkersSuccess bool

const (
	CommonResponseWorkersSuccessTrue CommonResponseWorkersSuccess = true
)

func (r CommonResponseWorkersSuccess) IsKnown() bool {
	switch r {
	case CommonResponseWorkersSuccessTrue:
		return true
	}
	return false
}

type CommonResponseWorkersParam struct {
	Errors   param.Field[[]WorkersMessagesParam] `json:"errors,required"`
	Messages param.Field[[]WorkersMessagesParam] `json:"messages,required"`
	// Whether the API call was successful
	Success param.Field[CommonResponseWorkersSuccess] `json:"success,required"`
}

func (r CommonResponseWorkersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TailResponse struct {
	Result TailResponseResult `json:"result"`
	JSON   tailResponseJSON   `json:"-"`
	CommonResponseWorkers
}

// tailResponseJSON contains the JSON metadata for the struct [TailResponse]
type tailResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TailResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tailResponseJSON) RawJSON() string {
	return r.raw
}

type TailResponseResult struct {
	ID        string                 `json:"id"`
	ExpiresAt string                 `json:"expires_at"`
	URL       string                 `json:"url"`
	JSON      tailResponseResultJSON `json:"-"`
}

// tailResponseResultJSON contains the JSON metadata for the struct
// [TailResponseResult]
type tailResponseResultJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TailResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tailResponseResultJSON) RawJSON() string {
	return r.raw
}

type WorkersMessages struct {
	Code    int64               `json:"code,required"`
	Message string              `json:"message,required"`
	JSON    workersMessagesJSON `json:"-"`
}

// workersMessagesJSON contains the JSON metadata for the struct [WorkersMessages]
type workersMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersMessagesJSON) RawJSON() string {
	return r.raw
}

type WorkersMessagesParam struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkersMessagesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptTailDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountWorkerScriptTailDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountWorkerScriptTailStartParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountWorkerScriptTailStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
