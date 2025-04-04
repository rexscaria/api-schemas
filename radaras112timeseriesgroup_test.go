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

func TestRadarAs112TimeseriesGroupGetDnssecWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.TimeseriesGroups.GetDnssec(context.TODO(), cfrex.RadarAs112TimeseriesGroupGetDnssecParams{
		AggInterval:  cfrex.F(cfrex.RadarAs112TimeseriesGroupGetDnssecParamsAggInterval15m),
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarAs112TimeseriesGroupGetDnssecParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Protocol:     cfrex.F(cfrex.RadarAs112TimeseriesGroupGetDnssecParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeNoerror),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAs112TimeseriesGroupGetEdnsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.TimeseriesGroups.GetEdns(context.TODO(), cfrex.RadarAs112TimeseriesGroupGetEdnsParams{
		AggInterval:  cfrex.F(cfrex.RadarAs112TimeseriesGroupGetEdnsParamsAggInterval15m),
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarAs112TimeseriesGroupGetEdnsParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Protocol:     cfrex.F(cfrex.RadarAs112TimeseriesGroupGetEdnsParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeNoerror),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAs112TimeseriesGroupGetIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.TimeseriesGroups.GetIPVersion(context.TODO(), cfrex.RadarAs112TimeseriesGroupGetIPVersionParams{
		AggInterval:  cfrex.F(cfrex.RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval15m),
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarAs112TimeseriesGroupGetIPVersionParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		Protocol:     cfrex.F(cfrex.RadarAs112TimeseriesGroupGetIPVersionParamsProtocolUdp),
		QueryType:    cfrex.F(cfrex.RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeNoerror),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAs112TimeseriesGroupGetProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.TimeseriesGroups.GetProtocol(context.TODO(), cfrex.RadarAs112TimeseriesGroupGetProtocolParams{
		AggInterval:  cfrex.F(cfrex.RadarAs112TimeseriesGroupGetProtocolParamsAggInterval15m),
		Asn:          cfrex.F([]string{"string"}),
		Continent:    cfrex.F([]string{"string"}),
		DateEnd:      cfrex.F([]time.Time{time.Now()}),
		DateRange:    cfrex.F([]string{"7d"}),
		DateStart:    cfrex.F([]time.Time{time.Now()}),
		Format:       cfrex.F(cfrex.RadarAs112TimeseriesGroupGetProtocolParamsFormatJson),
		Location:     cfrex.F([]string{"string"}),
		Name:         cfrex.F([]string{"main_series"}),
		QueryType:    cfrex.F(cfrex.RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeA),
		ResponseCode: cfrex.F(cfrex.RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeNoerror),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAs112TimeseriesGroupGetQueryTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.TimeseriesGroups.GetQueryType(context.TODO(), cfrex.RadarAs112TimeseriesGroupGetQueryTypeParams{
		AggInterval:   cfrex.F(cfrex.RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval15m),
		Asn:           cfrex.F([]string{"string"}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Format:        cfrex.F(cfrex.RadarAs112TimeseriesGroupGetQueryTypeParamsFormatJson),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Protocol:      cfrex.F(cfrex.RadarAs112TimeseriesGroupGetQueryTypeParamsProtocolUdp),
		ResponseCode:  cfrex.F(cfrex.RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeNoerror),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAs112TimeseriesGroupGetResponseCodesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.As112.TimeseriesGroups.GetResponseCodes(context.TODO(), cfrex.RadarAs112TimeseriesGroupGetResponseCodesParams{
		AggInterval:   cfrex.F(cfrex.RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval15m),
		Asn:           cfrex.F([]string{"string"}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Format:        cfrex.F(cfrex.RadarAs112TimeseriesGroupGetResponseCodesParamsFormatJson),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Protocol:      cfrex.F(cfrex.RadarAs112TimeseriesGroupGetResponseCodesParamsProtocolUdp),
		QueryType:     cfrex.F(cfrex.RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeA),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
