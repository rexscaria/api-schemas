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

func TestRadarTrafficAnomalyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.TrafficAnomalies.List(context.TODO(), cfrex.RadarTrafficAnomalyListParams{
		Asn:       cfrex.F(int64(174)),
		DateEnd:   cfrex.F(time.Now()),
		DateRange: cfrex.F("7d"),
		DateStart: cfrex.F(time.Now()),
		Format:    cfrex.F(cfrex.RadarTrafficAnomalyListParamsFormatJson),
		Limit:     cfrex.F(int64(5)),
		Location:  cfrex.F("US"),
		Offset:    cfrex.F(int64(0)),
		Status:    cfrex.F(cfrex.RadarTrafficAnomalyListParamsStatusVerified),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarTrafficAnomalyLocationsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.TrafficAnomalies.Locations(context.TODO(), cfrex.RadarTrafficAnomalyLocationsParams{
		DateEnd:   cfrex.F(time.Now()),
		DateRange: cfrex.F("7d"),
		DateStart: cfrex.F(time.Now()),
		Format:    cfrex.F(cfrex.RadarTrafficAnomalyLocationsParamsFormatJson),
		Limit:     cfrex.F(int64(5)),
		Status:    cfrex.F(cfrex.RadarTrafficAnomalyLocationsParamsStatusVerified),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
