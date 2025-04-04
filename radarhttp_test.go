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

func TestRadarHTTPGetTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.GetTimeseries(context.TODO(), cfrex.RadarHTTPGetTimeseriesParams{
		AggInterval:   cfrex.F(cfrex.RadarHTTPGetTimeseriesParamsAggInterval15m),
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPGetTimeseriesParamsBotClass{cfrex.RadarHTTPGetTimeseriesParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPGetTimeseriesParamsBrowserFamily{cfrex.RadarHTTPGetTimeseriesParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPGetTimeseriesParamsDeviceType{cfrex.RadarHTTPGetTimeseriesParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPGetTimeseriesParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPGetTimeseriesParamsHTTPProtocol{cfrex.RadarHTTPGetTimeseriesParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPGetTimeseriesParamsHTTPVersion{cfrex.RadarHTTPGetTimeseriesParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPGetTimeseriesParamsIPVersion{cfrex.RadarHTTPGetTimeseriesParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Normalization: cfrex.F(cfrex.RadarHTTPGetTimeseriesParamsNormalizationPercentageChange),
		Os:            cfrex.F([]cfrex.RadarHTTPGetTimeseriesParamsO{cfrex.RadarHTTPGetTimeseriesParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPGetTimeseriesParamsTlsVersion{cfrex.RadarHTTPGetTimeseriesParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
