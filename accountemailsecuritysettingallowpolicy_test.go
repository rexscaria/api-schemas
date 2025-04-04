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

func TestAccountEmailSecuritySettingAllowPolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.EmailSecurity.Settings.AllowPolicies.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountEmailSecuritySettingAllowPolicyNewParams{
			IsAcceptableSender: cfrex.F(false),
			IsExemptRecipient:  cfrex.F(false),
			IsRegex:            cfrex.F(false),
			IsTrustedSender:    cfrex.F(true),
			Pattern:            cfrex.F("test@example.com"),
			PatternType:        cfrex.F(cfrex.PatternTypeEmail),
			VerifySender:       cfrex.F(true),
			Comments:           cfrex.F("Trust all messages send from test@example.com"),
			IsRecipient:        cfrex.F(false),
			IsSender:           cfrex.F(true),
			IsSpoof:            cfrex.F(false),
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

func TestAccountEmailSecuritySettingAllowPolicyGet(t *testing.T) {
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
	_, err := client.Accounts.EmailSecurity.Settings.AllowPolicies.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		int64(2401),
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountEmailSecuritySettingAllowPolicyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.EmailSecurity.Settings.AllowPolicies.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		int64(2401),
		cfrex.AccountEmailSecuritySettingAllowPolicyUpdateParams{
			Comments:           cfrex.F("comments"),
			IsAcceptableSender: cfrex.F(true),
			IsExemptRecipient:  cfrex.F(true),
			IsRegex:            cfrex.F(true),
			IsTrustedSender:    cfrex.F(true),
			Pattern:            cfrex.F("x"),
			PatternType:        cfrex.F(cfrex.AccountEmailSecuritySettingAllowPolicyUpdateParamsPatternTypeEmail),
			VerifySender:       cfrex.F(true),
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

func TestAccountEmailSecuritySettingAllowPolicyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.EmailSecurity.Settings.AllowPolicies.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountEmailSecuritySettingAllowPolicyListParams{
			Direction:          cfrex.F(cfrex.SortingDirectionAsc),
			IsAcceptableSender: cfrex.F(true),
			IsExemptRecipient:  cfrex.F(true),
			IsRecipient:        cfrex.F(true),
			IsSender:           cfrex.F(true),
			IsSpoof:            cfrex.F(true),
			IsTrustedSender:    cfrex.F(true),
			Order:              cfrex.F(cfrex.AccountEmailSecuritySettingAllowPolicyListParamsOrderPattern),
			Page:               cfrex.F(int64(1)),
			PatternType:        cfrex.F(cfrex.PatternTypeEmail),
			PerPage:            cfrex.F(int64(1)),
			Search:             cfrex.F("search"),
			VerifySender:       cfrex.F(true),
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

func TestAccountEmailSecuritySettingAllowPolicyDelete(t *testing.T) {
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
	_, err := client.Accounts.EmailSecurity.Settings.AllowPolicies.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		int64(2401),
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
