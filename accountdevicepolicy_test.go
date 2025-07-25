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

func TestAccountDevicePolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policy.New(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.AccountDevicePolicyNewParams{
			Match:               cfrex.F(`identity.email == "test@cloudflare.com"`),
			Name:                cfrex.F("Allow Developers"),
			Precedence:          cfrex.F(100.000000),
			AllowModeSwitch:     cfrex.F(true),
			AllowUpdates:        cfrex.F(true),
			AllowedToLeave:      cfrex.F(true),
			AutoConnect:         cfrex.F(0.000000),
			CaptivePortal:       cfrex.F(180.000000),
			Description:         cfrex.F("Policy for test teams."),
			DisableAutoFallback: cfrex.F(true),
			Enabled:             cfrex.F(true),
			Exclude: cfrex.F([]cfrex.SplitTunnelUnionParam{cfrex.SplitTunnelTeamsDevicesExcludeSplitTunnelWithAddressParam{
				Address:     cfrex.F("192.0.2.0/24"),
				Description: cfrex.F("Exclude testing domains from the tunnel"),
			}}),
			ExcludeOfficeIPs: cfrex.F(true),
			Include: cfrex.F([]cfrex.SplitTunnelIncludeUnionParam{cfrex.SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressParam{
				Address:     cfrex.F("192.0.2.0/24"),
				Description: cfrex.F("Include testing domains in the tunnel"),
			}}),
			LanAllowMinutes:            cfrex.F(30.000000),
			LanAllowSubnetSize:         cfrex.F(24.000000),
			RegisterInterfaceIPWithDNS: cfrex.F(true),
			SccmVpnBoundarySupport:     cfrex.F(false),
			ServiceModeV2: cfrex.F(cfrex.ServiceModeV2Param{
				Mode: cfrex.F("proxy"),
				Port: cfrex.F(3000.000000),
			}),
			SupportURL:     cfrex.F("https://1.1.1.1/help"),
			SwitchLocked:   cfrex.F(true),
			TunnelProtocol: cfrex.F("wireguard"),
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

func TestAccountDevicePolicyGet(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policy.Get(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDevicePolicyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policy.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.AccountDevicePolicyUpdateParams{
			AllowModeSwitch:     cfrex.F(true),
			AllowUpdates:        cfrex.F(true),
			AllowedToLeave:      cfrex.F(true),
			AutoConnect:         cfrex.F(0.000000),
			CaptivePortal:       cfrex.F(180.000000),
			DisableAutoFallback: cfrex.F(true),
			Exclude: cfrex.F([]cfrex.SplitTunnelUnionParam{cfrex.SplitTunnelTeamsDevicesExcludeSplitTunnelWithAddressParam{
				Address:     cfrex.F("192.0.2.0/24"),
				Description: cfrex.F("Exclude testing domains from the tunnel"),
			}}),
			ExcludeOfficeIPs: cfrex.F(true),
			Include: cfrex.F([]cfrex.SplitTunnelIncludeUnionParam{cfrex.SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressParam{
				Address:     cfrex.F("192.0.2.0/24"),
				Description: cfrex.F("Include testing domains in the tunnel"),
			}}),
			LanAllowMinutes:            cfrex.F(30.000000),
			LanAllowSubnetSize:         cfrex.F(24.000000),
			RegisterInterfaceIPWithDNS: cfrex.F(true),
			SccmVpnBoundarySupport:     cfrex.F(false),
			ServiceModeV2: cfrex.F(cfrex.ServiceModeV2Param{
				Mode: cfrex.F("proxy"),
				Port: cfrex.F(3000.000000),
			}),
			SupportURL:     cfrex.F("https://1.1.1.1/help"),
			SwitchLocked:   cfrex.F(true),
			TunnelProtocol: cfrex.F("wireguard"),
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

func TestAccountDevicePolicyDelete(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policy.Delete(
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

func TestAccountDevicePolicyGetByID(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policy.GetByID(
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

func TestAccountDevicePolicyUpdateByIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Devices.Policy.UpdateByID(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.AccountDevicePolicyUpdateByIDParams{
			AllowModeSwitch:     cfrex.F(true),
			AllowUpdates:        cfrex.F(true),
			AllowedToLeave:      cfrex.F(true),
			AutoConnect:         cfrex.F(0.000000),
			CaptivePortal:       cfrex.F(180.000000),
			Description:         cfrex.F("Policy for test teams."),
			DisableAutoFallback: cfrex.F(true),
			Enabled:             cfrex.F(true),
			Exclude: cfrex.F([]cfrex.SplitTunnelUnionParam{cfrex.SplitTunnelTeamsDevicesExcludeSplitTunnelWithAddressParam{
				Address:     cfrex.F("192.0.2.0/24"),
				Description: cfrex.F("Exclude testing domains from the tunnel"),
			}}),
			ExcludeOfficeIPs: cfrex.F(true),
			Include: cfrex.F([]cfrex.SplitTunnelIncludeUnionParam{cfrex.SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressParam{
				Address:     cfrex.F("192.0.2.0/24"),
				Description: cfrex.F("Include testing domains in the tunnel"),
			}}),
			LanAllowMinutes:            cfrex.F(30.000000),
			LanAllowSubnetSize:         cfrex.F(24.000000),
			Match:                      cfrex.F(`identity.email == "test@cloudflare.com"`),
			Name:                       cfrex.F("Allow Developers"),
			Precedence:                 cfrex.F(100.000000),
			RegisterInterfaceIPWithDNS: cfrex.F(true),
			SccmVpnBoundarySupport:     cfrex.F(false),
			ServiceModeV2: cfrex.F(cfrex.ServiceModeV2Param{
				Mode: cfrex.F("proxy"),
				Port: cfrex.F(3000.000000),
			}),
			SupportURL:     cfrex.F("https://1.1.1.1/help"),
			SwitchLocked:   cfrex.F(true),
			TunnelProtocol: cfrex.F("wireguard"),
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
