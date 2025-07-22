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

// AccountCloudforceOneEventRelateService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneEventRelateService] method instead.
type AccountCloudforceOneEventRelateService struct {
	Options []option.RequestOption
}

// NewAccountCloudforceOneEventRelateService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCloudforceOneEventRelateService(opts ...option.RequestOption) (r *AccountCloudforceOneEventRelateService) {
	r = &AccountCloudforceOneEventRelateService{}
	r.Options = opts
	return
}

// Creates event references for a event
func (r *AccountCloudforceOneEventRelateService) New(ctx context.Context, accountID float64, eventID string, body AccountCloudforceOneEventRelateNewParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventRelateNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/relate/%s/create", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes an event reference
func (r *AccountCloudforceOneEventRelateService) Remove(ctx context.Context, accountID float64, eventID string, body AccountCloudforceOneEventRelateRemoveParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventRelateRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/relate/%s", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AccountCloudforceOneEventRelateNewResponse struct {
	Result  AccountCloudforceOneEventRelateNewResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    accountCloudforceOneEventRelateNewResponseJSON   `json:"-"`
}

// accountCloudforceOneEventRelateNewResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneEventRelateNewResponse]
type accountCloudforceOneEventRelateNewResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventRelateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventRelateNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventRelateNewResponseResult struct {
	Success bool                                                 `json:"success,required"`
	JSON    accountCloudforceOneEventRelateNewResponseResultJSON `json:"-"`
}

// accountCloudforceOneEventRelateNewResponseResultJSON contains the JSON metadata
// for the struct [AccountCloudforceOneEventRelateNewResponseResult]
type accountCloudforceOneEventRelateNewResponseResultJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventRelateNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventRelateNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventRelateRemoveResponse struct {
	Result  AccountCloudforceOneEventRelateRemoveResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    accountCloudforceOneEventRelateRemoveResponseJSON   `json:"-"`
}

// accountCloudforceOneEventRelateRemoveResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneEventRelateRemoveResponse]
type accountCloudforceOneEventRelateRemoveResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventRelateRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventRelateRemoveResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventRelateRemoveResponseResult struct {
	Success bool                                                    `json:"success,required"`
	JSON    accountCloudforceOneEventRelateRemoveResponseResultJSON `json:"-"`
}

// accountCloudforceOneEventRelateRemoveResponseResultJSON contains the JSON
// metadata for the struct [AccountCloudforceOneEventRelateRemoveResponseResult]
type accountCloudforceOneEventRelateRemoveResponseResultJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventRelateRemoveResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventRelateRemoveResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventRelateNewParams struct {
	Events param.Field[[]string] `json:"events,required"`
}

func (r AccountCloudforceOneEventRelateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCloudforceOneEventRelateRemoveParams struct {
	Events param.Field[[]string] `json:"events,required"`
}

func (r AccountCloudforceOneEventRelateRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
