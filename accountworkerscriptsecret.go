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

// AccountWorkerScriptSecretService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerScriptSecretService] method instead.
type AccountWorkerScriptSecretService struct {
	Options []option.RequestOption
}

// NewAccountWorkerScriptSecretService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptSecretService(opts ...option.RequestOption) (r *AccountWorkerScriptSecretService) {
	r = &AccountWorkerScriptSecretService{}
	r.Options = opts
	return
}

// List secrets bound to a script.
func (r *AccountWorkerScriptSecretService) List(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerScriptSecretListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/secrets", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove a secret from a script.
func (r *AccountWorkerScriptSecretService) Delete(ctx context.Context, accountID string, scriptName string, secretName string, opts ...option.RequestOption) (res *NullResult, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
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
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/secrets/%s", accountID, scriptName, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add a secret to a script.
func (r *AccountWorkerScriptSecretService) Add(ctx context.Context, accountID string, scriptName string, body AccountWorkerScriptSecretAddParams, opts ...option.RequestOption) (res *AccountWorkerScriptSecretAddResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/secrets", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get a given secret binding (value omitted) on a script.
func (r *AccountWorkerScriptSecretService) Get(ctx context.Context, accountID string, scriptName string, secretName string, opts ...option.RequestOption) (res *AccountWorkerScriptSecretGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
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
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/secrets/%s", accountID, scriptName, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountWorkerScriptSecretListResponse struct {
	Result []SecretResponse                          `json:"result"`
	JSON   accountWorkerScriptSecretListResponseJSON `json:"-"`
	CommonResponseWorkers
}

// accountWorkerScriptSecretListResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptSecretListResponse]
type accountWorkerScriptSecretListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSecretListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptSecretListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptSecretAddResponse struct {
	Result SecretResponse                           `json:"result"`
	JSON   accountWorkerScriptSecretAddResponseJSON `json:"-"`
	CommonResponseWorkers
}

// accountWorkerScriptSecretAddResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptSecretAddResponse]
type accountWorkerScriptSecretAddResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSecretAddResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptSecretAddResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptSecretGetResponse struct {
	Result SecretResponse                           `json:"result"`
	JSON   accountWorkerScriptSecretGetResponseJSON `json:"-"`
	CommonResponseWorkers
}

// accountWorkerScriptSecretGetResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptSecretGetResponse]
type accountWorkerScriptSecretGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSecretGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptSecretGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptSecretAddParams struct {
	Secret SecretParam `json:"secret,required"`
}

func (r AccountWorkerScriptSecretAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Secret)
}
