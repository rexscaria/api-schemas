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

// AccountWorkerDispatchNamespaceService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerDispatchNamespaceService] method instead.
type AccountWorkerDispatchNamespaceService struct {
	Options []option.RequestOption
	Scripts *AccountWorkerDispatchNamespaceScriptService
}

// NewAccountWorkerDispatchNamespaceService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerDispatchNamespaceService(opts ...option.RequestOption) (r *AccountWorkerDispatchNamespaceService) {
	r = &AccountWorkerDispatchNamespaceService{}
	r.Options = opts
	r.Scripts = NewAccountWorkerDispatchNamespaceScriptService(opts...)
	return
}

// Create a new Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceService) New(ctx context.Context, accountID string, body AccountWorkerDispatchNamespaceNewParams, opts ...option.RequestOption) (res *SingleNamespaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceService) Get(ctx context.Context, accountID string, dispatchNamespace string, opts ...option.RequestOption) (res *SingleNamespaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s", accountID, dispatchNamespace)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetch a list of Workers for Platforms namespaces.
func (r *AccountWorkerDispatchNamespaceService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceService) Delete(ctx context.Context, accountID string, dispatchNamespace string, opts ...option.RequestOption) (res *NullResult, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s", accountID, dispatchNamespace)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type NamespaceResponse struct {
	// Identifier
	CreatedBy string `json:"created_by"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Identifier
	ModifiedBy string `json:"modified_by"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// API Resource UUID tag.
	NamespaceID string `json:"namespace_id"`
	// Name of the Workers for Platforms dispatch namespace.
	NamespaceName string `json:"namespace_name"`
	// The current number of scripts in this Dispatch Namespace
	ScriptCount int64                 `json:"script_count"`
	JSON        namespaceResponseJSON `json:"-"`
}

// namespaceResponseJSON contains the JSON metadata for the struct
// [NamespaceResponse]
type namespaceResponseJSON struct {
	CreatedBy     apijson.Field
	CreatedOn     apijson.Field
	ModifiedBy    apijson.Field
	ModifiedOn    apijson.Field
	NamespaceID   apijson.Field
	NamespaceName apijson.Field
	ScriptCount   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *NamespaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceResponseJSON) RawJSON() string {
	return r.raw
}

type NullResult struct {
	Result interface{}    `json:"result,nullable"`
	JSON   nullResultJSON `json:"-"`
	CommonResponseWorkers
}

// nullResultJSON contains the JSON metadata for the struct [NullResult]
type nullResultJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NullResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nullResultJSON) RawJSON() string {
	return r.raw
}

type SingleNamespaceResponse struct {
	Result NamespaceResponse           `json:"result"`
	JSON   singleNamespaceResponseJSON `json:"-"`
	CommonResponseWorkers
}

// singleNamespaceResponseJSON contains the JSON metadata for the struct
// [SingleNamespaceResponse]
type singleNamespaceResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleNamespaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleNamespaceResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDispatchNamespaceListResponse struct {
	Result []NamespaceResponse                            `json:"result"`
	JSON   accountWorkerDispatchNamespaceListResponseJSON `json:"-"`
	CommonResponseWorkers
}

// accountWorkerDispatchNamespaceListResponseJSON contains the JSON metadata for
// the struct [AccountWorkerDispatchNamespaceListResponse]
type accountWorkerDispatchNamespaceListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDispatchNamespaceNewParams struct {
	// The name of the dispatch namespace
	Name param.Field[string] `json:"name"`
}

func (r AccountWorkerDispatchNamespaceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
