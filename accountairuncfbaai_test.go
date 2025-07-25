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

func TestAccountAIRunCfBaaiExecuteBgeBaseEnV1_5WithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Baai.ExecuteBgeBaseEnV1_5(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfBaaiExecuteBgeBaseEnV1_5Params{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObject{
				Text:    cfrex.F[cfrex.AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectTextUnion](shared.UnionString("x")),
				Pooling: cfrex.F(cfrex.AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectPoolingMean),
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

func TestAccountAIRunCfBaaiExecuteBgeLargeEnV1_5WithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Baai.ExecuteBgeLargeEnV1_5(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfBaaiExecuteBgeLargeEnV1_5Params{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObject{
				Text:    cfrex.F[cfrex.AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectTextUnion](shared.UnionString("x")),
				Pooling: cfrex.F(cfrex.AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectPoolingMean),
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

func TestAccountAIRunCfBaaiExecuteBgeM3WithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Baai.ExecuteBgeM3(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfBaaiExecuteBgeM3Params{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputQueryAndContexts{
				Contexts: cfrex.F([]cfrex.AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputQueryAndContextsContext{{
					Text: cfrex.F("x"),
				}}),
				Query:          cfrex.F("x"),
				TruncateInputs: cfrex.F(true),
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

func TestAccountAIRunCfBaaiExecuteBgeRerankerBaseWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Baai.ExecuteBgeRerankerBase(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfBaaiExecuteBgeRerankerBaseParams{
			Contexts: cfrex.F([]cfrex.AccountAIRunCfBaaiExecuteBgeRerankerBaseParamsContext{{
				Text: cfrex.F("x"),
			}}),
			Query:        cfrex.F("x"),
			QueueRequest: cfrex.F("true"),
			TopK:         cfrex.F(int64(1)),
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

func TestAccountAIRunCfBaaiExecuteBgeSmallEnV1_5WithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Baai.ExecuteBgeSmallEnV1_5(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfBaaiExecuteBgeSmallEnV1_5Params{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObject{
				Text:    cfrex.F[cfrex.AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectTextUnion](shared.UnionString("x")),
				Pooling: cfrex.F(cfrex.AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectPoolingMean),
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
