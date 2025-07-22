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

func TestZoneAccessOrganizationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Access.Organizations.New(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.ZoneAccessOrganizationNewParams{
			AuthDomain:   cfrex.F("test.cloudflareaccess.com"),
			Name:         cfrex.F("Widget Corps Internal Applications"),
			IsUiReadOnly: cfrex.F(true),
			LoginDesign: cfrex.F(cfrex.LoginDesignParam{
				BackgroundColor: cfrex.F("#c5ed1b"),
				FooterText:      cfrex.F("This is an example description."),
				HeaderText:      cfrex.F("This is an example description."),
				LogoPath:        cfrex.F("https://example.com/logo.png"),
				TextColor:       cfrex.F("#c5ed1b"),
			}),
			UiReadOnlyToggleReason:         cfrex.F("Temporarily turn off the UI read only lock to make a change via the UI"),
			UserSeatExpirationInactiveTime: cfrex.F("720h"),
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

func TestZoneAccessOrganizationGet(t *testing.T) {
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
	_, err := client.Zones.Access.Organizations.Get(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneAccessOrganizationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Access.Organizations.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.ZoneAccessOrganizationUpdateParams{
			AuthDomain:   cfrex.F("test.cloudflareaccess.com"),
			IsUiReadOnly: cfrex.F(true),
			LoginDesign: cfrex.F(cfrex.LoginDesignParam{
				BackgroundColor: cfrex.F("#c5ed1b"),
				FooterText:      cfrex.F("This is an example description."),
				HeaderText:      cfrex.F("This is an example description."),
				LogoPath:        cfrex.F("https://example.com/logo.png"),
				TextColor:       cfrex.F("#c5ed1b"),
			}),
			Name:                           cfrex.F("Widget Corps Internal Applications"),
			UiReadOnlyToggleReason:         cfrex.F("Temporarily turn off the UI read only lock to make a change via the UI"),
			UserSeatExpirationInactiveTime: cfrex.F("720h"),
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

func TestZoneAccessOrganizationRevokeUser(t *testing.T) {
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
	_, err := client.Zones.Access.Organizations.RevokeUser(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.ZoneAccessOrganizationRevokeUserParams{
			Email: cfrex.F("test@example.com"),
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
