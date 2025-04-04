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

func TestRadarAttackLayer7GetTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.GetTimeseries(context.TODO(), cfrex.RadarAttackLayer7GetTimeseriesParams{
		AggInterval:       cfrex.F(cfrex.RadarAttackLayer7GetTimeseriesParamsAggInterval15m),
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7GetTimeseriesParamsFormatJson),
		HTTPMethod:        cfrex.F([]cfrex.RadarAttackLayer7GetTimeseriesParamsHTTPMethod{cfrex.RadarAttackLayer7GetTimeseriesParamsHTTPMethodGet}),
		HTTPVersion:       cfrex.F([]cfrex.RadarAttackLayer7GetTimeseriesParamsHTTPVersion{cfrex.RadarAttackLayer7GetTimeseriesParamsHTTPVersionHttPv1}),
		IPVersion:         cfrex.F([]cfrex.RadarAttackLayer7GetTimeseriesParamsIPVersion{cfrex.RadarAttackLayer7GetTimeseriesParamsIPVersionIPv4}),
		Location:          cfrex.F([]string{"string"}),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7GetTimeseriesParamsMitigationProduct{cfrex.RadarAttackLayer7GetTimeseriesParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
		Normalization:     cfrex.F(cfrex.RadarAttackLayer7GetTimeseriesParamsNormalizationPercentageChange),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
