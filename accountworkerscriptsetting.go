// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountWorkerScriptSettingService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerScriptSettingService] method instead.
type AccountWorkerScriptSettingService struct {
	Options []option.RequestOption
}

// NewAccountWorkerScriptSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptSettingService(opts ...option.RequestOption) (r *AccountWorkerScriptSettingService) {
	r = &AccountWorkerScriptSettingService{}
	r.Options = opts
	return
}

// Get metadata and config, such as bindings or usage model.
func (r *AccountWorkerScriptSettingService) Get(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *ScriptVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch metadata or config, such as bindings or usage model.
func (r *AccountWorkerScriptSettingService) Patch(ctx context.Context, accountID string, scriptName string, body AccountWorkerScriptSettingPatchParams, opts ...option.RequestOption) (res *ScriptVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ScriptVersionResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	Result   ScriptVersionItem `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptVersionResponseSuccess `json:"success,required"`
	JSON    scriptVersionResponseJSON    `json:"-"`
}

// scriptVersionResponseJSON contains the JSON metadata for the struct
// [ScriptVersionResponse]
type scriptVersionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptVersionResponseSuccess bool

const (
	ScriptVersionResponseSuccessTrue ScriptVersionResponseSuccess = true
)

func (r ScriptVersionResponseSuccess) IsKnown() bool {
	switch r {
	case ScriptVersionResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptSettingPatchParams struct {
	Settings param.Field[ScriptVersionItemParam] `json:"settings"`
}

func (r AccountWorkerScriptSettingPatchParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
