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

func TestAccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.MetaLlama.ExecuteLlama2_7bChatHfLora(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPromptResponseFormatTypeJsonObject),
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
