// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/cf-rex-go"
	"github.com/stainless-sdks/cf-rex-go/internal/testutil"
	"github.com/stainless-sdks/cf-rex-go/option"
)

func TestZoneAPIGatewayUserSchemaGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.APIGateway.UserSchemas.Get(
		context.TODO(),
		"zone_id",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.ZoneAPIGatewayUserSchemaGetParams{
			OmitSource: cfrex.F(true),
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

func TestZoneAPIGatewayUserSchemaListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.APIGateway.UserSchemas.List(
		context.TODO(),
		"zone_id",
		cfrex.ZoneAPIGatewayUserSchemaListParams{
			OmitSource:        cfrex.F(true),
			Page:              cfrex.F(int64(1)),
			PerPage:           cfrex.F(int64(5)),
			ValidationEnabled: cfrex.F(true),
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

func TestZoneAPIGatewayUserSchemaDelete(t *testing.T) {
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
	_, err := client.Zones.APIGateway.UserSchemas.Delete(
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

func TestZoneAPIGatewayUserSchemaEnableValidationWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.APIGateway.UserSchemas.EnableValidation(
		context.TODO(),
		"zone_id",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.ZoneAPIGatewayUserSchemaEnableValidationParams{
			ValidationEnabled: cfrex.F(cfrex.ZoneAPIGatewayUserSchemaEnableValidationParamsValidationEnabledTrue),
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

func TestZoneAPIGatewayUserSchemaGetHostsWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.APIGateway.UserSchemas.GetHosts(
		context.TODO(),
		"zone_id",
		cfrex.ZoneAPIGatewayUserSchemaGetHostsParams{
			Page:    cfrex.F(int64(1)),
			PerPage: cfrex.F(int64(5)),
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

func TestZoneAPIGatewayUserSchemaGetOperationsWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.APIGateway.UserSchemas.GetOperations(
		context.TODO(),
		"zone_id",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.ZoneAPIGatewayUserSchemaGetOperationsParams{
			Endpoint:        cfrex.F("/api/v1"),
			Feature:         cfrex.F([]cfrex.ZoneAPIGatewayUserSchemaGetOperationsParamsFeature{cfrex.ZoneAPIGatewayUserSchemaGetOperationsParamsFeatureThresholds}),
			Host:            cfrex.F([]string{"api.cloudflare.com"}),
			Method:          cfrex.F([]string{"GET"}),
			OperationStatus: cfrex.F(cfrex.ZoneAPIGatewayUserSchemaGetOperationsParamsOperationStatusNew),
			Page:            cfrex.F(int64(1)),
			PerPage:         cfrex.F(int64(5)),
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

func TestZoneAPIGatewayUserSchemaUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.APIGateway.UserSchemas.Upload(
		context.TODO(),
		"zone_id",
		cfrex.ZoneAPIGatewayUserSchemaUploadParams{
			File:              cfrex.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
			Kind:              cfrex.F(cfrex.APIShieldKindOpenAPIV3),
			Name:              cfrex.F("petstore schema"),
			ValidationEnabled: cfrex.F(cfrex.ZoneAPIGatewayUserSchemaUploadParamsValidationEnabledTrue),
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
