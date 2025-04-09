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

func TestZonePageruleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Pagerules.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZonePageruleNewParams{
			Actions: cfrex.F([]cfrex.ZoneActionUnionParam{cfrex.ZoneActionBrowserCheckParam{
				ID:    cfrex.F(cfrex.ZoneActionBrowserCheckIDBrowserCheck),
				Value: cfrex.F(cfrex.ZoneActionBrowserCheckValueOn),
			}}),
			Targets: cfrex.F([]cfrex.RequestConditionTargetParam{{
				Constraint: cfrex.F(cfrex.RequestConditionTargetConstraintParam{
					Operator: cfrex.F(cfrex.RequestConditionTargetConstraintOperatorMatches),
					Value:    cfrex.F("*example.com/images/*"),
				}),
				Target: cfrex.F(cfrex.RequestConditionTargetTargetURL),
			}}),
			Priority: cfrex.F(int64(0)),
			Status:   cfrex.F(cfrex.PageRuleStatusActive),
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

func TestZonePageruleGet(t *testing.T) {
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
	_, err := client.Zones.Pagerules.Get(
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

func TestZonePageruleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Pagerules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZonePageruleUpdateParams{
			Actions: cfrex.F([]cfrex.ZoneActionUnionParam{cfrex.ZoneActionBrowserCheckParam{
				ID:    cfrex.F(cfrex.ZoneActionBrowserCheckIDBrowserCheck),
				Value: cfrex.F(cfrex.ZoneActionBrowserCheckValueOn),
			}}),
			Targets: cfrex.F([]cfrex.RequestConditionTargetParam{{
				Constraint: cfrex.F(cfrex.RequestConditionTargetConstraintParam{
					Operator: cfrex.F(cfrex.RequestConditionTargetConstraintOperatorMatches),
					Value:    cfrex.F("*example.com/images/*"),
				}),
				Target: cfrex.F(cfrex.RequestConditionTargetTargetURL),
			}}),
			Priority: cfrex.F(int64(0)),
			Status:   cfrex.F(cfrex.PageRuleStatusActive),
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

func TestZonePageruleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Pagerules.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZonePageruleListParams{
			Direction: cfrex.F(cfrex.ZonePageruleListParamsDirectionDesc),
			Match:     cfrex.F(cfrex.ZonePageruleListParamsMatchAny),
			Order:     cfrex.F(cfrex.ZonePageruleListParamsOrderStatus),
			Status:    cfrex.F(cfrex.ZonePageruleListParamsStatusActive),
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

func TestZonePageruleDelete(t *testing.T) {
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
	_, err := client.Zones.Pagerules.Delete(
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

func TestZonePageruleEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Pagerules.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZonePageruleEditParams{
			Actions: cfrex.F([]cfrex.ZoneActionUnionParam{cfrex.ZoneActionBrowserCheckParam{
				ID:    cfrex.F(cfrex.ZoneActionBrowserCheckIDBrowserCheck),
				Value: cfrex.F(cfrex.ZoneActionBrowserCheckValueOn),
			}}),
			Priority: cfrex.F(int64(0)),
			Status:   cfrex.F(cfrex.PageRuleStatusActive),
			Targets: cfrex.F([]cfrex.RequestConditionTargetParam{{
				Constraint: cfrex.F(cfrex.RequestConditionTargetConstraintParam{
					Operator: cfrex.F(cfrex.RequestConditionTargetConstraintOperatorMatches),
					Value:    cfrex.F("*example.com/images/*"),
				}),
				Target: cfrex.F(cfrex.RequestConditionTargetTargetURL),
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

func TestZonePageruleListSettings(t *testing.T) {
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
	_, err := client.Zones.Pagerules.ListSettings(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
