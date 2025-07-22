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

func TestAccountMnmConfigNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Mnm.Config.New(
		context.TODO(),
		"6f91088a406011ed95aed352566e8d4c",
		cfrex.AccountMnmConfigNewParams{
			DefaultSampling: cfrex.F(1.000000),
			Name:            cfrex.F("cloudflare user's account"),
			RouterIPs:       cfrex.F([]string{"203.0.113.1"}),
			WarpDevices: cfrex.F([]cfrex.WarpDeviceParam{{
				ID:       cfrex.F("5360368d-b351-4791-abe1-93550dabd351"),
				Name:     cfrex.F("My warp device"),
				RouterIP: cfrex.F("203.0.113.1"),
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

func TestAccountMnmConfigUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Mnm.Config.Update(
		context.TODO(),
		"6f91088a406011ed95aed352566e8d4c",
		cfrex.AccountMnmConfigUpdateParams{
			DefaultSampling: cfrex.F(1.000000),
			Name:            cfrex.F("cloudflare user's account"),
			RouterIPs:       cfrex.F([]string{"203.0.113.1"}),
			WarpDevices: cfrex.F([]cfrex.WarpDeviceParam{{
				ID:       cfrex.F("5360368d-b351-4791-abe1-93550dabd351"),
				Name:     cfrex.F("My warp device"),
				RouterIP: cfrex.F("203.0.113.1"),
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

func TestAccountMnmConfigList(t *testing.T) {
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
	_, err := client.Accounts.Mnm.Config.List(context.TODO(), "6f91088a406011ed95aed352566e8d4c")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountMnmConfigDelete(t *testing.T) {
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
	_, err := client.Accounts.Mnm.Config.Delete(
		context.TODO(),
		"6f91088a406011ed95aed352566e8d4c",
		cfrex.AccountMnmConfigDeleteParams{
			Body: map[string]interface{}{},
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

func TestAccountMnmConfigListFull(t *testing.T) {
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
	_, err := client.Accounts.Mnm.Config.ListFull(context.TODO(), "6f91088a406011ed95aed352566e8d4c")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountMnmConfigUpdateFieldsWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Mnm.Config.UpdateFields(
		context.TODO(),
		"6f91088a406011ed95aed352566e8d4c",
		cfrex.AccountMnmConfigUpdateFieldsParams{
			DefaultSampling: cfrex.F(1.000000),
			Name:            cfrex.F("cloudflare user's account"),
			RouterIPs:       cfrex.F([]string{"203.0.113.1"}),
			WarpDevices: cfrex.F([]cfrex.WarpDeviceParam{{
				ID:       cfrex.F("5360368d-b351-4791-abe1-93550dabd351"),
				Name:     cfrex.F("My warp device"),
				RouterIP: cfrex.F("203.0.113.1"),
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
