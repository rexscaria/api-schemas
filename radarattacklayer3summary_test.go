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

func TestRadarAttackLayer3SummaryGetBitrateSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.GetBitrateSummary(context.TODO(), cfrex.RadarAttackLayer3SummaryGetBitrateSummaryParams{
		Continent: cfrex.F([]string{"string"}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Direction: cfrex.F(cfrex.RadarAttackLayer3SummaryGetBitrateSummaryParamsDirectionOrigin),
		Format:    cfrex.F(cfrex.RadarAttackLayer3SummaryGetBitrateSummaryParamsFormatJson),
		IPVersion: cfrex.F([]cfrex.RadarAttackLayer3SummaryGetBitrateSummaryParamsIPVersion{cfrex.RadarAttackLayer3SummaryGetBitrateSummaryParamsIPVersionIPv4}),
		Location:  cfrex.F([]string{"string"}),
		Name:      cfrex.F([]string{"main_series"}),
		Protocol:  cfrex.F([]cfrex.RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocol{cfrex.RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3SummaryGetDurationSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.GetDurationSummary(context.TODO(), cfrex.RadarAttackLayer3SummaryGetDurationSummaryParams{
		Continent: cfrex.F([]string{"string"}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Direction: cfrex.F(cfrex.RadarAttackLayer3SummaryGetDurationSummaryParamsDirectionOrigin),
		Format:    cfrex.F(cfrex.RadarAttackLayer3SummaryGetDurationSummaryParamsFormatJson),
		IPVersion: cfrex.F([]cfrex.RadarAttackLayer3SummaryGetDurationSummaryParamsIPVersion{cfrex.RadarAttackLayer3SummaryGetDurationSummaryParamsIPVersionIPv4}),
		Location:  cfrex.F([]string{"string"}),
		Name:      cfrex.F([]string{"main_series"}),
		Protocol:  cfrex.F([]cfrex.RadarAttackLayer3SummaryGetDurationSummaryParamsProtocol{cfrex.RadarAttackLayer3SummaryGetDurationSummaryParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3SummaryGetIndustrySummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.GetIndustrySummary(context.TODO(), cfrex.RadarAttackLayer3SummaryGetIndustrySummaryParams{
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Direction:     cfrex.F(cfrex.RadarAttackLayer3SummaryGetIndustrySummaryParamsDirectionOrigin),
		Format:        cfrex.F(cfrex.RadarAttackLayer3SummaryGetIndustrySummaryParamsFormatJson),
		IPVersion:     cfrex.F([]cfrex.RadarAttackLayer3SummaryGetIndustrySummaryParamsIPVersion{cfrex.RadarAttackLayer3SummaryGetIndustrySummaryParamsIPVersionIPv4}),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Protocol:      cfrex.F([]cfrex.RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocol{cfrex.RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3SummaryGetIPVersionSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.GetIPVersionSummary(context.TODO(), cfrex.RadarAttackLayer3SummaryGetIPVersionSummaryParams{
		Continent: cfrex.F([]string{"string"}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Direction: cfrex.F(cfrex.RadarAttackLayer3SummaryGetIPVersionSummaryParamsDirectionOrigin),
		Format:    cfrex.F(cfrex.RadarAttackLayer3SummaryGetIPVersionSummaryParamsFormatJson),
		Location:  cfrex.F([]string{"string"}),
		Name:      cfrex.F([]string{"main_series"}),
		Protocol:  cfrex.F([]cfrex.RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocol{cfrex.RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3SummaryGetProtocolSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.GetProtocolSummary(context.TODO(), cfrex.RadarAttackLayer3SummaryGetProtocolSummaryParams{
		Continent: cfrex.F([]string{"string"}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Direction: cfrex.F(cfrex.RadarAttackLayer3SummaryGetProtocolSummaryParamsDirectionOrigin),
		Format:    cfrex.F(cfrex.RadarAttackLayer3SummaryGetProtocolSummaryParamsFormatJson),
		IPVersion: cfrex.F([]cfrex.RadarAttackLayer3SummaryGetProtocolSummaryParamsIPVersion{cfrex.RadarAttackLayer3SummaryGetProtocolSummaryParamsIPVersionIPv4}),
		Location:  cfrex.F([]string{"string"}),
		Name:      cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3SummaryGetVectorSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.GetVectorSummary(context.TODO(), cfrex.RadarAttackLayer3SummaryGetVectorSummaryParams{
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Direction:     cfrex.F(cfrex.RadarAttackLayer3SummaryGetVectorSummaryParamsDirectionOrigin),
		Format:        cfrex.F(cfrex.RadarAttackLayer3SummaryGetVectorSummaryParamsFormatJson),
		IPVersion:     cfrex.F([]cfrex.RadarAttackLayer3SummaryGetVectorSummaryParamsIPVersion{cfrex.RadarAttackLayer3SummaryGetVectorSummaryParamsIPVersionIPv4}),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Protocol:      cfrex.F([]cfrex.RadarAttackLayer3SummaryGetVectorSummaryParamsProtocol{cfrex.RadarAttackLayer3SummaryGetVectorSummaryParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3SummaryGetVerticalSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Summary.GetVerticalSummary(context.TODO(), cfrex.RadarAttackLayer3SummaryGetVerticalSummaryParams{
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Direction:     cfrex.F(cfrex.RadarAttackLayer3SummaryGetVerticalSummaryParamsDirectionOrigin),
		Format:        cfrex.F(cfrex.RadarAttackLayer3SummaryGetVerticalSummaryParamsFormatJson),
		IPVersion:     cfrex.F([]cfrex.RadarAttackLayer3SummaryGetVerticalSummaryParamsIPVersion{cfrex.RadarAttackLayer3SummaryGetVerticalSummaryParamsIPVersionIPv4}),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Protocol:      cfrex.F([]cfrex.RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocol{cfrex.RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
