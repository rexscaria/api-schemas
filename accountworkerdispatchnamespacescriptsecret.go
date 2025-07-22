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

// AccountWorkerDispatchNamespaceScriptSecretService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerDispatchNamespaceScriptSecretService] method instead.
type AccountWorkerDispatchNamespaceScriptSecretService struct {
	Options []option.RequestOption
}

// NewAccountWorkerDispatchNamespaceScriptSecretService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewAccountWorkerDispatchNamespaceScriptSecretService(opts ...option.RequestOption) (r *AccountWorkerDispatchNamespaceScriptSecretService) {
	r = &AccountWorkerDispatchNamespaceScriptSecretService{}
	r.Options = opts
	return
}

// List secrets bound to a script uploaded to a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptSecretService) List(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptSecretListResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/secrets", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove a secret from a script uploaded to a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptSecretService) Delete(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, secretName string, opts ...option.RequestOption) (res *NullResult, err error) {
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
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/secrets/%s", accountID, dispatchNamespace, scriptName, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add a secret to a script uploaded to a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptSecretService) Add(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, body AccountWorkerDispatchNamespaceScriptSecretAddParams, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptSecretAddResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/secrets", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get a given secret binding (value omitted) on a script uploaded to a Workers for
// Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptSecretService) Get(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, secretName string, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptSecretGetResponse, err error) {
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
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/secrets/%s", accountID, dispatchNamespace, scriptName, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SecretParam struct {
	// The name of this secret, this is what will be used to access it inside the
	// Worker.
	Name param.Field[string] `json:"name"`
	// The value of the secret.
	Text param.Field[string] `json:"text"`
	// The type of secret to put.
	Type param.Field[SecretType] `json:"type"`
}

func (r SecretParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of secret to put.
type SecretType string

const (
	SecretTypeSecretText SecretType = "secret_text"
)

func (r SecretType) IsKnown() bool {
	switch r {
	case SecretTypeSecretText:
		return true
	}
	return false
}

type SecretResponse struct {
	// The name of this secret, this is what will be used to access it inside the
	// Worker.
	Name string `json:"name"`
	// The type of secret.
	Type SecretResponseType `json:"type"`
	JSON secretResponseJSON `json:"-"`
}

// secretResponseJSON contains the JSON metadata for the struct [SecretResponse]
type secretResponseJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretResponseJSON) RawJSON() string {
	return r.raw
}

// The type of secret.
type SecretResponseType string

const (
	SecretResponseTypeSecretText SecretResponseType = "secret_text"
)

func (r SecretResponseType) IsKnown() bool {
	switch r {
	case SecretResponseTypeSecretText:
		return true
	}
	return false
}

type AccountWorkerDispatchNamespaceScriptSecretListResponse struct {
	Result []SecretResponse                                           `json:"result"`
	JSON   accountWorkerDispatchNamespaceScriptSecretListResponseJSON `json:"-"`
	CommonResponseWorkers
}

// accountWorkerDispatchNamespaceScriptSecretListResponseJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptSecretListResponse]
type accountWorkerDispatchNamespaceScriptSecretListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSecretListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptSecretListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDispatchNamespaceScriptSecretAddResponse struct {
	Result SecretResponse                                            `json:"result"`
	JSON   accountWorkerDispatchNamespaceScriptSecretAddResponseJSON `json:"-"`
	CommonResponseWorkers
}

// accountWorkerDispatchNamespaceScriptSecretAddResponseJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptSecretAddResponse]
type accountWorkerDispatchNamespaceScriptSecretAddResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSecretAddResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptSecretAddResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDispatchNamespaceScriptSecretGetResponse struct {
	Result SecretResponse                                            `json:"result"`
	JSON   accountWorkerDispatchNamespaceScriptSecretGetResponseJSON `json:"-"`
	CommonResponseWorkers
}

// accountWorkerDispatchNamespaceScriptSecretGetResponseJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptSecretGetResponse]
type accountWorkerDispatchNamespaceScriptSecretGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSecretGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptSecretGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDispatchNamespaceScriptSecretAddParams struct {
	Secret SecretParam `json:"secret,required"`
}

func (r AccountWorkerDispatchNamespaceScriptSecretAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Secret)
}
