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

func TestAccountWorkerScriptSettingGet(t *testing.T) {
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
	_, err := client.Accounts.Workers.Scripts.Settings.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is_my_script-01",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountWorkerScriptSettingPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Workers.Scripts.Settings.Patch(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is_my_script-01",
		cfrex.AccountWorkerScriptSettingPatchParams{
			Settings: cfrex.F(cfrex.ScriptVersionItemParam{
				Bindings: cfrex.F([]cfrex.BindingItemUnionParam{cfrex.BindingItemWorkersBindingKindAIParam{
					Name: cfrex.F("MY_ENV_VAR"),
					Type: cfrex.F(cfrex.BindingItemWorkersBindingKindAITypePlainText),
				}}),
				CompatibilityDate:  cfrex.F("2021-01-01"),
				CompatibilityFlags: cfrex.F([]string{"nodejs_compat"}),
				Limits: cfrex.F(cfrex.ScriptVersionItemLimitsParam{
					CPUMs: cfrex.F(int64(50)),
				}),
				Logpush: cfrex.F(false),
				Migrations: cfrex.F[cfrex.ScriptVersionItemMigrationsUnionParam](cfrex.ScriptVersionItemMigrationsWorkersSingleStepMigrationsParam(cfrex.ScriptVersionItemMigrationsWorkersSingleStepMigrationsParam{
					MigrationTagConditionsParam: cfrex.MigrationTagConditionsParam{
						NewTag: cfrex.F("v2"),
						OldTag: cfrex.F("v1"),
					},
					MigrationStepParam: cfrex.MigrationStepParam{
						DeletedClasses:   cfrex.F([]string{"string"}),
						NewClasses:       cfrex.F([]string{"string"}),
						NewSqliteClasses: cfrex.F([]string{"string"}),
						RenamedClasses: cfrex.F([]cfrex.MigrationStepRenamedClassParam{{
							From: cfrex.F("from"),
							To:   cfrex.F("to"),
						}}),
						TransferredClasses: cfrex.F([]cfrex.MigrationStepTransferredClassParam{{
							From:       cfrex.F("from"),
							FromScript: cfrex.F("from_script"),
							To:         cfrex.F("to"),
						}}),
					},
				})),
				Observability: cfrex.F(cfrex.ObservabilityParam{
					Enabled:          cfrex.F(true),
					HeadSamplingRate: cfrex.F(0.100000),
				}),
				Placement: cfrex.F(cfrex.ScriptVersionItemPlacementParam{
					Mode: cfrex.F(cfrex.PlacementModeSmart),
				}),
				Tags: cfrex.F([]string{"my-tag"}),
				TailConsumers: cfrex.F([]cfrex.TailConsumersScriptParam{{
					Service:     cfrex.F("my-log-consumer"),
					Environment: cfrex.F("production"),
					Namespace:   cfrex.F("my-namespace"),
				}}),
				UsageModel: cfrex.F(cfrex.UsageModelStandard),
			}),
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
