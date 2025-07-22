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

func TestAccountEmailSecuritySettingDomainGet(t *testing.T) {
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
	_, err := client.Accounts.EmailSecurity.Settings.Domains.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		int64(2400),
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountEmailSecuritySettingDomainUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.EmailSecurity.Settings.Domains.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		int64(2400),
		cfrex.AccountEmailSecuritySettingDomainUpdateParams{
			IPRestrictions:     cfrex.F([]string{"192.0.2.0/24", "2001:db8::/32"}),
			Domain:             cfrex.F("domain"),
			DropDispositions:   cfrex.F([]cfrex.DispositionLabel{cfrex.DispositionLabelMalicious}),
			Folder:             cfrex.F(cfrex.ScannableFolderAllItems),
			IntegrationID:      cfrex.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LookbackHops:       cfrex.F(int64(1)),
			RequireTlsInbound:  cfrex.F(true),
			RequireTlsOutbound: cfrex.F(true),
			Transport:          cfrex.F("transport"),
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

func TestAccountEmailSecuritySettingDomainListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.EmailSecurity.Settings.Domains.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountEmailSecuritySettingDomainListParams{
			ActiveDeliveryMode:  cfrex.F(cfrex.DeliveryModeDirect),
			AllowedDeliveryMode: cfrex.F(cfrex.DeliveryModeDirect),
			Direction:           cfrex.F(cfrex.SortingDirectionAsc),
			Domain:              cfrex.F([]string{"string"}),
			Order:               cfrex.F(cfrex.AccountEmailSecuritySettingDomainListParamsOrderDomain),
			Page:                cfrex.F(int64(1)),
			PerPage:             cfrex.F(int64(1)),
			Search:              cfrex.F("search"),
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

func TestAccountEmailSecuritySettingDomainUnprotect(t *testing.T) {
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
	_, err := client.Accounts.EmailSecurity.Settings.Domains.Unprotect(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		int64(2400),
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountEmailSecuritySettingDomainUnprotectMultiple(t *testing.T) {
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
	_, err := client.Accounts.EmailSecurity.Settings.Domains.UnprotectMultiple(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountEmailSecuritySettingDomainUnprotectMultipleParams{
			Body: []cfrex.AccountEmailSecuritySettingDomainUnprotectMultipleParamsBody{{
				ID: cfrex.F(int64(2400)),
			}},
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
