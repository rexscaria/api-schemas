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

func TestAccountMagicSiteLanNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.Sites.Lans.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountMagicSiteLanNewParams{
			Physport: cfrex.F(int64(1)),
			HaLink:   cfrex.F(true),
			Name:     cfrex.F("name"),
			Nat: cfrex.F(cfrex.MagicNatParam{
				StaticPrefix: cfrex.F("192.0.2.0/24"),
			}),
			RoutedSubnets: cfrex.F([]cfrex.MagicRoutedSubnetParam{{
				NextHop: cfrex.F("192.0.2.1"),
				Prefix:  cfrex.F("192.0.2.0/24"),
				Nat: cfrex.F(cfrex.MagicNatParam{
					StaticPrefix: cfrex.F("192.0.2.0/24"),
				}),
			}}),
			StaticAddressing: cfrex.F(cfrex.MagicLanStaticAddressingParam{
				Address: cfrex.F("192.0.2.0/24"),
				DhcpRelay: cfrex.F(cfrex.MagicLanStaticAddressingDhcpRelayParam{
					ServerAddresses: cfrex.F([]string{"192.0.2.1"}),
				}),
				DhcpServer: cfrex.F(cfrex.MagicLanStaticAddressingDhcpServerParam{
					DhcpPoolEnd:   cfrex.F("192.0.2.1"),
					DhcpPoolStart: cfrex.F("192.0.2.1"),
					DNSServer:     cfrex.F("192.0.2.1"),
					DNSServers:    cfrex.F([]string{"192.0.2.1"}),
					Reservations: cfrex.F(map[string]string{
						"00:11:22:33:44:55": "192.0.2.100",
						"AA:BB:CC:DD:EE:FF": "192.168.1.101",
					}),
				}),
				SecondaryAddress: cfrex.F("192.0.2.0/24"),
				VirtualAddress:   cfrex.F("192.0.2.0/24"),
			}),
			VlanTag: cfrex.F(int64(42)),
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

func TestAccountMagicSiteLanGet(t *testing.T) {
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
	_, err := client.Accounts.Magic.Sites.Lans.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
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

func TestAccountMagicSiteLanUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.Sites.Lans.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountMagicSiteLanUpdateParams{
			MagicLanUpdateRequest: cfrex.MagicLanUpdateRequestParam{
				Name: cfrex.F("name"),
				Nat: cfrex.F(cfrex.MagicNatParam{
					StaticPrefix: cfrex.F("192.0.2.0/24"),
				}),
				Physport: cfrex.F(int64(1)),
				RoutedSubnets: cfrex.F([]cfrex.MagicRoutedSubnetParam{{
					NextHop: cfrex.F("192.0.2.1"),
					Prefix:  cfrex.F("192.0.2.0/24"),
					Nat: cfrex.F(cfrex.MagicNatParam{
						StaticPrefix: cfrex.F("192.0.2.0/24"),
					}),
				}}),
				StaticAddressing: cfrex.F(cfrex.MagicLanStaticAddressingParam{
					Address: cfrex.F("192.0.2.0/24"),
					DhcpRelay: cfrex.F(cfrex.MagicLanStaticAddressingDhcpRelayParam{
						ServerAddresses: cfrex.F([]string{"192.0.2.1"}),
					}),
					DhcpServer: cfrex.F(cfrex.MagicLanStaticAddressingDhcpServerParam{
						DhcpPoolEnd:   cfrex.F("192.0.2.1"),
						DhcpPoolStart: cfrex.F("192.0.2.1"),
						DNSServer:     cfrex.F("192.0.2.1"),
						DNSServers:    cfrex.F([]string{"192.0.2.1"}),
						Reservations: cfrex.F(map[string]string{
							"00:11:22:33:44:55": "192.0.2.100",
							"AA:BB:CC:DD:EE:FF": "192.168.1.101",
						}),
					}),
					SecondaryAddress: cfrex.F("192.0.2.0/24"),
					VirtualAddress:   cfrex.F("192.0.2.0/24"),
				}),
				VlanTag: cfrex.F(int64(42)),
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

func TestAccountMagicSiteLanList(t *testing.T) {
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
	_, err := client.Accounts.Magic.Sites.Lans.List(
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

func TestAccountMagicSiteLanDelete(t *testing.T) {
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
	_, err := client.Accounts.Magic.Sites.Lans.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
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

func TestAccountMagicSiteLanPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.Sites.Lans.Patch(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountMagicSiteLanPatchParams{
			MagicLanUpdateRequest: cfrex.MagicLanUpdateRequestParam{
				Name: cfrex.F("name"),
				Nat: cfrex.F(cfrex.MagicNatParam{
					StaticPrefix: cfrex.F("192.0.2.0/24"),
				}),
				Physport: cfrex.F(int64(1)),
				RoutedSubnets: cfrex.F([]cfrex.MagicRoutedSubnetParam{{
					NextHop: cfrex.F("192.0.2.1"),
					Prefix:  cfrex.F("192.0.2.0/24"),
					Nat: cfrex.F(cfrex.MagicNatParam{
						StaticPrefix: cfrex.F("192.0.2.0/24"),
					}),
				}}),
				StaticAddressing: cfrex.F(cfrex.MagicLanStaticAddressingParam{
					Address: cfrex.F("192.0.2.0/24"),
					DhcpRelay: cfrex.F(cfrex.MagicLanStaticAddressingDhcpRelayParam{
						ServerAddresses: cfrex.F([]string{"192.0.2.1"}),
					}),
					DhcpServer: cfrex.F(cfrex.MagicLanStaticAddressingDhcpServerParam{
						DhcpPoolEnd:   cfrex.F("192.0.2.1"),
						DhcpPoolStart: cfrex.F("192.0.2.1"),
						DNSServer:     cfrex.F("192.0.2.1"),
						DNSServers:    cfrex.F([]string{"192.0.2.1"}),
						Reservations: cfrex.F(map[string]string{
							"00:11:22:33:44:55": "192.0.2.100",
							"AA:BB:CC:DD:EE:FF": "192.168.1.101",
						}),
					}),
					SecondaryAddress: cfrex.F("192.0.2.0/24"),
					VirtualAddress:   cfrex.F("192.0.2.0/24"),
				}),
				VlanTag: cfrex.F(int64(42)),
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
