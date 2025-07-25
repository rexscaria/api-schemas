// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apiform"
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
func (r *AccountWorkerScriptService) Delete(ctx context.Context, accountID string, scriptName string, body AccountWorkerScriptDeleteParams, opts ...option.RequestOption) (res *NullResult, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Start uploading a collection of assets for use in a Worker version. To learn
// more about the direct uploads of assets, see
// https://developers.cloudflare.com/workers/static-assets/direct-upload/.
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
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	Result   []ScriptResponse  `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerScriptListResponseSuccess `json:"success,required"`
	JSON    accountWorkerScriptListResponseJSON    `json:"-"`
}

// accountWorkerScriptListResponseJSON contains the JSON metadata for the struct
// [AccountWorkerScriptListResponse]
type accountWorkerScriptListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerScriptListResponseSuccess bool

const (
	AccountWorkerScriptListResponseSuccessTrue AccountWorkerScriptListResponseSuccess = true
)

func (r AccountWorkerScriptListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerScriptListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptUploadResponse struct {
	Errors   []WorkersMessages                       `json:"errors,required"`
	Messages []WorkersMessages                       `json:"messages,required"`
	Result   AccountWorkerScriptUploadResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerScriptUploadResponseSuccess `json:"success,required"`
	JSON    accountWorkerScriptUploadResponseJSON    `json:"-"`
}

// accountWorkerScriptUploadResponseJSON contains the JSON metadata for the struct
// [AccountWorkerScriptUploadResponse]
type accountWorkerScriptUploadResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptUploadResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptUploadResponseResult struct {
	StartupTimeMs int64 `json:"startup_time_ms,required"`
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
	Placement AccountWorkerScriptUploadResponseResultPlacement `json:"placement"`
	// Deprecated: deprecated
	PlacementMode AccountWorkerScriptUploadResponseResultPlacementMode `json:"placement_mode"`
	// Deprecated: deprecated
	PlacementStatus AccountWorkerScriptUploadResponseResultPlacementStatus `json:"placement_status"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []TailConsumersScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel UsageModel                                  `json:"usage_model"`
	JSON       accountWorkerScriptUploadResponseResultJSON `json:"-"`
}

// accountWorkerScriptUploadResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkerScriptUploadResponseResult]
type accountWorkerScriptUploadResponseResultJSON struct {
	StartupTimeMs   apijson.Field
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

func (r *AccountWorkerScriptUploadResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptUploadResponseResultJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type AccountWorkerScriptUploadResponseResultPlacement struct {
	// The last time the script was analyzed for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	LastAnalyzedAt time.Time `json:"last_analyzed_at" format:"date-time"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode PlacementMode `json:"mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Status PlacementStatus                                      `json:"status"`
	JSON   accountWorkerScriptUploadResponseResultPlacementJSON `json:"-"`
}

// accountWorkerScriptUploadResponseResultPlacementJSON contains the JSON metadata
// for the struct [AccountWorkerScriptUploadResponseResultPlacement]
type accountWorkerScriptUploadResponseResultPlacementJSON struct {
	LastAnalyzedAt apijson.Field
	Mode           apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountWorkerScriptUploadResponseResultPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptUploadResponseResultPlacementJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptUploadResponseResultPlacementMode string

const (
	AccountWorkerScriptUploadResponseResultPlacementModeSmart AccountWorkerScriptUploadResponseResultPlacementMode = "smart"
)

func (r AccountWorkerScriptUploadResponseResultPlacementMode) IsKnown() bool {
	switch r {
	case AccountWorkerScriptUploadResponseResultPlacementModeSmart:
		return true
	}
	return false
}

type AccountWorkerScriptUploadResponseResultPlacementStatus string

const (
	AccountWorkerScriptUploadResponseResultPlacementStatusSuccess                 AccountWorkerScriptUploadResponseResultPlacementStatus = "SUCCESS"
	AccountWorkerScriptUploadResponseResultPlacementStatusUnsupportedApplication  AccountWorkerScriptUploadResponseResultPlacementStatus = "UNSUPPORTED_APPLICATION"
	AccountWorkerScriptUploadResponseResultPlacementStatusInsufficientInvocations AccountWorkerScriptUploadResponseResultPlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r AccountWorkerScriptUploadResponseResultPlacementStatus) IsKnown() bool {
	switch r {
	case AccountWorkerScriptUploadResponseResultPlacementStatusSuccess, AccountWorkerScriptUploadResponseResultPlacementStatusUnsupportedApplication, AccountWorkerScriptUploadResponseResultPlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

// Whether the API call was successful.
type AccountWorkerScriptUploadResponseSuccess bool

const (
	AccountWorkerScriptUploadResponseSuccessTrue AccountWorkerScriptUploadResponseSuccess = true
)

func (r AccountWorkerScriptUploadResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerScriptUploadResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptDeleteParams struct {
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
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
	// An array of modules (often JavaScript files) comprising a Worker script. At
	// least one module must be present and referenced in the metadata as `main_module`
	// or `body_part` by filename.<br/>Possible Content-Type(s) are:
	// `application/javascript+module`, `text/javascript+module`,
	// `application/javascript`, `text/javascript`, `application/wasm`, `text/plain`,
	// `application/octet-stream`, `application/source-map`.
	Files param.Field[[]io.Reader] `json:"files" format:"binary"`
}

func (r AccountWorkerScriptUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type AccountWorkerScriptUploadParamsMetadata struct {
	// Configuration for assets within a Worker.
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

// Configuration for assets within a Worker.
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
	// responses).
	Headers param.Field[string] `json:"_headers"`
	// The contents of a \_redirects file (used to apply redirects or proxy paths ahead
	// of asset serving).
	Redirects param.Field[string] `json:"_redirects"`
	// Determines the redirects and rewrites of requests for HTML content.
	HTMLHandling param.Field[AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandling] `json:"html_handling"`
	// Determines the response when a request does not match a static asset, and there
	// is no Worker script.
	NotFoundHandling param.Field[AccountWorkerScriptUploadParamsMetadataAssetsConfigNotFoundHandling] `json:"not_found_handling"`
	// Contains a list path rules to control routing to either the Worker or assets.
	// Glob (\*) and negative (!) rules are supported. Rules must start with either '/'
	// or '!/'. At least one non-negative rule must be provided, and negative rules
	// have higher precedence than non-negative rules.
	RunWorkerFirst param.Field[AccountWorkerScriptUploadParamsMetadataAssetsConfigRunWorkerFirstUnion] `json:"run_worker_first"`
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

// Contains a list path rules to control routing to either the Worker or assets.
// Glob (\*) and negative (!) rules are supported. Rules must start with either '/'
// or '!/'. At least one non-negative rule must be provided, and negative rules
// have higher precedence than non-negative rules.
//
// Satisfied by
// [AccountWorkerScriptUploadParamsMetadataAssetsConfigRunWorkerFirstArray],
// [shared.UnionBool].
type AccountWorkerScriptUploadParamsMetadataAssetsConfigRunWorkerFirstUnion interface {
	ImplementsAccountWorkerScriptUploadParamsMetadataAssetsConfigRunWorkerFirstUnion()
}

type AccountWorkerScriptUploadParamsMetadataAssetsConfigRunWorkerFirstArray []string

func (r AccountWorkerScriptUploadParamsMetadataAssetsConfigRunWorkerFirstArray) ImplementsAccountWorkerScriptUploadParamsMetadataAssetsConfigRunWorkerFirstUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
type AccountWorkerScriptUploadParamsMetadataMigrations struct {
	DeletedClasses   param.Field[interface{}] `json:"deleted_classes"`
	NewClasses       param.Field[interface{}] `json:"new_classes"`
	NewSqliteClasses param.Field[interface{}] `json:"new_sqlite_classes"`
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag             param.Field[string]      `json:"old_tag"`
	RenamedClasses     param.Field[interface{}] `json:"renamed_classes"`
	Steps              param.Field[interface{}] `json:"steps"`
	TransferredClasses param.Field[interface{}] `json:"transferred_classes"`
}

func (r AccountWorkerScriptUploadParamsMetadataMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptUploadParamsMetadataMigrations) implementsAccountWorkerScriptUploadParamsMetadataMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations],
// [AccountWorkerScriptUploadParamsMetadataMigrationsWorkersMultipleStepMigrations],
// [AccountWorkerScriptUploadParamsMetadataMigrations].
type AccountWorkerScriptUploadParamsMetadataMigrationsUnion interface {
	implementsAccountWorkerScriptUploadParamsMetadataMigrationsUnion()
}

// A single set of migrations to apply.
type AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes to create Durable Object namespaces with SQLite from.
	NewSqliteClasses param.Field[[]string] `json:"new_sqlite_classes"`
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations) implementsAccountWorkerScriptUploadParamsMetadataMigrationsUnion() {
}

type AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptUploadParamsMetadataMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]MigrationStepParam] `json:"steps"`
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
