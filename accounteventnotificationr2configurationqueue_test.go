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

func TestAccountEventNotificationR2ConfigurationQueueNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.EventNotifications.R2.Configuration.Queues.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		"queue_id",
		cfrex.AccountEventNotificationR2ConfigurationQueueNewParams{
			Rules: cfrex.F([]cfrex.R2RuleParam{{
				Actions:     cfrex.F([]cfrex.R2RuleAction{cfrex.R2RuleActionPutObject, cfrex.R2RuleActionCopyObject}),
				Description: cfrex.F("Notifications from source bucket to queue"),
				Prefix:      cfrex.F("img/"),
				Suffix:      cfrex.F(".jpeg"),
			}}),
			Jurisdiction: cfrex.F(cfrex.AccountEventNotificationR2ConfigurationQueueNewParamsCfR2JurisdictionDefault),
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

func TestAccountEventNotificationR2ConfigurationQueueDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.EventNotifications.R2.Configuration.Queues.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		"queue_id",
		cfrex.AccountEventNotificationR2ConfigurationQueueDeleteParams{
			RuleIDs:      cfrex.F([]string{"string"}),
			Jurisdiction: cfrex.F(cfrex.AccountEventNotificationR2ConfigurationQueueDeleteParamsCfR2JurisdictionDefault),
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
