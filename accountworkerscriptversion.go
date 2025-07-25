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

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountWorkerScriptVersionService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerScriptVersionService] method instead.
type AccountWorkerScriptVersionService struct {
	Options []option.RequestOption
}

// NewAccountWorkerScriptVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptVersionService(opts ...option.RequestOption) (r *AccountWorkerScriptVersionService) {
	r = &AccountWorkerScriptVersionService{}
	r.Options = opts
	return
}

// List of Worker Versions. The first version in the list is the latest version.
func (r *AccountWorkerScriptVersionService) List(ctx context.Context, accountID string, scriptName string, query AccountWorkerScriptVersionListParams, opts ...option.RequestOption) (res *AccountWorkerScriptVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Version Detail
func (r *AccountWorkerScriptVersionService) GetDetail(ctx context.Context, accountID string, scriptName string, versionID string, opts ...option.RequestOption) (res *AccountWorkerScriptVersionGetDetailResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions/%s", accountID, scriptName, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a Worker Version without deploying to Cloudflare's network. You can find
// more about the multipart metadata on our docs:
// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.
func (r *AccountWorkerScriptVersionService) Upload(ctx context.Context, accountID string, scriptName string, body AccountWorkerScriptVersionUploadParams, opts ...option.RequestOption) (res *AccountWorkerScriptVersionUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type VersionItemFull struct {
	Resources VersionItemFullResources `json:"resources,required"`
	ID        string                   `json:"id"`
	Metadata  VersionItemFullMetadata  `json:"metadata"`
	Number    float64                  `json:"number"`
	JSON      versionItemFullJSON      `json:"-"`
}

// versionItemFullJSON contains the JSON metadata for the struct [VersionItemFull]
type versionItemFullJSON struct {
	Resources   apijson.Field
	ID          apijson.Field
	Metadata    apijson.Field
	Number      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionItemFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionItemFullJSON) RawJSON() string {
	return r.raw
}

type VersionItemFullResources struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings      []BindingItem                         `json:"bindings"`
	Script        VersionItemFullResourcesScript        `json:"script"`
	ScriptRuntime VersionItemFullResourcesScriptRuntime `json:"script_runtime"`
	JSON          versionItemFullResourcesJSON          `json:"-"`
}

// versionItemFullResourcesJSON contains the JSON metadata for the struct
// [VersionItemFullResources]
type versionItemFullResourcesJSON struct {
	Bindings      apijson.Field
	Script        apijson.Field
	ScriptRuntime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VersionItemFullResources) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionItemFullResourcesJSON) RawJSON() string {
	return r.raw
}

type VersionItemFullResourcesScript struct {
	Etag             string                                       `json:"etag"`
	Handlers         []string                                     `json:"handlers"`
	LastDeployedFrom string                                       `json:"last_deployed_from"`
	NamedHandlers    []VersionItemFullResourcesScriptNamedHandler `json:"named_handlers"`
	JSON             versionItemFullResourcesScriptJSON           `json:"-"`
}

// versionItemFullResourcesScriptJSON contains the JSON metadata for the struct
// [VersionItemFullResourcesScript]
type versionItemFullResourcesScriptJSON struct {
	Etag             apijson.Field
	Handlers         apijson.Field
	LastDeployedFrom apijson.Field
	NamedHandlers    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VersionItemFullResourcesScript) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionItemFullResourcesScriptJSON) RawJSON() string {
	return r.raw
}

type VersionItemFullResourcesScriptNamedHandler struct {
	Handlers []string                                       `json:"handlers"`
	Name     string                                         `json:"name"`
	JSON     versionItemFullResourcesScriptNamedHandlerJSON `json:"-"`
}

// versionItemFullResourcesScriptNamedHandlerJSON contains the JSON metadata for
// the struct [VersionItemFullResourcesScriptNamedHandler]
type versionItemFullResourcesScriptNamedHandlerJSON struct {
	Handlers    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionItemFullResourcesScriptNamedHandler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionItemFullResourcesScriptNamedHandlerJSON) RawJSON() string {
	return r.raw
}

