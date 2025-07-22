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

func TestUserFirewallAccessRuleRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Firewall.AccessRules.Rules.New(context.TODO(), cfrex.UserFirewallAccessRuleRuleNewParams{
		Configuration: cfrex.F[cfrex.FirewallRuleConfigurationUnionParam](cfrex.FirewallRuleConfigurationFirewallIPConfigurationParam{
			Target: cfrex.F(cfrex.FirewallRuleConfigurationFirewallIPConfigurationTargetIP),
			Value:  cfrex.F("198.51.100.4"),
		}),
		Mode:  cfrex.F(cfrex.FirewallSchemasModeChallenge),
		Notes: cfrex.F("This rule is enabled because of an event that occurred on date X."),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserFirewallAccessRuleRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Firewall.AccessRules.Rules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.UserFirewallAccessRuleRuleUpdateParams{
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

func TestUserFirewallAccessRuleRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Firewall.AccessRules.Rules.List(context.TODO(), cfrex.UserFirewallAccessRuleRuleListParams{
		Configuration: cfrex.F(cfrex.UserFirewallAccessRuleRuleListParamsConfiguration{
			Target: cfrex.F(cfrex.UserFirewallAccessRuleRuleListParamsConfigurationTargetIP),
			Value:  cfrex.F("198.51.100.4"),
		}),
		Direction: cfrex.F(cfrex.UserFirewallAccessRuleRuleListParamsDirectionDesc),
		Match:     cfrex.F(cfrex.UserFirewallAccessRuleRuleListParamsMatchAny),
		Mode:      cfrex.F(cfrex.FirewallSchemasModeChallenge),
		Notes:     cfrex.F("my note"),
		Order:     cfrex.F(cfrex.UserFirewallAccessRuleRuleListParamsOrderMode),
		Page:      cfrex.F(1.000000),
		PerPage:   cfrex.F(20.000000),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserFirewallAccessRuleRuleDelete(t *testing.T) {
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
	_, err := client.User.Firewall.AccessRules.Rules.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.UserFirewallAccessRuleRuleDeleteParams{
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
