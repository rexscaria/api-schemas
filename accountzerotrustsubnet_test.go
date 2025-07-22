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

func TestAccountZerotrustSubnetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Zerotrust.Subnets.List(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.AccountZerotrustSubnetListParams{
			AddressFamily:    cfrex.F(cfrex.AddressFamilyV4),
			Comment:          cfrex.F("example%20comment"),
			ExistedAt:        cfrex.F("2019-10-12T07%3A20%3A50.52Z"),
			IsDefaultNetwork: cfrex.F(true),
			IsDeleted:        cfrex.F(true),
			Name:             cfrex.F("IPv4%20Cloudflare%20Source%20IPs"),
			Network:          cfrex.F("172.16.0.0%2F16"),
			Page:             cfrex.F(1.000000),
			PerPage:          cfrex.F(1.000000),
			SortOrder:        cfrex.F(cfrex.AccountZerotrustSubnetListParamsSortOrderAsc),
			SubnetTypes:      cfrex.F(cfrex.AccountZerotrustSubnetListParamsSubnetTypesCloudflareSource),
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

func TestAccountZerotrustSubnetUpdateCloudflareSourceWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Zerotrust.Subnets.UpdateCloudflareSource(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.AddressFamilyV4,
		cfrex.AccountZerotrustSubnetUpdateCloudflareSourceParams{
			Comment: cfrex.F("example comment"),
			Name:    cfrex.F("IPv4 Cloudflare Source IPs"),
			Network: cfrex.F("100.64.0.0/12"),
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
