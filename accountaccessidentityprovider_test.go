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

func TestAccountAccessIdentityProviderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Access.IdentityProviders.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAccessIdentityProviderNewParams{
			AccountIdentityProviders: cfrex.AzureAdParam(cfrex.AzureAdParam{
				IdentityProviderAccountParam: cfrex.IdentityProviderAccountParam{
					Config: cfrex.F[any](map[string]interface{}{
						"client_id":     "<your client id>",
						"client_secret": "<your client secret>",
						"claims": map[string]interface{}{
							"0": "email_verified",
							"1": "preferred_username",
							"2": "custom_claim_name",
						},
						"email_claim_name":           "custom_claim_name",
						"conditional_access_enabled": true,
						"directory_id":               "<your azure directory uuid>",
						"prompt":                     "login",
						"support_groups":             true,
					}),
					Name: cfrex.F("Widget Corps IDP"),
					Type: cfrex.F(cfrex.IdentityProviderAccountTypeOnetimepin),
					ID:   cfrex.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					ScimConfig: cfrex.F(cfrex.IdentityProviderAccountScimConfigParam{
						Enabled:                cfrex.F(true),
						IdentityUpdateBehavior: cfrex.F(cfrex.IdentityProviderAccountScimConfigIdentityUpdateBehaviorAutomatic),
						SeatDeprovision:        cfrex.F(true),
						UserDeprovision:        cfrex.F(true),
					}),
				},
				Config: cfrex.F(map[string]interface{}{
					"client_id":     "<your client id>",
					"client_secret": "<your client secret>",
					"claims": map[string]interface{}{
						"0": "email_verified",
						"1": "preferred_username",
						"2": "custom_claim_name",
					},
					"email_claim_name":           "custom_claim_name",
					"conditional_access_enabled": true,
					"directory_id":               "<your azure directory uuid>",
					"prompt":                     "login",
					"support_groups":             true,
				}),
			}),
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

func TestAccountAccessIdentityProviderGet(t *testing.T) {
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
	_, err := client.Accounts.Access.IdentityProviders.Get(
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

func TestAccountAccessIdentityProviderUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Access.IdentityProviders.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.AccountAccessIdentityProviderUpdateParams{
			AccountIdentityProviders: cfrex.AzureAdParam(cfrex.AzureAdParam{
				IdentityProviderAccountParam: cfrex.IdentityProviderAccountParam{
					Config: cfrex.F[any](map[string]interface{}{
						"client_id":     "<your client id>",
						"client_secret": "<your client secret>",
						"claims": map[string]interface{}{
							"0": "email_verified",
							"1": "preferred_username",
							"2": "custom_claim_name",
						},
						"email_claim_name":           "custom_claim_name",
						"conditional_access_enabled": true,
						"directory_id":               "<your azure directory uuid>",
						"prompt":                     "login",
						"support_groups":             true,
					}),
					Name: cfrex.F("Widget Corps IDP"),
					Type: cfrex.F(cfrex.IdentityProviderAccountTypeOnetimepin),
					ID:   cfrex.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					ScimConfig: cfrex.F(cfrex.IdentityProviderAccountScimConfigParam{
						Enabled:                cfrex.F(true),
						IdentityUpdateBehavior: cfrex.F(cfrex.IdentityProviderAccountScimConfigIdentityUpdateBehaviorAutomatic),
						SeatDeprovision:        cfrex.F(true),
						UserDeprovision:        cfrex.F(true),
					}),
				},
				Config: cfrex.F(map[string]interface{}{
					"client_id":     "<your client id>",
					"client_secret": "<your client secret>",
					"claims": map[string]interface{}{
						"0": "email_verified",
						"1": "preferred_username",
						"2": "custom_claim_name",
					},
					"email_claim_name":           "custom_claim_name",
					"conditional_access_enabled": true,
					"directory_id":               "<your azure directory uuid>",
					"prompt":                     "login",
					"support_groups":             true,
				}),
			}),
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

func TestAccountAccessIdentityProviderListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Access.IdentityProviders.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAccessIdentityProviderListParams{
			ScimEnabled: cfrex.F("scim_enabled"),
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

func TestAccountAccessIdentityProviderDelete(t *testing.T) {
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
	_, err := client.Accounts.Access.IdentityProviders.Delete(
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
