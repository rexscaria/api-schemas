// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/rexscaria/api-schemas"
	"github.com/rexscaria/api-schemas/internal/testutil"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
)

func TestZoneAnalyticsListColosWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Analytics.ListColos(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneAnalyticsListColosParams{
			Continuous: cfrex.F(true),
			Since:      cfrex.F[cfrex.ZoneAnalyticsListColosParamsSinceUnion](shared.UnionString("2015-01-01T12:23:00Z")),
			Until:      cfrex.F[cfrex.ZoneAnalyticsUntilUnionParam](shared.UnionString("2015-01-02T12:23:00Z")),
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

func TestZoneAnalyticsGetDashboardWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Analytics.GetDashboard(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneAnalyticsGetDashboardParams{
			Continuous: cfrex.F(true),
			Since:      cfrex.F[cfrex.ZoneAnalyticsGetDashboardParamsSinceUnion](shared.UnionString("2015-01-01T12:23:00Z")),
			Until:      cfrex.F[cfrex.ZoneAnalyticsUntilUnionParam](shared.UnionString("2015-01-02T12:23:00Z")),
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
