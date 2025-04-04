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

func TestAccountGatewayRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Rules.New(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.AccountGatewayRuleNewParams{
			Action:        cfrex.F(cfrex.ActionPerformOn),
			Name:          cfrex.F("block bad websites"),
			Description:   cfrex.F("Block bad websites based on their host name."),
			DevicePosture: cfrex.F(`any(device_posture.checks.passed[*] in {"1308749e-fcfb-4ebc-b051-fe022b632644"})`),
			Enabled:       cfrex.F(true),
			Expiration: cfrex.F(cfrex.ExpirationParam{
				ExpiresAt: cfrex.F("2014-01-01T05:20:20Z"),
				Duration:  cfrex.F(int64(10)),
				Expired:   cfrex.F(false),
			}),
			Filters:    cfrex.F([]cfrex.GatewayFilters{cfrex.GatewayFiltersHTTP}),
			Identity:   cfrex.F(`any(identity.groups.name[*] in {"finance"})`),
			Precedence: cfrex.F(int64(0)),
			RuleSettings: cfrex.F(cfrex.RuleSettingsParam{
				AddHeaders: cfrex.F(map[string]string{
					"My-Next-Header":       "string",
					"X-Custom-Header-Name": "string",
				}),
				AllowChildBypass: cfrex.F(false),
				AuditSSH: cfrex.F(cfrex.RuleSettingsAuditSSHParam{
					CommandLogging: cfrex.F(false),
				}),
				BisoAdminControls: cfrex.F(cfrex.RuleSettingsBisoAdminControlsParam{
					Copy:     cfrex.F(cfrex.RuleSettingsBisoAdminControlsCopyEnabled),
					Dcp:      cfrex.F(false),
					Dd:       cfrex.F(false),
					Dk:       cfrex.F(false),
					Download: cfrex.F(cfrex.RuleSettingsBisoAdminControlsDownloadEnabled),
					Dp:       cfrex.F(false),
					Du:       cfrex.F(false),
					Keyboard: cfrex.F(cfrex.RuleSettingsBisoAdminControlsKeyboardEnabled),
					Paste:    cfrex.F(cfrex.RuleSettingsBisoAdminControlsPasteEnabled),
					Printing: cfrex.F(cfrex.RuleSettingsBisoAdminControlsPrintingEnabled),
					Upload:   cfrex.F(cfrex.RuleSettingsBisoAdminControlsUploadEnabled),
					Version:  cfrex.F(cfrex.RuleSettingsBisoAdminControlsVersionV1),
				}),
				BlockPageEnabled: cfrex.F(true),
				BlockReason:      cfrex.F("This website is a security risk"),
				BypassParentRule: cfrex.F(false),
				CheckSession: cfrex.F(cfrex.RuleSettingsCheckSessionParam{
					Duration: cfrex.F("300s"),
					Enforce:  cfrex.F(true),
				}),
				DNSResolvers: cfrex.F(cfrex.RuleSettingsDNSResolversParam{
					Ipv4: cfrex.F([]cfrex.RuleSettingsDNSResolversIpv4Param{{
						IP:                         cfrex.F("2.2.2.2"),
						Port:                       cfrex.F(int64(5053)),
						RouteThroughPrivateNetwork: cfrex.F(true),
						VnetID:                     cfrex.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
					Ipv6: cfrex.F([]cfrex.RuleSettingsDNSResolversIpv6Param{{
						IP:                         cfrex.F("2001:DB8::"),
						Port:                       cfrex.F(int64(5053)),
						RouteThroughPrivateNetwork: cfrex.F(true),
						VnetID:                     cfrex.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
				}),
				Egress: cfrex.F(cfrex.RuleSettingsEgressParam{
					Ipv4:         cfrex.F("192.0.2.2"),
					Ipv4Fallback: cfrex.F("192.0.2.3"),
					Ipv6:         cfrex.F("2001:DB8::/64"),
				}),
				IgnoreCnameCategoryMatches:      cfrex.F(true),
				InsecureDisableDnssecValidation: cfrex.F(false),
				IPCategories:                    cfrex.F(true),
				IPIndicatorFeeds:                cfrex.F(true),
				L4override: cfrex.F(cfrex.RuleSettingsL4overrideParam{
					IP:   cfrex.F("1.1.1.1"),
					Port: cfrex.F(int64(0)),
				}),
				NotificationSettings: cfrex.F(cfrex.RuleSettingsNotificationSettingsParam{
					Enabled:    cfrex.F(true),
					Msg:        cfrex.F("msg"),
					SupportURL: cfrex.F("support_url"),
				}),
				OverrideHost: cfrex.F("example.com"),
				OverrideIPs:  cfrex.F([]string{"1.1.1.1", "2.2.2.2"}),
				PayloadLog: cfrex.F(cfrex.RuleSettingsPayloadLogParam{
					Enabled: cfrex.F(true),
				}),
				Quarantine: cfrex.F(cfrex.RuleSettingsQuarantineParam{
					FileTypes: cfrex.F([]cfrex.RuleSettingsQuarantineFileType{cfrex.RuleSettingsQuarantineFileTypeExe}),
				}),
				Redirect: cfrex.F(cfrex.RuleSettingsRedirectParam{
					TargetUri:            cfrex.F("https://example.com"),
					PreservePathAndQuery: cfrex.F(true),
				}),
				ResolveDNSInternally: cfrex.F(cfrex.RuleSettingsResolveDNSInternallyParam{
					Fallback: cfrex.F(cfrex.RuleSettingsResolveDNSInternallyFallbackNone),
					ViewID:   cfrex.F("view_id"),
				}),
				ResolveDNSThroughCloudflare: cfrex.F(true),
				UntrustedCert: cfrex.F(cfrex.RuleSettingsUntrustedCertParam{
					Action: cfrex.F(cfrex.RuleSettingsUntrustedCertActionPassThrough),
				}),
			}),
			Schedule: cfrex.F(cfrex.ScheduleParam{
				Fri:      cfrex.F("08:00-12:30,13:30-17:00"),
				Mon:      cfrex.F("08:00-12:30,13:30-17:00"),
				Sat:      cfrex.F("08:00-12:30,13:30-17:00"),
				Sun:      cfrex.F("08:00-12:30,13:30-17:00"),
				Thu:      cfrex.F("08:00-12:30,13:30-17:00"),
				TimeZone: cfrex.F("America/New York"),
				Tue:      cfrex.F("08:00-12:30,13:30-17:00"),
				Wed:      cfrex.F("08:00-12:30,13:30-17:00"),
			}),
			Traffic: cfrex.F(`http.request.uri matches ".*a/partial/uri.*" and http.request.host in $01302951-49f9-47c9-a400-0297e60b6a10`),
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

func TestAccountGatewayRuleGet(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Rules.Get(
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

func TestAccountGatewayRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Rules.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.AccountGatewayRuleUpdateParams{
			Action:        cfrex.F(cfrex.ActionPerformOn),
			Name:          cfrex.F("block bad websites"),
			Description:   cfrex.F("Block bad websites based on their host name."),
			DevicePosture: cfrex.F(`any(device_posture.checks.passed[*] in {"1308749e-fcfb-4ebc-b051-fe022b632644"})`),
			Enabled:       cfrex.F(true),
			Expiration: cfrex.F(cfrex.ExpirationParam{
				ExpiresAt: cfrex.F("2014-01-01T05:20:20Z"),
				Duration:  cfrex.F(int64(10)),
				Expired:   cfrex.F(false),
			}),
			Filters:    cfrex.F([]cfrex.GatewayFilters{cfrex.GatewayFiltersHTTP}),
			Identity:   cfrex.F(`any(identity.groups.name[*] in {"finance"})`),
			Precedence: cfrex.F(int64(0)),
			RuleSettings: cfrex.F(cfrex.RuleSettingsParam{
				AddHeaders: cfrex.F(map[string]string{
					"My-Next-Header":       "string",
					"X-Custom-Header-Name": "string",
				}),
				AllowChildBypass: cfrex.F(false),
				AuditSSH: cfrex.F(cfrex.RuleSettingsAuditSSHParam{
					CommandLogging: cfrex.F(false),
				}),
				BisoAdminControls: cfrex.F(cfrex.RuleSettingsBisoAdminControlsParam{
					Copy:     cfrex.F(cfrex.RuleSettingsBisoAdminControlsCopyEnabled),
					Dcp:      cfrex.F(false),
					Dd:       cfrex.F(false),
					Dk:       cfrex.F(false),
					Download: cfrex.F(cfrex.RuleSettingsBisoAdminControlsDownloadEnabled),
					Dp:       cfrex.F(false),
					Du:       cfrex.F(false),
					Keyboard: cfrex.F(cfrex.RuleSettingsBisoAdminControlsKeyboardEnabled),
					Paste:    cfrex.F(cfrex.RuleSettingsBisoAdminControlsPasteEnabled),
					Printing: cfrex.F(cfrex.RuleSettingsBisoAdminControlsPrintingEnabled),
					Upload:   cfrex.F(cfrex.RuleSettingsBisoAdminControlsUploadEnabled),
					Version:  cfrex.F(cfrex.RuleSettingsBisoAdminControlsVersionV1),
				}),
				BlockPageEnabled: cfrex.F(true),
				BlockReason:      cfrex.F("This website is a security risk"),
				BypassParentRule: cfrex.F(false),
				CheckSession: cfrex.F(cfrex.RuleSettingsCheckSessionParam{
					Duration: cfrex.F("300s"),
					Enforce:  cfrex.F(true),
				}),
				DNSResolvers: cfrex.F(cfrex.RuleSettingsDNSResolversParam{
					Ipv4: cfrex.F([]cfrex.RuleSettingsDNSResolversIpv4Param{{
						IP:                         cfrex.F("2.2.2.2"),
						Port:                       cfrex.F(int64(5053)),
						RouteThroughPrivateNetwork: cfrex.F(true),
						VnetID:                     cfrex.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
					Ipv6: cfrex.F([]cfrex.RuleSettingsDNSResolversIpv6Param{{
						IP:                         cfrex.F("2001:DB8::"),
						Port:                       cfrex.F(int64(5053)),
						RouteThroughPrivateNetwork: cfrex.F(true),
						VnetID:                     cfrex.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
				}),
				Egress: cfrex.F(cfrex.RuleSettingsEgressParam{
					Ipv4:         cfrex.F("192.0.2.2"),
					Ipv4Fallback: cfrex.F("192.0.2.3"),
					Ipv6:         cfrex.F("2001:DB8::/64"),
				}),
				IgnoreCnameCategoryMatches:      cfrex.F(true),
				InsecureDisableDnssecValidation: cfrex.F(false),
				IPCategories:                    cfrex.F(true),
				IPIndicatorFeeds:                cfrex.F(true),
				L4override: cfrex.F(cfrex.RuleSettingsL4overrideParam{
					IP:   cfrex.F("1.1.1.1"),
					Port: cfrex.F(int64(0)),
				}),
				NotificationSettings: cfrex.F(cfrex.RuleSettingsNotificationSettingsParam{
					Enabled:    cfrex.F(true),
					Msg:        cfrex.F("msg"),
					SupportURL: cfrex.F("support_url"),
				}),
				OverrideHost: cfrex.F("example.com"),
				OverrideIPs:  cfrex.F([]string{"1.1.1.1", "2.2.2.2"}),
				PayloadLog: cfrex.F(cfrex.RuleSettingsPayloadLogParam{
					Enabled: cfrex.F(true),
				}),
				Quarantine: cfrex.F(cfrex.RuleSettingsQuarantineParam{
					FileTypes: cfrex.F([]cfrex.RuleSettingsQuarantineFileType{cfrex.RuleSettingsQuarantineFileTypeExe}),
				}),
				Redirect: cfrex.F(cfrex.RuleSettingsRedirectParam{
					TargetUri:            cfrex.F("https://example.com"),
					PreservePathAndQuery: cfrex.F(true),
				}),
				ResolveDNSInternally: cfrex.F(cfrex.RuleSettingsResolveDNSInternallyParam{
					Fallback: cfrex.F(cfrex.RuleSettingsResolveDNSInternallyFallbackNone),
					ViewID:   cfrex.F("view_id"),
				}),
				ResolveDNSThroughCloudflare: cfrex.F(true),
				UntrustedCert: cfrex.F(cfrex.RuleSettingsUntrustedCertParam{
					Action: cfrex.F(cfrex.RuleSettingsUntrustedCertActionPassThrough),
				}),
			}),
			Schedule: cfrex.F(cfrex.ScheduleParam{
				Fri:      cfrex.F("08:00-12:30,13:30-17:00"),
				Mon:      cfrex.F("08:00-12:30,13:30-17:00"),
				Sat:      cfrex.F("08:00-12:30,13:30-17:00"),
				Sun:      cfrex.F("08:00-12:30,13:30-17:00"),
				Thu:      cfrex.F("08:00-12:30,13:30-17:00"),
				TimeZone: cfrex.F("America/New York"),
				Tue:      cfrex.F("08:00-12:30,13:30-17:00"),
				Wed:      cfrex.F("08:00-12:30,13:30-17:00"),
			}),
			Traffic: cfrex.F(`http.request.uri matches ".*a/partial/uri.*" and http.request.host in $01302951-49f9-47c9-a400-0297e60b6a10`),
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

func TestAccountGatewayRuleList(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Rules.List(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountGatewayRuleDelete(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Rules.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cfrex.AccountGatewayRuleDeleteParams{
			Body: map[string]interface{}{},
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

func TestAccountGatewayRuleResetExpiration(t *testing.T) {
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
	_, err := client.Accounts.Gateway.Rules.ResetExpiration(
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
