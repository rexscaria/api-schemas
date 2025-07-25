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

// AccountEventNotificationR2ConfigurationService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountEventNotificationR2ConfigurationService] method instead.
type AccountEventNotificationR2ConfigurationService struct {
	Options []option.RequestOption
	Queues  *AccountEventNotificationR2ConfigurationQueueService
}

// NewAccountEventNotificationR2ConfigurationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountEventNotificationR2ConfigurationService(opts ...option.RequestOption) (r *AccountEventNotificationR2ConfigurationService) {
	r = &AccountEventNotificationR2ConfigurationService{}
	r.Options = opts
	r.Queues = NewAccountEventNotificationR2ConfigurationQueueService(opts...)
	return
}

// List all event notification rules for a bucket.
func (r *AccountEventNotificationR2ConfigurationService) List(ctx context.Context, accountID string, bucketName string, query AccountEventNotificationR2ConfigurationListParams, opts ...option.RequestOption) (res *AccountEventNotificationR2ConfigurationListResponse, err error) {
	if query.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", query.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_notifications/r2/%s/configuration", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountEventNotificationR2ConfigurationListResponse struct {
	Errors   []AccountEventNotificationR2ConfigurationListResponseError `json:"errors,required"`
	Messages []string                                                   `json:"messages,required"`
	Result   AccountEventNotificationR2ConfigurationListResponseResult  `json:"result,required"`
	// Whether the API call was successful.
	Success AccountEventNotificationR2ConfigurationListResponseSuccess `json:"success,required"`
	JSON    accountEventNotificationR2ConfigurationListResponseJSON    `json:"-"`
}

// accountEventNotificationR2ConfigurationListResponseJSON contains the JSON
// metadata for the struct [AccountEventNotificationR2ConfigurationListResponse]
type accountEventNotificationR2ConfigurationListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEventNotificationR2ConfigurationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEventNotificationR2ConfigurationListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEventNotificationR2ConfigurationListResponseError struct {
	Code             int64                                                           `json:"code,required"`
	Message          string                                                          `json:"message,required"`
	DocumentationURL string                                                          `json:"documentation_url"`
	Source           AccountEventNotificationR2ConfigurationListResponseErrorsSource `json:"source"`
	JSON             accountEventNotificationR2ConfigurationListResponseErrorJSON    `json:"-"`
}

// accountEventNotificationR2ConfigurationListResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountEventNotificationR2ConfigurationListResponseError]
type accountEventNotificationR2ConfigurationListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountEventNotificationR2ConfigurationListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEventNotificationR2ConfigurationListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountEventNotificationR2ConfigurationListResponseErrorsSource struct {
	Pointer string                                                              `json:"pointer"`
	JSON    accountEventNotificationR2ConfigurationListResponseErrorsSourceJSON `json:"-"`
}

// accountEventNotificationR2ConfigurationListResponseErrorsSourceJSON contains the
// JSON metadata for the struct
// [AccountEventNotificationR2ConfigurationListResponseErrorsSource]
type accountEventNotificationR2ConfigurationListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEventNotificationR2ConfigurationListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEventNotificationR2ConfigurationListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountEventNotificationR2ConfigurationListResponseResult struct {
	// Name of the bucket.
	BucketName string `json:"bucketName"`
	// List of queues associated with the bucket.
	Queues []AccountEventNotificationR2ConfigurationListResponseResultQueue `json:"queues"`
	JSON   accountEventNotificationR2ConfigurationListResponseResultJSON    `json:"-"`
}

// accountEventNotificationR2ConfigurationListResponseResultJSON contains the JSON
// metadata for the struct
// [AccountEventNotificationR2ConfigurationListResponseResult]
type accountEventNotificationR2ConfigurationListResponseResultJSON struct {
	BucketName  apijson.Field
	Queues      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEventNotificationR2ConfigurationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEventNotificationR2ConfigurationListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEventNotificationR2ConfigurationListResponseResultQueue struct {
	// Queue ID.
	QueueID string `json:"queueId"`
	// Name of the queue.
	QueueName string                                                                `json:"queueName"`
	Rules     []AccountEventNotificationR2ConfigurationListResponseResultQueuesRule `json:"rules"`
	JSON      accountEventNotificationR2ConfigurationListResponseResultQueueJSON    `json:"-"`
}

// accountEventNotificationR2ConfigurationListResponseResultQueueJSON contains the
// JSON metadata for the struct
// [AccountEventNotificationR2ConfigurationListResponseResultQueue]
type accountEventNotificationR2ConfigurationListResponseResultQueueJSON struct {
	QueueID     apijson.Field
	QueueName   apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEventNotificationR2ConfigurationListResponseResultQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEventNotificationR2ConfigurationListResponseResultQueueJSON) RawJSON() string {
	return r.raw
}

type AccountEventNotificationR2ConfigurationListResponseResultQueuesRule struct {
	// Array of R2 object actions that will trigger notifications.
	Actions []AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesAction `json:"actions,required"`
	// Timestamp when the rule was created.
	CreatedAt string `json:"createdAt"`
	// A description that can be used to identify the event notification rule after
	// creation.
	Description string `json:"description"`
	// Notifications will be sent only for objects with this prefix.
	Prefix string `json:"prefix"`
	// Rule ID.
	RuleID string `json:"ruleId"`
	// Notifications will be sent only for objects with this suffix.
	Suffix string                                                                  `json:"suffix"`
	JSON   accountEventNotificationR2ConfigurationListResponseResultQueuesRuleJSON `json:"-"`
}

// accountEventNotificationR2ConfigurationListResponseResultQueuesRuleJSON contains
// the JSON metadata for the struct
// [AccountEventNotificationR2ConfigurationListResponseResultQueuesRule]
type accountEventNotificationR2ConfigurationListResponseResultQueuesRuleJSON struct {
	Actions     apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Prefix      apijson.Field
	RuleID      apijson.Field
	Suffix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEventNotificationR2ConfigurationListResponseResultQueuesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEventNotificationR2ConfigurationListResponseResultQueuesRuleJSON) RawJSON() string {
	return r.raw
}

type AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesAction string

const (
	AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesActionPutObject               AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesAction = "PutObject"
	AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesActionCopyObject              AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesAction = "CopyObject"
	AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesActionDeleteObject            AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesAction = "DeleteObject"
	AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesActionCompleteMultipartUpload AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesAction = "CompleteMultipartUpload"
	AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesActionLifecycleDeletion       AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesAction = "LifecycleDeletion"
)

func (r AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesAction) IsKnown() bool {
	switch r {
	case AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesActionPutObject, AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesActionCopyObject, AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesActionDeleteObject, AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesActionCompleteMultipartUpload, AccountEventNotificationR2ConfigurationListResponseResultQueuesRulesActionLifecycleDeletion:
		return true
	}
	return false
}

// Whether the API call was successful.
type AccountEventNotificationR2ConfigurationListResponseSuccess bool

const (
	AccountEventNotificationR2ConfigurationListResponseSuccessTrue AccountEventNotificationR2ConfigurationListResponseSuccess = true
)

func (r AccountEventNotificationR2ConfigurationListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountEventNotificationR2ConfigurationListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountEventNotificationR2ConfigurationListParams struct {
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[AccountEventNotificationR2ConfigurationListParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type AccountEventNotificationR2ConfigurationListParamsCfR2Jurisdiction string

const (
	AccountEventNotificationR2ConfigurationListParamsCfR2JurisdictionDefault AccountEventNotificationR2ConfigurationListParamsCfR2Jurisdiction = "default"
	AccountEventNotificationR2ConfigurationListParamsCfR2JurisdictionEu      AccountEventNotificationR2ConfigurationListParamsCfR2Jurisdiction = "eu"
	AccountEventNotificationR2ConfigurationListParamsCfR2JurisdictionFedramp AccountEventNotificationR2ConfigurationListParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountEventNotificationR2ConfigurationListParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountEventNotificationR2ConfigurationListParamsCfR2JurisdictionDefault, AccountEventNotificationR2ConfigurationListParamsCfR2JurisdictionEu, AccountEventNotificationR2ConfigurationListParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}
