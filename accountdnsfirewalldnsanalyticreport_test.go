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

func TestAccountDNSFirewallDNSAnalyticReportGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.DNSFirewall.DNSAnalytics.Report.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountDNSFirewallDNSAnalyticReportGetParams{
			Dimensions: cfrex.F("queryType"),
			Filters:    cfrex.F("responseCode==NOERROR,queryType==A"),
			Limit:      cfrex.F(int64(100)),
			Metrics:    cfrex.F("queryCount,uncachedCount"),
			Since:      cfrex.F(time.Now()),
			Sort:       cfrex.F("+responseCode,-queryName"),
			Until:      cfrex.F(time.Now()),
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

func TestAccountDNSFirewallDNSAnalyticReportListByTimeWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.DNSFirewall.DNSAnalytics.Report.ListByTime(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountDNSFirewallDNSAnalyticReportListByTimeParams{
			Dimensions: cfrex.F("queryType"),
			Filters:    cfrex.F("responseCode==NOERROR,queryType==A"),
			Limit:      cfrex.F(int64(100)),
			Metrics:    cfrex.F("queryCount,uncachedCount"),
			Since:      cfrex.F(time.Now()),
			Sort:       cfrex.F("+responseCode,-queryName"),
			TimeDelta:  cfrex.F(cfrex.TimeDeltaAll),
			Until:      cfrex.F(time.Now()),
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
