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

func TestRadarBgpGetTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Bgp.GetTimeseries(context.TODO(), cfrex.RadarBgpGetTimeseriesParams{
		AggInterval: cfrex.F(cfrex.RadarBgpGetTimeseriesParamsAggInterval1h),
		Asn:         cfrex.F([]string{"string"}),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Format:      cfrex.F(cfrex.RadarBgpGetTimeseriesParamsFormatJson),
		Name:        cfrex.F([]string{"main_series"}),
		Prefix:      cfrex.F([]string{"1.1.1.0/24"}),
		UpdateType:  cfrex.F([]cfrex.RadarBgpGetTimeseriesParamsUpdateType{cfrex.RadarBgpGetTimeseriesParamsUpdateTypeAnnouncement}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
