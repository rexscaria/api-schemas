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
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountWorkerScriptSecretListResponseSuccess `json:"success,required"`
	Result  []Secret                                     `json:"result"`
	JSON    accountWorkerScriptSecretListResponseJSON    `json:"-"`
}

// accountWorkerScriptSecretListResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptSecretListResponse]
type accountWorkerScriptSecretListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful.
type AccountWorkerScriptSecretListResponseSuccess bool

const (
	AccountWorkerScriptSecretListResponseSuccessTrue AccountWorkerScriptSecretListResponseSuccess = true
)

func (r AccountWorkerScriptSecretListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerScriptSecretListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptSecretAddResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountWorkerScriptSecretAddResponseSuccess `json:"success,required"`
	// A secret value accessible through a binding.
	Result Secret                                   `json:"result"`
	JSON   accountWorkerScriptSecretAddResponseJSON `json:"-"`
}

// accountWorkerScriptSecretAddResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptSecretAddResponse]
type accountWorkerScriptSecretAddResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful.
type AccountWorkerScriptSecretAddResponseSuccess bool

const (
	AccountWorkerScriptSecretAddResponseSuccessTrue AccountWorkerScriptSecretAddResponseSuccess = true
)

func (r AccountWorkerScriptSecretAddResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerScriptSecretAddResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptSecretGetResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountWorkerScriptSecretGetResponseSuccess `json:"success,required"`
	// A secret value accessible through a binding.
	Result Secret                                   `json:"result"`
	JSON   accountWorkerScriptSecretGetResponseJSON `json:"-"`
}

// accountWorkerScriptSecretGetResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptSecretGetResponse]
type accountWorkerScriptSecretGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful.
type AccountWorkerScriptSecretGetResponseSuccess bool

const (
	AccountWorkerScriptSecretGetResponseSuccessTrue AccountWorkerScriptSecretGetResponseSuccess = true
)

func (r AccountWorkerScriptSecretGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerScriptSecretGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptSecretAddParams struct {
	// A secret value accessible through a binding.
	Secret SecretUnionParam `json:"secret,required"`
}

func (r AccountWorkerScriptSecretAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Secret)
}
