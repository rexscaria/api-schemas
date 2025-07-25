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
func (r *AccountWorkerScriptTailService) List(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerScriptTailListResponse, err error) {
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
func (r *AccountWorkerScriptTailService) Delete(ctx context.Context, accountID string, scriptName string, id string, opts ...option.RequestOption) (res *CommonResponseWorkers, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Starts a tail that receives logs and exception from a Worker.
func (r *AccountWorkerScriptTailService) Start(ctx context.Context, accountID string, scriptName string, body AccountWorkerScriptTailStartParams, opts ...option.RequestOption) (res *AccountWorkerScriptTailStartResponse, err error) {
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
	// Whether the API call was successful.
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

// Whether the API call was successful.
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

type WorkersMessages struct {
	Code             int64                 `json:"code,required"`
	Message          string                `json:"message,required"`
	DocumentationURL string                `json:"documentation_url"`
	Source           WorkersMessagesSource `json:"source"`
	JSON             workersMessagesJSON   `json:"-"`
}

// workersMessagesJSON contains the JSON metadata for the struct [WorkersMessages]
type workersMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WorkersMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersMessagesJSON) RawJSON() string {
	return r.raw
}

type WorkersMessagesSource struct {
	Pointer string                    `json:"pointer"`
	JSON    workersMessagesSourceJSON `json:"-"`
}

// workersMessagesSourceJSON contains the JSON metadata for the struct
// [WorkersMessagesSource]
type workersMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type WorkersMessagesParam struct {
	Code             param.Field[int64]                      `json:"code,required"`
	Message          param.Field[string]                     `json:"message,required"`
	DocumentationURL param.Field[string]                     `json:"documentation_url"`
	Source           param.Field[WorkersMessagesSourceParam] `json:"source"`
}

func (r WorkersMessagesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersMessagesSourceParam struct {
	Pointer param.Field[string] `json:"pointer"`
}

func (r WorkersMessagesSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptTailListResponse struct {
	Errors   []WorkersMessages                         `json:"errors,required"`
	Messages []WorkersMessages                         `json:"messages,required"`
	Result   AccountWorkerScriptTailListResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerScriptTailListResponseSuccess `json:"success,required"`
	JSON    accountWorkerScriptTailListResponseJSON    `json:"-"`
}

// accountWorkerScriptTailListResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptTailListResponse]
type accountWorkerScriptTailListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptTailListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptTailListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptTailListResponseResult struct {
	// Identifier.
	ID        string                                        `json:"id,required"`
	ExpiresAt string                                        `json:"expires_at,required"`
	URL       string                                        `json:"url,required"`
	JSON      accountWorkerScriptTailListResponseResultJSON `json:"-"`
}

// accountWorkerScriptTailListResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkerScriptTailListResponseResult]
type accountWorkerScriptTailListResponseResultJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptTailListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptTailListResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerScriptTailListResponseSuccess bool

const (
	AccountWorkerScriptTailListResponseSuccessTrue AccountWorkerScriptTailListResponseSuccess = true
)

func (r AccountWorkerScriptTailListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerScriptTailListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptTailStartResponse struct {
	Errors   []WorkersMessages                          `json:"errors,required"`
	Messages []WorkersMessages                          `json:"messages,required"`
	Result   AccountWorkerScriptTailStartResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerScriptTailStartResponseSuccess `json:"success,required"`
	JSON    accountWorkerScriptTailStartResponseJSON    `json:"-"`
}

// accountWorkerScriptTailStartResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptTailStartResponse]
type accountWorkerScriptTailStartResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptTailStartResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptTailStartResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptTailStartResponseResult struct {
	// Identifier.
	ID        string                                         `json:"id,required"`
	ExpiresAt string                                         `json:"expires_at,required"`
	URL       string                                         `json:"url,required"`
	JSON      accountWorkerScriptTailStartResponseResultJSON `json:"-"`
}

// accountWorkerScriptTailStartResponseResultJSON contains the JSON metadata for
// the struct [AccountWorkerScriptTailStartResponseResult]
type accountWorkerScriptTailStartResponseResultJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptTailStartResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptTailStartResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerScriptTailStartResponseSuccess bool

const (
	AccountWorkerScriptTailStartResponseSuccessTrue AccountWorkerScriptTailStartResponseSuccess = true
)

func (r AccountWorkerScriptTailStartResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerScriptTailStartResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptTailStartParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountWorkerScriptTailStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
