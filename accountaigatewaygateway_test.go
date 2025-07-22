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

func TestAccountAIGatewayGatewayNewGatewayWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.NewGateway(
		context.TODO(),
		"3ebbcb006d4d46d7bb6a8c7f14676cb0",
		cfrex.AccountAIGatewayGatewayNewGatewayParams{
			ID:                      cfrex.F("my-gateway"),
			CacheInvalidateOnUpdate: cfrex.F(true),
			CacheTtl:                cfrex.F(int64(0)),
			CollectLogs:             cfrex.F(true),
			RateLimitingInterval:    cfrex.F(int64(0)),
			RateLimitingLimit:       cfrex.F(int64(0)),
			RateLimitingTechnique:   cfrex.F(cfrex.AccountAIGatewayGatewayNewGatewayParamsRateLimitingTechniqueFixed),
			Authentication:          cfrex.F(true),
			LogManagement:           cfrex.F(int64(10000)),
			LogManagementStrategy:   cfrex.F(cfrex.AccountAIGatewayGatewayNewGatewayParamsLogManagementStrategyStopInserting),
			Logpush:                 cfrex.F(true),
			LogpushPublicKey:        cfrex.F("xxxxxxxxxxxxxxxx"),
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

func TestAccountAIGatewayGatewayDeleteGateway(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.DeleteGateway(
		context.TODO(),
		"3ebbcb006d4d46d7bb6a8c7f14676cb0",
		"my-gateway",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAIGatewayGatewayFetchGateway(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.FetchGateway(
		context.TODO(),
		"3ebbcb006d4d46d7bb6a8c7f14676cb0",
		"my-gateway",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAIGatewayGatewayGetGatewayURL(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.GetGatewayURL(
		context.TODO(),
		"0d37909e38d3e99c29fa2cd343ac421a",
		"my-gateway",
		"workers-ai",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAIGatewayGatewayListGatewaysWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.ListGateways(
		context.TODO(),
		"3ebbcb006d4d46d7bb6a8c7f14676cb0",
		cfrex.AccountAIGatewayGatewayListGatewaysParams{
			Page:    cfrex.F(int64(1)),
			PerPage: cfrex.F(int64(1)),
			Search:  cfrex.F("search"),
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

func TestAccountAIGatewayGatewayUpdateGatewayWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.UpdateGateway(
		context.TODO(),
		"3ebbcb006d4d46d7bb6a8c7f14676cb0",
		"my-gateway",
		cfrex.AccountAIGatewayGatewayUpdateGatewayParams{
			CacheInvalidateOnUpdate: cfrex.F(true),
			CacheTtl:                cfrex.F(int64(0)),
			CollectLogs:             cfrex.F(true),
			RateLimitingInterval:    cfrex.F(int64(0)),
			RateLimitingLimit:       cfrex.F(int64(0)),
			RateLimitingTechnique:   cfrex.F(cfrex.AccountAIGatewayGatewayUpdateGatewayParamsRateLimitingTechniqueFixed),
			Authentication:          cfrex.F(true),
			LogManagement:           cfrex.F(int64(10000)),
			LogManagementStrategy:   cfrex.F(cfrex.AccountAIGatewayGatewayUpdateGatewayParamsLogManagementStrategyStopInserting),
			Logpush:                 cfrex.F(true),
			LogpushPublicKey:        cfrex.F("xxxxxxxxxxxxxxxx"),
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
