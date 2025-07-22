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
	Result AccountEventNotificationR2ConfigurationListResponseResult `json:"result"`
	JSON   accountEventNotificationR2ConfigurationListResponseJSON   `json:"-"`
	R2V4Response
}

// accountEventNotificationR2ConfigurationListResponseJSON contains the JSON
// metadata for the struct [AccountEventNotificationR2ConfigurationListResponse]
type accountEventNotificationR2ConfigurationListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEventNotificationR2ConfigurationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEventNotificationR2ConfigurationListResponseJSON) RawJSON() string {
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
	// Queue ID
	QueueID string `json:"queueId"`
	// Name of the queue
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
	// Timestamp when the rule was created
	CreatedAt string `json:"createdAt"`
	// A description that can be used to identify the event notification rule after
	// creation
	Description string `json:"description"`
	// Rule ID
	RuleID string                                                                  `json:"ruleId"`
	JSON   accountEventNotificationR2ConfigurationListResponseResultQueuesRuleJSON `json:"-"`
	R2Rule
}

// accountEventNotificationR2ConfigurationListResponseResultQueuesRuleJSON contains
// the JSON metadata for the struct
// [AccountEventNotificationR2ConfigurationListResponseResultQueuesRule]
type accountEventNotificationR2ConfigurationListResponseResultQueuesRuleJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	RuleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEventNotificationR2ConfigurationListResponseResultQueuesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEventNotificationR2ConfigurationListResponseResultQueuesRuleJSON) RawJSON() string {
	return r.raw
}

type AccountEventNotificationR2ConfigurationListParams struct {
	// The bucket jurisdiction
	Jurisdiction param.Field[AccountEventNotificationR2ConfigurationListParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction
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
