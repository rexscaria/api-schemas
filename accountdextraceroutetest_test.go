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

func TestAccountDexTracerouteTestGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Dex.TracerouteTests.Get(
		context.TODO(),
		"01a7362d577a6c3019a474fd6f485823",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.AccountDexTracerouteTestGetParams{
			From:     cfrex.F("from"),
			Interval: cfrex.F(cfrex.AccountDexTracerouteTestGetParamsIntervalMinute),
			To:       cfrex.F("to"),
			Colo:     cfrex.F("colo"),
			DeviceID: cfrex.F([]string{"string"}),
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

func TestAccountDexTracerouteTestGetNetworkPath(t *testing.T) {
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
	_, err := client.Accounts.Dex.TracerouteTests.GetNetworkPath(
		context.TODO(),
		"01a7362d577a6c3019a474fd6f485823",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.AccountDexTracerouteTestGetNetworkPathParams{
			DeviceID: cfrex.F("deviceId"),
			From:     cfrex.F("from"),
			Interval: cfrex.F(cfrex.AccountDexTracerouteTestGetNetworkPathParamsIntervalMinute),
			To:       cfrex.F("to"),
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

func TestAccountDexTracerouteTestGetPercentilesWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Dex.TracerouteTests.GetPercentiles(
		context.TODO(),
		"01a7362d577a6c3019a474fd6f485823",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.AccountDexTracerouteTestGetPercentilesParams{
			From:     cfrex.F("2023-09-20T17:00:00Z"),
			To:       cfrex.F("2023-09-20T17:00:00Z"),
			Colo:     cfrex.F("colo"),
			DeviceID: cfrex.F([]string{"string"}),
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
