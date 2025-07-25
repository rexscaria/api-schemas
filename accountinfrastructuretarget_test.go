// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/rexscaria/api-schemas"
	"github.com/rexscaria/api-schemas/internal/testutil"
	"github.com/rexscaria/api-schemas/option"
)

func TestAccountInfrastructureTargetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Infrastructure.Targets.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountInfrastructureTargetNewParams{
			Hostname: cfrex.F("infra-access-target"),
			IP: cfrex.F(cfrex.IPInfoTargetParam{
				Ipv4: cfrex.F(cfrex.IPInfoTargetIpv4Param{
					IPAddr:           cfrex.F("187.26.29.249"),
					VirtualNetworkID: cfrex.F("c77b744e-acc8-428f-9257-6878c046ed55"),
				}),
				Ipv6: cfrex.F(cfrex.IPInfoTargetIpv6Param{
					IPAddr:           cfrex.F("64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0"),
					VirtualNetworkID: cfrex.F("c77b744e-acc8-428f-9257-6878c046ed55"),
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

func TestAccountInfrastructureTargetGet(t *testing.T) {
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
	_, err := client.Accounts.Infrastructure.Targets.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountInfrastructureTargetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Infrastructure.Targets.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cfrex.AccountInfrastructureTargetUpdateParams{
			Hostname: cfrex.F("infra-access-target"),
			IP: cfrex.F(cfrex.IPInfoTargetParam{
				Ipv4: cfrex.F(cfrex.IPInfoTargetIpv4Param{
					IPAddr:           cfrex.F("187.26.29.249"),
					VirtualNetworkID: cfrex.F("c77b744e-acc8-428f-9257-6878c046ed55"),
				}),
				Ipv6: cfrex.F(cfrex.IPInfoTargetIpv6Param{
					IPAddr:           cfrex.F("64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0"),
					VirtualNetworkID: cfrex.F("c77b744e-acc8-428f-9257-6878c046ed55"),
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

func TestAccountInfrastructureTargetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Infrastructure.Targets.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountInfrastructureTargetListParams{
			CreatedAfter:     cfrex.F(time.Now()),
			CreatedBefore:    cfrex.F(time.Now()),
			Direction:        cfrex.F(cfrex.AccountInfrastructureTargetListParamsDirectionAsc),
			Hostname:         cfrex.F("hostname"),
			HostnameContains: cfrex.F("hostname_contains"),
			IPLike:           cfrex.F("ip_like"),
			IPV4:             cfrex.F("ip_v4"),
			IPV6:             cfrex.F("ip_v6"),
			IPs:              cfrex.F([]string{"string"}),
			Ipv4End:          cfrex.F("ipv4_end"),
			Ipv4Start:        cfrex.F("ipv4_start"),
			Ipv6End:          cfrex.F("ipv6_end"),
			Ipv6Start:        cfrex.F("ipv6_start"),
			ModifiedAfter:    cfrex.F(time.Now()),
			ModifiedBefore:   cfrex.F(time.Now()),
			Order:            cfrex.F(cfrex.AccountInfrastructureTargetListParamsOrderHostname),
			Page:             cfrex.F(int64(1)),
			PerPage:          cfrex.F(int64(1)),
			TargetIDs:        cfrex.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			VirtualNetworkID: cfrex.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestAccountInfrastructureTargetDelete(t *testing.T) {
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
	err := client.Accounts.Infrastructure.Targets.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
