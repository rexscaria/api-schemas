// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountInfrastructureTargetBatchService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountInfrastructureTargetBatchService] method instead.
type AccountInfrastructureTargetBatchService struct {
	Options []option.RequestOption
}

// NewAccountInfrastructureTargetBatchService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountInfrastructureTargetBatchService(opts ...option.RequestOption) (r *AccountInfrastructureTargetBatchService) {
	r = &AccountInfrastructureTargetBatchService{}
	r.Options = opts
	return
}

// Adds one or more targets.
func (r *AccountInfrastructureTargetBatchService) New(ctx context.Context, accountID string, body AccountInfrastructureTargetBatchNewParams, opts ...option.RequestOption) (res *AccountInfrastructureTargetBatchNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/batch", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Removes one or more targets.
//
// Deprecated: deprecated
func (r *AccountInfrastructureTargetBatchService) Delete(ctx context.Context, accountID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return err
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/batch", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type TargetBatch struct {
	// Target identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Date and time at which the target was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A non-unique field that refers to a target
	Hostname string `json:"hostname" api:"required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP IPInfoTarget `json:"ip" api:"required"`
	// Date and time at which the target was modified
	ModifiedAt time.Time       `json:"modified_at" api:"required" format:"date-time"`
	JSON       targetBatchJSON `json:"-"`
}

// targetBatchJSON contains the JSON metadata for the struct [TargetBatch]
type targetBatchJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TargetBatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r targetBatchJSON) RawJSON() string {
	return r.raw
}

type AccountInfrastructureTargetBatchNewResponse struct {
	Errors   []MessagesInfraItem `json:"errors" api:"required"`
	Messages []MessagesInfraItem `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccountInfrastructureTargetBatchNewResponseSuccess `json:"success" api:"required"`
	Result  []TargetBatch                                      `json:"result"`
	JSON    accountInfrastructureTargetBatchNewResponseJSON    `json:"-"`
}

// accountInfrastructureTargetBatchNewResponseJSON contains the JSON metadata for
// the struct [AccountInfrastructureTargetBatchNewResponse]
type accountInfrastructureTargetBatchNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountInfrastructureTargetBatchNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountInfrastructureTargetBatchNewResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountInfrastructureTargetBatchNewResponseSuccess bool

const (
	AccountInfrastructureTargetBatchNewResponseSuccessTrue AccountInfrastructureTargetBatchNewResponseSuccess = true
)

func (r AccountInfrastructureTargetBatchNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountInfrastructureTargetBatchNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountInfrastructureTargetBatchNewParams struct {
	Body []AccountInfrastructureTargetBatchNewParamsBody `json:"body" api:"required"`
}

func (r AccountInfrastructureTargetBatchNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountInfrastructureTargetBatchNewParamsBody struct {
	// A non-unique field that refers to a target. Case insensitive, maximum length of
	// 255 characters, supports the use of special characters dash and period, does not
	// support spaces, and must start and end with an alphanumeric character.
	Hostname param.Field[string] `json:"hostname" api:"required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP param.Field[IPInfoTargetParam] `json:"ip" api:"required"`
}

func (r AccountInfrastructureTargetBatchNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
