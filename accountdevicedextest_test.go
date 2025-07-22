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

func TestAccountDeviceDexTestNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Devices.DexTests.New(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.AccountDeviceDexTestNewParams{
			DeviceDexTestHTTP: cfrex.DeviceDexTestHTTPParam{
				Data: cfrex.F(cfrex.DeviceDexTestHTTPDataParam{
					Host:   cfrex.F("https://dash.cloudflare.com"),
					Kind:   cfrex.F("http"),
					Method: cfrex.F("GET"),
				}),
				Enabled:     cfrex.F(true),
				Interval:    cfrex.F("30m"),
				Name:        cfrex.F("HTTP dash health check"),
				Description: cfrex.F("Checks the dash endpoint every 30 minutes"),
				TargetPolicies: cfrex.F([]cfrex.DeviceDexTestHTTPTargetPolicyParam{{
					ID:      cfrex.F("id"),
					Default: cfrex.F(true),
					Name:    cfrex.F("name"),
				}}),
				Targeted: cfrex.F(true),
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

func TestAccountDeviceDexTestGet(t *testing.T) {
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
	_, err := client.Accounts.Devices.DexTests.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"372e67954025e0ba6aaa6d586b9e0b59",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDeviceDexTestUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Devices.DexTests.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.AccountDeviceDexTestUpdateParams{
			DeviceDexTestHTTP: cfrex.DeviceDexTestHTTPParam{
				Data: cfrex.F(cfrex.DeviceDexTestHTTPDataParam{
					Host:   cfrex.F("https://dash.cloudflare.com"),
					Kind:   cfrex.F("http"),
					Method: cfrex.F("GET"),
				}),
				Enabled:     cfrex.F(true),
				Interval:    cfrex.F("30m"),
				Name:        cfrex.F("HTTP dash health check"),
				Description: cfrex.F("Checks the dash endpoint every 30 minutes"),
				TargetPolicies: cfrex.F([]cfrex.DeviceDexTestHTTPTargetPolicyParam{{
					ID:      cfrex.F("id"),
					Default: cfrex.F(true),
					Name:    cfrex.F("name"),
				}}),
				Targeted: cfrex.F(true),
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

func TestAccountDeviceDexTestList(t *testing.T) {
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
	_, err := client.Accounts.Devices.DexTests.List(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDeviceDexTestDelete(t *testing.T) {
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
	_, err := client.Accounts.Devices.DexTests.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
