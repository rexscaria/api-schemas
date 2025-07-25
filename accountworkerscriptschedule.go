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
func (r *AccountWorkerScriptScheduleService) Update(ctx context.Context, accountID string, scriptName string, body AccountWorkerScriptScheduleUpdateParams, opts ...option.RequestOption) (res *AccountWorkerScriptScheduleUpdateResponse, err error) {
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
func (r *AccountWorkerScriptScheduleService) Get(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerScriptScheduleGetResponse, err error) {
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

type AccountWorkerScriptScheduleUpdateResponse struct {
	Errors   []WorkersMessages                               `json:"errors,required"`
	Messages []WorkersMessages                               `json:"messages,required"`
	Result   AccountWorkerScriptScheduleUpdateResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerScriptScheduleUpdateResponseSuccess `json:"success,required"`
	JSON    accountWorkerScriptScheduleUpdateResponseJSON    `json:"-"`
}

// accountWorkerScriptScheduleUpdateResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptScheduleUpdateResponse]
type accountWorkerScriptScheduleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptScheduleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptScheduleUpdateResponseResult struct {
	Schedules []AccountWorkerScriptScheduleUpdateResponseResultSchedule `json:"schedules,required"`
	JSON      accountWorkerScriptScheduleUpdateResponseResultJSON       `json:"-"`
}

// accountWorkerScriptScheduleUpdateResponseResultJSON contains the JSON metadata
// for the struct [AccountWorkerScriptScheduleUpdateResponseResult]
type accountWorkerScriptScheduleUpdateResponseResultJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptScheduleUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptScheduleUpdateResponseResultSchedule struct {
	Cron       string                                                      `json:"cron,required"`
	CreatedOn  string                                                      `json:"created_on"`
	ModifiedOn string                                                      `json:"modified_on"`
	JSON       accountWorkerScriptScheduleUpdateResponseResultScheduleJSON `json:"-"`
}

// accountWorkerScriptScheduleUpdateResponseResultScheduleJSON contains the JSON
// metadata for the struct
// [AccountWorkerScriptScheduleUpdateResponseResultSchedule]
type accountWorkerScriptScheduleUpdateResponseResultScheduleJSON struct {
	Cron        apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleUpdateResponseResultSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptScheduleUpdateResponseResultScheduleJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerScriptScheduleUpdateResponseSuccess bool

const (
	AccountWorkerScriptScheduleUpdateResponseSuccessTrue AccountWorkerScriptScheduleUpdateResponseSuccess = true
)

func (r AccountWorkerScriptScheduleUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerScriptScheduleUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptScheduleGetResponse struct {
	Errors   []WorkersMessages                            `json:"errors,required"`
	Messages []WorkersMessages                            `json:"messages,required"`
	Result   AccountWorkerScriptScheduleGetResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerScriptScheduleGetResponseSuccess `json:"success,required"`
	JSON    accountWorkerScriptScheduleGetResponseJSON    `json:"-"`
}

// accountWorkerScriptScheduleGetResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptScheduleGetResponse]
type accountWorkerScriptScheduleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptScheduleGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptScheduleGetResponseResult struct {
	Schedules []AccountWorkerScriptScheduleGetResponseResultSchedule `json:"schedules,required"`
	JSON      accountWorkerScriptScheduleGetResponseResultJSON       `json:"-"`
}

// accountWorkerScriptScheduleGetResponseResultJSON contains the JSON metadata for
// the struct [AccountWorkerScriptScheduleGetResponseResult]
type accountWorkerScriptScheduleGetResponseResultJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptScheduleGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptScheduleGetResponseResultSchedule struct {
	Cron       string                                                   `json:"cron,required"`
	CreatedOn  string                                                   `json:"created_on"`
	ModifiedOn string                                                   `json:"modified_on"`
	JSON       accountWorkerScriptScheduleGetResponseResultScheduleJSON `json:"-"`
}

// accountWorkerScriptScheduleGetResponseResultScheduleJSON contains the JSON
// metadata for the struct [AccountWorkerScriptScheduleGetResponseResultSchedule]
type accountWorkerScriptScheduleGetResponseResultScheduleJSON struct {
	Cron        apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleGetResponseResultSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptScheduleGetResponseResultScheduleJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerScriptScheduleGetResponseSuccess bool

const (
	AccountWorkerScriptScheduleGetResponseSuccessTrue AccountWorkerScriptScheduleGetResponseSuccess = true
)

func (r AccountWorkerScriptScheduleGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerScriptScheduleGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptScheduleUpdateParams struct {
	Body []AccountWorkerScriptScheduleUpdateParamsBody `json:"body,required"`
}

func (r AccountWorkerScriptScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountWorkerScriptScheduleUpdateParamsBody struct {
	Cron param.Field[string] `json:"cron,required"`
}

func (r AccountWorkerScriptScheduleUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
