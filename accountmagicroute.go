// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountMagicRouteService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicRouteService] method instead.
type AccountMagicRouteService struct {
	Options []option.RequestOption
}

// NewAccountMagicRouteService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMagicRouteService(opts ...option.RequestOption) (r *AccountMagicRouteService) {
	r = &AccountMagicRouteService{}
	r.Options = opts
	return
}

// Creates a new Magic static route. Use `?validate_only=true` as an optional query
// parameter to run validation only without persisting changes.
func (r *AccountMagicRouteService) New(ctx context.Context, accountID string, body AccountMagicRouteNewParams, opts ...option.RequestOption) (res *MagicRoutesCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/routes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific Magic static route.
func (r *AccountMagicRouteService) Get(ctx context.Context, accountID string, routeID string, opts ...option.RequestOption) (res *AccountMagicRouteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if routeID == "" {
		err = errors.New("missing required route_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", accountID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a specific Magic static route. Use `?validate_only=true` as an optional
// query parameter to run validation only without persisting changes.
func (r *AccountMagicRouteService) Update(ctx context.Context, accountID string, routeID string, body AccountMagicRouteUpdateParams, opts ...option.RequestOption) (res *AccountMagicRouteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if routeID == "" {
		err = errors.New("missing required route_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", accountID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all Magic static routes.
func (r *AccountMagicRouteService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *MagicRoutesCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/routes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Disable and remove a specific Magic static route.
func (r *AccountMagicRouteService) Delete(ctx context.Context, accountID string, routeID string, body AccountMagicRouteDeleteParams, opts ...option.RequestOption) (res *AccountMagicRouteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if routeID == "" {
		err = errors.New("missing required route_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/routes/%s", accountID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Delete multiple Magic static routes.
func (r *AccountMagicRouteService) DeleteMany(ctx context.Context, accountID string, body AccountMagicRouteDeleteManyParams, opts ...option.RequestOption) (res *AccountMagicRouteDeleteManyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/routes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Update multiple Magic static routes. Use `?validate_only=true` as an optional
// query parameter to run validation only without persisting changes. Only fields
// for a route that need to be changed need be provided.
func (r *AccountMagicRouteService) UpdateMany(ctx context.Context, accountID string, body AccountMagicRouteUpdateManyParams, opts ...option.RequestOption) (res *AccountMagicRouteUpdateManyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/routes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type MagicRoute struct {
	// The next-hop IP Address for the static route.
	Nexthop string `json:"nexthop,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Prefix string `json:"prefix,required"`
	// Priority of the static route.
	Priority int64 `json:"priority,required"`
	// Identifier
	ID string `json:"id"`
	// When the route was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional human provided description of the static route.
	Description string `json:"description"`
	// When the route was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Used only for ECMP routes.
	Scope MagicScope `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight int64          `json:"weight"`
	JSON   magicRouteJSON `json:"-"`
}

// magicRouteJSON contains the JSON metadata for the struct [MagicRoute]
type magicRouteJSON struct {
	Nexthop     apijson.Field
	Prefix      apijson.Field
	Priority    apijson.Field
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Scope       apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicRouteJSON) RawJSON() string {
	return r.raw
}

type MagicRouteAddSingleRequestParam struct {
	// The next-hop IP Address for the static route.
	Nexthop param.Field[string] `json:"nexthop,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Prefix param.Field[string] `json:"prefix,required"`
	// Priority of the static route.
	Priority param.Field[int64] `json:"priority,required"`
	// An optional human provided description of the static route.
	Description param.Field[string] `json:"description"`
	// Used only for ECMP routes.
	Scope param.Field[MagicScopeParam] `json:"scope"`
	// Optional weight of the ECMP scope - if provided.
	Weight param.Field[int64] `json:"weight"`
}

func (r MagicRouteAddSingleRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicRoutesCollectionResponse struct {
	Result MagicRoutesCollectionResponseResult `json:"result"`
	JSON   magicRoutesCollectionResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// magicRoutesCollectionResponseJSON contains the JSON metadata for the struct
// [MagicRoutesCollectionResponse]
type magicRoutesCollectionResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRoutesCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicRoutesCollectionResponseJSON) RawJSON() string {
	return r.raw
}

type MagicRoutesCollectionResponseResult struct {
	Routes []MagicRoute                            `json:"routes"`
	JSON   magicRoutesCollectionResponseResultJSON `json:"-"`
}

// magicRoutesCollectionResponseResultJSON contains the JSON metadata for the
// struct [MagicRoutesCollectionResponseResult]
type magicRoutesCollectionResponseResultJSON struct {
	Routes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRoutesCollectionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicRoutesCollectionResponseResultJSON) RawJSON() string {
	return r.raw
}

// Used only for ECMP routes.
type MagicScope struct {
	// List of colo names for the ECMP scope.
	ColoNames []string `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions []string       `json:"colo_regions"`
	JSON        magicScopeJSON `json:"-"`
}

// magicScopeJSON contains the JSON metadata for the struct [MagicScope]
type magicScopeJSON struct {
	ColoNames   apijson.Field
	ColoRegions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicScopeJSON) RawJSON() string {
	return r.raw
}

// Used only for ECMP routes.
type MagicScopeParam struct {
	// List of colo names for the ECMP scope.
	ColoNames param.Field[[]string] `json:"colo_names"`
	// List of colo regions for the ECMP scope.
	ColoRegions param.Field[[]string] `json:"colo_regions"`
}

func (r MagicScopeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicRouteGetResponse struct {
	Result AccountMagicRouteGetResponseResult `json:"result"`
	JSON   accountMagicRouteGetResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicRouteGetResponseJSON contains the JSON metadata for the struct
// [AccountMagicRouteGetResponse]
type accountMagicRouteGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicRouteGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicRouteGetResponseResult struct {
	Route MagicRoute                             `json:"route"`
	JSON  accountMagicRouteGetResponseResultJSON `json:"-"`
}

// accountMagicRouteGetResponseResultJSON contains the JSON metadata for the struct
// [AccountMagicRouteGetResponseResult]
type accountMagicRouteGetResponseResultJSON struct {
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicRouteGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicRouteUpdateResponse struct {
	Result AccountMagicRouteUpdateResponseResult `json:"result"`
	JSON   accountMagicRouteUpdateResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicRouteUpdateResponseJSON contains the JSON metadata for the struct
// [AccountMagicRouteUpdateResponse]
type accountMagicRouteUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicRouteUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicRouteUpdateResponseResult struct {
	Modified      bool                                      `json:"modified"`
	ModifiedRoute MagicRoute                                `json:"modified_route"`
	JSON          accountMagicRouteUpdateResponseResultJSON `json:"-"`
}

// accountMagicRouteUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountMagicRouteUpdateResponseResult]
type accountMagicRouteUpdateResponseResultJSON struct {
	Modified      apijson.Field
	ModifiedRoute apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountMagicRouteUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicRouteUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicRouteDeleteResponse struct {
	Result AccountMagicRouteDeleteResponseResult `json:"result"`
	JSON   accountMagicRouteDeleteResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicRouteDeleteResponseJSON contains the JSON metadata for the struct
// [AccountMagicRouteDeleteResponse]
type accountMagicRouteDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicRouteDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicRouteDeleteResponseResult struct {
	Deleted      bool                                      `json:"deleted"`
	DeletedRoute MagicRoute                                `json:"deleted_route"`
	JSON         accountMagicRouteDeleteResponseResultJSON `json:"-"`
}

// accountMagicRouteDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountMagicRouteDeleteResponseResult]
type accountMagicRouteDeleteResponseResultJSON struct {
	Deleted      apijson.Field
	DeletedRoute apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountMagicRouteDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicRouteDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicRouteDeleteManyResponse struct {
	Result AccountMagicRouteDeleteManyResponseResult `json:"result"`
	JSON   accountMagicRouteDeleteManyResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicRouteDeleteManyResponseJSON contains the JSON metadata for the
// struct [AccountMagicRouteDeleteManyResponse]
type accountMagicRouteDeleteManyResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteDeleteManyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicRouteDeleteManyResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicRouteDeleteManyResponseResult struct {
	Deleted       bool                                          `json:"deleted"`
	DeletedRoutes []MagicRoute                                  `json:"deleted_routes"`
	JSON          accountMagicRouteDeleteManyResponseResultJSON `json:"-"`
}

// accountMagicRouteDeleteManyResponseResultJSON contains the JSON metadata for the
// struct [AccountMagicRouteDeleteManyResponseResult]
type accountMagicRouteDeleteManyResponseResultJSON struct {
	Deleted       apijson.Field
	DeletedRoutes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountMagicRouteDeleteManyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicRouteDeleteManyResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicRouteUpdateManyResponse struct {
	Result AccountMagicRouteUpdateManyResponseResult `json:"result"`
	JSON   accountMagicRouteUpdateManyResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicRouteUpdateManyResponseJSON contains the JSON metadata for the
// struct [AccountMagicRouteUpdateManyResponse]
type accountMagicRouteUpdateManyResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteUpdateManyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicRouteUpdateManyResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicRouteUpdateManyResponseResult struct {
	Modified       bool                                          `json:"modified"`
	ModifiedRoutes []MagicRoute                                  `json:"modified_routes"`
	JSON           accountMagicRouteUpdateManyResponseResultJSON `json:"-"`
}

// accountMagicRouteUpdateManyResponseResultJSON contains the JSON metadata for the
// struct [AccountMagicRouteUpdateManyResponseResult]
type accountMagicRouteUpdateManyResponseResultJSON struct {
	Modified       apijson.Field
	ModifiedRoutes apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountMagicRouteUpdateManyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicRouteUpdateManyResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicRouteNewParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountMagicRouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMagicRouteUpdateParams struct {
	MagicRouteAddSingleRequest MagicRouteAddSingleRequestParam `json:"magic_route_add_single_request,required"`
}

func (r AccountMagicRouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MagicRouteAddSingleRequest)
}

type AccountMagicRouteDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountMagicRouteDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMagicRouteDeleteManyParams struct {
	Routes param.Field[[]AccountMagicRouteDeleteManyParamsRoute] `json:"routes,required"`
}

func (r AccountMagicRouteDeleteManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicRouteDeleteManyParamsRoute struct {
}

func (r AccountMagicRouteDeleteManyParamsRoute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicRouteUpdateManyParams struct {
	Routes param.Field[[]AccountMagicRouteUpdateManyParamsRoute] `json:"routes,required"`
}

func (r AccountMagicRouteUpdateManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicRouteUpdateManyParamsRoute struct {
	MagicRouteAddSingleRequestParam
}

func (r AccountMagicRouteUpdateManyParamsRoute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
