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

func TestAccountStreamNewWithOptionalParams(t *testing.T) {
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
	err := client.Accounts.Stream.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountStreamNewParams{
			Body:           map[string]interface{}{},
			TusResumable:   cfrex.F(cfrex.AccountStreamNewParamsTusResumable1_0_0),
			UploadLength:   cfrex.F(int64(0)),
			DirectUser:     cfrex.F(true),
			UploadCreator:  cfrex.F("creator-id_abcde12345"),
			UploadMetadata: cfrex.F("name aGVsbG8gd29ybGQ=, requiresignedurls, allowedorigins ZXhhbXBsZS5jb20sdGVzdC5jb20="),
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

func TestAccountStreamGet(t *testing.T) {
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
	_, err := client.Accounts.Stream.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"ea95132c15732412d22c1476fa83f27a",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountStreamUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Stream.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"ea95132c15732412d22c1476fa83f27a",
		cfrex.AccountStreamUpdateParams{
			AllowedOrigins:     cfrex.F([]string{"example.com"}),
			Creator:            cfrex.F("creator-id_abcde12345"),
			MaxDurationSeconds: cfrex.F(int64(1)),
			Meta: cfrex.F[any](map[string]interface{}{
				"name": "video12345.mp4",
			}),
			RequireSignedURLs:     cfrex.F(true),
			ScheduledDeletion:     cfrex.F(time.Now()),
			ThumbnailTimestampPct: cfrex.F(0.529241),
			UploadExpiry:          cfrex.F(time.Now()),
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

func TestAccountStreamListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Stream.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountStreamListParams{
			Asc:           cfrex.F(true),
			Creator:       cfrex.F("creator-id_abcde12345"),
			End:           cfrex.F(time.Now()),
			IncludeCounts: cfrex.F(true),
			Search:        cfrex.F("puppy.mp4"),
			Start:         cfrex.F(time.Now()),
			Status:        cfrex.F(cfrex.MediaStateInprogress),
			Type:          cfrex.F("live"),
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

func TestAccountStreamDelete(t *testing.T) {
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
	err := client.Accounts.Stream.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"ea95132c15732412d22c1476fa83f27a",
		cfrex.AccountStreamDeleteParams{
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

func TestAccountStreamClipWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Stream.Clip(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountStreamClipParams{
			ClippedFromVideoUid:   cfrex.F("023e105f4ecef8ad9ca31a8372d0c353"),
			EndTimeSeconds:        cfrex.F(int64(0)),
			StartTimeSeconds:      cfrex.F(int64(0)),
			AllowedOrigins:        cfrex.F([]string{"example.com"}),
			Creator:               cfrex.F("creator-id_abcde12345"),
			MaxDurationSeconds:    cfrex.F(int64(1)),
			RequireSignedURLs:     cfrex.F(true),
			ThumbnailTimestampPct: cfrex.F(0.529241),
			Watermark: cfrex.F(cfrex.WatermarkAtUploadEventParam{
				Uid: cfrex.F("ea95132c15732412d22c1476fa83f27a"),
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

func TestAccountStreamCopyWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Stream.Copy(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountStreamCopyParams{
			URL:            cfrex.F("https://example.com/myvideo.mp4"),
			AllowedOrigins: cfrex.F([]string{"example.com"}),
			Creator:        cfrex.F("creator-id_abcde12345"),
			Meta: cfrex.F[any](map[string]interface{}{
				"name": "video12345.mp4",
			}),
			RequireSignedURLs:     cfrex.F(true),
			ScheduledDeletion:     cfrex.F(time.Now()),
			ThumbnailTimestampPct: cfrex.F(0.529241),
			Watermark: cfrex.F(cfrex.WatermarkAtUploadStreamParam{
				Uid: cfrex.F("ea95132c15732412d22c1476fa83f27a"),
			}),
			UploadCreator: cfrex.F("creator-id_abcde12345"),
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

func TestAccountStreamNewSignedURLWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Stream.NewSignedURL(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"ea95132c15732412d22c1476fa83f27a",
		cfrex.AccountStreamNewSignedURLParams{
			ID: cfrex.F("ab0d4ef71g4425f8dcba9041231813000"),
			AccessRules: cfrex.F([]cfrex.AccountStreamNewSignedURLParamsAccessRule{{
				Action:  cfrex.F(cfrex.AccountStreamNewSignedURLParamsAccessRulesActionBlock),
				Country: cfrex.F([]string{"US", "MX"}),
				IP:      cfrex.F([]string{"string"}),
				Type:    cfrex.F(cfrex.AccountStreamNewSignedURLParamsAccessRulesTypeIPGeoipCountry),
			}, {
				Action:  cfrex.F(cfrex.AccountStreamNewSignedURLParamsAccessRulesActionAllow),
				Country: cfrex.F([]string{"string"}),
				IP:      cfrex.F([]string{"93.184.216.0/24", "2400:cb00::/32"}),
				Type:    cfrex.F(cfrex.AccountStreamNewSignedURLParamsAccessRulesTypeIPSrc),
			}, {
				Action:  cfrex.F(cfrex.AccountStreamNewSignedURLParamsAccessRulesActionBlock),
				Country: cfrex.F([]string{"string"}),
				IP:      cfrex.F([]string{"string"}),
				Type:    cfrex.F(cfrex.AccountStreamNewSignedURLParamsAccessRulesTypeAny),
			}}),
			Downloadable: cfrex.F(true),
			Exp:          cfrex.F(int64(0)),
			Nbf:          cfrex.F(int64(0)),
			Pem:          cfrex.F("LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBc284dnBvOFpEWXRkOUgzbWlPaW1qYXAzVXlVM0oyZ3kwTUYvN1R4blJuRnkwRHpDCkxqUk9naFZsQ0hPQmxsd3NVaE9GU0lyYnN4K05tUTdBeS90TFpXSGxuVGF3UWJ5WGZGOStJeDhVSnNlSHBGV1oKNVF5Z1JYd2liSjh1MVVsZ2xlcmZHMkpueldjVXpZTzEySktZN3doSkw1ajROMWgxZFJNUXQ5Q1pkZFlCQWRzOQpCdk02cjRFMDcxQkhQekhWeDMrUTI1VWtubGdUNXIwS3FiM1E1Y0dlTlBXY1JreW1ybkJEWWR0OXR4eFFMb1dPCllzNXdsMnVYWFVYL0VGcDMwajU0Nmp6czllWExLYlNDbjJjTDZFVE96Y2x3aG9DRGx2a2VQT05rUE9LMDVKNUMKTm1TdFdhMG9hV1VGRzM0MFl3cVVrWGt4OU9tNndXd1JldU1uU1FJREFRQUJBb0lCQUFJOHo1ck5kOEdtOGJBMgo1S3pxQjI1R2lOVENwbUNJeW53NXRJWHZTQmNHcEdydUcvdlN2WG9kVlFVSVY0TWdHQkVXUEFrVzdsNWVBcHI4CnA1ZFd5SkRXYTNkdklFSE9vSEpYU3dBYksxZzZEMTNVa2NkZ1EyRGpoNVhuWDhHZCtBY2c2SmRTQWgxOWtYSHEKMk54RUtBVDB6Ri83a1g2MkRkREFBcWxmQkpGSXJodVIvZUdEVWh4L2piTTRhQ2JCcFdiM0pnRE9OYm5tS1ZoMwpxS2ZwZmRZZENZU1lzWUxrNTlxRDF2VFNwUVFUQ0VadW9VKzNzRVNhdkJzaUs1bU0vTzY5ZkRMRXNURG1MeTVQCmhEK3BMQXI0SlhNNjFwRGVBS0l3cUVqWWJybXlDRHRXTUdJNnZzZ0E1eXQzUUJaME9vV2w5QUkwdWxoZ3p4dXQKZ2ZFNTRRRUNnWUVBN0F3a0lhVEEzYmQ4Nk9jSVZnNFlrWGk1cm5aNDdsM1k4V24zcjIzUmVISXhLdkllRUtSbgp5bUlFNDFtRVBBSmlGWFpLK1VPTXdkeS9EcnFJUithT1JiT2NiV01jWUg2QzgvbG1wdVJFaXE3SW1Ub3VWcnA4CnlnUkprMWprVDA4cTIvNmg4eTBEdjJqMitsaHFXNzRNOUt0cmwxcTRlWmZRUFREL01tR1NnTWtDZ1lFQXdhY04KaSttN1p6dnJtL3NuekF2VlZ5SEtwZHVUUjNERk1naC9maC9tZ0ZHZ1RwZWtUOVV5b3FleGNYQXdwMVlhL01iQQoyNTVJVDZRbXZZTm5yNXp6Wmxic2tMV0hsYllvbWhmWnVXTHhXR3hRaEFORWdaMFVVdUVTRGMvbWx2UXZHbEtSCkZoaGhBUWlVSmdDamhPaHk1SlBiNGFldGRKd0UxK09lVWRFaE1vRUNnWUVBNG8yZ25CM1o4ck5xa3NzemlBek4KYmNuMlJVbDJOaW9pejBwS3JMaDFaT29NNE5BekpQdjJsaHRQMzdtS0htS1hLMHczRjFqTEgwSTBxZmxFVmVZbQpSU1huakdHazJjUnpBYUVzOGgrQzNheDE0Z01pZUtGU3BqNUpNOEFNbVVZOXQ1cUVhN2FYc3o0V1ZoOUlMYmVTCkRiNzlhKzVwd21LQVBrcnBsTHhyZFdrQ2dZQlNNSHVBWVdBbmJYZ1BDS2FZWklGVWJNUWNacmY0ZnpWQ2lmYksKYWZHampvRlNPZXdEOGdGK3BWdWJRTGwxbkFieU44ek1xVDRaaHhybUhpcFlqMjJDaHV2NmN3RXJtbGRiSnpwQwpBMnRaVXdkTk1ESFlMUG5lUHlZeGRJWnlsUXFVeW14SGkydElUQUxNcWtLOGV3ZWdXZHpkeGhQSlJScU5JazhrCmZIVHhnUUtCZ1FEUFc2UXIxY3F3QjNUdnVWdWR4WGRqUTdIcDFodXhrNEVWaEFJZllKNFhSTW1NUE5YS28wdHUKdUt6LzE0QW14R0dvSWJxYVc1bDMzeFNteUxhem84clNUN0tSTjVKME9JSHcrZkR5SFgxdHpVSjZCTldDcEFTcwpjbWdNK0htSzVON0w2bkNaZFJQY2IwU1hGaVRQUGhCUG1PVWFDUnpER0ZMK2JYM1VwajJKbWc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="),
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

func TestAccountStreamDirectUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Stream.DirectUpload(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountStreamDirectUploadParams{
			MaxDurationSeconds: cfrex.F(int64(1)),
			AllowedOrigins:     cfrex.F([]string{"example.com"}),
			Creator:            cfrex.F("creator-id_abcde12345"),
			Expiry:             cfrex.F(time.Now()),
			Meta: cfrex.F[any](map[string]interface{}{
				"name": "video12345.mp4",
			}),
			RequireSignedURLs:     cfrex.F(true),
			ScheduledDeletion:     cfrex.F(time.Now()),
			ThumbnailTimestampPct: cfrex.F(0.529241),
			Watermark: cfrex.F(cfrex.WatermarkAtUploadStreamParam{
				Uid: cfrex.F("ea95132c15732412d22c1476fa83f27a"),
			}),
			UploadCreator: cfrex.F("creator-id_abcde12345"),
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

func TestAccountStreamGetEmbedCode(t *testing.T) {
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
	_, err := client.Accounts.Stream.GetEmbedCode(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"ea95132c15732412d22c1476fa83f27a",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountStreamGetStorageUsageWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Stream.GetStorageUsage(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountStreamGetStorageUsageParams{
			Creator: cfrex.F("creator-id_abcde12345"),
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
