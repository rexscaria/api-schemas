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
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountWorkerDispatchNamespaceScriptService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerDispatchNamespaceScriptService] method instead.
type AccountWorkerDispatchNamespaceScriptService struct {
	Options  []option.RequestOption
	Content  *AccountWorkerDispatchNamespaceScriptContentService
	Secrets  *AccountWorkerDispatchNamespaceScriptSecretService
	Settings *AccountWorkerDispatchNamespaceScriptSettingService
	Tags     *AccountWorkerDispatchNamespaceScriptTagService
}

// NewAccountWorkerDispatchNamespaceScriptService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountWorkerDispatchNamespaceScriptService(opts ...option.RequestOption) (r *AccountWorkerDispatchNamespaceScriptService) {
	r = &AccountWorkerDispatchNamespaceScriptService{}
	r.Options = opts
	r.Content = NewAccountWorkerDispatchNamespaceScriptContentService(opts...)
	r.Secrets = NewAccountWorkerDispatchNamespaceScriptSecretService(opts...)
	r.Settings = NewAccountWorkerDispatchNamespaceScriptSettingService(opts...)
	r.Tags = NewAccountWorkerDispatchNamespaceScriptTagService(opts...)
	return
}

// Fetch information about a script uploaded to a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptService) Get(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptGetResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a worker from a Workers for Platforms namespace. This call has no
// response body on a successful delete.
func (r *AccountWorkerDispatchNamespaceScriptService) Delete(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, body AccountWorkerDispatchNamespaceScriptDeleteParams, opts ...option.RequestOption) (res *NullResult, err error) {
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Start uploading a collection of assets for use in a Worker version. To learn
// more about the direct uploads of assets, see
// https://developers.cloudflare.com/workers/static-assets/direct-upload/.
func (r *AccountWorkerDispatchNamespaceScriptService) NewAssetsUploadSession(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, body AccountWorkerDispatchNamespaceScriptNewAssetsUploadSessionParams, opts ...option.RequestOption) (res *UploadSessionResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/assets-upload-session", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch script bindings from a script uploaded to a Workers for Platforms
// namespace.
func (r *AccountWorkerDispatchNamespaceScriptService) GetBindings(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptGetBindingsResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/bindings", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a worker module to a Workers for Platforms namespace. You can find more
// about the multipart metadata on our docs:
// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.
func (r *AccountWorkerDispatchNamespaceScriptService) Upload(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, body AccountWorkerDispatchNamespaceScriptUploadParams, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptUploadResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// A binding to allow the Worker to communicate with resources.
type BindingItem struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemType `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID string `json:"id"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The name of the dataset to bind to.
	Dataset string `json:"dataset"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format BindingItemFormat `json:"format"`
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name"`
	// JSON data to use.
	Json string `json:"json"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// Namespace to bind to.
	Namespace string `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// This field can have the runtime type of
	// [BindingItemWorkersBindingKindDispatchNamespaceOutbound].
	Outbound interface{} `json:"outbound"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string `json:"script_name"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name"`
	// Name of Worker to bind to.
	Service string `json:"service"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id"`
	// The text value to use.
	Text string `json:"text"`
	// This field can have the runtime type of
	// [[]BindingItemWorkersBindingKindSecretKeyUsage].
	Usages interface{} `json:"usages"`
	// Name of the Workflow to bind to.
	WorkflowName string          `json:"workflow_name"`
	JSON         bindingItemJSON `json:"-"`
	union        BindingItemUnion
}

// bindingItemJSON contains the JSON metadata for the struct [BindingItem]
type bindingItemJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	ID            apijson.Field
	Algorithm     apijson.Field
	BucketName    apijson.Field
	CertificateID apijson.Field
	ClassName     apijson.Field
	Dataset       apijson.Field
	Environment   apijson.Field
	Format        apijson.Field
	IndexName     apijson.Field
	Json          apijson.Field
	KeyJwk        apijson.Field
	Namespace     apijson.Field
	NamespaceID   apijson.Field
	Outbound      apijson.Field
	Pipeline      apijson.Field
	QueueName     apijson.Field
	ScriptName    apijson.Field
	SecretName    apijson.Field
	Service       apijson.Field
	StoreID       apijson.Field
	Text          apijson.Field
	Usages        apijson.Field
	WorkflowName  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r bindingItemJSON) RawJSON() string {
	return r.raw
}

func (r *BindingItem) UnmarshalJSON(data []byte) (err error) {
	*r = BindingItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BindingItemUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [BindingItemWorkersBindingKindAI],
// [BindingItemWorkersBindingKindAnalyticsEngine],
// [BindingItemWorkersBindingKindAssets], [BindingItemWorkersBindingKindBrowser],
// [BindingItemWorkersBindingKindD1],
// [BindingItemWorkersBindingKindDispatchNamespace],
// [BindingItemWorkersBindingKindDurableObjectNamespace],
// [BindingItemWorkersBindingKindHyperdrive], [BindingItemWorkersBindingKindJson],
// [BindingItemWorkersBindingKindKvNamespace],
// [BindingItemWorkersBindingKindMtlsCertificate],
// [BindingItemWorkersBindingKindPlainText],
// [BindingItemWorkersBindingKindPipelines], [BindingItemWorkersBindingKindQueue],
// [BindingItemWorkersBindingKindR2Bucket],
// [BindingItemWorkersBindingKindSecretText],
// [BindingItemWorkersBindingKindService],
// [BindingItemWorkersBindingKindTailConsumer],
// [BindingItemWorkersBindingKindVectorize],
// [BindingItemWorkersBindingKindVersionMetadata],
// [BindingItemWorkersBindingKindSecretsStoreSecret],
// [BindingItemWorkersBindingKindSecretKey],
// [BindingItemWorkersBindingKindWorkflow].
func (r BindingItem) AsUnion() BindingItemUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources.
//
// Union satisfied by [BindingItemWorkersBindingKindAI],
// [BindingItemWorkersBindingKindAnalyticsEngine],
// [BindingItemWorkersBindingKindAssets], [BindingItemWorkersBindingKindBrowser],
// [BindingItemWorkersBindingKindD1],
// [BindingItemWorkersBindingKindDispatchNamespace],
// [BindingItemWorkersBindingKindDurableObjectNamespace],
// [BindingItemWorkersBindingKindHyperdrive], [BindingItemWorkersBindingKindJson],
// [BindingItemWorkersBindingKindKvNamespace],
// [BindingItemWorkersBindingKindMtlsCertificate],
// [BindingItemWorkersBindingKindPlainText],
// [BindingItemWorkersBindingKindPipelines], [BindingItemWorkersBindingKindQueue],
// [BindingItemWorkersBindingKindR2Bucket],
// [BindingItemWorkersBindingKindSecretText],
// [BindingItemWorkersBindingKindService],
// [BindingItemWorkersBindingKindTailConsumer],
// [BindingItemWorkersBindingKindVectorize],
// [BindingItemWorkersBindingKindVersionMetadata],
// [BindingItemWorkersBindingKindSecretsStoreSecret],
// [BindingItemWorkersBindingKindSecretKey] or
// [BindingItemWorkersBindingKindWorkflow].
type BindingItemUnion interface {
	implementsBindingItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BindingItemUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowser{}),
			DiscriminatorValue: "browser",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPipelines{}),
			DiscriminatorValue: "pipelines",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretsStoreSecret{}),
			DiscriminatorValue: "secrets_store_secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindWorkflow{}),
			DiscriminatorValue: "workflow",
		},
	)
}

type BindingItemWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindAIType `json:"type,required"`
	JSON bindingItemWorkersBindingKindAIJSON `json:"-"`
}

// bindingItemWorkersBindingKindAIJSON contains the JSON metadata for the struct
// [BindingItemWorkersBindingKindAI]
type bindingItemWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindAI) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindAIType string

const (
	BindingItemWorkersBindingKindAITypeAI BindingItemWorkersBindingKindAIType = "ai"
)

func (r BindingItemWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindAnalyticsEngineType `json:"type,required"`
	JSON bindingItemWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// bindingItemWorkersBindingKindAnalyticsEngineJSON contains the JSON metadata for
// the struct [BindingItemWorkersBindingKindAnalyticsEngine]
type bindingItemWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindAnalyticsEngine) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindAnalyticsEngineType string

const (
	BindingItemWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine BindingItemWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r BindingItemWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindAssetsType `json:"type,required"`
	JSON bindingItemWorkersBindingKindAssetsJSON `json:"-"`
}

// bindingItemWorkersBindingKindAssetsJSON contains the JSON metadata for the
// struct [BindingItemWorkersBindingKindAssets]
type bindingItemWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindAssets) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindAssetsType string

const (
	BindingItemWorkersBindingKindAssetsTypeAssets BindingItemWorkersBindingKindAssetsType = "assets"
)

func (r BindingItemWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindBrowserType `json:"type,required"`
	JSON bindingItemWorkersBindingKindBrowserJSON `json:"-"`
}

// bindingItemWorkersBindingKindBrowserJSON contains the JSON metadata for the
// struct [BindingItemWorkersBindingKindBrowser]
type bindingItemWorkersBindingKindBrowserJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindBrowserJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindBrowser) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindBrowserType string

const (
	BindingItemWorkersBindingKindBrowserTypeBrowser BindingItemWorkersBindingKindBrowserType = "browser"
)

func (r BindingItemWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindD1Type `json:"type,required"`
	JSON bindingItemWorkersBindingKindD1JSON `json:"-"`
}

// bindingItemWorkersBindingKindD1JSON contains the JSON metadata for the struct
// [BindingItemWorkersBindingKindD1]
type bindingItemWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindD1) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindD1Type string

const (
	BindingItemWorkersBindingKindD1TypeD1 BindingItemWorkersBindingKindD1Type = "d1"
)

func (r BindingItemWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindDispatchNamespaceType `json:"type,required"`
	// Outbound worker.
	Outbound BindingItemWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     bindingItemWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// bindingItemWorkersBindingKindDispatchNamespaceJSON contains the JSON metadata
// for the struct [BindingItemWorkersBindingKindDispatchNamespace]
type bindingItemWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindDispatchNamespace) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindDispatchNamespaceType string

const (
	BindingItemWorkersBindingKindDispatchNamespaceTypeDispatchNamespace BindingItemWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r BindingItemWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type BindingItemWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
	// Outbound worker.
	Worker BindingItemWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   bindingItemWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// bindingItemWorkersBindingKindDispatchNamespaceOutboundJSON contains the JSON
// metadata for the struct [BindingItemWorkersBindingKindDispatchNamespaceOutbound]
type bindingItemWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type BindingItemWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                           `json:"service"`
	JSON    bindingItemWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// bindingItemWorkersBindingKindDispatchNamespaceOutboundWorkerJSON contains the
// JSON metadata for the struct
// [BindingItemWorkersBindingKindDispatchNamespaceOutboundWorker]
type bindingItemWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type BindingItemWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                  `json:"script_name"`
	JSON       bindingItemWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// bindingItemWorkersBindingKindDurableObjectNamespaceJSON contains the JSON
// metadata for the struct [BindingItemWorkersBindingKindDurableObjectNamespace]
type bindingItemWorkersBindingKindDurableObjectNamespaceJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	ClassName   apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindDurableObjectNamespace) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindDurableObjectNamespaceType string

