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

func TestAccountCloudforceOneEventNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.CloudforceOne.Events.New(
		context.TODO(),
		"account_id",
		cfrex.AccountCloudforceOneEventNewParams{
			Category:      cfrex.F("Domain Resolution"),
			Date:          cfrex.F(time.Now()),
			Event:         cfrex.F("An attacker registered the domain domain.com"),
			IndicatorType: cfrex.F("domain"),
			Raw: cfrex.F(cfrex.AccountCloudforceOneEventNewParamsRaw{
				Data: cfrex.F(map[string]interface{}{
					"foo": "bar",
				}),
				Source: cfrex.F("example.com"),
				Tlp:    cfrex.F("amber"),
			}),
			Tlp:             cfrex.F("amber"),
			AccountID:       cfrex.F(123456.000000),
			Attacker:        cfrex.F("Flying Yeti"),
			AttackerCountry: cfrex.F("CN"),
			DatasetID:       cfrex.F("durableObjectName"),
			Indicator:       cfrex.F("domain.com"),
			Tags:            cfrex.F([]string{"malware"}),
			TargetCountry:   cfrex.F("US"),
			TargetIndustry:  cfrex.F("Agriculture"),
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

func TestAccountCloudforceOneEventGet(t *testing.T) {
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
	_, err := client.Accounts.CloudforceOne.Events.Get(
		context.TODO(),
		"account_id",
		"event_id",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountCloudforceOneEventUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.CloudforceOne.Events.Update(
		context.TODO(),
		"account_id",
		"event_id",
		cfrex.AccountCloudforceOneEventUpdateParams{
			Attacker:        cfrex.F("Flying Yeti"),
			AttackerCountry: cfrex.F("CN"),
			Category:        cfrex.F("Domain Resolution"),
			Date:            cfrex.F(time.Now()),
			Event:           cfrex.F("An attacker registered the domain domain.com"),
			Indicator:       cfrex.F("domain2.com"),
			IndicatorType:   cfrex.F("sha256"),
			Insight:         cfrex.F("new insight"),
			Raw: cfrex.F(cfrex.AccountCloudforceOneEventUpdateParamsRaw{
				Data: cfrex.F(map[string]interface{}{
					"foo": "bar",
				}),
				Source: cfrex.F("example.com"),
				Tlp:    cfrex.F("amber"),
			}),
			TargetCountry:  cfrex.F("US"),
			TargetIndustry: cfrex.F("Insurance"),
			Tlp:            cfrex.F("amber"),
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

func TestAccountCloudforceOneEventDelete(t *testing.T) {
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
	_, err := client.Accounts.CloudforceOne.Events.Delete(
		context.TODO(),
		"account_id",
		"event_id",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountCloudforceOneEventNewBulk(t *testing.T) {
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
	_, err := client.Accounts.CloudforceOne.Events.NewBulk(
		context.TODO(),
		"account_id",
		cfrex.AccountCloudforceOneEventNewBulkParams{
			Data: cfrex.F([]cfrex.AccountCloudforceOneEventNewBulkParamsData{{
				Category:      cfrex.F("Domain Resolution"),
				Date:          cfrex.F(time.Now()),
				Event:         cfrex.F("An attacker registered the domain domain.com"),
				IndicatorType: cfrex.F("domain"),
				Raw: cfrex.F(cfrex.AccountCloudforceOneEventNewBulkParamsDataRaw{
					Data: cfrex.F(map[string]interface{}{
						"foo": "bar",
					}),
					Source: cfrex.F("example.com"),
					Tlp:    cfrex.F("amber"),
				}),
				Tlp:             cfrex.F("amber"),
				AccountID:       cfrex.F(123456.000000),
				Attacker:        cfrex.F("Flying Yeti"),
				AttackerCountry: cfrex.F("CN"),
				DatasetID:       cfrex.F("durableObjectName"),
				Indicator:       cfrex.F("domain.com"),
				Tags:            cfrex.F([]string{"malware"}),
				TargetCountry:   cfrex.F("US"),
				TargetIndustry:  cfrex.F("Agriculture"),
			}}),
			DatasetID: cfrex.F("durableObjectName"),
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

func TestAccountCloudforceOneEventListAttackers(t *testing.T) {
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
	_, err := client.Accounts.CloudforceOne.Events.ListAttackers(context.TODO(), "account_id")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountCloudforceOneEventListCountries(t *testing.T) {
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
	_, err := client.Accounts.CloudforceOne.Events.ListCountries(context.TODO(), "account_id")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountCloudforceOneEventListIndicatorTypes(t *testing.T) {
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
	_, err := client.Accounts.CloudforceOne.Events.ListIndicatorTypes(context.TODO(), "account_id")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountCloudforceOneEventListTargetIndustries(t *testing.T) {
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
	_, err := client.Accounts.CloudforceOne.Events.ListTargetIndustries(context.TODO(), "account_id")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
