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

func TestRadarAttackLayer7TopGetTopAttacksWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Top.GetTopAttacks(context.TODO(), cfrex.RadarAttackLayer7TopGetTopAttacksParams{
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7TopGetTopAttacksParamsFormatJson),
		Limit:             cfrex.F(int64(5)),
		LimitDirection:    cfrex.F(cfrex.RadarAttackLayer7TopGetTopAttacksParamsLimitDirectionOrigin),
		LimitPerLocation:  cfrex.F(int64(10)),
		Location:          cfrex.F([]string{"string"}),
		Magnitude:         cfrex.F(cfrex.RadarAttackLayer7TopGetTopAttacksParamsMagnitudeAffectedZones),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7TopGetTopAttacksParamsMitigationProduct{cfrex.RadarAttackLayer7TopGetTopAttacksParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
		Normalization:     cfrex.F(cfrex.RadarAttackLayer7TopGetTopAttacksParamsNormalizationPercentage),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TopGetTopIndustryWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Top.GetTopIndustry(context.TODO(), cfrex.RadarAttackLayer7TopGetTopIndustryParams{
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7TopGetTopIndustryParamsFormatJson),
		HTTPMethod:        cfrex.F([]cfrex.RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod{cfrex.RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodGet}),
		HTTPVersion:       cfrex.F([]cfrex.RadarAttackLayer7TopGetTopIndustryParamsHTTPVersion{cfrex.RadarAttackLayer7TopGetTopIndustryParamsHTTPVersionHttPv1}),
		IPVersion:         cfrex.F([]cfrex.RadarAttackLayer7TopGetTopIndustryParamsIPVersion{cfrex.RadarAttackLayer7TopGetTopIndustryParamsIPVersionIPv4}),
		Limit:             cfrex.F(int64(5)),
		Location:          cfrex.F([]string{"string"}),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7TopGetTopIndustryParamsMitigationProduct{cfrex.RadarAttackLayer7TopGetTopIndustryParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarAttackLayer7TopGetTopVerticalsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Top.GetTopVerticals(context.TODO(), cfrex.RadarAttackLayer7TopGetTopVerticalsParams{
		Asn:               cfrex.F([]string{"string"}),
		Continent:         cfrex.F([]string{"string"}),
		DateEnd:           cfrex.F([]time.Time{time.Now()}),
		DateRange:         cfrex.F([]string{"7d"}),
		DateStart:         cfrex.F([]time.Time{time.Now()}),
		Format:            cfrex.F(cfrex.RadarAttackLayer7TopGetTopVerticalsParamsFormatJson),
		HTTPMethod:        cfrex.F([]cfrex.RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod{cfrex.RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodGet}),
		HTTPVersion:       cfrex.F([]cfrex.RadarAttackLayer7TopGetTopVerticalsParamsHTTPVersion{cfrex.RadarAttackLayer7TopGetTopVerticalsParamsHTTPVersionHttPv1}),
		IPVersion:         cfrex.F([]cfrex.RadarAttackLayer7TopGetTopVerticalsParamsIPVersion{cfrex.RadarAttackLayer7TopGetTopVerticalsParamsIPVersionIPv4}),
		Limit:             cfrex.F(int64(5)),
		Location:          cfrex.F([]string{"string"}),
		MitigationProduct: cfrex.F([]cfrex.RadarAttackLayer7TopGetTopVerticalsParamsMitigationProduct{cfrex.RadarAttackLayer7TopGetTopVerticalsParamsMitigationProductDdos}),
		Name:              cfrex.F([]string{"main_series"}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
