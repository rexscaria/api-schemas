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

func TestAccountR2BucketLifecycleGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.R2.Buckets.Lifecycle.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		cfrex.AccountR2BucketLifecycleGetParams{
			Jurisdiction: cfrex.F(cfrex.AccountR2BucketLifecycleGetParamsCfR2JurisdictionDefault),
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

func TestAccountR2BucketLifecycleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.R2.Buckets.Lifecycle.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-bucket",
		cfrex.AccountR2BucketLifecycleUpdateParams{
			Rules: cfrex.F([]cfrex.R2LifecycleRuleParam{{
				ID: cfrex.F("Expire all objects older than 24 hours"),
				Conditions: cfrex.F(cfrex.R2LifecycleRuleConditionsParam{
					Prefix: cfrex.F("prefix"),
				}),
				Enabled: cfrex.F(true),
				AbortMultipartUploadsTransition: cfrex.F(cfrex.R2LifecycleRuleAbortMultipartUploadsTransitionParam{
					Condition: cfrex.F(cfrex.R2LifecycleAgeConditionParam{
						MaxAge: cfrex.F(int64(0)),
						Type:   cfrex.F(cfrex.R2LifecycleAgeConditionTypeAge),
					}),
				}),
				DeleteObjectsTransition: cfrex.F(cfrex.R2LifecycleRuleDeleteObjectsTransitionParam{
					Condition: cfrex.F[cfrex.R2LifecycleRuleDeleteObjectsTransitionConditionUnionParam](cfrex.R2LifecycleAgeConditionParam{
						MaxAge: cfrex.F(int64(0)),
						Type:   cfrex.F(cfrex.R2LifecycleAgeConditionTypeAge),
					}),
				}),
				StorageClassTransitions: cfrex.F([]cfrex.R2LifecycleRuleStorageClassTransitionParam{{
					Condition: cfrex.F[cfrex.R2LifecycleRuleStorageClassTransitionsConditionUnionParam](cfrex.R2LifecycleAgeConditionParam{
						MaxAge: cfrex.F(int64(0)),
						Type:   cfrex.F(cfrex.R2LifecycleAgeConditionTypeAge),
					}),
					StorageClass: cfrex.F(cfrex.R2LifecycleRuleStorageClassTransitionsStorageClassInfrequentAccess),
				}}),
			}}),
			Jurisdiction: cfrex.F(cfrex.AccountR2BucketLifecycleUpdateParamsCfR2JurisdictionDefault),
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
