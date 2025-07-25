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

// AccountWorkerSubdomainService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerSubdomainService] method instead.
type AccountWorkerSubdomainService struct {
	Options []option.RequestOption
}

// NewAccountWorkerSubdomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountWorkerSubdomainService(opts ...option.RequestOption) (r *AccountWorkerSubdomainService) {
	r = &AccountWorkerSubdomainService{}
	r.Options = opts
	return
}

// Creates a Workers subdomain for an account.
func (r *AccountWorkerSubdomainService) New(ctx context.Context, accountID string, body AccountWorkerSubdomainNewParams, opts ...option.RequestOption) (res *AccountWorkerSubdomainNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/subdomain", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns a Workers subdomain for an account.
func (r *AccountWorkerSubdomainService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountWorkerSubdomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/subdomain", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountWorkerSubdomainNewResponse struct {
	Errors   []WorkersMessages                       `json:"errors,required"`
	Messages []WorkersMessages                       `json:"messages,required"`
	Result   AccountWorkerSubdomainNewResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerSubdomainNewResponseSuccess `json:"success,required"`
	JSON    accountWorkerSubdomainNewResponseJSON    `json:"-"`
}

// accountWorkerSubdomainNewResponseJSON contains the JSON metadata for the struct
// [AccountWorkerSubdomainNewResponse]
type accountWorkerSubdomainNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerSubdomainNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerSubdomainNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerSubdomainNewResponseResult struct {
	Subdomain string                                      `json:"subdomain,required"`
	JSON      accountWorkerSubdomainNewResponseResultJSON `json:"-"`
}

// accountWorkerSubdomainNewResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkerSubdomainNewResponseResult]
type accountWorkerSubdomainNewResponseResultJSON struct {
	Subdomain   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerSubdomainNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerSubdomainNewResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerSubdomainNewResponseSuccess bool

const (
	AccountWorkerSubdomainNewResponseSuccessTrue AccountWorkerSubdomainNewResponseSuccess = true
)

func (r AccountWorkerSubdomainNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerSubdomainNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerSubdomainGetResponse struct {
	Errors   []WorkersMessages                       `json:"errors,required"`
	Messages []WorkersMessages                       `json:"messages,required"`
	Result   AccountWorkerSubdomainGetResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerSubdomainGetResponseSuccess `json:"success,required"`
	JSON    accountWorkerSubdomainGetResponseJSON    `json:"-"`
}

// accountWorkerSubdomainGetResponseJSON contains the JSON metadata for the struct
// [AccountWorkerSubdomainGetResponse]
type accountWorkerSubdomainGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerSubdomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerSubdomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerSubdomainGetResponseResult struct {
	Subdomain string                                      `json:"subdomain,required"`
	JSON      accountWorkerSubdomainGetResponseResultJSON `json:"-"`
}

// accountWorkerSubdomainGetResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkerSubdomainGetResponseResult]
type accountWorkerSubdomainGetResponseResultJSON struct {
	Subdomain   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerSubdomainGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerSubdomainGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerSubdomainGetResponseSuccess bool

const (
	AccountWorkerSubdomainGetResponseSuccessTrue AccountWorkerSubdomainGetResponseSuccess = true
)

func (r AccountWorkerSubdomainGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerSubdomainGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerSubdomainNewParams struct {
	Subdomain param.Field[string] `json:"subdomain,required"`
}

func (r AccountWorkerSubdomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
