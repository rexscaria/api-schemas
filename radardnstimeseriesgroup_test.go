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

func TestRadarDNSTimeseriesGroupGetCacheHitWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.GetCacheHit(context.TODO(), cfrex.RadarDNSTimeseriesGroupGetCacheHitParams{
		AggInterval:  cfrex.F(cfrex.RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval15m),
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSTimeseriesGroupGetCacheHitParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		Protocol:     cfrex.F(cfrex.RadarDNSTimeseriesGroupGetCacheHitParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeNoerror),
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

func TestRadarDNSTimeseriesGroupGetDnssecWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.GetDnssec(context.TODO(), cfrex.RadarDNSTimeseriesGroupGetDnssecParams{
		AggInterval:  cfrex.F(cfrex.RadarDNSTimeseriesGroupGetDnssecParamsAggInterval15m),
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSTimeseriesGroupGetDnssecParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		Protocol:     cfrex.F(cfrex.RadarDNSTimeseriesGroupGetDnssecParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeNoerror),
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

func TestRadarDNSTimeseriesGroupGetDnssecAwareWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.GetDnssecAware(context.TODO(), cfrex.RadarDNSTimeseriesGroupGetDnssecAwareParams{
		AggInterval:  cfrex.F(cfrex.RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval15m),
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSTimeseriesGroupGetDnssecAwareParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		Protocol:     cfrex.F(cfrex.RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeNoerror),
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

func TestRadarDNSTimeseriesGroupGetDnssecE2EWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.GetDnssecE2E(context.TODO(), cfrex.RadarDNSTimeseriesGroupGetDnssecE2EParams{
		AggInterval:  cfrex.F(cfrex.RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval15m),
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSTimeseriesGroupGetDnssecE2EParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		Protocol:     cfrex.F(cfrex.RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeNoerror),
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

func TestRadarDNSTimeseriesGroupGetIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.GetIPVersion(context.TODO(), cfrex.RadarDNSTimeseriesGroupGetIPVersionParams{
		AggInterval:  cfrex.F(cfrex.RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval15m),
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSTimeseriesGroupGetIPVersionParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		Protocol:     cfrex.F(cfrex.RadarDNSTimeseriesGroupGetIPVersionParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeNoerror),
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

func TestRadarDNSTimeseriesGroupGetMatchingAnswerWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.GetMatchingAnswer(context.TODO(), cfrex.RadarDNSTimeseriesGroupGetMatchingAnswerParams{
		AggInterval:  cfrex.F(cfrex.RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval15m),
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSTimeseriesGroupGetMatchingAnswerParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		Protocol:     cfrex.F(cfrex.RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeNoerror),
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

func TestRadarDNSTimeseriesGroupGetProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.GetProtocol(context.TODO(), cfrex.RadarDNSTimeseriesGroupGetProtocolParams{
		AggInterval:  cfrex.F(cfrex.RadarDNSTimeseriesGroupGetProtocolParamsAggInterval15m),
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSTimeseriesGroupGetProtocolParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		QueryType:    cfrex.F(cfrex.RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeNoerror),
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

func TestRadarDNSTimeseriesGroupGetQueryTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.GetQueryType(context.TODO(), cfrex.RadarDNSTimeseriesGroupGetQueryTypeParams{
		AggInterval:   cfrex.F(cfrex.RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval15m),
		Asn:           cfrex.F([]string{"string"}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Format:        cfrex.F(cfrex.RadarDNSTimeseriesGroupGetQueryTypeParamsFormatJson),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Nodata:        cfrex.F(true),
		Protocol:      cfrex.F(cfrex.RadarDNSTimeseriesGroupGetQueryTypeParamsProtocolUdp),
		ResponseCode:  cfrex.F(cfrex.RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeNoerror),
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

func TestRadarDNSTimeseriesGroupGetResponseCodeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.GetResponseCode(context.TODO(), cfrex.RadarDNSTimeseriesGroupGetResponseCodeParams{
		AggInterval:   cfrex.F(cfrex.RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval15m),
		Asn:           cfrex.F([]string{"string"}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Format:        cfrex.F(cfrex.RadarDNSTimeseriesGroupGetResponseCodeParamsFormatJson),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Nodata:        cfrex.F(true),
		Protocol:      cfrex.F(cfrex.RadarDNSTimeseriesGroupGetResponseCodeParamsProtocolUdp),
		QueryType:     cfrex.F(cfrex.RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeA),
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

func TestRadarDNSTimeseriesGroupGetResponseTtlWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.DNS.TimeseriesGroups.GetResponseTtl(context.TODO(), cfrex.RadarDNSTimeseriesGroupGetResponseTtlParams{
		AggInterval:  cfrex.F(cfrex.RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval15m),
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarDNSTimeseriesGroupGetResponseTtlParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Nodata:       cfrex.F(true),
		Protocol:     cfrex.F(cfrex.RadarDNSTimeseriesGroupGetResponseTtlParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeNoerror),
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
