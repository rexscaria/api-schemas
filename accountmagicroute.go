// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
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
func (r *AccountMagicRouteService) New(ctx context.Context, accountID string, body AccountMagicRouteNewParams, opts ...option.RequestOption) (res *AccountMagicRouteNewResponse, err error) {
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
func (r *AccountMagicRouteService) Delete(ctx context.Context, accountID string, routeID string, opts ...option.RequestOption) (res *AccountMagicRouteDeleteResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Delete multiple Magic static routes.
func (r *AccountMagicRouteService) DeleteMany(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountMagicRouteDeleteManyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/routes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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
	// Identifier
	ID string `json:"id,required"`
	// The next-hop IP Address for the static route.
	Nexthop string `json:"nexthop,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Prefix string `json:"prefix,required"`
	// Priority of the static route.
	Priority int64 `json:"priority,required"`
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
	ID          apijson.Field
	Nexthop     apijson.Field
	Prefix      apijson.Field
	Priority    apijson.Field
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
	Errors   []MagicMessageItem                  `json:"errors,required"`
	Messages []MagicMessageItem                  `json:"messages,required"`
	Result   MagicRoutesCollectionResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success MagicRoutesCollectionResponseSuccess `json:"success,required"`
	JSON    magicRoutesCollectionResponseJSON    `json:"-"`
}

// magicRoutesCollectionResponseJSON contains the JSON metadata for the struct
// [MagicRoutesCollectionResponse]
type magicRoutesCollectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type MagicRoutesCollectionResponseSuccess bool

const (
	MagicRoutesCollectionResponseSuccessTrue MagicRoutesCollectionResponseSuccess = true
)

func (r MagicRoutesCollectionResponseSuccess) IsKnown() bool {
	switch r {
	case MagicRoutesCollectionResponseSuccessTrue:
		return true
	}
	return false
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

type AccountMagicRouteNewResponse struct {
	Errors   []AccountMagicRouteNewResponseError   `json:"errors,required"`
	Messages []AccountMagicRouteNewResponseMessage `json:"messages,required"`
	Result   MagicRoute                            `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicRouteNewResponseSuccess `json:"success,required"`
	JSON    accountMagicRouteNewResponseJSON    `json:"-"`
}

// accountMagicRouteNewResponseJSON contains the JSON metadata for the struct
// [AccountMagicRouteNewResponse]
type accountMagicRouteNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicRouteNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicRouteNewResponseError struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           AccountMagicRouteNewResponseErrorsSource `json:"source"`
	JSON             accountMagicRouteNewResponseErrorJSON    `json:"-"`
}

// accountMagicRouteNewResponseErrorJSON contains the JSON metadata for the struct
// [AccountMagicRouteNewResponseError]
type accountMagicRouteNewResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMagicRouteNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicRouteNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountMagicRouteNewResponseErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    accountMagicRouteNewResponseErrorsSourceJSON `json:"-"`
}

// accountMagicRouteNewResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountMagicRouteNewResponseErrorsSource]
type accountMagicRouteNewResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteNewResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicRouteNewResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountMagicRouteNewResponseMessage struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           AccountMagicRouteNewResponseMessagesSource `json:"source"`
	JSON             accountMagicRouteNewResponseMessageJSON    `json:"-"`
}

// accountMagicRouteNewResponseMessageJSON contains the JSON metadata for the
// struct [AccountMagicRouteNewResponseMessage]
type accountMagicRouteNewResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMagicRouteNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicRouteNewResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountMagicRouteNewResponseMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    accountMagicRouteNewResponseMessagesSourceJSON `json:"-"`
}

// accountMagicRouteNewResponseMessagesSourceJSON contains the JSON metadata for
// the struct [AccountMagicRouteNewResponseMessagesSource]
type accountMagicRouteNewResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicRouteNewResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicRouteNewResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountMagicRouteNewResponseSuccess bool

const (
	AccountMagicRouteNewResponseSuccessTrue AccountMagicRouteNewResponseSuccess = true
)

