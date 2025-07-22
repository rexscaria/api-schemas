// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/rexscaria/api-schemas"
	"github.com/rexscaria/api-schemas/internal/testutil"
	"github.com/rexscaria/api-schemas/option"
)

func TestZoneWaitingRoomNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneWaitingRoomNewParams{
			QueryWaitingroom: cfrex.QueryWaitingroomParam{
				Host:              cfrex.F("shop.example.com"),
				Name:              cfrex.F("production_webinar"),
				NewUsersPerMinute: cfrex.F(int64(200)),
				TotalActiveUsers:  cfrex.F(int64(200)),
				AdditionalRoutes: cfrex.F([]cfrex.AdditionalRouteParam{{
					Host: cfrex.F("shop2.example.com"),
					Path: cfrex.F("/shop2/checkout"),
				}}),
				CookieAttributes: cfrex.F(cfrex.CookieAttributesParam{
					Samesite: cfrex.F(cfrex.CookieAttributesSamesiteAuto),
					Secure:   cfrex.F(cfrex.CookieAttributesSecureAuto),
				}),
				CookieSuffix:            cfrex.F("abcd"),
				CustomPageHTML:          cfrex.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
				DefaultTemplateLanguage: cfrex.F(cfrex.DefaultTemplateLanguageEsEs),
				Description:             cfrex.F("Production - DO NOT MODIFY"),
				DisableSessionRenewal:   cfrex.F(false),
				EnabledOriginCommands:   cfrex.F([]cfrex.EnabledOriginCommand{cfrex.EnabledOriginCommandRevoke}),
				JsonResponseEnabled:     cfrex.F(false),
				Path:                    cfrex.F("/shop/checkout"),
				QueueAll:                cfrex.F(true),
				QueueingMethod:          cfrex.F(cfrex.QueueingMethodFifo),
				QueueingStatusCode:      cfrex.F(cfrex.QueueingStatusCode202),
				SessionDuration:         cfrex.F(int64(1)),
				Suspended:               cfrex.F(true),
				TurnstileAction:         cfrex.F(cfrex.TurnstileDetectionActionLog),
				TurnstileMode:           cfrex.F(cfrex.TurnstileWidgetModeOff),
			},
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

func TestZoneWaitingRoomGet(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneWaitingRoomUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		cfrex.ZoneWaitingRoomUpdateParams{
			QueryWaitingroom: cfrex.QueryWaitingroomParam{
				Host:              cfrex.F("shop.example.com"),
				Name:              cfrex.F("production_webinar"),
				NewUsersPerMinute: cfrex.F(int64(200)),
				TotalActiveUsers:  cfrex.F(int64(200)),
				AdditionalRoutes: cfrex.F([]cfrex.AdditionalRouteParam{{
					Host: cfrex.F("shop2.example.com"),
					Path: cfrex.F("/shop2/checkout"),
				}}),
				CookieAttributes: cfrex.F(cfrex.CookieAttributesParam{
					Samesite: cfrex.F(cfrex.CookieAttributesSamesiteAuto),
					Secure:   cfrex.F(cfrex.CookieAttributesSecureAuto),
				}),
				CookieSuffix:            cfrex.F("abcd"),
				CustomPageHTML:          cfrex.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
				DefaultTemplateLanguage: cfrex.F(cfrex.DefaultTemplateLanguageEsEs),
				Description:             cfrex.F("Production - DO NOT MODIFY"),
				DisableSessionRenewal:   cfrex.F(false),
				EnabledOriginCommands:   cfrex.F([]cfrex.EnabledOriginCommand{cfrex.EnabledOriginCommandRevoke}),
				JsonResponseEnabled:     cfrex.F(false),
				Path:                    cfrex.F("/shop/checkout"),
				QueueAll:                cfrex.F(true),
				QueueingMethod:          cfrex.F(cfrex.QueueingMethodFifo),
				QueueingStatusCode:      cfrex.F(cfrex.QueueingStatusCode202),
				SessionDuration:         cfrex.F(int64(1)),
				Suspended:               cfrex.F(true),
				TurnstileAction:         cfrex.F(cfrex.TurnstileDetectionActionLog),
				TurnstileMode:           cfrex.F(cfrex.TurnstileWidgetModeOff),
			},
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

func TestZoneWaitingRoomListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneWaitingRoomListParams{
			Page:    cfrex.F(1.000000),
			PerPage: cfrex.F(5.000000),
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

func TestZoneWaitingRoomDelete(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		cfrex.ZoneWaitingRoomDeleteParams{
			Body: map[string]interface{}{},
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

func TestZoneWaitingRoomPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Patch(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		cfrex.ZoneWaitingRoomPatchParams{
			QueryWaitingroom: cfrex.QueryWaitingroomParam{
				Host:              cfrex.F("shop.example.com"),
				Name:              cfrex.F("production_webinar"),
				NewUsersPerMinute: cfrex.F(int64(200)),
				TotalActiveUsers:  cfrex.F(int64(200)),
				AdditionalRoutes: cfrex.F([]cfrex.AdditionalRouteParam{{
					Host: cfrex.F("shop2.example.com"),
					Path: cfrex.F("/shop2/checkout"),
				}}),
				CookieAttributes: cfrex.F(cfrex.CookieAttributesParam{
					Samesite: cfrex.F(cfrex.CookieAttributesSamesiteAuto),
					Secure:   cfrex.F(cfrex.CookieAttributesSecureAuto),
				}),
				CookieSuffix:            cfrex.F("abcd"),
				CustomPageHTML:          cfrex.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
				DefaultTemplateLanguage: cfrex.F(cfrex.DefaultTemplateLanguageEsEs),
				Description:             cfrex.F("Production - DO NOT MODIFY"),
				DisableSessionRenewal:   cfrex.F(false),
				EnabledOriginCommands:   cfrex.F([]cfrex.EnabledOriginCommand{cfrex.EnabledOriginCommandRevoke}),
				JsonResponseEnabled:     cfrex.F(false),
				Path:                    cfrex.F("/shop/checkout"),
				QueueAll:                cfrex.F(true),
				QueueingMethod:          cfrex.F(cfrex.QueueingMethodFifo),
				QueueingStatusCode:      cfrex.F(cfrex.QueueingStatusCode202),
				SessionDuration:         cfrex.F(int64(1)),
				Suspended:               cfrex.F(true),
				TurnstileAction:         cfrex.F(cfrex.TurnstileDetectionActionLog),
				TurnstileMode:           cfrex.F(cfrex.TurnstileWidgetModeOff),
			},
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

func TestZoneWaitingRoomPreview(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Preview(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneWaitingRoomPreviewParams{
			CustomHTML: cfrex.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
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

func TestZoneWaitingRoomStatus(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Status(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
