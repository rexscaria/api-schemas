// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/tidwall/gjson"
)

// AccountQueueConsumerService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountQueueConsumerService] method instead.
type AccountQueueConsumerService struct {
	Options []option.RequestOption
}

// NewAccountQueueConsumerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountQueueConsumerService(opts ...option.RequestOption) (r *AccountQueueConsumerService) {
	r = &AccountQueueConsumerService{}
	r.Options = opts
	return
}

// Creates a new consumer for a Queue
func (r *AccountQueueConsumerService) New(ctx context.Context, accountID string, queueID string, body AccountQueueConsumerNewParams, opts ...option.RequestOption) (res *AccountQueueConsumerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers", accountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates the consumer for a queue, or creates one if it does not exist.
func (r *AccountQueueConsumerService) Update(ctx context.Context, accountID string, queueID string, consumerID string, body AccountQueueConsumerUpdateParams, opts ...option.RequestOption) (res *AccountQueueConsumerUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	if consumerID == "" {
		err = errors.New("missing required consumer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers/%s", accountID, queueID, consumerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns the consumers for a Queue
func (r *AccountQueueConsumerService) List(ctx context.Context, accountID string, queueID string, opts ...option.RequestOption) (res *AccountQueueConsumerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers", accountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes the consumer for a queue.
func (r *AccountQueueConsumerService) Delete(ctx context.Context, accountID string, queueID string, consumerID string, opts ...option.RequestOption) (res *MqAPIV4Success, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	if consumerID == "" {
		err = errors.New("missing required consumer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/consumers/%s", accountID, queueID, consumerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type MqConsumer struct {
	// A Resource identifier.
	ConsumerID string `json:"consumer_id"`
	CreatedOn  string `json:"created_on"`
	// A Resource identifier.
	QueueID string `json:"queue_id"`
	Script  string `json:"script"`
	// This field can have the runtime type of [MqConsumerMqWorkerConsumerSettings],
	// [MqConsumerMqHTTPConsumerSettings].
	Settings interface{}    `json:"settings"`
	Type     MqConsumerType `json:"type"`
	JSON     mqConsumerJSON `json:"-"`
	union    MqConsumerUnion
}

// mqConsumerJSON contains the JSON metadata for the struct [MqConsumer]
type mqConsumerJSON struct {
	ConsumerID  apijson.Field
	CreatedOn   apijson.Field
	QueueID     apijson.Field
	Script      apijson.Field
	Settings    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r mqConsumerJSON) RawJSON() string {
	return r.raw
}

func (r *MqConsumer) UnmarshalJSON(data []byte) (err error) {
	*r = MqConsumer{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MqConsumerUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [MqConsumerMqWorkerConsumer],
// [MqConsumerMqHTTPConsumer].
func (r MqConsumer) AsUnion() MqConsumerUnion {
	return r.union
}

// Union satisfied by [MqConsumerMqWorkerConsumer] or [MqConsumerMqHTTPConsumer].
type MqConsumerUnion interface {
	implementsMqConsumer()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MqConsumerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MqConsumerMqWorkerConsumer{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MqConsumerMqHTTPConsumer{}),
		},
	)
}

type MqConsumerMqWorkerConsumer struct {
	// A Resource identifier.
	ConsumerID string `json:"consumer_id"`
	CreatedOn  string `json:"created_on"`
	// A Resource identifier.
	QueueID  string                             `json:"queue_id"`
	Script   string                             `json:"script"`
	Settings MqConsumerMqWorkerConsumerSettings `json:"settings"`
	Type     MqConsumerMqWorkerConsumerType     `json:"type"`
	JSON     mqConsumerMqWorkerConsumerJSON     `json:"-"`
}

// mqConsumerMqWorkerConsumerJSON contains the JSON metadata for the struct
// [MqConsumerMqWorkerConsumer]
type mqConsumerMqWorkerConsumerJSON struct {
	ConsumerID  apijson.Field
	CreatedOn   apijson.Field
	QueueID     apijson.Field
	Script      apijson.Field
	Settings    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MqConsumerMqWorkerConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mqConsumerMqWorkerConsumerJSON) RawJSON() string {
	return r.raw
}

func (r MqConsumerMqWorkerConsumer) implementsMqConsumer() {}

type MqConsumerMqWorkerConsumerSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize float64 `json:"batch_size"`
	// Maximum number of concurrent consumers that may consume from this Queue. Set to
	// `null` to automatically opt in to the platform's maximum (recommended).
	MaxConcurrency float64 `json:"max_concurrency"`
	// The maximum number of retries
	MaxRetries float64 `json:"max_retries"`
	// The number of milliseconds to wait for a batch to fill up before attempting to
	// deliver it
	MaxWaitTimeMs float64 `json:"max_wait_time_ms"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay float64                                `json:"retry_delay"`
	JSON       mqConsumerMqWorkerConsumerSettingsJSON `json:"-"`
}

// mqConsumerMqWorkerConsumerSettingsJSON contains the JSON metadata for the struct
// [MqConsumerMqWorkerConsumerSettings]
type mqConsumerMqWorkerConsumerSettingsJSON struct {
	BatchSize      apijson.Field
	MaxConcurrency apijson.Field
	MaxRetries     apijson.Field
	MaxWaitTimeMs  apijson.Field
	RetryDelay     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *MqConsumerMqWorkerConsumerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mqConsumerMqWorkerConsumerSettingsJSON) RawJSON() string {
	return r.raw
}

type MqConsumerMqWorkerConsumerType string

const (
	MqConsumerMqWorkerConsumerTypeWorker MqConsumerMqWorkerConsumerType = "worker"
)

func (r MqConsumerMqWorkerConsumerType) IsKnown() bool {
	switch r {
	case MqConsumerMqWorkerConsumerTypeWorker:
		return true
	}
	return false
}

type MqConsumerMqHTTPConsumer struct {
	// A Resource identifier.
	ConsumerID string `json:"consumer_id"`
	CreatedOn  string `json:"created_on"`
	// A Resource identifier.
	QueueID  string                           `json:"queue_id"`
	Settings MqConsumerMqHTTPConsumerSettings `json:"settings"`
	Type     MqConsumerMqHTTPConsumerType     `json:"type"`
	JSON     mqConsumerMqHTTPConsumerJSON     `json:"-"`
}

// mqConsumerMqHTTPConsumerJSON contains the JSON metadata for the struct
// [MqConsumerMqHTTPConsumer]
type mqConsumerMqHTTPConsumerJSON struct {
	ConsumerID  apijson.Field
	CreatedOn   apijson.Field
	QueueID     apijson.Field
	Settings    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MqConsumerMqHTTPConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mqConsumerMqHTTPConsumerJSON) RawJSON() string {
	return r.raw
}

func (r MqConsumerMqHTTPConsumer) implementsMqConsumer() {}

type MqConsumerMqHTTPConsumerSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize float64 `json:"batch_size"`
	// The maximum number of retries
	MaxRetries float64 `json:"max_retries"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay float64 `json:"retry_delay"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs float64                              `json:"visibility_timeout_ms"`
	JSON                mqConsumerMqHTTPConsumerSettingsJSON `json:"-"`
}

// mqConsumerMqHTTPConsumerSettingsJSON contains the JSON metadata for the struct
// [MqConsumerMqHTTPConsumerSettings]
type mqConsumerMqHTTPConsumerSettingsJSON struct {
	BatchSize           apijson.Field
	MaxRetries          apijson.Field
	RetryDelay          apijson.Field
	VisibilityTimeoutMs apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MqConsumerMqHTTPConsumerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mqConsumerMqHTTPConsumerSettingsJSON) RawJSON() string {
	return r.raw
}

type MqConsumerMqHTTPConsumerType string

const (
	MqConsumerMqHTTPConsumerTypeHTTPPull MqConsumerMqHTTPConsumerType = "http_pull"
)

func (r MqConsumerMqHTTPConsumerType) IsKnown() bool {
	switch r {
	case MqConsumerMqHTTPConsumerTypeHTTPPull:
		return true
	}
	return false
}

type MqConsumerType string

const (
	MqConsumerTypeWorker   MqConsumerType = "worker"
	MqConsumerTypeHTTPPull MqConsumerType = "http_pull"
)

func (r MqConsumerType) IsKnown() bool {
	switch r {
	case MqConsumerTypeWorker, MqConsumerTypeHTTPPull:
		return true
	}
	return false
}

type MqConsumerParam struct {
	Script     param.Field[string]         `json:"script"`
	ScriptName param.Field[string]         `json:"script_name"`
	Settings   param.Field[interface{}]    `json:"settings"`
	Type       param.Field[MqConsumerType] `json:"type"`
}

func (r MqConsumerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MqConsumerParam) implementsMqConsumerUnionParam() {}

// Satisfied by [MqConsumerMqWorkerConsumerParam], [MqConsumerMqHTTPConsumerParam],
// [MqConsumerParam].
type MqConsumerUnionParam interface {
	implementsMqConsumerUnionParam()
}

type MqConsumerMqWorkerConsumerParam struct {
	Script     param.Field[string]                                  `json:"script"`
	ScriptName param.Field[string]                                  `json:"script_name"`
	Settings   param.Field[MqConsumerMqWorkerConsumerSettingsParam] `json:"settings"`
	Type       param.Field[MqConsumerMqWorkerConsumerType]          `json:"type"`
}

func (r MqConsumerMqWorkerConsumerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MqConsumerMqWorkerConsumerParam) implementsMqConsumerUnionParam() {}

type MqConsumerMqWorkerConsumerSettingsParam struct {
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// Maximum number of concurrent consumers that may consume from this Queue. Set to
	// `null` to automatically opt in to the platform's maximum (recommended).
	MaxConcurrency param.Field[float64] `json:"max_concurrency"`
	// The maximum number of retries
	MaxRetries param.Field[float64] `json:"max_retries"`
	// The number of milliseconds to wait for a batch to fill up before attempting to
	// deliver it
	MaxWaitTimeMs param.Field[float64] `json:"max_wait_time_ms"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay param.Field[float64] `json:"retry_delay"`
}

func (r MqConsumerMqWorkerConsumerSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MqConsumerMqHTTPConsumerParam struct {
	Settings param.Field[MqConsumerMqHTTPConsumerSettingsParam] `json:"settings"`
	Type     param.Field[MqConsumerMqHTTPConsumerType]          `json:"type"`
}

func (r MqConsumerMqHTTPConsumerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MqConsumerMqHTTPConsumerParam) implementsMqConsumerUnionParam() {}

type MqConsumerMqHTTPConsumerSettingsParam struct {
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// The maximum number of retries
	MaxRetries param.Field[float64] `json:"max_retries"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay param.Field[float64] `json:"retry_delay"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs param.Field[float64] `json:"visibility_timeout_ms"`
}

func (r MqConsumerMqHTTPConsumerSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountQueueConsumerNewResponse struct {
	Result MqConsumer                          `json:"result"`
	JSON   accountQueueConsumerNewResponseJSON `json:"-"`
	MqAPIV4Success
}

// accountQueueConsumerNewResponseJSON contains the JSON metadata for the struct
// [AccountQueueConsumerNewResponse]
type accountQueueConsumerNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueConsumerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueConsumerNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountQueueConsumerUpdateResponse struct {
	Result MqConsumer                             `json:"result"`
	JSON   accountQueueConsumerUpdateResponseJSON `json:"-"`
	MqAPIV4Success
}

// accountQueueConsumerUpdateResponseJSON contains the JSON metadata for the struct
// [AccountQueueConsumerUpdateResponse]
type accountQueueConsumerUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueConsumerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueConsumerUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountQueueConsumerListResponse struct {
	Result []MqConsumer                         `json:"result"`
	JSON   accountQueueConsumerListResponseJSON `json:"-"`
	MqAPIV4Success
}

// accountQueueConsumerListResponseJSON contains the JSON metadata for the struct
// [AccountQueueConsumerListResponse]
type accountQueueConsumerListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueConsumerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueConsumerListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountQueueConsumerNewParams struct {
	Body AccountQueueConsumerNewParamsBodyUnion `json:"body"`
}

func (r AccountQueueConsumerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountQueueConsumerNewParamsBody struct {
	DeadLetterQueue param.Field[string]                                `json:"dead_letter_queue"`
	Script          param.Field[string]                                `json:"script"`
	ScriptName      param.Field[string]                                `json:"script_name"`
	Settings        param.Field[interface{}]                           `json:"settings"`
	Type            param.Field[AccountQueueConsumerNewParamsBodyType] `json:"type"`
}

func (r AccountQueueConsumerNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountQueueConsumerNewParamsBody) implementsAccountQueueConsumerNewParamsBodyUnion() {}

// Satisfied by [AccountQueueConsumerNewParamsBodyMqWorkerConsumer],
// [AccountQueueConsumerNewParamsBodyMqHTTPConsumer],
// [AccountQueueConsumerNewParamsBody].
type AccountQueueConsumerNewParamsBodyUnion interface {
	implementsAccountQueueConsumerNewParamsBodyUnion()
}

type AccountQueueConsumerNewParamsBodyMqWorkerConsumer struct {
	DeadLetterQueue param.Field[string]                                                    `json:"dead_letter_queue"`
	Script          param.Field[string]                                                    `json:"script"`
	ScriptName      param.Field[string]                                                    `json:"script_name"`
	Settings        param.Field[AccountQueueConsumerNewParamsBodyMqWorkerConsumerSettings] `json:"settings"`
	Type            param.Field[AccountQueueConsumerNewParamsBodyMqWorkerConsumerType]     `json:"type"`
}

func (r AccountQueueConsumerNewParamsBodyMqWorkerConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountQueueConsumerNewParamsBodyMqWorkerConsumer) implementsAccountQueueConsumerNewParamsBodyUnion() {
}

type AccountQueueConsumerNewParamsBodyMqWorkerConsumerSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// Maximum number of concurrent consumers that may consume from this Queue. Set to
	// `null` to automatically opt in to the platform's maximum (recommended).
	MaxConcurrency param.Field[float64] `json:"max_concurrency"`
	// The maximum number of retries
	MaxRetries param.Field[float64] `json:"max_retries"`
	// The number of milliseconds to wait for a batch to fill up before attempting to
	// deliver it
	MaxWaitTimeMs param.Field[float64] `json:"max_wait_time_ms"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay param.Field[float64] `json:"retry_delay"`
}

func (r AccountQueueConsumerNewParamsBodyMqWorkerConsumerSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountQueueConsumerNewParamsBodyMqWorkerConsumerType string

const (
	AccountQueueConsumerNewParamsBodyMqWorkerConsumerTypeWorker AccountQueueConsumerNewParamsBodyMqWorkerConsumerType = "worker"
)

func (r AccountQueueConsumerNewParamsBodyMqWorkerConsumerType) IsKnown() bool {
	switch r {
	case AccountQueueConsumerNewParamsBodyMqWorkerConsumerTypeWorker:
		return true
	}
	return false
}

type AccountQueueConsumerNewParamsBodyMqHTTPConsumer struct {
	DeadLetterQueue param.Field[string]                                                  `json:"dead_letter_queue"`
	Settings        param.Field[AccountQueueConsumerNewParamsBodyMqHTTPConsumerSettings] `json:"settings"`
	Type            param.Field[AccountQueueConsumerNewParamsBodyMqHTTPConsumerType]     `json:"type"`
}

func (r AccountQueueConsumerNewParamsBodyMqHTTPConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountQueueConsumerNewParamsBodyMqHTTPConsumer) implementsAccountQueueConsumerNewParamsBodyUnion() {
}

type AccountQueueConsumerNewParamsBodyMqHTTPConsumerSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// The maximum number of retries
	MaxRetries param.Field[float64] `json:"max_retries"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay param.Field[float64] `json:"retry_delay"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs param.Field[float64] `json:"visibility_timeout_ms"`
}

func (r AccountQueueConsumerNewParamsBodyMqHTTPConsumerSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountQueueConsumerNewParamsBodyMqHTTPConsumerType string

const (
	AccountQueueConsumerNewParamsBodyMqHTTPConsumerTypeHTTPPull AccountQueueConsumerNewParamsBodyMqHTTPConsumerType = "http_pull"
)

func (r AccountQueueConsumerNewParamsBodyMqHTTPConsumerType) IsKnown() bool {
	switch r {
	case AccountQueueConsumerNewParamsBodyMqHTTPConsumerTypeHTTPPull:
		return true
	}
	return false
}

type AccountQueueConsumerNewParamsBodyType string

const (
	AccountQueueConsumerNewParamsBodyTypeWorker   AccountQueueConsumerNewParamsBodyType = "worker"
	AccountQueueConsumerNewParamsBodyTypeHTTPPull AccountQueueConsumerNewParamsBodyType = "http_pull"
)

func (r AccountQueueConsumerNewParamsBodyType) IsKnown() bool {
	switch r {
	case AccountQueueConsumerNewParamsBodyTypeWorker, AccountQueueConsumerNewParamsBodyTypeHTTPPull:
		return true
	}
	return false
}

type AccountQueueConsumerUpdateParams struct {
	Body AccountQueueConsumerUpdateParamsBodyUnion `json:"body,required"`
}

func (r AccountQueueConsumerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountQueueConsumerUpdateParamsBody struct {
	DeadLetterQueue param.Field[string]                                   `json:"dead_letter_queue"`
	Script          param.Field[string]                                   `json:"script"`
	ScriptName      param.Field[string]                                   `json:"script_name"`
	Settings        param.Field[interface{}]                              `json:"settings"`
	Type            param.Field[AccountQueueConsumerUpdateParamsBodyType] `json:"type"`
}

func (r AccountQueueConsumerUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountQueueConsumerUpdateParamsBody) implementsAccountQueueConsumerUpdateParamsBodyUnion() {}

// Satisfied by [AccountQueueConsumerUpdateParamsBodyMqWorkerConsumer],
// [AccountQueueConsumerUpdateParamsBodyMqHTTPConsumer],
// [AccountQueueConsumerUpdateParamsBody].
type AccountQueueConsumerUpdateParamsBodyUnion interface {
	implementsAccountQueueConsumerUpdateParamsBodyUnion()
}

type AccountQueueConsumerUpdateParamsBodyMqWorkerConsumer struct {
	DeadLetterQueue param.Field[string]                                                       `json:"dead_letter_queue"`
	Script          param.Field[string]                                                       `json:"script"`
	ScriptName      param.Field[string]                                                       `json:"script_name"`
	Settings        param.Field[AccountQueueConsumerUpdateParamsBodyMqWorkerConsumerSettings] `json:"settings"`
	Type            param.Field[AccountQueueConsumerUpdateParamsBodyMqWorkerConsumerType]     `json:"type"`
}

func (r AccountQueueConsumerUpdateParamsBodyMqWorkerConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountQueueConsumerUpdateParamsBodyMqWorkerConsumer) implementsAccountQueueConsumerUpdateParamsBodyUnion() {
}

type AccountQueueConsumerUpdateParamsBodyMqWorkerConsumerSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// Maximum number of concurrent consumers that may consume from this Queue. Set to
	// `null` to automatically opt in to the platform's maximum (recommended).
	MaxConcurrency param.Field[float64] `json:"max_concurrency"`
	// The maximum number of retries
	MaxRetries param.Field[float64] `json:"max_retries"`
	// The number of milliseconds to wait for a batch to fill up before attempting to
	// deliver it
	MaxWaitTimeMs param.Field[float64] `json:"max_wait_time_ms"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay param.Field[float64] `json:"retry_delay"`
}

func (r AccountQueueConsumerUpdateParamsBodyMqWorkerConsumerSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountQueueConsumerUpdateParamsBodyMqWorkerConsumerType string

const (
	AccountQueueConsumerUpdateParamsBodyMqWorkerConsumerTypeWorker AccountQueueConsumerUpdateParamsBodyMqWorkerConsumerType = "worker"
)

func (r AccountQueueConsumerUpdateParamsBodyMqWorkerConsumerType) IsKnown() bool {
	switch r {
	case AccountQueueConsumerUpdateParamsBodyMqWorkerConsumerTypeWorker:
		return true
	}
	return false
}

type AccountQueueConsumerUpdateParamsBodyMqHTTPConsumer struct {
	DeadLetterQueue param.Field[string]                                                     `json:"dead_letter_queue"`
	Settings        param.Field[AccountQueueConsumerUpdateParamsBodyMqHTTPConsumerSettings] `json:"settings"`
	Type            param.Field[AccountQueueConsumerUpdateParamsBodyMqHTTPConsumerType]     `json:"type"`
}

func (r AccountQueueConsumerUpdateParamsBodyMqHTTPConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountQueueConsumerUpdateParamsBodyMqHTTPConsumer) implementsAccountQueueConsumerUpdateParamsBodyUnion() {
}

type AccountQueueConsumerUpdateParamsBodyMqHTTPConsumerSettings struct {
	// The maximum number of messages to include in a batch.
	BatchSize param.Field[float64] `json:"batch_size"`
	// The maximum number of retries
	MaxRetries param.Field[float64] `json:"max_retries"`
	// The number of seconds to delay before making the message available for another
	// attempt.
	RetryDelay param.Field[float64] `json:"retry_delay"`
	// The number of milliseconds that a message is exclusively leased. After the
	// timeout, the message becomes available for another attempt.
	VisibilityTimeoutMs param.Field[float64] `json:"visibility_timeout_ms"`
}

func (r AccountQueueConsumerUpdateParamsBodyMqHTTPConsumerSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountQueueConsumerUpdateParamsBodyMqHTTPConsumerType string

const (
	AccountQueueConsumerUpdateParamsBodyMqHTTPConsumerTypeHTTPPull AccountQueueConsumerUpdateParamsBodyMqHTTPConsumerType = "http_pull"
)

func (r AccountQueueConsumerUpdateParamsBodyMqHTTPConsumerType) IsKnown() bool {
	switch r {
	case AccountQueueConsumerUpdateParamsBodyMqHTTPConsumerTypeHTTPPull:
		return true
	}
	return false
}

type AccountQueueConsumerUpdateParamsBodyType string

const (
	AccountQueueConsumerUpdateParamsBodyTypeWorker   AccountQueueConsumerUpdateParamsBodyType = "worker"
	AccountQueueConsumerUpdateParamsBodyTypeHTTPPull AccountQueueConsumerUpdateParamsBodyType = "http_pull"
)

func (r AccountQueueConsumerUpdateParamsBodyType) IsKnown() bool {
	switch r {
	case AccountQueueConsumerUpdateParamsBodyTypeWorker, AccountQueueConsumerUpdateParamsBodyTypeHTTPPull:
		return true
	}
	return false
}
