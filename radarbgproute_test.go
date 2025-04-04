// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/cf-rex-go"
	"github.com/stainless-sdks/cf-rex-go/internal/testutil"
	"github.com/stainless-sdks/cf-rex-go/option"
)

func TestRadarBgpRouteListAsesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Bgp.Routes.ListAses(context.TODO(), cfrex.RadarBgpRouteListAsesParams{
		Format:    cfrex.F(cfrex.RadarBgpRouteListAsesParamsFormatJson),
		Limit:     cfrex.F(int64(5)),
		Location:  cfrex.F("US"),
		SortBy:    cfrex.F(cfrex.RadarBgpRouteListAsesParamsSortByCone),
		SortOrder: cfrex.F(cfrex.RadarBgpRouteListAsesParamsSortOrderAsc),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarBgpRouteListMoasWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Bgp.Routes.ListMoas(context.TODO(), cfrex.RadarBgpRouteListMoasParams{
		Format:      cfrex.F(cfrex.RadarBgpRouteListMoasParamsFormatJson),
		InvalidOnly: cfrex.F(true),
		Origin:      cfrex.F(int64(0)),
		Prefix:      cfrex.F("1.1.1.0/24"),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarBgpRouteGetPrefixToAsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Bgp.Routes.GetPrefixToAs(context.TODO(), cfrex.RadarBgpRouteGetPrefixToAsParams{
		Format:             cfrex.F(cfrex.RadarBgpRouteGetPrefixToAsParamsFormatJson),
		LongestPrefixMatch: cfrex.F(true),
		Origin:             cfrex.F(int64(0)),
		Prefix:             cfrex.F("1.1.1.0/24"),
		RpkiStatus:         cfrex.F(cfrex.RadarBgpRouteGetPrefixToAsParamsRpkiStatusValid),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarBgpRouteGetRealtimeRoutesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Bgp.Routes.GetRealtimeRoutes(context.TODO(), cfrex.RadarBgpRouteGetRealtimeRoutesParams{
		Format: cfrex.F(cfrex.RadarBgpRouteGetRealtimeRoutesParamsFormatJson),
		Prefix: cfrex.F("1.1.1.0/24"),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarBgpRouteGetStatsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Bgp.Routes.GetStats(context.TODO(), cfrex.RadarBgpRouteGetStatsParams{
		Asn:      cfrex.F(int64(174)),
		Format:   cfrex.F(cfrex.RadarBgpRouteGetStatsParamsFormatJson),
		Location: cfrex.F("US"),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
