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

// AccountQueuePurgeService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountQueuePurgeService] method instead.
type AccountQueuePurgeService struct {
	Options []option.RequestOption
}

// NewAccountQueuePurgeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountQueuePurgeService(opts ...option.RequestOption) (r *AccountQueuePurgeService) {
	r = &AccountQueuePurgeService{}
	r.Options = opts
	return
}

// Deletes all messages from the Queue.
func (r *AccountQueuePurgeService) Execute(ctx context.Context, accountID string, queueID string, body AccountQueuePurgeExecuteParams, opts ...option.RequestOption) (res *AccountQueuePurgeExecuteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/purge", accountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details about a Queue's purge status.
func (r *AccountQueuePurgeService) Status(ctx context.Context, accountID string, queueID string, opts ...option.RequestOption) (res *AccountQueuePurgeStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/purge", accountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountQueuePurgeExecuteResponse struct {
	Errors   []AccountQueuePurgeExecuteResponseError `json:"errors"`
	Messages []string                                `json:"messages"`
	Result   MqQueue                                 `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountQueuePurgeExecuteResponseSuccess `json:"success"`
	JSON    accountQueuePurgeExecuteResponseJSON    `json:"-"`
}

// accountQueuePurgeExecuteResponseJSON contains the JSON metadata for the struct
// [AccountQueuePurgeExecuteResponse]
type accountQueuePurgeExecuteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueuePurgeExecuteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueuePurgeExecuteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountQueuePurgeExecuteResponseError struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           AccountQueuePurgeExecuteResponseErrorsSource `json:"source"`
	JSON             accountQueuePurgeExecuteResponseErrorJSON    `json:"-"`
}

// accountQueuePurgeExecuteResponseErrorJSON contains the JSON metadata for the
// struct [AccountQueuePurgeExecuteResponseError]
type accountQueuePurgeExecuteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountQueuePurgeExecuteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueuePurgeExecuteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountQueuePurgeExecuteResponseErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    accountQueuePurgeExecuteResponseErrorsSourceJSON `json:"-"`
}

// accountQueuePurgeExecuteResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountQueuePurgeExecuteResponseErrorsSource]
type accountQueuePurgeExecuteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueuePurgeExecuteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueuePurgeExecuteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountQueuePurgeExecuteResponseSuccess bool

const (
	AccountQueuePurgeExecuteResponseSuccessTrue AccountQueuePurgeExecuteResponseSuccess = true
)

func (r AccountQueuePurgeExecuteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountQueuePurgeExecuteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountQueuePurgeStatusResponse struct {
	Errors   []AccountQueuePurgeStatusResponseError `json:"errors"`
	Messages []string                               `json:"messages"`
	Result   AccountQueuePurgeStatusResponseResult  `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountQueuePurgeStatusResponseSuccess `json:"success"`
	JSON    accountQueuePurgeStatusResponseJSON    `json:"-"`
}

// accountQueuePurgeStatusResponseJSON contains the JSON metadata for the struct
// [AccountQueuePurgeStatusResponse]
type accountQueuePurgeStatusResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueuePurgeStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueuePurgeStatusResponseJSON) RawJSON() string {
	return r.raw
}

type AccountQueuePurgeStatusResponseError struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AccountQueuePurgeStatusResponseErrorsSource `json:"source"`
	JSON             accountQueuePurgeStatusResponseErrorJSON    `json:"-"`
}

// accountQueuePurgeStatusResponseErrorJSON contains the JSON metadata for the
// struct [AccountQueuePurgeStatusResponseError]
type accountQueuePurgeStatusResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountQueuePurgeStatusResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueuePurgeStatusResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountQueuePurgeStatusResponseErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    accountQueuePurgeStatusResponseErrorsSourceJSON `json:"-"`
}

// accountQueuePurgeStatusResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountQueuePurgeStatusResponseErrorsSource]
type accountQueuePurgeStatusResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueuePurgeStatusResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueuePurgeStatusResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountQueuePurgeStatusResponseResult struct {
	// Indicates if the last purge operation completed successfully.
	Completed string `json:"completed"`
	// Timestamp when the last purge operation started.
	StartedAt string                                    `json:"started_at"`
	JSON      accountQueuePurgeStatusResponseResultJSON `json:"-"`
}

// accountQueuePurgeStatusResponseResultJSON contains the JSON metadata for the
// struct [AccountQueuePurgeStatusResponseResult]
type accountQueuePurgeStatusResponseResultJSON struct {
	Completed   apijson.Field
	StartedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueuePurgeStatusResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueuePurgeStatusResponseResultJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountQueuePurgeStatusResponseSuccess bool

const (
	AccountQueuePurgeStatusResponseSuccessTrue AccountQueuePurgeStatusResponseSuccess = true
)

func (r AccountQueuePurgeStatusResponseSuccess) IsKnown() bool {
	switch r {
	case AccountQueuePurgeStatusResponseSuccessTrue:
		return true
	}
	return false
}

type AccountQueuePurgeExecuteParams struct {
	// Confimation that all messages will be deleted permanently.
	DeleteMessagesPermanently param.Field[bool] `json:"delete_messages_permanently"`
}

func (r AccountQueuePurgeExecuteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
