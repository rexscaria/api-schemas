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

// AccountWorkerScriptSubdomainService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerScriptSubdomainService] method instead.
type AccountWorkerScriptSubdomainService struct {
	Options []option.RequestOption
}

// NewAccountWorkerScriptSubdomainService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptSubdomainService(opts ...option.RequestOption) (r *AccountWorkerScriptSubdomainService) {
	r = &AccountWorkerScriptSubdomainService{}
	r.Options = opts
	return
}

// Get if the Worker is available on the workers.dev subdomain.
func (r *AccountWorkerScriptSubdomainService) Get(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerScriptSubdomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/subdomain", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable the Worker on the workers.dev subdomain.
func (r *AccountWorkerScriptSubdomainService) Post(ctx context.Context, accountID string, scriptName string, body AccountWorkerScriptSubdomainPostParams, opts ...option.RequestOption) (res *AccountWorkerScriptSubdomainPostResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/subdomain", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountWorkerScriptSubdomainGetResponse struct {
	// Whether the Worker is available on the workers.dev subdomain.
	Enabled bool `json:"enabled"`
	// Whether the Worker's Preview URLs should be available on the workers.dev
	// subdomain.
	PreviewsEnabled bool                                        `json:"previews_enabled"`
	JSON            accountWorkerScriptSubdomainGetResponseJSON `json:"-"`
}

// accountWorkerScriptSubdomainGetResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptSubdomainGetResponse]
type accountWorkerScriptSubdomainGetResponseJSON struct {
	Enabled         apijson.Field
	PreviewsEnabled apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountWorkerScriptSubdomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptSubdomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptSubdomainPostResponse struct {
	// Whether the Worker is available on the workers.dev subdomain.
	Enabled bool `json:"enabled"`
	// Whether the Worker's Preview URLs should be available on the workers.dev
	// subdomain.
	PreviewsEnabled bool                                         `json:"previews_enabled"`
	JSON            accountWorkerScriptSubdomainPostResponseJSON `json:"-"`
}

// accountWorkerScriptSubdomainPostResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptSubdomainPostResponse]
type accountWorkerScriptSubdomainPostResponseJSON struct {
	Enabled         apijson.Field
	PreviewsEnabled apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountWorkerScriptSubdomainPostResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptSubdomainPostResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptSubdomainPostParams struct {
	// Whether the Worker should be available on the workers.dev subdomain.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Whether the Worker's Preview URLs should be available on the workers.dev
	// subdomain.
	PreviewsEnabled param.Field[bool] `json:"previews_enabled"`
}

func (r AccountWorkerScriptSubdomainPostParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
