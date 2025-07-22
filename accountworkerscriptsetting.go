// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apiform"
	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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

// Get metadata and config, such as bindings or usage model
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

// Patch metadata or config, such as bindings or usage model
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
	Result ScriptVersionItem         `json:"result"`
	JSON   scriptVersionResponseJSON `json:"-"`
	CommonResponseWorkers
}

// scriptVersionResponseJSON contains the JSON metadata for the struct
// [ScriptVersionResponse]
type scriptVersionResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionResponseJSON) RawJSON() string {
	return r.raw
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
