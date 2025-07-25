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

func TestAccountAccessLogScimUpdatesWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Access.Logs.Scim.Updates(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAccessLogScimUpdatesParams{
			IdpID:             cfrex.F([]string{"df7e2w5f-02b7-4d9d-af26-8d1988fca630", "0194ae2c-efcf-7cfb-8884-055f1a161fa5"}),
			CfResourceID:      cfrex.F("bd97ef8d-7986-43e3-9ee0-c25dda33e4b0"),
			Direction:         cfrex.F(cfrex.AccountAccessLogScimUpdatesParamsDirectionDesc),
			IdpResourceID:     cfrex.F("idp_resource_id"),
			Limit:             cfrex.F(int64(10)),
			RequestMethod:     cfrex.F([]cfrex.AccountAccessLogScimUpdatesParamsRequestMethod{cfrex.AccountAccessLogScimUpdatesParamsRequestMethodDelete, cfrex.AccountAccessLogScimUpdatesParamsRequestMethodPatch}),
			ResourceGroupName: cfrex.F("ALL_EMPLOYEES"),
			ResourceType:      cfrex.F([]cfrex.AccountAccessLogScimUpdatesParamsResourceType{cfrex.AccountAccessLogScimUpdatesParamsResourceTypeUser, cfrex.AccountAccessLogScimUpdatesParamsResourceTypeGroup}),
			ResourceUserEmail: cfrex.F("john.smith@example.com"),
			Since:             cfrex.F(time.Now()),
			Status:            cfrex.F([]cfrex.AccountAccessLogScimUpdatesParamsStatus{cfrex.AccountAccessLogScimUpdatesParamsStatusFailure, cfrex.AccountAccessLogScimUpdatesParamsStatusSuccess}),
			Until:             cfrex.F(time.Now()),
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
