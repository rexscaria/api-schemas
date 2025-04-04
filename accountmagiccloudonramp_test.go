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

	"github.com/stainless-sdks/cf-rex-go"
	"github.com/stainless-sdks/cf-rex-go/internal/testutil"
	"github.com/stainless-sdks/cf-rex-go/option"
)

func TestAccountMagicCloudOnrampNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.Cloud.Onramps.New(
		context.TODO(),
		"account_id",
		cfrex.AccountMagicCloudOnrampNewParams{
			CloudType:                 cfrex.F(cfrex.McnOnrampCloudTypeAws),
			InstallRoutesInCloud:      cfrex.F(true),
			InstallRoutesInMagicWan:   cfrex.F(true),
			Name:                      cfrex.F("name"),
			Type:                      cfrex.F(cfrex.McnOnrampTypeOnrampTypeSingle),
			AdoptedHubID:              cfrex.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AttachedHubs:              cfrex.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			AttachedVpcs:              cfrex.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			Description:               cfrex.F("description"),
			HubProviderID:             cfrex.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ManageHubToHubAttachments: cfrex.F(true),
			ManageVpcToHubAttachments: cfrex.F(true),
			Region:                    cfrex.F("region"),
			Vpc:                       cfrex.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Forwarded:                 cfrex.F("forwarded"),
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

func TestAccountMagicCloudOnrampGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.Cloud.Onramps.Get(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cfrex.AccountMagicCloudOnrampGetParams{
			PlannedResources:   cfrex.F(true),
			PostApplyResources: cfrex.F(true),
			Status:             cfrex.F(true),
			Vpcs:               cfrex.F(true),
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

func TestAccountMagicCloudOnrampUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.Cloud.Onramps.Update(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cfrex.AccountMagicCloudOnrampUpdateParams{
			McnUpdateOnrampRequest: cfrex.McnUpdateOnrampRequestParam{
				AttachedHubs:              cfrex.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
				AttachedVpcs:              cfrex.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
				Description:               cfrex.F("description"),
				InstallRoutesInCloud:      cfrex.F(true),
				InstallRoutesInMagicWan:   cfrex.F(true),
				ManageHubToHubAttachments: cfrex.F(true),
				ManageVpcToHubAttachments: cfrex.F(true),
				Name:                      cfrex.F("name"),
				Vpc:                       cfrex.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestAccountMagicCloudOnrampListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.Cloud.Onramps.List(
		context.TODO(),
		"account_id",
		cfrex.AccountMagicCloudOnrampListParams{
			Desc:    cfrex.F(true),
			OrderBy: cfrex.F("order_by"),
			Status:  cfrex.F(true),
			Vpcs:    cfrex.F(true),
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

func TestAccountMagicCloudOnrampDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.Cloud.Onramps.Delete(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cfrex.AccountMagicCloudOnrampDeleteParams{
			Destroy: cfrex.F(true),
			Force:   cfrex.F(true),
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

func TestAccountMagicCloudOnrampApply(t *testing.T) {
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
	_, err := client.Accounts.Magic.Cloud.Onramps.Apply(
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

func TestAccountMagicCloudOnrampExport(t *testing.T) {
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
	resp, err := client.Accounts.Magic.Cloud.Onramps.Export(
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

func TestAccountMagicCloudOnrampPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.Cloud.Onramps.Patch(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cfrex.AccountMagicCloudOnrampPatchParams{
			McnUpdateOnrampRequest: cfrex.McnUpdateOnrampRequestParam{
				AttachedHubs:              cfrex.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
				AttachedVpcs:              cfrex.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
				Description:               cfrex.F("description"),
				InstallRoutesInCloud:      cfrex.F(true),
				InstallRoutesInMagicWan:   cfrex.F(true),
				ManageHubToHubAttachments: cfrex.F(true),
				ManageVpcToHubAttachments: cfrex.F(true),
				Name:                      cfrex.F("name"),
				Vpc:                       cfrex.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestAccountMagicCloudOnrampPlan(t *testing.T) {
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
	_, err := client.Accounts.Magic.Cloud.Onramps.Plan(
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
