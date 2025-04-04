// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneDNSAnalyticReportService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneDNSAnalyticReportService] method instead.
type ZoneDNSAnalyticReportService struct {
	Options []option.RequestOption
}

// NewZoneDNSAnalyticReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneDNSAnalyticReportService(opts ...option.RequestOption) (r *ZoneDNSAnalyticReportService) {
	r = &ZoneDNSAnalyticReportService{}
	r.Options = opts
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *ZoneDNSAnalyticReportService) Get(ctx context.Context, zoneID string, query ZoneDNSAnalyticReportGetParams, opts ...option.RequestOption) (res *ZoneDNSAnalyticReportGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_analytics/report", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *ZoneDNSAnalyticReportService) ByTime(ctx context.Context, zoneID string, query ZoneDNSAnalyticReportByTimeParams, opts ...option.RequestOption) (res *ZoneDNSAnalyticReportByTimeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_analytics/report/bytime", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneDNSAnalyticReportGetResponse struct {
	Result DataReport                           `json:"result"`
	JSON   zoneDNSAnalyticReportGetResponseJSON `json:"-"`
	APIResponseSingleDNSAnalytics
}

// zoneDNSAnalyticReportGetResponseJSON contains the JSON metadata for the struct
// [ZoneDNSAnalyticReportGetResponse]
type zoneDNSAnalyticReportGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSAnalyticReportGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDNSAnalyticReportGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneDNSAnalyticReportByTimeResponse struct {
	Result ReportByTime                            `json:"result"`
	JSON   zoneDNSAnalyticReportByTimeResponseJSON `json:"-"`
	APIResponseSingleDNSAnalytics
}

// zoneDNSAnalyticReportByTimeResponseJSON contains the JSON metadata for the
// struct [ZoneDNSAnalyticReportByTimeResponse]
type zoneDNSAnalyticReportByTimeResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSAnalyticReportByTimeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDNSAnalyticReportByTimeResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneDNSAnalyticReportGetParams struct {
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

// URLQuery serializes [ZoneDNSAnalyticReportGetParams]'s query parameters as
// `url.Values`.
func (r ZoneDNSAnalyticReportGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneDNSAnalyticReportByTimeParams struct {
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

// URLQuery serializes [ZoneDNSAnalyticReportByTimeParams]'s query parameters as
// `url.Values`.
func (r ZoneDNSAnalyticReportByTimeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
