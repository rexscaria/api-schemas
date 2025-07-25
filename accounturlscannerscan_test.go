// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/rexscaria/api-schemas"
	"github.com/rexscaria/api-schemas/internal/testutil"
	"github.com/rexscaria/api-schemas/option"
)

func TestAccountUrlscannerScanNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Urlscanner.Scan.New(
		context.TODO(),
		"account_id",
		cfrex.AccountUrlscannerScanNewParams{
			URL:     cfrex.F("https://www.example.com"),
			Country: cfrex.F(cfrex.AccountUrlscannerScanNewParamsCountryAf),
			CustomHeaders: cfrex.F(map[string]string{
				"foo": "string",
			}),
			ScreenshotsResolutions: cfrex.F([]cfrex.AccountUrlscannerScanNewParamsScreenshotsResolution{cfrex.AccountUrlscannerScanNewParamsScreenshotsResolutionDesktop}),
			Visibility:             cfrex.F(cfrex.AccountUrlscannerScanNewParamsVisibilityPublic),
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

func TestAccountUrlscannerScanGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Urlscanner.Scan.Get(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cfrex.AccountUrlscannerScanGetParams{
			Full: cfrex.F(true),
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

func TestAccountUrlscannerScanListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Urlscanner.Scan.List(
		context.TODO(),
		"account_id",
		cfrex.AccountUrlscannerScanListParams{
			AccountScans: cfrex.F(true),
			DateEnd:      cfrex.F(time.Now()),
			DateStart:    cfrex.F(time.Now()),
			Hash:         cfrex.F("hash"),
			Hostname:     cfrex.F("example.com"),
			IP:           cfrex.F("1.1.1.1"),
			IsMalicious:  cfrex.F(true),
			Limit:        cfrex.F(int64(100)),
			NextCursor:   cfrex.F("next_cursor"),
			PageAsn:      cfrex.F("page_asn"),
			PageHostname: cfrex.F("page_hostname"),
			PageIP:       cfrex.F("page_ip"),
			PagePath:     cfrex.F("page_path"),
			PageURL:      cfrex.F("page_url"),
			Path:         cfrex.F("/samples/subresource-integrity/"),
			ScanID:       cfrex.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			URL:          cfrex.F("https://example.com/?hello"),
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

func TestAccountUrlscannerScanGetHar(t *testing.T) {
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
	_, err := client.Accounts.Urlscanner.Scan.GetHar(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountUrlscannerScanGetScreenshotWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := cfrex.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("My API Email"),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.Accounts.Urlscanner.Scan.GetScreenshot(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cfrex.AccountUrlscannerScanGetScreenshotParams{
			Resolution: cfrex.F(cfrex.AccountUrlscannerScanGetScreenshotParamsResolutionDesktop),
		},
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
