// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountGatewayProxyEndpointService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountGatewayProxyEndpointService] method instead.
type AccountGatewayProxyEndpointService struct {
	Options []option.RequestOption
}

// NewAccountGatewayProxyEndpointService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountGatewayProxyEndpointService(opts ...option.RequestOption) (r *AccountGatewayProxyEndpointService) {
	r = &AccountGatewayProxyEndpointService{}
	r.Options = opts
	return
}

// Creates a new Zero Trust Gateway proxy endpoint.
func (r *AccountGatewayProxyEndpointService) New(ctx context.Context, accountID string, body AccountGatewayProxyEndpointNewParams, opts ...option.RequestOption) (res *SingleResponseProxy, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single Zero Trust Gateway proxy endpoint.
func (r *AccountGatewayProxyEndpointService) Get(ctx context.Context, accountID string, proxyEndpointID string, opts ...option.RequestOption) (res *AccountGatewayProxyEndpointGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if proxyEndpointID == "" {
		err = errors.New("missing required proxy_endpoint_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints/%s", accountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Zero Trust Gateway proxy endpoint.
func (r *AccountGatewayProxyEndpointService) Update(ctx context.Context, accountID string, proxyEndpointID string, body AccountGatewayProxyEndpointUpdateParams, opts ...option.RequestOption) (res *SingleResponseProxy, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if proxyEndpointID == "" {
		err = errors.New("missing required proxy_endpoint_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints/%s", accountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetches all Zero Trust Gateway proxy endpoints for an account.
func (r *AccountGatewayProxyEndpointService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *SingleResponseProxy, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a configured Zero Trust Gateway proxy endpoint.
func (r *AccountGatewayProxyEndpointService) Delete(ctx context.Context, accountID string, proxyEndpointID string, body AccountGatewayProxyEndpointDeleteParams, opts ...option.RequestOption) (res *ZeroTrustGatewayEmptyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if proxyEndpointID == "" {
		err = errors.New("missing required proxy_endpoint_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints/%s", accountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type Endpoint struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string       `json:"subdomain"`
	UpdatedAt time.Time    `json:"updated_at" format:"date-time"`
	JSON      endpointJSON `json:"-"`
}

// endpointJSON contains the JSON metadata for the struct [Endpoint]
type endpointJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Endpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointJSON) RawJSON() string {
	return r.raw
}

type SingleResponseProxy struct {
	Result Endpoint                `json:"result"`
	JSON   singleResponseProxyJSON `json:"-"`
	APIResponseSingleZeroTrustGateway
}

// singleResponseProxyJSON contains the JSON metadata for the struct
// [SingleResponseProxy]
type singleResponseProxyJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseProxy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseProxyJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayProxyEndpointGetResponse struct {
	Result []Endpoint                                 `json:"result"`
	JSON   accountGatewayProxyEndpointGetResponseJSON `json:"-"`
	APIResponseCollectionZeroTrustGateway
}

// accountGatewayProxyEndpointGetResponseJSON contains the JSON metadata for the
// struct [AccountGatewayProxyEndpointGetResponse]
type accountGatewayProxyEndpointGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayProxyEndpointGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayProxyEndpointNewParams struct {
	// A list of CIDRs to restrict ingress connections.
	IPs param.Field[[]string] `json:"ips,required"`
	// The name of the proxy endpoint.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountGatewayProxyEndpointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayProxyEndpointUpdateParams struct {
	// A list of CIDRs to restrict ingress connections.
	IPs param.Field[[]string] `json:"ips"`
	// The name of the proxy endpoint.
	Name param.Field[string] `json:"name"`
}

func (r AccountGatewayProxyEndpointUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayProxyEndpointDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountGatewayProxyEndpointDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
