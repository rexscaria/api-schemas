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

func TestRadarHTTPSummaryGetByBotClassWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.GetByBotClass(context.TODO(), cfrex.RadarHTTPSummaryGetByBotClassParams{
		Asn:           cfrex.F([]string{"string"}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPSummaryGetByBotClassParamsBrowserFamily{cfrex.RadarHTTPSummaryGetByBotClassParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPSummaryGetByBotClassParamsDeviceType{cfrex.RadarHTTPSummaryGetByBotClassParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPSummaryGetByBotClassParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPSummaryGetByBotClassParamsHTTPProtocol{cfrex.RadarHTTPSummaryGetByBotClassParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPSummaryGetByBotClassParamsHTTPVersion{cfrex.RadarHTTPSummaryGetByBotClassParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPSummaryGetByBotClassParamsIPVersion{cfrex.RadarHTTPSummaryGetByBotClassParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPSummaryGetByBotClassParamsO{cfrex.RadarHTTPSummaryGetByBotClassParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPSummaryGetByBotClassParamsTlsVersion{cfrex.RadarHTTPSummaryGetByBotClassParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPSummaryGetByDeviceTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.GetByDeviceType(context.TODO(), cfrex.RadarHTTPSummaryGetByDeviceTypeParams{
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPSummaryGetByDeviceTypeParamsBotClass{cfrex.RadarHTTPSummaryGetByDeviceTypeParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPSummaryGetByDeviceTypeParamsBrowserFamily{cfrex.RadarHTTPSummaryGetByDeviceTypeParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Format:        cfrex.F(cfrex.RadarHTTPSummaryGetByDeviceTypeParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPSummaryGetByDeviceTypeParamsHTTPProtocol{cfrex.RadarHTTPSummaryGetByDeviceTypeParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPSummaryGetByDeviceTypeParamsHTTPVersion{cfrex.RadarHTTPSummaryGetByDeviceTypeParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPSummaryGetByDeviceTypeParamsIPVersion{cfrex.RadarHTTPSummaryGetByDeviceTypeParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPSummaryGetByDeviceTypeParamsO{cfrex.RadarHTTPSummaryGetByDeviceTypeParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPSummaryGetByDeviceTypeParamsTlsVersion{cfrex.RadarHTTPSummaryGetByDeviceTypeParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPSummaryGetByHTTPProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.GetByHTTPProtocol(context.TODO(), cfrex.RadarHTTPSummaryGetByHTTPProtocolParams{
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPSummaryGetByHTTPProtocolParamsBotClass{cfrex.RadarHTTPSummaryGetByHTTPProtocolParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPSummaryGetByHTTPProtocolParamsBrowserFamily{cfrex.RadarHTTPSummaryGetByHTTPProtocolParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPSummaryGetByHTTPProtocolParamsDeviceType{cfrex.RadarHTTPSummaryGetByHTTPProtocolParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPSummaryGetByHTTPProtocolParamsFormatJson),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPSummaryGetByHTTPProtocolParamsHTTPVersion{cfrex.RadarHTTPSummaryGetByHTTPProtocolParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPSummaryGetByHTTPProtocolParamsIPVersion{cfrex.RadarHTTPSummaryGetByHTTPProtocolParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPSummaryGetByHTTPProtocolParamsO{cfrex.RadarHTTPSummaryGetByHTTPProtocolParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPSummaryGetByHTTPProtocolParamsTlsVersion{cfrex.RadarHTTPSummaryGetByHTTPProtocolParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPSummaryGetByHTTPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.GetByHTTPVersion(context.TODO(), cfrex.RadarHTTPSummaryGetByHTTPVersionParams{
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPSummaryGetByHTTPVersionParamsBotClass{cfrex.RadarHTTPSummaryGetByHTTPVersionParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPSummaryGetByHTTPVersionParamsBrowserFamily{cfrex.RadarHTTPSummaryGetByHTTPVersionParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPSummaryGetByHTTPVersionParamsDeviceType{cfrex.RadarHTTPSummaryGetByHTTPVersionParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPSummaryGetByHTTPVersionParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPSummaryGetByHTTPVersionParamsHTTPProtocol{cfrex.RadarHTTPSummaryGetByHTTPVersionParamsHTTPProtocolHTTP}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPSummaryGetByHTTPVersionParamsIPVersion{cfrex.RadarHTTPSummaryGetByHTTPVersionParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPSummaryGetByHTTPVersionParamsO{cfrex.RadarHTTPSummaryGetByHTTPVersionParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPSummaryGetByHTTPVersionParamsTlsVersion{cfrex.RadarHTTPSummaryGetByHTTPVersionParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPSummaryGetByIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.GetByIPVersion(context.TODO(), cfrex.RadarHTTPSummaryGetByIPVersionParams{
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPSummaryGetByIPVersionParamsBotClass{cfrex.RadarHTTPSummaryGetByIPVersionParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPSummaryGetByIPVersionParamsBrowserFamily{cfrex.RadarHTTPSummaryGetByIPVersionParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPSummaryGetByIPVersionParamsDeviceType{cfrex.RadarHTTPSummaryGetByIPVersionParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPSummaryGetByIPVersionParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPSummaryGetByIPVersionParamsHTTPProtocol{cfrex.RadarHTTPSummaryGetByIPVersionParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPSummaryGetByIPVersionParamsHTTPVersion{cfrex.RadarHTTPSummaryGetByIPVersionParamsHTTPVersionHttPv1}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPSummaryGetByIPVersionParamsO{cfrex.RadarHTTPSummaryGetByIPVersionParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPSummaryGetByIPVersionParamsTlsVersion{cfrex.RadarHTTPSummaryGetByIPVersionParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPSummaryGetByOsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.GetByOs(context.TODO(), cfrex.RadarHTTPSummaryGetByOsParams{
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPSummaryGetByOsParamsBotClass{cfrex.RadarHTTPSummaryGetByOsParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPSummaryGetByOsParamsBrowserFamily{cfrex.RadarHTTPSummaryGetByOsParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPSummaryGetByOsParamsDeviceType{cfrex.RadarHTTPSummaryGetByOsParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPSummaryGetByOsParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPSummaryGetByOsParamsHTTPProtocol{cfrex.RadarHTTPSummaryGetByOsParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPSummaryGetByOsParamsHTTPVersion{cfrex.RadarHTTPSummaryGetByOsParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPSummaryGetByOsParamsIPVersion{cfrex.RadarHTTPSummaryGetByOsParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPSummaryGetByOsParamsTlsVersion{cfrex.RadarHTTPSummaryGetByOsParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPSummaryGetByPostQuantumWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.GetByPostQuantum(context.TODO(), cfrex.RadarHTTPSummaryGetByPostQuantumParams{
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPSummaryGetByPostQuantumParamsBotClass{cfrex.RadarHTTPSummaryGetByPostQuantumParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPSummaryGetByPostQuantumParamsBrowserFamily{cfrex.RadarHTTPSummaryGetByPostQuantumParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPSummaryGetByPostQuantumParamsDeviceType{cfrex.RadarHTTPSummaryGetByPostQuantumParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPSummaryGetByPostQuantumParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPSummaryGetByPostQuantumParamsHTTPProtocol{cfrex.RadarHTTPSummaryGetByPostQuantumParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPSummaryGetByPostQuantumParamsHTTPVersion{cfrex.RadarHTTPSummaryGetByPostQuantumParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPSummaryGetByPostQuantumParamsIPVersion{cfrex.RadarHTTPSummaryGetByPostQuantumParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPSummaryGetByPostQuantumParamsO{cfrex.RadarHTTPSummaryGetByPostQuantumParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPSummaryGetByPostQuantumParamsTlsVersion{cfrex.RadarHTTPSummaryGetByPostQuantumParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPSummaryGetByTlsVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.GetByTlsVersion(context.TODO(), cfrex.RadarHTTPSummaryGetByTlsVersionParams{
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPSummaryGetByTlsVersionParamsBotClass{cfrex.RadarHTTPSummaryGetByTlsVersionParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPSummaryGetByTlsVersionParamsBrowserFamily{cfrex.RadarHTTPSummaryGetByTlsVersionParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPSummaryGetByTlsVersionParamsDeviceType{cfrex.RadarHTTPSummaryGetByTlsVersionParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPSummaryGetByTlsVersionParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPSummaryGetByTlsVersionParamsHTTPProtocol{cfrex.RadarHTTPSummaryGetByTlsVersionParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPSummaryGetByTlsVersionParamsHTTPVersion{cfrex.RadarHTTPSummaryGetByTlsVersionParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPSummaryGetByTlsVersionParamsIPVersion{cfrex.RadarHTTPSummaryGetByTlsVersionParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPSummaryGetByTlsVersionParamsO{cfrex.RadarHTTPSummaryGetByTlsVersionParamsOWindows}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
