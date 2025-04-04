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

// Returns a list of Hyperdrives
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

type HyperdriveAPIResponseCommon struct {
	Errors   []HyperdriveMessagesItem `json:"errors,required"`
	Messages []HyperdriveMessagesItem `json:"messages,required"`
	Result   interface{}              `json:"result,required"`
	// Whether the API call was successful
	Success HyperdriveAPIResponseCommonSuccess `json:"success,required"`
	JSON    hyperdriveAPIResponseCommonJSON    `json:"-"`
}

// hyperdriveAPIResponseCommonJSON contains the JSON metadata for the struct
// [HyperdriveAPIResponseCommon]
type hyperdriveAPIResponseCommonJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveAPIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveAPIResponseCommonJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HyperdriveAPIResponseCommonSuccess bool

const (
	HyperdriveAPIResponseCommonSuccessTrue HyperdriveAPIResponseCommonSuccess = true
)

func (r HyperdriveAPIResponseCommonSuccess) IsKnown() bool {
	switch r {
	case HyperdriveAPIResponseCommonSuccessTrue:
		return true
	}
	return false
}

type HyperdriveAPIResponseSingle struct {
	Errors   []HyperdriveMessagesItem `json:"errors,required"`
	Messages []HyperdriveMessagesItem `json:"messages,required"`
	Result   interface{}              `json:"result,required"`
	// Whether the API call was successful
	Success HyperdriveAPIResponseSingleSuccess `json:"success,required"`
	JSON    hyperdriveAPIResponseSingleJSON    `json:"-"`
}

// hyperdriveAPIResponseSingleJSON contains the JSON metadata for the struct
// [HyperdriveAPIResponseSingle]
type hyperdriveAPIResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveAPIResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HyperdriveAPIResponseSingleSuccess bool

const (
	HyperdriveAPIResponseSingleSuccessTrue HyperdriveAPIResponseSingleSuccess = true
)

func (r HyperdriveAPIResponseSingleSuccess) IsKnown() bool {
	switch r {
	case HyperdriveAPIResponseSingleSuccessTrue:
		return true
	}
	return false
}

