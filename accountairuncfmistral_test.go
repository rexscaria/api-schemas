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

func TestAccountAIRunCfMistralExecuteMistral7bInstructV0_1WithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Mistral.ExecuteMistral7bInstructV0_1(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMistralExecuteMistral7bInstructV0_1Params{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(-2.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(-2.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPromptResponseFormatTypeJsonObject),
				}),
				Seed:        cfrex.F(int64(1)),
				Stream:      cfrex.F(true),
				Temperature: cfrex.F(0.000000),
				TopK:        cfrex.F(int64(1)),
				TopP:        cfrex.F(0.001000),
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

func TestAccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Mistral.ExecuteMistral7bInstructV0_2Lora(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(-2.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(-2.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPromptResponseFormatTypeJsonObject),
				}),
				Seed:        cfrex.F(int64(1)),
				Stream:      cfrex.F(true),
				Temperature: cfrex.F(0.000000),
				TopK:        cfrex.F(int64(1)),
				TopP:        cfrex.F(0.001000),
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
