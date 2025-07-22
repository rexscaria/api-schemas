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

func TestRadarHTTPTopAseListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Ases.List(context.TODO(), cfrex.RadarHTTPTopAseListParams{
		Asn:           cfrex.F([]string{"string"}),
		BotClass:      cfrex.F([]cfrex.RadarHTTPTopAseListParamsBotClass{cfrex.RadarHTTPTopAseListParamsBotClassLikelyAutomated}),
		BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopAseListParamsBrowserFamily{cfrex.RadarHTTPTopAseListParamsBrowserFamilyChrome}),
		Continent:     cfrex.F([]string{"string"}),
		DateEnd:       cfrex.F([]time.Time{time.Now()}),
		DateRange:     cfrex.F([]string{"7d"}),
		DateStart:     cfrex.F([]time.Time{time.Now()}),
		DeviceType:    cfrex.F([]cfrex.RadarHTTPTopAseListParamsDeviceType{cfrex.RadarHTTPTopAseListParamsDeviceTypeDesktop}),
		Format:        cfrex.F(cfrex.RadarHTTPTopAseListParamsFormatJson),
		HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTopAseListParamsHTTPProtocol{cfrex.RadarHTTPTopAseListParamsHTTPProtocolHTTP}),
		HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTopAseListParamsHTTPVersion{cfrex.RadarHTTPTopAseListParamsHTTPVersionHttPv1}),
		IPVersion:     cfrex.F([]cfrex.RadarHTTPTopAseListParamsIPVersion{cfrex.RadarHTTPTopAseListParamsIPVersionIPv4}),
		Limit:         cfrex.F(int64(5)),
		Location:      cfrex.F([]string{"string"}),
		Name:          cfrex.F([]string{"main_series"}),
		Os:            cfrex.F([]cfrex.RadarHTTPTopAseListParamsO{cfrex.RadarHTTPTopAseListParamsOWindows}),
		TlsVersion:    cfrex.F([]cfrex.RadarHTTPTopAseListParamsTlsVersion{cfrex.RadarHTTPTopAseListParamsTlsVersionTlSv1_0}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarHTTPTopAseListByBotClassWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Ases.ListByBotClass(
		context.TODO(),
		cfrex.RadarHTTPTopAseListByBotClassParamsBotClassLikelyAutomated,
		cfrex.RadarHTTPTopAseListByBotClassParams{
			Asn:           cfrex.F([]string{"string"}),
			BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopAseListByBotClassParamsBrowserFamily{cfrex.RadarHTTPTopAseListByBotClassParamsBrowserFamilyChrome}),
			Continent:     cfrex.F([]string{"string"}),
			DateEnd:       cfrex.F([]time.Time{time.Now()}),
			DateRange:     cfrex.F([]string{"7d"}),
			DateStart:     cfrex.F([]time.Time{time.Now()}),
			DeviceType:    cfrex.F([]cfrex.RadarHTTPTopAseListByBotClassParamsDeviceType{cfrex.RadarHTTPTopAseListByBotClassParamsDeviceTypeDesktop}),
			Format:        cfrex.F(cfrex.RadarHTTPTopAseListByBotClassParamsFormatJson),
			HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTopAseListByBotClassParamsHTTPProtocol{cfrex.RadarHTTPTopAseListByBotClassParamsHTTPProtocolHTTP}),
			HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTopAseListByBotClassParamsHTTPVersion{cfrex.RadarHTTPTopAseListByBotClassParamsHTTPVersionHttPv1}),
			IPVersion:     cfrex.F([]cfrex.RadarHTTPTopAseListByBotClassParamsIPVersion{cfrex.RadarHTTPTopAseListByBotClassParamsIPVersionIPv4}),
			Limit:         cfrex.F(int64(5)),
			Location:      cfrex.F([]string{"string"}),
			Name:          cfrex.F([]string{"main_series"}),
			Os:            cfrex.F([]cfrex.RadarHTTPTopAseListByBotClassParamsO{cfrex.RadarHTTPTopAseListByBotClassParamsOWindows}),
			TlsVersion:    cfrex.F([]cfrex.RadarHTTPTopAseListByBotClassParamsTlsVersion{cfrex.RadarHTTPTopAseListByBotClassParamsTlsVersionTlSv1_0}),
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