type HyperdriveHyperdriveCaching struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
	Disabled bool `json:"disabled"`
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
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
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge int64 `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate int64                                                             `json:"stale_while_revalidate"`
	JSON                 hyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabledJSON `json:"-"`
	HyperdriveHyperdriveCachingCommon
}

// hyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabledJSON contains the
// JSON metadata for the struct
// [HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabled]
type hyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabledJSON struct {
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

// Satisfied by [HyperdriveHyperdriveCachingCommonParam],
// [HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabledParam].
type HyperdriveHyperdriveCachingUnionParam interface {
	implementsHyperdriveHyperdriveCachingUnionParam()
}

type HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabledParam struct {
	// When present, specifies max duration for which items should persist in the
	// cache. Not returned if set to default. (Default: 60)
	MaxAge param.Field[int64] `json:"max_age"`
	// When present, indicates the number of seconds cache may serve the response after
	// it becomes stale. Not returned if set to default. (Default: 15)
	StaleWhileRevalidate param.Field[int64] `json:"stale_while_revalidate"`
	HyperdriveHyperdriveCachingCommonParam
}

func (r HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabledParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveHyperdriveCachingHyperdriveHyperdriveCachingEnabledParam) implementsHyperdriveHyperdriveCachingUnionParam() {
}

type HyperdriveHyperdriveCachingCommon struct {
	// When set to true, disables the caching of SQL responses. (Default: false)
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
	// When set to true, disables the caching of SQL responses. (Default: false)
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
}

func (r HyperdriveHyperdriveConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [HyperdriveHyperdriveConfigOriginPublicDatabaseParam],
// [HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelParam].
type HyperdriveHyperdriveConfigOriginUnionParam interface {
	implementsHyperdriveHyperdriveConfigOriginUnionParam()
}

type HyperdriveHyperdriveConfigOriginPublicDatabaseParam struct {
	HyperdriveHyperdriveDatabaseFullParam
	HyperdriveInternetOriginParam
}

func (r HyperdriveHyperdriveConfigOriginPublicDatabaseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveHyperdriveConfigOriginPublicDatabaseParam) implementsHyperdriveHyperdriveConfigOriginUnionParam() {
}

type HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelParam struct {
	HyperdriveHyperdriveDatabaseFullParam
	HyperdriveOverAccessOriginParam
}

func (r HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveHyperdriveConfigOriginAccessProtectedDatabaseBehindCloudflareTunnelParam) implementsHyperdriveHyperdriveConfigOriginUnionParam() {
}

type HyperdriveHyperdriveConfigResponse struct {
	// Identifier
	ID      string                                   `json:"id,required"`
	Name    string                                   `json:"name,required"`
	Origin  HyperdriveHyperdriveConfigResponseOrigin `json:"origin,required"`
	Caching HyperdriveHyperdriveCaching              `json:"caching"`
	// When the Hyperdrive configuration was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// When the Hyperdrive configuration was last modified.
	ModifiedOn time.Time                              `json:"modified_on" format:"date-time"`
	JSON       hyperdriveHyperdriveConfigResponseJSON `json:"-"`
}

// hyperdriveHyperdriveConfigResponseJSON contains the JSON metadata for the struct
// [HyperdriveHyperdriveConfigResponse]
type hyperdriveHyperdriveConfigResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Origin      apijson.Field
	Caching     apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveHyperdriveConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveHyperdriveConfigResponseJSON) RawJSON() string {
	return r.raw
}

type HyperdriveHyperdriveConfigResponseOrigin struct {
	// The Client ID of the Access token to use when connecting to the origin database.
	AccessClientID string `json:"access_client_id"`
	// The name of your origin database.
	Database string `json:"database"`
	// The host (hostname or IP) of your origin database.
	Host string `json:"host"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64 `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveHyperdriveConfigResponseOriginScheme `json:"scheme"`
	// The user of your origin database.
	User  string                                       `json:"user"`
	JSON  hyperdriveHyperdriveConfigResponseOriginJSON `json:"-"`
	union HyperdriveHyperdriveConfigResponseOriginUnion
}

// hyperdriveHyperdriveConfigResponseOriginJSON contains the JSON metadata for the
// struct [HyperdriveHyperdriveConfigResponseOrigin]
type hyperdriveHyperdriveConfigResponseOriginJSON struct {
	AccessClientID apijson.Field
	Database       apijson.Field
	Host           apijson.Field
	Port           apijson.Field
	Scheme         apijson.Field
	User           apijson.Field
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
	JSON hyperdriveHyperdriveConfigResponseOriginPublicDatabaseJSON `json:"-"`
	HyperdriveHyperdriveDatabaseFull
	HyperdriveInternetOrigin
}

// hyperdriveHyperdriveConfigResponseOriginPublicDatabaseJSON contains the JSON
// metadata for the struct [HyperdriveHyperdriveConfigResponseOriginPublicDatabase]
type hyperdriveHyperdriveConfigResponseOriginPublicDatabaseJSON struct {
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

type HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel struct {
	JSON hyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON `json:"-"`
	HyperdriveHyperdriveDatabaseFull
	HyperdriveOverAccessOrigin
}

// hyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON
// contains the JSON metadata for the struct
// [HyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnel]
type hyperdriveHyperdriveConfigResponseOriginAccessProtectedDatabaseBehindCloudflareTunnelJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
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
type HyperdriveHyperdriveConfigResponseOriginScheme string

const (
	HyperdriveHyperdriveConfigResponseOriginSchemePostgres   HyperdriveHyperdriveConfigResponseOriginScheme = "postgres"
	HyperdriveHyperdriveConfigResponseOriginSchemePostgresql HyperdriveHyperdriveConfigResponseOriginScheme = "postgresql"
)

func (r HyperdriveHyperdriveConfigResponseOriginScheme) IsKnown() bool {
	switch r {
	case HyperdriveHyperdriveConfigResponseOriginSchemePostgres, HyperdriveHyperdriveConfigResponseOriginSchemePostgresql:
		return true
	}
	return false
}

type HyperdriveHyperdriveDatabaseParam struct {
	// The name of your origin database.
	Database param.Field[string] `json:"database"`
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveHyperdriveDatabaseScheme] `json:"scheme"`
	// The user of your origin database.
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
)

func (r HyperdriveHyperdriveDatabaseScheme) IsKnown() bool {
	switch r {
	case HyperdriveHyperdriveDatabaseSchemePostgres, HyperdriveHyperdriveDatabaseSchemePostgresql:
		return true
	}
	return false
}

type HyperdriveHyperdriveDatabaseFull struct {
	// The name of your origin database.
	Database string `json:"database,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme HyperdriveHyperdriveDatabaseFullScheme `json:"scheme,required"`
	// The user of your origin database.
	User string                               `json:"user,required"`
	JSON hyperdriveHyperdriveDatabaseFullJSON `json:"-"`
}

