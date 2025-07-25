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

func TestAccountAddressingPrefixBgpPrefixNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Addressing.Prefixes.Bgp.Prefixes.New(
		context.TODO(),
		"258def64c72dae45f3e4c8516e2111f2",
		"2af39739cc4e3b5910c918468bb89828",
		cfrex.AccountAddressingPrefixBgpPrefixNewParams{
			Cidr: cfrex.F("192.0.2.0/24"),
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

func TestAccountAddressingPrefixBgpPrefixGet(t *testing.T) {
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
	_, err := client.Accounts.Addressing.Prefixes.Bgp.Prefixes.Get(
		context.TODO(),
		"258def64c72dae45f3e4c8516e2111f2",
		"2af39739cc4e3b5910c918468bb89828",
		"7009ba364c7a5760798ceb430e603b74",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountAddressingPrefixBgpPrefixUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Addressing.Prefixes.Bgp.Prefixes.Update(
		context.TODO(),
		"258def64c72dae45f3e4c8516e2111f2",
		"2af39739cc4e3b5910c918468bb89828",
		"7009ba364c7a5760798ceb430e603b74",
		cfrex.AccountAddressingPrefixBgpPrefixUpdateParams{
			AsnPrependCount:       cfrex.F(int64(2)),
			AutoAdvertiseWithdraw: cfrex.F(true),
			OnDemand: cfrex.F(cfrex.AccountAddressingPrefixBgpPrefixUpdateParamsOnDemand{
				Advertised: cfrex.F(true),
			}),
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

func TestAccountAddressingPrefixBgpPrefixList(t *testing.T) {
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
	_, err := client.Accounts.Addressing.Prefixes.Bgp.Prefixes.List(
		context.TODO(),
		"258def64c72dae45f3e4c8516e2111f2",
		"2af39739cc4e3b5910c918468bb89828",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
