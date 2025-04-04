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
	"github.com/stainless-sdks/cf-rex-go/shared"
)

func TestUserGet(t *testing.T) {
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
	_, err := client.User.Get(context.TODO())
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Update(context.TODO(), cfrex.UserUpdateParams{
		Country:   cfrex.F("US"),
		FirstName: cfrex.F("John"),
		LastName:  cfrex.F("Appleseed"),
		Telephone: cfrex.F("+1 123-123-1234"),
		Zipcode:   cfrex.F("12345"),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserListAuditLogsWithOptionalParams(t *testing.T) {
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
	_, err := client.User.ListAuditLogs(context.TODO(), cfrex.UserListAuditLogsParams{
		ID: cfrex.F("f174be97-19b1-40d6-954d-70cd5fbd52db"),
		Action: cfrex.F(cfrex.UserListAuditLogsParamsAction{
			Type: cfrex.F("add"),
		}),
		Actor: cfrex.F(cfrex.UserListAuditLogsParamsActor{
			Email: cfrex.F("alice@example.com"),
			IP:    cfrex.F("17.168.228.63"),
		}),
		Before:       cfrex.F[cfrex.UserListAuditLogsParamsBeforeUnion](shared.UnionTime(time.Now())),
		Direction:    cfrex.F(cfrex.UserListAuditLogsParamsDirectionDesc),
		Export:       cfrex.F(true),
		HideUserLogs: cfrex.F(true),
		Page:         cfrex.F(50.000000),
		PerPage:      cfrex.F(25.000000),
		Since:        cfrex.F[cfrex.UserListAuditLogsParamsSinceUnion](shared.UnionTime(time.Now())),
		Zone: cfrex.F(cfrex.UserListAuditLogsParamsZone{
			Name: cfrex.F("example.com"),
		}),
	})
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
