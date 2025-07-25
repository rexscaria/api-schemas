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

// ZoneAnalyticsLatencyService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAnalyticsLatencyService] method instead.
type ZoneAnalyticsLatencyService struct {
	Options []option.RequestOption
}

// NewZoneAnalyticsLatencyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAnalyticsLatencyService(opts ...option.RequestOption) (r *ZoneAnalyticsLatencyService) {
	r = &ZoneAnalyticsLatencyService{}
	r.Options = opts
	return
}

// Argo Analytics for a zone
func (r *ZoneAnalyticsLatencyService) Get(ctx context.Context, zoneID string, query ZoneAnalyticsLatencyGetParams, opts ...option.RequestOption) (res *ArgoAnalyticsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/analytics/latency", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Argo Analytics for a zone at different PoPs
func (r *ZoneAnalyticsLatencyService) ListColos(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ArgoAnalyticsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/analytics/latency/colos", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ArgoAnalyticsMessage struct {
	Code             int64                      `json:"code,required"`
	Message          string                     `json:"message,required"`
	DocumentationURL string                     `json:"documentation_url"`
	Source           ArgoAnalyticsMessageSource `json:"source"`
	JSON             argoAnalyticsMessageJSON   `json:"-"`
}

// argoAnalyticsMessageJSON contains the JSON metadata for the struct
// [ArgoAnalyticsMessage]
type argoAnalyticsMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ArgoAnalyticsMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r argoAnalyticsMessageJSON) RawJSON() string {
	return r.raw
}

type ArgoAnalyticsMessageSource struct {
	Pointer string                         `json:"pointer"`
	JSON    argoAnalyticsMessageSourceJSON `json:"-"`
}

// argoAnalyticsMessageSourceJSON contains the JSON metadata for the struct
// [ArgoAnalyticsMessageSource]
type argoAnalyticsMessageSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoAnalyticsMessageSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r argoAnalyticsMessageSourceJSON) RawJSON() string {
	return r.raw
}

type ArgoAnalyticsResponse struct {
	Errors   []ArgoAnalyticsMessage `json:"errors,required"`
	Messages []ArgoAnalyticsMessage `json:"messages,required"`
	Result   interface{}            `json:"result,required"`
	// Whether the API call was successful
	Success ArgoAnalyticsResponseSuccess `json:"success,required"`
	JSON    argoAnalyticsResponseJSON    `json:"-"`
}

// argoAnalyticsResponseJSON contains the JSON metadata for the struct
// [ArgoAnalyticsResponse]
type argoAnalyticsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoAnalyticsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r argoAnalyticsResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ArgoAnalyticsResponseSuccess bool

const (
	ArgoAnalyticsResponseSuccessTrue ArgoAnalyticsResponseSuccess = true
)

func (r ArgoAnalyticsResponseSuccess) IsKnown() bool {
	switch r {
	case ArgoAnalyticsResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAnalyticsLatencyGetParams struct {
	Bins param.Field[string] `query:"bins"`
}

// URLQuery serializes [ZoneAnalyticsLatencyGetParams]'s query parameters as
// `url.Values`.
func (r ZoneAnalyticsLatencyGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
