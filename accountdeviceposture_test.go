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

func TestAccountDevicePostureNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Devices.Posture.New(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.AccountDevicePostureNewParams{
			Name:        cfrex.F("Admin Serial Numbers"),
			Type:        cfrex.F(cfrex.DeviceTypePostureRuleFile),
			Description: cfrex.F("The rule for admin serial numbers"),
			Expiration:  cfrex.F("1h"),
			Input: cfrex.F[cfrex.InputValueUnionParam](cfrex.InputValueTeamsDevicesFileInputRequestParam{
				OperatingSystem: cfrex.F(cfrex.InputValueTeamsDevicesFileInputRequestOperatingSystemLinux),
				Path:            cfrex.F("/bin/cat"),
				Exists:          cfrex.F(true),
				Sha256:          cfrex.F("https://api.us-2.crowdstrike.com"),
				Thumbprint:      cfrex.F("0aabab210bdb998e9cf45da2c9ce352977ab531c681b74cf1e487be1bbe9fe6e"),
			}),
			Match: cfrex.F([]cfrex.MatchItemParam{{
				Platform: cfrex.F(cfrex.PlatformWindows),
			}}),
			Schedule: cfrex.F("1h"),
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

func TestAccountDevicePostureGet(t *testing.T) {
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
	_, err := client.Accounts.Devices.Posture.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
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

func TestAccountDevicePostureUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Devices.Posture.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.AccountDevicePostureUpdateParams{
			Name:        cfrex.F("Admin Serial Numbers"),
			Type:        cfrex.F(cfrex.DeviceTypePostureRuleFile),
			Description: cfrex.F("The rule for admin serial numbers"),
			Expiration:  cfrex.F("1h"),
			Input: cfrex.F[cfrex.InputValueUnionParam](cfrex.InputValueTeamsDevicesFileInputRequestParam{
				OperatingSystem: cfrex.F(cfrex.InputValueTeamsDevicesFileInputRequestOperatingSystemLinux),
				Path:            cfrex.F("/bin/cat"),
				Exists:          cfrex.F(true),
				Sha256:          cfrex.F("https://api.us-2.crowdstrike.com"),
				Thumbprint:      cfrex.F("0aabab210bdb998e9cf45da2c9ce352977ab531c681b74cf1e487be1bbe9fe6e"),
			}),
			Match: cfrex.F([]cfrex.MatchItemParam{{
				Platform: cfrex.F(cfrex.PlatformWindows),
			}}),
			Schedule: cfrex.F("1h"),
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

func TestAccountDevicePostureList(t *testing.T) {
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
	_, err := client.Accounts.Devices.Posture.List(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDevicePostureDelete(t *testing.T) {
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
	_, err := client.Accounts.Devices.Posture.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
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
