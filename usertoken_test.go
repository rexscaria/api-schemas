// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/cf-rex-go"
	"github.com/stainless-sdks/cf-rex-go/internal/testutil"
	"github.com/stainless-sdks/cf-rex-go/option"
)

func TestUserTokenNewWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Tokens.New(context.TODO(), cfrex.UserTokenNewParams{
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
				Resources: cfrex.F(map[string]string{
					"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
					"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
				}),
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
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserTokenGet(t *testing.T) {
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
	_, err := client.User.Tokens.Get(context.TODO(), "ed17574386854bf78a67040be0a770b0")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserTokenUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Tokens.Update(
		context.TODO(),
		"ed17574386854bf78a67040be0a770b0",
		cfrex.UserTokenUpdateParams{
			IamTokenBody: cfrex.IamTokenBodyParam{
				IamTokenBaseParam: cfrex.IamTokenBaseParam{
					Condition: cfrex.F(cfrex.IamConditionParam{
						RequestIP: cfrex.F(cfrex.IamConditionRequestIPParam{
							In:    cfrex.F([]string{"123.123.123.0/24", "2606:4700::/32"}),
							NotIn: cfrex.F([]string{"123.123.123.100/24", "2606:4700:4700::/48"}),
						}),
					}),
					ExpiresOn: cfrex.F(time.Now()),
					Name:      cfrex.F("readonly token"),
					NotBefore: cfrex.F(time.Now()),
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
						Resources: cfrex.F(map[string]string{
							"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*",
							"com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4": "*",
						}),
					}}),
					Status: cfrex.F(cfrex.IamStatusActive),
				},
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

func TestUserTokenListWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Tokens.List(context.TODO(), cfrex.UserTokenListParams{
		Direction: cfrex.F(cfrex.UserTokenListParamsDirectionAsc),
		Page:      cfrex.F(1.000000),
		PerPage:   cfrex.F(5.000000),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserTokenDelete(t *testing.T) {
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
	_, err := client.User.Tokens.Delete(
		context.TODO(),
		"ed17574386854bf78a67040be0a770b0",
		cfrex.UserTokenDeleteParams{
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

func TestUserTokenListPermissionGroups(t *testing.T) {
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
	_, err := client.User.Tokens.ListPermissionGroups(context.TODO())
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserTokenRoll(t *testing.T) {
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
	_, err := client.User.Tokens.Roll(
		context.TODO(),
		"ed17574386854bf78a67040be0a770b0",
		cfrex.UserTokenRollParams{
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

func TestUserTokenVerify(t *testing.T) {
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
	_, err := client.User.Tokens.Verify(context.TODO())
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
