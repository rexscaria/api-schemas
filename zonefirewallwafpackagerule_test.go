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

func TestZoneFirewallWafPackageRuleGet(t *testing.T) {
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
	_, err := client.Zones.Firewall.Waf.Packages.Rules.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		"a25a9a7e9c00afc1fb2e0245519d725b",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneFirewallWafPackageRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewall.Waf.Packages.Rules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		cfrex.ZoneFirewallWafPackageRuleUpdateParams{
			Mode: cfrex.F(cfrex.ZoneFirewallWafPackageRuleUpdateParamsModeDefault),
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

func TestZoneFirewallWafPackageRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewall.Waf.Packages.Rules.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"a25a9a7e9c00afc1fb2e0245519d725b",
		cfrex.ZoneFirewallWafPackageRuleListParams{
			Description: cfrex.F("SQL injection prevention for SELECT statements"),
			Direction:   cfrex.F(cfrex.ZoneFirewallWafPackageRuleListParamsDirectionAsc),
			GroupID:     cfrex.F("de677e5818985db1285d0e80225f06e5"),
			Match:       cfrex.F(cfrex.ZoneFirewallWafPackageRuleListParamsMatchAny),
			Mode:        cfrex.F(cfrex.ZoneFirewallWafPackageRuleListParamsModeDis),
			Order:       cfrex.F(cfrex.ZoneFirewallWafPackageRuleListParamsOrderPriority),
			Page:        cfrex.F(1.000000),
			PerPage:     cfrex.F(5.000000),
			Priority:    cfrex.F("priority"),
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
