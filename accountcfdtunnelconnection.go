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

// AccountCfdTunnelConnectionService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCfdTunnelConnectionService] method instead.
type AccountCfdTunnelConnectionService struct {
	Options []option.RequestOption
}

// NewAccountCfdTunnelConnectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCfdTunnelConnectionService(opts ...option.RequestOption) (r *AccountCfdTunnelConnectionService) {
	r = &AccountCfdTunnelConnectionService{}
	r.Options = opts
	return
}

// Fetches connection details for a Cloudflare Tunnel.
func (r *AccountCfdTunnelConnectionService) List(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *AccountCfdTunnelConnectionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connections", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Removes a connection (aka Cloudflare Tunnel Connector) from a Cloudflare Tunnel
// independently of its current state. If no connector id (client_id) is provided
// all connectors will be removed. We recommend running this command after rotating
// tokens.
func (r *AccountCfdTunnelConnectionService) Cleanup(ctx context.Context, accountID string, tunnelID string, body AccountCfdTunnelConnectionCleanupParams, opts ...option.RequestOption) (res *AccountCfdTunnelConnectionCleanupResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connections", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AccountCfdTunnelConnectionListResponse struct {
	Errors   []MessagesTunnelItem `json:"errors,required"`
	Messages []MessagesTunnelItem `json:"messages,required"`
	Result   []TunnelClient       `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccountCfdTunnelConnectionListResponseSuccess    `json:"success,required"`
	ResultInfo AccountCfdTunnelConnectionListResponseResultInfo `json:"result_info"`
	JSON       accountCfdTunnelConnectionListResponseJSON       `json:"-"`
}

// accountCfdTunnelConnectionListResponseJSON contains the JSON metadata for the
// struct [AccountCfdTunnelConnectionListResponse]
type accountCfdTunnelConnectionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConnectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCfdTunnelConnectionListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountCfdTunnelConnectionListResponseSuccess bool

const (
	AccountCfdTunnelConnectionListResponseSuccessTrue AccountCfdTunnelConnectionListResponseSuccess = true
)

func (r AccountCfdTunnelConnectionListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountCfdTunnelConnectionListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountCfdTunnelConnectionListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       accountCfdTunnelConnectionListResponseResultInfoJSON `json:"-"`
}

// accountCfdTunnelConnectionListResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountCfdTunnelConnectionListResponseResultInfo]
type accountCfdTunnelConnectionListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConnectionListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCfdTunnelConnectionListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountCfdTunnelConnectionCleanupResponse struct {
	Errors   []AccountCfdTunnelConnectionCleanupResponseError   `json:"errors,required"`
	Messages []AccountCfdTunnelConnectionCleanupResponseMessage `json:"messages,required"`
	Result   interface{}                                        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AccountCfdTunnelConnectionCleanupResponseSuccess `json:"success,required"`
	JSON    accountCfdTunnelConnectionCleanupResponseJSON    `json:"-"`
}

// accountCfdTunnelConnectionCleanupResponseJSON contains the JSON metadata for the
// struct [AccountCfdTunnelConnectionCleanupResponse]
type accountCfdTunnelConnectionCleanupResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConnectionCleanupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCfdTunnelConnectionCleanupResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCfdTunnelConnectionCleanupResponseError struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           AccountCfdTunnelConnectionCleanupResponseErrorsSource `json:"source"`
	JSON             accountCfdTunnelConnectionCleanupResponseErrorJSON    `json:"-"`
}

// accountCfdTunnelConnectionCleanupResponseErrorJSON contains the JSON metadata
// for the struct [AccountCfdTunnelConnectionCleanupResponseError]
type accountCfdTunnelConnectionCleanupResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountCfdTunnelConnectionCleanupResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCfdTunnelConnectionCleanupResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountCfdTunnelConnectionCleanupResponseErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    accountCfdTunnelConnectionCleanupResponseErrorsSourceJSON `json:"-"`
}

// accountCfdTunnelConnectionCleanupResponseErrorsSourceJSON contains the JSON
// metadata for the struct [AccountCfdTunnelConnectionCleanupResponseErrorsSource]
type accountCfdTunnelConnectionCleanupResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConnectionCleanupResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCfdTunnelConnectionCleanupResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountCfdTunnelConnectionCleanupResponseMessage struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           AccountCfdTunnelConnectionCleanupResponseMessagesSource `json:"source"`
	JSON             accountCfdTunnelConnectionCleanupResponseMessageJSON    `json:"-"`
}

// accountCfdTunnelConnectionCleanupResponseMessageJSON contains the JSON metadata
// for the struct [AccountCfdTunnelConnectionCleanupResponseMessage]
type accountCfdTunnelConnectionCleanupResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountCfdTunnelConnectionCleanupResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCfdTunnelConnectionCleanupResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountCfdTunnelConnectionCleanupResponseMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    accountCfdTunnelConnectionCleanupResponseMessagesSourceJSON `json:"-"`
}

// accountCfdTunnelConnectionCleanupResponseMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccountCfdTunnelConnectionCleanupResponseMessagesSource]
type accountCfdTunnelConnectionCleanupResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConnectionCleanupResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCfdTunnelConnectionCleanupResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountCfdTunnelConnectionCleanupResponseSuccess bool

const (
	AccountCfdTunnelConnectionCleanupResponseSuccessTrue AccountCfdTunnelConnectionCleanupResponseSuccess = true
)

func (r AccountCfdTunnelConnectionCleanupResponseSuccess) IsKnown() bool {
	switch r {
	case AccountCfdTunnelConnectionCleanupResponseSuccessTrue:
		return true
	}
	return false
}

type AccountCfdTunnelConnectionCleanupParams struct {
	// UUID of the Cloudflare Tunnel connector.
	ClientID param.Field[string] `query:"client_id" format:"uuid"`
}

// URLQuery serializes [AccountCfdTunnelConnectionCleanupParams]'s query parameters
// as `url.Values`.
func (r AccountCfdTunnelConnectionCleanupParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