func (r AccountMagicRouteNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicRouteNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicRouteGetResponse struct {
	Errors   []MagicMessageItem                 `json:"errors,required"`
	Messages []MagicMessageItem                 `json:"messages,required"`
	Result   AccountMagicRouteGetResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicRouteGetResponseSuccess `json:"success,required"`
	JSON    accountMagicRouteGetResponseJSON    `json:"-"`
}

// accountMagicRouteGetResponseJSON contains the JSON metadata for the struct
// [AccountMagicRouteGetResponse]
type accountMagicRouteGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type AccountMagicRouteGetResponseSuccess bool

const (
	AccountMagicRouteGetResponseSuccessTrue AccountMagicRouteGetResponseSuccess = true
)

func (r AccountMagicRouteGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicRouteGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicRouteUpdateResponse struct {
	Errors   []MagicMessageItem                    `json:"errors,required"`
	Messages []MagicMessageItem                    `json:"messages,required"`
	Result   AccountMagicRouteUpdateResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicRouteUpdateResponseSuccess `json:"success,required"`
	JSON    accountMagicRouteUpdateResponseJSON    `json:"-"`
}

// accountMagicRouteUpdateResponseJSON contains the JSON metadata for the struct
// [AccountMagicRouteUpdateResponse]
type accountMagicRouteUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type AccountMagicRouteUpdateResponseSuccess bool

const (
	AccountMagicRouteUpdateResponseSuccessTrue AccountMagicRouteUpdateResponseSuccess = true
)

func (r AccountMagicRouteUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicRouteUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicRouteDeleteResponse struct {
	Errors   []MagicMessageItem                    `json:"errors,required"`
	Messages []MagicMessageItem                    `json:"messages,required"`
	Result   AccountMagicRouteDeleteResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicRouteDeleteResponseSuccess `json:"success,required"`
	JSON    accountMagicRouteDeleteResponseJSON    `json:"-"`
}

// accountMagicRouteDeleteResponseJSON contains the JSON metadata for the struct
// [AccountMagicRouteDeleteResponse]
type accountMagicRouteDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type AccountMagicRouteDeleteResponseSuccess bool

const (
	AccountMagicRouteDeleteResponseSuccessTrue AccountMagicRouteDeleteResponseSuccess = true
)

func (r AccountMagicRouteDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicRouteDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicRouteDeleteManyResponse struct {
	Errors   []MagicMessageItem                        `json:"errors,required"`
	Messages []MagicMessageItem                        `json:"messages,required"`
	Result   AccountMagicRouteDeleteManyResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicRouteDeleteManyResponseSuccess `json:"success,required"`
	JSON    accountMagicRouteDeleteManyResponseJSON    `json:"-"`
}

// accountMagicRouteDeleteManyResponseJSON contains the JSON metadata for the
// struct [AccountMagicRouteDeleteManyResponse]
type accountMagicRouteDeleteManyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type AccountMagicRouteDeleteManyResponseSuccess bool

const (
	AccountMagicRouteDeleteManyResponseSuccessTrue AccountMagicRouteDeleteManyResponseSuccess = true
)

func (r AccountMagicRouteDeleteManyResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicRouteDeleteManyResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicRouteUpdateManyResponse struct {
	Errors   []MagicMessageItem                        `json:"errors,required"`
	Messages []MagicMessageItem                        `json:"messages,required"`
	Result   AccountMagicRouteUpdateManyResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicRouteUpdateManyResponseSuccess `json:"success,required"`
	JSON    accountMagicRouteUpdateManyResponseJSON    `json:"-"`
}

// accountMagicRouteUpdateManyResponseJSON contains the JSON metadata for the
// struct [AccountMagicRouteUpdateManyResponse]
type accountMagicRouteUpdateManyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type AccountMagicRouteUpdateManyResponseSuccess bool

const (
	AccountMagicRouteUpdateManyResponseSuccessTrue AccountMagicRouteUpdateManyResponseSuccess = true
)

func (r AccountMagicRouteUpdateManyResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicRouteUpdateManyResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicRouteNewParams struct {
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

func (r AccountMagicRouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicRouteUpdateParams struct {
	MagicRouteAddSingleRequest MagicRouteAddSingleRequestParam `json:"magic_route_add_single_request,required"`
}

func (r AccountMagicRouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MagicRouteAddSingleRequest)
}

type AccountMagicRouteUpdateManyParams struct {
	Routes param.Field[[]AccountMagicRouteUpdateManyParamsRoute] `json:"routes,required"`
}

func (r AccountMagicRouteUpdateManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicRouteUpdateManyParamsRoute struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
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

func (r AccountMagicRouteUpdateManyParamsRoute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
