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

func TestRadarHTTPTopGetTopBrowserFamiliesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.GetTopBrowserFamilies(context.TODO(), cfrex.RadarHTTPTopGetTopBrowserFamiliesParams{
		Asn:          cfrex.F([]string{"string"}),
		BotClass:     cfrex.F([]cfrex.RadarHTTPTopGetTopBrowserFamiliesParamsBotClass{cfrex.RadarHTTPTopGetTopBrowserFamiliesParamsBotClassLikelyAutomated}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		DeviceType:   cfrex.F([]cfrex.RadarHTTPTopGetTopBrowserFamiliesParamsDeviceType{cfrex.RadarHTTPTopGetTopBrowserFamiliesParamsDeviceTypeDesktop}),
		Format:       cfrex.F(cfrex.RadarHTTPTopGetTopBrowserFamiliesParamsFormatJson),
		HTTPProtocol: cfrex.F([]cfrex.RadarHTTPTopGetTopBrowserFamiliesParamsHTTPProtocol{cfrex.RadarHTTPTopGetTopBrowserFamiliesParamsHTTPProtocolHTTP}),
		HTTPVersion:  cfrex.F([]cfrex.RadarHTTPTopGetTopBrowserFamiliesParamsHTTPVersion{cfrex.RadarHTTPTopGetTopBrowserFamiliesParamsHTTPVersionHttPv1}),
		IPVersion:    cfrex.F([]cfrex.RadarHTTPTopGetTopBrowserFamiliesParamsIPVersion{cfrex.RadarHTTPTopGetTopBrowserFamiliesParamsIPVersionIPv4}),
		Limit:        cfrex.F(int64(5)),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Os:           cfrex.F([]cfrex.RadarHTTPTopGetTopBrowserFamiliesParamsO{cfrex.RadarHTTPTopGetTopBrowserFamiliesParamsOWindows}),
		TlsVersion:   cfrex.F([]cfrex.RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersion{cfrex.RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTopGetTopBrowsersWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.GetTopBrowsers(context.TODO(), cfrex.RadarHTTPTopGetTopBrowsersParams{
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPTopGetTopBrowsersParamsBotClass{cfrex.RadarHTTPTopGetTopBrowsersParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopGetTopBrowsersParamsBrowserFamily{cfrex.RadarHTTPTopGetTopBrowsersParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPTopGetTopBrowsersParamsDeviceType{cfrex.RadarHTTPTopGetTopBrowsersParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPTopGetTopBrowsersParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTopGetTopBrowsersParamsHTTPProtocol{cfrex.RadarHTTPTopGetTopBrowsersParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTopGetTopBrowsersParamsHTTPVersion{cfrex.RadarHTTPTopGetTopBrowsersParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPTopGetTopBrowsersParamsIPVersion{cfrex.RadarHTTPTopGetTopBrowsersParamsIPVersionIPv4}),
		Limit:         cfrex.F(int64(5)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPTopGetTopBrowsersParamsO{cfrex.RadarHTTPTopGetTopBrowsersParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPTopGetTopBrowsersParamsTlsVersion{cfrex.RadarHTTPTopGetTopBrowsersParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
