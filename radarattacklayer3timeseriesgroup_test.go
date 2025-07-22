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

func TestRadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.GetBitrateTimeseries(context.TODO(), cfrex.RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParams{
		AggInterval:   cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval1h),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Direction:     cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsDirectionOrigin),
		Format:        cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsFormatJson),
		IPVersion:     cfrex.F([]cfrex.RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsIPVersion{cfrex.RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Normalization: cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsNormalizationPercentage),
		Protocol:      cfrex.F([]cfrex.RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocol{cfrex.RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3TimeseriesGroupGetDurationTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.GetDurationTimeseries(context.TODO(), cfrex.RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParams{
		AggInterval:   cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval1h),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Direction:     cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsDirectionOrigin),
		Format:        cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsFormatJson),
		IPVersion:     cfrex.F([]cfrex.RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsIPVersion{cfrex.RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Normalization: cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsNormalizationPercentage),
		Protocol:      cfrex.F([]cfrex.RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocol{cfrex.RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.GetIndustryTimeseries(context.TODO(), cfrex.RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParams{
		AggInterval:   cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval1h),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Direction:     cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsDirectionOrigin),
		Format:        cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsFormatJson),
		IPVersion:     cfrex.F([]cfrex.RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsIPVersion{cfrex.RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsIPVersionIPv4}),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Normalization: cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsNormalizationPercentage),
		Protocol:      cfrex.F([]cfrex.RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocol{cfrex.RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.GetIPVersionTimeseries(context.TODO(), cfrex.RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParams{
		AggInterval:   cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval1h),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Direction:     cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsDirectionOrigin),
		Format:        cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsFormatJson),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Normalization: cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsNormalizationPercentage),
		Protocol:      cfrex.F([]cfrex.RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocol{cfrex.RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.GetProtocolTimeseries(context.TODO(), cfrex.RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParams{
		AggInterval:   cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval1h),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Direction:     cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsDirectionOrigin),
		Format:        cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsFormatJson),
		IPVersion:     cfrex.F([]cfrex.RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsIPVersion{cfrex.RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Normalization: cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3TimeseriesGroupGetVectorTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.GetVectorTimeseries(context.TODO(), cfrex.RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParams{
		AggInterval:   cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval1h),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Direction:     cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsDirectionOrigin),
		Format:        cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsFormatJson),
		IPVersion:     cfrex.F([]cfrex.RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsIPVersion{cfrex.RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsIPVersionIPv4}),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Normalization: cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsNormalizationPercentage),
		Protocol:      cfrex.F([]cfrex.RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocol{cfrex.RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.TimeseriesGroups.GetVerticalTimeseries(context.TODO(), cfrex.RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParams{
		AggInterval:   cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval1h),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Direction:     cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsDirectionOrigin),
		Format:        cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsFormatJson),
		IPVersion:     cfrex.F([]cfrex.RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsIPVersion{cfrex.RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsIPVersionIPv4}),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Normalization: cfrex.F(cfrex.RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsNormalizationPercentage),
		Protocol:      cfrex.F([]cfrex.RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocol{cfrex.RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
