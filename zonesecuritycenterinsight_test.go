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

func TestZoneSecurityCenterInsightGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.SecurityCenter.Insights.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneSecurityCenterInsightGetParams{
			Dismissed:     cfrex.F(false),
			IssueClass:    cfrex.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueClassNeq: cfrex.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueType:     cfrex.F([]cfrex.IssueType{cfrex.IssueTypeComplianceViolation, cfrex.IssueTypeEmailSecurity}),
			IssueTypeNeq:  cfrex.F([]cfrex.IssueType{cfrex.IssueTypeComplianceViolation, cfrex.IssueTypeEmailSecurity}),
			Page:          cfrex.F(int64(1)),
			PerPage:       cfrex.F(int64(25)),
			Product:       cfrex.F([]string{"access", "dns"}),
			ProductNeq:    cfrex.F([]string{"access", "dns"}),
			Severity:      cfrex.F([]cfrex.SeverityQueryParam{cfrex.SeverityQueryParamLow, cfrex.SeverityQueryParamModerate}),
			SeverityNeq:   cfrex.F([]cfrex.SeverityQueryParam{cfrex.SeverityQueryParamLow, cfrex.SeverityQueryParamModerate}),
			Subject:       cfrex.F([]string{"example.com"}),
			SubjectNeq:    cfrex.F([]string{"example.com"}),
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

func TestZoneSecurityCenterInsightClassCountsWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.SecurityCenter.Insights.ClassCounts(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneSecurityCenterInsightClassCountsParams{
			Dismissed:     cfrex.F(false),
			IssueClass:    cfrex.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueClassNeq: cfrex.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueType:     cfrex.F([]cfrex.IssueType{cfrex.IssueTypeComplianceViolation, cfrex.IssueTypeEmailSecurity}),
			IssueTypeNeq:  cfrex.F([]cfrex.IssueType{cfrex.IssueTypeComplianceViolation, cfrex.IssueTypeEmailSecurity}),
			Product:       cfrex.F([]string{"access", "dns"}),
			ProductNeq:    cfrex.F([]string{"access", "dns"}),
			Severity:      cfrex.F([]cfrex.SeverityQueryParam{cfrex.SeverityQueryParamLow, cfrex.SeverityQueryParamModerate}),
			SeverityNeq:   cfrex.F([]cfrex.SeverityQueryParam{cfrex.SeverityQueryParamLow, cfrex.SeverityQueryParamModerate}),
			Subject:       cfrex.F([]string{"example.com"}),
			SubjectNeq:    cfrex.F([]string{"example.com"}),
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

func TestZoneSecurityCenterInsightDismissWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.SecurityCenter.Insights.Dismiss(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"issue_id",
		cfrex.ZoneSecurityCenterInsightDismissParams{
			Dismiss: cfrex.F(true),
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

func TestZoneSecurityCenterInsightSeverityCountsWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.SecurityCenter.Insights.SeverityCounts(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneSecurityCenterInsightSeverityCountsParams{
			Dismissed:     cfrex.F(false),
			IssueClass:    cfrex.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueClassNeq: cfrex.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueType:     cfrex.F([]cfrex.IssueType{cfrex.IssueTypeComplianceViolation, cfrex.IssueTypeEmailSecurity}),
			IssueTypeNeq:  cfrex.F([]cfrex.IssueType{cfrex.IssueTypeComplianceViolation, cfrex.IssueTypeEmailSecurity}),
			Product:       cfrex.F([]string{"access", "dns"}),
			ProductNeq:    cfrex.F([]string{"access", "dns"}),
			Severity:      cfrex.F([]cfrex.SeverityQueryParam{cfrex.SeverityQueryParamLow, cfrex.SeverityQueryParamModerate}),
			SeverityNeq:   cfrex.F([]cfrex.SeverityQueryParam{cfrex.SeverityQueryParamLow, cfrex.SeverityQueryParamModerate}),
			Subject:       cfrex.F([]string{"example.com"}),
			SubjectNeq:    cfrex.F([]string{"example.com"}),
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

func TestZoneSecurityCenterInsightTypeCountsWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.SecurityCenter.Insights.TypeCounts(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneSecurityCenterInsightTypeCountsParams{
			Dismissed:     cfrex.F(false),
			IssueClass:    cfrex.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueClassNeq: cfrex.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueType:     cfrex.F([]cfrex.IssueType{cfrex.IssueTypeComplianceViolation, cfrex.IssueTypeEmailSecurity}),
			IssueTypeNeq:  cfrex.F([]cfrex.IssueType{cfrex.IssueTypeComplianceViolation, cfrex.IssueTypeEmailSecurity}),
			Product:       cfrex.F([]string{"access", "dns"}),
			ProductNeq:    cfrex.F([]string{"access", "dns"}),
			Severity:      cfrex.F([]cfrex.SeverityQueryParam{cfrex.SeverityQueryParamLow, cfrex.SeverityQueryParamModerate}),
			SeverityNeq:   cfrex.F([]cfrex.SeverityQueryParam{cfrex.SeverityQueryParamLow, cfrex.SeverityQueryParamModerate}),
			Subject:       cfrex.F([]string{"example.com"}),
			SubjectNeq:    cfrex.F([]string{"example.com"}),
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
