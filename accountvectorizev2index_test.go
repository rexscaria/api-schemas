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

func TestAccountVectorizeV2IndexNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Vectorize.V2.Indexes.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountVectorizeV2IndexNewParams{
			VectorizeCreateIndexRequest: cfrex.VectorizeCreateIndexRequestParam{
				Config: cfrex.F[cfrex.VectorizeCreateIndexRequestConfigUnionParam](cfrex.VectorizeIndexDimensionConfigurationParam{
					Dimensions: cfrex.F(int64(768)),
					Metric:     cfrex.F(cfrex.VectorizeIndexDimensionConfigurationMetricCosine),
				}),
				Name:        cfrex.F("example-index"),
				Description: cfrex.F("This is my example index."),
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

func TestAccountVectorizeV2IndexGet(t *testing.T) {
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
	_, err := client.Accounts.Vectorize.V2.Indexes.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-index",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountVectorizeV2IndexList(t *testing.T) {
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
	_, err := client.Accounts.Vectorize.V2.Indexes.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountVectorizeV2IndexDelete(t *testing.T) {
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
	_, err := client.Accounts.Vectorize.V2.Indexes.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-index",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountVectorizeV2IndexDeleteByIDsWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Vectorize.V2.Indexes.DeleteByIDs(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-index",
		cfrex.AccountVectorizeV2IndexDeleteByIDsParams{
			VectorizeIndexDeleteVectorsByIDRequest: cfrex.VectorizeIndexDeleteVectorsByIDRequestParam{
				IDs: cfrex.F([]string{"5121db81354a40c6aedc3fe1ace51c59", "f90eb49c2107486abdfd78c67e853430"}),
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

func TestAccountVectorizeV2IndexGetByIDsWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Vectorize.V2.Indexes.GetByIDs(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-index",
		cfrex.AccountVectorizeV2IndexGetByIDsParams{
			VectorizeIndexGetVectorsByIDRequest: cfrex.VectorizeIndexGetVectorsByIDRequestParam{
				IDs: cfrex.F([]string{"5121db81354a40c6aedc3fe1ace51c59", "f90eb49c2107486abdfd78c67e853430"}),
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

func TestAccountVectorizeV2IndexGetInfo(t *testing.T) {
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
	_, err := client.Accounts.Vectorize.V2.Indexes.GetInfo(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-index",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountVectorizeV2IndexInsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Vectorize.V2.Indexes.Insert(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-index",
		cfrex.AccountVectorizeV2IndexInsertParams{
			Body:               "@/path/to/vectors.ndjson",
			UnparsableBehavior: cfrex.F(cfrex.AccountVectorizeV2IndexInsertParamsUnparsableBehaviorError),
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

func TestAccountVectorizeV2IndexQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Vectorize.V2.Indexes.Query(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-index",
		cfrex.AccountVectorizeV2IndexQueryParams{
			Vector: cfrex.F([]float64{0.500000, 0.500000, 0.500000}),
			Filter: cfrex.F[any](map[string]interface{}{
				"has_viewed": map[string]interface{}{
					"$ne": true,
				},
				"streaming_platform": "netflix",
			}),
			ReturnMetadata: cfrex.F(cfrex.AccountVectorizeV2IndexQueryParamsReturnMetadataNone),
			ReturnValues:   cfrex.F(true),
			TopK:           cfrex.F(5.000000),
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

func TestAccountVectorizeV2IndexUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Vectorize.V2.Indexes.Upsert(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"example-index",
		cfrex.AccountVectorizeV2IndexUpsertParams{
			Body:               "@/path/to/vectors.ndjson",
			UnparsableBehavior: cfrex.F(cfrex.AccountVectorizeV2IndexUpsertParamsUnparsableBehaviorError),
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
