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

func TestRadarHTTPTimeseriesGroupGetByBotClassWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.GetByBotClass(context.TODO(), cfrex.RadarHTTPTimeseriesGroupGetByBotClassParams{
		AggInterval:   cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsAggInterval1h),
		Asn:           cfrex.F([]string{"string"}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamily{cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsDeviceType{cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPProtocol{cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPVersion{cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsIPVersion{cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsO{cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersion{cfrex.RadarHTTPTimeseriesGroupGetByBotClassParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupGetByBrowserWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.GetByBrowser(context.TODO(), cfrex.RadarHTTPTimeseriesGroupGetByBrowserParams{
		AggInterval:   cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsAggInterval1h),
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsBotClass{cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamily{cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsDeviceType{cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPProtocol{cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPVersion{cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsIPVersion{cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsIPVersionIPv4}),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsO{cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersion{cfrex.RadarHTTPTimeseriesGroupGetByBrowserParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupGetByBrowserFamilyWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.GetByBrowserFamily(context.TODO(), cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParams{
		AggInterval:   cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsAggInterval1h),
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsBotClass{cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsBotClassLikelyAutomated}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsDeviceType{cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPProtocol{cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPVersion{cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsIPVersion{cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsIPVersionIPv4}),
		LimitPerGroup: cfrex.F(int64(10)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsO{cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersion{cfrex.RadarHTTPTimeseriesGroupGetByBrowserFamilyParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupGetByDeviceTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.GetByDeviceType(context.TODO(), cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParams{
		AggInterval:   cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsAggInterval1h),
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBotClass{cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamily{cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		Format:        cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPProtocol{cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPVersion{cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsIPVersion{cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsO{cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersion{cfrex.RadarHTTPTimeseriesGroupGetByDeviceTypeParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupGetByHTTPProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.GetByHTTPProtocol(context.TODO(), cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParams{
		AggInterval:   cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsAggInterval1h),
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBotClass{cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamily{cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsDeviceType{cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsFormatJson),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsHTTPVersion{cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsIPVersion{cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsO{cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersion{cfrex.RadarHTTPTimeseriesGroupGetByHTTPProtocolParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupGetByHTTPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.GetByHTTPVersion(context.TODO(), cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParams{
		AggInterval:   cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsAggInterval1h),
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBotClass{cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamily{cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsDeviceType{cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsHTTPProtocol{cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsHTTPProtocolHTTP}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsIPVersion{cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsO{cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersion{cfrex.RadarHTTPTimeseriesGroupGetByHTTPVersionParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupGetByIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.GetByIPVersion(context.TODO(), cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParams{
		AggInterval:   cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsAggInterval1h),
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsBotClass{cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamily{cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsDeviceType{cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPProtocol{cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPVersion{cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsHTTPVersionHttPv1}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsO{cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersion{cfrex.RadarHTTPTimeseriesGroupGetByIPVersionParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupGetByOsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.GetByOs(context.TODO(), cfrex.RadarHTTPTimeseriesGroupGetByOsParams{
		AggInterval:   cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByOsParamsAggInterval1h),
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByOsParamsBotClass{cfrex.RadarHTTPTimeseriesGroupGetByOsParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamily{cfrex.RadarHTTPTimeseriesGroupGetByOsParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByOsParamsDeviceType{cfrex.RadarHTTPTimeseriesGroupGetByOsParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByOsParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByOsParamsHTTPProtocol{cfrex.RadarHTTPTimeseriesGroupGetByOsParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByOsParamsHTTPVersion{cfrex.RadarHTTPTimeseriesGroupGetByOsParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByOsParamsIPVersion{cfrex.RadarHTTPTimeseriesGroupGetByOsParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByOsParamsTlsVersion{cfrex.RadarHTTPTimeseriesGroupGetByOsParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupGetByPostQuantumWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.GetByPostQuantum(context.TODO(), cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParams{
		AggInterval:   cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsAggInterval1h),
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsBotClass{cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamily{cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsDeviceType{cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPProtocol{cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPVersion{cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsIPVersion{cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsO{cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersion{cfrex.RadarHTTPTimeseriesGroupGetByPostQuantumParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTimeseriesGroupGetByTlsVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.GetByTlsVersion(context.TODO(), cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParams{
		AggInterval:   cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsAggInterval1h),
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsBotClass{cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamily{cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsDeviceType{cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPProtocol{cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPVersion{cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsIPVersion{cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsIPVersionIPv4}),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsO{cfrex.RadarHTTPTimeseriesGroupGetByTlsVersionParamsOWindows}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
