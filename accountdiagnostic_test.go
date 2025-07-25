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

func TestAccountDiagnosticRunTracerouteWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Diagnostics.RunTraceroute(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountDiagnosticRunTracerouteParams{
			Targets: cfrex.F([]string{"203.0.113.1", "cloudflare.com"}),
			Colos:   cfrex.F([]string{"den", "sin"}),
			Options: cfrex.F(cfrex.AccountDiagnosticRunTracerouteParamsOptions{
				MaxTtl:        cfrex.F(int64(15)),
				PacketType:    cfrex.F(cfrex.AccountDiagnosticRunTracerouteParamsOptionsPacketTypeIcmp),
				PacketsPerTtl: cfrex.F(int64(0)),
				Port:          cfrex.F(int64(0)),
				WaitTime:      cfrex.F(int64(1)),
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
