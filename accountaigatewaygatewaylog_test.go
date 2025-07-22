// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/cf-rex-go"
	"github.com/stainless-sdks/cf-rex-go/internal/testutil"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/stainless-sdks/cf-rex-go/shared"
)

func TestAccountAIGatewayGatewayLogDeleteGatewayLogsWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.Logs.DeleteGatewayLogs(
		context.TODO(),
		"0d37909e38d3e99c29fa2cd343ac421a",
		"my-gateway",
		cfrex.AccountAIGatewayGatewayLogDeleteGatewayLogsParams{
			Filters: cfrex.F([]cfrex.AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFilter{{
				Key:      cfrex.F(cfrex.AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyID),
				Operator: cfrex.F(cfrex.AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperatorEq),
				Value:    cfrex.F([]cfrex.AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersValueUnion{shared.UnionString("string")}),
			}}),
			Limit:            cfrex.F(int64(1)),
			OrderBy:          cfrex.F(cfrex.AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByCreatedAt),
			OrderByDirection: cfrex.F(cfrex.AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByDirectionAsc),
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

func TestAccountAIGatewayGatewayLogGetGatewayLogDetail(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.Logs.GetGatewayLogDetail(
		context.TODO(),
		"0d37909e38d3e99c29fa2cd343ac421a",
		"my-gateway",
		"id",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAIGatewayGatewayLogGetGatewayLogRequest(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.Logs.GetGatewayLogRequest(
		context.TODO(),
		"0d37909e38d3e99c29fa2cd343ac421a",
		"my-gateway",
		"id",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAIGatewayGatewayLogGetGatewayLogResponse(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.Logs.GetGatewayLogResponse(
		context.TODO(),
		"0d37909e38d3e99c29fa2cd343ac421a",
		"my-gateway",
		"id",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAIGatewayGatewayLogListGatewayLogsWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.Logs.ListGatewayLogs(
		context.TODO(),
		"0d37909e38d3e99c29fa2cd343ac421a",
		"my-gateway",
		cfrex.AccountAIGatewayGatewayLogListGatewayLogsParams{
			Cached:    cfrex.F(true),
			Direction: cfrex.F(cfrex.AccountAIGatewayGatewayLogListGatewayLogsParamsDirectionAsc),
			EndDate:   cfrex.F(time.Now()),
			Feedback:  cfrex.F(cfrex.AccountAIGatewayGatewayLogListGatewayLogsParamsFeedbackMinus1),
			Filters: cfrex.F([]cfrex.AccountAIGatewayGatewayLogListGatewayLogsParamsFilter{{
				Key:      cfrex.F(cfrex.AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyID),
				Operator: cfrex.F(cfrex.AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperatorEq),
				Value:    cfrex.F([]cfrex.AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersValueUnion{shared.UnionString("string")}),
			}}),
			MaxCost:             cfrex.F(0.000000),
			MaxDuration:         cfrex.F(0.000000),
			MaxTokensIn:         cfrex.F(0.000000),
			MaxTokensOut:        cfrex.F(0.000000),
			MaxTotalTokens:      cfrex.F(0.000000),
			MetaInfo:            cfrex.F(true),
			MinCost:             cfrex.F(0.000000),
			MinDuration:         cfrex.F(0.000000),
			MinTokensIn:         cfrex.F(0.000000),
			MinTokensOut:        cfrex.F(0.000000),
			MinTotalTokens:      cfrex.F(0.000000),
			Model:               cfrex.F("model"),
			ModelType:           cfrex.F("model_type"),
			OrderBy:             cfrex.F(cfrex.AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByCreatedAt),
			OrderByDirection:    cfrex.F(cfrex.AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByDirectionAsc),
			Page:                cfrex.F(int64(1)),
			PerPage:             cfrex.F(int64(1)),
			Provider:            cfrex.F("provider"),
			RequestContentType:  cfrex.F("request_content_type"),
			ResponseContentType: cfrex.F("response_content_type"),
			Search:              cfrex.F("search"),
			StartDate:           cfrex.F(time.Now()),
			Success:             cfrex.F(true),
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

func TestAccountAIGatewayGatewayLogPatchGatewayLogWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.Logs.PatchGatewayLog(
		context.TODO(),
		"0d37909e38d3e99c29fa2cd343ac421a",
		"my-gateway",
		"id",
		cfrex.AccountAIGatewayGatewayLogPatchGatewayLogParams{
			Feedback: cfrex.F(-1.000000),
			Metadata: cfrex.F(map[string]cfrex.AccountAIGatewayGatewayLogPatchGatewayLogParamsMetadataUnion{
				"foo": shared.UnionString("string"),
			}),
			Score: cfrex.F(0.000000),
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
