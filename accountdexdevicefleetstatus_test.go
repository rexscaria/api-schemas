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

func TestAccountDexDeviceFleetStatusGetLiveStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Dex.Devices.FleetStatus.GetLiveStatus(
		context.TODO(),
		"01a7362d577a6c3019a474fd6f485823",
		"cb49c27f-7f97-49c5-b6f3-f7c01ead0fd7",
		cfrex.AccountDexDeviceFleetStatusGetLiveStatusParams{
			SinceMinutes: cfrex.F(10.000000),
			Colo:         cfrex.F("SJC"),
			TimeNow:      cfrex.F("2023-10-11T00:00:00Z"),
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
