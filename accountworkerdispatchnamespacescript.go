// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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
func (r *AccountWorkerDispatchNamespaceScriptService) Delete(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, params AccountWorkerDispatchNamespaceScriptDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Start uploading a collection of assets for use in a Worker version. To learn
// more about the direct uploads of assets, see
// https://developers.cloudflare.com/workers/static-assets/direct-upload/
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

// A binding to allow the Worker to communicate with resources
type BindingItem struct {
	// A JavaScript variable name for the binding.
	Name string          `json:"name,required"`
	Type BindingItemType `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID string `json:"id"`
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
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name"`
	// This field can have the runtime type of [interface{}].
	Json interface{} `json:"json"`
	// Namespace to bind to.
	Namespace string `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// This field can have the runtime type of
	// [BindingItemWorkersBindingKindDispatchNamespaceOutbound].
	Outbound interface{} `json:"outbound"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string `json:"script_name"`
	// Name of Worker to bind to.
	Service string `json:"service"`
	// The text value to use.
	Text  string          `json:"text"`
	JSON  bindingItemJSON `json:"-"`
	union BindingItemUnion
}

// bindingItemJSON contains the JSON metadata for the struct [BindingItem]
type bindingItemJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	ID            apijson.Field
	BucketName    apijson.Field
	CertificateID apijson.Field
	ClassName     apijson.Field
	Dataset       apijson.Field
	Environment   apijson.Field
	IndexName     apijson.Field
	Json          apijson.Field
	Namespace     apijson.Field
	NamespaceID   apijson.Field
	Outbound      apijson.Field
	QueueName     apijson.Field
	ScriptName    apijson.Field
	Service       apijson.Field
	Text          apijson.Field
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
// [BindingItemWorkersBindingKindAssets],
// [BindingItemWorkersBindingKindBrowserRendering],
// [BindingItemWorkersBindingKindD1],
// [BindingItemWorkersBindingKindDispatchNamespace],
// [BindingItemWorkersBindingKindDurableObjectNamespace],
// [BindingItemWorkersBindingKindHyperdrive], [BindingItemWorkersBindingKindJson],
// [BindingItemWorkersBindingKindKvNamespace],
// [BindingItemWorkersBindingKindMtlsCertificate],
// [BindingItemWorkersBindingKindPlainText], [BindingItemWorkersBindingKindQueue],
// [BindingItemWorkersBindingKindR2Bucket],
// [BindingItemWorkersBindingKindSecretText],
// [BindingItemWorkersBindingKindService],
// [BindingItemWorkersBindingKindTailConsumer],
// [BindingItemWorkersBindingKindVectorize],
// [BindingItemWorkersBindingKindVersionMetadata].
func (r BindingItem) AsUnion() BindingItemUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by [BindingItemWorkersBindingKindAI],
// [BindingItemWorkersBindingKindAnalyticsEngine],
// [BindingItemWorkersBindingKindAssets],
// [BindingItemWorkersBindingKindBrowserRendering],
// [BindingItemWorkersBindingKindD1],
// [BindingItemWorkersBindingKindDispatchNamespace],
// [BindingItemWorkersBindingKindDurableObjectNamespace],
// [BindingItemWorkersBindingKindHyperdrive], [BindingItemWorkersBindingKindJson],
// [BindingItemWorkersBindingKindKvNamespace],
// [BindingItemWorkersBindingKindMtlsCertificate],
// [BindingItemWorkersBindingKindPlainText], [BindingItemWorkersBindingKindQueue],
// [BindingItemWorkersBindingKindR2Bucket],
// [BindingItemWorkersBindingKindSecretText],
// [BindingItemWorkersBindingKindService],
// [BindingItemWorkersBindingKindTailConsumer],
// [BindingItemWorkersBindingKindVectorize] or
// [BindingItemWorkersBindingKindVersionMetadata].
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
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAI{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindAssets{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindD1{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindJson{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindKvNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindMtlsCertificate{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindPlainText{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindQueue{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindSecretText{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindService{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVectorize{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BindingItemWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "vectorize",
		},
	)
}

type BindingItemWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string                              `json:"name,required"`
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

type BindingItemWorkersBindingKindAIType string

const (
	BindingItemWorkersBindingKindAITypeAI                     BindingItemWorkersBindingKindAIType = "ai"
	BindingItemWorkersBindingKindAITypeAnalyticsEngine        BindingItemWorkersBindingKindAIType = "analytics_engine"
	BindingItemWorkersBindingKindAITypeAssets                 BindingItemWorkersBindingKindAIType = "assets"
	BindingItemWorkersBindingKindAITypeBrowserRendering       BindingItemWorkersBindingKindAIType = "browser_rendering"
	BindingItemWorkersBindingKindAITypeD1                     BindingItemWorkersBindingKindAIType = "d1"
	BindingItemWorkersBindingKindAITypeDispatchNamespace      BindingItemWorkersBindingKindAIType = "dispatch_namespace"
	BindingItemWorkersBindingKindAITypeDurableObjectNamespace BindingItemWorkersBindingKindAIType = "durable_object_namespace"
	BindingItemWorkersBindingKindAITypeHyperdrive             BindingItemWorkersBindingKindAIType = "hyperdrive"
	BindingItemWorkersBindingKindAITypeJson                   BindingItemWorkersBindingKindAIType = "json"
	BindingItemWorkersBindingKindAITypeKvNamespace            BindingItemWorkersBindingKindAIType = "kv_namespace"
	BindingItemWorkersBindingKindAITypeMtlsCertificate        BindingItemWorkersBindingKindAIType = "mtls_certificate"
	BindingItemWorkersBindingKindAITypePlainText              BindingItemWorkersBindingKindAIType = "plain_text"
	BindingItemWorkersBindingKindAITypeQueue                  BindingItemWorkersBindingKindAIType = "queue"
	BindingItemWorkersBindingKindAITypeR2Bucket               BindingItemWorkersBindingKindAIType = "r2_bucket"
	BindingItemWorkersBindingKindAITypeSecretText             BindingItemWorkersBindingKindAIType = "secret_text"
	BindingItemWorkersBindingKindAITypeService                BindingItemWorkersBindingKindAIType = "service"
	BindingItemWorkersBindingKindAITypeTailConsumer           BindingItemWorkersBindingKindAIType = "tail_consumer"
	BindingItemWorkersBindingKindAITypeVectorize              BindingItemWorkersBindingKindAIType = "vectorize"
	BindingItemWorkersBindingKindAITypeVersionMetadata        BindingItemWorkersBindingKindAIType = "version_metadata"
)

func (r BindingItemWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindAITypeAI, BindingItemWorkersBindingKindAITypeAnalyticsEngine, BindingItemWorkersBindingKindAITypeAssets, BindingItemWorkersBindingKindAITypeBrowserRendering, BindingItemWorkersBindingKindAITypeD1, BindingItemWorkersBindingKindAITypeDispatchNamespace, BindingItemWorkersBindingKindAITypeDurableObjectNamespace, BindingItemWorkersBindingKindAITypeHyperdrive, BindingItemWorkersBindingKindAITypeJson, BindingItemWorkersBindingKindAITypeKvNamespace, BindingItemWorkersBindingKindAITypeMtlsCertificate, BindingItemWorkersBindingKindAITypePlainText, BindingItemWorkersBindingKindAITypeQueue, BindingItemWorkersBindingKindAITypeR2Bucket, BindingItemWorkersBindingKindAITypeSecretText, BindingItemWorkersBindingKindAITypeService, BindingItemWorkersBindingKindAITypeTailConsumer, BindingItemWorkersBindingKindAITypeVectorize, BindingItemWorkersBindingKindAITypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string                                           `json:"name,required"`
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

type BindingItemWorkersBindingKindAnalyticsEngineType string

const (
	BindingItemWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine        BindingItemWorkersBindingKindAnalyticsEngineType = "analytics_engine"
	BindingItemWorkersBindingKindAnalyticsEngineTypeAI                     BindingItemWorkersBindingKindAnalyticsEngineType = "ai"
	BindingItemWorkersBindingKindAnalyticsEngineTypeAssets                 BindingItemWorkersBindingKindAnalyticsEngineType = "assets"
	BindingItemWorkersBindingKindAnalyticsEngineTypeBrowserRendering       BindingItemWorkersBindingKindAnalyticsEngineType = "browser_rendering"
	BindingItemWorkersBindingKindAnalyticsEngineTypeD1                     BindingItemWorkersBindingKindAnalyticsEngineType = "d1"
	BindingItemWorkersBindingKindAnalyticsEngineTypeDispatchNamespace      BindingItemWorkersBindingKindAnalyticsEngineType = "dispatch_namespace"
	BindingItemWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace BindingItemWorkersBindingKindAnalyticsEngineType = "durable_object_namespace"
	BindingItemWorkersBindingKindAnalyticsEngineTypeHyperdrive             BindingItemWorkersBindingKindAnalyticsEngineType = "hyperdrive"
	BindingItemWorkersBindingKindAnalyticsEngineTypeJson                   BindingItemWorkersBindingKindAnalyticsEngineType = "json"
	BindingItemWorkersBindingKindAnalyticsEngineTypeKvNamespace            BindingItemWorkersBindingKindAnalyticsEngineType = "kv_namespace"
	BindingItemWorkersBindingKindAnalyticsEngineTypeMtlsCertificate        BindingItemWorkersBindingKindAnalyticsEngineType = "mtls_certificate"
	BindingItemWorkersBindingKindAnalyticsEngineTypePlainText              BindingItemWorkersBindingKindAnalyticsEngineType = "plain_text"
	BindingItemWorkersBindingKindAnalyticsEngineTypeQueue                  BindingItemWorkersBindingKindAnalyticsEngineType = "queue"
	BindingItemWorkersBindingKindAnalyticsEngineTypeR2Bucket               BindingItemWorkersBindingKindAnalyticsEngineType = "r2_bucket"
	BindingItemWorkersBindingKindAnalyticsEngineTypeSecretText             BindingItemWorkersBindingKindAnalyticsEngineType = "secret_text"
	BindingItemWorkersBindingKindAnalyticsEngineTypeService                BindingItemWorkersBindingKindAnalyticsEngineType = "service"
	BindingItemWorkersBindingKindAnalyticsEngineTypeTailConsumer           BindingItemWorkersBindingKindAnalyticsEngineType = "tail_consumer"
	BindingItemWorkersBindingKindAnalyticsEngineTypeVectorize              BindingItemWorkersBindingKindAnalyticsEngineType = "vectorize"
	BindingItemWorkersBindingKindAnalyticsEngineTypeVersionMetadata        BindingItemWorkersBindingKindAnalyticsEngineType = "version_metadata"
)

func (r BindingItemWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine, BindingItemWorkersBindingKindAnalyticsEngineTypeAI, BindingItemWorkersBindingKindAnalyticsEngineTypeAssets, BindingItemWorkersBindingKindAnalyticsEngineTypeBrowserRendering, BindingItemWorkersBindingKindAnalyticsEngineTypeD1, BindingItemWorkersBindingKindAnalyticsEngineTypeDispatchNamespace, BindingItemWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace, BindingItemWorkersBindingKindAnalyticsEngineTypeHyperdrive, BindingItemWorkersBindingKindAnalyticsEngineTypeJson, BindingItemWorkersBindingKindAnalyticsEngineTypeKvNamespace, BindingItemWorkersBindingKindAnalyticsEngineTypeMtlsCertificate, BindingItemWorkersBindingKindAnalyticsEngineTypePlainText, BindingItemWorkersBindingKindAnalyticsEngineTypeQueue, BindingItemWorkersBindingKindAnalyticsEngineTypeR2Bucket, BindingItemWorkersBindingKindAnalyticsEngineTypeSecretText, BindingItemWorkersBindingKindAnalyticsEngineTypeService, BindingItemWorkersBindingKindAnalyticsEngineTypeTailConsumer, BindingItemWorkersBindingKindAnalyticsEngineTypeVectorize, BindingItemWorkersBindingKindAnalyticsEngineTypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string                                  `json:"name,required"`
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

type BindingItemWorkersBindingKindAssetsType string

const (
	BindingItemWorkersBindingKindAssetsTypeAssets                 BindingItemWorkersBindingKindAssetsType = "assets"
	BindingItemWorkersBindingKindAssetsTypeAI                     BindingItemWorkersBindingKindAssetsType = "ai"
	BindingItemWorkersBindingKindAssetsTypeAnalyticsEngine        BindingItemWorkersBindingKindAssetsType = "analytics_engine"
	BindingItemWorkersBindingKindAssetsTypeBrowserRendering       BindingItemWorkersBindingKindAssetsType = "browser_rendering"
	BindingItemWorkersBindingKindAssetsTypeD1                     BindingItemWorkersBindingKindAssetsType = "d1"
	BindingItemWorkersBindingKindAssetsTypeDispatchNamespace      BindingItemWorkersBindingKindAssetsType = "dispatch_namespace"
	BindingItemWorkersBindingKindAssetsTypeDurableObjectNamespace BindingItemWorkersBindingKindAssetsType = "durable_object_namespace"
	BindingItemWorkersBindingKindAssetsTypeHyperdrive             BindingItemWorkersBindingKindAssetsType = "hyperdrive"
	BindingItemWorkersBindingKindAssetsTypeJson                   BindingItemWorkersBindingKindAssetsType = "json"
	BindingItemWorkersBindingKindAssetsTypeKvNamespace            BindingItemWorkersBindingKindAssetsType = "kv_namespace"
	BindingItemWorkersBindingKindAssetsTypeMtlsCertificate        BindingItemWorkersBindingKindAssetsType = "mtls_certificate"
	BindingItemWorkersBindingKindAssetsTypePlainText              BindingItemWorkersBindingKindAssetsType = "plain_text"
	BindingItemWorkersBindingKindAssetsTypeQueue                  BindingItemWorkersBindingKindAssetsType = "queue"
	BindingItemWorkersBindingKindAssetsTypeR2Bucket               BindingItemWorkersBindingKindAssetsType = "r2_bucket"
	BindingItemWorkersBindingKindAssetsTypeSecretText             BindingItemWorkersBindingKindAssetsType = "secret_text"
	BindingItemWorkersBindingKindAssetsTypeService                BindingItemWorkersBindingKindAssetsType = "service"
	BindingItemWorkersBindingKindAssetsTypeTailConsumer           BindingItemWorkersBindingKindAssetsType = "tail_consumer"
	BindingItemWorkersBindingKindAssetsTypeVectorize              BindingItemWorkersBindingKindAssetsType = "vectorize"
	BindingItemWorkersBindingKindAssetsTypeVersionMetadata        BindingItemWorkersBindingKindAssetsType = "version_metadata"
)

func (r BindingItemWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindAssetsTypeAssets, BindingItemWorkersBindingKindAssetsTypeAI, BindingItemWorkersBindingKindAssetsTypeAnalyticsEngine, BindingItemWorkersBindingKindAssetsTypeBrowserRendering, BindingItemWorkersBindingKindAssetsTypeD1, BindingItemWorkersBindingKindAssetsTypeDispatchNamespace, BindingItemWorkersBindingKindAssetsTypeDurableObjectNamespace, BindingItemWorkersBindingKindAssetsTypeHyperdrive, BindingItemWorkersBindingKindAssetsTypeJson, BindingItemWorkersBindingKindAssetsTypeKvNamespace, BindingItemWorkersBindingKindAssetsTypeMtlsCertificate, BindingItemWorkersBindingKindAssetsTypePlainText, BindingItemWorkersBindingKindAssetsTypeQueue, BindingItemWorkersBindingKindAssetsTypeR2Bucket, BindingItemWorkersBindingKindAssetsTypeSecretText, BindingItemWorkersBindingKindAssetsTypeService, BindingItemWorkersBindingKindAssetsTypeTailConsumer, BindingItemWorkersBindingKindAssetsTypeVectorize, BindingItemWorkersBindingKindAssetsTypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindBrowserRendering struct {
	// A JavaScript variable name for the binding.
	Name string                                            `json:"name,required"`
	Type BindingItemWorkersBindingKindBrowserRenderingType `json:"type,required"`
	JSON bindingItemWorkersBindingKindBrowserRenderingJSON `json:"-"`
}

// bindingItemWorkersBindingKindBrowserRenderingJSON contains the JSON metadata for
// the struct [BindingItemWorkersBindingKindBrowserRendering]
type bindingItemWorkersBindingKindBrowserRenderingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersBindingKindBrowserRendering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersBindingKindBrowserRenderingJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersBindingKindBrowserRendering) implementsBindingItem() {}

type BindingItemWorkersBindingKindBrowserRenderingType string

const (
	BindingItemWorkersBindingKindBrowserRenderingTypeBrowserRendering       BindingItemWorkersBindingKindBrowserRenderingType = "browser_rendering"
	BindingItemWorkersBindingKindBrowserRenderingTypeAI                     BindingItemWorkersBindingKindBrowserRenderingType = "ai"
	BindingItemWorkersBindingKindBrowserRenderingTypeAnalyticsEngine        BindingItemWorkersBindingKindBrowserRenderingType = "analytics_engine"
	BindingItemWorkersBindingKindBrowserRenderingTypeAssets                 BindingItemWorkersBindingKindBrowserRenderingType = "assets"
	BindingItemWorkersBindingKindBrowserRenderingTypeD1                     BindingItemWorkersBindingKindBrowserRenderingType = "d1"
	BindingItemWorkersBindingKindBrowserRenderingTypeDispatchNamespace      BindingItemWorkersBindingKindBrowserRenderingType = "dispatch_namespace"
	BindingItemWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace BindingItemWorkersBindingKindBrowserRenderingType = "durable_object_namespace"
	BindingItemWorkersBindingKindBrowserRenderingTypeHyperdrive             BindingItemWorkersBindingKindBrowserRenderingType = "hyperdrive"
	BindingItemWorkersBindingKindBrowserRenderingTypeJson                   BindingItemWorkersBindingKindBrowserRenderingType = "json"
	BindingItemWorkersBindingKindBrowserRenderingTypeKvNamespace            BindingItemWorkersBindingKindBrowserRenderingType = "kv_namespace"
	BindingItemWorkersBindingKindBrowserRenderingTypeMtlsCertificate        BindingItemWorkersBindingKindBrowserRenderingType = "mtls_certificate"
	BindingItemWorkersBindingKindBrowserRenderingTypePlainText              BindingItemWorkersBindingKindBrowserRenderingType = "plain_text"
	BindingItemWorkersBindingKindBrowserRenderingTypeQueue                  BindingItemWorkersBindingKindBrowserRenderingType = "queue"
	BindingItemWorkersBindingKindBrowserRenderingTypeR2Bucket               BindingItemWorkersBindingKindBrowserRenderingType = "r2_bucket"
	BindingItemWorkersBindingKindBrowserRenderingTypeSecretText             BindingItemWorkersBindingKindBrowserRenderingType = "secret_text"
	BindingItemWorkersBindingKindBrowserRenderingTypeService                BindingItemWorkersBindingKindBrowserRenderingType = "service"
	BindingItemWorkersBindingKindBrowserRenderingTypeTailConsumer           BindingItemWorkersBindingKindBrowserRenderingType = "tail_consumer"
	BindingItemWorkersBindingKindBrowserRenderingTypeVectorize              BindingItemWorkersBindingKindBrowserRenderingType = "vectorize"
	BindingItemWorkersBindingKindBrowserRenderingTypeVersionMetadata        BindingItemWorkersBindingKindBrowserRenderingType = "version_metadata"
)

func (r BindingItemWorkersBindingKindBrowserRenderingType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindBrowserRenderingTypeBrowserRendering, BindingItemWorkersBindingKindBrowserRenderingTypeAI, BindingItemWorkersBindingKindBrowserRenderingTypeAnalyticsEngine, BindingItemWorkersBindingKindBrowserRenderingTypeAssets, BindingItemWorkersBindingKindBrowserRenderingTypeD1, BindingItemWorkersBindingKindBrowserRenderingTypeDispatchNamespace, BindingItemWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace, BindingItemWorkersBindingKindBrowserRenderingTypeHyperdrive, BindingItemWorkersBindingKindBrowserRenderingTypeJson, BindingItemWorkersBindingKindBrowserRenderingTypeKvNamespace, BindingItemWorkersBindingKindBrowserRenderingTypeMtlsCertificate, BindingItemWorkersBindingKindBrowserRenderingTypePlainText, BindingItemWorkersBindingKindBrowserRenderingTypeQueue, BindingItemWorkersBindingKindBrowserRenderingTypeR2Bucket, BindingItemWorkersBindingKindBrowserRenderingTypeSecretText, BindingItemWorkersBindingKindBrowserRenderingTypeService, BindingItemWorkersBindingKindBrowserRenderingTypeTailConsumer, BindingItemWorkersBindingKindBrowserRenderingTypeVectorize, BindingItemWorkersBindingKindBrowserRenderingTypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string                              `json:"name,required"`
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

type BindingItemWorkersBindingKindD1Type string

const (
	BindingItemWorkersBindingKindD1TypeD1                     BindingItemWorkersBindingKindD1Type = "d1"
	BindingItemWorkersBindingKindD1TypeAI                     BindingItemWorkersBindingKindD1Type = "ai"
	BindingItemWorkersBindingKindD1TypeAnalyticsEngine        BindingItemWorkersBindingKindD1Type = "analytics_engine"
	BindingItemWorkersBindingKindD1TypeAssets                 BindingItemWorkersBindingKindD1Type = "assets"
	BindingItemWorkersBindingKindD1TypeBrowserRendering       BindingItemWorkersBindingKindD1Type = "browser_rendering"
	BindingItemWorkersBindingKindD1TypeDispatchNamespace      BindingItemWorkersBindingKindD1Type = "dispatch_namespace"
	BindingItemWorkersBindingKindD1TypeDurableObjectNamespace BindingItemWorkersBindingKindD1Type = "durable_object_namespace"
	BindingItemWorkersBindingKindD1TypeHyperdrive             BindingItemWorkersBindingKindD1Type = "hyperdrive"
	BindingItemWorkersBindingKindD1TypeJson                   BindingItemWorkersBindingKindD1Type = "json"
	BindingItemWorkersBindingKindD1TypeKvNamespace            BindingItemWorkersBindingKindD1Type = "kv_namespace"
	BindingItemWorkersBindingKindD1TypeMtlsCertificate        BindingItemWorkersBindingKindD1Type = "mtls_certificate"
	BindingItemWorkersBindingKindD1TypePlainText              BindingItemWorkersBindingKindD1Type = "plain_text"
	BindingItemWorkersBindingKindD1TypeQueue                  BindingItemWorkersBindingKindD1Type = "queue"
	BindingItemWorkersBindingKindD1TypeR2Bucket               BindingItemWorkersBindingKindD1Type = "r2_bucket"
	BindingItemWorkersBindingKindD1TypeSecretText             BindingItemWorkersBindingKindD1Type = "secret_text"
	BindingItemWorkersBindingKindD1TypeService                BindingItemWorkersBindingKindD1Type = "service"
	BindingItemWorkersBindingKindD1TypeTailConsumer           BindingItemWorkersBindingKindD1Type = "tail_consumer"
	BindingItemWorkersBindingKindD1TypeVectorize              BindingItemWorkersBindingKindD1Type = "vectorize"
	BindingItemWorkersBindingKindD1TypeVersionMetadata        BindingItemWorkersBindingKindD1Type = "version_metadata"
)

func (r BindingItemWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindD1TypeD1, BindingItemWorkersBindingKindD1TypeAI, BindingItemWorkersBindingKindD1TypeAnalyticsEngine, BindingItemWorkersBindingKindD1TypeAssets, BindingItemWorkersBindingKindD1TypeBrowserRendering, BindingItemWorkersBindingKindD1TypeDispatchNamespace, BindingItemWorkersBindingKindD1TypeDurableObjectNamespace, BindingItemWorkersBindingKindD1TypeHyperdrive, BindingItemWorkersBindingKindD1TypeJson, BindingItemWorkersBindingKindD1TypeKvNamespace, BindingItemWorkersBindingKindD1TypeMtlsCertificate, BindingItemWorkersBindingKindD1TypePlainText, BindingItemWorkersBindingKindD1TypeQueue, BindingItemWorkersBindingKindD1TypeR2Bucket, BindingItemWorkersBindingKindD1TypeSecretText, BindingItemWorkersBindingKindD1TypeService, BindingItemWorkersBindingKindD1TypeTailConsumer, BindingItemWorkersBindingKindD1TypeVectorize, BindingItemWorkersBindingKindD1TypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string                                             `json:"namespace,required"`
	Type      BindingItemWorkersBindingKindDispatchNamespaceType `json:"type,required"`
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

type BindingItemWorkersBindingKindDispatchNamespaceType string

const (
	BindingItemWorkersBindingKindDispatchNamespaceTypeDispatchNamespace      BindingItemWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
	BindingItemWorkersBindingKindDispatchNamespaceTypeAI                     BindingItemWorkersBindingKindDispatchNamespaceType = "ai"
	BindingItemWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine        BindingItemWorkersBindingKindDispatchNamespaceType = "analytics_engine"
	BindingItemWorkersBindingKindDispatchNamespaceTypeAssets                 BindingItemWorkersBindingKindDispatchNamespaceType = "assets"
	BindingItemWorkersBindingKindDispatchNamespaceTypeBrowserRendering       BindingItemWorkersBindingKindDispatchNamespaceType = "browser_rendering"
	BindingItemWorkersBindingKindDispatchNamespaceTypeD1                     BindingItemWorkersBindingKindDispatchNamespaceType = "d1"
	BindingItemWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace BindingItemWorkersBindingKindDispatchNamespaceType = "durable_object_namespace"
	BindingItemWorkersBindingKindDispatchNamespaceTypeHyperdrive             BindingItemWorkersBindingKindDispatchNamespaceType = "hyperdrive"
	BindingItemWorkersBindingKindDispatchNamespaceTypeJson                   BindingItemWorkersBindingKindDispatchNamespaceType = "json"
	BindingItemWorkersBindingKindDispatchNamespaceTypeKvNamespace            BindingItemWorkersBindingKindDispatchNamespaceType = "kv_namespace"
	BindingItemWorkersBindingKindDispatchNamespaceTypeMtlsCertificate        BindingItemWorkersBindingKindDispatchNamespaceType = "mtls_certificate"
	BindingItemWorkersBindingKindDispatchNamespaceTypePlainText              BindingItemWorkersBindingKindDispatchNamespaceType = "plain_text"
	BindingItemWorkersBindingKindDispatchNamespaceTypeQueue                  BindingItemWorkersBindingKindDispatchNamespaceType = "queue"
	BindingItemWorkersBindingKindDispatchNamespaceTypeR2Bucket               BindingItemWorkersBindingKindDispatchNamespaceType = "r2_bucket"
	BindingItemWorkersBindingKindDispatchNamespaceTypeSecretText             BindingItemWorkersBindingKindDispatchNamespaceType = "secret_text"
	BindingItemWorkersBindingKindDispatchNamespaceTypeService                BindingItemWorkersBindingKindDispatchNamespaceType = "service"
	BindingItemWorkersBindingKindDispatchNamespaceTypeTailConsumer           BindingItemWorkersBindingKindDispatchNamespaceType = "tail_consumer"
	BindingItemWorkersBindingKindDispatchNamespaceTypeVectorize              BindingItemWorkersBindingKindDispatchNamespaceType = "vectorize"
	BindingItemWorkersBindingKindDispatchNamespaceTypeVersionMetadata        BindingItemWorkersBindingKindDispatchNamespaceType = "version_metadata"
)

func (r BindingItemWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindDispatchNamespaceTypeDispatchNamespace, BindingItemWorkersBindingKindDispatchNamespaceTypeAI, BindingItemWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine, BindingItemWorkersBindingKindDispatchNamespaceTypeAssets, BindingItemWorkersBindingKindDispatchNamespaceTypeBrowserRendering, BindingItemWorkersBindingKindDispatchNamespaceTypeD1, BindingItemWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace, BindingItemWorkersBindingKindDispatchNamespaceTypeHyperdrive, BindingItemWorkersBindingKindDispatchNamespaceTypeJson, BindingItemWorkersBindingKindDispatchNamespaceTypeKvNamespace, BindingItemWorkersBindingKindDispatchNamespaceTypeMtlsCertificate, BindingItemWorkersBindingKindDispatchNamespaceTypePlainText, BindingItemWorkersBindingKindDispatchNamespaceTypeQueue, BindingItemWorkersBindingKindDispatchNamespaceTypeR2Bucket, BindingItemWorkersBindingKindDispatchNamespaceTypeSecretText, BindingItemWorkersBindingKindDispatchNamespaceTypeService, BindingItemWorkersBindingKindDispatchNamespaceTypeTailConsumer, BindingItemWorkersBindingKindDispatchNamespaceTypeVectorize, BindingItemWorkersBindingKindDispatchNamespaceTypeVersionMetadata:
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
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string                                                  `json:"name,required"`
	Type BindingItemWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
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
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
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

type BindingItemWorkersBindingKindDurableObjectNamespaceType string

const (
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace BindingItemWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeAI                     BindingItemWorkersBindingKindDurableObjectNamespaceType = "ai"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine        BindingItemWorkersBindingKindDurableObjectNamespaceType = "analytics_engine"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeAssets                 BindingItemWorkersBindingKindDurableObjectNamespaceType = "assets"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering       BindingItemWorkersBindingKindDurableObjectNamespaceType = "browser_rendering"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeD1                     BindingItemWorkersBindingKindDurableObjectNamespaceType = "d1"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace      BindingItemWorkersBindingKindDurableObjectNamespaceType = "dispatch_namespace"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeHyperdrive             BindingItemWorkersBindingKindDurableObjectNamespaceType = "hyperdrive"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeJson                   BindingItemWorkersBindingKindDurableObjectNamespaceType = "json"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeKvNamespace            BindingItemWorkersBindingKindDurableObjectNamespaceType = "kv_namespace"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeMtlsCertificate        BindingItemWorkersBindingKindDurableObjectNamespaceType = "mtls_certificate"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypePlainText              BindingItemWorkersBindingKindDurableObjectNamespaceType = "plain_text"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeQueue                  BindingItemWorkersBindingKindDurableObjectNamespaceType = "queue"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeR2Bucket               BindingItemWorkersBindingKindDurableObjectNamespaceType = "r2_bucket"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeSecretText             BindingItemWorkersBindingKindDurableObjectNamespaceType = "secret_text"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeService                BindingItemWorkersBindingKindDurableObjectNamespaceType = "service"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeTailConsumer           BindingItemWorkersBindingKindDurableObjectNamespaceType = "tail_consumer"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeVectorize              BindingItemWorkersBindingKindDurableObjectNamespaceType = "vectorize"
	BindingItemWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata        BindingItemWorkersBindingKindDurableObjectNamespaceType = "version_metadata"
)

func (r BindingItemWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace, BindingItemWorkersBindingKindDurableObjectNamespaceTypeAI, BindingItemWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine, BindingItemWorkersBindingKindDurableObjectNamespaceTypeAssets, BindingItemWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering, BindingItemWorkersBindingKindDurableObjectNamespaceTypeD1, BindingItemWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace, BindingItemWorkersBindingKindDurableObjectNamespaceTypeHyperdrive, BindingItemWorkersBindingKindDurableObjectNamespaceTypeJson, BindingItemWorkersBindingKindDurableObjectNamespaceTypeKvNamespace, BindingItemWorkersBindingKindDurableObjectNamespaceTypeMtlsCertificate, BindingItemWorkersBindingKindDurableObjectNamespaceTypePlainText, BindingItemWorkersBindingKindDurableObjectNamespaceTypeQueue, BindingItemWorkersBindingKindDurableObjectNamespaceTypeR2Bucket, BindingItemWorkersBindingKindDurableObjectNamespaceTypeSecretText, BindingItemWorkersBindingKindDurableObjectNamespaceTypeService, BindingItemWorkersBindingKindDurableObjectNamespaceTypeTailConsumer, BindingItemWorkersBindingKindDurableObjectNamespaceTypeVectorize, BindingItemWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string                                      `json:"name,required"`
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

type BindingItemWorkersBindingKindHyperdriveType string

const (
	BindingItemWorkersBindingKindHyperdriveTypeHyperdrive             BindingItemWorkersBindingKindHyperdriveType = "hyperdrive"
	BindingItemWorkersBindingKindHyperdriveTypeAI                     BindingItemWorkersBindingKindHyperdriveType = "ai"
	BindingItemWorkersBindingKindHyperdriveTypeAnalyticsEngine        BindingItemWorkersBindingKindHyperdriveType = "analytics_engine"
	BindingItemWorkersBindingKindHyperdriveTypeAssets                 BindingItemWorkersBindingKindHyperdriveType = "assets"
	BindingItemWorkersBindingKindHyperdriveTypeBrowserRendering       BindingItemWorkersBindingKindHyperdriveType = "browser_rendering"
	BindingItemWorkersBindingKindHyperdriveTypeD1                     BindingItemWorkersBindingKindHyperdriveType = "d1"
	BindingItemWorkersBindingKindHyperdriveTypeDispatchNamespace      BindingItemWorkersBindingKindHyperdriveType = "dispatch_namespace"
	BindingItemWorkersBindingKindHyperdriveTypeDurableObjectNamespace BindingItemWorkersBindingKindHyperdriveType = "durable_object_namespace"
	BindingItemWorkersBindingKindHyperdriveTypeJson                   BindingItemWorkersBindingKindHyperdriveType = "json"
	BindingItemWorkersBindingKindHyperdriveTypeKvNamespace            BindingItemWorkersBindingKindHyperdriveType = "kv_namespace"
	BindingItemWorkersBindingKindHyperdriveTypeMtlsCertificate        BindingItemWorkersBindingKindHyperdriveType = "mtls_certificate"
	BindingItemWorkersBindingKindHyperdriveTypePlainText              BindingItemWorkersBindingKindHyperdriveType = "plain_text"
	BindingItemWorkersBindingKindHyperdriveTypeQueue                  BindingItemWorkersBindingKindHyperdriveType = "queue"
	BindingItemWorkersBindingKindHyperdriveTypeR2Bucket               BindingItemWorkersBindingKindHyperdriveType = "r2_bucket"
	BindingItemWorkersBindingKindHyperdriveTypeSecretText             BindingItemWorkersBindingKindHyperdriveType = "secret_text"
	BindingItemWorkersBindingKindHyperdriveTypeService                BindingItemWorkersBindingKindHyperdriveType = "service"
	BindingItemWorkersBindingKindHyperdriveTypeTailConsumer           BindingItemWorkersBindingKindHyperdriveType = "tail_consumer"
	BindingItemWorkersBindingKindHyperdriveTypeVectorize              BindingItemWorkersBindingKindHyperdriveType = "vectorize"
	BindingItemWorkersBindingKindHyperdriveTypeVersionMetadata        BindingItemWorkersBindingKindHyperdriveType = "version_metadata"
)

func (r BindingItemWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindHyperdriveTypeHyperdrive, BindingItemWorkersBindingKindHyperdriveTypeAI, BindingItemWorkersBindingKindHyperdriveTypeAnalyticsEngine, BindingItemWorkersBindingKindHyperdriveTypeAssets, BindingItemWorkersBindingKindHyperdriveTypeBrowserRendering, BindingItemWorkersBindingKindHyperdriveTypeD1, BindingItemWorkersBindingKindHyperdriveTypeDispatchNamespace, BindingItemWorkersBindingKindHyperdriveTypeDurableObjectNamespace, BindingItemWorkersBindingKindHyperdriveTypeJson, BindingItemWorkersBindingKindHyperdriveTypeKvNamespace, BindingItemWorkersBindingKindHyperdriveTypeMtlsCertificate, BindingItemWorkersBindingKindHyperdriveTypePlainText, BindingItemWorkersBindingKindHyperdriveTypeQueue, BindingItemWorkersBindingKindHyperdriveTypeR2Bucket, BindingItemWorkersBindingKindHyperdriveTypeSecretText, BindingItemWorkersBindingKindHyperdriveTypeService, BindingItemWorkersBindingKindHyperdriveTypeTailConsumer, BindingItemWorkersBindingKindHyperdriveTypeVectorize, BindingItemWorkersBindingKindHyperdriveTypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindJson struct {
	// JSON data to use.
	Json interface{} `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string                                `json:"name,required"`
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

type BindingItemWorkersBindingKindJsonType string

const (
	BindingItemWorkersBindingKindJsonTypeJson                   BindingItemWorkersBindingKindJsonType = "json"
	BindingItemWorkersBindingKindJsonTypeAI                     BindingItemWorkersBindingKindJsonType = "ai"
	BindingItemWorkersBindingKindJsonTypeAnalyticsEngine        BindingItemWorkersBindingKindJsonType = "analytics_engine"
	BindingItemWorkersBindingKindJsonTypeAssets                 BindingItemWorkersBindingKindJsonType = "assets"
	BindingItemWorkersBindingKindJsonTypeBrowserRendering       BindingItemWorkersBindingKindJsonType = "browser_rendering"
	BindingItemWorkersBindingKindJsonTypeD1                     BindingItemWorkersBindingKindJsonType = "d1"
	BindingItemWorkersBindingKindJsonTypeDispatchNamespace      BindingItemWorkersBindingKindJsonType = "dispatch_namespace"
	BindingItemWorkersBindingKindJsonTypeDurableObjectNamespace BindingItemWorkersBindingKindJsonType = "durable_object_namespace"
	BindingItemWorkersBindingKindJsonTypeHyperdrive             BindingItemWorkersBindingKindJsonType = "hyperdrive"
	BindingItemWorkersBindingKindJsonTypeKvNamespace            BindingItemWorkersBindingKindJsonType = "kv_namespace"
	BindingItemWorkersBindingKindJsonTypeMtlsCertificate        BindingItemWorkersBindingKindJsonType = "mtls_certificate"
	BindingItemWorkersBindingKindJsonTypePlainText              BindingItemWorkersBindingKindJsonType = "plain_text"
	BindingItemWorkersBindingKindJsonTypeQueue                  BindingItemWorkersBindingKindJsonType = "queue"
	BindingItemWorkersBindingKindJsonTypeR2Bucket               BindingItemWorkersBindingKindJsonType = "r2_bucket"
	BindingItemWorkersBindingKindJsonTypeSecretText             BindingItemWorkersBindingKindJsonType = "secret_text"
	BindingItemWorkersBindingKindJsonTypeService                BindingItemWorkersBindingKindJsonType = "service"
	BindingItemWorkersBindingKindJsonTypeTailConsumer           BindingItemWorkersBindingKindJsonType = "tail_consumer"
	BindingItemWorkersBindingKindJsonTypeVectorize              BindingItemWorkersBindingKindJsonType = "vectorize"
	BindingItemWorkersBindingKindJsonTypeVersionMetadata        BindingItemWorkersBindingKindJsonType = "version_metadata"
)

func (r BindingItemWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindJsonTypeJson, BindingItemWorkersBindingKindJsonTypeAI, BindingItemWorkersBindingKindJsonTypeAnalyticsEngine, BindingItemWorkersBindingKindJsonTypeAssets, BindingItemWorkersBindingKindJsonTypeBrowserRendering, BindingItemWorkersBindingKindJsonTypeD1, BindingItemWorkersBindingKindJsonTypeDispatchNamespace, BindingItemWorkersBindingKindJsonTypeDurableObjectNamespace, BindingItemWorkersBindingKindJsonTypeHyperdrive, BindingItemWorkersBindingKindJsonTypeKvNamespace, BindingItemWorkersBindingKindJsonTypeMtlsCertificate, BindingItemWorkersBindingKindJsonTypePlainText, BindingItemWorkersBindingKindJsonTypeQueue, BindingItemWorkersBindingKindJsonTypeR2Bucket, BindingItemWorkersBindingKindJsonTypeSecretText, BindingItemWorkersBindingKindJsonTypeService, BindingItemWorkersBindingKindJsonTypeTailConsumer, BindingItemWorkersBindingKindJsonTypeVectorize, BindingItemWorkersBindingKindJsonTypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindKvNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string                                       `json:"namespace_id,required"`
	Type        BindingItemWorkersBindingKindKvNamespaceType `json:"type,required"`
	JSON        bindingItemWorkersBindingKindKvNamespaceJSON `json:"-"`
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

type BindingItemWorkersBindingKindKvNamespaceType string

const (
	BindingItemWorkersBindingKindKvNamespaceTypeKvNamespace            BindingItemWorkersBindingKindKvNamespaceType = "kv_namespace"
	BindingItemWorkersBindingKindKvNamespaceTypeAI                     BindingItemWorkersBindingKindKvNamespaceType = "ai"
	BindingItemWorkersBindingKindKvNamespaceTypeAnalyticsEngine        BindingItemWorkersBindingKindKvNamespaceType = "analytics_engine"
	BindingItemWorkersBindingKindKvNamespaceTypeAssets                 BindingItemWorkersBindingKindKvNamespaceType = "assets"
	BindingItemWorkersBindingKindKvNamespaceTypeBrowserRendering       BindingItemWorkersBindingKindKvNamespaceType = "browser_rendering"
	BindingItemWorkersBindingKindKvNamespaceTypeD1                     BindingItemWorkersBindingKindKvNamespaceType = "d1"
	BindingItemWorkersBindingKindKvNamespaceTypeDispatchNamespace      BindingItemWorkersBindingKindKvNamespaceType = "dispatch_namespace"
	BindingItemWorkersBindingKindKvNamespaceTypeDurableObjectNamespace BindingItemWorkersBindingKindKvNamespaceType = "durable_object_namespace"
	BindingItemWorkersBindingKindKvNamespaceTypeHyperdrive             BindingItemWorkersBindingKindKvNamespaceType = "hyperdrive"
	BindingItemWorkersBindingKindKvNamespaceTypeJson                   BindingItemWorkersBindingKindKvNamespaceType = "json"
	BindingItemWorkersBindingKindKvNamespaceTypeMtlsCertificate        BindingItemWorkersBindingKindKvNamespaceType = "mtls_certificate"
	BindingItemWorkersBindingKindKvNamespaceTypePlainText              BindingItemWorkersBindingKindKvNamespaceType = "plain_text"
	BindingItemWorkersBindingKindKvNamespaceTypeQueue                  BindingItemWorkersBindingKindKvNamespaceType = "queue"
	BindingItemWorkersBindingKindKvNamespaceTypeR2Bucket               BindingItemWorkersBindingKindKvNamespaceType = "r2_bucket"
	BindingItemWorkersBindingKindKvNamespaceTypeSecretText             BindingItemWorkersBindingKindKvNamespaceType = "secret_text"
	BindingItemWorkersBindingKindKvNamespaceTypeService                BindingItemWorkersBindingKindKvNamespaceType = "service"
	BindingItemWorkersBindingKindKvNamespaceTypeTailConsumer           BindingItemWorkersBindingKindKvNamespaceType = "tail_consumer"
	BindingItemWorkersBindingKindKvNamespaceTypeVectorize              BindingItemWorkersBindingKindKvNamespaceType = "vectorize"
	BindingItemWorkersBindingKindKvNamespaceTypeVersionMetadata        BindingItemWorkersBindingKindKvNamespaceType = "version_metadata"
)

func (r BindingItemWorkersBindingKindKvNamespaceType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindKvNamespaceTypeKvNamespace, BindingItemWorkersBindingKindKvNamespaceTypeAI, BindingItemWorkersBindingKindKvNamespaceTypeAnalyticsEngine, BindingItemWorkersBindingKindKvNamespaceTypeAssets, BindingItemWorkersBindingKindKvNamespaceTypeBrowserRendering, BindingItemWorkersBindingKindKvNamespaceTypeD1, BindingItemWorkersBindingKindKvNamespaceTypeDispatchNamespace, BindingItemWorkersBindingKindKvNamespaceTypeDurableObjectNamespace, BindingItemWorkersBindingKindKvNamespaceTypeHyperdrive, BindingItemWorkersBindingKindKvNamespaceTypeJson, BindingItemWorkersBindingKindKvNamespaceTypeMtlsCertificate, BindingItemWorkersBindingKindKvNamespaceTypePlainText, BindingItemWorkersBindingKindKvNamespaceTypeQueue, BindingItemWorkersBindingKindKvNamespaceTypeR2Bucket, BindingItemWorkersBindingKindKvNamespaceTypeSecretText, BindingItemWorkersBindingKindKvNamespaceTypeService, BindingItemWorkersBindingKindKvNamespaceTypeTailConsumer, BindingItemWorkersBindingKindKvNamespaceTypeVectorize, BindingItemWorkersBindingKindKvNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindMtlsCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string                                           `json:"name,required"`
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

type BindingItemWorkersBindingKindMtlsCertificateType string

const (
	BindingItemWorkersBindingKindMtlsCertificateTypeMtlsCertificate        BindingItemWorkersBindingKindMtlsCertificateType = "mtls_certificate"
	BindingItemWorkersBindingKindMtlsCertificateTypeAI                     BindingItemWorkersBindingKindMtlsCertificateType = "ai"
	BindingItemWorkersBindingKindMtlsCertificateTypeAnalyticsEngine        BindingItemWorkersBindingKindMtlsCertificateType = "analytics_engine"
	BindingItemWorkersBindingKindMtlsCertificateTypeAssets                 BindingItemWorkersBindingKindMtlsCertificateType = "assets"
	BindingItemWorkersBindingKindMtlsCertificateTypeBrowserRendering       BindingItemWorkersBindingKindMtlsCertificateType = "browser_rendering"
	BindingItemWorkersBindingKindMtlsCertificateTypeD1                     BindingItemWorkersBindingKindMtlsCertificateType = "d1"
	BindingItemWorkersBindingKindMtlsCertificateTypeDispatchNamespace      BindingItemWorkersBindingKindMtlsCertificateType = "dispatch_namespace"
	BindingItemWorkersBindingKindMtlsCertificateTypeDurableObjectNamespace BindingItemWorkersBindingKindMtlsCertificateType = "durable_object_namespace"
	BindingItemWorkersBindingKindMtlsCertificateTypeHyperdrive             BindingItemWorkersBindingKindMtlsCertificateType = "hyperdrive"
	BindingItemWorkersBindingKindMtlsCertificateTypeJson                   BindingItemWorkersBindingKindMtlsCertificateType = "json"
	BindingItemWorkersBindingKindMtlsCertificateTypeKvNamespace            BindingItemWorkersBindingKindMtlsCertificateType = "kv_namespace"
	BindingItemWorkersBindingKindMtlsCertificateTypePlainText              BindingItemWorkersBindingKindMtlsCertificateType = "plain_text"
	BindingItemWorkersBindingKindMtlsCertificateTypeQueue                  BindingItemWorkersBindingKindMtlsCertificateType = "queue"
	BindingItemWorkersBindingKindMtlsCertificateTypeR2Bucket               BindingItemWorkersBindingKindMtlsCertificateType = "r2_bucket"
	BindingItemWorkersBindingKindMtlsCertificateTypeSecretText             BindingItemWorkersBindingKindMtlsCertificateType = "secret_text"
	BindingItemWorkersBindingKindMtlsCertificateTypeService                BindingItemWorkersBindingKindMtlsCertificateType = "service"
	BindingItemWorkersBindingKindMtlsCertificateTypeTailConsumer           BindingItemWorkersBindingKindMtlsCertificateType = "tail_consumer"
	BindingItemWorkersBindingKindMtlsCertificateTypeVectorize              BindingItemWorkersBindingKindMtlsCertificateType = "vectorize"
	BindingItemWorkersBindingKindMtlsCertificateTypeVersionMetadata        BindingItemWorkersBindingKindMtlsCertificateType = "version_metadata"
)

func (r BindingItemWorkersBindingKindMtlsCertificateType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindMtlsCertificateTypeMtlsCertificate, BindingItemWorkersBindingKindMtlsCertificateTypeAI, BindingItemWorkersBindingKindMtlsCertificateTypeAnalyticsEngine, BindingItemWorkersBindingKindMtlsCertificateTypeAssets, BindingItemWorkersBindingKindMtlsCertificateTypeBrowserRendering, BindingItemWorkersBindingKindMtlsCertificateTypeD1, BindingItemWorkersBindingKindMtlsCertificateTypeDispatchNamespace, BindingItemWorkersBindingKindMtlsCertificateTypeDurableObjectNamespace, BindingItemWorkersBindingKindMtlsCertificateTypeHyperdrive, BindingItemWorkersBindingKindMtlsCertificateTypeJson, BindingItemWorkersBindingKindMtlsCertificateTypeKvNamespace, BindingItemWorkersBindingKindMtlsCertificateTypePlainText, BindingItemWorkersBindingKindMtlsCertificateTypeQueue, BindingItemWorkersBindingKindMtlsCertificateTypeR2Bucket, BindingItemWorkersBindingKindMtlsCertificateTypeSecretText, BindingItemWorkersBindingKindMtlsCertificateTypeService, BindingItemWorkersBindingKindMtlsCertificateTypeTailConsumer, BindingItemWorkersBindingKindMtlsCertificateTypeVectorize, BindingItemWorkersBindingKindMtlsCertificateTypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string                                     `json:"text,required"`
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

type BindingItemWorkersBindingKindPlainTextType string

const (
	BindingItemWorkersBindingKindPlainTextTypePlainText              BindingItemWorkersBindingKindPlainTextType = "plain_text"
	BindingItemWorkersBindingKindPlainTextTypeAI                     BindingItemWorkersBindingKindPlainTextType = "ai"
	BindingItemWorkersBindingKindPlainTextTypeAnalyticsEngine        BindingItemWorkersBindingKindPlainTextType = "analytics_engine"
	BindingItemWorkersBindingKindPlainTextTypeAssets                 BindingItemWorkersBindingKindPlainTextType = "assets"
	BindingItemWorkersBindingKindPlainTextTypeBrowserRendering       BindingItemWorkersBindingKindPlainTextType = "browser_rendering"
	BindingItemWorkersBindingKindPlainTextTypeD1                     BindingItemWorkersBindingKindPlainTextType = "d1"
	BindingItemWorkersBindingKindPlainTextTypeDispatchNamespace      BindingItemWorkersBindingKindPlainTextType = "dispatch_namespace"
	BindingItemWorkersBindingKindPlainTextTypeDurableObjectNamespace BindingItemWorkersBindingKindPlainTextType = "durable_object_namespace"
	BindingItemWorkersBindingKindPlainTextTypeHyperdrive             BindingItemWorkersBindingKindPlainTextType = "hyperdrive"
	BindingItemWorkersBindingKindPlainTextTypeJson                   BindingItemWorkersBindingKindPlainTextType = "json"
	BindingItemWorkersBindingKindPlainTextTypeKvNamespace            BindingItemWorkersBindingKindPlainTextType = "kv_namespace"
	BindingItemWorkersBindingKindPlainTextTypeMtlsCertificate        BindingItemWorkersBindingKindPlainTextType = "mtls_certificate"
	BindingItemWorkersBindingKindPlainTextTypeQueue                  BindingItemWorkersBindingKindPlainTextType = "queue"
	BindingItemWorkersBindingKindPlainTextTypeR2Bucket               BindingItemWorkersBindingKindPlainTextType = "r2_bucket"
	BindingItemWorkersBindingKindPlainTextTypeSecretText             BindingItemWorkersBindingKindPlainTextType = "secret_text"
	BindingItemWorkersBindingKindPlainTextTypeService                BindingItemWorkersBindingKindPlainTextType = "service"
	BindingItemWorkersBindingKindPlainTextTypeTailConsumer           BindingItemWorkersBindingKindPlainTextType = "tail_consumer"
	BindingItemWorkersBindingKindPlainTextTypeVectorize              BindingItemWorkersBindingKindPlainTextType = "vectorize"
	BindingItemWorkersBindingKindPlainTextTypeVersionMetadata        BindingItemWorkersBindingKindPlainTextType = "version_metadata"
)

func (r BindingItemWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindPlainTextTypePlainText, BindingItemWorkersBindingKindPlainTextTypeAI, BindingItemWorkersBindingKindPlainTextTypeAnalyticsEngine, BindingItemWorkersBindingKindPlainTextTypeAssets, BindingItemWorkersBindingKindPlainTextTypeBrowserRendering, BindingItemWorkersBindingKindPlainTextTypeD1, BindingItemWorkersBindingKindPlainTextTypeDispatchNamespace, BindingItemWorkersBindingKindPlainTextTypeDurableObjectNamespace, BindingItemWorkersBindingKindPlainTextTypeHyperdrive, BindingItemWorkersBindingKindPlainTextTypeJson, BindingItemWorkersBindingKindPlainTextTypeKvNamespace, BindingItemWorkersBindingKindPlainTextTypeMtlsCertificate, BindingItemWorkersBindingKindPlainTextTypeQueue, BindingItemWorkersBindingKindPlainTextTypeR2Bucket, BindingItemWorkersBindingKindPlainTextTypeSecretText, BindingItemWorkersBindingKindPlainTextTypeService, BindingItemWorkersBindingKindPlainTextTypeTailConsumer, BindingItemWorkersBindingKindPlainTextTypeVectorize, BindingItemWorkersBindingKindPlainTextTypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string                                 `json:"queue_name,required"`
	Type      BindingItemWorkersBindingKindQueueType `json:"type,required"`
	JSON      bindingItemWorkersBindingKindQueueJSON `json:"-"`
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

type BindingItemWorkersBindingKindQueueType string

const (
	BindingItemWorkersBindingKindQueueTypeQueue                  BindingItemWorkersBindingKindQueueType = "queue"
	BindingItemWorkersBindingKindQueueTypeAI                     BindingItemWorkersBindingKindQueueType = "ai"
	BindingItemWorkersBindingKindQueueTypeAnalyticsEngine        BindingItemWorkersBindingKindQueueType = "analytics_engine"
	BindingItemWorkersBindingKindQueueTypeAssets                 BindingItemWorkersBindingKindQueueType = "assets"
	BindingItemWorkersBindingKindQueueTypeBrowserRendering       BindingItemWorkersBindingKindQueueType = "browser_rendering"
	BindingItemWorkersBindingKindQueueTypeD1                     BindingItemWorkersBindingKindQueueType = "d1"
	BindingItemWorkersBindingKindQueueTypeDispatchNamespace      BindingItemWorkersBindingKindQueueType = "dispatch_namespace"
	BindingItemWorkersBindingKindQueueTypeDurableObjectNamespace BindingItemWorkersBindingKindQueueType = "durable_object_namespace"
	BindingItemWorkersBindingKindQueueTypeHyperdrive             BindingItemWorkersBindingKindQueueType = "hyperdrive"
	BindingItemWorkersBindingKindQueueTypeJson                   BindingItemWorkersBindingKindQueueType = "json"
	BindingItemWorkersBindingKindQueueTypeKvNamespace            BindingItemWorkersBindingKindQueueType = "kv_namespace"
	BindingItemWorkersBindingKindQueueTypeMtlsCertificate        BindingItemWorkersBindingKindQueueType = "mtls_certificate"
	BindingItemWorkersBindingKindQueueTypePlainText              BindingItemWorkersBindingKindQueueType = "plain_text"
	BindingItemWorkersBindingKindQueueTypeR2Bucket               BindingItemWorkersBindingKindQueueType = "r2_bucket"
	BindingItemWorkersBindingKindQueueTypeSecretText             BindingItemWorkersBindingKindQueueType = "secret_text"
	BindingItemWorkersBindingKindQueueTypeService                BindingItemWorkersBindingKindQueueType = "service"
	BindingItemWorkersBindingKindQueueTypeTailConsumer           BindingItemWorkersBindingKindQueueType = "tail_consumer"
	BindingItemWorkersBindingKindQueueTypeVectorize              BindingItemWorkersBindingKindQueueType = "vectorize"
	BindingItemWorkersBindingKindQueueTypeVersionMetadata        BindingItemWorkersBindingKindQueueType = "version_metadata"
)

func (r BindingItemWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindQueueTypeQueue, BindingItemWorkersBindingKindQueueTypeAI, BindingItemWorkersBindingKindQueueTypeAnalyticsEngine, BindingItemWorkersBindingKindQueueTypeAssets, BindingItemWorkersBindingKindQueueTypeBrowserRendering, BindingItemWorkersBindingKindQueueTypeD1, BindingItemWorkersBindingKindQueueTypeDispatchNamespace, BindingItemWorkersBindingKindQueueTypeDurableObjectNamespace, BindingItemWorkersBindingKindQueueTypeHyperdrive, BindingItemWorkersBindingKindQueueTypeJson, BindingItemWorkersBindingKindQueueTypeKvNamespace, BindingItemWorkersBindingKindQueueTypeMtlsCertificate, BindingItemWorkersBindingKindQueueTypePlainText, BindingItemWorkersBindingKindQueueTypeR2Bucket, BindingItemWorkersBindingKindQueueTypeSecretText, BindingItemWorkersBindingKindQueueTypeService, BindingItemWorkersBindingKindQueueTypeTailConsumer, BindingItemWorkersBindingKindQueueTypeVectorize, BindingItemWorkersBindingKindQueueTypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string                                    `json:"name,required"`
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

type BindingItemWorkersBindingKindR2BucketType string

const (
	BindingItemWorkersBindingKindR2BucketTypeR2Bucket               BindingItemWorkersBindingKindR2BucketType = "r2_bucket"
	BindingItemWorkersBindingKindR2BucketTypeAI                     BindingItemWorkersBindingKindR2BucketType = "ai"
	BindingItemWorkersBindingKindR2BucketTypeAnalyticsEngine        BindingItemWorkersBindingKindR2BucketType = "analytics_engine"
	BindingItemWorkersBindingKindR2BucketTypeAssets                 BindingItemWorkersBindingKindR2BucketType = "assets"
	BindingItemWorkersBindingKindR2BucketTypeBrowserRendering       BindingItemWorkersBindingKindR2BucketType = "browser_rendering"
	BindingItemWorkersBindingKindR2BucketTypeD1                     BindingItemWorkersBindingKindR2BucketType = "d1"
	BindingItemWorkersBindingKindR2BucketTypeDispatchNamespace      BindingItemWorkersBindingKindR2BucketType = "dispatch_namespace"
	BindingItemWorkersBindingKindR2BucketTypeDurableObjectNamespace BindingItemWorkersBindingKindR2BucketType = "durable_object_namespace"
	BindingItemWorkersBindingKindR2BucketTypeHyperdrive             BindingItemWorkersBindingKindR2BucketType = "hyperdrive"
	BindingItemWorkersBindingKindR2BucketTypeJson                   BindingItemWorkersBindingKindR2BucketType = "json"
	BindingItemWorkersBindingKindR2BucketTypeKvNamespace            BindingItemWorkersBindingKindR2BucketType = "kv_namespace"
	BindingItemWorkersBindingKindR2BucketTypeMtlsCertificate        BindingItemWorkersBindingKindR2BucketType = "mtls_certificate"
	BindingItemWorkersBindingKindR2BucketTypePlainText              BindingItemWorkersBindingKindR2BucketType = "plain_text"
	BindingItemWorkersBindingKindR2BucketTypeQueue                  BindingItemWorkersBindingKindR2BucketType = "queue"
	BindingItemWorkersBindingKindR2BucketTypeSecretText             BindingItemWorkersBindingKindR2BucketType = "secret_text"
	BindingItemWorkersBindingKindR2BucketTypeService                BindingItemWorkersBindingKindR2BucketType = "service"
	BindingItemWorkersBindingKindR2BucketTypeTailConsumer           BindingItemWorkersBindingKindR2BucketType = "tail_consumer"
	BindingItemWorkersBindingKindR2BucketTypeVectorize              BindingItemWorkersBindingKindR2BucketType = "vectorize"
	BindingItemWorkersBindingKindR2BucketTypeVersionMetadata        BindingItemWorkersBindingKindR2BucketType = "version_metadata"
)

func (r BindingItemWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindR2BucketTypeR2Bucket, BindingItemWorkersBindingKindR2BucketTypeAI, BindingItemWorkersBindingKindR2BucketTypeAnalyticsEngine, BindingItemWorkersBindingKindR2BucketTypeAssets, BindingItemWorkersBindingKindR2BucketTypeBrowserRendering, BindingItemWorkersBindingKindR2BucketTypeD1, BindingItemWorkersBindingKindR2BucketTypeDispatchNamespace, BindingItemWorkersBindingKindR2BucketTypeDurableObjectNamespace, BindingItemWorkersBindingKindR2BucketTypeHyperdrive, BindingItemWorkersBindingKindR2BucketTypeJson, BindingItemWorkersBindingKindR2BucketTypeKvNamespace, BindingItemWorkersBindingKindR2BucketTypeMtlsCertificate, BindingItemWorkersBindingKindR2BucketTypePlainText, BindingItemWorkersBindingKindR2BucketTypeQueue, BindingItemWorkersBindingKindR2BucketTypeSecretText, BindingItemWorkersBindingKindR2BucketTypeService, BindingItemWorkersBindingKindR2BucketTypeTailConsumer, BindingItemWorkersBindingKindR2BucketTypeVectorize, BindingItemWorkersBindingKindR2BucketTypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The secret value to use.
	Text string                                      `json:"text,required"`
	Type BindingItemWorkersBindingKindSecretTextType `json:"type,required"`
	JSON bindingItemWorkersBindingKindSecretTextJSON `json:"-"`
}

// bindingItemWorkersBindingKindSecretTextJSON contains the JSON metadata for the
// struct [BindingItemWorkersBindingKindSecretText]
type bindingItemWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
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

type BindingItemWorkersBindingKindSecretTextType string

const (
	BindingItemWorkersBindingKindSecretTextTypeSecretText             BindingItemWorkersBindingKindSecretTextType = "secret_text"
	BindingItemWorkersBindingKindSecretTextTypeAI                     BindingItemWorkersBindingKindSecretTextType = "ai"
	BindingItemWorkersBindingKindSecretTextTypeAnalyticsEngine        BindingItemWorkersBindingKindSecretTextType = "analytics_engine"
	BindingItemWorkersBindingKindSecretTextTypeAssets                 BindingItemWorkersBindingKindSecretTextType = "assets"
	BindingItemWorkersBindingKindSecretTextTypeBrowserRendering       BindingItemWorkersBindingKindSecretTextType = "browser_rendering"
	BindingItemWorkersBindingKindSecretTextTypeD1                     BindingItemWorkersBindingKindSecretTextType = "d1"
	BindingItemWorkersBindingKindSecretTextTypeDispatchNamespace      BindingItemWorkersBindingKindSecretTextType = "dispatch_namespace"
	BindingItemWorkersBindingKindSecretTextTypeDurableObjectNamespace BindingItemWorkersBindingKindSecretTextType = "durable_object_namespace"
	BindingItemWorkersBindingKindSecretTextTypeHyperdrive             BindingItemWorkersBindingKindSecretTextType = "hyperdrive"
	BindingItemWorkersBindingKindSecretTextTypeJson                   BindingItemWorkersBindingKindSecretTextType = "json"
	BindingItemWorkersBindingKindSecretTextTypeKvNamespace            BindingItemWorkersBindingKindSecretTextType = "kv_namespace"
	BindingItemWorkersBindingKindSecretTextTypeMtlsCertificate        BindingItemWorkersBindingKindSecretTextType = "mtls_certificate"
	BindingItemWorkersBindingKindSecretTextTypePlainText              BindingItemWorkersBindingKindSecretTextType = "plain_text"
	BindingItemWorkersBindingKindSecretTextTypeQueue                  BindingItemWorkersBindingKindSecretTextType = "queue"
	BindingItemWorkersBindingKindSecretTextTypeR2Bucket               BindingItemWorkersBindingKindSecretTextType = "r2_bucket"
	BindingItemWorkersBindingKindSecretTextTypeService                BindingItemWorkersBindingKindSecretTextType = "service"
	BindingItemWorkersBindingKindSecretTextTypeTailConsumer           BindingItemWorkersBindingKindSecretTextType = "tail_consumer"
	BindingItemWorkersBindingKindSecretTextTypeVectorize              BindingItemWorkersBindingKindSecretTextType = "vectorize"
	BindingItemWorkersBindingKindSecretTextTypeVersionMetadata        BindingItemWorkersBindingKindSecretTextType = "version_metadata"
)

func (r BindingItemWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindSecretTextTypeSecretText, BindingItemWorkersBindingKindSecretTextTypeAI, BindingItemWorkersBindingKindSecretTextTypeAnalyticsEngine, BindingItemWorkersBindingKindSecretTextTypeAssets, BindingItemWorkersBindingKindSecretTextTypeBrowserRendering, BindingItemWorkersBindingKindSecretTextTypeD1, BindingItemWorkersBindingKindSecretTextTypeDispatchNamespace, BindingItemWorkersBindingKindSecretTextTypeDurableObjectNamespace, BindingItemWorkersBindingKindSecretTextTypeHyperdrive, BindingItemWorkersBindingKindSecretTextTypeJson, BindingItemWorkersBindingKindSecretTextTypeKvNamespace, BindingItemWorkersBindingKindSecretTextTypeMtlsCertificate, BindingItemWorkersBindingKindSecretTextTypePlainText, BindingItemWorkersBindingKindSecretTextTypeQueue, BindingItemWorkersBindingKindSecretTextTypeR2Bucket, BindingItemWorkersBindingKindSecretTextTypeService, BindingItemWorkersBindingKindSecretTextTypeTailConsumer, BindingItemWorkersBindingKindSecretTextTypeVectorize, BindingItemWorkersBindingKindSecretTextTypeVersionMetadata:
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
	Service string                                   `json:"service,required"`
	Type    BindingItemWorkersBindingKindServiceType `json:"type,required"`
	JSON    bindingItemWorkersBindingKindServiceJSON `json:"-"`
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

type BindingItemWorkersBindingKindServiceType string

const (
	BindingItemWorkersBindingKindServiceTypeService                BindingItemWorkersBindingKindServiceType = "service"
	BindingItemWorkersBindingKindServiceTypeAI                     BindingItemWorkersBindingKindServiceType = "ai"
	BindingItemWorkersBindingKindServiceTypeAnalyticsEngine        BindingItemWorkersBindingKindServiceType = "analytics_engine"
	BindingItemWorkersBindingKindServiceTypeAssets                 BindingItemWorkersBindingKindServiceType = "assets"
	BindingItemWorkersBindingKindServiceTypeBrowserRendering       BindingItemWorkersBindingKindServiceType = "browser_rendering"
	BindingItemWorkersBindingKindServiceTypeD1                     BindingItemWorkersBindingKindServiceType = "d1"
	BindingItemWorkersBindingKindServiceTypeDispatchNamespace      BindingItemWorkersBindingKindServiceType = "dispatch_namespace"
	BindingItemWorkersBindingKindServiceTypeDurableObjectNamespace BindingItemWorkersBindingKindServiceType = "durable_object_namespace"
	BindingItemWorkersBindingKindServiceTypeHyperdrive             BindingItemWorkersBindingKindServiceType = "hyperdrive"
	BindingItemWorkersBindingKindServiceTypeJson                   BindingItemWorkersBindingKindServiceType = "json"
	BindingItemWorkersBindingKindServiceTypeKvNamespace            BindingItemWorkersBindingKindServiceType = "kv_namespace"
	BindingItemWorkersBindingKindServiceTypeMtlsCertificate        BindingItemWorkersBindingKindServiceType = "mtls_certificate"
	BindingItemWorkersBindingKindServiceTypePlainText              BindingItemWorkersBindingKindServiceType = "plain_text"
	BindingItemWorkersBindingKindServiceTypeQueue                  BindingItemWorkersBindingKindServiceType = "queue"
	BindingItemWorkersBindingKindServiceTypeR2Bucket               BindingItemWorkersBindingKindServiceType = "r2_bucket"
	BindingItemWorkersBindingKindServiceTypeSecretText             BindingItemWorkersBindingKindServiceType = "secret_text"
	BindingItemWorkersBindingKindServiceTypeTailConsumer           BindingItemWorkersBindingKindServiceType = "tail_consumer"
	BindingItemWorkersBindingKindServiceTypeVectorize              BindingItemWorkersBindingKindServiceType = "vectorize"
	BindingItemWorkersBindingKindServiceTypeVersionMetadata        BindingItemWorkersBindingKindServiceType = "version_metadata"
)

func (r BindingItemWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindServiceTypeService, BindingItemWorkersBindingKindServiceTypeAI, BindingItemWorkersBindingKindServiceTypeAnalyticsEngine, BindingItemWorkersBindingKindServiceTypeAssets, BindingItemWorkersBindingKindServiceTypeBrowserRendering, BindingItemWorkersBindingKindServiceTypeD1, BindingItemWorkersBindingKindServiceTypeDispatchNamespace, BindingItemWorkersBindingKindServiceTypeDurableObjectNamespace, BindingItemWorkersBindingKindServiceTypeHyperdrive, BindingItemWorkersBindingKindServiceTypeJson, BindingItemWorkersBindingKindServiceTypeKvNamespace, BindingItemWorkersBindingKindServiceTypeMtlsCertificate, BindingItemWorkersBindingKindServiceTypePlainText, BindingItemWorkersBindingKindServiceTypeQueue, BindingItemWorkersBindingKindServiceTypeR2Bucket, BindingItemWorkersBindingKindServiceTypeSecretText, BindingItemWorkersBindingKindServiceTypeTailConsumer, BindingItemWorkersBindingKindServiceTypeVectorize, BindingItemWorkersBindingKindServiceTypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string                                        `json:"service,required"`
	Type    BindingItemWorkersBindingKindTailConsumerType `json:"type,required"`
	JSON    bindingItemWorkersBindingKindTailConsumerJSON `json:"-"`
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

type BindingItemWorkersBindingKindTailConsumerType string

const (
	BindingItemWorkersBindingKindTailConsumerTypeTailConsumer           BindingItemWorkersBindingKindTailConsumerType = "tail_consumer"
	BindingItemWorkersBindingKindTailConsumerTypeAI                     BindingItemWorkersBindingKindTailConsumerType = "ai"
	BindingItemWorkersBindingKindTailConsumerTypeAnalyticsEngine        BindingItemWorkersBindingKindTailConsumerType = "analytics_engine"
	BindingItemWorkersBindingKindTailConsumerTypeAssets                 BindingItemWorkersBindingKindTailConsumerType = "assets"
	BindingItemWorkersBindingKindTailConsumerTypeBrowserRendering       BindingItemWorkersBindingKindTailConsumerType = "browser_rendering"
	BindingItemWorkersBindingKindTailConsumerTypeD1                     BindingItemWorkersBindingKindTailConsumerType = "d1"
	BindingItemWorkersBindingKindTailConsumerTypeDispatchNamespace      BindingItemWorkersBindingKindTailConsumerType = "dispatch_namespace"
	BindingItemWorkersBindingKindTailConsumerTypeDurableObjectNamespace BindingItemWorkersBindingKindTailConsumerType = "durable_object_namespace"
	BindingItemWorkersBindingKindTailConsumerTypeHyperdrive             BindingItemWorkersBindingKindTailConsumerType = "hyperdrive"
	BindingItemWorkersBindingKindTailConsumerTypeJson                   BindingItemWorkersBindingKindTailConsumerType = "json"
	BindingItemWorkersBindingKindTailConsumerTypeKvNamespace            BindingItemWorkersBindingKindTailConsumerType = "kv_namespace"
	BindingItemWorkersBindingKindTailConsumerTypeMtlsCertificate        BindingItemWorkersBindingKindTailConsumerType = "mtls_certificate"
	BindingItemWorkersBindingKindTailConsumerTypePlainText              BindingItemWorkersBindingKindTailConsumerType = "plain_text"
	BindingItemWorkersBindingKindTailConsumerTypeQueue                  BindingItemWorkersBindingKindTailConsumerType = "queue"
	BindingItemWorkersBindingKindTailConsumerTypeR2Bucket               BindingItemWorkersBindingKindTailConsumerType = "r2_bucket"
	BindingItemWorkersBindingKindTailConsumerTypeSecretText             BindingItemWorkersBindingKindTailConsumerType = "secret_text"
	BindingItemWorkersBindingKindTailConsumerTypeService                BindingItemWorkersBindingKindTailConsumerType = "service"
	BindingItemWorkersBindingKindTailConsumerTypeVectorize              BindingItemWorkersBindingKindTailConsumerType = "vectorize"
	BindingItemWorkersBindingKindTailConsumerTypeVersionMetadata        BindingItemWorkersBindingKindTailConsumerType = "version_metadata"
)

func (r BindingItemWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindTailConsumerTypeTailConsumer, BindingItemWorkersBindingKindTailConsumerTypeAI, BindingItemWorkersBindingKindTailConsumerTypeAnalyticsEngine, BindingItemWorkersBindingKindTailConsumerTypeAssets, BindingItemWorkersBindingKindTailConsumerTypeBrowserRendering, BindingItemWorkersBindingKindTailConsumerTypeD1, BindingItemWorkersBindingKindTailConsumerTypeDispatchNamespace, BindingItemWorkersBindingKindTailConsumerTypeDurableObjectNamespace, BindingItemWorkersBindingKindTailConsumerTypeHyperdrive, BindingItemWorkersBindingKindTailConsumerTypeJson, BindingItemWorkersBindingKindTailConsumerTypeKvNamespace, BindingItemWorkersBindingKindTailConsumerTypeMtlsCertificate, BindingItemWorkersBindingKindTailConsumerTypePlainText, BindingItemWorkersBindingKindTailConsumerTypeQueue, BindingItemWorkersBindingKindTailConsumerTypeR2Bucket, BindingItemWorkersBindingKindTailConsumerTypeSecretText, BindingItemWorkersBindingKindTailConsumerTypeService, BindingItemWorkersBindingKindTailConsumerTypeVectorize, BindingItemWorkersBindingKindTailConsumerTypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string                                     `json:"name,required"`
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

type BindingItemWorkersBindingKindVectorizeType string

const (
	BindingItemWorkersBindingKindVectorizeTypeVectorize              BindingItemWorkersBindingKindVectorizeType = "vectorize"
	BindingItemWorkersBindingKindVectorizeTypeAI                     BindingItemWorkersBindingKindVectorizeType = "ai"
	BindingItemWorkersBindingKindVectorizeTypeAnalyticsEngine        BindingItemWorkersBindingKindVectorizeType = "analytics_engine"
	BindingItemWorkersBindingKindVectorizeTypeAssets                 BindingItemWorkersBindingKindVectorizeType = "assets"
	BindingItemWorkersBindingKindVectorizeTypeBrowserRendering       BindingItemWorkersBindingKindVectorizeType = "browser_rendering"
	BindingItemWorkersBindingKindVectorizeTypeD1                     BindingItemWorkersBindingKindVectorizeType = "d1"
	BindingItemWorkersBindingKindVectorizeTypeDispatchNamespace      BindingItemWorkersBindingKindVectorizeType = "dispatch_namespace"
	BindingItemWorkersBindingKindVectorizeTypeDurableObjectNamespace BindingItemWorkersBindingKindVectorizeType = "durable_object_namespace"
	BindingItemWorkersBindingKindVectorizeTypeHyperdrive             BindingItemWorkersBindingKindVectorizeType = "hyperdrive"
	BindingItemWorkersBindingKindVectorizeTypeJson                   BindingItemWorkersBindingKindVectorizeType = "json"
	BindingItemWorkersBindingKindVectorizeTypeKvNamespace            BindingItemWorkersBindingKindVectorizeType = "kv_namespace"
	BindingItemWorkersBindingKindVectorizeTypeMtlsCertificate        BindingItemWorkersBindingKindVectorizeType = "mtls_certificate"
	BindingItemWorkersBindingKindVectorizeTypePlainText              BindingItemWorkersBindingKindVectorizeType = "plain_text"
	BindingItemWorkersBindingKindVectorizeTypeQueue                  BindingItemWorkersBindingKindVectorizeType = "queue"
	BindingItemWorkersBindingKindVectorizeTypeR2Bucket               BindingItemWorkersBindingKindVectorizeType = "r2_bucket"
	BindingItemWorkersBindingKindVectorizeTypeSecretText             BindingItemWorkersBindingKindVectorizeType = "secret_text"
	BindingItemWorkersBindingKindVectorizeTypeService                BindingItemWorkersBindingKindVectorizeType = "service"
	BindingItemWorkersBindingKindVectorizeTypeTailConsumer           BindingItemWorkersBindingKindVectorizeType = "tail_consumer"
	BindingItemWorkersBindingKindVectorizeTypeVersionMetadata        BindingItemWorkersBindingKindVectorizeType = "version_metadata"
)

func (r BindingItemWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindVectorizeTypeVectorize, BindingItemWorkersBindingKindVectorizeTypeAI, BindingItemWorkersBindingKindVectorizeTypeAnalyticsEngine, BindingItemWorkersBindingKindVectorizeTypeAssets, BindingItemWorkersBindingKindVectorizeTypeBrowserRendering, BindingItemWorkersBindingKindVectorizeTypeD1, BindingItemWorkersBindingKindVectorizeTypeDispatchNamespace, BindingItemWorkersBindingKindVectorizeTypeDurableObjectNamespace, BindingItemWorkersBindingKindVectorizeTypeHyperdrive, BindingItemWorkersBindingKindVectorizeTypeJson, BindingItemWorkersBindingKindVectorizeTypeKvNamespace, BindingItemWorkersBindingKindVectorizeTypeMtlsCertificate, BindingItemWorkersBindingKindVectorizeTypePlainText, BindingItemWorkersBindingKindVectorizeTypeQueue, BindingItemWorkersBindingKindVectorizeTypeR2Bucket, BindingItemWorkersBindingKindVectorizeTypeSecretText, BindingItemWorkersBindingKindVectorizeTypeService, BindingItemWorkersBindingKindVectorizeTypeTailConsumer, BindingItemWorkersBindingKindVectorizeTypeVersionMetadata:
		return true
	}
	return false
}

type BindingItemWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string                                           `json:"name,required"`
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

type BindingItemWorkersBindingKindVersionMetadataType string

const (
	BindingItemWorkersBindingKindVersionMetadataTypeVersionMetadata        BindingItemWorkersBindingKindVersionMetadataType = "version_metadata"
	BindingItemWorkersBindingKindVersionMetadataTypeAI                     BindingItemWorkersBindingKindVersionMetadataType = "ai"
	BindingItemWorkersBindingKindVersionMetadataTypeAnalyticsEngine        BindingItemWorkersBindingKindVersionMetadataType = "analytics_engine"
	BindingItemWorkersBindingKindVersionMetadataTypeAssets                 BindingItemWorkersBindingKindVersionMetadataType = "assets"
	BindingItemWorkersBindingKindVersionMetadataTypeBrowserRendering       BindingItemWorkersBindingKindVersionMetadataType = "browser_rendering"
	BindingItemWorkersBindingKindVersionMetadataTypeD1                     BindingItemWorkersBindingKindVersionMetadataType = "d1"
	BindingItemWorkersBindingKindVersionMetadataTypeDispatchNamespace      BindingItemWorkersBindingKindVersionMetadataType = "dispatch_namespace"
	BindingItemWorkersBindingKindVersionMetadataTypeDurableObjectNamespace BindingItemWorkersBindingKindVersionMetadataType = "durable_object_namespace"
	BindingItemWorkersBindingKindVersionMetadataTypeHyperdrive             BindingItemWorkersBindingKindVersionMetadataType = "hyperdrive"
	BindingItemWorkersBindingKindVersionMetadataTypeJson                   BindingItemWorkersBindingKindVersionMetadataType = "json"
	BindingItemWorkersBindingKindVersionMetadataTypeKvNamespace            BindingItemWorkersBindingKindVersionMetadataType = "kv_namespace"
	BindingItemWorkersBindingKindVersionMetadataTypeMtlsCertificate        BindingItemWorkersBindingKindVersionMetadataType = "mtls_certificate"
	BindingItemWorkersBindingKindVersionMetadataTypePlainText              BindingItemWorkersBindingKindVersionMetadataType = "plain_text"
	BindingItemWorkersBindingKindVersionMetadataTypeQueue                  BindingItemWorkersBindingKindVersionMetadataType = "queue"
	BindingItemWorkersBindingKindVersionMetadataTypeR2Bucket               BindingItemWorkersBindingKindVersionMetadataType = "r2_bucket"
	BindingItemWorkersBindingKindVersionMetadataTypeSecretText             BindingItemWorkersBindingKindVersionMetadataType = "secret_text"
	BindingItemWorkersBindingKindVersionMetadataTypeService                BindingItemWorkersBindingKindVersionMetadataType = "service"
	BindingItemWorkersBindingKindVersionMetadataTypeTailConsumer           BindingItemWorkersBindingKindVersionMetadataType = "tail_consumer"
	BindingItemWorkersBindingKindVersionMetadataTypeVectorize              BindingItemWorkersBindingKindVersionMetadataType = "vectorize"
)

func (r BindingItemWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case BindingItemWorkersBindingKindVersionMetadataTypeVersionMetadata, BindingItemWorkersBindingKindVersionMetadataTypeAI, BindingItemWorkersBindingKindVersionMetadataTypeAnalyticsEngine, BindingItemWorkersBindingKindVersionMetadataTypeAssets, BindingItemWorkersBindingKindVersionMetadataTypeBrowserRendering, BindingItemWorkersBindingKindVersionMetadataTypeD1, BindingItemWorkersBindingKindVersionMetadataTypeDispatchNamespace, BindingItemWorkersBindingKindVersionMetadataTypeDurableObjectNamespace, BindingItemWorkersBindingKindVersionMetadataTypeHyperdrive, BindingItemWorkersBindingKindVersionMetadataTypeJson, BindingItemWorkersBindingKindVersionMetadataTypeKvNamespace, BindingItemWorkersBindingKindVersionMetadataTypeMtlsCertificate, BindingItemWorkersBindingKindVersionMetadataTypePlainText, BindingItemWorkersBindingKindVersionMetadataTypeQueue, BindingItemWorkersBindingKindVersionMetadataTypeR2Bucket, BindingItemWorkersBindingKindVersionMetadataTypeSecretText, BindingItemWorkersBindingKindVersionMetadataTypeService, BindingItemWorkersBindingKindVersionMetadataTypeTailConsumer, BindingItemWorkersBindingKindVersionMetadataTypeVectorize:
		return true
	}
	return false
}

type BindingItemType string

const (
	BindingItemTypeAI                     BindingItemType = "ai"
	BindingItemTypeAnalyticsEngine        BindingItemType = "analytics_engine"
	BindingItemTypeAssets                 BindingItemType = "assets"
	BindingItemTypeBrowserRendering       BindingItemType = "browser_rendering"
	BindingItemTypeD1                     BindingItemType = "d1"
	BindingItemTypeDispatchNamespace      BindingItemType = "dispatch_namespace"
	BindingItemTypeDurableObjectNamespace BindingItemType = "durable_object_namespace"
	BindingItemTypeHyperdrive             BindingItemType = "hyperdrive"
	BindingItemTypeJson                   BindingItemType = "json"
	BindingItemTypeKvNamespace            BindingItemType = "kv_namespace"
	BindingItemTypeMtlsCertificate        BindingItemType = "mtls_certificate"
	BindingItemTypePlainText              BindingItemType = "plain_text"
	BindingItemTypeQueue                  BindingItemType = "queue"
	BindingItemTypeR2Bucket               BindingItemType = "r2_bucket"
	BindingItemTypeSecretText             BindingItemType = "secret_text"
	BindingItemTypeService                BindingItemType = "service"
	BindingItemTypeTailConsumer           BindingItemType = "tail_consumer"
	BindingItemTypeVectorize              BindingItemType = "vectorize"
	BindingItemTypeVersionMetadata        BindingItemType = "version_metadata"
)

func (r BindingItemType) IsKnown() bool {
	switch r {
	case BindingItemTypeAI, BindingItemTypeAnalyticsEngine, BindingItemTypeAssets, BindingItemTypeBrowserRendering, BindingItemTypeD1, BindingItemTypeDispatchNamespace, BindingItemTypeDurableObjectNamespace, BindingItemTypeHyperdrive, BindingItemTypeJson, BindingItemTypeKvNamespace, BindingItemTypeMtlsCertificate, BindingItemTypePlainText, BindingItemTypeQueue, BindingItemTypeR2Bucket, BindingItemTypeSecretText, BindingItemTypeService, BindingItemTypeTailConsumer, BindingItemTypeVectorize, BindingItemTypeVersionMetadata:
		return true
	}
	return false
}

// A binding to allow the Worker to communicate with resources
type BindingItemParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string]          `json:"name,required"`
	Type param.Field[BindingItemType] `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id"`
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
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string]      `json:"index_name"`
	Json      param.Field[interface{}] `json:"json"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID param.Field[string]      `json:"namespace_id"`
	Outbound    param.Field[interface{}] `json:"outbound"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName param.Field[string] `json:"script_name"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service"`
	// The text value to use.
	Text param.Field[string] `json:"text"`
}

func (r BindingItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemParam) implementsBindingItemUnionParam() {}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by [BindingItemWorkersBindingKindAIParam],
// [BindingItemWorkersBindingKindAnalyticsEngineParam],
// [BindingItemWorkersBindingKindAssetsParam],
// [BindingItemWorkersBindingKindBrowserRenderingParam],
// [BindingItemWorkersBindingKindD1Param],
// [BindingItemWorkersBindingKindDispatchNamespaceParam],
// [BindingItemWorkersBindingKindDurableObjectNamespaceParam],
// [BindingItemWorkersBindingKindHyperdriveParam],
// [BindingItemWorkersBindingKindJsonParam],
// [BindingItemWorkersBindingKindKvNamespaceParam],
// [BindingItemWorkersBindingKindMtlsCertificateParam],
// [BindingItemWorkersBindingKindPlainTextParam],
// [BindingItemWorkersBindingKindQueueParam],
// [BindingItemWorkersBindingKindR2BucketParam],
// [BindingItemWorkersBindingKindSecretTextParam],
// [BindingItemWorkersBindingKindServiceParam],
// [BindingItemWorkersBindingKindTailConsumerParam],
// [BindingItemWorkersBindingKindVectorizeParam],
// [BindingItemWorkersBindingKindVersionMetadataParam], [BindingItemParam].
type BindingItemUnionParam interface {
	implementsBindingItemUnionParam()
}

type BindingItemWorkersBindingKindAIParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string]                              `json:"name,required"`
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
	Name param.Field[string]                                           `json:"name,required"`
	Type param.Field[BindingItemWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindAnalyticsEngineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindAnalyticsEngineParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindAssetsParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string]                                  `json:"name,required"`
	Type param.Field[BindingItemWorkersBindingKindAssetsType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindAssetsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindAssetsParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindBrowserRenderingParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string]                                            `json:"name,required"`
	Type param.Field[BindingItemWorkersBindingKindBrowserRenderingType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindBrowserRenderingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindBrowserRenderingParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindD1Param struct {
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string]                              `json:"name,required"`
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
	Namespace param.Field[string]                                             `json:"namespace,required"`
	Type      param.Field[BindingItemWorkersBindingKindDispatchNamespaceType] `json:"type,required"`
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
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string]                                                  `json:"name,required"`
	Type param.Field[BindingItemWorkersBindingKindDurableObjectNamespaceType] `json:"type,required"`
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
	Name param.Field[string]                                      `json:"name,required"`
	Type param.Field[BindingItemWorkersBindingKindHyperdriveType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindHyperdriveParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindHyperdriveParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindJsonParam struct {
	// JSON data to use.
	Json param.Field[interface{}] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string]                                `json:"name,required"`
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
	NamespaceID param.Field[string]                                       `json:"namespace_id,required"`
	Type        param.Field[BindingItemWorkersBindingKindKvNamespaceType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindKvNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindKvNamespaceParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindMtlsCertificateParam struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string]                                           `json:"name,required"`
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
	Text param.Field[string]                                     `json:"text,required"`
	Type param.Field[BindingItemWorkersBindingKindPlainTextType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindPlainTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindPlainTextParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindQueueParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string]                                 `json:"queue_name,required"`
	Type      param.Field[BindingItemWorkersBindingKindQueueType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindQueueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindQueueParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindR2BucketParam struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string]                                    `json:"name,required"`
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
	Text param.Field[string]                                      `json:"text,required"`
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
	Service param.Field[string]                                   `json:"service,required"`
	Type    param.Field[BindingItemWorkersBindingKindServiceType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindServiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindServiceParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindTailConsumerParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service param.Field[string]                                        `json:"service,required"`
	Type    param.Field[BindingItemWorkersBindingKindTailConsumerType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindTailConsumerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindTailConsumerParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindVectorizeParam struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string]                                     `json:"name,required"`
	Type param.Field[BindingItemWorkersBindingKindVectorizeType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindVectorizeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindVectorizeParam) implementsBindingItemUnionParam() {}

type BindingItemWorkersBindingKindVersionMetadataParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string]                                           `json:"name,required"`
	Type param.Field[BindingItemWorkersBindingKindVersionMetadataType] `json:"type,required"`
}

func (r BindingItemWorkersBindingKindVersionMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersBindingKindVersionMetadataParam) implementsBindingItemUnionParam() {}

type UploadSessionObjectParam struct {
	// A manifest ([path]: {hash, size}) map of files to upload. As an example,
	// `/blog/hello-world.html` would be a valid path key.
	Manifest param.Field[map[string]UploadSessionObjectManifestParam] `json:"manifest"`
}

func (r UploadSessionObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UploadSessionObjectManifestParam struct {
	// The hash of the file.
	Hash param.Field[string] `json:"hash"`
	// The size of the file in bytes.
	Size param.Field[int64] `json:"size"`
}

func (r UploadSessionObjectManifestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UploadSessionResponse struct {
	Result UploadSessionResponseResult `json:"result"`
	JSON   uploadSessionResponseJSON   `json:"-"`
	CommonResponseWorkers
}

// uploadSessionResponseJSON contains the JSON metadata for the struct
// [UploadSessionResponse]
type uploadSessionResponseJSON struct {
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
	// Details about a worker uploaded to a Workers for Platforms namespace.
	Result AccountWorkerDispatchNamespaceScriptGetResponseResult `json:"result"`
	JSON   accountWorkerDispatchNamespaceScriptGetResponseJSON   `json:"-"`
	CommonResponseWorkers
}

// accountWorkerDispatchNamespaceScriptGetResponseJSON contains the JSON metadata
// for the struct [AccountWorkerDispatchNamespaceScriptGetResponse]
type accountWorkerDispatchNamespaceScriptGetResponseJSON struct {
	Result      apijson.Field
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

type AccountWorkerDispatchNamespaceScriptGetBindingsResponse struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Result []BindingItem                                               `json:"result"`
	JSON   accountWorkerDispatchNamespaceScriptGetBindingsResponseJSON `json:"-"`
	CommonResponseWorkers
}

// accountWorkerDispatchNamespaceScriptGetBindingsResponseJSON contains the JSON
// metadata for the struct
// [AccountWorkerDispatchNamespaceScriptGetBindingsResponse]
type accountWorkerDispatchNamespaceScriptGetBindingsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptGetBindingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptGetBindingsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDispatchNamespaceScriptUploadResponse struct {
	JSON accountWorkerDispatchNamespaceScriptUploadResponseJSON `json:"-"`
	APIResponseSingleWorkers
}

// accountWorkerDispatchNamespaceScriptUploadResponseJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptUploadResponse]
type accountWorkerDispatchNamespaceScriptUploadResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptUploadResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDispatchNamespaceScriptDeleteParams struct {
	Body interface{} `json:"body,required"`
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
}

func (r AccountWorkerDispatchNamespaceScriptDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
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
}

func (r AccountWorkerDispatchNamespaceScriptUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type AccountWorkerDispatchNamespaceScriptUploadParamsMetadata struct {
	// Configuration for assets within a Worker
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

// Configuration for assets within a Worker
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
	// responses)
	Headers param.Field[string] `json:"_headers"`
	// The contents of a \_redirects file (used to apply redirects or proxy paths ahead
	// of asset serving)
	Redirects param.Field[string] `json:"_redirects"`
	// Determines the redirects and rewrites of requests for HTML content.
	HTMLHandling param.Field[AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandling] `json:"html_handling"`
	// Determines the response when a request does not match a static asset, and there
	// is no Worker script.
	NotFoundHandling param.Field[AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigNotFoundHandling] `json:"not_found_handling"`
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

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations],
// [AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersMultipleStepMigrations].
type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsUnion interface {
	implementsAccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsUnion()
}

// A single set of migrations to apply.
type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations struct {
	MigrationTagConditionsParam
	MigrationStepParam
}

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations) implementsAccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsUnion() {
}

type AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersMultipleStepMigrations struct {
	// Migrations to apply in order.
	Steps param.Field[[]MigrationStepParam] `json:"steps"`
	MigrationTagConditionsParam
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
