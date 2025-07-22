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

func TestAccountCfdTunnelConfigurationGet(t *testing.T) {
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
	_, err := client.Accounts.CfdTunnel.Configurations.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountCfdTunnelConfigurationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.CfdTunnel.Configurations.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		cfrex.AccountCfdTunnelConfigurationUpdateParams{
			Config: cfrex.F(cfrex.TunnelConfigParam{
				Ingress: cfrex.F([]cfrex.TunnelConfigIngressParam{{
					Hostname: cfrex.F("tunnel.example.com"),
					Service:  cfrex.F("https://localhost:8001"),
					OriginRequest: cfrex.F(cfrex.OriginRequestParam{
						Access: cfrex.F(cfrex.OriginRequestAccessParam{
							AudTag:   cfrex.F([]string{"string"}),
							TeamName: cfrex.F("teamName"),
							Required: cfrex.F(true),
						}),
						CaPool:                 cfrex.F("caPool"),
						ConnectTimeout:         cfrex.F(int64(0)),
						DisableChunkedEncoding: cfrex.F(true),
						Http2Origin:            cfrex.F(true),
						HTTPHostHeader:         cfrex.F("httpHostHeader"),
						KeepAliveConnections:   cfrex.F(int64(0)),
						KeepAliveTimeout:       cfrex.F(int64(0)),
						NoHappyEyeballs:        cfrex.F(true),
						NoTlsVerify:            cfrex.F(true),
						OriginServerName:       cfrex.F("originServerName"),
						ProxyType:              cfrex.F("proxyType"),
						TcpKeepAlive:           cfrex.F(int64(0)),
						TlsTimeout:             cfrex.F(int64(0)),
					}),
					Path: cfrex.F("subpath"),
				}}),
				OriginRequest: cfrex.F(cfrex.OriginRequestParam{
					Access: cfrex.F(cfrex.OriginRequestAccessParam{
						AudTag:   cfrex.F([]string{"string"}),
						TeamName: cfrex.F("teamName"),
						Required: cfrex.F(true),
					}),
					CaPool:                 cfrex.F("caPool"),
					ConnectTimeout:         cfrex.F(int64(0)),
					DisableChunkedEncoding: cfrex.F(true),
					Http2Origin:            cfrex.F(true),
					HTTPHostHeader:         cfrex.F("httpHostHeader"),
					KeepAliveConnections:   cfrex.F(int64(0)),
					KeepAliveTimeout:       cfrex.F(int64(0)),
					NoHappyEyeballs:        cfrex.F(true),
					NoTlsVerify:            cfrex.F(true),
					OriginServerName:       cfrex.F("originServerName"),
					ProxyType:              cfrex.F("proxyType"),
					TcpKeepAlive:           cfrex.F(int64(0)),
					TlsTimeout:             cfrex.F(int64(0)),
				}),
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
