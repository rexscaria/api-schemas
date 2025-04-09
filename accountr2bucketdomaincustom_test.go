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

func TestAccountR2BucketDomainCustomGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.R2.Buckets.Domains.Custom.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		"example-domain/custom-domain.com",
		cfrex.AccountR2BucketDomainCustomGetParams{
			Jurisdiction: cfrex.F(cfrex.AccountR2BucketDomainCustomGetParamsCfR2JurisdictionDefault),
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

func TestAccountR2BucketDomainCustomUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.R2.Buckets.Domains.Custom.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		"example-domain/custom-domain.com",
		cfrex.AccountR2BucketDomainCustomUpdateParams{
			Enabled:      cfrex.F(true),
			MinTls:       cfrex.F(cfrex.AccountR2BucketDomainCustomUpdateParamsMinTls1_2),
			Jurisdiction: cfrex.F(cfrex.AccountR2BucketDomainCustomUpdateParamsCfR2JurisdictionDefault),
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

func TestAccountR2BucketDomainCustomListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.R2.Buckets.Domains.Custom.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		cfrex.AccountR2BucketDomainCustomListParams{
			Jurisdiction: cfrex.F(cfrex.AccountR2BucketDomainCustomListParamsCfR2JurisdictionDefault),
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

func TestAccountR2BucketDomainCustomAttachWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.R2.Buckets.Domains.Custom.Attach(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		cfrex.AccountR2BucketDomainCustomAttachParams{
			Domain:       cfrex.F("prefix.example-domain.com"),
			Enabled:      cfrex.F(true),
			ZoneID:       cfrex.F("36ca64a6d92827b8a6b90be344bb1bfd"),
			MinTls:       cfrex.F(cfrex.AccountR2BucketDomainCustomAttachParamsMinTls1_0),
			Jurisdiction: cfrex.F(cfrex.AccountR2BucketDomainCustomAttachParamsCfR2JurisdictionDefault),
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

func TestAccountR2BucketDomainCustomRemoveWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.R2.Buckets.Domains.Custom.Remove(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		"example-domain/custom-domain.com",
		cfrex.AccountR2BucketDomainCustomRemoveParams{
			Jurisdiction: cfrex.F(cfrex.AccountR2BucketDomainCustomRemoveParamsCfR2JurisdictionDefault),
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
