// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// AccountWorkflowInstanceService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkflowInstanceService] method instead.
type AccountWorkflowInstanceService struct {
	Options []option.RequestOption
}

// NewAccountWorkflowInstanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountWorkflowInstanceService(opts ...option.RequestOption) (r *AccountWorkflowInstanceService) {
	r = &AccountWorkflowInstanceService{}
	r.Options = opts
	return
}

// Create a new workflow instance
func (r *AccountWorkflowInstanceService) New(ctx context.Context, accountID string, workflowName string, body AccountWorkflowInstanceNewParams, opts ...option.RequestOption) (res *AccountWorkflowInstanceNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances", accountID, workflowName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get logs and status from instance
func (r *AccountWorkflowInstanceService) Get(ctx context.Context, accountID string, workflowName string, instanceID string, opts ...option.RequestOption) (res *AccountWorkflowInstanceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances/%s", accountID, workflowName, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List of workflow instances
func (r *AccountWorkflowInstanceService) List(ctx context.Context, accountID string, workflowName string, query AccountWorkflowInstanceListParams, opts ...option.RequestOption) (res *AccountWorkflowInstanceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances", accountID, workflowName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Batch create new Workflow instances
func (r *AccountWorkflowInstanceService) BatchNew(ctx context.Context, accountID string, workflowName string, body AccountWorkflowInstanceBatchNewParams, opts ...option.RequestOption) (res *AccountWorkflowInstanceBatchNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances/batch", accountID, workflowName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send event to instance
func (r *AccountWorkflowInstanceService) SendEvent(ctx context.Context, accountID string, workflowName string, instanceID string, eventType string, body AccountWorkflowInstanceSendEventParams, opts ...option.RequestOption) (res *AccountWorkflowInstanceSendEventResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	if eventType == "" {
		err = errors.New("missing required event_type parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances/%s/events/%s", accountID, workflowName, instanceID, eventType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change status of instance
func (r *AccountWorkflowInstanceService) UpdateStatus(ctx context.Context, accountID string, workflowName string, instanceID string, body AccountWorkflowInstanceUpdateStatusParams, opts ...option.RequestOption) (res *AccountWorkflowInstanceUpdateStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances/%s/status", accountID, workflowName, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AccountWorkflowInstanceNewResponse struct {
	Errors     []AccountWorkflowInstanceNewResponseError    `json:"errors,required"`
	Messages   []AccountWorkflowInstanceNewResponseMessage  `json:"messages,required"`
	Result     AccountWorkflowInstanceNewResponseResult     `json:"result,required"`
	Success    AccountWorkflowInstanceNewResponseSuccess    `json:"success,required"`
	ResultInfo AccountWorkflowInstanceNewResponseResultInfo `json:"result_info"`
	JSON       accountWorkflowInstanceNewResponseJSON       `json:"-"`
}

// accountWorkflowInstanceNewResponseJSON contains the JSON metadata for the struct
// [AccountWorkflowInstanceNewResponse]
type accountWorkflowInstanceNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceNewResponseError struct {
	Code    float64                                     `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountWorkflowInstanceNewResponseErrorJSON `json:"-"`
}

// accountWorkflowInstanceNewResponseErrorJSON contains the JSON metadata for the
// struct [AccountWorkflowInstanceNewResponseError]
type accountWorkflowInstanceNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceNewResponseMessage struct {
	Code    float64                                       `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountWorkflowInstanceNewResponseMessageJSON `json:"-"`
}

// accountWorkflowInstanceNewResponseMessageJSON contains the JSON metadata for the
// struct [AccountWorkflowInstanceNewResponseMessage]
type accountWorkflowInstanceNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceNewResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceNewResponseResult struct {
	ID         string                                         `json:"id,required"`
	Status     AccountWorkflowInstanceNewResponseResultStatus `json:"status,required"`
	VersionID  string                                         `json:"version_id,required" format:"uuid"`
	WorkflowID string                                         `json:"workflow_id,required" format:"uuid"`
	JSON       accountWorkflowInstanceNewResponseResultJSON   `json:"-"`
}

// accountWorkflowInstanceNewResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkflowInstanceNewResponseResult]
type accountWorkflowInstanceNewResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	VersionID   apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceNewResponseResultStatus string

const (
	AccountWorkflowInstanceNewResponseResultStatusQueued          AccountWorkflowInstanceNewResponseResultStatus = "queued"
	AccountWorkflowInstanceNewResponseResultStatusRunning         AccountWorkflowInstanceNewResponseResultStatus = "running"
	AccountWorkflowInstanceNewResponseResultStatusPaused          AccountWorkflowInstanceNewResponseResultStatus = "paused"
	AccountWorkflowInstanceNewResponseResultStatusErrored         AccountWorkflowInstanceNewResponseResultStatus = "errored"
	AccountWorkflowInstanceNewResponseResultStatusTerminated      AccountWorkflowInstanceNewResponseResultStatus = "terminated"
	AccountWorkflowInstanceNewResponseResultStatusComplete        AccountWorkflowInstanceNewResponseResultStatus = "complete"
	AccountWorkflowInstanceNewResponseResultStatusWaitingForPause AccountWorkflowInstanceNewResponseResultStatus = "waitingForPause"
	AccountWorkflowInstanceNewResponseResultStatusWaiting         AccountWorkflowInstanceNewResponseResultStatus = "waiting"
)

func (r AccountWorkflowInstanceNewResponseResultStatus) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceNewResponseResultStatusQueued, AccountWorkflowInstanceNewResponseResultStatusRunning, AccountWorkflowInstanceNewResponseResultStatusPaused, AccountWorkflowInstanceNewResponseResultStatusErrored, AccountWorkflowInstanceNewResponseResultStatusTerminated, AccountWorkflowInstanceNewResponseResultStatusComplete, AccountWorkflowInstanceNewResponseResultStatusWaitingForPause, AccountWorkflowInstanceNewResponseResultStatusWaiting:
		return true
	}
	return false
}

type AccountWorkflowInstanceNewResponseSuccess bool

const (
	AccountWorkflowInstanceNewResponseSuccessTrue AccountWorkflowInstanceNewResponseSuccess = true
)

func (r AccountWorkflowInstanceNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkflowInstanceNewResponseResultInfo struct {
	Count      float64                                          `json:"count,required"`
	Page       float64                                          `json:"page,required"`
	PerPage    float64                                          `json:"per_page,required"`
	TotalCount float64                                          `json:"total_count,required"`
	JSON       accountWorkflowInstanceNewResponseResultInfoJSON `json:"-"`
}

// accountWorkflowInstanceNewResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountWorkflowInstanceNewResponseResultInfo]
type accountWorkflowInstanceNewResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceNewResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceNewResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceGetResponse struct {
	Errors     []AccountWorkflowInstanceGetResponseError    `json:"errors,required"`
	Messages   []AccountWorkflowInstanceGetResponseMessage  `json:"messages,required"`
	Result     AccountWorkflowInstanceGetResponseResult     `json:"result,required"`
	Success    AccountWorkflowInstanceGetResponseSuccess    `json:"success,required"`
	ResultInfo AccountWorkflowInstanceGetResponseResultInfo `json:"result_info"`
	JSON       accountWorkflowInstanceGetResponseJSON       `json:"-"`
}

// accountWorkflowInstanceGetResponseJSON contains the JSON metadata for the struct
// [AccountWorkflowInstanceGetResponse]
type accountWorkflowInstanceGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceGetResponseError struct {
	Code    float64                                     `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountWorkflowInstanceGetResponseErrorJSON `json:"-"`
}

// accountWorkflowInstanceGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountWorkflowInstanceGetResponseError]
type accountWorkflowInstanceGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceGetResponseMessage struct {
	Code    float64                                       `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountWorkflowInstanceGetResponseMessageJSON `json:"-"`
}

// accountWorkflowInstanceGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountWorkflowInstanceGetResponseMessage]
type accountWorkflowInstanceGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceGetResponseResult struct {
	End       time.Time                                           `json:"end,required,nullable" format:"date-time"`
	Error     AccountWorkflowInstanceGetResponseResultError       `json:"error,required,nullable"`
	Output    AccountWorkflowInstanceGetResponseResultOutputUnion `json:"output,required"`
	Params    interface{}                                         `json:"params,required"`
	Queued    time.Time                                           `json:"queued,required" format:"date-time"`
	Start     time.Time                                           `json:"start,required,nullable" format:"date-time"`
	Status    AccountWorkflowInstanceGetResponseResultStatus      `json:"status,required"`
	Steps     []AccountWorkflowInstanceGetResponseResultStep      `json:"steps,required"`
	Success   bool                                                `json:"success,required,nullable"`
	Trigger   AccountWorkflowInstanceGetResponseResultTrigger     `json:"trigger,required"`
	VersionID string                                              `json:"versionId,required" format:"uuid"`
	JSON      accountWorkflowInstanceGetResponseResultJSON        `json:"-"`
}

// accountWorkflowInstanceGetResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkflowInstanceGetResponseResult]
type accountWorkflowInstanceGetResponseResultJSON struct {
	End         apijson.Field
	Error       apijson.Field
	Output      apijson.Field
	Params      apijson.Field
	Queued      apijson.Field
	Start       apijson.Field
	Status      apijson.Field
	Steps       apijson.Field
	Success     apijson.Field
	Trigger     apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceGetResponseResultError struct {
	Message string                                            `json:"message,required"`
	Name    string                                            `json:"name,required"`
	JSON    accountWorkflowInstanceGetResponseResultErrorJSON `json:"-"`
}

// accountWorkflowInstanceGetResponseResultErrorJSON contains the JSON metadata for
// the struct [AccountWorkflowInstanceGetResponseResultError]
type accountWorkflowInstanceGetResponseResultErrorJSON struct {
	Message     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceGetResponseResultError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceGetResponseResultErrorJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type AccountWorkflowInstanceGetResponseResultOutputUnion interface {
	ImplementsAccountWorkflowInstanceGetResponseResultOutputUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountWorkflowInstanceGetResponseResultOutputUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type AccountWorkflowInstanceGetResponseResultStatus string

const (
	AccountWorkflowInstanceGetResponseResultStatusQueued          AccountWorkflowInstanceGetResponseResultStatus = "queued"
	AccountWorkflowInstanceGetResponseResultStatusRunning         AccountWorkflowInstanceGetResponseResultStatus = "running"
	AccountWorkflowInstanceGetResponseResultStatusPaused          AccountWorkflowInstanceGetResponseResultStatus = "paused"
	AccountWorkflowInstanceGetResponseResultStatusErrored         AccountWorkflowInstanceGetResponseResultStatus = "errored"
	AccountWorkflowInstanceGetResponseResultStatusTerminated      AccountWorkflowInstanceGetResponseResultStatus = "terminated"
	AccountWorkflowInstanceGetResponseResultStatusComplete        AccountWorkflowInstanceGetResponseResultStatus = "complete"
	AccountWorkflowInstanceGetResponseResultStatusWaitingForPause AccountWorkflowInstanceGetResponseResultStatus = "waitingForPause"
	AccountWorkflowInstanceGetResponseResultStatusWaiting         AccountWorkflowInstanceGetResponseResultStatus = "waiting"
)

func (r AccountWorkflowInstanceGetResponseResultStatus) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceGetResponseResultStatusQueued, AccountWorkflowInstanceGetResponseResultStatusRunning, AccountWorkflowInstanceGetResponseResultStatusPaused, AccountWorkflowInstanceGetResponseResultStatusErrored, AccountWorkflowInstanceGetResponseResultStatusTerminated, AccountWorkflowInstanceGetResponseResultStatusComplete, AccountWorkflowInstanceGetResponseResultStatusWaitingForPause, AccountWorkflowInstanceGetResponseResultStatusWaiting:
		return true
	}
	return false
}

type AccountWorkflowInstanceGetResponseResultStep struct {
	Type AccountWorkflowInstanceGetResponseResultStepsType `json:"type,required"`
	// This field can have the runtime type of
	// [[]AccountWorkflowInstanceGetResponseResultStepsObjectAttempt].
	Attempts interface{} `json:"attempts"`
	// This field can have the runtime type of
	// [AccountWorkflowInstanceGetResponseResultStepsObjectConfig].
	Config interface{} `json:"config"`
	End    time.Time   `json:"end,nullable" format:"date-time"`
	// This field can have the runtime type of
	// [AccountWorkflowInstanceGetResponseResultStepsObjectError].
	Error    interface{} `json:"error"`
	Finished bool        `json:"finished"`
	Name     string      `json:"name"`
	// This field can have the runtime type of [interface{}].
	Output  interface{} `json:"output"`
	Start   time.Time   `json:"start" format:"date-time"`
	Success bool        `json:"success,nullable"`
	// This field can have the runtime type of
	// [AccountWorkflowInstanceGetResponseResultStepsObjectTrigger].
	Trigger interface{}                                      `json:"trigger"`
	JSON    accountWorkflowInstanceGetResponseResultStepJSON `json:"-"`
	union   AccountWorkflowInstanceGetResponseResultStepsUnion
}

// accountWorkflowInstanceGetResponseResultStepJSON contains the JSON metadata for
// the struct [AccountWorkflowInstanceGetResponseResultStep]
type accountWorkflowInstanceGetResponseResultStepJSON struct {
	Type        apijson.Field
	Attempts    apijson.Field
	Config      apijson.Field
	End         apijson.Field
	Error       apijson.Field
	Finished    apijson.Field
	Name        apijson.Field
	Output      apijson.Field
	Start       apijson.Field
	Success     apijson.Field
	Trigger     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r accountWorkflowInstanceGetResponseResultStepJSON) RawJSON() string {
	return r.raw
}

func (r *AccountWorkflowInstanceGetResponseResultStep) UnmarshalJSON(data []byte) (err error) {
	*r = AccountWorkflowInstanceGetResponseResultStep{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountWorkflowInstanceGetResponseResultStepsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountWorkflowInstanceGetResponseResultStepsObject],
// [AccountWorkflowInstanceGetResponseResultStepsObject],
// [AccountWorkflowInstanceGetResponseResultStepsObject],
// [AccountWorkflowInstanceGetResponseResultStepsObject].
func (r AccountWorkflowInstanceGetResponseResultStep) AsUnion() AccountWorkflowInstanceGetResponseResultStepsUnion {
	return r.union
}

// Union satisfied by [AccountWorkflowInstanceGetResponseResultStepsObject],
// [AccountWorkflowInstanceGetResponseResultStepsObject],
// [AccountWorkflowInstanceGetResponseResultStepsObject] or
// [AccountWorkflowInstanceGetResponseResultStepsObject].
type AccountWorkflowInstanceGetResponseResultStepsUnion interface {
	implementsAccountWorkflowInstanceGetResponseResultStep()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountWorkflowInstanceGetResponseResultStepsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountWorkflowInstanceGetResponseResultStepsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountWorkflowInstanceGetResponseResultStepsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountWorkflowInstanceGetResponseResultStepsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountWorkflowInstanceGetResponseResultStepsObject{}),
		},
	)
}

type AccountWorkflowInstanceGetResponseResultStepsObject struct {
	Attempts []AccountWorkflowInstanceGetResponseResultStepsObjectAttempt `json:"attempts,required"`
	Config   AccountWorkflowInstanceGetResponseResultStepsObjectConfig    `json:"config,required"`
	End      time.Time                                                    `json:"end,required,nullable" format:"date-time"`
	Name     string                                                       `json:"name,required"`
	Output   interface{}                                                  `json:"output,required"`
	Start    time.Time                                                    `json:"start,required" format:"date-time"`
	Success  bool                                                         `json:"success,required,nullable"`
	Type     AccountWorkflowInstanceGetResponseResultStepsObjectType      `json:"type,required"`
	JSON     accountWorkflowInstanceGetResponseResultStepsObjectJSON      `json:"-"`
}

// accountWorkflowInstanceGetResponseResultStepsObjectJSON contains the JSON
// metadata for the struct [AccountWorkflowInstanceGetResponseResultStepsObject]
type accountWorkflowInstanceGetResponseResultStepsObjectJSON struct {
	Attempts    apijson.Field
	Config      apijson.Field
	End         apijson.Field
	Name        apijson.Field
	Output      apijson.Field
	Start       apijson.Field
	Success     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceGetResponseResultStepsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceGetResponseResultStepsObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountWorkflowInstanceGetResponseResultStepsObject) implementsAccountWorkflowInstanceGetResponseResultStep() {
}

type AccountWorkflowInstanceGetResponseResultStepsObjectAttempt struct {
	End     time.Time                                                        `json:"end,required,nullable" format:"date-time"`
	Error   AccountWorkflowInstanceGetResponseResultStepsObjectAttemptsError `json:"error,required,nullable"`
	Start   time.Time                                                        `json:"start,required" format:"date-time"`
	Success bool                                                             `json:"success,required,nullable"`
	JSON    accountWorkflowInstanceGetResponseResultStepsObjectAttemptJSON   `json:"-"`
}

// accountWorkflowInstanceGetResponseResultStepsObjectAttemptJSON contains the JSON
// metadata for the struct
// [AccountWorkflowInstanceGetResponseResultStepsObjectAttempt]
type accountWorkflowInstanceGetResponseResultStepsObjectAttemptJSON struct {
	End         apijson.Field
	Error       apijson.Field
	Start       apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceGetResponseResultStepsObjectAttempt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceGetResponseResultStepsObjectAttemptJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceGetResponseResultStepsObjectAttemptsError struct {
	Message string                                                               `json:"message,required"`
	Name    string                                                               `json:"name,required"`
	JSON    accountWorkflowInstanceGetResponseResultStepsObjectAttemptsErrorJSON `json:"-"`
}

// accountWorkflowInstanceGetResponseResultStepsObjectAttemptsErrorJSON contains
// the JSON metadata for the struct
// [AccountWorkflowInstanceGetResponseResultStepsObjectAttemptsError]
type accountWorkflowInstanceGetResponseResultStepsObjectAttemptsErrorJSON struct {
	Message     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceGetResponseResultStepsObjectAttemptsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceGetResponseResultStepsObjectAttemptsErrorJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceGetResponseResultStepsObjectConfig struct {
	Retries AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetries `json:"retries,required"`
	Timeout interface{}                                                      `json:"timeout,required"`
	JSON    accountWorkflowInstanceGetResponseResultStepsObjectConfigJSON    `json:"-"`
}

// accountWorkflowInstanceGetResponseResultStepsObjectConfigJSON contains the JSON
// metadata for the struct
// [AccountWorkflowInstanceGetResponseResultStepsObjectConfig]
type accountWorkflowInstanceGetResponseResultStepsObjectConfigJSON struct {
	Retries     apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceGetResponseResultStepsObjectConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceGetResponseResultStepsObjectConfigJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetries struct {
	Delay   interface{}                                                             `json:"delay,required"`
	Limit   float64                                                                 `json:"limit,required"`
	Backoff AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesBackoff `json:"backoff"`
	JSON    accountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesJSON    `json:"-"`
}

// accountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesJSON contains
// the JSON metadata for the struct
// [AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetries]
type accountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesJSON struct {
	Delay       apijson.Field
	Limit       apijson.Field
	Backoff     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesBackoff string

const (
	AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesBackoffConstant    AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesBackoff = "constant"
	AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesBackoffLinear      AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesBackoff = "linear"
	AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesBackoffExponential AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesBackoff = "exponential"
)

func (r AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesBackoff) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesBackoffConstant, AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesBackoffLinear, AccountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesBackoffExponential:
		return true
	}
	return false
}

type AccountWorkflowInstanceGetResponseResultStepsObjectType string

const (
	AccountWorkflowInstanceGetResponseResultStepsObjectTypeStep AccountWorkflowInstanceGetResponseResultStepsObjectType = "step"
)

func (r AccountWorkflowInstanceGetResponseResultStepsObjectType) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceGetResponseResultStepsObjectTypeStep:
		return true
	}
	return false
}

type AccountWorkflowInstanceGetResponseResultStepsType string

const (
	AccountWorkflowInstanceGetResponseResultStepsTypeStep         AccountWorkflowInstanceGetResponseResultStepsType = "step"
	AccountWorkflowInstanceGetResponseResultStepsTypeSleep        AccountWorkflowInstanceGetResponseResultStepsType = "sleep"
	AccountWorkflowInstanceGetResponseResultStepsTypeTermination  AccountWorkflowInstanceGetResponseResultStepsType = "termination"
	AccountWorkflowInstanceGetResponseResultStepsTypeWaitForEvent AccountWorkflowInstanceGetResponseResultStepsType = "waitForEvent"
)

func (r AccountWorkflowInstanceGetResponseResultStepsType) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceGetResponseResultStepsTypeStep, AccountWorkflowInstanceGetResponseResultStepsTypeSleep, AccountWorkflowInstanceGetResponseResultStepsTypeTermination, AccountWorkflowInstanceGetResponseResultStepsTypeWaitForEvent:
		return true
	}
	return false
}

type AccountWorkflowInstanceGetResponseResultTrigger struct {
	Source AccountWorkflowInstanceGetResponseResultTriggerSource `json:"source,required"`
	JSON   accountWorkflowInstanceGetResponseResultTriggerJSON   `json:"-"`
}

// accountWorkflowInstanceGetResponseResultTriggerJSON contains the JSON metadata
// for the struct [AccountWorkflowInstanceGetResponseResultTrigger]
type accountWorkflowInstanceGetResponseResultTriggerJSON struct {
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceGetResponseResultTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceGetResponseResultTriggerJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceGetResponseResultTriggerSource string

const (
	AccountWorkflowInstanceGetResponseResultTriggerSourceUnknown AccountWorkflowInstanceGetResponseResultTriggerSource = "unknown"
	AccountWorkflowInstanceGetResponseResultTriggerSourceAPI     AccountWorkflowInstanceGetResponseResultTriggerSource = "api"
	AccountWorkflowInstanceGetResponseResultTriggerSourceBinding AccountWorkflowInstanceGetResponseResultTriggerSource = "binding"
	AccountWorkflowInstanceGetResponseResultTriggerSourceEvent   AccountWorkflowInstanceGetResponseResultTriggerSource = "event"
	AccountWorkflowInstanceGetResponseResultTriggerSourceCron    AccountWorkflowInstanceGetResponseResultTriggerSource = "cron"
)

func (r AccountWorkflowInstanceGetResponseResultTriggerSource) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceGetResponseResultTriggerSourceUnknown, AccountWorkflowInstanceGetResponseResultTriggerSourceAPI, AccountWorkflowInstanceGetResponseResultTriggerSourceBinding, AccountWorkflowInstanceGetResponseResultTriggerSourceEvent, AccountWorkflowInstanceGetResponseResultTriggerSourceCron:
		return true
	}
	return false
}

type AccountWorkflowInstanceGetResponseSuccess bool

const (
	AccountWorkflowInstanceGetResponseSuccessTrue AccountWorkflowInstanceGetResponseSuccess = true
)

func (r AccountWorkflowInstanceGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkflowInstanceGetResponseResultInfo struct {
	Count      float64                                          `json:"count,required"`
	Page       float64                                          `json:"page,required"`
	PerPage    float64                                          `json:"per_page,required"`
	TotalCount float64                                          `json:"total_count,required"`
	JSON       accountWorkflowInstanceGetResponseResultInfoJSON `json:"-"`
}

// accountWorkflowInstanceGetResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountWorkflowInstanceGetResponseResultInfo]
type accountWorkflowInstanceGetResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceGetResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceGetResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceListResponse struct {
	Errors     []AccountWorkflowInstanceListResponseError    `json:"errors,required"`
	Messages   []AccountWorkflowInstanceListResponseMessage  `json:"messages,required"`
	Result     []AccountWorkflowInstanceListResponseResult   `json:"result,required"`
	Success    AccountWorkflowInstanceListResponseSuccess    `json:"success,required"`
	ResultInfo AccountWorkflowInstanceListResponseResultInfo `json:"result_info"`
	JSON       accountWorkflowInstanceListResponseJSON       `json:"-"`
}

// accountWorkflowInstanceListResponseJSON contains the JSON metadata for the
// struct [AccountWorkflowInstanceListResponse]
type accountWorkflowInstanceListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceListResponseError struct {
	Code    float64                                      `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountWorkflowInstanceListResponseErrorJSON `json:"-"`
}

// accountWorkflowInstanceListResponseErrorJSON contains the JSON metadata for the
// struct [AccountWorkflowInstanceListResponseError]
type accountWorkflowInstanceListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceListResponseMessage struct {
	Code    float64                                        `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountWorkflowInstanceListResponseMessageJSON `json:"-"`
}

// accountWorkflowInstanceListResponseMessageJSON contains the JSON metadata for
// the struct [AccountWorkflowInstanceListResponseMessage]
type accountWorkflowInstanceListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceListResponseResult struct {
	ID         string                                          `json:"id,required"`
	CreatedOn  time.Time                                       `json:"created_on,required" format:"date-time"`
	EndedOn    time.Time                                       `json:"ended_on,required,nullable" format:"date-time"`
	ModifiedOn time.Time                                       `json:"modified_on,required" format:"date-time"`
	StartedOn  time.Time                                       `json:"started_on,required,nullable" format:"date-time"`
	Status     AccountWorkflowInstanceListResponseResultStatus `json:"status,required"`
	VersionID  string                                          `json:"version_id,required" format:"uuid"`
	WorkflowID string                                          `json:"workflow_id,required" format:"uuid"`
	JSON       accountWorkflowInstanceListResponseResultJSON   `json:"-"`
}

// accountWorkflowInstanceListResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkflowInstanceListResponseResult]
type accountWorkflowInstanceListResponseResultJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	EndedOn     apijson.Field
	ModifiedOn  apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	VersionID   apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceListResponseResultStatus string

const (
	AccountWorkflowInstanceListResponseResultStatusQueued          AccountWorkflowInstanceListResponseResultStatus = "queued"
	AccountWorkflowInstanceListResponseResultStatusRunning         AccountWorkflowInstanceListResponseResultStatus = "running"
	AccountWorkflowInstanceListResponseResultStatusPaused          AccountWorkflowInstanceListResponseResultStatus = "paused"
	AccountWorkflowInstanceListResponseResultStatusErrored         AccountWorkflowInstanceListResponseResultStatus = "errored"
	AccountWorkflowInstanceListResponseResultStatusTerminated      AccountWorkflowInstanceListResponseResultStatus = "terminated"
	AccountWorkflowInstanceListResponseResultStatusComplete        AccountWorkflowInstanceListResponseResultStatus = "complete"
	AccountWorkflowInstanceListResponseResultStatusWaitingForPause AccountWorkflowInstanceListResponseResultStatus = "waitingForPause"
	AccountWorkflowInstanceListResponseResultStatusWaiting         AccountWorkflowInstanceListResponseResultStatus = "waiting"
)

func (r AccountWorkflowInstanceListResponseResultStatus) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceListResponseResultStatusQueued, AccountWorkflowInstanceListResponseResultStatusRunning, AccountWorkflowInstanceListResponseResultStatusPaused, AccountWorkflowInstanceListResponseResultStatusErrored, AccountWorkflowInstanceListResponseResultStatusTerminated, AccountWorkflowInstanceListResponseResultStatusComplete, AccountWorkflowInstanceListResponseResultStatusWaitingForPause, AccountWorkflowInstanceListResponseResultStatusWaiting:
		return true
	}
	return false
}

type AccountWorkflowInstanceListResponseSuccess bool

const (
	AccountWorkflowInstanceListResponseSuccessTrue AccountWorkflowInstanceListResponseSuccess = true
)

func (r AccountWorkflowInstanceListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkflowInstanceListResponseResultInfo struct {
	Count      float64                                           `json:"count,required"`
	Page       float64                                           `json:"page,required"`
	PerPage    float64                                           `json:"per_page,required"`
	TotalCount float64                                           `json:"total_count,required"`
	JSON       accountWorkflowInstanceListResponseResultInfoJSON `json:"-"`
}

// accountWorkflowInstanceListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountWorkflowInstanceListResponseResultInfo]
type accountWorkflowInstanceListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceBatchNewResponse struct {
	Errors     []AccountWorkflowInstanceBatchNewResponseError    `json:"errors,required"`
	Messages   []AccountWorkflowInstanceBatchNewResponseMessage  `json:"messages,required"`
	Result     []AccountWorkflowInstanceBatchNewResponseResult   `json:"result,required"`
	Success    AccountWorkflowInstanceBatchNewResponseSuccess    `json:"success,required"`
	ResultInfo AccountWorkflowInstanceBatchNewResponseResultInfo `json:"result_info"`
	JSON       accountWorkflowInstanceBatchNewResponseJSON       `json:"-"`
}

// accountWorkflowInstanceBatchNewResponseJSON contains the JSON metadata for the
// struct [AccountWorkflowInstanceBatchNewResponse]
type accountWorkflowInstanceBatchNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceBatchNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceBatchNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceBatchNewResponseError struct {
	Code    float64                                          `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountWorkflowInstanceBatchNewResponseErrorJSON `json:"-"`
}

// accountWorkflowInstanceBatchNewResponseErrorJSON contains the JSON metadata for
// the struct [AccountWorkflowInstanceBatchNewResponseError]
type accountWorkflowInstanceBatchNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceBatchNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceBatchNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceBatchNewResponseMessage struct {
	Code    float64                                            `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accountWorkflowInstanceBatchNewResponseMessageJSON `json:"-"`
}

// accountWorkflowInstanceBatchNewResponseMessageJSON contains the JSON metadata
// for the struct [AccountWorkflowInstanceBatchNewResponseMessage]
type accountWorkflowInstanceBatchNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceBatchNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceBatchNewResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceBatchNewResponseResult struct {
	ID         string                                              `json:"id,required"`
	Status     AccountWorkflowInstanceBatchNewResponseResultStatus `json:"status,required"`
	VersionID  string                                              `json:"version_id,required" format:"uuid"`
	WorkflowID string                                              `json:"workflow_id,required" format:"uuid"`
	JSON       accountWorkflowInstanceBatchNewResponseResultJSON   `json:"-"`
}

// accountWorkflowInstanceBatchNewResponseResultJSON contains the JSON metadata for
// the struct [AccountWorkflowInstanceBatchNewResponseResult]
type accountWorkflowInstanceBatchNewResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	VersionID   apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceBatchNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceBatchNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceBatchNewResponseResultStatus string

const (
	AccountWorkflowInstanceBatchNewResponseResultStatusQueued          AccountWorkflowInstanceBatchNewResponseResultStatus = "queued"
	AccountWorkflowInstanceBatchNewResponseResultStatusRunning         AccountWorkflowInstanceBatchNewResponseResultStatus = "running"
	AccountWorkflowInstanceBatchNewResponseResultStatusPaused          AccountWorkflowInstanceBatchNewResponseResultStatus = "paused"
	AccountWorkflowInstanceBatchNewResponseResultStatusErrored         AccountWorkflowInstanceBatchNewResponseResultStatus = "errored"
	AccountWorkflowInstanceBatchNewResponseResultStatusTerminated      AccountWorkflowInstanceBatchNewResponseResultStatus = "terminated"
	AccountWorkflowInstanceBatchNewResponseResultStatusComplete        AccountWorkflowInstanceBatchNewResponseResultStatus = "complete"
	AccountWorkflowInstanceBatchNewResponseResultStatusWaitingForPause AccountWorkflowInstanceBatchNewResponseResultStatus = "waitingForPause"
	AccountWorkflowInstanceBatchNewResponseResultStatusWaiting         AccountWorkflowInstanceBatchNewResponseResultStatus = "waiting"
)

func (r AccountWorkflowInstanceBatchNewResponseResultStatus) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceBatchNewResponseResultStatusQueued, AccountWorkflowInstanceBatchNewResponseResultStatusRunning, AccountWorkflowInstanceBatchNewResponseResultStatusPaused, AccountWorkflowInstanceBatchNewResponseResultStatusErrored, AccountWorkflowInstanceBatchNewResponseResultStatusTerminated, AccountWorkflowInstanceBatchNewResponseResultStatusComplete, AccountWorkflowInstanceBatchNewResponseResultStatusWaitingForPause, AccountWorkflowInstanceBatchNewResponseResultStatusWaiting:
		return true
	}
	return false
}

