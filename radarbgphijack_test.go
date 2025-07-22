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

func TestRadarBgpHijackListEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Bgp.Hijacks.ListEvents(context.TODO(), cfrex.RadarBgpHijackListEventsParams{
		DateEnd:         cfrex.F(time.Now()),
		DateRange:       cfrex.F("7d"),
		DateStart:       cfrex.F(time.Now()),
		EventID:         cfrex.F(int64(0)),
		Format:          cfrex.F(cfrex.RadarBgpHijackListEventsParamsFormatJson),
		HijackerAsn:     cfrex.F(int64(0)),
		InvolvedAsn:     cfrex.F(int64(0)),
		InvolvedCountry: cfrex.F("involvedCountry"),
		MaxConfidence:   cfrex.F(int64(0)),
		MinConfidence:   cfrex.F(int64(0)),
		Page:            cfrex.F(int64(0)),
		PerPage:         cfrex.F(int64(0)),
		Prefix:          cfrex.F("1.1.1.0/24"),
		SortBy:          cfrex.F(cfrex.RadarBgpHijackListEventsParamsSortByTime),
		SortOrder:       cfrex.F(cfrex.RadarBgpHijackListEventsParamsSortOrderAsc),
		VictimAsn:       cfrex.F(int64(0)),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
