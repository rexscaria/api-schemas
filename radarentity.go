// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// RadarEntityService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarEntityService] method instead.
type RadarEntityService struct {
	Options   []option.RequestOption
	Asns      *RadarEntityAsnService
	Locations *RadarEntityLocationService
}

// NewRadarEntityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarEntityService(opts ...option.RequestOption) (r *RadarEntityService) {
	r = &RadarEntityService{}
	r.Options = opts
	r.Asns = NewRadarEntityAsnService(opts...)
	r.Locations = NewRadarEntityLocationService(opts...)
	return
}

// Retrieves IP address information.
func (r *RadarEntityService) GetIPDetails(ctx context.Context, query RadarEntityGetIPDetailsParams, opts ...option.RequestOption) (res *RadarEntityGetIPDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "radar/entities/ip"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type RadarEntityGetIPDetailsResponse struct {
	Result  RadarEntityGetIPDetailsResponseResult `json:"result" api:"required"`
	Success bool                                  `json:"success" api:"required"`
	JSON    radarEntityGetIPDetailsResponseJSON   `json:"-"`
}

// radarEntityGetIPDetailsResponseJSON contains the JSON metadata for the struct
// [RadarEntityGetIPDetailsResponse]
type radarEntityGetIPDetailsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityGetIPDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityGetIPDetailsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEntityGetIPDetailsResponseResult struct {
	IP   RadarEntityGetIPDetailsResponseResultIP   `json:"ip" api:"required"`
	JSON radarEntityGetIPDetailsResponseResultJSON `json:"-"`
}

// radarEntityGetIPDetailsResponseResultJSON contains the JSON metadata for the
// struct [RadarEntityGetIPDetailsResponseResult]
type radarEntityGetIPDetailsResponseResultJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityGetIPDetailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityGetIPDetailsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEntityGetIPDetailsResponseResultIP struct {
	Asn          string                                      `json:"asn" api:"required"`
	AsnLocation  string                                      `json:"asnLocation" api:"required"`
	AsnName      string                                      `json:"asnName" api:"required"`
	AsnOrgName   string                                      `json:"asnOrgName" api:"required"`
	IP           string                                      `json:"ip" api:"required"`
	IPVersion    string                                      `json:"ipVersion" api:"required"`
	Location     string                                      `json:"location" api:"required"`
	LocationName string                                      `json:"locationName" api:"required"`
	JSON         radarEntityGetIPDetailsResponseResultIPJSON `json:"-"`
}

// radarEntityGetIPDetailsResponseResultIPJSON contains the JSON metadata for the
// struct [RadarEntityGetIPDetailsResponseResultIP]
type radarEntityGetIPDetailsResponseResultIPJSON struct {
	Asn          apijson.Field
	AsnLocation  apijson.Field
	AsnName      apijson.Field
	AsnOrgName   apijson.Field
	IP           apijson.Field
	IPVersion    apijson.Field
	Location     apijson.Field
	LocationName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEntityGetIPDetailsResponseResultIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityGetIPDetailsResponseResultIPJSON) RawJSON() string {
	return r.raw
}

type RadarEntityGetIPDetailsParams struct {
	// IP address.
	IP param.Field[string] `query:"ip" api:"required" format:"ip"`
	// Format in which results will be returned.
	Format param.Field[RadarEntityGetIPDetailsParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityGetIPDetailsParams]'s query parameters as
// `url.Values`.
func (r RadarEntityGetIPDetailsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarEntityGetIPDetailsParamsFormat string

const (
	RadarEntityGetIPDetailsParamsFormatJson RadarEntityGetIPDetailsParamsFormat = "JSON"
	RadarEntityGetIPDetailsParamsFormatCsv  RadarEntityGetIPDetailsParamsFormat = "CSV"
)

func (r RadarEntityGetIPDetailsParamsFormat) IsKnown() bool {
	switch r {
	case RadarEntityGetIPDetailsParamsFormatJson, RadarEntityGetIPDetailsParamsFormatCsv:
		return true
	}
	return false
}
