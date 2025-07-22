// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/cf-rex-go"
	"github.com/stainless-sdks/cf-rex-go/internal/testutil"
	"github.com/stainless-sdks/cf-rex-go/option"
)

func TestAccountDlpDatasetVersionSetColumnInfo(t *testing.T) {
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
	_, err := client.Accounts.Dlp.Datasets.Versions.SetColumnInfo(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		int64(0),
		cfrex.AccountDlpDatasetVersionSetColumnInfoParams{
			Body: []cfrex.AccountDlpDatasetVersionSetColumnInfoParamsBodyUnion{cfrex.AccountDlpDatasetVersionSetColumnInfoParamsBodyExistingColumn{
				EntryID:    cfrex.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				HeaderName: cfrex.F("header_name"),
				NumCells:   cfrex.F(int64(0)),
			}},
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

func TestAccountDlpDatasetVersionUploadEntry(t *testing.T) {
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
	_, err := client.Accounts.Dlp.Datasets.Versions.UploadEntry(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		int64(0),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cfrex.AccountDlpDatasetVersionUploadEntryParams{
			Body: io.Reader(bytes.NewBuffer([]byte("some file contents"))),
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
