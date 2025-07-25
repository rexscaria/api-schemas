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

func TestAccountTokenNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Tokens.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountTokenNewParams{
			IamCreatePayload: cfrex.IamCreatePayloadParam{
				Name: cfrex.F("readonly token"),
				Policies: cfrex.F([]cfrex.IamPolicyWithPermissionGroupsAndResourcesParam{{
					Effect: cfrex.F(cfrex.IamPolicyWithPermissionGroupsAndResourcesEffectAllow),
					PermissionGroups: cfrex.F([]cfrex.IamPermissionGroupParam{{
						ID: cfrex.F("c8fed203ed3043cba015a93ad1616f1f"),
						Meta: cfrex.F(cfrex.IamPermissionGroupMetaParam{
							Key:   cfrex.F("key"),
							Value: cfrex.F("value"),
						}),
					}, {
						ID: cfrex.F("82e64a83756745bbbb1c9c2701bf816b"),
						Meta: cfrex.F(cfrex.IamPermissionGroupMetaParam{
							Key:   cfrex.F("key"),
							Value: cfrex.F("value"),
						}),
					}}),
					Resources: cfrex.F[cfrex.IamPolicyWithPermissionGroupsAndResourcesResourcesUnionParam](cfrex.IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectStringParam(map[string]string{
						"foo": "string",
					})),
				}}),
				Condition: cfrex.F(cfrex.IamConditionParam{
					RequestIP: cfrex.F(cfrex.IamConditionRequestIPParam{
						In:    cfrex.F([]string{"123.123.123.0/24", "2606:4700::/32"}),
						NotIn: cfrex.F([]string{"123.123.123.100/24", "2606:4700:4700::/48"}),
					}),
				}),
				ExpiresOn: cfrex.F(time.Now()),
				NotBefore: cfrex.F(time.Now()),
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

func TestAccountTokenGet(t *testing.T) {
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
	_, err := client.Accounts.Tokens.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"ed17574386854bf78a67040be0a770b0",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountTokenUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Tokens.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"ed17574386854bf78a67040be0a770b0",
		cfrex.AccountTokenUpdateParams{
			IamTokenBody: cfrex.IamTokenBodyParam{
				Name: cfrex.F("readonly token"),
				Policies: cfrex.F([]cfrex.IamPolicyWithPermissionGroupsAndResourcesParam{{
					Effect: cfrex.F(cfrex.IamPolicyWithPermissionGroupsAndResourcesEffectAllow),
					PermissionGroups: cfrex.F([]cfrex.IamPermissionGroupParam{{
						ID: cfrex.F("c8fed203ed3043cba015a93ad1616f1f"),
						Meta: cfrex.F(cfrex.IamPermissionGroupMetaParam{
							Key:   cfrex.F("key"),
							Value: cfrex.F("value"),
						}),
					}, {
						ID: cfrex.F("82e64a83756745bbbb1c9c2701bf816b"),
						Meta: cfrex.F(cfrex.IamPermissionGroupMetaParam{
							Key:   cfrex.F("key"),
							Value: cfrex.F("value"),
						}),
					}}),
					Resources: cfrex.F[cfrex.IamPolicyWithPermissionGroupsAndResourcesResourcesUnionParam](cfrex.IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectStringParam(map[string]string{
						"foo": "string",
					})),
				}}),
				Condition: cfrex.F(cfrex.IamConditionParam{
					RequestIP: cfrex.F(cfrex.IamConditionRequestIPParam{
						In:    cfrex.F([]string{"123.123.123.0/24", "2606:4700::/32"}),
						NotIn: cfrex.F([]string{"123.123.123.100/24", "2606:4700:4700::/48"}),
					}),
				}),
				ExpiresOn: cfrex.F(time.Now()),
				NotBefore: cfrex.F(time.Now()),
				Status:    cfrex.F(cfrex.IamTokenBodyStatusActive),
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

func TestAccountTokenListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Tokens.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountTokenListParams{
			Direction: cfrex.F(cfrex.AccountTokenListParamsDirectionDesc),
			Page:      cfrex.F(1.000000),
			PerPage:   cfrex.F(5.000000),
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

func TestAccountTokenDelete(t *testing.T) {
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
	_, err := client.Accounts.Tokens.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"ed17574386854bf78a67040be0a770b0",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountTokenListPermissionGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Tokens.ListPermissionGroups(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountTokenListPermissionGroupsParams{
			Name:  cfrex.F("Account%20Settings%20Write"),
			Scope: cfrex.F("com.cloudflare.api.account.zone"),
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

func TestAccountTokenRoll(t *testing.T) {
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
	_, err := client.Accounts.Tokens.Roll(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"ed17574386854bf78a67040be0a770b0",
		cfrex.AccountTokenRollParams{
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

func TestAccountTokenVerify(t *testing.T) {
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
	_, err := client.Accounts.Tokens.Verify(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
