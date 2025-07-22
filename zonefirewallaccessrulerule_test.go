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

func TestZoneFirewallAccessRuleRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewall.AccessRules.Rules.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneFirewallAccessRuleRuleNewParams{
			Configuration: cfrex.F[cfrex.FirewallRuleConfigurationUnionParam](cfrex.FirewallRuleConfigurationFirewallIPConfigurationParam{
				Target: cfrex.F(cfrex.FirewallRuleConfigurationFirewallIPConfigurationTargetIP),
				Value:  cfrex.F("198.51.100.4"),
			}),
			Mode:  cfrex.F(cfrex.FirewallSchemasModeChallenge),
			Notes: cfrex.F("This rule is enabled because of an event that occurred on date X."),
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

func TestZoneFirewallAccessRuleRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewall.AccessRules.Rules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneFirewallAccessRuleRuleUpdateParams{
			Mode:  cfrex.F(cfrex.FirewallSchemasModeChallenge),
			Notes: cfrex.F("This rule is enabled because of an event that occurred on date X."),
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

func TestZoneFirewallAccessRuleRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewall.AccessRules.Rules.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneFirewallAccessRuleRuleListParams{
			Configuration: cfrex.F(cfrex.ZoneFirewallAccessRuleRuleListParamsConfiguration{
				Target: cfrex.F(cfrex.ZoneFirewallAccessRuleRuleListParamsConfigurationTargetIP),
				Value:  cfrex.F("198.51.100.4"),
			}),
			Direction: cfrex.F(cfrex.ZoneFirewallAccessRuleRuleListParamsDirectionDesc),
			Match:     cfrex.F(cfrex.ZoneFirewallAccessRuleRuleListParamsMatchAny),
			Mode:      cfrex.F(cfrex.FirewallSchemasModeChallenge),
			Notes:     cfrex.F("my note"),
			Order:     cfrex.F(cfrex.ZoneFirewallAccessRuleRuleListParamsOrderMode),
			Page:      cfrex.F(1.000000),
			PerPage:   cfrex.F(20.000000),
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

func TestZoneFirewallAccessRuleRuleDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Firewall.AccessRules.Rules.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneFirewallAccessRuleRuleDeleteParams{
			Cascade: cfrex.F(cfrex.ZoneFirewallAccessRuleRuleDeleteParamsCascadeNone),
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
