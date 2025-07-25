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

func TestAccountWorkerServiceEnvironmentSettingGet(t *testing.T) {
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
	_, err := client.Accounts.Workers.Services.Environments.Settings.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"my-worker",
		"production",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountWorkerServiceEnvironmentSettingPatch(t *testing.T) {
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
	_, err := client.Accounts.Workers.Services.Environments.Settings.Patch(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"my-worker",
		"production",
		cfrex.AccountWorkerServiceEnvironmentSettingPatchParams{
			SettingsResponseScriptSettings: cfrex.SettingsResponseScriptSettingsParam{
				Errors: cfrex.F([]cfrex.WorkersMessagesParam{{
					Code:             cfrex.F(int64(1000)),
					Message:          cfrex.F("message"),
					DocumentationURL: cfrex.F("documentation_url"),
					Source: cfrex.F(cfrex.WorkersMessagesSourceParam{
						Pointer: cfrex.F("pointer"),
					}),
				}}),
				Messages: cfrex.F([]cfrex.WorkersMessagesParam{{
					Code:             cfrex.F(int64(1000)),
					Message:          cfrex.F("message"),
					DocumentationURL: cfrex.F("documentation_url"),
					Source: cfrex.F(cfrex.WorkersMessagesSourceParam{
						Pointer: cfrex.F("pointer"),
					}),
				}}),
				Result: cfrex.F(cfrex.SettingsItemParam{
					Logpush: cfrex.F(false),
					Observability: cfrex.F(cfrex.ObservabilityParam{
						Enabled:          cfrex.F(true),
						HeadSamplingRate: cfrex.F(0.100000),
						Logs: cfrex.F(cfrex.ObservabilityLogsParam{
							Enabled:          cfrex.F(true),
							InvocationLogs:   cfrex.F(true),
							HeadSamplingRate: cfrex.F(0.100000),
						}),
					}),
					TailConsumers: cfrex.F([]cfrex.TailConsumersScriptParam{{
						Service:     cfrex.F("my-log-consumer"),
						Environment: cfrex.F("production"),
						Namespace:   cfrex.F("my-namespace"),
					}}),
				}),
				Success: cfrex.F(cfrex.SettingsResponseScriptSettingsSuccessTrue),
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