// hyperdriveHyperdriveDatabaseFullJSON contains the JSON metadata for the struct
// [HyperdriveHyperdriveDatabaseFull]
type hyperdriveHyperdriveDatabaseFullJSON struct {
	Database    apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveHyperdriveDatabaseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveHyperdriveDatabaseFullJSON) RawJSON() string {
	return r.raw
}

// Specifies the URL scheme used to connect to your origin database.
type HyperdriveHyperdriveDatabaseFullScheme string

const (
	HyperdriveHyperdriveDatabaseFullSchemePostgres   HyperdriveHyperdriveDatabaseFullScheme = "postgres"
	HyperdriveHyperdriveDatabaseFullSchemePostgresql HyperdriveHyperdriveDatabaseFullScheme = "postgresql"
)

func (r HyperdriveHyperdriveDatabaseFullScheme) IsKnown() bool {
	switch r {
	case HyperdriveHyperdriveDatabaseFullSchemePostgres, HyperdriveHyperdriveDatabaseFullSchemePostgresql:
		return true
	}
	return false
}

type HyperdriveHyperdriveDatabaseFullParam struct {
	// The name of your origin database.
	Database param.Field[string] `json:"database,required"`
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password,required"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[HyperdriveHyperdriveDatabaseFullScheme] `json:"scheme,required"`
	// The user of your origin database.
	User param.Field[string] `json:"user,required"`
}

func (r HyperdriveHyperdriveDatabaseFullParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveInternetOrigin struct {
	// The host (hostname or IP) of your origin database.
	Host string `json:"host,required"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port int64                        `json:"port,required"`
	JSON hyperdriveInternetOriginJSON `json:"-"`
}

// hyperdriveInternetOriginJSON contains the JSON metadata for the struct
// [HyperdriveInternetOrigin]
type hyperdriveInternetOriginJSON struct {
	Host        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveInternetOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveInternetOriginJSON) RawJSON() string {
	return r.raw
}

type HyperdriveInternetOriginParam struct {
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port,required"`
}

func (r HyperdriveInternetOriginParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveInternetOriginParam) implementsAccountHyperdriveConfigPatchParamsOriginUnion() {}

type HyperdriveMessagesItem struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    hyperdriveMessagesItemJSON `json:"-"`
}

// hyperdriveMessagesItemJSON contains the JSON metadata for the struct
// [HyperdriveMessagesItem]
type hyperdriveMessagesItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveMessagesItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveMessagesItemJSON) RawJSON() string {
	return r.raw
}

type HyperdriveOverAccessOrigin struct {
	// The Client ID of the Access token to use when connecting to the origin database.
	AccessClientID string `json:"access_client_id,required"`
	// The host (hostname or IP) of your origin database.
	Host string                         `json:"host,required"`
	JSON hyperdriveOverAccessOriginJSON `json:"-"`
}

// hyperdriveOverAccessOriginJSON contains the JSON metadata for the struct
// [HyperdriveOverAccessOrigin]
type hyperdriveOverAccessOriginJSON struct {
	AccessClientID apijson.Field
	Host           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HyperdriveOverAccessOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hyperdriveOverAccessOriginJSON) RawJSON() string {
	return r.raw
}

type HyperdriveOverAccessOriginParam struct {
	// The Client ID of the Access token to use when connecting to the origin database.
	AccessClientID param.Field[string] `json:"access_client_id,required"`
	// The Client Secret of the Access token to use when connecting to the origin
	// database. This value is write-only and never returned by the API.
	AccessClientSecret param.Field[string] `json:"access_client_secret,required"`
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host,required"`
}

func (r HyperdriveOverAccessOriginParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HyperdriveOverAccessOriginParam) implementsAccountHyperdriveConfigPatchParamsOriginUnion() {}

