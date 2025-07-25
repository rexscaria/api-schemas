// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/rexscaria/api-schemas"
	"github.com/rexscaria/api-schemas/internal/testutil"
	"github.com/rexscaria/api-schemas/option"
)

func TestZoneSecurityCenterSecuritytxtGet(t *testing.T) {
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
	_, err := client.Zones.SecurityCenter.Securitytxt.Get(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneSecurityCenterSecuritytxtUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.SecurityCenter.Securitytxt.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneSecurityCenterSecuritytxtUpdateParams{
			SecurityTxt: cfrex.SecurityTxtParam{
				Acknowledgments:    cfrex.F([]string{"https://example.com/hall-of-fame.html"}),
				Canonical:          cfrex.F([]string{"https://www.example.com/.well-known/security.txt"}),
				Contact:            cfrex.F([]string{"mailto:security@example.com", "tel:+1-201-555-0123", "https://example.com/security-contact.html"}),
				Enabled:            cfrex.F(true),
				Encryption:         cfrex.F([]string{"https://example.com/pgp-key.txt", "dns:5d2d37ab76d47d36._openpgpkey.example.com?type=OPENPGPKEY", "openpgp4fpr:5f2de5521c63a801ab59ccb603d49de44b29100f"}),
				Expires:            cfrex.F(time.Now()),
				Hiring:             cfrex.F([]string{"https://example.com/jobs.html"}),
				Policy:             cfrex.F([]string{"https://example.com/disclosure-policy.html"}),
				PreferredLanguages: cfrex.F("en, es, fr"),
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

func TestZoneSecurityCenterSecuritytxtDelete(t *testing.T) {
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
	_, err := client.Zones.SecurityCenter.Securitytxt.Delete(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
