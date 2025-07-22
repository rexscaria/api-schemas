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

func TestZoneAccessAppPolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Access.Apps.Policies.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.ZoneAccessAppPolicyNewParams{
			Decision: cfrex.F(cfrex.AppDecisionAllow),
			Include: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
				Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
					ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
				}),
			}}),
			Name: cfrex.F("Allow devs"),
			ApprovalGroups: cfrex.F([]cfrex.ApprovalGroupAppParam{{
				ApprovalsNeeded: cfrex.F(1.000000),
				EmailAddresses:  cfrex.F([]interface{}{"test1@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUuid:   cfrex.F("email_list_uuid"),
			}, {
				ApprovalsNeeded: cfrex.F(3.000000),
				EmailAddresses:  cfrex.F([]interface{}{"test@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUuid:   cfrex.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
			}}),
			ApprovalRequired: cfrex.F(true),
			Exclude: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
				Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
					ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
				}),
			}}),
			IsolationRequired:            cfrex.F(false),
			Precedence:                   cfrex.F(int64(0)),
			PurposeJustificationPrompt:   cfrex.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cfrex.F(true),
			Require: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
				Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
					ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
				}),
			}}),
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

func TestZoneAccessAppPolicyGet(t *testing.T) {
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
	_, err := client.Zones.Access.Apps.Policies.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
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

func TestZoneAccessAppPolicyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Access.Apps.Policies.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.ZoneAccessAppPolicyUpdateParams{
			Decision: cfrex.F(cfrex.AppDecisionAllow),
			Include: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
				Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
					ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
				}),
			}}),
			Name: cfrex.F("Allow devs"),
			ApprovalGroups: cfrex.F([]cfrex.ApprovalGroupAppParam{{
				ApprovalsNeeded: cfrex.F(1.000000),
				EmailAddresses:  cfrex.F([]interface{}{"test1@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUuid:   cfrex.F("email_list_uuid"),
			}, {
				ApprovalsNeeded: cfrex.F(3.000000),
				EmailAddresses:  cfrex.F([]interface{}{"test@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUuid:   cfrex.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
			}}),
			ApprovalRequired: cfrex.F(true),
			Exclude: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
				Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
					ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
				}),
			}}),
			IsolationRequired:            cfrex.F(false),
			Precedence:                   cfrex.F(int64(0)),
			PurposeJustificationPrompt:   cfrex.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cfrex.F(true),
			Require: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
				Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
					ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
				}),
			}}),
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

func TestZoneAccessAppPolicyList(t *testing.T) {
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
	_, err := client.Zones.Access.Apps.Policies.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
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

func TestZoneAccessAppPolicyDelete(t *testing.T) {
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
	_, err := client.Zones.Access.Apps.Policies.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
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
