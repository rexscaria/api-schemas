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

func TestAccountLoadBalancerMonitorNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Monitors.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountLoadBalancerMonitorNewParams{
			EditableMonitor: cfrex.EditableMonitorParam{
				AllowInsecure:   cfrex.F(true),
				ConsecutiveDown: cfrex.F(int64(0)),
				ConsecutiveUp:   cfrex.F(int64(0)),
				Description:     cfrex.F("Login page monitor"),
				ExpectedBody:    cfrex.F("alive"),
				ExpectedCodes:   cfrex.F("2xx"),
				FollowRedirects: cfrex.F(true),
				Header: cfrex.F(map[string][]string{
					"Host":     {"example.com"},
					"X-App-ID": {"abc123"},
				}),
				Interval:  cfrex.F(int64(0)),
				Method:    cfrex.F("GET"),
				Path:      cfrex.F("/health"),
				Port:      cfrex.F(int64(0)),
				ProbeZone: cfrex.F("example.com"),
				Retries:   cfrex.F(int64(0)),
				Timeout:   cfrex.F(int64(0)),
				Type:      cfrex.F(cfrex.EditableMonitorTypeHTTP),
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

func TestAccountLoadBalancerMonitorGet(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Monitors.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountLoadBalancerMonitorUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Monitors.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		cfrex.AccountLoadBalancerMonitorUpdateParams{
			EditableMonitor: cfrex.EditableMonitorParam{
				AllowInsecure:   cfrex.F(true),
				ConsecutiveDown: cfrex.F(int64(0)),
				ConsecutiveUp:   cfrex.F(int64(0)),
				Description:     cfrex.F("Login page monitor"),
				ExpectedBody:    cfrex.F("alive"),
				ExpectedCodes:   cfrex.F("2xx"),
				FollowRedirects: cfrex.F(true),
				Header: cfrex.F(map[string][]string{
					"Host":     {"example.com"},
					"X-App-ID": {"abc123"},
				}),
				Interval:  cfrex.F(int64(0)),
				Method:    cfrex.F("GET"),
				Path:      cfrex.F("/health"),
				Port:      cfrex.F(int64(0)),
				ProbeZone: cfrex.F("example.com"),
				Retries:   cfrex.F(int64(0)),
				Timeout:   cfrex.F(int64(0)),
				Type:      cfrex.F(cfrex.EditableMonitorTypeHTTP),
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

func TestAccountLoadBalancerMonitorList(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Monitors.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountLoadBalancerMonitorDelete(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Monitors.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		cfrex.AccountLoadBalancerMonitorDeleteParams{
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

func TestAccountLoadBalancerMonitorListReferences(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Monitors.ListReferences(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountLoadBalancerMonitorPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Monitors.Patch(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		cfrex.AccountLoadBalancerMonitorPatchParams{
			EditableMonitor: cfrex.EditableMonitorParam{
				AllowInsecure:   cfrex.F(true),
				ConsecutiveDown: cfrex.F(int64(0)),
				ConsecutiveUp:   cfrex.F(int64(0)),
				Description:     cfrex.F("Login page monitor"),
				ExpectedBody:    cfrex.F("alive"),
				ExpectedCodes:   cfrex.F("2xx"),
				FollowRedirects: cfrex.F(true),
				Header: cfrex.F(map[string][]string{
					"Host":     {"example.com"},
					"X-App-ID": {"abc123"},
				}),
				Interval:  cfrex.F(int64(0)),
				Method:    cfrex.F("GET"),
				Path:      cfrex.F("/health"),
				Port:      cfrex.F(int64(0)),
				ProbeZone: cfrex.F("example.com"),
				Retries:   cfrex.F(int64(0)),
				Timeout:   cfrex.F(int64(0)),
				Type:      cfrex.F(cfrex.EditableMonitorTypeHTTP),
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

func TestAccountLoadBalancerMonitorPreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Monitors.Preview(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		cfrex.AccountLoadBalancerMonitorPreviewParams{
			EditableMonitor: cfrex.EditableMonitorParam{
				AllowInsecure:   cfrex.F(true),
				ConsecutiveDown: cfrex.F(int64(0)),
				ConsecutiveUp:   cfrex.F(int64(0)),
				Description:     cfrex.F("Login page monitor"),
				ExpectedBody:    cfrex.F("alive"),
				ExpectedCodes:   cfrex.F("2xx"),
				FollowRedirects: cfrex.F(true),
				Header: cfrex.F(map[string][]string{
					"Host":     {"example.com"},
					"X-App-ID": {"abc123"},
				}),
				Interval:  cfrex.F(int64(0)),
				Method:    cfrex.F("GET"),
				Path:      cfrex.F("/health"),
				Port:      cfrex.F(int64(0)),
				ProbeZone: cfrex.F("example.com"),
				Retries:   cfrex.F(int64(0)),
				Timeout:   cfrex.F(int64(0)),
				Type:      cfrex.F(cfrex.EditableMonitorTypeHTTP),
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
