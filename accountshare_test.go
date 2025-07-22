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

func TestAccountShareNew(t *testing.T) {
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
	_, err := client.Accounts.Shares.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountShareNewParams{
			Name: cfrex.F("My Shared WAF Managed Rule"),
			Recipients: cfrex.F([]cfrex.CreateShareRecipientRequestParam{{
				AccountID:      cfrex.F("023e105f4ecef8ad9ca31a8372d0c353"),
				OrganizationID: cfrex.F("023e105f4ecef8ad9ca31a8372d0c353"),
			}}),
			Resources: cfrex.F([]cfrex.CreateShareResourceRequestParam{{
				Meta:              cfrex.F[any](map[string]interface{}{}),
				ResourceAccountID: cfrex.F("023e105f4ecef8ad9ca31a8372d0c353"),
				ResourceID:        cfrex.F("023e105f4ecef8ad9ca31a8372d0c353"),
				ResourceType:      cfrex.F(cfrex.ResourceTypeCustomRuleset),
			}}),
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

func TestAccountShareGet(t *testing.T) {
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
	_, err := client.Accounts.Shares.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"3fd85f74b32742f1bff64a85009dda07",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountShareUpdate(t *testing.T) {
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
	_, err := client.Accounts.Shares.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"3fd85f74b32742f1bff64a85009dda07",
		cfrex.AccountShareUpdateParams{
			Name: cfrex.F("My Shared WAF Managed Rule"),
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

func TestAccountShareListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Shares.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountShareListParams{
			Direction:  cfrex.F(cfrex.AccountShareListParamsDirectionAsc),
			Kind:       cfrex.F(cfrex.AccountShareListParamsKindSent),
			Order:      cfrex.F(cfrex.AccountShareListParamsOrderName),
			Page:       cfrex.F(int64(2)),
			PerPage:    cfrex.F(int64(20)),
			Status:     cfrex.F(cfrex.AccountShareListParamsStatusActive),
			TargetType: cfrex.F(cfrex.AccountShareListParamsTargetTypeAccount),
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

func TestAccountShareDelete(t *testing.T) {
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
	_, err := client.Accounts.Shares.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"3fd85f74b32742f1bff64a85009dda07",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
