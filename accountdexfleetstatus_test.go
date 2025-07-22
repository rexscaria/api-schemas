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

func TestAccountDexFleetStatusListDevicesWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Dex.FleetStatus.ListDevices(
		context.TODO(),
		"01a7362d577a6c3019a474fd6f485823",
		cfrex.AccountDexFleetStatusListDevicesParams{
			From:     cfrex.F("2023-10-11T00:00:00Z"),
			Page:     cfrex.F(1.000000),
			PerPage:  cfrex.F(10.000000),
			To:       cfrex.F("2023-10-11T00:00:00Z"),
			Colo:     cfrex.F("SJC"),
			DeviceID: cfrex.F("cb49c27f-7f97-49c5-b6f3-f7c01ead0fd7"),
			Mode:     cfrex.F("proxy"),
			Platform: cfrex.F("windows"),
			SortBy:   cfrex.F(cfrex.AccountDexFleetStatusListDevicesParamsSortByColo),
			Source:   cfrex.F(cfrex.AccountDexFleetStatusListDevicesParamsSourceLastSeen),
			Status:   cfrex.F("connected"),
			Version:  cfrex.F("1.0.0"),
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

func TestAccountDexFleetStatusListLiveStatus(t *testing.T) {
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
	_, err := client.Accounts.Dex.FleetStatus.ListLiveStatus(
		context.TODO(),
		"01a7362d577a6c3019a474fd6f485823",
		cfrex.AccountDexFleetStatusListLiveStatusParams{
			SinceMinutes: cfrex.F(10.000000),
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

func TestAccountDexFleetStatusListOverTimeWithOptionalParams(t *testing.T) {
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
	err := client.Accounts.Dex.FleetStatus.ListOverTime(
		context.TODO(),
		"01a7362d577a6c3019a474fd6f485823",
		cfrex.AccountDexFleetStatusListOverTimeParams{
			From:     cfrex.F("2023-10-11T00:00:00Z"),
			To:       cfrex.F("2023-10-11T00:00:00Z"),
			Colo:     cfrex.F("SJC"),
			DeviceID: cfrex.F("cb49c27f-7f97-49c5-b6f3-f7c01ead0fd7"),
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
