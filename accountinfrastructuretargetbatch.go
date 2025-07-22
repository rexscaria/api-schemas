// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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
func (r *AccountInfrastructureTargetBatchService) New(ctx context.Context, accountID string, body AccountInfrastructureTargetBatchNewParams, opts ...option.RequestOption) (res *[]TargetBatch, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/batch", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Removes one or more targets.
func (r *AccountInfrastructureTargetBatchService) Delete(ctx context.Context, accountID string, body AccountInfrastructureTargetBatchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/batch", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type TargetBatch struct {
	// Target identifier
	ID string `json:"id,required" format:"uuid"`
	// Date and time at which the target was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A non-unique field that refers to a target
	Hostname string `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP IPInfoTarget `json:"ip,required"`
	// Date and time at which the target was modified
	ModifiedAt time.Time       `json:"modified_at,required" format:"date-time"`
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

type AccountInfrastructureTargetBatchNewParams struct {
	Body []AccountInfrastructureTargetBatchNewParamsBody `json:"body,required"`
}

func (r AccountInfrastructureTargetBatchNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountInfrastructureTargetBatchNewParamsBody struct {
	// A non-unique field that refers to a target. Case insensitive, maximum length of
	// 255 characters, supports the use of special characters dash and period, does not
	// support spaces, and must start and end with an alphanumeric character.
	Hostname param.Field[string] `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP param.Field[IPInfoTargetParam] `json:"ip,required"`
}

func (r AccountInfrastructureTargetBatchNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountInfrastructureTargetBatchDeleteParams struct {
	TargetIDs param.Field[[]string] `json:"target_ids,required" format:"uuid"`
}

func (r AccountInfrastructureTargetBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
