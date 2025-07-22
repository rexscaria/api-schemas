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
	"github.com/rexscaria/api-schemas/shared"
)

func TestZoneSettingZarazConfigGet(t *testing.T) {
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
	_, err := client.Zones.Settings.Zaraz.Config.Get(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneSettingZarazConfigUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Settings.Zaraz.Config.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneSettingZarazConfigUpdateParams{
			DataLayer: cfrex.F(true),
			DebugKey:  cfrex.F("debugKey"),
			Settings: cfrex.F(cfrex.ZoneSettingZarazConfigUpdateParamsSettings{
				AutoInjectScript: cfrex.F(true),
				ContextEnricher: cfrex.F(cfrex.ZoneSettingZarazConfigUpdateParamsSettingsContextEnricher{
					EscapedWorkerName: cfrex.F("escapedWorkerName"),
					WorkerTag:         cfrex.F("workerTag"),
				}),
				CookieDomain:        cfrex.F("cookieDomain"),
				Ecommerce:           cfrex.F(true),
				EventsAPIPath:       cfrex.F("eventsApiPath"),
				HideExternalReferer: cfrex.F(true),
				HideIPAddress:       cfrex.F(true),
				HideQueryParams:     cfrex.F(true),
				HideUserAgent:       cfrex.F(true),
				InitPath:            cfrex.F("initPath"),
				InjectIframes:       cfrex.F(true),
				McRootPath:          cfrex.F("mcRootPath"),
				ScriptPath:          cfrex.F("scriptPath"),
				TrackPath:           cfrex.F("trackPath"),
			}),
			Tools: cfrex.F(map[string]cfrex.ZoneSettingZarazConfigUpdateParamsToolsUnion{
				"foo": cfrex.ZoneSettingZarazConfigUpdateParamsToolsObject{
					BlockingTriggers: cfrex.F([]string{"string"}),
					Component:        cfrex.F("component"),
					DefaultFields: cfrex.F(map[string]cfrex.ZoneSettingZarazConfigUpdateParamsToolsObjectDefaultFieldsUnion{
						"foo": shared.UnionString("string"),
					}),
					Enabled:     cfrex.F(true),
					Name:        cfrex.F("name"),
					Permissions: cfrex.F([]string{"string"}),
					Settings: cfrex.F(map[string]cfrex.ZoneSettingZarazConfigUpdateParamsToolsObjectSettingsUnion{
						"foo": shared.UnionString("string"),
					}),
					Type: cfrex.F(cfrex.ZoneSettingZarazConfigUpdateParamsToolsObjectTypeComponent),
					Actions: cfrex.F(map[string]cfrex.ZoneSettingZarazConfigUpdateParamsToolsObjectActions{
						"foo": {
							ActionType:       cfrex.F("actionType"),
							BlockingTriggers: cfrex.F([]string{"string"}),
							Data:             cfrex.F[any](map[string]interface{}{}),
							FiringTriggers:   cfrex.F([]string{"string"}),
						},
					}),
					DefaultPurpose: cfrex.F("defaultPurpose"),
					NeoEvents: cfrex.F([]cfrex.ZoneSettingZarazConfigUpdateParamsToolsObjectNeoEvent{{
						ActionType:       cfrex.F("actionType"),
						BlockingTriggers: cfrex.F([]string{"string"}),
						Data:             cfrex.F[any](map[string]interface{}{}),
						FiringTriggers:   cfrex.F([]string{"string"}),
					}}),
					VendorName:      cfrex.F("vendorName"),
					VendorPolicyURL: cfrex.F("vendorPolicyUrl"),
				},
			}),
			Triggers: cfrex.F(map[string]cfrex.ZoneSettingZarazConfigUpdateParamsTriggers{
				"foo": {
					ExcludeRules: cfrex.F([]cfrex.ZoneSettingZarazConfigUpdateParamsTriggersExcludeRuleUnion{cfrex.ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRule{
						ID:    cfrex.F("id"),
						Match: cfrex.F("match"),
						Op:    cfrex.F(cfrex.ZoneSettingZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpContains),
						Value: cfrex.F("value"),
					}}),
					LoadRules: cfrex.F([]cfrex.ZoneSettingZarazConfigUpdateParamsTriggersLoadRuleUnion{cfrex.ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRule{
						ID:    cfrex.F("id"),
						Match: cfrex.F("match"),
						Op:    cfrex.F(cfrex.ZoneSettingZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpContains),
						Value: cfrex.F("value"),
					}}),
					Name:        cfrex.F("name"),
					Description: cfrex.F("description"),
					System:      cfrex.F(cfrex.ZoneSettingZarazConfigUpdateParamsTriggersSystemPageload),
				},
			}),
			Variables: cfrex.F(map[string]cfrex.ZoneSettingZarazConfigUpdateParamsVariablesUnion{
				"foo": cfrex.ZoneSettingZarazConfigUpdateParamsVariablesZarazStringVariable{
					Name:  cfrex.F("name"),
					Type:  cfrex.F(cfrex.ZoneSettingZarazConfigUpdateParamsVariablesZarazStringVariableTypeString),
					Value: cfrex.F("value"),
				},
			}),
			ZarazVersion: cfrex.F(int64(0)),
			Analytics: cfrex.F(cfrex.ZoneSettingZarazConfigUpdateParamsAnalytics{
				DefaultPurpose: cfrex.F("defaultPurpose"),
				Enabled:        cfrex.F(true),
				SessionExpTime: cfrex.F(int64(60)),
			}),
			Consent: cfrex.F(cfrex.ZoneSettingZarazConfigUpdateParamsConsent{
				Enabled: cfrex.F(true),
				ButtonTextTranslations: cfrex.F(cfrex.ZoneSettingZarazConfigUpdateParamsConsentButtonTextTranslations{
					AcceptAll: cfrex.F(map[string]string{
						"foo": "string",
					}),
					ConfirmMyChoices: cfrex.F(map[string]string{
						"foo": "string",
					}),
					RejectAll: cfrex.F(map[string]string{
						"foo": "string",
					}),
				}),
				CompanyEmail:          cfrex.F("companyEmail"),
				CompanyName:           cfrex.F("companyName"),
				CompanyStreetAddress:  cfrex.F("companyStreetAddress"),
				ConsentModalIntroHTML: cfrex.F("consentModalIntroHTML"),
				ConsentModalIntroHTMLWithTranslations: cfrex.F(map[string]string{
					"foo": "string",
				}),
				CookieName:                     cfrex.F("cookieName"),
				CustomCss:                      cfrex.F("customCSS"),
				CustomIntroDisclaimerDismissed: cfrex.F(true),
				DefaultLanguage:                cfrex.F("defaultLanguage"),
				HideModal:                      cfrex.F(true),
				Purposes: cfrex.F(map[string]cfrex.ZoneSettingZarazConfigUpdateParamsConsentPurposes{
					"foo": {
						Description: cfrex.F("description"),
						Name:        cfrex.F("name"),
					},
				}),
				PurposesWithTranslations: cfrex.F(map[string]cfrex.ZoneSettingZarazConfigUpdateParamsConsentPurposesWithTranslations{
					"foo": {
						Description: cfrex.F(map[string]string{
							"foo": "string",
						}),
						Name: cfrex.F(map[string]string{
							"foo": "string",
						}),
						Order: cfrex.F(int64(0)),
					},
				}),
				TcfCompliant: cfrex.F(true),
			}),
			HistoryChange: cfrex.F(true),
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