type VersionItemFullResourcesScriptRuntime struct {
	CompatibilityDate  string                                          `json:"compatibility_date"`
	CompatibilityFlags []string                                        `json:"compatibility_flags"`
	Limits             VersionItemFullResourcesScriptRuntimeLimits     `json:"limits"`
	MigrationTag       string                                          `json:"migration_tag"`
	UsageModel         VersionItemFullResourcesScriptRuntimeUsageModel `json:"usage_model"`
	JSON               versionItemFullResourcesScriptRuntimeJSON       `json:"-"`
}

// versionItemFullResourcesScriptRuntimeJSON contains the JSON metadata for the
// struct [VersionItemFullResourcesScriptRuntime]
type versionItemFullResourcesScriptRuntimeJSON struct {
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Limits             apijson.Field
	MigrationTag       apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VersionItemFullResourcesScriptRuntime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionItemFullResourcesScriptRuntimeJSON) RawJSON() string {
	return r.raw
}

type VersionItemFullResourcesScriptRuntimeLimits struct {
	CPUMs int64                                           `json:"cpu_ms"`
	JSON  versionItemFullResourcesScriptRuntimeLimitsJSON `json:"-"`
}

// versionItemFullResourcesScriptRuntimeLimitsJSON contains the JSON metadata for
// the struct [VersionItemFullResourcesScriptRuntimeLimits]
type versionItemFullResourcesScriptRuntimeLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionItemFullResourcesScriptRuntimeLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionItemFullResourcesScriptRuntimeLimitsJSON) RawJSON() string {
	return r.raw
}

type VersionItemFullResourcesScriptRuntimeUsageModel string

const (
	VersionItemFullResourcesScriptRuntimeUsageModelBundled  VersionItemFullResourcesScriptRuntimeUsageModel = "bundled"
	VersionItemFullResourcesScriptRuntimeUsageModelUnbound  VersionItemFullResourcesScriptRuntimeUsageModel = "unbound"
	VersionItemFullResourcesScriptRuntimeUsageModelStandard VersionItemFullResourcesScriptRuntimeUsageModel = "standard"
)

func (r VersionItemFullResourcesScriptRuntimeUsageModel) IsKnown() bool {
	switch r {
	case VersionItemFullResourcesScriptRuntimeUsageModelBundled, VersionItemFullResourcesScriptRuntimeUsageModelUnbound, VersionItemFullResourcesScriptRuntimeUsageModelStandard:
		return true
	}
	return false
}

type VersionItemFullMetadata struct {
	AuthorEmail string                        `json:"author_email"`
	AuthorID    string                        `json:"author_id"`
	CreatedOn   string                        `json:"created_on"`
	HasPreview  bool                          `json:"hasPreview"`
	ModifiedOn  string                        `json:"modified_on"`
	Source      VersionItemFullMetadataSource `json:"source"`
	JSON        versionItemFullMetadataJSON   `json:"-"`
}

// versionItemFullMetadataJSON contains the JSON metadata for the struct
// [VersionItemFullMetadata]
type versionItemFullMetadataJSON struct {
	AuthorEmail apijson.Field
	AuthorID    apijson.Field
	CreatedOn   apijson.Field
	HasPreview  apijson.Field
	ModifiedOn  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionItemFullMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionItemFullMetadataJSON) RawJSON() string {
	return r.raw
}

type VersionItemFullMetadataSource string

const (
	VersionItemFullMetadataSourceUnknown      VersionItemFullMetadataSource = "unknown"
	VersionItemFullMetadataSourceAPI          VersionItemFullMetadataSource = "api"
	VersionItemFullMetadataSourceWrangler     VersionItemFullMetadataSource = "wrangler"
	VersionItemFullMetadataSourceTerraform    VersionItemFullMetadataSource = "terraform"
	VersionItemFullMetadataSourceDash         VersionItemFullMetadataSource = "dash"
	VersionItemFullMetadataSourceDashTemplate VersionItemFullMetadataSource = "dash_template"
	VersionItemFullMetadataSourceIntegration  VersionItemFullMetadataSource = "integration"
	VersionItemFullMetadataSourceQuickEditor  VersionItemFullMetadataSource = "quick_editor"
	VersionItemFullMetadataSourcePlayground   VersionItemFullMetadataSource = "playground"
	VersionItemFullMetadataSourceWorkersci    VersionItemFullMetadataSource = "workersci"
)

