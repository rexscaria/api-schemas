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

	"github.com/stainless-sdks/cf-rex-go"
	"github.com/stainless-sdks/cf-rex-go/internal/testutil"
	"github.com/stainless-sdks/cf-rex-go/option"
)

func TestAccountDexCommandNew(t *testing.T) {
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
	_, err := client.Accounts.Dex.Commands.New(
		context.TODO(),
		"01a7362d577a6c3019a474fd6f485823",
		cfrex.AccountDexCommandNewParams{
			Commands: cfrex.F([]cfrex.AccountDexCommandNewParamsCommand{{
				CommandType: cfrex.F(cfrex.AccountDexCommandNewParamsCommandsCommandTypePcap),
				DeviceID:    cfrex.F("device_id"),
				UserEmail:   cfrex.F("user_email"),
				CommandArgs: cfrex.F(cfrex.AccountDexCommandNewParamsCommandsCommandArgs{
					Interfaces:      cfrex.F([]cfrex.AccountDexCommandNewParamsCommandsCommandArgsInterface{cfrex.AccountDexCommandNewParamsCommandsCommandArgsInterfaceDefault}),
					MaxFileSizeMB:   cfrex.F(1.000000),
					PacketSizeBytes: cfrex.F(1.000000),
					TestAllRoutes:   cfrex.F(true),
					TimeLimitMin:    cfrex.F(1.000000),
				}),
			}}),
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

func TestAccountDexCommandListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Dex.Commands.List(
		context.TODO(),
		"01a7362d577a6c3019a474fd6f485823",
		cfrex.AccountDexCommandListParams{
			Page:        cfrex.F(1.000000),
			PerPage:     cfrex.F(50.000000),
			CommandType: cfrex.F("command_type"),
			DeviceID:    cfrex.F("device_id"),
			From:        cfrex.F(time.Now()),
			Status:      cfrex.F(cfrex.AccountDexCommandListParamsStatusPendingExec),
			To:          cfrex.F(time.Now()),
			UserEmail:   cfrex.F("user_email"),
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

func TestAccountDexCommandDownloadOutput(t *testing.T) {
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
	resp, err := client.Accounts.Dex.Commands.DownloadOutput(
		context.TODO(),
		"01a7362d577a6c3019a474fd6f485823",
		"5758fefe-ae7e-4538-a39b-1fef6abcb909",
		"filename",
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

func TestAccountDexCommandGetQuota(t *testing.T) {
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
	_, err := client.Accounts.Dex.Commands.GetQuota(context.TODO(), "01a7362d577a6c3019a474fd6f485823")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDexCommandListEligibleDevicesWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Dex.Commands.ListEligibleDevices(
		context.TODO(),
		"01a7362d577a6c3019a474fd6f485823",
		cfrex.AccountDexCommandListEligibleDevicesParams{
			Page:    cfrex.F(1.000000),
			PerPage: cfrex.F(1.000000),
			Search:  cfrex.F("search"),
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
