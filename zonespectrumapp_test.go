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
	"github.com/rexscaria/api-schemas/shared"
)

func TestZoneSpectrumAppNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Spectrum.Apps.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneSpectrumAppNewParams{
			UpdateAppConfig: cfrex.AppConfigParam{
				DNS: cfrex.F(cfrex.DNSParam{
					Name: cfrex.F("ssh.example.com"),
					Type: cfrex.F(cfrex.DNSTypeCname),
				}),
				IPFirewall:       cfrex.F(true),
				Protocol:         cfrex.F("tcp/22"),
				ProxyProtocol:    cfrex.F(cfrex.AppConfigProxyProtocolOff),
				Tls:              cfrex.F(cfrex.AppConfigTlsFull),
				TrafficType:      cfrex.F(cfrex.AppConfigTrafficTypeDirect),
				ArgoSmartRouting: cfrex.F(true),
				EdgeIPs: cfrex.F[cfrex.AppConfigEdgeIPsUnionParam](cfrex.AppConfigEdgeIPsDynamicParam{
					Connectivity: cfrex.F(cfrex.AppConfigEdgeIPsDynamicConnectivityAll),
					Type:         cfrex.F(cfrex.AppConfigEdgeIPsDynamicTypeDynamic),
				}),
				OriginDirect: cfrex.F([]string{"tcp://127.0.0.1:8080"}),
				OriginDNS: cfrex.F(cfrex.AppConfigOriginDNSParam{
					Name: cfrex.F("origin.example.com"),
					Ttl:  cfrex.F(int64(600)),
					Type: cfrex.F(cfrex.AppConfigOriginDNSTypeEmpty),
				}),
				OriginPort: cfrex.F[cfrex.AppConfigOriginPortUnionParam](shared.UnionInt(int64(22))),
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

func TestZoneSpectrumAppGet(t *testing.T) {
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
	_, err := client.Zones.Spectrum.Apps.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneSpectrumAppUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Spectrum.Apps.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneSpectrumAppUpdateParams{
			UpdateAppConfig: cfrex.AppConfigParam{
				DNS: cfrex.F(cfrex.DNSParam{
					Name: cfrex.F("ssh.example.com"),
					Type: cfrex.F(cfrex.DNSTypeCname),
				}),
				IPFirewall:       cfrex.F(true),
				Protocol:         cfrex.F("tcp/22"),
				ProxyProtocol:    cfrex.F(cfrex.AppConfigProxyProtocolOff),
				Tls:              cfrex.F(cfrex.AppConfigTlsFull),
				TrafficType:      cfrex.F(cfrex.AppConfigTrafficTypeDirect),
				ArgoSmartRouting: cfrex.F(true),
				EdgeIPs: cfrex.F[cfrex.AppConfigEdgeIPsUnionParam](cfrex.AppConfigEdgeIPsDynamicParam{
					Connectivity: cfrex.F(cfrex.AppConfigEdgeIPsDynamicConnectivityAll),
					Type:         cfrex.F(cfrex.AppConfigEdgeIPsDynamicTypeDynamic),
				}),
				OriginDirect: cfrex.F([]string{"tcp://127.0.0.1:8080"}),
				OriginDNS: cfrex.F(cfrex.AppConfigOriginDNSParam{
					Name: cfrex.F("origin.example.com"),
					Ttl:  cfrex.F(int64(600)),
					Type: cfrex.F(cfrex.AppConfigOriginDNSTypeEmpty),
				}),
				OriginPort: cfrex.F[cfrex.AppConfigOriginPortUnionParam](shared.UnionInt(int64(22))),
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

func TestZoneSpectrumAppListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Spectrum.Apps.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneSpectrumAppListParams{
			Direction: cfrex.F(cfrex.ZoneSpectrumAppListParamsDirectionDesc),
			Order:     cfrex.F(cfrex.ZoneSpectrumAppListParamsOrderProtocol),
			Page:      cfrex.F(1.000000),
			PerPage:   cfrex.F(1.000000),
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

func TestZoneSpectrumAppDelete(t *testing.T) {
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
	_, err := client.Zones.Spectrum.Apps.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
