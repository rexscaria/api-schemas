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

func TestRadarEmailSecurityTopTldGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Top.Tlds.Get(context.TODO(), cfrex.RadarEmailSecurityTopTldGetParams{
		Arc:         cfrex.F([]cfrex.RadarEmailSecurityTopTldGetParamsArc{cfrex.RadarEmailSecurityTopTldGetParamsArcPass}),
		DateEnd:     cfrex.F([]time.Time{time.Now()}),
		DateRange:   cfrex.F([]string{"7d"}),
		DateStart:   cfrex.F([]time.Time{time.Now()}),
		Dkim:        cfrex.F([]cfrex.RadarEmailSecurityTopTldGetParamsDkim{cfrex.RadarEmailSecurityTopTldGetParamsDkimPass}),
		Dmarc:       cfrex.F([]cfrex.RadarEmailSecurityTopTldGetParamsDmarc{cfrex.RadarEmailSecurityTopTldGetParamsDmarcPass}),
		Format:      cfrex.F(cfrex.RadarEmailSecurityTopTldGetParamsFormatJson),
		Limit:       cfrex.F(int64(5)),
		Name:        cfrex.F([]string{"main_series"}),
		Spf:         cfrex.F([]cfrex.RadarEmailSecurityTopTldGetParamsSpf{cfrex.RadarEmailSecurityTopTldGetParamsSpfPass}),
		TldCategory: cfrex.F(cfrex.RadarEmailSecurityTopTldGetParamsTldCategoryClassic),
		TlsVersion:  cfrex.F([]cfrex.RadarEmailSecurityTopTldGetParamsTlsVersion{cfrex.RadarEmailSecurityTopTldGetParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarEmailSecurityTopTldGetMaliciousWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Top.Tlds.GetMalicious(
		context.TODO(),
		cfrex.RadarEmailSecurityTopTldGetMaliciousParamsMaliciousMalicious,
		cfrex.RadarEmailSecurityTopTldGetMaliciousParams{
			Arc:         cfrex.F([]cfrex.RadarEmailSecurityTopTldGetMaliciousParamsArc{cfrex.RadarEmailSecurityTopTldGetMaliciousParamsArcPass}),
			DateEnd:     cfrex.F([]time.Time{time.Now()}),
			DateRange:   cfrex.F([]string{"7d"}),
			DateStart:   cfrex.F([]time.Time{time.Now()}),
			Dkim:        cfrex.F([]cfrex.RadarEmailSecurityTopTldGetMaliciousParamsDkim{cfrex.RadarEmailSecurityTopTldGetMaliciousParamsDkimPass}),
			Dmarc:       cfrex.F([]cfrex.RadarEmailSecurityTopTldGetMaliciousParamsDmarc{cfrex.RadarEmailSecurityTopTldGetMaliciousParamsDmarcPass}),
			Format:      cfrex.F(cfrex.RadarEmailSecurityTopTldGetMaliciousParamsFormatJson),
			Limit:       cfrex.F(int64(5)),
			Name:        cfrex.F([]string{"main_series"}),
			Spf:         cfrex.F([]cfrex.RadarEmailSecurityTopTldGetMaliciousParamsSpf{cfrex.RadarEmailSecurityTopTldGetMaliciousParamsSpfPass}),
			TldCategory: cfrex.F(cfrex.RadarEmailSecurityTopTldGetMaliciousParamsTldCategoryClassic),
			TlsVersion:  cfrex.F([]cfrex.RadarEmailSecurityTopTldGetMaliciousParamsTlsVersion{cfrex.RadarEmailSecurityTopTldGetMaliciousParamsTlsVersionTlSv1_0}),
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

func TestRadarEmailSecurityTopTldGetSpamWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Top.Tlds.GetSpam(
		context.TODO(),
		cfrex.RadarEmailSecurityTopTldGetSpamParamsSpamSpam,
		cfrex.RadarEmailSecurityTopTldGetSpamParams{
			Arc:         cfrex.F([]cfrex.RadarEmailSecurityTopTldGetSpamParamsArc{cfrex.RadarEmailSecurityTopTldGetSpamParamsArcPass}),
			DateEnd:     cfrex.F([]time.Time{time.Now()}),
			DateRange:   cfrex.F([]string{"7d"}),
			DateStart:   cfrex.F([]time.Time{time.Now()}),
			Dkim:        cfrex.F([]cfrex.RadarEmailSecurityTopTldGetSpamParamsDkim{cfrex.RadarEmailSecurityTopTldGetSpamParamsDkimPass}),
			Dmarc:       cfrex.F([]cfrex.RadarEmailSecurityTopTldGetSpamParamsDmarc{cfrex.RadarEmailSecurityTopTldGetSpamParamsDmarcPass}),
			Format:      cfrex.F(cfrex.RadarEmailSecurityTopTldGetSpamParamsFormatJson),
			Limit:       cfrex.F(int64(5)),
			Name:        cfrex.F([]string{"main_series"}),
			Spf:         cfrex.F([]cfrex.RadarEmailSecurityTopTldGetSpamParamsSpf{cfrex.RadarEmailSecurityTopTldGetSpamParamsSpfPass}),
			TldCategory: cfrex.F(cfrex.RadarEmailSecurityTopTldGetSpamParamsTldCategoryClassic),
			TlsVersion:  cfrex.F([]cfrex.RadarEmailSecurityTopTldGetSpamParamsTlsVersion{cfrex.RadarEmailSecurityTopTldGetSpamParamsTlsVersionTlSv1_0}),
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

func TestRadarEmailSecurityTopTldGetSpoofWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Top.Tlds.GetSpoof(
		context.TODO(),
		cfrex.RadarEmailSecurityTopTldGetSpoofParamsSpoofSpoof,
		cfrex.RadarEmailSecurityTopTldGetSpoofParams{
			Arc:         cfrex.F([]cfrex.RadarEmailSecurityTopTldGetSpoofParamsArc{cfrex.RadarEmailSecurityTopTldGetSpoofParamsArcPass}),
			DateEnd:     cfrex.F([]time.Time{time.Now()}),
			DateRange:   cfrex.F([]string{"7d"}),
			DateStart:   cfrex.F([]time.Time{time.Now()}),
			Dkim:        cfrex.F([]cfrex.RadarEmailSecurityTopTldGetSpoofParamsDkim{cfrex.RadarEmailSecurityTopTldGetSpoofParamsDkimPass}),
			Dmarc:       cfrex.F([]cfrex.RadarEmailSecurityTopTldGetSpoofParamsDmarc{cfrex.RadarEmailSecurityTopTldGetSpoofParamsDmarcPass}),
			Format:      cfrex.F(cfrex.RadarEmailSecurityTopTldGetSpoofParamsFormatJson),
			Limit:       cfrex.F(int64(5)),
			Name:        cfrex.F([]string{"main_series"}),
			Spf:         cfrex.F([]cfrex.RadarEmailSecurityTopTldGetSpoofParamsSpf{cfrex.RadarEmailSecurityTopTldGetSpoofParamsSpfPass}),
			TldCategory: cfrex.F(cfrex.RadarEmailSecurityTopTldGetSpoofParamsTldCategoryClassic),
			TlsVersion:  cfrex.F([]cfrex.RadarEmailSecurityTopTldGetSpoofParamsTlsVersion{cfrex.RadarEmailSecurityTopTldGetSpoofParamsTlsVersionTlSv1_0}),
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