const (
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace BindingItemWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r BindingItemWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindHyperdriveType `json:"type,required"`
	JSON bindingItemWorkersBindingKindHyperdriveJSON `json:"-"`
}

// bindingItemWorkersBindingKindHyperdriveJSON contains the JSON metadata for the
// struct [BindingItemWorkersBindingKindHyperdrive]
type bindingItemWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindHyperdrive) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindHyperdriveType string

const (
	BindingItemWorkersBindingKindHyperdriveTypeHyperdrive BindingItemWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r BindingItemWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindJsonType `json:"type,required"`
	JSON bindingItemWorkersBindingKindJsonJSON `json:"-"`
}

// bindingItemWorkersBindingKindJsonJSON contains the JSON metadata for the struct
// [BindingItemWorkersBindingKindJson]
type bindingItemWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindJson) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindJsonType string

const (
	BindingItemWorkersBindingKindJsonTypeJson BindingItemWorkersBindingKindJsonType = "json"
)

func (r BindingItemWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindKvNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindKvNamespaceType `json:"type,required"`
	JSON bindingItemWorkersBindingKindKvNamespaceJSON `json:"-"`
}

// bindingItemWorkersBindingKindKvNamespaceJSON contains the JSON metadata for the
// struct [BindingItemWorkersBindingKindKvNamespace]
type bindingItemWorkersBindingKindKvNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindKvNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindKvNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindKvNamespace) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindKvNamespaceType string

const (
	BindingItemWorkersBindingKindKvNamespaceTypeKvNamespace BindingItemWorkersBindingKindKvNamespaceType = "kv_namespace"
)

func (r BindingItemWorkersBindingKindKvNamespaceType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindKvNamespaceTypeKvNamespace:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindMtlsCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindMtlsCertificateType `json:"type,required"`
	JSON bindingItemWorkersBindingKindMtlsCertificateJSON `json:"-"`
}

// bindingItemWorkersBindingKindMtlsCertificateJSON contains the JSON metadata for
// the struct [BindingItemWorkersBindingKindMtlsCertificate]
type bindingItemWorkersBindingKindMtlsCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindMtlsCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindMtlsCertificateJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindMtlsCertificate) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindMtlsCertificateType string

const (
	BindingItemWorkersBindingKindMtlsCertificateTypeMtlsCertificate BindingItemWorkersBindingKindMtlsCertificateType = "mtls_certificate"
)

func (r BindingItemWorkersBindingKindMtlsCertificateType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindMtlsCertificateTypeMtlsCertificate:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindPlainTextType `json:"type,required"`
	JSON bindingItemWorkersBindingKindPlainTextJSON `json:"-"`
}

// bindingItemWorkersBindingKindPlainTextJSON contains the JSON metadata for the
// struct [BindingItemWorkersBindingKindPlainText]
type bindingItemWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindPlainText) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindPlainTextType string

const (
	BindingItemWorkersBindingKindPlainTextTypePlainText BindingItemWorkersBindingKindPlainTextType = "plain_text"
)

func (r BindingItemWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindPipelinesType `json:"type,required"`
	JSON bindingItemWorkersBindingKindPipelinesJSON `json:"-"`
}

// bindingItemWorkersBindingKindPipelinesJSON contains the JSON metadata for the
// struct [BindingItemWorkersBindingKindPipelines]
type bindingItemWorkersBindingKindPipelinesJSON struct {
	Name        apijson.Field
	Pipeline    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindPipelines) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindPipelinesJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindPipelines) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindPipelinesType string

const (
	BindingItemWorkersBindingKindPipelinesTypePipelines BindingItemWorkersBindingKindPipelinesType = "pipelines"
)

func (r BindingItemWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindQueueType `json:"type,required"`
	JSON bindingItemWorkersBindingKindQueueJSON `json:"-"`
}

// bindingItemWorkersBindingKindQueueJSON contains the JSON metadata for the struct
// [BindingItemWorkersBindingKindQueue]
type bindingItemWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindQueue) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindQueueType string

const (
	BindingItemWorkersBindingKindQueueTypeQueue BindingItemWorkersBindingKindQueueType = "queue"
)

func (r BindingItemWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindR2BucketType `json:"type,required"`
	JSON bindingItemWorkersBindingKindR2BucketJSON `json:"-"`
}

// bindingItemWorkersBindingKindR2BucketJSON contains the JSON metadata for the
// struct [BindingItemWorkersBindingKindR2Bucket]
type bindingItemWorkersBindingKindR2BucketJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindR2Bucket) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindR2BucketType string

const (
	BindingItemWorkersBindingKindR2BucketTypeR2Bucket BindingItemWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r BindingItemWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindSecretTextType `json:"type,required"`
	JSON bindingItemWorkersBindingKindSecretTextJSON `json:"-"`
}

// bindingItemWorkersBindingKindSecretTextJSON contains the JSON metadata for the
// struct [BindingItemWorkersBindingKindSecretText]
type bindingItemWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindSecretText) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindSecretTextType string

const (
	BindingItemWorkersBindingKindSecretTextTypeSecretText BindingItemWorkersBindingKindSecretTextType = "secret_text"
)

func (r BindingItemWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindServiceType `json:"type,required"`
	JSON bindingItemWorkersBindingKindServiceJSON `json:"-"`
}

// bindingItemWorkersBindingKindServiceJSON contains the JSON metadata for the
// struct [BindingItemWorkersBindingKindService]
type bindingItemWorkersBindingKindServiceJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindService) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindServiceType string

const (
	BindingItemWorkersBindingKindServiceTypeService BindingItemWorkersBindingKindServiceType = "service"
)

