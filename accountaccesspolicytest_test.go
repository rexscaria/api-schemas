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

func TestAccountAccessPolicyTestGet(t *testing.T) {
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
	_, err := client.Accounts.Access.PolicyTests.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f1a8b3c9d4e5f6789a0b1c2d3e4f5678a9b0c1d2e3f4a5b67890c1d2e3f4b5a6",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAccessPolicyTestStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Access.PolicyTests.Start(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAccessPolicyTestStartParams{
			Policies: cfrex.F([]cfrex.AccountAccessPolicyTestStartParamsPolicyUnion{cfrex.PolicyRequestForAccessParam(cfrex.PolicyRequestForAccessParam{
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
			})}),
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

func TestAccountAccessPolicyTestUsersWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Access.PolicyTests.Users(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f1a8b3c9d4e5f6789a0b1c2d3e4f5678a9b0c1d2e3f4a5b67890c1d2e3f4b5a6",
		cfrex.AccountAccessPolicyTestUsersParams{
			Page:    cfrex.F(int64(0)),
			PerPage: cfrex.F(int64(0)),
			Status:  cfrex.F(cfrex.AccountAccessPolicyTestUsersParamsStatusSuccess),
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
