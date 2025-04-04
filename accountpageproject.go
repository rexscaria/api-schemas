// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/tidwall/gjson"
)

// AccountPageProjectService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPageProjectService] method instead.
type AccountPageProjectService struct {
	Options     []option.RequestOption
	Deployments *AccountPageProjectDeploymentService
	Domains     *AccountPageProjectDomainService
}

// NewAccountPageProjectService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountPageProjectService(opts ...option.RequestOption) (r *AccountPageProjectService) {
	r = &AccountPageProjectService{}
	r.Options = opts
	r.Deployments = NewAccountPageProjectDeploymentService(opts...)
	r.Domains = NewAccountPageProjectDomainService(opts...)
	return
}

// Create a new project.
func (r *AccountPageProjectService) New(ctx context.Context, accountID string, body AccountPageProjectNewParams, opts ...option.RequestOption) (res *ProjectResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a project by name.
func (r *AccountPageProjectService) Get(ctx context.Context, accountID string, projectName string, opts ...option.RequestOption) (res *ProjectResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", accountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set new attributes for an existing project. Modify environment variables. To
// delete an environment variable, set the key to null.
func (r *AccountPageProjectService) Update(ctx context.Context, accountID string, projectName string, body AccountPageProjectUpdateParams, opts ...option.RequestOption) (res *ProjectResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", accountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetch a list of all user projects.
func (r *AccountPageProjectService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountPageProjectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a project by name.
func (r *AccountPageProjectService) Delete(ctx context.Context, accountID string, projectName string, body AccountPageProjectDeleteParams, opts ...option.RequestOption) (res *AccountPageProjectDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", accountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Purge all cached build artifacts for a Pages project
func (r *AccountPageProjectService) PurgeBuildCache(ctx context.Context, accountID string, projectName string, opts ...option.RequestOption) (res *AccountPageProjectPurgeBuildCacheResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/purge_build_cache", accountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type APIResponsePages struct {
	Errors   []MessagesPageItem `json:"errors,required"`
	Messages []MessagesPageItem `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponsePagesSuccess `json:"success,required"`
	JSON    apiResponsePagesJSON    `json:"-"`
}

// apiResponsePagesJSON contains the JSON metadata for the struct
// [APIResponsePages]
type apiResponsePagesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponsePages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponsePagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponsePagesSuccess bool

const (
	APIResponsePagesSuccessFalse APIResponsePagesSuccess = false
	APIResponsePagesSuccessTrue  APIResponsePagesSuccess = true
)

func (r APIResponsePagesSuccess) IsKnown() bool {
	switch r {
	case APIResponsePagesSuccessFalse, APIResponsePagesSuccessTrue:
		return true
	}
	return false
}

type APIResponsePagination struct {
	ResultInfo APIResponsePaginationResultInfo `json:"result_info"`
	JSON       apiResponsePaginationJSON       `json:"-"`
}

// apiResponsePaginationJSON contains the JSON metadata for the struct
// [APIResponsePagination]
type apiResponsePaginationJSON struct {
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type APIResponsePaginationResultInfo struct {
	// The number of items on the current page.
	Count int64 `json:"count,required"`
	// The page currently being requested.
	Page int64 `json:"page,required"`
	// The number of items per page being returned.
	PerPage int64 `json:"per_page,required"`
	// The total count of items.
	TotalCount int64 `json:"total_count,required"`
	// The total count of pages.
	TotalPages int64                               `json:"total_pages"`
	JSON       apiResponsePaginationResultInfoJSON `json:"-"`
}

// apiResponsePaginationResultInfoJSON contains the JSON metadata for the struct
// [APIResponsePaginationResultInfo]
type apiResponsePaginationResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponsePaginationResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponsePaginationResultInfoJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type BuildConfig struct {
	// Enable build caching for the project.
	BuildCaching bool `json:"build_caching,nullable"`
	// Command used to build project.
	BuildCommand string `json:"build_command,nullable"`
	// Output directory of the build.
	DestinationDir string `json:"destination_dir,nullable"`
	// Directory to run the command.
	RootDir string `json:"root_dir,nullable"`
	// The classifying tag for analytics.
	WebAnalyticsTag string `json:"web_analytics_tag,nullable"`
	// The auth token for analytics.
	WebAnalyticsToken string          `json:"web_analytics_token,nullable"`
	JSON              buildConfigJSON `json:"-"`
}

// buildConfigJSON contains the JSON metadata for the struct [BuildConfig]
type buildConfigJSON struct {
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *BuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildConfigJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type BuildConfigParam struct {
	// Enable build caching for the project.
	BuildCaching param.Field[bool] `json:"build_caching"`
	// Command used to build project.
	BuildCommand param.Field[string] `json:"build_command"`
	// Output directory of the build.
	DestinationDir param.Field[string] `json:"destination_dir"`
	// Directory to run the command.
	RootDir param.Field[string] `json:"root_dir"`
	// The classifying tag for analytics.
	WebAnalyticsTag param.Field[string] `json:"web_analytics_tag"`
	// The auth token for analytics.
	WebAnalyticsToken param.Field[string] `json:"web_analytics_token"`
}

func (r BuildConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeploymentConfigsValues struct {
	// Constellation bindings used for Pages Functions.
	AIBindings map[string]DeploymentConfigsValuesAIBinding `json:"ai_bindings,nullable"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets map[string]DeploymentConfigsValuesAnalyticsEngineDataset `json:"analytics_engine_datasets,nullable"`
	// Browser bindings used for Pages Functions.
	Browsers map[string]DeploymentConfigsValuesBrowser `json:"browsers,nullable"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases map[string]DeploymentConfigsValuesD1Database `json:"d1_databases,nullable"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces map[string]DeploymentConfigsValuesDurableObjectNamespace `json:"durable_object_namespaces,nullable"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]DeploymentConfigsValuesEnvVar `json:"env_vars"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings map[string]DeploymentConfigsValuesHyperdriveBinding `json:"hyperdrive_bindings,nullable"`
	// KV namespaces used for Pages Functions.
	KvNamespaces map[string]DeploymentConfigsValuesKvNamespace `json:"kv_namespaces,nullable"`
	// mTLS bindings used for Pages Functions.
	MtlsCertificates map[string]DeploymentConfigsValuesMtlsCertificate `json:"mtls_certificates,nullable"`
	// Placement setting used for Pages Functions.
	Placement DeploymentConfigsValuesPlacement `json:"placement,nullable"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers map[string]DeploymentConfigsValuesQueueProducer `json:"queue_producers,nullable"`
	// R2 buckets used for Pages Functions.
	R2Buckets map[string]DeploymentConfigsValuesR2Bucket `json:"r2_buckets,nullable"`
	// Services used for Pages Functions.
	Services map[string]DeploymentConfigsValuesService `json:"services,nullable"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings map[string]DeploymentConfigsValuesVectorizeBinding `json:"vectorize_bindings,nullable"`
	JSON              deploymentConfigsValuesJSON                        `json:"-"`
}

// deploymentConfigsValuesJSON contains the JSON metadata for the struct
// [DeploymentConfigsValues]
type deploymentConfigsValuesJSON struct {
	AIBindings              apijson.Field
	AnalyticsEngineDatasets apijson.Field
	Browsers                apijson.Field
	CompatibilityDate       apijson.Field
	CompatibilityFlags      apijson.Field
	D1Databases             apijson.Field
	DurableObjectNamespaces apijson.Field
	EnvVars                 apijson.Field
	HyperdriveBindings      apijson.Field
	KvNamespaces            apijson.Field
	MtlsCertificates        apijson.Field
	Placement               apijson.Field
	QueueProducers          apijson.Field
	R2Buckets               apijson.Field
	Services                apijson.Field
	VectorizeBindings       apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *DeploymentConfigsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesJSON) RawJSON() string {
	return r.raw
}

// AI binding.
type DeploymentConfigsValuesAIBinding struct {
	ProjectID string                               `json:"project_id"`
	JSON      deploymentConfigsValuesAIBindingJSON `json:"-"`
}

// deploymentConfigsValuesAIBindingJSON contains the JSON metadata for the struct
// [DeploymentConfigsValuesAIBinding]
type deploymentConfigsValuesAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentConfigsValuesAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type DeploymentConfigsValuesAnalyticsEngineDataset struct {
	// Name of the dataset.
	Dataset string                                            `json:"dataset"`
	JSON    deploymentConfigsValuesAnalyticsEngineDatasetJSON `json:"-"`
}

// deploymentConfigsValuesAnalyticsEngineDatasetJSON contains the JSON metadata for
// the struct [DeploymentConfigsValuesAnalyticsEngineDataset]
type deploymentConfigsValuesAnalyticsEngineDatasetJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentConfigsValuesAnalyticsEngineDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesAnalyticsEngineDatasetJSON) RawJSON() string {
	return r.raw
}

// Browser binding.
type DeploymentConfigsValuesBrowser struct {
	JSON deploymentConfigsValuesBrowserJSON `json:"-"`
}

// deploymentConfigsValuesBrowserJSON contains the JSON metadata for the struct
// [DeploymentConfigsValuesBrowser]
type deploymentConfigsValuesBrowserJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentConfigsValuesBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesBrowserJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type DeploymentConfigsValuesD1Database struct {
	// UUID of the D1 database.
	ID   string                                `json:"id"`
	JSON deploymentConfigsValuesD1DatabaseJSON `json:"-"`
}

// deploymentConfigsValuesD1DatabaseJSON contains the JSON metadata for the struct
// [DeploymentConfigsValuesD1Database]
type deploymentConfigsValuesD1DatabaseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentConfigsValuesD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesD1DatabaseJSON) RawJSON() string {
	return r.raw
}

// Durable Object binding.
type DeploymentConfigsValuesDurableObjectNamespace struct {
	// ID of the Durable Object namespace.
	NamespaceID string                                            `json:"namespace_id"`
	JSON        deploymentConfigsValuesDurableObjectNamespaceJSON `json:"-"`
}

// deploymentConfigsValuesDurableObjectNamespaceJSON contains the JSON metadata for
// the struct [DeploymentConfigsValuesDurableObjectNamespace]
type deploymentConfigsValuesDurableObjectNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentConfigsValuesDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

// A plaintext environment variable.
type DeploymentConfigsValuesEnvVar struct {
	Type DeploymentConfigsValuesEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                            `json:"value,required"`
	JSON  deploymentConfigsValuesEnvVarJSON `json:"-"`
	union DeploymentConfigsValuesEnvVarsUnion
}

// deploymentConfigsValuesEnvVarJSON contains the JSON metadata for the struct
// [DeploymentConfigsValuesEnvVar]
type deploymentConfigsValuesEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r deploymentConfigsValuesEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *DeploymentConfigsValuesEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = DeploymentConfigsValuesEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DeploymentConfigsValuesEnvVarsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVar],
// [DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVar].
func (r DeploymentConfigsValuesEnvVar) AsUnion() DeploymentConfigsValuesEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by [DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVar] or
// [DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVar].
type DeploymentConfigsValuesEnvVarsUnion interface {
	implementsDeploymentConfigsValuesEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DeploymentConfigsValuesEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVar struct {
	Type DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                 `json:"value,required"`
	JSON  deploymentConfigsValuesEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// deploymentConfigsValuesEnvVarsPagesPlainTextEnvVarJSON contains the JSON
// metadata for the struct [DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVar]
type deploymentConfigsValuesEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVar) implementsDeploymentConfigsValuesEnvVar() {
}

type DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarType string

const (
	DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarTypePlainText DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVar struct {
	Type DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                  `json:"value,required"`
	JSON  deploymentConfigsValuesEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// deploymentConfigsValuesEnvVarsPagesSecretTextEnvVarJSON contains the JSON
// metadata for the struct [DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVar]
type deploymentConfigsValuesEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVar) implementsDeploymentConfigsValuesEnvVar() {
}

type DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVarType string

const (
	DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVarTypeSecretText DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type DeploymentConfigsValuesEnvVarsType string

const (
	DeploymentConfigsValuesEnvVarsTypePlainText  DeploymentConfigsValuesEnvVarsType = "plain_text"
	DeploymentConfigsValuesEnvVarsTypeSecretText DeploymentConfigsValuesEnvVarsType = "secret_text"
)

func (r DeploymentConfigsValuesEnvVarsType) IsKnown() bool {
	switch r {
	case DeploymentConfigsValuesEnvVarsTypePlainText, DeploymentConfigsValuesEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive binding.
type DeploymentConfigsValuesHyperdriveBinding struct {
	ID   string                                       `json:"id"`
	JSON deploymentConfigsValuesHyperdriveBindingJSON `json:"-"`
}

// deploymentConfigsValuesHyperdriveBindingJSON contains the JSON metadata for the
// struct [DeploymentConfigsValuesHyperdriveBinding]
type deploymentConfigsValuesHyperdriveBindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentConfigsValuesHyperdriveBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesHyperdriveBindingJSON) RawJSON() string {
	return r.raw
}

// KV namespace binding.
type DeploymentConfigsValuesKvNamespace struct {
	// ID of the KV namespace.
	NamespaceID string                                 `json:"namespace_id"`
	JSON        deploymentConfigsValuesKvNamespaceJSON `json:"-"`
}

// deploymentConfigsValuesKvNamespaceJSON contains the JSON metadata for the struct
// [DeploymentConfigsValuesKvNamespace]
type deploymentConfigsValuesKvNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentConfigsValuesKvNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesKvNamespaceJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type DeploymentConfigsValuesMtlsCertificate struct {
	CertificateID string                                     `json:"certificate_id"`
	JSON          deploymentConfigsValuesMtlsCertificateJSON `json:"-"`
}

// deploymentConfigsValuesMtlsCertificateJSON contains the JSON metadata for the
// struct [DeploymentConfigsValuesMtlsCertificate]
type deploymentConfigsValuesMtlsCertificateJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeploymentConfigsValuesMtlsCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesMtlsCertificateJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type DeploymentConfigsValuesPlacement struct {
	// Placement mode.
	Mode string                               `json:"mode"`
	JSON deploymentConfigsValuesPlacementJSON `json:"-"`
}

// deploymentConfigsValuesPlacementJSON contains the JSON metadata for the struct
// [DeploymentConfigsValuesPlacement]
type deploymentConfigsValuesPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentConfigsValuesPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type DeploymentConfigsValuesQueueProducer struct {
	// Name of the Queue.
	Name string                                   `json:"name"`
	JSON deploymentConfigsValuesQueueProducerJSON `json:"-"`
}

// deploymentConfigsValuesQueueProducerJSON contains the JSON metadata for the
// struct [DeploymentConfigsValuesQueueProducer]
type deploymentConfigsValuesQueueProducerJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentConfigsValuesQueueProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesQueueProducerJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type DeploymentConfigsValuesR2Bucket struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction string `json:"jurisdiction,nullable"`
	// Name of the R2 bucket.
	Name string                              `json:"name"`
	JSON deploymentConfigsValuesR2BucketJSON `json:"-"`
}

// deploymentConfigsValuesR2BucketJSON contains the JSON metadata for the struct
// [DeploymentConfigsValuesR2Bucket]
type deploymentConfigsValuesR2BucketJSON struct {
	Jurisdiction apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DeploymentConfigsValuesR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesR2BucketJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type DeploymentConfigsValuesService struct {
	// The entrypoint to bind to.
	Entrypoint string `json:"entrypoint,nullable"`
	// The Service environment.
	Environment string `json:"environment"`
	// The Service name.
	Service string                             `json:"service"`
	JSON    deploymentConfigsValuesServiceJSON `json:"-"`
}

// deploymentConfigsValuesServiceJSON contains the JSON metadata for the struct
// [DeploymentConfigsValuesService]
type deploymentConfigsValuesServiceJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentConfigsValuesService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesServiceJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type DeploymentConfigsValuesVectorizeBinding struct {
	IndexName string                                      `json:"index_name"`
	JSON      deploymentConfigsValuesVectorizeBindingJSON `json:"-"`
}

// deploymentConfigsValuesVectorizeBindingJSON contains the JSON metadata for the
// struct [DeploymentConfigsValuesVectorizeBinding]
type deploymentConfigsValuesVectorizeBindingJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentConfigsValuesVectorizeBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentConfigsValuesVectorizeBindingJSON) RawJSON() string {
	return r.raw
}

type DeploymentConfigsValuesParam struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[map[string]DeploymentConfigsValuesAIBindingParam] `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[map[string]DeploymentConfigsValuesAnalyticsEngineDatasetParam] `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers param.Field[map[string]DeploymentConfigsValuesBrowserParam] `json:"browsers"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[map[string]DeploymentConfigsValuesD1DatabaseParam] `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[map[string]DeploymentConfigsValuesDurableObjectNamespaceParam] `json:"durable_object_namespaces"`
	// Environment variables used for builds and Pages Functions.
	EnvVars param.Field[map[string]DeploymentConfigsValuesEnvVarsUnionParam] `json:"env_vars"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings param.Field[map[string]DeploymentConfigsValuesHyperdriveBindingParam] `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KvNamespaces param.Field[map[string]DeploymentConfigsValuesKvNamespaceParam] `json:"kv_namespaces"`
	// mTLS bindings used for Pages Functions.
	MtlsCertificates param.Field[map[string]DeploymentConfigsValuesMtlsCertificateParam] `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement param.Field[DeploymentConfigsValuesPlacementParam] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[map[string]DeploymentConfigsValuesQueueProducerParam] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[map[string]DeploymentConfigsValuesR2BucketParam] `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services param.Field[map[string]DeploymentConfigsValuesServiceParam] `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings param.Field[map[string]DeploymentConfigsValuesVectorizeBindingParam] `json:"vectorize_bindings"`
}

func (r DeploymentConfigsValuesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type DeploymentConfigsValuesAIBindingParam struct {
	ProjectID param.Field[string] `json:"project_id"`
}

func (r DeploymentConfigsValuesAIBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type DeploymentConfigsValuesAnalyticsEngineDatasetParam struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset"`
}

func (r DeploymentConfigsValuesAnalyticsEngineDatasetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser binding.
type DeploymentConfigsValuesBrowserParam struct {
}

func (r DeploymentConfigsValuesBrowserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type DeploymentConfigsValuesD1DatabaseParam struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id"`
}

func (r DeploymentConfigsValuesD1DatabaseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durable Object binding.
type DeploymentConfigsValuesDurableObjectNamespaceParam struct {
	// ID of the Durable Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r DeploymentConfigsValuesDurableObjectNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A plaintext environment variable.
type DeploymentConfigsValuesEnvVarParam struct {
	Type param.Field[DeploymentConfigsValuesEnvVarsType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r DeploymentConfigsValuesEnvVarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeploymentConfigsValuesEnvVarParam) implementsDeploymentConfigsValuesEnvVarsUnionParam() {}

// A plaintext environment variable.
//
// Satisfied by [DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarParam],
// [DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVarParam],
// [DeploymentConfigsValuesEnvVarParam].
type DeploymentConfigsValuesEnvVarsUnionParam interface {
	implementsDeploymentConfigsValuesEnvVarsUnionParam()
}

// A plaintext environment variable.
type DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarParam struct {
	Type param.Field[DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarParam) implementsDeploymentConfigsValuesEnvVarsUnionParam() {
}

// An encrypted environment variable.
type DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVarParam struct {
	Type param.Field[DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVarType] `json:"type,required"`
	// Secret value.
	Value param.Field[string] `json:"value,required"`
}

func (r DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeploymentConfigsValuesEnvVarsPagesSecretTextEnvVarParam) implementsDeploymentConfigsValuesEnvVarsUnionParam() {
}

// Hyperdrive binding.
type DeploymentConfigsValuesHyperdriveBindingParam struct {
	ID param.Field[string] `json:"id"`
}

func (r DeploymentConfigsValuesHyperdriveBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV namespace binding.
type DeploymentConfigsValuesKvNamespaceParam struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id"`
}

func (r DeploymentConfigsValuesKvNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS binding.
type DeploymentConfigsValuesMtlsCertificateParam struct {
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r DeploymentConfigsValuesMtlsCertificateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Placement setting used for Pages Functions.
type DeploymentConfigsValuesPlacementParam struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode"`
}

func (r DeploymentConfigsValuesPlacementParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type DeploymentConfigsValuesQueueProducerParam struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name"`
}

func (r DeploymentConfigsValuesQueueProducerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type DeploymentConfigsValuesR2BucketParam struct {
	// Jurisdiction of the R2 bucket.
	Jurisdiction param.Field[string] `json:"jurisdiction"`
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name"`
}

func (r DeploymentConfigsValuesR2BucketParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type DeploymentConfigsValuesServiceParam struct {
	// The entrypoint to bind to.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
	// The Service name.
	Service param.Field[string] `json:"service"`
}

func (r DeploymentConfigsValuesServiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Vectorize binding.
type DeploymentConfigsValuesVectorizeBindingParam struct {
	IndexName param.Field[string] `json:"index_name"`
}

func (r DeploymentConfigsValuesVectorizeBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Deployments struct {
	// Id of the deployment.
	ID string `json:"id"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,nullable"`
	// Configs for the project build process.
	BuildConfig BuildConfig `json:"build_config"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger DeploymentsDeploymentTrigger `json:"deployment_trigger"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]DeploymentsEnvVar `json:"env_vars"`
	// Type of deploy.
	Environment DeploymentsEnvironment `json:"environment"`
	// If the deployment has been skipped.
	IsSkipped bool `json:"is_skipped"`
	// The status of the deployment.
	LatestStage Stage `json:"latest_stage"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id"`
	// Name of the project.
	ProjectName string `json:"project_name"`
	// Short Id (8 character) of the deployment.
	ShortID string `json:"short_id"`
	Source  Source `json:"source"`
	// List of past stages.
	Stages []Stage `json:"stages"`
	// The live URL to view this deployment.
	URL  string          `json:"url"`
	JSON deploymentsJSON `json:"-"`
}

// deploymentsJSON contains the JSON metadata for the struct [Deployments]
type deploymentsJSON struct {
	ID                apijson.Field
	Aliases           apijson.Field
	BuildConfig       apijson.Field
	CreatedOn         apijson.Field
	DeploymentTrigger apijson.Field
	EnvVars           apijson.Field
	Environment       apijson.Field
	IsSkipped         apijson.Field
	LatestStage       apijson.Field
	ModifiedOn        apijson.Field
	ProjectID         apijson.Field
	ProjectName       apijson.Field
	ShortID           apijson.Field
	Source            apijson.Field
	Stages            apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Deployments) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentsJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type DeploymentsDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata DeploymentsDeploymentTriggerMetadata `json:"metadata"`
	// What caused the deployment.
	Type DeploymentsDeploymentTriggerType `json:"type"`
	JSON deploymentsDeploymentTriggerJSON `json:"-"`
}

// deploymentsDeploymentTriggerJSON contains the JSON metadata for the struct
// [DeploymentsDeploymentTrigger]
type deploymentsDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentsDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentsDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type DeploymentsDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash"`
	// Message of the deployment trigger commit.
	CommitMessage string                                   `json:"commit_message"`
	JSON          deploymentsDeploymentTriggerMetadataJSON `json:"-"`
}

// deploymentsDeploymentTriggerMetadataJSON contains the JSON metadata for the
// struct [DeploymentsDeploymentTriggerMetadata]
type deploymentsDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeploymentsDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentsDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type DeploymentsDeploymentTriggerType string

const (
	DeploymentsDeploymentTriggerTypePush  DeploymentsDeploymentTriggerType = "push"
	DeploymentsDeploymentTriggerTypeAdHoc DeploymentsDeploymentTriggerType = "ad_hoc"
)

func (r DeploymentsDeploymentTriggerType) IsKnown() bool {
	switch r {
	case DeploymentsDeploymentTriggerTypePush, DeploymentsDeploymentTriggerTypeAdHoc:
		return true
	}
	return false
}

// A plaintext environment variable.
type DeploymentsEnvVar struct {
	Type DeploymentsEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                `json:"value,required"`
	JSON  deploymentsEnvVarJSON `json:"-"`
	union DeploymentsEnvVarsUnion
}

// deploymentsEnvVarJSON contains the JSON metadata for the struct
// [DeploymentsEnvVar]
type deploymentsEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r deploymentsEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *DeploymentsEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = DeploymentsEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DeploymentsEnvVarsUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [DeploymentsEnvVarsPagesPlainTextEnvVar],
// [DeploymentsEnvVarsPagesSecretTextEnvVar].
func (r DeploymentsEnvVar) AsUnion() DeploymentsEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by [DeploymentsEnvVarsPagesPlainTextEnvVar] or
// [DeploymentsEnvVarsPagesSecretTextEnvVar].
type DeploymentsEnvVarsUnion interface {
	implementsDeploymentsEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DeploymentsEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DeploymentsEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DeploymentsEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type DeploymentsEnvVarsPagesPlainTextEnvVar struct {
	Type DeploymentsEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                     `json:"value,required"`
	JSON  deploymentsEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// deploymentsEnvVarsPagesPlainTextEnvVarJSON contains the JSON metadata for the
// struct [DeploymentsEnvVarsPagesPlainTextEnvVar]
type deploymentsEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentsEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentsEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r DeploymentsEnvVarsPagesPlainTextEnvVar) implementsDeploymentsEnvVar() {}

type DeploymentsEnvVarsPagesPlainTextEnvVarType string

const (
	DeploymentsEnvVarsPagesPlainTextEnvVarTypePlainText DeploymentsEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r DeploymentsEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case DeploymentsEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type DeploymentsEnvVarsPagesSecretTextEnvVar struct {
	Type DeploymentsEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                      `json:"value,required"`
	JSON  deploymentsEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// deploymentsEnvVarsPagesSecretTextEnvVarJSON contains the JSON metadata for the
// struct [DeploymentsEnvVarsPagesSecretTextEnvVar]
type deploymentsEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentsEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentsEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r DeploymentsEnvVarsPagesSecretTextEnvVar) implementsDeploymentsEnvVar() {}

type DeploymentsEnvVarsPagesSecretTextEnvVarType string

const (
	DeploymentsEnvVarsPagesSecretTextEnvVarTypeSecretText DeploymentsEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r DeploymentsEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case DeploymentsEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type DeploymentsEnvVarsType string

const (
	DeploymentsEnvVarsTypePlainText  DeploymentsEnvVarsType = "plain_text"
	DeploymentsEnvVarsTypeSecretText DeploymentsEnvVarsType = "secret_text"
)

func (r DeploymentsEnvVarsType) IsKnown() bool {
	switch r {
	case DeploymentsEnvVarsTypePlainText, DeploymentsEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type DeploymentsEnvironment string

const (
	DeploymentsEnvironmentPreview    DeploymentsEnvironment = "preview"
	DeploymentsEnvironmentProduction DeploymentsEnvironment = "production"
)

func (r DeploymentsEnvironment) IsKnown() bool {
	switch r {
	case DeploymentsEnvironmentPreview, DeploymentsEnvironmentProduction:
		return true
	}
	return false
}

type DeploymentsParam struct {
	// Configs for the project build process.
	BuildConfig param.Field[BuildConfigParam] `json:"build_config"`
	// Environment variables used for builds and Pages Functions.
	EnvVars param.Field[map[string]DeploymentsEnvVarsUnionParam] `json:"env_vars"`
}

func (r DeploymentsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Info about what caused the deployment.
type DeploymentsDeploymentTriggerParam struct {
	// Additional info about the trigger.
	Metadata param.Field[DeploymentsDeploymentTriggerMetadataParam] `json:"metadata"`
}

func (r DeploymentsDeploymentTriggerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional info about the trigger.
type DeploymentsDeploymentTriggerMetadataParam struct {
}

func (r DeploymentsDeploymentTriggerMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A plaintext environment variable.
type DeploymentsEnvVarParam struct {
	Type param.Field[DeploymentsEnvVarsType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r DeploymentsEnvVarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeploymentsEnvVarParam) implementsDeploymentsEnvVarsUnionParam() {}

// A plaintext environment variable.
//
// Satisfied by [DeploymentsEnvVarsPagesPlainTextEnvVarParam],
// [DeploymentsEnvVarsPagesSecretTextEnvVarParam], [DeploymentsEnvVarParam].
type DeploymentsEnvVarsUnionParam interface {
	implementsDeploymentsEnvVarsUnionParam()
}

// A plaintext environment variable.
type DeploymentsEnvVarsPagesPlainTextEnvVarParam struct {
	Type param.Field[DeploymentsEnvVarsPagesPlainTextEnvVarType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r DeploymentsEnvVarsPagesPlainTextEnvVarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeploymentsEnvVarsPagesPlainTextEnvVarParam) implementsDeploymentsEnvVarsUnionParam() {}

// An encrypted environment variable.
type DeploymentsEnvVarsPagesSecretTextEnvVarParam struct {
	Type param.Field[DeploymentsEnvVarsPagesSecretTextEnvVarType] `json:"type,required"`
	// Secret value.
	Value param.Field[string] `json:"value,required"`
}

func (r DeploymentsEnvVarsPagesSecretTextEnvVarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeploymentsEnvVarsPagesSecretTextEnvVarParam) implementsDeploymentsEnvVarsUnionParam() {}

type MessagesPageItem struct {
	Code    int64                `json:"code,required"`
	Message string               `json:"message,required"`
	JSON    messagesPageItemJSON `json:"-"`
}

// messagesPageItemJSON contains the JSON metadata for the struct
// [MessagesPageItem]
type messagesPageItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesPageItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesPageItemJSON) RawJSON() string {
	return r.raw
}

type Project struct {
	// Id of the project.
	ID string `json:"id"`
	// Configs for the project build process.
	BuildConfig BuildConfig `json:"build_config"`
	// Most recent deployment to the repo.
	CanonicalDeployment ProjectCanonicalDeployment `json:"canonical_deployment"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs ProjectDeploymentConfigs `json:"deployment_configs"`
	// A list of associated custom domains for the project.
	Domains []string `json:"domains"`
	// Most recent deployment to the repo.
	LatestDeployment ProjectLatestDeployment `json:"latest_deployment"`
	// Name of the project.
	Name string `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch string `json:"production_branch"`
	Source           Source `json:"source"`
	// The Cloudflare subdomain associated with the project.
	Subdomain string      `json:"subdomain"`
	JSON      projectJSON `json:"-"`
}

// projectJSON contains the JSON metadata for the struct [Project]
type projectJSON struct {
	ID                  apijson.Field
	BuildConfig         apijson.Field
	CanonicalDeployment apijson.Field
	CreatedOn           apijson.Field
	DeploymentConfigs   apijson.Field
	Domains             apijson.Field
	LatestDeployment    apijson.Field
	Name                apijson.Field
	ProductionBranch    apijson.Field
	Source              apijson.Field
	Subdomain           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Project) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectJSON) RawJSON() string {
	return r.raw
}

// Most recent deployment to the repo.
type ProjectCanonicalDeployment struct {
	JSON projectCanonicalDeploymentJSON `json:"-"`
	Deployments
}

// projectCanonicalDeploymentJSON contains the JSON metadata for the struct
// [ProjectCanonicalDeployment]
type projectCanonicalDeploymentJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectCanonicalDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectCanonicalDeploymentJSON) RawJSON() string {
	return r.raw
}

// Configs for deployments in a project.
type ProjectDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview DeploymentConfigsValues `json:"preview"`
	// Configs for production deploys.
	Production DeploymentConfigsValues      `json:"production"`
	JSON       projectDeploymentConfigsJSON `json:"-"`
}

// projectDeploymentConfigsJSON contains the JSON metadata for the struct
// [ProjectDeploymentConfigs]
type projectDeploymentConfigsJSON struct {
	Preview     apijson.Field
	Production  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsJSON) RawJSON() string {
	return r.raw
}

// Most recent deployment to the repo.
type ProjectLatestDeployment struct {
	JSON projectLatestDeploymentJSON `json:"-"`
	Deployments
}

// projectLatestDeploymentJSON contains the JSON metadata for the struct
// [ProjectLatestDeployment]
type projectLatestDeploymentJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectLatestDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectLatestDeploymentJSON) RawJSON() string {
	return r.raw
}

type ProjectParam struct {
	// Configs for the project build process.
	BuildConfig param.Field[BuildConfigParam] `json:"build_config"`
	// Most recent deployment to the repo.
	CanonicalDeployment param.Field[ProjectCanonicalDeploymentParam] `json:"canonical_deployment"`
	// Configs for deployments in a project.
	DeploymentConfigs param.Field[ProjectDeploymentConfigsParam] `json:"deployment_configs"`
	// Most recent deployment to the repo.
	LatestDeployment param.Field[ProjectLatestDeploymentParam] `json:"latest_deployment"`
	// Name of the project.
	Name param.Field[string] `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch param.Field[string] `json:"production_branch"`
}

func (r ProjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Most recent deployment to the repo.
type ProjectCanonicalDeploymentParam struct {
	DeploymentsParam
}

func (r ProjectCanonicalDeploymentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for deployments in a project.
type ProjectDeploymentConfigsParam struct {
	// Configs for preview deploys.
	Preview param.Field[DeploymentConfigsValuesParam] `json:"preview"`
	// Configs for production deploys.
	Production param.Field[DeploymentConfigsValuesParam] `json:"production"`
}

func (r ProjectDeploymentConfigsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Most recent deployment to the repo.
type ProjectLatestDeploymentParam struct {
	DeploymentsParam
}

func (r ProjectLatestDeploymentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectResponse struct {
	Result Project             `json:"result"`
	JSON   projectResponseJSON `json:"-"`
	APIResponsePages
}

// projectResponseJSON contains the JSON metadata for the struct [ProjectResponse]
type projectResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectResponseJSON) RawJSON() string {
	return r.raw
}

type Source struct {
	Config SourceConfig `json:"config"`
	Type   string       `json:"type"`
	JSON   sourceJSON   `json:"-"`
}

// sourceJSON contains the JSON metadata for the struct [Source]
type sourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Source) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sourceJSON) RawJSON() string {
	return r.raw
}

type SourceConfig struct {
	DeploymentsEnabled           bool                                 `json:"deployments_enabled"`
	Owner                        string                               `json:"owner"`
	PathExcludes                 []string                             `json:"path_excludes"`
	PathIncludes                 []string                             `json:"path_includes"`
	PrCommentsEnabled            bool                                 `json:"pr_comments_enabled"`
	PreviewBranchExcludes        []string                             `json:"preview_branch_excludes"`
	PreviewBranchIncludes        []string                             `json:"preview_branch_includes"`
	PreviewDeploymentSetting     SourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	ProductionBranch             string                               `json:"production_branch"`
	ProductionDeploymentsEnabled bool                                 `json:"production_deployments_enabled"`
	RepoName                     string                               `json:"repo_name"`
	JSON                         sourceConfigJSON                     `json:"-"`
}

// sourceConfigJSON contains the JSON metadata for the struct [SourceConfig]
type sourceConfigJSON struct {
	DeploymentsEnabled           apijson.Field
	Owner                        apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PrCommentsEnabled            apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionBranch             apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoName                     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sourceConfigJSON) RawJSON() string {
	return r.raw
}

type SourceConfigPreviewDeploymentSetting string

const (
	SourceConfigPreviewDeploymentSettingAll    SourceConfigPreviewDeploymentSetting = "all"
	SourceConfigPreviewDeploymentSettingNone   SourceConfigPreviewDeploymentSetting = "none"
	SourceConfigPreviewDeploymentSettingCustom SourceConfigPreviewDeploymentSetting = "custom"
)

func (r SourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case SourceConfigPreviewDeploymentSettingAll, SourceConfigPreviewDeploymentSettingNone, SourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

type SourceParam struct {
	Config param.Field[SourceConfigParam] `json:"config"`
	Type   param.Field[string]            `json:"type"`
}

func (r SourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SourceConfigParam struct {
	DeploymentsEnabled           param.Field[bool]                                 `json:"deployments_enabled"`
	Owner                        param.Field[string]                               `json:"owner"`
	PathExcludes                 param.Field[[]string]                             `json:"path_excludes"`
	PathIncludes                 param.Field[[]string]                             `json:"path_includes"`
	PrCommentsEnabled            param.Field[bool]                                 `json:"pr_comments_enabled"`
	PreviewBranchExcludes        param.Field[[]string]                             `json:"preview_branch_excludes"`
	PreviewBranchIncludes        param.Field[[]string]                             `json:"preview_branch_includes"`
	PreviewDeploymentSetting     param.Field[SourceConfigPreviewDeploymentSetting] `json:"preview_deployment_setting"`
	ProductionBranch             param.Field[string]                               `json:"production_branch"`
	ProductionDeploymentsEnabled param.Field[bool]                                 `json:"production_deployments_enabled"`
	RepoName                     param.Field[string]                               `json:"repo_name"`
}

func (r SourceConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the deployment.
type Stage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,nullable" format:"date-time"`
	// The current build stage.
	Name StageName `json:"name"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,nullable" format:"date-time"`
	// State of the current stage.
	Status StageStatus `json:"status"`
	JSON   stageJSON   `json:"-"`
}

// stageJSON contains the JSON metadata for the struct [Stage]
type stageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Stage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stageJSON) RawJSON() string {
	return r.raw
}

// The current build stage.
type StageName string

const (
	StageNameQueued     StageName = "queued"
	StageNameInitialize StageName = "initialize"
	StageNameCloneRepo  StageName = "clone_repo"
	StageNameBuild      StageName = "build"
	StageNameDeploy     StageName = "deploy"
)

func (r StageName) IsKnown() bool {
	switch r {
	case StageNameQueued, StageNameInitialize, StageNameCloneRepo, StageNameBuild, StageNameDeploy:
		return true
	}
	return false
}

// State of the current stage.
type StageStatus string

const (
	StageStatusSuccess  StageStatus = "success"
	StageStatusIdle     StageStatus = "idle"
	StageStatusActive   StageStatus = "active"
	StageStatusFailure  StageStatus = "failure"
	StageStatusCanceled StageStatus = "canceled"
)

func (r StageStatus) IsKnown() bool {
	switch r {
	case StageStatusSuccess, StageStatusIdle, StageStatusActive, StageStatusFailure, StageStatusCanceled:
		return true
	}
	return false
}

// The status of the deployment.
type StageParam struct {
	// The current build stage.
	Name param.Field[StageName] `json:"name"`
}

func (r StageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountPageProjectListResponse struct {
	Result []Deployments                      `json:"result"`
	JSON   accountPageProjectListResponseJSON `json:"-"`
	APIResponsePages
	APIResponsePagination
}

// accountPageProjectListResponseJSON contains the JSON metadata for the struct
// [AccountPageProjectListResponse]
type accountPageProjectListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDeleteResponse struct {
	Result interface{}                          `json:"result,nullable"`
	JSON   accountPageProjectDeleteResponseJSON `json:"-"`
	APIResponsePages
}

// accountPageProjectDeleteResponseJSON contains the JSON metadata for the struct
// [AccountPageProjectDeleteResponse]
type accountPageProjectDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectPurgeBuildCacheResponse struct {
	Result interface{}                                   `json:"result,nullable"`
	JSON   accountPageProjectPurgeBuildCacheResponseJSON `json:"-"`
	APIResponsePages
}

// accountPageProjectPurgeBuildCacheResponseJSON contains the JSON metadata for the
// struct [AccountPageProjectPurgeBuildCacheResponse]
type accountPageProjectPurgeBuildCacheResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectPurgeBuildCacheResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectPurgeBuildCacheResponseJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectNewParams struct {
	Project ProjectParam `json:"project,required"`
}

func (r AccountPageProjectNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Project)
}

type AccountPageProjectUpdateParams struct {
	Body AccountPageProjectUpdateParamsBody `json:"body,required"`
}

func (r AccountPageProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountPageProjectUpdateParamsBody struct {
	ProjectParam
}

func (r AccountPageProjectUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountPageProjectDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountPageProjectDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
