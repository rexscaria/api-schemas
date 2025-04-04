// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/cf-rex-go"
	"github.com/stainless-sdks/cf-rex-go/internal/testutil"
	"github.com/stainless-sdks/cf-rex-go/option"
)

func TestRadarRankingGetDomainRankWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Ranking.GetDomainRank(
		context.TODO(),
		"google.com",
		cfrex.RadarRankingGetDomainRankParams{
			Date:                cfrex.F([]time.Time{time.Now()}),
			Format:              cfrex.F(cfrex.RadarRankingGetDomainRankParamsFormatJson),
			IncludeTopLocations: cfrex.F(true),
			Limit:               cfrex.F(int64(5)),
			Name:                cfrex.F([]string{"main_series"}),
			RankingType:         cfrex.F(cfrex.RadarRankingGetDomainRankParamsRankingTypePopular),
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

func TestRadarRankingGetTimeseriesGroupsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Ranking.GetTimeseriesGroups(context.TODO(), cfrex.RadarRankingGetTimeseriesGroupsParams{
		DateEnd:        cfrex.F([]time.Time{time.Now()}),
		DateRange:      cfrex.F([]string{"7d"}),
		DateStart:      cfrex.F([]time.Time{time.Now()}),
		DomainCategory: cfrex.F([]string{"string"}),
		Domains:        cfrex.F([]string{"sfNk.l.iJR-Bu.u9JV-.3.h.U.s.l4b-Jdq-.zk-H-8ASjA2.1vKx-qF-3v.0-NI-5P54N-n.keyk"}),
		Format:         cfrex.F(cfrex.RadarRankingGetTimeseriesGroupsParamsFormatJson),
		Limit:          cfrex.F(int64(5)),
		Location:       cfrex.F([]string{"string"}),
		Name:           cfrex.F([]string{"main_series"}),
		RankingType:    cfrex.F(cfrex.RadarRankingGetTimeseriesGroupsParamsRankingTypePopular),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRadarRankingGetTopDomainsWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Ranking.GetTopDomains(context.TODO(), cfrex.RadarRankingGetTopDomainsParams{
		Date:           cfrex.F([]time.Time{time.Now()}),
		DomainCategory: cfrex.F([]string{"string"}),
		Format:         cfrex.F(cfrex.RadarRankingGetTopDomainsParamsFormatJson),
		Limit:          cfrex.F(int64(5)),
		Location:       cfrex.F([]string{"string"}),
		Name:           cfrex.F([]string{"main_series"}),
		RankingType:    cfrex.F(cfrex.RadarRankingGetTopDomainsParamsRankingTypePopular),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
