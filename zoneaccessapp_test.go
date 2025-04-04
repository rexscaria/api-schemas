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

func TestZoneAccessAppNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Access.Apps.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneAccessAppNewParams{
			AccessApps: cfrex.AccessAppsSelfHostedApplicationParam(cfrex.AccessAppsSelfHostedApplicationParam{
				BasicAppResponsePropsParam: cfrex.BasicAppResponsePropsParam{
					ID:        cfrex.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					CreatedAt: cfrex.F(time.Now()),
					ScimConfig: cfrex.F(cfrex.BasicAppResponsePropsScimConfigParam{
						IdpUid:    cfrex.F("idp_uid"),
						RemoteUri: cfrex.F("remote_uri"),
						Authentication: cfrex.F[cfrex.BasicAppResponsePropsScimConfigAuthenticationUnionParam](cfrex.ScimConfigAuthHTTPBasicParam{
							Password: cfrex.F("password"),
							Scheme:   cfrex.F(cfrex.ScimConfigAuthHTTPBasicSchemeHttpbasic),
							User:     cfrex.F("user"),
						}),
						DeactivateOnDelete: cfrex.F(true),
						Enabled:            cfrex.F(true),
						Mappings: cfrex.F([]cfrex.ScimConfigMappingParam{{
							Schema:  cfrex.F("urn:ietf:params:scim:schemas:core:2.0:User"),
							Enabled: cfrex.F(true),
							Filter:  cfrex.F(`title pr or userType eq "Intern"`),
							Operations: cfrex.F(cfrex.ScimConfigMappingOperationsParam{
								Create: cfrex.F(true),
								Delete: cfrex.F(true),
								Update: cfrex.F(true),
							}),
							Strictness:       cfrex.F(cfrex.ScimConfigMappingStrictnessStrict),
							TransformJsonata: cfrex.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
						}}),
					}),
					UpdatedAt: cfrex.F(time.Now()),
				},
				SelfHostedPropsZoneParam: cfrex.SelfHostedPropsZoneParam{
					Domain:                 cfrex.F("test.example.com/admin"),
					Type:                   cfrex.F("self_hosted"),
					AllowedIdps:            cfrex.F([]string{"699d98642c564d2e855e9661899b7252"}),
					AppLauncherVisible:     cfrex.F(true),
					AutoRedirectToIdentity: cfrex.F(true),
					CorsHeaders: cfrex.F(cfrex.SelfHostedPropsZoneCorsHeadersParam{
						AllowAllHeaders:  cfrex.F(true),
						AllowAllMethods:  cfrex.F(true),
						AllowAllOrigins:  cfrex.F(true),
						AllowCredentials: cfrex.F(true),
						AllowedHeaders:   cfrex.F([]interface{}{map[string]interface{}{}}),
						AllowedMethods:   cfrex.F([]cfrex.AllowedMethodsItem{cfrex.AllowedMethodsItemGet}),
						AllowedOrigins:   cfrex.F([]interface{}{"https://example.com"}),
						MaxAge:           cfrex.F(-1.000000),
					}),
					CustomDenyMessage:       cfrex.F("custom_deny_message"),
					CustomDenyURL:           cfrex.F("custom_deny_url"),
					EnableBindingCookie:     cfrex.F(true),
					HTTPOnlyCookieAttribute: cfrex.F(true),
					LogoURL:                 cfrex.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
					Name:                    cfrex.F("Admin Site"),
					OptionsPreflightBypass:  cfrex.F(true),
					SameSiteCookieAttribute: cfrex.F("strict"),
					ServiceAuth401Redirect:  cfrex.F(true),
					SessionDuration:         cfrex.F("24h"),
					SkipInterstitial:        cfrex.F(true),
				},
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

func TestZoneAccessAppGet(t *testing.T) {
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
	_, err := client.Zones.Access.Apps.Get(
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

func TestZoneAccessAppUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Access.Apps.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneAccessAppUpdateParams{
			AccessApps: cfrex.AccessAppsSelfHostedApplicationParam(cfrex.AccessAppsSelfHostedApplicationParam{
				BasicAppResponsePropsParam: cfrex.BasicAppResponsePropsParam{
					ID:        cfrex.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					CreatedAt: cfrex.F(time.Now()),
					ScimConfig: cfrex.F(cfrex.BasicAppResponsePropsScimConfigParam{
						IdpUid:    cfrex.F("idp_uid"),
						RemoteUri: cfrex.F("remote_uri"),
						Authentication: cfrex.F[cfrex.BasicAppResponsePropsScimConfigAuthenticationUnionParam](cfrex.ScimConfigAuthHTTPBasicParam{
							Password: cfrex.F("password"),
							Scheme:   cfrex.F(cfrex.ScimConfigAuthHTTPBasicSchemeHttpbasic),
							User:     cfrex.F("user"),
						}),
						DeactivateOnDelete: cfrex.F(true),
						Enabled:            cfrex.F(true),
						Mappings: cfrex.F([]cfrex.ScimConfigMappingParam{{
							Schema:  cfrex.F("urn:ietf:params:scim:schemas:core:2.0:User"),
							Enabled: cfrex.F(true),
							Filter:  cfrex.F(`title pr or userType eq "Intern"`),
							Operations: cfrex.F(cfrex.ScimConfigMappingOperationsParam{
								Create: cfrex.F(true),
								Delete: cfrex.F(true),
								Update: cfrex.F(true),
							}),
							Strictness:       cfrex.F(cfrex.ScimConfigMappingStrictnessStrict),
							TransformJsonata: cfrex.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
						}}),
					}),
					UpdatedAt: cfrex.F(time.Now()),
				},
				SelfHostedPropsZoneParam: cfrex.SelfHostedPropsZoneParam{
					Domain:                 cfrex.F("test.example.com/admin"),
					Type:                   cfrex.F("self_hosted"),
					AllowedIdps:            cfrex.F([]string{"699d98642c564d2e855e9661899b7252"}),
					AppLauncherVisible:     cfrex.F(true),
					AutoRedirectToIdentity: cfrex.F(true),
					CorsHeaders: cfrex.F(cfrex.SelfHostedPropsZoneCorsHeadersParam{
						AllowAllHeaders:  cfrex.F(true),
						AllowAllMethods:  cfrex.F(true),
						AllowAllOrigins:  cfrex.F(true),
						AllowCredentials: cfrex.F(true),
						AllowedHeaders:   cfrex.F([]interface{}{map[string]interface{}{}}),
						AllowedMethods:   cfrex.F([]cfrex.AllowedMethodsItem{cfrex.AllowedMethodsItemGet}),
						AllowedOrigins:   cfrex.F([]interface{}{"https://example.com"}),
						MaxAge:           cfrex.F(-1.000000),
					}),
					CustomDenyMessage:       cfrex.F("custom_deny_message"),
					CustomDenyURL:           cfrex.F("custom_deny_url"),
					EnableBindingCookie:     cfrex.F(true),
					HTTPOnlyCookieAttribute: cfrex.F(true),
					LogoURL:                 cfrex.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
					Name:                    cfrex.F("Admin Site"),
					OptionsPreflightBypass:  cfrex.F(true),
					SameSiteCookieAttribute: cfrex.F("strict"),
					ServiceAuth401Redirect:  cfrex.F(true),
					SessionDuration:         cfrex.F("24h"),
					SkipInterstitial:        cfrex.F(true),
				},
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

func TestZoneAccessAppList(t *testing.T) {
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
	_, err := client.Zones.Access.Apps.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneAccessAppDelete(t *testing.T) {
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
	_, err := client.Zones.Access.Apps.Delete(
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

func TestZoneAccessAppRevokeTokens(t *testing.T) {
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
	_, err := client.Zones.Access.Apps.RevokeTokens(
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

func TestZoneAccessAppUserPolicyChecks(t *testing.T) {
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
	_, err := client.Zones.Access.Apps.UserPolicyChecks(
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
