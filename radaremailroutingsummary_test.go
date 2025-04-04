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

func TestRadarEmailRoutingSummaryGetArcWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.GetArc(context.TODO(), cfrex.RadarEmailRoutingSummaryGetArcParams{
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Dkim:      cfrex.F([]cfrex.RadarEmailRoutingSummaryGetArcParamsDkim{cfrex.RadarEmailRoutingSummaryGetArcParamsDkimPass}),
		Dmarc:     cfrex.F([]cfrex.RadarEmailRoutingSummaryGetArcParamsDmarc{cfrex.RadarEmailRoutingSummaryGetArcParamsDmarcPass}),
		Encrypted: cfrex.F([]cfrex.RadarEmailRoutingSummaryGetArcParamsEncrypted{cfrex.RadarEmailRoutingSummaryGetArcParamsEncryptedEncrypted}),
		Format:    cfrex.F(cfrex.RadarEmailRoutingSummaryGetArcParamsFormatJson),
		IPVersion: cfrex.F([]cfrex.RadarEmailRoutingSummaryGetArcParamsIPVersion{cfrex.RadarEmailRoutingSummaryGetArcParamsIPVersionIPv4}),
		Name:      cfrex.F([]string{"main_series"}),
		Spf:       cfrex.F([]cfrex.RadarEmailRoutingSummaryGetArcParamsSpf{cfrex.RadarEmailRoutingSummaryGetArcParamsSpfPass}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingSummaryGetDkimWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.GetDkim(context.TODO(), cfrex.RadarEmailRoutingSummaryGetDkimParams{
		Arc:       cfrex.F([]cfrex.RadarEmailRoutingSummaryGetDkimParamsArc{cfrex.RadarEmailRoutingSummaryGetDkimParamsArcPass}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Dmarc:     cfrex.F([]cfrex.RadarEmailRoutingSummaryGetDkimParamsDmarc{cfrex.RadarEmailRoutingSummaryGetDkimParamsDmarcPass}),
		Encrypted: cfrex.F([]cfrex.RadarEmailRoutingSummaryGetDkimParamsEncrypted{cfrex.RadarEmailRoutingSummaryGetDkimParamsEncryptedEncrypted}),
		Format:    cfrex.F(cfrex.RadarEmailRoutingSummaryGetDkimParamsFormatJson),
		IPVersion: cfrex.F([]cfrex.RadarEmailRoutingSummaryGetDkimParamsIPVersion{cfrex.RadarEmailRoutingSummaryGetDkimParamsIPVersionIPv4}),
		Name:      cfrex.F([]string{"main_series"}),
		Spf:       cfrex.F([]cfrex.RadarEmailRoutingSummaryGetDkimParamsSpf{cfrex.RadarEmailRoutingSummaryGetDkimParamsSpfPass}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingSummaryGetDmarcWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.GetDmarc(context.TODO(), cfrex.RadarEmailRoutingSummaryGetDmarcParams{
		Arc:       cfrex.F([]cfrex.RadarEmailRoutingSummaryGetDmarcParamsArc{cfrex.RadarEmailRoutingSummaryGetDmarcParamsArcPass}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Dkim:      cfrex.F([]cfrex.RadarEmailRoutingSummaryGetDmarcParamsDkim{cfrex.RadarEmailRoutingSummaryGetDmarcParamsDkimPass}),
		Encrypted: cfrex.F([]cfrex.RadarEmailRoutingSummaryGetDmarcParamsEncrypted{cfrex.RadarEmailRoutingSummaryGetDmarcParamsEncryptedEncrypted}),
		Format:    cfrex.F(cfrex.RadarEmailRoutingSummaryGetDmarcParamsFormatJson),
		IPVersion: cfrex.F([]cfrex.RadarEmailRoutingSummaryGetDmarcParamsIPVersion{cfrex.RadarEmailRoutingSummaryGetDmarcParamsIPVersionIPv4}),
		Name:      cfrex.F([]string{"main_series"}),
		Spf:       cfrex.F([]cfrex.RadarEmailRoutingSummaryGetDmarcParamsSpf{cfrex.RadarEmailRoutingSummaryGetDmarcParamsSpfPass}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingSummaryGetEncryptedWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.GetEncrypted(context.TODO(), cfrex.RadarEmailRoutingSummaryGetEncryptedParams{
		Arc:       cfrex.F([]cfrex.RadarEmailRoutingSummaryGetEncryptedParamsArc{cfrex.RadarEmailRoutingSummaryGetEncryptedParamsArcPass}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Dkim:      cfrex.F([]cfrex.RadarEmailRoutingSummaryGetEncryptedParamsDkim{cfrex.RadarEmailRoutingSummaryGetEncryptedParamsDkimPass}),
		Dmarc:     cfrex.F([]cfrex.RadarEmailRoutingSummaryGetEncryptedParamsDmarc{cfrex.RadarEmailRoutingSummaryGetEncryptedParamsDmarcPass}),
		Format:    cfrex.F(cfrex.RadarEmailRoutingSummaryGetEncryptedParamsFormatJson),
		IPVersion: cfrex.F([]cfrex.RadarEmailRoutingSummaryGetEncryptedParamsIPVersion{cfrex.RadarEmailRoutingSummaryGetEncryptedParamsIPVersionIPv4}),
		Name:      cfrex.F([]string{"main_series"}),
		Spf:       cfrex.F([]cfrex.RadarEmailRoutingSummaryGetEncryptedParamsSpf{cfrex.RadarEmailRoutingSummaryGetEncryptedParamsSpfPass}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingSummaryGetIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.GetIPVersion(context.TODO(), cfrex.RadarEmailRoutingSummaryGetIPVersionParams{
		Arc:       cfrex.F([]cfrex.RadarEmailRoutingSummaryGetIPVersionParamsArc{cfrex.RadarEmailRoutingSummaryGetIPVersionParamsArcPass}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Dkim:      cfrex.F([]cfrex.RadarEmailRoutingSummaryGetIPVersionParamsDkim{cfrex.RadarEmailRoutingSummaryGetIPVersionParamsDkimPass}),
		Dmarc:     cfrex.F([]cfrex.RadarEmailRoutingSummaryGetIPVersionParamsDmarc{cfrex.RadarEmailRoutingSummaryGetIPVersionParamsDmarcPass}),
		Encrypted: cfrex.F([]cfrex.RadarEmailRoutingSummaryGetIPVersionParamsEncrypted{cfrex.RadarEmailRoutingSummaryGetIPVersionParamsEncryptedEncrypted}),
		Format:    cfrex.F(cfrex.RadarEmailRoutingSummaryGetIPVersionParamsFormatJson),
		Name:      cfrex.F([]string{"main_series"}),
		Spf:       cfrex.F([]cfrex.RadarEmailRoutingSummaryGetIPVersionParamsSpf{cfrex.RadarEmailRoutingSummaryGetIPVersionParamsSpfPass}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailRoutingSummaryGetSpfWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Routing.Summary.GetSpf(context.TODO(), cfrex.RadarEmailRoutingSummaryGetSpfParams{
		Arc:       cfrex.F([]cfrex.RadarEmailRoutingSummaryGetSpfParamsArc{cfrex.RadarEmailRoutingSummaryGetSpfParamsArcPass}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Dkim:      cfrex.F([]cfrex.RadarEmailRoutingSummaryGetSpfParamsDkim{cfrex.RadarEmailRoutingSummaryGetSpfParamsDkimPass}),
		Dmarc:     cfrex.F([]cfrex.RadarEmailRoutingSummaryGetSpfParamsDmarc{cfrex.RadarEmailRoutingSummaryGetSpfParamsDmarcPass}),
		Encrypted: cfrex.F([]cfrex.RadarEmailRoutingSummaryGetSpfParamsEncrypted{cfrex.RadarEmailRoutingSummaryGetSpfParamsEncryptedEncrypted}),
		Format:    cfrex.F(cfrex.RadarEmailRoutingSummaryGetSpfParamsFormatJson),
		IPVersion: cfrex.F([]cfrex.RadarEmailRoutingSummaryGetSpfParamsIPVersion{cfrex.RadarEmailRoutingSummaryGetSpfParamsIPVersionIPv4}),
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