type AccountWorkflowInstanceBatchNewResponseSuccess bool

const (
	AccountWorkflowInstanceBatchNewResponseSuccessTrue AccountWorkflowInstanceBatchNewResponseSuccess = true
)

func (r AccountWorkflowInstanceBatchNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceBatchNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkflowInstanceBatchNewResponseResultInfo struct {
	Count      float64                                               `json:"count,required"`
	Page       float64                                               `json:"page,required"`
	PerPage    float64                                               `json:"per_page,required"`
	TotalCount float64                                               `json:"total_count,required"`
	JSON       accountWorkflowInstanceBatchNewResponseResultInfoJSON `json:"-"`
}

// accountWorkflowInstanceBatchNewResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountWorkflowInstanceBatchNewResponseResultInfo]
type accountWorkflowInstanceBatchNewResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceBatchNewResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceBatchNewResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceSendEventResponse struct {
	Errors     []AccountWorkflowInstanceSendEventResponseError    `json:"errors,required"`
	Messages   []AccountWorkflowInstanceSendEventResponseMessage  `json:"messages,required"`
	Success    AccountWorkflowInstanceSendEventResponseSuccess    `json:"success,required"`
	Result     interface{}                                        `json:"result"`
	ResultInfo AccountWorkflowInstanceSendEventResponseResultInfo `json:"result_info"`
	JSON       accountWorkflowInstanceSendEventResponseJSON       `json:"-"`
}

// accountWorkflowInstanceSendEventResponseJSON contains the JSON metadata for the
// struct [AccountWorkflowInstanceSendEventResponse]
type accountWorkflowInstanceSendEventResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceSendEventResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceSendEventResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceSendEventResponseError struct {
	Code    float64                                           `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountWorkflowInstanceSendEventResponseErrorJSON `json:"-"`
}

// accountWorkflowInstanceSendEventResponseErrorJSON contains the JSON metadata for
// the struct [AccountWorkflowInstanceSendEventResponseError]
type accountWorkflowInstanceSendEventResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceSendEventResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceSendEventResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceSendEventResponseMessage struct {
	Code    float64                                             `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountWorkflowInstanceSendEventResponseMessageJSON `json:"-"`
}

// accountWorkflowInstanceSendEventResponseMessageJSON contains the JSON metadata
// for the struct [AccountWorkflowInstanceSendEventResponseMessage]
type accountWorkflowInstanceSendEventResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceSendEventResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceSendEventResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceSendEventResponseSuccess bool

const (
	AccountWorkflowInstanceSendEventResponseSuccessTrue AccountWorkflowInstanceSendEventResponseSuccess = true
)

func (r AccountWorkflowInstanceSendEventResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceSendEventResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkflowInstanceSendEventResponseResultInfo struct {
	Count      float64                                                `json:"count,required"`
	Page       float64                                                `json:"page,required"`
	PerPage    float64                                                `json:"per_page,required"`
	TotalCount float64                                                `json:"total_count,required"`
	JSON       accountWorkflowInstanceSendEventResponseResultInfoJSON `json:"-"`
}

// accountWorkflowInstanceSendEventResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountWorkflowInstanceSendEventResponseResultInfo]
type accountWorkflowInstanceSendEventResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceSendEventResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceSendEventResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceUpdateStatusResponse struct {
	Errors     []AccountWorkflowInstanceUpdateStatusResponseError    `json:"errors,required"`
	Messages   []AccountWorkflowInstanceUpdateStatusResponseMessage  `json:"messages,required"`
	Result     AccountWorkflowInstanceUpdateStatusResponseResult     `json:"result,required"`
	Success    AccountWorkflowInstanceUpdateStatusResponseSuccess    `json:"success,required"`
	ResultInfo AccountWorkflowInstanceUpdateStatusResponseResultInfo `json:"result_info"`
	JSON       accountWorkflowInstanceUpdateStatusResponseJSON       `json:"-"`
}

// accountWorkflowInstanceUpdateStatusResponseJSON contains the JSON metadata for
// the struct [AccountWorkflowInstanceUpdateStatusResponse]
type accountWorkflowInstanceUpdateStatusResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceUpdateStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceUpdateStatusResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceUpdateStatusResponseError struct {
	Code    float64                                              `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accountWorkflowInstanceUpdateStatusResponseErrorJSON `json:"-"`
}

// accountWorkflowInstanceUpdateStatusResponseErrorJSON contains the JSON metadata
// for the struct [AccountWorkflowInstanceUpdateStatusResponseError]
type accountWorkflowInstanceUpdateStatusResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceUpdateStatusResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceUpdateStatusResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceUpdateStatusResponseMessage struct {
	Code    float64                                                `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountWorkflowInstanceUpdateStatusResponseMessageJSON `json:"-"`
}

// accountWorkflowInstanceUpdateStatusResponseMessageJSON contains the JSON
// metadata for the struct [AccountWorkflowInstanceUpdateStatusResponseMessage]
type accountWorkflowInstanceUpdateStatusResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceUpdateStatusResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceUpdateStatusResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceUpdateStatusResponseResult struct {
	Status AccountWorkflowInstanceUpdateStatusResponseResultStatus `json:"status,required"`
	// Accepts ISO 8601 with no timezone offsets and in UTC.
	Timestamp time.Time                                             `json:"timestamp,required" format:"date-time"`
	JSON      accountWorkflowInstanceUpdateStatusResponseResultJSON `json:"-"`
}

// accountWorkflowInstanceUpdateStatusResponseResultJSON contains the JSON metadata
// for the struct [AccountWorkflowInstanceUpdateStatusResponseResult]
type accountWorkflowInstanceUpdateStatusResponseResultJSON struct {
	Status      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceUpdateStatusResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceUpdateStatusResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceUpdateStatusResponseResultStatus string

const (
	AccountWorkflowInstanceUpdateStatusResponseResultStatusQueued          AccountWorkflowInstanceUpdateStatusResponseResultStatus = "queued"
	AccountWorkflowInstanceUpdateStatusResponseResultStatusRunning         AccountWorkflowInstanceUpdateStatusResponseResultStatus = "running"
	AccountWorkflowInstanceUpdateStatusResponseResultStatusPaused          AccountWorkflowInstanceUpdateStatusResponseResultStatus = "paused"
	AccountWorkflowInstanceUpdateStatusResponseResultStatusErrored         AccountWorkflowInstanceUpdateStatusResponseResultStatus = "errored"
	AccountWorkflowInstanceUpdateStatusResponseResultStatusTerminated      AccountWorkflowInstanceUpdateStatusResponseResultStatus = "terminated"
	AccountWorkflowInstanceUpdateStatusResponseResultStatusComplete        AccountWorkflowInstanceUpdateStatusResponseResultStatus = "complete"
	AccountWorkflowInstanceUpdateStatusResponseResultStatusWaitingForPause AccountWorkflowInstanceUpdateStatusResponseResultStatus = "waitingForPause"
	AccountWorkflowInstanceUpdateStatusResponseResultStatusWaiting         AccountWorkflowInstanceUpdateStatusResponseResultStatus = "waiting"
)

func (r AccountWorkflowInstanceUpdateStatusResponseResultStatus) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceUpdateStatusResponseResultStatusQueued, AccountWorkflowInstanceUpdateStatusResponseResultStatusRunning, AccountWorkflowInstanceUpdateStatusResponseResultStatusPaused, AccountWorkflowInstanceUpdateStatusResponseResultStatusErrored, AccountWorkflowInstanceUpdateStatusResponseResultStatusTerminated, AccountWorkflowInstanceUpdateStatusResponseResultStatusComplete, AccountWorkflowInstanceUpdateStatusResponseResultStatusWaitingForPause, AccountWorkflowInstanceUpdateStatusResponseResultStatusWaiting:
		return true
	}
	return false
}

type AccountWorkflowInstanceUpdateStatusResponseSuccess bool

const (
	AccountWorkflowInstanceUpdateStatusResponseSuccessTrue AccountWorkflowInstanceUpdateStatusResponseSuccess = true
)

func (r AccountWorkflowInstanceUpdateStatusResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceUpdateStatusResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkflowInstanceUpdateStatusResponseResultInfo struct {
	Count      float64                                                   `json:"count,required"`
	Page       float64                                                   `json:"page,required"`
	PerPage    float64                                                   `json:"per_page,required"`
	TotalCount float64                                                   `json:"total_count,required"`
	JSON       accountWorkflowInstanceUpdateStatusResponseResultInfoJSON `json:"-"`
}

// accountWorkflowInstanceUpdateStatusResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountWorkflowInstanceUpdateStatusResponseResultInfo]
type accountWorkflowInstanceUpdateStatusResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowInstanceUpdateStatusResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowInstanceUpdateStatusResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowInstanceNewParams struct {
	InstanceID        param.Field[string]      `json:"instance_id"`
	InstanceRetention param.Field[interface{}] `json:"instance_retention"`
	Params            param.Field[interface{}] `json:"params"`
}

func (r AccountWorkflowInstanceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkflowInstanceListParams struct {
	// Accepts ISO 8601 with no timezone offsets and in UTC.
	DateEnd param.Field[time.Time] `query:"date_end" format:"date-time"`
	// Accepts ISO 8601 with no timezone offsets and in UTC.
	DateStart param.Field[time.Time]                               `query:"date_start" format:"date-time"`
	Page      param.Field[float64]                                 `query:"page"`
	PerPage   param.Field[float64]                                 `query:"per_page"`
	Status    param.Field[AccountWorkflowInstanceListParamsStatus] `query:"status"`
}

// URLQuery serializes [AccountWorkflowInstanceListParams]'s query parameters as
// `url.Values`.
func (r AccountWorkflowInstanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountWorkflowInstanceListParamsStatus string

const (
	AccountWorkflowInstanceListParamsStatusQueued          AccountWorkflowInstanceListParamsStatus = "queued"
	AccountWorkflowInstanceListParamsStatusRunning         AccountWorkflowInstanceListParamsStatus = "running"
	AccountWorkflowInstanceListParamsStatusPaused          AccountWorkflowInstanceListParamsStatus = "paused"
	AccountWorkflowInstanceListParamsStatusErrored         AccountWorkflowInstanceListParamsStatus = "errored"
	AccountWorkflowInstanceListParamsStatusTerminated      AccountWorkflowInstanceListParamsStatus = "terminated"
	AccountWorkflowInstanceListParamsStatusComplete        AccountWorkflowInstanceListParamsStatus = "complete"
	AccountWorkflowInstanceListParamsStatusWaitingForPause AccountWorkflowInstanceListParamsStatus = "waitingForPause"
	AccountWorkflowInstanceListParamsStatusWaiting         AccountWorkflowInstanceListParamsStatus = "waiting"
)

func (r AccountWorkflowInstanceListParamsStatus) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceListParamsStatusQueued, AccountWorkflowInstanceListParamsStatusRunning, AccountWorkflowInstanceListParamsStatusPaused, AccountWorkflowInstanceListParamsStatusErrored, AccountWorkflowInstanceListParamsStatusTerminated, AccountWorkflowInstanceListParamsStatusComplete, AccountWorkflowInstanceListParamsStatusWaitingForPause, AccountWorkflowInstanceListParamsStatusWaiting:
		return true
	}
	return false
}

type AccountWorkflowInstanceBatchNewParams struct {
	Body []AccountWorkflowInstanceBatchNewParamsBody `json:"body"`
}

func (r AccountWorkflowInstanceBatchNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountWorkflowInstanceBatchNewParamsBody struct {
	InstanceID        param.Field[string]      `json:"instance_id"`
	InstanceRetention param.Field[interface{}] `json:"instance_retention"`
	Params            param.Field[interface{}] `json:"params"`
}

func (r AccountWorkflowInstanceBatchNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkflowInstanceSendEventParams struct {
	Body interface{} `json:"body"`
}

func (r AccountWorkflowInstanceSendEventParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountWorkflowInstanceUpdateStatusParams struct {
	// Apply action to instance.
	Status param.Field[AccountWorkflowInstanceUpdateStatusParamsStatus] `json:"status,required"`
}

func (r AccountWorkflowInstanceUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Apply action to instance.
type AccountWorkflowInstanceUpdateStatusParamsStatus string

const (
	AccountWorkflowInstanceUpdateStatusParamsStatusResume    AccountWorkflowInstanceUpdateStatusParamsStatus = "resume"
	AccountWorkflowInstanceUpdateStatusParamsStatusPause     AccountWorkflowInstanceUpdateStatusParamsStatus = "pause"
	AccountWorkflowInstanceUpdateStatusParamsStatusTerminate AccountWorkflowInstanceUpdateStatusParamsStatus = "terminate"
)

func (r AccountWorkflowInstanceUpdateStatusParamsStatus) IsKnown() bool {
	switch r {
	case AccountWorkflowInstanceUpdateStatusParamsStatusResume, AccountWorkflowInstanceUpdateStatusParamsStatusPause, AccountWorkflowInstanceUpdateStatusParamsStatusTerminate:
		return true
	}
	return false
}
