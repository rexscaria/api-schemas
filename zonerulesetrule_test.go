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

func TestZoneRulesetRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Rulesets.Rules.New(
		context.TODO(),
		"9f1839b6152d298aca64c4e906b6d074",
		"2f2feab2026849078ba485f918791bdc",
		cfrex.ZoneRulesetRuleNewParams{
			Body: cfrex.ZoneRulesetRuleNewParamsBodyObject{
				ID:     cfrex.F("3a03d665bac047339bb530ecb439a90d"),
				Action: cfrex.F(cfrex.ZoneRulesetRuleNewParamsBodyObjectActionBlock),
				ActionParameters: cfrex.F(cfrex.ZoneRulesetRuleNewParamsBodyObjectActionParameters{
					Response: cfrex.F(cfrex.ZoneRulesetRuleNewParamsBodyObjectActionParametersResponse{
						Content:     cfrex.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cfrex.F("application/json"),
						StatusCode:  cfrex.F(int64(400)),
					}),
				}),
				Description: cfrex.F[any](map[string]interface{}{}),
				Enabled:     cfrex.F(true),
				ExposedCredentialCheck: cfrex.F(cfrex.ZoneRulesetRuleNewParamsBodyObjectExposedCredentialCheck{
					PasswordExpression: cfrex.F(`url_decode(http.request.body.form[\"password\"][0])`),
					UsernameExpression: cfrex.F(`url_decode(http.request.body.form[\"username\"][0])`),
				}),
				Expression: cfrex.F("ip.src ne 1.1.1.1"),
				Logging: cfrex.F(cfrex.ZoneRulesetRuleNewParamsBodyObjectLogging{
					Enabled: cfrex.F(true),
				}),
				Position: cfrex.F[cfrex.ZoneRulesetRuleNewParamsBodyObjectPositionUnion](cfrex.ZoneRulesetRuleNewParamsBodyObjectPositionBefore{
					Before: cfrex.F("da5e8e506c8e7877fe06cdf4c41add54"),
				}),
				Ratelimit: cfrex.F(cfrex.ZoneRulesetRuleNewParamsBodyObjectRatelimit{
					Characteristics:         cfrex.F([]string{"ip.src"}),
					Period:                  cfrex.F(cfrex.ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod60),
					CountingExpression:      cfrex.F(`http.request.body.raw eq "abcd"`),
					MitigationTimeout:       cfrex.F(int64(600)),
					RequestsPerPeriod:       cfrex.F(int64(1000)),
					RequestsToOrigin:        cfrex.F(true),
					ScorePerPeriod:          cfrex.F(int64(400)),
					ScoreResponseHeaderName: cfrex.F("my-score"),
				}),
				Ref: cfrex.F("my_ref"),
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

func TestZoneRulesetRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Rulesets.Rules.Update(
		context.TODO(),
		"9f1839b6152d298aca64c4e906b6d074",
		"2f2feab2026849078ba485f918791bdc",
		"3a03d665bac047339bb530ecb439a90d",
		cfrex.ZoneRulesetRuleUpdateParams{
			Body: cfrex.ZoneRulesetRuleUpdateParamsBodyObject{
				ID:     cfrex.F("3a03d665bac047339bb530ecb439a90d"),
				Action: cfrex.F(cfrex.ZoneRulesetRuleUpdateParamsBodyObjectActionBlock),
				ActionParameters: cfrex.F(cfrex.ZoneRulesetRuleUpdateParamsBodyObjectActionParameters{
					Response: cfrex.F(cfrex.ZoneRulesetRuleUpdateParamsBodyObjectActionParametersResponse{
						Content:     cfrex.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cfrex.F("application/json"),
						StatusCode:  cfrex.F(int64(400)),
					}),
				}),
				Description: cfrex.F[any](map[string]interface{}{}),
				Enabled:     cfrex.F(true),
				ExposedCredentialCheck: cfrex.F(cfrex.ZoneRulesetRuleUpdateParamsBodyObjectExposedCredentialCheck{
					PasswordExpression: cfrex.F(`url_decode(http.request.body.form[\"password\"][0])`),
					UsernameExpression: cfrex.F(`url_decode(http.request.body.form[\"username\"][0])`),
				}),
				Expression: cfrex.F("ip.src ne 1.1.1.1"),
				Logging: cfrex.F(cfrex.ZoneRulesetRuleUpdateParamsBodyObjectLogging{
					Enabled: cfrex.F(true),
				}),
				Position: cfrex.F[cfrex.ZoneRulesetRuleUpdateParamsBodyObjectPositionUnion](cfrex.ZoneRulesetRuleUpdateParamsBodyObjectPositionBefore{
					Before: cfrex.F("da5e8e506c8e7877fe06cdf4c41add54"),
				}),
				Ratelimit: cfrex.F(cfrex.ZoneRulesetRuleUpdateParamsBodyObjectRatelimit{
					Characteristics:         cfrex.F([]string{"ip.src"}),
					Period:                  cfrex.F(cfrex.ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod60),
					CountingExpression:      cfrex.F(`http.request.body.raw eq "abcd"`),
					MitigationTimeout:       cfrex.F(int64(600)),
					RequestsPerPeriod:       cfrex.F(int64(1000)),
					RequestsToOrigin:        cfrex.F(true),
					ScorePerPeriod:          cfrex.F(int64(400)),
					ScoreResponseHeaderName: cfrex.F("my-score"),
				}),
				Ref: cfrex.F("my_ref"),
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

func TestZoneRulesetRuleDelete(t *testing.T) {
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
	_, err := client.Zones.Rulesets.Rules.Delete(
		context.TODO(),
		"9f1839b6152d298aca64c4e906b6d074",
		"2f2feab2026849078ba485f918791bdc",
		"3a03d665bac047339bb530ecb439a90d",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
