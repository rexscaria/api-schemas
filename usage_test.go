// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"context"
	"os"
	"testing"

	"github.com/rexscaria/api-schemas"
	"github.com/rexscaria/api-schemas/internal/testutil"
	"github.com/rexscaria/api-schemas/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.Accounts.Access.UpdateSeats(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.AccountAccessUpdateSeatsParams{
			Body: []cfrex.AccountAccessUpdateSeatsParamsBody{{
				AccessSeat:  cfrex.F(false),
				GatewaySeat: cfrex.F(false),
				SeatUid:     cfrex.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
			}},
		},
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", response.Errors)
}
