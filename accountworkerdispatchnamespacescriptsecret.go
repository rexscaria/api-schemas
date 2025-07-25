// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountWorkerDispatchNamespaceScriptSecretService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerDispatchNamespaceScriptSecretService] method instead.
type AccountWorkerDispatchNamespaceScriptSecretService struct {
	Options []option.RequestOption
}

// NewAccountWorkerDispatchNamespaceScriptSecretService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewAccountWorkerDispatchNamespaceScriptSecretService(opts ...option.RequestOption) (r *AccountWorkerDispatchNamespaceScriptSecretService) {
	r = &AccountWorkerDispatchNamespaceScriptSecretService{}
	r.Options = opts
	return
}

// List secrets bound to a script uploaded to a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptSecretService) List(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptSecretListResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/secrets", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove a secret from a script uploaded to a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptSecretService) Delete(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, secretName string, opts ...option.RequestOption) (res *NullResult, err error) {
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
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/secrets/%s", accountID, dispatchNamespace, scriptName, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add a secret to a script uploaded to a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptSecretService) Add(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, body AccountWorkerDispatchNamespaceScriptSecretAddParams, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptSecretAddResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/secrets", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get a given secret binding (value omitted) on a script uploaded to a Workers for
// Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptSecretService) Get(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, secretName string, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptSecretGetResponse, err error) {
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
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/secrets/%s", accountID, dispatchNamespace, scriptName, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A secret value accessible through a binding.
type Secret struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type SecretType `json:"type,required"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format SecretFormat `json:"format"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// This field can have the runtime type of
	// [[]SecretWorkersBindingKindSecretKeyUsage].
	Usages interface{} `json:"usages"`
	JSON   secretJSON  `json:"-"`
	union  SecretUnion
}

// secretJSON contains the JSON metadata for the struct [Secret]
type secretJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Algorithm   apijson.Field
	Format      apijson.Field
	KeyJwk      apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r secretJSON) RawJSON() string {
	return r.raw
}

func (r *Secret) UnmarshalJSON(data []byte) (err error) {
	*r = Secret{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SecretUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [SecretWorkersBindingKindSecretText],
// [SecretWorkersBindingKindSecretKey].
func (r Secret) AsUnion() SecretUnion {
	return r.union
}

// A secret value accessible through a binding.
//
// Union satisfied by [SecretWorkersBindingKindSecretText] or
// [SecretWorkersBindingKindSecretKey].
type SecretUnion interface {
	implementsSecret()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SecretUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SecretWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SecretWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
	)
}

type SecretWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type SecretWorkersBindingKindSecretTextType `json:"type,required"`
	JSON secretWorkersBindingKindSecretTextJSON `json:"-"`
}

// secretWorkersBindingKindSecretTextJSON contains the JSON metadata for the struct
// [SecretWorkersBindingKindSecretText]
type secretWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r SecretWorkersBindingKindSecretText) implementsSecret() {}

// The kind of resource that the binding provides.
type SecretWorkersBindingKindSecretTextType string

const (
	SecretWorkersBindingKindSecretTextTypeSecretText SecretWorkersBindingKindSecretTextType = "secret_text"
)

func (r SecretWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case SecretWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type SecretWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format SecretWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type SecretWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []SecretWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   secretWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// secretWorkersBindingKindSecretKeyJSON contains the JSON metadata for the struct
// [SecretWorkersBindingKindSecretKey]
type secretWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r SecretWorkersBindingKindSecretKey) implementsSecret() {}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type SecretWorkersBindingKindSecretKeyFormat string

