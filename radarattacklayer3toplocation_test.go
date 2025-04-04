// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/cf-rex-go"
	"github.com/stainless-sdks/cf-rex-go/internal/testutil"
	"github.com/stainless-sdks/cf-rex-go/option"
)

func TestRadarAttackLayer3TopLocationGetTopOriginLocationsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.Locations.GetTopOriginLocations(context.TODO(), cfrex.RadarAttackLayer3TopLocationGetTopOriginLocationsParams{
		Continent: cfrex.F([]string{"string"}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Format:    cfrex.F(cfrex.RadarAttackLayer3TopLocationGetTopOriginLocationsParamsFormatJson),
		IPVersion: cfrex.F([]cfrex.RadarAttackLayer3TopLocationGetTopOriginLocationsParamsIPVersion{cfrex.RadarAttackLayer3TopLocationGetTopOriginLocationsParamsIPVersionIPv4}),
		Limit:     cfrex.F(int64(5)),
		Location:  cfrex.F([]string{"string"}),
		Name:      cfrex.F([]string{"main_series"}),
		Protocol:  cfrex.F([]cfrex.RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocol{cfrex.RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3TopLocationGetTopTargetLocationsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.Locations.GetTopTargetLocations(context.TODO(), cfrex.RadarAttackLayer3TopLocationGetTopTargetLocationsParams{
		Continent: cfrex.F([]string{"string"}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Format:    cfrex.F(cfrex.RadarAttackLayer3TopLocationGetTopTargetLocationsParamsFormatJson),
		IPVersion: cfrex.F([]cfrex.RadarAttackLayer3TopLocationGetTopTargetLocationsParamsIPVersion{cfrex.RadarAttackLayer3TopLocationGetTopTargetLocationsParamsIPVersionIPv4}),
		Limit:     cfrex.F(int64(5)),
		Location:  cfrex.F([]string{"string"}),
		Name:      cfrex.F([]string{"main_series"}),
		Protocol:  cfrex.F([]cfrex.RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocol{cfrex.RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
