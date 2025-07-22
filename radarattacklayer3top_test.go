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

func TestRadarAttackLayer3TopGetTopAttacksWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.GetTopAttacks(context.TODO(), cfrex.RadarAttackLayer3TopGetTopAttacksParams{
		Continent:        cfrex.F([]string{"string"}),
		DateEnd:          cfrex.F([]time.Time{time.Now()}),
		DateRange:        cfrex.F([]string{"7d"}),
		DateStart:        cfrex.F([]time.Time{time.Now()}),
		Format:           cfrex.F(cfrex.RadarAttackLayer3TopGetTopAttacksParamsFormatJson),
		IPVersion:        cfrex.F([]cfrex.RadarAttackLayer3TopGetTopAttacksParamsIPVersion{cfrex.RadarAttackLayer3TopGetTopAttacksParamsIPVersionIPv4}),
		Limit:            cfrex.F(int64(5)),
		LimitDirection:   cfrex.F(cfrex.RadarAttackLayer3TopGetTopAttacksParamsLimitDirectionOrigin),
		LimitPerLocation: cfrex.F(int64(10)),
		Location:         cfrex.F([]string{"string"}),
		Magnitude:        cfrex.F(cfrex.RadarAttackLayer3TopGetTopAttacksParamsMagnitudeMitigatedBytes),
		Name:             cfrex.F([]string{"main_series"}),
		Normalization:    cfrex.F(cfrex.RadarAttackLayer3TopGetTopAttacksParamsNormalizationPercentage),
		Protocol:         cfrex.F([]cfrex.RadarAttackLayer3TopGetTopAttacksParamsProtocol{cfrex.RadarAttackLayer3TopGetTopAttacksParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3TopGetTopIndustryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.GetTopIndustry(context.TODO(), cfrex.RadarAttackLayer3TopGetTopIndustryParams{
		Continent: cfrex.F([]string{"string"}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Format:    cfrex.F(cfrex.RadarAttackLayer3TopGetTopIndustryParamsFormatJson),
		IPVersion: cfrex.F([]cfrex.RadarAttackLayer3TopGetTopIndustryParamsIPVersion{cfrex.RadarAttackLayer3TopGetTopIndustryParamsIPVersionIPv4}),
		Limit:     cfrex.F(int64(5)),
		Location:  cfrex.F([]string{"string"}),
		Name:      cfrex.F([]string{"main_series"}),
		Protocol:  cfrex.F([]cfrex.RadarAttackLayer3TopGetTopIndustryParamsProtocol{cfrex.RadarAttackLayer3TopGetTopIndustryParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer3TopGetTopVerticalsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer3.Top.GetTopVerticals(context.TODO(), cfrex.RadarAttackLayer3TopGetTopVerticalsParams{
		Continent: cfrex.F([]string{"string"}),
		DateEnd:   cfrex.F([]time.Time{time.Now()}),
		DateRange: cfrex.F([]string{"7d"}),
		DateStart: cfrex.F([]time.Time{time.Now()}),
		Format:    cfrex.F(cfrex.RadarAttackLayer3TopGetTopVerticalsParamsFormatJson),
		IPVersion: cfrex.F([]cfrex.RadarAttackLayer3TopGetTopVerticalsParamsIPVersion{cfrex.RadarAttackLayer3TopGetTopVerticalsParamsIPVersionIPv4}),
		Limit:     cfrex.F(int64(5)),
		Location:  cfrex.F([]string{"string"}),
		Name:      cfrex.F([]string{"main_series"}),
		Protocol:  cfrex.F([]cfrex.RadarAttackLayer3TopGetTopVerticalsParamsProtocol{cfrex.RadarAttackLayer3TopGetTopVerticalsParamsProtocolUdp}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
