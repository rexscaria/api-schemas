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

// AccountWorkerScriptScheduleService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerScriptScheduleService] method instead.
type AccountWorkerScriptScheduleService struct {
	Options []option.RequestOption
}

// NewAccountWorkerScriptScheduleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptScheduleService(opts ...option.RequestOption) (r *AccountWorkerScriptScheduleService) {
	r = &AccountWorkerScriptScheduleService{}
	r.Options = opts
	return
}

// Updates Cron Triggers for a Worker.
func (r *AccountWorkerScriptScheduleService) Update(ctx context.Context, accountID string, scriptName string, body AccountWorkerScriptScheduleUpdateParams, opts ...option.RequestOption) (res *CronTriggerCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches Cron Triggers for a Worker.
func (r *AccountWorkerScriptScheduleService) Get(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *CronTriggerCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CronObject struct {
	CreatedOn  string         `json:"created_on"`
	Cron       string         `json:"cron"`
	ModifiedOn string         `json:"modified_on"`
	JSON       cronObjectJSON `json:"-"`
}

// cronObjectJSON contains the JSON metadata for the struct [CronObject]
type cronObjectJSON struct {
	CreatedOn   apijson.Field
	Cron        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CronObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cronObjectJSON) RawJSON() string {
	return r.raw
}

type CronObjectParam struct {
	Cron param.Field[string] `json:"cron"`
}

func (r CronObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CronTriggerCollection struct {
	Result CronTriggerCollectionResult `json:"result"`
	JSON   cronTriggerCollectionJSON   `json:"-"`
	CommonResponseWorkers
}

// cronTriggerCollectionJSON contains the JSON metadata for the struct
// [CronTriggerCollection]
type cronTriggerCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CronTriggerCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cronTriggerCollectionJSON) RawJSON() string {
	return r.raw
}

type CronTriggerCollectionResult struct {
	Schedules []CronObject                    `json:"schedules"`
	JSON      cronTriggerCollectionResultJSON `json:"-"`
}

// cronTriggerCollectionResultJSON contains the JSON metadata for the struct
// [CronTriggerCollectionResult]
type cronTriggerCollectionResultJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CronTriggerCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cronTriggerCollectionResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptScheduleUpdateParams struct {
	Body []CronObjectParam `json:"body,required"`
}

func (r AccountWorkerScriptScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
