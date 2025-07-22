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
)

func TestZoneRateLimitNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.RateLimits.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneRateLimitNewParams{
			Action: cfrex.F(cfrex.FirewallActionParam{
				Mode: cfrex.F(cfrex.FirewallActionModeChallenge),
				Response: cfrex.F(cfrex.FirewallActionResponseParam{
					Body:        cfrex.F("<error>This request has been rate-limited.</error>"),
					ContentType: cfrex.F("text/xml"),
				}),
				Timeout: cfrex.F(86400.000000),
			}),
			Match: cfrex.F(cfrex.FirewallMatchParam{
				Headers: cfrex.F([]cfrex.FirewallMatchHeaderParam{{
					Name:  cfrex.F("Cf-Cache-Status"),
					Op:    cfrex.F(cfrex.FirewallMatchHeadersOpNe),
					Value: cfrex.F("HIT"),
				}}),
				Request: cfrex.F(cfrex.FirewallMatchRequestParam{
					Methods: cfrex.F([]cfrex.FirewallMatchRequestMethod{cfrex.FirewallMatchRequestMethodGet, cfrex.FirewallMatchRequestMethodPost}),
					Schemes: cfrex.F([]string{"HTTP", "HTTPS"}),
					URL:     cfrex.F("*.example.org/path*"),
				}),
				Response: cfrex.F(cfrex.FirewallMatchResponseParam{
					OriginTraffic: cfrex.F(true),
				}),
			}),
			Period:    cfrex.F(900.000000),
			Threshold: cfrex.F(60.000000),
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

func TestZoneRateLimitGet(t *testing.T) {
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
	_, err := client.Zones.RateLimits.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"372e67954025e0ba6aaa6d586b9e0b59",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneRateLimitUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.RateLimits.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"372e67954025e0ba6aaa6d586b9e0b59",
		cfrex.ZoneRateLimitUpdateParams{
			Action: cfrex.F(cfrex.FirewallActionParam{
				Mode: cfrex.F(cfrex.FirewallActionModeChallenge),
				Response: cfrex.F(cfrex.FirewallActionResponseParam{
					Body:        cfrex.F("<error>This request has been rate-limited.</error>"),
					ContentType: cfrex.F("text/xml"),
				}),
				Timeout: cfrex.F(86400.000000),
			}),
			Match: cfrex.F(cfrex.FirewallMatchParam{
				Headers: cfrex.F([]cfrex.FirewallMatchHeaderParam{{
					Name:  cfrex.F("Cf-Cache-Status"),
					Op:    cfrex.F(cfrex.FirewallMatchHeadersOpNe),
					Value: cfrex.F("HIT"),
				}}),
				Request: cfrex.F(cfrex.FirewallMatchRequestParam{
					Methods: cfrex.F([]cfrex.FirewallMatchRequestMethod{cfrex.FirewallMatchRequestMethodGet, cfrex.FirewallMatchRequestMethodPost}),
					Schemes: cfrex.F([]string{"HTTP", "HTTPS"}),
					URL:     cfrex.F("*.example.org/path*"),
				}),
				Response: cfrex.F(cfrex.FirewallMatchResponseParam{
					OriginTraffic: cfrex.F(true),
				}),
			}),
			Period:    cfrex.F(900.000000),
			Threshold: cfrex.F(60.000000),
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

func TestZoneRateLimitListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.RateLimits.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneRateLimitListParams{
			Page:    cfrex.F(1.000000),
			PerPage: cfrex.F(1.000000),
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

func TestZoneRateLimitDelete(t *testing.T) {
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
	_, err := client.Zones.RateLimits.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"372e67954025e0ba6aaa6d586b9e0b59",
		cfrex.ZoneRateLimitDeleteParams{
			Body: map[string]interface{}{},
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
