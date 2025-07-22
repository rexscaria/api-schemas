// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountWorkerScriptService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerScriptService] method instead.
type AccountWorkerScriptService struct {
	Options        []option.RequestOption
	Content        *AccountWorkerScriptContentService
	Deployments    *AccountWorkerScriptDeploymentService
	Schedules      *AccountWorkerScriptScheduleService
	ScriptSettings *AccountWorkerScriptScriptSettingService
	Secrets        *AccountWorkerScriptSecretService
	Settings       *AccountWorkerScriptSettingService
	Subdomain      *AccountWorkerScriptSubdomainService
	Tails          *AccountWorkerScriptTailService
	UsageModel     *AccountWorkerScriptUsageModelService
	Versions       *AccountWorkerScriptVersionService
}

// NewAccountWorkerScriptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptService(opts ...option.RequestOption) (r *AccountWorkerScriptService) {
	r = &AccountWorkerScriptService{}
	r.Options = opts
	r.Content = NewAccountWorkerScriptContentService(opts...)
	r.Deployments = NewAccountWorkerScriptDeploymentService(opts...)
	r.Schedules = NewAccountWorkerScriptScheduleService(opts...)
	r.ScriptSettings = NewAccountWorkerScriptScriptSettingService(opts...)
	r.Secrets = NewAccountWorkerScriptSecretService(opts...)
	r.Settings = NewAccountWorkerScriptSettingService(opts...)
	r.Subdomain = NewAccountWorkerScriptSubdomainService(opts...)
	r.Tails = NewAccountWorkerScriptTailService(opts...)
	r.UsageModel = NewAccountWorkerScriptUsageModelService(opts...)
	r.Versions = NewAccountWorkerScriptVersionService(opts...)
	return
}

