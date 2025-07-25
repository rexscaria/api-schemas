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

	"github.com/rexscaria/api-schemas"
	"github.com/rexscaria/api-schemas/internal/testutil"
	"github.com/rexscaria/api-schemas/option"
)

func TestAccountMagicCloudResourceGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.Cloud.Resources.Get(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cfrex.AccountMagicCloudResourceGetParams{
			V2: cfrex.F(true),
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

func TestAccountMagicCloudResourceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.Cloud.Resources.List(
		context.TODO(),
		"account_id",
		cfrex.AccountMagicCloudResourceListParams{
			Cloudflare:    cfrex.F(true),
			Desc:          cfrex.F(true),
			Managed:       cfrex.F(true),
			OrderBy:       cfrex.F("order_by"),
			Page:          cfrex.F(int64(1)),
			PerPage:       cfrex.F(int64(1)),
			ProviderID:    cfrex.F("provider_id"),
			Region:        cfrex.F("region"),
			ResourceGroup: cfrex.F("resource_group"),
			ResourceID:    cfrex.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ResourceType:  cfrex.F([]cfrex.McnResourceType{cfrex.McnResourceTypeAwsCustomerGateway}),
			Search:        cfrex.F([]string{"string"}),
			V2:            cfrex.F(true),
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

func TestAccountMagicCloudResourceExportWithOptionalParams(t *testing.T) {
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
	resp, err := client.Accounts.Magic.Cloud.Resources.Export(
		context.TODO(),
		"account_id",
		cfrex.AccountMagicCloudResourceExportParams{
			Desc:          cfrex.F(true),
			OrderBy:       cfrex.F("order_by"),
			ProviderID:    cfrex.F("provider_id"),
			Region:        cfrex.F("region"),
			ResourceGroup: cfrex.F("resource_group"),
			ResourceID:    cfrex.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ResourceType:  cfrex.F([]cfrex.McnResourceType{cfrex.McnResourceTypeAwsCustomerGateway}),
			Search:        cfrex.F([]string{"string"}),
			V2:            cfrex.F(true),
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

func TestAccountMagicCloudResourcePreviewPolicy(t *testing.T) {
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
	_, err := client.Accounts.Magic.Cloud.Resources.PreviewPolicy(
		context.TODO(),
		"account_id",
		cfrex.AccountMagicCloudResourcePreviewPolicyParams{
			Policy: cfrex.F("policy"),
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
