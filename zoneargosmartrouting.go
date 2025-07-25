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

// ZoneArgoSmartRoutingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneArgoSmartRoutingService] method instead.
type ZoneArgoSmartRoutingService struct {
	Options []option.RequestOption
}

// NewZoneArgoSmartRoutingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneArgoSmartRoutingService(opts ...option.RequestOption) (r *ZoneArgoSmartRoutingService) {
	r = &ZoneArgoSmartRoutingService{}
	r.Options = opts
	return
}

// Retrieves the value of Argo Smart Routing enablement setting.
func (r *ZoneArgoSmartRoutingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ArgoConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/argo/smart_routing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Configures the value of the Argo Smart Routing enablement setting.
func (r *ZoneArgoSmartRoutingService) Update(ctx context.Context, zoneID string, body ZoneArgoSmartRoutingUpdateParams, opts ...option.RequestOption) (res *ArgoConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/argo/smart_routing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ArgoConfigMessage struct {
	Code             int64                   `json:"code,required"`
	Message          string                  `json:"message,required"`
	DocumentationURL string                  `json:"documentation_url"`
	Source           ArgoConfigMessageSource `json:"source"`
	JSON             argoConfigMessageJSON   `json:"-"`
}

// argoConfigMessageJSON contains the JSON metadata for the struct
// [ArgoConfigMessage]
type argoConfigMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ArgoConfigMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r argoConfigMessageJSON) RawJSON() string {
	return r.raw
}

type ArgoConfigMessageSource struct {
	Pointer string                      `json:"pointer"`
	JSON    argoConfigMessageSourceJSON `json:"-"`
}

// argoConfigMessageSourceJSON contains the JSON metadata for the struct
// [ArgoConfigMessageSource]
type argoConfigMessageSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoConfigMessageSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r argoConfigMessageSourceJSON) RawJSON() string {
	return r.raw
}

type ArgoConfigResponse struct {
	Errors   []ArgoConfigMessage `json:"errors,required"`
	Messages []ArgoConfigMessage `json:"messages,required"`
	// Describes a successful API response.
	Success ArgoConfigResponseSuccess `json:"success,required"`
	Result  interface{}               `json:"result"`
	JSON    argoConfigResponseJSON    `json:"-"`
}

// argoConfigResponseJSON contains the JSON metadata for the struct
// [ArgoConfigResponse]
type argoConfigResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r argoConfigResponseJSON) RawJSON() string {
	return r.raw
}

// Describes a successful API response.
type ArgoConfigResponseSuccess bool

const (
	ArgoConfigResponseSuccessTrue ArgoConfigResponseSuccess = true
)

func (r ArgoConfigResponseSuccess) IsKnown() bool {
	switch r {
	case ArgoConfigResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneArgoSmartRoutingUpdateParams struct {
	// Enables Argo Smart Routing.
	Value param.Field[ZoneArgoSmartRoutingUpdateParamsValue] `json:"value,required"`
}

func (r ZoneArgoSmartRoutingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Argo Smart Routing.
type ZoneArgoSmartRoutingUpdateParamsValue string

const (
	ZoneArgoSmartRoutingUpdateParamsValueOn  ZoneArgoSmartRoutingUpdateParamsValue = "on"
	ZoneArgoSmartRoutingUpdateParamsValueOff ZoneArgoSmartRoutingUpdateParamsValue = "off"
)

func (r ZoneArgoSmartRoutingUpdateParamsValue) IsKnown() bool {
	switch r {
	case ZoneArgoSmartRoutingUpdateParamsValueOn, ZoneArgoSmartRoutingUpdateParamsValueOff:
		return true
	}
	return false
}
