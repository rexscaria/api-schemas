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

func TestAccountRequestTracerTraceWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.RequestTracer.Trace(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountRequestTracerTraceParams{
			Method: cfrex.F("PUT"),
			URL:    cfrex.F("https://some.zone/some_path"),
			Body: cfrex.F(cfrex.AccountRequestTracerTraceParamsBody{
				Base64:    cfrex.F("c29tZV9yZXF1ZXN0X2JvZHk="),
				Json:      cfrex.F[any](map[string]interface{}{}),
				PlainText: cfrex.F("plain_text"),
			}),
			Context: cfrex.F(cfrex.AccountRequestTracerTraceParamsContext{
				BotScore: cfrex.F(int64(0)),
				Geoloc: cfrex.F(cfrex.AccountRequestTracerTraceParamsContextGeoloc{
					City:                cfrex.F("London"),
					Continent:           cfrex.F("continent"),
					IsEuCountry:         cfrex.F(true),
					ISOCode:             cfrex.F("iso_code"),
					Latitude:            cfrex.F(0.000000),
					Longitude:           cfrex.F(0.000000),
					PostalCode:          cfrex.F("postal_code"),
					RegionCode:          cfrex.F("region_code"),
					Subdivision2ISOCode: cfrex.F("subdivision_2_iso_code"),
					Timezone:            cfrex.F("timezone"),
				}),
				SkipChallenge: cfrex.F(true),
				ThreatScore:   cfrex.F(int64(0)),
			}),
			Cookies: cfrex.F(map[string]string{
				"cookie_name_1": "cookie_value_1",
				"cookie_name_2": "cookie_value_2",
			}),
			Headers: cfrex.F(map[string]string{
				"header_name_1": "header_value_1",
				"header_name_2": "header_value_2",
			}),
			Protocol:     cfrex.F("HTTP/1.1"),
			SkipResponse: cfrex.F(true),
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