func (r VersionItemFullMetadataSource) IsKnown() bool {
	switch r {
	case VersionItemFullMetadataSourceUnknown, VersionItemFullMetadataSourceAPI, VersionItemFullMetadataSourceWrangler, VersionItemFullMetadataSourceTerraform, VersionItemFullMetadataSourceDash, VersionItemFullMetadataSourceDashTemplate, VersionItemFullMetadataSourceIntegration, VersionItemFullMetadataSourceQuickEditor, VersionItemFullMetadataSourcePlayground, VersionItemFullMetadataSourceWorkersci:
		return true
	}
	return false
}

type VersionItemShort struct {
	ID       string                   `json:"id"`
	Metadata VersionItemShortMetadata `json:"metadata"`
	Number   float64                  `json:"number"`
	JSON     versionItemShortJSON     `json:"-"`
}

// versionItemShortJSON contains the JSON metadata for the struct
// [VersionItemShort]
type versionItemShortJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Number      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionItemShort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionItemShortJSON) RawJSON() string {
	return r.raw
}

type VersionItemShortMetadata struct {
	AuthorEmail string                         `json:"author_email"`
	AuthorID    string                         `json:"author_id"`
	CreatedOn   string                         `json:"created_on"`
	HasPreview  bool                           `json:"hasPreview"`
	ModifiedOn  string                         `json:"modified_on"`
	Source      VersionItemShortMetadataSource `json:"source"`
	JSON        versionItemShortMetadataJSON   `json:"-"`
}

// versionItemShortMetadataJSON contains the JSON metadata for the struct
// [VersionItemShortMetadata]
type versionItemShortMetadataJSON struct {
	AuthorEmail apijson.Field
	AuthorID    apijson.Field
	CreatedOn   apijson.Field
	HasPreview  apijson.Field
	ModifiedOn  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionItemShortMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionItemShortMetadataJSON) RawJSON() string {
	return r.raw
}

type VersionItemShortMetadataSource string

const (
	VersionItemShortMetadataSourceUnknown      VersionItemShortMetadataSource = "unknown"
	VersionItemShortMetadataSourceAPI          VersionItemShortMetadataSource = "api"
	VersionItemShortMetadataSourceWrangler     VersionItemShortMetadataSource = "wrangler"
	VersionItemShortMetadataSourceTerraform    VersionItemShortMetadataSource = "terraform"
	VersionItemShortMetadataSourceDash         VersionItemShortMetadataSource = "dash"
	VersionItemShortMetadataSourceDashTemplate VersionItemShortMetadataSource = "dash_template"
	VersionItemShortMetadataSourceIntegration  VersionItemShortMetadataSource = "integration"
	VersionItemShortMetadataSourceQuickEditor  VersionItemShortMetadataSource = "quick_editor"
	VersionItemShortMetadataSourcePlayground   VersionItemShortMetadataSource = "playground"
	VersionItemShortMetadataSourceWorkersci    VersionItemShortMetadataSource = "workersci"
)

func (r VersionItemShortMetadataSource) IsKnown() bool {
	switch r {
	case VersionItemShortMetadataSourceUnknown, VersionItemShortMetadataSourceAPI, VersionItemShortMetadataSourceWrangler, VersionItemShortMetadataSourceTerraform, VersionItemShortMetadataSourceDash, VersionItemShortMetadataSourceDashTemplate, VersionItemShortMetadataSourceIntegration, VersionItemShortMetadataSourceQuickEditor, VersionItemShortMetadataSourcePlayground, VersionItemShortMetadataSourceWorkersci:
		return true
	}
	return false
}

type AccountWorkerScriptVersionListResponse struct {
	Errors   []WorkersMessages                            `json:"errors,required"`
	Messages []WorkersMessages                            `json:"messages,required"`
	Result   AccountWorkerScriptVersionListResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerScriptVersionListResponseSuccess `json:"success,required"`
	JSON    accountWorkerScriptVersionListResponseJSON    `json:"-"`
}

// accountWorkerScriptVersionListResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptVersionListResponse]
type accountWorkerScriptVersionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptVersionListResponseResult struct {
	Items []VersionItemShort                               `json:"items"`
	JSON  accountWorkerScriptVersionListResponseResultJSON `json:"-"`
}

