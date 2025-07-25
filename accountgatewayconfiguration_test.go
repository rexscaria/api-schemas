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

func TestAccountGatewayConfigurationGet(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Configuration.Get(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountGatewayConfigurationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Configuration.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.AccountGatewayConfigurationUpdateParams{
			GatewayAccountSettings: cfrex.GatewayAccountSettingsParam{
				Settings: cfrex.F(cfrex.GatewayAccountSettingsSettingsParam{
					ActivityLog: cfrex.F(cfrex.GatewayAccountSettingsSettingsActivityLogParam{
						Enabled: cfrex.F(true),
					}),
					Antivirus: cfrex.F(cfrex.GatewayAccountSettingsSettingsAntivirusParam{
						EnabledDownloadPhase: cfrex.F(false),
						EnabledUploadPhase:   cfrex.F(false),
						FailClosed:           cfrex.F(false),
						NotificationSettings: cfrex.F(cfrex.GatewayAccountSettingsSettingsAntivirusNotificationSettingsParam{
							Enabled:        cfrex.F(true),
							IncludeContext: cfrex.F(true),
							Msg:            cfrex.F("msg"),
							SupportURL:     cfrex.F("support_url"),
						}),
					}),
					BlockPage: cfrex.F(cfrex.GatewayAccountSettingsSettingsBlockPageParam{
						Enabled:         cfrex.F(true),
						Mode:            cfrex.F(cfrex.GatewayAccountSettingsSettingsBlockPageModeCustomizedBlockPage),
						BackgroundColor: cfrex.F("background_color"),
						FooterText:      cfrex.F("--footer--"),
						HeaderText:      cfrex.F("--header--"),
						IncludeContext:  cfrex.F(true),
						LogoPath:        cfrex.F("https://logos.com/a.png"),
						MailtoAddress:   cfrex.F("admin@example.com"),
						MailtoSubject:   cfrex.F("Blocked User Inquiry"),
						Name:            cfrex.F("Cloudflare"),
						SuppressFooter:  cfrex.F(false),
						TargetUri:       cfrex.F("https://example.com"),
					}),
					BodyScanning: cfrex.F(cfrex.GatewayAccountSettingsSettingsBodyScanningParam{
						InspectionMode: cfrex.F(cfrex.GatewayAccountSettingsSettingsBodyScanningInspectionModeDeep),
					}),
					BrowserIsolation: cfrex.F(cfrex.GatewayAccountSettingsSettingsBrowserIsolationParam{
						NonIdentityEnabled:         cfrex.F(true),
						URLBrowserIsolationEnabled: cfrex.F(true),
					}),
					Certificate: cfrex.F(cfrex.GatewayAccountSettingsSettingsCertificateParam{
						ID: cfrex.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
					}),
					CustomCertificate: cfrex.F(cfrex.CustomCertificateSettingsParam{
						Enabled: cfrex.F(true),
						ID:      cfrex.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
					}),
					ExtendedEmailMatching: cfrex.F(cfrex.GatewayAccountSettingsSettingsExtendedEmailMatchingParam{
						Enabled: cfrex.F(true),
					}),
					Fips: cfrex.F(cfrex.GatewayAccountSettingsSettingsFipsParam{
						Tls: cfrex.F(true),
					}),
					HostSelector: cfrex.F(cfrex.GatewayAccountSettingsSettingsHostSelectorParam{
						Enabled: cfrex.F(false),
					}),
					Inspection: cfrex.F(cfrex.GatewayAccountSettingsSettingsInspectionParam{
						Mode: cfrex.F(cfrex.GatewayAccountSettingsSettingsInspectionModeStatic),
					}),
					ProtocolDetection: cfrex.F(cfrex.GatewayAccountSettingsSettingsProtocolDetectionParam{
						Enabled: cfrex.F(true),
					}),
					Sandbox: cfrex.F(cfrex.GatewayAccountSettingsSettingsSandboxParam{
						Enabled:        cfrex.F(true),
						FallbackAction: cfrex.F(cfrex.GatewayAccountSettingsSettingsSandboxFallbackActionAllow),
					}),
					TlsDecrypt: cfrex.F(cfrex.GatewayAccountSettingsSettingsTlsDecryptParam{
						Enabled: cfrex.F(true),
					}),
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

func TestAccountGatewayConfigurationPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Configuration.Patch(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.AccountGatewayConfigurationPatchParams{
			GatewayAccountSettings: cfrex.GatewayAccountSettingsParam{
				Settings: cfrex.F(cfrex.GatewayAccountSettingsSettingsParam{
					ActivityLog: cfrex.F(cfrex.GatewayAccountSettingsSettingsActivityLogParam{
						Enabled: cfrex.F(true),
					}),
					Antivirus: cfrex.F(cfrex.GatewayAccountSettingsSettingsAntivirusParam{
						EnabledDownloadPhase: cfrex.F(false),
						EnabledUploadPhase:   cfrex.F(false),
						FailClosed:           cfrex.F(false),
						NotificationSettings: cfrex.F(cfrex.GatewayAccountSettingsSettingsAntivirusNotificationSettingsParam{
							Enabled:        cfrex.F(true),
							IncludeContext: cfrex.F(true),
							Msg:            cfrex.F("msg"),
							SupportURL:     cfrex.F("support_url"),
						}),
					}),
					BlockPage: cfrex.F(cfrex.GatewayAccountSettingsSettingsBlockPageParam{
						Enabled:         cfrex.F(true),
						Mode:            cfrex.F(cfrex.GatewayAccountSettingsSettingsBlockPageModeCustomizedBlockPage),
						BackgroundColor: cfrex.F("background_color"),
						FooterText:      cfrex.F("--footer--"),
						HeaderText:      cfrex.F("--header--"),
						IncludeContext:  cfrex.F(true),
						LogoPath:        cfrex.F("https://logos.com/a.png"),
						MailtoAddress:   cfrex.F("admin@example.com"),
						MailtoSubject:   cfrex.F("Blocked User Inquiry"),
						Name:            cfrex.F("Cloudflare"),
						SuppressFooter:  cfrex.F(false),
						TargetUri:       cfrex.F("https://example.com"),
					}),
					BodyScanning: cfrex.F(cfrex.GatewayAccountSettingsSettingsBodyScanningParam{
						InspectionMode: cfrex.F(cfrex.GatewayAccountSettingsSettingsBodyScanningInspectionModeDeep),
					}),
					BrowserIsolation: cfrex.F(cfrex.GatewayAccountSettingsSettingsBrowserIsolationParam{
						NonIdentityEnabled:         cfrex.F(true),
						URLBrowserIsolationEnabled: cfrex.F(true),
					}),
					Certificate: cfrex.F(cfrex.GatewayAccountSettingsSettingsCertificateParam{
						ID: cfrex.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
					}),
					CustomCertificate: cfrex.F(cfrex.CustomCertificateSettingsParam{
						Enabled: cfrex.F(true),
						ID:      cfrex.F("d1b364c5-1311-466e-a194-f0e943e0799f"),
					}),
					ExtendedEmailMatching: cfrex.F(cfrex.GatewayAccountSettingsSettingsExtendedEmailMatchingParam{
						Enabled: cfrex.F(true),
					}),
					Fips: cfrex.F(cfrex.GatewayAccountSettingsSettingsFipsParam{
						Tls: cfrex.F(true),
					}),
					HostSelector: cfrex.F(cfrex.GatewayAccountSettingsSettingsHostSelectorParam{
						Enabled: cfrex.F(false),
					}),
					Inspection: cfrex.F(cfrex.GatewayAccountSettingsSettingsInspectionParam{
						Mode: cfrex.F(cfrex.GatewayAccountSettingsSettingsInspectionModeStatic),
					}),
					ProtocolDetection: cfrex.F(cfrex.GatewayAccountSettingsSettingsProtocolDetectionParam{
						Enabled: cfrex.F(true),
					}),
					Sandbox: cfrex.F(cfrex.GatewayAccountSettingsSettingsSandboxParam{
						Enabled:        cfrex.F(true),
						FallbackAction: cfrex.F(cfrex.GatewayAccountSettingsSettingsSandboxFallbackActionAllow),
					}),
					TlsDecrypt: cfrex.F(cfrex.GatewayAccountSettingsSettingsTlsDecryptParam{
						Enabled: cfrex.F(true),
					}),
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

func TestAccountGatewayConfigurationGetCustomCertificate(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Configuration.GetCustomCertificate(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
