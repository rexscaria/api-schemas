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
			Before:          cfrex.F(time.Now()),
			Since:           cfrex.F(time.Now()),
			AccountName:     cfrex.F("account_name"),
			ActionResult:    cfrex.F(cfrex.AccountLogGetAuditLogsParamsActionResultSuccess),
			ActionType:      cfrex.F(cfrex.AccountLogGetAuditLogsParamsActionTypeCreate),
			ActorContext:    cfrex.F(cfrex.AccountLogGetAuditLogsParamsActorContextAPIKey),
			ActorEmail:      cfrex.F("alice@example.com"),
			ActorID:         cfrex.F("1d20c3afe174f18b642710cec6298a9d"),
			ActorIPAddress:  cfrex.F("17.168.228.63"),
			ActorTokenID:    cfrex.F("144cdb2e39c55e203cf225d8d8208647"),
			ActorTokenName:  cfrex.F("Test Token"),
			ActorType:       cfrex.F(cfrex.AccountLogGetAuditLogsParamsActorTypeCloudflareAdmin),
			AuditLogID:      cfrex.F("f174be97-19b1-40d6-954d-70cd5fbd52db"),
			Cursor:          cfrex.F("Q1buH-__DQqqig7SVYXT-SsMOTGY2Z3Y80W-fGgva7yaDdmPKveucH5ddOcHsJRhNb-xUK8agZQqkJSMAENGO8NU6g=="),
			Direction:       cfrex.F(cfrex.AccountLogGetAuditLogsParamsDirectionDesc),
			Limit:           cfrex.F(25.000000),
			RawCfRayID:      cfrex.F("8e8dd2156ef28414"),
			RawMethod:       cfrex.F("GET"),
			RawStatusCode:   cfrex.F(int64(200)),
			RawUri:          cfrex.F("raw_uri"),
			ResourceID:      cfrex.F("resource_id"),
			ResourceProduct: cfrex.F("Stream"),
			ResourceScope:   cfrex.F(cfrex.AccountLogGetAuditLogsParamsResourceScopeAccounts),
			ResourceType:    cfrex.F("Video"),
			ZoneID:          cfrex.F("zone_id"),
			ZoneName:        cfrex.F("example.com"),
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
