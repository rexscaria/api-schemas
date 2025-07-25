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

func TestRadarDNSSummaryGetCacheHitWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.GetCacheHit(context.TODO(), cfrex.RadarDNSSummaryGetCacheHitParams{
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSSummaryGetCacheHitParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		Protocol:     cfrex.F(cfrex.RadarDNSSummaryGetCacheHitParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarDNSSummaryGetCacheHitParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSSummaryGetCacheHitParamsResponseCodeNoerror),
		Tld:          cfrex.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarDNSSummaryGetDnssecWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.GetDnssec(context.TODO(), cfrex.RadarDNSSummaryGetDnssecParams{
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSSummaryGetDnssecParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		Protocol:     cfrex.F(cfrex.RadarDNSSummaryGetDnssecParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarDNSSummaryGetDnssecParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSSummaryGetDnssecParamsResponseCodeNoerror),
		Tld:          cfrex.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarDNSSummaryGetDnssecAwareWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.GetDnssecAware(context.TODO(), cfrex.RadarDNSSummaryGetDnssecAwareParams{
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSSummaryGetDnssecAwareParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		Protocol:     cfrex.F(cfrex.RadarDNSSummaryGetDnssecAwareParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarDNSSummaryGetDnssecAwareParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSSummaryGetDnssecAwareParamsResponseCodeNoerror),
		Tld:          cfrex.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarDNSSummaryGetDnssecE2EWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.GetDnssecE2E(context.TODO(), cfrex.RadarDNSSummaryGetDnssecE2EParams{
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSSummaryGetDnssecE2EParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		Protocol:     cfrex.F(cfrex.RadarDNSSummaryGetDnssecE2EParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarDNSSummaryGetDnssecE2EParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSSummaryGetDnssecE2EParamsResponseCodeNoerror),
		Tld:          cfrex.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarDNSSummaryGetIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.GetIPVersion(context.TODO(), cfrex.RadarDNSSummaryGetIPVersionParams{
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSSummaryGetIPVersionParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		Protocol:     cfrex.F(cfrex.RadarDNSSummaryGetIPVersionParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarDNSSummaryGetIPVersionParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSSummaryGetIPVersionParamsResponseCodeNoerror),
		Tld:          cfrex.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarDNSSummaryGetMatchingAnswerWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.GetMatchingAnswer(context.TODO(), cfrex.RadarDNSSummaryGetMatchingAnswerParams{
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSSummaryGetMatchingAnswerParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		Protocol:     cfrex.F(cfrex.RadarDNSSummaryGetMatchingAnswerParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarDNSSummaryGetMatchingAnswerParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSSummaryGetMatchingAnswerParamsResponseCodeNoerror),
		Tld:          cfrex.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarDNSSummaryGetProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.GetProtocol(context.TODO(), cfrex.RadarDNSSummaryGetProtocolParams{
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSSummaryGetProtocolParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		QueryType:    cfrex.F(cfrex.RadarDNSSummaryGetProtocolParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSSummaryGetProtocolParamsResponseCodeNoerror),
		Tld:          cfrex.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarDNSSummaryGetQueryTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.GetQueryType(context.TODO(), cfrex.RadarDNSSummaryGetQueryTypeParams{
		Asn:           cfrex.F([]string{"string"}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Format:        cfrex.F(cfrex.RadarDNSSummaryGetQueryTypeParamsFormatJson),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Nodata:        cfrex.F(true),
		Protocol:      cfrex.F(cfrex.RadarDNSSummaryGetQueryTypeParamsProtocolUdp),
		ResponseCode:  cfrex.F(cfrex.RadarDNSSummaryGetQueryTypeParamsResponseCodeNoerror),
		Tld:           cfrex.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarDNSSummaryGetResponseCodeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.GetResponseCode(context.TODO(), cfrex.RadarDNSSummaryGetResponseCodeParams{
		Asn:           cfrex.F([]string{"string"}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Format:        cfrex.F(cfrex.RadarDNSSummaryGetResponseCodeParamsFormatJson),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Nodata:        cfrex.F(true),
		Protocol:      cfrex.F(cfrex.RadarDNSSummaryGetResponseCodeParamsProtocolUdp),
		QueryType:     cfrex.F(cfrex.RadarDNSSummaryGetResponseCodeParamsQueryTypeA),
		Tld:           cfrex.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarDNSSummaryGetResponseTtlWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.Summary.GetResponseTtl(context.TODO(), cfrex.RadarDNSSummaryGetResponseTtlParams{
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSSummaryGetResponseTtlParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		Protocol:     cfrex.F(cfrex.RadarDNSSummaryGetResponseTtlParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarDNSSummaryGetResponseTtlParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSSummaryGetResponseTtlParamsResponseCodeNoerror),
		Tld:          cfrex.F([]string{"string"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
