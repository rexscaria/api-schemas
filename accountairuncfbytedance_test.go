// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rexscaria/api-schemas"
	"github.com/rexscaria/api-schemas/option"
)

func TestAccountAIRunCfBytedanceExecuteStableDiffusionXlLightningWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := cfrex.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("My API Email"),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.Accounts.AI.Run.Cf.Bytedance.ExecuteStableDiffusionXlLightning(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountAIRunCfBytedanceExecuteStableDiffusionXlLightningParams{
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
