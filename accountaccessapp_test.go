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

func TestAccountAccessAppNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Access.Apps.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAccessAppNewParams{
			AppRequest: cfrex.AppRequestSelfHostedApplicationParam(cfrex.AppRequestSelfHostedApplicationParam{
				SelfHostedPropsAccountParam: cfrex.SelfHostedPropsAccountParam{
					Domain:                   cfrex.F("test.example.com/admin"),
					Type:                     cfrex.F("self_hosted"),
					AllowAuthenticateViaWarp: cfrex.F(true),
					AllowedIdps:              cfrex.F([]string{"699d98642c564d2e855e9661899b7252"}),
					AppLauncherVisible:       cfrex.F(true),
					AutoRedirectToIdentity:   cfrex.F(true),
					CorsHeaders: cfrex.F(cfrex.SelfHostedPropsAccountCorsHeadersParam{
						AllowAllHeaders:  cfrex.F(true),
						AllowAllMethods:  cfrex.F(true),
						AllowAllOrigins:  cfrex.F(true),
						AllowCredentials: cfrex.F(true),
						AllowedHeaders:   cfrex.F([]string{"string"}),
						AllowedMethods:   cfrex.F([]cfrex.AllowedMethodsItem{cfrex.AllowedMethodsItemGet}),
						AllowedOrigins:   cfrex.F([]string{"https://example.com"}),
						MaxAge:           cfrex.F(-1.000000),
					}),
					CustomDenyMessage:        cfrex.F("custom_deny_message"),
					CustomDenyURL:            cfrex.F("custom_deny_url"),
					CustomNonIdentityDenyURL: cfrex.F("custom_non_identity_deny_url"),
					CustomPages:              cfrex.F([]string{"699d98642c564d2e855e9661899b7252"}),
					Destinations: cfrex.F([]cfrex.SelfHostedPropsAccountDestinationsUnionParam{cfrex.SelfHostedPropsAccountDestinationsPublicDestinationParam{
						Type: cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPublicDestinationTypePublic),
						Uri:  cfrex.F("test.example.com/admin"),
					}, cfrex.SelfHostedPropsAccountDestinationsPublicDestinationParam{
						Type: cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPublicDestinationTypePublic),
						Uri:  cfrex.F("test.anotherexample.com/staff"),
					}, cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationParam{
						Cidr:       cfrex.F("10.5.0.0/24"),
						Hostname:   cfrex.F("hostname"),
						L4Protocol: cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationL4ProtocolTcp),
						PortRange:  cfrex.F("80-90"),
						Type:       cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationTypePrivate),
						VnetID:     cfrex.F("vnet_id"),
					}, cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationParam{
						Cidr:       cfrex.F("10.5.0.3/32"),
						Hostname:   cfrex.F("hostname"),
						L4Protocol: cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationL4ProtocolTcp),
						PortRange:  cfrex.F("80"),
						Type:       cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationTypePrivate),
						VnetID:     cfrex.F("vnet_id"),
					}, cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationParam{
						Cidr:       cfrex.F("cidr"),
						Hostname:   cfrex.F("hostname"),
						L4Protocol: cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationL4ProtocolTcp),
						PortRange:  cfrex.F("port_range"),
						Type:       cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationTypePrivate),
						VnetID:     cfrex.F("vnet_id"),
					}}),
					EnableBindingCookie:     cfrex.F(true),
					HTTPOnlyCookieAttribute: cfrex.F(true),
					LogoURL:                 cfrex.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
					Name:                    cfrex.F("Admin Site"),
					OptionsPreflightBypass:  cfrex.F(true),
					PathCookieAttribute:     cfrex.F(true),
					SameSiteCookieAttribute: cfrex.F("strict"),
					SelfHostedDomains:       cfrex.F([]string{"test.example.com/admin", "test.anotherexample.com/staff"}),
					ServiceAuth401Redirect:  cfrex.F(true),
					SessionDuration:         cfrex.F("24h"),
					SkipInterstitial:        cfrex.F(true),
					Tags:                    cfrex.F([]string{"engineers"}),
				},
				EmbeddedPoliciesParam: cfrex.EmbeddedPoliciesParam{
					Policies: cfrex.F([]cfrex.EmbeddedPoliciesPoliciesUnionParam{cfrex.EmbeddedPoliciesPoliciesAccessAppPolicyLinkParam{
						ID:         cfrex.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
						Precedence: cfrex.F(int64(0)),
					}}),
				},
				EmbeddedScimConfigParam: cfrex.EmbeddedScimConfigParam{
					ScimConfig: cfrex.F(cfrex.ScimConfigParam{
						IdpUid:    cfrex.F("idp_uid"),
						RemoteUri: cfrex.F("remote_uri"),
						Authentication: cfrex.F[cfrex.ScimConfigAuthenticationUnionParam](cfrex.ScimConfigAuthHTTPBasicParam{
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

func TestAccountAccessAppGet(t *testing.T) {
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
	_, err := client.Accounts.Access.Apps.Get(
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

func TestAccountAccessAppUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Access.Apps.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAccessAppUpdateParams{
			AppRequest: cfrex.AppRequestSelfHostedApplicationParam(cfrex.AppRequestSelfHostedApplicationParam{
				SelfHostedPropsAccountParam: cfrex.SelfHostedPropsAccountParam{
					Domain:                   cfrex.F("test.example.com/admin"),
					Type:                     cfrex.F("self_hosted"),
					AllowAuthenticateViaWarp: cfrex.F(true),
					AllowedIdps:              cfrex.F([]string{"699d98642c564d2e855e9661899b7252"}),
					AppLauncherVisible:       cfrex.F(true),
					AutoRedirectToIdentity:   cfrex.F(true),
					CorsHeaders: cfrex.F(cfrex.SelfHostedPropsAccountCorsHeadersParam{
						AllowAllHeaders:  cfrex.F(true),
						AllowAllMethods:  cfrex.F(true),
						AllowAllOrigins:  cfrex.F(true),
						AllowCredentials: cfrex.F(true),
						AllowedHeaders:   cfrex.F([]string{"string"}),
						AllowedMethods:   cfrex.F([]cfrex.AllowedMethodsItem{cfrex.AllowedMethodsItemGet}),
						AllowedOrigins:   cfrex.F([]string{"https://example.com"}),
						MaxAge:           cfrex.F(-1.000000),
					}),
					CustomDenyMessage:        cfrex.F("custom_deny_message"),
					CustomDenyURL:            cfrex.F("custom_deny_url"),
					CustomNonIdentityDenyURL: cfrex.F("custom_non_identity_deny_url"),
					CustomPages:              cfrex.F([]string{"699d98642c564d2e855e9661899b7252"}),
					Destinations: cfrex.F([]cfrex.SelfHostedPropsAccountDestinationsUnionParam{cfrex.SelfHostedPropsAccountDestinationsPublicDestinationParam{
						Type: cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPublicDestinationTypePublic),
						Uri:  cfrex.F("test.example.com/admin"),
					}, cfrex.SelfHostedPropsAccountDestinationsPublicDestinationParam{
						Type: cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPublicDestinationTypePublic),
						Uri:  cfrex.F("test.anotherexample.com/staff"),
					}, cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationParam{
						Cidr:       cfrex.F("10.5.0.0/24"),
						Hostname:   cfrex.F("hostname"),
						L4Protocol: cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationL4ProtocolTcp),
						PortRange:  cfrex.F("80-90"),
						Type:       cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationTypePrivate),
						VnetID:     cfrex.F("vnet_id"),
					}, cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationParam{
						Cidr:       cfrex.F("10.5.0.3/32"),
						Hostname:   cfrex.F("hostname"),
						L4Protocol: cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationL4ProtocolTcp),
						PortRange:  cfrex.F("80"),
						Type:       cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationTypePrivate),
						VnetID:     cfrex.F("vnet_id"),
					}, cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationParam{
						Cidr:       cfrex.F("cidr"),
						Hostname:   cfrex.F("hostname"),
						L4Protocol: cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationL4ProtocolTcp),
						PortRange:  cfrex.F("port_range"),
						Type:       cfrex.F(cfrex.SelfHostedPropsAccountDestinationsPrivateDestinationTypePrivate),
						VnetID:     cfrex.F("vnet_id"),
					}}),
					EnableBindingCookie:     cfrex.F(true),
					HTTPOnlyCookieAttribute: cfrex.F(true),
					LogoURL:                 cfrex.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
					Name:                    cfrex.F("Admin Site"),
					OptionsPreflightBypass:  cfrex.F(true),
					PathCookieAttribute:     cfrex.F(true),
					SameSiteCookieAttribute: cfrex.F("strict"),
					SelfHostedDomains:       cfrex.F([]string{"test.example.com/admin", "test.anotherexample.com/staff"}),
					ServiceAuth401Redirect:  cfrex.F(true),
					SessionDuration:         cfrex.F("24h"),
					SkipInterstitial:        cfrex.F(true),
					Tags:                    cfrex.F([]string{"engineers"}),
				},
				EmbeddedPoliciesParam: cfrex.EmbeddedPoliciesParam{
					Policies: cfrex.F([]cfrex.EmbeddedPoliciesPoliciesUnionParam{cfrex.EmbeddedPoliciesPoliciesAccessAppPolicyLinkParam{
						ID:         cfrex.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
						Precedence: cfrex.F(int64(0)),
					}}),
				},
				EmbeddedScimConfigParam: cfrex.EmbeddedScimConfigParam{
					ScimConfig: cfrex.F(cfrex.ScimConfigParam{
						IdpUid:    cfrex.F("idp_uid"),
						RemoteUri: cfrex.F("remote_uri"),
						Authentication: cfrex.F[cfrex.ScimConfigAuthenticationUnionParam](cfrex.ScimConfigAuthHTTPBasicParam{
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

func TestAccountAccessAppListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Access.Apps.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAccessAppListParams{
			Aud:    cfrex.F("aud"),
			Domain: cfrex.F("domain"),
			Name:   cfrex.F("name"),
			Search: cfrex.F("search"),
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

func TestAccountAccessAppDelete(t *testing.T) {
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
	_, err := client.Accounts.Access.Apps.Delete(
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

func TestAccountAccessAppRevokeTokens(t *testing.T) {
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
	_, err := client.Accounts.Access.Apps.RevokeTokens(
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

func TestAccountAccessAppUserPolicyChecks(t *testing.T) {
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
	_, err := client.Accounts.Access.Apps.UserPolicyChecks(
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
