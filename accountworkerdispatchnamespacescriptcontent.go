// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountWorkerDispatchNamespaceScriptContentService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerDispatchNamespaceScriptContentService] method instead.
type AccountWorkerDispatchNamespaceScriptContentService struct {
	Options []option.RequestOption
}

// NewAccountWorkerDispatchNamespaceScriptContentService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewAccountWorkerDispatchNamespaceScriptContentService(opts ...option.RequestOption) (r *AccountWorkerDispatchNamespaceScriptContentService) {
	r = &AccountWorkerDispatchNamespaceScriptContentService{}
	r.Options = opts
	return
}

// Fetch script content from a script uploaded to a Workers for Platforms
// namespace.
func (r *AccountWorkerDispatchNamespaceScriptContentService) Get(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "string")}, opts...)
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/content", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Put script content for a script uploaded to a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptContentService) Put(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, params AccountWorkerDispatchNamespaceScriptContentPutParams, opts ...option.RequestOption) (res *SingleScriptResponse, err error) {
	if params.CfWorkerBodyPart.Present {
		opts = append(opts, option.WithHeader("CF-WORKER-BODY-PART", fmt.Sprintf("%s", params.CfWorkerBodyPart)))
	}
	if params.CfWorkerMainModulePart.Present {
		opts = append(opts, option.WithHeader("CF-WORKER-MAIN-MODULE-PART", fmt.Sprintf("%s", params.CfWorkerMainModulePart)))
	}
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/content", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type PlacementStatus string

const (
	PlacementStatusSuccess                 PlacementStatus = "SUCCESS"
	PlacementStatusUnsupportedApplication  PlacementStatus = "UNSUPPORTED_APPLICATION"
	PlacementStatusInsufficientInvocations PlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r PlacementStatus) IsKnown() bool {
	switch r {
	case PlacementStatusSuccess, PlacementStatusUnsupportedApplication, PlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

type ScriptResponse struct {
	// The id of the script in the Workers system. Usually the script name.
	ID string `json:"id"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Hashed script content, can be used in a If-None-Match header when updating.
	Etag string `json:"etag"`
	// Whether a Worker contains assets.
	HasAssets bool `json:"has_assets"`
	// Whether a Worker contains modules.
	HasModules bool `json:"has_modules"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement ScriptResponsePlacement `json:"placement"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	//
	// Deprecated: deprecated
	PlacementMode PlacementMode `json:"placement_mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	//
	// Deprecated: deprecated
	PlacementStatus PlacementStatus `json:"placement_status"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []TailConsumersScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel UsageModel         `json:"usage_model"`
	JSON       scriptResponseJSON `json:"-"`
}

// scriptResponseJSON contains the JSON metadata for the struct [ScriptResponse]
type scriptResponseJSON struct {
	ID              apijson.Field
	CreatedOn       apijson.Field
	Etag            apijson.Field
	HasAssets       apijson.Field
	HasModules      apijson.Field
	Logpush         apijson.Field
	ModifiedOn      apijson.Field
	Placement       apijson.Field
	PlacementMode   apijson.Field
	PlacementStatus apijson.Field
	TailConsumers   apijson.Field
	UsageModel      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScriptResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptResponseJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode PlacementMode `json:"mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Status PlacementStatus             `json:"status"`
	JSON   scriptResponsePlacementJSON `json:"-"`
}

// scriptResponsePlacementJSON contains the JSON metadata for the struct
// [ScriptResponsePlacement]
type scriptResponsePlacementJSON struct {
	Mode        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptResponsePlacementJSON) RawJSON() string {
	return r.raw
}

type SingleScriptResponse struct {
	Result ScriptResponse           `json:"result"`
	JSON   singleScriptResponseJSON `json:"-"`
	APIResponseSingleWorkers
}

// singleScriptResponseJSON contains the JSON metadata for the struct
// [SingleScriptResponse]
type singleScriptResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleScriptResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleScriptResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDispatchNamespaceScriptContentPutParams struct {
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata               param.Field[AccountWorkerDispatchNamespaceScriptContentPutParamsMetadata] `json:"metadata,required"`
	CfWorkerBodyPart       param.Field[string]                                                       `header:"CF-WORKER-BODY-PART"`
	CfWorkerMainModulePart param.Field[string]                                                       `header:"CF-WORKER-MAIN-MODULE-PART"`
}

func (r AccountWorkerDispatchNamespaceScriptContentPutParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// JSON encoded metadata about the uploaded parts and Worker configuration.
type AccountWorkerDispatchNamespaceScriptContentPutParamsMetadata struct {
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
}

func (r AccountWorkerDispatchNamespaceScriptContentPutParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
