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

// RadarDatasetService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarDatasetService] method instead.
type RadarDatasetService struct {
	Options []option.RequestOption
}

// NewRadarDatasetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarDatasetService(opts ...option.RequestOption) (r *RadarDatasetService) {
	r = &RadarDatasetService{}
	r.Options = opts
	return
}

// Retrieves an URL to download a single dataset.
func (r *RadarDatasetService) GetDownloadURL(ctx context.Context, params RadarDatasetGetDownloadURLParams, opts ...option.RequestOption) (res *RadarDatasetGetDownloadURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "radar/datasets/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a list of datasets.
func (r *RadarDatasetService) ListDatasets(ctx context.Context, query RadarDatasetListDatasetsParams, opts ...option.RequestOption) (res *RadarDatasetListDatasetsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "radar/datasets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves the CSV content of a given dataset by alias or ID. When getting the
// content by alias the latest dataset is returned, optionally filtered by the
// latest available at a given date.
func (r *RadarDatasetService) GetCsv(ctx context.Context, alias string, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/csv")}, opts...)
	if alias == "" {
		err = errors.New("missing required alias parameter")
		return nil, err
	}
	path := fmt.Sprintf("radar/datasets/%s", alias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type RadarDatasetGetDownloadURLResponse struct {
	Result RadarDatasetGetDownloadURLResponseResult `json:"result" api:"required"`
	JSON   radarDatasetGetDownloadURLResponseJSON   `json:"-"`
}

// radarDatasetGetDownloadURLResponseJSON contains the JSON metadata for the struct
// [RadarDatasetGetDownloadURLResponse]
type radarDatasetGetDownloadURLResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetGetDownloadURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDatasetGetDownloadURLResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDatasetGetDownloadURLResponseResult struct {
	Dataset RadarDatasetGetDownloadURLResponseResultDataset `json:"dataset" api:"required"`
	JSON    radarDatasetGetDownloadURLResponseResultJSON    `json:"-"`
}

// radarDatasetGetDownloadURLResponseResultJSON contains the JSON metadata for the
// struct [RadarDatasetGetDownloadURLResponseResult]
type radarDatasetGetDownloadURLResponseResultJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetGetDownloadURLResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDatasetGetDownloadURLResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDatasetGetDownloadURLResponseResultDataset struct {
	URL  string                                              `json:"url" api:"required"`
	JSON radarDatasetGetDownloadURLResponseResultDatasetJSON `json:"-"`
}

// radarDatasetGetDownloadURLResponseResultDatasetJSON contains the JSON metadata
// for the struct [RadarDatasetGetDownloadURLResponseResultDataset]
type radarDatasetGetDownloadURLResponseResultDatasetJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetGetDownloadURLResponseResultDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDatasetGetDownloadURLResponseResultDatasetJSON) RawJSON() string {
	return r.raw
}

type RadarDatasetListDatasetsResponse struct {
	Result  RadarDatasetListDatasetsResponseResult `json:"result" api:"required"`
	Success bool                                   `json:"success" api:"required"`
	JSON    radarDatasetListDatasetsResponseJSON   `json:"-"`
}

// radarDatasetListDatasetsResponseJSON contains the JSON metadata for the struct
// [RadarDatasetListDatasetsResponse]
type radarDatasetListDatasetsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetListDatasetsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDatasetListDatasetsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDatasetListDatasetsResponseResult struct {
	Datasets []RadarDatasetListDatasetsResponseResultDataset `json:"datasets" api:"required"`
	JSON     radarDatasetListDatasetsResponseResultJSON      `json:"-"`
}

// radarDatasetListDatasetsResponseResultJSON contains the JSON metadata for the
// struct [RadarDatasetListDatasetsResponseResult]
type radarDatasetListDatasetsResponseResultJSON struct {
	Datasets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetListDatasetsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDatasetListDatasetsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDatasetListDatasetsResponseResultDataset struct {
	ID          int64                                             `json:"id" api:"required"`
	Description string                                            `json:"description" api:"required"`
	Meta        interface{}                                       `json:"meta" api:"required"`
	Tags        []string                                          `json:"tags" api:"required"`
	Title       string                                            `json:"title" api:"required"`
	Type        string                                            `json:"type" api:"required"`
	JSON        radarDatasetListDatasetsResponseResultDatasetJSON `json:"-"`
}

// radarDatasetListDatasetsResponseResultDatasetJSON contains the JSON metadata for
// the struct [RadarDatasetListDatasetsResponseResultDataset]
type radarDatasetListDatasetsResponseResultDatasetJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Meta        apijson.Field
	Tags        apijson.Field
	Title       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetListDatasetsResponseResultDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDatasetListDatasetsResponseResultDatasetJSON) RawJSON() string {
	return r.raw
}

type RadarDatasetGetDownloadURLParams struct {
	DatasetID param.Field[int64] `json:"datasetId" api:"required"`
	// Format in which results will be returned.
	Format param.Field[RadarDatasetGetDownloadURLParamsFormat] `query:"format"`
}

func (r RadarDatasetGetDownloadURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [RadarDatasetGetDownloadURLParams]'s query parameters as
// `url.Values`.
func (r RadarDatasetGetDownloadURLParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarDatasetGetDownloadURLParamsFormat string

const (
	RadarDatasetGetDownloadURLParamsFormatJson RadarDatasetGetDownloadURLParamsFormat = "JSON"
	RadarDatasetGetDownloadURLParamsFormatCsv  RadarDatasetGetDownloadURLParamsFormat = "CSV"
)

func (r RadarDatasetGetDownloadURLParamsFormat) IsKnown() bool {
	switch r {
	case RadarDatasetGetDownloadURLParamsFormatJson, RadarDatasetGetDownloadURLParamsFormatCsv:
		return true
	}
	return false
}

type RadarDatasetListDatasetsParams struct {
	// Filters results by dataset type.
	DatasetType param.Field[RadarDatasetListDatasetsParamsDatasetType] `query:"datasetType"`
	// Filters results by the specified date.
	Date param.Field[time.Time] `query:"date" format:"date"`
	// Format in which results will be returned.
	Format param.Field[RadarDatasetListDatasetsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [RadarDatasetListDatasetsParams]'s query parameters as
// `url.Values`.
func (r RadarDatasetListDatasetsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filters results by dataset type.
type RadarDatasetListDatasetsParamsDatasetType string

const (
	RadarDatasetListDatasetsParamsDatasetTypeRankingBucket RadarDatasetListDatasetsParamsDatasetType = "RANKING_BUCKET"
	RadarDatasetListDatasetsParamsDatasetTypeReport        RadarDatasetListDatasetsParamsDatasetType = "REPORT"
)

func (r RadarDatasetListDatasetsParamsDatasetType) IsKnown() bool {
	switch r {
	case RadarDatasetListDatasetsParamsDatasetTypeRankingBucket, RadarDatasetListDatasetsParamsDatasetTypeReport:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarDatasetListDatasetsParamsFormat string

const (
	RadarDatasetListDatasetsParamsFormatJson RadarDatasetListDatasetsParamsFormat = "JSON"
	RadarDatasetListDatasetsParamsFormatCsv  RadarDatasetListDatasetsParamsFormat = "CSV"
)

func (r RadarDatasetListDatasetsParamsFormat) IsKnown() bool {
	switch r {
	case RadarDatasetListDatasetsParamsFormatJson, RadarDatasetListDatasetsParamsFormatCsv:
		return true
	}
	return false
}
