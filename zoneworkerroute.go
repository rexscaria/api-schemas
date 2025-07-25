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

// ZoneWorkerRouteService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneWorkerRouteService] method instead.
type ZoneWorkerRouteService struct {
	Options []option.RequestOption
}

// NewZoneWorkerRouteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneWorkerRouteService(opts ...option.RequestOption) (r *ZoneWorkerRouteService) {
	r = &ZoneWorkerRouteService{}
	r.Options = opts
	return
}

// Creates a route that maps a URL pattern to a Worker.
func (r *ZoneWorkerRouteService) New(ctx context.Context, zoneID string, body ZoneWorkerRouteNewParams, opts ...option.RequestOption) (res *ZoneWorkerRouteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/workers/routes", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns information about a route, including URL pattern and Worker.
func (r *ZoneWorkerRouteService) Get(ctx context.Context, zoneID string, routeID string, opts ...option.RequestOption) (res *ZoneWorkerRouteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if routeID == "" {
		err = errors.New("missing required route_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/workers/routes/%s", zoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the URL pattern or Worker associated with a route.
func (r *ZoneWorkerRouteService) Update(ctx context.Context, zoneID string, routeID string, body ZoneWorkerRouteUpdateParams, opts ...option.RequestOption) (res *ZoneWorkerRouteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if routeID == "" {
		err = errors.New("missing required route_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/workers/routes/%s", zoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns routes for a zone.
func (r *ZoneWorkerRouteService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneWorkerRouteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/workers/routes", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a route.
func (r *ZoneWorkerRouteService) Delete(ctx context.Context, zoneID string, routeID string, opts ...option.RequestOption) (res *ZoneWorkerRouteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if routeID == "" {
		err = errors.New("missing required route_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/workers/routes/%s", zoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ZoneWorkerRouteNewResponse struct {
	Errors   []WorkersMessages                `json:"errors,required"`
	Messages []WorkersMessages                `json:"messages,required"`
	Result   ZoneWorkerRouteNewResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneWorkerRouteNewResponseSuccess `json:"success,required"`
	JSON    zoneWorkerRouteNewResponseJSON    `json:"-"`
}

// zoneWorkerRouteNewResponseJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteNewResponse]
type zoneWorkerRouteNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWorkerRouteNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWorkerRouteNewResponseResult struct {
	// Identifier.
	ID string `json:"id,required"`
	// Pattern to match incoming requests against.
	// [Learn more](https://developers.cloudflare.com/workers/configuration/routing/routes/#matching-behavior).
	Pattern string `json:"pattern,required"`
	// Name of the script to run if the route matches.
	Script string                               `json:"script"`
	JSON   zoneWorkerRouteNewResponseResultJSON `json:"-"`
}

// zoneWorkerRouteNewResponseResultJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteNewResponseResult]
type zoneWorkerRouteNewResponseResultJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWorkerRouteNewResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneWorkerRouteNewResponseSuccess bool

const (
	ZoneWorkerRouteNewResponseSuccessTrue ZoneWorkerRouteNewResponseSuccess = true
)

func (r ZoneWorkerRouteNewResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneWorkerRouteNewResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneWorkerRouteGetResponse struct {
	Errors   []WorkersMessages                `json:"errors,required"`
	Messages []WorkersMessages                `json:"messages,required"`
	Result   ZoneWorkerRouteGetResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneWorkerRouteGetResponseSuccess `json:"success,required"`
	JSON    zoneWorkerRouteGetResponseJSON    `json:"-"`
}

// zoneWorkerRouteGetResponseJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteGetResponse]
type zoneWorkerRouteGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWorkerRouteGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWorkerRouteGetResponseResult struct {
	// Identifier.
	ID string `json:"id,required"`
	// Pattern to match incoming requests against.
	// [Learn more](https://developers.cloudflare.com/workers/configuration/routing/routes/#matching-behavior).
	Pattern string `json:"pattern,required"`
	// Name of the script to run if the route matches.
	Script string                               `json:"script"`
	JSON   zoneWorkerRouteGetResponseResultJSON `json:"-"`
}

// zoneWorkerRouteGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteGetResponseResult]
type zoneWorkerRouteGetResponseResultJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWorkerRouteGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneWorkerRouteGetResponseSuccess bool

const (
	ZoneWorkerRouteGetResponseSuccessTrue ZoneWorkerRouteGetResponseSuccess = true
)

func (r ZoneWorkerRouteGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneWorkerRouteGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneWorkerRouteUpdateResponse struct {
	Errors   []WorkersMessages                   `json:"errors,required"`
	Messages []WorkersMessages                   `json:"messages,required"`
	Result   ZoneWorkerRouteUpdateResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneWorkerRouteUpdateResponseSuccess `json:"success,required"`
	JSON    zoneWorkerRouteUpdateResponseJSON    `json:"-"`
}

// zoneWorkerRouteUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteUpdateResponse]
type zoneWorkerRouteUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWorkerRouteUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWorkerRouteUpdateResponseResult struct {
	// Identifier.
	ID string `json:"id,required"`
	// Pattern to match incoming requests against.
	// [Learn more](https://developers.cloudflare.com/workers/configuration/routing/routes/#matching-behavior).
	Pattern string `json:"pattern,required"`
	// Name of the script to run if the route matches.
	Script string                                  `json:"script"`
	JSON   zoneWorkerRouteUpdateResponseResultJSON `json:"-"`
}

// zoneWorkerRouteUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneWorkerRouteUpdateResponseResult]
type zoneWorkerRouteUpdateResponseResultJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWorkerRouteUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneWorkerRouteUpdateResponseSuccess bool

const (
	ZoneWorkerRouteUpdateResponseSuccessTrue ZoneWorkerRouteUpdateResponseSuccess = true
)

func (r ZoneWorkerRouteUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneWorkerRouteUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneWorkerRouteListResponse struct {
	Errors   []WorkersMessages                   `json:"errors,required"`
	Messages []WorkersMessages                   `json:"messages,required"`
	Result   []ZoneWorkerRouteListResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneWorkerRouteListResponseSuccess `json:"success,required"`
	JSON    zoneWorkerRouteListResponseJSON    `json:"-"`
}

// zoneWorkerRouteListResponseJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteListResponse]
type zoneWorkerRouteListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWorkerRouteListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWorkerRouteListResponseResult struct {
	// Identifier.
	ID string `json:"id,required"`
	// Pattern to match incoming requests against.
	// [Learn more](https://developers.cloudflare.com/workers/configuration/routing/routes/#matching-behavior).
	Pattern string `json:"pattern,required"`
	// Name of the script to run if the route matches.
	Script string                                `json:"script"`
	JSON   zoneWorkerRouteListResponseResultJSON `json:"-"`
}

// zoneWorkerRouteListResponseResultJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteListResponseResult]
type zoneWorkerRouteListResponseResultJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWorkerRouteListResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneWorkerRouteListResponseSuccess bool

const (
	ZoneWorkerRouteListResponseSuccessTrue ZoneWorkerRouteListResponseSuccess = true
)

func (r ZoneWorkerRouteListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneWorkerRouteListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneWorkerRouteDeleteResponse struct {
	Errors   []WorkersMessages                   `json:"errors,required"`
	Messages []WorkersMessages                   `json:"messages,required"`
	Result   ZoneWorkerRouteDeleteResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneWorkerRouteDeleteResponseSuccess `json:"success,required"`
	JSON    zoneWorkerRouteDeleteResponseJSON    `json:"-"`
}

// zoneWorkerRouteDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteDeleteResponse]
type zoneWorkerRouteDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWorkerRouteDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWorkerRouteDeleteResponseResult struct {
	// Identifier.
	ID   string                                  `json:"id"`
	JSON zoneWorkerRouteDeleteResponseResultJSON `json:"-"`
}

// zoneWorkerRouteDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneWorkerRouteDeleteResponseResult]
type zoneWorkerRouteDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWorkerRouteDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneWorkerRouteDeleteResponseSuccess bool

const (
	ZoneWorkerRouteDeleteResponseSuccessTrue ZoneWorkerRouteDeleteResponseSuccess = true
)

func (r ZoneWorkerRouteDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneWorkerRouteDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneWorkerRouteNewParams struct {
	// Pattern to match incoming requests against.
	// [Learn more](https://developers.cloudflare.com/workers/configuration/routing/routes/#matching-behavior).
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script to run if the route matches.
	Script param.Field[string] `json:"script"`
}

func (r ZoneWorkerRouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneWorkerRouteUpdateParams struct {
	// Pattern to match incoming requests against.
	// [Learn more](https://developers.cloudflare.com/workers/configuration/routing/routes/#matching-behavior).
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script to run if the route matches.
	Script param.Field[string] `json:"script"`
}

func (r ZoneWorkerRouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
