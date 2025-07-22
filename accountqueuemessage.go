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

// AccountQueueMessageService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountQueueMessageService] method instead.
type AccountQueueMessageService struct {
	Options []option.RequestOption
}

// NewAccountQueueMessageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountQueueMessageService(opts ...option.RequestOption) (r *AccountQueueMessageService) {
	r = &AccountQueueMessageService{}
	r.Options = opts
	return
}

// Acknowledge + Retry messages from a Queue
func (r *AccountQueueMessageService) Acknowledge(ctx context.Context, accountID string, queueID string, body AccountQueueMessageAcknowledgeParams, opts ...option.RequestOption) (res *AccountQueueMessageAcknowledgeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/messages/ack", accountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Pull a batch of messages from a Queue
func (r *AccountQueueMessageService) Pull(ctx context.Context, accountID string, queueID string, body AccountQueueMessagePullParams, opts ...option.RequestOption) (res *AccountQueueMessagePullResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/messages/pull", accountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountQueueMessageAcknowledgeResponse struct {
	Result AccountQueueMessageAcknowledgeResponseResult `json:"result"`
	JSON   accountQueueMessageAcknowledgeResponseJSON   `json:"-"`
	MqAPIV4Success
}

// accountQueueMessageAcknowledgeResponseJSON contains the JSON metadata for the
// struct [AccountQueueMessageAcknowledgeResponse]
type accountQueueMessageAcknowledgeResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueMessageAcknowledgeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueMessageAcknowledgeResponseJSON) RawJSON() string {
	return r.raw
}

type AccountQueueMessageAcknowledgeResponseResult struct {
	// The number of messages that were succesfully acknowledged.
	AckCount float64 `json:"ackCount"`
	// The number of messages that were succesfully retried.
	RetryCount float64                                          `json:"retryCount"`
	Warnings   []string                                         `json:"warnings"`
	JSON       accountQueueMessageAcknowledgeResponseResultJSON `json:"-"`
}

// accountQueueMessageAcknowledgeResponseResultJSON contains the JSON metadata for
// the struct [AccountQueueMessageAcknowledgeResponseResult]
type accountQueueMessageAcknowledgeResponseResultJSON struct {
	AckCount    apijson.Field
	RetryCount  apijson.Field
	Warnings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueMessageAcknowledgeResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueMessageAcknowledgeResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountQueueMessagePullResponse struct {
	Result []AccountQueueMessagePullResponseResult `json:"result"`
	JSON   accountQueueMessagePullResponseJSON     `json:"-"`
	MqAPIV4Success
}

// accountQueueMessagePullResponseJSON contains the JSON metadata for the struct
// [AccountQueueMessagePullResponse]
type accountQueueMessagePullResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueMessagePullResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueMessagePullResponseJSON) RawJSON() string {
	return r.raw
}

type AccountQueueMessagePullResponseResult struct {
	ID       string  `json:"id"`
	Attempts float64 `json:"attempts"`
	Body     string  `json:"body"`
	// An ID that represents an "in-flight" message that has been pulled from a Queue.
	// You must hold on to this ID and use it to acknowledge this message.
	LeaseID     string                                    `json:"lease_id"`
	Metadata    interface{}                               `json:"metadata"`
	TimestampMs float64                                   `json:"timestamp_ms"`
	JSON        accountQueueMessagePullResponseResultJSON `json:"-"`
}

// accountQueueMessagePullResponseResultJSON contains the JSON metadata for the
// struct [AccountQueueMessagePullResponseResult]
type accountQueueMessagePullResponseResultJSON struct {
	ID          apijson.Field
	Attempts    apijson.Field
	Body        apijson.Field
	LeaseID     apijson.Field
	Metadata    apijson.Field
	TimestampMs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueMessagePullResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueMessagePullResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountQueueMessageAcknowledgeParams struct {
	Acks    param.Field[[]AccountQueueMessageAcknowledgeParamsAck]   `json:"acks"`
	Retries param.Field[[]AccountQueueMessageAcknowledgeParamsRetry] `json:"retries"`
}

func (r AccountQueueMessageAcknowledgeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountQueueMessageAcknowledgeParamsAck struct {
	// An ID that represents an "in-flight" message that has been pulled from a Queue.
	// You must hold on to this ID and use it to acknowledge this message.
	LeaseID param.Field[string] `json:"lease_id"`
}

func (r AccountQueueMessageAcknowledgeParamsAck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountQueueMessageAcknowledgeParamsRetry struct {
	// The number of seconds to delay before making the message available for another
	// attempt.
	DelaySeconds param.Field[float64] `json:"delay_seconds"`
	// An ID that represents an "in-flight" message that has been pulled from a Queue.
	// You must hold on to this ID and use it to acknowledge this message.
	LeaseID param.Field[string] `json:"lease_id"`
}

func (r AccountQueueMessageAcknowledgeParamsRetry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountQueueMessagePullParams struct {
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs param.Field[float64] `json:"visibility_timeout_ms"`
}

func (r AccountQueueMessagePullParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
