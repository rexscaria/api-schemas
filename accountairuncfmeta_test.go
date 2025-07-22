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

func TestAccountAIRunCfMetaExecuteLlama2_7bChatFp16WithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama2_7bChatFp16(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama2_7bChatFp16Params{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunCfMetaExecuteLlama2_7bChatInt8WithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama2_7bChatInt8(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama2_7bChatInt8Params{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunCfMetaExecuteLlama3_1_70bInstructWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama3_1_70bInstruct(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama3_1_70bInstructParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama3_1_70bInstructPreview(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunCfMetaExecuteLlama3_1_70bPreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama3_1_70bPreview(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunCfMetaExecuteLlama3_1_8bInstructWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama3_1_8bInstruct(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama3_1_8bInstructAwq(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunCfMetaExecuteLlama3_1_8bInstructFastWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama3_1_8bInstructFast(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8WithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama3_1_8bInstructFp8(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8Params{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunCfMetaExecuteLlama3_1_8bPreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama3_1_8bPreview(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama3_2_11bVisionInstruct(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Image:             cfrex.F[cfrex.AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPromptImageUnion](cfrex.AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPromptImageArray([]float64{0.000000})),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				Seed:              cfrex.F(int64(1)),
				Stream:            cfrex.F(true),
				Temperature:       cfrex.F(0.000000),
				TopK:              cfrex.F(int64(1)),
				TopP:              cfrex.F(0.000000),
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

func TestAccountAIRunCfMetaExecuteLlama3_2_1bInstructWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama3_2_1bInstruct(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama3_2_1bInstructParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunCfMetaExecuteLlama3_2_3bInstructWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama3_2_3bInstruct(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama3_2_3bInstructParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama3_3_70bInstructFp8Fast(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunCfMetaExecuteLlama3_8bInstructWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama3_8bInstruct(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama3_8bInstructParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunCfMetaExecuteLlama3_8bInstructAwqWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlama3_8bInstructAwq(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParams{
			QueueRequest: cfrex.F("true"),
			Body: cfrex.AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPrompt{
				Prompt:            cfrex.F("x"),
				FrequencyPenalty:  cfrex.F(0.000000),
				Lora:              cfrex.F("lora"),
				MaxTokens:         cfrex.F(int64(0)),
				PresencePenalty:   cfrex.F(0.000000),
				Raw:               cfrex.F(true),
				RepetitionPenalty: cfrex.F(0.000000),
				ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPromptResponseFormat{
					JsonSchema: cfrex.F[any](map[string]interface{}{}),
					Type:       cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPromptResponseFormatTypeJsonObject),
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

func TestAccountAIRunCfMetaExecuteLlamaGuard3_8bWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteLlamaGuard3_8b(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteLlamaGuard3_8bParams{
			Messages: cfrex.F([]cfrex.AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsMessage{{
				Content: cfrex.F("content"),
				Role:    cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsMessagesRoleUser),
			}}),
			QueueRequest: cfrex.F("true"),
			MaxTokens:    cfrex.F(int64(0)),
			ResponseFormat: cfrex.F(cfrex.AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsResponseFormat{
				Type: cfrex.F("type"),
			}),
			Temperature: cfrex.F(0.000000),
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

func TestAccountAIRunCfMetaExecuteM2m100_1_2bWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Meta.ExecuteM2m100_1_2b(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfMetaExecuteM2m100_1_2bParams{
			TargetLang:   cfrex.F("target_lang"),
			Text:         cfrex.F("x"),
			QueueRequest: cfrex.F("true"),
			SourceLang:   cfrex.F("source_lang"),
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
