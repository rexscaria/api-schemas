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

func TestAccountR2BucketLockGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.R2.Buckets.Lock.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		cfrex.AccountR2BucketLockGetParams{
			Jurisdiction: cfrex.F(cfrex.AccountR2BucketLockGetParamsCfR2JurisdictionDefault),
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

func TestAccountR2BucketLockUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.R2.Buckets.Lock.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		cfrex.AccountR2BucketLockUpdateParams{
			Rules: cfrex.F([]cfrex.R2BucketLockRuleParam{{
				ID: cfrex.F("Lock all objects for 24 hours"),
				Condition: cfrex.F[cfrex.R2BucketLockRuleConditionUnionParam](cfrex.R2BucketLockRuleConditionR2LockRuleAgeConditionParam{
					MaxAgeSeconds: cfrex.F(int64(100)),
					Type:          cfrex.F(cfrex.R2BucketLockRuleConditionR2LockRuleAgeConditionTypeAge),
				}),
				Enabled: cfrex.F(true),
				Prefix:  cfrex.F("prefix"),
			}}),
			Jurisdiction: cfrex.F(cfrex.AccountR2BucketLockUpdateParamsCfR2JurisdictionDefault),
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