func (r BindingItemWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindTailConsumerType `json:"type,required"`
	JSON bindingItemWorkersBindingKindTailConsumerJSON `json:"-"`
}

// bindingItemWorkersBindingKindTailConsumerJSON contains the JSON metadata for the
// struct [BindingItemWorkersBindingKindTailConsumer]
type bindingItemWorkersBindingKindTailConsumerJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindTailConsumerJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindTailConsumer) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindTailConsumerType string

const (
	BindingItemWorkersBindingKindTailConsumerTypeTailConsumer BindingItemWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r BindingItemWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindTailConsumerTypeTailConsumer:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindVectorizeType `json:"type,required"`
	JSON bindingItemWorkersBindingKindVectorizeJSON `json:"-"`
}

// bindingItemWorkersBindingKindVectorizeJSON contains the JSON metadata for the
// struct [BindingItemWorkersBindingKindVectorize]
type bindingItemWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindVectorize) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindVectorizeType string

const (
	BindingItemWorkersBindingKindVectorizeTypeVectorize BindingItemWorkersBindingKindVectorizeType = "vectorize"
)

func (r BindingItemWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindVersionMetadataType `json:"type,required"`
	JSON bindingItemWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// bindingItemWorkersBindingKindVersionMetadataJSON contains the JSON metadata for
// the struct [BindingItemWorkersBindingKindVersionMetadata]
type bindingItemWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindVersionMetadata) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindVersionMetadataType string

const (
	BindingItemWorkersBindingKindVersionMetadataTypeVersionMetadata BindingItemWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r BindingItemWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindSecretsStoreSecretType `json:"type,required"`
	JSON bindingItemWorkersBindingKindSecretsStoreSecretJSON `json:"-"`
}

// bindingItemWorkersBindingKindSecretsStoreSecretJSON contains the JSON metadata
// for the struct [BindingItemWorkersBindingKindSecretsStoreSecret]
type bindingItemWorkersBindingKindSecretsStoreSecretJSON struct {
	Name        apijson.Field
	SecretName  apijson.Field
	StoreID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindSecretsStoreSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindSecretsStoreSecretJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindSecretsStoreSecret) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindSecretsStoreSecretType string

const (
	BindingItemWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret BindingItemWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r BindingItemWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format BindingItemWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []BindingItemWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   bindingItemWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// bindingItemWorkersBindingKindSecretKeyJSON contains the JSON metadata for the
// struct [BindingItemWorkersBindingKindSecretKey]
type bindingItemWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindSecretKey) implementsBindingItem() {}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type BindingItemWorkersBindingKindSecretKeyFormat string

