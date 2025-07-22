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

func TestRadarHTTPTopLocationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Locations.List(context.TODO(), cfrex.RadarHTTPTopLocationListParams{
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPTopLocationListParamsBotClass{cfrex.RadarHTTPTopLocationListParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopLocationListParamsBrowserFamily{cfrex.RadarHTTPTopLocationListParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPTopLocationListParamsDeviceType{cfrex.RadarHTTPTopLocationListParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPTopLocationListParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTopLocationListParamsHTTPProtocol{cfrex.RadarHTTPTopLocationListParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTopLocationListParamsHTTPVersion{cfrex.RadarHTTPTopLocationListParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPTopLocationListParamsIPVersion{cfrex.RadarHTTPTopLocationListParamsIPVersionIPv4}),
		Limit:         cfrex.F(int64(5)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPTopLocationListParamsO{cfrex.RadarHTTPTopLocationListParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPTopLocationListParamsTlsVersion{cfrex.RadarHTTPTopLocationListParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTopLocationListByBotClassWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Locations.ListByBotClass(
		context.TODO(),
		cfrex.RadarHTTPTopLocationListByBotClassParamsBotClassLikelyAutomated,
		cfrex.RadarHTTPTopLocationListByBotClassParams{
			Asn:           cfrex.F([]string{"string"}),
			BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopLocationListByBotClassParamsBrowserFamily{cfrex.RadarHTTPTopLocationListByBotClassParamsBrowserFamilyChrome}),
			Continent:     cfrex.F([]string{"string"}),
			DateEnd:       cfrex.F([]time.Time{time.Now()}),
			DateRange:     cfrex.F([]string{"7d"}),
			DateStart:     cfrex.F([]time.Time{time.Now()}),
			DeviceType:    cfrex.F([]cfrex.RadarHTTPTopLocationListByBotClassParamsDeviceType{cfrex.RadarHTTPTopLocationListByBotClassParamsDeviceTypeDesktop}),
			Format:        cfrex.F(cfrex.RadarHTTPTopLocationListByBotClassParamsFormatJson),
			HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTopLocationListByBotClassParamsHTTPProtocol{cfrex.RadarHTTPTopLocationListByBotClassParamsHTTPProtocolHTTP}),
			HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTopLocationListByBotClassParamsHTTPVersion{cfrex.RadarHTTPTopLocationListByBotClassParamsHTTPVersionHttPv1}),
			IPVersion:     cfrex.F([]cfrex.RadarHTTPTopLocationListByBotClassParamsIPVersion{cfrex.RadarHTTPTopLocationListByBotClassParamsIPVersionIPv4}),
			Limit:         cfrex.F(int64(5)),
			Location:      cfrex.F([]string{"string"}),
			Name:          cfrex.F([]string{"main_series"}),
			Os:            cfrex.F([]cfrex.RadarHTTPTopLocationListByBotClassParamsO{cfrex.RadarHTTPTopLocationListByBotClassParamsOWindows}),
			TlsVersion:    cfrex.F([]cfrex.RadarHTTPTopLocationListByBotClassParamsTlsVersion{cfrex.RadarHTTPTopLocationListByBotClassParamsTlsVersionTlSv1_0}),
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

func TestRadarHTTPTopLocationListByBrowserFamilyWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Locations.ListByBrowserFamily(
		context.TODO(),
		cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamilyChrome,
		cfrex.RadarHTTPTopLocationListByBrowserFamilyParams{
			Asn:          cfrex.F([]string{"string"}),
			BotClass:     cfrex.F([]cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsBotClass{cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsBotClassLikelyAutomated}),
			Continent:    cfrex.F([]string{"string"}),
			DateEnd:      cfrex.F([]time.Time{time.Now()}),
			DateRange:    cfrex.F([]string{"7d"}),
			DateStart:    cfrex.F([]time.Time{time.Now()}),
			DeviceType:   cfrex.F([]cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsDeviceType{cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsDeviceTypeDesktop}),
			Format:       cfrex.F(cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsFormatJson),
			HTTPProtocol: cfrex.F([]cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsHTTPProtocol{cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsHTTPProtocolHTTP}),
			HTTPVersion:  cfrex.F([]cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsHTTPVersion{cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsHTTPVersionHttPv1}),
			IPVersion:    cfrex.F([]cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsIPVersion{cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsIPVersionIPv4}),
			Limit:        cfrex.F(int64(5)),
			Location:     cfrex.F([]string{"string"}),
			Name:         cfrex.F([]string{"main_series"}),
			Os:           cfrex.F([]cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsO{cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsOWindows}),
			TlsVersion:   cfrex.F([]cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersion{cfrex.RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersionTlSv1_0}),
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

func TestRadarHTTPTopLocationListByDeviceTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Locations.ListByDeviceType(
		context.TODO(),
		cfrex.RadarHTTPTopLocationListByDeviceTypeParamsDeviceTypeDesktop,
		cfrex.RadarHTTPTopLocationListByDeviceTypeParams{
			Asn:           cfrex.F([]string{"string"}),
			BotClass:      cfrex.F([]cfrex.RadarHTTPTopLocationListByDeviceTypeParamsBotClass{cfrex.RadarHTTPTopLocationListByDeviceTypeParamsBotClassLikelyAutomated}),
			BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamily{cfrex.RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamilyChrome}),
			Continent:     cfrex.F([]string{"string"}),
			DateEnd:       cfrex.F([]time.Time{time.Now()}),
			DateRange:     cfrex.F([]string{"7d"}),
			DateStart:     cfrex.F([]time.Time{time.Now()}),
			Format:        cfrex.F(cfrex.RadarHTTPTopLocationListByDeviceTypeParamsFormatJson),
			HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTopLocationListByDeviceTypeParamsHTTPProtocol{cfrex.RadarHTTPTopLocationListByDeviceTypeParamsHTTPProtocolHTTP}),
			HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTopLocationListByDeviceTypeParamsHTTPVersion{cfrex.RadarHTTPTopLocationListByDeviceTypeParamsHTTPVersionHttPv1}),
			IPVersion:     cfrex.F([]cfrex.RadarHTTPTopLocationListByDeviceTypeParamsIPVersion{cfrex.RadarHTTPTopLocationListByDeviceTypeParamsIPVersionIPv4}),
			Limit:         cfrex.F(int64(5)),
			Location:      cfrex.F([]string{"string"}),
			Name:          cfrex.F([]string{"main_series"}),
			Os:            cfrex.F([]cfrex.RadarHTTPTopLocationListByDeviceTypeParamsO{cfrex.RadarHTTPTopLocationListByDeviceTypeParamsOWindows}),
			TlsVersion:    cfrex.F([]cfrex.RadarHTTPTopLocationListByDeviceTypeParamsTlsVersion{cfrex.RadarHTTPTopLocationListByDeviceTypeParamsTlsVersionTlSv1_0}),
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

func TestRadarHTTPTopLocationListByHTTPProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Locations.ListByHTTPProtocol(
		context.TODO(),
		cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsHTTPProtocolHTTP,
		cfrex.RadarHTTPTopLocationListByHTTPProtocolParams{
			Asn:           cfrex.F([]string{"string"}),
			BotClass:      cfrex.F([]cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsBotClass{cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsBotClassLikelyAutomated}),
			BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamily{cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamilyChrome}),
			Continent:     cfrex.F([]string{"string"}),
			DateEnd:       cfrex.F([]time.Time{time.Now()}),
			DateRange:     cfrex.F([]string{"7d"}),
			DateStart:     cfrex.F([]time.Time{time.Now()}),
			DeviceType:    cfrex.F([]cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsDeviceType{cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsDeviceTypeDesktop}),
			Format:        cfrex.F(cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsFormatJson),
			HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsHTTPVersion{cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsHTTPVersionHttPv1}),
			IPVersion:     cfrex.F([]cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsIPVersion{cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsIPVersionIPv4}),
			Limit:         cfrex.F(int64(5)),
			Location:      cfrex.F([]string{"string"}),
			Name:          cfrex.F([]string{"main_series"}),
			Os:            cfrex.F([]cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsO{cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsOWindows}),
			TlsVersion:    cfrex.F([]cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersion{cfrex.RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersionTlSv1_0}),
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

func TestRadarHTTPTopLocationListByHTTPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Locations.ListByHTTPVersion(
		context.TODO(),
		cfrex.RadarHTTPTopLocationListByHTTPVersionParamsHTTPVersionHttPv1,
		cfrex.RadarHTTPTopLocationListByHTTPVersionParams{
			Asn:           cfrex.F([]string{"string"}),
			BotClass:      cfrex.F([]cfrex.RadarHTTPTopLocationListByHTTPVersionParamsBotClass{cfrex.RadarHTTPTopLocationListByHTTPVersionParamsBotClassLikelyAutomated}),
			BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamily{cfrex.RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamilyChrome}),
			Continent:     cfrex.F([]string{"string"}),
			DateEnd:       cfrex.F([]time.Time{time.Now()}),
			DateRange:     cfrex.F([]string{"7d"}),
			DateStart:     cfrex.F([]time.Time{time.Now()}),
			DeviceType:    cfrex.F([]cfrex.RadarHTTPTopLocationListByHTTPVersionParamsDeviceType{cfrex.RadarHTTPTopLocationListByHTTPVersionParamsDeviceTypeDesktop}),
			Format:        cfrex.F(cfrex.RadarHTTPTopLocationListByHTTPVersionParamsFormatJson),
			HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTopLocationListByHTTPVersionParamsHTTPProtocol{cfrex.RadarHTTPTopLocationListByHTTPVersionParamsHTTPProtocolHTTP}),
			IPVersion:     cfrex.F([]cfrex.RadarHTTPTopLocationListByHTTPVersionParamsIPVersion{cfrex.RadarHTTPTopLocationListByHTTPVersionParamsIPVersionIPv4}),
			Limit:         cfrex.F(int64(5)),
			Location:      cfrex.F([]string{"string"}),
			Name:          cfrex.F([]string{"main_series"}),
			Os:            cfrex.F([]cfrex.RadarHTTPTopLocationListByHTTPVersionParamsO{cfrex.RadarHTTPTopLocationListByHTTPVersionParamsOWindows}),
			TlsVersion:    cfrex.F([]cfrex.RadarHTTPTopLocationListByHTTPVersionParamsTlsVersion{cfrex.RadarHTTPTopLocationListByHTTPVersionParamsTlsVersionTlSv1_0}),
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

func TestRadarHTTPTopLocationListByIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Locations.ListByIPVersion(
		context.TODO(),
		cfrex.RadarHTTPTopLocationListByIPVersionParamsIPVersionIPv4,
		cfrex.RadarHTTPTopLocationListByIPVersionParams{
			Asn:           cfrex.F([]string{"string"}),
			BotClass:      cfrex.F([]cfrex.RadarHTTPTopLocationListByIPVersionParamsBotClass{cfrex.RadarHTTPTopLocationListByIPVersionParamsBotClassLikelyAutomated}),
			BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopLocationListByIPVersionParamsBrowserFamily{cfrex.RadarHTTPTopLocationListByIPVersionParamsBrowserFamilyChrome}),
			Continent:     cfrex.F([]string{"string"}),
			DateEnd:       cfrex.F([]time.Time{time.Now()}),
			DateRange:     cfrex.F([]string{"7d"}),
			DateStart:     cfrex.F([]time.Time{time.Now()}),
			DeviceType:    cfrex.F([]cfrex.RadarHTTPTopLocationListByIPVersionParamsDeviceType{cfrex.RadarHTTPTopLocationListByIPVersionParamsDeviceTypeDesktop}),
			Format:        cfrex.F(cfrex.RadarHTTPTopLocationListByIPVersionParamsFormatJson),
			HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTopLocationListByIPVersionParamsHTTPProtocol{cfrex.RadarHTTPTopLocationListByIPVersionParamsHTTPProtocolHTTP}),
			HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTopLocationListByIPVersionParamsHTTPVersion{cfrex.RadarHTTPTopLocationListByIPVersionParamsHTTPVersionHttPv1}),
			Limit:         cfrex.F(int64(5)),
			Location:      cfrex.F([]string{"string"}),
			Name:          cfrex.F([]string{"main_series"}),
			Os:            cfrex.F([]cfrex.RadarHTTPTopLocationListByIPVersionParamsO{cfrex.RadarHTTPTopLocationListByIPVersionParamsOWindows}),
			TlsVersion:    cfrex.F([]cfrex.RadarHTTPTopLocationListByIPVersionParamsTlsVersion{cfrex.RadarHTTPTopLocationListByIPVersionParamsTlsVersionTlSv1_0}),
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

func TestRadarHTTPTopLocationListByOsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Locations.ListByOs(
		context.TODO(),
		cfrex.RadarHTTPTopLocationListByOsParamsOsWindows,
		cfrex.RadarHTTPTopLocationListByOsParams{
			Asn:           cfrex.F([]string{"string"}),
			BotClass:      cfrex.F([]cfrex.RadarHTTPTopLocationListByOsParamsBotClass{cfrex.RadarHTTPTopLocationListByOsParamsBotClassLikelyAutomated}),
			BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopLocationListByOsParamsBrowserFamily{cfrex.RadarHTTPTopLocationListByOsParamsBrowserFamilyChrome}),
			Continent:     cfrex.F([]string{"string"}),
			DateEnd:       cfrex.F([]time.Time{time.Now()}),
			DateRange:     cfrex.F([]string{"7d"}),
			DateStart:     cfrex.F([]time.Time{time.Now()}),
			DeviceType:    cfrex.F([]cfrex.RadarHTTPTopLocationListByOsParamsDeviceType{cfrex.RadarHTTPTopLocationListByOsParamsDeviceTypeDesktop}),
			Format:        cfrex.F(cfrex.RadarHTTPTopLocationListByOsParamsFormatJson),
			HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTopLocationListByOsParamsHTTPProtocol{cfrex.RadarHTTPTopLocationListByOsParamsHTTPProtocolHTTP}),
			HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTopLocationListByOsParamsHTTPVersion{cfrex.RadarHTTPTopLocationListByOsParamsHTTPVersionHttPv1}),
			IPVersion:     cfrex.F([]cfrex.RadarHTTPTopLocationListByOsParamsIPVersion{cfrex.RadarHTTPTopLocationListByOsParamsIPVersionIPv4}),
			Limit:         cfrex.F(int64(5)),
			Location:      cfrex.F([]string{"string"}),
			Name:          cfrex.F([]string{"main_series"}),
			TlsVersion:    cfrex.F([]cfrex.RadarHTTPTopLocationListByOsParamsTlsVersion{cfrex.RadarHTTPTopLocationListByOsParamsTlsVersionTlSv1_0}),
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

func TestRadarHTTPTopLocationListByTlsVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Locations.ListByTlsVersion(
		context.TODO(),
		cfrex.RadarHTTPTopLocationListByTlsVersionParamsTlsVersionTlSv1_0,
		cfrex.RadarHTTPTopLocationListByTlsVersionParams{
			Asn:           cfrex.F([]string{"string"}),
			BotClass:      cfrex.F([]cfrex.RadarHTTPTopLocationListByTlsVersionParamsBotClass{cfrex.RadarHTTPTopLocationListByTlsVersionParamsBotClassLikelyAutomated}),
			BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopLocationListByTlsVersionParamsBrowserFamily{cfrex.RadarHTTPTopLocationListByTlsVersionParamsBrowserFamilyChrome}),
			Continent:     cfrex.F([]string{"string"}),
			DateEnd:       cfrex.F([]time.Time{time.Now()}),
			DateRange:     cfrex.F([]string{"7d"}),
			DateStart:     cfrex.F([]time.Time{time.Now()}),
			DeviceType:    cfrex.F([]cfrex.RadarHTTPTopLocationListByTlsVersionParamsDeviceType{cfrex.RadarHTTPTopLocationListByTlsVersionParamsDeviceTypeDesktop}),
			Format:        cfrex.F(cfrex.RadarHTTPTopLocationListByTlsVersionParamsFormatJson),
			HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTopLocationListByTlsVersionParamsHTTPProtocol{cfrex.RadarHTTPTopLocationListByTlsVersionParamsHTTPProtocolHTTP}),
			HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTopLocationListByTlsVersionParamsHTTPVersion{cfrex.RadarHTTPTopLocationListByTlsVersionParamsHTTPVersionHttPv1}),
			IPVersion:     cfrex.F([]cfrex.RadarHTTPTopLocationListByTlsVersionParamsIPVersion{cfrex.RadarHTTPTopLocationListByTlsVersionParamsIPVersionIPv4}),
			Limit:         cfrex.F(int64(5)),
			Location:      cfrex.F([]string{"string"}),
			Name:          cfrex.F([]string{"main_series"}),
			Os:            cfrex.F([]cfrex.RadarHTTPTopLocationListByTlsVersionParamsO{cfrex.RadarHTTPTopLocationListByTlsVersionParamsOWindows}),
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
