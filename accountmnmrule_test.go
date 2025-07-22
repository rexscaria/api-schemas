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

func TestAccountMnmRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Mnm.Rules.New(
		context.TODO(),
		"6f91088a406011ed95aed352566e8d4c",
		cfrex.AccountMnmRuleNewParams{
			Duration:               cfrex.F(cfrex.RuleDuration1m),
			Name:                   cfrex.F("my_rule_1"),
			AutomaticAdvertisement: cfrex.F(true),
			Bandwidth:              cfrex.F(1000.000000),
			PacketThreshold:        cfrex.F(10000.000000),
			Prefixes:               cfrex.F([]string{"203.0.113.1/32"}),
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

func TestAccountMnmRuleGet(t *testing.T) {
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
	_, err := client.Accounts.Mnm.Rules.Get(
		context.TODO(),
		"6f91088a406011ed95aed352566e8d4c",
		"2890e6fa406311ed9b5a23f70f6fb8cf",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountMnmRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Mnm.Rules.Update(
		context.TODO(),
		"6f91088a406011ed95aed352566e8d4c",
		cfrex.AccountMnmRuleUpdateParams{
			Duration:               cfrex.F(cfrex.RuleDuration1m),
			Name:                   cfrex.F("my_rule_1"),
			ID:                     cfrex.F("2890e6fa406311ed9b5a23f70f6fb8cf"),
			AutomaticAdvertisement: cfrex.F(true),
			Bandwidth:              cfrex.F(1000.000000),
			PacketThreshold:        cfrex.F(10000.000000),
			Prefixes:               cfrex.F([]string{"203.0.113.1/32"}),
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

func TestAccountMnmRuleList(t *testing.T) {
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
	_, err := client.Accounts.Mnm.Rules.List(context.TODO(), "6f91088a406011ed95aed352566e8d4c")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountMnmRuleDelete(t *testing.T) {
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
	_, err := client.Accounts.Mnm.Rules.Delete(
		context.TODO(),
		"6f91088a406011ed95aed352566e8d4c",
		"2890e6fa406311ed9b5a23f70f6fb8cf",
		cfrex.AccountMnmRuleDeleteParams{
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

func TestAccountMnmRuleUpdateAdvertisement(t *testing.T) {
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
	_, err := client.Accounts.Mnm.Rules.UpdateAdvertisement(
		context.TODO(),
		"6f91088a406011ed95aed352566e8d4c",
		"2890e6fa406311ed9b5a23f70f6fb8cf",
		cfrex.AccountMnmRuleUpdateAdvertisementParams{
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

func TestAccountMnmRuleUpdateRuleWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Mnm.Rules.UpdateRule(
		context.TODO(),
		"6f91088a406011ed95aed352566e8d4c",
		"2890e6fa406311ed9b5a23f70f6fb8cf",
		cfrex.AccountMnmRuleUpdateRuleParams{
			AutomaticAdvertisement: cfrex.F(true),
			Bandwidth:              cfrex.F(1000.000000),
			Duration:               cfrex.F(cfrex.RuleDuration1m),
			Name:                   cfrex.F("my_rule_1"),
			PacketThreshold:        cfrex.F(10000.000000),
			Prefixes:               cfrex.F([]string{"203.0.113.1/32"}),
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
