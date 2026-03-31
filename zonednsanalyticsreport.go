// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneDNSAnalyticsReportService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneDNSAnalyticsReportService] method instead.
type ZoneDNSAnalyticsReportService struct {
	Options []option.RequestOption
}

// NewZoneDNSAnalyticsReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneDNSAnalyticsReportService(opts ...option.RequestOption) (r *ZoneDNSAnalyticsReportService) {
	r = &ZoneDNSAnalyticsReportService{}
	r.Options = opts
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *ZoneDNSAnalyticsReportService) Get(ctx context.Context, zoneID string, query ZoneDNSAnalyticsReportGetParams, opts ...option.RequestOption) (res *ZoneDNSAnalyticsReportGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/dns_analytics/report", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves a list of aggregate metrics grouped by time interval.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *ZoneDNSAnalyticsReportService) ByTime(ctx context.Context, zoneID string, query ZoneDNSAnalyticsReportByTimeParams, opts ...option.RequestOption) (res *ZoneDNSAnalyticsReportByTimeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/dns_analytics/report/bytime", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ZoneDNSAnalyticsReportGetResponse struct {
	Errors   []MessagesDNSAnalyticsItem `json:"errors" api:"required"`
	Messages []MessagesDNSAnalyticsItem `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success ZoneDNSAnalyticsReportGetResponseSuccess `json:"success" api:"required"`
	Result  DataReport                               `json:"result"`
	JSON    zoneDNSAnalyticsReportGetResponseJSON    `json:"-"`
}

// zoneDNSAnalyticsReportGetResponseJSON contains the JSON metadata for the struct
// [ZoneDNSAnalyticsReportGetResponse]
type zoneDNSAnalyticsReportGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSAnalyticsReportGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDNSAnalyticsReportGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneDNSAnalyticsReportGetResponseSuccess bool

const (
	ZoneDNSAnalyticsReportGetResponseSuccessTrue ZoneDNSAnalyticsReportGetResponseSuccess = true
)

func (r ZoneDNSAnalyticsReportGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneDNSAnalyticsReportGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneDNSAnalyticsReportByTimeResponse struct {
	Errors   []MessagesDNSAnalyticsItem `json:"errors" api:"required"`
	Messages []MessagesDNSAnalyticsItem `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success ZoneDNSAnalyticsReportByTimeResponseSuccess `json:"success" api:"required"`
	Result  ReportByTime                                `json:"result"`
	JSON    zoneDNSAnalyticsReportByTimeResponseJSON    `json:"-"`
}

// zoneDNSAnalyticsReportByTimeResponseJSON contains the JSON metadata for the
// struct [ZoneDNSAnalyticsReportByTimeResponse]
type zoneDNSAnalyticsReportByTimeResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSAnalyticsReportByTimeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDNSAnalyticsReportByTimeResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneDNSAnalyticsReportByTimeResponseSuccess bool

const (
	ZoneDNSAnalyticsReportByTimeResponseSuccessTrue ZoneDNSAnalyticsReportByTimeResponseSuccess = true
)

func (r ZoneDNSAnalyticsReportByTimeResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneDNSAnalyticsReportByTimeResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneDNSAnalyticsReportGetParams struct {
	// A comma-separated list of dimensions to group results by.
	Dimensions param.Field[string] `query:"dimensions"`
	// Segmentation filter in 'attribute operator value' format.
	Filters param.Field[string] `query:"filters"`
	// Limit number of returned metrics.
	Limit param.Field[int64] `query:"limit"`
	// A comma-separated list of metrics to query.
	Metrics param.Field[string] `query:"metrics"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// A comma-separated list of dimensions to sort by, where each dimension may be
	// prefixed by - (descending) or + (ascending).
	Sort param.Field[string] `query:"sort"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [ZoneDNSAnalyticsReportGetParams]'s query parameters as
// `url.Values`.
func (r ZoneDNSAnalyticsReportGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneDNSAnalyticsReportByTimeParams struct {
	// A comma-separated list of dimensions to group results by.
	Dimensions param.Field[string] `query:"dimensions"`
	// Segmentation filter in 'attribute operator value' format.
	Filters param.Field[string] `query:"filters"`
	// Limit number of returned metrics.
	Limit param.Field[int64] `query:"limit"`
	// A comma-separated list of metrics to query.
	Metrics param.Field[string] `query:"metrics"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// A comma-separated list of dimensions to sort by, where each dimension may be
	// prefixed by - (descending) or + (ascending).
	Sort param.Field[string] `query:"sort"`
	// Unit of time to group data by.
	TimeDelta param.Field[TimeDelta] `query:"time_delta"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [ZoneDNSAnalyticsReportByTimeParams]'s query parameters as
// `url.Values`.
func (r ZoneDNSAnalyticsReportByTimeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
