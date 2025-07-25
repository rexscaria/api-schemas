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

func TestAccountMagicSiteACLNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.Sites.ACLs.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountMagicSiteACLNewParams{
			Lan1: cfrex.F(cfrex.MagicLanACLConfigurationParam{
				LanID:      cfrex.F("lan_id"),
				LanName:    cfrex.F("lan_name"),
				PortRanges: cfrex.F([]string{"8080-9000"}),
				Ports:      cfrex.F([]int64{int64(1)}),
				Subnets:    cfrex.F([]string{"192.0.2.1"}),
			}),
			Lan2: cfrex.F(cfrex.MagicLanACLConfigurationParam{
				LanID:      cfrex.F("lan_id"),
				LanName:    cfrex.F("lan_name"),
				PortRanges: cfrex.F([]string{"8080-9000"}),
				Ports:      cfrex.F([]int64{int64(1)}),
				Subnets:    cfrex.F([]string{"192.0.2.1"}),
			}),
			Name:           cfrex.F("PIN Pad - Cash Register"),
			Description:    cfrex.F("Allows local traffic between PIN pads and cash register."),
			ForwardLocally: cfrex.F(true),
			Protocols:      cfrex.F([]cfrex.AccountMagicSiteACLNewParamsProtocol{cfrex.AccountMagicSiteACLNewParamsProtocolTcp}),
			Unidirectional: cfrex.F(true),
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

func TestAccountMagicSiteACLGet(t *testing.T) {
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
	_, err := client.Accounts.Magic.Sites.ACLs.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
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

func TestAccountMagicSiteACLUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.Sites.ACLs.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountMagicSiteACLUpdateParams{
			MagicACLUpdateRequest: cfrex.MagicACLUpdateRequestParam{
				Description:    cfrex.F("Allows local traffic between PIN pads and cash register."),
				ForwardLocally: cfrex.F(true),
				Lan1: cfrex.F(cfrex.MagicLanACLConfigurationParam{
					LanID:      cfrex.F("lan_id"),
					LanName:    cfrex.F("lan_name"),
					PortRanges: cfrex.F([]string{"8080-9000"}),
					Ports:      cfrex.F([]int64{int64(1)}),
					Subnets:    cfrex.F([]string{"192.0.2.1"}),
				}),
				Lan2: cfrex.F(cfrex.MagicLanACLConfigurationParam{
					LanID:      cfrex.F("lan_id"),
					LanName:    cfrex.F("lan_name"),
					PortRanges: cfrex.F([]string{"8080-9000"}),
					Ports:      cfrex.F([]int64{int64(1)}),
					Subnets:    cfrex.F([]string{"192.0.2.1"}),
				}),
				Name:           cfrex.F("PIN Pad - Cash Register"),
				Protocols:      cfrex.F([]cfrex.MagicACLUpdateRequestProtocol{cfrex.MagicACLUpdateRequestProtocolTcp}),
				Unidirectional: cfrex.F(true),
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

func TestAccountMagicSiteACLList(t *testing.T) {
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
	_, err := client.Accounts.Magic.Sites.ACLs.List(
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

func TestAccountMagicSiteACLDelete(t *testing.T) {
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
	_, err := client.Accounts.Magic.Sites.ACLs.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
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

func TestAccountMagicSiteACLPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Magic.Sites.ACLs.Patch(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountMagicSiteACLPatchParams{
			MagicACLUpdateRequest: cfrex.MagicACLUpdateRequestParam{
				Description:    cfrex.F("Allows local traffic between PIN pads and cash register."),
				ForwardLocally: cfrex.F(true),
				Lan1: cfrex.F(cfrex.MagicLanACLConfigurationParam{
					LanID:      cfrex.F("lan_id"),
					LanName:    cfrex.F("lan_name"),
					PortRanges: cfrex.F([]string{"8080-9000"}),
					Ports:      cfrex.F([]int64{int64(1)}),
					Subnets:    cfrex.F([]string{"192.0.2.1"}),
				}),
				Lan2: cfrex.F(cfrex.MagicLanACLConfigurationParam{
					LanID:      cfrex.F("lan_id"),
					LanName:    cfrex.F("lan_name"),
					PortRanges: cfrex.F([]string{"8080-9000"}),
					Ports:      cfrex.F([]int64{int64(1)}),
					Subnets:    cfrex.F([]string{"192.0.2.1"}),
				}),
				Name:           cfrex.F("PIN Pad - Cash Register"),
				Protocols:      cfrex.F([]cfrex.MagicACLUpdateRequestProtocol{cfrex.MagicACLUpdateRequestProtocolTcp}),
				Unidirectional: cfrex.F(true),
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
