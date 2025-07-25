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

func TestAccountR2BucketCorGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.R2.Buckets.Cors.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		cfrex.AccountR2BucketCorGetParams{
			Jurisdiction: cfrex.F(cfrex.AccountR2BucketCorGetParamsCfR2JurisdictionDefault),
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

func TestAccountR2BucketCorUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.R2.Buckets.Cors.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		cfrex.AccountR2BucketCorUpdateParams{
			Rules: cfrex.F([]cfrex.R2CorsRuleParam{{
				Allowed: cfrex.F(cfrex.R2CorsRuleAllowedParam{
					Methods: cfrex.F([]cfrex.R2CorsRuleAllowedMethod{cfrex.R2CorsRuleAllowedMethodGet}),
					Origins: cfrex.F([]string{"http://localhost:3000"}),
					Headers: cfrex.F([]string{"x-requested-by"}),
				}),
				ID:            cfrex.F("Allow Local Development"),
				ExposeHeaders: cfrex.F([]string{"Content-Encoding"}),
				MaxAgeSeconds: cfrex.F(3600.000000),
			}}),
			Jurisdiction: cfrex.F(cfrex.AccountR2BucketCorUpdateParamsCfR2JurisdictionDefault),
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

func TestAccountR2BucketCorDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.R2.Buckets.Cors.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		cfrex.AccountR2BucketCorDeleteParams{
			Jurisdiction: cfrex.F(cfrex.AccountR2BucketCorDeleteParamsCfR2JurisdictionDefault),
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
