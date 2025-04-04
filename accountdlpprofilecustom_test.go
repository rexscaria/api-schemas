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

func TestAccountDlpProfileCustomNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Dlp.Profiles.Custom.New(
		context.TODO(),
		"account_id",
		cfrex.AccountDlpProfileCustomNewParams{
			Body: cfrex.AccountDlpProfileCustomNewParamsBodyProfiles{
				Profiles: cfrex.F([]cfrex.NewCustomProfileParam{{
					Entries: cfrex.F([]cfrex.NewCustomProfileEntriesUnionParam{cfrex.NewCustomEntryParam{
						Enabled: cfrex.F(true),
						Name:    cfrex.F("name"),
						Pattern: cfrex.F(cfrex.PatternParam{
							Regex:      cfrex.F("regex"),
							Validation: cfrex.F(cfrex.PatternValidationLuhn),
						}),
					}}),
					Name:                cfrex.F("name"),
					AIContextEnabled:    cfrex.F(true),
					AllowedMatchCount:   cfrex.F(int64(5)),
					ConfidenceThreshold: cfrex.F("confidence_threshold"),
					ContextAwareness: cfrex.F(cfrex.ContextAwarenessParam{
						Enabled: cfrex.F(true),
						Skip: cfrex.F(cfrex.ContextAwarenessSkipParam{
							Files: cfrex.F(true),
						}),
					}),
					Description: cfrex.F("description"),
					OcrEnabled:  cfrex.F(true),
					SharedEntries: cfrex.F([]cfrex.NewCustomProfileSharedEntriesUnionParam{cfrex.NewCustomProfileSharedEntriesObjectParam{
						Enabled:   cfrex.F(true),
						EntryID:   cfrex.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
						EntryType: cfrex.F(cfrex.NewCustomProfileSharedEntriesObjectEntryTypeCustom),
					}}),
				}}),
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

func TestAccountDlpProfileCustomGet(t *testing.T) {
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
	_, err := client.Accounts.Dlp.Profiles.Custom.Get(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountDlpProfileCustomUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Dlp.Profiles.Custom.Update(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cfrex.AccountDlpProfileCustomUpdateParams{
			Name:                cfrex.F("name"),
			AIContextEnabled:    cfrex.F(true),
			AllowedMatchCount:   cfrex.F(int64(0)),
			ConfidenceThreshold: cfrex.F("confidence_threshold"),
			ContextAwareness: cfrex.F(cfrex.ContextAwarenessParam{
				Enabled: cfrex.F(true),
				Skip: cfrex.F(cfrex.ContextAwarenessSkipParam{
					Files: cfrex.F(true),
				}),
			}),
			Description: cfrex.F("description"),
			Entries: cfrex.F([]cfrex.AccountDlpProfileCustomUpdateParamsEntryUnion{cfrex.AccountDlpProfileCustomUpdateParamsEntriesDlpNewCustomEntryWithID(cfrex.AccountDlpProfileCustomUpdateParamsEntriesDlpNewCustomEntryWithID{
				NewCustomEntryParam: cfrex.NewCustomEntryParam{
					Enabled: cfrex.F(true),
					Name:    cfrex.F("name"),
					Pattern: cfrex.F(cfrex.PatternParam{
						Regex:      cfrex.F("regex"),
						Validation: cfrex.F(cfrex.PatternValidationLuhn),
					}),
				},
				EntryID: cfrex.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			})}),
			OcrEnabled: cfrex.F(true),
			SharedEntries: cfrex.F([]cfrex.AccountDlpProfileCustomUpdateParamsSharedEntryUnion{cfrex.AccountDlpProfileCustomUpdateParamsSharedEntriesObject{
				Enabled:   cfrex.F(true),
				EntryID:   cfrex.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				EntryType: cfrex.F(cfrex.AccountDlpProfileCustomUpdateParamsSharedEntriesObjectEntryTypePredefined),
			}}),
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

func TestAccountDlpProfileCustomDelete(t *testing.T) {
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
	_, err := client.Accounts.Dlp.Profiles.Custom.Delete(
		context.TODO(),
		"account_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
