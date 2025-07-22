// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountR2Service contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountR2Service] method instead.
type AccountR2Service struct {
	Options []option.RequestOption
	Buckets *AccountR2BucketService
}

// NewAccountR2Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountR2Service(opts ...option.RequestOption) (r *AccountR2Service) {
	r = &AccountR2Service{}
	r.Options = opts
	r.Buckets = NewAccountR2BucketService(opts...)
	return
}

// Creates temporary access credentials on a bucket that can be optionally scoped
// to prefixes or objects.
func (r *AccountR2Service) NewTempAccessCredentials(ctx context.Context, accountID string, body AccountR2NewTempAccessCredentialsParams, opts ...option.RequestOption) (res *AccountR2NewTempAccessCredentialsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/temp-access-credentials", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Storage/Object Count Metrics across all buckets in your account. Note that
// Account-Level Metrics may not immediately reflect the latest data.
func (r *AccountR2Service) GetMetrics(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountR2GetMetricsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/metrics", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Metrics based on what state they are in(uploaded or published)
type R2ClassBasedMetrics struct {
	// Metrics on number of objects/amount of storage used
	Published R2ObjectSizeMetrics `json:"published"`
	// Metrics on number of objects/amount of storage used
	Uploaded R2ObjectSizeMetrics     `json:"uploaded"`
	JSON     r2ClassBasedMetricsJSON `json:"-"`
}

// r2ClassBasedMetricsJSON contains the JSON metadata for the struct
// [R2ClassBasedMetrics]
type r2ClassBasedMetricsJSON struct {
	Published   apijson.Field
	Uploaded    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2ClassBasedMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2ClassBasedMetricsJSON) RawJSON() string {
	return r.raw
}

// Metrics on number of objects/amount of storage used
type R2ObjectSizeMetrics struct {
	// Amount of
	MetadataSize float64 `json:"metadataSize"`
	// Number of objects stored
	Objects float64 `json:"objects"`
	// Amount of storage used by object data
	PayloadSize float64                 `json:"payloadSize"`
	JSON        r2ObjectSizeMetricsJSON `json:"-"`
}

// r2ObjectSizeMetricsJSON contains the JSON metadata for the struct
// [R2ObjectSizeMetrics]
type r2ObjectSizeMetricsJSON struct {
	MetadataSize apijson.Field
	Objects      apijson.Field
	PayloadSize  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *R2ObjectSizeMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2ObjectSizeMetricsJSON) RawJSON() string {
	return r.raw
}

type AccountR2NewTempAccessCredentialsResponse struct {
	Result AccountR2NewTempAccessCredentialsResponseResult `json:"result"`
	JSON   accountR2NewTempAccessCredentialsResponseJSON   `json:"-"`
	R2V4Response
}

// accountR2NewTempAccessCredentialsResponseJSON contains the JSON metadata for the
// struct [AccountR2NewTempAccessCredentialsResponse]
type accountR2NewTempAccessCredentialsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2NewTempAccessCredentialsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2NewTempAccessCredentialsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2NewTempAccessCredentialsResponseResult struct {
	// ID for new access key
	AccessKeyID string `json:"accessKeyId"`
	// Secret access key
	SecretAccessKey string `json:"secretAccessKey"`
	// Security token
	SessionToken string                                              `json:"sessionToken"`
	JSON         accountR2NewTempAccessCredentialsResponseResultJSON `json:"-"`
}

// accountR2NewTempAccessCredentialsResponseResultJSON contains the JSON metadata
// for the struct [AccountR2NewTempAccessCredentialsResponseResult]
type accountR2NewTempAccessCredentialsResponseResultJSON struct {
	AccessKeyID     apijson.Field
	SecretAccessKey apijson.Field
	SessionToken    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountR2NewTempAccessCredentialsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2NewTempAccessCredentialsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountR2GetMetricsResponse struct {
	// Metrics based on the class they belong to
	Result AccountR2GetMetricsResponseResult `json:"result"`
	JSON   accountR2GetMetricsResponseJSON   `json:"-"`
	R2V4Response
}

// accountR2GetMetricsResponseJSON contains the JSON metadata for the struct
// [AccountR2GetMetricsResponse]
type accountR2GetMetricsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2GetMetricsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2GetMetricsResponseJSON) RawJSON() string {
	return r.raw
}

// Metrics based on the class they belong to
type AccountR2GetMetricsResponseResult struct {
	// Metrics based on what state they are in(uploaded or published)
	InfrequentAccess R2ClassBasedMetrics `json:"infrequentAccess"`
	// Metrics based on what state they are in(uploaded or published)
	Standard R2ClassBasedMetrics                   `json:"standard"`
	JSON     accountR2GetMetricsResponseResultJSON `json:"-"`
}

// accountR2GetMetricsResponseResultJSON contains the JSON metadata for the struct
// [AccountR2GetMetricsResponseResult]
type accountR2GetMetricsResponseResultJSON struct {
	InfrequentAccess apijson.Field
	Standard         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountR2GetMetricsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2GetMetricsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountR2NewTempAccessCredentialsParams struct {
	// Name of the R2 bucket
	Bucket param.Field[string] `json:"bucket,required"`
	// The parent access key id to use for signing
	ParentAccessKeyID param.Field[string] `json:"parentAccessKeyId,required"`
	// Permissions allowed on the credentials
	Permission param.Field[AccountR2NewTempAccessCredentialsParamsPermission] `json:"permission,required"`
	// How long the credentials will live for in seconds
	TtlSeconds param.Field[float64] `json:"ttlSeconds,required"`
	// Optional object paths to scope the credentials to
	Objects param.Field[[]string] `json:"objects"`
	// Optional prefix paths to scope the credentials to
	Prefixes param.Field[[]string] `json:"prefixes"`
}

func (r AccountR2NewTempAccessCredentialsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Permissions allowed on the credentials
type AccountR2NewTempAccessCredentialsParamsPermission string

const (
	AccountR2NewTempAccessCredentialsParamsPermissionAdminReadWrite  AccountR2NewTempAccessCredentialsParamsPermission = "admin-read-write"
	AccountR2NewTempAccessCredentialsParamsPermissionAdminReadOnly   AccountR2NewTempAccessCredentialsParamsPermission = "admin-read-only"
	AccountR2NewTempAccessCredentialsParamsPermissionObjectReadWrite AccountR2NewTempAccessCredentialsParamsPermission = "object-read-write"
	AccountR2NewTempAccessCredentialsParamsPermissionObjectReadOnly  AccountR2NewTempAccessCredentialsParamsPermission = "object-read-only"
)

func (r AccountR2NewTempAccessCredentialsParamsPermission) IsKnown() bool {
	switch r {
	case AccountR2NewTempAccessCredentialsParamsPermissionAdminReadWrite, AccountR2NewTempAccessCredentialsParamsPermissionAdminReadOnly, AccountR2NewTempAccessCredentialsParamsPermissionObjectReadWrite, AccountR2NewTempAccessCredentialsParamsPermissionObjectReadOnly:
		return true
	}
	return false
}
