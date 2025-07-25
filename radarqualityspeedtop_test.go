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

func TestRadarQualitySpeedTopGetTopAsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Quality.Speed.Top.GetTopAs(context.TODO(), cfrex.RadarQualitySpeedTopGetTopAsParams{
		Asn:       cfrex.F([]string{"string"}),
		Continent: cfrex.F([]string{"string"}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		Format:    cfrex.F(cfrex.RadarQualitySpeedTopGetTopAsParamsFormatJson),
		Limit:     cfrex.F(int64(5)),
		Location:  cfrex.F([]string{"string"}),
		Name:      cfrex.F([]string{"main_series"}),
		OrderBy:   cfrex.F(cfrex.RadarQualitySpeedTopGetTopAsParamsOrderByBandwidthDownload),
		Reverse:   cfrex.F(true),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarQualitySpeedTopGetTopLocationsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Quality.Speed.Top.GetTopLocations(context.TODO(), cfrex.RadarQualitySpeedTopGetTopLocationsParams{
		Asn:       cfrex.F([]string{"string"}),
		Continent: cfrex.F([]string{"string"}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		Format:    cfrex.F(cfrex.RadarQualitySpeedTopGetTopLocationsParamsFormatJson),
		Limit:     cfrex.F(int64(5)),
		Location:  cfrex.F([]string{"string"}),
		Name:      cfrex.F([]string{"main_series"}),
		OrderBy:   cfrex.F(cfrex.RadarQualitySpeedTopGetTopLocationsParamsOrderByBandwidthDownload),
		Reverse:   cfrex.F(true),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
