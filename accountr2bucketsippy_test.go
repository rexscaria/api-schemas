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

func TestAccountR2BucketSippyGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.R2.Buckets.Sippy.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		cfrex.AccountR2BucketSippyGetParams{
			Jurisdiction: cfrex.F(cfrex.AccountR2BucketSippyGetParamsCfR2JurisdictionDefault),
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

func TestAccountR2BucketSippyDisableWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.R2.Buckets.Sippy.Disable(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		cfrex.AccountR2BucketSippyDisableParams{
			Jurisdiction: cfrex.F(cfrex.AccountR2BucketSippyDisableParamsCfR2JurisdictionDefault),
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

func TestAccountR2BucketSippyEnableWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.R2.Buckets.Sippy.Enable(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		cfrex.AccountR2BucketSippyEnableParams{
			Body: cfrex.AccountR2BucketSippyEnableParamsBodyR2EnableSippyAws{
				Destination: cfrex.F(cfrex.AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsDestination{
					AccessKeyID:     cfrex.F("accessKeyId"),
					Provider:        cfrex.F(cfrex.AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsDestinationProviderR2),
					SecretAccessKey: cfrex.F("secretAccessKey"),
				}),
				Source: cfrex.F(cfrex.AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsSource{
					AccessKeyID:     cfrex.F("accessKeyId"),
					Bucket:          cfrex.F("bucket"),
					Provider:        cfrex.F(cfrex.AccountR2BucketSippyEnableParamsBodyR2EnableSippyAwsSourceProviderAws),
					Region:          cfrex.F("region"),
					SecretAccessKey: cfrex.F("secretAccessKey"),
				}),
			},
			Jurisdiction: cfrex.F(cfrex.AccountR2BucketSippyEnableParamsCfR2JurisdictionDefault),
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
