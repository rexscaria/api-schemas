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

func TestAccountSlurperSourceCheckConnectivityWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Slurper.Source.CheckConnectivity(
		context.TODO(),
		"account_id",
		cfrex.AccountSlurperSourceCheckConnectivityParams{
			SourceJobSchema: cfrex.SourceJobSchemaR2SlurperS3SourceSchemaParam{
				Bucket:   cfrex.F("bucket"),
				Endpoint: cfrex.F("endpoint"),
				Secret: cfrex.F(cfrex.S3LikeCredsSchemaParam{
					AccessKeyID:     cfrex.F("accessKeyId"),
					SecretAccessKey: cfrex.F("secretAccessKey"),
				}),
				Vendor: cfrex.F(cfrex.SourceJobSchemaR2SlurperS3SourceSchemaVendorS3),
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
