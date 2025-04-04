// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/cf-rex-go"
	"github.com/stainless-sdks/cf-rex-go/internal/testutil"
	"github.com/stainless-sdks/cf-rex-go/option"
)

func TestZoneWaitingRoomEventNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Events.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		cfrex.ZoneWaitingRoomEventNewParams{
			QueryEvent: cfrex.QueryEventParam{
				EventEndTime:          cfrex.F("2021-09-28T17:00:00.000Z"),
				EventStartTime:        cfrex.F("2021-09-28T15:30:00.000Z"),
				Name:                  cfrex.F("production_webinar_event"),
				CustomPageHTML:        cfrex.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Event is prequeueing / Queue all enabled {{/waitTimeKnown}}"),
				Description:           cfrex.F("Production event - DO NOT MODIFY"),
				DisableSessionRenewal: cfrex.F(true),
				NewUsersPerMinute:     cfrex.F(int64(200)),
				PrequeueStartTime:     cfrex.F("2021-09-28T15:00:00.000Z"),
				QueueingMethod:        cfrex.F("random"),
				SessionDuration:       cfrex.F(int64(1)),
				ShuffleAtEventStart:   cfrex.F(true),
				Suspended:             cfrex.F(true),
				TotalActiveUsers:      cfrex.F(int64(200)),
				TurnstileAction:       cfrex.F(cfrex.TurnstileEventActionLog),
				TurnstileMode:         cfrex.F(cfrex.TurnstileEventModeOff),
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

func TestZoneWaitingRoomEventGet(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Events.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneWaitingRoomEventUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Events.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		cfrex.ZoneWaitingRoomEventUpdateParams{
			QueryEvent: cfrex.QueryEventParam{
				EventEndTime:          cfrex.F("2021-09-28T17:00:00.000Z"),
				EventStartTime:        cfrex.F("2021-09-28T15:30:00.000Z"),
				Name:                  cfrex.F("production_webinar_event"),
				CustomPageHTML:        cfrex.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Event is prequeueing / Queue all enabled {{/waitTimeKnown}}"),
				Description:           cfrex.F("Production event - DO NOT MODIFY"),
				DisableSessionRenewal: cfrex.F(true),
				NewUsersPerMinute:     cfrex.F(int64(200)),
				PrequeueStartTime:     cfrex.F("2021-09-28T15:00:00.000Z"),
				QueueingMethod:        cfrex.F("random"),
				SessionDuration:       cfrex.F(int64(1)),
				ShuffleAtEventStart:   cfrex.F(true),
				Suspended:             cfrex.F(true),
				TotalActiveUsers:      cfrex.F(int64(200)),
				TurnstileAction:       cfrex.F(cfrex.TurnstileEventActionLog),
				TurnstileMode:         cfrex.F(cfrex.TurnstileEventModeOff),
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

func TestZoneWaitingRoomEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Events.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		cfrex.ZoneWaitingRoomEventListParams{
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

func TestZoneWaitingRoomEventDelete(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Events.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		cfrex.ZoneWaitingRoomEventDeleteParams{
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

func TestZoneWaitingRoomEventPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Events.Patch(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
		cfrex.ZoneWaitingRoomEventPatchParams{
			QueryEvent: cfrex.QueryEventParam{
				EventEndTime:          cfrex.F("2021-09-28T17:00:00.000Z"),
				EventStartTime:        cfrex.F("2021-09-28T15:30:00.000Z"),
				Name:                  cfrex.F("production_webinar_event"),
				CustomPageHTML:        cfrex.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Event is prequeueing / Queue all enabled {{/waitTimeKnown}}"),
				Description:           cfrex.F("Production event - DO NOT MODIFY"),
				DisableSessionRenewal: cfrex.F(true),
				NewUsersPerMinute:     cfrex.F(int64(200)),
				PrequeueStartTime:     cfrex.F("2021-09-28T15:00:00.000Z"),
				QueueingMethod:        cfrex.F("random"),
				SessionDuration:       cfrex.F(int64(1)),
				ShuffleAtEventStart:   cfrex.F(true),
				Suspended:             cfrex.F(true),
				TotalActiveUsers:      cfrex.F(int64(200)),
				TurnstileAction:       cfrex.F(cfrex.TurnstileEventActionLog),
				TurnstileMode:         cfrex.F(cfrex.TurnstileEventModeOff),
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

func TestZoneWaitingRoomEventPreview(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Events.Preview(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		"25756b2dfe6e378a06b033b670413757",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
