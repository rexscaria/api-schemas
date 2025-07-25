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

func TestAccountLoadBalancerPoolNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Pools.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountLoadBalancerPoolNewParams{
			Name: cfrex.F("primary-dc-1"),
			Origins: cfrex.F([]cfrex.OriginParam{{
				Address: cfrex.F("0.0.0.0"),
				Enabled: cfrex.F(true),
				Header: cfrex.F(cfrex.OriginHeaderParam{
					Host: cfrex.F([]string{"example.com"}),
				}),
				Name:             cfrex.F("app-server-1"),
				Port:             cfrex.F(int64(0)),
				VirtualNetworkID: cfrex.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cfrex.F(0.600000),
			}}),
			Description: cfrex.F("Primary data center - Provider XYZ"),
			Enabled:     cfrex.F(false),
			Latitude:    cfrex.F(0.000000),
			LoadShedding: cfrex.F(cfrex.LoadSheddingParam{
				DefaultPercent: cfrex.F(0.000000),
				DefaultPolicy:  cfrex.F(cfrex.LoadSheddingDefaultPolicyRandom),
				SessionPercent: cfrex.F(0.000000),
				SessionPolicy:  cfrex.F(cfrex.LoadSheddingSessionPolicyHash),
			}),
			Longitude:         cfrex.F(0.000000),
			MinimumOrigins:    cfrex.F(int64(0)),
			Monitor:           cfrex.F("monitor"),
			NotificationEmail: cfrex.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: cfrex.F(cfrex.NotificationFilterParam{
				Origin: cfrex.F(cfrex.FilterOptionsParam{
					Disable: cfrex.F(true),
					Healthy: cfrex.F(true),
				}),
				Pool: cfrex.F(cfrex.FilterOptionsParam{
					Disable: cfrex.F(true),
					Healthy: cfrex.F(false),
				}),
			}),
			OriginSteering: cfrex.F(cfrex.OriginSteeringParam{
				Policy: cfrex.F(cfrex.OriginSteeringPolicyRandom),
			}),
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

func TestAccountLoadBalancerPoolGet(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Pools.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"17b5962d775c646f3f9725cbc7a53df4",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountLoadBalancerPoolUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Pools.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"17b5962d775c646f3f9725cbc7a53df4",
		cfrex.AccountLoadBalancerPoolUpdateParams{
			Name: cfrex.F("primary-dc-1"),
			Origins: cfrex.F([]cfrex.OriginParam{{
				Address: cfrex.F("0.0.0.0"),
				Enabled: cfrex.F(true),
				Header: cfrex.F(cfrex.OriginHeaderParam{
					Host: cfrex.F([]string{"example.com"}),
				}),
				Name:             cfrex.F("app-server-1"),
				Port:             cfrex.F(int64(0)),
				VirtualNetworkID: cfrex.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cfrex.F(0.600000),
			}}),
			CheckRegions: cfrex.F([]cfrex.CheckRegions{cfrex.CheckRegionsWeu, cfrex.CheckRegionsEnam}),
			Description:  cfrex.F("Primary data center - Provider XYZ"),
			Enabled:      cfrex.F(false),
			Latitude:     cfrex.F(0.000000),
			LoadShedding: cfrex.F(cfrex.LoadSheddingParam{
				DefaultPercent: cfrex.F(0.000000),
				DefaultPolicy:  cfrex.F(cfrex.LoadSheddingDefaultPolicyRandom),
				SessionPercent: cfrex.F(0.000000),
				SessionPolicy:  cfrex.F(cfrex.LoadSheddingSessionPolicyHash),
			}),
			Longitude:         cfrex.F(0.000000),
			MinimumOrigins:    cfrex.F(int64(0)),
			Monitor:           cfrex.F("monitor"),
			NotificationEmail: cfrex.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: cfrex.F(cfrex.NotificationFilterParam{
				Origin: cfrex.F(cfrex.FilterOptionsParam{
					Disable: cfrex.F(true),
					Healthy: cfrex.F(true),
				}),
				Pool: cfrex.F(cfrex.FilterOptionsParam{
					Disable: cfrex.F(true),
					Healthy: cfrex.F(false),
				}),
			}),
			OriginSteering: cfrex.F(cfrex.OriginSteeringParam{
				Policy: cfrex.F(cfrex.OriginSteeringPolicyRandom),
			}),
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

func TestAccountLoadBalancerPoolListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Pools.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountLoadBalancerPoolListParams{
			Monitor: cfrex.F("monitor"),
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

func TestAccountLoadBalancerPoolDelete(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Pools.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"17b5962d775c646f3f9725cbc7a53df4",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountLoadBalancerPoolHealth(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Pools.Health(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"17b5962d775c646f3f9725cbc7a53df4",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountLoadBalancerPoolListReferences(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Pools.ListReferences(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"17b5962d775c646f3f9725cbc7a53df4",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountLoadBalancerPoolPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Pools.Patch(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"17b5962d775c646f3f9725cbc7a53df4",
		cfrex.AccountLoadBalancerPoolPatchParams{
			CheckRegions: cfrex.F([]cfrex.CheckRegions{cfrex.CheckRegionsWeu, cfrex.CheckRegionsEnam}),
			Description:  cfrex.F("Primary data center - Provider XYZ"),
			Enabled:      cfrex.F(false),
			Latitude:     cfrex.F(0.000000),
			LoadShedding: cfrex.F(cfrex.LoadSheddingParam{
				DefaultPercent: cfrex.F(0.000000),
				DefaultPolicy:  cfrex.F(cfrex.LoadSheddingDefaultPolicyRandom),
				SessionPercent: cfrex.F(0.000000),
				SessionPolicy:  cfrex.F(cfrex.LoadSheddingSessionPolicyHash),
			}),
			Longitude:         cfrex.F(0.000000),
			MinimumOrigins:    cfrex.F(int64(0)),
			Monitor:           cfrex.F("monitor"),
			Name:              cfrex.F("primary-dc-1"),
			NotificationEmail: cfrex.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: cfrex.F(cfrex.NotificationFilterParam{
				Origin: cfrex.F(cfrex.FilterOptionsParam{
					Disable: cfrex.F(true),
					Healthy: cfrex.F(true),
				}),
				Pool: cfrex.F(cfrex.FilterOptionsParam{
					Disable: cfrex.F(true),
					Healthy: cfrex.F(false),
				}),
			}),
			OriginSteering: cfrex.F(cfrex.OriginSteeringParam{
				Policy: cfrex.F(cfrex.OriginSteeringPolicyRandom),
			}),
			Origins: cfrex.F([]cfrex.OriginParam{{
				Address: cfrex.F("0.0.0.0"),
				Enabled: cfrex.F(true),
				Header: cfrex.F(cfrex.OriginHeaderParam{
					Host: cfrex.F([]string{"example.com"}),
				}),
				Name:             cfrex.F("app-server-1"),
				Port:             cfrex.F(int64(0)),
				VirtualNetworkID: cfrex.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cfrex.F(0.600000),
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

func TestAccountLoadBalancerPoolPreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.LoadBalancers.Pools.Preview(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"17b5962d775c646f3f9725cbc7a53df4",
		cfrex.AccountLoadBalancerPoolPreviewParams{
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
				Type:      cfrex.F(cfrex.EditableMonitorTypeHTTPS),
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
