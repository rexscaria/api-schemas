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

func TestAccountGatewayLocationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Locations.New(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.AccountGatewayLocationNewParams{
			Name:                cfrex.F("Austin Office Location"),
			ClientDefault:       cfrex.F(false),
			DNSDestinationIPsID: cfrex.F("0e4a32c6-6fb8-4858-9296-98f51631e8e6"),
			EcsSupport:          cfrex.F(false),
			Endpoints: cfrex.F(cfrex.EndpointsParam{
				Doh: cfrex.F(cfrex.EndpointsDohParam{
					Enabled: cfrex.F(true),
					Networks: cfrex.F([]cfrex.IPNetworkParam{{
						Network: cfrex.F("2001:85a3::/64"),
					}}),
					RequireToken: cfrex.F(true),
				}),
				Dot: cfrex.F(cfrex.EndpointsDotParam{
					Enabled: cfrex.F(true),
					Networks: cfrex.F([]cfrex.IPNetworkParam{{
						Network: cfrex.F("2001:85a3::/64"),
					}}),
				}),
				Ipv4: cfrex.F(cfrex.EndpointsIpv4Param{
					Enabled: cfrex.F(true),
				}),
				Ipv6: cfrex.F(cfrex.EndpointsIpv6Param{
					Enabled: cfrex.F(true),
					Networks: cfrex.F([]cfrex.EndpointsIpv6NetworkParam{{
						Network: cfrex.F("2001:85a3::/64"),
					}}),
				}),
			}),
			Networks: cfrex.F([]cfrex.Ipv4NetworkParam{{
				Network: cfrex.F("192.0.2.1/32"),
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

func TestAccountGatewayLocationGet(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Locations.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"ed35569b41ce4d1facfe683550f54086",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountGatewayLocationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Locations.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"ed35569b41ce4d1facfe683550f54086",
		cfrex.AccountGatewayLocationUpdateParams{
			Name:                cfrex.F("Austin Office Location"),
			ClientDefault:       cfrex.F(false),
			DNSDestinationIPsID: cfrex.F("0e4a32c6-6fb8-4858-9296-98f51631e8e6"),
			EcsSupport:          cfrex.F(false),
			Endpoints: cfrex.F(cfrex.EndpointsParam{
				Doh: cfrex.F(cfrex.EndpointsDohParam{
					Enabled: cfrex.F(true),
					Networks: cfrex.F([]cfrex.IPNetworkParam{{
						Network: cfrex.F("2001:85a3::/64"),
					}}),
					RequireToken: cfrex.F(true),
				}),
				Dot: cfrex.F(cfrex.EndpointsDotParam{
					Enabled: cfrex.F(true),
					Networks: cfrex.F([]cfrex.IPNetworkParam{{
						Network: cfrex.F("2001:85a3::/64"),
					}}),
				}),
				Ipv4: cfrex.F(cfrex.EndpointsIpv4Param{
					Enabled: cfrex.F(true),
				}),
				Ipv6: cfrex.F(cfrex.EndpointsIpv6Param{
					Enabled: cfrex.F(true),
					Networks: cfrex.F([]cfrex.EndpointsIpv6NetworkParam{{
						Network: cfrex.F("2001:85a3::/64"),
					}}),
				}),
			}),
			Networks: cfrex.F([]cfrex.Ipv4NetworkParam{{
				Network: cfrex.F("192.0.2.1/32"),
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

func TestAccountGatewayLocationList(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Locations.List(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountGatewayLocationDelete(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Locations.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"ed35569b41ce4d1facfe683550f54086",
		cfrex.AccountGatewayLocationDeleteParams{
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
