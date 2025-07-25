// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/rexscaria/api-schemas"
	"github.com/rexscaria/api-schemas/internal/testutil"
	"github.com/rexscaria/api-schemas/option"
)

func TestAccountWorkerScriptList(t *testing.T) {
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
	_, err := client.Accounts.Workers.Scripts.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountWorkerScriptDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Workers.Scripts.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is_my_script-01",
		cfrex.AccountWorkerScriptDeleteParams{
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

func TestAccountWorkerScriptNewAssetsUploadSession(t *testing.T) {
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
	_, err := client.Accounts.Workers.Scripts.NewAssetsUploadSession(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is_my_script-01",
		cfrex.AccountWorkerScriptNewAssetsUploadSessionParams{
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

func TestAccountWorkerScriptDownload(t *testing.T) {
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
	_, err := client.Accounts.Workers.Scripts.Download(
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

func TestAccountWorkerScriptUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Workers.Scripts.Upload(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is_my_script-01",
		cfrex.AccountWorkerScriptUploadParams{
			Metadata: cfrex.F(cfrex.AccountWorkerScriptUploadParamsMetadata{
				Assets: cfrex.F(cfrex.AccountWorkerScriptUploadParamsMetadataAssets{
					Config: cfrex.F(cfrex.AccountWorkerScriptUploadParamsMetadataAssetsConfig{
						Headers:          cfrex.F("/dashboard/*\nX-Frame-Options: DENY\n\n/static/*\nAccess-Control-Allow-Origin: *"),
						Redirects:        cfrex.F("/foo /bar 301\n/news/* /blog/:splat"),
						HTMLHandling:     cfrex.F(cfrex.AccountWorkerScriptUploadParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash),
						NotFoundHandling: cfrex.F(cfrex.AccountWorkerScriptUploadParamsMetadataAssetsConfigNotFoundHandling404Page),
						RunWorkerFirst:   cfrex.F[cfrex.AccountWorkerScriptUploadParamsMetadataAssetsConfigRunWorkerFirstUnion](cfrex.AccountWorkerScriptUploadParamsMetadataAssetsConfigRunWorkerFirstArray([]string{"string"})),
						ServeDirectly:    cfrex.F(true),
					}),
					Jwt: cfrex.F("jwt"),
				}),
				Bindings: cfrex.F([]cfrex.BindingItemUnionParam{cfrex.BindingItemWorkersBindingKindPlainTextParam{
					Name: cfrex.F("MY_ENV_VAR"),
					Text: cfrex.F("my_data"),
					Type: cfrex.F(cfrex.BindingItemWorkersBindingKindPlainTextTypePlainText),
				}}),
				BodyPart:           cfrex.F("worker.js"),
				CompatibilityDate:  cfrex.F("2021-01-01"),
				CompatibilityFlags: cfrex.F([]string{"nodejs_compat"}),
				KeepAssets:         cfrex.F(false),
				KeepBindings:       cfrex.F([]string{"string"}),
				Logpush:            cfrex.F(false),
				MainModule:         cfrex.F("worker.js"),
				Migrations: cfrex.F[cfrex.AccountWorkerScriptUploadParamsMetadataMigrationsUnion](cfrex.AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrations{
					DeletedClasses:   cfrex.F([]string{"string"}),
					NewClasses:       cfrex.F([]string{"string"}),
					NewSqliteClasses: cfrex.F([]string{"string"}),
					NewTag:           cfrex.F("v2"),
					OldTag:           cfrex.F("v1"),
					RenamedClasses: cfrex.F([]cfrex.AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass{{
						From: cfrex.F("from"),
						To:   cfrex.F("to"),
					}}),
					TransferredClasses: cfrex.F([]cfrex.AccountWorkerScriptUploadParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass{{
						From:       cfrex.F("from"),
						FromScript: cfrex.F("from_script"),
						To:         cfrex.F("to"),
					}}),
				}),
				Observability: cfrex.F(cfrex.ObservabilityParam{
					Enabled:          cfrex.F(true),
					HeadSamplingRate: cfrex.F(0.100000),
					Logs: cfrex.F(cfrex.ObservabilityLogsParam{
						Enabled:          cfrex.F(true),
						InvocationLogs:   cfrex.F(true),
						HeadSamplingRate: cfrex.F(0.100000),
					}),
				}),
				Placement: cfrex.F(cfrex.AccountWorkerScriptUploadParamsMetadataPlacement{
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
			Files: cfrex.F([]io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents")))}),
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
