// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/rexscaria/api-schemas"
	"github.com/rexscaria/api-schemas/internal/testutil"
	"github.com/rexscaria/api-schemas/option"
)

func TestRadarAttackLayer7TopLocationGetTopOriginLocationsWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cfrex.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("My API Email"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Radar.Attacks.Layer7.Top.Locations.GetTopOriginLocations(context.TODO(), cfrex.RadarAttackLayer7TopLocationGetTopOriginLocationsParams{
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7TopLocationGetTopOriginLocationsParamsFormatJson),
		HTTPMethod:        cfrex.F([]cfrex.RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod{cfrex.RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodGet}),
		HTTPVersion:       cfrex.F([]cfrex.RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPVersion{cfrex.RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPVersionHttPv1}),
		IPVersion:         cfrex.F([]cfrex.RadarAttackLayer7TopLocationGetTopOriginLocationsParamsIPVersion{cfrex.RadarAttackLayer7TopLocationGetTopOriginLocationsParamsIPVersionIPv4}),
		Limit:             cfrex.F(int64(5)),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProduct{cfrex.RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TopLocationGetTopTargetLocationsWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cfrex.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("My API Email"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Radar.Attacks.Layer7.Top.Locations.GetTopTargetLocations(context.TODO(), cfrex.RadarAttackLayer7TopLocationGetTopTargetLocationsParams{
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7TopLocationGetTopTargetLocationsParamsFormatJson),
		Limit:             cfrex.F(int64(5)),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProduct{cfrex.RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
