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

func TestZonePageShieldCookieGet(t *testing.T) {
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
	_, err := client.Zones.PageShield.Cookies.Get(
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

func TestZonePageShieldCookieListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.PageShield.Cookies.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZonePageShieldCookieListParams{
			Direction: cfrex.F(cfrex.ZonePageShieldCookieListParamsDirectionAsc),
			Domain:    cfrex.F("example.com"),
			Export:    cfrex.F(cfrex.ZonePageShieldCookieListParamsExportCsv),
			Hosts:     cfrex.F("blog.cloudflare.com,www.example*,*cloudflare.com"),
			HTTPOnly:  cfrex.F(true),
			Name:      cfrex.F("session_id"),
			OrderBy:   cfrex.F(cfrex.ZonePageShieldCookieListParamsOrderByFirstSeenAt),
			Page:      cfrex.F("page"),
			PageURL:   cfrex.F("example.com/page,*/checkout,example.com/*,*checkout*"),
			Path:      cfrex.F("/"),
			PerPage:   cfrex.F(100.000000),
			SameSite:  cfrex.F(cfrex.ZonePageShieldCookieListParamsSameSiteStrict),
			Secure:    cfrex.F(true),
			Type:      cfrex.F(cfrex.ZonePageShieldCookieListParamsTypeFirstParty),
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
