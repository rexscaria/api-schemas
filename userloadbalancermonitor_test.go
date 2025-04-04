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

func TestUserLoadBalancerMonitorNewWithOptionalParams(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Monitors.New(context.TODO(), cfrex.UserLoadBalancerMonitorNewParams{
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
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerMonitorGet(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Monitors.Get(context.TODO(), "f1aba936b94213e5b8dca0c0dbf1f9cc")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerMonitorUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Monitors.Update(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		cfrex.UserLoadBalancerMonitorUpdateParams{
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

func TestUserLoadBalancerMonitorList(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Monitors.List(context.TODO())
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerMonitorDelete(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Monitors.Delete(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		cfrex.UserLoadBalancerMonitorDeleteParams{
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

func TestUserLoadBalancerMonitorListReferences(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Monitors.ListReferences(context.TODO(), "f1aba936b94213e5b8dca0c0dbf1f9cc")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerMonitorPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Monitors.Patch(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		cfrex.UserLoadBalancerMonitorPatchParams{
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

func TestUserLoadBalancerMonitorPreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.User.LoadBalancers.Monitors.Preview(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		cfrex.UserLoadBalancerMonitorPreviewParams{
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
