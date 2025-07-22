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

func TestRadarAttackLayer7SummaryGetHTTPMethodSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.GetHTTPMethodSummary(context.TODO(), cfrex.RadarAttackLayer7SummaryGetHTTPMethodSummaryParams{
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsFormatJson),
		HTTPVersion:       cfrex.F([]cfrex.RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsHTTPVersion{cfrex.RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsHTTPVersionHttPv1}),
		IPVersion:         cfrex.F([]cfrex.RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsIPVersion{cfrex.RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsIPVersionIPv4}),
		LimitPerGroup:     cfrex.F(int64(10)),
		Location:          cfrex.F([]string{"string"}),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProduct{cfrex.RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7SummaryGetHTTPVersionSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.GetHTTPVersionSummary(context.TODO(), cfrex.RadarAttackLayer7SummaryGetHTTPVersionSummaryParams{
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsFormatJson),
		HTTPMethod:        cfrex.F([]cfrex.RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod{cfrex.RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodGet}),
		IPVersion:         cfrex.F([]cfrex.RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsIPVersion{cfrex.RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsIPVersionIPv4}),
		Location:          cfrex.F([]string{"string"}),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProduct{cfrex.RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7SummaryGetIndustrySummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.GetIndustrySummary(context.TODO(), cfrex.RadarAttackLayer7SummaryGetIndustrySummaryParams{
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7SummaryGetIndustrySummaryParamsFormatJson),
		HTTPMethod:        cfrex.F([]cfrex.RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod{cfrex.RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodGet}),
		HTTPVersion:       cfrex.F([]cfrex.RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPVersion{cfrex.RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPVersionHttPv1}),
		IPVersion:         cfrex.F([]cfrex.RadarAttackLayer7SummaryGetIndustrySummaryParamsIPVersion{cfrex.RadarAttackLayer7SummaryGetIndustrySummaryParamsIPVersionIPv4}),
		LimitPerGroup:     cfrex.F(int64(10)),
		Location:          cfrex.F([]string{"string"}),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProduct{cfrex.RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7SummaryGetIPVersionSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.GetIPVersionSummary(context.TODO(), cfrex.RadarAttackLayer7SummaryGetIPVersionSummaryParams{
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7SummaryGetIPVersionSummaryParamsFormatJson),
		HTTPMethod:        cfrex.F([]cfrex.RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod{cfrex.RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodGet}),
		HTTPVersion:       cfrex.F([]cfrex.RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPVersion{cfrex.RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPVersionHttPv1}),
		Location:          cfrex.F([]string{"string"}),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProduct{cfrex.RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7SummaryGetManagedRulesSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.GetManagedRulesSummary(context.TODO(), cfrex.RadarAttackLayer7SummaryGetManagedRulesSummaryParams{
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7SummaryGetManagedRulesSummaryParamsFormatJson),
		HTTPMethod:        cfrex.F([]cfrex.RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod{cfrex.RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodGet}),
		HTTPVersion:       cfrex.F([]cfrex.RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPVersion{cfrex.RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPVersionHttPv1}),
		IPVersion:         cfrex.F([]cfrex.RadarAttackLayer7SummaryGetManagedRulesSummaryParamsIPVersion{cfrex.RadarAttackLayer7SummaryGetManagedRulesSummaryParamsIPVersionIPv4}),
		LimitPerGroup:     cfrex.F(int64(10)),
		Location:          cfrex.F([]string{"string"}),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProduct{cfrex.RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7SummaryGetMitigationProductSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.GetMitigationProductSummary(context.TODO(), cfrex.RadarAttackLayer7SummaryGetMitigationProductSummaryParams{
		Asn:           cfrex.F([]string{"string"}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Format:        cfrex.F(cfrex.RadarAttackLayer7SummaryGetMitigationProductSummaryParamsFormatJson),
		HTTPMethod:    cfrex.F([]cfrex.RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod{cfrex.RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodGet}),
		HTTPVersion:   cfrex.F([]cfrex.RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPVersion{cfrex.RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarAttackLayer7SummaryGetMitigationProductSummaryParamsIPVersion{cfrex.RadarAttackLayer7SummaryGetMitigationProductSummaryParamsIPVersionIPv4}),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7SummaryGetVerticalSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summary.GetVerticalSummary(context.TODO(), cfrex.RadarAttackLayer7SummaryGetVerticalSummaryParams{
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7SummaryGetVerticalSummaryParamsFormatJson),
		HTTPMethod:        cfrex.F([]cfrex.RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod{cfrex.RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodGet}),
		HTTPVersion:       cfrex.F([]cfrex.RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPVersion{cfrex.RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPVersionHttPv1}),
		IPVersion:         cfrex.F([]cfrex.RadarAttackLayer7SummaryGetVerticalSummaryParamsIPVersion{cfrex.RadarAttackLayer7SummaryGetVerticalSummaryParamsIPVersionIPv4}),
		LimitPerGroup:     cfrex.F(int64(10)),
		Location:          cfrex.F([]string{"string"}),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProduct{cfrex.RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
