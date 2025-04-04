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

func TestRadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.GetHTTPMethodTimeseries(context.TODO(), cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParams{
		AggInterval:       cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval15m),
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsFormatJson),
		HTTPVersion:       cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsHTTPVersion{cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsHTTPVersionHttPv1}),
		IPVersion:         cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsIPVersion{cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsIPVersionIPv4}),
		LimitPerGroup:     cfrex.F(int64(10)),
		Location:          cfrex.F([]string{"string"}),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProduct{cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
		Normalization:     cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.GetHTTPVersionTimeseries(context.TODO(), cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParams{
		AggInterval:       cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval15m),
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsFormatJson),
		HTTPMethod:        cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod{cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodGet}),
		IPVersion:         cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsIPVersion{cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsIPVersionIPv4}),
		Location:          cfrex.F([]string{"string"}),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProduct{cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
		Normalization:     cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.GetIndustryTimeseries(context.TODO(), cfrex.RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParams{
		AggInterval:       cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval15m),
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsFormatJson),
		HTTPMethod:        cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod{cfrex.RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodGet}),
		HTTPVersion:       cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPVersion{cfrex.RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPVersionHttPv1}),
		IPVersion:         cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsIPVersion{cfrex.RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsIPVersionIPv4}),
		LimitPerGroup:     cfrex.F(int64(10)),
		Location:          cfrex.F([]string{"string"}),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProduct{cfrex.RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
		Normalization:     cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.GetIPVersionTimeseries(context.TODO(), cfrex.RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParams{
		AggInterval:       cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval15m),
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsFormatJson),
		HTTPMethod:        cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod{cfrex.RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodGet}),
		HTTPVersion:       cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPVersion{cfrex.RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPVersionHttPv1}),
		Location:          cfrex.F([]string{"string"}),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProduct{cfrex.RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
		Normalization:     cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.GetManagedRulesTimeseries(context.TODO(), cfrex.RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParams{
		AggInterval:       cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval15m),
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsFormatJson),
		HTTPMethod:        cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod{cfrex.RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodGet}),
		HTTPVersion:       cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPVersion{cfrex.RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPVersionHttPv1}),
		IPVersion:         cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsIPVersion{cfrex.RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsIPVersionIPv4}),
		LimitPerGroup:     cfrex.F(int64(10)),
		Location:          cfrex.F([]string{"string"}),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProduct{cfrex.RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
		Normalization:     cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.GetMitigationProductTimeseries(context.TODO(), cfrex.RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParams{
		AggInterval:   cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval15m),
		Asn:           cfrex.F([]string{"string"}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Format:        cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsFormatJson),
		HTTPMethod:    cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod{cfrex.RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodGet}),
		HTTPVersion:   cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPVersion{cfrex.RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsIPVersion{cfrex.RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsIPVersionIPv4}),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Normalization: cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.TimeseriesGroups.GetVerticalTimeseries(context.TODO(), cfrex.RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParams{
		AggInterval:       cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval15m),
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsFormatJson),
		HTTPMethod:        cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod{cfrex.RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodGet}),
		HTTPVersion:       cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPVersion{cfrex.RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPVersionHttPv1}),
		IPVersion:         cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsIPVersion{cfrex.RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsIPVersionIPv4}),
		LimitPerGroup:     cfrex.F(int64(10)),
		Location:          cfrex.F([]string{"string"}),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProduct{cfrex.RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
		Normalization:     cfrex.F(cfrex.RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
