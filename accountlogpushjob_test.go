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

func TestAccountLogpushJobNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Logpush.Jobs.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountLogpushJobNewParams{
			DestinationConf:          cfrex.F("s3://mybucket/logs?region=us-west-2"),
			Dataset:                  cfrex.F("gateway_dns"),
			Enabled:                  cfrex.F(false),
			Frequency:                cfrex.F(cfrex.FrequencyHigh),
			Kind:                     cfrex.F(cfrex.LogpushKindEdge),
			LogpullOptions:           cfrex.F("fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"),
			MaxUploadBytes:           cfrex.F(int64(5000000)),
			MaxUploadIntervalSeconds: cfrex.F(int64(30)),
			MaxUploadRecords:         cfrex.F(int64(1000)),
			Name:                     cfrex.F("example.com"),
			OutputOptions: cfrex.F(cfrex.OutputOptionsParam{
				BatchPrefix:     cfrex.F(""),
				BatchSuffix:     cfrex.F(""),
				Cve2021_44228:   cfrex.F(false),
				FieldDelimiter:  cfrex.F(","),
				FieldNames:      cfrex.F([]string{"Datetime", "DstIP", "SrcIP"}),
				OutputType:      cfrex.F(cfrex.OutputOptionsOutputTypeNdjson),
				RecordDelimiter: cfrex.F(""),
				RecordPrefix:    cfrex.F("{"),
				RecordSuffix:    cfrex.F("}\n"),
				RecordTemplate:  cfrex.F("record_template"),
				SampleRate:      cfrex.F(1.000000),
				TimestampFormat: cfrex.F(cfrex.OutputOptionsTimestampFormatUnixnano),
			}),
			OwnershipChallenge: cfrex.F("00000000000000000000"),
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

func TestAccountLogpushJobGet(t *testing.T) {
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
	_, err := client.Accounts.Logpush.Jobs.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		int64(1),
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountLogpushJobUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Logpush.Jobs.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		int64(1),
		cfrex.AccountLogpushJobUpdateParams{
			DestinationConf:          cfrex.F("s3://mybucket/logs?region=us-west-2"),
			Enabled:                  cfrex.F(false),
			Frequency:                cfrex.F(cfrex.FrequencyHigh),
			Kind:                     cfrex.F(cfrex.LogpushKindEdge),
			LogpullOptions:           cfrex.F("fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"),
			MaxUploadBytes:           cfrex.F(int64(5000000)),
			MaxUploadIntervalSeconds: cfrex.F(int64(30)),
			MaxUploadRecords:         cfrex.F(int64(1000)),
			Name:                     cfrex.F("example.com"),
			OutputOptions: cfrex.F(cfrex.OutputOptionsParam{
				BatchPrefix:     cfrex.F(""),
				BatchSuffix:     cfrex.F(""),
				Cve2021_44228:   cfrex.F(false),
				FieldDelimiter:  cfrex.F(","),
				FieldNames:      cfrex.F([]string{"Datetime", "DstIP", "SrcIP"}),
				OutputType:      cfrex.F(cfrex.OutputOptionsOutputTypeNdjson),
				RecordDelimiter: cfrex.F(""),
				RecordPrefix:    cfrex.F("{"),
				RecordSuffix:    cfrex.F("}\n"),
				RecordTemplate:  cfrex.F("record_template"),
				SampleRate:      cfrex.F(1.000000),
				TimestampFormat: cfrex.F(cfrex.OutputOptionsTimestampFormatUnixnano),
			}),
			OwnershipChallenge: cfrex.F("00000000000000000000"),
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

func TestAccountLogpushJobList(t *testing.T) {
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
	_, err := client.Accounts.Logpush.Jobs.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountLogpushJobDelete(t *testing.T) {
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
	_, err := client.Accounts.Logpush.Jobs.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		int64(1),
		cfrex.AccountLogpushJobDeleteParams{
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