const (
	BindingItemWorkersBindingKindSecretKeyFormatRaw   BindingItemWorkersBindingKindSecretKeyFormat = "raw"
	BindingItemWorkersBindingKindSecretKeyFormatPkcs8 BindingItemWorkersBindingKindSecretKeyFormat = "pkcs8"
	BindingItemWorkersBindingKindSecretKeyFormatSpki  BindingItemWorkersBindingKindSecretKeyFormat = "spki"
	BindingItemWorkersBindingKindSecretKeyFormatJwk   BindingItemWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r BindingItemWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindSecretKeyFormatRaw, BindingItemWorkersBindingKindSecretKeyFormatPkcs8, BindingItemWorkersBindingKindSecretKeyFormatSpki, BindingItemWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindSecretKeyType string

const (
	BindingItemWorkersBindingKindSecretKeyTypeSecretKey BindingItemWorkersBindingKindSecretKeyType = "secret_key"
)

func (r BindingItemWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindSecretKeyUsage string

const (
	BindingItemWorkersBindingKindSecretKeyUsageEncrypt    BindingItemWorkersBindingKindSecretKeyUsage = "encrypt"
	BindingItemWorkersBindingKindSecretKeyUsageDecrypt    BindingItemWorkersBindingKindSecretKeyUsage = "decrypt"
	BindingItemWorkersBindingKindSecretKeyUsageSign       BindingItemWorkersBindingKindSecretKeyUsage = "sign"
	BindingItemWorkersBindingKindSecretKeyUsageVerify     BindingItemWorkersBindingKindSecretKeyUsage = "verify"
	BindingItemWorkersBindingKindSecretKeyUsageDeriveKey  BindingItemWorkersBindingKindSecretKeyUsage = "deriveKey"
	BindingItemWorkersBindingKindSecretKeyUsageDeriveBits BindingItemWorkersBindingKindSecretKeyUsage = "deriveBits"
	BindingItemWorkersBindingKindSecretKeyUsageWrapKey    BindingItemWorkersBindingKindSecretKeyUsage = "wrapKey"
	BindingItemWorkersBindingKindSecretKeyUsageUnwrapKey  BindingItemWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r BindingItemWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindSecretKeyUsageEncrypt, BindingItemWorkersBindingKindSecretKeyUsageDecrypt, BindingItemWorkersBindingKindSecretKeyUsageSign, BindingItemWorkersBindingKindSecretKeyUsageVerify, BindingItemWorkersBindingKindSecretKeyUsageDeriveKey, BindingItemWorkersBindingKindSecretKeyUsageDeriveBits, BindingItemWorkersBindingKindSecretKeyUsageWrapKey, BindingItemWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type BindingItemWorkersBindingKindWorkflowType `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName string `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName string `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName string                                    `json:"script_name"`
	JSON       bindingItemWorkersBindingKindWorkflowJSON `json:"-"`
}

// bindingItemWorkersBindingKindWorkflowJSON contains the JSON metadata for the
// struct [BindingItemWorkersBindingKindWorkflow]
type bindingItemWorkersBindingKindWorkflowJSON struct {
	Name         apijson.Field
	Type         apijson.Field
	WorkflowName apijson.Field
	ClassName    apijson.Field
	ScriptName   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindWorkflow) implementsBindingItem() {}

// The kind of resource that the binding provides.
type BindingItemWorkersBindingKindWorkflowType string

const (
	BindingItemWorkersBindingKindWorkflowTypeWorkflow BindingItemWorkersBindingKindWorkflowType = "workflow"
)

func (r BindingItemWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type BindingItemType string

const (
	BindingItemTypeAI                     BindingItemType = "ai"
	BindingItemTypeAnalyticsEngine        BindingItemType = "analytics_engine"
	BindingItemTypeAssets                 BindingItemType = "assets"
	BindingItemTypeBrowser                BindingItemType = "browser"
	BindingItemTypeD1                     BindingItemType = "d1"
	BindingItemTypeDispatchNamespace      BindingItemType = "dispatch_namespace"
	BindingItemTypeDurableObjectNamespace BindingItemType = "durable_object_namespace"
	BindingItemTypeHyperdrive             BindingItemType = "hyperdrive"
	BindingItemTypeJson                   BindingItemType = "json"
	BindingItemTypeKvNamespace            BindingItemType = "kv_namespace"
	BindingItemTypeMtlsCertificate        BindingItemType = "mtls_certificate"
	BindingItemTypePlainText              BindingItemType = "plain_text"
	BindingItemTypePipelines              BindingItemType = "pipelines"
	BindingItemTypeQueue                  BindingItemType = "queue"
	BindingItemTypeR2Bucket               BindingItemType = "r2_bucket"
	BindingItemTypeSecretText             BindingItemType = "secret_text"
	BindingItemTypeService                BindingItemType = "service"
	BindingItemTypeTailConsumer           BindingItemType = "tail_consumer"
	BindingItemTypeVectorize              BindingItemType = "vectorize"
	BindingItemTypeVersionMetadata        BindingItemType = "version_metadata"
	BindingItemTypeSecretsStoreSecret     BindingItemType = "secrets_store_secret"
	BindingItemTypeSecretKey              BindingItemType = "secret_key"
	BindingItemTypeWorkflow               BindingItemType = "workflow"
)

func (r BindingItemType) IsKnown() bool {
	switch r {
	case BindingItemTypeAI, BindingItemTypeAnalyticsEngine, BindingItemTypeAssets, BindingItemTypeBrowser, BindingItemTypeD1, BindingItemTypeDispatchNamespace, BindingItemTypeDurableObjectNamespace, BindingItemTypeHyperdrive, BindingItemTypeJson, BindingItemTypeKvNamespace, BindingItemTypeMtlsCertificate, BindingItemTypePlainText, BindingItemTypePipelines, BindingItemTypeQueue, BindingItemTypeR2Bucket, BindingItemTypeSecretText, BindingItemTypeService, BindingItemTypeTailConsumer, BindingItemTypeVectorize, BindingItemTypeVersionMetadata, BindingItemTypeSecretsStoreSecret, BindingItemTypeSecretKey, BindingItemTypeWorkflow:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type BindingItemFormat string

const (
	BindingItemFormatRaw   BindingItemFormat = "raw"
	BindingItemFormatPkcs8 BindingItemFormat = "pkcs8"
	BindingItemFormatSpki  BindingItemFormat = "spki"
	BindingItemFormatJwk   BindingItemFormat = "jwk"
)

func (r BindingItemFormat) IsKnown() bool {
	switch r {
	case BindingItemFormatRaw, BindingItemFormatPkcs8, BindingItemFormatSpki, BindingItemFormatJwk:
		return true
	}
	return false
}

// A binding to allow the Worker to communicate with resources.
type BindingItemParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemType] `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID        param.Field[string]      `json:"id"`
	Algorithm param.Field[interface{}] `json:"algorithm"`
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name"`
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset"`
	// The environment of the script_name to bind to.
	Environment param.Field[string] `json:"environment"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[BindingItemFormat] `json:"format"`
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name"`
	// JSON data to use.
	Json param.Field[string] `json:"json"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string]      `json:"key_base64"`
	KeyJwk    param.Field[interface{}] `json:"key_jwk"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID param.Field[string]      `json:"namespace_id"`
	Outbound    param.Field[interface{}] `json:"outbound"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName param.Field[string] `json:"script_name"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id"`
	// The text value to use.
	Text   param.Field[string]      `json:"text"`
	Usages param.Field[interface{}] `json:"usages"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name"`
}

func (r BindingItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemParam) implementsBindingItemUnionParam() {}

// A binding to allow the Worker to communicate with resources.
//
// Satisfied by [BindingItemWorkersBindingKindAIParam],
// [BindingItemWorkersBindingKindAnalyticsEngineParam],
// [BindingItemWorkersBindingKindAssetsParam],
// [BindingItemWorkersBindingKindBrowserParam],
// [BindingItemWorkersBindingKindD1Param],
// [BindingItemWorkersBindingKindDispatchNamespaceParam],
// [BindingItemWorkersBindingKindDurableObjectNamespaceParam],
// [BindingItemWorkersBindingKindHyperdriveParam],
// [BindingItemWorkersBindingKindJsonParam],
// [BindingItemWorkersBindingKindKvNamespaceParam],
// [BindingItemWorkersBindingKindMtlsCertificateParam],
// [BindingItemWorkersBindingKindPlainTextParam],
// [BindingItemWorkersBindingKindPipelinesParam],
// [BindingItemWorkersBindingKindQueueParam],
// [BindingItemWorkersBindingKindR2BucketParam],
// [BindingItemWorkersBindingKindSecretTextParam],
// [BindingItemWorkersBindingKindServiceParam],
// [BindingItemWorkersBindingKindTailConsumerParam],
// [BindingItemWorkersBindingKindVectorizeParam],
// [BindingItemWorkersBindingKindVersionMetadataParam],
// [BindingItemWorkersBindingKindSecretsStoreSecretParam],
// [BindingItemWorkersBindingKindSecretKeyParam],
// [BindingItemWorkersBindingKindWorkflowParam], [BindingItemParam].
type BindingItemUnionParam interface {
	implementsBindingItemUnionParam()
}

type BindingItemWorkersBindingKindAIParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindAIType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindAIParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindAIParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindAnalyticsEngineParam struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindAnalyticsEngineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindAnalyticsEngineParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindAssetsParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindAssetsType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindAssetsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindAssetsParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindBrowserParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindBrowserType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindBrowserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindBrowserParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindD1Param struct {
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindD1Type] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindD1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindD1Param) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindDispatchNamespaceParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindDispatchNamespaceType] `json:"type,required"`
	// Outbound worker.
	Outbound param.Field[BindingItemWorkersBindingKindDispatchNamespaceOutboundParam] `json:"outbound"`
}

