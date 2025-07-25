// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountWorkerDomainService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerDomainService] method instead.
type AccountWorkerDomainService struct {
	Options []option.RequestOption
}

// NewAccountWorkerDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountWorkerDomainService(opts ...option.RequestOption) (r *AccountWorkerDomainService) {
	r = &AccountWorkerDomainService{}
	r.Options = opts
	return
}

// Gets a Worker domain.
func (r *AccountWorkerDomainService) Get(ctx context.Context, accountID string, domainID string, opts ...option.RequestOption) (res *DomainResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/domains/%s", accountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all Worker Domains for an account.
func (r *AccountWorkerDomainService) List(ctx context.Context, accountID string, query AccountWorkerDomainListParams, opts ...option.RequestOption) (res *AccountWorkerDomainListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Attaches a Worker to a zone and hostname.
func (r *AccountWorkerDomainService) Attach(ctx context.Context, accountID string, body AccountWorkerDomainAttachParams, opts ...option.RequestOption) (res *DomainResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Detaches a Worker from a zone and hostname.
func (r *AccountWorkerDomainService) Detach(ctx context.Context, accountID string, domainID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/domains/%s", accountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type DomainResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DomainResponseSuccess `json:"success,required"`
	Result  DomainWorker          `json:"result"`
	JSON    domainResponseJSON    `json:"-"`
}

// domainResponseJSON contains the JSON metadata for the struct [DomainResponse]
type domainResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DomainResponseSuccess bool

const (
	DomainResponseSuccessTrue DomainResponseSuccess = true
)

func (r DomainResponseSuccess) IsKnown() bool {
	switch r {
	case DomainResponseSuccessTrue:
		return true
	}
	return false
}

type DomainWorker struct {
	// Identifer of the Worker Domain.
	ID string `json:"id"`
	// Worker environment associated with the zone and hostname.
	Environment string `json:"environment"`
	// Hostname of the Worker Domain.
	Hostname string `json:"hostname"`
	// Worker service associated with the zone and hostname.
	Service string `json:"service"`
	// Identifier of the zone.
	ZoneID string `json:"zone_id"`
	// Name of the zone.
	ZoneName string           `json:"zone_name"`
	JSON     domainWorkerJSON `json:"-"`
}

// domainWorkerJSON contains the JSON metadata for the struct [DomainWorker]
type domainWorkerJSON struct {
	ID          apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainWorkerJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDomainListResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountWorkerDomainListResponseSuccess `json:"success,required"`
	Result  []DomainWorker                         `json:"result"`
	JSON    accountWorkerDomainListResponseJSON    `json:"-"`
}

// accountWorkerDomainListResponseJSON contains the JSON metadata for the struct
// [AccountWorkerDomainListResponse]
type accountWorkerDomainListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDomainListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerDomainListResponseSuccess bool

const (
	AccountWorkerDomainListResponseSuccessTrue AccountWorkerDomainListResponseSuccess = true
)

func (r AccountWorkerDomainListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerDomainListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerDomainListParams struct {
	// Worker environment associated with the zone and hostname.
	Environment param.Field[string] `query:"environment"`
	// Hostname of the Worker Domain.
	Hostname param.Field[string] `query:"hostname"`
	// Worker service associated with the zone and hostname.
	Service param.Field[string] `query:"service"`
	// Identifier of the zone.
	ZoneID param.Field[string] `query:"zone_id"`
	// Name of the zone.
	ZoneName param.Field[string] `query:"zone_name"`
}

// URLQuery serializes [AccountWorkerDomainListParams]'s query parameters as
// `url.Values`.
func (r AccountWorkerDomainListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountWorkerDomainAttachParams struct {
	// Worker environment associated with the zone and hostname.
	Environment param.Field[string] `json:"environment,required"`
	// Hostname of the Worker Domain.
	Hostname param.Field[string] `json:"hostname,required"`
	// Worker service associated with the zone and hostname.
	Service param.Field[string] `json:"service,required"`
	// Identifier of the zone.
	ZoneID param.Field[string] `json:"zone_id,required"`
}

func (r AccountWorkerDomainAttachParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
