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

func TestZoneAPIGatewayOperationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.APIGateway.Operations.Get(
		context.TODO(),
		"zone_id",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.ZoneAPIGatewayOperationGetParams{
			Feature: cfrex.F([]cfrex.ZoneAPIGatewayOperationGetParamsFeature{cfrex.ZoneAPIGatewayOperationGetParamsFeatureThresholds}),
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

func TestZoneAPIGatewayOperationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.APIGateway.Operations.List(
		context.TODO(),
		"zone_id",
		cfrex.ZoneAPIGatewayOperationListParams{
			Direction: cfrex.F(cfrex.ZoneAPIGatewayOperationListParamsDirectionDesc),
			Endpoint:  cfrex.F("/api/v1"),
			Feature:   cfrex.F([]cfrex.ZoneAPIGatewayOperationListParamsFeature{cfrex.ZoneAPIGatewayOperationListParamsFeatureThresholds}),
			Host:      cfrex.F([]string{"api.cloudflare.com"}),
			Method:    cfrex.F([]string{"GET"}),
			Order:     cfrex.F(cfrex.ZoneAPIGatewayOperationListParamsOrderMethod),
			Page:      cfrex.F(int64(1)),
			PerPage:   cfrex.F(int64(5)),
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

func TestZoneAPIGatewayOperationAddMultiple(t *testing.T) {
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
	_, err := client.Zones.APIGateway.Operations.AddMultiple(
		context.TODO(),
		"zone_id",
		cfrex.ZoneAPIGatewayOperationAddMultipleParams{
			Body: []cfrex.BasicOperationParam{{
				Endpoint: cfrex.F("/api/v1/users/{var1}"),
				Host:     cfrex.F("www.example.com"),
				Method:   cfrex.F(cfrex.BasicOperationMethodGet),
			}},
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

func TestZoneAPIGatewayOperationAddSingle(t *testing.T) {
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
	_, err := client.Zones.APIGateway.Operations.AddSingle(
		context.TODO(),
		"zone_id",
		cfrex.ZoneAPIGatewayOperationAddSingleParams{
			BasicOperation: cfrex.BasicOperationParam{
				Endpoint: cfrex.F("/api/v1/users/{var1}"),
				Host:     cfrex.F("www.example.com"),
				Method:   cfrex.F(cfrex.BasicOperationMethodGet),
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

func TestZoneAPIGatewayOperationDeleteMultiple(t *testing.T) {
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
	_, err := client.Zones.APIGateway.Operations.DeleteMultiple(
		context.TODO(),
		"zone_id",
		cfrex.ZoneAPIGatewayOperationDeleteMultipleParams{
			Body: []cfrex.ZoneAPIGatewayOperationDeleteMultipleParamsBody{{
				OperationID: cfrex.F("b17c8043-99a0-4202-b7d9-8f7cdbee02cd"),
			}, {
				OperationID: cfrex.F("3818d821-5901-4147-a474-f5f5aec1d54e"),
			}},
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

func TestZoneAPIGatewayOperationDeleteSingle(t *testing.T) {
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
	_, err := client.Zones.APIGateway.Operations.DeleteSingle(
		context.TODO(),
		"zone_id",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
