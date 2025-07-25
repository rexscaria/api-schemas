// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/rexscaria/api-schemas"
	"github.com/rexscaria/api-schemas/internal/testutil"
	"github.com/rexscaria/api-schemas/option"
)

func TestAccountImageV2ListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Images.V2.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountImageV2ListParams{
			ContinuationToken: cfrex.F("continuation_token"),
			Creator:           cfrex.F("creator"),
			PerPage:           cfrex.F(10.000000),
			SortOrder:         cfrex.F(cfrex.AccountImageV2ListParamsSortOrderAsc),
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

func TestAccountImageV2NewDirectUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Images.V2.NewDirectUpload(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountImageV2NewDirectUploadParams{
			ID:                cfrex.F("this/is/my-customid"),
			Creator:           cfrex.F("creator"),
			Expiry:            cfrex.F(time.Now()),
			Metadata:          cfrex.F[any](map[string]interface{}{}),
			RequireSignedURLs: cfrex.F(true),
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
