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

func TestAccountAlertingV3PolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Alerting.V3.Policies.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAlertingV3PolicyNewParams{
			AlertType: cfrex.F(cfrex.AlertTypeAccessCustomCertificateExpirationType),
			Enabled:   cfrex.F(true),
			Mechanisms: cfrex.F(cfrex.MechanismsParam{
				Email: cfrex.F([]cfrex.MechanismsEmailParam{{
					ID: cfrex.F("test@example.com"),
				}}),
				Pagerduty: cfrex.F([]cfrex.MechanismsPagerdutyParam{{
					ID: cfrex.F("e8133a15-00a4-4d69-aec1-32f70c51f6e5"),
				}}),
				Webhooks: cfrex.F([]cfrex.MechanismsWebhookParam{{
					ID: cfrex.F("14cc1190-5d2b-4b98-a696-c424cb2ad05f"),
				}}),
			}),
			Name:          cfrex.F("SSL Notification Event Policy"),
			AlertInterval: cfrex.F("30m"),
			Description:   cfrex.F("Something describing the policy."),
			Filters: cfrex.F(cfrex.AlertFiltersParam{
				Actions:                      cfrex.F([]string{"string"}),
				AffectedAsns:                 cfrex.F([]string{"string"}),
				AffectedComponents:           cfrex.F([]string{"string"}),
				AffectedLocations:            cfrex.F([]string{"string"}),
				AirportCode:                  cfrex.F([]string{"string"}),
				AlertTriggerPreferences:      cfrex.F([]string{"string"}),
				AlertTriggerPreferencesValue: cfrex.F([]string{"string"}),
				Enabled:                      cfrex.F([]string{"string"}),
				Environment:                  cfrex.F([]string{"string"}),
				Event:                        cfrex.F([]string{"string"}),
				EventSource:                  cfrex.F([]string{"string"}),
				EventType:                    cfrex.F([]string{"string"}),
				GroupBy:                      cfrex.F([]string{"string"}),
				HealthCheckID:                cfrex.F([]string{"string"}),
				IncidentImpact:               cfrex.F([]cfrex.AlertFiltersIncidentImpact{cfrex.AlertFiltersIncidentImpactIncidentImpactNone}),
				InputID:                      cfrex.F([]string{"string"}),
				InsightClass:                 cfrex.F([]string{"string"}),
				Limit:                        cfrex.F([]string{"string"}),
				LogoTag:                      cfrex.F([]string{"string"}),
				MegabitsPerSecond:            cfrex.F([]string{"string"}),
				NewHealth:                    cfrex.F([]string{"string"}),
				NewStatus:                    cfrex.F([]string{"string"}),
				PacketsPerSecond:             cfrex.F([]string{"string"}),
				PoolID:                       cfrex.F([]string{"string"}),
				PopNames:                     cfrex.F([]string{"string"}),
				Product:                      cfrex.F([]string{"string"}),
				ProjectID:                    cfrex.F([]string{"string"}),
				Protocol:                     cfrex.F([]string{"string"}),
				QueryTag:                     cfrex.F([]string{"string"}),
				RequestsPerSecond:            cfrex.F([]string{"string"}),
				Selectors:                    cfrex.F([]string{"string"}),
				Services:                     cfrex.F([]string{"string"}),
				Slo:                          cfrex.F([]string{"99.9"}),
				Status:                       cfrex.F([]string{"string"}),
				TargetHostname:               cfrex.F([]string{"string"}),
				TargetIP:                     cfrex.F([]string{"string"}),
				TargetZoneName:               cfrex.F([]string{"string"}),
				TrafficExclusions:            cfrex.F([]cfrex.AlertFiltersTrafficExclusion{cfrex.AlertFiltersTrafficExclusionSecurityEvents}),
				TunnelID:                     cfrex.F([]string{"string"}),
				TunnelName:                   cfrex.F([]string{"string"}),
				Where:                        cfrex.F([]string{"string"}),
				Zones:                        cfrex.F([]string{"string"}),
			}),
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

func TestAccountAlertingV3PolicyGet(t *testing.T) {
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
	_, err := client.Accounts.Alerting.V3.Policies.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0da2b59e-f118-439d-8097-bdfb215203c9",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAlertingV3PolicyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Alerting.V3.Policies.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0da2b59e-f118-439d-8097-bdfb215203c9",
		cfrex.AccountAlertingV3PolicyUpdateParams{
			AlertInterval: cfrex.F("30m"),
			AlertType:     cfrex.F(cfrex.AlertTypeAccessCustomCertificateExpirationType),
			Description:   cfrex.F("Something describing the policy."),
			Enabled:       cfrex.F(true),
			Filters: cfrex.F(cfrex.AlertFiltersParam{
				Actions:                      cfrex.F([]string{"string"}),
				AffectedAsns:                 cfrex.F([]string{"string"}),
				AffectedComponents:           cfrex.F([]string{"string"}),
				AffectedLocations:            cfrex.F([]string{"string"}),
				AirportCode:                  cfrex.F([]string{"string"}),
				AlertTriggerPreferences:      cfrex.F([]string{"string"}),
				AlertTriggerPreferencesValue: cfrex.F([]string{"string"}),
				Enabled:                      cfrex.F([]string{"string"}),
				Environment:                  cfrex.F([]string{"string"}),
				Event:                        cfrex.F([]string{"string"}),
				EventSource:                  cfrex.F([]string{"string"}),
				EventType:                    cfrex.F([]string{"string"}),
				GroupBy:                      cfrex.F([]string{"string"}),
				HealthCheckID:                cfrex.F([]string{"string"}),
				IncidentImpact:               cfrex.F([]cfrex.AlertFiltersIncidentImpact{cfrex.AlertFiltersIncidentImpactIncidentImpactNone}),
				InputID:                      cfrex.F([]string{"string"}),
				InsightClass:                 cfrex.F([]string{"string"}),
				Limit:                        cfrex.F([]string{"string"}),
				LogoTag:                      cfrex.F([]string{"string"}),
				MegabitsPerSecond:            cfrex.F([]string{"string"}),
				NewHealth:                    cfrex.F([]string{"string"}),
				NewStatus:                    cfrex.F([]string{"string"}),
				PacketsPerSecond:             cfrex.F([]string{"string"}),
				PoolID:                       cfrex.F([]string{"string"}),
				PopNames:                     cfrex.F([]string{"string"}),
				Product:                      cfrex.F([]string{"string"}),
				ProjectID:                    cfrex.F([]string{"string"}),
				Protocol:                     cfrex.F([]string{"string"}),
				QueryTag:                     cfrex.F([]string{"string"}),
				RequestsPerSecond:            cfrex.F([]string{"string"}),
				Selectors:                    cfrex.F([]string{"string"}),
				Services:                     cfrex.F([]string{"string"}),
				Slo:                          cfrex.F([]string{"99.9"}),
				Status:                       cfrex.F([]string{"string"}),
				TargetHostname:               cfrex.F([]string{"string"}),
				TargetIP:                     cfrex.F([]string{"string"}),
				TargetZoneName:               cfrex.F([]string{"string"}),
				TrafficExclusions:            cfrex.F([]cfrex.AlertFiltersTrafficExclusion{cfrex.AlertFiltersTrafficExclusionSecurityEvents}),
				TunnelID:                     cfrex.F([]string{"string"}),
				TunnelName:                   cfrex.F([]string{"string"}),
				Where:                        cfrex.F([]string{"string"}),
				Zones:                        cfrex.F([]string{"string"}),
			}),
			Mechanisms: cfrex.F(cfrex.MechanismsParam{
				Email: cfrex.F([]cfrex.MechanismsEmailParam{{
					ID: cfrex.F("test@example.com"),
				}}),
				Pagerduty: cfrex.F([]cfrex.MechanismsPagerdutyParam{{
					ID: cfrex.F("e8133a15-00a4-4d69-aec1-32f70c51f6e5"),
				}}),
				Webhooks: cfrex.F([]cfrex.MechanismsWebhookParam{{
					ID: cfrex.F("14cc1190-5d2b-4b98-a696-c424cb2ad05f"),
				}}),
			}),
			Name: cfrex.F("SSL Notification Event Policy"),
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

func TestAccountAlertingV3PolicyList(t *testing.T) {
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
	_, err := client.Accounts.Alerting.V3.Policies.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAlertingV3PolicyDelete(t *testing.T) {
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
	_, err := client.Accounts.Alerting.V3.Policies.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0da2b59e-f118-439d-8097-bdfb215203c9",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
