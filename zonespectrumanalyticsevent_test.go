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

func TestZoneSpectrumAnalyticsEventGetByTimeWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Spectrum.Analytics.Events.GetByTime(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneSpectrumAnalyticsEventGetByTimeParams{
			TimeDelta:  cfrex.F(cfrex.ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaYear),
			Dimensions: cfrex.F([]cfrex.DimensionItem{cfrex.DimensionItemEvent, cfrex.DimensionItemAppID}),
			Filters:    cfrex.F("event==disconnect%20AND%20coloName!=SFO"),
			Metrics:    cfrex.F([]cfrex.MetricItem{cfrex.MetricItemCount, cfrex.MetricItemBytesIngress}),
			Since:      cfrex.F(time.Now()),
			Sort:       cfrex.F([]string{"+count", "-bytesIngress"}),
			Until:      cfrex.F(time.Now()),
		},
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneSpectrumAnalyticsEventGetSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Spectrum.Analytics.Events.GetSummary(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneSpectrumAnalyticsEventGetSummaryParams{
			Dimensions: cfrex.F([]cfrex.DimensionItem{cfrex.DimensionItemEvent, cfrex.DimensionItemAppID}),
			Filters:    cfrex.F("event==disconnect%20AND%20coloName!=SFO"),
			Metrics:    cfrex.F([]cfrex.MetricItem{cfrex.MetricItemCount, cfrex.MetricItemBytesIngress}),
			Since:      cfrex.F(time.Now()),
			Sort:       cfrex.F([]string{"+count", "-bytesIngress"}),
			Until:      cfrex.F(time.Now()),
		},
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
