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

func TestAccountWorkerScriptVersionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Workers.Scripts.Versions.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is_my_script-01",
		cfrex.AccountWorkerScriptVersionListParams{
			Deployable: cfrex.F(true),
			Page:       cfrex.F(int64(0)),
			PerPage:    cfrex.F(int64(0)),
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

func TestAccountWorkerScriptVersionGetDetail(t *testing.T) {
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
	_, err := client.Accounts.Workers.Scripts.Versions.GetDetail(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is_my_script-01",
		"bcf48806-b317-4351-9ee7-36e7d557d4de",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountWorkerScriptVersionUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Workers.Scripts.Versions.Upload(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is_my_script-01",
		cfrex.AccountWorkerScriptVersionUploadParams{
			Metadata: cfrex.F(cfrex.AccountWorkerScriptVersionUploadParamsMetadata{
				MainModule: cfrex.F("worker.js"),
				Annotations: cfrex.F(cfrex.AccountWorkerScriptVersionUploadParamsMetadataAnnotations{
					WorkersMessage: cfrex.F("Fixed worker code."),
					WorkersTag:     cfrex.F("workers/tag"),
				}),
				Bindings: cfrex.F([]cfrex.BindingItemUnionParam{cfrex.BindingItemWorkersBindingKindPlainTextParam{
					Name: cfrex.F("MY_ENV_VAR"),
					Text: cfrex.F("my_data"),
					Type: cfrex.F(cfrex.BindingItemWorkersBindingKindPlainTextTypePlainText),
				}}),
				CompatibilityDate:  cfrex.F("2021-01-01"),
				CompatibilityFlags: cfrex.F([]string{"nodejs_compat"}),
				KeepBindings:       cfrex.F([]string{"string"}),
				UsageModel:         cfrex.F(cfrex.UsageModelStandard),
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
