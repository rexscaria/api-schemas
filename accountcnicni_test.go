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

func TestAccountCniCniNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Cni.Cnis.New(
		context.TODO(),
		"account_id",
		cfrex.AccountCniCniNewParams{
			Account:      cfrex.F("account"),
			Interconnect: cfrex.F("interconnect"),
			Magic: cfrex.F(cfrex.NscMagicSettingsParam{
				ConduitName: cfrex.F("conduit_name"),
				Description: cfrex.F("description"),
				Mtu:         cfrex.F(int64(0)),
			}),
			Bgp: cfrex.F(cfrex.NscBgpControlParam{
				CustomerAsn:   cfrex.F(int64(0)),
				ExtraPrefixes: cfrex.F([]string{"string"}),
				Md5Key:        cfrex.F("md5_key"),
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

func TestAccountCniCniGet(t *testing.T) {
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
	_, err := client.Accounts.Cni.Cnis.Get(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountCniCniUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Cni.Cnis.Update(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cfrex.AccountCniCniUpdateParams{
			NscCni: cfrex.NscCniParam{
				ID:           cfrex.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Account:      cfrex.F("account"),
				CustIP:       cfrex.F("192.168.3.4/31"),
				Interconnect: cfrex.F("interconnect"),
				Magic: cfrex.F(cfrex.NscMagicSettingsParam{
					ConduitName: cfrex.F("conduit_name"),
					Description: cfrex.F("description"),
					Mtu:         cfrex.F(int64(0)),
				}),
				P2pIP: cfrex.F("192.168.3.4/31"),
				Bgp: cfrex.F(cfrex.NscBgpControlParam{
					CustomerAsn:   cfrex.F(int64(0)),
					ExtraPrefixes: cfrex.F([]string{"string"}),
					Md5Key:        cfrex.F("md5_key"),
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

func TestAccountCniCniListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Cni.Cnis.List(
		context.TODO(),
		"account_id",
		cfrex.AccountCniCniListParams{
			Cursor: cfrex.F(int64(0)),
			Limit:  cfrex.F(int64(0)),
			Slot:   cfrex.F("slot"),
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

func TestAccountCniCniDelete(t *testing.T) {
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
	err := client.Accounts.Cni.Cnis.Delete(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
