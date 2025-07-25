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

func TestAccountAIRunCfRunwaymlExecuteStableDiffusionV1_5Img2imgWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Runwayml.ExecuteStableDiffusionV1_5Img2img(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfRunwaymlExecuteStableDiffusionV1_5Img2imgParams{
			Prompt:         cfrex.F("x"),
			QueueRequest:   cfrex.F("true"),
			Guidance:       cfrex.F(0.000000),
			Height:         cfrex.F(int64(256)),
			Image:          cfrex.F([]float64{0.000000}),
			ImageB64:       cfrex.F("image_b64"),
			Mask:           cfrex.F([]float64{0.000000}),
			NegativePrompt: cfrex.F("negative_prompt"),
			NumSteps:       cfrex.F(int64(20)),
			Seed:           cfrex.F(int64(0)),
			Strength:       cfrex.F(0.000000),
			Width:          cfrex.F(int64(256)),
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

func TestAccountAIRunCfRunwaymlExecuteStableDiffusionV1_5InpaintingWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.AI.Run.Cf.Runwayml.ExecuteStableDiffusionV1_5Inpainting(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfRunwaymlExecuteStableDiffusionV1_5InpaintingParams{
			Prompt:         cfrex.F("x"),
			QueueRequest:   cfrex.F("true"),
			Guidance:       cfrex.F(0.000000),
			Height:         cfrex.F(int64(256)),
			Image:          cfrex.F([]float64{0.000000}),
			ImageB64:       cfrex.F("image_b64"),
			Mask:           cfrex.F([]float64{0.000000}),
			NegativePrompt: cfrex.F("negative_prompt"),
			NumSteps:       cfrex.F(int64(20)),
			Seed:           cfrex.F(int64(0)),
			Strength:       cfrex.F(0.000000),
			Width:          cfrex.F(int64(256)),
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
