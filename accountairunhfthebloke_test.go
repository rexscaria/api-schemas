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

func TestAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Hf.Thebloke.ExecuteDeepseekCoder6_7bBaseAwq(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Hf.Thebloke.ExecuteDeepseekCoder6_7bInstructAwq(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunHfTheblokeExecuteLlama2_13bChatAwqWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Hf.Thebloke.ExecuteLlama2_13bChatAwq(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunHfTheblokeExecuteLlamaguard7bAwqWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Hf.Thebloke.ExecuteLlamaguard7bAwq(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Hf.Thebloke.ExecuteMistral7bInstructV0_1Awq(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Hf.Thebloke.ExecuteNeuralChat7bV3_1Awq(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Hf.Thebloke.ExecuteOpenhermes2_5Mistral7bAwq(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunHfTheblokeExecuteZephyr7bBetaAwqWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Hf.Thebloke.ExecuteZephyr7bBetaAwq(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPromptResponseFormatTypeJsonObject),
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
