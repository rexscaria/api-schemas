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

func TestAccountAccessPolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Access.Policies.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAccessPolicyNewParams{
			PolicyRequestForAccess: cfrex.PolicyRequestForAccessParam{
				BasePolicyRequestParam: cfrex.BasePolicyRequestParam{
					Decision: cfrex.F(cfrex.AccessDecisionAllow),
					Include: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
						Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
							ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
						}),
					}}),
					Name: cfrex.F("Allow devs"),
					Exclude: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
						Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
							ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
						}),
					}}),
					Require: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
						Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
							ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
						}),
					}}),
				},
				ApprovalGroups: cfrex.F([]cfrex.ApprovalGroupEmailParam{{
					ApprovalsNeeded: cfrex.F(1.000000),
					EmailAddresses:  cfrex.F([]string{"test1@cloudflare.com", "test2@cloudflare.com"}),
					EmailListUuid:   cfrex.F("email_list_uuid"),
				}, {
					ApprovalsNeeded: cfrex.F(3.000000),
					EmailAddresses:  cfrex.F([]string{"test@cloudflare.com", "test2@cloudflare.com"}),
					EmailListUuid:   cfrex.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
				}}),
				ApprovalRequired:             cfrex.F(true),
				IsolationRequired:            cfrex.F(false),
				PurposeJustificationPrompt:   cfrex.F("Please enter a justification for entering this protected domain."),
				PurposeJustificationRequired: cfrex.F(true),
				SessionDuration:              cfrex.F("24h"),
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

func TestAccountAccessPolicyGet(t *testing.T) {
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
	_, err := client.Accounts.Access.Policies.Get(
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

func TestAccountAccessPolicyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Access.Policies.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.AccountAccessPolicyUpdateParams{
			PolicyRequestForAccess: cfrex.PolicyRequestForAccessParam{
				BasePolicyRequestParam: cfrex.BasePolicyRequestParam{
					Decision: cfrex.F(cfrex.AccessDecisionAllow),
					Include: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
						Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
							ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
						}),
					}}),
					Name: cfrex.F("Allow devs"),
					Exclude: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
						Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
							ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
						}),
					}}),
					Require: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
						Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
							ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
						}),
					}}),
				},
				ApprovalGroups: cfrex.F([]cfrex.ApprovalGroupEmailParam{{
					ApprovalsNeeded: cfrex.F(1.000000),
					EmailAddresses:  cfrex.F([]string{"test1@cloudflare.com", "test2@cloudflare.com"}),
					EmailListUuid:   cfrex.F("email_list_uuid"),
				}, {
					ApprovalsNeeded: cfrex.F(3.000000),
					EmailAddresses:  cfrex.F([]string{"test@cloudflare.com", "test2@cloudflare.com"}),
					EmailListUuid:   cfrex.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
				}}),
				ApprovalRequired:             cfrex.F(true),
				IsolationRequired:            cfrex.F(false),
				PurposeJustificationPrompt:   cfrex.F("Please enter a justification for entering this protected domain."),
				PurposeJustificationRequired: cfrex.F(true),
				SessionDuration:              cfrex.F("24h"),
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

func TestAccountAccessPolicyList(t *testing.T) {
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
	_, err := client.Accounts.Access.Policies.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAccessPolicyDelete(t *testing.T) {
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
	_, err := client.Accounts.Access.Policies.Delete(
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
