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

func TestZoneAccessGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Access.Groups.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneAccessGroupNewParams{
			Include: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
				Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
					ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
				}),
			}}),
			Name: cfrex.F("Allow devs"),
			Exclude: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
				Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
					ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
				}),
			}}),
			Require: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
				Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
					ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
				}),
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

func TestZoneAccessGroupGet(t *testing.T) {
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
	_, err := client.Zones.Access.Groups.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneAccessGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Access.Groups.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.ZoneAccessGroupUpdateParams{
			Include: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
				Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
					ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
				}),
			}}),
			Name: cfrex.F("Allow devs"),
			Exclude: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
				Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
					ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
				}),
			}}),
			Require: cfrex.F([]cfrex.AccessRuleUnionParam{cfrex.AccessRuleAccessAccessGroupRuleParam{
				Group: cfrex.F(cfrex.AccessRuleAccessAccessGroupRuleGroupParam{
					ID: cfrex.F("aa0a4aab-672b-4bdb-bc33-a59f1130a11f"),
				}),
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

func TestZoneAccessGroupList(t *testing.T) {
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
	_, err := client.Zones.Access.Groups.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneAccessGroupDelete(t *testing.T) {
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
	_, err := client.Zones.Access.Groups.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
