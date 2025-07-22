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
	"github.com/stainless-sdks/cf-rex-go/shared"
)

func TestAccountAIGatewayGatewayDatasetNewDataset(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.Datasets.NewDataset(
		context.TODO(),
		"3ebbcb006d4d46d7bb6a8c7f14676cb0",
		"my-gateway",
		cfrex.AccountAIGatewayGatewayDatasetNewDatasetParams{
			Enable: cfrex.F(true),
			Filters: cfrex.F([]cfrex.AccountAIGatewayGatewayDatasetNewDatasetParamsFilter{{
				Key:      cfrex.F(cfrex.AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyCreatedAt),
				Operator: cfrex.F(cfrex.AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperatorEq),
				Value:    cfrex.F([]cfrex.AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersValueUnion{shared.UnionString("string")}),
			}}),
			Name: cfrex.F("name"),
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

func TestAccountAIGatewayGatewayDatasetDeleteDataset(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.Datasets.DeleteDataset(
		context.TODO(),
		"3ebbcb006d4d46d7bb6a8c7f14676cb0",
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

func TestAccountAIGatewayGatewayDatasetFetchDataset(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.Datasets.FetchDataset(
		context.TODO(),
		"3ebbcb006d4d46d7bb6a8c7f14676cb0",
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

func TestAccountAIGatewayGatewayDatasetListDatasetsWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.Datasets.ListDatasets(
		context.TODO(),
		"3ebbcb006d4d46d7bb6a8c7f14676cb0",
		"my-gateway",
		cfrex.AccountAIGatewayGatewayDatasetListDatasetsParams{
			Enable:  cfrex.F(true),
			Name:    cfrex.F("name"),
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

func TestAccountAIGatewayGatewayDatasetUpdateDataset(t *testing.T) {
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
	_, err := client.Accounts.AIGateway.Gateways.Datasets.UpdateDataset(
		context.TODO(),
		"3ebbcb006d4d46d7bb6a8c7f14676cb0",
		"my-gateway",
		"id",
		cfrex.AccountAIGatewayGatewayDatasetUpdateDatasetParams{
			Enable: cfrex.F(true),
			Filters: cfrex.F([]cfrex.AccountAIGatewayGatewayDatasetUpdateDatasetParamsFilter{{
				Key:      cfrex.F(cfrex.AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyCreatedAt),
				Operator: cfrex.F(cfrex.AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperatorEq),
				Value:    cfrex.F([]cfrex.AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersValueUnion{shared.UnionString("string")}),
			}}),
			Name: cfrex.F("name"),
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
