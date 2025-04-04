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

func TestZonePageShieldConnectionGet(t *testing.T) {
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
	_, err := client.Zones.PageShield.Connections.Get(
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

func TestZonePageShieldConnectionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PageShield.Connections.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZonePageShieldConnectionListParams{
			Direction:           cfrex.F(cfrex.ZonePageShieldConnectionListParamsDirectionAsc),
			ExcludeCdnCgi:       cfrex.F(true),
			ExcludeURLs:         cfrex.F("blog.cloudflare.com,www.example"),
			Export:              cfrex.F(cfrex.ZonePageShieldConnectionListParamsExportCsv),
			Hosts:               cfrex.F("blog.cloudflare.com,www.example*,*cloudflare.com"),
			OrderBy:             cfrex.F(cfrex.ZonePageShieldConnectionListParamsOrderByFirstSeenAt),
			Page:                cfrex.F("page"),
			PageURL:             cfrex.F("example.com/page,*/checkout,example.com/*,*checkout*"),
			PerPage:             cfrex.F(100.000000),
			PrioritizeMalicious: cfrex.F(true),
			Status:              cfrex.F("active,inactive"),
			URLs:                cfrex.F("blog.cloudflare.com,www.example"),
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