// Fetch a list of uploaded workers.
func (r *AccountWorkerScriptService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountWorkerScriptListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete your worker. This call has no response body on a successful delete.
func (r *AccountWorkerScriptService) Delete(ctx context.Context, accountID string, scriptName string, params AccountWorkerScriptDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Start uploading a collection of assets for use in a Worker version. To learn
// more about the direct uploads of assets, see
// https://developers.cloudflare.com/workers/static-assets/direct-upload/
func (r *AccountWorkerScriptService) NewAssetsUploadSession(ctx context.Context, accountID string, scriptName string, body AccountWorkerScriptNewAssetsUploadSessionParams, opts ...option.RequestOption) (res *UploadSessionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/assets-upload-session", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch raw script content for your worker. Note this is the original script
// content, not JSON encoded.
func (r *AccountWorkerScriptService) Download(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/javascript")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a worker module. You can find more about the multipart metadata on our
// docs:
// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.
func (r *AccountWorkerScriptService) Upload(ctx context.Context, accountID string, scriptName string, body AccountWorkerScriptUploadParams, opts ...option.RequestOption) (res *AccountWorkerScriptUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountWorkerScriptListResponse struct {
	Result []ScriptResponse                    `json:"result"`
	JSON   accountWorkerScriptListResponseJSON `json:"-"`
	CommonResponseWorkers
}

// accountWorkerScriptListResponseJSON contains the JSON metadata for the struct
// [AccountWorkerScriptListResponse]
type accountWorkerScriptListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptUploadResponse struct {
	JSON accountWorkerScriptUploadResponseJSON `json:"-"`
	APIResponseSingleWorkers
}

// accountWorkerScriptUploadResponseJSON contains the JSON metadata for the struct
// [AccountWorkerScriptUploadResponse]
type accountWorkerScriptUploadResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptUploadResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptDeleteParams struct {
	Body interface{} `json:"body,required"`
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
}

func (r AccountWorkerScriptDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountWorkerScriptDeleteParams]'s query parameters as
// `url.Values`.
func (r AccountWorkerScriptDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountWorkerScriptNewAssetsUploadSessionParams struct {
	UploadSessionObject UploadSessionObjectParam `json:"upload_session_object,required"`
}

func (r AccountWorkerScriptNewAssetsUploadSessionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UploadSessionObject)
}

type AccountWorkerScriptUploadParams struct {
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[AccountWorkerScriptUploadParamsMetadata] `json:"metadata,required"`
}

func (r AccountWorkerScriptUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type AccountWorkerScriptUploadParamsMetadata struct {
	// Configuration for assets within a Worker
	Assets param.Field[AccountWorkerScriptUploadParamsMetadataAssets] `json:"assets"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]BindingItemUnionParam] `json:"bindings"`
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Retain assets which exist for a previously uploaded Worker version; used in lieu
	// of providing a completion token.
	KeepAssets param.Field[bool] `json:"keep_assets"`
	// List of binding types to keep from previous_upload.
	KeepBindings param.Field[[]string] `json:"keep_bindings"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[AccountWorkerScriptUploadParamsMetadataMigrationsUnion] `json:"migrations"`
	// Observability settings for the Worker.
	Observability param.Field[ObservabilityParam] `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement param.Field[AccountWorkerScriptUploadParamsMetadataPlacement] `json:"placement"`
	// List of strings to use as tags for this Worker.
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]TailConsumersScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[UsageModel] `json:"usage_model"`
}

func (r AccountWorkerScriptUploadParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker
type AccountWorkerScriptUploadParamsMetadataAssets struct {
	// Configuration for assets within a Worker.
	Config param.Field[AccountWorkerScriptUploadParamsMetadataAssetsConfig] `json:"config"`
	// Token provided upon successful upload of all files from a registered manifest.
	Jwt param.Field[string] `json:"jwt"`
}

func (r AccountWorkerScriptUploadParamsMetadataAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type AccountWorkerScriptUploadParamsMetadataAssetsConfig struct {
	// The contents of a \_headers file (used to attach custom headers on asset
	// responses)
	Headers param.Field[string] `json:"_headers"`
	// The contents of a \_redirects file (used to apply redirects or proxy paths ahead
	// of asset serving)
	Redirects param.Field[string] `json:"_redirects"`
	// Determines the redirects and rewrites of requests for HTML content.
	HTMLHandling param.Field[AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandling] `json:"html_handling"`
	// Determines the response when a request does not match a static asset, and there
	// is no Worker script.
	NotFoundHandling param.Field[AccountWorkerScriptUploadParamsMetadataAssetsConfigNotFoundHandling] `json:"not_found_handling"`
	// When true, requests will always invoke the Worker script. Otherwise, attempt to
	// serve an asset matching the request, falling back to the Worker script.
	RunWorkerFirst param.Field[bool] `json:"run_worker_first"`
	// When true and the incoming request matches an asset, that will be served instead
	// of invoking the Worker script. When false, requests will always invoke the
	// Worker script.
	//
	// Deprecated: deprecated
	ServeDirectly param.Field[bool] `json:"serve_directly"`
}

func (r AccountWorkerScriptUploadParamsMetadataAssetsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the redirects and rewrites of requests for HTML content.
type AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandling string

const (
	AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash  AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandling = "auto-trailing-slash"
	AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandlingForceTrailingSlash AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandling = "force-trailing-slash"
	AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandlingDropTrailingSlash  AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandling = "drop-trailing-slash"
	AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandlingNone               AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandling = "none"
)

func (r AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandling) IsKnown() bool {
	switch r {
	case AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash, AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandlingForceTrailingSlash, AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandlingDropTrailingSlash, AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandlingNone:
		return true
	}
	return false
}

// Determines the response when a request does not match a static asset, and there
// is no Worker script.
type AccountWorkerScriptUploadParamsMetadataAssetsConfigNotFoundHandling string

const (
	AccountWorkerScriptUploadParamsMetadataAssetsConfigNotFoundHandlingNone                  AccountWorkerScriptUploadParamsMetadataAssetsConfigNotFoundHandling = "none"
	AccountWorkerScriptUploadParamsMetadataAssetsConfigNotFoundHandling404Page               AccountWorkerScriptUploadParamsMetadataAssetsConfigNotFoundHandling = "404-page"
	AccountWorkerScriptUploadParamsMetadataAssetsConfigNotFoundHandlingSinglePageApplication AccountWorkerScriptUploadParamsMetadataAssetsConfigNotFoundHandling = "single-page-application"
)

func (r AccountWorkerScriptUploadParamsMetadataAssetsConfigNotFoundHandling) IsKnown() bool {
	switch r {
	case AccountWorkerScriptUploadParamsMetadataAssetsConfigNotFoundHandlingNone, AccountWorkerScriptUploadParamsMetadataAssetsConfigNotFoundHandling404Page, AccountWorkerScriptUploadParamsMetadataAssetsConfigNotFoundHandlingSinglePageApplication:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations],
// [AccountWorkerScriptUploadParamsMetadataMigrationsWorkersMultipleStepMigrations].
type AccountWorkerScriptUploadParamsMetadataMigrationsUnion interface {
	implementsAccountWorkerScriptUploadParamsMetadataMigrationsUnion()
}

// A single set of migrations to apply.
type AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations struct {
	MigrationTagConditionsParam
	MigrationStepParam
}

func (r AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations) implementsAccountWorkerScriptUploadParamsMetadataMigrationsUnion() {
}

type AccountWorkerScriptUploadParamsMetadataMigrationsWorkersMultipleStepMigrations struct {
	// Migrations to apply in order.
	Steps param.Field[[]MigrationStepParam] `json:"steps"`
	MigrationTagConditionsParam
}

func (r AccountWorkerScriptUploadParamsMetadataMigrationsWorkersMultipleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptUploadParamsMetadataMigrationsWorkersMultipleStepMigrations) implementsAccountWorkerScriptUploadParamsMetadataMigrationsUnion() {
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type AccountWorkerScriptUploadParamsMetadataPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[PlacementMode] `json:"mode"`
}

func (r AccountWorkerScriptUploadParamsMetadataPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
