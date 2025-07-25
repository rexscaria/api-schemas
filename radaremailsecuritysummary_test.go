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

func TestRadarEmailSecuritySummaryGetArcWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.GetArc(context.TODO(), cfrex.RadarEmailSecuritySummaryGetArcParams{
		DateEnd:    cfrex.F([]time.Time{time.Now()}),
		DateRange:  cfrex.F([]string{"7d"}),
		DateStart:  cfrex.F([]time.Time{time.Now()}),
		Dkim:       cfrex.F([]cfrex.RadarEmailSecuritySummaryGetArcParamsDkim{cfrex.RadarEmailSecuritySummaryGetArcParamsDkimPass}),
		Dmarc:      cfrex.F([]cfrex.RadarEmailSecuritySummaryGetArcParamsDmarc{cfrex.RadarEmailSecuritySummaryGetArcParamsDmarcPass}),
		Format:     cfrex.F(cfrex.RadarEmailSecuritySummaryGetArcParamsFormatJson),
		Name:       cfrex.F([]string{"main_series"}),
		Spf:        cfrex.F([]cfrex.RadarEmailSecuritySummaryGetArcParamsSpf{cfrex.RadarEmailSecuritySummaryGetArcParamsSpfPass}),
		TlsVersion: cfrex.F([]cfrex.RadarEmailSecuritySummaryGetArcParamsTlsVersion{cfrex.RadarEmailSecuritySummaryGetArcParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummaryGetDkimWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.GetDkim(context.TODO(), cfrex.RadarEmailSecuritySummaryGetDkimParams{
		Arc:        cfrex.F([]cfrex.RadarEmailSecuritySummaryGetDkimParamsArc{cfrex.RadarEmailSecuritySummaryGetDkimParamsArcPass}),
		DateEnd:    cfrex.F([]time.Time{time.Now()}),
		DateRange:  cfrex.F([]string{"7d"}),
		DateStart:  cfrex.F([]time.Time{time.Now()}),
		Dmarc:      cfrex.F([]cfrex.RadarEmailSecuritySummaryGetDkimParamsDmarc{cfrex.RadarEmailSecuritySummaryGetDkimParamsDmarcPass}),
		Format:     cfrex.F(cfrex.RadarEmailSecuritySummaryGetDkimParamsFormatJson),
		Name:       cfrex.F([]string{"main_series"}),
		Spf:        cfrex.F([]cfrex.RadarEmailSecuritySummaryGetDkimParamsSpf{cfrex.RadarEmailSecuritySummaryGetDkimParamsSpfPass}),
		TlsVersion: cfrex.F([]cfrex.RadarEmailSecuritySummaryGetDkimParamsTlsVersion{cfrex.RadarEmailSecuritySummaryGetDkimParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummaryGetDmarcWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.GetDmarc(context.TODO(), cfrex.RadarEmailSecuritySummaryGetDmarcParams{
		Arc:        cfrex.F([]cfrex.RadarEmailSecuritySummaryGetDmarcParamsArc{cfrex.RadarEmailSecuritySummaryGetDmarcParamsArcPass}),
		DateEnd:    cfrex.F([]time.Time{time.Now()}),
		DateRange:  cfrex.F([]string{"7d"}),
		DateStart:  cfrex.F([]time.Time{time.Now()}),
		Dkim:       cfrex.F([]cfrex.RadarEmailSecuritySummaryGetDmarcParamsDkim{cfrex.RadarEmailSecuritySummaryGetDmarcParamsDkimPass}),
		Format:     cfrex.F(cfrex.RadarEmailSecuritySummaryGetDmarcParamsFormatJson),
		Name:       cfrex.F([]string{"main_series"}),
		Spf:        cfrex.F([]cfrex.RadarEmailSecuritySummaryGetDmarcParamsSpf{cfrex.RadarEmailSecuritySummaryGetDmarcParamsSpfPass}),
		TlsVersion: cfrex.F([]cfrex.RadarEmailSecuritySummaryGetDmarcParamsTlsVersion{cfrex.RadarEmailSecuritySummaryGetDmarcParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummaryGetMaliciousWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.GetMalicious(context.TODO(), cfrex.RadarEmailSecuritySummaryGetMaliciousParams{
		Arc:        cfrex.F([]cfrex.RadarEmailSecuritySummaryGetMaliciousParamsArc{cfrex.RadarEmailSecuritySummaryGetMaliciousParamsArcPass}),
		DateEnd:    cfrex.F([]time.Time{time.Now()}),
		DateRange:  cfrex.F([]string{"7d"}),
		DateStart:  cfrex.F([]time.Time{time.Now()}),
		Dkim:       cfrex.F([]cfrex.RadarEmailSecuritySummaryGetMaliciousParamsDkim{cfrex.RadarEmailSecuritySummaryGetMaliciousParamsDkimPass}),
		Dmarc:      cfrex.F([]cfrex.RadarEmailSecuritySummaryGetMaliciousParamsDmarc{cfrex.RadarEmailSecuritySummaryGetMaliciousParamsDmarcPass}),
		Format:     cfrex.F(cfrex.RadarEmailSecuritySummaryGetMaliciousParamsFormatJson),
		Name:       cfrex.F([]string{"main_series"}),
		Spf:        cfrex.F([]cfrex.RadarEmailSecuritySummaryGetMaliciousParamsSpf{cfrex.RadarEmailSecuritySummaryGetMaliciousParamsSpfPass}),
		TlsVersion: cfrex.F([]cfrex.RadarEmailSecuritySummaryGetMaliciousParamsTlsVersion{cfrex.RadarEmailSecuritySummaryGetMaliciousParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummaryGetSpamWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.GetSpam(context.TODO(), cfrex.RadarEmailSecuritySummaryGetSpamParams{
		Arc:        cfrex.F([]cfrex.RadarEmailSecuritySummaryGetSpamParamsArc{cfrex.RadarEmailSecuritySummaryGetSpamParamsArcPass}),
		DateEnd:    cfrex.F([]time.Time{time.Now()}),
		DateRange:  cfrex.F([]string{"7d"}),
		DateStart:  cfrex.F([]time.Time{time.Now()}),
		Dkim:       cfrex.F([]cfrex.RadarEmailSecuritySummaryGetSpamParamsDkim{cfrex.RadarEmailSecuritySummaryGetSpamParamsDkimPass}),
		Dmarc:      cfrex.F([]cfrex.RadarEmailSecuritySummaryGetSpamParamsDmarc{cfrex.RadarEmailSecuritySummaryGetSpamParamsDmarcPass}),
		Format:     cfrex.F(cfrex.RadarEmailSecuritySummaryGetSpamParamsFormatJson),
		Name:       cfrex.F([]string{"main_series"}),
		Spf:        cfrex.F([]cfrex.RadarEmailSecuritySummaryGetSpamParamsSpf{cfrex.RadarEmailSecuritySummaryGetSpamParamsSpfPass}),
		TlsVersion: cfrex.F([]cfrex.RadarEmailSecuritySummaryGetSpamParamsTlsVersion{cfrex.RadarEmailSecuritySummaryGetSpamParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummaryGetSpfWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.GetSpf(context.TODO(), cfrex.RadarEmailSecuritySummaryGetSpfParams{
		Arc:        cfrex.F([]cfrex.RadarEmailSecuritySummaryGetSpfParamsArc{cfrex.RadarEmailSecuritySummaryGetSpfParamsArcPass}),
		DateEnd:    cfrex.F([]time.Time{time.Now()}),
		DateRange:  cfrex.F([]string{"7d"}),
		DateStart:  cfrex.F([]time.Time{time.Now()}),
		Dkim:       cfrex.F([]cfrex.RadarEmailSecuritySummaryGetSpfParamsDkim{cfrex.RadarEmailSecuritySummaryGetSpfParamsDkimPass}),
		Dmarc:      cfrex.F([]cfrex.RadarEmailSecuritySummaryGetSpfParamsDmarc{cfrex.RadarEmailSecuritySummaryGetSpfParamsDmarcPass}),
		Format:     cfrex.F(cfrex.RadarEmailSecuritySummaryGetSpfParamsFormatJson),
		Name:       cfrex.F([]string{"main_series"}),
		TlsVersion: cfrex.F([]cfrex.RadarEmailSecuritySummaryGetSpfParamsTlsVersion{cfrex.RadarEmailSecuritySummaryGetSpfParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummaryGetSpoofWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.GetSpoof(context.TODO(), cfrex.RadarEmailSecuritySummaryGetSpoofParams{
		Arc:        cfrex.F([]cfrex.RadarEmailSecuritySummaryGetSpoofParamsArc{cfrex.RadarEmailSecuritySummaryGetSpoofParamsArcPass}),
		DateEnd:    cfrex.F([]time.Time{time.Now()}),
		DateRange:  cfrex.F([]string{"7d"}),
		DateStart:  cfrex.F([]time.Time{time.Now()}),
		Dkim:       cfrex.F([]cfrex.RadarEmailSecuritySummaryGetSpoofParamsDkim{cfrex.RadarEmailSecuritySummaryGetSpoofParamsDkimPass}),
		Dmarc:      cfrex.F([]cfrex.RadarEmailSecuritySummaryGetSpoofParamsDmarc{cfrex.RadarEmailSecuritySummaryGetSpoofParamsDmarcPass}),
		Format:     cfrex.F(cfrex.RadarEmailSecuritySummaryGetSpoofParamsFormatJson),
		Name:       cfrex.F([]string{"main_series"}),
		Spf:        cfrex.F([]cfrex.RadarEmailSecuritySummaryGetSpoofParamsSpf{cfrex.RadarEmailSecuritySummaryGetSpoofParamsSpfPass}),
		TlsVersion: cfrex.F([]cfrex.RadarEmailSecuritySummaryGetSpoofParamsTlsVersion{cfrex.RadarEmailSecuritySummaryGetSpoofParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummaryGetThreatCategoryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.GetThreatCategory(context.TODO(), cfrex.RadarEmailSecuritySummaryGetThreatCategoryParams{
		Arc:        cfrex.F([]cfrex.RadarEmailSecuritySummaryGetThreatCategoryParamsArc{cfrex.RadarEmailSecuritySummaryGetThreatCategoryParamsArcPass}),
		DateEnd:    cfrex.F([]time.Time{time.Now()}),
		DateRange:  cfrex.F([]string{"7d"}),
		DateStart:  cfrex.F([]time.Time{time.Now()}),
		Dkim:       cfrex.F([]cfrex.RadarEmailSecuritySummaryGetThreatCategoryParamsDkim{cfrex.RadarEmailSecuritySummaryGetThreatCategoryParamsDkimPass}),
		Dmarc:      cfrex.F([]cfrex.RadarEmailSecuritySummaryGetThreatCategoryParamsDmarc{cfrex.RadarEmailSecuritySummaryGetThreatCategoryParamsDmarcPass}),
		Format:     cfrex.F(cfrex.RadarEmailSecuritySummaryGetThreatCategoryParamsFormatJson),
		Name:       cfrex.F([]string{"main_series"}),
		Spf:        cfrex.F([]cfrex.RadarEmailSecuritySummaryGetThreatCategoryParamsSpf{cfrex.RadarEmailSecuritySummaryGetThreatCategoryParamsSpfPass}),
		TlsVersion: cfrex.F([]cfrex.RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersion{cfrex.RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecuritySummaryGetTlsVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summary.GetTlsVersion(context.TODO(), cfrex.RadarEmailSecuritySummaryGetTlsVersionParams{
		Arc:       cfrex.F([]cfrex.RadarEmailSecuritySummaryGetTlsVersionParamsArc{cfrex.RadarEmailSecuritySummaryGetTlsVersionParamsArcPass}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Dkim:      cfrex.F([]cfrex.RadarEmailSecuritySummaryGetTlsVersionParamsDkim{cfrex.RadarEmailSecuritySummaryGetTlsVersionParamsDkimPass}),
		Dmarc:     cfrex.F([]cfrex.RadarEmailSecuritySummaryGetTlsVersionParamsDmarc{cfrex.RadarEmailSecuritySummaryGetTlsVersionParamsDmarcPass}),
		Format:    cfrex.F(cfrex.RadarEmailSecuritySummaryGetTlsVersionParamsFormatJson),
		Name:      cfrex.F([]string{"main_series"}),
		Spf:       cfrex.F([]cfrex.RadarEmailSecuritySummaryGetTlsVersionParamsSpf{cfrex.RadarEmailSecuritySummaryGetTlsVersionParamsSpfPass}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
