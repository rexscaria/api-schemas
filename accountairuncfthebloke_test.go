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

func TestAccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Thebloke.ExecuteDiscolmGerman7bV1Awq(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPromptResponseFormatTypeJsonObject),
				}),
				Seed:        cfrex.F(int64(1)),
				Stream:      cfrex.F(true),
				Temperature: cfrex.F(0.000000),
				TopK:        cfrex.F(int64(1)),
				TopP:        cfrex.F(0.000000),
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
