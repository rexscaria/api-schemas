// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/rexscaria/api-schemas"
	"github.com/rexscaria/api-schemas/internal/testutil"
	"github.com/rexscaria/api-schemas/option"
)

func TestAccountAIRunCfFacebookExecuteBartLargeCnnWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
		option.WithAccessToken("My Access Token"),
	)
	_, err := client.Accounts.AI.Run.Cf.Facebook.ExecuteBartLargeCnn(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfFacebookExecuteBartLargeCnnParams{
			InputText:    cfrex.F("x"),
			QueueRequest: cfrex.F("true"),
			MaxLength:    cfrex.F(int64(0)),
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

func TestAccountAIRunCfFacebookExecuteDetrResnet50WithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
		option.WithAccessToken("My Access Token"),
	)
	_, err := client.Accounts.AI.Run.Cf.Facebook.ExecuteDetrResnet50(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		io.Reader(bytes.NewBuffer([]byte("Example data"))),
		cfrex.AccountAIRunCfFacebookExecuteDetrResnet50Params{
			QueueRequest: cfrex.F("true"),
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