func (r BindingItemWorkersBindingKindDispatchNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindDispatchNamespaceParam) implementsBindingItemUnionParam() {}

// Outbound worker.
type BindingItemWorkersBindingKindDispatchNamespaceOutboundParam struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params param.Field[[]string] `json:"params"`
	// Outbound worker.
	Worker param.Field[BindingItemWorkersBindingKindDispatchNamespaceOutboundWorkerParam] `json:"worker"`
}

func (r BindingItemWorkersBindingKindDispatchNamespaceOutboundParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type BindingItemWorkersBindingKindDispatchNamespaceOutboundWorkerParam struct {
	// Environment of the outbound worker.
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker.
	Service param.Field[string] `json:"service"`
}

func (r BindingItemWorkersBindingKindDispatchNamespaceOutboundWorkerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BindingItemWorkersBindingKindDurableObjectNamespaceParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindDurableObjectNamespaceType] `json:"type,required"`
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name"`
	// The environment of the script_name to bind to.
	Environment param.Field[string] `json:"environment"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r BindingItemWorkersBindingKindDurableObjectNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindDurableObjectNamespaceParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindHyperdriveParam struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindHyperdriveType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindHyperdriveParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindHyperdriveParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindJsonParam struct {
	// JSON data to use.
	Json param.Field[string] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindJsonType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindJsonParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindJsonParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindKvNamespaceParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindKvNamespaceType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindKvNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindKvNamespaceParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindMtlsCertificateParam struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindMtlsCertificateType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindMtlsCertificateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindMtlsCertificateParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindPlainTextParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The text value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindPlainTextType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindPlainTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindPlainTextParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindPipelinesParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindPipelinesType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindPipelinesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindPipelinesParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindQueueParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindQueueType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindQueueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindQueueParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindR2BucketParam struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindR2BucketType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindR2BucketParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindR2BucketParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindSecretTextParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindSecretTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindSecretTextParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindServiceParam struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindServiceType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindServiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindServiceParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindTailConsumerParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindTailConsumerType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindTailConsumerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindTailConsumerParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindVectorizeParam struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindVectorizeType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindVectorizeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindVectorizeParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindVersionMetadataParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindVersionMetadataType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindVersionMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindVersionMetadataParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindSecretsStoreSecretParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindSecretsStoreSecretType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindSecretsStoreSecretParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindSecretsStoreSecretParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindSecretKeyParam struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[BindingItemWorkersBindingKindSecretKeyFormat] `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindSecretKeyType] `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]BindingItemWorkersBindingKindSecretKeyUsage] `json:"usages,required"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string] `json:"key_base64"`
	// Key data in
	// [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key)
	// format. Required if `format` is "jwk".
	KeyJwk param.Field[interface{}] `json:"key_jwk"`
}

func (r BindingItemWorkersBindingKindSecretKeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindSecretKeyParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindWorkflowParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[BindingItemWorkersBindingKindWorkflowType] `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName param.Field[string] `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r BindingItemWorkersBindingKindWorkflowParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindWorkflowParam) implementsBindingItemUnionParam() {}

type UploadSessionObjectParam struct {
	// A manifest ([path]: {hash, size}) map of files to upload. As an example,
	// `/blog/hello-world.html` would be a valid path key.
	Manifest param.Field[map[string]UploadSessionObjectManifestParam] `json:"manifest,required"`
}

func (r UploadSessionObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UploadSessionObjectManifestParam struct {
	// The hash of the file.
	Hash param.Field[string] `json:"hash,required"`
	// The size of the file in bytes.
	Size param.Field[int64] `json:"size,required"`
}

func (r UploadSessionObjectManifestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UploadSessionResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success UploadSessionResponseSuccess `json:"success,required"`
	Result  UploadSessionResponseResult  `json:"result"`
	JSON    uploadSessionResponseJSON    `json:"-"`
}

// uploadSessionResponseJSON contains the JSON metadata for the struct
// [UploadSessionResponse]
type uploadSessionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UploadSessionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadSessionResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UploadSessionResponseSuccess bool

const (
	UploadSessionResponseSuccessTrue UploadSessionResponseSuccess = true
)

func (r UploadSessionResponseSuccess) IsKnown() bool {
	switch r {
	case UploadSessionResponseSuccessTrue:
		return true
	}
	return false
}

type UploadSessionResponseResult struct {
	// The requests to make to upload assets.
	Buckets [][]string `json:"buckets"`
	// A JWT to use as authentication for uploading assets.
	Jwt  string                          `json:"jwt"`
	JSON uploadSessionResponseResultJSON `json:"-"`
}

// uploadSessionResponseResultJSON contains the JSON metadata for the struct
// [UploadSessionResponseResult]
type uploadSessionResponseResultJSON struct {
	Buckets     apijson.Field
	Jwt         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UploadSessionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadSessionResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDispatchNamespaceScriptGetResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// Details about a worker uploaded to a Workers for Platforms namespace.
	Result AccountWorkerDispatchNamespaceScriptGetResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerDispatchNamespaceScriptGetResponseSuccess `json:"success,required"`
	JSON    accountWorkerDispatchNamespaceScriptGetResponseJSON    `json:"-"`
}

// accountWorkerDispatchNamespaceScriptGetResponseJSON contains the JSON metadata
// for the struct [AccountWorkerDispatchNamespaceScriptGetResponse]
type accountWorkerDispatchNamespaceScriptGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptGetResponseJSON) RawJSON() string {
	return r.raw
}

