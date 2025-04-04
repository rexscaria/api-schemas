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

func TestAccountIamResourceGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Iam.ResourceGroups.New(
		context.TODO(),
		"eb78d65290b24279ba6f44721b3ea3c4",
		cfrex.AccountIamResourceGroupNewParams{
			IamCreateResourceGroup: cfrex.IamCreateResourceGroupParam{
				Scope: cfrex.F(cfrex.IamCreateScopeParam{
					Key: cfrex.F[any]("com.cloudflare.api.account.eb78d65290b24279ba6f44721b3ea3c4"),
					Objects: cfrex.F([]cfrex.IamCreateScopeObjectParam{{
						Key: cfrex.F[any]("com.cloudflare.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
					}}),
				}),
				Meta: cfrex.F[any](map[string]interface{}{
					"editable": "false",
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

func TestAccountIamResourceGroupGet(t *testing.T) {
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
	_, err := client.Accounts.Iam.ResourceGroups.Get(
		context.TODO(),
		"eb78d65290b24279ba6f44721b3ea3c4",
		"6d7f2f5f5b1d4a0e9081fdc98d432fd1",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountIamResourceGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Iam.ResourceGroups.Update(
		context.TODO(),
		"eb78d65290b24279ba6f44721b3ea3c4",
		"6d7f2f5f5b1d4a0e9081fdc98d432fd1",
		cfrex.AccountIamResourceGroupUpdateParams{
			IamCreateResourceGroup: cfrex.IamCreateResourceGroupParam{
				Scope: cfrex.F(cfrex.IamCreateScopeParam{
					Key: cfrex.F[any]("com.cloudflare.api.account.eb78d65290b24279ba6f44721b3ea3c4"),
					Objects: cfrex.F([]cfrex.IamCreateScopeObjectParam{{
						Key: cfrex.F[any]("com.cloudflare.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
					}}),
				}),
				Meta: cfrex.F[any](map[string]interface{}{
					"editable": "false",
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

func TestAccountIamResourceGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Iam.ResourceGroups.List(
		context.TODO(),
		"eb78d65290b24279ba6f44721b3ea3c4",
		cfrex.AccountIamResourceGroupListParams{
			ID:      cfrex.F("6d7f2f5f5b1d4a0e9081fdc98d432fd1"),
			Name:    cfrex.F("NameOfTheResourceGroup"),
			Page:    cfrex.F(1.000000),
			PerPage: cfrex.F(5.000000),
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

func TestAccountIamResourceGroupDelete(t *testing.T) {
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
	_, err := client.Accounts.Iam.ResourceGroups.Delete(
		context.TODO(),
		"eb78d65290b24279ba6f44721b3ea3c4",
		"6d7f2f5f5b1d4a0e9081fdc98d432fd1",
		cfrex.AccountIamResourceGroupDeleteParams{
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
