// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountWorkerDispatchNamespaceScriptTagService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerDispatchNamespaceScriptTagService] method instead.
type AccountWorkerDispatchNamespaceScriptTagService struct {
	Options []option.RequestOption
}

// NewAccountWorkerDispatchNamespaceScriptTagService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountWorkerDispatchNamespaceScriptTagService(opts ...option.RequestOption) (r *AccountWorkerDispatchNamespaceScriptTagService) {
	r = &AccountWorkerDispatchNamespaceScriptTagService{}
	r.Options = opts
	return
}

// Delete script tag for a script uploaded to a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptTagService) Delete(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, tag string, opts ...option.RequestOption) (res *NullResult, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	if tag == "" {
		err = errors.New("missing required tag parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/tags/%s", accountID, dispatchNamespace, scriptName, tag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Fetch tags from a script uploaded to a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptTagService) Get(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptTagGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/tags", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Put a single tag on a script uploaded to a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptTagService) Put(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, tag string, opts ...option.RequestOption) (res *NullResult, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	if tag == "" {
		err = errors.New("missing required tag parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/tags/%s", accountID, dispatchNamespace, scriptName, tag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type AccountWorkerDispatchNamespaceScriptTagGetResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountWorkerDispatchNamespaceScriptTagGetResponseSuccess `json:"success,required"`
	Result  []string                                                  `json:"result"`
	JSON    accountWorkerDispatchNamespaceScriptTagGetResponseJSON    `json:"-"`
}

// accountWorkerDispatchNamespaceScriptTagGetResponseJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptTagGetResponse]
type accountWorkerDispatchNamespaceScriptTagGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptTagGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptTagGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerDispatchNamespaceScriptTagGetResponseSuccess bool

const (
	AccountWorkerDispatchNamespaceScriptTagGetResponseSuccessTrue AccountWorkerDispatchNamespaceScriptTagGetResponseSuccess = true
)

func (r AccountWorkerDispatchNamespaceScriptTagGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerDispatchNamespaceScriptTagGetResponseSuccessTrue:
		return true
	}
	return false
}
