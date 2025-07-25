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

func TestAccountChallengeWidgetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Challenges.Widgets.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountChallengeWidgetNewParams{
			Domains:        cfrex.F([]string{"203.0.113.1", "cloudflare.com", "blog.example.com"}),
			Mode:           cfrex.F(cfrex.TurnstileWidgetModeInvisible),
			Name:           cfrex.F("blog.cloudflare.com login form"),
			Direction:      cfrex.F(cfrex.AccountChallengeWidgetNewParamsDirectionAsc),
			Order:          cfrex.F(cfrex.AccountChallengeWidgetNewParamsOrderID),
			Page:           cfrex.F(1.000000),
			PerPage:        cfrex.F(5.000000),
			BotFightMode:   cfrex.F(false),
			ClearanceLevel: cfrex.F(cfrex.TurnstileClearanceLevelInteractive),
			EphemeralID:    cfrex.F(false),
			Offlabel:       cfrex.F(false),
			Region:         cfrex.F(cfrex.TurnstileRegionWorld),
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

func TestAccountChallengeWidgetGet(t *testing.T) {
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
	_, err := client.Accounts.Challenges.Widgets.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0x4AAF00AAAABn0R22HWm-YUc",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountChallengeWidgetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Challenges.Widgets.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0x4AAF00AAAABn0R22HWm-YUc",
		cfrex.AccountChallengeWidgetUpdateParams{
			Domains:        cfrex.F([]string{"203.0.113.1", "cloudflare.com", "blog.example.com"}),
			Mode:           cfrex.F(cfrex.TurnstileWidgetModeInvisible),
			Name:           cfrex.F("blog.cloudflare.com login form"),
			BotFightMode:   cfrex.F(false),
			ClearanceLevel: cfrex.F(cfrex.TurnstileClearanceLevelInteractive),
			EphemeralID:    cfrex.F(false),
			Offlabel:       cfrex.F(false),
			Region:         cfrex.F(cfrex.TurnstileRegionWorld),
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

func TestAccountChallengeWidgetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Challenges.Widgets.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountChallengeWidgetListParams{
			Direction: cfrex.F(cfrex.AccountChallengeWidgetListParamsDirectionAsc),
			Order:     cfrex.F(cfrex.AccountChallengeWidgetListParamsOrderID),
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

func TestAccountChallengeWidgetDelete(t *testing.T) {
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
	_, err := client.Accounts.Challenges.Widgets.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0x4AAF00AAAABn0R22HWm-YUc",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountChallengeWidgetRotateSecretWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Challenges.Widgets.RotateSecret(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0x4AAF00AAAABn0R22HWm-YUc",
		cfrex.AccountChallengeWidgetRotateSecretParams{
			InvalidateImmediately: cfrex.F(true),
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
