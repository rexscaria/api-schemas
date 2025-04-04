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

func TestRadarAs112SummaryGetDnssecWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.Summary.GetDnssec(context.TODO(), cfrex.RadarAs112SummaryGetDnssecParams{
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarAs112SummaryGetDnssecParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Protocol:     cfrex.F(cfrex.RadarAs112SummaryGetDnssecParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarAs112SummaryGetDnssecParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarAs112SummaryGetDnssecParamsResponseCodeNoerror),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAs112SummaryGetEdnsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.Summary.GetEdns(context.TODO(), cfrex.RadarAs112SummaryGetEdnsParams{
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarAs112SummaryGetEdnsParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Protocol:     cfrex.F(cfrex.RadarAs112SummaryGetEdnsParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarAs112SummaryGetEdnsParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarAs112SummaryGetEdnsParamsResponseCodeNoerror),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAs112SummaryGetIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.Summary.GetIPVersion(context.TODO(), cfrex.RadarAs112SummaryGetIPVersionParams{
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarAs112SummaryGetIPVersionParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Protocol:     cfrex.F(cfrex.RadarAs112SummaryGetIPVersionParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarAs112SummaryGetIPVersionParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarAs112SummaryGetIPVersionParamsResponseCodeNoerror),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAs112SummaryGetProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.Summary.GetProtocol(context.TODO(), cfrex.RadarAs112SummaryGetProtocolParams{
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarAs112SummaryGetProtocolParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		QueryType:    cfrex.F(cfrex.RadarAs112SummaryGetProtocolParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarAs112SummaryGetProtocolParamsResponseCodeNoerror),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAs112SummaryGetQueryTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.Summary.GetQueryType(context.TODO(), cfrex.RadarAs112SummaryGetQueryTypeParams{
		Asn:           cfrex.F([]string{"string"}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Format:        cfrex.F(cfrex.RadarAs112SummaryGetQueryTypeParamsFormatJson),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Protocol:      cfrex.F(cfrex.RadarAs112SummaryGetQueryTypeParamsProtocolUdp),
		ResponseCode:  cfrex.F(cfrex.RadarAs112SummaryGetQueryTypeParamsResponseCodeNoerror),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAs112SummaryGetResponseCodesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.Summary.GetResponseCodes(context.TODO(), cfrex.RadarAs112SummaryGetResponseCodesParams{
		Asn:           cfrex.F([]string{"string"}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Format:        cfrex.F(cfrex.RadarAs112SummaryGetResponseCodesParamsFormatJson),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Protocol:      cfrex.F(cfrex.RadarAs112SummaryGetResponseCodesParamsProtocolUdp),
		QueryType:     cfrex.F(cfrex.RadarAs112SummaryGetResponseCodesParamsQueryTypeA),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