const (
	SecretWorkersBindingKindSecretKeyFormatRaw   SecretWorkersBindingKindSecretKeyFormat = "raw"
	SecretWorkersBindingKindSecretKeyFormatPkcs8 SecretWorkersBindingKindSecretKeyFormat = "pkcs8"
	SecretWorkersBindingKindSecretKeyFormatSpki  SecretWorkersBindingKindSecretKeyFormat = "spki"
	SecretWorkersBindingKindSecretKeyFormatJwk   SecretWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r SecretWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case SecretWorkersBindingKindSecretKeyFormatRaw, SecretWorkersBindingKindSecretKeyFormatPkcs8, SecretWorkersBindingKindSecretKeyFormatSpki, SecretWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type SecretWorkersBindingKindSecretKeyType string

const (
	SecretWorkersBindingKindSecretKeyTypeSecretKey SecretWorkersBindingKindSecretKeyType = "secret_key"
)

func (r SecretWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case SecretWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type SecretWorkersBindingKindSecretKeyUsage string

const (
	SecretWorkersBindingKindSecretKeyUsageEncrypt    SecretWorkersBindingKindSecretKeyUsage = "encrypt"
	SecretWorkersBindingKindSecretKeyUsageDecrypt    SecretWorkersBindingKindSecretKeyUsage = "decrypt"
	SecretWorkersBindingKindSecretKeyUsageSign       SecretWorkersBindingKindSecretKeyUsage = "sign"
	SecretWorkersBindingKindSecretKeyUsageVerify     SecretWorkersBindingKindSecretKeyUsage = "verify"
	SecretWorkersBindingKindSecretKeyUsageDeriveKey  SecretWorkersBindingKindSecretKeyUsage = "deriveKey"
	SecretWorkersBindingKindSecretKeyUsageDeriveBits SecretWorkersBindingKindSecretKeyUsage = "deriveBits"
	SecretWorkersBindingKindSecretKeyUsageWrapKey    SecretWorkersBindingKindSecretKeyUsage = "wrapKey"
	SecretWorkersBindingKindSecretKeyUsageUnwrapKey  SecretWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r SecretWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case SecretWorkersBindingKindSecretKeyUsageEncrypt, SecretWorkersBindingKindSecretKeyUsageDecrypt, SecretWorkersBindingKindSecretKeyUsageSign, SecretWorkersBindingKindSecretKeyUsageVerify, SecretWorkersBindingKindSecretKeyUsageDeriveKey, SecretWorkersBindingKindSecretKeyUsageDeriveBits, SecretWorkersBindingKindSecretKeyUsageWrapKey, SecretWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type SecretType string

const (
	SecretTypeSecretText SecretType = "secret_text"
	SecretTypeSecretKey  SecretType = "secret_key"
)

func (r SecretType) IsKnown() bool {
	switch r {
	case SecretTypeSecretText, SecretTypeSecretKey:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type SecretFormat string

const (
	SecretFormatRaw   SecretFormat = "raw"
	SecretFormatPkcs8 SecretFormat = "pkcs8"
	SecretFormatSpki  SecretFormat = "spki"
	SecretFormatJwk   SecretFormat = "jwk"
)

func (r SecretFormat) IsKnown() bool {
	switch r {
	case SecretFormatRaw, SecretFormatPkcs8, SecretFormatSpki, SecretFormatJwk:
		return true
	}
	return false
}

// A secret value accessible through a binding.
type SecretParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type      param.Field[SecretType]  `json:"type,required"`
	Algorithm param.Field[interface{}] `json:"algorithm"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[SecretFormat] `json:"format"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string]      `json:"key_base64"`
	KeyJwk    param.Field[interface{}] `json:"key_jwk"`
	// The secret value to use.
	Text   param.Field[string]      `json:"text"`
	Usages param.Field[interface{}] `json:"usages"`
}

func (r SecretParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SecretParam) implementsSecretUnionParam() {}

// A secret value accessible through a binding.
//
// Satisfied by [SecretWorkersBindingKindSecretTextParam],
// [SecretWorkersBindingKindSecretKeyParam], [SecretParam].
type SecretUnionParam interface {
	implementsSecretUnionParam()
}

type SecretWorkersBindingKindSecretTextParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[SecretWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r SecretWorkersBindingKindSecretTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SecretWorkersBindingKindSecretTextParam) implementsSecretUnionParam() {}

type SecretWorkersBindingKindSecretKeyParam struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[SecretWorkersBindingKindSecretKeyFormat] `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[SecretWorkersBindingKindSecretKeyType] `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]SecretWorkersBindingKindSecretKeyUsage] `json:"usages,required"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string] `json:"key_base64"`
	// Key data in
	// [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key)
	// format. Required if `format` is "jwk".
	KeyJwk param.Field[interface{}] `json:"key_jwk"`
}

func (r SecretWorkersBindingKindSecretKeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SecretWorkersBindingKindSecretKeyParam) implementsSecretUnionParam() {}

type AccountWorkerDispatchNamespaceScriptSecretListResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	Result   []Secret          `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerDispatchNamespaceScriptSecretListResponseSuccess `json:"success,required"`
	JSON    accountWorkerDispatchNamespaceScriptSecretListResponseJSON    `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSecretListResponseJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptSecretListResponse]
type accountWorkerDispatchNamespaceScriptSecretListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSecretListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptSecretListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerDispatchNamespaceScriptSecretListResponseSuccess bool

const (
	AccountWorkerDispatchNamespaceScriptSecretListResponseSuccessTrue AccountWorkerDispatchNamespaceScriptSecretListResponseSuccess = true
)

func (r AccountWorkerDispatchNamespaceScriptSecretListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerDispatchNamespaceScriptSecretListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerDispatchNamespaceScriptSecretAddResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// A secret value accessible through a binding.
	Result Secret `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerDispatchNamespaceScriptSecretAddResponseSuccess `json:"success,required"`
	JSON    accountWorkerDispatchNamespaceScriptSecretAddResponseJSON    `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSecretAddResponseJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptSecretAddResponse]
type accountWorkerDispatchNamespaceScriptSecretAddResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSecretAddResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptSecretAddResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerDispatchNamespaceScriptSecretAddResponseSuccess bool

const (
	AccountWorkerDispatchNamespaceScriptSecretAddResponseSuccessTrue AccountWorkerDispatchNamespaceScriptSecretAddResponseSuccess = true
)

func (r AccountWorkerDispatchNamespaceScriptSecretAddResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerDispatchNamespaceScriptSecretAddResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerDispatchNamespaceScriptSecretGetResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// A secret value accessible through a binding.
	Result Secret `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerDispatchNamespaceScriptSecretGetResponseSuccess `json:"success,required"`
	JSON    accountWorkerDispatchNamespaceScriptSecretGetResponseJSON    `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSecretGetResponseJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptSecretGetResponse]
type accountWorkerDispatchNamespaceScriptSecretGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSecretGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptSecretGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerDispatchNamespaceScriptSecretGetResponseSuccess bool

const (
	AccountWorkerDispatchNamespaceScriptSecretGetResponseSuccessTrue AccountWorkerDispatchNamespaceScriptSecretGetResponseSuccess = true
)

func (r AccountWorkerDispatchNamespaceScriptSecretGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerDispatchNamespaceScriptSecretGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerDispatchNamespaceScriptSecretAddParams struct {
	// A secret value accessible through a binding.
	Secret SecretUnionParam `json:"secret,required"`
}

func (r AccountWorkerDispatchNamespaceScriptSecretAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Secret)
}
