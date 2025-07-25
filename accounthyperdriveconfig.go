// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountHyperdriveConfigService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountHyperdriveConfigService] method instead.
type AccountHyperdriveConfigService struct {
	Options []option.RequestOption
}

// NewAccountHyperdriveConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountHyperdriveConfigService(opts ...option.RequestOption) (r *AccountHyperdriveConfigService) {
	r = &AccountHyperdriveConfigService{}
	r.Options = opts
	return
}

// Creates and returns a new Hyperdrive configuration.
func (r *AccountHyperdriveConfigService) New(ctx context.Context, accountID string, body AccountHyperdriveConfigNewParams, opts ...option.RequestOption) (res *AccountHyperdriveConfigNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the specified Hyperdrive configuration.
func (r *AccountHyperdriveConfigService) Get(ctx context.Context, accountID string, hyperdriveID string, opts ...option.RequestOption) (res *AccountHyperdriveConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if hyperdriveID == "" {
		err = errors.New("missing required hyperdrive_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", accountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates and returns the specified Hyperdrive configuration.
func (r *AccountHyperdriveConfigService) Update(ctx context.Context, accountID string, hyperdriveID string, body AccountHyperdriveConfigUpdateParams, opts ...option.RequestOption) (res *AccountHyperdriveConfigUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if hyperdriveID == "" {
		err = errors.New("missing required hyperdrive_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", accountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns a list of Hyperdrives.
func (r *AccountHyperdriveConfigService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountHyperdriveConfigListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes the specified Hyperdrive.
func (r *AccountHyperdriveConfigService) Delete(ctx context.Context, accountID string, hyperdriveID string, opts ...option.RequestOption) (res *AccountHyperdriveConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if hyperdriveID == "" {
		err = errors.New("missing required hyperdrive_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", accountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Patches and returns the specified Hyperdrive configuration. Custom caching
// settings are not kept if caching is disabled.
func (r *AccountHyperdriveConfigService) Patch(ctx context.Context, accountID string, hyperdriveID string, body AccountHyperdriveConfigPatchParams, opts ...option.RequestOption) (res *AccountHyperdriveConfigPatchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if hyperdriveID == "" {
		err = errors.New("missing required hyperdrive_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", accountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type HyperdriveHyperdriveCaching struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled bool `json:"disabled"`
	// Specify the maximum duration items should persist in the cache. Not returned if
	// set to the default (60).
	MaxAge int64 `json:"max_age"`
	// Specify the number of seconds the cache may serve a stale response. Omitted if
	// set to the default (15).
	StaleWhileRevalidate int64                           `json:"stale_while_revalidate"`
	JSON                 hyperdriveHyperdriveCachingJSON `json:"-"`
	union                HyperdriveHyperdriveCachingUnion
}

// hyperdriveHyperdriveCachingJSON contains the JSON metadata for the struct
// [HyperdriveHyperdriveCaching]
type hyperdriveHyperdriveCachingJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r hyperdriveHyperdriveCachingJSON) RawJSON() string {
	return r.raw
}

func (r *HyperdriveHyperdriveCaching) UnmarshalJSON(data []byte) (err error) {
	*r = HyperdriveHyperdriveCaching{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [HyperdriveHyperdriveCachingUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [HyperdriveHyperdriveCachingCommon],
// [HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabled].
func (r HyperdriveHyperdriveCaching) AsUnion() HyperdriveHyperdriveCachingUnion {
	return r.union
}

// Union satisfied by [HyperdriveHyperdriveCachingCommon] or
// [HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabled].
type HyperdriveHyperdriveCachingUnion interface {
	implementsHyperdriveHyperdriveCaching()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HyperdriveHyperdriveCachingUnion)(nil)).Elem(),
		"disabled",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HyperdriveHyperdriveCachingCommon{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabled{}),
		},
	)
}

type HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabled struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled bool `json:"disabled"`
	// Specify the maximum duration items should persist in the cache. Not returned if
	// set to the default (60).
	MaxAge int64 `json:"max_age"`
	// Specify the number of seconds the cache may serve a stale response. Omitted if
	// set to the default (15).
	StaleWhileRevalidate int64                                                             `json:"stale_while_revalidate"`
	JSON                 hyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabledJSON `json:"-"`
}

// hyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabledJSON contains the
// JSON metadata for the struct
// [HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabled]
type hyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabledJSON struct {
	Disabled             apijson.Field
	MaxAge               apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabledJSON) RawJSON() string {
	return r.raw
}

func (r HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabled) implementsHyperdriveHyperdriveCaching() {
}

type HyperdriveHyperdriveCachingParam struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled param.Field[bool] `json:"disabled"`
	// Specify the maximum duration items should persist in the cache. Not returned if
	// set to the default (60).
	MaxAge param.Field[int64] `json:"max_age"`
	// Specify the number of seconds the cache may serve a stale response. Omitted if
	// set to the default (15).
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r HyperdriveHyperdriveCachingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveHyperdriveCachingParam) implementsHyperdriveHyperdriveCachingUnionParam() {}

// Satisfied by [HyperdriveHyperdriveCachingCommonParam],
// [HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabledParam],
// [HyperdriveHyperdriveCachingParam].
type HyperdriveHyperdriveCachingUnionParam interface {
	implementsHyperdriveHyperdriveCachingUnionParam()
}

type HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabledParam struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled param.Field[bool] `json:"disabled"`
	// Specify the maximum duration items should persist in the cache. Not returned if
	// set to the default (60).
	MaxAge param.Field[int64] `json:"max_age"`
	// Specify the number of seconds the cache may serve a stale response. Omitted if
	// set to the default (15).
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
}

func (r HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabledParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabledParam) implementsHyperdriveHyperdriveCachingUnionParam() {
}

type HyperdriveHyperdriveCachingCommon struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled bool                                  `json:"disabled"`
	JSON     hyperdriveHyperdriveCachingCommonJSON `json:"-"`
}

// hyperdriveHyperdriveCachingCommonJSON contains the JSON metadata for the struct
// [HyperdriveHyperdriveCachingCommon]
type hyperdriveHyperdriveCachingCommonJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveHyperdriveCachingCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveHyperdriveCachingCommonJSON) RawJSON() string {
	return r.raw
}

func (r HyperdriveHyperdriveCachingCommon) implementsHyperdriveHyperdriveCaching() {}

type HyperdriveHyperdriveCachingCommonParam struct {
	// Set to true to disable caching of SQL responses. Default is false.
	Disabled param.Field[bool] `json:"disabled"`
}

func (r HyperdriveHyperdriveCachingCommonParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveHyperdriveCachingCommonParam) implementsHyperdriveHyperdriveCachingUnionParam() {}

type HyperdriveHyperdriveConfigParam struct {
	Name    param.Field[string]                                     `json:"name,required"`
	Origin  param.Field[HyperdriveHyperdriveConfigOriginUnionParam] `json:"origin,required"`
	Caching param.Field[HyperdriveHyperdriveCachingUnionParam]      `json:"caching"`
	Mtls    param.Field[HyperdriveHyperdriveConfigMtlsParam]        `json:"mtls"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to
	// the origin database.
	OriginConnectionLimit param.Field[int64] `json:"origin_connection_limit"`
}

func (r HyperdriveHyperdriveConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveHyperdriveConfigOriginParam struct {
	// Set the name of your origin database.
	Database param.Field[string] `json:"database,required"`
	// Defines the host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveHyperdriveConfigOriginScheme] `json:"scheme,required"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user,required"`
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID param.Field[string] `json:"access_client_id"`
	// Defines the Client Secret of the Access Token to use when connecting to the
	// origin database. The API never returns this write-only value.
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
	// Defines the port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port"`
}

func (r HyperdriveHyperdriveConfigOriginParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveHyperdriveConfigOriginParam) implementsHyperdriveHyperdriveConfigOriginUnionParam() {
}

// Satisfied by [HyperdriveHyperdriveConfigOriginPublicDatabaseParam],
// [HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelParam],
// [HyperdriveHyperdriveConfigOriginParam].
type HyperdriveHyperdriveConfigOriginUnionParam interface {
	implementsHyperdriveHyperdriveConfigOriginUnionParam()
}

type HyperdriveHyperdriveConfigOriginPublicDatabaseParam struct {
	// Set the name of your origin database.
	Database param.Field[string] `json:"database,required"`
	// Defines the host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password,required"`
	// Defines the port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveHyperdriveConfigOriginPublicDatabaseScheme] `json:"scheme,required"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user,required"`
}

func (r HyperdriveHyperdriveConfigOriginPublicDatabaseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveHyperdriveConfigOriginPublicDatabaseParam) implementsHyperdriveHyperdriveConfigOriginUnionParam() {
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveHyperdriveConfigOriginPublicDatabaseScheme string

const (
	HyperdriveHyperdriveConfigOriginPublicDatabaseSchemePostgres   HyperdriveHyperdriveConfigOriginPublicDatabaseScheme = "postgres"
	HyperdriveHyperdriveConfigOriginPublicDatabaseSchemePostgresql HyperdriveHyperdriveConfigOriginPublicDatabaseScheme = "postgresql"
	HyperdriveHyperdriveConfigOriginPublicDatabaseSchemeMysql      HyperdriveHyperdriveConfigOriginPublicDatabaseScheme = "mysql"
)

func (r HyperdriveHyperdriveConfigOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case HyperdriveHyperdriveConfigOriginPublicDatabaseSchemePostgres, HyperdriveHyperdriveConfigOriginPublicDatabaseSchemePostgresql, HyperdriveHyperdriveConfigOriginPublicDatabaseSchemeMysql:
		return true
	}
	return false
}

type HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelParam struct {
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID param.Field[string] `json:"access_client_id,required"`
	// Defines the Client Secret of the Access Token to use when connecting to the
	// origin database. The API never returns this write-only value.
	AccessClientSecret param.Field[string] `json:"access_client_secret,required"`
	// Set the name of your origin database.
	Database param.Field[string] `json:"database,required"`
	// Defines the host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme] `json:"scheme,required"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user,required"`
}

func (r HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelParam) implementsHyperdriveHyperdriveConfigOriginUnionParam() {
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
	HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql      HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "mysql"
)

func (r HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql, HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveHyperdriveConfigOriginScheme string

const (
	HyperdriveHyperdriveConfigOriginSchemePostgres   HyperdriveHyperdriveConfigOriginScheme = "postgres"
	HyperdriveHyperdriveConfigOriginSchemePostgresql HyperdriveHyperdriveConfigOriginScheme = "postgresql"
	HyperdriveHyperdriveConfigOriginSchemeMysql      HyperdriveHyperdriveConfigOriginScheme = "mysql"
)

func (r HyperdriveHyperdriveConfigOriginScheme) IsKnown() bool {
	switch r {
	case HyperdriveHyperdriveConfigOriginSchemePostgres, HyperdriveHyperdriveConfigOriginSchemePostgresql, HyperdriveHyperdriveConfigOriginSchemeMysql:
		return true
	}
	return false
}

type HyperdriveHyperdriveConfigMtlsParam struct {
	// Define CA certificate ID obtained after uploading CA cert.
	CaCertificateID param.Field[string] `json:"ca_certificate_id"`
	// Define mTLS certificate ID obtained after uploading client cert.
	MtlsCertificateID param.Field[string] `json:"mtls_certificate_id"`
	// Set SSL mode to 'require', 'verify-ca', or 'verify-full' to verify the CA.
	Sslmode param.Field[string] `json:"sslmode"`
}

func (r HyperdriveHyperdriveConfigMtlsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveHyperdriveConfigResponse struct {
	Caching HyperdriveHyperdriveCaching `json:"caching,required"`
	// Define configurations using a unique string identifier.
	ID string `json:"id"`
	// Defines the creation time of the Hyperdrive configuration.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Defines the last modified time of the Hyperdrive configuration.
	ModifiedOn time.Time                                `json:"modified_on" format:"date-time"`
	Mtls       HyperdriveHyperdriveConfigResponseMtls   `json:"mtls"`
	Name       string                                   `json:"name"`
	Origin     HyperdriveHyperdriveConfigResponseOrigin `json:"origin"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to
	// the origin database.
	OriginConnectionLimit int64                                  `json:"origin_connection_limit"`
	JSON                  hyperdriveHyperdriveConfigResponseJSON `json:"-"`
}

// hyperdriveHyperdriveConfigResponseJSON contains the JSON metadata for the struct
// [HyperdriveHyperdriveConfigResponse]
type hyperdriveHyperdriveConfigResponseJSON struct {
	Caching               apijson.Field
	ID                    apijson.Field
	CreatedOn             apijson.Field
	ModifiedOn            apijson.Field
	Mtls                  apijson.Field
	Name                  apijson.Field
	Origin                apijson.Field
	OriginConnectionLimit apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *HyperdriveHyperdriveConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveHyperdriveConfigResponseJSON) RawJSON() string {
	return r.raw
}

type HyperdriveHyperdriveConfigResponseMtls struct {
	// Define CA certificate ID obtained after uploading CA cert.
	CaCertificateID string `json:"ca_certificate_id"`
	// Define mTLS certificate ID obtained after uploading client cert.
	MtlsCertificateID string `json:"mtls_certificate_id"`
	// Set SSL mode to 'require', 'verify-ca', or 'verify-full' to verify the CA.
	Sslmode string                                     `json:"sslmode"`
	JSON    hyperdriveHyperdriveConfigResponseMtlsJSON `json:"-"`
}

// hyperdriveHyperdriveConfigResponseMtlsJSON contains the JSON metadata for the
// struct [HyperdriveHyperdriveConfigResponseMtls]
type hyperdriveHyperdriveConfigResponseMtlsJSON struct {
	CaCertificateID   apijson.Field
	MtlsCertificateID apijson.Field
	Sslmode           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *HyperdriveHyperdriveConfigResponseMtls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveHyperdriveConfigResponseMtlsJSON) RawJSON() string {
	return r.raw
}

type HyperdriveHyperdriveConfigResponseOrigin struct {
	// Set the name of your origin database.
	Database string `json:"database,required"`
	// Defines the host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveHyperdriveConfigResponseOriginScheme `json:"scheme,required"`
	// Set the user of your origin database.
	User string `json:"user,required"`
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID string `json:"access_client_id"`
	// Defines the port (default: 5432 for Postgres) of your origin database.
	Port  int64                                        `json:"port"`
	JSON  hyperdriveHyperdriveConfigResponseOriginJSON `json:"-"`
	union HyperdriveHyperdriveConfigResponseOriginUnion
}

// hyperdriveHyperdriveConfigResponseOriginJSON contains the JSON metadata for the
// struct [HyperdriveHyperdriveConfigResponseOrigin]
type hyperdriveHyperdriveConfigResponseOriginJSON struct {
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	AccessClientID apijson.Field
	Port           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r hyperdriveHyperdriveConfigResponseOriginJSON) RawJSON() string {
	return r.raw
}

func (r *HyperdriveHyperdriveConfigResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	*r = HyperdriveHyperdriveConfigResponseOrigin{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [HyperdriveHyperdriveConfigResponseOriginUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [HyperdriveHyperdriveConfigResponseOriginPublicDatabase],
// [HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel].
func (r HyperdriveHyperdriveConfigResponseOrigin) AsUnion() HyperdriveHyperdriveConfigResponseOriginUnion {
	return r.union
}

// Union satisfied by [HyperdriveHyperdriveConfigResponseOriginPublicDatabase] or
// [HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel].
type HyperdriveHyperdriveConfigResponseOriginUnion interface {
	implementsHyperdriveHyperdriveConfigResponseOrigin()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HyperdriveHyperdriveConfigResponseOriginUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HyperdriveHyperdriveConfigResponseOriginPublicDatabase{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel{}),
		},
	)
}

type HyperdriveHyperdriveConfigResponseOriginPublicDatabase struct {
	// Set the name of your origin database.
	Database string `json:"database,required"`
	// Defines the host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Defines the port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveHyperdriveConfigResponseOriginPublicDatabaseScheme `json:"scheme,required"`
	// Set the user of your origin database.
	User string                                                     `json:"user,required"`
	JSON hyperdriveHyperdriveConfigResponseOriginPublicDatabaseJSON `json:"-"`
}

// hyperdriveHyperdriveConfigResponseOriginPublicDatabaseJSON contains the JSON
// metadata for the struct [HyperdriveHyperdriveConfigResponseOriginPublicDatabase]
type hyperdriveHyperdriveConfigResponseOriginPublicDatabaseJSON struct {
	Database    apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveHyperdriveConfigResponseOriginPublicDatabase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveHyperdriveConfigResponseOriginPublicDatabaseJSON) RawJSON() string {
	return r.raw
}

func (r HyperdriveHyperdriveConfigResponseOriginPublicDatabase) implementsHyperdriveHyperdriveConfigResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveHyperdriveConfigResponseOriginPublicDatabaseScheme string

const (
	HyperdriveHyperdriveConfigResponseOriginPublicDatabaseSchemePostgres   HyperdriveHyperdriveConfigResponseOriginPublicDatabaseScheme = "postgres"
	HyperdriveHyperdriveConfigResponseOriginPublicDatabaseSchemePostgresql HyperdriveHyperdriveConfigResponseOriginPublicDatabaseScheme = "postgresql"
	HyperdriveHyperdriveConfigResponseOriginPublicDatabaseSchemeMysql      HyperdriveHyperdriveConfigResponseOriginPublicDatabaseScheme = "mysql"
)

func (r HyperdriveHyperdriveConfigResponseOriginPublicDatabaseScheme) IsKnown() bool {
	switch r {
	case HyperdriveHyperdriveConfigResponseOriginPublicDatabaseSchemePostgres, HyperdriveHyperdriveConfigResponseOriginPublicDatabaseSchemePostgresql, HyperdriveHyperdriveConfigResponseOriginPublicDatabaseSchemeMysql:
		return true
	}
	return false
}

type HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID string `json:"access_client_id,required"`
	// Set the name of your origin database.
	Database string `json:"database,required"`
	// Defines the host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme `json:"scheme,required"`
	// Set the user of your origin database.
	User string                                                                                    `json:"user,required"`
	JSON hyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON `json:"-"`
}

// hyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON
// contains the JSON metadata for the struct
// [HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel]
type hyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON struct {
	AccessClientID apijson.Field
	Database       apijson.Field
	Host           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON) RawJSON() string {
	return r.raw
}

func (r HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel) implementsHyperdriveHyperdriveConfigResponseOrigin() {
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme string

const (
	HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres   HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgres"
	HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "postgresql"
	HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql      HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme = "mysql"
)

func (r HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelScheme) IsKnown() bool {
	switch r {
	case HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgres, HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemePostgresql, HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelSchemeMysql:
		return true
	}
	return false
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveHyperdriveConfigResponseOriginScheme string

const (
	HyperdriveHyperdriveConfigResponseOriginSchemePostgres   HyperdriveHyperdriveConfigResponseOriginScheme = "postgres"
	HyperdriveHyperdriveConfigResponseOriginSchemePostgresql HyperdriveHyperdriveConfigResponseOriginScheme = "postgresql"
	HyperdriveHyperdriveConfigResponseOriginSchemeMysql      HyperdriveHyperdriveConfigResponseOriginScheme = "mysql"
)

func (r HyperdriveHyperdriveConfigResponseOriginScheme) IsKnown() bool {
	switch r {
	case HyperdriveHyperdriveConfigResponseOriginSchemePostgres, HyperdriveHyperdriveConfigResponseOriginSchemePostgresql, HyperdriveHyperdriveConfigResponseOriginSchemeMysql:
		return true
	}
	return false
}

type HyperdriveHyperdriveDatabaseParam struct {
	// Set the name of your origin database.
	Database param.Field[string] `json:"database"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveHyperdriveDatabaseScheme] `json:"scheme"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user"`
}

func (r HyperdriveHyperdriveDatabaseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveHyperdriveDatabaseParam) implementsAccountHyperdriveConfigPatchParamsOriginUnion() {
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveHyperdriveDatabaseScheme string

const (
	HyperdriveHyperdriveDatabaseSchemePostgres   HyperdriveHyperdriveDatabaseScheme = "postgres"
	HyperdriveHyperdriveDatabaseSchemePostgresql HyperdriveHyperdriveDatabaseScheme = "postgresql"
	HyperdriveHyperdriveDatabaseSchemeMysql      HyperdriveHyperdriveDatabaseScheme = "mysql"
)

func (r HyperdriveHyperdriveDatabaseScheme) IsKnown() bool {
	switch r {
	case HyperdriveHyperdriveDatabaseSchemePostgres, HyperdriveHyperdriveDatabaseSchemePostgresql, HyperdriveHyperdriveDatabaseSchemeMysql:
		return true
	}
	return false
}

type HyperdriveInternetOriginParam struct {
	// Defines the host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// Defines the port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port,required"`
}

func (r HyperdriveInternetOriginParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveInternetOriginParam) implementsAccountHyperdriveConfigPatchParamsOriginUnion() {}

type HyperdriveMessagesItem struct {
	Code             int64                        `json:"code,required"`
	Message          string                       `json:"message,required"`
	DocumentationURL string                       `json:"documentation_url"`
	Source           HyperdriveMessagesItemSource `json:"source"`
	JSON             hyperdriveMessagesItemJSON   `json:"-"`
}

// hyperdriveMessagesItemJSON contains the JSON metadata for the struct
// [HyperdriveMessagesItem]
type hyperdriveMessagesItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HyperdriveMessagesItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveMessagesItemJSON) RawJSON() string {
	return r.raw
}

type HyperdriveMessagesItemSource struct {
	Pointer string                           `json:"pointer"`
	JSON    hyperdriveMessagesItemSourceJSON `json:"-"`
}

// hyperdriveMessagesItemSourceJSON contains the JSON metadata for the struct
// [HyperdriveMessagesItemSource]
type hyperdriveMessagesItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveMessagesItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveMessagesItemSourceJSON) RawJSON() string {
	return r.raw
}

type HyperdriveOverAccessOriginParam struct {
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID param.Field[string] `json:"access_client_id,required"`
	// Defines the Client Secret of the Access Token to use when connecting to the
	// origin database. The API never returns this write-only value.
	AccessClientSecret param.Field[string] `json:"access_client_secret,required"`
	// Defines the host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
}

func (r HyperdriveOverAccessOriginParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveOverAccessOriginParam) implementsAccountHyperdriveConfigPatchParamsOriginUnion() {}

type AccountHyperdriveConfigNewResponse struct {
	Errors   []HyperdriveMessagesItem           `json:"errors,required"`
	Messages []HyperdriveMessagesItem           `json:"messages,required"`
	Result   HyperdriveHyperdriveConfigResponse `json:"result,required"`
	// Return the status of the API call success.
	Success AccountHyperdriveConfigNewResponseSuccess `json:"success,required"`
	JSON    accountHyperdriveConfigNewResponseJSON    `json:"-"`
}

// accountHyperdriveConfigNewResponseJSON contains the JSON metadata for the struct
// [AccountHyperdriveConfigNewResponse]
type accountHyperdriveConfigNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHyperdriveConfigNewResponseJSON) RawJSON() string {
	return r.raw
}

// Return the status of the API call success.
type AccountHyperdriveConfigNewResponseSuccess bool

const (
	AccountHyperdriveConfigNewResponseSuccessTrue AccountHyperdriveConfigNewResponseSuccess = true
)

func (r AccountHyperdriveConfigNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountHyperdriveConfigNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountHyperdriveConfigGetResponse struct {
	Errors   []HyperdriveMessagesItem           `json:"errors,required"`
	Messages []HyperdriveMessagesItem           `json:"messages,required"`
	Result   HyperdriveHyperdriveConfigResponse `json:"result,required"`
	// Return the status of the API call success.
	Success AccountHyperdriveConfigGetResponseSuccess `json:"success,required"`
	JSON    accountHyperdriveConfigGetResponseJSON    `json:"-"`
}

// accountHyperdriveConfigGetResponseJSON contains the JSON metadata for the struct
// [AccountHyperdriveConfigGetResponse]
type accountHyperdriveConfigGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHyperdriveConfigGetResponseJSON) RawJSON() string {
	return r.raw
}

// Return the status of the API call success.
type AccountHyperdriveConfigGetResponseSuccess bool

const (
	AccountHyperdriveConfigGetResponseSuccessTrue AccountHyperdriveConfigGetResponseSuccess = true
)

func (r AccountHyperdriveConfigGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountHyperdriveConfigGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountHyperdriveConfigUpdateResponse struct {
	Errors   []HyperdriveMessagesItem           `json:"errors,required"`
	Messages []HyperdriveMessagesItem           `json:"messages,required"`
	Result   HyperdriveHyperdriveConfigResponse `json:"result,required"`
	// Return the status of the API call success.
	Success AccountHyperdriveConfigUpdateResponseSuccess `json:"success,required"`
	JSON    accountHyperdriveConfigUpdateResponseJSON    `json:"-"`
}

// accountHyperdriveConfigUpdateResponseJSON contains the JSON metadata for the
// struct [AccountHyperdriveConfigUpdateResponse]
type accountHyperdriveConfigUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHyperdriveConfigUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Return the status of the API call success.
type AccountHyperdriveConfigUpdateResponseSuccess bool

const (
	AccountHyperdriveConfigUpdateResponseSuccessTrue AccountHyperdriveConfigUpdateResponseSuccess = true
)

func (r AccountHyperdriveConfigUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountHyperdriveConfigUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountHyperdriveConfigListResponse struct {
	Errors   []HyperdriveMessagesItem             `json:"errors,required"`
	Messages []HyperdriveMessagesItem             `json:"messages,required"`
	Result   []HyperdriveHyperdriveConfigResponse `json:"result,required"`
	// Return the status of the API call success.
	Success    AccountHyperdriveConfigListResponseSuccess    `json:"success,required"`
	ResultInfo AccountHyperdriveConfigListResponseResultInfo `json:"result_info"`
	JSON       accountHyperdriveConfigListResponseJSON       `json:"-"`
}

// accountHyperdriveConfigListResponseJSON contains the JSON metadata for the
// struct [AccountHyperdriveConfigListResponse]
type accountHyperdriveConfigListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHyperdriveConfigListResponseJSON) RawJSON() string {
	return r.raw
}

// Return the status of the API call success.
type AccountHyperdriveConfigListResponseSuccess bool

const (
	AccountHyperdriveConfigListResponseSuccessTrue AccountHyperdriveConfigListResponseSuccess = true
)

func (r AccountHyperdriveConfigListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountHyperdriveConfigListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountHyperdriveConfigListResponseResultInfo struct {
	// Defines the total number of results for the requested service.
	Count float64 `json:"count"`
	// Defines the current page within paginated list of results.
	Page float64 `json:"page"`
	// Defines the number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Defines the total results available without any search parameters.
	TotalCount float64                                           `json:"total_count"`
	JSON       accountHyperdriveConfigListResponseResultInfoJSON `json:"-"`
}

// accountHyperdriveConfigListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountHyperdriveConfigListResponseResultInfo]
type accountHyperdriveConfigListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHyperdriveConfigListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountHyperdriveConfigDeleteResponse struct {
	Errors   []HyperdriveMessagesItem `json:"errors,required"`
	Messages []HyperdriveMessagesItem `json:"messages,required"`
	Result   interface{}              `json:"result,required,nullable"`
	// Return the status of the API call success.
	Success AccountHyperdriveConfigDeleteResponseSuccess `json:"success,required"`
	JSON    accountHyperdriveConfigDeleteResponseJSON    `json:"-"`
}

// accountHyperdriveConfigDeleteResponseJSON contains the JSON metadata for the
// struct [AccountHyperdriveConfigDeleteResponse]
type accountHyperdriveConfigDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHyperdriveConfigDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Return the status of the API call success.
type AccountHyperdriveConfigDeleteResponseSuccess bool

const (
	AccountHyperdriveConfigDeleteResponseSuccessTrue AccountHyperdriveConfigDeleteResponseSuccess = true
)

func (r AccountHyperdriveConfigDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountHyperdriveConfigDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountHyperdriveConfigPatchResponse struct {
	Errors   []HyperdriveMessagesItem           `json:"errors,required"`
	Messages []HyperdriveMessagesItem           `json:"messages,required"`
	Result   HyperdriveHyperdriveConfigResponse `json:"result,required"`
	// Return the status of the API call success.
	Success AccountHyperdriveConfigPatchResponseSuccess `json:"success,required"`
	JSON    accountHyperdriveConfigPatchResponseJSON    `json:"-"`
}

// accountHyperdriveConfigPatchResponseJSON contains the JSON metadata for the
// struct [AccountHyperdriveConfigPatchResponse]
type accountHyperdriveConfigPatchResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigPatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHyperdriveConfigPatchResponseJSON) RawJSON() string {
	return r.raw
}

// Return the status of the API call success.
type AccountHyperdriveConfigPatchResponseSuccess bool

const (
	AccountHyperdriveConfigPatchResponseSuccessTrue AccountHyperdriveConfigPatchResponseSuccess = true
)

func (r AccountHyperdriveConfigPatchResponseSuccess) IsKnown() bool {
	switch r {
	case AccountHyperdriveConfigPatchResponseSuccessTrue:
		return true
	}
	return false
}

type AccountHyperdriveConfigNewParams struct {
	HyperdriveHyperdriveConfig HyperdriveHyperdriveConfigParam `json:"hyperdrive_hyperdrive_config,required"`
}

func (r AccountHyperdriveConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.HyperdriveHyperdriveConfig)
}

type AccountHyperdriveConfigUpdateParams struct {
	HyperdriveHyperdriveConfig HyperdriveHyperdriveConfigParam `json:"hyperdrive_hyperdrive_config,required"`
}

func (r AccountHyperdriveConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.HyperdriveHyperdriveConfig)
}

type AccountHyperdriveConfigPatchParams struct {
	Caching param.Field[HyperdriveHyperdriveCachingUnionParam]         `json:"caching"`
	Mtls    param.Field[AccountHyperdriveConfigPatchParamsMtls]        `json:"mtls"`
	Name    param.Field[string]                                        `json:"name"`
	Origin  param.Field[AccountHyperdriveConfigPatchParamsOriginUnion] `json:"origin"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to
	// the origin database.
	OriginConnectionLimit param.Field[int64] `json:"origin_connection_limit"`
}

func (r AccountHyperdriveConfigPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountHyperdriveConfigPatchParamsMtls struct {
	// Define CA certificate ID obtained after uploading CA cert.
	CaCertificateID param.Field[string] `json:"ca_certificate_id"`
	// Define mTLS certificate ID obtained after uploading client cert.
	MtlsCertificateID param.Field[string] `json:"mtls_certificate_id"`
	// Set SSL mode to 'require', 'verify-ca', or 'verify-full' to verify the CA.
	Sslmode param.Field[string] `json:"sslmode"`
}

func (r AccountHyperdriveConfigPatchParamsMtls) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountHyperdriveConfigPatchParamsOrigin struct {
	// Defines the Client ID of the Access token to use when connecting to the origin
	// database.
	AccessClientID param.Field[string] `json:"access_client_id"`
	// Defines the Client Secret of the Access Token to use when connecting to the
	// origin database. The API never returns this write-only value.
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
	// Set the name of your origin database.
	Database param.Field[string] `json:"database"`
	// Defines the host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host"`
	// Set the password needed to access your origin database. The API never returns
	// this write-only value.
	Password param.Field[string] `json:"password"`
	// Defines the port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[AccountHyperdriveConfigPatchParamsOriginScheme] `json:"scheme"`
	// Set the user of your origin database.
	User param.Field[string] `json:"user"`
}

func (r AccountHyperdriveConfigPatchParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountHyperdriveConfigPatchParamsOrigin) implementsAccountHyperdriveConfigPatchParamsOriginUnion() {
}

// Satisfied by [HyperdriveHyperdriveDatabaseParam],
// [HyperdriveInternetOriginParam], [HyperdriveOverAccessOriginParam],
// [AccountHyperdriveConfigPatchParamsOrigin].
type AccountHyperdriveConfigPatchParamsOriginUnion interface {
	implementsAccountHyperdriveConfigPatchParamsOriginUnion()
}

// Specifies the URL scheme used to connect to your origin database.
type AccountHyperdriveConfigPatchParamsOriginScheme string

const (
	AccountHyperdriveConfigPatchParamsOriginSchemePostgres   AccountHyperdriveConfigPatchParamsOriginScheme = "postgres"
	AccountHyperdriveConfigPatchParamsOriginSchemePostgresql AccountHyperdriveConfigPatchParamsOriginScheme = "postgresql"
	AccountHyperdriveConfigPatchParamsOriginSchemeMysql      AccountHyperdriveConfigPatchParamsOriginScheme = "mysql"
)

func (r AccountHyperdriveConfigPatchParamsOriginScheme) IsKnown() bool {
	switch r {
	case AccountHyperdriveConfigPatchParamsOriginSchemePostgres, AccountHyperdriveConfigPatchParamsOriginSchemePostgresql, AccountHyperdriveConfigPatchParamsOriginSchemeMysql:
		return true
	}
	return false
}
