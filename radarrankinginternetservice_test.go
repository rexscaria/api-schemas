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

func TestRadarRankingInternetServiceListCategoriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Ranking.InternetServices.ListCategories(context.TODO(), cfrex.RadarRankingInternetServiceListCategoriesParams{
		Date:   cfrex.F([]time.Time{time.Now()}),
		Format: cfrex.F(cfrex.RadarRankingInternetServiceListCategoriesParamsFormatJson),
		Limit:  cfrex.F(int64(5)),
		Name:   cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarRankingInternetServiceGetTimeseriesGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Ranking.InternetServices.GetTimeseriesGroups(context.TODO(), cfrex.RadarRankingInternetServiceGetTimeseriesGroupsParams{
		DateEnd:         cfrex.F([]time.Time{time.Now()}),
		DateRange:       cfrex.F([]string{"7d"}),
		DateStart:       cfrex.F([]time.Time{time.Now()}),
		Format:          cfrex.F(cfrex.RadarRankingInternetServiceGetTimeseriesGroupsParamsFormatJson),
		Limit:           cfrex.F(int64(5)),
		Name:            cfrex.F([]string{"main_series"}),
		ServiceCategory: cfrex.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarRankingInternetServiceGetTopServicesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Ranking.InternetServices.GetTopServices(context.TODO(), cfrex.RadarRankingInternetServiceGetTopServicesParams{
		Date:            cfrex.F([]time.Time{time.Now()}),
		Format:          cfrex.F(cfrex.RadarRankingInternetServiceGetTopServicesParamsFormatJson),
		Limit:           cfrex.F(int64(5)),
		Name:            cfrex.F([]string{"main_series"}),
		ServiceCategory: cfrex.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
