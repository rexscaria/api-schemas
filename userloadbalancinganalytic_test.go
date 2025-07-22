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

func TestUserLoadBalancingAnalyticListEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.User.LoadBalancingAnalytics.ListEvents(context.TODO(), cfrex.UserLoadBalancingAnalyticListEventsParams{
		OriginHealthy: cfrex.F(true),
		OriginName:    cfrex.F("primary-dc-1"),
		PoolHealthy:   cfrex.F(true),
		PoolID:        cfrex.F("17b5962d775c646f3f9725cbc7a53df4"),
		PoolName:      cfrex.F("primary-dc"),
		Since:         cfrex.F(time.Now()),
		Until:         cfrex.F(time.Now()),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