// accountWorkerScriptVersionListResponseResultJSON contains the JSON metadata for
// the struct [AccountWorkerScriptVersionListResponseResult]
type accountWorkerScriptVersionListResponseResultJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionListResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerScriptVersionListResponseSuccess bool

const (
	AccountWorkerScriptVersionListResponseSuccessTrue AccountWorkerScriptVersionListResponseSuccess = true
)

func (r AccountWorkerScriptVersionListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerScriptVersionListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptVersionGetDetailResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	Result   VersionItemFull   `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerScriptVersionGetDetailResponseSuccess `json:"success,required"`
	JSON    accountWorkerScriptVersionGetDetailResponseJSON    `json:"-"`
}

// accountWorkerScriptVersionGetDetailResponseJSON contains the JSON metadata for
// the struct [AccountWorkerScriptVersionGetDetailResponse]
type accountWorkerScriptVersionGetDetailResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionGetDetailResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionGetDetailResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerScriptVersionGetDetailResponseSuccess bool

const (
	AccountWorkerScriptVersionGetDetailResponseSuccessTrue AccountWorkerScriptVersionGetDetailResponseSuccess = true
)

func (r AccountWorkerScriptVersionGetDetailResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerScriptVersionGetDetailResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptVersionUploadResponse struct {
	Errors   []WorkersMessages                              `json:"errors,required"`
	Messages []WorkersMessages                              `json:"messages,required"`
	Result   AccountWorkerScriptVersionUploadResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerScriptVersionUploadResponseSuccess `json:"success,required"`
	JSON    accountWorkerScriptVersionUploadResponseJSON    `json:"-"`
}

// accountWorkerScriptVersionUploadResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptVersionUploadResponse]
type accountWorkerScriptVersionUploadResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionUploadResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptVersionUploadResponseResult struct {
	Resources     AccountWorkerScriptVersionUploadResponseResultResources `json:"resources,required"`
	ID            string                                                  `json:"id"`
	Metadata      AccountWorkerScriptVersionUploadResponseResultMetadata  `json:"metadata"`
	Number        float64                                                 `json:"number"`
	StartupTimeMs int64                                                   `json:"startup_time_ms"`
	JSON          accountWorkerScriptVersionUploadResponseResultJSON      `json:"-"`
}

// accountWorkerScriptVersionUploadResponseResultJSON contains the JSON metadata
// for the struct [AccountWorkerScriptVersionUploadResponseResult]
type accountWorkerScriptVersionUploadResponseResultJSON struct {
	Resources     apijson.Field
	ID            apijson.Field
	Metadata      apijson.Field
	Number        apijson.Field
	StartupTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionUploadResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionUploadResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptVersionUploadResponseResultResources struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings      []BindingItem                                                        `json:"bindings"`
	Script        AccountWorkerScriptVersionUploadResponseResultResourcesScript        `json:"script"`
	ScriptRuntime AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntime `json:"script_runtime"`
	JSON          accountWorkerScriptVersionUploadResponseResultResourcesJSON          `json:"-"`
}

// accountWorkerScriptVersionUploadResponseResultResourcesJSON contains the JSON
// metadata for the struct
// [AccountWorkerScriptVersionUploadResponseResultResources]
type accountWorkerScriptVersionUploadResponseResultResourcesJSON struct {
	Bindings      apijson.Field
	Script        apijson.Field
	ScriptRuntime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionUploadResponseResultResources) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionUploadResponseResultResourcesJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptVersionUploadResponseResultResourcesScript struct {
	Etag             string                                                                      `json:"etag"`
	Handlers         []string                                                                    `json:"handlers"`
	LastDeployedFrom string                                                                      `json:"last_deployed_from"`
	NamedHandlers    []AccountWorkerScriptVersionUploadResponseResultResourcesScriptNamedHandler `json:"named_handlers"`
	JSON             accountWorkerScriptVersionUploadResponseResultResourcesScriptJSON           `json:"-"`
}

// accountWorkerScriptVersionUploadResponseResultResourcesScriptJSON contains the
// JSON metadata for the struct
// [AccountWorkerScriptVersionUploadResponseResultResourcesScript]
type accountWorkerScriptVersionUploadResponseResultResourcesScriptJSON struct {
	Etag             apijson.Field
	Handlers         apijson.Field
	LastDeployedFrom apijson.Field
	NamedHandlers    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionUploadResponseResultResourcesScript) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionUploadResponseResultResourcesScriptJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptVersionUploadResponseResultResourcesScriptNamedHandler struct {
	Handlers []string                                                                      `json:"handlers"`
	Name     string                                                                        `json:"name"`
	JSON     accountWorkerScriptVersionUploadResponseResultResourcesScriptNamedHandlerJSON `json:"-"`
}

// accountWorkerScriptVersionUploadResponseResultResourcesScriptNamedHandlerJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptVersionUploadResponseResultResourcesScriptNamedHandler]
type accountWorkerScriptVersionUploadResponseResultResourcesScriptNamedHandlerJSON struct {
	Handlers    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionUploadResponseResultResourcesScriptNamedHandler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionUploadResponseResultResourcesScriptNamedHandlerJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntime struct {
	CompatibilityDate  string                                                                         `json:"compatibility_date"`
	CompatibilityFlags []string                                                                       `json:"compatibility_flags"`
	Limits             AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeLimits     `json:"limits"`
	MigrationTag       string                                                                         `json:"migration_tag"`
	UsageModel         AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeUsageModel `json:"usage_model"`
	JSON               accountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeJSON       `json:"-"`
}

// accountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntime]
type accountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeJSON struct {
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Limits             apijson.Field
	MigrationTag       apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeLimits struct {
	CPUMs int64                                                                          `json:"cpu_ms"`
	JSON  accountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeLimitsJSON `json:"-"`
}

// accountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeLimitsJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeLimits]
type accountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeLimitsJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeUsageModel string

const (
	AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeUsageModelBundled  AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeUsageModel = "bundled"
	AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeUsageModelUnbound  AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeUsageModel = "unbound"
	AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeUsageModelStandard AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeUsageModel = "standard"
)

func (r AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeUsageModel) IsKnown() bool {
	switch r {
	case AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeUsageModelBundled, AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeUsageModelUnbound, AccountWorkerScriptVersionUploadResponseResultResourcesScriptRuntimeUsageModelStandard:
		return true
	}
	return false
}

type AccountWorkerScriptVersionUploadResponseResultMetadata struct {
	AuthorEmail string                                                       `json:"author_email"`
	AuthorID    string                                                       `json:"author_id"`
	CreatedOn   string                                                       `json:"created_on"`
	HasPreview  bool                                                         `json:"hasPreview"`
	ModifiedOn  string                                                       `json:"modified_on"`
	Source      AccountWorkerScriptVersionUploadResponseResultMetadataSource `json:"source"`
	JSON        accountWorkerScriptVersionUploadResponseResultMetadataJSON   `json:"-"`
}

// accountWorkerScriptVersionUploadResponseResultMetadataJSON contains the JSON
// metadata for the struct [AccountWorkerScriptVersionUploadResponseResultMetadata]
type accountWorkerScriptVersionUploadResponseResultMetadataJSON struct {
	AuthorEmail apijson.Field
	AuthorID    apijson.Field
	CreatedOn   apijson.Field
	HasPreview  apijson.Field
	ModifiedOn  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionUploadResponseResultMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionUploadResponseResultMetadataJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptVersionUploadResponseResultMetadataSource string

const (
	AccountWorkerScriptVersionUploadResponseResultMetadataSourceUnknown      AccountWorkerScriptVersionUploadResponseResultMetadataSource = "unknown"
	AccountWorkerScriptVersionUploadResponseResultMetadataSourceAPI          AccountWorkerScriptVersionUploadResponseResultMetadataSource = "api"
	AccountWorkerScriptVersionUploadResponseResultMetadataSourceWrangler     AccountWorkerScriptVersionUploadResponseResultMetadataSource = "wrangler"
	AccountWorkerScriptVersionUploadResponseResultMetadataSourceTerraform    AccountWorkerScriptVersionUploadResponseResultMetadataSource = "terraform"
	AccountWorkerScriptVersionUploadResponseResultMetadataSourceDash         AccountWorkerScriptVersionUploadResponseResultMetadataSource = "dash"
	AccountWorkerScriptVersionUploadResponseResultMetadataSourceDashTemplate AccountWorkerScriptVersionUploadResponseResultMetadataSource = "dash_template"
	AccountWorkerScriptVersionUploadResponseResultMetadataSourceIntegration  AccountWorkerScriptVersionUploadResponseResultMetadataSource = "integration"
	AccountWorkerScriptVersionUploadResponseResultMetadataSourceQuickEditor  AccountWorkerScriptVersionUploadResponseResultMetadataSource = "quick_editor"
	AccountWorkerScriptVersionUploadResponseResultMetadataSourcePlayground   AccountWorkerScriptVersionUploadResponseResultMetadataSource = "playground"
	AccountWorkerScriptVersionUploadResponseResultMetadataSourceWorkersci    AccountWorkerScriptVersionUploadResponseResultMetadataSource = "workersci"
)

func (r AccountWorkerScriptVersionUploadResponseResultMetadataSource) IsKnown() bool {
	switch r {
	case AccountWorkerScriptVersionUploadResponseResultMetadataSourceUnknown, AccountWorkerScriptVersionUploadResponseResultMetadataSourceAPI, AccountWorkerScriptVersionUploadResponseResultMetadataSourceWrangler, AccountWorkerScriptVersionUploadResponseResultMetadataSourceTerraform, AccountWorkerScriptVersionUploadResponseResultMetadataSourceDash, AccountWorkerScriptVersionUploadResponseResultMetadataSourceDashTemplate, AccountWorkerScriptVersionUploadResponseResultMetadataSourceIntegration, AccountWorkerScriptVersionUploadResponseResultMetadataSourceQuickEditor, AccountWorkerScriptVersionUploadResponseResultMetadataSourcePlayground, AccountWorkerScriptVersionUploadResponseResultMetadataSourceWorkersci:
		return true
	}
	return false
}

// Whether the API call was successful.
type AccountWorkerScriptVersionUploadResponseSuccess bool

const (
	AccountWorkerScriptVersionUploadResponseSuccessTrue AccountWorkerScriptVersionUploadResponseSuccess = true
)

func (r AccountWorkerScriptVersionUploadResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerScriptVersionUploadResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptVersionListParams struct {
	// Only return versions that can be used in a deployment. Ignores pagination.
	Deployable param.Field[bool] `query:"deployable"`
	// Current page.
	Page param.Field[int64] `query:"page"`
	// Items per-page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [AccountWorkerScriptVersionListParams]'s query parameters as
// `url.Values`.
func (r AccountWorkerScriptVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountWorkerScriptVersionUploadParams struct {
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[AccountWorkerScriptVersionUploadParamsMetadata] `json:"metadata,required"`
	// An array of modules (often JavaScript files) comprising a Worker script. At
	// least one module must be present and referenced in the metadata as `main_module`
	// or `body_part` by filename.<br/>Possible Content-Type(s) are:
	// `application/javascript+module`, `text/javascript+module`,
	// `application/javascript`, `text/javascript`, `application/wasm`, `text/plain`,
	// `application/octet-stream`, `application/source-map`.
	Files param.Field[[]io.Reader] `json:"files" format:"binary"`
}

func (r AccountWorkerScriptVersionUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type AccountWorkerScriptVersionUploadParamsMetadata struct {
	// Name of the uploaded file that contains the main module (e.g. the file exporting
	// a `fetch` handler). Indicates a `module syntax` Worker, which is required for
	// Version Upload.
	MainModule  param.Field[string]                                                    `json:"main_module,required"`
	Annotations param.Field[AccountWorkerScriptVersionUploadParamsMetadataAnnotations] `json:"annotations"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]BindingItemUnionParam] `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// List of binding types to keep from previous_upload.
	KeepBindings param.Field[[]string] `json:"keep_bindings"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[UsageModel] `json:"usage_model"`
}

func (r AccountWorkerScriptVersionUploadParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptVersionUploadParamsMetadataAnnotations struct {
	// Human-readable message about the version. Truncated to 100 bytes.
	WorkersMessage param.Field[string] `json:"workers/message"`
	// User-provided identifier for the version.
	WorkersTag param.Field[string] `json:"workers/tag"`
}

func (r AccountWorkerScriptVersionUploadParamsMetadataAnnotations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