func TestRadarHTTPTopAseListByBrowserFamilyWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Ases.ListByBrowserFamily(
		context.TODO(),
		cfrex.RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamilyChrome,
		cfrex.RadarHTTPTopAseListByBrowserFamilyParams{
			Asn:          cfrex.F([]string{"string"}),
			BotClass:     cfrex.F([]cfrex.RadarHTTPTopAseListByBrowserFamilyParamsBotClass{cfrex.RadarHTTPTopAseListByBrowserFamilyParamsBotClassLikelyAutomated}),
			Continent:    cfrex.F([]string{"string"}),
			DateEnd:      cfrex.F([]time.Time{time.Now()}),
			DateRange:    cfrex.F([]string{"7d"}),
			DateStart:    cfrex.F([]time.Time{time.Now()}),
			DeviceType:   cfrex.F([]cfrex.RadarHTTPTopAseListByBrowserFamilyParamsDeviceType{cfrex.RadarHTTPTopAseListByBrowserFamilyParamsDeviceTypeDesktop}),
			Format:       cfrex.F(cfrex.RadarHTTPTopAseListByBrowserFamilyParamsFormatJson),
			HTTPProtocol: cfrex.F([]cfrex.RadarHTTPTopAseListByBrowserFamilyParamsHTTPProtocol{cfrex.RadarHTTPTopAseListByBrowserFamilyParamsHTTPProtocolHTTP}),
			HTTPVersion:  cfrex.F([]cfrex.RadarHTTPTopAseListByBrowserFamilyParamsHTTPVersion{cfrex.RadarHTTPTopAseListByBrowserFamilyParamsHTTPVersionHttPv1}),
			IPVersion:    cfrex.F([]cfrex.RadarHTTPTopAseListByBrowserFamilyParamsIPVersion{cfrex.RadarHTTPTopAseListByBrowserFamilyParamsIPVersionIPv4}),
			Limit:        cfrex.F(int64(5)),
			Location:     cfrex.F([]string{"string"}),
			Name:         cfrex.F([]string{"main_series"}),
			Os:           cfrex.F([]cfrex.RadarHTTPTopAseListByBrowserFamilyParamsO{cfrex.RadarHTTPTopAseListByBrowserFamilyParamsOWindows}),
			TlsVersion:   cfrex.F([]cfrex.RadarHTTPTopAseListByBrowserFamilyParamsTlsVersion{cfrex.RadarHTTPTopAseListByBrowserFamilyParamsTlsVersionTlSv1_0}),
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

func TestRadarHTTPTopAseListByDeviceTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Ases.ListByDeviceType(
		context.TODO(),
		cfrex.RadarHTTPTopAseListByDeviceTypeParamsDeviceTypeDesktop,
		cfrex.RadarHTTPTopAseListByDeviceTypeParams{
			Asn:           cfrex.F([]string{"string"}),
			BotClass:      cfrex.F([]cfrex.RadarHTTPTopAseListByDeviceTypeParamsBotClass{cfrex.RadarHTTPTopAseListByDeviceTypeParamsBotClassLikelyAutomated}),
			BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopAseListByDeviceTypeParamsBrowserFamily{cfrex.RadarHTTPTopAseListByDeviceTypeParamsBrowserFamilyChrome}),
			Continent:     cfrex.F([]string{"string"}),
			DateEnd:       cfrex.F([]time.Time{time.Now()}),
			DateRange:     cfrex.F([]string{"7d"}),
			DateStart:     cfrex.F([]time.Time{time.Now()}),
			Format:        cfrex.F(cfrex.RadarHTTPTopAseListByDeviceTypeParamsFormatJson),
			HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTopAseListByDeviceTypeParamsHTTPProtocol{cfrex.RadarHTTPTopAseListByDeviceTypeParamsHTTPProtocolHTTP}),
			HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTopAseListByDeviceTypeParamsHTTPVersion{cfrex.RadarHTTPTopAseListByDeviceTypeParamsHTTPVersionHttPv1}),
			IPVersion:     cfrex.F([]cfrex.RadarHTTPTopAseListByDeviceTypeParamsIPVersion{cfrex.RadarHTTPTopAseListByDeviceTypeParamsIPVersionIPv4}),
			Limit:         cfrex.F(int64(5)),
			Location:      cfrex.F([]string{"string"}),
			Name:          cfrex.F([]string{"main_series"}),
			Os:            cfrex.F([]cfrex.RadarHTTPTopAseListByDeviceTypeParamsO{cfrex.RadarHTTPTopAseListByDeviceTypeParamsOWindows}),
			TlsVersion:    cfrex.F([]cfrex.RadarHTTPTopAseListByDeviceTypeParamsTlsVersion{cfrex.RadarHTTPTopAseListByDeviceTypeParamsTlsVersionTlSv1_0}),
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

func TestRadarHTTPTopAseListByHTTPProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Ases.ListByHTTPProtocol(
		context.TODO(),
		cfrex.RadarHTTPTopAseListByHTTPProtocolParamsHTTPProtocolHTTP,
		cfrex.RadarHTTPTopAseListByHTTPProtocolParams{
			Asn:           cfrex.F([]string{"string"}),
			BotClass:      cfrex.F([]cfrex.RadarHTTPTopAseListByHTTPProtocolParamsBotClass{cfrex.RadarHTTPTopAseListByHTTPProtocolParamsBotClassLikelyAutomated}),
			BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamily{cfrex.RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamilyChrome}),
			Continent:     cfrex.F([]string{"string"}),
			DateEnd:       cfrex.F([]time.Time{time.Now()}),
			DateRange:     cfrex.F([]string{"7d"}),
			DateStart:     cfrex.F([]time.Time{time.Now()}),
			DeviceType:    cfrex.F([]cfrex.RadarHTTPTopAseListByHTTPProtocolParamsDeviceType{cfrex.RadarHTTPTopAseListByHTTPProtocolParamsDeviceTypeDesktop}),
			Format:        cfrex.F(cfrex.RadarHTTPTopAseListByHTTPProtocolParamsFormatJson),
			HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTopAseListByHTTPProtocolParamsHTTPVersion{cfrex.RadarHTTPTopAseListByHTTPProtocolParamsHTTPVersionHttPv1}),
			IPVersion:     cfrex.F([]cfrex.RadarHTTPTopAseListByHTTPProtocolParamsIPVersion{cfrex.RadarHTTPTopAseListByHTTPProtocolParamsIPVersionIPv4}),
			Limit:         cfrex.F(int64(5)),
			Location:      cfrex.F([]string{"string"}),
			Name:          cfrex.F([]string{"main_series"}),
			Os:            cfrex.F([]cfrex.RadarHTTPTopAseListByHTTPProtocolParamsO{cfrex.RadarHTTPTopAseListByHTTPProtocolParamsOWindows}),
			TlsVersion:    cfrex.F([]cfrex.RadarHTTPTopAseListByHTTPProtocolParamsTlsVersion{cfrex.RadarHTTPTopAseListByHTTPProtocolParamsTlsVersionTlSv1_0}),
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

func TestRadarHTTPTopAseListByHTTPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Ases.ListByHTTPVersion(
		context.TODO(),
		cfrex.RadarHTTPTopAseListByHTTPVersionParamsHTTPVersionHttPv1,
		cfrex.RadarHTTPTopAseListByHTTPVersionParams{
			Asn:           cfrex.F([]string{"string"}),
			BotClass:      cfrex.F([]cfrex.RadarHTTPTopAseListByHTTPVersionParamsBotClass{cfrex.RadarHTTPTopAseListByHTTPVersionParamsBotClassLikelyAutomated}),
			BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopAseListByHTTPVersionParamsBrowserFamily{cfrex.RadarHTTPTopAseListByHTTPVersionParamsBrowserFamilyChrome}),
			Continent:     cfrex.F([]string{"string"}),
			DateEnd:       cfrex.F([]time.Time{time.Now()}),
			DateRange:     cfrex.F([]string{"7d"}),
			DateStart:     cfrex.F([]time.Time{time.Now()}),
			DeviceType:    cfrex.F([]cfrex.RadarHTTPTopAseListByHTTPVersionParamsDeviceType{cfrex.RadarHTTPTopAseListByHTTPVersionParamsDeviceTypeDesktop}),
			Format:        cfrex.F(cfrex.RadarHTTPTopAseListByHTTPVersionParamsFormatJson),
			HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTopAseListByHTTPVersionParamsHTTPProtocol{cfrex.RadarHTTPTopAseListByHTTPVersionParamsHTTPProtocolHTTP}),
			IPVersion:     cfrex.F([]cfrex.RadarHTTPTopAseListByHTTPVersionParamsIPVersion{cfrex.RadarHTTPTopAseListByHTTPVersionParamsIPVersionIPv4}),
			Limit:         cfrex.F(int64(5)),
			Location:      cfrex.F([]string{"string"}),
			Name:          cfrex.F([]string{"main_series"}),
			Os:            cfrex.F([]cfrex.RadarHTTPTopAseListByHTTPVersionParamsO{cfrex.RadarHTTPTopAseListByHTTPVersionParamsOWindows}),
			TlsVersion:    cfrex.F([]cfrex.RadarHTTPTopAseListByHTTPVersionParamsTlsVersion{cfrex.RadarHTTPTopAseListByHTTPVersionParamsTlsVersionTlSv1_0}),
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

func TestRadarHTTPTopAseListByIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Ases.ListByIPVersion(
		context.TODO(),
		cfrex.RadarHTTPTopAseListByIPVersionParamsIPVersionIPv4,
		cfrex.RadarHTTPTopAseListByIPVersionParams{
			Asn:           cfrex.F([]string{"string"}),
			BotClass:      cfrex.F([]cfrex.RadarHTTPTopAseListByIPVersionParamsBotClass{cfrex.RadarHTTPTopAseListByIPVersionParamsBotClassLikelyAutomated}),
			BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopAseListByIPVersionParamsBrowserFamily{cfrex.RadarHTTPTopAseListByIPVersionParamsBrowserFamilyChrome}),
			Continent:     cfrex.F([]string{"string"}),
			DateEnd:       cfrex.F([]time.Time{time.Now()}),
			DateRange:     cfrex.F([]string{"7d"}),
			DateStart:     cfrex.F([]time.Time{time.Now()}),
			DeviceType:    cfrex.F([]cfrex.RadarHTTPTopAseListByIPVersionParamsDeviceType{cfrex.RadarHTTPTopAseListByIPVersionParamsDeviceTypeDesktop}),
			Format:        cfrex.F(cfrex.RadarHTTPTopAseListByIPVersionParamsFormatJson),
			HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTopAseListByIPVersionParamsHTTPProtocol{cfrex.RadarHTTPTopAseListByIPVersionParamsHTTPProtocolHTTP}),
			HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTopAseListByIPVersionParamsHTTPVersion{cfrex.RadarHTTPTopAseListByIPVersionParamsHTTPVersionHttPv1}),
			Limit:         cfrex.F(int64(5)),
			Location:      cfrex.F([]string{"string"}),
			Name:          cfrex.F([]string{"main_series"}),
			Os:            cfrex.F([]cfrex.RadarHTTPTopAseListByIPVersionParamsO{cfrex.RadarHTTPTopAseListByIPVersionParamsOWindows}),
			TlsVersion:    cfrex.F([]cfrex.RadarHTTPTopAseListByIPVersionParamsTlsVersion{cfrex.RadarHTTPTopAseListByIPVersionParamsTlsVersionTlSv1_0}),
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

func TestRadarHTTPTopAseListByOsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Ases.ListByOs(
		context.TODO(),
		cfrex.RadarHTTPTopAseListByOsParamsOsWindows,
		cfrex.RadarHTTPTopAseListByOsParams{
			Asn:           cfrex.F([]string{"string"}),
			BotClass:      cfrex.F([]cfrex.RadarHTTPTopAseListByOsParamsBotClass{cfrex.RadarHTTPTopAseListByOsParamsBotClassLikelyAutomated}),
			BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopAseListByOsParamsBrowserFamily{cfrex.RadarHTTPTopAseListByOsParamsBrowserFamilyChrome}),
			Continent:     cfrex.F([]string{"string"}),
			DateEnd:       cfrex.F([]time.Time{time.Now()}),
			DateRange:     cfrex.F([]string{"7d"}),
			DateStart:     cfrex.F([]time.Time{time.Now()}),
			DeviceType:    cfrex.F([]cfrex.RadarHTTPTopAseListByOsParamsDeviceType{cfrex.RadarHTTPTopAseListByOsParamsDeviceTypeDesktop}),
			Format:        cfrex.F(cfrex.RadarHTTPTopAseListByOsParamsFormatJson),
			HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTopAseListByOsParamsHTTPProtocol{cfrex.RadarHTTPTopAseListByOsParamsHTTPProtocolHTTP}),
			HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTopAseListByOsParamsHTTPVersion{cfrex.RadarHTTPTopAseListByOsParamsHTTPVersionHttPv1}),
			IPVersion:     cfrex.F([]cfrex.RadarHTTPTopAseListByOsParamsIPVersion{cfrex.RadarHTTPTopAseListByOsParamsIPVersionIPv4}),
			Limit:         cfrex.F(int64(5)),
			Location:      cfrex.F([]string{"string"}),
			Name:          cfrex.F([]string{"main_series"}),
			TlsVersion:    cfrex.F([]cfrex.RadarHTTPTopAseListByOsParamsTlsVersion{cfrex.RadarHTTPTopAseListByOsParamsTlsVersionTlSv1_0}),
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

func TestRadarHTTPTopAseListByTlsVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Ases.ListByTlsVersion(
		context.TODO(),
		cfrex.RadarHTTPTopAseListByTlsVersionParamsTlsVersionTlSv1_0,
		cfrex.RadarHTTPTopAseListByTlsVersionParams{
			Asn:           cfrex.F([]string{"string"}),
			BotClass:      cfrex.F([]cfrex.RadarHTTPTopAseListByTlsVersionParamsBotClass{cfrex.RadarHTTPTopAseListByTlsVersionParamsBotClassLikelyAutomated}),
			BrowserFamily: cfrex.F([]cfrex.RadarHTTPTopAseListByTlsVersionParamsBrowserFamily{cfrex.RadarHTTPTopAseListByTlsVersionParamsBrowserFamilyChrome}),
			Continent:     cfrex.F([]string{"string"}),
			DateEnd:       cfrex.F([]time.Time{time.Now()}),
			DateRange:     cfrex.F([]string{"7d"}),
			DateStart:     cfrex.F([]time.Time{time.Now()}),
			DeviceType:    cfrex.F([]cfrex.RadarHTTPTopAseListByTlsVersionParamsDeviceType{cfrex.RadarHTTPTopAseListByTlsVersionParamsDeviceTypeDesktop}),
			Format:        cfrex.F(cfrex.RadarHTTPTopAseListByTlsVersionParamsFormatJson),
			HTTPProtocol:  cfrex.F([]cfrex.RadarHTTPTopAseListByTlsVersionParamsHTTPProtocol{cfrex.RadarHTTPTopAseListByTlsVersionParamsHTTPProtocolHTTP}),
			HTTPVersion:   cfrex.F([]cfrex.RadarHTTPTopAseListByTlsVersionParamsHTTPVersion{cfrex.RadarHTTPTopAseListByTlsVersionParamsHTTPVersionHttPv1}),
			IPVersion:     cfrex.F([]cfrex.RadarHTTPTopAseListByTlsVersionParamsIPVersion{cfrex.RadarHTTPTopAseListByTlsVersionParamsIPVersionIPv4}),
			Limit:         cfrex.F(int64(5)),
			Location:      cfrex.F([]string{"string"}),
			Name:          cfrex.F([]string{"main_series"}),
			Os:            cfrex.F([]cfrex.RadarHTTPTopAseListByTlsVersionParamsO{cfrex.RadarHTTPTopAseListByTlsVersionParamsOWindows}),
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
