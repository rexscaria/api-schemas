// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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
func (r *AccountWorkerSubdomainService) New(ctx context.Context, accountID string, body AccountWorkerSubdomainNewParams, opts ...option.RequestOption) (res *SubdomainResponse, err error) {
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
func (r *AccountWorkerSubdomainService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *SubdomainResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/subdomain", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SubdomainObject struct {
	Subdomain string              `json:"subdomain"`
	JSON      subdomainObjectJSON `json:"-"`
}

// subdomainObjectJSON contains the JSON metadata for the struct [SubdomainObject]
type subdomainObjectJSON struct {
	Subdomain   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainObjectJSON) RawJSON() string {
	return r.raw
}

type SubdomainObjectParam struct {
	Subdomain param.Field[string] `json:"subdomain"`
}

func (r SubdomainObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubdomainResponse struct {
	Result SubdomainObject       `json:"result"`
	JSON   subdomainResponseJSON `json:"-"`
	CommonResponseWorkers
}

// subdomainResponseJSON contains the JSON metadata for the struct
// [SubdomainResponse]
type subdomainResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerSubdomainNewParams struct {
	SubdomainObject SubdomainObjectParam `json:"subdomain_object,required"`
}

func (r AccountWorkerSubdomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SubdomainObject)
}
