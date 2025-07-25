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

func TestRadarAs112TopLocationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.Top.Locations.Get(context.TODO(), cfrex.RadarAs112TopLocationGetParams{
		Continent: cfrex.F([]string{"string"}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Format:    cfrex.F(cfrex.RadarAs112TopLocationGetParamsFormatJson),
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

func TestRadarAs112TopLocationGetByDnssecWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.Top.Locations.GetByDnssec(
		context.TODO(),
		cfrex.RadarAs112TopLocationGetByDnssecParamsDnssecSupported,
		cfrex.RadarAs112TopLocationGetByDnssecParams{
			Continent: cfrex.F([]string{"string"}),
			DateEnd:   cfrex.F([]time.Time{time.Now()}),
			DateRange: cfrex.F([]string{"7d"}),
			DateStart: cfrex.F([]time.Time{time.Now()}),
			Format:    cfrex.F(cfrex.RadarAs112TopLocationGetByDnssecParamsFormatJson),
			Limit:     cfrex.F(int64(5)),
			Location:  cfrex.F([]string{"string"}),
			Name:      cfrex.F([]string{"main_series"}),
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

func TestRadarAs112TopLocationGetByEdnsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.Top.Locations.GetByEdns(
		context.TODO(),
		cfrex.RadarAs112TopLocationGetByEdnsParamsEdnsSupported,
		cfrex.RadarAs112TopLocationGetByEdnsParams{
			Continent: cfrex.F([]string{"string"}),
			DateEnd:   cfrex.F([]time.Time{time.Now()}),
			DateRange: cfrex.F([]string{"7d"}),
			DateStart: cfrex.F([]time.Time{time.Now()}),
			Format:    cfrex.F(cfrex.RadarAs112TopLocationGetByEdnsParamsFormatJson),
			Limit:     cfrex.F(int64(5)),
			Location:  cfrex.F([]string{"string"}),
			Name:      cfrex.F([]string{"main_series"}),
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

func TestRadarAs112TopLocationGetByIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.Top.Locations.GetByIPVersion(
		context.TODO(),
		cfrex.RadarAs112TopLocationGetByIPVersionParamsIPVersionIPv4,
		cfrex.RadarAs112TopLocationGetByIPVersionParams{
			Continent: cfrex.F([]string{"string"}),
			DateEnd:   cfrex.F([]time.Time{time.Now()}),
			DateRange: cfrex.F([]string{"7d"}),
			DateStart: cfrex.F([]time.Time{time.Now()}),
			Format:    cfrex.F(cfrex.RadarAs112TopLocationGetByIPVersionParamsFormatJson),
			Limit:     cfrex.F(int64(5)),
			Location:  cfrex.F([]string{"string"}),
			Name:      cfrex.F([]string{"main_series"}),
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
