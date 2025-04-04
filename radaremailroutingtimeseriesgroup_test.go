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

func TestRadarEmailRoutingTimeseriesGroupGetArcWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.GetArc(context.TODO(), cfrex.RadarEmailRoutingTimeseriesGroupGetArcParams{
		AggInterval: cfrex.F(cfrex.RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval15m),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dkim:        cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetArcParamsDkim{cfrex.RadarEmailRoutingTimeseriesGroupGetArcParamsDkimPass}),
		Dmarc:       cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetArcParamsDmarc{cfrex.RadarEmailRoutingTimeseriesGroupGetArcParamsDmarcPass}),
		Encrypted:   cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetArcParamsEncrypted{cfrex.RadarEmailRoutingTimeseriesGroupGetArcParamsEncryptedEncrypted}),
		Format:      cfrex.F(cfrex.RadarEmailRoutingTimeseriesGroupGetArcParamsFormatJson),
		IPVersion:   cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersion{cfrex.RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersionIPv4}),
		Name:        cfrex.F([]string{"main_series"}),
		Spf:         cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetArcParamsSpf{cfrex.RadarEmailRoutingTimeseriesGroupGetArcParamsSpfPass}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingTimeseriesGroupGetDkimWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.GetDkim(context.TODO(), cfrex.RadarEmailRoutingTimeseriesGroupGetDkimParams{
		AggInterval: cfrex.F(cfrex.RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval15m),
		Arc:         cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetDkimParamsArc{cfrex.RadarEmailRoutingTimeseriesGroupGetDkimParamsArcPass}),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dmarc:       cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarc{cfrex.RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarcPass}),
		Encrypted:   cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetDkimParamsEncrypted{cfrex.RadarEmailRoutingTimeseriesGroupGetDkimParamsEncryptedEncrypted}),
		Format:      cfrex.F(cfrex.RadarEmailRoutingTimeseriesGroupGetDkimParamsFormatJson),
		IPVersion:   cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersion{cfrex.RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersionIPv4}),
		Name:        cfrex.F([]string{"main_series"}),
		Spf:         cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetDkimParamsSpf{cfrex.RadarEmailRoutingTimeseriesGroupGetDkimParamsSpfPass}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingTimeseriesGroupGetDmarcWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.GetDmarc(context.TODO(), cfrex.RadarEmailRoutingTimeseriesGroupGetDmarcParams{
		AggInterval: cfrex.F(cfrex.RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval15m),
		Arc:         cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetDmarcParamsArc{cfrex.RadarEmailRoutingTimeseriesGroupGetDmarcParamsArcPass}),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dkim:        cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkim{cfrex.RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkimPass}),
		Encrypted:   cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncrypted{cfrex.RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncryptedEncrypted}),
		Format:      cfrex.F(cfrex.RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormatJson),
		IPVersion:   cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersion{cfrex.RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersionIPv4}),
		Name:        cfrex.F([]string{"main_series"}),
		Spf:         cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpf{cfrex.RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpfPass}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingTimeseriesGroupGetEncryptedWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.GetEncrypted(context.TODO(), cfrex.RadarEmailRoutingTimeseriesGroupGetEncryptedParams{
		AggInterval: cfrex.F(cfrex.RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval15m),
		Arc:         cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArc{cfrex.RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArcPass}),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dkim:        cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkim{cfrex.RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkimPass}),
		Dmarc:       cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarc{cfrex.RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarcPass}),
		Format:      cfrex.F(cfrex.RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormatJson),
		IPVersion:   cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersion{cfrex.RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersionIPv4}),
		Name:        cfrex.F([]string{"main_series"}),
		Spf:         cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpf{cfrex.RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpfPass}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingTimeseriesGroupGetIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.GetIPVersion(context.TODO(), cfrex.RadarEmailRoutingTimeseriesGroupGetIPVersionParams{
		AggInterval: cfrex.F(cfrex.RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval15m),
		Arc:         cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArc{cfrex.RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArcPass}),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dkim:        cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkim{cfrex.RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkimPass}),
		Dmarc:       cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarc{cfrex.RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarcPass}),
		Encrypted:   cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncrypted{cfrex.RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncryptedEncrypted}),
		Format:      cfrex.F(cfrex.RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormatJson),
		Name:        cfrex.F([]string{"main_series"}),
		Spf:         cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpf{cfrex.RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpfPass}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingTimeseriesGroupGetSpfWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.TimeseriesGroups.GetSpf(context.TODO(), cfrex.RadarEmailRoutingTimeseriesGroupGetSpfParams{
		AggInterval: cfrex.F(cfrex.RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval15m),
		Arc:         cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetSpfParamsArc{cfrex.RadarEmailRoutingTimeseriesGroupGetSpfParamsArcPass}),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dkim:        cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetSpfParamsDkim{cfrex.RadarEmailRoutingTimeseriesGroupGetSpfParamsDkimPass}),
		Dmarc:       cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarc{cfrex.RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarcPass}),
		Encrypted:   cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetSpfParamsEncrypted{cfrex.RadarEmailRoutingTimeseriesGroupGetSpfParamsEncryptedEncrypted}),
		Format:      cfrex.F(cfrex.RadarEmailRoutingTimeseriesGroupGetSpfParamsFormatJson),
		IPVersion:   cfrex.F([]cfrex.RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersion{cfrex.RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersionIPv4}),
		Name:        cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
