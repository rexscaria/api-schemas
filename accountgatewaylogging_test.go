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

func TestAccountGatewayLoggingGet(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Logging.Get(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountGatewayLoggingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Logging.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.AccountGatewayLoggingUpdateParams{
			GatewayAccountLoggingSettings: cfrex.GatewayAccountLoggingSettingsParam{
				RedactPii: cfrex.F(true),
				SettingsByRuleType: cfrex.F(cfrex.GatewayAccountLoggingSettingsSettingsByRuleTypeParam{
					DNS: cfrex.F(cfrex.AccountLogOptionsParam{
						LogAll:    cfrex.F(false),
						LogBlocks: cfrex.F(true),
					}),
					HTTP: cfrex.F(cfrex.AccountLogOptionsParam{
						LogAll:    cfrex.F(false),
						LogBlocks: cfrex.F(true),
					}),
					L4: cfrex.F(cfrex.AccountLogOptionsParam{
						LogAll:    cfrex.F(false),
						LogBlocks: cfrex.F(true),
					}),
				}),
			},
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
