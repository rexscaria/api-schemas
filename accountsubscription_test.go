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

func TestAccountSubscriptionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Subscriptions.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountSubscriptionNewParams{
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

func TestAccountSubscriptionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Subscriptions.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"506e3185e9c882d175a2d0cb0093d9f2",
		cfrex.AccountSubscriptionUpdateParams{
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

func TestAccountSubscriptionList(t *testing.T) {
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
	_, err := client.Accounts.Subscriptions.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountSubscriptionDelete(t *testing.T) {
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
	_, err := client.Accounts.Subscriptions.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"506e3185e9c882d175a2d0cb0093d9f2",
		cfrex.AccountSubscriptionDeleteParams{
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
