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

func TestRadarEmailSecurityTimeseriesGroupGetArcWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.GetArc(context.TODO(), cfrex.RadarEmailSecurityTimeseriesGroupGetArcParams{
		AggInterval: cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval1h),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dkim:        cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetArcParamsDkim{cfrex.RadarEmailSecurityTimeseriesGroupGetArcParamsDkimPass}),
		Dmarc:       cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetArcParamsDmarc{cfrex.RadarEmailSecurityTimeseriesGroupGetArcParamsDmarcPass}),
		Format:      cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetArcParamsFormatJson),
		Name:        cfrex.F([]string{"main_series"}),
		Spf:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetArcParamsSpf{cfrex.RadarEmailSecurityTimeseriesGroupGetArcParamsSpfPass}),
		TlsVersion:  cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersion{cfrex.RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupGetDkimWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.GetDkim(context.TODO(), cfrex.RadarEmailSecurityTimeseriesGroupGetDkimParams{
		AggInterval: cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval1h),
		Arc:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetDkimParamsArc{cfrex.RadarEmailSecurityTimeseriesGroupGetDkimParamsArcPass}),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dmarc:       cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetDkimParamsDmarc{cfrex.RadarEmailSecurityTimeseriesGroupGetDkimParamsDmarcPass}),
		Format:      cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetDkimParamsFormatJson),
		Name:        cfrex.F([]string{"main_series"}),
		Spf:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetDkimParamsSpf{cfrex.RadarEmailSecurityTimeseriesGroupGetDkimParamsSpfPass}),
		TlsVersion:  cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersion{cfrex.RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupGetDmarcWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.GetDmarc(context.TODO(), cfrex.RadarEmailSecurityTimeseriesGroupGetDmarcParams{
		AggInterval: cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval1h),
		Arc:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetDmarcParamsArc{cfrex.RadarEmailSecurityTimeseriesGroupGetDmarcParamsArcPass}),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dkim:        cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetDmarcParamsDkim{cfrex.RadarEmailSecurityTimeseriesGroupGetDmarcParamsDkimPass}),
		Format:      cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetDmarcParamsFormatJson),
		Name:        cfrex.F([]string{"main_series"}),
		Spf:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetDmarcParamsSpf{cfrex.RadarEmailSecurityTimeseriesGroupGetDmarcParamsSpfPass}),
		TlsVersion:  cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersion{cfrex.RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupGetMaliciousWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.GetMalicious(context.TODO(), cfrex.RadarEmailSecurityTimeseriesGroupGetMaliciousParams{
		AggInterval: cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval1h),
		Arc:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetMaliciousParamsArc{cfrex.RadarEmailSecurityTimeseriesGroupGetMaliciousParamsArcPass}),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dkim:        cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDkim{cfrex.RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDkimPass}),
		Dmarc:       cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDmarc{cfrex.RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDmarcPass}),
		Format:      cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetMaliciousParamsFormatJson),
		Name:        cfrex.F([]string{"main_series"}),
		Spf:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetMaliciousParamsSpf{cfrex.RadarEmailSecurityTimeseriesGroupGetMaliciousParamsSpfPass}),
		TlsVersion:  cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersion{cfrex.RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupGetSpamWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.GetSpam(context.TODO(), cfrex.RadarEmailSecurityTimeseriesGroupGetSpamParams{
		AggInterval: cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval1h),
		Arc:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetSpamParamsArc{cfrex.RadarEmailSecurityTimeseriesGroupGetSpamParamsArcPass}),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dkim:        cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetSpamParamsDkim{cfrex.RadarEmailSecurityTimeseriesGroupGetSpamParamsDkimPass}),
		Dmarc:       cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetSpamParamsDmarc{cfrex.RadarEmailSecurityTimeseriesGroupGetSpamParamsDmarcPass}),
		Format:      cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetSpamParamsFormatJson),
		Name:        cfrex.F([]string{"main_series"}),
		Spf:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetSpamParamsSpf{cfrex.RadarEmailSecurityTimeseriesGroupGetSpamParamsSpfPass}),
		TlsVersion:  cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersion{cfrex.RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupGetSpfWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.GetSpf(context.TODO(), cfrex.RadarEmailSecurityTimeseriesGroupGetSpfParams{
		AggInterval: cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval1h),
		Arc:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetSpfParamsArc{cfrex.RadarEmailSecurityTimeseriesGroupGetSpfParamsArcPass}),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dkim:        cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetSpfParamsDkim{cfrex.RadarEmailSecurityTimeseriesGroupGetSpfParamsDkimPass}),
		Dmarc:       cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetSpfParamsDmarc{cfrex.RadarEmailSecurityTimeseriesGroupGetSpfParamsDmarcPass}),
		Format:      cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetSpfParamsFormatJson),
		Name:        cfrex.F([]string{"main_series"}),
		TlsVersion:  cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersion{cfrex.RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupGetSpoofWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.GetSpoof(context.TODO(), cfrex.RadarEmailSecurityTimeseriesGroupGetSpoofParams{
		AggInterval: cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval1h),
		Arc:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetSpoofParamsArc{cfrex.RadarEmailSecurityTimeseriesGroupGetSpoofParamsArcPass}),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dkim:        cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetSpoofParamsDkim{cfrex.RadarEmailSecurityTimeseriesGroupGetSpoofParamsDkimPass}),
		Dmarc:       cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetSpoofParamsDmarc{cfrex.RadarEmailSecurityTimeseriesGroupGetSpoofParamsDmarcPass}),
		Format:      cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetSpoofParamsFormatJson),
		Name:        cfrex.F([]string{"main_series"}),
		Spf:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetSpoofParamsSpf{cfrex.RadarEmailSecurityTimeseriesGroupGetSpoofParamsSpfPass}),
		TlsVersion:  cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersion{cfrex.RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupGetThreatCategoryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.GetThreatCategory(context.TODO(), cfrex.RadarEmailSecurityTimeseriesGroupGetThreatCategoryParams{
		AggInterval: cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval1h),
		Arc:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsArc{cfrex.RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsArcPass}),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dkim:        cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDkim{cfrex.RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDkimPass}),
		Dmarc:       cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDmarc{cfrex.RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDmarcPass}),
		Format:      cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsFormatJson),
		Name:        cfrex.F([]string{"main_series"}),
		Spf:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsSpf{cfrex.RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsSpfPass}),
		TlsVersion:  cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersion{cfrex.RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTimeseriesGroupGetTlsVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.TimeseriesGroups.GetTlsVersion(context.TODO(), cfrex.RadarEmailSecurityTimeseriesGroupGetTlsVersionParams{
		AggInterval: cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval1h),
		Arc:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsArc{cfrex.RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsArcPass}),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dkim:        cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDkim{cfrex.RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDkimPass}),
		Dmarc:       cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDmarc{cfrex.RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDmarcPass}),
		Format:      cfrex.F(cfrex.RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsFormatJson),
		Name:        cfrex.F([]string{"main_series"}),
		Spf:         cfrex.F([]cfrex.RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsSpf{cfrex.RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsSpfPass}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
