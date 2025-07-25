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

func TestAccountMagicIpsecTunnelNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.IpsecTunnels.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountMagicIpsecTunnelNewParams{
			MagicIpsecTunnelAddSingleRequest: cfrex.MagicIpsecTunnelAddSingleRequestParam{
				CloudflareEndpoint: cfrex.F("203.0.113.1"),
				InterfaceAddress:   cfrex.F("192.0.2.0/31"),
				Name:               cfrex.F("IPsec_1"),
				CustomerEndpoint:   cfrex.F("203.0.113.1"),
				Description:        cfrex.F("Tunnel for ISP X"),
				HealthCheck: cfrex.F(cfrex.MagicTunnelHealthCheckParam{
					Direction: cfrex.F(cfrex.MagicTunnelHealthCheckDirectionBidirectional),
					Enabled:   cfrex.F(true),
					Rate:      cfrex.F(cfrex.MagicTunnelHealthCheckRateLow),
					Target: cfrex.F[cfrex.MagicTunnelHealthCheckTargetUnionParam](cfrex.MagicTunnelHealthCheckTargetMagicHealthCheckTargetParam{
						Saved: cfrex.F("203.0.113.1"),
					}),
					Type: cfrex.F(cfrex.MagicTunnelHealthCheckTypeRequest),
				}),
				Psk:              cfrex.F("O3bwKSjnaoCxDoUxjcq4Rk8ZKkezQUiy"),
				ReplayProtection: cfrex.F(false),
			},
			XMagicNewHcTarget: cfrex.F(true),
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

func TestAccountMagicIpsecTunnelGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.IpsecTunnels.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountMagicIpsecTunnelGetParams{
			XMagicNewHcTarget: cfrex.F(true),
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

func TestAccountMagicIpsecTunnelUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.IpsecTunnels.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountMagicIpsecTunnelUpdateParams{
			MagicIpsecTunnelAddSingleRequest: cfrex.MagicIpsecTunnelAddSingleRequestParam{
				CloudflareEndpoint: cfrex.F("203.0.113.1"),
				InterfaceAddress:   cfrex.F("192.0.2.0/31"),
				Name:               cfrex.F("IPsec_1"),
				CustomerEndpoint:   cfrex.F("203.0.113.1"),
				Description:        cfrex.F("Tunnel for ISP X"),
				HealthCheck: cfrex.F(cfrex.MagicTunnelHealthCheckParam{
					Direction: cfrex.F(cfrex.MagicTunnelHealthCheckDirectionBidirectional),
					Enabled:   cfrex.F(true),
					Rate:      cfrex.F(cfrex.MagicTunnelHealthCheckRateLow),
					Target: cfrex.F[cfrex.MagicTunnelHealthCheckTargetUnionParam](cfrex.MagicTunnelHealthCheckTargetMagicHealthCheckTargetParam{
						Saved: cfrex.F("203.0.113.1"),
					}),
					Type: cfrex.F(cfrex.MagicTunnelHealthCheckTypeRequest),
				}),
				Psk:              cfrex.F("O3bwKSjnaoCxDoUxjcq4Rk8ZKkezQUiy"),
				ReplayProtection: cfrex.F(false),
			},
			XMagicNewHcTarget: cfrex.F(true),
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

func TestAccountMagicIpsecTunnelListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.IpsecTunnels.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountMagicIpsecTunnelListParams{
			XMagicNewHcTarget: cfrex.F(true),
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

func TestAccountMagicIpsecTunnelDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.IpsecTunnels.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountMagicIpsecTunnelDeleteParams{
			XMagicNewHcTarget: cfrex.F(true),
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

func TestAccountMagicIpsecTunnelGeneratePsk(t *testing.T) {
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
	_, err := client.Accounts.Magic.IpsecTunnels.GeneratePsk(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountMagicIpsecTunnelGeneratePskParams{
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
