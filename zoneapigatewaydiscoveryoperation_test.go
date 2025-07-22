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

func TestZoneAPIGatewayDiscoveryOperationUpdate(t *testing.T) {
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
	_, err := client.Zones.APIGateway.Discovery.Operations.Update(
		context.TODO(),
		"zone_id",
		cfrex.ZoneAPIGatewayDiscoveryOperationUpdateParams{
			Body: map[string]cfrex.ZoneAPIGatewayDiscoveryOperationUpdateParamsBody{
				"3818d821-5901-4147-a474-f5f5aec1d54e": {
					State: cfrex.F(cfrex.ZoneAPIGatewayDiscoveryOperationUpdateParamsBodyStateIgnored),
				},
				"b17c8043-99a0-4202-b7d9-8f7cdbee02cd": {
					State: cfrex.F(cfrex.ZoneAPIGatewayDiscoveryOperationUpdateParamsBodyStateReview),
				},
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

func TestZoneAPIGatewayDiscoveryOperationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.APIGateway.Discovery.Operations.List(
		context.TODO(),
		"zone_id",
		cfrex.ZoneAPIGatewayDiscoveryOperationListParams{
			Diff:      cfrex.F(true),
			Direction: cfrex.F(cfrex.ZoneAPIGatewayDiscoveryOperationListParamsDirectionDesc),
			Endpoint:  cfrex.F("/api/v1"),
			Host:      cfrex.F([]string{"api.cloudflare.com"}),
			Method:    cfrex.F([]string{"GET"}),
			Order:     cfrex.F(cfrex.ZoneAPIGatewayDiscoveryOperationListParamsOrderMethod),
			Origin:    cfrex.F(cfrex.ZoneAPIGatewayDiscoveryOperationListParamsOriginMl),
			Page:      cfrex.F(int64(1)),
			PerPage:   cfrex.F(int64(5)),
			State:     cfrex.F(cfrex.APIDiscoveryStateReview),
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

func TestZoneAPIGatewayDiscoveryOperationUpdateSingleWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.APIGateway.Discovery.Operations.UpdateSingle(
		context.TODO(),
		"zone_id",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.ZoneAPIGatewayDiscoveryOperationUpdateSingleParams{
			State: cfrex.F(cfrex.ZoneAPIGatewayDiscoveryOperationUpdateSingleParamsStateReview),
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
