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

func TestAccountWorkerDispatchNamespaceScriptGet(t *testing.T) {
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
	_, err := client.Accounts.Workers.Dispatch.Namespaces.Scripts.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"my-dispatch-namespace",
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

func TestAccountWorkerDispatchNamespaceScriptDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Accounts.Workers.Dispatch.Namespaces.Scripts.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"my-dispatch-namespace",
		"this-is_my_script-01",
		cfrex.AccountWorkerDispatchNamespaceScriptDeleteParams{
			Body:  map[string]interface{}{},
			Force: cfrex.F(true),
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

func TestAccountWorkerDispatchNamespaceScriptNewAssetsUploadSessionWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Workers.Dispatch.Namespaces.Scripts.NewAssetsUploadSession(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"my-dispatch-namespace",
		"this-is_my_script-01",
		cfrex.AccountWorkerDispatchNamespaceScriptNewAssetsUploadSessionParams{
			UploadSessionObject: cfrex.UploadSessionObjectParam{
				Manifest: cfrex.F(map[string]cfrex.UploadSessionObjectManifestParam{
					"foo": {
						Hash: cfrex.F("hash"),
						Size: cfrex.F(int64(0)),
					},
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

func TestAccountWorkerDispatchNamespaceScriptGetBindings(t *testing.T) {
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
	_, err := client.Accounts.Workers.Dispatch.Namespaces.Scripts.GetBindings(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"my-dispatch-namespace",
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

func TestAccountWorkerDispatchNamespaceScriptUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Workers.Dispatch.Namespaces.Scripts.Upload(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"my-dispatch-namespace",
		"this-is_my_script-01",
		cfrex.AccountWorkerDispatchNamespaceScriptUploadParams{
			Metadata: cfrex.F(cfrex.AccountWorkerDispatchNamespaceScriptUploadParamsMetadata{
				Assets: cfrex.F(cfrex.AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssets{
					Config: cfrex.F(cfrex.AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfig{
						Headers:          cfrex.F("/dashboard/*\nX-Frame-Options: DENY\n\n/static/*\nAccess-Control-Allow-Origin: *"),
						Redirects:        cfrex.F("/foo /bar 301\n/news/* /blog/:splat"),
						HTMLHandling:     cfrex.F(cfrex.AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash),
						NotFoundHandling: cfrex.F(cfrex.AccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigNotFoundHandlingNone),
						RunWorkerFirst:   cfrex.F(false),
						ServeDirectly:    cfrex.F(true),
					}),
					Jwt: cfrex.F("jwt"),
				}),
				Bindings: cfrex.F([]cfrex.BindingItemUnionParam{cfrex.BindingItemWorkersBindingKindAIParam{
					Name: cfrex.F("MY_ENV_VAR"),
					Type: cfrex.F(cfrex.BindingItemWorkersBindingKindAITypeAI),
				}}),
				BodyPart:           cfrex.F("worker.js"),
				CompatibilityDate:  cfrex.F("2021-01-01"),
				CompatibilityFlags: cfrex.F([]string{"nodejs_compat"}),
				KeepAssets:         cfrex.F(false),
				KeepBindings:       cfrex.F([]string{"string"}),
				Logpush:            cfrex.F(false),
				MainModule:         cfrex.F("worker.js"),
				Migrations: cfrex.F[cfrex.AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsUnion](cfrex.AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations(cfrex.AccountWorkerDispatchNamespaceScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations{
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
				Placement: cfrex.F(cfrex.AccountWorkerDispatchNamespaceScriptUploadParamsMetadataPlacement{
					Mode: cfrex.F(cfrex.PlacementModeSmart),
				}),
				Tags: cfrex.F([]string{"string"}),
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
