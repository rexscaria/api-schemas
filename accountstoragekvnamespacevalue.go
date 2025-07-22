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

// AccountStorageKvNamespaceValueService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountStorageKvNamespaceValueService] method instead.
type AccountStorageKvNamespaceValueService struct {
	Options []option.RequestOption
}

// NewAccountStorageKvNamespaceValueService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountStorageKvNamespaceValueService(opts ...option.RequestOption) (r *AccountStorageKvNamespaceValueService) {
	r = &AccountStorageKvNamespaceValueService{}
	r.Options = opts
	return
}

// Returns the value associated with the given key in the given namespace. Use
// URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key
// name. If the KV-pair is set to expire at some point, the expiration time as
// measured in seconds since the UNIX epoch will be returned in the `expiration`
// response header.
func (r *AccountStorageKvNamespaceValueService) Get(ctx context.Context, accountID string, namespaceID string, keyName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	if keyName == "" {
		err = errors.New("missing required key_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", accountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove a KV pair from the namespace. Use URL-encoding to use special characters
// (for example, `:`, `!`, `%`) in the key name.
func (r *AccountStorageKvNamespaceValueService) Delete(ctx context.Context, accountID string, namespaceID string, keyName string, body AccountStorageKvNamespaceValueDeleteParams, opts ...option.RequestOption) (res *APIResponseCommonNoResult, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	if keyName == "" {
		err = errors.New("missing required key_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", accountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Write a value identified by a key. Use URL-encoding to use special characters
// (for example, `:`, `!`, `%`) in the key name. Body should be the value to be
// stored. If JSON metadata to be associated with the key/value pair is needed, use
// `multipart/form-data` content type for your PUT request (see dropdown below in
// `REQUEST BODY SCHEMA`). Existing values, expirations, and metadata will be
// overwritten. If neither `expiration` nor `expiration_ttl` is specified, the
// key-value pair will never expire. If both are set, `expiration_ttl` is used and
// `expiration` is ignored.
func (r *AccountStorageKvNamespaceValueService) Write(ctx context.Context, accountID string, namespaceID string, keyName string, params AccountStorageKvNamespaceValueWriteParams, opts ...option.RequestOption) (res *APIResponseCommonNoResult, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	if keyName == "" {
		err = errors.New("missing required key_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/values/%s", accountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type AccountStorageKvNamespaceValueDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountStorageKvNamespaceValueDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountStorageKvNamespaceValueWriteParams struct {
	// Arbitrary JSON to be associated with a key/value pair.
	Metadata param.Field[string] `json:"metadata,required"`
	// A byte sequence to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value,required"`
	// The time, measured in number of seconds since the UNIX epoch, at which the key
	// should expire.
	Expiration param.Field[float64] `query:"expiration"`
	// The number of seconds for which the key should be visible before it expires. At
	// least 60.
	ExpirationTtl param.Field[float64] `query:"expiration_ttl"`
}

func (r AccountStorageKvNamespaceValueWriteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountStorageKvNamespaceValueWriteParams]'s query
// parameters as `url.Values`.
func (r AccountStorageKvNamespaceValueWriteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
