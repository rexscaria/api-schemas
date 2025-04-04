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

func TestZoneDNSSettingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.DNSSettings.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneDNSSettingUpdateParams{
			SettingsZone: cfrex.SettingsZoneParam{
				DNSSettingsParam: cfrex.DNSSettingsParam{
					FlattenAllCnames: cfrex.F(false),
					FoundationDNS:    cfrex.F(false),
					InternalDNS: cfrex.F(cfrex.DNSSettingsInternalDNSParam{
						ReferenceZoneID: cfrex.F("reference_zone_id"),
					}),
					MultiProvider:      cfrex.F(false),
					NsTtl:              cfrex.F(86400.000000),
					SecondaryOverrides: cfrex.F(false),
					Soa: cfrex.F(cfrex.DNSSettingsSoaParam{
						Expire:  cfrex.F(604800.000000),
						MinTtl:  cfrex.F(1800.000000),
						Mname:   cfrex.F("kristina.ns.cloudflare.com"),
						Refresh: cfrex.F(10000.000000),
						Retry:   cfrex.F(2400.000000),
						Rname:   cfrex.F("admin.example.com"),
						Ttl:     cfrex.F(3600.000000),
					}),
					ZoneMode: cfrex.F(cfrex.DNSSettingsZoneModeStandard),
				},
				Nameservers: cfrex.F(cfrex.SettingsZoneNameserversParam{
					Type:  cfrex.F(cfrex.SettingsZoneNameserversTypeCloudflareStandard),
					NsSet: cfrex.F(int64(1)),
				}),
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

func TestZoneDNSSettingShow(t *testing.T) {
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
	_, err := client.Zones.DNSSettings.Show(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
