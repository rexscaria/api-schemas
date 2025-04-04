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
	Result MqQueue                              `json:"result"`
	JSON   accountQueuePurgeExecuteResponseJSON `json:"-"`
	MqAPIV4Success
}

// accountQueuePurgeExecuteResponseJSON contains the JSON metadata for the struct
// [AccountQueuePurgeExecuteResponse]
type accountQueuePurgeExecuteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueuePurgeExecuteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueuePurgeExecuteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountQueuePurgeStatusResponse struct {
	Result AccountQueuePurgeStatusResponseResult `json:"result"`
	JSON   accountQueuePurgeStatusResponseJSON   `json:"-"`
	MqAPIV4Success
}

// accountQueuePurgeStatusResponseJSON contains the JSON metadata for the struct
// [AccountQueuePurgeStatusResponse]
type accountQueuePurgeStatusResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueuePurgeStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueuePurgeStatusResponseJSON) RawJSON() string {
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

type AccountQueuePurgeExecuteParams struct {
	// Confimation that all messages will be deleted permanently.
	DeleteMessagesPermanently param.Field[bool] `json:"delete_messages_permanently"`
}

func (r AccountQueuePurgeExecuteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