// Details about a worker uploaded to a Workers for Platforms namespace.
type AccountWorkerDispatchNamespaceScriptGetResponseResult struct {
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Name of the Workers for Platforms dispatch namespace.
	DispatchNamespace string `json:"dispatch_namespace"`
	// When the script was last modified.
	ModifiedOn time.Time                                                 `json:"modified_on" format:"date-time"`
	Script     ScriptResponse                                            `json:"script"`
	JSON       accountWorkerDispatchNamespaceScriptGetResponseResultJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptGetResponseResultJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptGetResponseResult]
type accountWorkerDispatchNamespaceScriptGetResponseResultJSON struct {
	CreatedOn         apijson.Field
	DispatchNamespace apijson.Field
	ModifiedOn        apijson.Field
	Script            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerDispatchNamespaceScriptGetResponseSuccess bool

const (
	AccountWorkerDispatchNamespaceScriptGetResponseSuccessTrue AccountWorkerDispatchNamespaceScriptGetResponseSuccess = true
)

func (r AccountWorkerDispatchNamespaceScriptGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerDispatchNamespaceScriptGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerDispatchNamespaceScriptGetBindingsResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Result []BindingItem `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerDispatchNamespaceScriptGetBindingsResponseSuccess `json:"success,required"`
	JSON    accountWorkerDispatchNamespaceScriptGetBindingsResponseJSON    `json:"-"`
}

// accountWorkerDispatchNamespaceScriptGetBindingsResponseJSON contains the JSON
// metadata for the struct
// [AccountWorkerDispatchNamespaceScriptGetBindingsResponse]
type accountWorkerDispatchNamespaceScriptGetBindingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptGetBindingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptGetBindingsResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerDispatchNamespaceScriptGetBindingsResponseSuccess bool

const (
	AccountWorkerDispatchNamespaceScriptGetBindingsResponseSuccessTrue AccountWorkerDispatchNamespaceScriptGetBindingsResponseSuccess = true
)

func (r AccountWorkerDispatchNamespaceScriptGetBindingsResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerDispatchNamespaceScriptGetBindingsResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerDispatchNamespaceScriptUploadResponse struct {
	Errors   []WorkersMessages                                        `json:"errors,required"`
	Messages []WorkersMessages                                        `json:"messages,required"`
	Result   AccountWorkerDispatchNamespaceScriptUploadResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerDispatchNamespaceScriptUploadResponseSuccess `json:"success,required"`
	JSON    accountWorkerDispatchNamespaceScriptUploadResponseJSON    `json:"-"`
}

// accountWorkerDispatchNamespaceScriptUploadResponseJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptUploadResponse]
type accountWorkerDispatchNamespaceScriptUploadResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptUploadResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDispatchNamespaceScriptUploadResponseResult struct {
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
	Placement AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacement `json:"placement"`
	// Deprecated: deprecated
	PlacementMode AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementMode `json:"placement_mode"`
	// Deprecated: deprecated
	PlacementStatus AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementStatus `json:"placement_status"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []TailConsumersScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel UsageModel                                                   `json:"usage_model"`
	JSON       accountWorkerDispatchNamespaceScriptUploadResponseResultJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptUploadResponseResultJSON contains the JSON
// metadata for the struct
// [AccountWorkerDispatchNamespaceScriptUploadResponseResult]
type accountWorkerDispatchNamespaceScriptUploadResponseResultJSON struct {
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

func (r *AccountWorkerDispatchNamespaceScriptUploadResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptUploadResponseResultJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacement struct {
	// The last time the script was analyzed for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	LastAnalyzedAt time.Time `json:"last_analyzed_at" format:"date-time"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode PlacementMode `json:"mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Status PlacementStatus                                                       `json:"status"`
	JSON   accountWorkerDispatchNamespaceScriptUploadResponseResultPlacementJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptUploadResponseResultPlacementJSON contains
// the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacement]
type accountWorkerDispatchNamespaceScriptUploadResponseResultPlacementJSON struct {
	LastAnalyzedAt apijson.Field
	Mode           apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptUploadResponseResultPlacementJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementMode string

const (
	AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementModeSmart AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementMode = "smart"
)

func (r AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementMode) IsKnown() bool {
	switch r {
	case AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementModeSmart:
		return true
	}
	return false
}

type AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementStatus string

const (
	AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementStatusSuccess                 AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementStatus = "SUCCESS"
	AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementStatusUnsupportedApplication  AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementStatus = "UNSUPPORTED_APPLICATION"
	AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementStatusInsufficientInvocations AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementStatus) IsKnown() bool {
	switch r {
	case AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementStatusSuccess, AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementStatusUnsupportedApplication, AccountWorkerDispatchNamespaceScriptUploadResponseResultPlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

// Whether the API call was successful.
type AccountWorkerDispatchNamespaceScriptUploadResponseSuccess bool

const (
	AccountWorkerDispatchNamespaceScriptUploadResponseSuccessTrue AccountWorkerDispatchNamespaceScriptUploadResponseSuccess = true
)

func (r AccountWorkerDispatchNamespaceScriptUploadResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerDispatchNamespaceScriptUploadResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerDispatchNamespaceScriptDeleteParams struct {
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
}

// URLQuery serializes [AccountWorkerDispatchNamespaceScriptDeleteParams]'s query
// parameters as `url.Values`.
func (r AccountWorkerDispatchNamespaceScriptDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountWorkerDispatchNamespaceScriptNewAssetsUploadSessionParams struct {
	UploadSessionObject UploadSessionObjectParam `json:"upload_session_object,required"`
}

func (r AccountWorkerDispatchNamespaceScriptNewAssetsUploadSessionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UploadSessionObject)
}

type AccountWorkerDispatchNamespaceScriptUploadParams struct {
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[AccountWorkerDispatchNamespaceScriptUploadParamsMetadata] `json:"metadata,required"`
	// An array of modules (often JavaScript files) comprising a Worker script. At
	// least one module must be present and referenced in the metadata as `main_module`
	// or `body_part` by filename.<br/>Possible Content-Type(s) are:
	// `application/javascript+module`, `text/javascript+module`,
	// `application/javascript`, `text/javascript`, `application/wasm`, `text/plain`,
	// `application/octet-stream`, `application/source-map`.
	Files param.Field[[]io.Reader] `json:"files" format:"binary"`
}

func (r AccountWorkerDispatchNamespaceScriptUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type AccountWorkerDispatchNamespaceScriptUploadParamsMetadata struct {
	// Configuration for assets within a Worker.
	Assets param.Field[AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssets] `json:"assets"`
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
	Migrations param.Field[AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsUnion] `json:"migrations"`
	// Observability settings for the Worker.
	Observability param.Field[ObservabilityParam] `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement param.Field[AccountWorkerDispatchNamespaceScriptUploadParamsMetadataPlacement] `json:"placement"`
	// List of strings to use as tags for this Worker.
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]TailConsumersScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[UsageModel] `json:"usage_model"`
}

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssets struct {
	// Configuration for assets within a Worker.
	Config param.Field[AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfig] `json:"config"`
	// Token provided upon successful upload of all files from a registered manifest.
	Jwt param.Field[string] `json:"jwt"`
}

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfig struct {
	// The contents of a \_headers file (used to attach custom headers on asset
	// responses).
	Headers param.Field[string] `json:"_headers"`
	// The contents of a \_redirects file (used to apply redirects or proxy paths ahead
	// of asset serving).
	Redirects param.Field[string] `json:"_redirects"`
	// Determines the redirects and rewrites of requests for HTML content.
	HTMLHandling param.Field[AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandling] `json:"html_handling"`
	// Determines the response when a request does not match a static asset, and there
	// is no Worker script.
	NotFoundHandling param.Field[AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigNotFoundHandling] `json:"not_found_handling"`
	// Contains a list path rules to control routing to either the Worker or assets.
	// Glob (\*) and negative (!) rules are supported. Rules must start with either '/'
	// or '!/'. At least one non-negative rule must be provided, and negative rules
	// have higher precedence than non-negative rules.
	RunWorkerFirst param.Field[AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigRunWorkerFirstUnion] `json:"run_worker_first"`
	// When true and the incoming request matches an asset, that will be served instead
	// of invoking the Worker script. When false, requests will always invoke the
	// Worker script.
	//
	// Deprecated: deprecated
	ServeDirectly param.Field[bool] `json:"serve_directly"`
}

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the redirects and rewrites of requests for HTML content.
type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandling string

const (
	AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash  AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandling = "auto-trailing-slash"
	AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandlingForceTrailingSlash AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandling = "force-trailing-slash"
	AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandlingDropTrailingSlash  AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandling = "drop-trailing-slash"
	AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandlingNone               AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandling = "none"
)

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandling) IsKnown() bool {
	switch r {
	case AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash, AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandlingForceTrailingSlash, AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandlingDropTrailingSlash, AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandlingNone:
		return true
	}
	return false
}

// Determines the response when a request does not match a static asset, and there
// is no Worker script.
type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigNotFoundHandling string

const (
	AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigNotFoundHandlingNone                  AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigNotFoundHandling = "none"
	AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigNotFoundHandling404Page               AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigNotFoundHandling = "404-page"
	AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigNotFoundHandlingSinglePageApplication AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigNotFoundHandling = "single-page-application"
)

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigNotFoundHandling) IsKnown() bool {
	switch r {
	case AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigNotFoundHandlingNone, AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigNotFoundHandling404Page, AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigNotFoundHandlingSinglePageApplication:
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
// [AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigRunWorkerFirstArray],
// [shared.UnionBool].
type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigRunWorkerFirstUnion interface {
	ImplementsAccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigRunWorkerFirstUnion()
}

type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigRunWorkerFirstArray []string

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigRunWorkerFirstArray) ImplementsAccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigRunWorkerFirstUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrations struct {
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

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrations) implementsAccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations],
// [AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersMultipleStepMigrations],
// [AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrations].
type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsUnion interface {
	implementsAccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsUnion()
}

// A single set of migrations to apply.
type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations) implementsAccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsUnion() {
}

type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]MigrationStepParam] `json:"steps"`
}

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersMultipleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersMultipleStepMigrations) implementsAccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsUnion() {
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[PlacementMode] `json:"mode"`
}

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
