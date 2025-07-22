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

func TestZoneDNSRecordNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneDNSRecordNewParams{
			DNSRecordPost: cfrex.DNSRecordPostDNSRecordsARecordParam(cfrex.DNSRecordPostDNSRecordsARecordParam{
				SharedFieldsParam: cfrex.SharedFieldsParam{
					Comment: cfrex.F("Domain verification record"),
					Name:    cfrex.F("example.com"),
					Proxied: cfrex.F(true),
					Settings: cfrex.F(cfrex.SharedFieldsSettingsParam{
						Ipv4Only: cfrex.F(true),
						Ipv6Only: cfrex.F(true),
					}),
					Tags: cfrex.F([]string{"owner:dns-team"}),
					Ttl:  cfrex.F(cfrex.SharedFieldsTtl1),
				},
				Content: cfrex.F("198.51.100.4"),
				Type:    cfrex.F(cfrex.DNSRecordPostDNSRecordsARecordTypeA),
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

func TestZoneDNSRecordGet(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.Get(
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

func TestZoneDNSRecordUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneDNSRecordUpdateParams{
			DNSRecordPatch: cfrex.DNSRecordPatchDNSRecordsARecordParam(cfrex.DNSRecordPatchDNSRecordsARecordParam{
				SharedFieldsParam: cfrex.SharedFieldsParam{
					Comment: cfrex.F("Domain verification record"),
					Name:    cfrex.F("example.com"),
					Proxied: cfrex.F(true),
					Settings: cfrex.F(cfrex.SharedFieldsSettingsParam{
						Ipv4Only: cfrex.F(true),
						Ipv6Only: cfrex.F(true),
					}),
					Tags: cfrex.F([]string{"owner:dns-team"}),
					Ttl:  cfrex.F(cfrex.SharedFieldsTtl1),
				},
				Content: cfrex.F("198.51.100.4"),
				Type:    cfrex.F(cfrex.DNSRecordPatchDNSRecordsARecordTypeA),
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

func TestZoneDNSRecordListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneDNSRecordListParams{
			Comment: cfrex.F(cfrex.ZoneDNSRecordListParamsComment{
				Absent:     cfrex.F("absent"),
				Contains:   cfrex.F("ello, worl"),
				Endswith:   cfrex.F("o, world"),
				Exact:      cfrex.F("Hello, world"),
				Present:    cfrex.F("present"),
				Startswith: cfrex.F("Hello, w"),
			}),
			Content: cfrex.F(cfrex.ZoneDNSRecordListParamsContent{
				Contains:   cfrex.F("7.0.0."),
				Endswith:   cfrex.F(".0.1"),
				Exact:      cfrex.F("127.0.0.1"),
				Startswith: cfrex.F("127.0."),
			}),
			Direction: cfrex.F(cfrex.ZoneDNSRecordListParamsDirectionAsc),
			Match:     cfrex.F(cfrex.ZoneDNSRecordListParamsMatchAny),
			Name: cfrex.F(cfrex.ZoneDNSRecordListParamsName{
				Contains:   cfrex.F("w.example."),
				Endswith:   cfrex.F(".example.com"),
				Exact:      cfrex.F("www.example.com"),
				Startswith: cfrex.F("www.example"),
			}),
			Order:   cfrex.F(cfrex.ZoneDNSRecordListParamsOrderType),
			Page:    cfrex.F(1.000000),
			PerPage: cfrex.F(5.000000),
			Proxied: cfrex.F(true),
			Search:  cfrex.F("www.cloudflare.com"),
			Tag: cfrex.F(cfrex.ZoneDNSRecordListParamsTag{
				Absent:     cfrex.F("important"),
				Contains:   cfrex.F("greeting:ello, worl"),
				Endswith:   cfrex.F("greeting:o, world"),
				Exact:      cfrex.F("greeting:Hello, world"),
				Present:    cfrex.F("important"),
				Startswith: cfrex.F("greeting:Hello, w"),
			}),
			TagMatch: cfrex.F(cfrex.ZoneDNSRecordListParamsTagMatchAny),
			Type:     cfrex.F(cfrex.ZoneDNSRecordListParamsTypeA),
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

func TestZoneDNSRecordDelete(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneDNSRecordDeleteParams{
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

func TestZoneDNSRecordBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.Batch(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneDNSRecordBatchParams{
			Deletes: cfrex.F([]cfrex.ZoneDNSRecordBatchParamsDelete{{
				ID: cfrex.F("023e105f4ecef8ad9ca31a8372d0c353"),
			}}),
			Patches: cfrex.F([]cfrex.ZoneDNSRecordBatchParamsPatchUnion{cfrex.ZoneDNSRecordBatchParamsPatchesDNSRecordsARecord(cfrex.ZoneDNSRecordBatchParamsPatchesDNSRecordsARecord{
				SharedFieldsParam: cfrex.SharedFieldsParam{
					Comment: cfrex.F("Domain verification record"),
					Name:    cfrex.F("example.com"),
					Proxied: cfrex.F(true),
					Settings: cfrex.F(cfrex.SharedFieldsSettingsParam{
						Ipv4Only: cfrex.F(true),
						Ipv6Only: cfrex.F(true),
					}),
					Tags: cfrex.F([]string{"owner:dns-team"}),
					Ttl:  cfrex.F(cfrex.SharedFieldsTtl1),
				},
				ID:      cfrex.F("023e105f4ecef8ad9ca31a8372d0c353"),
				Content: cfrex.F("198.51.100.4"),
				Type:    cfrex.F(cfrex.ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordTypeA),
			})}),
			Posts: cfrex.F([]cfrex.DNSRecordPostUnionParam{cfrex.DNSRecordPostDNSRecordsARecordParam(cfrex.DNSRecordPostDNSRecordsARecordParam{
				SharedFieldsParam: cfrex.SharedFieldsParam{
					Comment: cfrex.F("Domain verification record"),
					Name:    cfrex.F("example.com"),
					Proxied: cfrex.F(true),
					Settings: cfrex.F(cfrex.SharedFieldsSettingsParam{
						Ipv4Only: cfrex.F(true),
						Ipv6Only: cfrex.F(true),
					}),
					Tags: cfrex.F([]string{"owner:dns-team"}),
					Ttl:  cfrex.F(cfrex.SharedFieldsTtl1),
				},
				Content: cfrex.F("198.51.100.4"),
				Type:    cfrex.F(cfrex.DNSRecordPostDNSRecordsARecordTypeA),
			})}),
			Puts: cfrex.F([]cfrex.ZoneDNSRecordBatchParamsPutUnion{cfrex.ZoneDNSRecordBatchParamsPutsDNSRecordsARecord(cfrex.ZoneDNSRecordBatchParamsPutsDNSRecordsARecord{
				SharedFieldsParam: cfrex.SharedFieldsParam{
					Comment: cfrex.F("Domain verification record"),
					Name:    cfrex.F("example.com"),
					Proxied: cfrex.F(true),
					Settings: cfrex.F(cfrex.SharedFieldsSettingsParam{
						Ipv4Only: cfrex.F(true),
						Ipv6Only: cfrex.F(true),
					}),
					Tags: cfrex.F([]string{"owner:dns-team"}),
					Ttl:  cfrex.F(cfrex.SharedFieldsTtl1),
				},
				ID:      cfrex.F("023e105f4ecef8ad9ca31a8372d0c353"),
				Content: cfrex.F("198.51.100.4"),
				Type:    cfrex.F(cfrex.ZoneDNSRecordBatchParamsPutsDNSRecordsARecordTypeA),
			})}),
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

func TestZoneDNSRecordExport(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.Export(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneDNSRecordImportWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.Import(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneDNSRecordImportParams{
			File:    cfrex.F("www.example.com. 300 IN  A 127.0.0.1"),
			Proxied: cfrex.F("true"),
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

func TestZoneDNSRecordOverwriteWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.Overwrite(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneDNSRecordOverwriteParams{
			DNSRecordPost: cfrex.DNSRecordPostDNSRecordsARecordParam(cfrex.DNSRecordPostDNSRecordsARecordParam{
				SharedFieldsParam: cfrex.SharedFieldsParam{
					Comment: cfrex.F("Domain verification record"),
					Name:    cfrex.F("example.com"),
					Proxied: cfrex.F(true),
					Settings: cfrex.F(cfrex.SharedFieldsSettingsParam{
						Ipv4Only: cfrex.F(true),
						Ipv6Only: cfrex.F(true),
					}),
					Tags: cfrex.F([]string{"owner:dns-team"}),
					Ttl:  cfrex.F(cfrex.SharedFieldsTtl1),
				},
				Content: cfrex.F("198.51.100.4"),
				Type:    cfrex.F(cfrex.DNSRecordPostDNSRecordsARecordTypeA),
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

func TestZoneDNSRecordScan(t *testing.T) {
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
	_, err := client.Zones.DNSRecords.Scan(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.ZoneDNSRecordScanParams{
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
