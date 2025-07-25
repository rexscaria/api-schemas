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

func TestAccountLogGetAuditLogsWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Logs.GetAuditLogs(
		context.TODO(),
		"a67e14daa5f8dceeb91fe5449ba496ef",
		cfrex.AccountLogGetAuditLogsParams{
			Before: cfrex.F(time.Now()),
			Since:  cfrex.F(time.Now()),
			AccountName: cfrex.F(cfrex.AccountLogGetAuditLogsParamsAccountName{
				Not: cfrex.F([]string{"string"}),
			}),
			ActionResult: cfrex.F(cfrex.AccountLogGetAuditLogsParamsActionResult{
				Not: cfrex.F([]cfrex.AccountLogGetAuditLogsParamsActionResultNot{cfrex.AccountLogGetAuditLogsParamsActionResultNotSuccess}),
			}),
			ActionType: cfrex.F(cfrex.AccountLogGetAuditLogsParamsActionType{
				Not: cfrex.F([]cfrex.AccountLogGetAuditLogsParamsActionTypeNot{cfrex.AccountLogGetAuditLogsParamsActionTypeNotCreate}),
			}),
			ActorContext: cfrex.F(cfrex.AccountLogGetAuditLogsParamsActorContext{
				Not: cfrex.F([]cfrex.AccountLogGetAuditLogsParamsActorContextNot{cfrex.AccountLogGetAuditLogsParamsActorContextNotAPIKey}),
			}),
			ActorEmail: cfrex.F(cfrex.AccountLogGetAuditLogsParamsActorEmail{
				Not: cfrex.F([]string{"alice@example.com"}),
			}),
			ActorID: cfrex.F(cfrex.AccountLogGetAuditLogsParamsActorID{
				Not: cfrex.F([]string{"1d20c3afe174f18b642710cec6298a9d"}),
			}),
			ActorIPAddress: cfrex.F(cfrex.AccountLogGetAuditLogsParamsActorIPAddress{
				Not: cfrex.F([]string{"17.168.228.63"}),
			}),
			ActorTokenID: cfrex.F(cfrex.AccountLogGetAuditLogsParamsActorTokenID{
				Not: cfrex.F([]string{"144cdb2e39c55e203cf225d8d8208647"}),
			}),
			ActorTokenName: cfrex.F(cfrex.AccountLogGetAuditLogsParamsActorTokenName{
				Not: cfrex.F([]string{"Test Token"}),
			}),
			ActorType: cfrex.F(cfrex.AccountLogGetAuditLogsParamsActorType{
				Not: cfrex.F([]cfrex.AccountLogGetAuditLogsParamsActorTypeNot{cfrex.AccountLogGetAuditLogsParamsActorTypeNotAccount}),
			}),
			AuditLogID: cfrex.F(cfrex.AccountLogGetAuditLogsParamsAuditLogID{
				Not: cfrex.F([]string{"f174be97-19b1-40d6-954d-70cd5fbd52db"}),
			}),
			Cursor:    cfrex.F("Q1buH-__DQqqig7SVYXT-SsMOTGY2Z3Y80W-fGgva7yaDdmPKveucH5ddOcHsJRhNb-xUK8agZQqkJSMAENGO8NU6g=="),
			Direction: cfrex.F(cfrex.AccountLogGetAuditLogsParamsDirectionDesc),
			Limit:     cfrex.F(25.000000),
			RawCfRayID: cfrex.F(cfrex.AccountLogGetAuditLogsParamsRawCfRayID{
				Not: cfrex.F([]string{"8e8dd2156ef28414"}),
			}),
			RawMethod: cfrex.F(cfrex.AccountLogGetAuditLogsParamsRawMethod{
				Not: cfrex.F([]string{"GET"}),
			}),
			RawStatusCode: cfrex.F(cfrex.AccountLogGetAuditLogsParamsRawStatusCode{
				Not: cfrex.F([]int64{int64(200)}),
			}),
			RawUri: cfrex.F(cfrex.AccountLogGetAuditLogsParamsRawUri{
				Not: cfrex.F([]string{"string"}),
			}),
			ResourceID: cfrex.F(cfrex.AccountLogGetAuditLogsParamsResourceID{
				Not: cfrex.F([]string{"string"}),
			}),
			ResourceProduct: cfrex.F(cfrex.AccountLogGetAuditLogsParamsResourceProduct{
				Not: cfrex.F([]string{"Stream"}),
			}),
			ResourceScope: cfrex.F(cfrex.AccountLogGetAuditLogsParamsResourceScope{
				Not: cfrex.F([]cfrex.AccountLogGetAuditLogsParamsResourceScopeNot{cfrex.AccountLogGetAuditLogsParamsResourceScopeNotAccounts}),
			}),
			ResourceType: cfrex.F(cfrex.AccountLogGetAuditLogsParamsResourceType{
				Not: cfrex.F([]string{"Video"}),
			}),
			ZoneID: cfrex.F(cfrex.AccountLogGetAuditLogsParamsZoneID{
				Not: cfrex.F([]string{"string"}),
			}),
			ZoneName: cfrex.F(cfrex.AccountLogGetAuditLogsParamsZoneName{
				Not: cfrex.F([]string{"example.com"}),
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
