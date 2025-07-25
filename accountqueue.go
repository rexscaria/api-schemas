// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountQueueService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountQueueService] method instead.
type AccountQueueService struct {
	Options   []option.RequestOption
	Consumers *AccountQueueConsumerService
	Messages  *AccountQueueMessageService
	Purge     *AccountQueuePurgeService
}

// NewAccountQueueService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountQueueService(opts ...option.RequestOption) (r *AccountQueueService) {
	r = &AccountQueueService{}
	r.Options = opts
	r.Consumers = NewAccountQueueConsumerService(opts...)
	r.Messages = NewAccountQueueMessageService(opts...)
	r.Purge = NewAccountQueuePurgeService(opts...)
	return
}

// Create a new queue
func (r *AccountQueueService) New(ctx context.Context, accountID string, body AccountQueueNewParams, opts ...option.RequestOption) (res *AccountQueueNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details about a specific queue.
func (r *AccountQueueService) Get(ctx context.Context, accountID string, queueID string, opts ...option.RequestOption) (res *AccountQueueGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s", accountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a Queue. Note that this endpoint does not support partial updates. If
// successful, the Queue's configuration is overwritten with the supplied
// configuration.
func (r *AccountQueueService) Update(ctx context.Context, accountID string, queueID string, body AccountQueueUpdateParams, opts ...option.RequestOption) (res *AccountQueueUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s", accountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns the queues owned by an account.
func (r *AccountQueueService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountQueueListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a queue
func (r *AccountQueueService) Delete(ctx context.Context, accountID string, queueID string, opts ...option.RequestOption) (res *MqAPIV4Success, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s", accountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Updates a Queue.
func (r *AccountQueueService) UpdatePartial(ctx context.Context, accountID string, queueID string, body AccountQueueUpdatePartialParams, opts ...option.RequestOption) (res *AccountQueueUpdatePartialResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s", accountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type MqAPIV4Success struct {
	Errors   []MqApiv4SuccessError `json:"errors"`
	Messages []string              `json:"messages"`
	// Indicates if the API call was successful or not.
	Success MqAPIV4SuccessSuccess `json:"success"`
	JSON    mqApiv4SuccessJSON    `json:"-"`
}

// mqApiv4SuccessJSON contains the JSON metadata for the struct [MqAPIV4Success]
type mqApiv4SuccessJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MqAPIV4Success) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mqApiv4SuccessJSON) RawJSON() string {
	return r.raw
}

type MqApiv4SuccessError struct {
	Code             int64                      `json:"code,required"`
	Message          string                     `json:"message,required"`
	DocumentationURL string                     `json:"documentation_url"`
	Source           MqAPIV4SuccessErrorsSource `json:"source"`
	JSON             mqApiv4SuccessErrorJSON    `json:"-"`
}

// mqApiv4SuccessErrorJSON contains the JSON metadata for the struct
// [MqApiv4SuccessError]
type mqApiv4SuccessErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MqApiv4SuccessError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mqApiv4SuccessErrorJSON) RawJSON() string {
	return r.raw
}

type MqAPIV4SuccessErrorsSource struct {
	Pointer string                         `json:"pointer"`
	JSON    mqApiv4SuccessErrorsSourceJSON `json:"-"`
}

// mqApiv4SuccessErrorsSourceJSON contains the JSON metadata for the struct
// [MqAPIV4SuccessErrorsSource]
type mqApiv4SuccessErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MqAPIV4SuccessErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mqApiv4SuccessErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type MqAPIV4SuccessSuccess bool

const (
	MqAPIV4SuccessSuccessTrue MqAPIV4SuccessSuccess = true
)

func (r MqAPIV4SuccessSuccess) IsKnown() bool {
	switch r {
	case MqAPIV4SuccessSuccessTrue:
		return true
	}
	return false
}

type MqQueue struct {
	Consumers           []MqConsumer      `json:"consumers"`
	ConsumersTotalCount float64           `json:"consumers_total_count"`
	CreatedOn           string            `json:"created_on"`
	ModifiedOn          string            `json:"modified_on"`
	Producers           []MqQueueProducer `json:"producers"`
	ProducersTotalCount float64           `json:"producers_total_count"`
	QueueID             string            `json:"queue_id"`
	QueueName           string            `json:"queue_name"`
	Settings            MqQueueSettings   `json:"settings"`
	JSON                mqQueueJSON       `json:"-"`
}

// mqQueueJSON contains the JSON metadata for the struct [MqQueue]
type mqQueueJSON struct {
	Consumers           apijson.Field
	ConsumersTotalCount apijson.Field
	CreatedOn           apijson.Field
	ModifiedOn          apijson.Field
	Producers           apijson.Field
	ProducersTotalCount apijson.Field
	QueueID             apijson.Field
	QueueName           apijson.Field
	Settings            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MqQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mqQueueJSON) RawJSON() string {
	return r.raw
}

type MqQueueProducer struct {
	BucketName string               `json:"bucket_name"`
	Script     string               `json:"script"`
	Type       MqQueueProducersType `json:"type"`
	JSON       mqQueueProducerJSON  `json:"-"`
	union      MqQueueProducersUnion
}

// mqQueueProducerJSON contains the JSON metadata for the struct [MqQueueProducer]
type mqQueueProducerJSON struct {
	BucketName  apijson.Field
	Script      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r mqQueueProducerJSON) RawJSON() string {
	return r.raw
}

func (r *MqQueueProducer) UnmarshalJSON(data []byte) (err error) {
	*r = MqQueueProducer{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MqQueueProducersUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [MqQueueProducersMqWorkerProducer],
// [MqQueueProducersMqR2Producer].
func (r MqQueueProducer) AsUnion() MqQueueProducersUnion {
	return r.union
}

// Union satisfied by [MqQueueProducersMqWorkerProducer] or
// [MqQueueProducersMqR2Producer].
type MqQueueProducersUnion interface {
	implementsMqQueueProducer()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MqQueueProducersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MqQueueProducersMqWorkerProducer{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MqQueueProducersMqR2Producer{}),
		},
	)
}

type MqQueueProducersMqWorkerProducer struct {
	Script string                               `json:"script"`
	Type   MqQueueProducersMqWorkerProducerType `json:"type"`
	JSON   mqQueueProducersMqWorkerProducerJSON `json:"-"`
}

// mqQueueProducersMqWorkerProducerJSON contains the JSON metadata for the struct
// [MqQueueProducersMqWorkerProducer]
type mqQueueProducersMqWorkerProducerJSON struct {
	Script      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MqQueueProducersMqWorkerProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mqQueueProducersMqWorkerProducerJSON) RawJSON() string {
	return r.raw
}

func (r MqQueueProducersMqWorkerProducer) implementsMqQueueProducer() {}

type MqQueueProducersMqWorkerProducerType string

const (
	MqQueueProducersMqWorkerProducerTypeWorker MqQueueProducersMqWorkerProducerType = "worker"
)

func (r MqQueueProducersMqWorkerProducerType) IsKnown() bool {
	switch r {
	case MqQueueProducersMqWorkerProducerTypeWorker:
		return true
	}
	return false
}

type MqQueueProducersMqR2Producer struct {
	BucketName string                           `json:"bucket_name"`
	Type       MqQueueProducersMqR2ProducerType `json:"type"`
	JSON       mqQueueProducersMqR2ProducerJSON `json:"-"`
}

// mqQueueProducersMqR2ProducerJSON contains the JSON metadata for the struct
// [MqQueueProducersMqR2Producer]
type mqQueueProducersMqR2ProducerJSON struct {
	BucketName  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MqQueueProducersMqR2Producer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mqQueueProducersMqR2ProducerJSON) RawJSON() string {
	return r.raw
}

func (r MqQueueProducersMqR2Producer) implementsMqQueueProducer() {}

type MqQueueProducersMqR2ProducerType string

const (
	MqQueueProducersMqR2ProducerTypeR2Bucket MqQueueProducersMqR2ProducerType = "r2_bucket"
)

func (r MqQueueProducersMqR2ProducerType) IsKnown() bool {
	switch r {
	case MqQueueProducersMqR2ProducerTypeR2Bucket:
		return true
	}
	return false
}

type MqQueueProducersType string

const (
	MqQueueProducersTypeWorker   MqQueueProducersType = "worker"
	MqQueueProducersTypeR2Bucket MqQueueProducersType = "r2_bucket"
)

func (r MqQueueProducersType) IsKnown() bool {
	switch r {
	case MqQueueProducersTypeWorker, MqQueueProducersTypeR2Bucket:
		return true
	}
	return false
}

type MqQueueSettings struct {
	// Number of seconds to delay delivery of all messages to consumers.
	DeliveryDelay float64 `json:"delivery_delay"`
	// Indicates if message delivery to consumers is currently paused.
	DeliveryPaused bool `json:"delivery_paused"`
	// Number of seconds after which an unconsumed message will be delayed.
	MessageRetentionPeriod float64             `json:"message_retention_period"`
	JSON                   mqQueueSettingsJSON `json:"-"`
}

// mqQueueSettingsJSON contains the JSON metadata for the struct [MqQueueSettings]
type mqQueueSettingsJSON struct {
	DeliveryDelay          apijson.Field
	DeliveryPaused         apijson.Field
	MessageRetentionPeriod apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *MqQueueSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mqQueueSettingsJSON) RawJSON() string {
	return r.raw
}

type MqQueueParam struct {
	QueueName param.Field[string]               `json:"queue_name"`
	Settings  param.Field[MqQueueSettingsParam] `json:"settings"`
}

func (r MqQueueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MqQueueProducerParam struct {
	BucketName param.Field[string]               `json:"bucket_name"`
	Script     param.Field[string]               `json:"script"`
	Type       param.Field[MqQueueProducersType] `json:"type"`
}

func (r MqQueueProducerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MqQueueProducerParam) implementsMqQueueProducersUnionParam() {}

// Satisfied by [MqQueueProducersMqWorkerProducerParam],
// [MqQueueProducersMqR2ProducerParam], [MqQueueProducerParam].
type MqQueueProducersUnionParam interface {
	implementsMqQueueProducersUnionParam()
}

type MqQueueProducersMqWorkerProducerParam struct {
	Script param.Field[string]                               `json:"script"`
	Type   param.Field[MqQueueProducersMqWorkerProducerType] `json:"type"`
}

func (r MqQueueProducersMqWorkerProducerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MqQueueProducersMqWorkerProducerParam) implementsMqQueueProducersUnionParam() {}

type MqQueueProducersMqR2ProducerParam struct {
	BucketName param.Field[string]                           `json:"bucket_name"`
	Type       param.Field[MqQueueProducersMqR2ProducerType] `json:"type"`
}

func (r MqQueueProducersMqR2ProducerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MqQueueProducersMqR2ProducerParam) implementsMqQueueProducersUnionParam() {}

type MqQueueSettingsParam struct {
	// Number of seconds to delay delivery of all messages to consumers.
	DeliveryDelay param.Field[float64] `json:"delivery_delay"`
	// Indicates if message delivery to consumers is currently paused.
	DeliveryPaused param.Field[bool] `json:"delivery_paused"`
	// Number of seconds after which an unconsumed message will be delayed.
	MessageRetentionPeriod param.Field[float64] `json:"message_retention_period"`
}

func (r MqQueueSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountQueueNewResponse struct {
	Errors   []AccountQueueNewResponseError `json:"errors"`
	Messages []string                       `json:"messages"`
	Result   MqQueue                        `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountQueueNewResponseSuccess `json:"success"`
	JSON    accountQueueNewResponseJSON    `json:"-"`
}

// accountQueueNewResponseJSON contains the JSON metadata for the struct
// [AccountQueueNewResponse]
type accountQueueNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountQueueNewResponseError struct {
	Code             int64                               `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           AccountQueueNewResponseErrorsSource `json:"source"`
	JSON             accountQueueNewResponseErrorJSON    `json:"-"`
}

// accountQueueNewResponseErrorJSON contains the JSON metadata for the struct
// [AccountQueueNewResponseError]
type accountQueueNewResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountQueueNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountQueueNewResponseErrorsSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    accountQueueNewResponseErrorsSourceJSON `json:"-"`
}

// accountQueueNewResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountQueueNewResponseErrorsSource]
type accountQueueNewResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueNewResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueNewResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountQueueNewResponseSuccess bool

const (
	AccountQueueNewResponseSuccessTrue AccountQueueNewResponseSuccess = true
)

func (r AccountQueueNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountQueueNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountQueueGetResponse struct {
	Errors   []AccountQueueGetResponseError `json:"errors"`
	Messages []string                       `json:"messages"`
	Result   MqQueue                        `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountQueueGetResponseSuccess `json:"success"`
	JSON    accountQueueGetResponseJSON    `json:"-"`
}

// accountQueueGetResponseJSON contains the JSON metadata for the struct
// [AccountQueueGetResponse]
type accountQueueGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountQueueGetResponseError struct {
	Code             int64                               `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           AccountQueueGetResponseErrorsSource `json:"source"`
	JSON             accountQueueGetResponseErrorJSON    `json:"-"`
}

// accountQueueGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountQueueGetResponseError]
type accountQueueGetResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountQueueGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountQueueGetResponseErrorsSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    accountQueueGetResponseErrorsSourceJSON `json:"-"`
}

// accountQueueGetResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountQueueGetResponseErrorsSource]
type accountQueueGetResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueGetResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueGetResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountQueueGetResponseSuccess bool

const (
	AccountQueueGetResponseSuccessTrue AccountQueueGetResponseSuccess = true
)

func (r AccountQueueGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountQueueGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountQueueUpdateResponse struct {
	Errors   []AccountQueueUpdateResponseError `json:"errors"`
	Messages []string                          `json:"messages"`
	Result   MqQueue                           `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountQueueUpdateResponseSuccess `json:"success"`
	JSON    accountQueueUpdateResponseJSON    `json:"-"`
}

// accountQueueUpdateResponseJSON contains the JSON metadata for the struct
// [AccountQueueUpdateResponse]
type accountQueueUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountQueueUpdateResponseError struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           AccountQueueUpdateResponseErrorsSource `json:"source"`
	JSON             accountQueueUpdateResponseErrorJSON    `json:"-"`
}

// accountQueueUpdateResponseErrorJSON contains the JSON metadata for the struct
// [AccountQueueUpdateResponseError]
type accountQueueUpdateResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountQueueUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueUpdateResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountQueueUpdateResponseErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    accountQueueUpdateResponseErrorsSourceJSON `json:"-"`
}

// accountQueueUpdateResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountQueueUpdateResponseErrorsSource]
type accountQueueUpdateResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueUpdateResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueUpdateResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountQueueUpdateResponseSuccess bool

const (
	AccountQueueUpdateResponseSuccessTrue AccountQueueUpdateResponseSuccess = true
)

func (r AccountQueueUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountQueueUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountQueueListResponse struct {
	Errors     []AccountQueueListResponseError    `json:"errors"`
	Messages   []string                           `json:"messages"`
	Result     []MqQueue                          `json:"result"`
	ResultInfo AccountQueueListResponseResultInfo `json:"result_info"`
	// Indicates if the API call was successful or not.
	Success AccountQueueListResponseSuccess `json:"success"`
	JSON    accountQueueListResponseJSON    `json:"-"`
}

// accountQueueListResponseJSON contains the JSON metadata for the struct
// [AccountQueueListResponse]
type accountQueueListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountQueueListResponseError struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           AccountQueueListResponseErrorsSource `json:"source"`
	JSON             accountQueueListResponseErrorJSON    `json:"-"`
}

// accountQueueListResponseErrorJSON contains the JSON metadata for the struct
// [AccountQueueListResponseError]
type accountQueueListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountQueueListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountQueueListResponseErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    accountQueueListResponseErrorsSourceJSON `json:"-"`
}

// accountQueueListResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountQueueListResponseErrorsSource]
type accountQueueListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountQueueListResponseResultInfo struct {
	// Total number of queues
	Count float64 `json:"count"`
	// Current page within paginated list of queues
	Page float64 `json:"page"`
	// Number of queues per page
	PerPage float64 `json:"per_page"`
	// Total queues available without any search parameters
	TotalCount float64 `json:"total_count"`
	// Total pages available without any search parameters
	TotalPages float64                                `json:"total_pages"`
	JSON       accountQueueListResponseResultInfoJSON `json:"-"`
}

// accountQueueListResponseResultInfoJSON contains the JSON metadata for the struct
// [AccountQueueListResponseResultInfo]
type accountQueueListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountQueueListResponseSuccess bool

const (
	AccountQueueListResponseSuccessTrue AccountQueueListResponseSuccess = true
)

func (r AccountQueueListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountQueueListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountQueueUpdatePartialResponse struct {
	Errors   []AccountQueueUpdatePartialResponseError `json:"errors"`
	Messages []string                                 `json:"messages"`
	Result   MqQueue                                  `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountQueueUpdatePartialResponseSuccess `json:"success"`
	JSON    accountQueueUpdatePartialResponseJSON    `json:"-"`
}

// accountQueueUpdatePartialResponseJSON contains the JSON metadata for the struct
// [AccountQueueUpdatePartialResponse]
type accountQueueUpdatePartialResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueUpdatePartialResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueUpdatePartialResponseJSON) RawJSON() string {
	return r.raw
}

type AccountQueueUpdatePartialResponseError struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           AccountQueueUpdatePartialResponseErrorsSource `json:"source"`
	JSON             accountQueueUpdatePartialResponseErrorJSON    `json:"-"`
}

// accountQueueUpdatePartialResponseErrorJSON contains the JSON metadata for the
// struct [AccountQueueUpdatePartialResponseError]
type accountQueueUpdatePartialResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountQueueUpdatePartialResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueUpdatePartialResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountQueueUpdatePartialResponseErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    accountQueueUpdatePartialResponseErrorsSourceJSON `json:"-"`
}

// accountQueueUpdatePartialResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountQueueUpdatePartialResponseErrorsSource]
type accountQueueUpdatePartialResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountQueueUpdatePartialResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountQueueUpdatePartialResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountQueueUpdatePartialResponseSuccess bool

const (
	AccountQueueUpdatePartialResponseSuccessTrue AccountQueueUpdatePartialResponseSuccess = true
)

func (r AccountQueueUpdatePartialResponseSuccess) IsKnown() bool {
	switch r {
	case AccountQueueUpdatePartialResponseSuccessTrue:
		return true
	}
	return false
}

type AccountQueueNewParams struct {
	QueueName param.Field[string] `json:"queue_name,required"`
}

func (r AccountQueueNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountQueueUpdateParams struct {
	MqQueue MqQueueParam `json:"mq_queue"`
}

func (r AccountQueueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MqQueue)
}

type AccountQueueUpdatePartialParams struct {
	MqQueue MqQueueParam `json:"mq_queue"`
}

func (r AccountQueueUpdatePartialParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MqQueue)
}