type AccountHyperdriveConfigNewResponse struct {
	Result HyperdriveHyperdriveConfigResponse     `json:"result"`
	JSON   accountHyperdriveConfigNewResponseJSON `json:"-"`
	HyperdriveAPIResponseSingle
}

// accountHyperdriveConfigNewResponseJSON contains the JSON metadata for the struct
// [AccountHyperdriveConfigNewResponse]
type accountHyperdriveConfigNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHyperdriveConfigNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountHyperdriveConfigGetResponse struct {
	Result HyperdriveHyperdriveConfigResponse     `json:"result"`
	JSON   accountHyperdriveConfigGetResponseJSON `json:"-"`
	HyperdriveAPIResponseSingle
}

// accountHyperdriveConfigGetResponseJSON contains the JSON metadata for the struct
// [AccountHyperdriveConfigGetResponse]
type accountHyperdriveConfigGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHyperdriveConfigGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountHyperdriveConfigUpdateResponse struct {
	Result HyperdriveHyperdriveConfigResponse        `json:"result"`
	JSON   accountHyperdriveConfigUpdateResponseJSON `json:"-"`
	HyperdriveAPIResponseSingle
}

// accountHyperdriveConfigUpdateResponseJSON contains the JSON metadata for the
// struct [AccountHyperdriveConfigUpdateResponse]
type accountHyperdriveConfigUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHyperdriveConfigUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountHyperdriveConfigListResponse struct {
	Result     []HyperdriveHyperdriveConfigResponse          `json:"result"`
	ResultInfo AccountHyperdriveConfigListResponseResultInfo `json:"result_info"`
	JSON       accountHyperdriveConfigListResponseJSON       `json:"-"`
	HyperdriveAPIResponseCommon
}

// accountHyperdriveConfigListResponseJSON contains the JSON metadata for the
// struct [AccountHyperdriveConfigListResponse]
type accountHyperdriveConfigListResponseJSON struct {
	Result      apijson.Field
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

type AccountHyperdriveConfigListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
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
	Result interface{}                               `json:"result,nullable"`
	JSON   accountHyperdriveConfigDeleteResponseJSON `json:"-"`
	HyperdriveAPIResponseSingle
}

// accountHyperdriveConfigDeleteResponseJSON contains the JSON metadata for the
// struct [AccountHyperdriveConfigDeleteResponse]
type accountHyperdriveConfigDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHyperdriveConfigDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountHyperdriveConfigPatchResponse struct {
	Result HyperdriveHyperdriveConfigResponse       `json:"result"`
	JSON   accountHyperdriveConfigPatchResponseJSON `json:"-"`
	HyperdriveAPIResponseCommon
}

// accountHyperdriveConfigPatchResponseJSON contains the JSON metadata for the
// struct [AccountHyperdriveConfigPatchResponse]
type accountHyperdriveConfigPatchResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountHyperdriveConfigPatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountHyperdriveConfigPatchResponseJSON) RawJSON() string {
	return r.raw
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
	Name    param.Field[string]                                        `json:"name"`
	Origin  param.Field[AccountHyperdriveConfigPatchParamsOriginUnion] `json:"origin"`
}

func (r AccountHyperdriveConfigPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountHyperdriveConfigPatchParamsOrigin struct {
	// The Client ID of the Access token to use when connecting to the origin database.
	AccessClientID param.Field[string] `json:"access_client_id"`
	// The Client Secret of the Access token to use when connecting to the origin
	// database. This value is write-only and never returned by the API.
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
	// The name of your origin database.
	Database param.Field[string] `json:"database"`
	// The host (hostname or IP) of your origin database.
	Host param.Field[string] `json:"host"`
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password"`
	// The port (default: 5432 for Postgres) of your origin database.
	Port param.Field[int64] `json:"port"`
	// Specifies the URL scheme used to connect to your origin database.
	Scheme param.Field[AccountHyperdriveConfigPatchParamsOriginScheme] `json:"scheme"`
	// The user of your origin database.
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
)

func (r AccountHyperdriveConfigPatchParamsOriginScheme) IsKnown() bool {
	switch r {
	case AccountHyperdriveConfigPatchParamsOriginSchemePostgres, AccountHyperdriveConfigPatchParamsOriginSchemePostgresql:
		return true
	}
	return false
}
