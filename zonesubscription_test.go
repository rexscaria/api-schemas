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

func TestZoneSubscriptionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Subscription.New(
		context.TODO(),
		"506e3185e9c882d175a2d0cb0093d9f2",
		cfrex.ZoneSubscriptionNewParams{
			SubscriptionV2: cfrex.SubscriptionV2Param{
				App: cfrex.F(cfrex.SubscriptionV2AppParam{
					InstallID: cfrex.F("install_id"),
				}),
				ComponentValues: cfrex.F([]cfrex.SubscriptionV2ComponentValueParam{{
					Default: cfrex.F(5.000000),
					Name:    cfrex.F("page_rules"),
					Price:   cfrex.F(5.000000),
					Value:   cfrex.F(20.000000),
				}}),
				Frequency: cfrex.F(cfrex.SubscriptionV2FrequencyMonthly),
				RatePlan: cfrex.F(cfrex.SubscriptionV2RatePlanParam{
					ID:                cfrex.F("free"),
					Currency:          cfrex.F("USD"),
					ExternallyManaged: cfrex.F(false),
					IsContract:        cfrex.F(false),
					PublicName:        cfrex.F("Business Plan"),
					Scope:             cfrex.F("zone"),
					Sets:              cfrex.F([]string{"string"}),
				}),
				Zone: cfrex.F(cfrex.SubscriptionV2ZoneParam{}),
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

func TestZoneSubscriptionGet(t *testing.T) {
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
	_, err := client.Zones.Subscription.Get(context.TODO(), "506e3185e9c882d175a2d0cb0093d9f2")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneSubscriptionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Subscription.Update(
		context.TODO(),
		"506e3185e9c882d175a2d0cb0093d9f2",
		cfrex.ZoneSubscriptionUpdateParams{
			SubscriptionV2: cfrex.SubscriptionV2Param{
				App: cfrex.F(cfrex.SubscriptionV2AppParam{
					InstallID: cfrex.F("install_id"),
				}),
				ComponentValues: cfrex.F([]cfrex.SubscriptionV2ComponentValueParam{{
					Default: cfrex.F(5.000000),
					Name:    cfrex.F("page_rules"),
					Price:   cfrex.F(5.000000),
					Value:   cfrex.F(20.000000),
				}}),
				Frequency: cfrex.F(cfrex.SubscriptionV2FrequencyMonthly),
				RatePlan: cfrex.F(cfrex.SubscriptionV2RatePlanParam{
					ID:                cfrex.F("free"),
					Currency:          cfrex.F("USD"),
					ExternallyManaged: cfrex.F(false),
					IsContract:        cfrex.F(false),
					PublicName:        cfrex.F("Business Plan"),
					Scope:             cfrex.F("zone"),
					Sets:              cfrex.F([]string{"string"}),
				}),
				Zone: cfrex.F(cfrex.SubscriptionV2ZoneParam{}),
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
