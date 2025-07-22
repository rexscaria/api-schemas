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

func TestRadarVerifiedBotTopBotsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.VerifiedBots.Top.Bots(context.TODO(), cfrex.RadarVerifiedBotTopBotsParams{
		Asn:       cfrex.F([]string{"string"}),
		Continent: cfrex.F([]string{"string"}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Format:    cfrex.F(cfrex.RadarVerifiedBotTopBotsParamsFormatJson),
		Limit:     cfrex.F(int64(5)),
		Location:  cfrex.F([]string{"string"}),
		Name:      cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarVerifiedBotTopCategoriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.VerifiedBots.Top.Categories(context.TODO(), cfrex.RadarVerifiedBotTopCategoriesParams{
		Asn:       cfrex.F([]string{"string"}),
		Continent: cfrex.F([]string{"string"}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Format:    cfrex.F(cfrex.RadarVerifiedBotTopCategoriesParamsFormatJson),
		Limit:     cfrex.F(int64(5)),
		Location:  cfrex.F([]string{"string"}),
		Name:      cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
