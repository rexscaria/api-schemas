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

func TestAccountInfrastructureTargetBatchNew(t *testing.T) {
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
	_, err := client.Accounts.Infrastructure.Targets.Batch.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountInfrastructureTargetBatchNewParams{
			Body: []cfrex.AccountInfrastructureTargetBatchNewParamsBody{{
				Hostname: cfrex.F("infra-access-target"),
				IP: cfrex.F(cfrex.IPInfoTargetParam{
					Ipv4: cfrex.F(cfrex.IPInfoTargetIpv4Param{
						IPAddr:           cfrex.F("187.26.29.249"),
						VirtualNetworkID: cfrex.F("c77b744e-acc8-428f-9257-6878c046ed55"),
					}),
					Ipv6: cfrex.F(cfrex.IPInfoTargetIpv6Param{
						IPAddr:           cfrex.F("64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0"),
						VirtualNetworkID: cfrex.F("c77b744e-acc8-428f-9257-6878c046ed55"),
					}),
				}),
			}},
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

func TestAccountInfrastructureTargetBatchDelete(t *testing.T) {
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
	err := client.Accounts.Infrastructure.Targets.Batch.Delete(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
