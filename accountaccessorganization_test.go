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

func TestAccountAccessOrganizationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Access.Organizations.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAccessOrganizationNewParams{
			AuthDomain:               cfrex.F("test.cloudflareaccess.com"),
			Name:                     cfrex.F("Widget Corps Internal Applications"),
			AllowAuthenticateViaWarp: cfrex.F(true),
			AutoRedirectToIdentity:   cfrex.F(true),
			IsUiReadOnly:             cfrex.F(true),
			LoginDesign: cfrex.F(cfrex.LoginDesignParam{
				BackgroundColor: cfrex.F("#c5ed1b"),
				FooterText:      cfrex.F("This is an example description."),
				HeaderText:      cfrex.F("This is an example description."),
				LogoPath:        cfrex.F("https://example.com/logo.png"),
				TextColor:       cfrex.F("#c5ed1b"),
			}),
			SessionDuration:                cfrex.F("24h"),
			UiReadOnlyToggleReason:         cfrex.F("Temporarily turn off the UI read only lock to make a change via the UI"),
			UserSeatExpirationInactiveTime: cfrex.F("730h"),
			WarpAuthSessionDuration:        cfrex.F("24h"),
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

func TestAccountAccessOrganizationGet(t *testing.T) {
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
	_, err := client.Accounts.Access.Organizations.Get(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAccessOrganizationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Access.Organizations.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAccessOrganizationUpdateParams{
			AllowAuthenticateViaWarp: cfrex.F(true),
			AuthDomain:               cfrex.F("test.cloudflareaccess.com"),
			AutoRedirectToIdentity:   cfrex.F(true),
			CustomPages: cfrex.F(cfrex.CustomPagesParam{
				Forbidden:      cfrex.F("699d98642c564d2e855e9661899b7252"),
				IdentityDenied: cfrex.F("699d98642c564d2e855e9661899b7252"),
			}),
			IsUiReadOnly: cfrex.F(true),
			LoginDesign: cfrex.F(cfrex.LoginDesignParam{
				BackgroundColor: cfrex.F("#c5ed1b"),
				FooterText:      cfrex.F("This is an example description."),
				HeaderText:      cfrex.F("This is an example description."),
				LogoPath:        cfrex.F("https://example.com/logo.png"),
				TextColor:       cfrex.F("#c5ed1b"),
			}),
			Name:                           cfrex.F("Widget Corps Internal Applications"),
			SessionDuration:                cfrex.F("24h"),
			UiReadOnlyToggleReason:         cfrex.F("Temporarily turn off the UI read only lock to make a change via the UI"),
			UserSeatExpirationInactiveTime: cfrex.F("730h"),
			WarpAuthSessionDuration:        cfrex.F("24h"),
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

func TestAccountAccessOrganizationRevokeUserWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Access.Organizations.RevokeUser(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAccessOrganizationRevokeUserParams{
			Email:             cfrex.F("test@example.com"),
			QueryDevices:      cfrex.F(true),
			BodyDevices:       cfrex.F(true),
			UserUid:           cfrex.F("699d98642c564d2e855e9661899b7252"),
			WarpSessionReauth: cfrex.F(true),
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
