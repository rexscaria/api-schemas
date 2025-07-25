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

// AccountEventNotificationR2ConfigurationQueueService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountEventNotificationR2ConfigurationQueueService] method instead.
type AccountEventNotificationR2ConfigurationQueueService struct {
	Options []option.RequestOption
}

// NewAccountEventNotificationR2ConfigurationQueueService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewAccountEventNotificationR2ConfigurationQueueService(opts ...option.RequestOption) (r *AccountEventNotificationR2ConfigurationQueueService) {
	r = &AccountEventNotificationR2ConfigurationQueueService{}
	r.Options = opts
	return
}

// Create event notification rule.
func (r *AccountEventNotificationR2ConfigurationQueueService) New(ctx context.Context, accountID string, bucketName string, queueID string, params AccountEventNotificationR2ConfigurationQueueNewParams, opts ...option.RequestOption) (res *R2V4Response, err error) {
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
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
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_notifications/r2/%s/configuration/queues/%s", accountID, bucketName, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Delete an event notification rule. **If no body is provided, all rules for
// specified queue will be deleted**.
func (r *AccountEventNotificationR2ConfigurationQueueService) Delete(ctx context.Context, accountID string, bucketName string, queueID string, body AccountEventNotificationR2ConfigurationQueueDeleteParams, opts ...option.RequestOption) (res *R2V4Response, err error) {
	if body.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", body.Jurisdiction)))
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
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_notifications/r2/%s/configuration/queues/%s", accountID, bucketName, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type R2RuleParam struct {
	// Array of R2 object actions that will trigger notifications.
	Actions param.Field[[]R2RuleAction] `json:"actions,required"`
	// A description that can be used to identify the event notification rule after
	// creation.
	Description param.Field[string] `json:"description"`
	// Notifications will be sent only for objects with this prefix.
	Prefix param.Field[string] `json:"prefix"`
	// Notifications will be sent only for objects with this suffix.
	Suffix param.Field[string] `json:"suffix"`
}

func (r R2RuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type R2RuleAction string

const (
	R2RuleActionPutObject               R2RuleAction = "PutObject"
	R2RuleActionCopyObject              R2RuleAction = "CopyObject"
	R2RuleActionDeleteObject            R2RuleAction = "DeleteObject"
	R2RuleActionCompleteMultipartUpload R2RuleAction = "CompleteMultipartUpload"
	R2RuleActionLifecycleDeletion       R2RuleAction = "LifecycleDeletion"
)

func (r R2RuleAction) IsKnown() bool {
	switch r {
	case R2RuleActionPutObject, R2RuleActionCopyObject, R2RuleActionDeleteObject, R2RuleActionCompleteMultipartUpload, R2RuleActionLifecycleDeletion:
		return true
	}
	return false
}

type AccountEventNotificationR2ConfigurationQueueNewParams struct {
	// Array of rules to drive notifications.
	Rules param.Field[[]R2RuleParam] `json:"rules"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[AccountEventNotificationR2ConfigurationQueueNewParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r AccountEventNotificationR2ConfigurationQueueNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type AccountEventNotificationR2ConfigurationQueueNewParamsCfR2Jurisdiction string

const (
	AccountEventNotificationR2ConfigurationQueueNewParamsCfR2JurisdictionDefault AccountEventNotificationR2ConfigurationQueueNewParamsCfR2Jurisdiction = "default"
	AccountEventNotificationR2ConfigurationQueueNewParamsCfR2JurisdictionEu      AccountEventNotificationR2ConfigurationQueueNewParamsCfR2Jurisdiction = "eu"
	AccountEventNotificationR2ConfigurationQueueNewParamsCfR2JurisdictionFedramp AccountEventNotificationR2ConfigurationQueueNewParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountEventNotificationR2ConfigurationQueueNewParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountEventNotificationR2ConfigurationQueueNewParamsCfR2JurisdictionDefault, AccountEventNotificationR2ConfigurationQueueNewParamsCfR2JurisdictionEu, AccountEventNotificationR2ConfigurationQueueNewParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type AccountEventNotificationR2ConfigurationQueueDeleteParams struct {
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[AccountEventNotificationR2ConfigurationQueueDeleteParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type AccountEventNotificationR2ConfigurationQueueDeleteParamsCfR2Jurisdiction string

const (
	AccountEventNotificationR2ConfigurationQueueDeleteParamsCfR2JurisdictionDefault AccountEventNotificationR2ConfigurationQueueDeleteParamsCfR2Jurisdiction = "default"
	AccountEventNotificationR2ConfigurationQueueDeleteParamsCfR2JurisdictionEu      AccountEventNotificationR2ConfigurationQueueDeleteParamsCfR2Jurisdiction = "eu"
	AccountEventNotificationR2ConfigurationQueueDeleteParamsCfR2JurisdictionFedramp AccountEventNotificationR2ConfigurationQueueDeleteParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountEventNotificationR2ConfigurationQueueDeleteParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountEventNotificationR2ConfigurationQueueDeleteParamsCfR2JurisdictionDefault, AccountEventNotificationR2ConfigurationQueueDeleteParamsCfR2JurisdictionEu, AccountEventNotificationR2ConfigurationQueueDeleteParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}
