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

// AccountCloudforceOneEventInsightService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneEventInsightService] method instead.
type AccountCloudforceOneEventInsightService struct {
	Options []option.RequestOption
}

// NewAccountCloudforceOneEventInsightService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCloudforceOneEventInsightService(opts ...option.RequestOption) (r *AccountCloudforceOneEventInsightService) {
	r = &AccountCloudforceOneEventInsightService{}
	r.Options = opts
	return
}

// Adds an insight to an event
func (r *AccountCloudforceOneEventInsightService) New(ctx context.Context, accountID float64, eventID string, body AccountCloudforceOneEventInsightNewParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventInsightNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/%s/insight/create", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Reads an event insight
func (r *AccountCloudforceOneEventInsightService) Get(ctx context.Context, accountID float64, eventID string, insightID string, opts ...option.RequestOption) (res *AccountCloudforceOneEventInsightGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/%s/insight/%s", accountID, eventID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an event insight
func (r *AccountCloudforceOneEventInsightService) Update(ctx context.Context, accountID float64, eventID string, insightID string, body AccountCloudforceOneEventInsightUpdateParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventInsightUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/%s/insight/%s", accountID, eventID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes an event insight
func (r *AccountCloudforceOneEventInsightService) Delete(ctx context.Context, accountID float64, eventID string, insightID string, opts ...option.RequestOption) (res *AccountCloudforceOneEventInsightDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/%s/insight/%s", accountID, eventID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountCloudforceOneEventInsightNewResponse struct {
	Result  AccountCloudforceOneEventInsightNewResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    accountCloudforceOneEventInsightNewResponseJSON   `json:"-"`
}

// accountCloudforceOneEventInsightNewResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneEventInsightNewResponse]
type accountCloudforceOneEventInsightNewResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventInsightNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventInsightNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventInsightNewResponseResult struct {
	Content string                                                `json:"content,required"`
	Uuid    string                                                `json:"uuid,required"`
	JSON    accountCloudforceOneEventInsightNewResponseResultJSON `json:"-"`
}

// accountCloudforceOneEventInsightNewResponseResultJSON contains the JSON metadata
// for the struct [AccountCloudforceOneEventInsightNewResponseResult]
type accountCloudforceOneEventInsightNewResponseResultJSON struct {
	Content     apijson.Field
	Uuid        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventInsightNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventInsightNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventInsightGetResponse struct {
	Result  AccountCloudforceOneEventInsightGetResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    accountCloudforceOneEventInsightGetResponseJSON   `json:"-"`
}

// accountCloudforceOneEventInsightGetResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneEventInsightGetResponse]
type accountCloudforceOneEventInsightGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventInsightGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventInsightGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventInsightGetResponseResult struct {
	Content string                                                `json:"content,required"`
	Uuid    string                                                `json:"uuid,required"`
	JSON    accountCloudforceOneEventInsightGetResponseResultJSON `json:"-"`
}

// accountCloudforceOneEventInsightGetResponseResultJSON contains the JSON metadata
// for the struct [AccountCloudforceOneEventInsightGetResponseResult]
type accountCloudforceOneEventInsightGetResponseResultJSON struct {
	Content     apijson.Field
	Uuid        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventInsightGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventInsightGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventInsightUpdateResponse struct {
	Result  AccountCloudforceOneEventInsightUpdateResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    accountCloudforceOneEventInsightUpdateResponseJSON   `json:"-"`
}

// accountCloudforceOneEventInsightUpdateResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneEventInsightUpdateResponse]
type accountCloudforceOneEventInsightUpdateResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventInsightUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventInsightUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventInsightUpdateResponseResult struct {
	Content string                                                   `json:"content,required"`
	Uuid    string                                                   `json:"uuid,required"`
	JSON    accountCloudforceOneEventInsightUpdateResponseResultJSON `json:"-"`
}

// accountCloudforceOneEventInsightUpdateResponseResultJSON contains the JSON
// metadata for the struct [AccountCloudforceOneEventInsightUpdateResponseResult]
type accountCloudforceOneEventInsightUpdateResponseResultJSON struct {
	Content     apijson.Field
	Uuid        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventInsightUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventInsightUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventInsightDeleteResponse struct {
	Result  AccountCloudforceOneEventInsightDeleteResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    accountCloudforceOneEventInsightDeleteResponseJSON   `json:"-"`
}

// accountCloudforceOneEventInsightDeleteResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneEventInsightDeleteResponse]
type accountCloudforceOneEventInsightDeleteResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventInsightDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventInsightDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventInsightDeleteResponseResult struct {
	Success bool                                                     `json:"success,required"`
	JSON    accountCloudforceOneEventInsightDeleteResponseResultJSON `json:"-"`
}

// accountCloudforceOneEventInsightDeleteResponseResultJSON contains the JSON
// metadata for the struct [AccountCloudforceOneEventInsightDeleteResponseResult]
type accountCloudforceOneEventInsightDeleteResponseResultJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventInsightDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventInsightDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventInsightNewParams struct {
	Content param.Field[string] `json:"content,required"`
}

func (r AccountCloudforceOneEventInsightNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCloudforceOneEventInsightUpdateParams struct {
	Content param.Field[string] `json:"content,required"`
}

func (r AccountCloudforceOneEventInsightUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
