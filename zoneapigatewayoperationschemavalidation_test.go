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

func TestZoneAPIGatewayOperationSchemaValidationGet(t *testing.T) {
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
	_, err := client.Zones.APIGateway.Operations.SchemaValidation.Get(
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

func TestZoneAPIGatewayOperationSchemaValidationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.APIGateway.Operations.SchemaValidation.Update(
		context.TODO(),
		"zone_id",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.ZoneAPIGatewayOperationSchemaValidationUpdateParams{
			SchemaValidationSettings: cfrex.SchemaValidationSettingsParam{
				MitigationAction: cfrex.F(cfrex.SchemaValidationSettingsMitigationActionLog),
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

func TestZoneAPIGatewayOperationSchemaValidationUpdateMultiple(t *testing.T) {
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
	_, err := client.Zones.APIGateway.Operations.SchemaValidation.UpdateMultiple(
		context.TODO(),
		"zone_id",
		cfrex.ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParams{
			Body: map[string]cfrex.ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBody{
				"3818d821-5901-4147-a474-f5f5aec1d54e": {
					MitigationAction: cfrex.F(cfrex.ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBodyMitigationActionLog),
				},
				"b17c8043-99a0-4202-b7d9-8f7cdbee02cd": {
					MitigationAction: cfrex.F(cfrex.ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBodyMitigationActionLog),
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
