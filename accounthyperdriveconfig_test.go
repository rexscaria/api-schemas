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

func TestAccountHyperdriveConfigNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Hyperdrive.Configs.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountHyperdriveConfigNewParams{
			HyperdriveHyperdriveConfig: cfrex.HyperdriveHyperdriveConfigParam{
				Name: cfrex.F("example-hyperdrive"),
				Origin: cfrex.F[cfrex.HyperdriveHyperdriveConfigOriginUnionParam](cfrex.HyperdriveHyperdriveConfigOriginPublicDatabaseParam(cfrex.HyperdriveHyperdriveConfigOriginPublicDatabaseParam{
					HyperdriveHyperdriveDatabaseFullParam: cfrex.HyperdriveHyperdriveDatabaseFullParam{
						Database: cfrex.F("postgres"),
						Password: cfrex.F("password"),
						Scheme:   cfrex.F(cfrex.HyperdriveHyperdriveDatabaseFullSchemePostgres),
						User:     cfrex.F("postgres"),
					},
					HyperdriveInternetOriginParam: cfrex.HyperdriveInternetOriginParam{
						Host: cfrex.F("database.example.com"),
						Port: cfrex.F(int64(0)),
					},
				})),
				Caching: cfrex.F[cfrex.HyperdriveHyperdriveCachingUnionParam](cfrex.HyperdriveHyperdriveCachingCommonParam{
					Disabled: cfrex.F(true),
				}),
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

func TestAccountHyperdriveConfigGet(t *testing.T) {
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
	_, err := client.Accounts.Hyperdrive.Configs.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountHyperdriveConfigUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Hyperdrive.Configs.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountHyperdriveConfigUpdateParams{
			HyperdriveHyperdriveConfig: cfrex.HyperdriveHyperdriveConfigParam{
				Name: cfrex.F("example-hyperdrive"),
				Origin: cfrex.F[cfrex.HyperdriveHyperdriveConfigOriginUnionParam](cfrex.HyperdriveHyperdriveConfigOriginPublicDatabaseParam(cfrex.HyperdriveHyperdriveConfigOriginPublicDatabaseParam{
					HyperdriveHyperdriveDatabaseFullParam: cfrex.HyperdriveHyperdriveDatabaseFullParam{
						Database: cfrex.F("postgres"),
						Password: cfrex.F("password"),
						Scheme:   cfrex.F(cfrex.HyperdriveHyperdriveDatabaseFullSchemePostgres),
						User:     cfrex.F("postgres"),
					},
					HyperdriveInternetOriginParam: cfrex.HyperdriveInternetOriginParam{
						Host: cfrex.F("database.example.com"),
						Port: cfrex.F(int64(0)),
					},
				})),
				Caching: cfrex.F[cfrex.HyperdriveHyperdriveCachingUnionParam](cfrex.HyperdriveHyperdriveCachingCommonParam{
					Disabled: cfrex.F(true),
				}),
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

func TestAccountHyperdriveConfigList(t *testing.T) {
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
	_, err := client.Accounts.Hyperdrive.Configs.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountHyperdriveConfigDelete(t *testing.T) {
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
	_, err := client.Accounts.Hyperdrive.Configs.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountHyperdriveConfigPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Hyperdrive.Configs.Patch(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountHyperdriveConfigPatchParams{
			Caching: cfrex.F[cfrex.HyperdriveHyperdriveCachingUnionParam](cfrex.HyperdriveHyperdriveCachingCommonParam{
				Disabled: cfrex.F(true),
			}),
			Name: cfrex.F("example-hyperdrive"),
			Origin: cfrex.F[cfrex.AccountHyperdriveConfigPatchParamsOriginUnion](cfrex.HyperdriveHyperdriveDatabaseParam{
				Database: cfrex.F("postgres"),
				Password: cfrex.F("password"),
				Scheme:   cfrex.F(cfrex.HyperdriveHyperdriveDatabaseSchemePostgres),
				User:     cfrex.F("postgres"),
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
