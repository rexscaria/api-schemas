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
func (r *ZoneWorkerRouteService) New(ctx context.Context, zoneID string, body ZoneWorkerRouteNewParams, opts ...option.RequestOption) (res *APIResponseSingleWorkers, err error) {
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
func (r *ZoneWorkerRouteService) Get(ctx context.Context, zoneID string, routeID string, opts ...option.RequestOption) (res *WorkersRouteResponseSingle, err error) {
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
func (r *ZoneWorkerRouteService) Update(ctx context.Context, zoneID string, routeID string, body ZoneWorkerRouteUpdateParams, opts ...option.RequestOption) (res *WorkersRouteResponseSingle, err error) {
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
func (r *ZoneWorkerRouteService) Delete(ctx context.Context, zoneID string, routeID string, body ZoneWorkerRouteDeleteParams, opts ...option.RequestOption) (res *APIResponseSingleWorkers, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type APIResponseSingleWorkers struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseSingleWorkersSuccess `json:"success,required"`
	JSON    apiResponseSingleWorkersJSON    `json:"-"`
}

// apiResponseSingleWorkersJSON contains the JSON metadata for the struct
// [APIResponseSingleWorkers]
type apiResponseSingleWorkersJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleWorkers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleWorkersJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseSingleWorkersSuccess bool

const (
	APIResponseSingleWorkersSuccessTrue APIResponseSingleWorkersSuccess = true
)

func (r APIResponseSingleWorkersSuccess) IsKnown() bool {
	switch r {
	case APIResponseSingleWorkersSuccessTrue:
		return true
	}
	return false
}

type RouteNoIDParam struct {
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script param.Field[string] `json:"script"`
}

func (r RouteNoIDParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Routes struct {
	// Identifier
	ID      string `json:"id,required"`
	Pattern string `json:"pattern,required"`
	// Name of the script, used in URLs and route configuration.
	Script string     `json:"script,required"`
	JSON   routesJSON `json:"-"`
}

// routesJSON contains the JSON metadata for the struct [Routes]
type routesJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Routes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routesJSON) RawJSON() string {
	return r.raw
}

type WorkersRouteResponseSingle struct {
	Result Routes                         `json:"result"`
	JSON   workersRouteResponseSingleJSON `json:"-"`
	APIResponseSingleWorkers
}

// workersRouteResponseSingleJSON contains the JSON metadata for the struct
// [WorkersRouteResponseSingle]
type workersRouteResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersRouteResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersRouteResponseSingleJSON) RawJSON() string {
	return r.raw
}

type ZoneWorkerRouteListResponse struct {
	Result []Routes                        `json:"result"`
	JSON   zoneWorkerRouteListResponseJSON `json:"-"`
	CommonResponseWorkers
}

// zoneWorkerRouteListResponseJSON contains the JSON metadata for the struct
// [ZoneWorkerRouteListResponse]
type zoneWorkerRouteListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerRouteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWorkerRouteListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWorkerRouteNewParams struct {
	RouteNoID RouteNoIDParam `json:"route_no_id,required"`
}

func (r ZoneWorkerRouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RouteNoID)
}

type ZoneWorkerRouteUpdateParams struct {
	RouteNoID RouteNoIDParam `json:"route_no_id,required"`
}

func (r ZoneWorkerRouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RouteNoID)
}

type ZoneWorkerRouteDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneWorkerRouteDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
