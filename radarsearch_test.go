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

func TestRadarSearchGlobalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Search.Global(context.TODO(), cfrex.RadarSearchGlobalParams{
		Query:         cfrex.F("United"),
		Exclude:       cfrex.F([]cfrex.RadarSearchGlobalParamsExclude{cfrex.RadarSearchGlobalParamsExcludeSpecialEvents}),
		Format:        cfrex.F(cfrex.RadarSearchGlobalParamsFormatJson),
		Include:       cfrex.F([]cfrex.RadarSearchGlobalParamsInclude{cfrex.RadarSearchGlobalParamsIncludeSpecialEvents}),
		Limit:         cfrex.F(int64(5)),
		LimitPerGroup: cfrex.F(0.000000),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
