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

func TestZoneHealthcheckNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Healthchecks.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneHealthcheckNewParams{
			HealthcheckQuery: cfrex.HealthcheckQueryParam{
				Address:              cfrex.F("www.example.com"),
				Name:                 cfrex.F("server-1"),
				CheckRegions:         cfrex.F([]cfrex.HealthcheckCheckRegion{cfrex.HealthcheckCheckRegionWeu, cfrex.HealthcheckCheckRegionEnam}),
				ConsecutiveFails:     cfrex.F(int64(0)),
				ConsecutiveSuccesses: cfrex.F(int64(0)),
				Description:          cfrex.F("Health check for www.example.com"),
				HTTPConfig: cfrex.F(cfrex.HealthcheckHTTPConfigParam{
					AllowInsecure:   cfrex.F(true),
					ExpectedBody:    cfrex.F("success"),
					ExpectedCodes:   cfrex.F([]string{"2xx", "302"}),
					FollowRedirects: cfrex.F(true),
					Header: cfrex.F(map[string][]string{
						"Host":     {"example.com"},
						"X-App-ID": {"abc123"},
					}),
					Method: cfrex.F(cfrex.HealthcheckHTTPConfigMethodGet),
					Path:   cfrex.F("/health"),
					Port:   cfrex.F(int64(0)),
				}),
				Interval:  cfrex.F(int64(0)),
				Retries:   cfrex.F(int64(0)),
				Suspended: cfrex.F(true),
				TcpConfig: cfrex.F(cfrex.HealthcheckTcpConfigParam{
					Method: cfrex.F(cfrex.HealthcheckTcpConfigMethodConnectionEstablished),
					Port:   cfrex.F(int64(0)),
				}),
				Timeout: cfrex.F(int64(0)),
				Type:    cfrex.F("HTTPS"),
			},
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

func TestZoneHealthcheckGet(t *testing.T) {
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
	_, err := client.Zones.Healthchecks.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneHealthcheckUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Healthchecks.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneHealthcheckUpdateParams{
			HealthcheckQuery: cfrex.HealthcheckQueryParam{
				Address:              cfrex.F("www.example.com"),
				Name:                 cfrex.F("server-1"),
				CheckRegions:         cfrex.F([]cfrex.HealthcheckCheckRegion{cfrex.HealthcheckCheckRegionWeu, cfrex.HealthcheckCheckRegionEnam}),
				ConsecutiveFails:     cfrex.F(int64(0)),
				ConsecutiveSuccesses: cfrex.F(int64(0)),
				Description:          cfrex.F("Health check for www.example.com"),
				HTTPConfig: cfrex.F(cfrex.HealthcheckHTTPConfigParam{
					AllowInsecure:   cfrex.F(true),
					ExpectedBody:    cfrex.F("success"),
					ExpectedCodes:   cfrex.F([]string{"2xx", "302"}),
					FollowRedirects: cfrex.F(true),
					Header: cfrex.F(map[string][]string{
						"Host":     {"example.com"},
						"X-App-ID": {"abc123"},
					}),
					Method: cfrex.F(cfrex.HealthcheckHTTPConfigMethodGet),
					Path:   cfrex.F("/health"),
					Port:   cfrex.F(int64(0)),
				}),
				Interval:  cfrex.F(int64(0)),
				Retries:   cfrex.F(int64(0)),
				Suspended: cfrex.F(true),
				TcpConfig: cfrex.F(cfrex.HealthcheckTcpConfigParam{
					Method: cfrex.F(cfrex.HealthcheckTcpConfigMethodConnectionEstablished),
					Port:   cfrex.F(int64(0)),
				}),
				Timeout: cfrex.F(int64(0)),
				Type:    cfrex.F("HTTPS"),
			},
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

func TestZoneHealthcheckListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Healthchecks.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneHealthcheckListParams{
			Page:    cfrex.F(1.000000),
			PerPage: cfrex.F(5.000000),
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

func TestZoneHealthcheckDelete(t *testing.T) {
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
	_, err := client.Zones.Healthchecks.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneHealthcheckDeleteParams{
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

func TestZoneHealthcheckPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Healthchecks.Patch(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneHealthcheckPatchParams{
			HealthcheckQuery: cfrex.HealthcheckQueryParam{
				Address:              cfrex.F("www.example.com"),
				Name:                 cfrex.F("server-1"),
				CheckRegions:         cfrex.F([]cfrex.HealthcheckCheckRegion{cfrex.HealthcheckCheckRegionWeu, cfrex.HealthcheckCheckRegionEnam}),
				ConsecutiveFails:     cfrex.F(int64(0)),
				ConsecutiveSuccesses: cfrex.F(int64(0)),
				Description:          cfrex.F("Health check for www.example.com"),
				HTTPConfig: cfrex.F(cfrex.HealthcheckHTTPConfigParam{
					AllowInsecure:   cfrex.F(true),
					ExpectedBody:    cfrex.F("success"),
					ExpectedCodes:   cfrex.F([]string{"2xx", "302"}),
					FollowRedirects: cfrex.F(true),
					Header: cfrex.F(map[string][]string{
						"Host":     {"example.com"},
						"X-App-ID": {"abc123"},
					}),
					Method: cfrex.F(cfrex.HealthcheckHTTPConfigMethodGet),
					Path:   cfrex.F("/health"),
					Port:   cfrex.F(int64(0)),
				}),
				Interval:  cfrex.F(int64(0)),
				Retries:   cfrex.F(int64(0)),
				Suspended: cfrex.F(true),
				TcpConfig: cfrex.F(cfrex.HealthcheckTcpConfigParam{
					Method: cfrex.F(cfrex.HealthcheckTcpConfigMethodConnectionEstablished),
					Port:   cfrex.F(int64(0)),
				}),
				Timeout: cfrex.F(int64(0)),
				Type:    cfrex.F("HTTPS"),
			},
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
